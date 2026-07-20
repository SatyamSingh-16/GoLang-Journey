package config
import "os"
type Config struct {
	Port         string
	DatabasePath string
	JWTSecret    string
	AppEnv       string
}
func Load() Config {

	return Config{

		Port: os.Getenv("PORT"),

		DatabasePath: os.Getenv("DATABASE_PATH"),

		JWTSecret: os.Getenv("JWT_SECRET"),

		AppEnv: os.Getenv("APP_ENV"),
	}
}