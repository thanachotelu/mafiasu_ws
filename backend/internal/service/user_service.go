package service

import (
    "context"
  //  "fmt"
   // "mafiasu_ws/external/keycloak"
    "mafiasu_ws/internal/interfaces"
    "mafiasu_ws/internal/models"
)

// userService implements business logic around users
type userService struct {
    repo            interfaces.UserRepository
    keycloakService interfaces.KeycloakService
}

// NewUserService constructs a new UserService
//func NewUserService(repo interfaces.UserRepository, keycloakService interfaces.KeycloakService) interfaces.UserService {
   // return &userService{repo: repo, keycloakService: keycloakService}
//}

// GetUserByID retrieves a user by its ID
func (s *userService) GetUserByID(ctx context.Context, id string) (models.User, error) {
    return s.repo.GetUserByID(ctx, id)
}

// RegisterUser registers a new user in Postgres and Keycloak
//func (s *userService) RegisterUser(ctx context.Context, user *models.User) error {
    // 1) บันทึกลง Postgres ก่อน
   // if err := s.repo.AddUser(ctx, user); err != nil {
    //    return fmt.Errorf("db add user: %w", err)
  //  }

    // 2) สร้าง user ใน Keycloak (รับแค่ error)
  //  err := s.keycloakService.CreateUser(ctx, keycloak.CreateUserRequest{
    //    Username:  user.Username,
    //    Email:     user.Email,
     //   FirstName: user.Firstname,
    //    LastName:  user.Lastname,
     //   Credentials: []keycloak.Credential{
          //  {
          //      Type:     "password",
           //     Value:    user.RawPassword,
           //     Temporary: false,
         //   },
      //  },
  //  })
   // if err != nil {
   //     return fmt.Errorf("keycloak create user: %w", err)
   // }

    // 3) Assign role "user" ให้กับ user ที่เพิ่งสร้าง (ใช้ user.UserID)
   // if err := s.keycloakService.AssignRole(ctx, user.UserID, "user"); err != nil {
   //     return fmt.Errorf("assign role failed: %w", err)
   // }

   // return nil
//}




