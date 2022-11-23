package internal

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	attrs  = "/attrs"
	token  = "/authorize"
	reg    = "/registration"
	update = "/update"
)

type Handler interface {
	Register(store Storage, url, web, csrf string)
}

type handler struct {
	router *mux.Router
	skey   []byte
	store  Storage
}

func NewHandler(key string) Handler {
	return &handler{
		router: mux.NewRouter(),
		skey:   []byte(key),
	}
}

func (h *handler) Register(store Storage, url, web, csrfkey string) {
	h.store = store
	h.store.Connect(url)

	api1 := h.router.PathPrefix("/api/v1").Subrouter()

	api1.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Origin", "Authorization"}),
		handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS"}),
	))

	api1.HandleFunc(token, h.Authorize()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(reg, h.Registration()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(attrs, h.GetTokenAttrs()).Methods(http.MethodGet)
	api1.HandleFunc(update, h.UpdateCredentials()).Methods(http.MethodGet, http.MethodOptions)

	log.Println("app has been started successfully")

	log.Fatal(http.ListenAndServe(web, h.router))
}

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

		user := &User{
			name: req.Name,
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

		user := &User{
			name:     req.Name,
			password: req.Password,
			email:    req.Email,
			role:     "user",
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
			return
		}

		refresh, err := createRefreshToken(user, h.skey)
		if err != nil {
			return
		}

		sendCredentials(w, access, refresh)
	}
}

func (h *handler) GetTokenAttrs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenstr := r.URL.Query().Get("token")
		if tokenstr == "" {
			errJSON(w, http.StatusUnauthorized, http.ErrAbortHandler)
			return
		}

		name, err := ParseAccessToken(tokenstr, h.skey)
		if err != nil {
			errJSON(w, http.StatusConflict, err)
			return
		}

		toJSON(w, http.StatusOK, map[string]string{"user": name})
	}
}

func (h *handler) UpdateCredentials() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		refresh := r.URL.Query().Get("ref")
		if refresh == "" {
			errJSON(w, http.StatusBadRequest, errEmptyHeader)
			return
		}

		email, err := ParseRefreshToken(refresh, h.skey)
		if err != nil {
			errJSON(w, http.StatusUnprocessableEntity, err)
			return
		}

		u := &User{
			email: email,
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
