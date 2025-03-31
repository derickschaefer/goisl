// mask_last4.go

/*
Purpose: An example of masking sensitive input using goisl. This example takes
any input string (e.g., API key, token, secret) and replaces all characters
with `#`, leaving only the last 4 characters visible for reference or logging.

✅ Example:
    go run mask_last4.go --secret="sk_test_4eC39HqLyjWDarjtT1zdp7dc"
    Output: Masked: ###########################p7dc

✅ Another example:
    go run mask_last4.go --secret="123456789"
    Output: Masked: #####6789

❌ Edge case:
    go run mask_last4.go --secret="abc"
    Output: Masked: *** (Fully masked — too short)
*/

package main

import (
	"fmt"
	"strings"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// maskLast4 replaces all characters with '#' except the last 4.
func maskLast4(input string) (string, error) {
	input = strings.TrimSpace(input)
	if len(input) <= 4 {
		return strings.Repeat("*", len(input)), nil // Fully mask short inputs
	}

	visible := input[len(input)-4:]
	masked := strings.Repeat("#", len(input)-4) + visible
	return masked, nil
}

func main() {
	secretFlag := isl.BindSanitizedFlag("secret", "", "Sensitive value to mask", maskLast4)
	pflag.Parse()

	masked, err := secretFlag.Get()
	if err != nil {
		fmt.Println("❌ Error masking secret:", err)
	} else {
		fmt.Println("🔒 Masked:", masked)
	}
}
