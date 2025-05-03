package interfaces

import (
    "context"
    "mafiasu_ws/external/models"
)

type KeycloakService interface {
    CreateUser(ctx context.Context, user models.CreateUserRequest) (string, error)
    GetAdminToken() (string, error)
    AssignRole(ctx context.Context, userID string, roleName string) error
    GetRoleIDByName(roleName, token string) (string, error)
    CreateRoleIfNotExists(roleName string) error
}
