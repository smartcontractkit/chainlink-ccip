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

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/burn_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/burn_to_address_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/burn_with_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/token_pool"
)

// tokenPoolTypes lists contract types that use the TokenPool ABI for withdrawFeeTokens.
// All subtypes inherit from the base TokenPool Solidity contract and share the same
// withdrawFeeTokens(address[],address) signature.
var tokenPoolTypes = map[datastore.ContractType]bool{
	datastore.ContractType(token_pool.ContractType):                                    true,
	datastore.ContractType(burn_mint_token_pool.ContractType):                          true,
	datastore.ContractType(burn_from_mint_token_pool.ContractType):                     true,
	datastore.ContractType(burn_with_from_mint_token_pool.ContractType):                true,
	datastore.ContractType(burn_to_address_mint_token_pool.ContractType):               true,
	datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType):   true,
	datastore.ContractType(lock_release_token_pool.ContractType):                       true,
	datastore.ContractType(siloed_lock_release_token_pool.ContractType):                true,
	datastore.ContractType(lombard_token_pool.ContractType):                            true,
	datastore.ContractType(cctp_through_ccv_token_pool.ContractType):                   true,
	datastore.ContractType(siloed_usdc_token_pool.ContractType):                        true,
}

// feeTokenHandlerTypes is the union of non-pool handlers (OnRamp, CommitteeVerifier)
// and all TokenPool variants. Built once at init to avoid duplication with tokenPoolTypes.
var feeTokenHandlerTypes map[datastore.ContractType]bool

func init() {
	feeTokenHandlerTypes = map[datastore.ContractType]bool{
		datastore.ContractType(onramp.ContractType):             true,
		datastore.ContractType(committee_verifier.ContractType): true,
	}
	for ct := range tokenPoolTypes {
		feeTokenHandlerTypes[ct] = true
	}
}

// IsFeeTokenHandler returns true if the given contract type supports withdrawFeeTokens.
// This is used by both the sequence and the changeset to validate user-supplied refs.
func IsFeeTokenHandler(contractType datastore.ContractType) bool {
	return feeTokenHandlerTypes[contractType]
}

// IsTokenPoolType returns true if the given contract type is any variant of TokenPool.
func IsTokenPoolType(contractType datastore.ContractType) bool {
	return tokenPoolTypes[contractType]
}

// WithdrawFeeTokensInput is the resolved input for the WithdrawFeeTokens sequence.
// All AddressRefs should already have their Address field populated (done by the
// changeset's ResolveInput via datastore lookup).
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

		writes := make([]contract_utils.WriteOutput, 0, len(input.ContractRefs))

		for _, ref := range input.ContractRefs {
			if !IsFeeTokenHandler(ref.Type) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"contract type %q is not a supported FeeTokenHandler",
					ref.Type,
				)
			}
			if !common.IsHexAddress(ref.Address) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"invalid contract address %q for type %s", ref.Address, ref.Type,
				)
			}
			addr := common.HexToAddress(ref.Address)

			switch {
			case ref.Type == datastore.ContractType(onramp.ContractType):
				report, err := cldf_ops.ExecuteOperation(b, onramp.WithdrawFeeTokens, chain, contract_utils.FunctionInput[[]common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       addr,
					Args:          input.FeeTokens,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw fee tokens from OnRamp %s: %w", ref.Address, err)
				}
				writes = append(writes, report.Output)

			case ref.Type == datastore.ContractType(committee_verifier.ContractType):
				report, err := cldf_ops.ExecuteOperation(b, committee_verifier.WithdrawFeeTokens, chain, contract_utils.FunctionInput[[]common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       addr,
					Args:          input.FeeTokens,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw fee tokens from CommitteeVerifier %s: %w", ref.Address, err)
				}
				writes = append(writes, report.Output)

			case IsTokenPoolType(ref.Type):
				if input.Recipient == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("recipient is required when withdrawing fee tokens from %s %s", ref.Type, ref.Address)
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
					return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw fee tokens from %s %s: %w", ref.Type, ref.Address, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
