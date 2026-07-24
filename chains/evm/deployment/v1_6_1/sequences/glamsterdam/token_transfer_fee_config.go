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
)

// applyFeeQuoterTokenTransferFeeConfigUpdates mirrors fee_quoter.ApplyTokenTransferFeeConfigUpdates,
// but always routes the write through MCMS regardless of who the deployer key is, per this
// feature's "always MCMS" decision.
var applyFeeQuoterTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs, *fee_quoter.FeeQuoterContract]{
	Name:            "glamsterdam:fee-quoter:apply-token-transfer-fee-config-updates",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Calls applyTokenTransferFeeConfigUpdates on FeeQuoter, always producing an MCMS proposal",
	ContractType:    fee_quoter.ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoterContract,
	IsAllowedCaller: contract.NoCallersAllowed[*fee_quoter.FeeQuoterContract, fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(c *fee_quoter.FeeQuoterContract, opts *bind.TransactOpts, args fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return c.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

// TokenTransferFeeConfigLane is one source chain with a confirmed lane pointed at the Glamsterdam
// target chain, along with the candidate tokens to check for an existing
// FeeQuoter.TokenTransferFeeConfig override on that lane.
type TokenTransferFeeConfigLane struct {
	ChainSelector    uint64
	FeeQuoterAddress common.Address
	// CandidateTokens are token addresses to check for an existing TokenTransferFeeConfig
	// override on this lane (e.g. every token this chain's TokenAdminRegistry knows about). Only
	// tokens with an enabled override are resolved and updated; tokens with no override
	// configured for this lane are skipped silently, since that's the common case.
	CandidateTokens []common.Address
}

// UpdateTokenTransferFeeConfigInput is the input to UpdateTokenTransferFeeConfig.
type UpdateTokenTransferFeeConfigInput struct {
	// TargetChainSelector is the chain selector moving to Glamsterdam.
	TargetChainSelector uint64
	// Lanes are the source chains with a confirmed lane pointed at the target chain.
	Lanes []TokenTransferFeeConfigLane
}

// UpdateTokenTransferFeeConfigOutput is the output of UpdateTokenTransferFeeConfig.
type UpdateTokenTransferFeeConfigOutput struct {
	// BatchOps contains at most one MCMS batch operation per chain in Lanes (chains with no
	// enabled token transfer fee override for any candidate token produce no batch operation).
	BatchOps []mcms_types.BatchOperation
	// Report is a human-readable summary of every field resolved on every lane, for inclusion in
	// the changeset's MCMS proposal description.
	Report *glamsterdamutils.Report
}

// UpdateTokenTransferFeeConfig reads the current FeeQuoter.TokenTransferFeeConfig for every
// candidate token on every confirmed lane. Tokens with no enabled override on that lane are
// skipped. For tokens with an enabled override, DestGasOverhead is resolved against its expected
// Prague baseline (applying the literal Glamsterdam value on a match, or the fallback rule on a
// mismatch), and all resolved tokens for one chain are grouped into a single MCMS batch
// operation. Every write is routed through MCMS regardless of who the deployer key is.
var UpdateTokenTransferFeeConfig = cldf_ops.NewSequence(
	"UpdateTokenTransferFeeConfigV1_6",
	semver.MustParse("1.6.0"),
	"Updates v1.6 FeeQuoter token transfer fee config for lanes pointed at the Glamsterdam target chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input UpdateTokenTransferFeeConfigInput) (UpdateTokenTransferFeeConfigOutput, error) {
		output := UpdateTokenTransferFeeConfigOutput{Report: glamsterdamutils.NewReport()}

		for _, lane := range input.Lanes {
			chain, ok := chains.EVMChains()[lane.ChainSelector]
			if !ok {
				return UpdateTokenTransferFeeConfigOutput{}, fmt.Errorf("chain with selector %d not found", lane.ChainSelector)
			}

			var singleTokenArgs []fee_quoter.TokenTransferFeeConfigSingleTokenArgs

			for _, token := range lane.CandidateTokens {
				cur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetTokenTransferFeeConfig, chain, contract.FunctionInput[fee_quoter.GetTokenTransferFeeConfigArgs]{
					ChainSelector: lane.ChainSelector,
					Address:       lane.FeeQuoterAddress,
					Args: fee_quoter.GetTokenTransferFeeConfigArgs{
						DestChainSelector: input.TargetChainSelector,
						Token:             token,
					},
				})
				if err != nil {
					return UpdateTokenTransferFeeConfigOutput{}, fmt.Errorf(
						"failed to read FeeQuoter token transfer fee config for src %d, dst %d, token %s: %w",
						lane.ChainSelector, input.TargetChainSelector, token, err,
					)
				}
				if !cur.Output.IsEnabled {
					// No override configured for this token on this lane — nothing to update.
					continue
				}

				result := glamsterdamutils.Resolve(USDCTokenPoolDestGasOverhead, cur.Output.DestGasOverhead)
				glamsterdamutils.AddField(output.Report, lane.ChainSelector, result)

				newConfig := cur.Output
				newConfig.DestGasOverhead = result.AppliedValue
				singleTokenArgs = append(singleTokenArgs, fee_quoter.TokenTransferFeeConfigSingleTokenArgs{
					Token:                  token,
					TokenTransferFeeConfig: newConfig,
				})
			}

			if len(singleTokenArgs) == 0 {
				continue
			}

			write, err := cldf_ops.ExecuteOperation(b, applyFeeQuoterTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.FeeQuoterAddress,
				Args: fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs{
					TokenTransferFeeConfigArgs: []fee_quoter.TokenTransferFeeConfigArgs{
						{DestChainSelector: input.TargetChainSelector, TokenTransferFeeConfigs: singleTokenArgs},
					},
				},
			})
			if err != nil {
				return UpdateTokenTransferFeeConfigOutput{}, fmt.Errorf("failed to apply FeeQuoter token transfer fee config update for src %d: %w", lane.ChainSelector, err)
			}

			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{write.Output})
			if err != nil {
				return UpdateTokenTransferFeeConfigOutput{}, fmt.Errorf("failed to build batch operation for src %d: %w", lane.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)
		}

		return output, nil
	},
)
