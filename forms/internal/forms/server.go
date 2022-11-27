package forms

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/L0nelyleob/UNITEAM/golang-forms/internal/forms/redisclient"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

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
	redis      *redisclient.RedisRepo
}

func Spawn(key []byte) *Server {
	return &Server{
		signingKey: key,
		handler:    mux.NewRouter(),
		redis:      redisclient.SetRedis(),
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

	s.setMigrations(db)

	s.store = NewStore(db)
}

func (s *Server) setMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"service", driver)
	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}

// set up routes to get requests
func (s *Server) setupRoutes() {
	api1 := s.handler.PathPrefix("/api/v1").Subrouter()

	api1.Use(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedHeaders([]string{"Origin", "Authorization"}),
			handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodOptions, http.MethodDelete}),
		), s.Authorize_Middleware)

	api1.HandleFunc(createForm, s.FormsCreatingHttp()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(createField, s.FieldCreatingForm()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(createAnswer, s.AnswerCreatingField()).Methods(http.MethodPost, http.MethodOptions)

	api1.HandleFunc(getForms, s.GetFormsByAuthorUUID()).Methods(http.MethodGet, http.MethodOptions)

	api1.HandleFunc(deleteForm, s.DeleteFormByFormUuid()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc(deleteField, s.DeleteFieldById()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc(deleteAnswer, s.DeleteAnswerById()).Methods(http.MethodDelete, http.MethodOptions)
}
