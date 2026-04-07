package fees

import (
	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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
// If Override is nil, the existing on-chain config is re-applied as-is (idempotent re-apply).
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
	DestBytesOverhead utils.Optional[uint32] `json:"destBytesOverhead" yaml:"destBytesOverhead"`
	DestGasOverhead   utils.Optional[uint32] `json:"destGasOverhead" yaml:"destGasOverhead"`
	MinFeeUSDCents    utils.Optional[uint32] `json:"minFeeUSDCents" yaml:"minFeeUSDCents"`
	MaxFeeUSDCents    utils.Optional[uint32] `json:"maxFeeUSDCents" yaml:"maxFeeUSDCents"`
	DeciBps           utils.Optional[uint16] `json:"deciBps" yaml:"deciBps"`
	IsEnabled         utils.Optional[bool]   `json:"isEnabled" yaml:"isEnabled"`
}

// Resolve fills in any unset fields in the unresolved configuration using the provided fallback values.
func (cfg UnresolvedTokenTransferFeeArgs) Resolve(fallbacks TokenTransferFeeArgs) *TokenTransferFeeArgs {
	return &TokenTransferFeeArgs{
		DestBytesOverhead: cfg.DestBytesOverhead.GetOrDefault(fallbacks.DestBytesOverhead),
		DestGasOverhead:   cfg.DestGasOverhead.GetOrDefault(fallbacks.DestGasOverhead),
		MinFeeUSDCents:    cfg.MinFeeUSDCents.GetOrDefault(fallbacks.MinFeeUSDCents),
		MaxFeeUSDCents:    cfg.MaxFeeUSDCents.GetOrDefault(fallbacks.MaxFeeUSDCents),
		IsEnabled:         cfg.IsEnabled.GetOrDefault(fallbacks.IsEnabled),
		DeciBps:           cfg.DeciBps.GetOrDefault(fallbacks.DeciBps),
	}
}
