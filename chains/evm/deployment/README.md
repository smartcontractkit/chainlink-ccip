# CCIP Deployments for EVM Chains

This module defines tooling for CCIP contracts on EVM-compatible chains. It provides a structured approach for deploying and configuring CCIP contracts through operations, sequences, and changesets.

## Core Components

- **Operations**: Produce a single-side effect (deploy contract, call function)
- **Sequences**: Ordered collections of operations that represent a complete workflow
- **Changesets**: Integration of a sequence with a deployment environment (MCMS, datastore, etc.)

Consumers can use the level of granularity they require.
- Want to execute MCMS proposals in chainlink-deployments (or a similar deployment environment)? Use a changeset.
- Want to complete a full operational story without integrating with a full-fledged deployment environment? Use a sequence.
- Want to make a single contract call? Use an operation. 

## Hierarchy

```
deployment/
├── utils/
│   ├── datastore/       # Utilities for working with datastore
│   ├── changesets/      # Utilities for building changesets
│   └── operations/      # Utilities for building operations
│       ├── call/
│       └── deployment/
├── v1_7_0/              # CCIP 1.7.0 operations, sequences, & changesets
│   ├── changesets/
│   ├── sequences/
│   └── operations/
├── v1_6_0/
└── ...
```

## Development Guide

Reference this guide when writing tooling to support contracts.

### Operations

Gethwrapper methods map 1:1 with operations. Operations yield reports that enable stateful retries, which are critical when you have a sequence that runs many operations. Use `contract.NewRead`, `contract.NewWrite`, and `contract.NewDeploy`. When creating operations, alias any gethwrapper types necessary such that consumers don't need to import gethwrappers when calling operations.

#### Write

```golang
var ApplySourceChainConfigUpdates = contract.NewWrite(
	"ccv-aggregator:apply-source-chain-config-updates", // Operation name - contract:method
	semver.MustParse("1.7.0"), // The contract version
	"Applies updates to source chain configurations on the CCVAggregator", // Operation description
	ContractType, // The contract type ("CCVAggregator" in this case)
	ccv_aggregator.CCVAggregatorABI, // Contract ABI - used to decode errors
	ccv_aggregator.NewCCVAggregator, // Contract constructor from gethwrappers
	contract.OnlyOwner, // Allowed callers check - used to determine whether or not the deployer key can make the call
	func([]SourceChainConfigArgs) error { return nil }, // Perform simple argument validations here (i.e. acceptable ranges)
	func(ccvAggregator *ccv_aggregator.CCVAggregator, opts *bind.TransactOpts, args []SourceChainConfigArgs) (*types.Transaction, error) {
		return ccvAggregator.ApplySourceChainConfigUpdates(opts, args)
	}, // Wrapper around gethwrappers call
)
```

#### Read

```golang
var GetStaticConfig = contract.NewRead(
	"ccv-aggregator:get-static-config", // Operation name - contract:method
	semver.MustParse("1.7.0"), // The contract version
	"Reads the static config of the CCVAggregator", // Operation description
	ContractType, // The contract type ("CCVAggregator" in this case)
	ccv_aggregator.NewCCVAggregator, // Contract constructor from gethwrappers
	func(ccvAggregator *ccv_aggregator.CCVAggregator, opts *bind.CallOpts, args any) (StaticConfig, error) {
		return ccvAggregator.GetStaticConfig(opts)
	}, // Wrapper around gethwrappers call
)
```

#### Deployment

```golang
var Deploy = contract.NewDeploy(
	"ccv-aggregator:deploy", // Operation name - contract:method
	semver.MustParse("1.7.0"), // The contract version
	"Deploys the CCVAggregator contract", // Operation description
	ContractType, // The contract type ("CCVAggregator" in this case)
	ccv_aggregator.CCVAggregatorABI, // Contract ABI - used to decode errors
	func(ConstructorArgs) error { return nil }, // Perform simple argument validations here (i.e. acceptable ranges)
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ccv_aggregator.DeployCCVAggregator(opts, backend, args)
			return address, tx, err
		},
		DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error) {...}
	}, // EVM & ZkSync deployment methodologies
)
```

### Sequences

A sequence is a composition of operations. Sequences accept a serializable input and a minimal set of dependencies. In the case of EVM, most sequences should only depend on `cldf_evm.Chain`. Coupling sequences too closely with the deployment environment makes them less portable.

For a reference implementation, see the [DeployChainContracts Sequence](/chains/evm/deployment/v1_7_0/sequences/deploy_chain_contracts.go). Notice how this sequence only targets one chain. It is simplest to keep sequence logic focused on synchronous steps. Leave it to another routine to handle the execution of multiple sequences concurrently.

### Changesets

Changesets essentially wrap sequences with the context of a deployment environment. For example, they can read addresses from a datastore, pass said addresses into sequences as input, and produce MCMS proposals based on the combination of sequence output and known MCMS addresses.

To create a changeset from a sequence that just makes calls and/or deploys contracts, use `changesets.NewFromOnChainSequence` combined with `datastore_utils.FindAndFormatEachRef` to avoid having to rewrite boilerplate for MCMS and datastore. For reference implementations, see [DeployChainContracts Changeset](/chains/evm/deployment/v1_7_0/changesets/deploy_chain_contracts.go) and [ConfigureChainForLanes Changeset](/chains/evm/deployment/v1_7_0/changesets/configure_chain_for_lanes.go).

## North Star

- **Define product-level APIs**: Every chain family should implement the same APIs (likely at the sequence-level), making it simple for engineers to operate cross-family.
- **Simplify changeset creation**: Since changesets primarily just wrap sequences, we should be able to create changesets from sequences in a simple way without having to rewrite a bunch of boilerplate.
- **Autogenerate operations**: Because they are so coupled with contract bindings, we should be able to autogenerate operations.
- **Avoid EVM coupling where possible**: Utilities like `changesets.NewFromOnChainSequence` can eventually be moved to a common deployment package. The outputs produced by each of the three operation utilities are already EVM-agnostic, so we may be able to abstract them further.
