package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures an EVM chain as a source & destination for multiple remote chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input adapters.ConfigureChainForLanesInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Create inputs for each operation
		offRampArgs := make([]offramp.SourceChainConfigArgs, 0, len(input.RemoteChains))
		onRampArgs := make([]onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		committeeVerifierDestConfigArgs := make([]committee_verifier.DestChainConfigArgs, 0, len(input.RemoteChains))
		committeeVerifierAllowlistArgs := make([]committee_verifier.AllowlistConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter.DestChainConfigArgs, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		destChainSelectorsPerExecutor := make(map[common.Address][]executor.RemoteChainConfigArgs)
		for remoteSelector, remoteConfig := range input.RemoteChains {
			defaultInboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultInboundCCVs))
			for _, ccv := range remoteConfig.DefaultInboundCCVs {
				defaultInboundCCVs = append(defaultInboundCCVs, common.HexToAddress(ccv))
			}
			laneMandatedInboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedInboundCCVs))
			for _, ccv := range remoteConfig.LaneMandatedInboundCCVs {
				laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, common.HexToAddress(ccv))
			}
			offRampArgs = append(offRampArgs, offramp.SourceChainConfigArgs{
				Router:              common.HexToAddress(input.Router),
				SourceChainSelector: remoteSelector,
				IsEnabled:           remoteConfig.AllowTrafficFrom,
				OnRamp:              remoteConfig.OnRamp,
				DefaultCCV:          defaultInboundCCVs,
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
				DefaultExecutor:      common.HexToAddress(remoteConfig.DefaultExecutor),
				OffRamp:              remoteConfig.OffRamp,
			})
			committeeVerifierDestConfigArgs = append(committeeVerifierDestConfigArgs, committee_verifier.DestChainConfigArgs{
				Router:             common.HexToAddress(input.Router),
				DestChainSelector:  remoteSelector,
				AllowlistEnabled:   remoteConfig.CommitteeVerifierDestChainConfig.AllowlistEnabled,
				FeeUSDCents:        remoteConfig.CommitteeVerifierDestChainConfig.FeeUSDCents,
				GasForVerification: remoteConfig.CommitteeVerifierDestChainConfig.GasForVerification,
				PayloadSizeBytes:   remoteConfig.CommitteeVerifierDestChainConfig.PayloadSizeBytes,
			})
			addedAllowlistedSenders := make([]common.Address, 0, len(remoteConfig.CommitteeVerifierDestChainConfig.AddedAllowlistedSenders))
			for _, sender := range remoteConfig.CommitteeVerifierDestChainConfig.AddedAllowlistedSenders {
				addedAllowlistedSenders = append(addedAllowlistedSenders, common.HexToAddress(sender))
			}
			removedAllowlistedSenders := make([]common.Address, 0, len(remoteConfig.CommitteeVerifierDestChainConfig.RemovedAllowlistedSenders))
			for _, sender := range remoteConfig.CommitteeVerifierDestChainConfig.RemovedAllowlistedSenders {
				removedAllowlistedSenders = append(removedAllowlistedSenders, common.HexToAddress(sender))
			}
			committeeVerifierAllowlistArgs = append(committeeVerifierAllowlistArgs, committee_verifier.AllowlistConfigArgs{
				AllowlistEnabled:          remoteConfig.CommitteeVerifierDestChainConfig.AllowlistEnabled,
				AddedAllowlistedSenders:   addedAllowlistedSenders,
				RemovedAllowlistedSenders: removedAllowlistedSenders,
				DestChainSelector:         remoteSelector,
			})
			feeQuoterArgs = append(feeQuoterArgs, fee_quoter.DestChainConfigArgs{
				DestChainSelector: remoteSelector,
				DestChainConfig:   remoteConfig.FeeQuoterDestChainConfig,
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
			if destChainSelectorsPerExecutor[defaultExecutor] == nil {
				destChainSelectorsPerExecutor[defaultExecutor] = []executor.RemoteChainConfigArgs{}
			}
			destChainSelectorsPerExecutor[defaultExecutor] = append(destChainSelectorsPerExecutor[defaultExecutor], executor.RemoteChainConfigArgs{
				DestChainSelector: remoteSelector,
				Config:            remoteConfig.ExecutorDestChainConfig,
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

		for _, committeeVerifier := range input.CommitteeVerifiers {
			// ApplyDestChainConfigUpdates on each CommitteeVerifier
			committeeVerifierReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier.Implementation),
				Args:          committeeVerifierDestConfigArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to CommitteeVerifier(%s) on chain %s: %w", committeeVerifier, chain, err)
			}
			writes = append(writes, committeeVerifierReport.Output)

			// ApplyInboundImplementationUpdates on each CommitteeVerifierResolver, first fetching the version tag
			versionTagReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetVersionTag, chain, contract.FunctionInput[any]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier.Implementation),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from CommitteeVerifier(%s) on chain %s: %w", committeeVerifier, chain, err)
			}
			inboundImplementationArgs := []versioned_verifier_resolver.InboundImplementationArgs{
				{
					Version:  versionTagReport.Output,
					Verifier: common.HexToAddress(committeeVerifier.Implementation),
				},
			}
			applyInboundUpdatesReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier.Resolver),
				Args:          inboundImplementationArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply inbound implementation updates to CommitteeVerifierResolver(%s) on chain %s: %w", committeeVerifier.Resolver, chain, err)
			}
			writes = append(writes, applyInboundUpdatesReport.Output)

			// ApplyOutboundImplementationUpdates on each CommitteeVerifierResolver
			outboundImplementationArgs := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0, len(input.RemoteChains))
			for remoteChainSelector := range input.RemoteChains {
				outboundImplementationArgs = append(outboundImplementationArgs, versioned_verifier_resolver.OutboundImplementationArgs{
					DestChainSelector: remoteChainSelector,
					Verifier:          common.HexToAddress(committeeVerifier.Implementation),
				})
			}
			applyOutboundUpdatesReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier.Resolver),
				Args:          outboundImplementationArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply outbound implementation updates to CommitteeVerifierResolver(%s) on chain %s: %w", committeeVerifier.Resolver, chain, err)
			}
			writes = append(writes, applyOutboundUpdatesReport.Output)
		}

		// ApplyDestChainUpdates on each Executor
		for ExecutorAddr, destChainSelectorsToAdd := range destChainSelectorsPerExecutor {
			ExecutorReport, err := cldf_ops.ExecuteOperation(b, executor.ApplyDestChainUpdates, chain, contract.FunctionInput[executor.ApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       ExecutorAddr,
				Args: executor.ApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: destChainSelectorsToAdd,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to Executor(%s) on chain %s: %w", ExecutorAddr, chain, err)
			}
			writes = append(writes, ExecutorReport.Output)
		}

		// ApplyAllowlistUpdates on CommitteeVerifier
		for _, committeeVerifier := range input.CommitteeVerifiers {
			committeeVerifierAllowlistReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyAllowlistUpdates, chain, contract.FunctionInput[[]committee_verifier.AllowlistConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(committeeVerifier.Implementation),
				Args:          committeeVerifierAllowlistArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allowlist updates to CommitteeVerifier(%s) on chain %s: %w", committeeVerifier, chain, err)
			}
			writes = append(writes, committeeVerifierAllowlistReport.Output)
		}

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

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
