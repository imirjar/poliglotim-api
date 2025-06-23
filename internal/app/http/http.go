package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imirjar/poliglotim-api/internal/models"
)

type HttpServer struct {
	Port    string
	Service Service
}

type Service interface {
	GetCourses() ([]models.Course, error)
	// GetProgress() (map[models.Course]map[models.Chapter][]models.Lesson, error)
	GetLesson() (models.Lesson, error)
}

func New(port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

func (srv *HttpServer) Run() error {

	router := mux.NewRouter()

	router.Handle("/courses", srv.getCourses())
	router.Handle("/lesson", srv.getLesson())
	return http.ListenAndServe(srv.Port, router)
}

func (srv *HttpServer) getCourses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		courses, err := srv.Service.GetCourses()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(courses); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (srv *HttpServer) getLesson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lesson, err := srv.Service.GetLesson()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(lesson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
