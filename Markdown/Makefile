BINARY_NAME=md2html
BUILD_DIR=build
VERSION=0.1.0

.PHONY: all build clean run help run-compress build-all

all: build

build:
	@mkdir -p $(BUILD_DIR)
	@go build -v -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	@rm -rf $(BUILD_DIR) output

run:
	@go run -v main.go

run-compress:
	@go run -v main.go -compress

build-all: clean
	@mkdir -p $(BUILD_DIR)
	
	@GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_amd64 .
	
	@GOOS=linux GOARCH=arm64 go build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_arm64 .
	
	@GOOS=windows GOARCH=amd64 go build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_windows_amd64.exe .
	
	@GOOS=darwin GOARCH=amd64 go build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_darwin_amd64 .
	
	@GOOS=darwin GOARCH=arm64 go build -v -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_darwin_arm64 .
	
	@ls -la $(BUILD_DIR)/

help:
	@echo "Available commands:"
	@echo "  make build        - Build the program"
	@echo "  make clean        - Clean build and output files"
	@echo "  make run          - Run the program (without building)"
	@echo "  make run-compress - Run the program with HTML compression (without building)"
	@echo "  make build-all    - Cross-platform build (Linux/Windows/macOS)"
	@echo "  make help         - Show this help information" 