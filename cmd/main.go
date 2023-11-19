package main

import (
	"log"

	farmsage "github.com/hel7/Atark-backend"
)

func main() {
	srv := new(farmsage.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error running server", err.Error())
	}
}
