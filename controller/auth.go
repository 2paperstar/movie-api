package controller

import "github.com/gofiber/fiber/v2"

// @Summary Authorize with credential
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
