package main

import (
	"github.com/GalahadKingsman/messenger_users/internal/app"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"github.com/caarlos0/env"
	"log"
)

func main() {
	config := new(config.Config)
	if err := env.Parse(config); err != nil {
		log.Fatal("can not parce config")
	}

	app.Run(config)
}
