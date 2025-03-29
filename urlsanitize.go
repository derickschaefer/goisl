/*
Package isl provides all escape and sanitize functions for the goisl library.

Version: 1.0.3

File: urlsanitize.go

Description:
    This file provides a simple wrapper for URL sanitization using the "display" context.
    The SanitizeURL function calls EscapeURL with a "display" context to apply standard URL escaping.

Change Log:
    - v1.0.3: Rename pkg to isl and bump version numbers
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Added documentation header and refined SanitizeURL.
    - v1.0.0: Initial implementation of URL sanitization wrapper.

License:
    MIT License
*/

package isl

// SanitizeURL sanitizes the input URL using the "display" context.
func SanitizeURL(input string) (string, error) {
	// Call EscapeURL with "display" context to apply standard escaping
	return EscapeURL(input, "display", nil)
}
