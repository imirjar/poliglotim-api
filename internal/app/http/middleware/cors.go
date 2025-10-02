package middleware

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CORS() mux.MiddlewareFunc {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Разрешить все источники
		// handlers.AllowedOrigins([]string{"http://localhost:59889"}), // Или конкретный источник
		// handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedMethods([]string{"GET"}),
		handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}),
		handlers.ExposedHeaders([]string{"Link"}),
		// handlers.AllowCredentials(),
		handlers.MaxAge(300),
	)
}
