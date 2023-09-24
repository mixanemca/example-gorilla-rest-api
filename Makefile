# Makefile

.PHONY: test lint build

test:
	go test ./internal/logger

lint:
	golangci-lint run

build:
	go build -o example-gorilla-rest-api cmd/example-gorilla-rest-api/main.go

all: test lint build
