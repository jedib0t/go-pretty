# Declare all phony targets (targets that don't create files)
.PHONY: all bench cyclo default demo-colors demo-list demo-progress demo-table fmt help lint profile test test-race tools vet

default: help

# ============================================================================
# Main targets
# ============================================================================

## all: Run all checks: tests and benchmarks
all: test bench

# ============================================================================
# Testing targets
# ============================================================================

## bench: Run benchmark tests with memory profiling
bench:
	go test -bench=. -benchmem

## test: Run tests with coverage (runs fmt, vet, lint, and cyclo first)
test: fmt vet lint cyclo
	go test -cover -coverprofile=.coverprofile ./...

## test-no-lint: Run tests with coverage (runs fmt, vet, and cyclo first)
test-no-lint: fmt vet cyclo
	go test -cover -coverprofile=.coverprofile ./...

## test-race: Run progress demo with race detector
test-race:
	go run -race ./cmd/demo-progress/demo.go

# ============================================================================
# Code quality targets
# ============================================================================

## cyclo: Check cyclomatic complexity (warns if complexity > 13)
cyclo:
	gocyclo -over 13 ./*/*.go

## fmt: Format code and organize imports
fmt:
	go fmt ./...
	gosimports -w .

## lint: Run golangci-lint static analysis
lint:
	golangci-lint run ./...

## vet: Run go vet static analysis
vet:
	go vet ./...

# ============================================================================
# Demo targets
# ============================================================================

## demo-colors: Run the colors demo
demo-colors:
	go run cmd/demo-colors/demo.go

## demo-list: Run the list demo
demo-list:
	go run cmd/demo-list/demo.go

## demo-progress: Run the progress demo
demo-progress:
	go run cmd/demo-progress/demo.go

## demo-table: Run the table demo
demo-table:
	go run cmd/demo-table/demo.go

# ============================================================================
# Utility targets
# ============================================================================

## help: Display help information for all available targets
help:
	@echo "\033[1mAvailable targets:\033[0m"
	@awk '/^# [A-Z].*targets$$/ {gsub(/^# /, ""); print "\n\033[1;36m" $$0 "\033[0m"} /^##/ {gsub(/^##\s*/, ""); idx=match($$0, /: /); if(idx) {target=substr($$0,1,idx-1); desc=substr($$0,idx+2); printf "  \033[36mmake\033[0m \033[32m%-15s\033[0m \033[33m- %s\033[0m\n", target, desc}}' $(MAKEFILE_LIST)

## profile: Run profiling script
profile:
	sh profile.sh

## tools: Install required development tools
tools: tools-ci
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.7.2

## tools-ci: Install required development tools for CI
tools-ci:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0
	go install github.com/mattn/goveralls@v0.0.12
	go install github.com/rinchsan/gosimports/cmd/gosimports@v0.3.8
