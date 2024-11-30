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
	#go test $(TEST_DIR) -v
	@go test $(TEST_DIR) -v | sed \
		-e 's/^=== RUN\(.*\)/\x1b[34m=== RUN\1\x1b[0m/' \
		-e 's/^--- PASS\(.*\)/\x1b[32m--- PASS\1\x1b[0m/' \
		-e 's/^--- FAIL\(.*\)/\x1b[31m--- FAIL\1\x1b[0m/' \
		-e 's/^PASS\(.*\)/\x1b[32mPASS\1\x1b[0m/' \
		-e 's/^FAIL\(.*\)/\x1b[31mFAIL\1\x1b[0m/'

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
