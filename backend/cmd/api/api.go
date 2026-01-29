package main

import (
	"fmt"

	"github.com/PrinceNarteh/pos/internal/config"
	"github.com/gofiber/fiber/v2"
)

func createServer(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("POS System API - %s", cfg.App.Version),
		ServerHeader: "POS-System-Server",
		BodyLimit:    cfg.File.MaxSize,
	})

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	// 	AllowCredentials: true,
	// }))

	return app
}
