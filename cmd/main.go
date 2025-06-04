package main

import (
	"github.com/GalahadKingsman/messenger_users/internal/app"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"log"
)

func main() {
	cfg := config.GetConfig()
	if err := app.Run(cfg); err != nil {
		log.Fatalf("Application failed: %v", err)
	}
}
