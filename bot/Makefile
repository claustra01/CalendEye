.PHONY: build env lint

build:
	go build -o build/calendeye

env:
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

test:
	go test ./...

lint:
	$$(go env GOPATH)/bin/golangci-lint run ./...

ci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$$(go env GOPATH)/bin/golangci-lint run ./...
