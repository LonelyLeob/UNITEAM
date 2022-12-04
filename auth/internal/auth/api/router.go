package api

import (
	"authenticate/internal/auth/store"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	UserReg_Route        = "/registration"
	UserAuth_Route       = "/authorize"
	UserUpdateCred_Route = "/update"
	ForgetPwd_Route      = "/forget/pwd"
	UpdatePwd_Route      = "/update/pwd"
	DeleteUser_Route     = "/delete"
	GetUser_Route        = "/user"
	LogoutUser_Route     = "/logout"

	debug    = os.Getenv("DEBUG")
	account  = os.Getenv("MAIL_ACCOUNT")
	password = os.Getenv("MAIL_PASSWORD")
	port     = os.Getenv("REMOTE_PORT")
)

type Server struct {
	tg         *TokenGiver
	router     *mux.Router
	pgstore    *store.PostgresStore
	rstore     *store.RedisStore
	mailClient *MailClient
}

func NewServer(pgaddr, raddr, pwd, key string) *Server {
	return &Server{
		NewGiver(15*time.Minute, 72*time.Hour, key),
		mux.NewRouter(),
		store.NewStore(pgaddr),
		store.NewRedis(raddr, pwd),
		NewMailClient(account, password),
	}
}

func (s *Server) StartUp() {
	if err := s.pgstore.InitConnect(context.Background()); err != nil {
		if debug == "True" {
			logrus.Error(err)
		}

		os.Exit(1)
	}

	api1 := s.router.PathPrefix("/api/v1").Subrouter()
	api1.Use(handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Origin", "Authorization", "X-Refresh"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodOptions, http.MethodDelete}),
	))

	go api1.HandleFunc(UserReg_Route, s.RegistrationUser_Handler()).Methods(http.MethodPost, http.MethodOptions)
	go api1.HandleFunc(UserAuth_Route, s.AuthenticateUser_Handler()).Methods(http.MethodPost, http.MethodOptions)
	go api1.HandleFunc(UserUpdateCred_Route, s.UpdateTokenUser_Handler()).Methods(http.MethodGet, http.MethodOptions)
	go api1.HandleFunc(ForgetPwd_Route, s.ForgetPassword_Handler()).Methods(http.MethodPost, http.MethodOptions)
	go api1.HandleFunc(DeleteUser_Route, s.DeleteUser_Handler()).Methods(http.MethodDelete, http.MethodOptions)
	go api1.HandleFunc(GetUser_Route, s.GetUser_Handler()).Methods(http.MethodGet, http.MethodOptions)
	go api1.HandleFunc(UpdatePwd_Route, s.UpdatePassword_Handler()).Methods(http.MethodGet, http.MethodOptions)
	go api1.HandleFunc(LogoutUser_Route, s.LogoutUser_Handler()).Methods(http.MethodGet, http.MethodOptions)

	http.ListenAndServe(":"+port, s.router)
}

func BindJSON(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			if os.Getenv("DEBUG") == "True" {
				logrus.Error(err)
			}
		}
	}
}

func errJSON(w http.ResponseWriter, err error, code int) {
	BindJSON(w, &errorResponse{
		Message: err.Error(),
	}, code)
}
