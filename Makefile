.PHONY: build test run install generate

build:
	go build -o tide cmd/tide/main.go

test:
	go test ./...

run:
	go run cmd/tide/main.go

install:
	./install.sh

generate:
	go run cmd/generator/main.go
