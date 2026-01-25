// Package services
package services

import "github.com/PrinceNarteh/pos/internal/repositories"

type Services struct{}

func NewServices(r *repositories.Repositories) *Services {
	return &Services{}
}
