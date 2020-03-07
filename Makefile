.PHONY: help

.DEFAULT_GOAL := help

vet: ## Run vet
	go vet ./...

lint: ## Run golint
ifeq ($(shell command -v golint 2> /dev/null),)
	GO111MODULE=off go get golang.org/x/lint/golint
endif
	golint -set_exit_status $$(go list ./...)

fmt: ## Run gofmt
	gofmt -s -w .

test: ## Run test
	go test -race -v -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html
