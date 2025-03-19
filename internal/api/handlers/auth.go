package handlers

import (
	"encoding/json"
	"net/http"

	"inventory-app/internal/models"
	"inventory-app/internal/services"
	"inventory-app/internal/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *services.AuthService
	validator   *utils.Validator
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validator:   utils.NewValidator(),
	}
}

// Login handles user login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Validate the login request
	if err := h.validator.Validate(loginReq); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation error", err)
		return
	}

	// Attempt to login
	response, err := h.authService.Login(loginReq)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Register handles user registration
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Validate the user data
	if err := h.validator.Validate(user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation error", err)
		return
	}

	// Register the user
	createdUser, err := h.authService.RegisterUser(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdUser)
}
