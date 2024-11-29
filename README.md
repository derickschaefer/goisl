# goisl: Go Input Sanitization Library

## Overview

`goisl` (Go Input Sanitization Library) is a lightweight and secure library designed to simplify input sanitization and escaping in Go applications. Inspired by WordPress's input-handling functions, `goisl` provides developers with intuitive tools to ensure safe and clean input processing.

The library includes functions for:
- **Sanitization**: Cleaning and validating user inputs such as text, email, and URLs.
- **Escaping**: Safely encoding output for various contexts such as HTML, JavaScript, and URLs.

By using `goisl`, you can reduce common vulnerabilities like XSS, SQL injection, and malformed input issues.

---

## Features

### Sanitization Functions
- `SanitizeTextField`: Clean text inputs by removing invalid characters and trimming whitespace.
- `SanitizeEmail`: Validate and sanitize email addresses.
- `SanitizeURL`: Validate and sanitize URLs for safe usage.

### Escaping Functions
- `EscapeHTML`: Escape special characters in user-generated HTML.
- `EscapeAttribute`: Escape input for safe use in HTML attributes.
- `EscapeURL`: Safely encode URL query parameters.
- `EscapeJS`: Escape strings for safe inclusion in JavaScript.

---

## Installation

Install the library using `go get`:
```bash
go get github.com/yourusername/goisl

## Usage - Basic Example

### Sanitizing Email Input
The `SanitizeEmail` function trims whitespace, validates email structure, and allows optional hooks for custom sanitization behavior.

#### Basic Usage
```go
package main

import (
    "fmt"
    "log"
    "github.com/derickschaefer/goisl/pkg"
)

func main() {
    email, err := pkg.SanitizeEmail("  user@example.com  ", nil) // No custom hook
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Println("Sanitized Email:", email) // Output: user@example.com
}

## Usage = Custom Hook Example

package main

import (
    "errors"
    "fmt"
    "log"
    "strings"
    "github.com/derickschaefer/goisl/pkg"
)

func main() {
    // Define a custom hook
    customHook := func(local, domain string) (string, string, error) {
        // Remove tags after '+'
        if plusIndex := strings.Index(local, "+"); plusIndex != -1 {
            local = local[:plusIndex]
        }

        // Block specific domains
        blockedDomains := []string{"tempmail.com"}
        domainLower := strings.ToLower(domain) // Normalize to lowercase
        for _, blocked := range blockedDomains {
            if domainLower == strings.ToLower(blocked) {
                return "", "", errors.New("blocked domain")
            }
        }

        return local, domain, nil
    }

    email, err := pkg.SanitizeEmail("user+tag@tempmail.com", customHook)
    if err != nil {
        log.Printf("Error: %v", err) // Output: Error: blocked domain
    } else {
        fmt.Println("Sanitized Email:", email)
    }
}

package main

import (
    "fmt"
    "goisl"
)

func main() {
    // Escape HTML
    escapedHTML := goisl.EscapeHTML("<script>alert('XSS!')</script>")
    fmt.Println("Escaped HTML:", escapedHTML)

    // Escape URL
    escapedURL := goisl.EscapeURL("https://example.com?param=<value>")
    fmt.Println("Escaped URL:", escapedURL)
}

Contributing

Contributions are welcome! If you’d like to help improve goisl, please:
	1.	Fork the repository.
	2.	Submit a pull request with a detailed description of the changes.

This project is licensed under the MIT License. See the LICENSE file for details.
