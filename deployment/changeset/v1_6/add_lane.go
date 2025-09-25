package v1_6

import (
	"github.com/smartcontractkit/chainlink-ccip/deployment/config/v1_6"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var _ cldf.ChangeSetV2[v1_6.AddLanesConfig] = AddLanes{}

type AddLanes struct{}

func (cs AddLanes) VerifyPreconditions(env cldf.Environment, cfg v1_6.AddLanesConfig) error {
	// TODO: implement this
	return nil
}

func (cs AddLanes) Apply(env cldf.Environment, cfg v1_6.AddLanesConfig) (cldf.ChangesetOutput, error) {
	// inputsByChain := make(map[uint64]v1_6.UpdateLanesInput)
	// for each lane in cfg.Lanes build UpdateLanesInput
	// for each inputsByChain call the respective sequence to addLanesSide
	return cldf.ChangesetOutput{}, nil
}
