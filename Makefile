.PHONY: run
all: clean build run
say_hello:
	@echo "Hello, world"
clean:
	@rm -rf build
build:
	@go build -o build/go-web ./cmd/go-web
run:
	@./build/go-web
run-web:
	@go run ./cmd/htmx-web/*.go
