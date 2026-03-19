package handlers

import (
	"net/http"

	"github.com/PrinceNarteh/pos/internal/dto"
	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v3"
)

var _ AuthHandler = (*authHandler)(nil)

type AuthHandler interface {
	Login(fiber.Ctx) error
	Register(fiber.Ctx) error
}

type authHandler struct {
	svc *services.Services
}

func (h *authHandler) Login(c fiber.Ctx) error {
	reqBody := new(dto.LoginDTO)
	if err := c.Bind().Body(reqBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, "request body missing"),
		)
	}

	if err := reqBody.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err),
		)
	}

	userResponse, err := h.svc.Auth.Login(c.Context(), reqBody)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err.Error()),
		)
	}

	return c.JSON(fiber.Map{
		"data": userResponse,
	})
}

func (h *authHandler) Register(c fiber.Ctx) error {
	registerDTO := new(dto.RegisterUserDTO)
	if err := c.Bind().Body(registerDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err.Error()),
		)
	}

	if err := registerDTO.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err.Error()),
		)
	}

	userResponse, err := h.svc.Auth.Register(c.Context(), registerDTO)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err.Error()),
		)
	}

	return c.JSON(fiber.Map{
		"data": userResponse,
	})
}
