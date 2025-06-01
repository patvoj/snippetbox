# Name of the binary
BINARY_NAME=snippetbox
BIN_DIR=bin
BINARY_PATH=$(BIN_DIR)/$(BINARY_NAME)

# Go entry point
MAIN=cmd/web/main.go

# Default target
all: generate build run

# Generate Templ files
generate:
	@echo "🔧 Generating Templ components..."
	templ generate

# Build the Go binary
build: | $(BIN_DIR)
	@echo "🏗️  Building Go app..."
	go build -o $(BINARY_PATH) ./cmd/web

# Ensure bin directory exists
$(BIN_DIR):
	@mkdir -p $(BIN_DIR)

# Run the server
run: build
	@echo "🚀 Running server..."
	./$(BINARY_PATH)

# Clean generated files and binary
clean:
	@echo "🧹 Cleaning up..."
	go clean
	rm -rf $(BIN_DIR)

# Rebuild everything from scratch
rebuild: clean all
