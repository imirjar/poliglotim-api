package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/imirjar/poliglotim-api/internal/domain/models"
)

func (s *Storage) GetCourseWithContent(ctx context.Context, courseID string) (models.Course, error) {
	var course models.Course

	query := `
        SELECT 
            c.id, c.name, c.description, 
            to_char(c.updated, 'YYYY-MM-DD"T"HH24:MI:SS.MS"Z"') as updated,
            c.logo_path, c.is_published,
            COALESCE(
                json_agg(
                    json_build_object(
                        'id', ch.id,
                        'name', ch.name,
                        'description', ch.description,
                        'updated', to_char(ch.updated, 'YYYY-MM-DD"T"HH24:MI:SS.MS"Z"'),
                        'position', ch.position,
                        'lessons', (
                            SELECT COALESCE(
                                json_agg(
                                    json_build_object(
                                        'id', l.id,
                                        'title', l.title,
                                        'text', l.text,
                                        'updated', to_char(l.updated, 'YYYY-MM-DD"T"HH24:MI:SS.MS"Z"')
                                    ) ORDER BY l.updated
                                ), '[]'::json
                            )
                            FROM lessons l
                            WHERE l.chapter_id = ch.id
                        )
                    ) ORDER BY ch.position
                ) FILTER (WHERE ch.id IS NOT NULL),
                '[]'::json
            ) as chapters
        FROM courses c
        LEFT JOIN chapters ch ON ch.course_id = c.id
        WHERE c.id = $1 AND c.is_published = true
        GROUP BY c.id
    `

	var chaptersJSON []byte
	var logoPath sql.NullString
	var updatedStr string

	err := s.psql.QueryRow(ctx, query, courseID).Scan(
		&course.ID, &course.Name, &course.Description,
		&updatedStr, &logoPath, &course.IsPublished,
		&chaptersJSON,
	)

	if err != nil {
		return course, err
	}

	// Парсим время
	course.Updated, err = time.Parse("2006-01-02T15:04:05.999Z", updatedStr)
	if err != nil {
		return course, fmt.Errorf("failed to parse updated time: %w", err)
	}

	if logoPath.Valid {
		course.LogoPath = &logoPath.String
	}

	if err := json.Unmarshal(chaptersJSON, &course.Chapters); err != nil {
		return course, fmt.Errorf("failed to unmarshal chapters: %w", err)
	}

	return course, nil
}

func (s *Storage) GetCourses(ctx context.Context) ([]models.Course, error) {
	var courses []models.Course

	query := `
		SELECT 
			c.id, 
			c.name, 
			c.description, 
			c.updated, 
			c.logo_path 
		FROM 
			courses c
		WHERE 
			is_published = true
		ORDER BY 
			c.name;
	`

	rows, err := s.psql.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		var course models.Course
		var logoPath sql.NullString
		var updated time.Time

		err := rows.Scan(
			&course.ID,
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

	return courses, nil
}

func (s *Storage) GetLesson(ctx context.Context, id string) (models.Lesson, error) {
	query := `
		SELECT 
			l.id, 
			l.title, 
			l.text, 
			l.updated
		FROM 
			lessons l
		WHERE
			l.id = $1
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
