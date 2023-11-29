package main

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/handlers"
	"github.com/hel7/Atark-backend/pkg/repository"
	"github.com/hel7/Atark-backend/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(farmsage.Server)
	log.Printf("Starting the server on port %s", viper.GetString("server.port"))

	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.SetDefault("server.port", "8000")
	return nil
}
