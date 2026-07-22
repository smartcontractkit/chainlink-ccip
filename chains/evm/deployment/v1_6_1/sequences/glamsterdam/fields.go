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
)

// OffRamp.GasForCallExactCheck (table row 3) is immutable, set only in the constructor, with no
// setter. It is read-only sanity checked (never written) against its expected Prague baseline,
// which should equal its Glamsterdam value.
const OffRampExpectedGasForCallExactCheck = uint16(5_000)
