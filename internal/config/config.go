package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ListeningPort         string
	StudentServiceURL     string
	CsrfTokenSecret       string
	CsrfTokenTimeInMs     string
	JwtAccessTokenSecret  string
	JwtRefreshTokenSecret string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found:", err)
	}

	cfg := &Config{
		ListeningPort:         getEnvOrPanic("LISTENING_PORT"),
		StudentServiceURL:     getEnvOrPanic("STUDENT_SERVICE_URL"),
		CsrfTokenSecret:       os.Getenv("CSRF_TOKEN_SECRET"),
		CsrfTokenTimeInMs:     os.Getenv("CSRF_TOKEN_TIME_IN_MS"),
		JwtAccessTokenSecret:  os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
		JwtRefreshTokenSecret: os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
	}

	return cfg
}

func getEnvOrPanic(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required env: %s", key)
	}
	return val
}
