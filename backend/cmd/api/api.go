package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func createServer() *fiber.App {
	version := "v1.0.0"

	app := fiber.New(fiber.Config{
		AppName:      fmt.Sprintf("POS System API - %s", version),
		ServerHeader: "POS-System-Server",
	})

	return app
}
