TEST_COUNT ?= 10
COVERAGE_FILE ?= coverage.out

build:
	go build -v ./...

test:
	go test -race -fullpath -shuffle on -count $(TEST_COUNT) -coverprofile=$(COVERAGE_FILE) ./...

lint:
	golangci-lint run -c ../../.golangci.yml

lint-fix:
	golangci-lint run -c ../../.golangci.yml --fix

checks: test lint

$(GOPATH)/bin/go-enum:
	go get -u github.com/abice/go-enum
	go install github.com/abice/go-enum

.PHONY: generate
generate: $(GOPATH)/bin/go-enum
	go generate ./...