.PHONY: test build run-standalone install-golangci build-plugin lint clean

test:
	go test -v ./...

build:
	go build -o bin/loglint cmd/loglint/main.go

run-standalone: build
	./bin/loglint ./...

install-golangci:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

build-plugin: install-golangci
	golangci-lint custom

lint: build-plugin
	./custom-gcl run --config .golangci.yml ./...

clean:
	rm -rf bin/ custom-gcl custom-gcl.exe