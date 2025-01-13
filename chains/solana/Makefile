BIN_DIR = bin
export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on
export ANCHOR_VERSION ?=v0.29.0
export ANCHOR_IMAGE ?= backpackapp/build:$(ANCHOR_VERSION)

.PHONY: projectserum_version
anchor_version:
	@echo "${ANCHOR_VERSION}"

.PHONY: clippy
clippy:
	cargo clippy --manifest-path ./contracts/programs/ccip-router/Cargo.toml

.PHONY: gomods
gomods: ## Install gomods
	go install github.com/jmank88/gomods@v0.1.3

.PHONY: gomodtidy
gomodtidy: gomods
	gomods tidy

.PHONY: format-contracts
format-contracts:
	cd ./contracts && cargo fmt && go fmt ./...

.PHONY: rust-tests
rust-tests:
	cd ./contracts && cargo test

.PHONY: lint-go
lint-go:
	golangci-lint --max-issues-per-linter 0 --max-same-issues 0 --color=always --exclude=dot-imports --timeout 15m --out-format checkstyle:golangci-lint-report.xml run

.PHONY: lint-go-fix
lint-go-fix:
	golangci-lint --max-issues-per-linter 0 --max-same-issues 0 --color=always --exclude=dot-imports --timeout 15m run --verbose --fix

.PHONY: anchor-go-gen
anchor-go-gen:
	cd ./contracts && rm -rf ./target && anchor build && cd .. && ./scripts/anchor-go-gen.sh

.PHONY: format
format:
	go fmt ./... && cd ./contracts && cargo fmt

.PHONY: go-tests
go-tests:
	go test -v ./... -json -covermode=atomic -coverpkg=./... -coverprofile=integration_coverage.txt 2>&1 | tee /tmp/gotest.log | gotestloghelper -ci=true -singlepackage=true -hidepassingtests=false -hidepassinglogs=false

.PHONY: ccip-router-contracts-go-tests
ccip-router-contracts-go-tests:
	go test -run TestCCIPRouter -v -failfast ./... -json -covermode=atomic -coverpkg=./... -coverprofile=integration_coverage.txt 2>&1 | tee /tmp/gotest.log | gotestloghelper -ci=true -singlepackage=true -hidepassingtests=false -hidepassinglogs=false

.PHONY: build-contracts
build-contracts:
	cd ./contracts && anchor build

.PHONY: solana-checks
solana-checks: clippy anchor-go-gen format gomodtidy lint-go rust-tests go-tests build-contracts
