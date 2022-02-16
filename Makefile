## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [Y/n] ' && read ans && [ $${ans:-n} = Y ]
## lint: run linter checks
.PHONY: lint
lint: lint/basic
## lint/basic: run basic (mandatory) linter check. You should always run it before committing
.PHONY: lint/basic
lint/basic:
	golangci-lint run -v -c .golangci.yml
## lint/advanced: run advanced (optional) linter check
.PHONY: lint/advanced
lint/advanced:
	golangci-lint run -v -c .golangci.advanced.yml
## run/agent: run agent
.PHONY: run/server
run/agent:
	go run ./cmd/agent/
## run/server: run server
.PHONY: run/server
run/server:
	go run ./cmd/server/
