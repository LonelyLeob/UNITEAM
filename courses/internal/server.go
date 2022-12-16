package internal

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	getShortCourses = "/get/courses/short"
	getCourse       = "/get/course"
	updateSection   = "/update/section"
	addCourse       = "/add/course"
	deleteCourse    = "/delete/course"
	addSection      = "/add/section"
)

type server struct {
	handler *mux.Router
	config  *ConfigRemote
	store   *Storage
}

func NewServer(cd *ConfigDatabase, cr *ConfigRemote) *server {
	return &server{
		handler: mux.NewRouter(),
		config:  cr,
		store:   NewStorage(cd),
	}
}

func (s *server) Start() {
	if err := s.store.Connect(); err != nil {
		log.Fatal(err)
	}

	api1 := s.handler.PathPrefix("/api/v1").Subrouter()
	api1.Use(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedHeaders([]string{"Origin", "Authorization"}),
			handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodOptions, http.MethodDelete, http.MethodPatch}),
		))

	api1.HandleFunc(getShortCourses, s.getShorthandCourses_Handler()).Methods(http.MethodGet, http.MethodOptions)
	api1.HandleFunc(getCourse, s.getCourseData_Handler()).Methods(http.MethodGet, http.MethodOptions)
	api1.HandleFunc(addCourse, s.addCourse_Handler()).Methods(http.MethodPost, http.MethodOptions)
	api1.HandleFunc(updateSection, s.updateSectionContent_Handler()).Methods(http.MethodPatch, http.MethodOptions)
	api1.HandleFunc(deleteCourse, s.deleteCourse_Handler()).Methods(http.MethodDelete, http.MethodOptions)
	api1.HandleFunc(addSection, s.addSectionCourse_Handler()).Methods(http.MethodPost, http.MethodOptions)

	http.ListenAndServe(s.config.CreateDomainAddr(), s.handler)
}
