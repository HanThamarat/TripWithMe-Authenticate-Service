package adapter

import (
	core "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/auth"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/response"

	"github.com/gofiber/fiber/v2"
)

type HttpAuthHandler struct {
	service core.AuthService
}

func NewHttpAuthHandler(service core.AuthService) *HttpAuthHandler {
	return &HttpAuthHandler{service: service}
}

// @Summary Authenticate a user
// @Description Authenticate a user with the provided credentials
// @Tags Auth
// @Accept json
// @Produce json
// @Param auth body core.Auth true "Authentication credentials"
// @Success 200 {object} response.Response{data=core.Auth} "Authenticated successfully"
// @Failure 401 {object} response.Response{} "Invalid credentials"
// @Router /auth [post]
func (h *HttpAuthHandler) Authenticate(c *fiber.Ctx) error {
	var auth core.Auth

	if err := c.BodyParser(&auth); err != nil {
		return response.SendErrorHandler(c, fiber.StatusBadRequest, "Invalid your request", err.Error());
	}

	authenticatedUser, err := h.service.Authenticate(auth)

	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusUnauthorized, "Invalid credentials", err.Error());
	}

	return response.SendResponseHandler(c, fiber.StatusOK, "Authenticated successfully", authenticatedUser);
}

