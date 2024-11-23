package boot

import (
	"log"
	"net"

	pb "github.com/aparnasukesh/inter-communication/notification"
	"github.com/aparnasukesh/notification-svc/config"
	"github.com/aparnasukesh/notification-svc/internal/app/chat"
	"github.com/aparnasukesh/notification-svc/internal/app/email"
	"google.golang.org/grpc"
)

func NewGrpcServer(config config.Config, emailHandler email.GrpcHandler) (func() error, error) {
	//lis, err := net.Listen("tcp", ":"+config.GrpcPort)
	lis, err := net.Listen("tcp", "0.0.0.0:"+config.GRPCPORT)

	if err != nil {
		return nil, err
	}
	s := grpc.NewServer()
	pb.RegisterEmailServiceServer(s, &emailHandler)
	srv := func() error {
		log.Printf("gRPC server started on port %s", config.GRPCPORT)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			return err
		}
		return nil
	}
	return srv, nil
}

func NewRabbitMQConsumer(chatConsumer *chat.NewRabbitMqConsumer) error {
	return chatConsumer.ChatConsumer()
}
