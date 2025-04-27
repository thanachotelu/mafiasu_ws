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

func (s *userService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *userService) AddUser(ctx context.Context, user models.CreateUserRequest) (models.User, error) {
	return s.repo.AddUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, id string, user models.User) (models.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *userService) DeleteUser(ctx context.Context, id string) (models.User, error) {
	return s.repo.DeleteUser(ctx, id)
}
