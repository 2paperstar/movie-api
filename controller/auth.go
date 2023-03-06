package controller

import (
	"github.com/2paperstar/movie-api/database"
	"github.com/2paperstar/movie-api/model"
	"github.com/2paperstar/movie-api/service"
	"github.com/gofiber/fiber/v2"
)

// @Summary Register with credential
// @Description Register with id and password
// @Tags auth
// @Accept json
// @Product json
// @Param request body model.RegisterForm true "Register"
// @Success 200 {object} model.AuthResponse
// @Failure 409 {object} model.Error "Duplicated id"
// @Router /auth/register [post]
func RegisterWithCredential(c *fiber.Ctx) error {
	form := new(model.RegisterForm)
	if err := c.BodyParser(form); err != nil {
		return err

	}

	user, err := service.RegisterWithCredential(form)
	if err != nil {
		if err == database.ErrDuplicatedUser {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}

// @Summary Authorize with credential - not implemented
// @Description Authorize with id and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.Credential true "Credential"
// @Success 200 {object} model.AuthResponse
// @Router /auth/authorize [post]
func AuthorizeWithCredential(c *fiber.Ctx) error {
	return nil
}

// @Summary Authorize with refresh token - not implemented
// @Description Authorize with refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RefreshToken true "RefreshToken"
// @Success 200 {object} model.AuthResponse
// @Router /auth/refresh [post]
func AuthorizeWithRefreshToken(c *fiber.Ctx) error {
	return nil
}
