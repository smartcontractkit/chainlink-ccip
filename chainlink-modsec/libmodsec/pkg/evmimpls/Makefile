.PHONY: all compile-contract generate-go

SOLC_VERSION ?= 0.8.19
SOLC_IMG = ethereum/solc:$(SOLC_VERSION)
CONTRACT_NAME = CCIPMessageSentEmitter
CONTRACT_PATH = testcontracts/$(CONTRACT_NAME).sol
ABI_DIR = testcontracts/abis
BIN_DIR = testcontracts/bin
GETH_WRAPPERS_DIR = gethwrappers
GO_PACKAGE = gethwrappers
ABIGEN = $(shell go env GOPATH)/bin/abigen

all: compile-contract generate-go test

compile-contract:
	@echo "Compiling $(CONTRACT_NAME)..."
	@mkdir -p $(ABI_DIR) $(BIN_DIR)
	@docker run --rm -v $(shell pwd)/testcontracts:/src $(SOLC_IMG) --abi --bin -o /src/abis --overwrite /src/$(CONTRACT_NAME).sol
	@mv $(ABI_DIR)/$(CONTRACT_NAME).abi $(ABI_DIR)/$(CONTRACT_NAME).json
	@docker run --rm -v $(shell pwd)/testcontracts:/src $(SOLC_IMG) --bin -o /src/bin --overwrite /src/$(CONTRACT_NAME).sol

generate-go: install-abigen
	@echo "Generating Go bindings..."
	@mkdir -p $(GETH_WRAPPERS_DIR)
	@$(ABIGEN) --abi=$(ABI_DIR)/$(CONTRACT_NAME).json --bin=$(BIN_DIR)/$(CONTRACT_NAME).bin --pkg=$(GO_PACKAGE) --type $(CONTRACT_NAME) --out=$(GETH_WRAPPERS_DIR)/$(CONTRACT_NAME).go

install-abigen:
	@echo "Installing abigen..."
	@go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.2

clean:
	@echo "Cleaning up..."
	@rm -rf $(ABI_DIR) $(BIN_DIR) $(GETH_WRAPPERS_DIR)

test:
	@echo "Running Go tests..."
	@go test -cover ./...
