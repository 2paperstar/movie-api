package router

import (
	"context"
	"strings"

	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/service"
	"github.com/gofiber/fiber/v2"
)

func RequireAuth(c *fiber.Ctx) error {
	authorization := c.GetReqHeaders()["Authorization"]
	authorizationToken := strings.Split(authorization, " ")
	if len(authorizationToken) < 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token := authorizationToken[1]

	payload, err := service.VerifyJwt(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	user, err := database.GetUserByUID(context.TODO(), payload.UID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	c.Locals("user", user)
	return c.Next()
}
