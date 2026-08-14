TEST_COUNT ?= 1
COVERAGE_FILE ?= coverage.out

.PHONY: build
build:
	go build -o bin/runbook .

.PHONY: run
run:
	go run . list

.PHONY: test
test:
	go test ./...

.PHONY: test-race
test-race:
	go test -race -count $(TEST_COUNT) -coverprofile=$(COVERAGE_FILE) ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: checks
checks: fmt vet test

# Backfill go.sum after go.mod changes.
.PHONY: tidy
tidy:
	go mod tidy
