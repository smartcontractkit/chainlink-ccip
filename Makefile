TEST_COUNT ?= 10
COVERAGE_FILE ?= coverage.out

build: ensure_go_version
	go build -v ./...

generate: ensure_go_version
	go install github.com/vektra/mockery/v2@v2.43.2
	mockery

test: ensure_go_version
	go test -race -fullpath -shuffle on -count $(TEST_COUNT) -coverprofile=$(COVERAGE_FILE) ./...

lint: ensure_go_version
	golangci-lint run -c .golangci.yml

ensure_go_version:
	@go version | grep -q 'go1.22' || (echo "Please use go1.22" && exit 1)

ensure_golangcilint_1_59:
	@golangci-lint --version | grep -q '1.59' || (echo "Please use golangci-lint 1.59" && exit 1)
