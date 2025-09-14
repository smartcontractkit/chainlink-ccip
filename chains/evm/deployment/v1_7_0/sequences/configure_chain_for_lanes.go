package sequences

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/defensive_example_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type CommitOnRampDestChainConfig struct {
	// Whether or not to allow traffic TO the remote chain
	AllowlistEnabled bool
	// Addresses that are allowed to send messages TO the remote chain
	AddedAllowlistedSenders []common.Address
	// Addresses that are no longer allowed to send messages TO the remote chain
	RemovedAllowlistedSenders []common.Address
}

type RemoteChainConfig struct {
	// Whether or not to allow traffic FROM this remote chain
	AllowTrafficFrom bool
	// The address on the remote chain from which the message is emitted
	// For example, on EVM chains, this is the CCVProxy
	CCIPMessageSource []byte
	// The default CCVs that will be applied to messages FROM this remote chain if no receiver is specified
	DefaultCCVOffRamps []common.Address
	// Any CCVs that must always be used for messages FROM this remote chain
	LaneMandatedCCVOffRamps []common.Address
	// The CCVs that will be used for messages TO this remote chain if none are specified
	DefaultCCVOnRamps []common.Address
	// The CCVs that will always be applied to messages TO this remote chain
	LaneMandatedCCVOnRamps []common.Address
	// The executor that will be used for messages TO this remote chain if none is specified
	// The address corresponds to the ExecutorOnRamp contract
	DefaultExecutor common.Address
	// CommitOnRampDestChainConfig configures the CommitOnRamp for this remote chain
	CommitOnRampDestChainConfig CommitOnRampDestChainConfig
	// FeeQuoterDestChainConfig configures the FeeQuoter for this remote chain
	FeeQuoterDestChainConfig fee_quoter_v2.DestChainConfig
}

type ConfigureChainForLanesInput struct {
	// The selector of the EVM chain being configured
	ChainSelector uint64
	// The router on the EVM chain being configured
	// We assume that all connections will use the same router, either test or production
	Router common.Address
	// The CCVProxy on the EVM chain being configured
	// Similarly, we assume that all connections will use the same CCVProxy
	CCVProxy common.Address
	// The CommitOnRamp on the EVM chain being configured
	CommitOnRamp common.Address
	// The CommitOffRamp on the EVM chain being configured
	CommitOffRamp common.Address
	// The DefensiveExampleReceiver on the EVM chain being configured
	DefensiveExampleReceiver common.Address
	// The FeeQuoter on the EVM chain being configured
	FeeQuoter common.Address
	// The CCVAggregator on the EVM chain being configured
	CCVAggregator common.Address
	// The configuration of each remote chain to configure
	RemoteChains map[uint64]RemoteChainConfig
}

var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures an EVM chain as a source & destination for multiple remote chains",
	func(b operations.Bundle, chain evm.Chain, input ConfigureChainForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)

		// Create inputs for each operation
		ccvAggregatorArgs := make([]ccv_aggregator.SourceChainConfigArgs, 0, len(input.RemoteChains))
		ccvProxyArgs := make([]ccv_proxy.DestChainConfigArgs, 0, len(input.RemoteChains))
		commitOnRampDestConfigArgs := make([]commit_onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		commitOnRampAllowlistArgs := make([]commit_onramp.AllowlistConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter_v2.DestChainConfigArgs, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		destChainSelectorsPerExecutor := make(map[common.Address][]uint64)
		remoteChainConfigsForReceiver := make(map[uint64]defensive_example_receiver.RemoteChainConfig, len(input.RemoteChains))
		for remoteSelector, remoteConfig := range input.RemoteChains {
			extraArgs, err := newGenericExtraArgsV2(big.NewInt(500_000), true)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to encode GenericExtraArgsV2 for remote chain %d: %w", remoteSelector, err)
			}

			ccvAggregatorArgs = append(ccvAggregatorArgs, ccv_aggregator.SourceChainConfigArgs{
				Router:              input.Router,
				SourceChainSelector: remoteSelector,
				IsEnabled:           remoteConfig.AllowTrafficFrom,
				OnRamp:              remoteConfig.CCIPMessageSource,
				DefaultCCV:          remoteConfig.DefaultCCVOffRamps,
				LaneMandatedCCVs:    remoteConfig.LaneMandatedCCVOffRamps,
			})
			ccvProxyArgs = append(ccvProxyArgs, ccv_proxy.DestChainConfigArgs{
				Router:            input.Router,
				DestChainSelector: remoteSelector,
				DefaultCCVs:       remoteConfig.DefaultCCVOnRamps,
				LaneMandatedCCVs:  remoteConfig.LaneMandatedCCVOnRamps,
				DefaultExecutor:   remoteConfig.DefaultExecutor,
			})
			commitOnRampDestConfigArgs = append(commitOnRampDestConfigArgs, commit_onramp.DestChainConfigArgs{
				CcvProxy:          input.CCVProxy,
				DestChainSelector: remoteSelector,
				AllowlistEnabled:  remoteConfig.CommitOnRampDestChainConfig.AllowlistEnabled,
			})
			commitOnRampAllowlistArgs = append(commitOnRampAllowlistArgs, commit_onramp.AllowlistConfigArgs{
				AllowlistEnabled:          remoteConfig.CommitOnRampDestChainConfig.AllowlistEnabled,
				AddedAllowlistedSenders:   remoteConfig.CommitOnRampDestChainConfig.AddedAllowlistedSenders,
				RemovedAllowlistedSenders: remoteConfig.CommitOnRampDestChainConfig.RemovedAllowlistedSenders,
			})
			feeQuoterArgs = append(feeQuoterArgs, fee_quoter_v2.DestChainConfigArgs{
				DestChainSelector: remoteSelector,
				DestChainConfig:   remoteConfig.FeeQuoterDestChainConfig,
			})
			onRampAdds = append(onRampAdds, router.OnRamp{
				DestChainSelector: remoteSelector,
				OnRamp:            input.CCVProxy,
			})
			offRampAdds = append(offRampAdds, router.OffRamp{
				SourceChainSelector: remoteSelector,
				OffRamp:             input.CCVAggregator,
			})
			if destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor] == nil {
				destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor] = []uint64{}
			}
			destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor] = append(destChainSelectorsPerExecutor[remoteConfig.DefaultExecutor], remoteSelector)
			remoteChainConfigsForReceiver[remoteSelector] = defensive_example_receiver.RemoteChainConfig{
				RemoteChainSelector: remoteSelector,
				ExtraArgs:           extraArgs,
				RequiredCCVs:        []common.Address{input.CommitOffRamp},
				OptionalCCVs:        []common.Address{},
				OptionalThreshold:   0,
			}
		}

		// EnableRemoteChain on DefensiveExampleReceiver
		for remoteChainSel := range input.RemoteChains {
			receiverReport, err := cldf_ops.ExecuteOperation(b, defensive_example_receiver.EnableRemoteChain, chain, contract.FunctionInput[defensive_example_receiver.RemoteChainConfig]{
				Address:       input.DefensiveExampleReceiver,
				ChainSelector: chain.Selector,
				Args:          remoteChainConfigsForReceiver[remoteChainSel],
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to enable remote chain %d on DefensiveExampleReceiver(%s) on chain %s: %w", remoteChainSel, input.CommitOffRamp, chain, err)
			}
			writes = append(writes, receiverReport.Output)
		}

		// ApplySourceChainConfigUpdates on CCVAggregator
		ccvAggregatorReport, err := cldf_ops.ExecuteOperation(b, ccv_aggregator.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]ccv_aggregator.SourceChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CCVAggregator,
			Args:          ccvAggregatorArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to CCVAggregator(%s) on chain %s: %w", input.CCVAggregator, chain, err)
		}
		writes = append(writes, ccvAggregatorReport.Output)

		// ApplyDestChainConfigUpdates on CCVProxy
		ccvProxyReport, err := cldf_ops.ExecuteOperation(b, ccv_proxy.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]ccv_proxy.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CCVProxy,
			Args:          ccvProxyArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to CCVProxy(%s) on chain %s: %w", input.CCVProxy, chain, err)
		}
		writes = append(writes, ccvProxyReport.Output)

		// ApplyDestChainConfigUpdates on CommitOnRamp
		commitOnRampReport, err := cldf_ops.ExecuteOperation(b, commit_onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]commit_onramp.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CommitOnRamp,
			Args:          commitOnRampDestConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to CommitOnRamp(%s) on chain %s: %w", input.CommitOnRamp, chain, err)
		}
		writes = append(writes, commitOnRampReport.Output)

		// ApplyDestChainUpdates on each ExecutorOnRamp
		for executorOnRampAddr, destChainSelectorsToAdd := range destChainSelectorsPerExecutor {
			executorOnRampReport, err := cldf_ops.ExecuteOperation(b, executor_onramp.ApplyDestChainUpdates, chain, contract.FunctionInput[executor_onramp.ApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       executorOnRampAddr,
				Args: executor_onramp.ApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: destChainSelectorsToAdd,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to ExecutorOnRamp(%s) on chain %s: %w", executorOnRampAddr, chain, err)
			}
			writes = append(writes, executorOnRampReport.Output)
		}

		// ApplyAllowlistUpdates on CommitOnRamp
		commitOnRampAllowlistReport, err := cldf_ops.ExecuteOperation(b, commit_onramp.ApplyAllowlistUpdates, chain, contract.FunctionInput[[]commit_onramp.AllowlistConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CommitOnRamp,
			Args:          commitOnRampAllowlistArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allowlist updates to CommitOnRamp(%s) on chain %s: %w", input.CommitOnRamp, chain, err)
		}
		writes = append(writes, commitOnRampAllowlistReport.Output)

		// ApplyDestChainConfigUpdates on FeeQuoter
		feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter_v2.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter_v2.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.FeeQuoter,
			Args:          feeQuoterArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
		}
		writes = append(writes, feeQuoterReport.Output)

		// ApplyRampUpdates on Router
		routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.Router,
			Args: router.ApplyRampsUpdatesArgs{
				OnRampUpdates:  onRampAdds,
				OffRampRemoves: []router.OffRamp{}, // removals should be processed by a separate sequence responsible for disconnecting lanes
				OffRampAdds:    offRampAdds,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", input.Router, chain, err)
		}
		writes = append(writes, routerReport.Output)

		return sequences.OnChainOutput{
			Writes: writes,
		}, nil
	},
)

// newGenericExtraArgsV2 encodes the GenericExtraArgsV2 struct according to the ABI
func newGenericExtraArgsV2(gasLimit *big.Int, allowOutOfOrderExecution bool) ([]byte, error) {
	clientABI := `
			[
				{
					"name": "encodeGenericExtraArgsV2",
					"type": "function",
					"inputs": [
						{
							"components": [
								{
									"name": "gasLimit",
									"type": "uint256"
								},
								{
									"name": "allowOutOfOrderExecution",
									"type": "bool"
								}
							],
							"name": "args",
							"type": "tuple"
						}
					],
					"outputs": [],
					"stateMutability": "pure"
				}
			]
			`

	parsedABI, err := abi.JSON(bytes.NewReader([]byte(clientABI)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	genericExtraArgsV2 := struct {
		GasLimit                 *big.Int
		AllowOutOfOrderExecution bool
	}{
		GasLimit:                 gasLimit,
		AllowOutOfOrderExecution: allowOutOfOrderExecution,
	}
	encoded, err := parsedABI.Methods["encodeGenericExtraArgsV2"].Inputs.Pack(genericExtraArgsV2)
	if err != nil {
		return nil, fmt.Errorf("failed to ABI encode GenericExtraArgsV2: %w", err)
	}

	tag := []byte{0x18, 0x1d, 0xcf, 0x10} // GENERIC_EXTRA_ARGS_V2_TAG
	return append(tag, encoded...), nil
}

/*
// newEVMExtraArgsV3 encodes the EVMExtraArgsV3 struct according to the ABI
// TODO: Use later once FeeQuoter supports it
func newEVMExtraArgsV3(requiredCCVs, optionalCCVs []common.Address, optionalThreshold uint8, finalityConfig uint32, executor common.Address, executorArgs, tokenArgs []byte) ([]byte, error) {
	clientABI := `
			[
				{
					"name": "encodeEVMExtraArgsV3",
					"type": "function",
					"inputs": [
						{
							"components": [
								{
									"internalType": "address[]",
									"name": "requiredCCV",
									"type": "address[]"
								},
								{
									"internalType": "address[]",
									"name": "optionalCCV",
									"type": "address[]"
								},
								{
									"internalType": "uint8",
									"name": "optionalThreshold",
									"type": "uint8"
								},
								{
									"internalType": "uint32",
									"name": "finalityConfig",
									"type": "uint32"
								},
								{
									"internalType": "address",
									"name": "executor",
									"type": "address"
								},
								{
									"internalType": "bytes",
									"name": "executorArgs",
									"type": "bytes"
								},
								{
									"internalType": "bytes",
									"name": "tokenArgs",
									"type": "bytes"
								}
							],
							"internalType": "struct EVMExtraArgsV3",
							"name": "args",
							"type": "tuple"
						}
					],
					"outputs": [],
					"stateMutability": "pure"
				}
			]
			`

	parsedABI, err := abi.JSON(bytes.NewReader([]byte(clientABI)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	evmExtraArgsV3 := struct {
		RequiredCCV       []common.Address
		OptionalCCV       []common.Address
		OptionalThreshold uint8
		FinalityConfig    uint32
		Executor          common.Address
		ExecutorArgs      []byte
		TokenArgs         []byte
	}{
		RequiredCCV:       requiredCCVs,
		OptionalCCV:       optionalCCVs,
		OptionalThreshold: optionalThreshold,
		FinalityConfig:    finalityConfig,
		Executor:          executor,
		ExecutorArgs:      executorArgs,
		TokenArgs:         tokenArgs,
	}
	encoded, err := parsedABI.Methods["encodeEVMExtraArgsV3"].Inputs.Pack(evmExtraArgsV3)
	if err != nil {
		return nil, fmt.Errorf("failed to ABI encode EVMExtraArgsV3: %w", err)
	}

	tag := []byte{0x30, 0x23, 0x26, 0xcb} // GENERIC_EXTRA_ARGS_V3_TAG
	return append(tag, encoded...), nil
}
*/
