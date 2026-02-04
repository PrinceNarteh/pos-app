package main

import (
	"fmt"

	"github.com/PrinceNarteh/pos/internal/config"
	"github.com/gofiber/fiber/v2"
)

func createServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("POS System API - %s", config.Envs.App.Version),
		ServerHeader: "POS-System-Server",
		BodyLimit:    config.Envs.File.MaxSize,
	})

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	// 	AllowCredentials: true,
	// }))

	return app
}
