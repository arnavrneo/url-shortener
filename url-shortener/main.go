package main

import (
	"log"
	"net/http"
	"os"
	"time"
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
	router := applications.LoadRoutes()

	server := &http.Server{
		Addr:         ":" + os.Args[2],
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting the server %s", err)
	} else {
		log.Println("server running on ")
	}
}
