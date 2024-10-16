package chat

import (
	"context"

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

func (h *GrpcHandler) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	chat, err := h.svc.CreateChat(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &pb.CreateChatResponse{
		ChatId:    int32(chat.ID[]),
		StartedAt: chat.StartedAt,
	}
}
