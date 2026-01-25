// Package handlers
package handlers

import "github.com/PrinceNarteh/pos/internal/services"

type Handlers struct {
	Auth AuthHandler
}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{
		Auth: &authHandler{Services: svc},
	}
}
