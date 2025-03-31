/*
Package isl provides all escape and sanitize functions for the goisl library.

Version: 1.1.0

File: escape.go

Description:
    This file contains functions for escaping and sanitizing content.
    The EscapePlainText function cleans plain text by removing unwanted characters
    and allows customization via an optional hook. SafeEscapeHTML escapes specific
    HTML characters without affecting '%' for use in URL contexts. Additionally,
    EscapeURL sanitizes URLs by normalizing schemes, hosts, query parameters, and
    applying optional custom hooks for further transformations.

Change Log:
    - v1.1.0: Added pflag integration for CLI support, custom hook examples, improved validation hooks, and expanded documentation.
    - v1.0.4: Rename pkg to isl and bump version numbers
    - v1.0.3: Remove conflicting license.txt file
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Improved documentation and refined escaping functions.
    - v1.0.0: Initial implementation of escaping utilities.

License:
    MIT License
*/

package isl

import (
    "errors"
    "net/url"
    "path"
    "regexp"
    "strings"
)

// URLHook defines a function signature for custom URL processing.
type URLHook func(parsedURL *url.URL) (*url.URL, error)

// EscapeAllowedProtocols defines the list of acceptable URL schemes for escaping.
var EscapeAllowedProtocols = []string{"http", "https", "mailto", "ftp"}

// EscapePlainTextHook defines a function signature for custom behavior.
type EscapePlainTextHook func() []rune

// EscapePlainText sanitizes plain text by removing unwanted characters.
// It allows customization through an optional hook to permit additional characters.
func EscapePlainText(input string, hook EscapePlainTextHook) string {
	// Trim leading and trailing whitespace
	input = strings.TrimSpace(input)

	// Replace sequences of whitespace with a single space
	whitespaceRegex := regexp.MustCompile(`\s+`)
	input = whitespaceRegex.ReplaceAllString(input, " ")

	// Base regex to allow alphanumeric characters and spaces
	allowedChars := "a-zA-Z0-9\\s"

	// Add custom characters from the hook
	if hook != nil {
		additionalChars := string(hook()) // Convert rune slice to string
		allowedChars += regexp.QuoteMeta(additionalChars) // Escape any special regex characters
	}

	// Regex to identify disallowed characters
	disallowedRegex := regexp.MustCompile(`[^` + allowedChars + `]`)

	// Replace disallowed characters with an empty string
	input = disallowedRegex.ReplaceAllString(input, "")

	// Ensure spaces between words are not removed
	input = whitespaceRegex.ReplaceAllString(input, " ")

	// Trim any residual spaces
	input = strings.TrimSpace(input)

	return input
}

// SafeEscapeHTML escapes only specific characters, excluding '%'.
func SafeEscapeHTML(input string) string {
    replacer := strings.NewReplacer(
        "&", "&amp;",
        "<", "&lt;",
        ">", "&gt;",
        "\"", "&quot;",
        "'", "&#39;",
    )
    return replacer.Replace(input)
}

// EscapeURL sanitizes and escapes a URL, applying an optional custom hook.
func EscapeURL(input string, context string, hook URLHook) (string, error) {
    // Step 1: Trim and replace spaces with %20
    input = strings.TrimSpace(input)
    input = strings.ReplaceAll(input, " ", "%20")

    // Early return if input is empty after trimming
    if input == "" {
        return "", nil
    }

    // Step 3: Parse the URL
    parsed, err := url.Parse(input)
    if err != nil {
        return "", errors.New("invalid URL")
    }

    // Step 4: Add a scheme if missing
    if parsed.Scheme == "" {
        if strings.HasPrefix(input, "/") || strings.HasPrefix(input, "#") || strings.HasPrefix(input, "?") {
            // Relative URL
            return input, nil
        }
        input = "http://" + input
        parsed, err = url.Parse(input)
        if err != nil {
            return "", errors.New("invalid URL after adding scheme")
        }
    }

    // Step 5: Normalize scheme and host
    parsed.Scheme = strings.ToLower(parsed.Scheme)
    parsed.Host = strings.ToLower(parsed.Host)

    // Step 6: Validate the domain without Punycode
    domainRegex := regexp.MustCompile(`^[a-zA-Z0-9.-]+(:[0-9]+)?$`)
    if !domainRegex.MatchString(parsed.Host) {
        return "", errors.New("invalid characters in domain")
    }

    // Step 7: Apply the custom hook, if provided
    if hook != nil {
        parsed, err = hook(parsed)
        if err != nil {
            return "", err
        }
    }

    // Step 8: Sanitize query parameters
    if err := sanitizeQueryParameters(parsed); err != nil {
        return "", err
    }

    // Step 9: Clean the path
    parsed.Path = path.Clean(parsed.Path)
    if parsed.Path == "." || parsed.Path == "/." {
        parsed.Path = ""
    }

    // Step 10: Escape fragment if present
    if parsed.Fragment != "" {
        parsed.Fragment = url.PathEscape(parsed.Fragment)
    }

    // Step 11: Context-based transformations
    finalURL := parsed.String()
    if context == "display" {
        // Use SafeEscapeHTML to avoid escaping '%'
        finalURL = SafeEscapeHTML(finalURL)
    }

    // Step 12: Encode brackets
    finalURL = strings.ReplaceAll(finalURL, "[", "%5B")
    finalURL = strings.ReplaceAll(finalURL, "]", "%5D")

    return finalURL, nil
}

// sanitizeQueryParameters sanitizes query parameters in the URL.
func sanitizeQueryParameters(parsed *url.URL) error {
    if len(parsed.RawQuery) > 0 {
        query := parsed.Query()        // Parse query into a map
        rawQuery := make([]string, 0) // To store sanitized key-value pairs
        for key, values := range query {
            for _, value := range values {
                // Decode existing encoded values to avoid double encoding
                decodedValue, err := url.QueryUnescape(value)
                if err != nil {
                    return errors.New("failed to decode query parameter")
                }

                // Properly escape the value using url.QueryEscape
                escapedValue := url.QueryEscape(decodedValue)

                // Replace '+' with '%20' to match test expectations
                escapedValue = strings.ReplaceAll(escapedValue, "+", "%20")

                // Add sanitized key-value pair to raw query
                rawQuery = append(rawQuery, key+"="+escapedValue)
            }
        }
        // Manually construct the RawQuery
        parsed.RawQuery = strings.Join(rawQuery, "&")
    }
    return nil
}
