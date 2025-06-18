package course

import "log"

type CourseStore struct {
	*Psql
}

func New() *CourseStore {
	log.Print("CourseStore")
	return &CourseStore{}
}
