package v1_6

import (
	"fmt"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type ChainAdapter interface {
	ConfigureLaneLeg(env cldf.Environment, cfg UpdateLanesInput) (cldf.ChangesetOutput, error)
}

var registeredChainAdapters = make(map[string]ChainAdapter)

// RegisterChainAdapter allows chains to register their changeset logic.
func RegisterChainAdapter(chainFamily string, adapter ChainAdapter) {
	if _, exists := registeredChainAdapters[chainFamily]; exists {
		panic(fmt.Sprintf("ChainAdapter '%s' already registered", chainFamily))
	}
	registeredChainAdapters[chainFamily] = adapter
}
