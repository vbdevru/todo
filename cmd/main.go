package main

import (
	"log"

	"github.com/spf13/viper"
	"github.com/vbdevru/todo"
	"github.com/vbdevru/todo/pkg/handler"
	"github.com/vbdevru/todo/pkg/repository"
	"github.com/vbdevru/todo/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
    // test
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("errors occured while running https server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}