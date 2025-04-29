package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
	
)


type keycloakService interface {
	CreateUser(ctx context.Context, user models.CreateUserRequest) (string, error)
	getAdminToken() (string, error)
	AssignRole(userID string, roleName string) error
	getRoleIDByName(roleName, token string) (string, error)
}
