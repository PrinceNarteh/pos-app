package main

import (
	"fmt"

	"github.com/PrinceNarteh/pos/internal/config"
	"github.com/gofiber/fiber/v2"
)

func createServer() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("POS System API - %s", config.Env.App.Version),
		ServerHeader: "POS-System-Server",
		BodyLimit:    config.Env.File.MaxSize,
	})

	return app
}
