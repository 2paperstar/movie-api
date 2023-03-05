package main

import (
	_ "github.com/2paperstar/movie-api/docs"
	"github.com/2paperstar/movie-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Movie API
// @description This is a sample server for a movie API.
// @version 1.0.1

// @host test.paperst.ar
func setup() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	router.SetupRouter(app)
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	return app
}
