#!/usr/bin/env make -f

all: test lint build

## test: Run the go test command.
.PHONY: test
test:
	go test -v ./...

## lint: Run the linting.
.PHONY: lint
lint:
	golangci-lint run ./...

## build: Compile the binary.
.PHONY: build
build:
	go build -o example-gorilla-rest-api cmd/example-gorilla-rest-api/main.go

## help: Show this message.
.PHONY: help
help: Makefile
	@echo "Available targets:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
