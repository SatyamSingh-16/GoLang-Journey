package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(
				w,
				"Authorization header is missing",
				http.StatusUnauthorized,
			)
			return
		}
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {

			http.Error(
				w,
				"Invalid Authorization header",
				http.StatusUnauthorized,
			)

			return
		}
		tokenString := parts[1]
		token, err := jwt.Parse(
			tokenString,

			func(token *jwt.Token) (interface{}, error) {

				return []byte("my-super-secret-key"), nil

			},
		)
		if err != nil || !token.Valid {

			http.Error(
				w,
				"Invalid token",
				http.StatusUnauthorized,
			)

			return
		}
		claims := token.Claims.(jwt.MapClaims)
		ctx := context.WithValue(
			r.Context(),
			"userId",
			claims["userId"],
		)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
