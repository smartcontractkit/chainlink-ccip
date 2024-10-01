TEST_COUNT ?= 10
COVERAGE_FILE ?= coverage.out

# Detect the system architecture
ARCH := $(shell uname -m)

# Find 'protoc' download URL based on the architecture
ifeq ($(ARCH),x86_64)
  PROTOC_ZIP := protoc-28.0-linux-x86_64.zip
else ifeq ($(ARCH),arm64)
  PROTOC_ZIP := protoc-28.0-osx-aarch_64.zip
else
  $(error Unsupported architecture: $(ARCH))
endif
PROTOC_URL := https://github.com/protocolbuffers/protobuf/releases/download/v28.0/$(PROTOC_ZIP)

PROTOC_BIN := /usr/local/bin/protoc

build: ensure_go_version
	go build -v ./...

generate: ensure_go_version clean-generate generate-protobuf generate-mocks

generate-mocks: ensure_go_version
	go install github.com/vektra/mockery/v2@v2.43.2
	mockery

generate-protobuf: ensure_go_version ensure_protoc_28_0
	$(PROTOC_BIN) --go_out=./commit/merkleroot/rmn/rmnpb --go_opt=paths=source_relative rmn_offchain.proto

clean-generate: ensure_go_version
	rm -rf ./commit/merkleroot/rmn/rmnpb/*
	rm -rf ./mocks/

test: ensure_go_version
	go test -race -fullpath -shuffle on -count $(TEST_COUNT) -coverprofile=$(COVERAGE_FILE) \
		`go list ./... | grep -Ev 'chainlink-ccip/internal/mocks|chainlink-ccip/mocks|chainlink-ccip/commit/merkleroot/rmn/rmnpb'`

lint: ensure_go_version ensure_golangcilint_1_59
	golangci-lint run -c .golangci.yml

checks: test lint

install-protoc:
	@echo "Downloading and installing protoc for $(ARCH)..."
	curl -OL $(PROTOC_URL)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)
	sudo chmod +x /usr/local/bin/protoc
	@echo "Installed protoc version:"
	$(PROTOC_BIN) --version
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31

install-golangcilint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.59

ensure_go_version:
	@go version | grep -q 'go1.22' || (echo "Please use go1.22" && exit 1)

ensure_golangcilint_1_59:
	@golangci-lint --version | grep -q '1.59' || (echo "Please use golangci-lint 1.59" && exit 1)

ensure_protoc_28_0:
	@$(PROTOC_BIN) --version | grep -q 'libprotoc 28.0' || (echo "Please use protoc 28.0, (make install-protoc)" && exit 1)
