package router

import (
	"github.com/2paperstar/movie-api/controller"
	"github.com/gofiber/fiber/v2"
)

func users(api fiber.Router) {
	api.Get("/me", RequireAuth, controller.GetMyInformation)
	api.Put("/me", RequireAuth, controller.UpdateMyInformation)
}
