package changesets

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
	tar_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// getAllConfiguredTokensArgs is the input to getAllConfiguredTokens.
type getAllConfiguredTokensArgs struct {
	StartIndex uint64
	MaxCount   uint64
}

// getAllConfiguredTokensMaxCount is large enough to fetch every configured token on a chain in a
// single call for any realistic v1.6 lane fan-out.
const getAllConfiguredTokensMaxCount = 1000

// getAllConfiguredTokens reads every token TokenAdminRegistry knows about on a chain. Used to
// build the candidate token list for row 5 of the v1.6 Glamsterdam mapping table
// (FeeQuoter.TokenTransferFeeConfig.DestGasOverhead, keyed by (destChainSelector, token)).
var getAllConfiguredTokens = contract.NewRead(contract.ReadParams[getAllConfiguredTokensArgs, []common.Address, *tar_bindings.TokenAdminRegistry]{
	Name:         "glamsterdam:token-admin-registry:get-all-configured-tokens",
	Version:      token_admin_registry.Version,
	Description:  "Calls getAllConfiguredTokens on TokenAdminRegistry",
	ContractType: token_admin_registry.ContractType,
	NewContract:  tar_bindings.NewTokenAdminRegistry,
	CallContract: func(c *tar_bindings.TokenAdminRegistry, opts *bind.CallOpts, args getAllConfiguredTokensArgs) ([]common.Address, error) {
		return c.GetAllConfiguredTokens(opts, args.StartIndex, args.MaxCount)
	},
})

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

			tarRef := datastore_utils.GetAddressRef(addrs, sel, token_admin_registry.ContractType, token_admin_registry.Version, "")
			if datastore_utils.IsAddressRefEmpty(tarRef) {
				report.AddUnresolvedContract(sel, "TokenAdminRegistry")
				continue
			}

			chain, ok := e.BlockChains.EVMChains()[sel]
			if !ok {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("chain with selector %d not found", sel)
			}
			tokensReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, getAllConfiguredTokens, chain, contract.FunctionInput[getAllConfiguredTokensArgs]{
				ChainSelector: sel,
				Address:       common.HexToAddress(tarRef.Address),
				Args:          getAllConfiguredTokensArgs{StartIndex: 0, MaxCount: getAllConfiguredTokensMaxCount},
			})
			if err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to read configured tokens for src %d: %w", sel, err)
			}
			if len(tokensReport.Output) > 0 {
				tokenLanes = append(tokenLanes, glamsterdamseq.TokenTransferFeeConfigLane{
					ChainSelector:    sel,
					FeeQuoterAddress: feeQuoterAddrByChain[sel],
					CandidateTokens:  tokensReport.Output,
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
