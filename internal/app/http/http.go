package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/imirjar/poliglotim-api/docs" // важно: импорт сгенерированной документации
	"github.com/imirjar/poliglotim-api/internal/app/http/middleware"
	"github.com/imirjar/poliglotim-api/internal/models"
	httpSwagger "github.com/swaggo/http-swagger"
)

// HttpServer представляет HTTP сервер приложения
// @title Poliglotim API Gateway
// @version 1.0
// @description API Gateway для образовательной платформы Poliglotim
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@poliglotim.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
type HttpServer struct {
	Port    string
	Service Service
}

type Service interface {
	GetAllCourses(context.Context) ([]models.Course, error)
	GetFullCourse(context.Context, string) (models.Course, error)
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

	// Swagger документация
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	router.Handle("/courses", srv.getCourses()).Methods("GET")
	router.Handle("/course/{course_id}", srv.getCourse()).Methods("GET")
	router.Handle("/lesson/{lesson_id}", srv.getLesson()).Methods("GET")

	return http.ListenAndServe(fmt.Sprintf(":%s", srv.Port), router)
}

// getCourses возвращает список всех курсов
// @Summary Получить все курсы
// @Description Возвращает список всех доступных курсов
// @Tags courses
// @Accept json
// @Produce json
// @Success 200 {array} models.Course "Список курсов"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /courses [get]
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

// getCourse возвращает полную информацию о курсе
// @Summary Получить курс по ID
// @Description Возвращает полную информацию о курсе включая главы и уроки
// @Tags courses
// @Accept json
// @Produce json
// @Param course_id path string true "ID курса"
// @Success 200 {object} models.Course "Информация о курсе"
// @Failure 400 {object} ErrorResponse "Неверный ID курса"
// @Failure 404 {object} ErrorResponse "Курс не найден"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /course/{course_id} [get]
func (srv *HttpServer) getCourse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		courseID := vars["course_id"]

		course, err := srv.Service.GetFullCourse(r.Context(), courseID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(course); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// getLesson возвращает информацию об уроке
// @Summary Получить урок по ID
// @Description Возвращает информацию об уроке по временной ссылке
// @Tags lessons
// @Accept json
// @Produce json
// @Param lesson_id path string true "ID урока"
// @Success 200 {object} models.Lesson "Информация об уроке"
// @Failure 400 {object} ErrorResponse "Неверный ID урока"
// @Failure 403 {object} ErrorResponse "Доступ запрещен"
// @Failure 404 {object} ErrorResponse "Урок не найден"
// @Failure 500 {object} ErrorResponse "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /lesson/{lesson_id} [get]
func (srv *HttpServer) getLesson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		lessonID := vars["lesson_id"]

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

// ErrorResponse представляет стандартный ответ об ошибке
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}
