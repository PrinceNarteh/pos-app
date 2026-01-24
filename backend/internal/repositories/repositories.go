// Package repositories
package repositories

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var QueryTimeoutDuration = time.Second * 5

type Repo struct{}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{}
}
