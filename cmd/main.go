package main

import (
	"log"

	"github.com/aparnasukesh/notification-svc/config"
	"github.com/aparnasukesh/notification-svc/internal/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	server, err := di.InitResources(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := server(); err != nil {
		log.Fatal(err)
	}
}
