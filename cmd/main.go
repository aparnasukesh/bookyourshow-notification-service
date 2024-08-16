package main

import (
	"log"

	"github.com/aparnasukesh/notification-svc/config"
	"github.com/aparnasukesh/notification-svc/internal/boot"
	"github.com/aparnasukesh/notification-svc/internal/di"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	handler := di.InitEmailNotification(cfg)
	server, err := boot.NewGrpcServer(cfg, handler)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := server(); err != nil {
		log.Fatal(err)
		return
	}

}
