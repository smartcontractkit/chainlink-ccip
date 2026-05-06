package authorizedcallers

import (
	"github.com/Masterminds/semver/v3"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// Caller is the byte-encoded identity of a caller address. Each chain family
// encodes its native address type into bytes (EVM: 20 bytes, Solana: 32 bytes).
type Caller = []byte

// CallerUpdate describes the set of callers to add and remove in a single
// applyAuthorizedCallerUpdates call.
type CallerUpdate struct {
	AddedCallers   []Caller `json:"addedCallers"   yaml:"addedCallers"`
	RemovedCallers []Caller `json:"removedCallers"  yaml:"removedCallers"`
}

// ApplyInput is the per-chain, per-contract input for ConfigureAuthorizedCallersChangeset.
// ContractType and Version together identify which AuthorizedCallers-inheriting
// contract on ChainSelector should receive the update.
type ApplyInput struct {
	ChainSelector uint64            `json:"chainSelector,string" yaml:"chainSelector"`
	ContractType  cldf.ContractType `json:"contractType"         yaml:"contractType"`
	Version       *semver.Version   `json:"version"              yaml:"version"`
	Update        CallerUpdate      `json:"update"               yaml:"update"`
}
