package handlers

import (
	"errors"
	"net/http"

	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UserHandler interface {
	FindByID(fiber.Ctx) error
}

type userHandler struct {
	svc *services.Services
}

func (h *userHandler) FindByID(c fiber.Ctx) error {
	id := fiber.Params(c, "id", "")

	user, err := h.svc.User.FindByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(
				ErrResponse(http.StatusNotFound, "user not found"),
			)
		}
		return c.Status(http.StatusInternalServerError).JSON(
			ErrResponse(http.StatusInternalServerError, err.Error()),
		)
	}

	return c.JSON(SuccessResponse(http.StatusOK, user))
}
