package service

import (
	"github.com/imirjar/poliglotim-api/internal/models"
)

func New() *Service {
	return &Service{}
}

type Service struct {
	Storage Storage
}

type Storage interface {
	GetLesson(id string) (models.Lesson, error)
	GetCourses() ([]models.Course, error)
	// GetChaptersFromCourse(courseId string) ([]models.Chapter, error)
}

func (s *Service) GetCourses() ([]models.Course, error) {
	return []models.Course{}, nil
}

func (s *Service) GetLesson() (models.Lesson, error) {
	lesson, err := s.Storage.GetLesson("1")
	if err != nil {
		return models.Lesson{}, err
	}
	return lesson, nil
}
