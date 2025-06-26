package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"io"
	"log"
	"mafiasu_ws/config"
	"mafiasu_ws/external/interfaces"
	"mafiasu_ws/external/models"
	"net/http"
	"time"
)

type KeycloakServices struct {
	httpClient *http.Client
	cfg        config.KeycloakConfig
	BaseURL    string
	Realm      string
	AdminToken string
}

func NewKeycloakService(cfg config.KeycloakConfig) interfaces.KeycloakService {
	log.Printf("Keycloak BaseURL: %s, Realm: %s", cfg.BaseURL, cfg.Realm)
	return &KeycloakServices{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		cfg:        cfg,
		BaseURL:    cfg.BaseURL, // ตั้งค่า BaseURL
		Realm:      cfg.Realm,   // ตั้งค่า Realm
	}
}
func (s *KeycloakServices) CreateUser(ctx context.Context, user models.CreateUserRequest) (string, error) {
	url := fmt.Sprintf("%s/admin/realms/%s/users", s.cfg.BaseURL, s.cfg.Realm)

	token, err := s.GetAdminToken()
	if err != nil {
		return "", fmt.Errorf("failed to get admin token: %w", err)
	}

	body, _ := json.Marshal(user)
	log.Printf("Request Body to Keycloak: %s", string(body)) // Log ข้อมูลที่ส่งไปยัง Keycloak

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to Keycloak: %w", err)
	}
	defer resp.Body.Close()

	// Log Response Status และ Body
	log.Printf("Response Status from Keycloak: %d", resp.StatusCode)
	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("Response Body from Keycloak: %s", string(respBody))

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create user: %d", resp.StatusCode)
	}

	location := resp.Header.Get("Location")
	var userID string
	fmt.Sscanf(location, s.cfg.BaseURL+"/admin/realms/"+s.cfg.Realm+"/users/%s", &userID)

	return userID, nil
}

func (s *KeycloakServices) GetAdminToken() (string, error) {
	data := "grant_type=password&client_id=admin-cli&username=admin&password=admin"
	req, err := http.NewRequest("POST", s.cfg.BaseURL+"/realms/"+s.cfg.Realm+"/protocol/openid-connect/token", bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get admin token, status code: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse access token")
	}

	// เพิ่ม Log เพื่อตรวจสอบ Token
	log.Printf("Admin Token retrieved successfully: %s", token)

	return token, nil
}

func (s *KeycloakServices) AssignRole(ctx context.Context, userID string, roleName string) error {
	token, err := s.GetAdminToken()
	if err != nil {
		return fmt.Errorf("get token: %w", err)
	}

	roleID, err := s.GetRoleIDByName(roleName, token)
	if err != nil {
		return fmt.Errorf("get role id: %w", err)
	}

	url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/realm", s.BaseURL, s.Realm, userID)

	role := []models.RoleRepresentation{
		{
			ID:   roleID,
			Name: roleName,
		},
	}

	body, _ := json.Marshal(role)
	log.Printf("Assigning Role: URL=%s, Body=%s", url, string(body)) // Log URL และ Body

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Response Status: %d", resp.StatusCode) // Log Response Status

	if resp.StatusCode != http.StatusNoContent {
		respBody, _ := io.ReadAll(resp.Body)
		log.Printf("Response Body: %s", string(respBody)) // Log Response Body
		return fmt.Errorf("failed to assign role: %d", resp.StatusCode)
	}

	return nil
}

func (s *KeycloakServices) GetRoleIDByName(roleName, token string) (string, error) {
	url := fmt.Sprintf("%s/admin/realms/%s/roles/%s", s.BaseURL, s.Realm, roleName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get role id: %s", resp.Status)
	}

	// รองรับทั้ง Object และ Array
	var role struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&role); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if role.Name == roleName {
		return role.ID, nil
	}

	return "", fmt.Errorf("role not found: %s", roleName)
}
func (s *KeycloakServices) CreateRoleIfNotExists(roleName string) error {
	token, err := s.GetAdminToken()
	if err != nil {
		return fmt.Errorf("failed to get admin token: %w", err)
	}

	// ตรวจสอบว่า Role มีอยู่แล้วหรือไม่
	url := fmt.Sprintf("%s/admin/realms/%s/roles/%s", s.BaseURL, s.Realm, roleName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to check role existence: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("Role '%s' already exists", roleName)
		return nil // Role มีอยู่แล้ว
	}

	// สร้าง Role ใหม่
	url = fmt.Sprintf("%s/admin/realms/%s/roles", s.BaseURL, s.Realm)
	role := map[string]string{
		"name": roleName,
	}
	body, _ := json.Marshal(role)

	req, _ = http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create role: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create role: %s", string(respBody))
	}

	log.Printf("Role '%s' created successfully", roleName)
	return nil
}

func (s *KeycloakServices) Login(ctx context.Context, username, password string) (string, string, error) {
    urlStr := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", s.BaseURL, s.Realm)

    form := url.Values{}
    form.Set("grant_type", "password")
    form.Set("client_id", s.cfg.ClientID)
    form.Set("username", username)
    form.Set("password", password)

    req, err := http.NewRequestWithContext(ctx, "POST", urlStr, bytes.NewBufferString(form.Encode()))
    if err != nil {
        return "", "", fmt.Errorf("failed to create request: %w", err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    log.Printf("Attempting login for user: %s", username)

    resp, err := s.httpClient.Do(req)
    if err != nil {
        return "", "", fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", "", fmt.Errorf("failed to read response body: %w", err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Printf("Login failed. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
        return "", "", fmt.Errorf("authentication failed with status: %d", resp.StatusCode)
    }

    var result struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
    }

    if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&result); err != nil {
        return "", "", fmt.Errorf("failed to decode response: %w", err)
    }

    return result.AccessToken, result.RefreshToken, nil
}
func (s *KeycloakServices) CreateClientIfNotExists(clientID string) error {
    token, err := s.GetAdminToken()
    if err != nil {
        return err
    }
    // เช็คว่ามี client นี้อยู่แล้วหรือยัง
    url := fmt.Sprintf("%s/admin/realms/%s/clients?clientId=%s", s.cfg.BaseURL, s.cfg.Realm, clientID)
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+token)
    resp, err := s.httpClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    var clients []interface{}
    json.NewDecoder(resp.Body).Decode(&clients)
    if len(clients) > 0 {
        return nil // มีแล้ว
    }
    // ถ้ายังไม่มี ให้สร้างใหม่
    createUrl := fmt.Sprintf("%s/admin/realms/%s/clients", s.cfg.BaseURL, s.cfg.Realm)
    body := map[string]interface{}{
        "clientId":                 clientID,
        "enabled":                  true,
        "publicClient":             true,
        "directAccessGrantsEnabled": true,
    }
    b, _ := json.Marshal(body)
    req, _ = http.NewRequest("POST", createUrl, bytes.NewReader(b))
    req.Header.Set("Authorization", "Bearer "+token)
    req.Header.Set("Content-Type", "application/json")
    resp, err = s.httpClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != 201 && resp.StatusCode != 204 {
        return fmt.Errorf("failed to create client: %s", resp.Status)
    }
    return nil
}
func (s *KeycloakServices) RefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
    urlStr := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", s.BaseURL, s.Realm)

    form := url.Values{}
    form.Set("grant_type", "refresh_token")
    form.Set("client_id", s.cfg.ClientID)
    form.Set("refresh_token", refreshToken)

    req, err := http.NewRequestWithContext(ctx, "POST", urlStr, bytes.NewBufferString(form.Encode()))
    if err != nil {
        return "", "", fmt.Errorf("failed to create request: %w", err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    resp, err := s.httpClient.Do(req)
    if err != nil {
        return "", "", fmt.Errorf("failed to send request: %w", err)
    }
    defer resp.Body.Close()

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", "", fmt.Errorf("failed to read response body: %w", err)
    }

    if resp.StatusCode != http.StatusOK {
        log.Printf("Refresh token failed. Status: %d, Body: %s", resp.StatusCode, string(bodyBytes))
        return "", "", fmt.Errorf("refresh failed with status: %d", resp.StatusCode)
    }

    var result struct {
        AccessToken  string `json:"access_token"`
        RefreshToken string `json:"refresh_token"`
    }

    if err := json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&result); err != nil {
        return "", "", fmt.Errorf("failed to decode response: %w", err)
    }

    return result.AccessToken, result.RefreshToken, nil
}
