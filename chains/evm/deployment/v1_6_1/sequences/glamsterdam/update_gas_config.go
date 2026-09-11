package glamsterdam

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
)

// applyFeeQuoterDestChainConfigUpdates mirrors fee_quoter.ApplyDestChainConfigUpdates, but always
// routes the write through MCMS regardless of who the deployer key is, per this feature's
// "always MCMS" decision.
var applyFeeQuoterDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]fee_quoter.DestChainConfigArgs, *fee_quoter.FeeQuoterContract]{
	Name:            "glamsterdam:fee-quoter:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Calls applyDestChainConfigUpdates on FeeQuoter, always producing an MCMS proposal",
	ContractType:    fee_quoter.ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoterContract,
	IsAllowedCaller: contract.NoCallersAllowed[*fee_quoter.FeeQuoterContract, []fee_quoter.DestChainConfigArgs],
	Validate:        func([]fee_quoter.DestChainConfigArgs) error { return nil },
	CallContract: func(c *fee_quoter.FeeQuoterContract, opts *bind.TransactOpts, args []fee_quoter.DestChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyDestChainConfigUpdates(opts, args)
	},
})

// LaneAddresses is the set of contract addresses to update on one source chain that has a
// confirmed lane pointed at the Glamsterdam target chain.
type LaneAddresses struct {
	ChainSelector    uint64
	FeeQuoterAddress common.Address
	// OffRampAddress is optional. If set, OffRamp's immutable GasForCallExactCheck field is read
	// and sanity checked against its expected Prague baseline (no write is ever made — there is
	// no setter). Leave as the zero address to skip this chain's check.
	OffRampAddress common.Address
}

// UpdateGasConfigInput is the input to UpdateGasConfig.
type UpdateGasConfigInput struct {
	// TargetChainSelector is the chain selector moving to Glamsterdam.
	TargetChainSelector uint64
	// Lanes are the source chains with a confirmed lane pointed at the target chain.
	Lanes []LaneAddresses
}

// UpdateGasConfigOutput is the output of UpdateGasConfig.
type UpdateGasConfigOutput struct {
	// BatchOps contains one MCMS batch operation per chain in Lanes.
	BatchOps []mcms_types.BatchOperation
	// Report is a human-readable summary of every field resolved on every lane, for inclusion in
	// the changeset's MCMS proposal description.
	Report *glamsterdamutils.Report
}

// UpdateGasConfig reads the current on-chain gas config for every confirmed lane, resolves each
// field in the v1.6 Glamsterdam mapping table against its expected Prague baseline (applying the
// literal Glamsterdam value on a match, or the field's fallback rule on a mismatch), and packages
// the resulting write into one MCMS batch operation per chain. Every write is routed through MCMS
// regardless of who the deployer key is.
var UpdateGasConfig = cldf_ops.NewSequence(
	"UpdateGasConfigV1_6",
	semver.MustParse("1.6.0"),
	"Updates v1.6 source-side gas config for lanes pointed at the Glamsterdam target chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input UpdateGasConfigInput) (UpdateGasConfigOutput, error) {
		output := UpdateGasConfigOutput{Report: glamsterdamutils.NewReport()}

		for _, lane := range input.Lanes {
			chain, ok := chains.EVMChains()[lane.ChainSelector]
			if !ok {
				return UpdateGasConfigOutput{}, fmt.Errorf("chain with selector %d not found", lane.ChainSelector)
			}

			var writes []contract.WriteOutput

			// --- FeeQuoter: DestGasOverhead, DefaultTokenDestGasOverhead (rows 1-2) ---
			fqCur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
				Args:          input.TargetChainSelector,
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf(
					"failed to read FeeQuoter dest chain config for src %d, dst %d: %w", lane.ChainSelector, input.TargetChainSelector, err,
				)
			}

			destGasOverheadResult := glamsterdamutils.Resolve(FeeQuoterDestGasOverhead, fqCur.Output.DestGasOverhead)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, destGasOverheadResult)

			defaultTokenDestGasOverheadResult := glamsterdamutils.Resolve(FeeQuoterDefaultTokenDestGasOverhead, fqCur.Output.DefaultTokenDestGasOverhead)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, defaultTokenDestGasOverheadResult)

			newFQConfig := fqCur.Output
			newFQConfig.DestGasOverhead = destGasOverheadResult.AppliedValue
			newFQConfig.DefaultTokenDestGasOverhead = defaultTokenDestGasOverheadResult.AppliedValue

			fqWrite, err := cldf_ops.ExecuteOperation(b, applyFeeQuoterDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
				Args: []fee_quoter.DestChainConfigArgs{
					{DestChainSelector: input.TargetChainSelector, DestChainConfig: newFQConfig},
				},
			})
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf("failed to apply FeeQuoter dest chain config update for src %d: %w", lane.ChainSelector, err)
			}
			writes = append(writes, fqWrite.Output)

			// --- OffRamp: GasForCallExactCheck (row 3) ---
			// Read-only sanity check: this field is immutable with no setter, so no write is ever
			// produced here, but an unexpected on-chain value is still worth surfacing.
			if lane.OffRampAddress != (common.Address{}) {
				offRampCur, err := cldf_ops.ExecuteOperation(b, offramp.GetStaticConfig, chain, contract.FunctionInput[struct{}]{
					ChainSelector: lane.ChainSelector,
					Address:       lane.OffRampAddress,
				})
				if err != nil {
					return UpdateGasConfigOutput{}, fmt.Errorf("failed to read OffRamp static config for src %d: %w", lane.ChainSelector, err)
				}

				if offRampCur.Output.GasForCallExactCheck != OffRampExpectedGasForCallExactCheck {
					line := fmt.Sprintf(
						"chain %d: OffRamp.StaticConfig.GasForCallExactCheck is immutable with no setter, but current "+
							"value %d does not match expected Prague baseline %d - unexpected deployment, investigate manually",
						lane.ChainSelector, offRampCur.Output.GasForCallExactCheck, OffRampExpectedGasForCallExactCheck,
					)
					b.Logger.Warn(line)
					output.Report.AddLine(line)
				}
			}

			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return UpdateGasConfigOutput{}, fmt.Errorf("failed to build batch operation for src %d: %w", lane.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}

		return output, nil
	},
)
