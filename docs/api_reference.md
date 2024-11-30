# API Reference

## `SanitizeEmail`

### Description:
Sanitizes an email address by:
- Trimming whitespace.
- Validating its structure.
- Allowing optional hooks for custom behavior.

### Usage:
```go
result, err := SanitizeEmail(input, hook)
```
SanitizeEmail
-------------
Description:
  Sanitizes an email address by trimming whitespace, validating its structure,
  and applying optional custom logic through a user-defined hook.

Usage:
  result, err := SanitizeEmail(input, hook)

Parameters:
  input (string)  - The email address to sanitize.
  hook (EmailHook) [optional] - A function for custom sanitization logic.
                                Defaults to nil if not provided.

Return Values:
  - result (string): The sanitized email address.
  - err (error): An error if the input is invalid or fails custom hook validation.

Custom Hook Example:
  EmailHook type:
    func(local, domain string) (string, string, error)

Example Usage:
  result, err := SanitizeEmail("  user@example.com  ", nil)
  if err != nil {
      log.Fatalf("Error: %v", err)
  }
  fmt.Println(result)  # Output: user@example.com

EscapeURL
---------
Description:
  Sanitizes and escapes a URL for safe use, trimming whitespace, validating its structure,
  and optionally applying custom logic through a user-defined hook.

Usage:
  result, err := EscapeURL(input, context, hook)

Parameters:
  input (string)   - The URL to sanitize and escape.
  context (string) - The context for escaping (e.g., "display" for HTML output).
  hook (URLHook) [optional] - A function for custom URL transformation logic.
                              Defaults to nil if not provided.

Return Values:
  - result (string): The sanitized and escaped URL.
  - err (error): An error if the input URL is invalid or fails custom hook validation.

Custom Hook Example:
  URLHook type:
    func(parsedURL *url.URL) (*url.URL, error)

Example Usage:
  Basic:
    result, err := EscapeURL("  http://example.com/path?query=<script>  ", "display", nil)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Println(result)  # Output: http://example.com/path?query=%3Cscript%3E

  With Custom Hook:
    customHook := func(parsedURL *url.URL) (*url.URL, error) {
        if parsedURL.Scheme == "http" {
            parsedURL.Scheme = "https"
        }
        query := parsedURL.Query()
        if _, exists := query["tracking_id"]; exists {
            return nil, errors.New("tracking_id is not allowed")
        }
        return parsedURL, nil
    }

    result, err := EscapeURL("http://example.com/path?tracking_id=12345", "display", customHook)
    if err != nil {
        fmt.Println("Error:", err)  # Output: Error: tracking_id is not allowed
    } else {
        fmt.Println(result)
    }
