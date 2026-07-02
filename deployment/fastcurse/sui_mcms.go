package fastcurse

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// mcmsInputForUncurseProposal returns MCMS input for uncurse proposals.
// Sui chains always use slow MCMS for uncurse even when the global qualifier is RMNMCMS.
func mcmsInputForUncurseProposal(cfg mcms_utils.Input, batchOps []mcms_types.BatchOperation) mcms_utils.Input {
	input := cfg
	chainQualifiers := cloneChainQualifiers(cfg.ChainQualifiers)

	for _, op := range batchOps {
		selector := uint64(op.ChainSelector)
		family, err := chain_selectors.GetSelectorFamily(selector)
		if err != nil || family != chain_selectors.FamilySui {
			continue
		}
		chainQualifiers[selector] = ""
	}

	if len(chainQualifiers) > 0 {
		input.ChainQualifiers = chainQualifiers
	}
	return input
}

func cloneChainQualifiers(in map[uint64]string) map[uint64]string {
	if len(in) == 0 {
		return make(map[uint64]string)
	}
	out := make(map[uint64]string, len(in))
	for selector, qualifier := range in {
		out[selector] = qualifier
	}
	return out
}
