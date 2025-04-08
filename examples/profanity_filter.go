// Example: Censor Profanity in Plain Text with a Custom Hook
//
// This CLI-style example accepts a --text flag and sanitizes the input using
// EscapePlainText with a custom hook that filters out profanity by removing
// offensive words.
//
// Usage:
//
//	go run censor_profanity.go --text="what the hell is this damn thing"
//
// Example Output:
//
//	✅ Sanitized Text: what the h*** is this d*** thing?
package main

import (
	"flag"
	"fmt"
	"strings"
)

// Custom profanity sanitizer using simple replacement
func ProfanityFilter(input string) string {
	badWords := map[string]string{
		"damn": "d***",
		"hell": "h***",
		"shit": "s***",
		"fuck": "f***",
	}
	inputLower := strings.ToLower(input)
	for word, replacement := range badWords {
		inputLower = strings.ReplaceAll(inputLower, word, replacement)
	}
	return inputLower
}

func main() {
	input := flag.String("input", "", "A string to clean (e.g., 'what the hell is this damn thing')")
	flag.Parse()

	cleaned := ProfanityFilter(*input)
	fmt.Println(cleaned)
}
