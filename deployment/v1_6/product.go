package v1_6

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
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
	GetMCMSMetadata(e cldf.Environment, chainSelector uint64, action types.TimelockAction) (types.ChainMetadata, error)
}

var registeredChainAdapters = make(map[string]ChainAdapter)

func newAdapterID(chainFamily string, version *semver.Version) string {
	return fmt.Sprintf("%s-%s", chainFamily, version.String())
}

// RegisterChainAdapter allows chains to register their changeset logic.
func RegisterChainAdapter(chainFamily string, version *semver.Version, adapter ChainAdapter) {
	id := newAdapterID(chainFamily, version)
	if _, exists := registeredChainAdapters[id]; exists {
		panic(fmt.Sprintf("ChainAdapter '%s' already registered", chainFamily))
	}
	registeredChainAdapters[id] = adapter
}
