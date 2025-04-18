package tests

import (
	"errors"
	"os"
	"testing"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

func resetFlags() {
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
}

func TestBindSanitizedFlag_EmailSuccess(t *testing.T) {
	resetFlags()
	os.Args = []string{"cmd", "--email= user@example.com "}

	emailFlag := isl.BindSanitizedFlag("email", "", "Email input", isl.SanitizeEmailBasic)
	pflag.Parse()

	result, err := emailFlag.Get()
	if err != nil {
		t.Fatalf("Expected valid email, got error: %v", err)
	}
	if result != "user@example.com" {
		t.Errorf("Expected 'user@example.com', got '%s'", result)
	}
}

func TestBindSanitizedFlag_EmailFail(t *testing.T) {
	resetFlags()
	os.Args = []string{"cmd", "--email=invalid_email"}

	emailFlag := isl.BindSanitizedFlag("email", "", "Email input", isl.SanitizeEmailBasic)
	pflag.Parse()

	_, err := emailFlag.Get()
	if err == nil {
		t.Errorf("Expected error for invalid email, got none")
	}
}

func TestBindSanitizedFlag_EmailMustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, but did not panic")
		}
	}()

	resetFlags()
	os.Args = []string{"cmd", "--email=invalid_email"}

	emailFlag := isl.BindSanitizedFlag("email", "", "Email input", isl.SanitizeEmailBasic)
	pflag.Parse()

	_ = emailFlag.MustGet() // should panic
}

func TestBindSanitizedTextFlag_Success(t *testing.T) {
	resetFlags()
	os.Args = []string{"cmd", "--comment= Hello!!! 💡🚀"}

	commentFlag := isl.BindSanitizedTextFlag("comment", "", "Text input", nil)
	pflag.Parse()

	result := commentFlag.MustGet()
	if result != "Hello" {
		t.Errorf("Expected 'Hello', got '%s'", result)
	}
}

func TestMustGetPanicsOnError(t *testing.T) {
	// Set up a pflag set to isolate test flags
	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)

	// Bad sanitizer that always returns an error
	badSanitizer := func(input string) (string, error) {
		return "", errors.New("forced sanitizer failure")
	}

	// Bind a sanitized flag with the bad sanitizer
	f := isl.BindSanitizedFlag("failflag", "badinput", "should fail", badSanitizer)
	flags.Parse([]string{}) // Required to initialize pflag.String

	// Defer panic check
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic in MustGet but did not panic")
		}
	}()

	_ = f.MustGet() // This should panic due to forced sanitizer failure
}

func stringPointer(s string) *string {
	return &s
}
