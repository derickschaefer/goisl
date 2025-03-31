// guid_format.go
//
// Example CLI: Validate a GUID-style identifier.
//
// Usage:
//   go run guid_format.go --guid=550e8400-e29b-41d4-a716-446655440000
//
// This program sanitizes input using isl.EscapePlainText and validates
// it against a basic GUID pattern.
//
// Author: Derick Schaefer
// License: MIT

package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

func allowHyphen() []rune {
	return []rune("-")
}

func main() {
	guid := pflag.String("guid", "", "GUID to validate")
	pflag.Parse()

	if *guid == "" {
		fmt.Println("❌ Error: --guid is required")
		os.Exit(1)
	}

	sanitized := isl.EscapePlainText(*guid, allowHyphen)

	matched, _ := regexp.MatchString(`^[a-fA-F0-9\-]{36}$`, sanitized)
	if matched {
		fmt.Println("✅ Valid GUID:", sanitized)
	} else {
		fmt.Println("❌ Invalid GUID format")
	}
}
