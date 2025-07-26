package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	Id          string    `json:"id" `
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Updated     time.Time `json:"updated"`
	Logo        File      `json:"logo,omitempty"`
}

type Chapter struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Updated     time.Time `json:"updated"`
	Course      string    `json:"course_id,omitempty"`
	Lessons     []Lesson  `json:"lessons,omitempty"`
}

type Lesson struct {
	Id      string    `json:"id" bson:"_id"`
	Chapter string    `json:"chapter_id,omitempty"`
	Title   string    `json:"title"`
	Text    string    `json:"text,omitempty"`
	Updated time.Time `json:"updated"`
}

type File struct {
	Id      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Link    string             `json:"link"`
	Updated time.Time          `json:"updated"`
}
