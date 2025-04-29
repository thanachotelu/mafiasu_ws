package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	//"time"
	"net/url"
	"strings"
	"io"
	"mafiasu_ws/config"
	"mafiasu_ws/external/keycloak"
	//"mafiasu_ws/internal/interfaces" // ตรวจสอบการ import ที่ถูกต้อง
)

// keycloakService implements interfaces.KeycloakService
type keycloakService struct {
	httpClient *http.Client
	cfg        config.KeycloakConfig
}

// NewKeycloakService constructs a KeycloakService
//func NewKeycloakService(cfg config.KeycloakConfig) interfaces.KeycloakService {
//	return &keycloakService{
	//	httpClient: &http.Client{Timeout: 10 * time.Second},
	//	cfg:        cfg,
//	}
//}


func (k *keycloakService) getAdminToken(ctx context.Context) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", k.cfg.AdminClientID)
	data.Set("username", k.cfg.AdminUser)
	data.Set("password", k.cfg.AdminPass)

	url := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", k.cfg.BaseURL, k.cfg.Realm)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, io.NopCloser(strings.NewReader(data.Encode())))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := k.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token endpoint returned %d", resp.StatusCode)
	}

	var body struct{ AccessToken string `json:"access_token"` }
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", err
	}
	return body.AccessToken, nil
}

// CreateUser sends a CreateUserRequest to Keycloak to register a new user
func (k *keycloakService) CreateUser(ctx context.Context, req keycloak.CreateUserRequest) error {
    token, err := k.getAdminToken(ctx)
    if err != nil {
        return fmt.Errorf("cannot fetch admin token: %w", err)
    }

    // Marshal the request into JSON to send to Keycloak
    bodyBytes, err := json.Marshal(req)
    if err != nil {
        return err
    }

    url := fmt.Sprintf("%s/admin/realms/%s/users", k.cfg.BaseURL, k.cfg.Realm)
    httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
    if err != nil {
        return err
    }
    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Authorization", "Bearer "+token)

    resp, err := k.httpClient.Do(httpReq)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return fmt.Errorf("keycloak create-user returned %d", resp.StatusCode)
    }

    return nil
}


// AssignRole grants a realm role to an existing Keycloak user
func (k *keycloakService) AssignRole(ctx context.Context, userID, roleName string) error {
	token, err := k.getAdminToken(ctx)
	if err != nil {
		return fmt.Errorf("cannot fetch admin token: %w", err)
	}

	// Create the payload for role assignment
	rolePayload := []keycloak.RoleRepresentation{{Name: roleName}}
	bodyBytes, err := json.Marshal(rolePayload)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/realm", k.cfg.BaseURL, k.cfg.Realm, userID)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+token)

	resp, err := k.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("keycloak assign-role returned %d", resp.StatusCode)
	}
	return nil
}
