package main

import (
	"fmt"
	"net/http"

	"student-api/routes"
)

func main() {

	routes.RegisterRoutes()

	fmt.Println("🚀 Server Running on :8080")

	http.ListenAndServe(":8080", nil)

}