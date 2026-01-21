// Package handlers
package handlers

import "github.com/PrinceNarteh/pos/internal/services"

type Handlers struct{}

func NewHandlers(svc *services.Services) *Handlers {
	return &Handlers{}
}
