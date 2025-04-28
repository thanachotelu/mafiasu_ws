package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (models.User, error)
	AddUser(ctx context.Context, user *models.User) error
	// GetAllUsers(ctx context.Context) ([]models.User, error)
	// AddUser(ctx context.Context, user models.CreateUserRequest) (models.User, error)
	// UpdateUser(ctx context.Context, id string) (models.User, error)
	// DeleteUser(ctx context.Context, id string) (models.User, error)
}
