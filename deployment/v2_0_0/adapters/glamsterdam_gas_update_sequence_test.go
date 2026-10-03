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
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// fakeAdapter is a hand-rolled GasUpdateAdapter for exercising GlamsterdamGasUpdateSequence's
// orchestration logic without any real chain family behind it. Each method delegates to an
// optional closure field; tests only set the closures relevant to the scenario they exercise, and
// any unset closure that ends up called fails the test loudly rather than panicking on a nil
// dereference, so a test's assumptions about which methods should run are explicit and
// self-checking.
type fakeAdapter struct {
	t *testing.T

	hasLaneToTarget           func(chainSel uint64) (bool, error)
	readDestGasFields         func(chainSel uint64) (map[string]uint32, error)
	writeDestGasFields        func(chainSel uint64, resolved map[string]uint32) ([]mcms_types.BatchOperation, error)
	readImmutableSanityFields func(chainSel uint64) (map[string]uint32, error)
	discoverCandidateTokens   func(chainSel uint64) ([][]byte, error)
	readTokenGasField         func(chainSel uint64, token []byte) (uint32, bool, error)
	writeTokenGasField        func(chainSel uint64, token []byte, value uint32) (mcms_types.BatchOperation, error)
	tokenFieldSpec            func(chainSel uint64, token []byte) (glamsterdamutils.FieldSpec[uint32], error)
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

func (f *fakeAdapter) TokenFieldSpec(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, srcChainSelector uint64, token []byte) (glamsterdamutils.FieldSpec[uint32], error) {
	if f.tokenFieldSpec == nil {
		f.t.Fatal("TokenFieldSpec called but not stubbed")
	}
	return f.tokenFieldSpec(srcChainSelector, token)
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

func TestGlamsterdamGasUpdateSequence_HasLaneError(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(uint64) (bool, error) {
			return false, errors.New("rpc unavailable")
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "chain 111: ERROR - failed to check lane to target: rpc unavailable, skipping this chain")
}

// TestGlamsterdamGasUpdateSequence_Uint32AndUint8FieldsResolved verifies both the uint32 field
// specs and the one uint8 field spec (FeeQuoterDestGasPerPayloadByteBase) are read, resolved, and
// passed to WriteDestGasFields with matching types/values.
func TestGlamsterdamGasUpdateSequence_Uint32AndUint8FieldsResolved(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	var gotResolved map[string]uint32
	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(uint64) (map[string]uint32, error) {
			return map[string]uint32{
				adapters.OnRampBaseExecutionGasCost.Name:         200_000, // matches Prague
				adapters.FeeQuoterDestGasPerPayloadByteBase.Name: 20,      // matches Prague, uint8-backed
			}, nil
		},
		writeDestGasFields: func(sel uint64, resolved map[string]uint32) ([]mcms_types.BatchOperation, error) {
			gotResolved = resolved
			return []mcms_types.BatchOperation{batchOp(sel)}, nil
		},
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens:   func(uint64) ([][]byte, error) { return nil, nil },
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Len(t, out.BatchOps, 1)
	require.Equal(t, uint32(400_000), gotResolved[adapters.OnRampBaseExecutionGasCost.Name])
	require.Equal(t, uint32(64), gotResolved[adapters.FeeQuoterDestGasPerPayloadByteBase.Name])
	require.Contains(t, report.String(), "OnRamp.DestChainConfig.BaseExecutionGasCost matched expected Prague value 200000, applying Glamsterdam value 400000")
	require.Contains(t, report.String(), "FeeQuoter.DestChainConfig.DestGasPerPayloadByteBase matched expected Prague value 20, applying Glamsterdam value 64")
}

// TestGlamsterdamGasUpdateSequence_ImmutableSanityMismatchWarns verifies a mismatched immutable
// field produces a warning report line without attempting any write (there is no setter). v2.0.0
// has two immutable sanity fields (OffRamp GasForCallExactCheck and MaxGasBufferToUpdateState);
// this exercises both.
func TestGlamsterdamGasUpdateSequence_ImmutableSanityMismatchWarns(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:   func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) {
			return map[string]uint32{
				"OffRamp.GasForCallExactCheck":      9_999,
				"OffRamp.MaxGasBufferToUpdateState": 1,
			}, nil
		},
		discoverCandidateTokens: func(uint64) ([][]byte, error) { return nil, nil },
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "WARNING - OffRamp.GasForCallExactCheck is 9999, expected 5000 (immutable, cannot be changed)")
	require.Contains(t, report.String(), "WARNING - OffRamp.MaxGasBufferToUpdateState is 1, expected 12000 (immutable, cannot be changed)")
}

// TestGlamsterdamGasUpdateSequence_TokenFieldSpecDistinguishesLombardAndUSDC is the key
// regression test for the "always defaults to USDC" bug: two discovered token-pool candidates,
// one Lombard and one USDC, must each resolve against their own FieldSpec (different Prague and
// Glamsterdam values), not a single shared one.
func TestGlamsterdamGasUpdateSequence_TokenFieldSpecDistinguishesLombardAndUSDC(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)
	lombardPool := []byte{0x01}
	usdcPool := []byte{0x02}

	writtenValues := make(map[string]uint32)
	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:           func(uint64) (bool, error) { return true, nil },
		readDestGasFields:         func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens: func(uint64) ([][]byte, error) {
			return [][]byte{lombardPool, usdcPool}, nil
		},
		readTokenGasField: func(_ uint64, token []byte) (uint32, bool, error) {
			if token[0] == lombardPool[0] {
				return adapters.LombardTokenPoolDestGasOverhead.ExpectedPrague, true, nil
			}
			return adapters.USDCTokenPoolDestGasOverhead.ExpectedPrague, true, nil
		},
		tokenFieldSpec: func(_ uint64, token []byte) (glamsterdamutils.FieldSpec[uint32], error) {
			if token[0] == lombardPool[0] {
				return adapters.LombardTokenPoolDestGasOverhead, nil
			}
			return adapters.USDCTokenPoolDestGasOverhead, nil
		},
		writeTokenGasField: func(_ uint64, token []byte, value uint32) (mcms_types.BatchOperation, error) {
			if token[0] == lombardPool[0] {
				writtenValues["lombard"] = value
			} else {
				writtenValues["usdc"] = value
			}
			return batchOp(chainSel), nil
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Len(t, out.BatchOps, 2)
	require.Equal(t, adapters.LombardTokenPoolDestGasOverhead.GlamsterdamValue, writtenValues["lombard"])
	require.Equal(t, adapters.USDCTokenPoolDestGasOverhead.GlamsterdamValue, writtenValues["usdc"])
	require.Contains(t, report.String(), "TokenPool.TokenTransferFeeConfig.DestGasOverhead (Lombard) matched")
	require.Contains(t, report.String(), "TokenPool.TokenTransferFeeConfig.DestGasOverhead (USDC) matched")
}

// TestGlamsterdamGasUpdateSequence_TokenFieldSpecErrorIsolated verifies a failure resolving one
// token's field spec (e.g. an unrecognized pool kind) is reported and that token is skipped,
// without blocking other tokens on the same chain.
func TestGlamsterdamGasUpdateSequence_TokenFieldSpecErrorIsolated(t *testing.T) {
	const chainSel = uint64(111)
	const target = uint64(999)
	badPool := []byte{0x01}
	goodPool := []byte{0x02}

	var wrote []byte
	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget:           func(uint64) (bool, error) { return true, nil },
		readDestGasFields:         func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		readImmutableSanityFields: func(uint64) (map[string]uint32, error) { return map[string]uint32{}, nil },
		discoverCandidateTokens: func(uint64) ([][]byte, error) {
			return [][]byte{badPool, goodPool}, nil
		},
		readTokenGasField: func(_ uint64, token []byte) (uint32, bool, error) {
			return adapters.USDCTokenPoolDestGasOverhead.ExpectedPrague, true, nil
		},
		tokenFieldSpec: func(_ uint64, token []byte) (glamsterdamutils.FieldSpec[uint32], error) {
			if token[0] == badPool[0] {
				return glamsterdamutils.FieldSpec[uint32]{}, errors.New("unrecognized pool kind")
			}
			return adapters.USDCTokenPoolDestGasOverhead, nil
		},
		writeTokenGasField: func(_ uint64, token []byte, value uint32) (mcms_types.BatchOperation, error) {
			wrote = token
			return batchOp(chainSel), nil
		},
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Len(t, out.BatchOps, 1)
	require.Equal(t, goodPool, wrote)
	require.Contains(t, report.String(), "resolve token field spec for token 01: unrecognized pool kind")
}

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
		tokenFieldSpec: func(uint64, []byte) (glamsterdamutils.FieldSpec[uint32], error) {
			return adapters.USDCTokenPoolDestGasOverhead, nil
		},
		// writeTokenGasField deliberately left nil: calling it fails the test.
	}

	out, report := runSequence(t, adapter, target, []uint64{chainSel})
	require.Empty(t, out.BatchOps)
	require.Contains(t, report.String(), "already matches Glamsterdam value")
}

func TestGlamsterdamGasUpdateSequence_MultipleChainsIndependent(t *testing.T) {
	const chainA = uint64(111)
	const chainB = uint64(222)
	const target = uint64(999)

	adapter := &fakeAdapter{
		t: t,
		hasLaneToTarget: func(uint64) (bool, error) { return true, nil },
		readDestGasFields: func(uint64) (map[string]uint32, error) {
			return map[string]uint32{adapters.OnRampBaseExecutionGasCost.Name: 200_000}, nil
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
