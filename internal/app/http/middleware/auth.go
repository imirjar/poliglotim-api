package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/imirjar/poliglotim-api/internal/domain/entities"
)

func Auth() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := extractToken(r)

			if token != "" {
				user, err := validateUser(token)
				if err != nil {
					log.Printf("Invalid user: %s %s", r.Method, r.URL.Path)
					http.Error(w, "", http.StatusUnauthorized)
					return
				}

				// log.Print(r.URL)
				useredPath := fmt.Sprintf(r.URL.Path + "?user=" + user.ID)
				r.URL.Path = useredPath
				log.Print(useredPath)

				next.ServeHTTP(w, r)
				return

			} else {
				// log.Printf("No JWT token found in request: %s %s", r.Method, r.URL.Path)
				// http.Error(w, "", http.StatusUnauthorized)
				next.ServeHTTP(w, r)
				return
			}

		})
	}
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

func validateUser(token string) (*entities.User, error) {
	user := &entities.User{
		ID: token,
	}
	return user, nil
}
