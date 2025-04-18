# API Reference (v1.1.3)

This reference documents all available functions in the `goisl` sanitization library, including new additions in v1.1.0.

---

SanitizeEmail
-------------

Description:
  Trims, validates, and sanitizes an email address. Supports optional custom hooks for domain filtering or alias stripping.

Usage:
  result, err := SanitizeEmail(input, hook)

Parameters:
  - input (string): The email address to sanitize.
  - hook (EmailHook): Optional. Function: func(local, domain string) (string, string, error)

Returns:
  - string: Sanitized email
  - error: Error if invalid or rejected by hook

Basic Wrapper:
  result, err := SanitizeEmailBasic(input)
  result := MustSanitizeEmailBasic(input)

Example:
  email, err := SanitizeEmail("  user@example.com  ", nil)

---

SanitizeFileName
-----------------

Description:
  Sanitizes a file name by removing unsafe characters, limiting length, and enforcing structure. Supports custom hooks for extension whitelists.

Usage:
  result, err := SanitizeFileName(input, hook)

Parameters:
  - input (string): Raw file name
  - hook (FileNameHook): Optional. Function: func(filename string) (string, error)

Returns:
  - string: Cleaned file name
  - error: Validation failure

Basic Wrapper:
  result, err := SanitizeFileNameBasic(input)
  result := MustSanitizeFileNameBasic(input)

---

SanitizeURL
-----------

Description:
  Validates and sanitizes a URL string. Applies optional transformation logic via hook.

Usage:
  result, err := SanitizeURL(input)

Returns:
  - string: Sanitized URL
  - error: Error if URL is invalid

Basic Wrapper:
  result, err := SanitizeURLBasic(input)
  result := MustSanitizeURLBasic(input)

---

EscapeURL
---------

Description:
  Escapes and normalizes a URL for safe use in UI or logic. Optionally allows URL mutation through a hook.

Usage:
  result, err := EscapeURL(input, context, hook)

Parameters:
  - input (string): Raw URL
  - context (string): Escaping context (e.g., "display")
  - hook (URLHook): Optional. Function: func(parsedURL *url.URL) (*url.URL, error)

Returns:
  - string: Escaped and transformed URL
  - error: Validation error or hook rejection

Example:
  url, err := EscapeURL("http://example.com?q=<script>", "display", nil)

---

HTMLSanitize
------------

Description:
  Sanitizes HTML content by removing disallowed tags and attributes.

Usage:
  result := HTMLSanitize(input, allowedTags)

Parameters:
  - input (string): Raw HTML
  - allowedTags (map[string][]string): Tag whitelist, e.g., {"b": nil, "a": {"href"}}

Returns:
  - string: Sanitized HTML

Basic Wrapper:
  HTMLSanitizeBasic(input)

Must Wrapper:
  MustHTMLSanitizeBasic(input)

---

EscapePlainText
---------------

Description:
  Escapes a string to plain, printable characters. Optional hook allows certain characters.

Usage:
  result := EscapePlainText(input, hook)

Parameters:
  - input (string): Input text
  - hook (EscapePlainTextHook): Optional. Function: func() []rune

Returns:
  - string: Escaped plain text

---

CLI Flag Bindings (pflag)
--------------------------

BindSanitizedFlag
  Binds and sanitizes a CLI flag with any compatible `SanitizeX` function.

  Usage:
    email := isl.BindSanitizedFlag("email", "", "Email", SanitizeEmailBasic)

BindSanitizedTextFlag
  Variant for plain text with `EscapePlainText`.

  Usage:
    comment := isl.BindSanitizedTextFlag("comment", "", "Comment")

Flag Access:
  - `flag.Get()` returns (value, error)
  - `flag.MustGet()` panics on error

---

IsAllowedProtocol
-----------------

Description:
  Validates that a URL scheme is in the allowlist.

Usage:
  result := IsAllowedProtocol("ftp", allowedList)

Returns:
  - bool

