package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/imirjar/poliglotim-api/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *Storage) GetChapterLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {
	collection := s.mongo.Database("PoliglotimCourses").Collection("lessons")

	// Создаем фильтр для поиска уроков по chapter_id
	filter := bson.M{"chapter_id": chapterID}

	// Правильный вариант сортировки
	opts := options.Find().SetSort(bson.M{"position": 1}) // Используем bson.M вместо bson.D

	// Выполняем запрос
	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Printf("ERROR finding lessons: %v", err)
		return nil, fmt.Errorf("failed to find lessons: %v", err)
	}
	defer cursor.Close(ctx)

	// Декодируем результаты
	var lessons []models.Lesson
	if err = cursor.All(ctx, &lessons); err != nil {
		log.Printf("ERROR decoding lessons: %v", err)
		return nil, fmt.Errorf("failed to decode lessons: %v", err)
	}

	return lessons, nil
}
func (s *Storage) GetLesson(ctx context.Context, id string) (models.Lesson, error) {

	collection := s.mongo.Database("PoliglotimCourses").Collection("lessons")

	var lesson models.Lesson

	// Преобразуем строковый ID в ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return lesson, fmt.Errorf("invalid report ID format: %v", err)
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&lesson)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return lesson, fmt.Errorf("report not found")
		}
		log.Printf("ERROR getting report: %v", err)
		return lesson, fmt.Errorf("failed to get report: %v", err)
	}

	return lesson, nil
}
