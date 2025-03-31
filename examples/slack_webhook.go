// slack_webhook.go

/*
Purpose: An example of a custom hook designed to validate Slack webhook URLs using
the Go Input Sanitization Library (goisl). This protects against incorrect formats and
ensures only properly structured webhooks are accepted.

✅ Valid input examples:
    go run slack_webhook.go --webhook="https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
    go run slack_webhook.go --webhook="https://hooks.slack.com/services/ABCD1234/EFGH5678/IJKL9012mnop3456"

❌ Invalid input examples:
    go run slack_webhook.go --webhook="http://hooks.slack.com/services/invalid"      // Not HTTPS
    go run slack_webhook.go --webhook="https://example.com/services/123"             // Wrong host
    go run slack_webhook.go --webhook="https://hooks.slack.com/bad/format"           // Incorrect path
*/

package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/derickschaefer/goisl"
	"github.com/spf13/pflag"
)

// Slack webhook pattern: must match Slack's official format
var slackWebhookPattern = regexp.MustCompile(`^https://hooks\.slack\.com/services/[A-Za-z0-9]+/[A-Za-z0-9]+/[A-Za-z0-9]+$`)

func slackWebhookHook(input string) (string, error) {
	url := strings.TrimSpace(input)

	if !strings.HasPrefix(url, "https://hooks.slack.com/services/") {
		return "", errors.New("webhook must begin with 'https://hooks.slack.com/services/'")
	}

	if !slackWebhookPattern.MatchString(url) {
		return "", errors.New("invalid Slack webhook format")
	}

	return url, nil
}

func main() {
	webhookFlag := isl.BindSanitizedFlag("webhook", "", "Slack webhook URL to validate", slackWebhookHook)
	pflag.Parse()

	webhook, err := webhookFlag.Get()
	if err != nil {
		fmt.Println("❌ Invalid webhook:", err)
	} else {
		fmt.Println("✅ Slack webhook is valid:", webhook)
	}
}
