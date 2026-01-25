// Package services
package services

import "github.com/PrinceNarteh/pos/internal/repositories"

type Services struct {
	Auth AuthService
}

func NewServices(r *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: r},
	}
}
