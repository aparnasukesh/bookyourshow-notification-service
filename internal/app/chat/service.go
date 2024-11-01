package chat

import (
	"encoding/json"
	"log"
)

type service struct {
	repo Repository
}

type Service interface {
	chat(body []byte) (string, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) chat(body []byte) (string, error) {
	revMsg := &ReceiveMessage{}
	if err := json.Unmarshal(body, revMsg); err != nil {
		return "", err
	}
	output := ResponseMsg(revMsg.Message)

	if err := s.repo.AppendMessage(revMsg.UserID, Message{
		Message:         revMsg.Message,
		ResponseMessage: output,
		SentAt:          revMsg.SentAt,
	}); err != nil {
		log.Println(err)
	}

	return output, nil
}
