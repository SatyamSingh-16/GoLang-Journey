package middleware

import (
	"net/http"
)

func ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			apiKey := r.Header.Get("X-API-Key")

			if apiKey == "" {

				http.Error(w, "API Key Missing", http.StatusUnauthorized)
				return

			}
		}

		next(w, r)

	}
}