// pkg/escape.go

package pkg

import (
    "errors"
    "net/url"
    "path"
    "regexp"
    "strings"
)

// URLHook defines a function signature for custom URL processing.
type URLHook func(parsedURL *url.URL) (*url.URL, error)

// EscapeAllowedProtocols defines the list of acceptable URL schemes for escaping.
var EscapeAllowedProtocols = []string{"http", "https", "mailto", "ftp"}

// EscapePlainText sanitizes plain text by removing punctuation, special characters,
// and non-alphanumeric symbols while condensing whitespace into a single space.
func EscapePlainText(input string) string {
    // Trim leading and trailing whitespace
    input = strings.TrimSpace(input)

    // Replace sequences of whitespace with a single space
    whitespaceRegex := regexp.MustCompile(`\s+`)
    input = whitespaceRegex.ReplaceAllString(input, " ")

    // Remove non-alphanumeric characters and punctuation
    alphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9\s]`)
    input = alphanumericRegex.ReplaceAllString(input, "")

    // Trim any residual spaces after special character removal
    input = strings.TrimSpace(input)

    return input
}

// SafeEscapeHTML escapes only specific characters, excluding '%'.
func SafeEscapeHTML(input string) string {
    replacer := strings.NewReplacer(
        "&", "&amp;",
        "<", "&lt;",
        ">", "&gt;",
        "\"", "&quot;",
        "'", "&#39;",
    )
    return replacer.Replace(input)
}

// EscapeURL sanitizes and escapes a URL, applying an optional custom hook.
func EscapeURL(input string, context string, hook URLHook) (string, error) {
    // Step 1: Trim and replace spaces with %20
    input = strings.TrimSpace(input)
    input = strings.ReplaceAll(input, " ", "%20")

    // Early return if input is empty after trimming
    if input == "" {
        return "", nil
    }

    // Step 3: Parse the URL
    parsed, err := url.Parse(input)
    if err != nil {
        return "", errors.New("invalid URL")
    }

    // Step 4: Add a scheme if missing
    if parsed.Scheme == "" {
        if strings.HasPrefix(input, "/") || strings.HasPrefix(input, "#") || strings.HasPrefix(input, "?") {
            // Relative URL
            return input, nil
        }
        input = "http://" + input
        parsed, err = url.Parse(input)
        if err != nil {
            return "", errors.New("invalid URL after adding scheme")
        }
    }

    // Step 5: Normalize scheme and host
    parsed.Scheme = strings.ToLower(parsed.Scheme)
    parsed.Host = strings.ToLower(parsed.Host)

    // Step 6: Validate the domain without Punycode
    domainRegex := regexp.MustCompile(`^[a-zA-Z0-9.-]+(:[0-9]+)?$`)
    if !domainRegex.MatchString(parsed.Host) {
        return "", errors.New("invalid characters in domain")
    }

    // Step 7: Apply the custom hook, if provided
    if hook != nil {
        parsed, err = hook(parsed)
        if err != nil {
            return "", err
        }
    }

    // Step 8: Sanitize query parameters
    if err := sanitizeQueryParameters(parsed); err != nil {
        return "", err
    }

    // Step 9: Clean the path
    parsed.Path = path.Clean(parsed.Path)
    if parsed.Path == "." || parsed.Path == "/." {
        parsed.Path = ""
    }

    // Step 10: Escape fragment if present
    if parsed.Fragment != "" {
        parsed.Fragment = url.PathEscape(parsed.Fragment)
    }

    // Step 11: Context-based transformations
    finalURL := parsed.String()
    if context == "display" {
        // Use SafeEscapeHTML to avoid escaping '%'
        finalURL = SafeEscapeHTML(finalURL)
    }

    // Step 12: Encode brackets
    finalURL = strings.ReplaceAll(finalURL, "[", "%5B")
    finalURL = strings.ReplaceAll(finalURL, "]", "%5D")

    return finalURL, nil
}

// sanitizeQueryParameters sanitizes query parameters in the URL.
func sanitizeQueryParameters(parsed *url.URL) error {
    if len(parsed.RawQuery) > 0 {
        query := parsed.Query()        // Parse query into a map
        rawQuery := make([]string, 0) // To store sanitized key-value pairs
        for key, values := range query {
            for _, value := range values {
                // Decode existing encoded values to avoid double encoding
                decodedValue, err := url.QueryUnescape(value)
                if err != nil {
                    return errors.New("failed to decode query parameter")
                }

                // Properly escape the value using url.QueryEscape
                escapedValue := url.QueryEscape(decodedValue)

                // Replace '+' with '%20' to match test expectations
                escapedValue = strings.ReplaceAll(escapedValue, "+", "%20")

                // Add sanitized key-value pair to raw query
                rawQuery = append(rawQuery, key+"="+escapedValue)
            }
        }
        // Manually construct the RawQuery
        parsed.RawQuery = strings.Join(rawQuery, "&")
    }
    return nil
}
