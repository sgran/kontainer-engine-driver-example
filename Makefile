.DEFAULT_GOAL := help
SHELL    := /bin/bash

.PHONY: build
build:  ## Build a binary for linux
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin ./...

.PHONY: help
help: ## Display this help
	@echo -e "Usage:\n  make \033[36m<target>\033[0m"
	@awk 'BEGIN {FS = ":.*##"}; \
		/^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } \
		/^###@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
