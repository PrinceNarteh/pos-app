package handlers

import (
	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v2"
)

var _ AuthHandler = (*authHandler)(nil)

type AuthHandler interface{}

type authHandler struct {
	Services *services.Services
}

func (h *authHandler) Login(c *fiber.Ctx) error {
}
