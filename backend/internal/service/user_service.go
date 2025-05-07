package service

import (
	"context"
	"fmt"
	"log"
	extInterfaces "mafiasu_ws/external/interfaces" // ตั้งชื่อ Alias ให้ external/interfaces
	exModels "mafiasu_ws/external/models"
	intInterfaces "mafiasu_ws/internal/interfaces" // ตั้งชื่อ Alias ให้ internal/interfaces
	intModels "mafiasu_ws/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// userService implements business logic around users
type userService struct {
	repo            intInterfaces.UserRepository
	keycloakService extInterfaces.KeycloakService
}

// NewUserService constructs a new UserService
func NewUserService(repo intInterfaces.UserRepository, keycloakService extInterfaces.KeycloakService) intInterfaces.UserService {
	return &userService{repo: repo, keycloakService: keycloakService}
}

// GetUserByID retrieves a user by its ID
func (s *userService) GetUserByID(ctx context.Context, id string) (intModels.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]intModels.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *userService) AddUser(ctx context.Context, user intModels.CreateUserRequest) (intModels.User, error) {
	log.Println("Starting AddUser process...")

	// ตรวจสอบ Role ที่ส่งมาจาก Client
	validRoles := map[string]bool{
		"user":       true,
		"admin":      true,
		"Affiliator": true,
	}

	if !validRoles[user.Role] {
		return intModels.User{}, fmt.Errorf("invalid role: %s", user.Role)
	}

	// Log ข้อมูลที่กำลังจะส่งไปยัง Keycloak
	log.Printf("Sending to Keycloak: Username=%s, Email=%s, FirstName=%s, LastName=%s, Role=%s",
		user.Username, user.Email, user.Firstname, user.Lastname, user.Role)

	// สร้างผู้ใช้ใน Keycloak
	userID, err := s.keycloakService.CreateUser(ctx, exModels.CreateUserRequest{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Enabled:   true,
		Credentials: []exModels.Credential{
			{
				Type:      "password",
				Value:     user.Password,
				Temporary: false,
			},
		},
	})
	if err != nil {
		log.Printf("Error creating user in Keycloak: %v", err)
		return intModels.User{}, fmt.Errorf("failed to create user in Keycloak: %w", err)
	}

	log.Printf("User created in Keycloak with ID: %s", userID)

	// แฮชรหัสผ่านก่อนเก็บในฐานข้อมูล
	log.Println("Hashing password...")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return intModels.User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	// Prepare user data for database
	userToSave := intModels.User{
		Username:     user.Username,
		PasswordHash: string(hashedPassword),
		Firstname:    user.Firstname,
		Lastname:     user.Lastname,
		Email:        user.Email,
		Phonenumber:  user.Phonenumber,
		Role:         user.Role,
		Status:       "active", // Add default status if needed
	}

	// Add user to database
	log.Printf("Adding user to database with data: %+v", userToSave)
	userInDB, err := s.repo.AddUser(ctx, userToSave)
	if err != nil {
		log.Printf("Error adding user to database: %v", err)
		return intModels.User{}, fmt.Errorf("failed to add user in database: %w", err)
	}

	// Assign Role in Keycloak
	if err := s.keycloakService.AssignRole(ctx, userID, user.Role); err != nil {
		return intModels.User{}, fmt.Errorf("failed to assign role in Keycloak: %w", err)
	}

	return userInDB, nil
}

func (s *userService) UpdateUser(ctx context.Context, id string, user intModels.User) (intModels.User, error) {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *userService) DeleteUser(ctx context.Context, id string) (intModels.User, error) {
	return s.repo.DeleteUser(ctx, id)
}
