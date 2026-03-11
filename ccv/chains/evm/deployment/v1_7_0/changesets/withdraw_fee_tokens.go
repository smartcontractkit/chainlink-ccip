package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
)

// WithdrawFeeTokensCfg is the YAML/pipeline input for the WithdrawFeeTokens changeset.
type WithdrawFeeTokensCfg struct {
	ChainSel uint64
	// ContractRefs identifies the contracts to withdraw from.
	ContractRefs []datastore.AddressRef
	// FeeTokens is the list of ERC-20 token addresses to withdraw from each contract.
	FeeTokens []common.Address
	// Recipient is required when any ref is a TokenPool. Ignored for OnRamp/CommitteeVerifier.
	Recipient common.Address
}

// ChainSelector implements the single-chain config interface required by
// ResolveEVMChainDep, which looks up the evm.Chain from the environment.
func (c WithdrawFeeTokensCfg) ChainSelector() uint64 {
	return c.ChainSel
}

// WithdrawFeeTokens wraps the withdraw-fee-tokens sequence into a changeset.
// It resolves each user-supplied AddressRef against the datastore to obtain on-chain
// addresses, validates that every ref is a known FeeTokenHandler, and delegates
// execution to the sequence. The result is an MCMS proposal containing all withdrawals.
var WithdrawFeeTokens = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.WithdrawFeeTokensInput,
	evm.Chain,
	WithdrawFeeTokensCfg,
]{
	Sequence: sequences.WithdrawFeeTokens,

	// ResolveInput converts the user-facing config into the sequence's input by:
	// 1. Validating each ref is a supported FeeTokenHandler type.
	// 2. Looking up the deployed address in the environment's datastore.
	// 3. Building fully-resolved AddressRefs with the on-chain address populated.
	ResolveInput: func(e cldf_deployment.Environment, cfg WithdrawFeeTokensCfg) (sequences.WithdrawFeeTokensInput, error) {
		resolvedRefs := make([]datastore.AddressRef, 0, len(cfg.ContractRefs))
		for _, ref := range cfg.ContractRefs {
			// Reject unknown contract types early, before hitting the datastore.
			if !sequences.IsFeeTokenHandler(ref.Type) {
				return sequences.WithdrawFeeTokensInput{}, fmt.Errorf(
					"contract type %q is not a supported FeeTokenHandler", ref.Type,
				)
			}
			// Look up the contract's deployed address from the datastore.
			resolvedAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return sequences.WithdrawFeeTokensInput{}, fmt.Errorf(
					"failed to resolve contract ref (type=%s, version=%v, qualifier=%s) on chain %d: %w",
					ref.Type, ref.Version, ref.Qualifier, cfg.ChainSel, err,
				)
			}
			resolvedRefs = append(resolvedRefs, datastore.AddressRef{
				Address:       resolvedAddr.Hex(),
				ChainSelector: cfg.ChainSel,
				Type:          ref.Type,
				Version:       ref.Version,
				Qualifier:     ref.Qualifier,
			})
		}

		return sequences.WithdrawFeeTokensInput{
			ChainSelector: cfg.ChainSel,
			ContractRefs:  resolvedRefs,
			FeeTokens:     cfg.FeeTokens,
			Recipient:     cfg.Recipient,
		}, nil
	},

	// ResolveDep looks up the evm.Chain object from the environment using ChainSel.
	ResolveDep: evm_sequences.ResolveEVMChainDep[WithdrawFeeTokensCfg],
})
