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
# test: run tests
.PHONY: test
test: test/server/controllers/metrics
## test/coverage/profile: generate coverage profile
.PHONY: test/coverage/profile
test/coverage/profile:
	@[ -d tmp ] || mkdir tmp
	@go clean -testcache && \
	go test -race -cover -p 1 \
		github.com/gitalek/go-runtime-monitor/internal/server/controllers/metrics \
		-coverprofile=tmp/profile.out
## test/coverage: show detailed (per function) test coverage in terminal
.PHONY: test/coverage
test/coverage: test/coverage/profile
	go tool cover -func=tmp/profile.out
## test/coverage/html: generate and show detailed test coverage (html-format) in browser
.PHONY: test/coverage/html
test/coverage/html: test/coverage/profile
	go tool cover -html=tmp/profile.out
# test/server/controllers/metrics: run metrics tests
.PHONY: test/server/controllers/metrics
test/server/controllers/metrics:
	go clean -testcache && go test -race ./internal/server/controllers/metrics
