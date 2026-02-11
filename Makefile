.PHONY: help build run test clean dev lint format

help:
	@echo "Available commands:"
	@echo "  build    - Build the application"
	@echo "  run      - Run the application"
	@echo "  dev      - Start development environment with docker-compose"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  lint     - Run linter"
	@echo "  format   - Format code"

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go

dev:
	docker compose up --build

test:
	go test -v ./...

clean:
	rm -rf bin/
	docker compose down --volumes

lint:
	golangci-lint run

format:
	go fmt ./...

mod-tidy:
	go mod tidy