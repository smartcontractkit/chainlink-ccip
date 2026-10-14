package changesets

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// resolveUSDCTokenPoolRef finds the chain's non-canonical USDC token pool address ref, if any.
// v1.6.1 has no USDC-specific ContractType; by this version's convention (see
// adapters/non_canonical_usdc_chain.go), the USDC pool is deployed as a
// BurnMintWithLockReleaseFlagTokenPool, and there is at most one per chain, so the first match by
// type is returned regardless of version or qualifier.
func resolveUSDCTokenPoolRef(addrs []datastore.AddressRef, sel uint64) datastore.AddressRef {
	for _, ref := range addrs {
		if ref.ChainSelector == sel && ref.Type == datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType) {
			return ref
		}
	}
	return datastore.AddressRef{}
}

// feeQuoterUse classifies the FeeQuoter a chain's v1.6 OnRamp actually prices messages with.
type feeQuoterUse int

const (
	// feeQuoterV16 is a 1.6.x FeeQuoter, which this changeset's v1.6 ABI can read and write.
	feeQuoterV16 feeQuoterUse = iota
	// feeQuoterOtherVersion is a FeeQuoter of another version (typically 2.0.0, after an in-place
	// upgrade). It has a different DestChainConfig struct/selectors, so this changeset cannot write it;
	// its gas config is migrated by the v2.0 changeset.
	feeQuoterOtherVersion
	// feeQuoterUnknown is a FeeQuoter address that is not in the datastore at all.
	feeQuoterUnknown
)

// classifyUsedFeeQuoter looks up the FeeQuoter address that the OnRamp points at in the chain's
// datastore entries and reports whether it is a 1.6.x FeeQuoter this changeset can update, and its
// datastore version (empty if unknown).
func classifyUsedFeeQuoter(addrs []datastore.AddressRef, sel uint64, used common.Address) (feeQuoterUse, string) {
	for _, ref := range addrs {
		if ref.ChainSelector != sel || ref.Type != datastore.ContractType(fee_quoter.ContractType) {
			continue
		}
		if common.HexToAddress(ref.Address) != used {
			continue
		}
		if ref.Version.Major() == 1 && ref.Version.Minor() == 6 {
			return feeQuoterV16, ref.Version.String()
		}
		return feeQuoterOtherVersion, ref.Version.String()
	}
	return feeQuoterUnknown, ""
}

// resolveFeeQuoter returns the FeeQuoter to update for a chain, or false (with a report line) if the
// chain must be skipped. The FeeQuoter that matters is the one the chain's v1.6 OnRamp is wired to
// (dynamicConfig.feeQuoter), NOT necessarily the datastore's FeeQuoter 1.6.0: on a chain whose
// FeeQuoter was upgraded in place the OnRamp uses the 2.0.0 one, and updating the old 1.6.0 would
// change a contract nothing prices with. If the chain has no v1.6 OnRamp in the datastore, fall back
// to the datastore's FeeQuoter 1.6.0.
func resolveFeeQuoter(e cldf_deployment.Environment, sel uint64, addrs []datastore.AddressRef, report *glamsterdamutils.Report) (common.Address, bool) {
	onRampRef := datastore_utils.GetAddressRef(addrs, sel, onramp.ContractType, onramp.Version, "")
	if datastore_utils.IsAddressRefEmpty(onRampRef) {
		fqRef := datastore_utils.GetAddressRef(addrs, sel, fee_quoter.ContractType, fee_quoter.Version, "")
		if datastore_utils.IsAddressRefEmpty(fqRef) {
			report.AddUnresolvedContract(sel, "FeeQuoter")
			return common.Address{}, false
		}
		return common.HexToAddress(fqRef.Address), true
	}

	chain, ok := e.BlockChains.EVMChains()[sel]
	if !ok {
		report.AddReadError(sel, "find chain for OnRamp dynamic config read", fmt.Errorf("chain %d not loaded", sel))
		return common.Address{}, false
	}
	dyn, err := cldf_ops.ExecuteOperation(e.OperationsBundle, onramp.GetDynamicConfig, chain, contract.FunctionInput[struct{}]{
		ChainSelector: sel,
		Address:       common.HexToAddress(onRampRef.Address),
	})
	if err != nil {
		report.AddReadError(sel, "read v1.6 OnRamp dynamic config", err)
		return common.Address{}, false
	}

	used := dyn.Output.FeeQuoter
	switch kind, version := classifyUsedFeeQuoter(addrs, sel, used); kind {
	case feeQuoterV16:
		return used, true
	case feeQuoterOtherVersion:
		line := fmt.Sprintf(
			"chain %d: v1.6 OnRamp %s is wired to FeeQuoter %s (%s), not a 1.6.x FeeQuoter - its gas config is "+
				"migrated by the v2.0 changeset, skipping this chain",
			sel, onRampRef.Address, used, version,
		)
		e.Logger.Info(line)
		report.AddLine(line)
	default:
		line := fmt.Sprintf(
			"chain %d: ERROR - v1.6 OnRamp %s is wired to FeeQuoter %s, which is not in the datastore, skipping this chain",
			sel, onRampRef.Address, used,
		)
		e.Logger.Warn(line)
		report.AddLine(line)
	}
	return common.Address{}, false
}

// GlamsterdamGasUpdateV16Cfg is configuration for the UpdateGasConfigForGlamsterdamV16 changeset.
type GlamsterdamGasUpdateV16Cfg struct {
	// TargetChainSelector is the chain selector of the chain moving to Glamsterdam.
	TargetChainSelector uint64
	// SkipChainSelectors are chain selectors to unconditionally skip — no lane-to-target check is
	// even performed. Used to batch a high-fanout chain (e.g. mainnet, ~80 lanes) into smaller
	// runs.
	SkipChainSelectors []uint64
}

// UpdateGasConfigForGlamsterdamV16 discovers every v1.6 chain with a lane pointed at
// cfg.TargetChainSelector, reads the current on-chain gas config for every field in the v1.6
// Glamsterdam mapping table, resolves each field against its expected Prague baseline, and
// packages the resulting writes into an MCMS timelock proposal. It never executes directly, even
// if the deployer key happens to be owner.
func UpdateGasConfigForGlamsterdamV16(mcmsRegistry *cs_core.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[cs_core.WithMCMS[GlamsterdamGasUpdateV16Cfg]] {
	validate := func(_ cldf_deployment.Environment, cfg cs_core.WithMCMS[GlamsterdamGasUpdateV16Cfg]) error {
		if cfg.Cfg.TargetChainSelector == 0 {
			return errors.New("target chain selector must be set")
		}
		return nil
	}

	apply := func(e cldf_deployment.Environment, cfg cs_core.WithMCMS[GlamsterdamGasUpdateV16Cfg]) (cldf_deployment.ChangesetOutput, error) {
		target := cfg.Cfg.TargetChainSelector
		report := glamsterdamutils.NewReport()

		for _, sel := range cfg.Cfg.SkipChainSelectors {
			report.AddSkipped(sel)
		}

		excluded := make([]uint64, 0, len(cfg.Cfg.SkipChainSelectors)+1)
		excluded = append(excluded, cfg.Cfg.SkipChainSelectors...)
		excluded = append(excluded, target)

		candidates := e.BlockChains.ListChainSelectors(
			cldf_chain.WithFamily(chain_selectors.FamilyEVM),
			cldf_chain.WithChainSelectorsExclusion(excluded),
		)

		addressesByChain := make(map[uint64][]datastore.AddressRef, len(candidates))
		feeQuoterAddrByChain := make(map[uint64]common.Address, len(candidates))
		for _, sel := range candidates {
			addrs := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(sel))
			addressesByChain[sel] = addrs

			fqAddr, ok := resolveFeeQuoter(e, sel, addrs, report)
			if !ok {
				continue
			}
			feeQuoterAddrByChain[sel] = fqAddr
		}

		discoveryReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.DiscoverLanesToTarget, e.BlockChains, glamsterdamseq.DiscoverLanesToTargetInput{
			TargetChainSelector:     target,
			FeeQuoterAddressByChain: feeQuoterAddrByChain,
		})
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to discover lanes to target chain %d: %w", target, err)
		}
		for _, sel := range discoveryReport.Output.NoLane {
			report.AddNoLane(sel)
		}
		report.Lines = append(report.Lines, discoveryReport.Output.Report.Lines...)

		var (
			lanes      []glamsterdamseq.LaneAddresses
			tokenLanes []glamsterdamseq.TokenTransferFeeConfigLane
		)

		for _, sel := range discoveryReport.Output.LanesToUpdate {
			addrs := addressesByChain[sel]

			lane := glamsterdamseq.LaneAddresses{
				ChainSelector:    sel,
				FeeQuoterAddress: feeQuoterAddrByChain[sel],
			}
			if offRampRef := datastore_utils.GetAddressRef(addrs, sel, offramp.ContractType, offramp.Version, ""); !datastore_utils.IsAddressRefEmpty(offRampRef) {
				lane.OffRampAddress = common.HexToAddress(offRampRef.Address)
			}
			lanes = append(lanes, lane)

			usdcPoolRef := resolveUSDCTokenPoolRef(addrs, sel)
			if datastore_utils.IsAddressRefEmpty(usdcPoolRef) {
				// No non-canonical USDC pool deployed on this chain — nothing to update for the
				// USDC-specific row of the v1.6 mapping table.
				continue
			}

			chain, ok := e.BlockChains.EVMChains()[sel]
			if !ok {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain with selector %d not found", sel)
			}
			usdcTokenReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, token_pool.GetToken, chain, contract.FunctionInput[struct{}]{
				ChainSelector: sel,
				Address:       common.HexToAddress(usdcPoolRef.Address),
				Args:          struct{}{},
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to read underlying token for USDC pool %s on src %d: %w", usdcPoolRef.Address, sel, err)
			}
			tokenLanes = append(tokenLanes, glamsterdamseq.TokenTransferFeeConfigLane{
				ChainSelector:    sel,
				FeeQuoterAddress: feeQuoterAddrByChain[sel],
				CandidateTokens:  []common.Address{usdcTokenReport.Output},
			})
		}

		var batchOps []mcms_types.BatchOperation

		if len(lanes) > 0 {
			gasCfgReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateGasConfig, e.BlockChains, glamsterdamseq.UpdateGasConfigInput{
				TargetChainSelector: target,
				Lanes:               lanes,
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to update gas config for target chain %d: %w", target, err)
			}
			batchOps = append(batchOps, gasCfgReport.Output.BatchOps...)
			report.Lines = append(report.Lines, gasCfgReport.Output.Report.Lines...)
		}

		if len(tokenLanes) > 0 {
			ttfcReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateTokenTransferFeeConfig, e.BlockChains, glamsterdamseq.UpdateTokenTransferFeeConfigInput{
				TargetChainSelector: target,
				Lanes:               tokenLanes,
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to update token transfer fee config for target chain %d: %w", target, err)
			}
			batchOps = append(batchOps, ttfcReport.Output.BatchOps...)
			report.Lines = append(report.Lines, ttfcReport.Output.Report.Lines...)
		}

		// The report is also attached to the proposal description, but when every chain is skipped there is
		// no proposal, so always log it as well.
		e.Logger.Infof("Glamsterdam v1.6 gas config report (%d batch operations):\n%s", len(batchOps), report.String())

		mcmsInput := cfg.MCMS
		if mcmsInput.Description == "" {
			mcmsInput.Description = report.String()
		} else {
			mcmsInput.Description = mcmsInput.Description + "\n\n" + report.String()
		}

		return cs_core.NewOutputBuilder(e, mcmsRegistry).WithBatchOps(batchOps).Build(mcmsInput)
	}

	return cldf_deployment.CreateChangeSet(apply, validate)
}
