# Variables
PROJECT_NAME := goisl
PKG := ./...
TEST_DIR := ./tests/...
BIN_DIR := ./bin
LINTER := golangci-lint

# Default target
.PHONY: all
all: build

# Build the project
.PHONY: build
build:
	go build -o $(BIN_DIR)/$(PROJECT_NAME) $(PKG)

# Run tests
.PHONY: test
test:
	@go test $(TEST_DIR) -v | sed \
		-e 's/^=== RUN\(.*\)/\x1b[34m=== RUN\1\x1b[0m/' \
		-e 's/^--- PASS\(.*\)/\x1b[32m--- PASS\1\x1b[0m/' \
		-e 's/^--- FAIL\(.*\)/\x1b[31m--- FAIL\1\x1b[0m/' \
		-e 's/^PASS\(.*\)/\x1b[32mPASS\1\x1b[0m/' \
		-e 's/^FAIL\(.*\)/\x1b[31mFAIL\1\x1b[0m/'
	@$(MAKE) examples-test

# Run linting
.PHONY: lint
lint:
	$(LINTER) run

# Clean generated files
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# Install dependencies
.PHONY: install
install:
	go mod tidy
	$(LINTER) install

# Run CLI example validations
.PHONY: examples-test

examples-test:
	@echo "🔍 Running CLI example validations..."
	@go run examples/api_key_format.go --apikey=sk-1234567890abcdef
	@go run examples/api_key_format.go --apikey=invalid-key                || echo "❌ Expected failure"
	@go run examples/slack_webhook.go --webhook=https://hooks.slack.com/services/T000/B000/XXXX
	@go run examples/slack_webhook.go --webhook=ftp://not-valid-url        || echo "❌ Expected failure"
	@go run examples/uuid_format.go --uuid=550e8400-e29b-41d4-a716-446655440000
	@go run examples/uuid_format.go --uuid=invalid-uuid                    || echo "❌ Expected failure"
	@go run examples/guid_format.go --guid=550e8400-e29b-41d4-a716-446655440000
	@go run examples/guid_format.go --guid=bad-guid                        || echo "❌ Expected failure"
	@go run examples/country_code.go --country=US
	@go run examples/country_code.go --country=U$                          || echo "❌ Expected failure"
	@go run examples/hex_token.go --token=abcdef123456
	@go run examples/hex_token.go --token=xyz123                           || echo "❌ Expected failure"
	@go run examples/mask_last4.go --secret=4111111111111234
	@go run examples/mask_last4.go --secret=123                            || echo "❌ Expected failure"
	@go run examples/twitter_handle.go --handle=@jack
	@go run examples/twitter_handle.go --handle=@@too-many-ats             || echo "❌ Expected failure"
	@go run examples/ip_address.go --ip=192.168.0.1
	@go run examples/ip_address.go --ip=999.999.999.999                    || echo "❌ Expected failure"
	@go run examples/block_shorteners.go --url=https://bit.ly/3xyzABC      || echo "❌ Expected failure"
	@go run examples/block_shorteners.go --url=https://moderncli.dev
	@go run examples/cli_flags.go --email="alice@example.com" --url="https://example.com"
	@go run examples/cli_flags.go --email="" --url=""                      || echo "❌ Expected failure"
	@go run examples/plaintext_escape.go --input="Hello, World!"
	@go run examples/plaintext_escape.go --input="<script>alert(1)</script>"
	@echo "✅ CLI example validations completed."
