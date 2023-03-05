//go:build debug

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	app := setup()
	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
