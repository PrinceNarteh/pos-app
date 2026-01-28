package handlers

import (
	"errors"
	"net/http"

	"github.com/PrinceNarteh/pos/internal/repositories"
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
		return c.Status(http.StatusBadRequest).JSON(
			ErrResponse(http.StatusBadRequest, "invalid user id"),
		)
	}

	user, err := h.svc.User.FindByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
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
