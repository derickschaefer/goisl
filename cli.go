/*
Package isl provides command-line interface helpers and sanitized flag bindings
for the goisl library.

Version: 1.1.0

File: cli.go

Description:
    This file contains helper types and functions to support sanitized CLI input using the pflag package.
    It includes BindSanitizedFlag and BindSanitizedTextFlag, which simplify flag registration while ensuring
    input is properly cleaned or validated. These helpers are designed to integrate cleanly with Go CLI applications.

Change Log:
    - v1.1.0 (cli.go): Initial implementation of CLI input sanitization helpers using BindSanitizedFlag and BindSanitizedTextFlag.
    - v1.1.0 (release): Added pflag integration for CLI support, custom hook examples, improved validation hooks, and expanded documentation.

License:
    MIT License
*/

package isl

import (
	"fmt"

	"github.com/spf13/pflag"
)

// SanitizedStringFlag represents a sanitized string flag bound to pflag.
type SanitizedStringFlag struct {
	raw       *string
	sanitizer func(string) (string, error)
}

// BindSanitizedFlag binds a string flag and attaches a sanitizer function.
func BindSanitizedFlag(name, defaultValue, usage string, sanitizer func(string) (string, error)) *SanitizedStringFlag {
	raw := pflag.String(name, defaultValue, usage)
	return &SanitizedStringFlag{
		raw:       raw,
		sanitizer: sanitizer,
	}
}

// Get returns the sanitized value or an error.
func (f *SanitizedStringFlag) Get() (string, error) {
	return f.sanitizer(*f.raw)
}

// MustGet returns the sanitized value or panics on error.
func (f *SanitizedStringFlag) MustGet() string {
	val, err := f.Get()
	if err != nil {
		panic(fmt.Sprintf("Invalid flag input: %v", err))
	}
	return val
}

// SanitizedTextFlag represents a plain text flag bound to pflag and auto-sanitized.
type SanitizedTextFlag struct {
	raw   *string
	hook  EscapePlainTextHook
}

// BindSanitizedTextFlag registers a flag that will be sanitized using EscapePlainText.
// The hook argument can be nil for default sanitization.
func BindSanitizedTextFlag(name, defaultValue, usage string, hook EscapePlainTextHook) *SanitizedTextFlag {
	raw := pflag.String(name, defaultValue, usage)
	return &SanitizedTextFlag{
		raw:  raw,
		hook: hook,
	}
}

// Get returns the sanitized text value using EscapePlainText and the optional hook.
func (f *SanitizedTextFlag) Get() string {
	return EscapePlainText(*f.raw, f.hook)
}

// MustGet is an alias for Get to match other sanitized flag helpers.
func (f *SanitizedTextFlag) MustGet() string {
	return f.Get()
}
