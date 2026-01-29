package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

var (
	_                    UserRepository = (*userRepository)(nil)
	ErrDuplicateEmail                   = errors.New("email already in used")
	ErrDuplicateUsername                = errors.New("username already in used")
)

type UserRepository interface {
	FindAll(context.Context) ([]models.User, error)
	FindByID(context.Context, int) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	Create(context.Context, *models.User) error
	Delete(context.Context, int) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) usersTbl() gorm.Interface[models.User] {
	return gorm.G[models.User](u.db)
}

func (u *userRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	user, err := u.usersTbl().Where("id = $1", id).First(ctx)
	return &user, err
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.usersTbl().Where("email = $1", email).First(ctx)
	return &user, err
}

func (u *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := u.usersTbl().Where("username = $1", username).First(ctx)
	return &user, err
}

func (u *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	return u.usersTbl().Find(ctx)
}

func (u *userRepository) Create(ctx context.Context, dto *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := u.usersTbl().Create(ctx, dto)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // unique_violation
				if pgErr.ConstraintName == "users_email_key" {
					return ErrDuplicateEmail
				}
				if pgErr.ConstraintName == "users_username_key" {
					return ErrDuplicateUsername
				}
			}
		}
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}

func (u *userRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rowsAffected, err := u.usersTbl().Where("id = ?", id).Delete(ctx)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}
