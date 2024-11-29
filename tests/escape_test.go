package tests

import (
	"errors"
	"net/url"
	"testing"

	"github.com/derickschaefer/goisl/pkg"
)

func TestEscapeURLWithCustomHook(t *testing.T) {
	customHook := func(parsedURL *url.URL) (*url.URL, error) {
		// Force HTTPS
		if parsedURL.Scheme == "http" {
			parsedURL.Scheme = "https"
		}

		// Reject URLs with tracking_id query parameter
		query := parsedURL.Query()
		if _, exists := query["tracking_id"]; exists {
			return nil, errors.New("tracking_id is not allowed")
		}

		return parsedURL, nil
	}

	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"http://example.com", "https://example.com", true},
		{"http://example.com?tracking_id=12345", "", false}, // Blocked due to tracking_id
		{"https://example.com/path", "https://example.com/path", true},
	}

	for _, test := range tests {
		result, err := pkg.EscapeURL(test.input, "display", customHook)
		if test.isValid && err != nil {
			t.Errorf("Expected valid URL, got error: %v", err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Expected error, got result: %v", result)
		}
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
