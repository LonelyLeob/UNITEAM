package forms

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
		))

	api1.HandleFunc("/create", s.FormsCreatingHttp()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc("/create/field", s.FieldCreatingForm()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc("/create/answer", s.AnswerCreatingField()).Methods(http.MethodPost, http.MethodOptions)

	api1.HandleFunc("/get/forms", s.GetFormsByAuthorUUID()).Methods(http.MethodGet, http.MethodOptions)

	api1.HandleFunc("/delete", s.DeleteFormByFormUuid()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc("/delete/field", s.DeleteFieldById()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc("/delete/answer", s.DeleteAnswerById()).Methods(http.MethodDelete, http.MethodOptions)
}
