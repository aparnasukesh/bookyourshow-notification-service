package chat

import (
	pb "github.com/aparnasukesh/inter-communication/notification"
)

type GrpcHandler struct {
	svc Service
	pb.UnimplementedChatServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}
