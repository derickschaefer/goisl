// crypto_btc_address.go
//
// Example CLI: Validate Bitcoin address input.
//
// Usage:
//   go run crypto_btc_address.go --btc=bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kygt080
//
// This program sanitizes a BTC address and checks for basic format validity
// based on common BTC address prefixes.
//
// Author: Derick Schaefer
// License: MIT

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

func allowBTCChars() []rune {
	return []rune("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
}

func main() {
	address := pflag.String("btc", "", "Bitcoin address to validate")
	pflag.Parse()

	if *address == "" {
		fmt.Println("❌ Error: --btc is required")
		os.Exit(1)
	}

	sanitized := isl.EscapePlainText(*address, allowBTCChars)

	if strings.HasPrefix(sanitized, "1") ||
		strings.HasPrefix(sanitized, "3") ||
		strings.HasPrefix(sanitized, "bc1") {
		fmt.Println("✅ Valid BTC address:", sanitized)
	} else {
		fmt.Println("❌ Invalid BTC address format")
	}
}
