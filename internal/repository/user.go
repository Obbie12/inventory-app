package repository

import (
	"database/sql"
	"fmt"
	"time"

	"inventory-app/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository handles all database operations for users
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create adds a new user to the database
func (r *UserRepository) Create(user models.User) (models.User, error) {
	// Generate UUID for user ID
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	query := `
		INSERT INTO users (id, username, password, email, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = r.db.Exec(
		query,
		user.ID,
		user.Username,
		string(hashedPassword),
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	// Clear the password before returning
	user.Password = ""
	return user, nil
}

// GetByUsername retrieves a user by username
func (r *UserRepository) GetByUsername(username string) (models.User, error) {
	var user models.User
	query := `
		SELECT id, username, password, email, created_at, updated_at
		FROM users
		WHERE username = ?
	`
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user with username %s not found", username)
		}
		return models.User{}, err
	}

	return user, nil
}
