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
	go test $(TEST_DIR) -v

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
