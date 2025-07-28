package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/imirjar/poliglotim-api/internal/models"
)

type HttpServer struct {
	Port    string
	Service Service
}

type Service interface {
	GetCourses(context.Context) ([]models.Course, error)
	GetChapterLessons(context.Context, string) ([]models.Lesson, error)
	GetCourseChapters(context.Context, string) ([]models.Chapter, error)
	GetLesson(context.Context, string) (models.Lesson, error)
}

func New(port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

func (srv *HttpServer) Run() error {
	router := mux.NewRouter()

	// Настраиваем CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Разрешить все источники
		// handlers.AllowedOrigins([]string{"http://localhost:59889"}), // Или конкретный источник
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}),
		handlers.ExposedHeaders([]string{"Link"}),
		// handlers.AllowCredentials(),
		handlers.MaxAge(300),
	)

	// Применяем CORS middleware ко всем маршрутам
	router.Use(corsMiddleware)

	router.Handle("/courses", srv.getCourses())
	router.Handle("/chapters/{course_id}", srv.getCourseChapters())
	router.Handle("/lessons/{chapter_id}", srv.getChapterLessons())
	router.Handle("/lesson/{lesson_id}", srv.getLesson())
	return http.ListenAndServe(fmt.Sprintf(":%s", srv.Port), router)
}

func (srv *HttpServer) getCourses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		courses, err := srv.Service.GetCourses(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(courses); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (srv *HttpServer) getCourseChapters() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		courseID := vars["course_id"]
		// log.Print(lessonID)

		lesson, err := srv.Service.GetCourseChapters(r.Context(), courseID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(lesson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (srv *HttpServer) getChapterLessons() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		chapterID := vars["chapter_id"]
		// log.Print(lessonID)

		lesson, err := srv.Service.GetChapterLessons(r.Context(), chapterID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(lesson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (srv *HttpServer) getLesson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		lessonID := vars["lesson_id"]
		// log.Print(lessonID)

		lesson, err := srv.Service.GetLesson(r.Context(), lessonID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(lesson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
