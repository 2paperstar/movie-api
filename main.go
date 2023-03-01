package main

import (
	"log"
	"os"

	"github.com/2paperstar/movie-api/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := router.SetupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
