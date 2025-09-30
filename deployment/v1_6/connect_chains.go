package v1_6

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var _ cldf.ChangeSetV2[ConnectChainsConfig] = ConnectChains{}

type ConnectChains struct{}

func (cs ConnectChains) VerifyPreconditions(env cldf.Environment, cfg ConnectChainsConfig) error {
	// TODO: implement this
	return nil
}

func (cs ConnectChains) Apply(env cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
	for _, lane := range cfg.Lanes {
		// todo: fill this in for dest
		src := lane.Source
		family, err := chain_selectors.GetSelectorFamily(src.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if _, exists := registeredChainAdapters[family]; !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", family)
		}
		// coalesce all outputs
		registeredChainAdapters[family].ConfigureLaneLeg(env, UpdateLanesInput{
			//todo fill this in
		})
	}
	return cldf.ChangesetOutput{}, nil
}


