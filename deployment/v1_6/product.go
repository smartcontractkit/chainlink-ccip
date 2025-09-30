package v1_6

import (
	"fmt"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type ChainAdapter interface {
	// high level API
	ConfigureLaneLeg(env cldf.Environment, cfg UpdateLanesInput) (cldf.ChangesetOutput, error)

	// helpers to expose lower level functionality if needed
	// needed for populating values in chain specific configs
	GetOnRampAddress(env cldf.Environment) ([]byte, error)
}

var registeredChainAdapters = make(map[string]ChainAdapter)

// RegisterChainAdapter allows chains to register their changeset logic.
func RegisterChainAdapter(chainFamily string, adapter ChainAdapter) {
	if _, exists := registeredChainAdapters[chainFamily]; exists {
		panic(fmt.Sprintf("ChainAdapter '%s' already registered", chainFamily))
	}
	registeredChainAdapters[chainFamily] = adapter
}
