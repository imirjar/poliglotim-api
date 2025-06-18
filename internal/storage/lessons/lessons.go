package lessons

import (
	"log"

	"github.com/imirjar/poliglotim-api/internal/models"
)

type LessonsStore struct {
	*Mongo
}

func New(dbConn string) *LessonsStore {
	log.Print("LessonsStore")
	return &LessonsStore{
		Mongo: NewDB(dbConn),
	}
}

func (ls *LessonsStore) GetLessons() ([]models.Lesson, error) {
	return []models.Lesson{}, nil
}

func (ls *LessonsStore) GetLesson(id string) (models.Lesson, error) {
	return models.Lesson{
		Title: "Урок1",
		Text:  "Учиться, учиться и еще раз учиться!",
	}, nil
}
