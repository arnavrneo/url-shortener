package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"url-shortener/routes"
)

func init() {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set.")
	}
}

func main() {
	router := routes.LoadRoutes()

	server := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
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
