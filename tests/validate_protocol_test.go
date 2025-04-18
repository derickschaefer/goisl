package tests

import (
	"testing"

	"github.com/derickschaefer/goisl"
)

func TestValidateProtocolHelper(t *testing.T) {
	valid := []string{"http", "https", "mailto"}
	invalid := "gopher"

	if !isl.IsAllowedProtocol("mailto", valid) {
		t.Error("Expected mailto to be allowed")
	}
	if isl.IsAllowedProtocol(invalid, valid) {
		t.Errorf("Expected %s to be disallowed", invalid)
	}
}
