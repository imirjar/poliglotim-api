package service

import (
	"context"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func New() *Service {
	return &Service{}
}

type Service struct {
	Storage Storage
}

type Storage interface {
	GetLesson(context.Context, string) (models.Lesson, error)
	GetCourses(context.Context) ([]models.Course, error)
	GetChaptersFromCourse(context.Context, string) ([]models.Chapter, error)
	GetChapterLessons(context.Context, string) ([]models.Lesson, error)
}

func (s *Service) GetCourses(ctx context.Context) ([]models.Course, error) {
	return s.Storage.GetCourses(ctx)
}

func (s *Service) GetChapterLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	return s.Storage.GetChapterLessons(ctx, chapterID)
}

func (s *Service) GetLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	lesson, err := s.Storage.GetLesson(ctx, lessonID)
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}

func (s *Service) GetCourseChapters(ctx context.Context, courseID string) ([]models.Chapter, error) {
	lesson, err := s.Storage.GetChaptersFromCourse(ctx, courseID)
	if err != nil {
		return lesson, err
	}
	return lesson, nil
}
