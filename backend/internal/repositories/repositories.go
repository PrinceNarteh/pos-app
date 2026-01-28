// Package repositories
package repositories

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrNotFound          = errors.New("record not found")
)

type Repositories struct {
	User UserRepository
}

func NewRepo(connPool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User: &userRepository{pool: connPool},
	}
}
