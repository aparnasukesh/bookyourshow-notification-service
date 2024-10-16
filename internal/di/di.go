package di

import (
	"log"

	"github.com/aparnasukesh/notification-svc/config"
	"github.com/aparnasukesh/notification-svc/internal/app/chat"
	"github.com/aparnasukesh/notification-svc/internal/app/email"
	"github.com/aparnasukesh/notification-svc/internal/boot"
	"github.com/aparnasukesh/notification-svc/pkg/mongodb"
)

func InitResources(cfg config.Config) (func() error, error) {
	smtpEmail := email.NewSMTPEmail(cfg)
	svc := email.NewService(smtpEmail)
	emailHandler := email.NewGrpcHandler(svc)

	// Db initialization
	db, err := mongodb.NewMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// // Movie Module Initialization
	chatRepo := chat.NewRepository(db)
	chatService := chat.NewService(chatRepo)
	chatGrpcHandler := chat.NewGrpcHandler(chatService)

	// Server initialization
	server, err := boot.NewGrpcServer(cfg, emailHandler, chatGrpcHandler)
	if err != nil {
		log.Fatal(err)
	}
	return server, nil
}
