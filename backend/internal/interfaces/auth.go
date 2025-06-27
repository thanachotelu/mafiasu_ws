package interfaces

import (
	"context"
	//"mafiasu_ws/external/keycloak"
)

// AuthRepository handles logging and API key validation
type AuthRepository interface {
	LogRequest(ctx context.Context, clientID *int, userID *string, endpoint, method string) error
	ValidateJWTToken(ctx context.Context, token string) (map[string]interface{}, error)
	ValidateAPIKey(ctx context.Context, apiKey string) (int, error)
}

// KeycloakService defines methods for user management in Keycloak
// Updated CreateUser to return the created userID and error
// AssignRole now takes context and roleName
type KeycloakService interface {
	// CreateUser(ctx context.Context, req keycloak.CreateUserRequest) error
	AssignRole(ctx context.Context, userID, roleName string) error
	Login(ctx context.Context, username, password string) (string, error)
}
