package sequences

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var DisableRemoteChainSequence = operations.NewSequence(
	"DisableRemoteChain",
	semver.MustParse("1.6.0"),
	"Disables both sending to and receiving from a remote chain on an EVM chain via router updates",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input lanes.DisableRemoteChainInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		b.Logger.Infof("EVM Disabling remote chain %d on chain %d", input.RemoteChainSelector, input.LocalChainSelector)

		result, err := sequences.RunAndMergeSequence(b, chains, RouterApplyRampUpdatesSequence, RouterApplyRampUpdatesSequenceInput{
			Address:       common.BytesToAddress(input.Router),
			ChainSelector: input.LocalChainSelector,
			UpdatesByChain: router.ApplyRampsUpdatesArgs{
				OnRampUpdates: []router.OnRamp{
					{
						DestChainSelector: input.RemoteChainSelector,
						OnRamp:            common.HexToAddress("0x0"),
					},
				},
				OffRampRemoves: []router.OffRamp{
					{
						SourceChainSelector: input.RemoteChainSelector,
						OffRamp:             common.BytesToAddress(input.OffRamp),
					},
				},
			},
		}, result)
		if err != nil {
			return result, err
		}
		b.Logger.Infof("Remote chain %d disabled on EVM chain %d", input.RemoteChainSelector, input.LocalChainSelector)

		return result, nil
	},
)

func (a *EVMAdapter) DisableRemoteChain() *operations.Sequence[lanes.DisableRemoteChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DisableRemoteChainSequence
}
