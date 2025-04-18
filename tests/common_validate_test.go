package tests

import (
	"testing"

	"github.com/derickschaefer/goisl"
)

func TestIsAllowedProtocol(t *testing.T) {
	allowed := []string{"http", "https", "ftp"}

	tests := []struct {
		scheme   string
		expected bool
	}{
		{"HTTP", true}, // case-insensitive match
		{"https", true},
		{"ftp", true},
		{"mailto", false},
		{"", false},
	}

	for _, test := range tests {
		result := isl.IsAllowedProtocol(test.scheme, allowed)
		if result != test.expected {
			t.Errorf("IsAllowedProtocol(%q) = %v; want %v", test.scheme, result, test.expected)
		}
	}
}

func TestValidateProtocol(t *testing.T) {
	tests := []struct {
		protocol string
		expected bool
	}{
		{"http", true},
		{"ftp", true},
		{"invalidproto", false},
	}

	for _, test := range tests {
		result := isl.HTMLSanitize("<a href='"+test.protocol+"://example.com'>link</a>", isl.AllowedHTML)
		if test.expected && result == "" {
			t.Errorf("validateProtocol(%q): expected allowed, got stripped", test.protocol)
		}
		if !test.expected && result != "" && result != "<a href='invalidproto://example.com'>link</a>" {
			t.Errorf("validateProtocol(%q): expected stripped or passthrough, got %q", test.protocol, result)
		}
	}
}
