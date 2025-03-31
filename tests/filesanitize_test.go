// filesanitize_test.go
package tests

import (
	"errors"
	"strings"
	"testing"

	"github.com/derickschaefer/goisl"
)

// FileTypeRestrictionHook creates a hook that restricts file types based on allowed extensions.
// It appends an underscore to disallowed intermediate extensions and ensures the final extension is allowed.
func FileTypeRestrictionHook(allowedExtensions []string) isl.FileNameHook {
	extensionSet := make(map[string]bool)
	for _, ext := range allowedExtensions {
		extensionSet[strings.ToLower(ext)] = true
	}

	return func(filename string) (string, error) {
		parts := strings.Split(filename, ".")
		if len(parts) < 2 {
			return "", errors.New("file name must contain an extension")
		}
		base := parts[0]
		extensions := parts[1:]

		// Process intermediate extensions
		for i, ext := range extensions[:len(extensions)-1] { // Exclude the final extension
			lowerExt := strings.ToLower(ext)
			if !extensionSet["."+lowerExt] {
				extensions[i] = ext + "_"
			}
		}

		// Check final extension
		finalExt := "." + strings.ToLower(extensions[len(extensions)-1])
		if !extensionSet[finalExt] {
			return "", errors.New("final file type not allowed")
		}

		// Rejoin the filename
		newFilename := base
		for _, ext := range extensions {
			newFilename += "." + ext
		}
		return newFilename, nil
	}
}

func TestSanitizeFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		// Basic Valid Filename
		{"  example  .txt  ", "example.txt", true},

		// Replace '+' with '-'
		{"file+name.txt", "file-name.txt", true},

		// Remove Special Characters
		{"example#@!.txt", "example.txt", true},

		// Remove Hyphens before Dot
		{"file-name-.txt", "file-name.txt", true}, // Adjusted expectation

		// Replace Spaces with Hyphens
		{"my file name.docx", "my-file-name.docx", true},

		// Multiple Hyphens and Dots
		{"my--file..name..pdf", "my-file.name.pdf", true}, // Expected to pass after sanitization

		// Trim Leading and Trailing Characters
		{"---myfile.txt---", "myfile.txt", true},

		// Prevent Directory Traversal
		{"../../etc/passwd", "", false},

		// Empty Filename After Trimming
		{"   ", "", false},

		// Missing Extension
		{"filename", "", false},

		// Exceeding Maximum Length (255 characters)
		{strings.Repeat("a", 251) + ".txt", strings.Repeat("a", 251) + ".txt", true}, // 255 chars
		{strings.Repeat("a", 252) + ".txt", "", false},                             // 256 chars
	}

	for _, test := range tests {
		result, err := isl.SanitizeFileName(test.input, nil)
		if test.isValid && err != nil {
			t.Errorf("Input: '%s', Expected valid file name, got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: '%s', Expected error, got result: '%s'", test.input, result)
		}
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestSanitizeFileNameWithCustomHook(t *testing.T) {
	allowedExtensions := []string{".txt", ".js", ".docx", ".pdf"}
	customHook := FileTypeRestrictionHook(allowedExtensions)

	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		// Allowed File Types
		{"valid-file.txt", "valid-file.txt", true},
		{"script.js", "script.js", true},
		{"document.docx", "document.docx", true},
		{"report.pdf", "report.pdf", true},

		// Disallowed File Types
		{"image.png", "", false},
		{"archive.zip", "", false},
		{"executable.exe", "", false},
		{"style.css", "", false},

		// Mixed Validity: Intermediate disallowed extension
		{"bad-file.exe.js", "bad-file.exe_.js", true}, // Adjusted isValid to true
	}

	for _, test := range tests {
		result, err := isl.SanitizeFileName(test.input, customHook)
		if test.isValid && err != nil {
			t.Errorf("Input: '%s', Expected valid file name, got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: '%s', Expected error, got result: '%s'", test.input, result)
		}
		if test.isValid && result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
		if !test.isValid && test.expected != "" && result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestSanitizeFileNameAdditional(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		// Multiple Special Characters
		{"my@#file!!.docx", "myfile.docx", true},

		// Already Sanitized Filenames
		{"valid-file.pdf", "valid-file.pdf", true},

		// Complex Multiple Extensions (Assuming "tar" and "gz" are allowed)
		{"archive.tar.gz", "archive.tar.gz", true},

		// Filenames with Spaces and Special Characters
		{"  my file @2023!!.xlsx  ", "my-file-2023.xlsx", true},

		// Directory Traversal Attempts
		{"../../etc/passwd", "", false},

		// Exceeding Maximum Length (256 'a's + ".txt" = 259 characters)
		{strings.Repeat("a", 256) + ".txt", "", false},

		// Filenames with Unicode Characters
		{"café-document.pdf", "cafe-document.pdf", true},

		// Filenames with Backticks
		{"file`name`.txt", "filename.txt", true},
	}

	for _, test := range tests {
		result, err := isl.SanitizeFileName(test.input, nil)
		if test.isValid && err != nil {
			t.Errorf("Input: '%s', Expected valid file name, got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: '%s', Expected error, got result: '%s'", test.input, result)
		}
		if test.isValid && result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
		if !test.isValid && test.expected != "" && result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestSanitizeFileNameOnlyDisallowedExtensions(t *testing.T) {
	allowedExtensions := []string{".txt", ".js", ".docx", ".pdf"}
	customHook := FileTypeRestrictionHook(allowedExtensions)

	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"dangerous.exe", "", false},
		{"malware.bat", "", false},
	}

	for _, test := range tests {
		result, err := isl.SanitizeFileName(test.input, customHook)
		if test.isValid && err != nil {
			t.Errorf("Input: '%s', Expected valid file name, got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: '%s', Expected error, got result: '%s'", test.input, result)
		}
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestSanitizeFileNameComplexPatterns(t *testing.T) {
	allowedExtensions := []string{".txt", ".js", ".docx", ".pdf"}
	customHook := FileTypeRestrictionHook(allowedExtensions)

	tests := []struct {
		input    string
		expected string
		isValid  bool
	}{
		{"complex---name__with@@@symbols...pdf", "complex-name_withsymbols.pdf", true},
		{"../secret/.env", "", false},
	}

	for _, test := range tests {
		result, err := isl.SanitizeFileName(test.input, customHook)
		if test.isValid && err != nil {
			t.Errorf("Input: '%s', Expected valid file name, got error: %v", test.input, err)
		}
		if !test.isValid && err == nil {
			t.Errorf("Input: '%s', Expected error, got result: '%s'", test.input, result)
		}
		if test.isValid && result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
		if !test.isValid && test.expected != "" && result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestSanitizeFileNameBasic(t *testing.T) {
	input := "  good_file--name..txt "
	expected := "good_file-name.txt"

	result, err := isl.SanitizeFileNameBasic(input)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if result != expected {
		t.Errorf("Expected: '%s', Got: '%s'", expected, result)
	}
}

func TestMustSanitizeFileNameBasic(t *testing.T) {
	// Valid file name (should not panic)
	valid := "clean-name.pdf"
	expected := "clean-name.pdf"
	result := isl.MustSanitizeFileNameBasic(valid)
	if result != expected {
		t.Errorf("Expected: '%s', Got: '%s'", expected, result)
	}

	// Invalid file name (should panic)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid file name, but did not panic")
		}
	}()

	_ = isl.MustSanitizeFileNameBasic("..") // should trigger panic
}
