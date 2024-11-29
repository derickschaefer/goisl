package pkg

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

// URLHook defines a function signature for custom URL processing
type URLHook func(parsedURL *url.URL) (*url.URL, error)

// AllowedProtocols defines the list of acceptable URL schemes
var AllowedProtocols = []string{"http", "https", "mailto", "ftp"}

// EscapeURL sanitizes and escapes a URL, applying an optional custom hook.
func EscapeURL(input string, context string, hook URLHook) (string, error) {
	if input == "" {
		return "", nil
	}

	// Replace leading spaces with %20
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, " ", "%20")

	// Remove invalid characters
	re := regexp.MustCompile(`[^a-zA-Z0-9-~+_.?#=!&;,/:%@$\|*'()$begin:math:display$$end:math:display$\\x80-\\xff]`)
	input = re.ReplaceAllString(input, "")

	if input == "" {
		return "", nil
	}

	// Deep replace dangerous sequences
	dangerous := []string{"%0d", "%0a", "%0D", "%0A"}
	for _, seq := range dangerous {
		input = strings.ReplaceAll(input, seq, "")
	}

	// Replace double slashes in schemes
	input = strings.Replace(input, "://", "://", 1)

	// Ensure the URL has a valid scheme
	parsed, err := url.Parse(input)
	if err != nil {
		return "", errors.New("invalid URL")
	}

	// If no scheme is present, assume http://
	if parsed.Scheme == "" {
		if strings.HasPrefix(input, "/") || strings.HasPrefix(input, "#") || strings.HasPrefix(input, "?") {
			// Relative URL
			return input, nil
		}
		input = "http://" + input
		parsed, _ = url.Parse(input)
	}

	// Validate the scheme against allowed protocols
	if !isAllowedProtocol(parsed.Scheme) {
		return "", errors.New("disallowed protocol")
	}

	// Apply the custom hook, if provided
	if hook != nil {
		parsed, err = hook(parsed)
		if err != nil {
			return "", err
		}
	}

	// Context-based transformations
	finalURL := parsed.String()
	if context == "display" {
		finalURL = strings.ReplaceAll(finalURL, "&", "&#038;")
		finalURL = strings.ReplaceAll(finalURL, "'", "&#039;")
	}

	// Encode brackets
	finalURL = strings.ReplaceAll(finalURL, "[", "%5B")
	finalURL = strings.ReplaceAll(finalURL, "]", "%5D")

	return finalURL, nil
}

// isAllowedProtocol checks if a URL scheme is allowed
func isAllowedProtocol(scheme string) bool {
	for _, protocol := range AllowedProtocols {
		if strings.EqualFold(scheme, protocol) {
			return true
		}
	}
	return false
}
