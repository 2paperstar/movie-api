package controller

import (
	"github.com/2paperstar/movie-api/model"
	"github.com/2paperstar/movie-api/service"
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

// @Summary Update my information
// @Description Update information of logged in user
// @Tags users
// @Accept json
// @Produce json
// @Param body body model.UserInfoForm true "User information"
// @Success 200 {object} model.UserInfo
// @Router /users/me [patch]
// @Security JWT
func UpdateMyInformation(c *fiber.Ctx) error {
	user := c.Locals("user").(*model.UserInfo)
	form := new(model.UserInfoForm)
	if err := c.BodyParser(form); err != nil {
		return err
	}

	user, err := service.UpdateUser(user.UID, form)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}
