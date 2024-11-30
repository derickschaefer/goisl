# goisl: Go Input Sanitization Library

## Overview

`goisl` (Go Input Sanitization Library) is a lightweight and secure library designed to simplify input sanitization and escaping in Go applications. Inspired by WordPress's input-handling functions, `goisl` provides developers with intuitive tools to ensure safe and clean input processing.

The library includes functions for:
- **Sanitization**: Cleaning and validating user inputs such as text, email, file names, and URLs.
- **Escaping**: Safely encoding output for various contexts such as HTML, JavaScript, and URLs.

By using `goisl`, you can reduce common vulnerabilities like XSS, SQL injection, and malformed input issues.

---

## Features

### Core Functions
- **Sanitization**:
  - `SanitizeEmail`: Validate and sanitize email addresses.
  - `SanitizeURL`: Validate and sanitize URLs.
  - `SanitizeFileName`: Clean and validate file names.
  - `HTMLSanitize`: Safely sanitize HTML content.

- **Escaping**:
  - `EscapeHTML`: Escape special characters for safe HTML rendering.
  - `EscapeURL`: Safely encode URL query parameters.

### Modular Design
- Each function is self-contained and can be used independently.
- Support for optional custom hooks to tailor sanitization or validation rules.

---

## Installation

Install the library using `go get`:

```bash
go get github.com/derickschaefer/goisl
```

## Usage

### Email Sanitization

```bash
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl/pkg"
)

func main() {
    email, err := pkg.SanitizeEmail(" user@example.com ", nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sanitized Email:", email)
    }
}
```

### URL Escaping

```bash
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl/pkg"
)

func main() {
    url, err := pkg.EscapeURL("  http://example.com/path?query=<script>  ", "display", nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Escaped URL:", url)
    }
}
```

### File Name Sanitization

```bash
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl/pkg"
)

func main() {
    fileName, err := pkg.SanitizeFileName("example#@!.txt", nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sanitized File Name:", fileName)
    }
}
```

## Contributing

Contributions are welcome! Please:
	1.	Fork the repository.
	2.	Submit a pull request with a detailed description of the changes.

This project is licensed under the MIT License. See the LICENSE file for details.
