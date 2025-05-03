package repository

import (
    "log"
	"context"
	"fmt"
	"mafiasu_ws/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *userRepository {
	return &userRepository{db: db}
}

// CreateUser inserts a new user into the users table
// repository/user_repository.go





// GetUserByID retrieves a user by their UUID
func (r *userRepository) GetUserByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, `
	SELECT user_id, username, firstname, lastname, email, status
	FROM users WHERE user_id = $1
`, id).Scan(
		&user.UserID,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Status,
	)
	err = r.db.QueryRow(ctx, `
	SELECT user_id, username, firstname, lastname, email, status
	FROM users WHERE user_id = $1
`, id).Scan(
		&user.UserID,
		&user.Username,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.Status,
	)
	return user, err
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.Query(ctx, `SELECT user_id, username, firstname, lastname, email, status FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.Username, &u.Firstname, &u.Lastname, &u.Email, &u.Status); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func (r *userRepository) AddUser(ctx context.Context, user models.CreateUserRequest) (models.User, error) {
    query := `
        INSERT INTO users (username, password_hash, firstname, lastname, phonenumber, email, role, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING user_id, created_at, updated_at
    `
    log.Printf("Executing query: %s", query)
    log.Printf("With values: Username=%s, PasswordHash=%s, Firstname=%s, Lastname=%s, Phonenumber=%s, Email=%s",
        user.Username, user.Password, user.Firstname, user.Lastname, user.Phonenumber, user.Email)

    var newUser models.User
    err := r.db.QueryRow(ctx, query,
        user.Username,
        user.Password, // ต้องเป็นรหัสผ่านที่ถูกแฮชแล้ว
        user.Firstname,
        user.Lastname,
        user.Phonenumber,
        user.Email,
		user.Role,       // ใช้ Role ที่ส่งมาจาก Client
        "active",     // ค่า status เริ่มต้น
    ).Scan(&newUser.UserID, &newUser.CreatedAt, &newUser.UpdatedAt)
    if err != nil {
        log.Printf("Error executing query: %v", err)
        return models.User{}, fmt.Errorf("failed to insert user: %w", err)
    }
    return newUser, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id string, input models.User) (models.User, error) {
	// ตรวจสอบ user
	var existing models.User
	err := r.db.QueryRow(ctx, `
		SELECT user_id FROM users WHERE user_id = $1
	`, id).Scan(&existing.UserID)
	if err != nil {
		return models.User{}, fmt.Errorf("user not found: %w", err)
	}

	// เช็ก email ซ้ำ
	var count int
	err = r.db.QueryRow(ctx, `
		SELECT COUNT(1) FROM users WHERE email = $1 AND user_id != $2
	`, input.Email, id).Scan(&count)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to check duplicate email: %w", err)
	}
	if count > 0 {
		return models.User{}, fmt.Errorf("email is already in use")
	}

	// อัปเดตข้อมูล
	var updated models.User
	err = r.db.QueryRow(ctx, `
		UPDATE users
		SET username = $1,
			password_hash = $2,
			firstname = $3,
			lastname = $4,
			phonenumber = $5,
			email = $6,
			role = $7,
			status = $8,
			updated_at = NOW()
		WHERE user_id = $9
		RETURNING user_id, username, password_hash, firstname, lastname, phonenumber, email, role, created_at, updated_at, status
	`,
		input.Username,
		input.PasswordHash,
		input.Firstname,
		input.Lastname,
		input.Phonenumber,
		input.Email,
		input.Role,
		input.Status,
		id,
	).Scan(
		&updated.UserID,
		&updated.Username,
		&updated.PasswordHash,
		&updated.Firstname,
		&updated.Lastname,
		&updated.Phonenumber,
		&updated.Email,
		&updated.Role,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.Status,
	)

	return updated, err
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) (models.User, error) {
	var deleted models.User
	err := r.db.QueryRow(ctx, `
		DELETE FROM users
		WHERE user_id = $1
		RETURNING user_id, username, password_hash, firstname, lastname, phonenumber, email, role, created_at, updated_at, status
	`, id).Scan(
		&deleted.UserID,
		&deleted.Username,
		&deleted.PasswordHash,
		&deleted.Firstname,
		&deleted.Lastname,
		&deleted.Phonenumber,
		&deleted.Email,
		&deleted.Role,
		&deleted.CreatedAt,
		&deleted.UpdatedAt,
		&deleted.Status,
	)

	if err != nil {
		return models.User{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return deleted, nil
}
