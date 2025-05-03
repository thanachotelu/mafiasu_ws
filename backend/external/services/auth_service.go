package services

import (
	"context"
	"mafiasu_ws/external/interfaces"
	"mafiasu_ws/external/models"
)

type AuthService struct {
    keycloakService interfaces.KeycloakService
}

func NewAuthService(ks interfaces.KeycloakService) *AuthService {
    return &AuthService{keycloakService: ks}
}

func (s *AuthService) RegisterToKeycloak(ctx context.Context, username, password, email, firstname, lastname string) error {
    req := models.CreateUserRequest{
        Username:  username,
        Email:     email,
        FirstName: firstname,
        LastName:  lastname,
        Credentials: []models.Credential{
            {
                Type:      "password",
                Value:     password,
                Temporary: false,
            },
        },
    }

    userID, err := s.keycloakService.CreateUser(ctx, req)
    if err != nil {
        return err
    }

    // Assign default role เช่น "user"
    if err := s.keycloakService.AssignRole(ctx, userID, "user"); err != nil {
        return err
    }

    return nil
}


func (s *AuthService) AssignDefaultRole(ctx context.Context, userID string) error {
    defaultRoleID := "YOUR_ROLE_ID" // หา role id จริง ๆ จาก keycloak หรือเก็บใน config
    // ส่ง context เข้าไป
    return s.keycloakService.AssignRole(ctx, userID, defaultRoleID)
}
