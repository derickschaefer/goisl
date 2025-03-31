/*
Package isl provides common utility functions used throughout the goisl project.

Version: 1.1.0

File: common.go

Description:
    This file contains utility functions for checking URL protocols.
    The IsAllowedProtocol function determines whether a given URL scheme is
    present in a list of allowed protocols in a case-insensitive manner.
    This file is lightweight but essential for ensuring security and consistent
    handling of input across modules.

Change Log:
    - v1.1.0: Added pflag integration for CLI support, custom hook examples, improved validation hooks, and expanded documentation.
    - v1.0.4: Rename pkg to isl and bump version numbers
    - v1.0.3: Remove conflicting license.txt file 
    - v1.0.2: Licensing file modifications for publication
    - v1.0.1: Improved documentation and refined the IsAllowedProtocol function.
    - v1.0.0: Initial implementation of protocol checking utility.

License:
    MIT License
*/

package isl

import "strings"

// IsAllowedProtocol checks if a URL scheme is in the provided allowed list.
func IsAllowedProtocol(scheme string, allowedProtocols []string) bool {
	for _, protocol := range allowedProtocols {
		if strings.EqualFold(scheme, protocol) {
			return true
		}
	}
	return false
}
