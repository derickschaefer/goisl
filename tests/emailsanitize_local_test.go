package tests

import (
	"testing"

	"github.com/derickschaefer/goisl"
)

func TestSanitizeLocalPart_UnicodeGarbage(t *testing.T) {
	_, err := isl.SanitizeEmail("カタカナ@example.com", nil)
	if err == nil {
		t.Errorf("Expected error for unicode characters in local part, got none")
	}
}

func TestSanitizeLocalPart(t *testing.T) {
	tests := []struct {
		input     string
		expected  string
		shouldErr bool
	}{
		// ✅ Valid unquoted
		{"john.doe", "john.doe", false},

		// ✅ Collapses multiple dots
		{"john..doe", "john.doe", false},

		// ✅ Trims dots
		{".johndoe.", "johndoe", false},

		// ❌ Invalid after sanitization
		{"...", "", true},

		// ✅ Valid quoted input
		{`"user.name"`, "user.name", false},

		// ❌ Invalid quoted (invalid chars inside)
		{`"user<name>"`, "", true},
	}

	for _, test := range tests {
		full := test.input + "@example.com"
		result, err := isl.SanitizeEmail(full, nil)

		if test.shouldErr {
			if err == nil {
				t.Errorf("Expected error for input '%s', got result: %s", test.input, result)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect error for input '%s', got: %v", test.input, err)
			}
		}
	}
}
