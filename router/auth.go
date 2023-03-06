package router

import (
	"github.com/2paperstar/movie-api/controller"
	"github.com/gofiber/fiber/v2"
)

func auth(api fiber.Router) {
	api.Post("/authorize", controller.AuthorizeWithCredential)
	api.Post("/refresh", controller.AuthorizeWithRefreshToken)
	api.Post("/register", controller.RegisterWithCredential)
}
