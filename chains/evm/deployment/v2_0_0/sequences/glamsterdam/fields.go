package glamsterdam

import (
	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/glamsterdam"
)

// Field specs for the v2.0 Glamsterdam gas-config mapping table. Values sourced from
// "Glamsterdam Gas Configuration.md" (embedded in GLAMSTERDAM_GAS_UPDATE_PLAN.md §6, V2.0 table).
var (
	// OnRampBaseExecutionGasCost is table row 1.
	OnRampBaseExecutionGasCost = glamsterdamutils.FieldSpec[uint32]{
		Name:             "OnRamp.DestChainConfig.BaseExecutionGasCost",
		ExpectedPrague:   200_000,
		GlamsterdamValue: 400_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](200_000, 400_000),
	}

	// FeeQuoterDefaultTokenDestGasOverhead is table row 2.
	FeeQuoterDefaultTokenDestGasOverhead = glamsterdamutils.FieldSpec[uint32]{
		Name:             "FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead",
		ExpectedPrague:   90_000,
		GlamsterdamValue: 270_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](90_000, 270_000),
	}

	// FeeQuoterMaxPerMsgGasLimit is table row 3. Prague and Glamsterdam values are identical per
	// the doc ("needs to be re-tested once testnet is live"), so this is a no-op for now: on a
	// mismatch, the current value is left unchanged rather than scaled.
	FeeQuoterMaxPerMsgGasLimit = glamsterdamutils.FieldSpec[uint32]{
		Name:             "FeeQuoter.DestChainConfig.MaxPerMsgGasLimit",
		ExpectedPrague:   15_000_000,
		GlamsterdamValue: 15_000_000,
		Fallback:         func(current uint32) uint32 { return current },
	}

	// FeeQuoterDestGasPerPayloadByteBase is table row 4. The "+DA" components
	// (DestGasPerDataAvailabilityByte et al.) are confirmed unchanged by the doc and are not
	// part of this field spec.
	FeeQuoterDestGasPerPayloadByteBase = glamsterdamutils.FieldSpec[uint8]{
		Name:             "FeeQuoter.DestChainConfig.DestGasPerPayloadByteBase",
		ExpectedPrague:   20,
		GlamsterdamValue: 64,
		Fallback:         glamsterdamutils.ApplyRatio[uint8](20, 64),
	}

	// FeeQuoterDefaultTxGasLimit is table row 5.
	FeeQuoterDefaultTxGasLimit = glamsterdamutils.FieldSpec[uint32]{
		Name:             "FeeQuoter.DestChainConfig.DefaultTxGasLimit",
		ExpectedPrague:   200_000,
		GlamsterdamValue: 400_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](200_000, 400_000),
	}

	// CommitteeVerifierGasForVerification is table row 8.
	CommitteeVerifierGasForVerification = glamsterdamutils.FieldSpec[uint32]{
		Name:             "CommitteeVerifier.RemoteChainConfigArgs.GasForVerification",
		ExpectedPrague:   75_000,
		GlamsterdamValue: 85_000,
		Fallback:         glamsterdamutils.ApplyRatio[uint32](75_000, 85_000),
	}
)

// OffRamp fields (table rows 6-7) are immutable, set only in the constructor, with no setter —
// StaticConfigCannotBeChanged reverts any attempt to change them. They are read-only sanity
// checked (never written) against their expected Prague baseline, which should equal their
// Glamsterdam value.
const (
	OffRampExpectedGasForCallExactCheck      = uint16(5_000)
	OffRampExpectedMaxGasBufferToUpdateState = uint32(12_000)
)
