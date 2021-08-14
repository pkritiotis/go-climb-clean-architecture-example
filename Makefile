.DEFAULT_GOAL := help

# `make help` generates a help message for each target that
# has a comment starting with ##
help:
	@echo "Please use 'make <target>' where <target> is one of the following:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	GO111MODULE=on go run -mod=vendor ./cmd/main.go

lint: ## Perform linting
	golangci-lint run --disable-all -E revive  --exclude-use-default=false --modules-download-mode=vendor

test: ## Run unit tests
	go test -mod=vendor `go list ./... | grep -v 'docs'` -race

test-int: ## Run all tests
	go test -mod=vendor `go list ./... | grep -v 'docs'` -race -tags=integration

build: ## Build the app executable for Linux
	CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -mod=vendor -a -installsuffix cgo -o ./go-climb ./cmd/main.go

fmt: ## Format the source code
	go fmt ./...


