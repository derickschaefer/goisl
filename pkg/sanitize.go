package pkg

import (
	"errors"
	"net/mail"
	"strings"
)

// SanitizeEmail trims whitespace and validates the email format.
func SanitizeEmail(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	addr, err := mail.ParseAddress(trimmed)
	if err != nil {
		return "", errors.New("invalid email address")
	}
	return addr.Address, nil
}
