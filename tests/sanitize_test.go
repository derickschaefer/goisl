package tests

import (
	"testing"
	"goisl/pkg"
)

func TestSanitizeEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		// Valid email cases
		{"user@example.com", "user@example.com", true},
		{"  user@example.com  ", "user@example.com", true},
		{"UPPERCASE@example.com", "UPPERCASE@example.com", true},

		// Invalid email cases
		{"invalid-email", "", false},
		{"@example.com", "", false},
		{"user@.com", "", false},
		{"", "", false}, // Empty input
	}

	for _, test := range tests {
		result, err := pkg.SanitizeEmail(test.input)
		if test.isValid && err != nil {
			t.Errorf("Input: %v, Expected valid email but got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: %v, Expected error for invalid email, got: %v", test.input, result)
		}
		if result != test.expected {
			t.Errorf("Input: %v, Expected: %v, Got: %v", test.input, test.expected, result)
		}
	}
}
