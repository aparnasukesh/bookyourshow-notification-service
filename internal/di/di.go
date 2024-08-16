package di

import (
	"github.com/aparnasukesh/notification-svc/config"
	"github.com/aparnasukesh/notification-svc/internal/app/email"
)

func InitEmailNotification(cfg config.Config) email.GrpcHandler {
	smtpEmail := email.NewSMTPEmail(cfg)
	svc := email.NewService(smtpEmail)
	handler := email.NewGrpcHandler(svc)
	return handler
}
