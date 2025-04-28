package service

import (
	
	"context"
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
	return &userService{repo}
}

func (s *userService) GetUserByID(ctx context.Context, id string) (models.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

// Implement the RegisterUser method
func (s *userService) RegisterUser(ctx context.Context, user *models.User) error {
	err := s.repo.AddUser(ctx, user) // เรียกใช้ AddUser จาก repository
	if err != nil {
		return err
	}
	return nil
}
