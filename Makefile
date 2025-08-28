.PHONY: build run clean test help

# Build the merge tool
build:
	@echo "Building merge tool..."
	@go build -o bin/merge cmd/merge/main.go
	@echo "Build complete! Binary available at bin/merge"

# Run the merge tool
run: build
	@echo "Running merge tool..."
	@./bin/merge configs

# Run with custom output file
run-custom: build
	@echo "Running merge tool with custom output..."
	@./bin/merge configs .coderabbit.merged.yaml

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@echo "Clean complete!"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download
	@echo "Dependencies installed!"

# Run tests
test:
	@echo "Running tests..."
	@go test ./...
	@echo "Tests complete!"

# Show help
help:
	@echo "Available commands:"
	@echo "  build       - Build the merge tool"
	@echo "  run         - Build and run the merge tool (default output: .coderabbit.yaml)"
	@echo "  run-custom  - Build and run with custom output file"
	@echo "  clean       - Clean build artifacts"
	@echo "  deps        - Install dependencies"
	@echo "  test        - Run tests"
	@echo "  help        - Show this help message"
	@echo ""
	@echo "Usage examples:"
	@echo "  make run                    # Merge configs to .coderabbit.yaml"
	@echo "  make run-custom             # Merge configs to .coderabbit.merged.yaml"
	@echo "  ./bin/merge configs         # Manual run"
	@echo "  ./bin/merge configs output.yaml  # Manual run with custom output"

# Default target
all: build
