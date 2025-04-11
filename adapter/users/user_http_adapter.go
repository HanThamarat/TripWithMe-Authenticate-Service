package adapter

import (
	core "github.com/HanThamarat/TripWithMe-Authenticate-Service/core/users"
	"github.com/HanThamarat/TripWithMe-Authenticate-Service/packages/response"
	"github.com/gofiber/fiber/v2"
)

type HttpUserHandler struct {
	service core.UserService
}

// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body core.User true "User object"
// @Success 201 {object} response.Response{data=core.User} "User created successfully"
// @Failure 400 {object} response.Response{} "Invalid request"
// @Router /user [post]
func (h *HttpUserHandler) CreateUser(c *fiber.Ctx) error {
	var user core.User;

	if err := c.BodyParser(&user); err != nil {
		return response.SendErrorHandler(c, fiber.StatusBadRequest, "Invalid your request", err.Error());
	}
	
	createdUser, err := h.service.Save(user);

	if err != nil {
		return response.SendErrorHandler(c, fiber.StatusBadRequest, "Invalid your request", err.Error());
	}

	return response.SendResponseHandler(c, fiber.StatusCreated, "Createing a new user successfully", createdUser);
}
func NewHttpUserHandler(service core.UserService) *HttpUserHandler {
	return &HttpUserHandler{service: service}
}