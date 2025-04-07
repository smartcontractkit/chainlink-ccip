# ALL_FOUNDRY_PRODUCTS contains a list of all products that have a foundry
# profile defined and use the Foundry snapshots.
ALL_FOUNDRY_PRODUCTS = ccip functions keystone l2ep llo-feeds operatorforwarder shared workflow data-feeds

# To make a snapshot for a specific product, either set the `FOUNDRY_PROFILE` env var
# or call the target with `FOUNDRY_PROFILE=product`
# When developing with Foundry, you'll most likely already have the env var set
# to run the tests for the product you're working on. In that case, you can just
# call `make snapshot` and it will use the env var.
# env var example
#		export FOUNDRY_PROFILE=llo-feeds
#		make snapshot
# make call example
# 		make FOUNDRY_PROFILE=llo-feeds snapshot
# note make snapshot skips fuzz tests named according to best practices, although forge uses
# a static fuzz seed by default, flaky gas results per platform are still observed.
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
	./scripts/lcov_prune ./lcov.info ./lcov.info.pruned
	genhtml -o report lcov.info.pruned --branch-coverage

# To generate gethwrappers for a specific product, either set the `FOUNDRY_PROFILE`
# env var or call the target with `FOUNDRY_PROFILE=product`
# This uses FOUNDRY_PROFILE, even though it does support non-foundry products. This
# is to improve the workflow for developers working with Foundry, which is the
# recommended way to develop Solidity for CL products.
# env var example
#		export FOUNDRY_PROFILE=llo-feeds
#		make wrappers
# make call example
# 		make FOUNDRY_PROFILE=llo-feeds wrappers
.PHONY: wrappers
wrappers: pnpmdep mockery abigen ## Recompiles solidity contracts and their go wrappers.
	./scripts/compile_all
	go generate ../core/gethwrappers/$(FOUNDRY_PROFILE)

# Use this to generate compiled JSON artifacts for gauntlet-plus-plus.
# This is currently only used for CCIP.
# example: make artifact-generate contract=LockReleaseTokenPoolAndProxy solcversion=0.8.24 artifactpath=../../gauntlet-plus-plus/packages-ethereum/operations-ccip/src/artifacts/1.5.0/lock-release-token-pool-and-proxy.json
artifact-generate: export FOUNDRY_PROFILE=ccip
.PHONY: artifact-generate
artifact-generate:
	chmod +x ./scripts/generate_compiled_json_ccip.sh
	./scripts/generate_compiled_json_ccip.sh $(contract) $(solcversion) $(artifactpath)


.PHONY: pnpmdep
pnpmdep: ## Install solidity contract dependencies through pnpm
	 pnpm i

.PHONY: abigen
abigen: ## Build & install abigen.
	./scripts/build_abigen

.PHONY: mockery
mockery: $(mockery) ## Install mockery.
	go install github.com/vektra/mockery/v2@v2.52.3
