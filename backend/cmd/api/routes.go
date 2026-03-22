package main

import (
	"github.com/PrinceNarteh/pos/internal/handlers"
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

	// users
	users := app.Group("/users")
	users.Get("/:id", r.Handlers.User.FindByID)

	// categories
	categories := app.Group("/categories")
	categories.Get("/", r.Handlers.Category.FindAllCategories)
	categories.Post("/", r.Handlers.Category.CreateCategory)
	categories.Get("/:id", r.Handlers.Category.FindCategoryByID)
	categories.Patch("/:id", r.Handlers.Category.UpdateCategory)
	categories.Delete("/:id", r.Handlers.Category.DeleteCategory)
}
