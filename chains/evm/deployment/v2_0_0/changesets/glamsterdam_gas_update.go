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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/glamsterdam"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// GlamsterdamGasUpdateCfg is configuration for the UpdateGasConfigForGlamsterdamV2 changeset.
type GlamsterdamGasUpdateCfg struct {
	// TargetChainSelector is the chain selector of the chain moving to Glamsterdam.
	TargetChainSelector uint64
	// SkipChainSelectors are chain selectors to unconditionally skip — no lane-to-target check is
	// even performed. Used to batch a high-fanout chain (e.g. mainnet, ~80 lanes) into smaller
	// runs.
	SkipChainSelectors []uint64
}

// UpdateGasConfigForGlamsterdamV2 discovers every v2.0 chain with a lane pointed at
// cfg.TargetChainSelector, reads the current on-chain gas config for every field in the v2.0
// Glamsterdam mapping table, resolves each field against its expected Prague baseline, and
// packages the resulting writes into an MCMS timelock proposal. It never executes directly, even
// if the deployer key happens to be owner.
func UpdateGasConfigForGlamsterdamV2(mcmsRegistry *cs_core.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[cs_core.WithMCMS[GlamsterdamGasUpdateCfg]] {
	validate := func(_ cldf_deployment.Environment, cfg cs_core.WithMCMS[GlamsterdamGasUpdateCfg]) error {
		if cfg.Cfg.TargetChainSelector == 0 {
			return errors.New("target chain selector must be set")
		}
		return nil
	}

	apply := func(e cldf_deployment.Environment, cfg cs_core.WithMCMS[GlamsterdamGasUpdateCfg]) (cldf_deployment.ChangesetOutput, error) {
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

			fqRef := datastore_utils.GetAddressRef(addrs, sel, fee_quoter.ContractType, fee_quoter.Version, "")
			if datastore_utils.IsAddressRefEmpty(fqRef) {
				report.AddUnresolvedContract(sel, "FeeQuoter")
				continue
			}
			feeQuoterAddrByChain[sel] = common.HexToAddress(fqRef.Address)
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

		var (
			lanes        []glamsterdamseq.LaneAddresses
			lombardPools []glamsterdamseq.TokenPoolLane
			usdcPools    []glamsterdamseq.TokenPoolLane
		)

		for _, sel := range discoveryReport.Output.LanesToUpdate {
			addrs := addressesByChain[sel]

			onRampRef := datastore_utils.GetAddressRef(addrs, sel, onramp.ContractType, onramp.Version, "")
			if datastore_utils.IsAddressRefEmpty(onRampRef) {
				report.AddUnresolvedContract(sel, "OnRamp")
				continue
			}
			cvRef := datastore_utils.GetAddressRef(addrs, sel, committee_verifier.ContractType, committee_verifier.Version, "")
			if datastore_utils.IsAddressRefEmpty(cvRef) {
				report.AddUnresolvedContract(sel, "CommitteeVerifier")
				continue
			}

			lane := glamsterdamseq.LaneAddresses{
				ChainSelector:            sel,
				OnRampAddress:            common.HexToAddress(onRampRef.Address),
				FeeQuoterAddress:         feeQuoterAddrByChain[sel],
				CommitteeVerifierAddress: common.HexToAddress(cvRef.Address),
			}
			if offRampRef := datastore_utils.GetAddressRef(addrs, sel, offramp.ContractType, offramp.Version, ""); !datastore_utils.IsAddressRefEmpty(offRampRef) {
				lane.OffRampAddress = common.HexToAddress(offRampRef.Address)
			}
			lanes = append(lanes, lane)

			if lombardRef := datastore_utils.GetAddressRef(addrs, sel, lombard_token_pool.ContractType, lombard_token_pool.Version, ""); !datastore_utils.IsAddressRefEmpty(lombardRef) {
				lombardPools = append(lombardPools, glamsterdamseq.TokenPoolLane{
					ChainSelector: sel,
					PoolAddress:   common.HexToAddress(lombardRef.Address),
				})
			}
			if usdcRef := datastore_utils.GetAddressRef(addrs, sel, siloed_usdc_token_pool.ContractType, siloed_usdc_token_pool.Version, ""); !datastore_utils.IsAddressRefEmpty(usdcRef) {
				usdcPools = append(usdcPools, glamsterdamseq.TokenPoolLane{
					ChainSelector: sel,
					PoolAddress:   common.HexToAddress(usdcRef.Address),
				})
			}
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

		if len(lombardPools) > 0 || len(usdcPools) > 0 {
			tokenPoolReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateTokenPoolGasConfig, e.BlockChains, glamsterdamseq.UpdateTokenPoolGasConfigInput{
				TargetChainSelector: target,
				LombardPools:        lombardPools,
				USDCPools:           usdcPools,
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to update token pool gas config for target chain %d: %w", target, err)
			}
			batchOps = append(batchOps, tokenPoolReport.Output.BatchOps...)
			report.Lines = append(report.Lines, tokenPoolReport.Output.Report.Lines...)
		}

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
