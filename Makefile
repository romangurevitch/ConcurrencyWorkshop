# Makefile for Golang Project

# Project Variables
PROJECT_NAME := myproject
GO := go
GOFMT := gofmt -s
GOFILES := $(shell find . -name "*.go" -type f)
GOTEST := $(GO) test
GOBENCH := $(GO) test -run=^$$ -bench .
GOLANGCI_LINT_VERSION := v1.61.0
GOLANGCI_LINT_FILE := bin/golangci-lint
GOLANGCI_LINT_VERSIONED := $(GOLANGCI_LINT_FILE)-$(GOLANGCI_LINT_VERSION)
GOLINT := $(GOLANGCI_LINT_VERSIONED) run

# Phony Targets
.PHONY: help build test bench lint fmt tidy setup clean update-deps fmt

# Help
help:
	@echo "Choose a command to run in $(PROJECT_NAME):"
	@echo "  build"
	@echo "    Compile the project into a binary."
	@echo "  test"
	@echo "    Run tests on a compiled project. This includes different levels of tests for various components."
	@echo "  bench"
	@echo "    Run benchmarks on a compiled project. This includes benchmarks for concurrency and pattern components."
	@echo "  lint"
	@echo "    Run linters on the project using a specific version of golangci-lint."
	@echo "  fmt"
	@echo "    Format the Go source files."
	@echo "  tidy"
	@echo "    Remove unused dependencies from the project."
	@echo "  setup"
	@echo "    Set up the necessary tools for linting."
	@echo "  clean"
	@echo "    Remove temporary files and compiled binaries."
	@echo "  update-deps"
	@echo "    Update all dependencies to their latest versions and tidy up the project."
	@echo "  fixme-basic"
	@echo "    Run basic tests for the 'fixme' challenge."
	@echo "  fixme-intermediate"
	@echo "    Run intermediate tests for the 'fixme' challenge."
	@echo "  fixme-advanced"
	@echo "    Run advanced tests for the 'fixme' challenge."
	@echo "  implme-basic"
	@echo "    Run basic tests and benchmarks for the 'implme' basic challenge."
	@echo "  implme-intermediate"
	@echo "    Run intermediate tests for the 'implme' intermediate challenge."
	@echo "  implme-advanced"
	@echo "    Run advanced tests for the 'implme' advanced challenge."

# Test Targets for different components
test:
	@echo "Testing concurrency components..."
	@$(GOTEST) ./internal/concurrency/...
	@echo "Testing pattern components..."
	@$(GOTEST) ./internal/pattern/...

# Test Targets for different components
fixme-basic:
	@echo "Running basic tests..."
	@$(GOTEST) -v -run=^TestBasic ./internal/challenge/fixme/...

# Test Targets for different components
fixme-intermediate:
	@echo "Running intermediate tests..."
	@$(GOTEST) -v -run=^TestIntermediate ./internal/challenge/fixme/...

# Test Targets for different components
fixme-advanced:
	@echo "Running advanced tests..."
	@$(GOTEST) -v -run=^TestAdvanced ./internal/challenge/fixme/...

# Test Targets for different components
implme-basic:
	@echo "Running basic tests..."
	@$(GOTEST) -v ./internal/challenge/implme/basic...
	@echo "Running basic benchmarks..."
	@$(GOBENCH) ./internal/challenge/implme/basic...

# Test Targets for different components
implme-intermediate:
	@echo "Running intermediate tests..."
	@$(GOTEST) -v ./internal/challenge/implme/intermediate...

# Test Targets for different components
implme-advanced:
	@echo "Running advanced tests..."
	@$(GOTEST) -v ./internal/challenge/implme/advanced...

# Benchmark Targets for different components
bench:
	@echo "Benchmarking concurrency components..."
	@$(GOBENCH) ./internal/concurrency/...
	@echo "Benchmarking pattern components..."
	@$(GOBENCH) ./internal/pattern/...

# Linting
lint: $(GOLANGCI_LINT_VERSIONED)
	@echo "Linting the code..."
	@$(GOLINT)

# Formatting Go source files
fmt:
	@echo "Formatting Go source files..."
	@$(GOFMT) -w $(GOFILES)

# Go mod tidy
tidy: fmt
	@echo "Tidying up the go.mod and go.sum files..."
	@$(GO) mod tidy

# Update dependencies and tidy
update-deps:
	@echo "Updating all dependencies to their latest versions..."
	@$(GO) get -u ./...
	@echo "Tidying up the go.mod and go.sum files..."
	@$(GO) mod tidy

# Setup
$(GOLANGCI_LINT_VERSIONED):
	@echo "Setting up..."
	@mkdir -p bin
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin $(GOLANGCI_LINT_VERSION)
	@mv $(GOLANGCI_LINT_FILE) $(GOLANGCI_LINT_VERSIONED)

# Cleaning up
clean:
	@echo "Cleaning up..."
	@rm -f $(PROJECT_NAME)
	@rm -f $(GOLANGCI_LINT_VERSIONED)
	@rm -f coverage.out
