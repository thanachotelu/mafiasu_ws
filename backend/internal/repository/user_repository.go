package repository

import (
    
	"context"
	"mafiasu_ws/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{db}
}

// CreateUser inserts a new user into the users table
// repository/user_repository.go

func (r *userRepository) AddUser(ctx context.Context, user *models.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users (username, password_hash, firstname, lastname, phonenumber, email, role, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.Username, user.PasswordHash, user.Firstname, user.Lastname, user.Phonenumber, user.Email, user.Role, user.Status)
	return err
}



// GetUserByID retrieves a user by their UUID
func (r *userRepository) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, `
		SELECT user_id, username, firstname, lastname, phonenumber, email, role, status
		FROM users WHERE user_id = $1
	`, id).Scan(
		&user.UserID,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Phonenumber,
		&user.Email,
		&user.Role,
		&user.Status,
	)
	return user, err
}
