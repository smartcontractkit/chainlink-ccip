BIN_DIR = bin
export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on
export ANCHOR_VERSION ?=v0.29.0
export ANCHOR_IMAGE ?= backpackapp/build:$(ANCHOR_VERSION)

.PHONY: projectserum_version
anchor_version:
	@echo "${ANCHOR_VERSION}"

.PHONY: gomods
gomods: ## Install gomods
	go install github.com/jmank88/gomods@v0.1.3

.PHONY: gomodtidy
gomodtidy: gomods
	gomods tidy

.PHONY: lint-go
lint-go:
	golangci-lint --max-issues-per-linter 0 --max-same-issues 0 --color=always --exclude=dot-imports --timeout 15m --out-format checkstyle:golangci-lint-report.xml run

.PHONY: anchor-go-gen
anchor-go-gen:
	cd ./contracts && rm -rf ./target && anchor build && cd .. && ./scripts/anchor-go-gen.sh

.PHONY: format
format:
	go fmt ./... && cd ./contracts && cargo fmt

.PHONY: go-tests
go-tests:
	go test -v ./... -json -covermode=atomic -coverpkg=./... -coverprofile=integration_coverage.txt 2>&1 | tee /tmp/gotest.log | gotestloghelper -ci=true -singlepackage=true -hidepassingtests=false -hidepassinglogs=false

.PHONY: generate-idl
generate-idl:
	cd ./contracts && anchor build
