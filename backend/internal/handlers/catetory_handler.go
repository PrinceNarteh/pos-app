package handlers

import (
	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v3"
)

var _ CategoryHandler = (*categoryHandler)(nil)

type CategoryHandler interface {
	FindAll(fiber.Ctx) error
}

type categoryHandler struct {
	svc *services.Services
}

func (h *categoryHandler) FindAll(c fiber.Ctx) error {
	categories, err := h.svc.Category.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			ErrResponse(fiber.StatusInternalServerError, `internal server error`),
		)
	}

	return c.JSON(SuccessResponse(fiber.StatusOK, categories))
}
