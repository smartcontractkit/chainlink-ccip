package v1_6

import (
	changeset_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
)

var ConnectChains = changeset_utils.NewFromOnChainSequence(changeset_utils.NewFromOnChainSequenceParams[
	UpdateLanesInput,
	cldf_chain.BlockChains,
	ConnectChainsConfig,
]{
	// inputsByChain := make(map[uint64]v1_6.UpdateLanesInput)
	// for each lane in cfg.Lanes build UpdateLanesInput
	// for each inputsByChain call the respective sequence to addLanesSide
})
