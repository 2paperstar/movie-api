package main

import (
	"log"
	"os"

	"github.com/2paperstar/movie-api/router"

	"github.com/joho/godotenv"

	_ "github.com/2paperstar/movie-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Movie API
// @description This is a sample server for a movie API.
// @version 1.0.0

// @host https://test.com:8080
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + os.Getenv("PORT"))
}
