package merkleroot

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var rmnRemoteCfg = testhelpers.CreateRMNRemoteCfg()

//nolint:dupl // test cases can be similar, we don't want to overcomplicate with functions for deduping
func Test_Processor_Outcome(t *testing.T) {
	const (
		chainA = 1
		chainB = 2
		chainC = 3
		chainD = 100
		chainE = 4
		chainF = 5
	)

	var (
		bytes32a = [32]byte{0xa}
		bytes32b = [32]byte{0xb}
	)

	type testCase struct {
		name                               string
		prevOutcome                        Outcome
		q                                  Query
		observations                       []func(tc testCase) Observation
		observers                          []commontypes.OracleID
		bigF                               int
		destChainSel                       cciptypes.ChainSelector
		maxMerkleTreeSize                  uint64
		rmnEnabled                         bool
		multipleReports                    bool
		maxReportTransmissionCheckAttempts int
		expOutcome                         Outcome
		expErr                             bool
	}

	testCases := []testCase{
		{
			name:        "empty previous outcome this should be a ranges outcome",
			prevOutcome: Outcome{},
			q:           Query{},
			observations: []func(tc testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OnRampMaxSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 20},
							{ChainSel: chainC, SeqNum: 30},
						},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10}, // we have to execute 10
							{ChainSel: chainB, SeqNum: 21}, // this is still at 20, nothing to execute
							{ChainSel: chainC, SeqNum: 35}, // this is an unexpected state but should not lead to issues
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1, // 2f+1 observations required for each chain
							chainB: 1,
							chainC: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers: []commontypes.OracleID{
				commontypes.OracleID(1),
				commontypes.OracleID(2),
				commontypes.OracleID(3),
			},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 256,
			rmnEnabled:        false,
			expOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: chainA, SeqNumRange: cciptypes.NewSeqNumRange(10, 10)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 21},
					{ChainSel: chainC, SeqNum: 35},
				},
			},
			expErr: false,
		},
		{
			name:        "no fChain consensus should lead to offRamp sequence numbers not being reported for a chain",
			prevOutcome: Outcome{},
			q:           Query{},
			observations: []func(tc testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OnRampMaxSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 20},
							{ChainSel: chainC, SeqNum: 30},
						},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10}, // we have to execute 10
							{ChainSel: chainB, SeqNum: 21}, // this is still at 20, nothing to execute
							{ChainSel: chainC, SeqNum: 35}, // this is an unexpected state but should not lead to issues
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1, // 2f+1 observations required for each chain
							chainB: 1,
							chainC: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation {
					baseObs := tc.observations[0](tc)
					baseObs.OffRampNextSeqNums = []plugintypes.SeqNumChain{
						{ChainSel: chainA, SeqNum: 10},
						// <-------------- chainB is missing
						{ChainSel: chainC, SeqNum: 35},
					}
					return baseObs
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers: []commontypes.OracleID{
				commontypes.OracleID(1),
				commontypes.OracleID(2),
				commontypes.OracleID(3),
			},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 256,
			rmnEnabled:        false,
			expOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: chainA, SeqNumRange: cciptypes.NewSeqNumRange(10, 10)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainC, SeqNum: 35},
				},
			},
			expErr: false,
		},
		{
			name:        "multiple source chain ranges selected",
			prevOutcome: Outcome{},
			q:           Query{},
			observations: []func(tc testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OnRampMaxSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 20},
							{ChainSel: chainC, SeqNum: 30},
						},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 5},
							{ChainSel: chainC, SeqNum: 1},
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1, // 2f+1 observations required for each chain
							chainB: 1,
							chainC: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers: []commontypes.OracleID{
				commontypes.OracleID(1),
				commontypes.OracleID(2),
				commontypes.OracleID(3),
			},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 256,
			rmnEnabled:        false,
			expOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: chainA, SeqNumRange: cciptypes.NewSeqNumRange(10, 10)},
					{ChainSel: chainB, SeqNumRange: cciptypes.NewSeqNumRange(5, 20)},
					{ChainSel: chainC, SeqNumRange: cciptypes.NewSeqNumRange(1, 30)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 5},
					{ChainSel: chainC, SeqNum: 1},
				},
			},
			expErr: false,
		},
		{
			name:        "multiple source chain ranges but some of them do not reach consensus",
			prevOutcome: Outcome{},
			q:           Query{},
			observations: []func(tc testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OnRampMaxSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 20},
							{ChainSel: chainC, SeqNum: 30},
						},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 5},
							{ChainSel: chainC, SeqNum: 1},
							{ChainSel: chainE, SeqNum: 4}, // <---- but "onRamp" seqNum for chainE does not exist
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1, // 2f+1 observations required for each chain
							chainB: 2, // <----------------------- chainB requires 2f+1=5 observations which not exist
							chainC: 1,
							chainD: 1,
							chainE: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers: []commontypes.OracleID{
				commontypes.OracleID(1),
				commontypes.OracleID(2),
				commontypes.OracleID(3),
			},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10, // <------ notice that this will lead to range truncation
			rmnEnabled:        false,
			expOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: chainA, SeqNumRange: cciptypes.NewSeqNumRange(10, 10)},
					// chainB missing due to not reach 2f+1
					{ChainSel: chainC, SeqNumRange: cciptypes.NewSeqNumRange(1, 10)}, // <--- truncated
					// chainE missing due to onRamp seqNums not observed
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					// chainB missing
					{ChainSel: chainC, SeqNum: 1},
					{ChainSel: chainE, SeqNum: 4},
				},
			},
			expErr: false,
		},
		{
			name: "we are in the building report next phase but leader said we want to retry rmn sigs",
			prevOutcome: Outcome{
				OutcomeType:                     ReportIntervalsSelected,
				ReportTransmissionCheckAttempts: 123, // <--- random value to verify if the same outcome is sent
			},
			q: Query{
				RetryRMNSignatures: true,
			},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OnRampMaxSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
						},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:         []commontypes.OracleID{1, 2, 3},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10,
			rmnEnabled:        true,
			expOutcome: Outcome{
				OutcomeType:                     ReportIntervalsSelected,
				ReportTransmissionCheckAttempts: 123,
			},
			expErr: false,
		},
		{
			name: "outcome of merkle roots and rmn signatures checking",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			q: Query{
				RMNSignatures: &rmn.ReportSignatures{
					Signatures: []*rmnpb.EcdsaSignature{
						{R: bytes32a[:], S: bytes32b[:]},
						{R: bytes32a[:], S: bytes32b[:]},
					},
					LaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainA,
								OnrampAddress:       []byte{0xa},
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 10,
								MaxMsgNr: 15,
							},
							Root: bytes32a[:],
						},
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainF,
								OnrampAddress:       []byte{0xb},
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 10,
								MaxMsgNr: 20,
							},
							Root: bytes32b[:],
						},
					},
				},
			},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:      chainA,
								OnRampAddress: []byte{0xa},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 15),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainB,
								OnRampAddress: []byte{0xb},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 15),
								MerkleRoot:    cciptypes.Bytes32{0xa},
							},
							{
								ChainSel:      chainC,
								OnRampAddress: []byte{0xc},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 15),
								MerkleRoot:    cciptypes.Bytes32{0xa},
							},
							{
								ChainSel:      chainE,
								OnRampAddress: []byte{0xd},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 15),
								MerkleRoot:    cciptypes.Bytes32{0xa},
							},
							{
								ChainSel:      chainF,
								OnRampAddress: []byte{0xb},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
								MerkleRoot:    bytes32b,
							},
						},
						RMNEnabledChains: map[cciptypes.ChainSelector]bool{
							chainA: true,
							chainB: true,
							chainC: true,
							chainD: true,
							chainE: true,
							chainF: true,
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 1,
							chainC: 1,
							chainD: 1,
							chainE: 1,
							chainF: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation {
					baseObs := tc.observations[0](tc)
					baseObs.MerkleRoots = append(baseObs.MerkleRoots[:1], baseObs.MerkleRoots[2:]...) // skip chainB

					// report a different onRamp address for chainC this leads to no consensus for chainC merkle roots
					baseObs.MerkleRoots[1].OnRampAddress = []byte{0xd}

					// report a different seqNum for chainE
					baseObs.MerkleRoots[2].SeqNumsRange = cciptypes.NewSeqNumRange(10, 16)

					return baseObs
				},
			},
			observers:         []commontypes.OracleID{1, 2, 3},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10,
			rmnEnabled:        true,
			expOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{
						ChainSel:      chainA,
						OnRampAddress: []byte{0xa},
						SeqNumsRange:  cciptypes.NewSeqNumRange(10, 15),
						MerkleRoot:    cciptypes.Bytes32{0xa},
					},
					{
						ChainSel:      chainF,
						OnRampAddress: []byte{0xb},
						SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
						MerkleRoot:    bytes32b,
					},
				},
				RMNEnabledChains: map[cciptypes.ChainSelector]bool{
					chainA: true,
					chainB: true,
					chainC: true,
					chainD: true,
					chainE: true,
					chainF: true,
				},
				RMNReportSignatures: []cciptypes.RMNECDSASignature{
					{R: bytes32a, S: bytes32b},
					{R: bytes32a, S: bytes32b},
				},
			},
			expErr: false,
		},
		{
			name: "waiting for merkleRoots report transmission",
			prevOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: chainA},
					{ChainSel: chainB},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
			},
			q: Query{},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 20},
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:                          []commontypes.OracleID{1, 2, 3},
			bigF:                               1,
			destChainSel:                       chainD,
			maxMerkleTreeSize:                  20,
			rmnEnabled:                         true,
			maxReportTransmissionCheckAttempts: 5,
			expOutcome: Outcome{
				OutcomeType: ReportInFlight,
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: chainA},
					{ChainSel: chainB},
				},
				ReportTransmissionCheckAttempts: 0x1,
			},
			expErr: false,
		},
		{
			name: "waiting for merkleRoots report transmission 2",
			prevOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: chainA},
					{ChainSel: chainB},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
			},
			q: Query{},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 21}, // <--- indicates report transmission (also for chainA)
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:                          []commontypes.OracleID{1, 2, 3},
			bigF:                               1,
			destChainSel:                       chainD,
			maxMerkleTreeSize:                  20,
			rmnEnabled:                         true,
			maxReportTransmissionCheckAttempts: 5,
			expOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
			expErr: false,
		},
		{
			// Same as the previous test but with multiple reports, this time it remains
			// inflight because only one of the seqNums was updated.
			name: "waiting for merkleRoots report transmission 3: multiple reports",
			prevOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: chainA},
					{ChainSel: chainB},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
			},
			q: Query{},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 21}, // <--- indicates report transmission (also for chainA)
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:                          []commontypes.OracleID{1, 2, 3},
			bigF:                               1,
			destChainSel:                       chainD,
			maxMerkleTreeSize:                  20,
			rmnEnabled:                         true,
			multipleReports:                    true,
			maxReportTransmissionCheckAttempts: 5,
			expOutcome: Outcome{
				OutcomeType: ReportInFlight,
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: chainA},
					{ChainSel: chainB},
				},
				ReportTransmissionCheckAttempts: 0x1,
			},
			expErr: false,
		},
		{
			name: "reached all report transmission check attempts",
			prevOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: chainA},
					{ChainSel: chainB},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: 10},
					{ChainSel: chainB, SeqNum: 20},
				},
				ReportTransmissionCheckAttempts: 10, // <------------
			},
			q: Query{},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						OffRampNextSeqNums: []plugintypes.SeqNumChain{
							{ChainSel: chainA, SeqNum: 10},
							{ChainSel: chainB, SeqNum: 20},
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 1,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:                          []commontypes.OracleID{1, 2, 3},
			bigF:                               1,
			destChainSel:                       chainD,
			maxMerkleTreeSize:                  20,
			rmnEnabled:                         true,
			maxReportTransmissionCheckAttempts: 10, // <-----------------
			expOutcome: Outcome{
				OutcomeType: ReportTransmissionFailed, // <----------- we don't want to retry checking
			},
			expErr: false,
		},
		{
			name: "if one signed root is missing then we make progress with the non-rmn roots",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			q: Query{
				RMNSignatures: &rmn.ReportSignatures{
					Signatures: []*rmnpb.EcdsaSignature{
						{R: bytes32a[:], S: bytes32b[:]},
						{R: bytes32a[:], S: bytes32b[:]},
					},
					LaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainA,
								OnrampAddress:       []byte{0xa},
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 10,
								MaxMsgNr: 15,
							},
							Root: bytes32a[:],
						},
					},
				},
			},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:      chainA,
								OnRampAddress: []byte{0xa},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainB,
								OnRampAddress: []byte{0xb},
								SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainC,
								OnRampAddress: []byte{0xc},
								SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainD,
								OnRampAddress: []byte{0xd},
								SeqNumsRange:  cciptypes.NewSeqNumRange(13, 23),
								MerkleRoot:    bytes32a,
							},
						},
						RMNEnabledChains: map[cciptypes.ChainSelector]bool{
							chainA: true,
							chainB: false,
							chainC: false,
							chainD: true,
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 0,
							chainC: 0,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:         []commontypes.OracleID{1, 2, 3},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10,
			rmnEnabled:        true,
			expOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{
						ChainSel:      chainB,
						OnRampAddress: []byte{0xb},
						SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
						MerkleRoot:    bytes32a,
					},
					{
						ChainSel:      chainC,
						OnRampAddress: []byte{0xc},
						SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
						MerkleRoot:    bytes32a,
					},
				},
				RMNEnabledChains: map[cciptypes.ChainSelector]bool{
					chainA: true,
					chainB: false,
					chainC: false,
					chainD: true,
				},
				RMNReportSignatures: []cciptypes.RMNECDSASignature{
					{R: bytes32a, S: bytes32b},
					{R: bytes32a, S: bytes32b},
				},
			},
			expErr: false,
		},
		{
			name: "if all signed roots are present we make progress with both rmn and non-rmn roots",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			q: Query{
				RMNSignatures: &rmn.ReportSignatures{
					Signatures: []*rmnpb.EcdsaSignature{
						{R: bytes32a[:], S: bytes32b[:]},
						{R: bytes32a[:], S: bytes32b[:]},
					},
					LaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainA,
								OnrampAddress:       []byte{0xa},
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 10,
								MaxMsgNr: 20,
							},
							Root: bytes32a[:],
						},
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainD,
								OnrampAddress:       []byte{0xd},
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 13,
								MaxMsgNr: 23,
							},
							Root: bytes32a[:],
						},
					},
				},
			},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:      chainA,
								OnRampAddress: []byte{0xa},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainB,
								OnRampAddress: []byte{0xb},
								SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainC,
								OnRampAddress: []byte{0xc},
								SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainD,
								OnRampAddress: []byte{0xd},
								SeqNumsRange:  cciptypes.NewSeqNumRange(13, 23),
								MerkleRoot:    bytes32a,
							},
						},
						RMNEnabledChains: map[cciptypes.ChainSelector]bool{
							chainA: true,
							chainB: false,
							chainC: false,
							chainD: true,
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 0,
							chainC: 0,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:         []commontypes.OracleID{1, 2, 3},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10,
			rmnEnabled:        true,
			expOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{
						ChainSel:      chainA,
						OnRampAddress: []byte{0xa},
						SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
						MerkleRoot:    bytes32a,
					},
					{
						ChainSel:      chainB,
						OnRampAddress: []byte{0xb},
						SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
						MerkleRoot:    bytes32a,
					},
					{
						ChainSel:      chainC,
						OnRampAddress: []byte{0xc},
						SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
						MerkleRoot:    bytes32a,
					},
					{
						ChainSel:      chainD,
						OnRampAddress: []byte{0xd},
						SeqNumsRange:  cciptypes.NewSeqNumRange(13, 23),
						MerkleRoot:    bytes32a,
					},
				},
				RMNEnabledChains: map[cciptypes.ChainSelector]bool{
					chainA: true,
					chainB: false,
					chainC: false,
					chainD: true,
				},
				RMNReportSignatures: []cciptypes.RMNECDSASignature{
					{R: bytes32a, S: bytes32b},
					{R: bytes32a, S: bytes32b},
				},
			},
			expErr: false,
		},
		{
			name: "if the leader sent one wrong root then we make progress only with the non-rmn roots",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			q: Query{
				RMNSignatures: &rmn.ReportSignatures{
					Signatures: []*rmnpb.EcdsaSignature{
						{R: bytes32a[:], S: bytes32b[:]},
						{R: bytes32a[:], S: bytes32b[:]},
					},
					LaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainA,
								OnrampAddress:       []byte{0xa},
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 10,
								MaxMsgNr: 20,
							},
							Root: bytes32a[:],
						},
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: chainD,
								OnrampAddress:       []byte("this is not an address  <------------------"),
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 13,
								MaxMsgNr: 23,
							},
							Root: bytes32a[:],
						},
					},
				},
			},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:      chainA,
								OnRampAddress: []byte{0xa},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainB,
								OnRampAddress: []byte{0xb},
								SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainC,
								OnRampAddress: []byte{0xc},
								SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainD,
								OnRampAddress: []byte{0xd},
								SeqNumsRange:  cciptypes.NewSeqNumRange(13, 23),
								MerkleRoot:    bytes32a,
							},
						},
						RMNEnabledChains: map[cciptypes.ChainSelector]bool{
							chainA: true,
							chainB: false,
							chainC: false,
							chainD: true,
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 0,
							chainC: 0,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:         []commontypes.OracleID{1, 2, 3},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10,
			rmnEnabled:        true,
			expOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{
						ChainSel:      chainB,
						OnRampAddress: []byte{0xb},
						SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
						MerkleRoot:    bytes32a,
					},
					{
						ChainSel:      chainC,
						OnRampAddress: []byte{0xc},
						SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
						MerkleRoot:    bytes32a,
					},
				},
				RMNEnabledChains: map[cciptypes.ChainSelector]bool{
					chainA: true,
					chainB: false,
					chainC: false,
					chainD: true,
				},
				RMNReportSignatures: []cciptypes.RMNECDSASignature{
					{R: bytes32a, S: bytes32b},
					{R: bytes32a, S: bytes32b},
				},
			},
			expErr: false,
		},
		{
			name: "if query is empty then progress should be made with rmn-disabled chains",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			q: Query{},
			observations: []func(testCase) Observation{
				func(tc testCase) Observation {
					return Observation{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:      chainA,
								OnRampAddress: []byte{0xa},
								SeqNumsRange:  cciptypes.NewSeqNumRange(10, 20),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainB,
								OnRampAddress: []byte{0xb},
								SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainC,
								OnRampAddress: []byte{0xc},
								SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
								MerkleRoot:    bytes32a,
							},
							{
								ChainSel:      chainD,
								OnRampAddress: []byte{0xd},
								SeqNumsRange:  cciptypes.NewSeqNumRange(13, 23),
								MerkleRoot:    bytes32a,
							},
						},
						RMNEnabledChains: map[cciptypes.ChainSelector]bool{
							chainA: true,
							chainB: false,
							chainC: false,
							chainD: true,
						},
						FChain: map[cciptypes.ChainSelector]int{
							chainA: 1,
							chainB: 0,
							chainC: 0,
							chainD: 1,
						},
					}
				},
				func(tc testCase) Observation { return tc.observations[0](tc) },
				func(tc testCase) Observation { return tc.observations[0](tc) },
			},
			observers:         []commontypes.OracleID{1, 2, 3},
			bigF:              1,
			destChainSel:      chainD,
			maxMerkleTreeSize: 10,
			rmnEnabled:        true,
			expOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []cciptypes.MerkleRootChain{
					{
						ChainSel:      chainB,
						OnRampAddress: []byte{0xb},
						SeqNumsRange:  cciptypes.NewSeqNumRange(11, 21),
						MerkleRoot:    bytes32a,
					},
					{
						ChainSel:      chainC,
						OnRampAddress: []byte{0xc},
						SeqNumsRange:  cciptypes.NewSeqNumRange(12, 22),
						MerkleRoot:    bytes32a,
					},
				},
				RMNEnabledChains: map[cciptypes.ChainSelector]bool{
					chainA: true,
					chainB: false,
					chainC: false,
					chainD: true,
				},
				RMNReportSignatures: []cciptypes.RMNECDSASignature{},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		require.Equal(t, len(tc.observations), len(tc.observers), "test case is wrong")
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			ctx := tests.Context(t)

			p := &Processor{
				lggr: lggr,
				reportingCfg: ocr3types.ReportingPluginConfig{
					F: tc.bigF,
				},
				destChain: tc.destChainSel,
				offchainCfg: pluginconfig.CommitOffchainConfig{
					MaxMerkleTreeSize:                  tc.maxMerkleTreeSize,
					RMNEnabled:                         tc.rmnEnabled,
					MaxReportTransmissionCheckAttempts: uint(tc.maxReportTransmissionCheckAttempts),
					MultipleReportsEnabled:             tc.multipleReports,
				},
				metricsReporter: NoopMetrics{},
				addressCodec:    internal.NewMockAddressCodecHex(t),
			}

			aos := make([]plugincommon.AttributedObservation[Observation], 0, len(tc.observations))
			for i, o := range tc.observations {
				aos = append(aos, plugincommon.AttributedObservation[Observation]{
					Observation: o(tc),
					OracleID:    tc.observers[i],
				})
			}

			outc, err := p.Outcome(ctx, tc.prevOutcome, tc.q, aos)
			if tc.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expOutcome, outc)
		})
	}
}

func Test_buildMerkleRootsOutcome(t *testing.T) {
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	t.Run("determinism check", func(t *testing.T) {
		const rounds = 50

		obs := consensusObservation{
			MerkleRoots: map[cciptypes.ChainSelector]cciptypes.MerkleRootChain{
				cciptypes.ChainSelector(1): {
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(10, 20),
					MerkleRoot:   cciptypes.Bytes32{1},
				},
				cciptypes.ChainSelector(2): {
					ChainSel:     2,
					SeqNumsRange: cciptypes.NewSeqNumRange(20, 30),
					MerkleRoot:   cciptypes.Bytes32{2},
				},
			},
			RMNRemoteConfig: map[cciptypes.ChainSelector]cciptypes.RemoteConfig{
				cciptypes.ChainSelector(1): rmnRemoteCfg,
			},
		}

		lggr := logger.Test(t)
		for i := 0; i < rounds; i++ {
			report1, err := buildMerkleRootsOutcome(rmn.ReportSignatures{}, false, lggr, obs, Outcome{}, mockAddrCodec)
			require.NoError(t, err)
			report2, err := buildMerkleRootsOutcome(rmn.ReportSignatures{}, false, lggr, obs, Outcome{}, mockAddrCodec)
			require.NoError(t, err)
			require.Equal(t, report1, report2)
		}
	})
}

func Test_reportRangesOutcome(t *testing.T) {
	lggr := logger.Test(t)

	destChain := cciptypes.ChainSelector(4)

	testCases := []struct {
		name                 string
		consensusObservation consensusObservation
		merkleTreeSizeLimit  uint64
		expectedOutcome      Outcome
	}{
		{
			name:            "base empty outcome",
			expectedOutcome: Outcome{OutcomeType: ReportEmpty},
		},
		{
			name: "simple scenario with one chain",
			consensusObservation: consensusObservation{
				OnRampMaxSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 20,
				},
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 18, // off ramp next is 18, on ramp max is 20 so new msgs are: [18, 19, 20]
				},
				RMNRemoteConfig: map[cciptypes.ChainSelector]cciptypes.RemoteConfig{
					destChain: rmnRemoteCfg,
				},
			},
			merkleTreeSizeLimit: 256, // default limit should be used
			expectedOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 1, SeqNumRange: cciptypes.NewSeqNumRange(18, 20)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1, SeqNum: 18},
				},
				RMNRemoteCfg: rmnRemoteCfg,
			},
		},
		{
			name: "simple scenario with one chain",
			consensusObservation: consensusObservation{
				OnRampMaxSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 20,
					2: 1000,
					3: 10000,
				},
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 18,  // off ramp next is 18, on ramp max is 20 so new msgs are: [18, 19, 20]
					2: 995, // off ramp next is 995, on ramp max is 1000 so new msgs are: [995, 996, 997, 998, 999, 1000]
					3: 500, // off ramp next is 500, we have new messages up to 10000 (default limit applied)
				},
				RMNRemoteConfig: map[cciptypes.ChainSelector]cciptypes.RemoteConfig{
					destChain: rmnRemoteCfg,
				},
			},
			merkleTreeSizeLimit: 5,
			expectedOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 1, SeqNumRange: cciptypes.NewSeqNumRange(18, 20)},
					{ChainSel: 2, SeqNumRange: cciptypes.NewSeqNumRange(995, 999)},
					{ChainSel: 3, SeqNumRange: cciptypes.NewSeqNumRange(500, 504)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1, SeqNum: 18},
					{ChainSel: 2, SeqNum: 995},
					{ChainSel: 3, SeqNum: 500},
				},
				RMNRemoteCfg: rmnRemoteCfg,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outcome := reportRangesOutcome(Query{}, lggr, tc.consensusObservation, tc.merkleTreeSizeLimit, destChain)
			require.Equal(t, tc.expectedOutcome, outcome)
		})
	}
}

func TestCheckForReportTransmission(t *testing.T) {
	source1, source2 := cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)
	tests := []struct {
		name                          string
		maxReportTransmissionAttempts uint
		multipleReports               bool
		previousOutcome               Outcome
		consensusObservation          consensusObservation
		expectedOutcome               Outcome
	}{
		{
			name:                          "Report transmitted, multiReports disabled",
			maxReportTransmissionAttempts: 3,
			multipleReports:               false,
			previousOutcome: Outcome{
				ReportTransmissionCheckAttempts: 1,
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 10)},
					{ChainSel: source2, SeqNumsRange: cciptypes.NewSeqNumRange(20, 20)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 10},
					{ChainSel: source2, SeqNum: 20},
				},
			},
			consensusObservation: consensusObservation{
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					source1: 11,
					source2: 21,
				},
			},
			expectedOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
		},
		{
			name:                          "Report transmitted, multiReports enabled",
			maxReportTransmissionAttempts: 3,
			multipleReports:               true,
			previousOutcome: Outcome{
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 10)},
					{ChainSel: source2, SeqNumsRange: cciptypes.NewSeqNumRange(20, 20)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 10},
					{ChainSel: source2, SeqNum: 20},
				},
			},
			consensusObservation: consensusObservation{
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					source1: 11,
					source2: 21,
				},
			},
			expectedOutcome: Outcome{
				OutcomeType: ReportTransmitted,
			},
		},
		{
			name:                          "Report in flight",
			maxReportTransmissionAttempts: 3,
			multipleReports:               true,
			previousOutcome: Outcome{
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 10)},
					{ChainSel: source2, SeqNumsRange: cciptypes.NewSeqNumRange(20, 20)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 10},
					{ChainSel: source2, SeqNum: 20},
				},
				ReportTransmissionCheckAttempts: 1,
			},
			consensusObservation: consensusObservation{
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					source1: 10,
					source2: 20,
				},
			},
			expectedOutcome: Outcome{
				OutcomeType: ReportInFlight,
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 10},
					{ChainSel: source2, SeqNum: 20},
				},
				ReportTransmissionCheckAttempts: 2, // previous + 1
				// same as before, still in-flight.
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 10)},
					{ChainSel: source2, SeqNumsRange: cciptypes.NewSeqNumRange(20, 20)},
				},
			},
		},
		{
			name:                          "Report transmission failed",
			maxReportTransmissionAttempts: 2,
			multipleReports:               true,
			previousOutcome: Outcome{
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 10)},
					{ChainSel: source2, SeqNumsRange: cciptypes.NewSeqNumRange(20, 20)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 10},
					{ChainSel: source2, SeqNum: 20},
				},
				ReportTransmissionCheckAttempts: 2,
			},
			consensusObservation: consensusObservation{
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					source1: 10,
					source2: 20,
				},
			},
			expectedOutcome: Outcome{
				OutcomeType: ReportTransmissionFailed,
			},
		},
		{
			name:                          "Impossible state",
			maxReportTransmissionAttempts: 3,
			multipleReports:               false,
			previousOutcome: Outcome{
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(15, 15)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 15},
				},
			},
			consensusObservation: consensusObservation{
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					source1: 10,
				},
			},
			expectedOutcome: Outcome{
				OutcomeType: ReportInFlight,
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: source1, SeqNum: 15},
				},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: source1, SeqNumsRange: cciptypes.NewSeqNumRange(15, 15)},
				},
				ReportTransmissionCheckAttempts: 1,
			}, // The function logs an error but continues execution.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lggr := logger.Test(t)
			outcome := checkForReportTransmission(
				lggr,
				tt.maxReportTransmissionAttempts,
				tt.multipleReports,
				tt.previousOutcome,
				tt.consensusObservation,
			)
			require.Equal(t, tt.expectedOutcome, outcome)
		})
	}
}
