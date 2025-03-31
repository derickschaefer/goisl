# goisl: Go Input Sanitization Library

## Overview

`goisl` (Go Input Sanitization Library) is a lightweight, secure, and developer-friendly Go package designed to simplify input sanitization and safe output escaping. It supports three core use cases:

1. **Built-in Sanitization & Escaping**: Quickly sanitize and escape common input types including email addresses, file names, URLs, and HTML for safe use in CLI tools, APIs, or web templates.
2. **Custom Format Overrides**: Use hookable functions or wrappers to validate and clean highly specific input formats such as crypto addresses, UUIDs, IP ranges, Twitter handles, and vendor-specific API keys.
3. **CLI Integration with `pflag`**: Seamlessly bind sanitized values to command-line flags using `BindSanitizedFlag` and `BindSanitizedTextFlag`, with panic-safe variants for required inputs.

Inspired by WordPress’s input-handling philosophy, `goisl` encourages a “sanitize on input, escape on output” model with modular, testable, and override-ready functions.

By using `goisl`, you can reduce the risk of malformed input, injection attacks, and XSS vulnerabilities while maintaining clean, consistent input handling across your Go applications.

---

## Features

### Core Functions

- **Sanitization**:
  - `SanitizeEmail`: Validate and sanitize email addresses with optional hooks.
  - `SanitizeFileName`: Clean and normalize filenames while enforcing safe character sets.
  - `SanitizeURL`: Validate and escape user-supplied URLs.
  - `HTMLSanitize`: Strip unsafe tags from HTML strings.

- **Escaping**:
  - `EscapePlainText`: Strip unsafe characters from plain text input (e.g., names, labels).
  - `EscapeURL`: Sanitize URLs and encode query parameters.
  - `SafeEscapeHTML`: Escape characters for safe HTML display (excluding `%`).

---

## New in v1.1.0

- 🧩 `SanitizeXBasic` helpers: Simple wrappers with default rules and no hooks.
- 🚨 `MustSanitizeXBasic`: Panic-on-failure versions for CLI defaults or enforced logic.
- ✅ `BindSanitizedFlag` and `BindSanitizedTextFlag`: Bind input flags with automatic sanitization via `pflag`.
- 🧪 100% unit test coverage with extensive edge case testing.
- 🌟 Dozens of real-world override examples (API keys, crypto, social handles, and more).

---

## Installation

```bash
go get github.com/derickschaefer/goisl
## Installation

Install the library by importing it and using `go mod tidy`:

```bash
go get github.com/derickschaefer/goisl
```

## Usage

### Email Sanitization

```go
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl"
)

func main() {
    email, err := isl.SanitizeEmail(" user@example.com ", nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sanitized Email:", email)
    }
}
```

### URL Escaping

```go
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl"
)

func main() {
    url, err := isl.EscapeURL("  http://example.com/path?query=<script>  ", "display", nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Escaped URL:", url)
    }
}
```

### File Name Sanitization

```go
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl"
)

func main() {
    fileName, err := isl.SanitizeFileName("example#@!.txt", nil)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sanitized File Name:", fileName)
    }
}
```

### CLI Integration with pflag

```go
package main

import (
    "fmt"
    "github.com/derickschaefer/goisl"
    "github.com/spf13/pflag"
)

func main() {
    emailFlag := isl.BindSanitizedFlag("email", "", "Email address", isl.SanitizeEmailBasic)
    urlFlag := isl.BindSanitizedFlag("url", "", "URL to sanitize", isl.SanitizeURLBasic)
    commentFlag := isl.BindSanitizedTextFlag("comment", "", "Plain text comment", nil)

    pflag.Parse()

    fmt.Println("✅ Email:", emailFlag.MustGet())
    fmt.Println("✅ URL:", urlFlag.MustGet())
    fmt.Println("✅ Comment:", commentFlag.MustGet())
}
```

### Custom hook to block disposable email domains

```go
package main

import (
    "errors"
    "fmt"
    "strings"

    "github.com/derickschaefer/goisl"
)

func main() {
    // Define a hook that strips email tags and blocks disposable domains
    customHook := func(local, domain string) (string, string, error) {
        // Remove anything after a '+' (e.g., user+tag@example.com → user@example.com)
        if plus := strings.Index(local, "+"); plus != -1 {
            local = local[:plus]
        }

        // Block disposable email domains
        blocked := []string{"tempmail.com", "mailinator.com"}
        domainLower := strings.ToLower(domain)
        for _, b := range blocked {
            if domainLower == b {
                return "", "", errors.New("disposable email domain not allowed")
            }
        }

        return local, domain, nil
    }

    input := "User+promo@tempmail.com"
    sanitized, err := isl.SanitizeEmail(input, customHook)
    if err != nil {
        fmt.Println("❌ Error:", err)
    } else {
        fmt.Println("✅ Sanitized:", sanitized)
    }
}
```

### Custom hook to block known URL shortners

```go
package main

import (
    "errors"
    "fmt"
    "net/url"
    "strings"

    "github.com/derickschaefer/goisl"
)

func main() {
    // Hook that blocks bit.ly and other known shorteners
    blockShorteners := func(parsed *url.URL) (*url.URL, error) {
        shorteners := []string{"bit.ly", "tinyurl.com", "t.co"}

        host := strings.ToLower(parsed.Host)
        for _, s := range shorteners {
            if host == s {
                return nil, errors.New("URL shorteners are not allowed")
            }
        }
        return parsed, nil
    }

    input := "https://bit.ly/3xyzABC"
    result, err := isl.EscapeURL(input, "display", blockShorteners)
    if err != nil {
        fmt.Println("❌ Error:", err)
    } else {
        fmt.Println("✅ Escaped URL:", result)
    }
}
```

## Contributing

Contributions are welcome! Please:
	1.	Fork the repository. (not preferred but permissible)
	2.	Submit a pull request with a detailed description of the changes.
    3.  Please share custom hook examples to grow that repository.

This project is licensed under the MIT License. See the LICENSE file for details.
