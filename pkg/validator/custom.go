package validator

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterValidation("lowercase", validateLowercase)
	v.RegisterValidation("strong_password", validateStrongPassword)
	v.RegisterValidation("slug", validateSlug)
}

func validateLowercase(fl validator.FieldLevel) bool {
	return strings.ToLower(fl.Field().String()) == fl.Field().String()
}

func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	
	// At least 8 characters, 1 uppercase, 1 lowercase, 1 number, 1 special char
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)
	
	return len(password) >= 8 && hasUpper && hasLower && hasNumber && hasSpecial
}

func validateSlug(fl validator.FieldLevel) bool {
	slug := fl.Field().String()
	matched, _ := regexp.MatchString(`^[a-z0-9-]+$`, slug)
	return matched
}