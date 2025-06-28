.PHONY: build run fmt lint test

build:
	go build -o bin/go-service ./cmd

run:
	go run ./cmd

fmt:
	goimports -w .
	gofmt -w .

lint:
	golangci-lint run

test:
	go test ./...
