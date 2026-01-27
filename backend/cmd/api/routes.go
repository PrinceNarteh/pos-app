package main

import (
	"github.com/PrinceNarteh/pos/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type routes struct {
	Handlers *handlers.Handlers
}

func NewRoutes(handlers *handlers.Handlers) *routes {
	return &routes{
		Handlers: handlers,
	}
}

func (r *routes) initRoutes(app fiber.Router) {
	auth := app.Group("/auth")
	auth.Post("/login", r.Handlers.Auth.Login)
	auth.Post("/register", r.Handlers.Auth.Register)

	user := app.Group("/users")
	user.Get("/:id", r.Handlers.User.FindByID)
}
