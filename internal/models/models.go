package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Updated     time.Time          `json:"updated"`
	Logo        File               `json:"logo"`
}

type Chapter struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name"`
	Course      Course             `json:"course"`
	Description string             `json:"description"`
	Updated     time.Time          `json:"updated"`
}

type File struct {
	Id      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Link    string             `json:"link"`
	Updated time.Time          `json:"updated"`
}

type Lesson struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Title   string             `json:"title"`
	Text    string             `json:"text"`
	Updated time.Time          `json:"updated"`
}
