package main

import (
	"log"

	"github.com/PrinceNarteh/pos/internal/config"
	"github.com/PrinceNarteh/pos/internal/db"
	"github.com/PrinceNarteh/pos/internal/handlers"
	"github.com/PrinceNarteh/pos/internal/repositories"
	"github.com/PrinceNarteh/pos/internal/services"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := createServer()

	repo := repositories.NewRepo(db.DB)
	svc := services.NewServices(repo)
	handlers := handlers.NewHandlers(svc)
	api := app.Group("/api")
	NewRoutes(handlers).initRoutes(api)

	log.Fatal(app.Listen(config.Envs.App.Port))
}
