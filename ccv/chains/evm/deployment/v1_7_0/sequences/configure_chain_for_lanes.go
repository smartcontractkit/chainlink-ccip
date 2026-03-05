package sequences

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	executor_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/executor"
	fqc "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
)

// ConfigureChainForLanesInput is the input for the ConfigureChainForLanes sequence.
type ConfigureChainForLanesInput struct {
	// The selector of the chain being configured.
	ChainSelector uint64
	// The Router address on the chain being configured.
	// We assume that all connections defined will use the same router, either test or production.
	Router string
	// The OnRamp address on the chain being configured.
	// Similarly, we assume that all connections will use the same OnRamp.
	OnRamp string
	// The CommitteeVerifiers on the chain being configured.
	// There can be multiple committee verifiers on a chain, each controlled by a different entity.
	CommitteeVerifiers []lanes.CommitteeVerifierConfig[datastore.AddressRef]
	// The FeeQuoter address on the chain being configured.
	FeeQuoter string
	// The OffRamp address on the chain being configured
	OffRamp string
}

var ConfigureLaneLegAsSource = cldf_ops.NewSequence(
	"ConfigureLaneLegAsSource",
	semver.MustParse("1.0.0"),
	"Configures lane leg as source on CCIP 1.6.0",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		b.Logger.Infof("EVM Configuring lane leg as source. src: %+v, dest: %+v", input.Source, input.Dest)

		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.Source.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.Source.Selector)
		}

		sourceRouter := common.BytesToAddress(input.Source.Router).Hex()
		sourceOnRamp := common.BytesToAddress(input.Source.OnRamp).Hex()
		sourceFeeQuoter := common.BytesToAddress(input.Source.FeeQuoter).Hex()
		remoteSelector := input.Dest.Selector

		remoteConfig := input.Dest

		// Apply OnRamp dest chain configs (only entries that differ from on-chain).
		onRampArgs := make([]onramp.DestChainConfigArgs, 0, 1)
		onRampArgs, err := maybeAddOnRampDestChainConfigArg(b, chain, chain.Selector, input, remoteSelector, remoteConfig, onRampArgs)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if len(onRampArgs) > 0 {
			onRampReport, err := cldf_ops.ExecuteOperation(b, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(sourceOnRamp),
				Args:          onRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to OnRamp(%s) on chain %s: %w", sourceOnRamp, chain, err)
			}
			writes = append(writes, onRampReport.Output)
		}

		// Apply Executor dest chain updates (only chains that need to be added/updated).
		destChainSelectorsPerExecutor := make(map[common.Address][]executor.RemoteChainConfigArgs)
		defaultExecutor := common.HexToAddress(remoteConfig.DefaultExecutor.Address)
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
		destChainSelectorsPerExecutor, err = filterExecutorDestChains(b, chain, chain.Selector, destChainSelectorsPerExecutor)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
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

		// ApplyDestChainConfigUpdates on FeeQuoter
		feeQContract, err := fqc.NewFeeQuoter(common.HexToAddress(sourceFeeQuoter), chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to bind fee quoter contract at address %s on chain %s: %w",
				sourceFeeQuoter, chain.String(), err)
		}
		feeQuoterArgs := make([]fee_quoter.DestChainConfigArgs, 0, 1)
		if !remoteConfig.FeeQuoterDestChainConfig.OverrideExistingConfig {
			destChainCfg, err := feeQContract.GetDestChainConfig(&bind.CallOpts{
				Context: b.GetContext(),
			}, remoteSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for "+
					"remote chain selector %d from fee quoter at address %s on chain %s: %w",
					remoteSelector, sourceFeeQuoter, chain.String(), err)
			}
			if !destChainCfg.IsEnabled {
				feeQuoterArgs, err = maybeAddFeeQuoterDestChainConfigArg(feeQContract, b, sourceFeeQuoter, chain, remoteSelector, remoteConfig, feeQuoterArgs)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
			}
		} else {
			feeQuoterArgs, err = maybeAddFeeQuoterDestChainConfigArg(feeQContract, b, sourceFeeQuoter, chain, remoteSelector, remoteConfig, feeQuoterArgs)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		}
		if len(feeQuoterArgs) > 0 {
			feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(sourceFeeQuoter),
				Args:          feeQuoterArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", sourceFeeQuoter, chain, err)
			}
			writes = append(writes, feeQuoterReport.Output)
		}

		// UpdatePrices on FeeQuoter (gas prices only, as these are per dest chain)
		gasPriceUpdates := make([]fee_quoter.GasPriceUpdate, 0, 1)
		if remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas != nil {
			gasPriceReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestinationChainGasPrice, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(sourceFeeQuoter),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get gas prices on FeeQuoter(%s) on chain %s: %w", sourceFeeQuoter, chain, err)
			}
			if remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas.Cmp(gasPriceReport.Output.Value) != 0 {
				gasPriceUpdates = append(gasPriceUpdates, fee_quoter.GasPriceUpdate{
					DestChainSelector: remoteSelector,
					UsdPerUnitGas:     remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas,
				})
			}
		}
		if len(gasPriceUpdates) > 0 {
			feeQuoterUpdatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract.FunctionInput[fee_quoter.PriceUpdates]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(sourceFeeQuoter),
				Args: fee_quoter.PriceUpdates{
					GasPriceUpdates: gasPriceUpdates,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update gas prices on FeeQuoter(%s) on chain %s: %w", sourceFeeQuoter, chain, err)
			}
			writes = append(writes, feeQuoterUpdatePricesReport.Output)
		}

		// Apply Router ramp updates (only when there are changes).
		onRampAdds := make([]router.OnRamp, 0, 1)
		onRampAddrReport, err := cldf_ops.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(sourceRouter),
			Args:          remoteSelector,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get on ramp for dest %d from Router(%s) on chain %s: %w", remoteSelector, sourceRouter, chain, err)
		}
		if onRampAddrReport.Output != common.HexToAddress(sourceOnRamp) {
			onRampAdds = append(onRampAdds, router.OnRamp{
				DestChainSelector: remoteSelector,
				OnRamp:            common.HexToAddress(sourceOnRamp),
			})
		}
		if len(onRampAdds) > 0 {
			routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(sourceRouter),
				Args: router.ApplyRampsUpdatesArgs{
					OnRampUpdates:  onRampAdds,
					OffRampRemoves: []router.OffRamp{},
					OffRampAdds:    []router.OffRamp{},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", sourceRouter, chain, err)
			}
			writes = append(writes, routerReport.Output)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		batchOps := []mcms_types.BatchOperation{batchOp}

		for _, cv := range input.Source.CommitteeVerifiers {
			filtered := filterCommitteeVerifierForRemote(cv, remoteSelector)
			if len(filtered.RemoteChains) == 0 {
				continue
			}
			committeeVerifierReport, err := cldf_ops.ExecuteSequence(b, ConfigureCommitteeVerifierForLanes, chains, ConfigureCommitteeVerifierForLanesInput{
				ChainSelector:           chain.Selector,
				Router:                  sourceRouter,
				CommitteeVerifierConfig: filtered,
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

var ConfigureLaneLegAsDest = cldf_ops.NewSequence(
	"ConfigureLaneLegAsDest",
	semver.MustParse("1.6.0"),
	"Configures lane leg as destination on CCIP 1.6.0",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input lanes.UpdateLanesInput) (sequences.OnChainOutput, error) {
		b.Logger.Infof("EVM Configuring lane leg as destination. src: %+v, dest: %+v", input.Source, input.Dest)

		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.Dest.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.Dest.Selector)
		}

		destRouter := common.BytesToAddress(input.Dest.Router).Hex()
		destOffRamp := common.BytesToAddress(input.Dest.OffRamp).Hex()
		remoteSelector := input.Source.Selector

		remoteConfig := input.Source

		// Apply OffRamp source chain configs (only entries that differ from on-chain).
		offRampArgs := make([]offramp.SourceChainConfigArgs, 0, 1)
		offRampArgs, err := maybeAddSourceChainConfigArg(b, chain, chain.Selector, input, remoteSelector, remoteConfig, offRampArgs)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if len(offRampArgs) > 0 {
			offRampReport, err := cldf_ops.ExecuteOperation(b, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(destOffRamp),
				Args:          offRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to OffRamp(%s) on chain %s: %w", destOffRamp, chain, err)
			}
			writes = append(writes, offRampReport.Output)
		}

		// Apply Router ramp updates (only when there are changes).
		offRampAdds := []router.OffRamp{
			{
				SourceChainSelector: remoteSelector,
				OffRamp:             common.HexToAddress(destOffRamp),
			},
		}
		offRampAdds, err = filterOffRampAdds(b, chain, chain.Selector, common.HexToAddress(destRouter), offRampAdds)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if len(offRampAdds) > 0 {
			routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(destRouter),
				Args: router.ApplyRampsUpdatesArgs{
					OnRampUpdates:  []router.OnRamp{},
					OffRampRemoves: []router.OffRamp{},
					OffRampAdds:    offRampAdds,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", destRouter, chain, err)
			}
			writes = append(writes, routerReport.Output)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		batchOps := []mcms_types.BatchOperation{batchOp}

		for _, cv := range input.Dest.CommitteeVerifiers {
			filtered := filterCommitteeVerifierForRemote(cv, remoteSelector)
			if len(filtered.RemoteChains) == 0 {
				continue
			}
			committeeVerifierReport, err := cldf_ops.ExecuteSequence(b, ConfigureCommitteeVerifierForLanes, chains, ConfigureCommitteeVerifierForLanesInput{
				ChainSelector:           chain.Selector,
				Router:                  destRouter,
				CommitteeVerifierConfig: filtered,
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
func maybeAddSourceChainConfigArg(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, input lanes.UpdateLanesInput, remoteSelector uint64, remoteConfig *lanes.ChainDefinition, offRampArgs []offramp.SourceChainConfigArgs) ([]offramp.SourceChainConfigArgs, error) {
	defaultInboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultInboundCCVs))
	for _, ccv := range remoteConfig.DefaultInboundCCVs {
		defaultInboundCCVs = append(defaultInboundCCVs, common.HexToAddress(ccv.Address))
	}
	laneMandatedInboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedInboundCCVs))
	for _, ccv := range remoteConfig.LaneMandatedInboundCCVs {
		laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, common.HexToAddress(ccv.Address))
	}
	onRamps := make([][]byte, 0, len(remoteConfig.OnRampsForRemotes))
	for _, onRamp := range remoteConfig.OnRampsForRemotes {
		onRamps = append(onRamps, common.LeftPadBytes(onRamp, 32))
	}
	desiredOffRampArg := offramp.SourceChainConfigArgs{
		Router:              common.BytesToAddress(input.Dest.Router),
		SourceChainSelector: remoteSelector,
		IsEnabled:           !input.IsDisabled,
		OnRamps:             onRamps,
		DefaultCCVs:         defaultInboundCCVs,
		LaneMandatedCCVs:    laneMandatedInboundCCVs,
	}
	offRampCurrentReport, err := cldf_ops.ExecuteOperation(b, offramp.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       common.BytesToAddress(input.Dest.OffRamp),
		Args:          remoteSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get source chain config for selector %d from OffRamp(%s) on chain %v: %w", remoteSelector, input.Dest.OffRamp, chain, err)
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
func maybeAddOnRampDestChainConfigArg(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, input lanes.UpdateLanesInput, remoteSelector uint64, remoteConfig *lanes.ChainDefinition, onRampArgs []onramp.DestChainConfigArgs) ([]onramp.DestChainConfigArgs, error) {
	defaultOutboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultOutboundCCVs))
	for _, ccv := range remoteConfig.DefaultOutboundCCVs {
		defaultOutboundCCVs = append(defaultOutboundCCVs, common.HexToAddress(ccv.Address))
	}
	laneMandatedOutboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedOutboundCCVs))
	for _, ccv := range remoteConfig.LaneMandatedOutboundCCVs {
		laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, common.HexToAddress(ccv.Address))
	}
	desiredOnRampArg := onramp.DestChainConfigArgs{
		Router:                    common.BytesToAddress(input.Source.Router),
		DestChainSelector:         remoteSelector,
		AddressBytesLength:        remoteConfig.AddressBytesLength,
		BaseExecutionGasCost:      remoteConfig.BaseExecutionGasCost,
		MessageNetworkFeeUSDCents: remoteConfig.MessageNetworkFeeUSDCents,
		TokenNetworkFeeUSDCents:   remoteConfig.TokenNetworkFeeUSDCents,
		DefaultCCVs:               defaultOutboundCCVs,
		LaneMandatedCCVs:          laneMandatedOutboundCCVs,
		DefaultExecutor:           common.HexToAddress(remoteConfig.DefaultExecutor.Address),
		OffRamp:                   remoteConfig.OffRamp,
	}
	onRampCurrentReport, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       common.BytesToAddress(input.Source.OnRamp),
		Args:          remoteSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get dest chain config for selector %d from OnRamp(%s) on chain %v: %w", remoteSelector, input.Source.OnRamp, chain, err)
	}
	curOn := onRampCurrentReport.Output
	// Fall back to on-chain value if inputted value is empty
	if remoteConfig.TokenReceiverAllowed == nil {
		desiredOnRampArg.TokenReceiverAllowed = curOn.TokenReceiverAllowed
	} else {
		desiredOnRampArg.TokenReceiverAllowed = *remoteConfig.TokenReceiverAllowed
	}
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
func feeQuoterDestChainConfigEqual(cur fqc.FeeQuoterDestChainConfig, desired lanes.FeeQuoterDestChainConfig) bool {
	return cur.IsEnabled == desired.IsEnabled &&
		cur.MaxDataBytes == desired.MaxDataBytes &&
		cur.MaxPerMsgGasLimit == desired.MaxPerMsgGasLimit &&
		cur.DestGasOverhead == desired.DestGasOverhead &&
		cur.DestGasPerPayloadByteBase == desired.DestGasPerPayloadByteBase &&
		cur.ChainFamilySelector == [4]byte(binary.BigEndian.AppendUint32(nil, desired.ChainFamilySelector)) &&
		cur.DefaultTokenFeeUSDCents == desired.DefaultTokenFeeUSDCents &&
		cur.DefaultTokenDestGasOverhead == desired.DefaultTokenDestGasOverhead &&
		cur.DefaultTxGasLimit == desired.DefaultTxGasLimit &&
		cur.NetworkFeeUSDCents == desired.NetworkFeeUSDCents &&
		cur.LinkFeeMultiplierPercent == desired.LinkFeeMultiplierPercent
}

// maybeAddFeeQuoterDestChainConfigArg fetches current FeeQuoter dest chain config and appends to feeQuoterArgs
// only when the config differs from desired (idempotent). Call only when OverrideExistingConfig is false.
// When a desired field is zero, the on-chain value is used so we do not overwrite with zero.
func maybeAddFeeQuoterDestChainConfigArg(feeQContract *fqc.FeeQuoter, b cldf_ops.Bundle, feeQuoterAddr string, chain evm.Chain, remoteSelector uint64, remoteConfig *lanes.ChainDefinition, feeQuoterArgs []fee_quoter.DestChainConfigArgs) ([]fee_quoter.DestChainConfigArgs, error) {
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
	if desired.ChainFamilySelector == 0 {
		desired.ChainFamilySelector = binary.BigEndian.Uint32(cur.ChainFamilySelector[:])
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
		DestChainConfig:   adapterDestChainConfigToFeeQuoter(desired),
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

func adapterDestChainConfigToFeeQuoter(cfg lanes.FeeQuoterDestChainConfig) fee_quoter.DestChainConfig {
	return fee_quoter.DestChainConfig{
		IsEnabled:                   cfg.IsEnabled,
		MaxDataBytes:                cfg.MaxDataBytes,
		MaxPerMsgGasLimit:           cfg.MaxPerMsgGasLimit,
		DestGasOverhead:             cfg.DestGasOverhead,
		DestGasPerPayloadByteBase:   cfg.DestGasPerPayloadByteBase,
		ChainFamilySelector:         [4]byte(binary.BigEndian.AppendUint32(nil, cfg.ChainFamilySelector)),
		DefaultTokenFeeUSDCents:     cfg.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead: cfg.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:           cfg.DefaultTxGasLimit,
		NetworkFeeUSDCents:          cfg.NetworkFeeUSDCents,
		LinkFeeMultiplierPercent:    cfg.LinkFeeMultiplierPercent,
	}
}

// filterCommitteeVerifierForRemote converts a lanes CommitteeVerifierConfig to the adapters
// type, filtering its RemoteChains to only include the given remote selector.
func filterCommitteeVerifierForRemote(cv lanes.CommitteeVerifierConfig[datastore.AddressRef], remoteSelector uint64) lanes.CommitteeVerifierConfig[datastore.AddressRef] {
	remoteChains := make(map[uint64]lanes.CommitteeVerifierRemoteChainConfig)
	if rc, ok := cv.RemoteChains[remoteSelector]; ok {
		remoteChains[remoteSelector] = lanes.CommitteeVerifierRemoteChainConfig{
			AllowlistEnabled:          rc.AllowlistEnabled,
			AddedAllowlistedSenders:   rc.AddedAllowlistedSenders,
			RemovedAllowlistedSenders: rc.RemovedAllowlistedSenders,
			FeeUSDCents:               rc.FeeUSDCents,
			GasForVerification:        rc.GasForVerification,
			PayloadSizeBytes:          rc.PayloadSizeBytes,
			SignatureConfig: lanes.CommitteeVerifierSignatureQuorumConfig{
				Signers:   rc.SignatureConfig.Signers,
				Threshold: rc.SignatureConfig.Threshold,
			},
		}
	}
	return lanes.CommitteeVerifierConfig[datastore.AddressRef]{
		CommitteeVerifier: cv.CommitteeVerifier,
		RemoteChains:      remoteChains,
	}
}
