# Makefile for CCIP v1.7 Protocol Development

# Go parameters
GOCMD = go
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOMOD = $(GOCMD) mod
GOFMT = gofumpt
GOLINT = golangci-lint

# Project parameters
PROJECT_NAME = chainlink-ccv-protocol

# Test flags
TEST_FLAGS = -v -race -timeout=30s

# Default target
.PHONY: help
help: ## Display this help message
	@echo "CCIP v1.7 Protocol Development Commands"
	@echo "======================================"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development targets
.PHONY: dev
dev: clean format lint test ## Run full development cycle (clean, format, lint, test)

.PHONY: install-tools
install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/vektra/mockery/v2@latest
	@echo "Development tools installed successfully"

# Testing
.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	@$(GOTEST) $(TEST_FLAGS) ./...

# Mocking
.PHONY: mock
mock: ## Delete and regenerate all mocks using mockery
	@echo "Cleaning existing mocks..."
	@find . -path "*/mocks/*.go" -delete 2>/dev/null || true
	@echo "Generating mocks with mockery..."
	@if command -v mockery >/dev/null 2>&1; then \
		mockery; \
	else \
		echo "mockery not found. Install with: go install github.com/vektra/mockery/v2@latest"; \
		exit 1; \
	fi
	@echo "Mocks regenerated successfully"

# Code quality
.PHONY: format
format: ## Format code using gofumpt (or gofmt as fallback)
	@echo "Formatting code..."
	@if command -v $(GOFMT) >/dev/null 2>&1; then \
		$(GOFMT) -l -w .; \
	else \
		echo "gofumpt not found, using gofmt instead..."; \
		$(GOCMD) fmt ./...; \
	fi
	@$(GOCMD) mod tidy

.PHONY: format-check
format-check: ## Check if code is formatted
	@echo "Checking code formatting..."
	@$(GOFMT) -l . | tee /tmp/gofmt-output
	@test ! -s /tmp/gofmt-output || (echo "Code is not formatted. Run 'make format' to fix." && exit 1)

.PHONY: lint
lint: ## Run linter (if available)
	@echo "Running linter..."
	@if command -v $(GOLINT) >/dev/null 2>&1; then \
		$(GOLINT) run ./...; \
	else \
		echo "golangci-lint not found, skipping linting. Install with: make install-tools"; \
		echo "Running go vet instead..."; \
		$(GOCMD) vet ./...; \
	fi

.PHONY: lint-fix
lint-fix: ## Run linter with auto-fix
	@echo "Running linter with auto-fix..."
	@$(GOLINT) run --fix ./...

.PHONY: vet
vet: ## Run go vet
	@echo "Running go vet..."
	@$(GOCMD) vet ./...

.PHONY: check
check: format-check lint vet ## Run all code quality checks

# Dependencies
.PHONY: deps
deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@$(GOMOD) download

.PHONY: deps-update
deps-update: ## Update dependencies
	@echo "Updating dependencies..."
	@$(GOMOD) get -u ./...
	@$(GOMOD) tidy

.PHONY: deps-verify
deps-verify: ## Verify dependencies
	@echo "Verifying dependencies..."
	@$(GOMOD) verify

.PHONY: deps-clean
deps-clean: ## Clean module cache
	@echo "Cleaning module cache..."
	@$(GOCMD) clean -modcache

# Cleaning
.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@$(GOCLEAN)
	@rm -f *.out
	@rm -f *.test

.PHONY: clean-all
clean-all: clean deps-clean ## Clean everything including module cache

# Quick development workflow
.PHONY: quick
quick: format test ## Quick development check (format + test)

.PHONY: full
full: clean format lint test ## Full development check

# Show current status
.PHONY: status
status: ## Show project status
	@echo "Project: $(PROJECT_NAME)"
	@echo "Packages: $(shell $(GOCMD) list ./... | wc -l)"
	@echo "Go Version: $(shell $(GOCMD) version | cut -d' ' -f3)"
	@echo "Test files: $(shell find . -name '*_test.go' | wc -l)"
	@echo "Lines of code: $(shell find . -name '*.go' -not -path './vendor/*' | xargs wc -l | tail -1 | awk '{print $$1}')"
