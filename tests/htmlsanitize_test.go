package tests

import (
	"testing"

	"github.com/derickschaefer/goisl" // Import the package containing the functions to test
)

func TestHTMLSanitize(t *testing.T) {
	tests := []struct {
		input       string
		allowedHTML map[string][]string
		expected    string
	}{
		{
			input:       "<script>alert('XSS')</script>",
			allowedHTML: isl.AllowedHTML,
			expected:    "alert('XSS')", // Script tags should be removed
		},
		{
			input:       "Hello &amp; welcome!",
			allowedHTML: isl.AllowedHTML,
			expected:    "Hello &amp; welcome!", // Ensure no double-escaping
		},
		{
			input:       "<b>Bold</b> and <i>italic</i>",
			allowedHTML: isl.AllowedHTML,
			expected:    "<b>Bold</b> and italic", // Only <b> is allowed
		},
	}

	for _, test := range tests {
		result := isl.HTMLSanitize(test.input, test.allowedHTML)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}

func TestHTMLSanitizeBasic(t *testing.T) {
	input := "<img src='bad.jpg' onerror='alert(1)'><b>Hello</b>"
	expected := "<img src='bad.jpg' onerror='alert(1)'><b>Hello</b>" // Currently allowed

	result := isl.HTMLSanitizeBasic(input)
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// TODO: Strip unsafe attributes from allowed tags (e.g., onerror in <img>) in future version
}

func TestMustHTMLSanitizeBasic(t *testing.T) {
	input := "<a href='http://x.com'>Click</a>"
	expected := "<a href='http://x.com'>Click</a>"

	result := isl.MustHTMLSanitizeBasic(input)
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
