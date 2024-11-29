## SanitizeEmail

### Description
Sanitizes an email address by trimming whitespace, validating its structure, and optionally applying custom logic through a user-defined hook.

### Parameters
- **`input (string)`**: The email address to sanitize.
- **`hook (EmailHook)`** *(optional)*: A user-defined function to customize sanitization behavior. If `nil`, the default sanitization logic is applied.

### Return Values
- **`string`**: The sanitized email address.
- **`error`**: An error if the input is invalid or fails custom hook validation.

### Custom Hook
The `EmailHook` type is defined as:
```go
type EmailHook func(local, domain string) (string, string, error)

## Example - Basic Usage

email, err := SanitizeEmail("  user@example.com  ", nil)
if err != nil {
    log.Fatalf("Error: %v", err)
}
fmt.Println(email) // Output: user@example.com

## Example - Custom Hook Example

customHook := func(local, domain string) (string, string, error) {
    if plusIndex := strings.Index(local, "+"); plusIndex != -1 {
        local = local[:plusIndex]
    }
    if domain == "tempmail.com" {
        return "", "", errors.New("blocked domain")
    }
    return local, domain, nil
}

email, err := SanitizeEmail("user+tag@tempmail.com", customHook)
if err != nil {
    fmt.Println("Error:", err) // Output: Error: blocked domain
} else {
    fmt.Println("Sanitized Email:", email)
}
