// Package services
package services

import "github.com/PrinceNarteh/pos/internal/repositories"

type Services struct {
	Auth AuthService
	User UserService
}

func NewServices(r *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: r},
		User: &userService{repo: r},
	}
}
