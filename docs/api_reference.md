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
```go
  result, err := SanitizeEmail(input, hook)
```
Parameters:
  input (string)  - The email address to sanitize.
  hook (EmailHook) [optional] - A function for custom sanitization logic.
                                Defaults to nil if not provided.

Return Values:
  - result (string): The sanitized email address.
  - err (error): An error if the input is invalid or fails custom hook validation.

Custom Hook Example:
  EmailHook type:
  ```go
    func(local, domain string) (string, string, error)
  ```

Example Usage:
```go
  result, err := SanitizeEmail("  user@example.com  ", nil)
  if err != nil {
      log.Fatalf("Error: %v", err)
  }
  fmt.Println(result)  # Output: user@example.com
```

EscapeURL
---------
Description:
  Sanitizes and escapes a URL for safe use, trimming whitespace, validating its structure,
  and optionally applying custom logic through a user-defined hook.

Usage:
```go
  result, err := EscapeURL(input, context, hook)
```

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
  ```go
    func(parsedURL *url.URL) (*url.URL, error)
```

Example Usage:
  Basic:
  ```go
    result, err := EscapeURL("  http://example.com/path?query=<script>  ", "display", nil)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Println(result)  # Output: http://example.com/path?query=%3Cscript%3E
  ```
  With Custom Hook:
  ```go
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
  ```
## SanitizeURL

### Description:

Validates and escapes a URL by:
	•	Trimming whitespace.
	•	Validating its structure.
	•	Optionally applying a custom hook.

### Usage:
```go
result, err := EscapeURL(input, context, hook)
```
Parameters:

	•	input (string): The URL to sanitize.
	•	context (string): Escaping context (e.g., “display”).
	•	hook (URLHook, optional): A custom function for additional validation.

Return Values:

	•	result (string): The sanitized URL.
	•	err (error): An error if validation fails.

## SanitizeFileName

### Description:

Sanitizes a file name by:
	•	Removing unsafe characters.
	•	Normalizing Unicode characters.
	•	Ensuring file name constraints.

### Usage:
 
```go
result, err := SanitizeFileName(input, hook)
```
Parameters:

	•	input (string): The file name to sanitize.
	•	hook (FileNameHook, optional): A custom function for additional validation.

Return Values:

	•	result (string): The sanitized file name.
	•	err (error): An error if validation fails.

## HTMLSanitize

### Description:

Sanitizes HTML content by:
	•	Removing unsafe tags and attributes.
	•	Normalizing HTML entities.

### Usage:
```go
result := HTMLSanitize(input)
```
Parameters:

	•	input (string): The HTML content to sanitize.

Return Values:

	•	result (string): The sanitized HTML content.

## IsAllowedProtocol

### Description:

Validates URL schemes against an allowed list.

### Usage:
```go
result := IsAllowedProtocol(scheme, allowedProtocols)
```
Parameters:

	•	scheme (string): The URL scheme to validate.
	•	allowedProtocols ([]string): A list of allowed protocols.

Return Values:

	•	result (bool): True if the scheme is allowed, false otherwise.
