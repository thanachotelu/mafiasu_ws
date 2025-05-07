package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	AddUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, id string, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id string) (models.User, error)
}
