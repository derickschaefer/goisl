/*
Package isl provides all escape and sanitize functions for the goisl library.

Version: 1.0.3

File: emailsanitize.go

Description:
    This file contains functions for sanitizing email addresses.
    The SanitizeEmail function validates and cleans an email address,
    splitting it into local and domain parts, and optionally applying a
    custom hook for further sanitization.

Change Log:
    - v1.0.3: Rename pkg to isl and bump version numbers
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Enhanced documentation and refined sanitization functions.
    - v1.0.0: Initial implementation of email sanitization utilities.

License:
    MIT License
*/

package isl

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
	var err error
	local, err = sanitizeLocalPart(local)
	if err != nil {
		return "", err
	}

	// Sanitize domain part
	domain, err = sanitizeDomainPart(domain)
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

// sanitizeLocalPart removes invalid characters and validates the local part of the email
func sanitizeLocalPart(local string) (string, error) {
	// Allow quoted local parts (e.g., "test")
	if strings.HasPrefix(local, "\"") && strings.HasSuffix(local, "\"") {
		// Strip quotes for normalization but ensure internal validity
		quotedContent := local[1 : len(local)-1]
		re := regexp.MustCompile(`[^a-zA-Z0-9 !#$%&'*+/=?^_` + "`{|}~.@]" + ``)
		normalized := re.ReplaceAllString(quotedContent, "")
		if quotedContent != normalized {
			return "", errors.New("invalid characters in quoted local part")
		}
		return quotedContent, nil
	}

	// Validate unquoted local parts
	re := regexp.MustCompile(`[^a-zA-Z0-9!#$%&'*+/=?^_` + "`{|}~.-]" + ``)
	cleaned := re.ReplaceAllString(local, "")
	if cleaned == "" {
		return "", errors.New("invalid characters in local part")
	}

	// Remove consecutive dots
	cleaned = strings.ReplaceAll(cleaned, "..", ".")

	// Ensure no leading or trailing dots
	cleaned = strings.Trim(cleaned, ".")

	if cleaned == "" {
		return "", errors.New("local part cannot be empty after sanitization")
	}

	return cleaned, nil
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
