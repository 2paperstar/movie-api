package controller

import "github.com/gofiber/fiber/v2"

// @Summary Register with credential - not implemented
// @Description Register with id and password
// @Tags auth
// @Accept json
// @Product json
// @Param request body model.RegisterForm true "Register"
// @Success 200 {object} model.AuthResponse
// @Router /auth/register [post]
func RegisterWithCredential(c *fiber.Ctx) error {
	return nil
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
