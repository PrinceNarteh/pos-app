package services

import (
	"context"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/PrinceNarteh/pos/internal/repositories"
)

var _ UserService = (*userService)(nil)

type UserService interface {
	FindByID(context.Context, string) (*models.User, error)
}

type userService struct {
	repo *repositories.Repositories
}

func (s *userService) FindByID(ctx context.Context, id string) (*models.User, error) {
	return s.repo.User.FindByID(ctx, id)
}
