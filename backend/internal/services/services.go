// Package services
package services

import "github.com/PrinceNarteh/pos/internal/repo"

type Services struct{}

func NewServices(r *repo.Repo) *Services {
	return &Services{}
}
