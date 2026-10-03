package adapters_test

import (
	"errors"
	"testing"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
)

// fakeAdapter is a hand-rolled GasUpdateAdapter for exercising GlamsterdamGasUpdateSequence's
// orchestration logic without any real chain family behind it. Each method delegates to an
// optional closure field; tests only set the closures relevant to the scenario they exercise, and
// any unset closure that ends up called fails the test loudly (via requireNotCalled) rather than
// panicking on a nil dereference, so a test's assumptions about which methods should run are
// explicit and self-checking.
type fakeAdapter struct {
	t *testing.T

	hasLaneToTarget           func(chainSel uint64) (bool, error)
	readDestGasFields         func(chainSel uint64) (map[string]uint32, error)
	writeDestGasFields        func(chainSel uint64, resolved map[string]uint32) ([]mcms_types.BatchOperation, error)
	readImmutableSanityFields func(chainSel uint64) (map[string]uint32, error)
	discoverCandidateTokens   func(chainSel uint64) ([][]byte, error)
	readTokenGasField         func(chainSel uint64, token []byte) (uint32, bool, error)
	writeTokenGasField        func(chainSel uint64, token []byte, value uint32) (mcms_types.BatchOperation, error)
}

func (f *fakeAdapter) HasLaneToTarget(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector, _ uint64) (bool, error) {
	if f.hasLaneToTarget == nil {
		f.t.Fatal("HasLaneToTarget called but not stubbed")
	}
	return f.hasLaneToTarget(srcChainSelector)
}

func (f *fakeAdapter) ReadDestGasFields(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector, _ uint64) (map[string]uint32, error) {
	if f.readDestGasFields == nil {
		f.t.Fatal("ReadDestGasFields called but not stubbed")
	}
	return f.readDestGasFields(srcChainSelector)
}

func (f *fakeAdapter) WriteDestGasFields(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector, _ uint64, resolved map[string]uint32) ([]mcms_types.BatchOperation, error) {
	if f.writeDestGasFields == nil {
		f.t.Fatal("WriteDestGasFields called but not stubbed")
	}
	return f.writeDestGasFields(srcChainSelector, resolved)
}

func (f *fakeAdapter) ReadImmutableSanityFields(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector uint64) (map[string]uint32, error) {
	if f.readImmutableSanityFields == nil {
		f.t.Fatal("ReadImmutableSanityFields called but not stubbed")
	}
	return f.readImmutableSanityFields(srcChainSelector)
}

func (f *fakeAdapter) DiscoverCandidateTokens(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector uint64) ([][]byte, error) {
	if f.discoverCandidateTokens == nil {
		f.t.Fatal("DiscoverCandidateTokens called but not stubbed")
	}
	return f.discoverCandidateTokens(srcChainSelector)
}

func (f *fakeAdapter) ReadTokenGasField(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector, _ uint64, token []byte) (uint32, bool, error) {
	if f.readTokenGasField == nil {
		f.t.Fatal("ReadTokenGasField called but not stubbed")
	}
	return f.readTokenGasField(srcChainSelector, token)
}

func (f *fakeAdapter) WriteTokenGasField(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector, _ uint64, token []byte, value uint32) (mcms_types.BatchOperation, error) {
	if f.writeTokenGasField == nil {
		f.t.Fatal("WriteTokenGasField called but not stubbed")
	}
	return f.writeTokenGasField(srcChainSelector, token, value)
}

func batchOp(chainSel uint64) mcms_types.BatchOperation {
	return mcms_types.BatchOperation{
		ChainSelector: mcms_types.ChainSelector(chainSel),
		Transactions:  []mcms_types.Transaction{{}},
	}
}

func runSequence(t *testing.T, adapter adapters.GasUpdateAdapter, target uint64, candidates []uint64) (adapters.GlamsterdamGasUpdateSequenceOutput, *glamsterdamutils.Report) {
	t.Helper()
	report := glamsterdamutils.NewReport()
	out, err := adapters.GlamsterdamGasUpdateSequence(
		cldf_ops.Bundle{},
		cldf_chain.BlockChains{},
		datastore.NewMemoryDataStore().Seal(),
		adapters.GlamsterdamGasUpdateSequenceInput{
			Adapter:                 adapter,
			TargetChainSelector:     target,
			CandidateChainSelectors: candidates,
			Report:                  report,
		},
	)
	require.NoError(t, err)
	return out, report
}

// TestGlamsterdamGasUpdateSequence_NoLane verifies a chain with no lane to the target is skipped
// entirely — nothing past HasLaneToTarget should be called for it.
func TestGlamsterdamGasUpdateSequence_NoLane(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(sel uint64) (bool, error) {
			require.Equal(t, chainSel, sel)
			return false, nil
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "chain 111: no lane to target chain, skipped")
}

// TestGlamsterdamGasUpdateSequence_HasLaneError verifies a hard read error checking the lane is
// reported and the chain is skipped, without aborting the whole run.
func TestGlamsterdamGasUpdateSequence_HasLaneError(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(sel uint64) (bool, error) {
			return false, errors.New("rpc unavailable")
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "chain 111: ERROR - failed to check lane to target: rpc unavailable, skipping this chain")
}

// TestGlamsterdamGasUpdateSequence_DestGasFieldsMatchAndFallback exercises the full dest-gas-field
// resolution and write path: one field exactly at its Prague baseline (literal Glamsterdam value
// applied) and one field mismatched (fallback ratio applied), verifying WriteDestGasFields
// receives the correctly resolved values and its returned batch op is propagated.
func TestGlamsterdamGasUpdateSequence_DestGasFieldsMatchAndFallback(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	var gotResolved map[string]uint32
	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(uint64) (map[string]uint32, error) {
			return map[string]uint32{
				adapters.FeeQuoterDestGasOverhead.Name:             300_000, // matches Prague exactly
				adapters.FeeQuoterDefaultTokenDestGasOverhead.Name: 100_000, // mismatched -> fallback
			}, nil
		},
		writeDestGasFields: func(sel uint64, resolved map[string]uint32) ([]mcms_types.BatchOperation, error) {
			require.Equal(t, chainSel, sel)
			gotResolved = resolved
			return []mcms_types.BatchOperation{batchOp(sel)}, nil
		},
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens:   func(uint64) ([][]byte, error) { return nil, nil },
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Len(t, out.BatchOps, 1)

	require.Equal(t, uint32(500_000), gotResolved[adapters.FeeQuoterDestGasOverhead.Name])
	// 100_000 current * (270_000/90_000 Glamsterdam/Prague ratio) = 300_000
	require.Equal(t, uint32(300_000), gotResolved[adapters.FeeQuoterDefaultTokenDestGasOverhead.Name])

	require.Contains(t, report.String(), "FeeQuoter.DestChainConfig.DestGasOverhead matched expected Prague value 300000, applying Glamsterdam value 500000")
	require.Contains(t, report.String(), "FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead MISMATCH")
}

// TestGlamsterdamGasUpdateSequence_ImmutableSanityMismatchWarns verifies a mismatched immutable
// field produces a warning report line without attempting any write (there is no setter).
func TestGlamsterdamGasUpdateSequence_ImmutableSanityMismatchWarns(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:   func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) {
			return map[string]uint32{"OffRamp.GasForCallExactCheck": 9_999}, nil
		},
		discoverCandidateTokens: func(uint64) ([][]byte, error) { return nil, nil },
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "WARNING - OffRamp.GasForCallExactCheck is 9999, expected 5000 (immutable, cannot be changed)")
}

// TestGlamsterdamGasUpdateSequence_TokenFieldResolvedAndWritten verifies a discovered token with a
// mismatched current value gets resolved against the USDC field spec and written back.
func TestGlamsterdamGasUpdateSequence_TokenFieldResolvedAndWritten(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)
	token := []byte{0xAA, 0xBB}

	var wroteValue uint32
	var wroteToken []byte
	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:           func(uint64) (bool, error) { return true, nil },
		readDestGasFields:         func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens:   func(uint64) ([][]byte, error) { return [][]byte{token}, nil },
		readTokenGasField: func(sel uint64, tok []byte) (uint32, bool, error) {
			require.Equal(t, token, tok)
			return 180_000, true, nil // matches USDC Prague baseline exactly
		},
		writeTokenGasField: func(sel uint64, tok []byte, value uint32) (mcms_types.BatchOperation, error) {
			wroteToken = tok
			wroteValue = value
			return batchOp(sel), nil
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Len(t, out.BatchOps, 1)
	require.Equal(t, token, wroteToken)
	require.Equal(t, uint32(540_000), wroteValue)
	require.Contains(t, report.String(), "FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC) matched expected Prague value 180000, applying Glamsterdam value 540000")
}

// TestGlamsterdamGasUpdateSequence_TokenAlreadyAppliedSkipsWrite verifies that a token whose
// current value already equals the literal Glamsterdam value (e.g. a rerun after a prior
// successful proposal) is treated as a no-op and never reaches WriteTokenGasField.
func TestGlamsterdamGasUpdateSequence_TokenAlreadyAppliedSkipsWrite(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)
	token := []byte{0xAA, 0xBB}

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:           func(uint64) (bool, error) { return true, nil },
		readDestGasFields:         func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens:   func(uint64) ([][]byte, error) { return [][]byte{token}, nil },
		readTokenGasField: func(uint64, []byte) (uint32, bool, error) {
			return adapters.USDCTokenPoolDestGasOverhead.GlamsterdamValue, true, nil
		},
		// writeTokenGasField deliberately left nil: calling it fails the test.
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "already matches Glamsterdam value 540000")
}

// TestGlamsterdamGasUpdateSequence_TokenNotConfiguredSkipped verifies a discovered token with no
// enabled override is skipped silently, without attempting to resolve or write it.
func TestGlamsterdamGasUpdateSequence_TokenNotConfiguredSkipped(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)
	token := []byte{0xAA, 0xBB}

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:           func(uint64) (bool, error) { return true, nil },
		readDestGasFields:         func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens:   func(uint64) ([][]byte, error) { return [][]byte{token}, nil },
		readTokenGasField: func(uint64, []byte) (uint32, bool, error) {
			return 0, false, nil
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.NotContains(t, report.String(), "TokenTransferFeeConfig")
}

// TestGlamsterdamGasUpdateSequence_DiscoverTokensErrorIsolated verifies a failure discovering
// candidate tokens is reported but doesn't discard dest-gas-field writes already produced for
// that same chain.
func TestGlamsterdamGasUpdateSequence_DiscoverTokensErrorIsolated(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(uint64) (map[string]uint32, error) {
			return map[string]uint32{adapters.FeeQuoterDestGasOverhead.Name: 300_000}, nil
		},
		writeDestGasFields: func(sel uint64, _ map[string]uint32) ([]mcms_types.BatchOperation, error) {
			return []mcms_types.BatchOperation{batchOp(sel)}, nil
		},
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens: func(uint64) ([][]byte, error) {
			return nil, errors.New("token registry unreachable")
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Len(t, out.BatchOps, 1) // the dest-gas write still made it through
	require.Contains(t, report.String(), "chain 111: ERROR - failed to discover candidate tokens: token registry unreachable, skipping this chain")
}

// TestGlamsterdamGasUpdateSequence_MultipleChainsIndependent verifies chains are processed
// independently: one chain's write failure doesn't block another chain's successful update.
func TestGlamsterdamGasUpdateSequence_MultipleChainsIndependent(t *testing.T) {
	const chainA = uint64(111)
	const chainB = uint64(222)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(sel uint64) (map[string]uint32, error) {
			return map[string]uint32{adapters.FeeQuoterDestGasOverhead.Name: 300_000}, nil
		},
		writeDestGasFields: func(sel uint64, _ map[string]uint32) ([]mcms_types.BatchOperation, error) {
			if sel == chainA {
				return nil, errors.New("mcms simulation failed")
			}
			return []mcms_types.BatchOperation{batchOp(sel)}, nil
		},
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens:   func(uint64) ([][]byte, error) { return nil, nil },
	}

	out, report := runSequence(t, adapter, target, []uint64{chainA, chainB})
	require.Len(t, out.BatchOps, 1)
	require.Equal(t, mcms_types.ChainSelector(chainB), out.BatchOps[0].ChainSelector)
	require.Contains(t, report.String(), "chain 111: ERROR - failed to write dest gas fields: mcms simulation failed, skipping this chain")
}
