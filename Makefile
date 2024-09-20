.DEFAULT_GOAL := test

vet: fmt
	@go vet ./...

fmt:
	@go fmt ./...

test: vet
	@go test ./...
