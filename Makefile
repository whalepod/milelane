.PHONY: help

.DEFAULT_GOAL := help
_spanner='spanner://projects/$(PROJECT_ID)/instances/$(INSTANCE)/databases/$(DATABASE)' -path ./spanner
_mysql=${USER_NAME}:${PASSWARD}@/${DATABASE}?parseTime=true

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

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

migrate-create: ## create miration file ex) make migrate-create NAME=xxx_yyy_zzz
	goose -dir "db/migration" mysql ${_mysql} create ${NAME} sql

migrate-up: ## apply migraion
	goose -dir "db/migration" mysql ${_mysql} up
