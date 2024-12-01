# Variables
PACKAGE = ./...

# Default target
all: test benchmark

# Run all unit tests recursively
test:
	@echo "Running unit tests recursively..."
	@go test $(PACKAGE) -v

# Run all benchmark tests recursively
benchmark:
	@echo "Running benchmark tests recursively..."
	@go test $(PACKAGE) -bench=. -benchmem

# Clean the generated files (optional)
clean:
	@echo "Cleaning up..."
	@rm -rf ./bin ./coverage.out

# Help target to display usage
help:
	@echo "Makefile targets:"
	@echo "  all        - Run tests and benchmarks recursively"
	@echo "  test       - Run all unit tests recursively"
	@echo "  benchmark  - Run all benchmark tests recursively"
	@echo "  clean      - Clean generated files"