package controller

import "github.com/gofiber/fiber/v2"

// @Summary Authorize with credential - 구현 안 됨
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

// @Summary Authorize with refresh token - 구현 안 됨
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
