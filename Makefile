BINARY = utify
BUILD_DIR = bin
SRC_DIR = .
TEST_DIR = .
GOFLAGS = -mod=readonly
LDFLAGS = -s -w

OS := $(shell uname -s)

ifeq ($(OS), Darwin)
    SED_I = sed -i ''
else
    SED_I = sed -i
endif

LINTER = golangci-lint run

build:
	@echo "🔨 Building $(BINARY)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY) -ldflags "$(LDFLAGS)" $(SRC_DIR)
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY)"

run: build
	@echo "🚀 Running $(BINARY)..."
	@./$(BUILD_DIR)/$(BINARY)

test:
	@echo "🧪 Running tests..."
	@go test -v ./...

coverage:
	@echo "📊 Running tests with coverage (profile: coverage.out)..."
	@go test -coverprofile=coverage.out ./...

coverage-html: coverage
	@echo "🌐 Opening HTML coverage report..."
	@go tool cover -html=coverage.out

lint:
	@echo "🔍 Running linters..."
	@$(LINTER)

check: lint coverage
	@echo "✅ All checks passed."

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR) coverage.out
	@echo "✅ Cleanup complete."

help:
	@echo "📌 Available commands:"
	@echo "  make build          - Build the binary"
	@echo "  make run            - Build and run the application"
	@echo "  make test           - Run tests"
	@echo "  make coverage       - Run tests with coverage (to coverage.out)"
	@echo "  make coverage-html  - Open HTML report for coverage"
	@echo "  make lint           - Run linters (requires golangci-lint)"
	@echo "  make check          - Run lint and coverage (CI)"
	@echo "  make clean          - Remove generated files"

