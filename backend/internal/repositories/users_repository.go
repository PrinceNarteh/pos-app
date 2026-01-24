package repositories

import (
	"context"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ UserRepository = (*userRepository)(nil)

type UserRepository interface {
	FindAll(context.Context) ([]models.User, error)
	FindByID(context.Context, int) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	Delete(context.Context, int) error
}

type userRepository struct {
	pool *pgxpool.Pool
}

func (u *userRepository) findBy(ctx context.Context, key, value any) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sql := `
		SELECT id, name, username, email, password, role
		FROM users
		WHERE $1 = $2
	`

	user := new(models.User)
	if err := u.pool.QueryRow(ctx, sql, key, value).
		Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
		); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	return u.findBy(ctx, "id", id)
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	return u.findBy(ctx, "email", email)
}

func (u *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	return u.findBy(ctx, "username", username)
}

func (u *userRepository) FindAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sql := `
		SELECT id, name, username, email, password, role
		FROM users
	`

	rows, err := u.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
		); err != nil {
			return nil, fmt.Errorf("error scanning user row: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sql := `
		DELETE FROM users WHERE id = $1
	`

	u.pool.

	return nil
}
