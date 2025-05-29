# Name of the binary
BINARY_NAME=snippetbox

# Go entry point
MAIN=cmd/web/main.go

# Default target
all: generate build run

# Generate Templ files
generate:
	@echo "🔧 Generating Templ components..."
	templ generate

# Build the Go binary
build:
	@echo "🏗️  Building Go app..."
	go build -o $(BINARY_NAME) ./cmd/web

# Run the server
run: build
	@echo "🚀 Running server..."
	./$(BINARY_NAME)

# Clean generated files and binary
clean:
	@echo "🧹 Cleaning up..."
	go clean
	rm -f $(BINARY_NAME)

# Rebuild everything from scratch
rebuild: clean all
