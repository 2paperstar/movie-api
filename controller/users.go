package controller

import (
	"github.com/2paperstar/movie-api/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get my information
// @Description Get information of logged in user
// @Tags users
// @Produce json
// @Success 200 {object} model.UserInfo
// @Router /users/me [get]
// @Security JWT
func GetMyInformation(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.UserInfo)
	return c.JSON(user)
}
