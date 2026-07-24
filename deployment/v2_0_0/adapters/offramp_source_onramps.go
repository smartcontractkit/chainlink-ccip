package adapters

import (
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// OffRampSetSourceOnRampsEntry sets the full source-chain onramp whitelist on a chain's OffRamp.
// OffRamp.applySourceChainConfigUpdates replaces onRamps entirely; pass every address that should
// remain allowed (e.g. [new, old] during drain, [new] after drain).
type OffRampSetSourceOnRampsEntry struct {
	LocalChainSelector  uint64 `json:"localChainSelector" yaml:"localChainSelector"`
	SourceChainSelector uint64 `json:"sourceChainSelector" yaml:"sourceChainSelector"`
	// OnRamps are hex strings (20- or 32-byte) in durable-pipeline input, not [][]byte: operators
	// author yaml/json as human-readable hex, and each chain-family adapter encodes to on-chain bytes.
	OnRamps []string `json:"onRamps" yaml:"onRamps"`
}

// OffRampSourceOnRampSetter is implemented by chain family adapters that support surgical
// OffRamp source onramp whitelist updates.
type OffRampSourceOnRampSetter interface {
	SetOffRampSourceOnRamps(e cldf.Environment, update OffRampSetSourceOnRampsEntry) (*mcms_types.BatchOperation, bool, error)
}
