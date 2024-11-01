package di

import (
	"log"

	"github.com/aparnasukesh/notification-svc/config"
	"github.com/aparnasukesh/notification-svc/internal/app/chat"
	"github.com/aparnasukesh/notification-svc/internal/app/email"
	"github.com/aparnasukesh/notification-svc/internal/boot"
	"github.com/aparnasukesh/notification-svc/pkg/mongodb"
	"github.com/aparnasukesh/notification-svc/pkg/rabbitmq"
)

func InitResources(cfg config.Config) (func() error, error) {
	smtpEmail := email.NewSMTPEmail(cfg)
	svc := email.NewService(smtpEmail)
	emailHandler := email.NewGrpcHandler(svc)

	// // Db initialization
	db, err := mongodb.NewMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := rabbitmq.NewRabbitMQConnection()
	if err != nil {
		log.Fatal(err)
	}
	// // Movie Module Initialization
	repo := chat.NewRepository(db)
	chatService := chat.NewService(repo)
	chatConsumer := chat.NewRabbitMQConsumer(chatService, conn)

	// Server initialization
	server, err := boot.NewGrpcServer(cfg, emailHandler)
	if err != nil {
		log.Fatal(err)
	}
	err = boot.NewRabbitMQConsumer(chatConsumer)
	if err != nil {
		log.Fatal(err)
	}
	return server, nil
}
