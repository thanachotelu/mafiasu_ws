package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
)

type UserService interface {
	GetUserByID(ctx context.Context, id string) (models.User, error)
	// GetAllUsers(ctx context.Context) ([]models.User, error)
	// AddUser(ctx context.Context, user models.CreateUserRequest) (models.User, error)
	// UpdateUser(ctx context.Context, id string) (models.User, error)
	// DeleteUser(ctx context.Context, id string) (models.User, error)
}
