---

### **`architecture.md`**
```markdown
# Architecture of `goisl`

## Design Principles

`goisl` is designed with the following principles in mind:

1. **Simplicity**:
   - Provide intuitive and easy-to-use functions for sanitization and escaping.
   - Focus on common use cases like handling text, email, URLs, and HTML.

2. **Security**:
   - Ensure that input sanitization and escaping meet modern security standards.
   - Protect against vulnerabilities such as XSS, SQL injection, and malformed inputs.

3. **Modularity**:
   - Separate concerns by organizing sanitization and escaping functions into dedicated files.
   - Allow developers to use only the functions they need.

4. **Test-Driven Development**:
   - All functions are built with comprehensive unit tests covering valid and invalid inputs.
   - Edge cases and real-world scenarios are prioritized in test coverage.

---

## File Structure

goisl/
│
├── README.md              # Project overview and examples
├── LICENSE                # License file (e.g., MIT or Apache 2.0)
├── go.mod                 # Go module file for dependency management
├── examples/              # Example applications using the library
│   ├── sanitize_email.go  # Email sanitization example
│   ├── sanitize_url.go    # URL sanitization example
│   ├── escape_html.go     # HTML escaping example
│   └── …
├── pkg/                   # Core package implementation
│   ├── sanitize.go        # Functions for sanitization
│   ├── escape.go          # Functions for escaping
│   ├── utils.go           # Utility/helper functions
│   └── errors.go          # Custom error types and constants
├── tests/                 # Unit tests for sanitization and escaping
│   ├── sanitize_test.go   # Tests for sanitization functions
│   ├── escape_test.go     # Tests for escaping functions
│   └── …
└── docs/                  # Documentation
├── architecture.md        # Explanation of library design and architecture
└── api_reference.md       # Detailed API reference


---

## Key Components

### 1. **Sanitization**
Sanitization functions clean and validate user inputs, ensuring they conform to expected formats:
- `SanitizeTextField`: Removes invalid characters and trims whitespace from text fields.
- `SanitizeEmail`: Validates email addresses using `net/mail` and trims unnecessary spaces.
- `SanitizeURL`: Validates URLs using `net/url` and ensures proper encoding.

---

### 2. **Escaping**
Escaping functions safely encode data for output in different contexts:
- `EscapeHTML`: Encodes special characters for safe HTML rendering.
- `EscapeAttribute`: Escapes input for safe use in HTML attributes.
- `EscapeURL`: Encodes query parameters and special characters in URLs.
- `EscapeJS`: Escapes strings for safe inclusion in JavaScript.

---

## Workflow

### 1. Input Sanitization
Sanitize user input at the point of collection to ensure it is well-formed and safe for processing.

### 2. Output Escaping
Escape sanitized data before displaying it in specific contexts (e.g., HTML, JavaScript, URLs).

### 3. Test-Driven Development
Each function in `goisl` is developed with unit tests covering:
- Valid inputs
- Invalid inputs
- Edge cases (e.g., empty strings, malformed data)

---

## Design Rationale

1. **Sanitization and Escaping Separation**:
   - Sanitization prepares data for internal processing (e.g., validating an email).
   - Escaping prepares data for specific output contexts (e.g., rendering HTML).

2. **Extensibility**:
   - Functions are modular and easy to extend with additional features (e.g., new escaping methods for XML or JSON).

3. **Adherence to Go Best Practices**:
   - Follows idiomatic Go patterns for error handling and modular design.
   - Leverages Go's standard library whenever possible (e.g., `net/mail`, `net/url`).

---

## Future Goals

- Add support for JSON and XML escaping.
- Provide configurable policies for sanitization and escaping.
- Improve performance for large-scale sanitization and escaping use cases.

---

## Conclusion

`goisl` is designed to be a lightweight, secure, and modular library for input handling in Go. By separating sanitization and escaping into distinct, reusable functions, it helps developers write safer applications with minimal effort.
