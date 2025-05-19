package main

import (
	"log"

	"github.com/joho/godotenv"

	server "github.com/borisdvlpr/epicshelter/cmd"
	"github.com/borisdvlpr/epicshelter/pkg/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Unable to load env file: %v. Loading default values.", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load application configuration: %v", err)
	}

	server.Run(cfg)
}
