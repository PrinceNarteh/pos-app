package main

import (
	"github.com/PrinceNarteh/pos/internal/handlers"
	"github.com/PrinceNarteh/pos/internal/middleware"
	"github.com/gofiber/fiber/v3"
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
	// auth
	auth := app.Group("/auth")
	auth.Post("/login", r.Handlers.Auth.Login)
	auth.Post("/register", r.Handlers.Auth.Register)

	app.Use(middleware.AuthMiddleware())
	// users
	user := app.Group("/users")
	user.Get("/:id", r.Handlers.User.FindByID)
}
