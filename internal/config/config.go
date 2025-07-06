package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var StudentServiceURL string
var ListeningPort string

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found:", err)
	}

	ListeningPort = os.Getenv("LISTENING_PORT")
	if ListeningPort == "" {
		log.Fatal("Listening port required env: LISTENING_PORT")
	}

	StudentServiceURL = os.Getenv("STUDENT_SERVICE_URL")
	if StudentServiceURL == "" {
		log.Fatal("Missing required env: STUDENT_SERVICE_URL")
	}
}
