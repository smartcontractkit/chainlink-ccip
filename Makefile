TEST_COUNT ?= 10
COVERAGE_FILE ?= coverage.out

# Detect the system architecture
ARCH := $(shell uname -m)

build: ensure_go_version
	go build -v ./...

generate: ensure_go_version clean-generate generate-mocks

generate-mocks: ensure_go_version
	go install github.com/vektra/mockery/v2@v2.43.2
	mockery

clean-generate: ensure_go_version
	rm -rf ./mocks/

test:
	go test -race -fullpath -shuffle on -count $(TEST_COUNT) -coverprofile=$(COVERAGE_FILE) \
		`go list ./... | grep -Ev 'chainlink-ccip/internal/mocks|chainlink-ccip/mocks'`

lint: ensure_go_version ensure_golangcilint_1_62_2
	golangci-lint run -c .golangci.yml

lint-fix: ensure_go_version ensure_golangcilint_1_62_2
	golangci-lint run -c .golangci.yml --fix

checks: test lint

install-golangcilint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

ensure_go_version:
	@go version | grep -q 'go1.23' || (echo "Please use go1.23" && exit 1)

ensure_golangcilint_1_62_2:
	@golangci-lint --version | grep -q '1.62.2' || (echo "Please use golangci-lint 1.62.2" && exit 1)
