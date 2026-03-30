package fqdests

import (
	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// DestChainConfigForDst represents a destination chain config update for a single destination.
type DestChainConfigForDst struct {
	Selector uint64                        `json:"selector" yaml:"selector"`
	Config   lanes.FeeQuoterDestChainConfig `json:"config" yaml:"config"`
}

// DestChainConfigForSrc represents all destination chain config updates originating from a single source.
type DestChainConfigForSrc struct {
	Selector uint64                  `json:"selector" yaml:"selector"`
	Settings []DestChainConfigForDst `json:"settings" yaml:"settings"`
}

// UpdateFeeQuoterDestsInput is the top-level input for the UpdateFeeQuoterDests changeset.
type UpdateFeeQuoterDestsInput struct {
	// Version is the lane version used for initial adapter lookup.
	Version *semver.Version          `json:"version" yaml:"version"`
	Args    []DestChainConfigForSrc  `json:"args" yaml:"args"`
	MCMS    mcms.Input               `json:"mcms" yaml:"mcms"`
}

// ApplyDestChainConfigSequenceInput is the input for a chain-specific adapter sequence.
type ApplyDestChainConfigSequenceInput struct {
	// Selector is the source chain selector.
	Selector uint64 `json:"selector" yaml:"selector"`
	// Settings maps destination chain selector -> FeeQuoterDestChainConfig.
	Settings map[uint64]lanes.FeeQuoterDestChainConfig `json:"settings" yaml:"settings"`
}
