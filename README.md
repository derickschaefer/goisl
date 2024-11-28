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

package main

import (
    "fmt"
    "log"
    "goisl"
)

func main() {
    // Sanitize email
    email, err := goisl.SanitizeEmail("  user@example.com  ")
    if err != nil {
        log.Fatalf("Invalid email: %v", err)
    }
    fmt.Println("Sanitized Email:", email)

    // Sanitize text field
    cleanText := goisl.SanitizeTextField("   Hello <World>!   ")
    fmt.Println("Sanitized Text:", cleanText)
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
