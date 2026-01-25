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
	app := createServer()

	connPool, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer connPool.Close()

	repo := repositories.NewRepo(connPool)
	svc := services.NewServices(repo)
	handlers := handlers.NewHandlers(svc)
	NewRoutes(handlers).initRoutes(app)

	log.Fatal(app.Listen(config.Env.APP.Port))
}
