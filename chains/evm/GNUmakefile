# Creates a gas snapshot
# note `make snapshot` skips fuzz/fork/reverting tests.
.PHONY: snapshot
snapshot: ## Make a snapshot for a specific product.
	export FOUNDRY_PROFILE=ccip && forge snapshot --nmt "test?(Fuzz|Fork|.*_RevertWhen)_.*"

.PHONY: snapshot-diff
snapshot-diff: ## Make a snapshot for a specific product.
	export FOUNDRY_PROFILE=ccip && forge snapshot --nmt "test?(Fuzz|Fork|.*_RevertWhen)_.*" --diff

.PHONY: foundry
foundry: ## Install foundry.
	foundryup --install v1.0.0

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

.PHONY: pnpmdep
pnpmdep: ## Install solidity contract dependencies through pnpm
	 pnpm i

.PHONY: abigen
abigen: ## Build & install abigen.
	./scripts/build_abigen

.PHONY: mockery
mockery: $(mockery) ## Install mockery.
	go install github.com/vektra/mockery/v2@v2.52.3
