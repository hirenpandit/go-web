.PHONY: run
all: clean build run
say_hello:
	@echo "Hello, world"
clean:
	@rm -rf build
build:
	@go build -o build/ -v ./...
run:
	@./build/htmx-web
run-dev:
	@go run ./cmd/htmx-web/*.go
