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
	// getCourses() ([]models.Course, error)
	// getChapters() ([]models.Chapter, error)
	// getLessons() ([]models.Lesson, error)
	// getLesson() (models.Lesson, error)
}

func (s *Service) GetCourses() ([]models.Course, error) {
	return []models.Course{}, nil
}

func (s *Service) GetProgress() (map[models.Course]map[models.Chapter][]models.Lesson, error) {
	return nil, nil
}

func (s *Service) GetLesson() (models.Lesson, error) {
	lesson, err := s.Storage.GetLesson("1")
	if err != nil {
		return models.Lesson{}, err
	}
	return lesson, nil
}
