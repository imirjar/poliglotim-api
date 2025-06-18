package storage

import (
	"github.com/imirjar/poliglotim-api/internal/models"
	"github.com/imirjar/poliglotim-api/internal/storage/course"
	"github.com/imirjar/poliglotim-api/internal/storage/lessons"
)

type CourseStorage interface{}

type LessonStorage interface {
	GetLesson(id string) (models.Lesson, error)
}

type Storage struct {
	CourseStorage
	LessonStorage
}

func New(mongoConf, psqlConf string) *Storage {
	return &Storage{
		CourseStorage: course.New(),
		LessonStorage: lessons.New(mongoConf),
	}
}
