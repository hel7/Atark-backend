package main

import (
	farmsage "github.com/hel7/Atark-backend"
	"github.com/hel7/Atark-backend/pkg/handlers"
	"log"
)

func main() {
	srv := new(farmsage.Server)
	handlers := new(handlers.Handlers)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
