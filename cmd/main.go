package main

import (
	"log"

	"github.com/vbdevru/todo"
	"github.com/vbdevru/todo/pkg/handler"
	"github.com/vbdevru/todo/pkg/repository"
	"github.com/vbdevru/todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
    // test
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running https server: %s", err.Error())
	}
}