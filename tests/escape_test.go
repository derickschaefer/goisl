package tests

import (
	"net/url"
	"testing"

	"github.com/derickschaefer/goisl"
)

func TestEscapePlainText(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hook     isl.EscapePlainTextHook
	}{
		// Without hook
		{"Hello, World!", "Hello World", nil},
		{"Special chars: @#$%^&*()", "Special chars", nil}, // Fixed: removed trailing space

		// With hook allowing '@', '#', '$'
		{"Special chars: @#$%^&*()", "Special chars @#$", func() []rune { return []rune{'@', '#', '$'} }},
		{"Emojis 🎉🚀😊", "Emojis", func() []rune { return []rune{'@', '#', '$'} }},
		{"Tabs\tand\nnewlines", "Tabs and newlines", func() []rune { return []rune{} }},
		{"Keep commas, periods, and hyphens.", "Keep commas, periods, and hyphens.", func() []rune { return []rune{',', '.', '-'} }},
	}

	for _, test := range tests {
		result := isl.EscapePlainText(test.input, test.hook)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}

func TestEscapeURLWithCustomHook(t *testing.T) {
	hook := func(parsedURL *url.URL) (*url.URL, error) {
		parsedURL.Path = "/path"
		return parsedURL, nil
	}

	input := "https://example.com/originalPath"
	expected := "https://example.com/path"

	result, err := isl.EscapeURL(input, "", hook)
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
		{"", ""},
		{"Hello, World!", "Hello, World!"},
		{"<script>alert('XSS')</script>", "&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;"},
		{"& < > \" '", "&amp; &lt; &gt; &quot; &#39;"},
		{"Text with % character", "Text with % character"},
		{"&lt;safe&gt;", "&amp;lt;safe&amp;gt;"},
		{"Quotes: \"double\" 'single'", "Quotes: &quot;double&quot; &#39;single&#39;"},
		{"Mixed text & <tags> \"inside\" 'quotes'", "Mixed text &amp; &lt;tags&gt; &quot;inside&quot; &#39;quotes&#39;"},
	}

	for _, test := range tests {
		result := isl.SafeEscapeHTML(test.input)
		if result != test.expected {
			t.Errorf("Input: '%s', Expected: '%s', Got: '%s'", test.input, test.expected, result)
		}
	}
}
