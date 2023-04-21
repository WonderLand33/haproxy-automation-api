# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory

GO_MODULE_NAME = haproxy-automation-api
VERSION_FLAG=-X '$(GO_MODULE_NAME)/pkg/version.gitBranch=`git branch --show-current`' \
-X '$(GO_MODULE_NAME)/pkg/version.gitCommit=`git rev-parse HEAD`' \
-X '$(GO_MODULE_NAME)/pkg/version.gitTag=`git describe --always`' \
-X '$(GO_MODULE_NAME)/pkg/version.buildUser=`whoami`' \
-X '$(GO_MODULE_NAME)/pkg/version.buildDate=`date +'%Y-%m-%dT%H:%M:%SZ'`'

GO_LDFLAGS=-ldflags "-s $(VERSION_FLAG)"

.PHONY: help
help: ## Describe useful make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'

.PHONY: all
all: ## Lint, and build (default)
	$(MAKE) lint
	$(MAKE) build

.PHONY: clean
clean: ## Delete intermediate build artifacts
	@# -X only removes untracked files, -d recurses into directories, -f actually removes files/dirs
	git clean -Xdf

.PHONY: build
build: ## Build all packages
	rm -rf dist
	GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o dist/haproxy-automation-api-linux-x86 main.go
	GOOS=darwin GOARCH=arm64 go build $(GO_LDFLAGS) -o dist/haproxy-automation-api-darwin-arm64 main.go
	GOOS=darwin GOARCH=amd64 go build $(GO_LDFLAGS) -o dist/haproxy-automation-api-darwin-x86 main.go

.PHONY: lint
lint: ## Lint Go and protobuf
	test -z "$$(buf format -d . | tee /dev/stderr)"
	go vet ./...
	golangci-lint run
	staticcheck ./...

.PHONY: lintfix
lintfix: ## Automatically fix some lint errors
	golangci-lint run --fix

.PHONY: upgrade
upgrade: ## Upgrade dependencies
	go get -u -t ./... && go mod tidy -v

.PHONY: init
init: ## Install cmd dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
