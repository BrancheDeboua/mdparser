.DEFAULT_GOAL := vet

vet: fmt
	go vet ./...

fmt:
	go fmt ./...