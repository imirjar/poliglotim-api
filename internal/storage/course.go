package storage

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/imirjar/poliglotim-api/internal/models"
)

func (s *Storage) GetCourses(ctx context.Context) ([]models.Course, error) {
	query := `
		SELECT 
			c.id, 
			c.name, 
			c.description, 
			c.updated, 
			c.logo_path 
		FROM 
			courses c
		ORDER BY 
			c.name
	`

	rows, err := s.psql.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		var logoPath sql.NullString
		var updated time.Time

		err := rows.Scan(
			&course.Id,
			&course.Name,
			&course.Description,
			&updated,
			&logoPath,
		)
		if err != nil {
			return nil, err
		}

		course.Updated = updated
		// if logoPath.Valid {
		// 	course.Logo = models.File{Path: logoPath.String}
		// }

		// // Получаем главы для курса
		// chapters, err := p.GetChaptersFromCourse(course.Id)
		// if err != nil {
		// 	return nil, err
		// }
		// course.Chapters = chapters

		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	log.Print()
	return courses, nil
}
