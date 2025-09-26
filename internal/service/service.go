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
	GetCourses(context.Context) ([]models.Course, error)
	GetCourseWithContent(context.Context, string) (models.Course, error)
	GetLesson(context.Context, string) (models.Lesson, error)
}

// Fetch all of 'published' courses.
func (s *Service) GetAllCourses(ctx context.Context) ([]models.Course, error) {
	return s.Storage.GetCourses(ctx)
}

// Get course plan with chapters and lessons name without lessons text.
func (s *Service) GetFullCourse(ctx context.Context, courseID string) (models.Course, error) {
	// Check user permisson for this course
	course, err := s.Storage.GetCourseWithContent(ctx, courseID)
	if err != nil {
		return course, err
	}

	return course, nil
}

func (s *Service) GetLesson(ctx context.Context, lessonID string) (models.Lesson, error) {
	// Check user progress for this course
	return s.Storage.GetLesson(ctx, lessonID)
}
