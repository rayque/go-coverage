.PHONY: build test clean install coverage example help

# Binary name
BINARY_NAME=go-coverage

# Build directory
BUILD_DIR=bin

help: ## Show this help
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the CLI binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/go-coverage
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

install: ## Install the CLI binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@go install ./cmd/go-coverage
	@echo "$(BINARY_NAME) installed successfully!"

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

coverage: ## Generate coverage report
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

coverage-html: build ## Generate HTML coverage report using the built tool
	@echo "Generating HTML coverage report..."
	@go test -coverprofile=coverage.out ./...
	@./$(BUILD_DIR)/$(BINARY_NAME) -input=coverage.out -output=coverage.html
	@echo "Coverage report generated: coverage.html"

example: build ## Run the example
	@echo "Running example..."
	@cd examples/basic && go run main.go

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

lint: fmt vet ## Run formatters and linters

all: clean lint test build ## Run all checks and build

.DEFAULT_GOAL := help

