package main

import (
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
	err := app.Start()
	if err != nil {
		log.Fatal("error starting the server")
	}
}
