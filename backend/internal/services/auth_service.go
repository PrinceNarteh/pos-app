package services

import (
	"context"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/PrinceNarteh/pos/internal/repositories"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gohugoio/hugo/tpl/fmt"
)

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Login(context.Context, models.LoginDTO) (*models.UserResponse, error)
}

type authService struct {
	repo *repositories.Repositories
}

func (s *authService) Login(ctx context.Context, loginDTO models.LoginDTO) (*models.UserResponse, error) {
	user := new(models.User)

	err := validation.Validate(loginDTO.UsernameOrEmail, is.Email)
	if err == nil {
		user, err = s.repo.User.FindByEmail(ctx, loginDTO.UsernameOrEmail)
		if err != nil {
			return &models.UserResponse{}, fmt.Errorf("user with email %q not found", loginDTO.UsernameOrEmail)
		}
	} else {
		user, err = s.repo.User.FindByUsername(ctx, loginDTO.UsernameOrEmail)
		if err != nil {
			return &models.UserResponse{}, fmt.Errorf("user with username %q not found", loginDTO.UsernameOrEmail)
		}
	}

	userResponse := &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		Token:    "",
	}

	return userResponse, nil
}
