package middleware

import (
	"context"
	"net/http"
	"strings"
	"student-api/config"
	"student-api/response"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			response.Error(w, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(w, http.StatusUnauthorized, "Invalid Authorization header")
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return config.JWTSecret, nil
			},
		)

		if err != nil || !token.Valid {
			response.Error(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(
			r.Context(),
			"userId",
			claims["userId"],
		)

		r = r.WithContext(ctx)

		next(w, r)
	}
}
