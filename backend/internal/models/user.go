package models

import "time"

type User struct {
	UserID       string    `db:"user_id" json:"user_id"`
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
	Firstname    string    `db:"firstname" json:"firstname"`
	Lastname     string    `db:"lastname" json:"lastname"`
	Phonenumber  string    `db:"phonenumber" json:"phonenumber"`
	Email        string    `db:"email" json:"email"`
	Role         string    `db:"role" json:"role"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Status       string    `db:"status" json:"status"`
}

type UserResponse struct {
	UserID      string `json:"user_id"`
	Username    string `json:"username"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      string `json:"status"`
}

type CreateUserRequest struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
}

type LoginRequest struct {
	Username     string `db:"username" binding:"required"`
	PasswordHash string `db:"password_hash" binding:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
