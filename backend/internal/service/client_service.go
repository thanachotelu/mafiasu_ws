package service

import (
	"context"
	"mafiasu_ws/internal/interfaces"
)

type clientService struct {
	repo interfaces.ClientRepository
}

func NewClientService(repo interfaces.ClientRepository) interfaces.ClientService {
	return &clientService{repo: repo}
}

func (s *clientService) CreateClient(ctx context.Context, name, email string) (string, error) {
	return s.repo.CreateClient(ctx, name, email)
}

func (s *clientService) RevokeClient(ctx context.Context, apiKey string) error {
	return s.repo.RevokeClient(ctx, apiKey)
}
