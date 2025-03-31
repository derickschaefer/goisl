// country_code.go

/*
Purpose: An example of a custom hook designed to validate ISO 3166-1 alpha-2 country codes
(e.g., "US", "DE", "JP") using the Go Input Sanitization Library (goisl).

Valid input examples:
    go run country_code.go --country="US"
    go run country_code.go --country="de"
    go run country_code.go --country="JP"

Invalid input examples:
    go run country_code.go --country="USA"      // too long
    go run country_code.go --country="u$"       // invalid characters
    go run country_code.go --country="XX"       // unsupported/unknown code
*/

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// Example whitelist of allowed ISO 3166-1 alpha-2 codes
var allowedCountryCodes = map[string]bool{
	"US": true,
	"DE": true,
	"JP": true,
	"FR": true,
	"IN": true,
	"BR": true,
	"CN": true,
	"GB": true,
}

func countryCodeHook(input string) (string, error) {
	code := strings.ToUpper(strings.TrimSpace(input))
	if len(code) != 2 {
		return "", errors.New("country code must be exactly two characters")
	}
	if !allowedCountryCodes[code] {
		return "", fmt.Errorf("unsupported country code: %s", code)
	}
	return code, nil
}

func main() {
	countryFlag := isl.BindSanitizedFlag("country", "", "Country code (ISO 3166-1 alpha-2)", countryCodeHook)
	pflag.Parse()

	code, err := countryFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid country code:", err)
	} else {
		fmt.Println("✅ Country code is valid:", code)
	}
}
