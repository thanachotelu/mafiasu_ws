package keycloak

import (
    "mafiasu_ws/config"
    "context"
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
   // "mafiasu_ws/internal/interfaces"
)
type keycloakService struct {
	httpClient *http.Client
	cfg        config.KeycloakConfig
}
type Service struct {
    BaseURL   string
    Realm     string
    AdminToken string
}
//func NewKeycloakService(cfg config.KeycloakConfig) interfaces.KeycloakService {
	//return &keycloakService{
	//	httpClient: &http.Client{Timeout: 10 * time.Second},
		//cfg:        cfg,
	//}
//}
func (s *Service) CreateUser(ctx context.Context, user CreateUserRequest) (string, error) {
    url := fmt.Sprintf("%s/admin/realms/%s/users", s.BaseURL, s.Realm)

    body, _ := json.Marshal(user)
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + s.AdminToken)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusCreated {
        return "", fmt.Errorf("failed to create user: %d", resp.StatusCode)
    }

    location := resp.Header.Get("Location")
    var userID string
    fmt.Sscanf(location, s.BaseURL+"/admin/realms/"+s.Realm+"/users/%s", &userID)

    return userID, nil
}


func (s *Service) getAdminToken() (string, error) {
    data := "grant_type=password&client_id=admin-cli&username=admin&password=admin"
    req, err := http.NewRequest("POST", s.BaseURL+"/realms/master/protocol/openid-connect/token", bytes.NewBufferString(data))
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

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    token, ok := result["access_token"].(string)
    if !ok {
        return "", fmt.Errorf("no access token in response")
    }

    return token, nil
}

func (s *Service) AssignRole(userID string, roleName string) error {
    token, err := s.getAdminToken()
    if err != nil {
        return fmt.Errorf("get token: %w", err)
    }

    roleID, err := s.getRoleIDByName(roleName, token)
    if err != nil {
        return fmt.Errorf("get role id: %w", err)
    }

    url := fmt.Sprintf("%s/admin/realms/%s/users/%s/role-mappings/realm", s.BaseURL, s.Realm, userID)

    role := []RoleRepresentation{
        {
            ID:   roleID,
            Name: roleName,
        },
    }

    body, _ := json.Marshal(role)
    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusNoContent {
        return fmt.Errorf("failed to assign role: %d", resp.StatusCode)
    }

    return nil
}

func (s *Service) getRoleIDByName(roleName, token string) (string, error) {
    url := fmt.Sprintf("%s/admin/realms/%s/roles/%s", s.BaseURL, s.Realm, roleName)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer " + token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to get role: %d", resp.StatusCode)
    }

    var role RoleRepresentation
    if err := json.NewDecoder(resp.Body).Decode(&role); err != nil {
        return "", err
    }

    return role.ID, nil
}
