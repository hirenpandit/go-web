run: clean build run
say_hello:
	@echo "Hello, world"
clean:
	@rm -rf bin 
build:
	@go build -o bin/ -v ./...
run:
	@./bin/htmx-web
air:
	@/Users/hirenpandit/go/bin/air
