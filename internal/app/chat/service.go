package chat

import (
	"encoding/json"
)

type service struct {
	repo Repository
}

type Service interface {
	chat(body []byte, done chan struct{}) (string, error)
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) chat(body []byte, done chan struct{}) (string, error) {
	revMsg := &ReceiveMessage{}
	if err := json.Unmarshal(body, revMsg); err != nil {
		return "", err
	}
	output := ResponseMsg(revMsg.Message)

	return output, nil
}
