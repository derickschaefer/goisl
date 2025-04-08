// Example: Transliterate German Characters in Filenames with a Custom Hook
//
// This CLI-style example uses a --file flag to sanitize file names. A custom
// hook transliterates German umlauts and ß into ASCII equivalents.
//
// Usage:
//   go run german_filename.go --file="Schloßgärtenüberwachungsdienst.xlsx"
//
// Example Output:
//   ✅ Sanitized Filename: Muellerstrasse-Bericht_Uebersicht.xlsx

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	isl "github.com/derickschaefer/goisl"
)

func transliterateGermanHook(input string) (string, error) {
	replacer := strings.NewReplacer(
		"ä", "ae", "ö", "oe", "ü", "ue", "ß", "ss",
		"Ä", "Ae", "Ö", "Oe", "Ü", "Ue",
	)
	return replacer.Replace(input), nil
}

func main() {
	fileFlag := flag.String("file", "", "Filename to sanitize (e.g., Schloßgärtenüberwachungsdienst.xlsx)")
	flag.Parse()

	if *fileFlag == "" {
		fmt.Println("❌ Error: --file flag is required")
		os.Exit(1)
	}

	result, err := isl.SanitizeFileName(*fileFlag, transliterateGermanHook)
	if err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Sanitized Filename:", result)
}
