package main

import (
	"log"
	"os"
	"url-shortener/applications"
	"url-shortener/initializers"
)

func init() {
	// load up the .env
	initializers.LoadEnvVar()

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set.")
	}

	// connect to db
	initializers.ConnectToDb(uri)
}

func main() {
	app := applications.New()
	err := app.Start()
	if err != nil {
		log.Fatal("error starting the server")
	}
}
