package services

import (
	"fmt"
	"time"

	"inventory-app/internal/models"
	"inventory-app/internal/repository"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthService handles authentication operations
type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(loginReq models.LoginRequest) (*models.LoginResponse, error) {
	user, err := s.userRepo.GetByUsername(loginReq.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	// Clear password
	user.Password = ""

	return &models.LoginResponse{
		Token: token,
		User:  user,
	}, nil
}

// RegisterUser creates a new user
func (s *AuthService) RegisterUser(user models.User) (models.User, error) {
	return s.userRepo.Create(user)
}

// GenerateToken creates a new JWT token for a user
func (s *AuthService) generateToken(user models.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token valid for 24 hours

	// Generate signed token
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates a JWT token
func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.jwtSecret), nil
	})
}
