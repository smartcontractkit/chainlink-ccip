package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type RMNRemoteSetCurserSequenceInput struct {
	DataStore datastore.DataStore
	Curser    solana.PublicKey
	Selector  uint64
}

var SetRMNRemoteCurser = cldf_ops.NewSequence(
	"rmn-remote:set-curser",
	semver.MustParse("1.6.3"),
	"Sets curser on the RMNRemote 1.6.3 contract",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input RMNRemoteSetCurserSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.SolanaChains()[input.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not found", input.Selector)
		}

		rmnAddrBytes, err := (&SolanaAdapter{}).GetRMNRemoteAddress(input.DataStore, input.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get RMNRemote address: %w", err)
		}
		rmnPubkey := solana.PublicKeyFromBytes(rmnAddrBytes)

		rmnConfigPDA, _, err := state.FindRMNRemoteConfigPDA(rmnPubkey)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find RMNRemote config PDA: %w", err)
		}

		out, err := cldf_ops.ExecuteOperation(b,
			rmnops.SetCurser,
			chain,
			rmnops.SetCurserInput{
				Curser:             input.Curser,
				RMNRemote:          rmnPubkey,
				RMNRemoteConfigPDA: rmnConfigPDA,
			},
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set curser: %w", err)
		}
		return out.Output, nil
	},
)
