package main

import (
	"context"
	"log"
	"url-shortener/applications"
	"url-shortener/initializers"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDb()
}

func main() {
	app := applications.New()
	err := app.Start(context.TODO())
	if err != nil {
		log.Fatal("error starting the server")
	}
}
