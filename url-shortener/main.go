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
	err := initializers.ConnectToDb(uri)
	if err != nil {
		panic("cannot connect to the monogodb cluster.")
	}

}

func main() {
	app := applications.New()
	err := app.Start()
	if err != nil {
		log.Fatal("error starting the server")
	}
}
