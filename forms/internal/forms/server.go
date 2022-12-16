package forms

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	createForm   = "/create"
	createField  = "/create/field"
	createAnswer = "/create/answer"
	getForms     = "/get/forms"
	deleteForm   = "/delete"
	deleteField  = "/delete/field"
	deleteAnswer = "/delete/answer"
)

type Server struct {
	signingKey []byte
	handler    *mux.Router
	store      *Store
}

func Spawn(key []byte) *Server {
	return &Server{
		signingKey: key,
		handler:    mux.NewRouter(),
	}
}

func (s *Server) Start(web, url string) {
	s.setupRoutes()
	s.configureDB(url)

	log.Fatal(http.ListenAndServe(":8080", s.handler))
}

func (s *Server) configureDB(url string) {

	//open connection via database
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}

	//ping to test database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	s.store = NewStore(db)
}

// set up routes to get requests
func (s *Server) setupRoutes() {
	api1 := s.handler.PathPrefix("/api/v1").Subrouter()

	api1.Use(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Origin", "Authorization"}),
			handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodOptions, http.MethodDelete}),
		))

	api1.HandleFunc(createForm, s.FormsCreatingHttp()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(createField, s.FieldCreatingForm()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(createAnswer, s.AnswerCreatingField()).Methods(http.MethodPost, http.MethodOptions)

	api1.HandleFunc(getForms, s.GetFormsByAuthorUUID()).Methods(http.MethodGet, http.MethodOptions)

	api1.HandleFunc(deleteForm, s.DeleteFormByFormUuid()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc(deleteField, s.DeleteFieldById()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc(deleteAnswer, s.DeleteAnswerById()).Methods(http.MethodDelete, http.MethodOptions)
}
