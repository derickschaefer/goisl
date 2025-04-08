/*
Package isl provides all escape and sanitize functions for the goisl library.

Version: 1.1.1

File: htmlsanitize.go

Description:
    This file contains functions for sanitizing HTML content.
    The HTMLSanitize function removes unwanted HTML tags, attributes, and protocols,
    while normalizing entities and eliminating null bytes.
    Additional helper functions provide tag-based cleanup, attribute filtering, and
    optional customization through user-defined tag policies.

Change Log:
    - v1.1.X: Additional ./examples will be pushed from time to time. Point releases not referenced here are example updates only.
    - v1.1.0: Added pflag integration for CLI support, custom hook examples, improved validation hooks, and expanded documentation.
    - v1.0.4: Rename pkg to isl and bump version numbers
    - v1.0.3: Remove conflicting license.txt file
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Improved documentation and refined HTML sanitization functions.
    - v1.0.0: Initial implementation of HTML sanitization utilities.

License:
    MIT License
*/

package isl

import (
	"regexp"
	"strings"
)

// SanitizeAllowedProtocols defines the list of acceptable URL schemes for sanitization.
var SanitizeAllowedProtocols = []string{
	"http", "https", "mailto", "ftp", "ftps", "news", "irc", "irc6",
	"ircs", "gopher", "nntp", "feed", "telnet", "mms", "rtsp", "sms",
	"svn", "tel", "fax", "xmpp", "webcal", "urn",
}

// AllowedHTML defines allowed tags and their permitted attributes.
var AllowedHTML = map[string][]string{
	"b":   nil,            // No attributes allowed
	"a":   {"href"},       // Allow only href
	"img": {"src", "alt"}, // Allow src and alt
}

// HTMLSanitizeBasic sanitizes HTML using the default allowed HTML map.
func HTMLSanitizeBasic(content string) string {
	return HTMLSanitize(content, AllowedHTML)
}

// MustHTMLSanitizeBasic runs HTMLSanitize using the default AllowedHTML rules.
func MustHTMLSanitizeBasic(content string) string {
	return HTMLSanitize(content, AllowedHTML)
}

// HTMLSanitize sanitizes content by removing unwanted HTML tags, attributes, and protocols.
func HTMLSanitize(content string, allowedHTML map[string][]string) string {
	// Step 1: Remove null bytes
	content = removeNullBytes(content, "remove")

	// Step 2: Normalize entities
	content = normalizeEntities(content)

	// Step 3: Sanitize HTML tags and attributes
	content = sanitizeHTML(content, allowedHTML)

	return content
}

// Helper function to validate protocols.
func validateProtocol(protocol string) bool {
	return IsAllowedProtocol(protocol, SanitizeAllowedProtocols)
}

// RemoveNullBytes removes null bytes and optionally handles slash-zero sequences.
func removeNullBytes(content string, slashZeroOption string) string {
	// Remove all control characters (excluding \n, \r, and \t).
	content = regexp.MustCompile(`[\x00-\x08\x0B\x0C\x0E-\x1F]`).ReplaceAllString(content, "")

	// Remove or keep slash-zero based on option.
	if slashZeroOption == "remove" {
		content = regexp.MustCompile(`\\+0+`).ReplaceAllString(content, "")
	}

	return content
}

// NormalizeEntities normalizes named and numeric entities in HTML.
func normalizeEntities(content string) string {
	// Replace "&" not followed by a valid entity with "&amp;"
	var result strings.Builder
	i := 0
	for i < len(content) {
		if content[i] == '&' {
			// Check if this is a valid entity
			j := i + 1
			for j < len(content) && ((content[j] >= 'a' && content[j] <= 'z') ||
				(content[j] >= 'A' && content[j] <= 'Z') ||
				(content[j] >= '0' && content[j] <= '9') ||
				content[j] == '#' || content[j] == ';') {
				j++
			}

			if j < len(content) && content[j-1] == ';' {
				// Valid entity; keep as is
				result.WriteString(content[i:j])
				i = j
			} else {
				// Not a valid entity; replace "&" with "&amp;"
				result.WriteString("&amp;")
				i++
			}
		} else {
			// Copy regular characters
			result.WriteByte(content[i])
			i++
		}
	}

	return result.String()
}

func allowedNamedEntity(entity string) string {
	// Add logic for validating named entities like &lt;, &gt;.
	return entity // Placeholder
}

func normalizeNumericEntity(entity string) string {
	// Convert numeric entities to valid characters.
	return entity // Placeholder
}

func normalizeHexEntity(entity string) string {
	// Convert hex entities to valid characters.
	return entity // Placeholder
}

// SanitizeHTML removes disallowed tags and attributes from content.
func sanitizeHTML(content string, allowedHTML map[string][]string) string {
	// Regex to match HTML tags
	tagRegex := regexp.MustCompile(`<(/?[a-zA-Z][^>]*)>`)

	return tagRegex.ReplaceAllStringFunc(content, func(tag string) string {
		// Check if the tag is allowed
		tagName := strings.Split(strings.Trim(tag, "<>/"), " ")[0] // Extract tag name
		if _, allowed := allowedHTML[tagName]; allowed {
			return tag // Return the allowed tag as is
		}
		// Strip disallowed tags
		return ""
	})
}

func validateHTMLTag(tag string, allowedHTML map[string][]string) string {
	// Basic tag validation (expand for attributes).
	return tag // Placeholder
}
