// api_key_format.go

/*
Purpose: An example of a custom hook designed to validate Stripe-style **test-only** API keys
using the Go Input Sanitization Library (goisl). This ensures that developers don't accidentally
use live production keys in development or test CLIs.

✅ Valid input examples (test only):
    go run api_key_format.go --apikey="sk_test_51H6u5eK3NjFspnKAEyB8uDUIpJnXr7v98XsJQw4Z6xuPbT8c"
    go run api_key_format.go --apikey="pk_test_abc123XYZ456morecharacters"

❌ Invalid input examples (production keys or wrong format):
    go run api_key_format.go --apikey="sk_live_abc123XYZ456"      // ❌ production key
    go run api_key_format.go --apikey="pk_live_something"         // ❌ production key
    go run api_key_format.go --apikey="api_key_123456"            // ❌ invalid prefix
    go run api_key_format.go --apikey="short"                     // ❌ too short
*/

package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// Only allow test-mode keys: sk_test_ or pk_test_ followed by 16+ alphanumeric characters
var stripeTestKeyPattern = regexp.MustCompile(`^(sk|pk)_test_[a-zA-Z0-9]{16,}$`)

func stripeTestOnlyHook(input string) (string, error) {
	key := strings.TrimSpace(input)

	// Block any live keys explicitly
	if strings.HasPrefix(key, "sk_live_") || strings.HasPrefix(key, "pk_live_") {
		return "", errors.New("❌ live production API keys are not allowed in this CLI")
	}

	// Enforce test-only key format
	if !stripeTestKeyPattern.MatchString(key) {
		return "", errors.New("invalid Stripe test API key format")
	}

	return key, nil
}

func main() {
	apiKeyFlag := isl.BindSanitizedFlag("apikey", "", "Stripe test API key to validate", stripeTestOnlyHook)
	pflag.Parse()

	key, err := apiKeyFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid API key:", err)
	} else {
		fmt.Println("✅ Stripe test API key is valid:", key)
	}
}
