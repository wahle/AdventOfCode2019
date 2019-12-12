DAY ?= day04

.PHONY: build test coverage clean lint format

build:
	go build -o ${DAY} ./days/${DAY}

test:
	go test -race -v ./...

coverage:
	gocov test ./... | gocov-xml -pwd > reports/coverage.xml

lint:
	golangci-lint run ./...

format:
	gofmt -w .
