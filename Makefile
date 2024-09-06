TEST_COUNT ?= 10
COVERAGE_FILE ?= coverage.out

build: ensure_go_version
	go build -v ./...

generate: ensure_go_version clean-generate generate-protobuf generate-mocks

generate-mocks: ensure_go_version
	go install github.com/vektra/mockery/v2@v2.43.2
	mockery

generate-protobuf: ensure_go_version
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=./commit/merkleroot/rmn --go_opt=paths=source_relative rmn_offchain.proto

clean-generate: ensure_go_version
	rm -rf ./commit/merkleroot/rmn/rmn_offchain.pb.go
	rm -rf ./mocks/

test: ensure_go_version
	go test -race -fullpath -shuffle on -count $(TEST_COUNT) -coverprofile=$(COVERAGE_FILE) `go list ./... | grep -Ev 'chainlink-ccip/internal/mocks|chainlink-ccip/mocks'`

lint: ensure_go_version ensure_golangcilint_1_59
	golangci-lint run -c .golangci.yml

checks: test lint

ensure_go_version:
	@go version | grep -q 'go1.22' || (echo "Please use go1.22" && exit 1)

ensure_golangcilint_1_59:
	@golangci-lint --version | grep -q '1.59' || (echo "Please use golangci-lint 1.59" && exit 1)
