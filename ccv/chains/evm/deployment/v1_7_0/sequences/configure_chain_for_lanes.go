package sequences

import (
	"bytes"
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
	seqtypes "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	changesetadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/proxy"
	fqc "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
)

// ConfigureChainForLanes is the canonical sequence for configuring an EVM chain to participate
// in CCIP 2.0 lanes with multiple remote chains. It is self-contained: all contract writes
// (OffRamp, OnRamp, Executor, FeeQuoter, CommitteeVerifier, Router) are handled here.
//
// Every write operation is idempotent — current on-chain state is read first and a write is
// only emitted when the desired state differs. This matters because the output BatchOps may
// be submitted as an MCMS proposal; redundant transactions would waste gas and clutter the
// proposal.
//
// Execution order is deliberate for safety:
//  1. OffRamp, OnRamp, Executor, FeeQuoter — infrastructure contracts that must be configured
//     before traffic can flow.
//  2. CommitteeVerifier — source-side outbound config and dest-side signature/inbound config.
//  3. Router (last) — wiring the router enables traffic. By doing this last we guarantee that
//     if any earlier step fails, the router is not wired to an incompletely-configured lane.
//
// The router write is placed in a separate BatchOperation to preserve ordering in MCMS
// proposals (batch N executes after batch N-1).
var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("2.0.0"),
	"Configures an EVM chain as a source and destination for multiple remote chains",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input changesetadapters.ConfigureChainForLanesInput) (seqtypes.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return seqtypes.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// When AllowOnrampOverride is false, refuse to replace an existing
		// OnRamp mapping in the Router with a different OnRamp address. This
		// prevents accidental overwrites of prod router state — use the
		// migration changeset to swap OnRamp versions. Switching which
		// router the OnRamp/OffRamp points to (e.g. test router to prod
		// router promotion) is always allowed.
		if !input.AllowOnrampOverride {
			for remoteSelector := range input.RemoteChains {
				existing, err := cldf_ops.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(input.Router),
					Args:          remoteSelector,
				})
				if err != nil {
					return seqtypes.OnChainOutput{}, fmt.Errorf(
						"failed to read onRamp for dest %d from Router(%s): %w",
						remoteSelector, input.Router, err,
					)
				}
				if existing.Output != (common.Address{}) && existing.Output != common.HexToAddress(input.OnRamp) {
					return seqtypes.OnChainOutput{}, fmt.Errorf(
						"router %s already has onRamp %s for dest chain %d; "+
							"refusing to overwrite with %s (AllowOnrampOverride is false) -- "+
							"use the migration changeset to update router mappings",
						input.Router, existing.Output.Hex(), remoteSelector, input.OnRamp,
					)
				}
			}
		}

		// ── Phase 1: Collect desired state per remote chain ──────────────────────
		// For each remote chain we build the desired args for every contract.
		// The "maybe" helpers read on-chain state and only append when a diff exists
		// (idempotency). This avoids emitting no-op transactions in MCMS proposals.
		offRampArgs := make([]offramp.SourceChainConfigArgs, 0, len(input.RemoteChains))
		onRampArgs := make([]onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter.DestChainConfigArgs, 0, len(input.RemoteChains))
		gasPriceUpdates := make([]fee_quoter.GasPriceUpdate, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		destChainSelectorsPerExecutor := make(map[common.Address][]ExecutorRemoteChainConfigArgs)

		feeQContract, err := fqc.NewFeeQuoter(common.HexToAddress(input.FeeQuoter), chain.Client)
		if err != nil {
			return seqtypes.OnChainOutput{}, fmt.Errorf("failed to bind fee quoter contract at address %s on chain %s: %w", input.FeeQuoter, chain.String(), err)
		}

		for remoteSelector, remoteConfig := range input.RemoteChains {
			// OffRamp: tells the local OffRamp which source chains to accept messages from.
			offRampArgs, err = maybeAddSourceChainConfigArgOnLocalChain(b, chain, input, remoteSelector, remoteConfig, offRampArgs)
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("remote chain %d: %w", remoteSelector, err)
			}
			// OnRamp: tells the local OnRamp how to send messages to this remote chain.
			onRampArgs, err = maybeAddOnRampDestChainConfigArgOnLocalChain(b, chain, input, remoteSelector, remoteConfig, onRampArgs)
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("remote chain %d: %w", remoteSelector, err)
			}

			if remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas != nil {
				gasPriceReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestinationChainGasPrice, chain, contract.FunctionInput[uint64]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(input.FeeQuoter),
					Args:          remoteSelector,
				})
				if err != nil {
					return seqtypes.OnChainOutput{}, fmt.Errorf("failed to get gas prices on FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
				}
				if remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas.Cmp(gasPriceReport.Output.Value) != 0 {
					gasPriceUpdates = append(gasPriceUpdates, fee_quoter.GasPriceUpdate{
						DestChainSelector: remoteSelector,
						UsdPerUnitGas:     remoteConfig.FeeQuoterDestChainConfig.USDPerUnitGas,
					})
				}
			}

			// Router OnRamp: only add if the router doesn't already point to our OnRamp
			// for this destination.
			onRampAddrReport, err := cldf_ops.ExecuteOperation(b, router.GetOnRamp, chain, contract.FunctionInput[uint64]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.Router),
				Args:          remoteSelector,
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to get on ramp for dest %d from Router(%s) on chain %s: %w", remoteSelector, input.Router, chain, err)
			}
			if onRampAddrReport.Output != common.HexToAddress(input.OnRamp) {
				onRampAdds = append(onRampAdds, router.OnRamp{
					DestChainSelector: remoteSelector,
					OnRamp:            common.HexToAddress(input.OnRamp),
				})
			}

			// Router OffRamp: always add — duplicates are filtered out in bulk below
			// via FilterOffRampAdds (one RPC call for all remotes).
			offRampAdds = append(offRampAdds, router.OffRamp{
				SourceChainSelector: remoteSelector,
				OffRamp:             common.HexToAddress(input.OffRamp),
			})

			// Executor: the input references the proxy address; we resolve through the
			// proxy to group dest chains by their actual implementation, since multiple
			// proxies may point to the same implementation.
			defaultExecutor := common.HexToAddress(remoteConfig.DefaultExecutor)
			getTargetReport, err := cldf_ops.ExecuteOperation(b, proxy.GetTarget, chain, contract.FunctionInput[struct{}]{
				ChainSelector: chain.Selector,
				Address:       defaultExecutor,
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to get target address of Executor(%s) on chain %s: %w", defaultExecutor, chain, err)
			}
			destChainSelectorsPerExecutor[getTargetReport.Output] = append(destChainSelectorsPerExecutor[getTargetReport.Output], ExecutorRemoteChainConfigArgs{
				DestChainSelector: remoteSelector,
				Config:            remoteConfig.ExecutorDestChainConfig,
			})

			// FeeQuoter dest chain config: when OverrideExistingConfig is false, we skip
			// chains that already have an enabled config to avoid accidentally overwriting
			// production parameters. When true, we always diff and update.
			if !remoteConfig.FeeQuoterDestChainConfig.OverrideExistingConfig {
				destChainCfg, err := feeQContract.GetDestChainConfig(&bind.CallOpts{Context: b.GetContext()}, remoteSelector)
				if err != nil {
					return seqtypes.OnChainOutput{}, fmt.Errorf("failed to get dest chain config for remote chain selector %d from fee quoter at address %s on chain %s: %w", remoteSelector, input.FeeQuoter, chain.String(), err)
				}
				if !destChainCfg.IsEnabled {
					feeQuoterArgs, err = maybeAddFeeQuoterDestChainConfigArgOnLocalChain(feeQContract, b, input.FeeQuoter, chain, remoteSelector, remoteConfig, feeQuoterArgs, &destChainCfg)
					if err != nil {
						return seqtypes.OnChainOutput{}, fmt.Errorf("remote chain %d: %w", remoteSelector, err)
					}
				}
			} else {
				feeQuoterArgs, err = maybeAddFeeQuoterDestChainConfigArgOnLocalChain(feeQContract, b, input.FeeQuoter, chain, remoteSelector, remoteConfig, feeQuoterArgs, nil)
				if err != nil {
					return seqtypes.OnChainOutput{}, fmt.Errorf("remote chain %d: %w", remoteSelector, err)
				}
			}
		}

		// ── Phase 2: Bulk-filter already-configured entries ──────────────────────
		// OffRamp adds and Executor dest chains are filtered in bulk (one RPC each)
		// rather than per-remote-chain, since the contracts expose list-all getters.
		offRampAdds, err = FilterOffRampAdds(b, chain, common.HexToAddress(input.Router), offRampAdds)
		if err != nil {
			return seqtypes.OnChainOutput{}, err
		}

		destChainSelectorsPerExecutor, err = FilterExecutorDestChains(b, chain, destChainSelectorsPerExecutor)
		if err != nil {
			return seqtypes.OnChainOutput{}, err
		}

		// ── Phase 3: Apply writes ────────────────────────────────────────────────
		// Each block only emits a write when there are actual changes to apply.
		if len(offRampArgs) > 0 {
			offRampReport, err := cldf_ops.ExecuteOperation(b, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.OffRamp),
				Args:          offRampArgs,
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to apply source chain config updates to OffRamp(%s) on chain %s: %w", input.OffRamp, chain, err)
			}
			writes = append(writes, offRampReport.Output)
		}

		if len(onRampArgs) > 0 {
			onRampReport, err := cldf_ops.ExecuteOperation(b, onramp.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.OnRamp),
				Args:          onRampArgs,
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to OnRamp(%s) on chain %s: %w", input.OnRamp, chain, err)
			}
			writes = append(writes, onRampReport.Output)
		}

		for executorAddr, toAdd := range destChainSelectorsPerExecutor {
			if len(toAdd) == 0 {
				continue
			}
			executorReport, err := cldf_ops.ExecuteOperation(b, ExecutorApplyDestChainUpdates, chain, contract.FunctionInput[ExecutorApplyDestChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       executorAddr,
				Args: ExecutorApplyDestChainUpdatesArgs{
					DestChainSelectorsToAdd: toAdd,
				},
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to Executor(%s) on chain %s: %w", executorAddr, chain, err)
			}
			writes = append(writes, executorReport.Output)
		}

		if len(feeQuoterArgs) > 0 {
			feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.FeeQuoter),
				Args:          feeQuoterArgs,
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
			}
			writes = append(writes, feeQuoterReport.Output)
		}

		// Gas price updates live in a separate on-chain mapping (not dest chain config),
		// so they must be applied via FeeQuoter.UpdatePrices.
		if len(gasPriceUpdates) > 0 {
			gasPriceReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract.FunctionInput[fee_quoter.PriceUpdates]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.FeeQuoter),
				Args: fee_quoter.PriceUpdates{
					GasPriceUpdates: gasPriceUpdates,
				},
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to update gas prices on FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
			}
			writes = append(writes, gasPriceReport.Output)
		}

		// CommitteeVerifier: for each verifier contract, configure outbound settings
		// (remote chain config, allowlist, resolver routing) and inbound settings
		// (signature quorum config, resolver version registration). Each verifier may
		// serve a different set of remote chains via its own RemoteChains map.
		for _, cv := range input.CommitteeVerifiers {
			cvWrites, err := configureCommitteeVerifierAsSource(b, chain, input.Router, input.ChainSelector, cv)
			if err != nil {
				return seqtypes.OnChainOutput{}, err
			}
			writes = append(writes, cvWrites...)

			cvDestWrites, err := configureCommitteeVerifierAsDest(b, chain, input.ChainSelector, cv)
			if err != nil {
				return seqtypes.OnChainOutput{}, err
			}
			writes = append(writes, cvDestWrites...)
		}

		// ── Phase 4: Build BatchOps ─────────────────────────────────────────────
		batchOps := make([]mcms_types.BatchOperation, 0, 2)
		if len(writes) > 0 {
			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			batchOps = append(batchOps, batchOp)
		}

		if len(onRampAdds) > 0 || len(offRampAdds) > 0 {
			routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, contract.FunctionInput[router.ApplyRampsUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(input.Router),
				Args: router.ApplyRampsUpdatesArgs{
					OnRampUpdates:  onRampAdds,
					OffRampRemoves: []router.OffRamp{},
					OffRampAdds:    offRampAdds,
				},
			})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", input.Router, chain, err)
			}
			routerBatchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{routerReport.Output})
			if err != nil {
				return seqtypes.OnChainOutput{}, fmt.Errorf("failed to create batch operation for router writes: %w", err)
			}
			batchOps = append(batchOps, routerBatchOp)
		}

		return seqtypes.OnChainOutput{BatchOps: batchOps}, nil
	},
)

// maybeAddSourceChainConfigArgOnLocalChain reads the current OffRamp source chain config
// and appends to offRampArgs only when the desired state differs. Zero/empty fields in the
// input are treated as "keep current" so callers only need to specify fields they want to change.
func maybeAddSourceChainConfigArgOnLocalChain(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input changesetadapters.ConfigureChainForLanesInput,
	remoteSelector uint64,
	remoteConfig changesetadapters.RemoteChainConfig[[]byte, string],
	offRampArgs []offramp.SourceChainConfigArgs,
) ([]offramp.SourceChainConfigArgs, error) {
	defaultInboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultInboundCCVs))
	for _, ccv := range remoteConfig.DefaultInboundCCVs {
		defaultInboundCCVs = append(defaultInboundCCVs, common.HexToAddress(ccv))
	}
	laneMandatedInboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedInboundCCVs))
	for _, ccv := range remoteConfig.LaneMandatedInboundCCVs {
		laneMandatedInboundCCVs = append(laneMandatedInboundCCVs, common.HexToAddress(ccv))
	}
	onRamps := make([][]byte, 0, len(remoteConfig.OnRamps))
	for _, onRampAddress := range remoteConfig.OnRamps {
		onRamps = append(onRamps, common.LeftPadBytes(onRampAddress, 32))
	}

	currentReport, err := cldf_ops.ExecuteOperation(b, offramp.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       common.HexToAddress(input.OffRamp),
		Args:          remoteSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get source chain config for selector %d from OffRamp(%s) on chain %v: %w", remoteSelector, input.OffRamp, chain, err)
	}
	current := currentReport.Output

	isEnabled := current.IsEnabled
	if remoteConfig.AllowTrafficFrom != nil {
		isEnabled = *remoteConfig.AllowTrafficFrom
	}

	desired := offramp.SourceChainConfigArgs{
		Router:              common.HexToAddress(input.Router),
		SourceChainSelector: remoteSelector,
		IsEnabled:           isEnabled,
		OnRamps:             onRamps,
		DefaultCCVs:         defaultInboundCCVs,
		LaneMandatedCCVs:    laneMandatedInboundCCVs,
	}
	if len(desired.OnRamps) == 0 {
		desired.OnRamps = current.OnRamps
	}
	if len(desired.DefaultCCVs) == 0 {
		desired.DefaultCCVs = current.DefaultCCVs
	}
	if len(desired.LaneMandatedCCVs) == 0 {
		desired.LaneMandatedCCVs = current.LaneMandatedCCVs
	}

	if current.IsEnabled != desired.IsEnabled ||
		current.Router != desired.Router ||
		!UnorderedSliceEqual(current.OnRamps, desired.OnRamps, bytes.Equal) ||
		!UnorderedSliceEqual(current.DefaultCCVs, desired.DefaultCCVs, func(x, y common.Address) bool { return x == y }) ||
		!UnorderedSliceEqual(current.LaneMandatedCCVs, desired.LaneMandatedCCVs, func(x, y common.Address) bool { return x == y }) {
		offRampArgs = append(offRampArgs, desired)
	}
	return offRampArgs, nil
}

// maybeAddOnRampDestChainConfigArgOnLocalChain reads current OnRamp dest chain config
// and appends to onRampArgs only when the desired state differs. Same zero-means-keep-current
// semantics as the OffRamp helper.
func maybeAddOnRampDestChainConfigArgOnLocalChain(
	b cldf_ops.Bundle,
	chain evm.Chain,
	input changesetadapters.ConfigureChainForLanesInput,
	remoteSelector uint64,
	remoteConfig changesetadapters.RemoteChainConfig[[]byte, string],
	onRampArgs []onramp.DestChainConfigArgs,
) ([]onramp.DestChainConfigArgs, error) {
	defaultOutboundCCVs := make([]common.Address, 0, len(remoteConfig.DefaultOutboundCCVs))
	for _, ccv := range remoteConfig.DefaultOutboundCCVs {
		defaultOutboundCCVs = append(defaultOutboundCCVs, common.HexToAddress(ccv))
	}
	laneMandatedOutboundCCVs := make([]common.Address, 0, len(remoteConfig.LaneMandatedOutboundCCVs))
	for _, ccv := range remoteConfig.LaneMandatedOutboundCCVs {
		laneMandatedOutboundCCVs = append(laneMandatedOutboundCCVs, common.HexToAddress(ccv))
	}

	desired := onramp.DestChainConfigArgs{
		Router:                    common.HexToAddress(input.Router),
		DestChainSelector:         remoteSelector,
		AddressBytesLength:        remoteConfig.AddressBytesLength,
		BaseExecutionGasCost:      remoteConfig.BaseExecutionGasCost,
		MessageNetworkFeeUSDCents: remoteConfig.MessageNetworkFeeUSDCents,
		TokenNetworkFeeUSDCents:   remoteConfig.TokenNetworkFeeUSDCents,
		DefaultCCVs:               defaultOutboundCCVs,
		LaneMandatedCCVs:          laneMandatedOutboundCCVs,
		DefaultExecutor:           common.HexToAddress(remoteConfig.DefaultExecutor),
		OffRamp:                   remoteConfig.OffRamp,
	}
	currentReport, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       common.HexToAddress(input.OnRamp),
		Args:          remoteSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get dest chain config for selector %d from OnRamp(%s) on chain %v: %w", remoteSelector, input.OnRamp, chain, err)
	}
	current := currentReport.Output

	if remoteConfig.TokenReceiverAllowed == nil {
		desired.TokenReceiverAllowed = current.TokenReceiverAllowed
	} else {
		desired.TokenReceiverAllowed = *remoteConfig.TokenReceiverAllowed
	}
	if desired.MessageNetworkFeeUSDCents == 0 {
		desired.MessageNetworkFeeUSDCents = current.MessageNetworkFeeUSDCents
	}
	if desired.TokenNetworkFeeUSDCents == 0 {
		desired.TokenNetworkFeeUSDCents = current.TokenNetworkFeeUSDCents
	}
	if desired.BaseExecutionGasCost == 0 {
		desired.BaseExecutionGasCost = current.BaseExecutionGasCost
	}
	if desired.AddressBytesLength == 0 {
		desired.AddressBytesLength = current.AddressBytesLength
	}
	if len(desired.DefaultCCVs) == 0 {
		desired.DefaultCCVs = current.DefaultCCVs
	}
	if len(desired.LaneMandatedCCVs) == 0 {
		desired.LaneMandatedCCVs = current.LaneMandatedCCVs
	}

	if current.Router != desired.Router || current.DefaultExecutor != desired.DefaultExecutor ||
		!bytes.Equal(current.OffRamp, desired.OffRamp) ||
		current.TokenReceiverAllowed != desired.TokenReceiverAllowed ||
		current.MessageNetworkFeeUSDCents != desired.MessageNetworkFeeUSDCents ||
		current.TokenNetworkFeeUSDCents != desired.TokenNetworkFeeUSDCents ||
		current.BaseExecutionGasCost != desired.BaseExecutionGasCost ||
		current.AddressBytesLength != desired.AddressBytesLength ||
		!UnorderedSliceEqual(current.DefaultCCVs, desired.DefaultCCVs, func(x, y common.Address) bool { return x == y }) ||
		!UnorderedSliceEqual(current.LaneMandatedCCVs, desired.LaneMandatedCCVs, func(x, y common.Address) bool { return x == y }) {
		onRampArgs = append(onRampArgs, desired)
	}
	return onRampArgs, nil
}

// USDPerUnitGas is intentionally excluded because it is updated via a separate
// FeeQuoter.UpdatePrices call (gas prices live in a different on-chain mapping).
func feeQuoterDestChainConfigEqualTo(cur fqc.FeeQuoterDestChainConfig, desired changesetadapters.FeeQuoterDestChainConfig) bool {
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

// maybeAddFeeQuoterDestChainConfigArgOnLocalChain compares the desired FeeQuoter dest chain
// config against the current on-chain state and appends to feeQuoterArgs only when they differ.
// Zero fields are filled from on-chain state so partial updates are safe.
// Pass a non-nil prefetchedCur to reuse a previously fetched config and avoid a redundant RPC.
func maybeAddFeeQuoterDestChainConfigArgOnLocalChain(
	feeQContract *fqc.FeeQuoter,
	b cldf_ops.Bundle,
	feeQuoterAddr string,
	chain evm.Chain,
	remoteSelector uint64,
	remoteConfig changesetadapters.RemoteChainConfig[[]byte, string],
	feeQuoterArgs []fee_quoter.DestChainConfigArgs,
	prefetchedCur *fqc.FeeQuoterDestChainConfig,
) ([]fee_quoter.DestChainConfigArgs, error) {
	var cur fqc.FeeQuoterDestChainConfig
	if prefetchedCur != nil {
		cur = *prefetchedCur
	} else {
		fetched, err := feeQContract.GetDestChainConfig(&bind.CallOpts{Context: b.GetContext()}, remoteSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get dest chain config for remote chain selector %d from fee quoter at address %s on chain %s: %w", remoteSelector, feeQuoterAddr, chain.String(), err)
		}
		cur = fetched
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
	if feeQuoterDestChainConfigEqualTo(cur, desired) {
		return feeQuoterArgs, nil
	}

	return append(feeQuoterArgs, fee_quoter.DestChainConfigArgs{
		DestChainSelector: remoteSelector,
		DestChainConfig:   adapterDestChainConfigToFeeQuoterV2(desired),
	}), nil
}

// configureCommitteeVerifierAsSource configures the outbound (source-side) settings for a
// CommitteeVerifier contract. For each remote chain it:
//   - Sets remote chain config (router, fee, allowlist flag) — only if on-chain differs.
//     Fee fields (FeeUSDCents, GasForVerification, PayloadSizeBytes) live in a separate
//     getter (GetFee) so we check both GetRemoteChainConfig and GetFee before deciding.
//   - Computes allowlist diffs (adds/removes) against the current on-chain allowlist.
//   - Registers the verifier as the outbound implementation on the VersionedVerifierResolver,
//     so the resolver knows which verifier to use when sending to each remote chain.
func configureCommitteeVerifierAsSource(
	b cldf_ops.Bundle,
	chain evm.Chain,
	routerAddr string,
	chainSelector uint64,
	cv changesetadapters.CommitteeVerifierConfig[datastore.AddressRef],
) ([]contract.WriteOutput, error) {
	cvAddr, resolverAddr, err := extractCommitteeVerifierAddresses(cv.CommitteeVerifier, chainSelector)
	if err != nil {
		return nil, err
	}

	remoteChainConfigArgs := make([]committee_verifier.RemoteChainConfigArgs, 0, len(cv.RemoteChains))
	allowlistArgs := make([]committee_verifier.AllowlistConfigArgs, 0, len(cv.RemoteChains))

	for remoteSelector, remoteConfig := range cv.RemoteChains {
		desired := committee_verifier.RemoteChainConfigArgs{
			Router:              common.HexToAddress(routerAddr),
			RemoteChainSelector: remoteSelector,
			AllowlistEnabled:    remoteConfig.AllowlistEnabled,
			FeeUSDCents:         remoteConfig.FeeUSDCents,
			GasForVerification:  remoteConfig.GasForVerification,
			PayloadSizeBytes:    remoteConfig.PayloadSizeBytes,
		}
		currentRemoteReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cvAddr),
			Args:          remoteSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get remote chain config for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
		}
		cur := currentRemoteReport.Output

		if cur.Router != desired.Router || cur.AllowlistEnabled != desired.AllowlistEnabled {
			remoteChainConfigArgs = append(remoteChainConfigArgs, desired)
		} else {
			getFeeReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetFee, chain, contract.FunctionInput[committee_verifier.GetFeeArgs]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(cvAddr),
				Args: committee_verifier.GetFeeArgs{
					DestChainSelector: remoteSelector,
				},
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get fee for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
			}
			curFee := getFeeReport.Output
			if curFee.FeeUSDCents != desired.FeeUSDCents ||
				curFee.GasForVerification != desired.GasForVerification ||
				curFee.PayloadSizeBytes != desired.PayloadSizeBytes {
				remoteChainConfigArgs = append(remoteChainConfigArgs, desired)
			}
		}

		toAdd, toRemove, err := makeAllowlistUpdates(cur.AllowedSendersList, remoteConfig.AddedAllowlistedSenders, remoteConfig.RemovedAllowlistedSenders)
		if err != nil {
			return nil, fmt.Errorf("invalid allowlist addresses for remote chain %d: %w", remoteSelector, err)
		}
		if len(toAdd) > 0 || len(toRemove) > 0 {
			allowlistArgs = append(allowlistArgs, committee_verifier.AllowlistConfigArgs{
				AllowlistEnabled:          remoteConfig.AllowlistEnabled,
				AddedAllowlistedSenders:   toAdd,
				RemovedAllowlistedSenders: toRemove,
				DestChainSelector:         remoteSelector,
			})
		}
	}

	currentOutboundReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.GetAllOutboundImplementations, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       common.HexToAddress(resolverAddr),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get outbound implementations from CommitteeVerifierResolver on chain %s: %w", chain, err)
	}
	currentOutbound := make(map[uint64]common.Address)
	for _, o := range currentOutboundReport.Output {
		currentOutbound[o.DestChainSelector] = o.Verifier
	}
	outboundArgs := make([]versioned_verifier_resolver.OutboundImplementationArgs, 0, len(cv.RemoteChains))
	cvEthAddr := common.HexToAddress(cvAddr)
	for remoteSelector := range cv.RemoteChains {
		if currentOutbound[remoteSelector] != cvEthAddr {
			outboundArgs = append(outboundArgs, versioned_verifier_resolver.OutboundImplementationArgs{
				DestChainSelector: remoteSelector,
				Verifier:          cvEthAddr,
			})
		}
	}

	var writes []contract.WriteOutput

	if len(remoteChainConfigArgs) > 0 {
		report, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyRemoteChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.RemoteChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cvAddr),
			Args:          remoteChainConfigArgs,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply remote chain config updates to CommitteeVerifier on chain %s: %w", chain, err)
		}
		writes = append(writes, report.Output)
	}

	if len(allowlistArgs) > 0 {
		report, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplyAllowlistUpdates, chain, contract.FunctionInput[[]committee_verifier.AllowlistConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cvAddr),
			Args:          allowlistArgs,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply allowlist updates to CommitteeVerifier on chain %s: %w", chain, err)
		}
		writes = append(writes, report.Output)
	}

	if len(outboundArgs) > 0 {
		report, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyOutboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.OutboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(resolverAddr),
			Args:          outboundArgs,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply outbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
		}
		writes = append(writes, report.Output)
	}

	return writes, nil
}

// configureCommitteeVerifierAsDest configures the inbound (destination-side) settings for a
// CommitteeVerifier contract. For each remote chain it:
//   - Sets the signature quorum config (signers + threshold) — only if on-chain differs.
//   - Registers the verifier's version tag as an inbound implementation on the
//     VersionedVerifierResolver, so incoming messages can be verified. The version tag is
//     read from the contract itself to ensure consistency.
func configureCommitteeVerifierAsDest(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	cv changesetadapters.CommitteeVerifierConfig[datastore.AddressRef],
) ([]contract.WriteOutput, error) {
	cvAddr, resolverAddr, err := extractCommitteeVerifierAddresses(cv.CommitteeVerifier, chainSelector)
	if err != nil {
		return nil, err
	}

	signatureConfigs := make([]committee_verifier.SignatureConfig, 0, len(cv.RemoteChains))
	for remoteSelector, remoteConfig := range cv.RemoteChains {
		signers := make([]common.Address, 0, len(remoteConfig.SignatureConfig.Signers))
		for _, signer := range remoteConfig.SignatureConfig.Signers {
			signers = append(signers, common.HexToAddress(signer))
		}
		desired := committee_verifier.SignatureConfig{
			SourceChainSelector: remoteSelector,
			Threshold:           remoteConfig.SignatureConfig.Threshold,
			Signers:             signers,
		}
		currentSigReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetSignatureConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cvAddr),
			Args:          remoteSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get signature config for selector %d from CommitteeVerifier on chain %s: %w", remoteSelector, chain, err)
		}
		curSig := currentSigReport.Output
		if curSig.Threshold != desired.Threshold || !UnorderedSliceEqual(curSig.Signers, desired.Signers, func(x, y common.Address) bool { return x == y }) {
			signatureConfigs = append(signatureConfigs, desired)
		}
	}

	var writes []contract.WriteOutput

	if len(signatureConfigs) > 0 {
		report, err := cldf_ops.ExecuteOperation(b, committee_verifier.ApplySignatureConfigs, chain, contract.FunctionInput[committee_verifier.ApplySignatureConfigsArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(cvAddr),
			Args: committee_verifier.ApplySignatureConfigsArgs{
				SignatureConfigs: signatureConfigs,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply signature configs to CommitteeVerifier on chain %s: %w", chain, err)
		}
		writes = append(writes, report.Output)
	}

	versionTagReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.VersionTag, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       common.HexToAddress(cvAddr),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get version tag from CommitteeVerifier on chain %s: %w", chain, err)
	}
	desiredInbound := versioned_verifier_resolver.InboundImplementationArgs{
		Version:  versionTagReport.Output,
		Verifier: common.HexToAddress(cvAddr),
	}
	currentInboundReport, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.GetAllInboundImplementations, chain, contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       common.HexToAddress(resolverAddr),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inbound implementations from CommitteeVerifierResolver on chain %s: %w", chain, err)
	}
	inboundAlreadySet := false
	for _, cur := range currentInboundReport.Output {
		if cur.Version == desiredInbound.Version && cur.Verifier == desiredInbound.Verifier {
			inboundAlreadySet = true
			break
		}
	}
	if !inboundAlreadySet {
		report, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(resolverAddr),
			Args:          []versioned_verifier_resolver.InboundImplementationArgs{desiredInbound},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply inbound implementation updates to CommitteeVerifierResolver on chain %s: %w", chain, err)
		}
		writes = append(writes, report.Output)
	}

	return writes, nil
}

func adapterDestChainConfigToFeeQuoterV2(cfg changesetadapters.FeeQuoterDestChainConfig) fee_quoter.DestChainConfig {
	return fee_quoter.DestChainConfig{
		IsEnabled:                   cfg.IsEnabled,
		MaxDataBytes:                cfg.MaxDataBytes,
		MaxPerMsgGasLimit:           cfg.MaxPerMsgGasLimit,
		DestGasOverhead:             cfg.DestGasOverhead,
		DestGasPerPayloadByteBase:   cfg.DestGasPerPayloadByteBase,
		ChainFamilySelector:         cfg.ChainFamilySelector,
		DefaultTokenFeeUSDCents:     cfg.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead: cfg.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:           cfg.DefaultTxGasLimit,
		NetworkFeeUSDCents:          cfg.NetworkFeeUSDCents,
		LinkFeeMultiplierPercent:    cfg.LinkFeeMultiplierPercent,
	}
}
