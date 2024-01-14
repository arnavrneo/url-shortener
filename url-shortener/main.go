package main

import (
	"context"
	"fmt"
	"url-shortener/applications"
)

func main() {
	app := applications.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Errorf("error starting the server: %w", err)
	}
}
