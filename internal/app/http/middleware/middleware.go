package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func New() *Middleware {
	return &Middleware{}
}

type Middleware struct {
}

func (m *Middleware) CORS() mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Разрешить все источники
		// handlers.AllowedOrigins([]string{"http://localhost:59889"}), // Или конкретный источник
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}),
		handlers.ExposedHeaders([]string{"Link"}),
		// handlers.AllowCredentials(),
		handlers.MaxAge(300),
	)
}

func Auth() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractToken(r)

			if token != "" {
				log.Printf("JWT Token received: %s (from %s %s)",
					maskToken(token), r.Method, r.URL.Path)
			} else {
				log.Printf("No JWT token found in request: %s %s", r.Method, r.URL.Path)
			}
			log.Print("AAAUUUTTTTHHH")
			next.ServeHTTP(w, r)
		})
	}
}

// maskToken маскирует токен для безопасности
func maskToken(token string) string {
	if len(token) <= 8 {
		return "***"
	}
	return token[:4] + "..." + token[len(token)-4:]
}

// extractToken извлекает JWT токен из запроса
func extractToken(r *http.Request) string {
	// Из заголовка Authorization
	if authHeader := r.Header.Get("Authorization"); authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			return parts[1]
		}
		return authHeader
	}

	// Из query параметра
	if token := r.URL.Query().Get("token"); token != "" {
		return token
	}

	return ""
}
