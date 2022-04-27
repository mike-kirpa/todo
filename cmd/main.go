package main

import (
	"github.com/mike-kirpa/just-todo"
	"github.com/mike-kirpa/just-todo/pkg/handler"
	"github.com/mike-kirpa/just-todo/pkg/repository"
	"github.com/mike-kirpa/just-todo/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error uccured while runnung http server: %s", err.Error())
	}
}
