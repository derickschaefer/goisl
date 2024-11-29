package tests

import (
	"net/url" // Added import for the url package
	"testing"

	"github.com/derickschaefer/goisl/pkg"
)

func TestEscapeURLWithCustomHook(t *testing.T) {
	// Define a custom hook that modifies the URL without introducing additional encoding
	hook := func(parsedURL *url.URL) (*url.URL, error) {
		// Example: Change the path to "/path" without encoding
		parsedURL.Path = "/path"
		return parsedURL, nil
	}

	input := "https://example.com/originalPath"
	expected := "https://example.com/path"

	result, err := pkg.EscapeURL(input, "", hook)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if result != expected {
		t.Errorf("Input: %q, Expected: %q, Got: %q", input, expected, result)
	}
}
