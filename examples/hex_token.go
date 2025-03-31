// hex_token.go

/*
Purpose: An example of a custom hook designed to validate hex-encoded tokens
(commonly used in APIs, email confirmation links, password resets, etc.)
using the Go Input Sanitization Library (goisl).

Valid token examples:
    go run hex_token.go --token="a3f5b9e7c0124d89b56f10ae7db394c8"
    go run hex_token.go --token="ABCDEF1234567890abcdef1234567890"

Invalid token examples:
    go run hex_token.go --token="not-a-token"
    go run hex_token.go --token="12345"
    go run hex_token.go --token="g1234567890abcdef1234567890"   // contains invalid 'g'
*/

package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// Custom hook to validate hex-encoded tokens
func hexTokenHook(input string) (string, error) {
	matched, err := regexp.MatchString(`^[a-fA-F0-9]{32,64}$`, input)
	if err != nil {
		return "", err
	}
	if !matched {
		return "", errors.New("invalid hex token format")
	}
	return input, nil
}

func main() {
	tokenFlag := isl.BindSanitizedFlag("token", "", "Hex-encoded token", func(input string) (string, error) {
		sanitized := isl.EscapePlainText(input, nil)
		return hexTokenHook(sanitized)
	})
	pflag.Parse()

	token, err := tokenFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid token:", err)
	} else {
		fmt.Println("✅ Token is valid:", token)
	}
}
