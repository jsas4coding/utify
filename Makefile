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
	@go test -v ./...

# Run tests with coverage and generate report
coverage:
	@echo "📊 Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out

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
	@echo "  make test       - Run tests"
	@echo "  make coverage   - Run tests with coverage"
	@echo "  make lint       - Run linters (requires golangci-lint)"
	@echo "  make docs       - Run revive and go vet to validate docs"
	@echo "  make clean      - Remove generated files"

