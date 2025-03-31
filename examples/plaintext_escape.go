// plaintext_escape.go

/*
Purpose: A basic example demonstrating how to escape plain text input using
the Go Input Sanitization Library (goisl)

Examples:
    go run plaintext_escape.go --text="Hello, world!"
    go run plaintext_escape.go --text="Clean me up 💥💣🔥!"
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/derickschaefer/goisl"
)

// ANSI escape codes for colored text
const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func main() {
	// Check if input is provided
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// Concatenate input if it spans multiple arguments
	input := strings.Join(os.Args[1:], " ")

	// Escape plain text
	escaped := isl.EscapePlainText(input, nil)

	// Highlight the parts removed in RED
	highlighted := highlightDifferences(input, escaped)

	// Print the before and after
	fmt.Println("Before:", highlighted)
	fmt.Println("After: ", escaped)
}

// highlightDifferences highlights removed characters in the original string
func highlightDifferences(original, escaped string) string {
	originalTrimmed := strings.TrimSpace(original)
	escapedTrimmed := strings.TrimSpace(escaped)

	builder := strings.Builder{}
	oi, ei := 0, 0

	for oi < len(originalTrimmed) {
		if ei < len(escapedTrimmed) && originalTrimmed[oi] == escapedTrimmed[ei] {
			builder.WriteByte(originalTrimmed[oi])
			ei++
		} else {
			builder.WriteString(Red + string(originalTrimmed[oi]) + Reset)
		}
		oi++
	}

	return builder.String()
}

// printHelp displays help information
func printHelp() {
	fmt.Println("Usage: plaintext_escaper <input>")
	fmt.Println("Example:")
	fmt.Println(`  plaintext_escaper "   Hello, World!   "`)
	fmt.Println("Output:")
	fmt.Println(`  Before:    "   Hello, World!   "`)
	fmt.Println(`  After:     "Hello World"`)
}
