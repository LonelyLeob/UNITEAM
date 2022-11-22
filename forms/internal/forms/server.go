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

	s.store = NewStore(db)
}

// set up routes to get requests
func (s *Server) setupRoutes() {
	s.handler.Use(
		s.SkipOptionsReq,
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Origin", "Content-Type", "Authorization"}),
		))

	api1 := s.handler.PathPrefix("/api/v1").Subrouter()

	api1.HandleFunc("/create", s.FormsCreatingHttp()).Methods("POST", "OPTIONS")
	api1.HandleFunc("/create/field", s.FieldCreatingForm()).Methods("POST", "OPTIONS")
	api1.HandleFunc("/create/answer", s.AnswerCreatingField()).Methods("POST", "OPTIONS")

	api1.HandleFunc("/get/forms", s.GetFormsByAuthorUUID()).Methods("GET", "OPTIONS")

	api1.HandleFunc("/delete", s.DeleteFormByFormUuid()).Methods("DELETE", "OPTIONS")
	api1.HandleFunc("/delete/field", s.DeleteFieldById()).Methods("DELETE", "OPTIONS")
	api1.HandleFunc("/delete/answer", s.DeleteAnswerById()).Methods("DELETE", "OPTIONS")
}

func (s *Server) SkipOptionsReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
	})
}
