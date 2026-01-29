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
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := createServer(cfg)

	repo := repositories.NewRepo(db)
	svc := services.NewServices(repo)
	handlers := handlers.NewHandlers(svc)
	api := app.Group("/api")
	NewRoutes(handlers).initRoutes(api)

	log.Fatal(app.Listen(config.Env.App.Port))
}
