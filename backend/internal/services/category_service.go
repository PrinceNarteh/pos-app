package services

import (
	"context"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/PrinceNarteh/pos/internal/repositories"
)

var _ CategoryService = (*categoryService)(nil)

type CategoryService interface {
	FindAll(context.Context) ([]models.Category, error)
}

type categoryService struct {
	repo *repositories.Repositories
}

func (s *categoryService) FindAll(ctx context.Context) ([]models.Category, error) {
	return s.repo.Category.FindAll(ctx)
}
