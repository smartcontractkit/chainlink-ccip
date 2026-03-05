package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
)

// feeTokenHandlerTypes enumerates the contract types that expose a withdrawFeeTokens
// method on-chain. Used by IsFeeTokenHandler to gate which contract refs are accepted.
var feeTokenHandlerTypes = map[datastore.ContractType]bool{
	datastore.ContractType(onramp.ContractType):             true,
	datastore.ContractType(committee_verifier.ContractType): true,
	datastore.ContractType(token_pool.ContractType):         true,
}

// IsFeeTokenHandler returns true if the given contract type supports WithdrawFeeTokens.
func IsFeeTokenHandler(contractType datastore.ContractType) bool {
	return feeTokenHandlerTypes[contractType]
}

// WithdrawFeeTokensInput is the input for the WithdrawFeeTokens sequence.
type WithdrawFeeTokensInput struct {
	ChainSelector uint64
	ContractRefs  []datastore.AddressRef
	FeeTokens     []common.Address
	// Recipient receives withdrawn tokens for TokenPool contracts.
	// OnRamp and CommitteeVerifier ignore this; they send to their configured feeAggregator.
	Recipient common.Address
}

// WithdrawFeeTokens is the core sequence that iterates over the supplied contract refs,
// dispatches the appropriate withdrawFeeTokens operation for each contract type, and
// collects all resulting write outputs into a single MCMS BatchOperation.
//
// TokenPool has a different Solidity signature (requires a recipient address), so it
// is handled as a separate case from OnRamp and CommitteeVerifier.
var WithdrawFeeTokens = cldf_ops.NewSequence(
	"withdraw-fee-tokens",
	semver.MustParse("1.7.0"),
	"Withdraws fee tokens from one or more fee-handling contracts on an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input WithdrawFeeTokensInput) (sequences.OnChainOutput, error) {
		if len(input.ContractRefs) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("at least one contract ref is required")
		}
		if len(input.FeeTokens) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("at least one fee token address is required")
		}

		// Accumulate a WriteOutput per contract so they can be batched into one MCMS proposal.
		writes := make([]contract_utils.WriteOutput, 0, len(input.ContractRefs))

		for _, ref := range input.ContractRefs {
			if !IsFeeTokenHandler(ref.Type) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"contract type %q is not a supported FeeTokenHandler (supported: OnRamp, CommitteeVerifier, TokenPool)",
					ref.Type,
				)
			}
			addr := common.HexToAddress(ref.Address)

			// Dispatch to the correct operation based on contract type.
			// OnRamp and CommitteeVerifier share the same signature (just feeTokens),
			// while TokenPool additionally requires a recipient address.
			switch ref.Type {
			case datastore.ContractType(onramp.ContractType):
				report, err := cldf_ops.ExecuteOperation(b, onramp.WithdrawFeeTokens, chain, contract_utils.FunctionInput[onramp.WithdrawFeeTokensArgs]{
					ChainSelector: input.ChainSelector,
					Address:       addr,
					Args: onramp.WithdrawFeeTokensArgs{
						FeeTokens: input.FeeTokens,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw fee tokens from OnRamp %s: %w", ref.Address, err)
				}
				writes = append(writes, report.Output)

			case datastore.ContractType(committee_verifier.ContractType):
				report, err := cldf_ops.ExecuteOperation(b, committee_verifier.WithdrawFeeTokens, chain, contract_utils.FunctionInput[committee_verifier.WithdrawFeeTokensArgs]{
					ChainSelector: input.ChainSelector,
					Address:       addr,
					Args: committee_verifier.WithdrawFeeTokensArgs{
						FeeTokens: input.FeeTokens,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw fee tokens from CommitteeVerifier %s: %w", ref.Address, err)
				}
				writes = append(writes, report.Output)

			case datastore.ContractType(token_pool.ContractType):
				// TokenPool's withdrawFeeTokens(address, address[]) requires a recipient.
				if input.Recipient == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("recipient is required when withdrawing fee tokens from TokenPool %s", ref.Address)
				}
				report, err := cldf_ops.ExecuteOperation(b, token_pool.WithdrawFeeTokens, chain, contract_utils.FunctionInput[token_pool.WithdrawFeeTokensArgs]{
					ChainSelector: input.ChainSelector,
					Address:       addr,
					Args: token_pool.WithdrawFeeTokensArgs{
						FeeTokens: input.FeeTokens,
						Recipient: input.Recipient,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw fee tokens from TokenPool %s: %w", ref.Address, err)
				}
				writes = append(writes, report.Output)
			}
		}

		// Bundle all write outputs into a single MCMS batch operation so they can be
		// proposed and executed atomically.
		batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
