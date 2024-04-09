# go.mk is a Go project general Makefile, encapsulated some common Target.
# Project repository: https://github.com/elliotxx/go-makefile

APPROOT     		?= $(shell basename $(PWD))
GOPKG       		?= $(shell go list 2>/dev/null)
GOPKGS      		?= $(shell go list ./... 2>/dev/null)
GOSOURCES   		?= $(shell find . -type f -name '*.go' ! -path '*Godeps/_workspace*')
# You can also customize GOSOURCE_PATHS, e.g. ./pkg/... ./cmd/...
GOSOURCE_PATHS		?= ././...
COVERAGEOUT 		?= coverage.out
COVERAGETMP 		?= coverage.tmp


# Go tools
GOFORMATER			?= gofumpt
GOFORMATER_VERSION	?= v0.2.0
GOLINTER			?= golangci-lint
GOLINTER_VERSION	?= v1.52.2


# To generate help information
.DEFAULT_GOAL := help
.PHONY: help
help:  ## This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' go.mk | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: cover-html
cover-html:  ## Generates coverage report and displays it in the browser
	go tool cover -html=$(COVERAGEOUT)

.PHONY: lint
lint:  ## Lint, will not fix but sets exit code on error
	@which $(GOLINTER) > /dev/null || (echo "Installing $(GOLINTER)@$(GOLINTER_VERSION) ..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINTER_VERSION) && echo -e "Installation complete!\n")
	$(GOLINTER) run --deadline=10m $(GOSOURCE_PATHS)

.PHONY: lint-fix
lint-fix:  ## Lint, will try to fix errors and modify code
	@which $(GOLINTER) > /dev/null || (echo "Installing $(GOLINTER)@$(GOLINTER_VERSION) ..."; go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLINTER_VERSION) && echo -e "Installation complete!\n")
	$(GOLINTER) run --deadline=10m $(GOSOURCE_PATHS) --fix

.PHONY: doc
doc:  ## Start the documentation server with godoc
	@which godoc > /dev/null || (echo "Installing godoc@latest ..."; go install golang.org/x/tools/cmd/godoc@latest && echo -e "Installation complete!\n")
	godoc -http=:6060
