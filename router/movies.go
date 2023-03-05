package router

import (
	"github.com/2paperstar/movie-api/controller"
	"github.com/gofiber/fiber/v2"
)

func movies(api fiber.Router) {
	api.Get("/", controller.GetMovies)
	api.Get("/:id", controller.GetMovieDetail)

	api.Get("/:id/comments", controller.GetMovieComments)
}
