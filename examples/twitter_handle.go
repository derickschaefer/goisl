// twitter_handle.go

/*
Purpose: An example of a custom hook designed to validate Twitter handles using
the Go Input Sanitization Library (goisl)

Valid handle example:
    go run twitter_handle.go --handle="@golang"

Invalid handle examples:
    go run twitter_handle.go --handle="golang"
    go run twitter_handle.go --handle="@thishandleiswaytoolongtobevalid"
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

// TwitterHandleHook validates and sanitizes a Twitter handle.
func TwitterHandleHook(input string) (string, error) {
	input = strings.TrimSpace(input)

	// Ensure it starts with '@'
	if !strings.HasPrefix(input, "@") {
		return "", errors.New("Twitter handle must start with '@'")
	}

	// Remove the '@' for validation
	handle := input[1:]

	// Validate using regex: 1–15 chars, alphanumeric + underscores
	matched, err := regexp.MatchString(`^[A-Za-z0-9_]{1,15}$`, handle)
	if err != nil || !matched {
		return "", errors.New("invalid Twitter handle format")
	}

	// Re-add '@' to sanitized output
	return "@" + handle, nil
}

func main() {
	twitterFlag := isl.BindSanitizedFlag("handle", "", "Twitter handle (e.g. @gopher)", TwitterHandleHook)
	pflag.Parse()

	handle, err := twitterFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid Twitter handle:", err)
		return
	}

	fmt.Println("✅ Twitter handle is valid:", handle)
}
