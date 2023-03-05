package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api fiber.Router) {
	movies(api.Group("/movies"))
}
