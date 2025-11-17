package fees

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
