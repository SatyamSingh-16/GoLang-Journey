package middleware

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("========== REQUEST ==========")
		fmt.Println("Method :", r.Method)
		fmt.Println("Path   :", r.URL.Path)
		fmt.Println("=============================")

		next(w, r)
	}
}