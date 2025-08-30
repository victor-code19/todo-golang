.PHONY: test test-unit test-integration test-all test-coverage clean build run

test-unit:
	@echo "Running unit tests..."
	go test ./models -v
	go test ./controllers -v -short

test-integration:
	@echo "Running integration tests..."
	go test . -v

test-all: test-unit test-integration

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

test-coverage-terminal:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

clean:
	rm -f coverage.out coverage.html

build:
	go build -o bin/todo-app main.go

run:
	go run main.go

test: test-unit

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down

docker-test: docker-up
	sleep 5
	$(MAKE) test-all
	$(MAKE) docker-down
