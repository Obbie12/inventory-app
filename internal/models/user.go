package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Password  string    `json:"password,omitempty" validate:"required,min=8"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginRequest represents a login request body
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}