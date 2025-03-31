// uuid_format.go

/*
Purpose: An example of a custom hook designed to validate UUID input using
the Go Input Sanitization Library (goisl)

Valid UUID example:
    go run uuid_format.go --uuid="123e4567-e89b-12d3-a456-426614174000"

Invalid UUID examples:
    go run uuid_format.go --uuid="bad-uuid"
    go run uuid_format.go --uuid="123e4567-e89b-12d3-a456-4266141"
*/

package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	isl "github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// UUIDRegex matches any valid UUID (version-agnostic, case-insensitive)
var UUIDRegex = regexp.MustCompile(`(?i)^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)

// UUIDHook validates the UUID format
func UUIDHook(input string) (string, error) {
	input = strings.TrimSpace(input)
	if !UUIDRegex.MatchString(input) {
		return "", errors.New("invalid UUID format")
	}
	return input, nil
}

func main() {
	uuidFlag := isl.BindSanitizedFlag("uuid", "", "A valid UUID (e.g. 123e4567-e89b-12d3-a456-426614174000)", UUIDHook)
	pflag.Parse()

	uuid, err := uuidFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid UUID:", err)
		return
	}

	fmt.Println("✅ UUID is valid:", uuid)
}
