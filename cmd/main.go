package main

import (
	"github.com/mike-kirpa/just-todo"
	"github.com/mike-kirpa/just-todo/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error uccured while runnung http server: %s", err.Error())
	}
}
