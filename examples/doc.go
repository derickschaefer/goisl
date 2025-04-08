// Package main contains standalone example programs that demonstrate how to use
// the goisl (Go Input Sanitization Library) package.
//
// These CLI-style examples show how to sanitize and validate various types of input,
// such as email addresses, URLs, file names, UUIDs, and plain text. Most examples
// showcase goisl's custom hook capability for domain-specific input logic,
// including format validation, character transliteration, and filtering.
//
// Each example is designed to be executed directly using `go run` and is built for
// educational use, testing, or integration into real-world CLI tools.
//
// See README.md in this directory for a full list of available examples.
//
// ---
//
// Example: Transliterate German Characters in Filenames with a Custom Hook
//
// This CLI-style example accepts a --file flag and uses goisl.SanitizeFileName
// with a custom hook to transliterate German umlauts and the ß character into
// ASCII-friendly equivalents.
//
// Usage:
//
//   go run german_filename.go --file="Schloßgärtenüberwachungsdienst.xlsx"
//
// Example Output:
//
//   ✅ Sanitized Filename: Schlossgaertenueberwachungsdienst.xlsx
package main
