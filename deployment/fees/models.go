package fees

import (
	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// ApplyDestChainConfigSequenceInput is the input for a chain-specific adapter sequence
// that applies FeeQuoter destination chain config updates.
type ApplyDestChainConfigSequenceInput struct {
	// Selector is the source chain selector.
	Selector uint64 `json:"selector" yaml:"selector"`
	// Settings maps destination chain selector to its FeeQuoterDestChainConfig.
	Settings map[uint64]lanes.FeeQuoterDestChainConfig `json:"settings" yaml:"settings"`
}

// DestChainConfigForDst represents a destination chain config override for a single destination.
// Override is a functional option that mutates the base config (read from on-chain or defaults).
// If Override is nil, the existing on-chain config is re-applied as-is (noop).
type DestChainConfigForDst struct {
	Selector uint64                                  `json:"selector" yaml:"selector"`
	Override *lanes.FeeQuoterDestChainConfigOverride `json:"-" yaml:"-"`
}

// DestChainConfigForSrc represents all destination chain config updates originating from a single source.
type DestChainConfigForSrc struct {
	Selector uint64                  `json:"selector" yaml:"selector"`
	Settings []DestChainConfigForDst `json:"settings" yaml:"settings"`
}

// UpdateFeeQuoterDestsInput is the top-level input for the UpdateFeeQuoterDests changeset.
type UpdateFeeQuoterDestsInput struct {
	// Version is the lane version used for initial adapter lookup.
	Version *semver.Version         `json:"version" yaml:"version"`
	Args    []DestChainConfigForSrc `json:"args" yaml:"args"`
	MCMS    mcms.Input              `json:"mcms" yaml:"mcms"`
}

// SetTokenTransferFeeSequenceInput defines the input for setting token transfer fee configurations in a sequence.
type SetTokenTransferFeeSequenceInput struct {
	Settings map[uint64]map[string]*TokenTransferFeeArgs `json:"settings" yaml:"settings"`
	Selector uint64                                      `json:"selector" yaml:"selector"`
}

// TokenTransferFeeArgs defines the standardized configuration for token transfer fees for all chain families.
type TokenTransferFeeArgs struct {
	DestBytesOverhead uint32 `json:"destBytesOverhead" yaml:"destBytesOverhead"`
	DestGasOverhead   uint32 `json:"destGasOverhead" yaml:"destGasOverhead"`
	MinFeeUSDCents    uint32 `json:"minFeeUSDCents" yaml:"minFeeUSDCents"`
	MaxFeeUSDCents    uint32 `json:"maxFeeUSDCents" yaml:"maxFeeUSDCents"`
	DeciBps           uint16 `json:"deciBps" yaml:"deciBps"`
	IsEnabled         bool   `json:"isEnabled" yaml:"isEnabled"`
}

// UnresolvedTokenTransferFeeArgs allows for partial specification of token transfer fee configurations.
type UnresolvedTokenTransferFeeArgs struct {
	DestBytesOverhead TokenTransferFeeValue[uint32] `json:"destBytesOverhead" yaml:"destBytesOverhead"`
	DestGasOverhead   TokenTransferFeeValue[uint32] `json:"destGasOverhead" yaml:"destGasOverhead"`
	MinFeeUSDCents    TokenTransferFeeValue[uint32] `json:"minFeeUSDCents" yaml:"minFeeUSDCents"`
	MaxFeeUSDCents    TokenTransferFeeValue[uint32] `json:"maxFeeUSDCents" yaml:"maxFeeUSDCents"`
	DeciBps           TokenTransferFeeValue[uint16] `json:"deciBps" yaml:"deciBps"`
	IsEnabled         TokenTransferFeeValue[bool]   `json:"isEnabled" yaml:"isEnabled"`
}

// Infer fills in any unset fields in the unresolved configuration using the provided fallback values.
func (cfg UnresolvedTokenTransferFeeArgs) Infer(fallbacks TokenTransferFeeArgs) *TokenTransferFeeArgs {
	return &TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead.Infer(fallbacks.DestBytesOverhead),
		DestGasOverhead:   cfg.DestGasOverhead.Infer(fallbacks.DestGasOverhead),
		MinFeeUSDCents:    cfg.MinFeeUSDCents.Infer(fallbacks.MinFeeUSDCents),
		MaxFeeUSDCents:    cfg.MaxFeeUSDCents.Infer(fallbacks.MaxFeeUSDCents),
		IsEnabled:         cfg.IsEnabled.Infer(fallbacks.IsEnabled),
		DeciBps:           cfg.DeciBps.Infer(fallbacks.DeciBps),
	}
}

// TokenTransferFeeValue represents a value that may or may not be explicitly set.
type TokenTransferFeeValue[T any] struct {
	// If set to false (the default), then `Value` will be autofilled from
	// on-chain data if it exists. If no on-chain data exists for it, then
	// a pre-selected sensible default will be used as a fallback
	Valid bool `json:"valid" yaml:"valid"`

	// This only has an effect when `Valid` is set to true. If this is the
	// case, then the provided value will overwrite existing on-chain data
	Value T `json:"value" yaml:"value"`
}

// Infer returns the contained value if `Valid` is true; otherwise, it returns the provided fallback.
func (cfg TokenTransferFeeValue[T]) Infer(fallback T) T {
	if cfg.Valid {
		return cfg.Value
	}

	return fallback
}
