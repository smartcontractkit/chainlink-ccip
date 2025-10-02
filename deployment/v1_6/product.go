package v1_6

import (
	"fmt"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/mcms/types"
)

type ChainAdapter interface {
	// high level API
	ConfigureLaneLegAsSource(e cldf.Environment, cfg UpdateLanesInput) (cldf.ChangesetOutput, error)
	ConfigureLaneLegAsDest(e cldf.Environment, cfg UpdateLanesInput) (cldf.ChangesetOutput, error)
	ConfigureLaneAsSourceAndDest(e cldf.Environment, cfg UpdateLanesInput) (cldf.ChangesetOutput, error)

	// helpers to expose lower level functionality if needed
	// needed for populating values in chain specific configs
	GetOnRampAddress(e cldf.Environment, chainSelector uint64) ([]byte, error)
	GetOffRampAddress(e cldf.Environment, chainSelector uint64) ([]byte, error)
	GetTimelockAddress(e cldf.Environment, chainSelector uint64) (string, error)
	GetMCMSMetadata(e cldf.Environment, chainSelector uint64) (types.ChainMetadata, error)
}

var registeredChainAdapters = make(map[string]ChainAdapter)

// RegisterChainAdapter allows chains to register their changeset logic.
func RegisterChainAdapter(chainFamily string, adapter ChainAdapter) {
	if _, exists := registeredChainAdapters[chainFamily]; exists {
		panic(fmt.Sprintf("ChainAdapter '%s' already registered", chainFamily))
	}
	registeredChainAdapters[chainFamily] = adapter
}
