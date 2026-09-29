package tokens

import (
	"fmt"
	"math/big"
	"strconv"
)

// LiquidityMigrationAmountFormat selects how a liquidity migration amount is interpreted.
type LiquidityMigrationAmountFormat string

const (
	// LiquidityMigrationAmountFormatRAW specifies an exact token amount to migrate, in the smallest unit of the token (wei for ETH, or the token's decimals).
	LiquidityMigrationAmountFormatRAW LiquidityMigrationAmountFormat = "raw"
	// LiquidityMigrationAmountFormatBPS specifies a percentage (1-10000, where 10000 = 100%) of the old pool's balance to migrate.
	LiquidityMigrationAmountFormatBPS LiquidityMigrationAmountFormat = "bps"
)

// LockReleasePoolLiquidityMigrationAmount expresses a liquidity migration amount as either an exact
// raw token amount (LiquidityMigrationAmountFormatRAW) or a basis-points percentage of the old
// pool's balance (LiquidityMigrationAmountFormatBPS).
type LockReleasePoolLiquidityMigrationAmount struct {
	Format LiquidityMigrationAmountFormat `yaml:"format" json:"format"`
	Value  string                         `yaml:"value" json:"value"`
}

// BasisPoints returns the amount as basis points when the format is
// LiquidityMigrationAmountFormatBPS, or (nil, nil) for any other format.
func (a *LockReleasePoolLiquidityMigrationAmount) BasisPoints() (*uint16, error) {
	if a.Format != LiquidityMigrationAmountFormatBPS {
		return nil, nil
	}
	val, err := strconv.ParseUint(a.Value, 10, 16)
	if err != nil {
		return nil, fmt.Errorf("invalid value for BasisPoints: %w", err)
	}
	if val == 0 || val > 10000 {
		return nil, fmt.Errorf("invalid value for BasisPoints: must be between 1 and 10000, got %d", val)
	}
	return new(uint16(val)), nil
}

// RawAmount returns the amount as a raw token amount when the format is
// LiquidityMigrationAmountFormatRAW, or (nil, nil) for any other format.
func (a *LockReleasePoolLiquidityMigrationAmount) RawAmount() (*big.Int, error) {
	if a.Format != LiquidityMigrationAmountFormatRAW {
		return nil, nil
	}
	val, ok := new(big.Int).SetString(a.Value, 10)
	if !ok {
		return nil, fmt.Errorf("invalid value for RawAmount: %s", a.Value)
	}
	if val.Sign() <= 0 {
		return nil, fmt.Errorf("invalid value for RawAmount: must be positive, got %s", a.Value)
	}
	return val, nil
}

// Validate reports whether the amount is well-formed for its format: the format must be known, the
// value must be present, and the value must parse and satisfy the format's constraints.
func (a *LockReleasePoolLiquidityMigrationAmount) Validate() error {
	if a.Value == "" {
		return fmt.Errorf("value must be provided")
	}
	switch a.Format {
	case LiquidityMigrationAmountFormatBPS:
		if _, err := a.BasisPoints(); err != nil {
			return err
		}
	case LiquidityMigrationAmountFormatRAW:
		if _, err := a.RawAmount(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid format: must be 'bps' or 'raw', got '%s'", a.Format)
	}
	return nil
}
