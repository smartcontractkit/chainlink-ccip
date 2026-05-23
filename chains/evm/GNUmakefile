.PHONY: test
test: ## Run the full Foundry test suite (always uses the `ccip` profile).
	FOUNDRY_PROFILE=ccip forge test

# Creates a gas snapshot
# note `make snapshot` skips fuzz/fork/reverting tests.
.PHONY: snapshot
snapshot: ## Make a snapshot for a specific product.
	FOUNDRY_PROFILE=ccip forge snapshot --nmt "test?(Fuzz|Fork|.*_RevertWhen)_.*"

.PHONY: snapshot-diff
snapshot-diff: ## Make a snapshot for a specific product.
	FOUNDRY_PROFILE=ccip forge snapshot --nmt "test?(Fuzz|Fork|.*_RevertWhen)_.*" --diff

.PHONY: foundry
foundry: ## Install foundry.
	foundryup --install v1.7.1

# it may be necessary to install libusb on your machine as a dependency for this script to work
.PHONY: foundry-zksync
foundry-zksync: ## Install foundry.
	./scripts/install_foundry_zksync

ccip-precommit: export FOUNDRY_PROFILE=ccip
.PHONY: ccip-precommit
ccip-precommit:
	forge test
	make snapshot
	forge fmt
	pnpm solhint

ccip-lcov: export FOUNDRY_PROFILE=ccip
.PHONY: ccip-lcov
ccip-lcov:
	forge coverage --report lcov
	./scripts/lcov_prune
	genhtml -o report lcov.info.pruned --branch-coverage --ignore-errors inconsistent,corrupt

# Generated the gethwrappers
.PHONY: wrappers
wrappers: pnpmdep mockery abigen
	./scripts/compile_all
	go generate ./gobindings

wrappers-release: export metadata=true
.PHONY: wrappers-release
wrappers-release: ## Generate the gethwrappers for release which includes metadata.
	 make wrappers

.PHONY: wrappers-zksync
wrappers-zksync: pnpmdep foundry-zksync
	ZKSYNC=true make wrappers

.PHONY: pnpmdep
pnpmdep: ## Install solidity contract dependencies through pnpm
	 pnpm i

.PHONY: abigen
abigen: ## Build & install abigen.
	./scripts/build_abigen

.PHONY: mockery
mockery: $(mockery) ## Install mockery.
	go install github.com/vektra/mockery/v2@v2.52.3

.PHONY: extract-bytecode-abi
extract-bytecode-abi: ## Extract bytecode and abi from versioned gobindings to bytecode and abi directories
	cd gobindings && go run ./cmd/extract_bytecode/main.go

OPERATIONS_GEN_V2 := github.com/smartcontractkit/chainlink-deployments-framework/tools/operations-gen@v0.0.2-0.20260515181606-62d38f5a89b0

.PHONY: operations-legacy
operations-legacy:
	go run cmd/operations-gen/main.go -config deployment/operations_gen_config.yaml

.PHONY: operations-1-5
operations-1-5:
	go run $(OPERATIONS_GEN_V2) -config deployment/operations_gen_config_1_5.yaml

.PHONY: operations
operations: pnpmdep wrappers operations-legacy operations-1-5

.PHONY: operations-fast
operations-fast: operations-legacy operations-1-5
