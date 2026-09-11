package glamsterdam

import (
	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
)

// Field specs for the v1.6 Glamsterdam gas-config mapping table. Values sourced from
// "Glamsterdam Gas Configuration.md" (embedded in GLAMSTERDAM_GAS_UPDATE_PLAN.md §6, V1.6 table).
var (
	// FeeQuoterDestGasOverhead is table row 1. The doc labels this "OnRamp::baseExecutionGasCost",
	// but v1.6 OnRamp has no gas field at all — this concept lives on FeeQuoter for v1.6 (see
	// GLAMSTERDAM_GAS_UPDATE_PLAN.md §2.1).
	FeeQuoterDestGasOverhead = glamsterdamutils.FieldSpec[uint32]{
		Name:             "FeeQuoter.DestChainConfig.DestGasOverhead",
		ExpectedPrague:   300_000,
		GlamsterdamValue: 500_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](300_000, 500_000),
	}

	// FeeQuoterDefaultTokenDestGasOverhead is table row 2.
	FeeQuoterDefaultTokenDestGasOverhead = glamsterdamutils.FieldSpec[uint32]{
		Name:             "FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead",
		ExpectedPrague:   90_000,
		GlamsterdamValue: 270_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](90_000, 270_000),
	}

	// USDCTokenPoolDestGasOverhead is table row 5. The doc labels this "USDCTokenPool", but it is
	// actually a FeeQuoter.TokenTransferFeeConfig.DestGasOverhead setting for v1.6, keyed by
	// (destChainSelector, token) — see GLAMSTERDAM_GAS_UPDATE_PLAN.md §2.3. A single FeeQuoter
	// setting governs the gas overhead for all v1.6 pools of that token on the lane; there is no
	// per-pool override in v1.6. Guesstimate value (real value needs testnet measurement), applied
	// as-is for the first (testnet) run.
	USDCTokenPoolDestGasOverhead = glamsterdamutils.FieldSpec[uint32]{
		Name:             "FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC)",
		ExpectedPrague:   180_000,
		GlamsterdamValue: 540_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](180_000, 540_000),
	}
)

// OffRamp.GasForCallExactCheck (table row 3) is immutable, set only in the constructor, with no
// setter. It is read-only sanity checked (never written) against its expected Prague baseline,
// which should equal its Glamsterdam value.
const OffRampExpectedGasForCallExactCheck = uint16(5_000)
