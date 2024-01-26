package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

// LoadEnvVar loads up the environment file
func LoadEnvVar() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env loaded: ERROR")
	}
	fmt.Println(".env loaded: OK")
}
