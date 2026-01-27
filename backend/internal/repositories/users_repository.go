package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/PrinceNarteh/pos/internal/models"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	_                    UserRepository = (*userRepository)(nil)
	ErrDuplicateEmail                   = errors.New("a user with this email already exists")
	ErrDuplicateUsername                = errors.New("a user with this username already exists")
)

type UserRepository interface {
	FindAll(context.Context) ([]models.User, error)
	FindByID(context.Context, int) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	Create(context.Context, *models.RegisterUserDTO) (*models.User, error)
	Delete(context.Context, int) error
}

type userRepository struct {
	pool *pgxpool.Pool
}

func (u *userRepository) findBy(ctx context.Context, sql string, args ...any) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	user := new(models.User)
	if err := u.pool.QueryRow(ctx, sql, args...).
		Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
		); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	sql := `
		SELECT id, name, username, email, password, role
		FROM users
		WHERE id = $1
	`
	return u.findBy(ctx, sql, id)
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	sql := `
		SELECT id, name, username, email, password, role
		FROM users
		WHERE email = $1
	`
	return u.findBy(ctx, sql, email)
}

func (u *userRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	sql := `
		SELECT id, name, username, email, password, role
		FROM users
		WHERE username = $1
	`
	return u.findBy(ctx, sql, username)
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

func (u *userRepository) Create(ctx context.Context, dto *models.RegisterUserDTO) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sql := `
		INSERT INTO users (name, username, email, password, role)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, username, email, password, role
	`

	user := new(models.User)
	if err := u.pool.QueryRow(ctx, sql,
		dto.Name,
		dto.Username,
		dto.Email,
		dto.Password,
		dto.Role,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" { // unique_violation
				if pgErr.ConstraintName == "users_email_key" {
					return nil, ErrDuplicateEmail
				}
				if pgErr.ConstraintName == "users_username_key" {
					return nil, ErrDuplicateUsername
				}
			}
		}
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (u *userRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	sql := `
		DELETE FROM users WHERE id = $1
	`

	res, err := u.pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}
