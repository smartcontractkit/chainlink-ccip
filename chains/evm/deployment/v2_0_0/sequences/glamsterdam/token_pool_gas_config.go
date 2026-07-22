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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
)

// applyTokenPoolTokenTransferFeeConfigUpdates mirrors
// token_pool.ApplyTokenTransferFeeConfigUpdates, but always routes the write through MCMS. Both
// LombardTokenPool and (Siloed)USDCTokenPool inherit this function unmodified from the shared
// TokenPool base contract, so the base TokenPoolABI binds to either address correctly.
var applyTokenPoolTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs, *token_pool.TokenPoolContract]{
	Name:            "glamsterdam:token-pool:apply-token-transfer-fee-config-updates",
	Version:         semver.MustParse("2.0.0"),
	Description:     "Calls applyTokenTransferFeeConfigUpdates on TokenPool, always producing an MCMS proposal",
	ContractType:    token_pool.ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPoolContract,
	IsAllowedCaller: contract.NoCallersAllowed[*token_pool.TokenPoolContract, token_pool.ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(token_pool.ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(c *token_pool.TokenPoolContract, opts *bind.TransactOpts, args token_pool.ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return c.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.DisableTokenTransferFeeConfigs)
	},
})

// TokenPoolLane is one token pool (Lombard or USDC) on a source chain with a confirmed lane
// pointed at the Glamsterdam target chain.
type TokenPoolLane struct {
	ChainSelector uint64
	PoolAddress   common.Address
}

// UpdateTokenPoolGasConfigInput is the input to UpdateTokenPoolGasConfig.
type UpdateTokenPoolGasConfigInput struct {
	// TargetChainSelector is the chain selector moving to Glamsterdam.
	TargetChainSelector uint64
	// LombardPools are LombardTokenPool addresses per chain (table row 9).
	LombardPools []TokenPoolLane
	// USDCPools are (Siloed)USDCTokenPool addresses per chain (table row 10).
	USDCPools []TokenPoolLane
}

// UpdateTokenPoolGasConfigOutput is the output of UpdateTokenPoolGasConfig.
type UpdateTokenPoolGasConfigOutput struct {
	// BatchOps contains one MCMS batch operation per pool (a chain with both a Lombard and a
	// USDC pool produces two batch operations, since they are different contracts).
	BatchOps []mcms_types.BatchOperation
	// Report is a human-readable summary of every field resolved on every pool, for inclusion in
	// the changeset's MCMS proposal description.
	Report *glamsterdamutils.Report
}

// UpdateTokenPoolGasConfig reads the current DestGasOverhead for each Lombard/USDC token pool
// pointed at the target chain, resolves it against its expected Prague baseline (applying the
// literal Glamsterdam value on a match, or the field's fallback rule on a mismatch), and packages
// the write into an MCMS batch operation per pool. Every write is routed through MCMS regardless
// of who the deployer key is.
var UpdateTokenPoolGasConfig = cldf_ops.NewSequence(
	"UpdateTokenPoolGasConfigV2",
	semver.MustParse("2.0.0"),
	"Updates v2.0 Lombard/USDC token pool gas config for lanes pointed at the Glamsterdam target chain",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input UpdateTokenPoolGasConfigInput) (UpdateTokenPoolGasConfigOutput, error) {
		output := UpdateTokenPoolGasConfigOutput{Report: glamsterdamutils.NewReport()}

		processLane := func(lane TokenPoolLane, spec glamsterdamutils.FieldSpec[uint32]) error {
			chain, ok := chains.EVMChains()[lane.ChainSelector]
			if !ok {
				return fmt.Errorf("chain with selector %d not found", lane.ChainSelector)
			}

			cur, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.PoolAddress,
				Args:          token_pool.GetTokenTransferFeeConfigArgs{DestChainSelector: input.TargetChainSelector},
			})
			if err != nil {
				return fmt.Errorf(
					"failed to read TokenPool(%s) transfer fee config for src %d, dst %d: %w",
					lane.PoolAddress, lane.ChainSelector, input.TargetChainSelector, err,
				)
			}

			result := glamsterdamutils.Resolve(spec, cur.Output.DestGasOverhead)
			glamsterdamutils.AddField(output.Report, lane.ChainSelector, result)

			newConfig := cur.Output
			newConfig.DestGasOverhead = result.AppliedValue

			write, err := cldf_ops.ExecuteOperation(b, applyTokenPoolTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
				ChainSelector: lane.ChainSelector,
				Address:       lane.PoolAddress,
				Args: token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
					TokenTransferFeeConfigArgs: []token_pool.TokenTransferFeeConfigArgs{
						{DestChainSelector: input.TargetChainSelector, TokenTransferFeeConfig: newConfig},
					},
				},
			})
			if err != nil {
				return fmt.Errorf("failed to apply TokenPool(%s) transfer fee config update for src %d: %w", lane.PoolAddress, lane.ChainSelector, err)
			}

			batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{write.Output})
			if err != nil {
				return fmt.Errorf("failed to build batch operation for TokenPool(%s) on src %d: %w", lane.PoolAddress, lane.ChainSelector, err)
			}
			output.BatchOps = append(output.BatchOps, batchOp)

			return nil
		}

		for _, lane := range input.LombardPools {
			if err := processLane(lane, LombardTokenPoolDestGasOverhead); err != nil {
				return UpdateTokenPoolGasConfigOutput{}, err
			}
		}
		for _, lane := range input.USDCPools {
			if err := processLane(lane, USDCTokenPoolDestGasOverhead); err != nil {
				return UpdateTokenPoolGasConfigOutput{}, err
			}
		}

		return output, nil
	},
)
