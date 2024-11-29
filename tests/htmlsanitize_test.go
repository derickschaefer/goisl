package tests

import (
	"testing"

	"github.com/derickschaefer/goisl/pkg" // Import the package containing the functions to test
)

func TestHTMLSanitize(t *testing.T) {
	tests := []struct {
		input       string
		allowedHTML map[string][]string
		expected    string
	}{
		{
			input:       "<script>alert('XSS')</script>",
			allowedHTML: pkg.AllowedHTML,
			expected:    "alert('XSS')", // Script tags should be removed
		},
		{
			input:       "Hello &amp; welcome!",
			allowedHTML: pkg.AllowedHTML,
			expected:    "Hello &amp; welcome!", // Ensure no double-escaping
		},
		{
			input:       "<b>Bold</b> and <i>italic</i>",
			allowedHTML: pkg.AllowedHTML,
			expected:    "<b>Bold</b> and italic", // Only <b> is allowed
		},
	}

	for _, test := range tests {
		result := pkg.HTMLSanitize(test.input, test.allowedHTML)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
