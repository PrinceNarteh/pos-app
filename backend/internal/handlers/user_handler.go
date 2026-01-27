package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	FindByID(*fiber.Ctx) error
}

type userHandler struct {
	svc *services.Services
}

func (h *userHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id", 0)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid iD",
		})
	}

	user, err := h.svc.User.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}
