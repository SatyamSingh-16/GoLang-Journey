package middleware

import (
	"net/http"
	"student-api/response"
)

func ValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			apiKey := r.Header.Get("X-API-Key")

			if apiKey == "" {

				response.Error(w, http.StatusUnauthorized, "API Key Missing")
				return

			}
		}

		next(w, r)

	}
}
