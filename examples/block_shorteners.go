// Example: Block URL Shorteners with a Custom Hook (Flag-Based)
//
// This CLI-style example accepts a --url flag and checks the input against
// a list of known URL shorteners using a custom goisl hook. If the URL is
// from a disallowed domain, the program will return an error.
//
// Usage:
//   go run block_shorteners.go --url=https://bit.ly/3xyzABC
//
// Example Output:
//   ❌ Error: URL shorteners are not allowed
//
//   go run block_shorteners.go --url=https://example.com/page
//   ✅ Escaped URL: https://example.com/page

package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/derickschaefer/goisl"
)

func main() {
	// Define the input flag
	urlFlag := flag.String("url", "", "The URL to sanitize and validate")
	flag.Parse()

	if *urlFlag == "" {
		fmt.Println("❌ Error: --url flag is required")
		os.Exit(1)
	}

	// Hook that blocks known URL shortener domains
	blockShorteners := func(parsed *url.URL) (*url.URL, error) {
		shorteners := []string{"bit.ly", "tinyurl.com", "t.co"}
		host := strings.ToLower(parsed.Host)
		for _, s := range shorteners {
			if host == s {
				return nil, errors.New("URL shorteners are not allowed")
			}
		}
		return parsed, nil
	}

	// Sanitize and validate the input
	result, err := isl.EscapeURL(*urlFlag, "display", blockShorteners)
	if err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Escaped URL:", result)
}
