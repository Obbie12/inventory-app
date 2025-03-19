package middleware

import (
	"context"
	"net/http"
	"strings"

	"inventory-app/internal/services"
	"inventory-app/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware is a middleware to handle JWT authentication
type AuthMiddleware struct {
	authService *services.AuthService
}

// NewAuthMiddleware creates a new auth middleware
func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// Authenticate verifies the JWT token
func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is empty or doesn't start with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authentication required", nil)
			return
		}

		// Extract the token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		token, err := m.authService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token", err)
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to parse token claims", nil)
			return
		}

		// Add user ID to the request context
		userID := claims["user_id"].(string)
		ctx := context.WithValue(r.Context(), "user_id", userID)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
