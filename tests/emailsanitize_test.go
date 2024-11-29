package tests

import (
	"errors"
	"testing"
	"strings"

	"github.com/derickschaefer/goisl/pkg"
)

func TestSanitizeEmail(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"user@example.com", "user@example.com", true},
		{"  user@example.com  ", "user@example.com", true},
		{"UPPERCASE@EXAMPLE.com", "UPPERCASE@EXAMPLE.com", true},
		{"invalid-email", "", false},
		{"user@.com", "", false},
		{"user@example", "", false},
		{"@example.com", "", false},
		{"", "", false},
	}

	for _, test := range tests {
		result, err := pkg.SanitizeEmail(test.input, nil) // No custom hook
		if test.isValid && err != nil {
			t.Errorf("Expected valid email, got error: %v", err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Expected error, got result: %v", result)
		}
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}

func TestSanitizeEmailWithCustomHook(t *testing.T) {
    // Custom hook to remove tags and block specific domains
    customHook := func(local, domain string) (string, string, error) {
        // Remove tags after '+'
        if plusIndex := strings.Index(local, "+"); plusIndex != -1 {
            local = local[:plusIndex]
        }

        // Block specific domains (case-insensitive)
        blockedDomains := []string{"tempmail.com"} // Removed "example.com"
        domainLower := strings.ToLower(domain)    // Normalize domain to lowercase
        for _, blocked := range blockedDomains {
            if domainLower == strings.ToLower(blocked) {
                return "", "", errors.New("blocked domain")
            }
        }

        return local, domain, nil
    }

    tests := []struct {
        input    string
        expected string
        isValid  bool
    }{
        {"user+tag@example.com", "user@example.com", true},   // Tag removed, valid domain
        {"admin+test@tempmail.com", "", false},              // Blocked domain
        {"info@allowed.com", "info@allowed.com", true},      // Valid domain
        {"test@EXAMPLE.com", "test@EXAMPLE.com", true},      // Allowed domain
	{"\"test\"@EXAMPLE.com", "test@EXAMPLE.com", true},
	{"test.user@example.com", "test.user@example.com", true},      // Allowed domain
	{"test..user@EXAMPLE.com", "test.user@EXAMPLE.com", true},

    }

    for _, test := range tests {
        result, err := pkg.SanitizeEmail(test.input, customHook) // With custom hook
        if test.isValid && err != nil {
            t.Errorf("Input: %s, Expected valid email, got error: %v", test.input, err)
        }
        if !test.isValid && err == nil {
            t.Errorf("Input: %s, Expected error, got result: %v", test.input, result)
        }
        if result != test.expected {
            t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
        }
    }
}
