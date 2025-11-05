.PHONY: build test run clean fmt vet tidy

build:
	go build ./cmd/api

test:
	go test ./...

run:
	go run ./cmd/api

clean:
	go clean
	rm -f api

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy