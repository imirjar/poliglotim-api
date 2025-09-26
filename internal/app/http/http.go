package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imirjar/poliglotim-api/internal/app/http/middleware"
	"github.com/imirjar/poliglotim-api/internal/models"
)

type HttpServer struct {
	Port    string
	Service Service
}

type Service interface {
	GetAllCourses(context.Context) ([]models.Course, error)
	// GetCourse(context.Context, string) (models.Course, error)
	GetFullCourse(context.Context, string) (models.Course, error)
	// GetCourseChapters(context.Context, string) ([]models.Chapter, error)
	GetLesson(context.Context, string) (models.Lesson, error)
}

func New(port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}

func (srv *HttpServer) Run() error {
	router := mux.NewRouter()
	mdlwr := middleware.New()

	// Применяем CORS middleware ко всем маршрутам
	router.Use(mdlwr.CORS(), middleware.Auth())

	router.Handle("/courses", srv.getCourses())
	router.Handle("/course/{course_id}", srv.getCourse())
	// router.Handle("/lessons/{chapter_id}", srv.getChapterLessons())
	router.Handle("/lesson/{lesson_id}", srv.getLesson())
	return http.ListenAndServe(fmt.Sprintf(":%s", srv.Port), router)
}

func (srv *HttpServer) getCourses() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		courses, err := srv.Service.GetAllCourses(r.Context())
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

func (srv *HttpServer) getCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		courseID := vars["course_id"]
		// log.Print(lessonID)

		lesson, err := srv.Service.GetFullCourse(r.Context(), courseID)
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

// необходимо получать урок по временной ссылке
// так как не факт что у пользователя есть доступ к этому уроку
// так как его прогресс по курсу не позволяет или у него нет доступа к курсу
// или курс этого урока не опубликован
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
