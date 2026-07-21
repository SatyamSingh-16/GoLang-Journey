package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {

			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return

		}

		next(w, r)

	}
}