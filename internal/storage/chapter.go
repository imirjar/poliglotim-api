package storage

import (
	"context"
	"time"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func (s *Storage) GetChaptersFromCourse(ctx context.Context, courseId string) ([]models.Chapter, error) {

	// Главы идут по порядку, поэтому им нужна "позиция"
	// отображать ее в API не нужно
	query := `
		SELECT 
			id, 
			name, 
			description, 
			updated
		FROM 
			chapters
		WHERE 
			course_id = $1
		ORDER BY 
			position
	`

	rows, err := s.psql.Query(ctx, query, courseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []models.Chapter
	for rows.Next() {
		var chapter models.Chapter
		var updated time.Time

		err := rows.Scan(
			&chapter.Id,
			&chapter.Name,
			&chapter.Description,
			&updated,
		)
		if err != nil {
			return nil, err
		}

		chapter.Updated = updated
		chapters = append(chapters, chapter)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return chapters, nil
}
