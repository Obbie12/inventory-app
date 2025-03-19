package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Validator is a wrapper around the validator package
type Validator struct {
	validator *validator.Validate
}

// NewValidator creates a new validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validate validates a struct based on validate tags
func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)
	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok && len(validationErrors) > 0 {
			firstError := validationErrors[0]
			return fmt.Errorf("validation error on field '%s': %s", firstError.Field(), getValidationErrorMsg(firstError))
		}
		return err
	}
	return nil
}

// getValidationErrorMsg returns a user-friendly error message for validation errors
func getValidationErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email format"
	case "min":
		return fmt.Sprintf("should be at least %s characters", e.Param())
	case "gte":
		return fmt.Sprintf("should be greater than or equal to %s", e.Param())
	default:
		return fmt.Sprintf("failed on '%s' validation", e.Tag())
	}
}