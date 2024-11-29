package pkg

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
