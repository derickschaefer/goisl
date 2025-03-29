/*
Package isl provides all escape and sanitize functions for the goisl library.

Version: 1.0.4

File: filesanitize.go

Description:
    This file contains functions for sanitizing file names.
    The SanitizeFileName function cleans a file name by removing unwanted characters,
    handling multiple extensions, preventing directory traversal, normalizing Unicode,
    and enforcing filename length constraints. Custom hooks can be applied for further validation.

Change Log:
    - v1.0.4: Rename pkg to isl and bump version numbers
    - v1.0.3: Remove conflicting license.txt file
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Enhanced documentation and added precompiled regex optimizations.
    - v1.0.0: Initial implementation of file name sanitization utilities.

License:
    MIT License
*/

// filesanitize.go
package isl

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// FileNameHook defines a function signature for custom filename validation or transformation.
// It receives the sanitized filename and can perform additional checks or modifications.
type FileNameHook func(filename string) (string, error)

// Constants for sanitization
const (
	MaxFileNameLength = 255
)

// Precompiled regular expressions for efficiency
var (
	// Removes special characters except for alphanumerics, hyphens, underscores, and dots
	specialCharsRegex = regexp.MustCompile(`[?[\]\/\\=<>:;,'"&$#*()|~` + "`" + `!{}%@‘«»”“]`)
	
	// Replaces any sequence of whitespace characters with a single hyphen
	whitespaceRegex = regexp.MustCompile(`[\s\r\n\t]+`)
	
	// Condenses multiple hyphens into a single hyphen
	multipleHyphensRegex = regexp.MustCompile(`\-{2,}`)
	
	// Removes hyphens immediately preceding a dot
	hyphenBeforeDotRegex = regexp.MustCompile(`-\.`)
	
	// Condenses multiple dots into a single dot
	multipleDotsRegex = regexp.MustCompile(`\.{2,}`)

	// Condenses multiple underscores into a single underscore
	multipleUnderscoresRegex = regexp.MustCompile(`_{2,}`)
	
	// Ensures the sanitized filename contains only allowed characters
	sanitizationRegexMatch = regexp.MustCompile(`^[a-zA-Z0-9\-_\.]+$`)
)

// SanitizeFileName sanitizes a filename by removing unwanted characters, handling multiple extensions,
// preventing directory traversal, normalizing Unicode characters, and enforcing filename length constraints.
// An optional custom hook can be applied for additional validation or transformation.
func SanitizeFileName(input string, hook FileNameHook) (string, error) {
	
	// Step 0: Check for '..' in the original input to prevent directory traversal
	//if strings.Contains(input, "..") {
	//	return "", errors.New("invalid file name: directory traversal detected")
	//}

	// Step 1: Trim leading and trailing spaces
	input = strings.TrimSpace(input)

	// Step 2: Remove accents and normalize Unicode
	input = removeAccents(input)

	// Step 3: Replace '+' with '-' for consistency
	input = strings.ReplaceAll(input, "+", "-")

	// Step 4: Remove special characters
	input = specialCharsRegex.ReplaceAllString(input, "")

	// Step 5: Replace sequences of whitespace with hyphens
	input = whitespaceRegex.ReplaceAllString(input, "-")

	// Step 6: Replace multiple hyphens with a single hyphen
	input = multipleHyphensRegex.ReplaceAllString(input, "-")

	// Step 7: Remove hyphens immediately preceding a dot
	input = hyphenBeforeDotRegex.ReplaceAllString(input, ".")

	// Step 8: Replace sequences of dots with a single dot
	input = multipleDotsRegex.ReplaceAllString(input, ".")

	// Step 9: Condense multiple underscores into a single underscore
	input = multipleUnderscoresRegex.ReplaceAllString(input, "_")

	// Step 10: Trim leading and trailing dots, dashes, and underscores
	input = strings.Trim(input, ".-_")

	// Step 11: Ensure the filename is not too long
	if len(input) > MaxFileNameLength {
		return "", errors.New("file name too long")
	}

	// Step 12: Ensure the filename contains an extension
	if !strings.Contains(input, ".") {
		return "", errors.New("file name must contain an extension")
	}

	// Step 13: Check for '..' in the sanitized filename
	if strings.Contains(input, "..") {
		return "", errors.New("invalid file name: directory traversal detected")
	}

	// Step 14: Apply custom hook if provided
	if hook != nil {
		sanitized, err := hook(input)
		if err != nil {
			return "", err
		}
		input = sanitized
	}

	// Optional: Verify that the sanitized filename matches expected patterns
	if !sanitizationRegexMatch.MatchString(input) {
		return "", errors.New("file name contains invalid characters after sanitization")
	}

	return input, nil
}

// removeAccents removes accents from a string using Unicode normalization.
func removeAccents(input string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Remove all nonspacing marks
	}), norm.NFC)
	output, _, err := transform.String(t, input)
	if err != nil {
		// If transformation fails, return the original input
		return input
	}
	return output
}
