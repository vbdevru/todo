package main

import (
	"log"

	"github.com/vbdevru/todo"
	"github.com/vbdevru/todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
    // test
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}