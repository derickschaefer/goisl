package pkg

// SanitizeURL sanitizes the input URL using the "display" context.
func SanitizeURL(input string) (string, error) {
	// Call EscapeURL with "display" context to apply standard escaping
	return EscapeURL(input, "display", nil)
}
