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
PROTOC_BIN ?= /usr/local/bin/protoc

BUF_BIN := $(shell go env GOPATH)/bin/buf
COMPARE_AGAINST_BRANCH := main

build: ensure_go_version
	go build -v ./...

# If you have a different version of protoc installed, you can use the following command to generate the protobuf files
# make generate PROTOC_BIN=/path/to/protoc
generate: ensure_go_version clean-generate generate-protobuf generate-mocks

generate-mocks: ensure_go_version
	go install github.com/vektra/mockery/v2@v2.43.2
	mockery

# If you have a different version of protoc installed, you can use the following command to generate the protobuf files
# make generate-protobuf PROTOC_BIN=/path/to/protoc
proto-generate: ensure_go_version ensure_protoc_28_0
	$(PROTOC_BIN) --go_out=./pkg/ocrtypecodec/ocrtypecodecpb ./pkg/ocrtypecodec/ocrtypes.proto

proto-lint: ensure_go_version ensure_buf_version
	$(BUF_BIN) lint

proto-check-breaking: ensure_go_version ensure_buf_version
	$(BUF_BIN) breaking --against '.git#branch=$(COMPARE_AGAINST_BRANCH)'

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

install-protoc:
	@echo "Downloading and installing protoc for $(ARCH)..."
	curl -OL $(PROTOC_URL)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)
	sudo chmod +x $(PROTOC_BIN)
	@echo "Installed protoc version:"
	$(PROTOC_BIN) --version
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31

install-golangcilint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

ensure_go_version:
	@go version | grep -q 'go1.23' || (echo "Please use go1.23" && exit 1)

ensure_golangcilint_1_62_2:
	@golangci-lint --version | grep -q '1.62.2' || (echo "Please use golangci-lint 1.62.2" && exit 1)

ensure_protoc_28_0:
	@$(PROTOC_BIN) --version | grep -q 'libprotoc 28.0' || (echo "Please use protoc 28.0, (make install-protoc)" && exit 1)

install_buf:
	go install github.com/bufbuild/buf/cmd/buf@v1.50.0

ensure_buf_version:
	@$(BUF_BIN) --version | grep -q '1.50.0' || (echo "Please use buf 1.50.0" && exit 1)
