package tests

import (
	"testing"

	"github.com/derickschaefer/goisl"
)

func TestSanitizeURL(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		// Basic Valid URLs
		{"http://example.com", "http://example.com", true},
		{"  http://example.com/path?query=<script>  ", "http://example.com/path?query=%3Cscript%3E", true},
		{"", "", true}, // Empty input
		{"javascript:alert('XSS')", "", false}, // Invalid protocol

		// Edge Cases
		{"   https://example.com   ", "https://example.com", true}, // Trim spaces
		{"HtTp://ExAmPlE.CoM", "http://example.com", true}, // Normalize scheme

		// Relative URLs
		{"/path/to/resource", "/path/to/resource", true}, // Valid relative URL
		{"?query=param", "?query=param", true}, // Query string without scheme

		// Security
		{"http://<invalid>.com", "", false}, // Invalid characters in domain
		{"https://example.com/%", "", false}, // Incomplete percent encoding
		{"https://example.com:javascript:alert('XSS')", "", false}, // Protocol injection
		{"data:text/plain;base64,SGVsbG8sIFdvcmxkIQ==", "", false}, // Disallowed protocol

		// Special Characters
		{"http://example.com?query=hello%20world", "http://example.com?query=hello%20world", true}, // Encoded query
		{"https://example.com/✓", "https://example.com/%E2%9C%93", true}, // Encode special characters

		// Complex URLs
		{"http://example.com/path#section", "http://example.com/path#section", true}, // URL with fragment
		{"http://example.com:8080", "http://example.com:8080", true}, // URL with port
		{"http://user:pass@example.com", "http://user:pass@example.com", true}, // URL with user info

		// Empty or Null Input
		{"   ", "", true}, // Empty after trimming
	}

	for _, test := range tests {
		result, err := isl.SanitizeURL(test.input)
		if test.isValid && err != nil {
			t.Errorf("Input: %q, Expected valid URL, got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: %q, Expected error, got result: %v", test.input, result)
		}
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}
