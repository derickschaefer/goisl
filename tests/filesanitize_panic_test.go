package tests

import (
	"testing"

	"github.com/derickschaefer/goisl"
)

func TestMustSanitizeFileNameBasic_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid filename, but did not panic")
		}
	}()

	// This should panic due to missing extension
	isl.MustSanitizeFileNameBasic("invalidfilename")
}
