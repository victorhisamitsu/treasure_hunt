package main

import (
	"context"

	"github.com/Hitsa/CacaTesouro/internal/api"
	"github.com/Hitsa/CacaTesouro/internal/repository"
	"github.com/Hitsa/CacaTesouro/internal/service"
)

func main() {
	// Conectar com o DB
	db := repository.ConexaoDb()
	repository.CreateTables(context.Background(), db)

	// Fornecer o acesso ao banco de dados para a Repository
	caminhoRepository := repository.NewRepositoryCaminho(db)

	pistaRepository := repository.NewRepositoryPista(db)

	userRepository := repository.NewRepositoryUser(db)

	// Forcener a repository para a service
	caminhoService := service.NewCaminhoService(caminhoRepository)

	pistaService := service.NewPistaService(pistaRepository)

	userService := service.NewUserService(userRepository)

	// Fornecer a service para a handler

	api.API(caminhoService, userService, pistaService)

}
