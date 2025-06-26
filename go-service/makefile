.PHONY: build run fmt lint

build:
	go build -o bin/go-service ./cmd

run:
	go run ./cmd

fmt:
	goimports -w .
	gofmt -w .

lint:
	golangci-lint run
