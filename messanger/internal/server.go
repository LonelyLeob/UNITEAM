package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	cfgRemote *ConfigRemote
	router    *mux.Router
	mongoose  *Mongoose
}

func Initialize(cr *ConfigRemote, cm *ConfigMongo) *server {
	return &server{
		cfgRemote: cr,
		router:    mux.NewRouter(),
		mongoose:  InitMongoose(cm),
	}
}

func (s *server) Serve() {
	fmt.Printf("try connect mongo at: %s \n", s.mongoose.config.CreateURI())
	if err := s.mongoose.Connect(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("database connected")

	fmt.Printf("app listening on port %s", s.cfgRemote.CreateDomainAddr())

	s.router.HandleFunc("/send", s.SendMessageHandler()).Methods("POST")
	s.router.HandleFunc("/messages", s.GetBatchMessages()).Methods("GET")

	log.Fatal(http.ListenAndServe(s.cfgRemote.CreateDomainAddr(), s.router))
}
