package handlers

import (
	"errors"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/dto"
	"github.com/PrinceNarteh/pos/internal/repositories"
	"github.com/PrinceNarteh/pos/internal/services"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var _ CategoryHandler = (*categoryHandler)(nil)

type CategoryHandler interface {
	FindAllCategories(fiber.Ctx) error
	FindCategoryByID(fiber.Ctx) error
	CreateCategory(fiber.Ctx) error
	UpdateCategory(fiber.Ctx) error
	DeleteCategory(fiber.Ctx) error
}

type categoryHandler struct {
	svc *services.Services
}

func (h *categoryHandler) FindAllCategories(c fiber.Ctx) error {
	categories, err := h.svc.Category.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			ErrResponse(fiber.StatusInternalServerError, `internal server error`),
		)
	}

	return c.JSON(SuccessResponse(fiber.StatusOK, categories))
}

func (h *categoryHandler) FindCategoryByID(c fiber.Ctx) error {
	id := c.Params("id", "")
	category, err := h.svc.Category.FindByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(
				ErrResponse(fiber.StatusNotFound, "category not found"),
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusInternalServerError, "internale server error"),
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		SuccessResponse(fiber.StatusOK, category),
	)
}

func (h *categoryHandler) CreateCategory(c fiber.Ctx) error {
	categoryDTO := new(dto.CreateCategoryDTO)
	if err := c.Bind().Body(categoryDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, ErrRequestBodyNotFound),
		)
	}
	if err := categoryDTO.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err),
		)
	}

	fmt.Println(categoryDTO)

	category, err := h.svc.Category.Create(c.Context(), categoryDTO)
	if err != nil {
		if errors.Is(err, repositories.ErrDuplicateCategoryName) {
			return c.Status(fiber.StatusConflict).JSON(
				ErrResponse(fiber.StatusConflict, repositories.ErrDuplicateCategoryName.Error()),
			)
		}
		if errors.Is(err, repositories.ErrDuplicateCategoryCode) {
			return c.Status(fiber.StatusConflict).JSON(
				ErrResponse(fiber.StatusConflict, repositories.ErrDuplicateCategoryCode.Error()),
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			ErrResponse(fiber.StatusInternalServerError, `internal server error`),
		)
	}
	return c.Status(fiber.StatusCreated).JSON(
		SuccessResponse(fiber.StatusCreated, category),
	)
}

func (h *categoryHandler) UpdateCategory(c fiber.Ctx) error {
	id := c.Params("id", "")
	updateDTO := new(dto.UpdateCategoryDTO)
	if err := c.Bind().Body(updateDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, ErrRequestBodyNotFound),
		)
	}
	if err := updateDTO.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, err),
		)
	}
	if updateDTO.Name == "" && updateDTO.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrResponse(fiber.StatusBadRequest, "both name and update cannot be empty"),
		)
	}
	category, err := h.svc.Category.Update(c.Context(), id, updateDTO)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			ErrResponse(fiber.StatusInternalServerError, err.Error()),
		)
	}
	return c.Status(fiber.StatusOK).JSON(SuccessResponse(fiber.StatusOK, category))
}

func (h *categoryHandler) DeleteCategory(c fiber.Ctx) error {
	id := c.Params("id", "")
	if err := h.svc.Category.Delete(c.Context(), id); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(
				ErrResponse(fiber.StatusNotFound, "category not found"),
			)
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(
				ErrResponse(fiber.StatusInternalServerError, "internal server error"),
			)
		}
	}
	return c.Status(fiber.StatusNoContent).JSON(
		SuccessResponse(fiber.StatusNoContent, "category deleted successfully"),
	)
}
