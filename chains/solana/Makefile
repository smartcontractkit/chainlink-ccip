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

.PHONY: solana-checks
solana-checks:
	bash -c "echo 'Running cargo clippy'; cargo clippy --manifest-path ./contracts/programs/ccip-router/Cargo.toml" && \
	bash -c "echo 'Running make anchor-go-gen'; make anchor-go-gen" && \
	bash -c "echo 'Running make format'; make format" && \
	bash -c "echo 'Running make gomodtidy'; make gomodtidy" && \
	bash -c "echo 'Running make lint-go'; make lint-go" && \
	bash -c "echo 'Running Rust Tests'; cd ./contracts && make rust-tests" && \
	bash -c "echo 'Running Go E2E Tests'; make go-tests" && \
	bash -c "echo 'Running make generate-idl'; make generate-idl"
