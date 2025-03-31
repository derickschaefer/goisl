// ip_address.go

/*
Purpose: An example of a custom hook designed to validate ip addresses using
the Go Input Sanitization Library (goisl)

Valid address examples
    go run ip_address.go --ip="192.168.1.1"
    go run ip_address.go --ip="2001:0db8:85a3:0000:0000:8a2e:0370:7334"

Invalid address examples

    go run ip_address.go --ip="192.168.1.1:443"
    go run ip_address.go --ip="2001:0db8:85a3:0000:0000:8a2e:03707334"
*/

package main

import (
	"errors"
	"fmt"
	"net"
	"strings"

	isl "github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// SanitizeIPAddressHook validates IPv4 and IPv6 addresses.
func SanitizeIPAddressHook(input string) (string, error) {
	input = strings.TrimSpace(input)
	if net.ParseIP(input) == nil {
		return "", errors.New("invalid IP address format")
	}
	return input, nil
}

func main() {
	ipFlag := isl.BindSanitizedFlag("ip", "", "IP address (IPv4 or IPv6)", SanitizeIPAddressHook)
	pflag.Parse()

	ip, err := ipFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid IP address:", err)
	} else {
		fmt.Println("✅ IP address is valid:", ip)
	}
}
