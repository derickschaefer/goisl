// cli_flags.go

/*
Purpose: A CLI-focused example showing how to bind sanitized input flags
using BindSanitizedFlag() and BindSanitizedTextFlag() in goisl.

Valid example:
    go run cli_flags.go --email=" derick@gmail.com " --url="https://x.com/test?x=<script>" --comment=" Nice post! 💡🚀"

Invalid examples:
    go run cli_flags.go --email="not-an-email"
    go run cli_flags.go --url="javascript:alert('xss')"
*/

package main

import (
	"fmt"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

func main() {
	// Bind sanitized input flags
	emailFlag := isl.BindSanitizedFlag("email", "", "Email address", isl.SanitizeEmailBasic)
	urlFlag := isl.BindSanitizedFlag("url", "", "URL to sanitize", isl.SanitizeURLBasic)
	commentFlag := isl.BindSanitizedTextFlag("comment", "", "Optional comment input", nil)

	// Parse flags
	pflag.Parse()

	fmt.Println("Sanitizing inputs...\n")

	// Email
	email, err := emailFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid email:", err)
	} else {
		fmt.Println("✅ Email:", email)
	}

	// URL
	url, err := urlFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid URL:", err)
	} else {
		fmt.Println("✅ URL:", url)
	}

	// Comment (sanitized plain text)
	comment := commentFlag.MustGet()
	fmt.Println("✅ Comment:", comment)
}
