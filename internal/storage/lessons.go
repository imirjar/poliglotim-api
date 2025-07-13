package storage

import (
	"context"
	"log"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func (s *Storage) GetChapterLessons(ctx context.Context, chapterID string) ([]models.Lesson, error) {

	query := `
		SELECT 
			l.id, 
			l.title, 
			l.text, 
			l.updated
		FROM 
			Lessons l
		WHERE 
			chapter_id = $1
		ORDER BY 
			l.title
	`

	rows, err := s.psql.Query(ctx, query, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson

		err := rows.Scan(
			&lesson.Id,
			&lesson.Title,
			&lesson.Text,
			&lesson.Updated,
		)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	log.Print()
	return lessons, nil

}
func (s *Storage) GetLesson(ctx context.Context, id string) (models.Lesson, error) {
	query := `
		SELECT 
			c.id, 
			c.name, 
			c.description, 
			c.updated
		FROM 
			courses c
		WHERE
			id = $1
	`
	row := s.psql.QueryRow(ctx, query, id)

	var lesson models.Lesson
	err := row.Scan(
		&lesson.Id,
		&lesson.Title,
		&lesson.Text,
		&lesson.Updated,
	)
	if err != nil {
		return lesson, err
	}

	return lesson, nil
}
