package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

// LoadEnvVar loads up the environment file
func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("cannot load .env file")
	}
}
