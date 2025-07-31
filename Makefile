BINARY     = utify
BUILD_DIR  = bin
SRC_DIR    = .
TEST_DIR   = .
GOFLAGS    = -mod=readonly
LDFLAGS    = -s -w

# OS-specific configuration
OS := $(shell uname -s)

ifeq ($(OS), Darwin)
	SED_I = sed -i ''
else
	SED_I = sed -i
endif

# Linter setup
LINTER = golangci-lint run

# Build the binary
build:
	@echo "🔨 Building $(BINARY)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY) -ldflags "$(LDFLAGS)" $(SRC_DIR)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY)"

# Build and run
run: build
	@echo "🚀 Running $(BINARY)..."
	@./$(BUILD_DIR)/$(BINARY)

# Run tests with verbose output
test:
	@echo "🧪 Running tests..."
	@go test -v ./tests/... ./pkg/... ./internal/...

# Run unit tests only
test-unit:
	@echo "🧪 Running unit tests..."
	@go test -v ./pkg/*/tests.go ./pkg/*/tests/*.go

# Run integration tests only
test-integration:
	@echo "🧪 Running integration tests..."
	@go test -v ./tests/integration/...

# Run benchmarks
bench:
	@echo "⚡ Running benchmarks..."
	@go test -bench=. ./tests/benchmarks/...

# Run tests with coverage and generate report
coverage:
	@echo "📊 Running tests with coverage..."
	@go test -covermode=atomic -coverprofile=coverage.out $$(go list ./... | grep -v /examples)
	@go tool cover -func=coverage.out

# Generate HTML coverage report
coverage-html:
	@echo "📊 Generating HTML coverage report..."
	@go test -covermode=atomic -coverprofile=coverage.out $$(go list ./... | grep -v /examples)
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

# Clean up build and coverage files
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR) coverage.out
	@echo "✅ Cleanup complete."

# Run linters
lint:
	@echo "🔍 Running linters..."
	@$(LINTER)

# Validate documentation and code style
docs:
	@echo "📚 Validating documentation and static checks..."
	@revive -config revive.toml ./...
	@go vet ./...
	@echo "✅ Docs and vet checks passed."

# Display available commands
help:
	@echo "📌 Available commands:"
	@echo "  make build      - Build the binary"
	@echo "  make run        - Build and run the application"
	@echo "  make test       - Run all tests"
	@echo "  make test-unit  - Run unit tests only"
	@echo "  make test-integration - Run integration tests only"
	@echo "  make bench      - Run benchmarks"
	@echo "  make coverage   - Run tests with coverage"
	@echo "  make coverage-html - Generate HTML coverage report"
	@echo "  make lint       - Run linters (requires golangci-lint)"
	@echo "  make docs       - Run revive and go vet to validate docs"
	@echo "  make clean      - Remove generated files"

