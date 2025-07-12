.PHONY: run test build clean help

## run: 🚀 Run application
run:
	@echo "Starting application..."
	go run ./cmd/app

## test: 🧪 Run tests
test:
	@echo "Running tests..."
	go test -v -cover ./...

## build: 🔨 Build binary
build:
	@echo "Building application..."
	go build -o bin/app ./cmd/app

## clean: 🧹 Clean artifacts
clean:
	@echo "Cleaning up..."
	rm -r -force bin
	go clean

## help: ℹ️ Show help
help:
	@echo "Available commands:"
	@grep -E '^## [a-zA-Z_-]+:' Makefile | cut -d " " -f 2- | awk 'BEGIN {FS = ":"}; {printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2}'