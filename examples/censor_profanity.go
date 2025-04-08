// Example: Censor Profanity in Text Input Using a Custom Escape Hook
//
// This CLI-style example uses a --text flag and censors profane words
// using a pre-filter before running through EscapePlainText.
//
// Usage:
//   go run censor_profanity.go --text="damn this is shit"
//
// Example Output:
//   ✅ Sanitized Text: d*** this is s***

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	isl "github.com/derickschaefer/goisl"
)

// Pre-filter for known profanity
func censorProfanity(input string) string {
	replacer := strings.NewReplacer(
		"damn", "d***",
		"shit", "s***",
		"jackass", "jack**",
		"hell", "h***",
	)
	return replacer.Replace(input)
}

func profanityHook() []rune {
	// Allow a few extra symbols that are not profane
	return []rune("!@#$%^&*()-_=+[]{}|;:,.?/") 
}

func main() {
	textFlag := flag.String("text", "", "Text to sanitize")
	flag.Parse()

	if *textFlag == "" {
		fmt.Println("❌ Error: --text flag is required")
		os.Exit(1)
	}

	// Step 1: Censor profanity
	cleaned := censorProfanity(*textFlag)

	// Step 2: Sanitize further using goisl's EscapePlainText with hook
	sanitized := isl.EscapePlainText(cleaned, profanityHook)

	fmt.Println("✅ Sanitized Text:", sanitized)
}
