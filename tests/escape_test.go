package tests

import (
	"net/url" // Added import for the url package
	"testing"

	"github.com/derickschaefer/goisl/pkg"
)

func TestEscapePlainText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"   Hello, World!   ", "Hello World"},
		{"Special chars: @#$%^&*()", "Special chars"},
		{"Emojis 🎉🚀😊", "Emojis"},
		{"Tabs\tand\nnewlines", "Tabs and newlines"},
		{" 123 Main St., Apt #4 ", "123 Main St Apt 4"},
		{"Symbols like ≠ and ≥ should go", "Symbols like  and  should go"},
		{"MiXeD CaSe AnD SpAcEs", "MiXeD CaSe AnD SpAcEs"},
	}

	for _, test := range tests {
		result := pkg.EscapePlainText(test.input)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestEscapeURLWithCustomHook(t *testing.T) {
	// Define a custom hook that modifies the URL without introducing additional encoding
	hook := func(parsedURL *url.URL) (*url.URL, error) {
		// Example: Change the path to "/path" without encoding
		parsedURL.Path = "/path"
		return parsedURL, nil
	}

	input := "https://example.com/originalPath"
	expected := "https://example.com/path"

	result, err := pkg.EscapeURL(input, "", hook)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if result != expected {
		t.Errorf("Input: %q, Expected: %q, Got: %q", input, expected, result)
	}
}

func TestSafeEscapeHTML(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""}, // Empty input
		{"Hello, World!", "Hello, World!"}, // No escaping needed
		{"<script>alert('XSS')</script>", "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;"},
		{"& < > \" '", "&amp; &lt; &gt; &quot; &#39;"}, // Common special characters
		{"Text with % character", "Text with % character"}, // '%' should remain untouched
		{"&lt;safe&gt;", "&amp;lt;safe&amp;gt;"}, // Double-escaped input
		{"Quotes: \"double\" 'single'", "Quotes: &quot;double&quot; &#39;single&#39;"}, // Mixed quotes
		{"Mixed text & <tags> \"inside\" 'quotes'", "Mixed text &amp; &lt;tags&gt; &quot;inside&quot; &#39;quotes&#39;"},
	}

	for _, test := range tests {
		result := pkg.SafeEscapeHTML(test.input)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}
