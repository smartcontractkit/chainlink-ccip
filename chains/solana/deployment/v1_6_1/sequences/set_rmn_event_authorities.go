package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type RMNRemoteSetEventAuthoritiesSequenceInput struct {
	DataStore datastore.DataStore
	Selector  uint64
}

func (a *SolanaAdapter) SetRMNRemoteEventAuthorities() *cldf_ops.Sequence[RMNRemoteSetEventAuthoritiesSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return SetRMNRemoteEventAuthoritiesConfig
}

var SetRMNRemoteEventAuthoritiesConfig = cldf_ops.NewSequence(
	"rmn-remote:set-event-authorities",
	semver.MustParse("1.6.1"),
	"Sets event authorities on the RMNRemote 1.6.1 contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input RMNRemoteSetEventAuthoritiesSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.SolanaChains()[input.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not found", input.Selector)
		}

		rmnAddrBytes, err := (&SolanaAdapter{}).GetRMNRemoteAddress(input.DataStore, input.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get RMNRemote address: %w", err)
		}
		rmnPubkey := solana.PublicKeyFromBytes(rmnAddrBytes)

		routerAddrBytes, err := (&SolanaAdapter{}).GetRouterAddress(input.DataStore, input.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get Router address: %w", err)
		}
		routerPubkey := solana.PublicKeyFromBytes(routerAddrBytes)

		rmnConfigPDA, _, err := state.FindRMNRemoteConfigPDA(rmnPubkey)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find RMNRemote config PDA: %w", err)
		}
		routerSignerPDA, _, err := state.FindFeeBillingSignerPDA(routerPubkey)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find Router Signer PDA: %w", err)
		}

		out, err := operations.ExecuteOperation(b,
			rmnops.SetEventAuthorities,
			chain,
			rmnops.EventAuthoritiesInput{
				EventAuthorities:   []solana.PublicKey{routerSignerPDA},
				RMNRemote:          rmnPubkey,
				RMNRemoteConfigPDA: rmnConfigPDA,
			},
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set event authorities: %w", err)
		}
		return out.Output, nil
	},
)
