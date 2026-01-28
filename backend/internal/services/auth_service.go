package services

import (
	"context"
	"database/sql"
	"errors"
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
	Register(context.Context, *models.RegisterUserDTO) (*models.UserResponse, error)
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

	token, err := utils.GenerateAccessToken(user)
	if err != nil {
		return &models.UserResponse{}, fmt.Errorf("error: %w", err)
	}

	userResponse := &models.UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		Token:    token,
	}

	return userResponse, nil
}

func (s *authService) Register(ctx context.Context, registerDTO *models.RegisterUserDTO) (*models.UserResponse, error) {
	emailExists, err := s.repo.User.FindByEmail(ctx, registerDTO.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.UserResponse{}, err
	}
	if emailExists != nil {
		return &models.UserResponse{}, fmt.Errorf("user with email %q exists", registerDTO.Email)
	}

	usernameExists, err := s.repo.User.FindByUsername(ctx, registerDTO.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return &models.UserResponse{}, err
	}
	if usernameExists != nil {
		return &models.UserResponse{}, fmt.Errorf("user with username %q exists", registerDTO.Email)
	}

	hashedPassword, err := utils.Hash(registerDTO.Password)
	if err != nil {
		return &models.UserResponse{}, err
	}

	registerDTO.Password = hashedPassword
	userResponse, err := s.repo.User.Create(ctx, registerDTO)
	if err != nil {
		return &models.UserResponse{}, err
	}

	return &models.UserResponse{
		ID:       userResponse.ID,
		Email:    userResponse.Email,
		Username: userResponse.Username,
		Name:     userResponse.Name,
		Role:     userResponse.Role,
	}, nil
}
