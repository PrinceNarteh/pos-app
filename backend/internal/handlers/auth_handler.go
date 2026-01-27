package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v2"
)

var _ AuthHandler = (*authHandler)(nil)

type AuthHandler interface {
	Login(*fiber.Ctx) error
	Register(*fiber.Ctx) error
}

type authHandler struct {
	svc *services.Services
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	reqBody := new(models.LoginDTO)
	if err := c.BodyParser(reqBody); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if err := reqBody.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	userResponse, err := h.svc.Auth.Login(c.Context(), reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": userResponse,
	})
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	registerDTO := new(models.RegisterUserDTO)
	if err := c.BodyParser(registerDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := registerDTO.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userResponse, err := h.svc.Auth.Register(c.Context(), registerDTO)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": userResponse,
	})
}
