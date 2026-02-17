# Justfile for katapultpro-go-sdk
# See https://github.com/casey/just
# SDK lives in versioned modules (v3, future v4). Use go.work from repo root.

# Default recipe: list available commands
default:
    @just --list

# Build all versioned modules (v3, ...)
build:
    go build ./v3/...

# Run all tests
test:
    go test ./v3/...

# Run tests with verbose output
test-verbose:
    go test -v ./v3/...

# Run tests with race detector
test-race:
    go test -race ./v3/...

# Run tests with coverage
test-cover:
    go test -cover ./v3/...

# Run tests with coverage report (outputs coverage to coverage.out)
test-cover-profile:
    go test -coverprofile=coverage.out ./v3/...
    go tool cover -func=coverage.out

# Run only short tests (skip long-running)
test-short:
    go test -short ./v3/...

# Vet the codebase
vet:
    go vet ./v3/...

# Format code (v3 and any other version dirs)
fmt:
    gofmt -s -w ./v3

# Run go mod tidy in each versioned module
tidy:
    cd v3 && go mod tidy

# Full check: build, vet, test
check: build vet test
    @echo "All checks passed."

# CI: tidy, build, vet, test with race
ci: tidy build vet test-race
    @echo "CI checks passed."
