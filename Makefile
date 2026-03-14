.PHONY: test build run-standalone install-gcl build-plugin lint clean

test:
	go test -v ./...

build:
	go build -o bin/loglint cmd/loglint/main.go

run-standalone: build
	./bin/loglint ./...

install-gcl:
	go install github.com/golangci/golangci-lint/cmd/custom-gcl@latest

build-plugin: install-gcl
	custom-gcl build .custom-gcl.yml

lint: build-plugin
	./custom-gcl run --config .golangci.yml ./...

clean:
	rm -rf bin/ custom-gcl custom-gcl.exe