package sequences

import (
	"bytes"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/proxy"
	executor_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/executor"
	fqc "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
)

var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures an EVM chain as a source & destination for multiple remote chains",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input adapters.ConfigureChainForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Create inputs for each operation
		offRampArgs := make([]offramp.SourceChainConfigArgs, 0, len(input.RemoteChains))
		onRampArgs := make([]onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter.DestChainConfigArgs, 0, len(input.RemoteChains))
		gasPriceUpdates := make([]fee_quoter.GasPriceUpdate, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		destChainSelectorsPerExecutor := make(map[common.Address][]executor.RemoteChainConfigArgs)
		feeQContract, err := fqc.NewFeeQuoter(common.HexToAddress(input.FeeQuoter), chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to bind fee quoter contract at address %s on chain %s: %w",
				input.FeeQuoter, chain.String(), err)
		}
		for remoteSelector, remoteConfig := range input.RemoteChains {
			var err error
			offRampArgs, err = maybeAddSourceChainConfigArg(b, chain, chain.Selector, input, remoteSelector, remoteConfig, offRampArgs)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			onRampArgs, err = maybeAddOnRampDestChainConfigArg(b, chain, chain.Selector, input, remoteSelector, remoteConfig, onRampArgs)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			if remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas != nil {
				gasPriceReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestinationChainGasPrice, chain, contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(input.FeeQuoter),
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get gas prices on FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
				}
				if remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas.Cmp(gasPriceReport.Output.Value) != 0 {
					gasPriceUpdates = append(gasPriceUpdates, fee_quoter.GasPriceUpdate{
						DestChainSelector: remoteSelector,
						UsdPerUnitGas:     remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas,
					})
				}
			}

			onRampAddrReport, err := cldf_ops.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.Router),
				Args:          remoteSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get on ramp for dest %d from Router(%s) on chain %s: %w", remoteSelector, input.Router, chain, err)
			}
			if onRampAddrReport.Output != common.HexToAddress(input.OnRamp) {
				onRampAdds = append(onRampAdds, router.OnRamp{
					DestChainSelector: remoteSelector,
					OnRamp:            common.HexToAddress(input.OnRamp),
				})
			}

			offRampAdds = append(offRampAdds, router.OffRamp{
				SourceChainSelector: remoteSelector,
				OffRamp:             common.HexToAddress(input.OffRamp),
			})
			defaultExecutor := common.HexToAddress(remoteConfig.DefaultExecutor)
			getTargetReport, err := cldf_ops.ExecuteOperation(b, proxy.GetTarget, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       defaultExecutor,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get target address of Executor(%s) on chain %s: %w", defaultExecutor, chain, err)
			}
			if destChainSelectorsPerExecutor[getTargetReport.Output] == nil {
				destChainSelectorsPerExecutor[getTargetReport.Output] = []executor.RemoteChainConfigArgs{}
			}
			destChainSelectorsPerExecutor[getTargetReport.Output] = append(destChainSelectorsPerExecutor[getTargetReport.Output], executor.RemoteChainConfigArgs{
				DestChainSelector: remoteSelector,
				Config:            remoteConfig.ExecutorDestChainConfig,
			})
			// Only add dest chain config for fee quoter if OverrideExistingConfig is true, or
			// the config is not already set or enabled
			// otherwise we assume the fee quoter dest chain config is already set up correctly, and we don't want to override it
			if !remoteConfig.FeeQuoterDestChainConfig.OverrideExistingConfig {
				// fetch dest chain config from fee quoter
				// not using operation api as operation api defaults to already executed reports if run multiple times in same environment,
				// it might return stale data if feequoter config is updated after the last operation execution report was generated.
				// therefore call the onchain function directly to ensure we get the latest config for the dest chain
				destChainCfg, err := feeQContract.GetDestChainConfig(&bind.CallOpts{
					Context: b.GetContext(),
				}, remoteSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for "+
						"remote chain selector %d from fee quoter at address %s on chain %s: %w",
						remoteSelector, input.FeeQuoter, chain.String(), err)
				}
				if !destChainCfg.IsEnabled {
					feeQuoterArgs, err = maybeAddFeeQuoterDestChainConfigArg(feeQContract, b, input.FeeQuoter, chain, remoteSelector, remoteConfig, feeQuoterArgs)
					if err != nil {
						return sequences.OnChainOutput{}, err
					}
				}
			} else {
				var err error
				feeQuoterArgs, err = maybeAddFeeQuoterDestChainConfigArg(feeQContract, b, input.FeeQuoter, chain, remoteSelector, remoteConfig, feeQuoterArgs)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
			}
		}

		offRampAdds, err = filterOffRampAdds(b, chain, chain.Selector, common.HexToAddress(input.Router), offRampAdds)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		destChainSelectorsPerExecutor, err = filterExecutorDestChains(b, chain, chain.Selector, destChainSelectorsPerExecutor)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Apply OffRamp source chain configs (only entries that differ from on-chain).
		if len(offRampArgs) > 0 {
			offRampReport, err := cldf_ops.ExecuteOperation(b, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.OffRamp),
				Args:          offRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to OffRamp(%s) on chain %s: %w", input.OffRamp, chain, err)
			}
			writes = append(writes, offRampReport.Output)
		}

		// Apply OnRamp dest chain configs (only entries that differ from on-chain).
		if len(onRampArgs) > 0 {
			onRampReport, err := cldf_ops.ExecuteOperation(b, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.OnRamp),
				Args:          onRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to OnRamp(%s) on chain %s: %w", input.OnRamp, chain, err)
			}
			writes = append(writes, onRampReport.Output)
		}

		// Apply Executor dest chain updates (only chains that need to be added/updated).
		for executorAddr, toAdd := range destChainSelectorsPerExecutor {
			if len(toAdd) == 0 {
				continue
			}
			executorReport, err := cldf_ops.ExecuteOperation(b, executor.ApplyDestChainUpdates, chain, contract.FunctionInput[executor.ApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       executorAddr,
				Args: executor.ApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: toAdd,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to Executor(%s) on chain %s: %w", executorAddr, chain, err)
			}
			writes = append(writes, executorReport.Output)
		}
		if len(feeQuoterArgs) > 0 {
			// ApplyDestChainConfigUpdates on FeeQuoter
			feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.FeeQuoter),
				Args:          feeQuoterArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
			}
			writes = append(writes, feeQuoterReport.Output)
		}

		// UpdatePrices on FeeQuoter (gas prices only, as these are per dest chain)
		if len(gasPriceUpdates) > 0 {
			feeQuoterUpdatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract.FunctionInput[fee_quoter.PriceUpdates]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.FeeQuoter),
				Args: fee_quoter.PriceUpdates{
					GasPriceUpdates: gasPriceUpdates,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update gas prices on FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
			}
			writes = append(writes, feeQuoterUpdatePricesReport.Output)
		}

		// Apply Router ramp updates (only when there are changes).
		if len(onRampAdds) > 0 || len(offRampAdds) > 0 {
			routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.Router),
				Args: router.ApplyRampsUpdatesArgs{
					OnRampUpdates:  onRampAdds,
					OffRampRemoves: []router.OffRamp{}, // Removals should be processed by a separate sequence responsible for disconnecting lanes
					OffRampAdds:    offRampAdds,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", input.Router, chain, err)
			}
			writes = append(writes, routerReport.Output)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		batchOps := []mcms_types.BatchOperation{batchOp}

		for _, committeeVerifier := range input.CommitteeVerifiers {
			committeeVerifierReport, err := cldf_ops.ExecuteSequence(b, ConfigureCommitteeVerifierForLanes, chains, ConfigureCommitteeVerifierForLanesInput{
				ChainSelector:           chain.Selector,
				Router:                  input.Router,
				CommitteeVerifierConfig: committeeVerifier,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure committee verifier for lanes: %w", err)
			}
			batchOps = append(batchOps, committeeVerifierReport.Output.BatchOps...)
		}

		return sequences.OnChainOutput{
			BatchOps: batchOps,
		}, nil
	},
)

// maybeAddSourceChainConfigArg fetches current OffRamp source chain config, builds desired arg from input,
// and appends to offRampArgs only when the lane is enabled and config differs (idempotent).
func maybeAddSourceChainConfigArg(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, input adapters.ConfigureChainForLanesInput, remoteSelector uint64, remoteConfig adapters.RemoteChainConfig[[]byte, string], offRampArgs []offramp.SourceChainConfigArgs) ([]offramp.SourceChainConfigArgs, error) {
	defaultInboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultInboundCCVs))
	for _, ccv := range remoteConfig.DefaultInboundCCVs {
		defaultInboundCCVs = append(defaultInboundCCVs, common.HexToAddress(ccv))
	}
	laneMandatedInboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedInboundCCVs))
	for _, ccv := range remoteConfig.LaneMandatedInboundCCVs {
		laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, common.HexToAddress(ccv))
	}
	onRamps := make([][]byte, 0, len(remoteConfig.OnRamps))
	for _, onRamp := range remoteConfig.OnRamps {
		onRamps = append(onRamps, common.LeftPadBytes(onRamp, 32))
	}
	desiredOffRampArg := offramp.SourceChainConfigArgs{
		Router:              common.HexToAddress(input.Router),
		SourceChainSelector: remoteSelector,
		IsEnabled:           remoteConfig.AllowTrafficFrom,
		OnRamps:             onRamps,
		DefaultCCVs:         defaultInboundCCVs,
		LaneMandatedCCVs:    laneMandatedInboundCCVs,
	}
	offRampCurrentReport, err := cldf_ops.ExecuteOperation(b, offramp.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(input.OffRamp),
		Args:          remoteSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get source chain config for selector %d from OffRamp(%s) on chain %v: %w", remoteSelector, input.OffRamp, chain, err)
	}
	curOff := offRampCurrentReport.Output

	// Fall back to on-chain values if inputted values are empty
	if !desiredOffRampArg.IsEnabled {
		desiredOffRampArg.IsEnabled = curOff.IsEnabled
	}
	if len(desiredOffRampArg.OnRamps) == 0 {
		desiredOffRampArg.OnRamps = curOff.OnRamps
	}
	if len(desiredOffRampArg.DefaultCCVs) == 0 {
		desiredOffRampArg.DefaultCCVs = curOff.DefaultCCVs
	}
	if len(desiredOffRampArg.LaneMandatedCCVs) == 0 {
		desiredOffRampArg.LaneMandatedCCVs = curOff.LaneMandatedCCVs
	}
	if curOff.IsEnabled != desiredOffRampArg.IsEnabled ||
		curOff.Router != desiredOffRampArg.Router ||
		!UnorderedSliceEqual(curOff.OnRamps, desiredOffRampArg.OnRamps, bytes.Equal) ||
		!UnorderedSliceEqual(curOff.DefaultCCVs, desiredOffRampArg.DefaultCCVs, func(x, y common.Address) bool { return x == y }) ||
		!UnorderedSliceEqual(curOff.LaneMandatedCCVs, desiredOffRampArg.LaneMandatedCCVs, func(x, y common.Address) bool { return x == y }) {
		offRampArgs = append(offRampArgs, desiredOffRampArg)
	}
	return offRampArgs, nil
}

// maybeAddOnRampDestChainConfigArg fetches current OnRamp dest chain config, builds desired arg from input,
// and appends to onRampArgs only when config differs (idempotent).
func maybeAddOnRampDestChainConfigArg(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, input adapters.ConfigureChainForLanesInput, remoteSelector uint64, remoteConfig adapters.RemoteChainConfig[[]byte, string], onRampArgs []onramp.DestChainConfigArgs) ([]onramp.DestChainConfigArgs, error) {
	defaultOutboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultOutboundCCVs))
	for _, ccv := range remoteConfig.DefaultOutboundCCVs {
		defaultOutboundCCVs = append(defaultOutboundCCVs, common.HexToAddress(ccv))
	}
	laneMandatedOutboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedOutboundCCVs))
	for _, ccv := range remoteConfig.LaneMandatedOutboundCCVs {
		laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, common.HexToAddress(ccv))
	}
	desiredOnRampArg := onramp.DestChainConfigArgs{
		Router:                    common.HexToAddress(input.Router),
		DestChainSelector:         remoteSelector,
		AddressBytesLength:        remoteConfig.AddressBytesLength,
		BaseExecutionGasCost:      remoteConfig.BaseExecutionGasCost,
		TokenReceiverAllowed:      false, // TODO @kylesmartin: Add to core deployment input
		MessageNetworkFeeUSDCents: 0,     // TODO @kylesmartin: Add to core deployment input
		TokenNetworkFeeUSDCents:   0,     // TODO @kylesmartin: Add to core deployment input
		DefaultCCVs:               defaultOutboundCCVs,
		LaneMandatedCCVs:          laneMandatedOutboundCCVs,
		DefaultExecutor:           common.HexToAddress(remoteConfig.DefaultExecutor),
		OffRamp:                   remoteConfig.OffRamp,
	}
	onRampCurrentReport, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress(input.OnRamp),
		Args:          remoteSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get dest chain config for selector %d from OnRamp(%s) on chain %v: %w", remoteSelector, input.OnRamp, chain, err)
	}
	curOn := onRampCurrentReport.Output
	// Fall back to on-chain value if inputted value is empty
	// TODO @kylesmartin: Check tokenReceiverAllowed (should be pointer)
	if desiredOnRampArg.MessageNetworkFeeUSDCents == 0 {
		desiredOnRampArg.MessageNetworkFeeUSDCents = curOn.MessageNetworkFeeUSDCents
	}
	if desiredOnRampArg.TokenNetworkFeeUSDCents == 0 {
		desiredOnRampArg.TokenNetworkFeeUSDCents = curOn.TokenNetworkFeeUSDCents
	}
	if desiredOnRampArg.BaseExecutionGasCost == 0 {
		desiredOnRampArg.BaseExecutionGasCost = curOn.BaseExecutionGasCost
	}
	if desiredOnRampArg.AddressBytesLength == 0 {
		desiredOnRampArg.AddressBytesLength = curOn.AddressBytesLength
	}
	desiredDefaultCCVs := desiredOnRampArg.DefaultCCVs
	if len(desiredDefaultCCVs) == 0 {
		desiredDefaultCCVs = curOn.DefaultCCVs
	}
	desiredLaneMandatedCCVs := desiredOnRampArg.LaneMandatedCCVs
	if len(desiredLaneMandatedCCVs) == 0 {
		desiredLaneMandatedCCVs = curOn.LaneMandatedCCVs
	}
	if curOn.Router != desiredOnRampArg.Router || curOn.DefaultExecutor != desiredOnRampArg.DefaultExecutor ||
		!bytes.Equal(curOn.OffRamp, desiredOnRampArg.OffRamp) ||
		curOn.TokenReceiverAllowed != desiredOnRampArg.TokenReceiverAllowed ||
		curOn.MessageNetworkFeeUSDCents != desiredOnRampArg.MessageNetworkFeeUSDCents ||
		curOn.TokenNetworkFeeUSDCents != desiredOnRampArg.TokenNetworkFeeUSDCents ||
		curOn.BaseExecutionGasCost != desiredOnRampArg.BaseExecutionGasCost ||
		curOn.AddressBytesLength != desiredOnRampArg.AddressBytesLength ||
		!UnorderedSliceEqual(curOn.DefaultCCVs, desiredDefaultCCVs, func(x, y common.Address) bool { return x == y }) ||
		!UnorderedSliceEqual(curOn.LaneMandatedCCVs, desiredLaneMandatedCCVs, func(x, y common.Address) bool { return x == y }) {
		onRampArgs = append(onRampArgs, desiredOnRampArg)
	}
	return onRampArgs, nil
}

// feeQuoterDestChainConfigEqual reports whether the on-chain config matches the desired adapter config (binding struct has no USDPerUnitGas; that is updated via UpdatePrices).
func feeQuoterDestChainConfigEqual(cur fqc.FeeQuoterDestChainConfig, desired adapters.FeeQuoterDestChainConfig) bool {
	return cur.IsEnabled == desired.IsEnabled &&
		cur.MaxDataBytes == desired.MaxDataBytes &&
		cur.MaxPerMsgGasLimit == desired.MaxPerMsgGasLimit &&
		cur.DestGasOverhead == desired.DestGasOverhead &&
		cur.DestGasPerPayloadByteBase == desired.DestGasPerPayloadByteBase &&
		cur.ChainFamilySelector == desired.ChainFamilySelector &&
		cur.DefaultTokenFeeUSDCents == desired.DefaultTokenFeeUSDCents &&
		cur.DefaultTokenDestGasOverhead == desired.DefaultTokenDestGasOverhead &&
		cur.DefaultTxGasLimit == desired.DefaultTxGasLimit &&
		cur.NetworkFeeUSDCents == desired.NetworkFeeUSDCents &&
		cur.LinkFeeMultiplierPercent == desired.LinkFeeMultiplierPercent
}

// maybeAddFeeQuoterDestChainConfigArg fetches current FeeQuoter dest chain config and appends to feeQuoterArgs
// only when the config differs from desired (idempotent). Call only when OverrideExistingConfig is false.
// When a desired field is zero, the on-chain value is used so we do not overwrite with zero.
func maybeAddFeeQuoterDestChainConfigArg(feeQContract *fqc.FeeQuoter, b cldf_ops.Bundle, feeQuoterAddr string, chain evm.Chain, remoteSelector uint64, remoteConfig adapters.RemoteChainConfig[[]byte, string], feeQuoterArgs []fee_quoter.DestChainConfigArgs) ([]fee_quoter.DestChainConfigArgs, error) {
	cur, err := feeQContract.GetDestChainConfig(&bind.CallOpts{
		Context: b.GetContext(),
	}, remoteSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get dest chain config for remote chain selector %d from fee quoter at address %s on chain %s: %w",
			remoteSelector, feeQuoterAddr, chain.String(), err)
	}
	desired := remoteConfig.FeeQuoterDestChainConfig
	if !desired.IsEnabled {
		desired.IsEnabled = cur.IsEnabled
	}
	if desired.MaxDataBytes == 0 {
		desired.MaxDataBytes = cur.MaxDataBytes
	}
	if desired.MaxPerMsgGasLimit == 0 {
		desired.MaxPerMsgGasLimit = cur.MaxPerMsgGasLimit
	}
	if desired.DestGasOverhead == 0 {
		desired.DestGasOverhead = cur.DestGasOverhead
	}
	if desired.DestGasPerPayloadByteBase == 0 {
		desired.DestGasPerPayloadByteBase = cur.DestGasPerPayloadByteBase
	}
	if desired.ChainFamilySelector == [4]byte{} {
		desired.ChainFamilySelector = cur.ChainFamilySelector
	}
	if desired.DefaultTokenFeeUSDCents == 0 {
		desired.DefaultTokenFeeUSDCents = cur.DefaultTokenFeeUSDCents
	}
	if desired.DefaultTokenDestGasOverhead == 0 {
		desired.DefaultTokenDestGasOverhead = cur.DefaultTokenDestGasOverhead
	}
	if desired.DefaultTxGasLimit == 0 {
		desired.DefaultTxGasLimit = cur.DefaultTxGasLimit
	}
	if desired.NetworkFeeUSDCents == 0 {
		desired.NetworkFeeUSDCents = cur.NetworkFeeUSDCents
	}
	if desired.LinkFeeMultiplierPercent == 0 {
		desired.LinkFeeMultiplierPercent = cur.LinkFeeMultiplierPercent
	}
	if feeQuoterDestChainConfigEqual(cur, desired) {
		return feeQuoterArgs, nil
	}
	return append(feeQuoterArgs, fee_quoter.DestChainConfigArgs{
		DestChainSelector: remoteSelector,
		DestChainConfig:   desired,
	}), nil
}

// filterOffRampAdds returns only those OffRamp entries not already present on the Router.
func filterOffRampAdds(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, routerAddr common.Address, offRampAdds []router.OffRamp) ([]router.OffRamp, error) {
	currentReport, err := cldf_ops.ExecuteOperation(b, router.GetOffRamps, chain, contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       routerAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramps from Router(%s) on chain %v: %w", routerAddr, chain, err)
	}
	currentSet := make(map[router.OffRamp]struct{})
	for _, o := range currentReport.Output {
		currentSet[o] = struct{}{}
	}
	filtered := offRampAdds[:0]
	for _, add := range offRampAdds {
		if _, ok := currentSet[add]; !ok {
			filtered = append(filtered, add)
		}
	}
	return filtered, nil
}

// filterExecutorDestChains returns a copy of destChainSelectorsPerExecutor with each executor's list
// filtered to only dest chains that are not already configured or whose config differs.
func filterExecutorDestChains(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, destChainSelectorsPerExecutor map[common.Address][]executor.RemoteChainConfigArgs) (map[common.Address][]executor.RemoteChainConfigArgs, error) {
	out := make(map[common.Address][]executor.RemoteChainConfigArgs, len(destChainSelectorsPerExecutor))
	for executorAddr, toAdd := range destChainSelectorsPerExecutor {
		currentReport, err := cldf_ops.ExecuteOperation(b, executor.GetDestChains, chain, contract.FunctionInput[any]{
			ChainSelector: chainSelector,
			Address:       executorAddr,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get dest chains from Executor(%s) on chain %v: %w", executorAddr, chain, err)
		}
		currentMap := make(map[uint64]executor_bindings.ExecutorRemoteChainConfigArgs)
		for _, c := range currentReport.Output {
			currentMap[c.DestChainSelector] = c
		}
		filtered := toAdd[:0]
		for _, add := range toAdd {
			cur, ok := currentMap[add.DestChainSelector]
			if ok && cur.Config.UsdCentsFee == add.Config.USDCentsFee && cur.Config.Enabled == add.Config.Enabled {
				continue
			}
			filtered = append(filtered, add)
		}
		out[executorAddr] = filtered
	}
	return out, nil
}

