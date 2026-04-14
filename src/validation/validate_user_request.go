package validation

import (
	"errors"
	"golang-crud/src/models"
	"strings"
)

func ValidateUser(user models.User) error {
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" {
		return errors.New("Name or Email is missing")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("Password is required")
	}
	return nil
}
