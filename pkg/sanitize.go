package pkg

import (
	"errors"
	"regexp"
	"strings"
)

// EmailHook defines a function signature for custom email sanitization.
type EmailHook func(local, domain string) (string, string, error)

// SanitizeEmail sanitizes an email address with optional hooks for custom behavior.
func SanitizeEmail(input string, hook EmailHook) (string, error) {
	input = strings.TrimSpace(input)

	// Check minimum length
	if len(input) < 6 {
		return "", errors.New("email too short")
	}

	// Ensure presence of '@' after the first character
	atIndex := strings.Index(input, "@")
	if atIndex < 1 {
		return "", errors.New("email must contain '@' after the first character")
	}

	// Split into local and domain parts
	local := input[:atIndex]
	domain := input[atIndex+1:]

	// Sanitize local part
	local = sanitizeLocalPart(local)
	if local == "" {
		return "", errors.New("invalid characters in local part")
	}

	// Sanitize domain part
	domain, err := sanitizeDomainPart(domain)
	if err != nil {
		return "", err
	}

	// Apply custom hook if provided
	if hook != nil {
		local, domain, err = hook(local, domain)
		if err != nil {
			return "", err
		}
	}

	// Reassemble sanitized email
	return local + "@" + domain, nil
}

// sanitizeLocalPart removes invalid characters from the local part of the email
func sanitizeLocalPart(local string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9!#$%&'*+/=?^_` + "`{|}~.-]" + ``)
	return re.ReplaceAllString(local, "")
}

// sanitizeDomainPart sanitizes and validates the domain part of the email
func sanitizeDomainPart(domain string) (string, error) {
	// Remove sequences of periods
	domain = strings.ReplaceAll(domain, "..", "")

	// Trim leading/trailing periods and hyphens
	domain = strings.Trim(domain, ".-")
	if domain == "" {
		return "", errors.New("invalid domain: empty after sanitization")
	}

	// Split into subdomains and validate each
	subdomains := strings.Split(domain, ".")
	if len(subdomains) < 2 {
		return "", errors.New("domain must contain at least two subdomains")
	}

	for i, sub := range subdomains {
		// Remove invalid characters from subdomains
		subdomains[i] = regexp.MustCompile(`[^a-zA-Z0-9-]`).ReplaceAllString(sub, "")
		if subdomains[i] == "" {
			return "", errors.New("invalid subdomain: empty after sanitization")
		}
	}

	return strings.Join(subdomains, "."), nil
}

func SanitizeURL(input string) (string, error) {
    // Call EscapeURL with "display" context to apply standard escaping
    return EscapeURL(input, "display", nil)
}
