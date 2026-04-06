package offchain

import (
	"fmt"
	"strconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
)

// FeeQuoterDefaultConfig is the complete fee quoter configuration used as a family-level
// default. All fields are required. ChainFamilySelector is excluded because it is derived
// at conversion time from the destination chain selector. OverrideExistingConfig is a
// runtime flag that lives on the changeset input, not in the topology.
type FeeQuoterDefaultConfig struct {
	IsEnabled                   bool   `toml:"is_enabled"`
	MaxDataBytes                uint32 `toml:"max_data_bytes"`
	MaxPerMsgGasLimit           uint32 `toml:"max_per_msg_gas_limit"`
	DestGasOverhead             uint32 `toml:"dest_gas_overhead"`
	DestGasPerPayloadByteBase   uint8  `toml:"dest_gas_per_payload_byte_base"`
	DefaultTokenFeeUSDCents     uint16 `toml:"default_token_fee_usd_cents"`
	DefaultTokenDestGasOverhead uint32 `toml:"default_token_dest_gas_overhead"`
	DefaultTxGasLimit           uint32 `toml:"default_tx_gas_limit"`
	NetworkFeeUSDCents          uint16 `toml:"network_fee_usd_cents"`
	LinkFeeMultiplierPercent    uint8  `toml:"link_fee_multiplier_percent"`
	USDPerUnitGas               int64  `toml:"usd_per_unit_gas"`
}

// FeeQuoterOverrideConfig is a partial fee quoter configuration used for dest-chain-level,
// source-chain-level, and lane-level overrides. Nil pointer fields inherit from the lower-priority level.
type FeeQuoterOverrideConfig struct {
	IsEnabled                   *bool   `toml:"is_enabled,omitempty"`
	MaxDataBytes                *uint32 `toml:"max_data_bytes,omitempty"`
	MaxPerMsgGasLimit           *uint32 `toml:"max_per_msg_gas_limit,omitempty"`
	DestGasOverhead             *uint32 `toml:"dest_gas_overhead,omitempty"`
	DestGasPerPayloadByteBase   *uint8  `toml:"dest_gas_per_payload_byte_base,omitempty"`
	DefaultTokenFeeUSDCents     *uint16 `toml:"default_token_fee_usd_cents,omitempty"`
	DefaultTokenDestGasOverhead *uint32 `toml:"default_token_dest_gas_overhead,omitempty"`
	DefaultTxGasLimit           *uint32 `toml:"default_tx_gas_limit,omitempty"`
	NetworkFeeUSDCents          *uint16 `toml:"network_fee_usd_cents,omitempty"`
	LinkFeeMultiplierPercent    *uint8  `toml:"link_fee_multiplier_percent,omitempty"`
	USDPerUnitGas               *int64  `toml:"usd_per_unit_gas,omitempty"`
}

// ApplyOverride returns a new FeeQuoterDefaultConfig with non-nil fields from the
// override merged on top of the receiver. The receiver is not modified.
func (c FeeQuoterDefaultConfig) ApplyOverride(o *FeeQuoterOverrideConfig) FeeQuoterDefaultConfig {
	if o == nil {
		return c
	}
	if o.IsEnabled != nil {
		c.IsEnabled = *o.IsEnabled
	}
	if o.MaxDataBytes != nil {
		c.MaxDataBytes = *o.MaxDataBytes
	}
	if o.MaxPerMsgGasLimit != nil {
		c.MaxPerMsgGasLimit = *o.MaxPerMsgGasLimit
	}
	if o.DestGasOverhead != nil {
		c.DestGasOverhead = *o.DestGasOverhead
	}
	if o.DestGasPerPayloadByteBase != nil {
		c.DestGasPerPayloadByteBase = *o.DestGasPerPayloadByteBase
	}
	if o.DefaultTokenFeeUSDCents != nil {
		c.DefaultTokenFeeUSDCents = *o.DefaultTokenFeeUSDCents
	}
	if o.DefaultTokenDestGasOverhead != nil {
		c.DefaultTokenDestGasOverhead = *o.DefaultTokenDestGasOverhead
	}
	if o.DefaultTxGasLimit != nil {
		c.DefaultTxGasLimit = *o.DefaultTxGasLimit
	}
	if o.NetworkFeeUSDCents != nil {
		c.NetworkFeeUSDCents = *o.NetworkFeeUSDCents
	}
	if o.LinkFeeMultiplierPercent != nil {
		c.LinkFeeMultiplierPercent = *o.LinkFeeMultiplierPercent
	}
	if o.USDPerUnitGas != nil {
		c.USDPerUnitGas = *o.USDPerUnitGas
	}
	return c
}

// FeeQuoterTopology groups the hierarchical fee quoter configuration.
// Resolution priority: family default (lowest) → dest chain → source chain → lane (highest).
//
// When resolving a lane (local/source chain S → remote/destination chain D), the merged config is
// built for configuring D on S's FeeQuoter. Use dest_chain_defaults for values shared by every
// source that talks to D. Use source_chain_defaults for values that apply to every lane from S
// (e.g. NetworkFeeUSDCents, DefaultTokenFeeUSDCents) and must override dest for that source.
// Use lane_defaults for the most specific S→D-only tweaks.
type FeeQuoterTopology struct {
	FamilyDefaults map[string]FeeQuoterDefaultConfig `toml:"family_defaults"`
	// DestChainDefaults is keyed by destination chain selector (decimal string). The same entry
	// applies to every lane whose remote chain is that selector, before source_chain_defaults and
	// lane_defaults.
	DestChainDefaults map[string]*FeeQuoterOverrideConfig `toml:"dest_chain_defaults"`
	// SourceChainDefaults is keyed by source (local) chain selector (decimal string). Applies to
	// every lane originating from that source; overrides dest_chain_defaults for the same field.
	SourceChainDefaults map[string]*FeeQuoterOverrideConfig `toml:"source_chain_defaults"`
	// LaneDefaults is keyed by source (local) chain selector, then destination chain selector.
	// Overrides source_chain_defaults and dest_chain_defaults for the matching S→D lane.
	LaneDefaults map[string]map[string]*FeeQuoterOverrideConfig `toml:"lane_defaults"`
}

// ResolveFeeQuoterConfigForLane resolves the fee quoter configuration for a specific lane
// by layering overrides on top of the family default:
//  1. Family default (all fields required, keyed by family name)
//  2. Dest chain override (partial, keyed by dest chain selector)
//  3. Source chain override (partial, keyed by source chain selector; wins over dest for same field)
//  4. Lane override (partial, keyed by src→dest chain selector pair)
func (t *FeeQuoterTopology) ResolveFeeQuoterConfigForLane(srcSelector, destSelector uint64) (FeeQuoterDefaultConfig, error) {
	if t == nil {
		return FeeQuoterDefaultConfig{}, fmt.Errorf("fee quoter topology is nil")
	}

	destFamily, err := chainsel.GetSelectorFamily(destSelector)
	if err != nil {
		return FeeQuoterDefaultConfig{}, fmt.Errorf("unknown chain family for dest selector %d: %w", destSelector, err)
	}

	familyDefault, ok := t.FamilyDefaults[destFamily]
	if !ok {
		return FeeQuoterDefaultConfig{}, fmt.Errorf("no fee quoter family default for family %q (dest selector %d)", destFamily, destSelector)
	}

	destKey := strconv.FormatUint(destSelector, 10)
	if destOverride, ok := t.DestChainDefaults[destKey]; ok {
		familyDefault = familyDefault.ApplyOverride(destOverride)
	}

	srcKey := strconv.FormatUint(srcSelector, 10)
	if srcOverride, ok := t.SourceChainDefaults[srcKey]; ok {
		familyDefault = familyDefault.ApplyOverride(srcOverride)
	}

	if lanesByDest, ok := t.LaneDefaults[srcKey]; ok {
		if laneOverride, ok := lanesByDest[destKey]; ok {
			familyDefault = familyDefault.ApplyOverride(laneOverride)
		}
	}

	return familyDefault, nil
}

// Validate checks the FeeQuoterTopology for consistency.
func (t *FeeQuoterTopology) Validate() error {
	if t == nil {
		return fmt.Errorf("fee quoter topology is nil")
	}

	if len(t.FamilyDefaults) == 0 {
		return fmt.Errorf("fee_quoter.family_defaults must contain at least one entry")
	}

	for family, cfg := range t.FamilyDefaults {
		if family == "" {
			return fmt.Errorf("fee_quoter.family_defaults contains empty family key")
		}
		if cfg.MaxDataBytes == 0 {
			return fmt.Errorf("fee_quoter.family_defaults[%q].max_data_bytes must be > 0", family)
		}
		if cfg.MaxPerMsgGasLimit == 0 {
			return fmt.Errorf("fee_quoter.family_defaults[%q].max_per_msg_gas_limit must be > 0", family)
		}
		if cfg.DefaultTxGasLimit == 0 {
			return fmt.Errorf("fee_quoter.family_defaults[%q].default_tx_gas_limit must be > 0", family)
		}
	}

	for key := range t.DestChainDefaults {
		if _, err := strconv.ParseUint(key, 10, 64); err != nil {
			return fmt.Errorf("fee_quoter.dest_chain_defaults has invalid chain selector key %q: %w", key, err)
		}
	}

	for key := range t.SourceChainDefaults {
		if _, err := strconv.ParseUint(key, 10, 64); err != nil {
			return fmt.Errorf("fee_quoter.source_chain_defaults has invalid chain selector key %q: %w", key, err)
		}
	}

	for srcKey, innerMap := range t.LaneDefaults {
		if _, err := strconv.ParseUint(srcKey, 10, 64); err != nil {
			return fmt.Errorf("fee_quoter.lane_defaults has invalid source chain selector key %q: %w", srcKey, err)
		}
		for destKey := range innerMap {
			if _, err := strconv.ParseUint(destKey, 10, 64); err != nil {
				return fmt.Errorf("fee_quoter.lane_defaults[%q] has invalid dest chain selector key %q: %w", srcKey, destKey, err)
			}
		}
	}

	return nil
}
