# goisl Examples

This directory contains practical, copy/paste-ready examples demonstrating the use of the `goisl` input sanitization and escaping library.

Each file is self-contained and includes a helpful comment block describing its purpose, along with valid and invalid input examples.

---

## 🔧 Core Functionality

These examples show basic usage of `goisl` with default sanitization and escaping:

- [`plaintext_escape.go`](./plaintext_escape.go): Escape plain text while preserving only safe characters.
- [`cli_flags.go`](./cli_flags.go): Bind input flags using `pflag` and sanitize values using `BindSanitizedFlag` and `BindSanitizedTextFlag`.

---

## 🧩 Custom Override Hooks (Real-World Formats)

These examples demonstrate how to create custom sanitization hooks for specific formats:

- [`uuid_format.go`](./uuid_format.go): Require input to match a valid UUID pattern.
- [`twitter_handle.go`](./twitter_handle.go): Validate Twitter handles (e.g., `@username`).
- [`ip_address.go`](./ip_address.go): Validate and sanitize IPv4 and IPv6 addresses.
- [`hex_token.go`](./hex_token.go): Validate secure hex tokens (e.g., for email confirmations).
- [`country_code.go`](./country_code.go): Restrict ISO 3166-1 country codes (e.g., "US", "DE").
- [`api_key_format.go`](./api_key_format.go): Validate test only API Stripe and prohibit production or invalid keys.
- [`slack_webhook.go`](./slack_webhook.go): Validate Slack webhook path and format.
- [`mask_last4.go`](./mast_last4.go): Uses isl to mask all but the last for digits of a sensative string (e.g. logging purposes).
- [`block_shortener.go`](./block_shortener.go): Disallows a URL input that is a known URL shortener.
- [`crypto_btc_address.go`](./crypto_btc_address.go): Validates Bitcoin addresses starting with `1`, `3`, or `bc1`.
- [`guid_format.go`](./guid_format.go): Validates GUID-style identifiers using a custom format check.

---

## ▶️ How to Run

The header comments in each example file contains testing examples. Use the Go CLI to test an example with your own inputs:

```bash
go run twitter_handle.go --handle="@derick"
go run ip_address.go --ip="192.168.1.1"
go run uuid_format.go --uuid="123e4567-e89b-12d3-a456-426614174000"
