package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(api *fiber.App) {
	movies(api.Group("/movies"))
	auth(api.Group("/auth"))
}
