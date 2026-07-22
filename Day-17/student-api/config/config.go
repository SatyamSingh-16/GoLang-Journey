package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWTSecret []byte

func LoadConfig() {
	err := godotenv.Load()

	if err != nil {

		log.Fatal("Error loading .env file")

	}
	JWTSecret = []byte(
		os.Getenv("JWT_SECRET"),
	)
}
