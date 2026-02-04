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
	Login(context.Context, *models.LoginDTO) (*models.UserWithToken, error)
	Register(context.Context, *models.RegisterUserDTO) (*models.UserWithToken, error)
}

type authService struct {
	repo *repositories.Repositories
}

func (s *authService) Login(ctx context.Context, loginDTO *models.LoginDTO) (*models.UserWithToken, error) {
	user := new(models.User)
	const errMsg = "invalid email/username or password"

	err := validation.Validate(loginDTO.UsernameOrEmail, is.Email)
	if err == nil {
		user, err = s.repo.User.FindByEmail(ctx, loginDTO.UsernameOrEmail)
		if err != nil {
			return nil, fmt.Errorf(errMsg)
		}
	} else {
		user, err = s.repo.User.FindByUsername(ctx, loginDTO.UsernameOrEmail)
		if err != nil {
			return nil, fmt.Errorf(errMsg)
		}
	}

	if !utils.CompareHashAndPassword(loginDTO.Password, user.Password) {
		return nil, fmt.Errorf(errMsg)
	}

	token, err := utils.GenerateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	userResponse := &models.UserWithToken{
		User:  user,
		Token: token,
	}

	return userResponse, nil
}

func (s *authService) Register(ctx context.Context, registerDTO *models.RegisterUserDTO) (*models.UserWithToken, error) {
	emailExists, err := s.repo.User.FindByEmail(ctx, registerDTO.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if emailExists != nil {
		return nil, fmt.Errorf("user with email %q exists", registerDTO.Email)
	}

	usernameExists, err := s.repo.User.FindByUsername(ctx, registerDTO.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if usernameExists != nil {
		return nil, fmt.Errorf("user with username %q exists", registerDTO.Email)
	}

	hashedPassword, err := utils.Hash(registerDTO.Password)
	if err != nil {
		return nil, err
	}

	registerDTO.Password = hashedPassword
	user := models.User{
		Email:    registerDTO.Email,
		Username: registerDTO.Username,
		Name:     registerDTO.Name,
		Password: registerDTO.Password,
		Role:     "user",
	}
	err = s.repo.User.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateAccessToken(&user)
	if err != nil {
		return nil, fmt.Errorf("error: %w", err)
	}

	return &models.UserWithToken{
		User:  &user,
		Token: token,
	}, nil
}
