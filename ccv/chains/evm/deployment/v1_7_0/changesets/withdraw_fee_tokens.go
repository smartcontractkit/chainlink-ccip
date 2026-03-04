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

// WithdrawFeeTokensCfg is the configuration for the WithdrawFeeTokens changeset.
type WithdrawFeeTokensCfg struct {
	ChainSel uint64
	// ContractRefs identifies the contracts to withdraw fee tokens from.
	// Each ref must specify Type (and optionally Version/Qualifier) so the contract
	// can be looked up in the datastore. The Type must be a supported FeeTokenHandler
	// (OnRamp, CommitteeVerifier, or TokenPool).
	ContractRefs []datastore.AddressRef
	// FeeTokens is the list of fee token addresses to withdraw.
	FeeTokens []common.Address
	// Recipient is the address that receives withdrawn fee tokens for TokenPool contracts.
	// Ignored for OnRamp and CommitteeVerifier (they send to their configured feeAggregator).
	Recipient common.Address
}

func (c WithdrawFeeTokensCfg) ChainSelector() uint64 {
	return c.ChainSel
}

// WithdrawFeeTokens creates a changeset that withdraws fee tokens from one or more
// fee-handling contracts (OnRamp, CommitteeVerifier, TokenPool) on an EVM chain.
var WithdrawFeeTokens = changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
	sequences.WithdrawFeeTokensInput,
	evm.Chain,
	WithdrawFeeTokensCfg,
]{
	Sequence: sequences.WithdrawFeeTokens,
	ResolveInput: func(e cldf_deployment.Environment, cfg WithdrawFeeTokensCfg) (sequences.WithdrawFeeTokensInput, error) {
		resolvedRefs := make([]datastore.AddressRef, 0, len(cfg.ContractRefs))
		for _, ref := range cfg.ContractRefs {
			if !sequences.IsFeeTokenHandler(ref.Type) {
				return sequences.WithdrawFeeTokensInput{}, fmt.Errorf(
					"contract type %q is not a supported FeeTokenHandler", ref.Type,
				)
			}
			resolvedAddr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, cfg.ChainSel, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return sequences.WithdrawFeeTokensInput{}, fmt.Errorf(
					"failed to resolve contract ref (type=%s, version=%s, qualifier=%s) on chain %d: %w",
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
	ResolveDep: evm_sequences.ResolveEVMChainDep[WithdrawFeeTokensCfg],
})
