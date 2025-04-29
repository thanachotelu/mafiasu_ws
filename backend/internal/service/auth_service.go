package service

import (
    "context"
   // "mafiasu_ws/external/keycloak"
    "mafiasu_ws/internal/interfaces"
)

type AuthService struct {
    keycloakService interfaces.KeycloakService
}

func NewAuthService(ks interfaces.KeycloakService) *AuthService {
    return &AuthService{keycloakService: ks}
}

//func (s *AuthService) RegisterToKeycloak(ctx context.Context, username, password, email, firstname, lastname string) error {
   // req := keycloak.CreateUserRequest{
    //    Username:  username,
    //    Email:     email,
    //    FirstName: firstname,
    //    LastName:  lastname,
        // ลบฟิลด์ "Enabled" ออกหรือเพิ่มเข้าไปใน keycloak.CreateUserRequest ถ้าจำเป็น
    //    Credentials: []keycloak.Credential{
      //      {
     //           Type:      "password",
     //           Value:     password,
     //           Temporary: false,
  //          },
     //   },
  //  }
    // ส่ง context เข้าไป
   // return s.keycloakService.CreateUser(ctx, req)
//}

func (s *AuthService) AssignDefaultRole(ctx context.Context, userID string) error {
    defaultRoleID := "YOUR_ROLE_ID" // หา role id จริง ๆ จาก keycloak หรือเก็บใน config
    // ส่ง context เข้าไป
    return s.keycloakService.AssignRole(ctx, userID, defaultRoleID)
}
