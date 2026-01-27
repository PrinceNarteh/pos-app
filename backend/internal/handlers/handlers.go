// Package handlers
package handlers

import "github.com/PrinceNarteh/pos/internal/services"

type Handlers struct {
	Auth AuthHandler
	User UserHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth: &authHandler{svc: svc},
		User: &userHandler{svc: svc},
	}
}
