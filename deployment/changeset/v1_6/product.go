package v1_6

import (
	"fmt"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type Product interface { // introduce generics as needed
	ConnectChains(env cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error)
}

type ChainAdapter interface {
	ConfigureLaneLeg(env cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error)
}

var registeredChainAdapters = make(map[string]ChainAdapter)

// RegisterChainAdapter allows chains to register their changeset logic.
func RegisterChainAdapter(chainFamily string, adapter ChainAdapter) {
	if _, exists := registeredChainAdapters[chainFamily]; exists {
		panic(fmt.Sprintf("ChainAdapter '%s' already registered", chainFamily))
	}
	registeredChainAdapters[chainFamily] = adapter
}
