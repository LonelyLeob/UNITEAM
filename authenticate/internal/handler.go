package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/L0nelyleob/UNITEAM/oauth-service/internal/cache"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mileusna/useragent"
)

const (
	signIn     = "/authorize"
	signUp     = "/registration"
	signUpdate = "/update"
	signOut    = "/out"
)

type Handler interface {
	Register(store Storage, url, web string)
}

type handler struct {
	router *mux.Router
	skey   []byte
	store  Storage
	redis  *cache.RedisRepo
}

func NewHandler(key string) Handler {
	return &handler{
		router: mux.NewRouter(),
		skey:   []byte(key),
	}
}

func (h *handler) Register(store Storage, url, web string) {
	h.store = store
	h.store.Connect(url)

	client, err := cache.SetRedis()
	if err != nil {
		log.Fatal(err)
	}

	h.redis = client

	api1 := h.router.PathPrefix("/api/v1").Subrouter()

	api1.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Origin", "Authorization"}),
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}),
	))

	api1.HandleFunc(signIn, h.Authorize()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(signUp, h.Registration()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(signUpdate, h.UpdateCredentials()).Methods(http.MethodGet, http.MethodOptions)
	api1.HandleFunc(signOut, h.signOut()).Methods(http.MethodGet, http.MethodOptions)

	log.Println("app has been started successfully")

	log.Fatal(http.ListenAndServe(web, h.router))
}

// TODO: do redis checking for users tokens
func (h *handler) Authorize() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		device := useragent.Parse(r.UserAgent())
		user := &User{
			name: req.Name,
			meta: &UserMeta{
				device: device.Device,
			},
		}

		user, err := h.store.Web().GetUser(user, req.Password)
		if err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		access, err := createAccessToken(user, h.skey)
		if err != nil {
			return
		}
		if err := h.redis.SetToken(access); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		refresh, err := createRefreshToken(user, h.skey)
		if err != nil {
			return
		}

		sendCredentials(w, access, refresh)
	}
}

func (h *handler) Registration() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		device := useragent.Parse(r.UserAgent())
		user := &User{
			name:     req.Name,
			password: req.Password,
			email:    req.Email,
			role:     "user",
			meta: &UserMeta{
				device:   device.Device,
				lastSign: time.Now().Unix(),
			},
		}

		if err := user.CheckForRequiredParams(); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		if err := h.store.Web().CreateUser(user); err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		access, err := createAccessToken(user, h.skey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}
		if err := h.redis.SetToken(access); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		refresh, err := createRefreshToken(user, h.skey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		sendCredentials(w, access, refresh)
	}
}

func (h *handler) signOut() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerparts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(headerparts) > 2 || headerparts[0] != "Bearer" {
			errJSON(w, http.StatusUnauthorized, errEmptyHeader)
			return
		}
		if err := h.redis.UnsetToken(headerparts[1]); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"status": "logout"})
	}
}

func (h *handler) UpdateCredentials() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh := r.URL.Query().Get("ref")
		if refresh == "" {
			errJSON(w, http.StatusBadRequest, errEmptyHeader)
			return
		}

		claims, err := ParseRefreshToken(refresh, h.skey)
		if err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		device := useragent.Parse(r.UserAgent())

		u := &User{
			email: claims.Email,
			meta: &UserMeta{
				device:   device.Device,
				lastSign: time.Now().Unix(),
			},
		}

		if err := h.store.Web().SearchUserByEmail(u); err != nil {
			errJSON(w, http.StatusBadRequest, errInfo)
			return
		}

		updacc, err := createAccessToken(u, h.skey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}
		if err := h.redis.SetToken(updacc); err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}

		updref, err := createRefreshToken(u, h.skey)
		if err != nil {
			errJSON(w, http.StatusBadRequest, err)
			return
		}
		toJSON(w, http.StatusOK, map[string]string{
			"access":  updacc,
			"refresh": updref,
		})
	}
}
