package entities

import "time"

// User - упрощенная модель пользователя, которая приходит из JWT токена
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
