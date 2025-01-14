package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"users/internal/config"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB

// InitDB initializes the database connection
func InitDB(cfg config.Config) *sql.DB {
	var err error

	// PostgreSQL connection string (DSN)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Open database connection
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test database connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	log.Println("Database connected successfully")
	return db
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	return db
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(db *sql.DB, userID int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := db.Exec(query, userID)
	if err != nil {
		log.Printf("Error deleting user with ID %d: %v", userID, err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user found with the given ID")
	}

	log.Printf("User with ID %d deleted successfully", userID)
	return nil
}

// AddUser adds a new user to the database
func AddUser(db *sql.DB, user User) error {
	// Check if email already exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", user.Email).Scan(&exists)
	if err != nil {
		log.Printf("Error checking existing email: %v", err)
		return err
	}

	if exists {
		return errors.New("email already exists")
	}

	// Insert the user into the database
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	_, err = db.Exec(query, user.Name, user.Email)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return nil
}

func UpdateUser(db *sql.DB, user User) error {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" {
		return errors.New("name and email are required")
	}
	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	return nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, name, email FROM users ORDER BY id"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetuserbyID(db *sql.DB, id int) ([]User, error) {
	query := "SELECT id, name, email FROM users where id =$1"

	rows, err := db.Query(query, id)

	if err != nil {
		log.Printf("Error querying users: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
