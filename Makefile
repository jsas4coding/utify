BINARY=utify
BUILD_DIR=bin
SRC_DIR=.
TEST_DIR=.
GOFLAGS=-mod=readonly
LDFLAGS=-s -w

OS := $(shell uname -s)

ifeq ($(OS), Darwin)
    SED_I=sed -i ''
else
    SED_I=sed -i
endif

LINTER=golangci-lint run

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
	@echo "📊 Running tests with coverage..."
	@go test -cover ./... | tee coverage.out

clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BUILD_DIR) coverage.out
	@echo "✅ Cleanup complete."

lint:
	@echo "🔍 Running linters..."
	@$(LINTER)

help:
	@echo "📌 Available commands:"
	@echo "  make build      - Build the binary"
	@echo "  make run        - Build and run the application"
	@echo "  make test       - Run tests"
	@echo "  make coverage   - Run tests with coverage"
	@echo "  make lint       - Run linters (requires golangci-lint)"
	@echo "  make clean      - Remove generated files"
