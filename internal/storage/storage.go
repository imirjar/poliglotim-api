package storage

import (
	"github.com/imirjar/poliglotim-api/internal/models"
	"github.com/imirjar/poliglotim-api/internal/storage/course"
	"github.com/imirjar/poliglotim-api/internal/storage/lessons"
)

type CourseStorage interface {
	GetCourses() ([]models.Course, error)
	GetChaptersFromCourse(courseId string) ([]models.Chapter, error)
	Disconnect() error
}

type LessonStorage interface {
	GetLesson(id string) (models.Lesson, error)
	Disconnect() error
}

type Storage struct {
	CourseStorage
	LessonStorage
}

func New(mongoConf, psqlConf string) *Storage {
	return &Storage{
		CourseStorage: course.New(psqlConf),
		LessonStorage: lessons.New(mongoConf),
	}
}

func (s *Storage) Disconnect() {
	s.CourseStorage.Disconnect()
	s.LessonStorage.Disconnect()
}
