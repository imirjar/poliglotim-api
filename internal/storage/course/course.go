package course

import (
	"log"

	"github.com/imirjar/poliglotim-api/internal/models"
)

type CourseStore struct {
	*Psql
}

func New(dbConn string) *CourseStore {
	log.Print("CourseStore")
	return &CourseStore{
		Psql: NewDB(dbConn),
	}
}

func (c *CourseStore) GetCourses() ([]models.Course, error) {
	return []models.Course{}, nil
}
func (c *CourseStore) GetChaptersFromCourse(courseId string) ([]models.Chapter, error) {
	return []models.Chapter{}, nil
}
