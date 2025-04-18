/*
Package isl provides all escape and sanitize functions for the goisl library.

Version: 1.1.2

File: urlsanitize.go

Description:
    This file provides a convenience wrapper around EscapeURL for common URL
    sanitization tasks using the "display" context. The SanitizeURL function applies
    standard URL cleaning and normalization for safe output in user interfaces.

Change Log:
    - v1.1.2: Improved unit testing coverage and fmt fixes.
    - v1.1.X: Additional ./examples will be pushed from time to time. Point releases not referenced here are example updates only.
    - v1.1.0: Added pflag integration for CLI support, custom hook examples, improved validation hooks, and expanded documentation.
    - v1.0.4: Rename pkg to isl and bump version numbers
    - v1.0.3: Remove conflicting license.txt file
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Added documentation header and refined SanitizeURL.
    - v1.0.0: Initial implementation of URL sanitization wrapper.

License:
    MIT License
*/

package isl

import (
	"fmt"
)

// SanitizeURLBasic sanitizes the URL using the default display context and no custom hook.
func SanitizeURLBasic(input string) (string, error) {
	return EscapeURL(input, "display", nil)
}

// MustSanitizeURLBasic is a fail-fast wrapper that panics if SanitizeURLBasic returns an error.
func MustSanitizeURLBasic(input string) string {
	result, err := EscapeURL(input, "display", nil)
	if err != nil {
		panic(fmt.Sprintf("invalid URL input: %v", err))
	}
	return result
}

// SanitizeURL sanitizes the input URL using the "display" context.
func SanitizeURL(input string) (string, error) {
	// Call EscapeURL with "display" context to apply standard escaping
	return EscapeURL(input, "display", nil)
}
