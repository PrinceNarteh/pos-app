package services

import (
	"context"

	"github.com/PrinceNarteh/pos/internal/dto"
	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/PrinceNarteh/pos/internal/repositories"
)

var _ CategoryService = (*categoryService)(nil)

type CategoryService interface {
	FindAll(ctx context.Context) ([]models.Category, error)
	FindByID(ctx context.Context, categoryID string) (*models.Category, error)
	FindByName(ctx context.Context, name string) (*models.Category, error)
	Create(ctx context.Context, data *dto.CategoryDTO) (*models.Category, error)
	Update(ctx context.Context, categoryID string, data *dto.CategoryDTO) (*models.Category, error)
	Delete(ctx context.Context, categoryID string) error
}

type categoryService struct {
	repo *repositories.Repositories
}

func (s *categoryService) FindAll(ctx context.Context) ([]models.Category, error) {
	return s.repo.Category.FindAll(ctx)
}

func (s *categoryService) FindByID(ctx context.Context, categoryID string) (*models.Category, error) {
	return s.repo.Category.FindByID(ctx, categoryID)
}

func (s *categoryService) FindByName(ctx context.Context, name string) (*models.Category, error) {
	return s.repo.Category.FindByName(ctx, name)
}

func (s *categoryService) Create(ctx context.Context, data *dto.CategoryDTO) (*models.Category, error) {
	category := models.Category{
		Name: data.Name,
	}
	if err := s.repo.Category.Create(ctx, &category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *categoryService) Update(ctx context.Context, categoryID string, data *dto.CategoryDTO) (*models.Category, error) {
	return s.repo.Category.Update(ctx, categoryID, data.Name)
}

func (s *categoryService) Delete(ctx context.Context, categoryID string) error {
	return s.repo.Category.Delete(ctx, categoryID)
}
