package services

import (
	"context"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/PrinceNarteh/pos/internal/repositories"
	"github.com/PrinceNarteh/pos/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Login(context.Context, *models.LoginDTO) (*models.UserResponse, error)
}

type authService struct {
	repo *repositories.Repositories
}

func (s *authService) Login(ctx context.Context, loginDTO *models.LoginDTO) (*models.UserResponse, error) {
	user := new(models.User)
	const errMsg = "invalid email/username or password"

	err := validation.Validate(loginDTO.UsernameOrEmail, is.Email)
	if err == nil {
		user, err = s.repo.User.FindByEmail(ctx, loginDTO.UsernameOrEmail)
		if err != nil {
			return &models.UserResponse{}, fmt.Errorf(errMsg)
		}
	} else {
		user, err = s.repo.User.FindByUsername(ctx, loginDTO.UsernameOrEmail)
		if err != nil {
			return &models.UserResponse{}, fmt.Errorf(errMsg)
		}
	}

	if !utils.CompareHashAndPassword(loginDTO.Password, user.Password) {
		return &models.UserResponse{}, fmt.Errorf(errMsg)
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
