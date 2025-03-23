.PHONY: all build run clean static tidy

APP_NAME = todoapp
GENERATE_DIR = cmd/generate
GENERATE_NAME = generate
DIST_DIR = dist
PORT = 8080

all: build

tidy:
	go mod tidy

build: tidy
	go build -o $(APP_NAME) main.go

build-generator: tidy
	cd $(GENERATE_DIR) && go build -o $(GENERATE_NAME)

run: build
	./$(APP_NAME) -port $(PORT)

static: build-generator
	cd $(GENERATE_DIR) && ./$(GENERATE_NAME)

all-in-one: build build-generator static

clean:
	rm -f $(APP_NAME)
	rm -f $(GENERATE_DIR)/$(GENERATE_NAME)
	rm -rf $(DIST_DIR)

help:
	@echo "Available commands:"
	@echo "  make tidy             - Tidy module dependencies"
	@echo "  make build            - Build server application"
	@echo "  make build-generator  - Build static generator tool"
	@echo "  make run              - Build and run server (default port 8080)"
	@echo "  make run PORT=3000    - Run server on specified port"
	@echo "  make static           - Generate static website"
	@echo "  make all-in-one       - Build all components and generate static site"
	@echo "  make clean            - Clean generated files"
	@echo "  make help             - Show this help message" 