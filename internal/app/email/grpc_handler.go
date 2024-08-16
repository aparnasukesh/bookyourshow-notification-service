package email

import (
	"context"

	pb "github.com/aparnasukesh/inter-communication/notification"
)

type GrpcHandler struct {
	svc Service
	pb.UnimplementedEmailServiceServer
}

func NewGrpcHandler(svc Service) GrpcHandler {
	return GrpcHandler{
		svc: svc,
	}
}
func (h *GrpcHandler) SendEmail(ctx context.Context, req *pb.EmailRequest) (*pb.EmailResponse, error) {

	if err := h.svc.SendEmail(req.Otp, req.Email); err != nil {
		return nil, err
	}
	return &pb.EmailResponse{
		Message: "email send successfully",
		Error:   "",
	}, nil
}
