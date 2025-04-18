# Architecture of `goisl`

## Design Principles

`goisl` is designed with the following principles in mind:

1. **Simplicity**
   - Provide intuitive, easy-to-use functions for sanitization and escaping.
   - Focus on common formats: emails, URLs, file names, HTML, and text.

2. **Security**
   - Assist in preventing XSS, SQL injection, and malformed input vulnerabilities.
   - Built on trusted Go libraries like `net/mail`, `net/url`, and `html`.

3. **Modularity**
   - Functions are separated into focused files.
   - Developers can implement extended use cases by hooking core functionality.

4. **CLI-first Focus**
   - Introduces `BindSanitizedFlag()` and `BindSanitizedTextFlag()` for seamless integration with `pflag`.
   - Simplifies input safety for command-line applications.

5. **Test-Driven Development**
   - Every feature is fully unit tested.
   - Makefile runs all custom hook example code in addition to unit tests.
   - Real-world edge cases are prioritized (e.g., malformed inputs, injection attempts).

---

## File Structure (v1.1.2)
```
goisl/
├── LICENSE                      # MIT license
├── Makefile                     # Build, test, and lint automation
├── README.md                    # Project overview and usage
├── go.mod / go.sum              # Go module dependencies
├── cli.go                       # CLI helper functions (BindSanitizedFlag)
├── common.go                    # Shared types and constants
├── doc.go                       # Package-level documentation
├── emailsanitize.go             # Email sanitization logic and hooks
├── escape.go                    # URL and text escaping functions
├── filesanitize.go              # File name sanitization logic
├── htmlsanitize.go              # Safe HTML sanitization
├── urlsanitize.go               # URL sanitization and validation
│
├── docs/                        # Developer documentation
│   ├── architecture.md          # Internal architecture overview
│   └── api_reference.md         # Function-by-function documentation
│
├── examples/                    # Copy-paste-ready usage patterns
│   ├── README.md
│   ├── api_key_format.go
│   ├── block_shorteners.go
│   ├── censor_profanity.go
│   ├── cli_flags.go
│   ├── country_code.go
│   ├── crypto_btc_address.go
│   ├── doc.go                   # Package documentation for examples
│   ├── german_filename.go
│   ├── guid_format.go
│   ├── hex_token.go
│   ├── ip_address.go
│   ├── mask_last4.go
│   ├── plaintext_esc           # Likely a temp/test file — consider removing
│   ├── plaintext_escape.go
│   ├── slack_webhook.go
│   ├── twitter_handle.go
│   └── uuid_format.go
│
└── tests/                       # Unit tests (90.4%+ coverage in v1.1.2)
    ├── cli_flag_helpers_test.go
    ├── common_validate_test.go
    ├── emailsanitize_local_test.go
    ├── emailsanitize_test.go
    ├── escape_test.go
    ├── filesanitize_panic_test.go
    ├── filesanitize_test.go
    ├── htmlsanitize_test.go
    ├── urlsanitize_test.go
    └── validate_protocol_test.go
```
---

## Key Components

### 1. Sanitization
Sanitization functions clean and validate input for safe internal use:
- `SanitizeEmail`: Trims and validates email addresses.
- `SanitizeFileName`: Cleans up unsafe file names.
- `SanitizeURL`: Parses and normalizes URLs.
- `HTMLSanitize`: Removes dangerous tags and attributes.

Each function supports optional **custom hooks** to allow format-specific overrides.

### 2. Escaping
Escaping functions prepare output for display in secure formats:
- `EscapePlainText`: Removes dangerous characters from user-submitted text.
- `EscapeURL`: Encodes query strings and unsafe URL characters.
- `SafeEscapeHTML`: Escapes characters for safe HTML output.

### 3. CLI Integration
A new feature in v1.1.0, CLI helpers bind input directly to flags with sanitization:
- `BindSanitizedFlag()`: Binds and sanitizes an input flag using a validator.
- `BindSanitizedTextFlag()`: Binds and escapes plain text safely.
- Also includes `.MustGet()` helpers for panic-on-error CLI logic.

---

## Workflow

### 1. Input Sanitization
Inputs are sanitized when received — from CLI flags, HTTP forms, or elsewhere.

### 2. Output Escaping
Before rendering (e.g., in HTML, logs, or terminal), data is safely escaped.

### 3. Testing
Each function has corresponding test coverage:
- Positive and negative test cases
- Custom hook validation
- Edge case behavior (empty values, unicode, malformed formats)

---

## Design Rationale

1. **Hook Support**:
   - All sanitizers accept optional hooks for flexible override logic.
   - This encourages safe customization without touching core logic.

2. **CLI-first Approach**:
   - Most Go libraries ignore CLI input safety — `goisl` brings sanitization to the command line with zero friction.

3. **Idiomatic Go**:
   - Uses idiomatic error returns.
   - Avoids magic or reflection-heavy solutions.

---

## Future Goals

- Add support for JSON and XML escaping.
- Continue to build the custom hook examples library.

---

## Conclusion

`goisl` provides a powerful yet simple way to sanitize and escape input in Go applications and CLIs. Its modular architecture, safety-first mindset, and real-world hook support make it an ideal building block for secure software.
