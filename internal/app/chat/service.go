package chat

import "context"

type service struct {
	repo Repository
}

type Service interface {
	CreateChat(ctx context.Context, userId int) (*Chat, error)
}

func NewService(repo Repository) Service {
	return service{
		repo: repo,
	}
}

func (s *service) CreateChat(ctx context.Context, userId int) (*Chat, error) {

}
