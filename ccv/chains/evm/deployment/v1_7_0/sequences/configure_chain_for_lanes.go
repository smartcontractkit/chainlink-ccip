package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
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
			defaultInboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultInboundCCVs))
			for _, ccv := range remoteConfig.DefaultInboundCCVs {
				defaultInboundCCVs = append(defaultInboundCCVs, common.HexToAddress(ccv))
			}
			laneMandatedInboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedInboundCCVs))
			for _, ccv := range remoteConfig.LaneMandatedInboundCCVs {
				laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, common.HexToAddress(ccv))
			}
			// Left-pad remoteConfig.OnRamps with zeros to the left to match the address bytes length
			onRamps := make([][]byte, 0, len(remoteConfig.OnRamps))
			for _, onRamp := range remoteConfig.OnRamps {
				onRamps = append(onRamps, common.LeftPadBytes(onRamp, 32))
			}
			offRampArgs = append(offRampArgs, offramp.SourceChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				SourceChainSelector: remoteSelector,
				IsEnabled:           remoteConfig.AllowTrafficFrom,
				OnRamps:             onRamps,
				DefaultCCVs:         defaultInboundCCVs,
				LaneMandatedCCVs:    laneMandatedInboundCCVs,
			})
			defaultOutboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultOutboundCCVs))
			for _, ccv := range remoteConfig.DefaultOutboundCCVs {
				defaultOutboundCCVs = append(defaultOutboundCCVs, common.HexToAddress(ccv))
			}
			laneMandatedOutboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedOutboundCCVs))
			for _, ccv := range remoteConfig.LaneMandatedOutboundCCVs {
				laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, common.HexToAddress(ccv))
			}
			onRampArgs = append(onRampArgs, onramp.DestChainConfigArgs{
				Router:               common.HexToAddress(input.Router),
				DestChainSelector:    remoteSelector,
				AddressBytesLength:   remoteConfig.AddressBytesLength,
				BaseExecutionGasCost: remoteConfig.BaseExecutionGasCost,
				DefaultCCVs:          defaultOutboundCCVs,
				LaneMandatedCCVs:     laneMandatedOutboundCCVs,
				DefaultExecutor:      common.HexToAddress(remoteConfig.DefaultExecutor), // The proxy address
				OffRamp:              remoteConfig.OffRamp,
			})
			gasPriceUpdates = append(gasPriceUpdates, fee_quoter.GasPriceUpdate{
				DestChainSelector: remoteSelector,
				UsdPerUnitGas:     remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas,
			})
			onRampAdds = append(onRampAdds, router.OnRamp{
				DestChainSelector: remoteSelector,
				OnRamp:            common.HexToAddress(input.OnRamp),
			})
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
				if destChainCfg != (fqc.FeeQuoterDestChainConfig{}) && destChainCfg.IsEnabled {
					continue
				}
			}
			feeQuoterArgs = append(feeQuoterArgs, fee_quoter.DestChainConfigArgs{
				DestChainSelector: remoteSelector,
				DestChainConfig:   remoteConfig.FeeQuoterDestChainConfig,
			})
		}

		// ApplySourceChainConfigUpdates on OffRamp
		offRampReport, err := cldf_ops.ExecuteOperation(b, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.OffRamp),
			Args:          offRampArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to OffRamp(%s) on chain %s: %w", input.OffRamp, chain, err)
		}
		writes = append(writes, offRampReport.Output)

		// ApplyDestChainConfigUpdates on OnRamp
		onRampReport, err := cldf_ops.ExecuteOperation(b, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.OnRamp),
			Args:          onRampArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to OnRamp(%s) on chain %s: %w", input.OnRamp, chain, err)
		}
		writes = append(writes, onRampReport.Output)

		// ApplyDestChainUpdates on each Executor
		for executorAddr, destChainSelectorsToAdd := range destChainSelectorsPerExecutor {
			executorReport, err := cldf_ops.ExecuteOperation(b, executor.ApplyDestChainUpdates, chain, contract.FunctionInput[executor.ApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       executorAddr,
				Args: executor.ApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: destChainSelectorsToAdd,
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

		// ApplyRampUpdates on Router
		routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(input.Router),
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
