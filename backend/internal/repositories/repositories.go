// Package repositories
package repositories

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var QueryTimeoutDuration = time.Second * 5

type Repositories struct {
	User UserRepository
}

func NewRepo(db *pgxpool.Pool) *Repositories {
	return &Repositories{}
}
