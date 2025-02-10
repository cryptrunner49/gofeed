# Variables
BINARY_NAME=gofeed
SRC_DIR=./src
BUILD_DIR=./bin

# Default target
.PHONY: all
all: test build

# Build the Go project
.PHONY: build
build:
	@echo "Building the project..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Run the Go project
.PHONY: run
run: build
	@echo "Running the project..."
	@$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

# Display help
.PHONY: help
help:
	@echo "Makefile for Go project"
	@echo
	@echo "Usage:"
	@echo "  make <target>"
	@echo
	@echo "Targets:"
	@echo "  all     - Run tests and build the project"
	@echo "  build   - Build the project"
	@echo "  run     - Run the project"
	@echo "  test    - Run tests"
	@echo "  clean   - Clean build artifacts"
	@echo "  help    - Display this help message"
