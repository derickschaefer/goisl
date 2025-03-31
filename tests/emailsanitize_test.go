package tests

import (
	"errors"
	"strings"
	"testing"

	"github.com/derickschaefer/goisl"
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
		result, err := isl.SanitizeEmail(test.input, nil)
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
	customHook := func(local, domain string) (string, string, error) {
		if plusIndex := strings.Index(local, "+"); plusIndex != -1 {
			local = local[:plusIndex]
		}
		blockedDomains := []string{"tempmail.com"}
		domainLower := strings.ToLower(domain)
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
		{"user+tag@example.com", "user@example.com", true},
		{"admin+test@tempmail.com", "", false},
		{"info@allowed.com", "info@allowed.com", true},
		{"test@EXAMPLE.com", "test@EXAMPLE.com", true},
		{"\"test\"@EXAMPLE.com", "test@EXAMPLE.com", true},
		{"test.user@example.com", "test.user@example.com", true},
		{"test..user@EXAMPLE.com", "test.user@EXAMPLE.com", true},
	}

	for _, test := range tests {
		result, err := isl.SanitizeEmail(test.input, customHook)
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

func TestSanitizeEmailBasic(t *testing.T) {
	input := "  test.user@example.com "
	expected := "test.user@example.com"
	result, err := isl.SanitizeEmailBasic(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestMustSanitizeEmailBasic(t *testing.T) {
	valid := "user@example.com"
	expected := "user@example.com"
	result := isl.MustSanitizeEmailBasic(valid)
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid email, but did not panic")
		}
	}()
	// Should panic
	isl.MustSanitizeEmailBasic("invalid")
}
