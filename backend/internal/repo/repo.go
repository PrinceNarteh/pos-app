// Package repo
package repo

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct{}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{}
}
