.PHONY: build test help fmt

default: help

build: fmt test ## Builds an executable
	env GOOS=linux GOARCH=amd64 GO111MODULE=on go build

test: ## Runs go test with coverage
	GO111MODULE=on go test ./... -cover

fmt: ## Verifies all files have been `gofmt`ed
	@echo "+ $@"
	@gofmt -s -l . | grep -v '.pb.go:' | grep -v vendor | tee /dev/stderr

install: test ## Installs the executable or package
	env GOOS=linux GOARCH=amd64 GO111MODULE=on go install

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
