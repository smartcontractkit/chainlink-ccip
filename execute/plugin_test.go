package execute

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/execute/internal/cache"
	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/observer"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	codec_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/ccipocr3"
	plugincommon_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func genRandomChainReports(numReports, numMsgsPerReport int) []cciptypes.ExecutePluginReportSingleChain {
	// generate random chain reports with unique (random) source
	// chain selectors and unique (random) sequence number ranges.
	// don't need to set anything else.
	chainReports := make([]cciptypes.ExecutePluginReportSingleChain, numReports)
	for i := 0; i < numReports; i++ {
		scc := cciptypes.ChainSelector(rand.RandomUint64())
		chainReports[i] = cciptypes.ExecutePluginReportSingleChain{
			SourceChainSelector: scc,
			Messages:            make([]cciptypes.Message, 0, numMsgsPerReport),
		}
		start := rand.RandomUint32()
		for j := 0; j < numMsgsPerReport; j++ {
			chainReports[i].Messages = append(chainReports[i].Messages, cciptypes.Message{
				Header: cciptypes.RampMessageHeader{
					SequenceNumber:      cciptypes.SeqNum(start + uint32(j)),
					SourceChainSelector: scc,
				},
			})
		}
	}
	return chainReports
}

func Test_getMinMaxSeqNrRangesBySource_emptyMsgs(t *testing.T) {
	chainReports := genRandomChainReports(1, 0)
	minMaxSeqNrRanges := getSeqNrRangesBySource(chainReports)
	require.Len(t, minMaxSeqNrRanges, 0)
}

func Test_getSeqNrRangesBySource(t *testing.T) {
	chainReports := genRandomChainReports(10, 15)
	minMaxSeqNrRanges := getSeqNrRangesBySource(chainReports)
	require.Len(t, minMaxSeqNrRanges, len(chainReports))
	for _, chainReport := range chainReports {
		seqNrRanges, ok := minMaxSeqNrRanges[chainReport.SourceChainSelector]
		require.True(t, ok)
		require.NotNil(t, seqNrRanges)
		require.Equal(t, 1, len(seqNrRanges))

		seqNrRange := seqNrRanges[0]
		// check the range.
		expectedMin := chainReport.Messages[0].Header.SequenceNumber
		expectedMax := chainReport.Messages[len(chainReport.Messages)-1].Header.SequenceNumber
		require.Equal(t, expectedMin, seqNrRange.Start())
		require.Equal(t, expectedMax, seqNrRange.End())

		// check the set.
		expectedSet := make(map[cciptypes.SeqNum]struct{})
		for _, msg := range chainReport.Messages {
			expectedSet[msg.Header.SequenceNumber] = struct{}{}
		}
		expectedSetSlice := maps.Keys(expectedSet)
		actualSetSlice := seqNrRange.ToSlice()
		sort.Slice(expectedSetSlice, func(i, j int) bool {
			return expectedSetSlice[i] < expectedSetSlice[j]
		})
		sort.Slice(actualSetSlice, func(i, j int) bool {
			return actualSetSlice[i] < actualSetSlice[j]
		})
		require.Equal(t, expectedSetSlice, actualSetSlice)
	}
}

func Test_checkAlreadyExecuted(t *testing.T) {
	testCases := []struct {
		name           string
		mockReaderFunc func(
			t *testing.T,
			snRangeSetPairBySource map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
		) *readerpkg_mock.MockCCIPReader
		shouldErr bool
	}{
		{
			name: "full range executed, rest unexecuted, should not error",
			mockReaderFunc: func(
				t *testing.T,
				snRangeSetPairBySource map[cciptypes.ChainSelector][]cciptypes.SeqNumRange) *readerpkg_mock.MockCCIPReader {
				ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)

				// map key order is undefined, but we can select any random one
				sourceSel := maps.Keys(snRangeSetPairBySource)[0]
				ccipReaderMock.
					EXPECT().
					ExecutedMessages(
						mock.Anything,
						snRangeSetPairBySource,
						primitives.Unconfirmed,
					).Return(
					map[cciptypes.ChainSelector][]cciptypes.SeqNum{
						sourceSel: snRangeSetPairBySource[sourceSel][0].ToSlice(),
					}, nil,
				).Maybe()
				return ccipReaderMock
			},
			shouldErr: false,
		},
		{
			name: "subset of range executed, rest unexecuted, should not error",
			mockReaderFunc: func(
				t *testing.T,
				snRangeSetPairBySource map[cciptypes.ChainSelector][]cciptypes.SeqNumRange) *readerpkg_mock.MockCCIPReader {
				ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)

				// map key order is undefined, but we can select any random one
				sourceSel := maps.Keys(snRangeSetPairBySource)[0]
				fullRange := snRangeSetPairBySource[sourceSel][0].ToSlice()
				ccipReaderMock.
					EXPECT().
					ExecutedMessages(
						mock.Anything,
						snRangeSetPairBySource,
						primitives.Unconfirmed,
					).Return(
					map[cciptypes.ChainSelector][]cciptypes.SeqNum{sourceSel: fullRange[:len(fullRange)/2]},
					nil,
				).Maybe()
				return ccipReaderMock
			},
			shouldErr: false,
		},
		{
			name: "none executed, should not error",
			mockReaderFunc: func(
				t *testing.T,
				snRangeSetPairBySource map[cciptypes.ChainSelector][]cciptypes.SeqNumRange) *readerpkg_mock.MockCCIPReader {
				ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
				ccipReaderMock.
					EXPECT().
					ExecutedMessages(
						mock.Anything,
						snRangeSetPairBySource,
						primitives.Unconfirmed,
					).Return(nil, nil).Maybe()
				return ccipReaderMock
			},
			shouldErr: false,
		},
		{
			name: "all executed, should error",
			mockReaderFunc: func(
				t *testing.T,
				snRangeSetPairBySource map[cciptypes.ChainSelector][]cciptypes.SeqNumRange) *readerpkg_mock.MockCCIPReader {
				ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)

				var allExecuted = make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
				for source, seqNrRanges := range snRangeSetPairBySource {
					if _, exists := allExecuted[source]; !exists {
						allExecuted[source] = make([]cciptypes.SeqNum, 0)
					}

					for _, seqNrRange := range seqNrRanges {
						allExecuted[source] = append(allExecuted[source], seqNrRange.ToSlice()...)
					}
				}
				ccipReaderMock.
					EXPECT().
					ExecutedMessages(
						mock.Anything,
						snRangeSetPairBySource,
						primitives.Unconfirmed,
					).Return(
					allExecuted,
					nil,
				)
				return ccipReaderMock
			},
			shouldErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chainReports := genRandomChainReports(10, 15)
			snRangeSetPairs := getSeqNrRangesBySource(chainReports)
			ccipReaderMock := tc.mockReaderFunc(t, snRangeSetPairs)
			p := &Plugin{
				lggr:       logger.Test(t),
				ccipReader: ccipReaderMock,
			}
			err := p.checkAlreadyExecuted(t.Context(), p.lggr, chainReports)
			if tc.shouldErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "already executed")
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_getPendingReportsForExecution(t *testing.T) {
	canExecute := func(ret bool) CanExecuteHandle {
		return func(cciptypes.ChainSelector, cciptypes.Bytes32) bool { return ret }
	}

	tcs := []struct {
		name                    string
		reports                 []cciptypes.CommitPluginReportWithMeta
		ranges                  map[cciptypes.ChainSelector][]cciptypes.SeqNum // finalized
		unfinalizedRanges       map[cciptypes.ChainSelector][]cciptypes.SeqNum // unfinalized
		cursedSourceChains      map[cciptypes.ChainSelector]bool               // cursed chains
		canExec                 CanExecuteHandle
		fetchFrom               time.Time
		wantObs                 exectypes.CommitObservations
		wantExecutedFinalized   []exectypes.CommitData
		wantExecutedUnfinalized []exectypes.CommitData
		wantErr                 assert.ErrorAssertionFunc
	}{
		{
			name:                    "empty",
			reports:                 nil,
			ranges:                  nil,
			unfinalizedRanges:       nil,
			cursedSourceChains:      nil,
			fetchFrom:               time.Time{},
			wantObs:                 exectypes.CommitObservations{},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "single non-executed report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: nil,
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: nil,
			},
			cursedSourceChains: nil,
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "single partially-executed report with finalized and unfinalized",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {1, 2, 3},
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {7, 8},
			},
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 7, 8},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "single fully-executed finalized report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(1, 10).ToSlice(),
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: nil,
			},
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: nil,
			},
			wantExecutedFinalized: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
					Timestamp:           time.UnixMilli(10101010101),
					BlockNum:            999,
				},
			},
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "single fully-executed unfinalized report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: nil,
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(1, 10).ToSlice(),
			},
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: nil,
			},
			wantExecutedFinalized: nil,
			wantExecutedUnfinalized: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
					Timestamp:           time.UnixMilli(10101010101),
					BlockNum:            999,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "multiple reports with both finalized and unfinalized fully executed",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
								MerkleRoot:   cciptypes.Bytes32{1}, // Unique root
							},
						},
					},
				},
				{
					BlockNum:  1000,
					Timestamp: time.UnixMilli(10101010102),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(11, 20),
								MerkleRoot:   cciptypes.Bytes32{2}, // Unique root
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(1, 10).ToSlice(),
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(11, 20).ToSlice(),
			},
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: nil,
			},
			wantExecutedFinalized: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
					Timestamp:           time.UnixMilli(10101010101),
					BlockNum:            999,
					MerkleRoot:          cciptypes.Bytes32{1}, // Match unique root
				},
			},
			wantExecutedUnfinalized: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20),
					Timestamp:           time.UnixMilli(10101010102),
					BlockNum:            1000,
					MerkleRoot:          cciptypes.Bytes32{2}, // Match unique root
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "single known-executed report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			canExec:            canExecute(false),
			ranges:             nil,
			unfinalizedRanges:  nil,
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: nil,
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "mixed execution states across three reports",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  1000,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
								MerkleRoot:   cciptypes.Bytes32{1}, // Unique root
							},
						},
					},
				},
				{
					BlockNum:  1001,
					Timestamp: time.UnixMilli(10101010102),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(11, 20),
								MerkleRoot:   cciptypes.Bytes32{2}, // Unique root
							},
						},
					},
				},
				{
					BlockNum:  1002,
					Timestamp: time.UnixMilli(10101010103),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(21, 30),
								MerkleRoot:   cciptypes.Bytes32{3}, // Unique root
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(11, 20).ToSlice(), // Second report fully finalized
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {1, 2, 3, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
				// First report partially executed, third report fully executed unfinalized
			},
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            1000,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
						MerkleRoot:          cciptypes.Bytes32{1}, // Match unique root
					},
				},
			},
			wantExecutedFinalized: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20),
					Timestamp:           time.UnixMilli(10101010102),
					BlockNum:            1001,
					MerkleRoot:          cciptypes.Bytes32{2}, // Match unique root
				},
			},
			wantExecutedUnfinalized: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(21, 30),
					Timestamp:           time.UnixMilli(10101010103),
					BlockNum:            1002,
					MerkleRoot:          cciptypes.Bytes32{3}, // Match unique root
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "single partially-executed report with finalized and unfinalized",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {1, 2, 3},
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {7, 8},
			},
			cursedSourceChains: nil, // No cursed chains
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            999,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 7, 8},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "single cursed chain report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges:            nil,
			unfinalizedRanges: nil,
			cursedSourceChains: map[cciptypes.ChainSelector]bool{
				1: true, // Chain 1 is cursed
			},
			fetchFrom:               time.Time{},
			wantObs:                 exectypes.CommitObservations{}, // Empty observations since all chains are cursed
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "multiple chains with one cursed",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  1000,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1, // Cursed chain
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
							{
								ChainSel:     2, // Non-cursed chain
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {1, 2},
				2: {1, 2},
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {3, 4},
				2: {3, 4},
			},
			cursedSourceChains: map[cciptypes.ChainSelector]bool{
				1: true,  // Chain 1 is cursed
				2: false, // Chain 2 is not cursed
			},
			fetchFrom: time.Time{},
			wantObs: exectypes.CommitObservations{
				2: []exectypes.CommitData{
					{
						SourceChain:         2,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            1000,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 4},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "mixed blessed and unblessed roots with cursed chains",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  1000,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1, // Non-cursed chain
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 5),
							},
						},
						UnblessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     2, // Cursed chain
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 5),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {1, 2},
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {3, 4},
			},
			cursedSourceChains: map[cciptypes.ChainSelector]bool{
				1: false, // Chain 1 is not cursed
				2: true,  // Chain 2 is cursed
			},
			fetchFrom: time.Time{},
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 5),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            1000,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 4},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "nil cursed source chains map",
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  1000,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {1, 2, 3},
			},
			unfinalizedRanges: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {4, 5},
			},
			cursedSourceChains: nil, // No curse information available
			fetchFrom:          time.Time{},
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010101),
						BlockNum:            1000,
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3, 4, 5},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{
			name: "filters_price_reports_and_processes_merkle_report",
			reports: []cciptypes.CommitPluginReportWithMeta{
				// Price report 1 (no Merkle roots)
				{
					BlockNum:  1000,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: nil, // Indicates a price-only report or similar non-executable
					},
				},
				// Price report 2 (no Merkle roots)
				{
					BlockNum:  1001,
					Timestamp: time.UnixMilli(10101010102),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{}, // Also indicates non-executable for messages
					},
				},
				// Report with an actual Merkle root
				{
					BlockNum:  1002,
					Timestamp: time.UnixMilli(10101010103),
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1, // Source chain selector
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
								MerkleRoot:   cciptypes.Bytes32{0x01}, // Example Merkle root
							},
						},
					},
				},
			},
			ranges:             nil, // No finalized messages for this new report
			unfinalizedRanges:  nil, // No unfinalized messages for this new report
			cursedSourceChains: nil,
			fetchFrom:          time.Time{},
			canExec:            canExecute(true), // Assume the Merkle root is executable
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{ // Expecting only the report with the Merkle root to be processed
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010103),
						BlockNum:            1002,
						MerkleRoot:          cciptypes.Bytes32{0x01},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
		{ // New test case for progress with low limit due to cache advancing
			name: "progress_with_low_limit_due_to_cache_advancing",
			// reports that the mockReader.CommitReportsGTETimestamp should return in this specific call
			reports: []cciptypes.CommitPluginReportWithMeta{
				{
					BlockNum:  1002,
					Timestamp: time.UnixMilli(10101010103), // Timestamp of the executable report
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
								MerkleRoot:   cciptypes.Bytes32{0x01},
							},
						},
					},
				},
			},
			ranges:             nil,
			unfinalizedRanges:  nil,
			cursedSourceChains: nil,
			fetchFrom:          time.UnixMilli(10101010000), // Broader visibility window
			canExec:            canExecute(true),
			wantObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.UnixMilli(10101010103),
						BlockNum:            1002,
						MerkleRoot:          cciptypes.Bytes32{0x01},
					},
				},
			},
			wantExecutedFinalized:   nil,
			wantExecutedUnfinalized: nil,
			wantErr:                 assert.NoError,
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			currentTestLogger := logger.Test(t)
			if tt.canExec == nil {
				tt.canExec = canExecute(true)
			}

			mockReader := readerpkg_mock.NewMockCCIPReader(t)
			var mockCache cache.CommitReportCache = &noopCommitReportCache{}

			offchainConfigForTest := pluginconfig.ExecuteOffchainConfig{MaxCommitReportsToFetch: 10} // Default for most tests

			if tt.name == "progress_with_low_limit_due_to_cache_advancing" {
				configurableCache := &configurableCommitReportCache{
					reportsToQueryFrom:    time.UnixMilli(10101010102), // Advanced timestamp
					cachedReportsToReturn: []cciptypes.CommitPluginReportWithMeta{},
					lggr:                  currentTestLogger,
				}
				// For this specific test, DeduplicateReports should directly return what CommitReportsGTETimestamp returns
				// as we assume cachedReportsToReturn is empty and the new reports are tt.reports.
				configurableCache.deduplicateShouldReturn = tt.reports
				mockCache = configurableCache
				offchainConfigForTest.MaxCommitReportsToFetch = 1 // Crucial for this test case

				mockReader.EXPECT().CommitReportsGTETimestamp(
					mock.Anything,               // context
					time.UnixMilli(10101010102), // Expecting the advanced timestamp from cache
					primitives.Unconfirmed,
					int(1), // Expecting the low fetch limit
				).Return(tt.reports, nil)
			} else {
				// Default mock setup for other tests
				mockReader.EXPECT().CommitReportsGTETimestamp(
					mock.Anything,
					mock.Anything,
					primitives.Unconfirmed,
					mock.Anything, // Most tests use a default MaxCommitReportsToFetch or don't care about this specific arg
				).Return(tt.reports, nil)
			}

			executed := make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
			for k, v := range tt.ranges {
				if _, exists := executed[k]; !exists {
					executed[k] = make([]cciptypes.SeqNum, 0)
				}
				executed[k] = append(executed[k], v...)
			}
			mockReader.EXPECT().ExecutedMessages(mock.Anything, mock.Anything, primitives.Finalized).Return(executed, nil)

			unfinalized := make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
			for k, v := range tt.unfinalizedRanges {
				if _, exists := unfinalized[k]; !exists {
					unfinalized[k] = make([]cciptypes.SeqNum, 0)
				}
				unfinalized[k] = append(unfinalized[k], v...)
			}
			mockReader.EXPECT().ExecutedMessages(mock.Anything, mock.Anything, primitives.Unconfirmed).Return(unfinalized, nil)

			got, gotFinalized, gotUnfinalized, err := getPendingReportsForExecution(
				t.Context(),
				mockReader,
				mockCache,
				tt.canExec,
				tt.fetchFrom,
				tt.cursedSourceChains,
				int(offchainConfigForTest.MaxCommitReportsToFetch), // limit int
				currentTestLogger,
			)
			if !tt.wantErr(t, err, "getPendingReportsForExecution(...)") {
				return
			}
			assert.Equalf(t, tt.wantObs, got, "getPendingReportsForExecution(...)")
			assert.Equalf(t, tt.wantExecutedFinalized, gotFinalized, "getPendingReportsForExecution(...)")
			assert.Equalf(t, tt.wantExecutedUnfinalized, gotUnfinalized, "getPendingReportsForExecution(...)")
		})
	}
}

func TestPlugin_Close(t *testing.T) {
	lggr := logger.Test(t)
	mockReader := readerpkg_mock.NewMockCCIPReader(t)

	// Set up expectation for Close() method
	mockReader.On("Close").Return(nil)
	p := &Plugin{
		lggr:              lggr,
		tokenDataObserver: &observer.NoopTokenDataObserver{},
		ccipReader:        mockReader}
	require.NoError(t, p.Close())
	mockReader.AssertExpectations(t)
}

func TestPlugin_Query(t *testing.T) {
	p := &Plugin{}
	q, err := p.Query(t.Context(), ocr3types.OutcomeContext{})
	require.NoError(t, err)
	require.Equal(t, types.Query{}, q)
}

func TestPlugin_ObservationQuorum(t *testing.T) {
	ctx := t.Context()
	p := &Plugin{
		reportingCfg: ocr3types.ReportingPluginConfig{F: 1},
	}
	got, err := p.ObservationQuorum(ctx, ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{Observation: []byte{}},
		{Observation: []byte{}},
		{Observation: []byte{}},
	})
	require.NoError(t, err)
	assert.Equal(t, true, got)
}

func TestPlugin_ValidateObservation_NonDecodable(t *testing.T) {
	ctx := t.Context()
	p := &Plugin{ocrTypeCodec: ocrTypeCodec}
	err := p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte("not a valid observation"),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observation")
}

func TestPlugin_ValidateObservation_SupportedChainsError(t *testing.T) {
	ctx := t.Context()
	p := &Plugin{
		ocrTypeCodec: ocrTypeCodec,
	}
	err := p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte{},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error finding supported chains by node: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_ValidateObservation_IneligibleMessageObserver(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet[cciptypes.ChainSelector](), nil)
	defer mockHomeChain.AssertExpectations(t)

	p := &Plugin{
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		ocrTypeCodec: ocrTypeCodec,
		lggr:         lggr,
	}

	observation := exectypes.NewObservation(nil, exectypes.MessageObservations{
		0: map[cciptypes.SeqNum]cciptypes.Message{
			1: {
				Header: cciptypes.RampMessageHeader{
					SourceChainSelector: 1,
				},
			},
		},
	}, nil, nil, dt.Observation{}, nil)
	encoded, err := ocrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)

	prevOutcome := exectypes.Outcome{
		State: exectypes.GetCommitReports,
	}
	encodedPrevOutcome, err := ocrTypeCodec.EncodeOutcome(prevOutcome)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{PreviousOutcome: encodedPrevOutcome}, types.Query{},
		types.AttributedObservation{
			Observation: encoded,
		})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observer reading eligibility: observer not allowed to read from chain 0")
}

func TestPlugin_ValidateObservation_IneligibleCommitReportsObserver(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet[cciptypes.ChainSelector](), nil)
	defer mockHomeChain.AssertExpectations(t)

	p := &Plugin{
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		lggr:         lggr,
		ocrTypeCodec: ocrTypeCodec,
	}

	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{
				MerkleRoot:          cciptypes.Bytes32{123},
				SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2),
				SourceChain:         1,
			},
		},
	}
	observation := exectypes.NewObservation(
		commitReports, nil, nil, nil, dt.Observation{}, nil,
	)
	encoded, err := ocrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t,
		err.Error(),
		"destination chain not supported but observed commit report")
}

func TestPlugin_ValidateObservation_ValidateObservedSeqNum_Error(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		ocrTypeCodec: ocrTypeCodec,
	}

	// Reports with duplicate roots.
	root := cciptypes.Bytes32{}
	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{MerkleRoot: root, SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2), SourceChain: 1},
			{MerkleRoot: root, SequenceNumberRange: cciptypes.NewSeqNumRange(1, 5), SourceChain: 1},
		},
	}
	observation := exectypes.NewObservation(
		commitReports, nil, nil, nil, dt.Observation{}, nil,
	)
	encoded, err := ocrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observed sequence numbers: duplicate merkle root")
}

func TestPlugin_ValidateObservation_CallsDiscoveryValidateObservation(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)
	mockDiscoveryProcessor := plugincommon_mock.NewMockPluginProcessor[dt.Query, dt.Observation, dt.Outcome](t)
	mockDiscoveryProcessor.EXPECT().ValidateObservation(dt.Outcome{}, dt.Query{}, mock.Anything).Return(nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		discovery:    mockDiscoveryProcessor,
		ocrTypeCodec: ocrTypeCodec,
	}

	// Reports with duplicate roots.
	root := cciptypes.Bytes32{}
	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{MerkleRoot: root, SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2), SourceChain: 1},
		},
	}
	observation := exectypes.NewObservation(
		commitReports, nil, nil, nil, dt.Observation{}, nil,
	)
	encoded, err := ocrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.NoError(t, err)
}

func TestPlugin_Observation_BadPreviousOutcome(t *testing.T) {
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}
	_, err := p.Observation(t.Context(), ocr3types.OutcomeContext{
		PreviousOutcome: []byte("not a valid observation"),
	}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode previous outcome: proto unmarshal ExecOutcome")
}

func TestPlugin_Observation_EligibilityCheckFailure(t *testing.T) {
	lggr := logger.Test(t)

	// Mock all necessary dependencies
	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetFChain().Return(map[cciptypes.ChainSelector]int{}, nil)

	// Create a mock CCIP reader with a relaxed approach - any calls will return a non-error response
	mockCCIPReader := readerpkg_mock.NewMockCCIPReader(t)

	// Use .Maybe() to make expectations optional, so the test is more robust
	mockCCIPReader.EXPECT().
		CommitReportsGTETimestamp(
			mock.Anything,
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return([]cciptypes.CommitPluginReportWithMeta{}, nil).Maybe()

	mockCCIPReader.EXPECT().
		GetRmnCurseInfo(mock.Anything).
		Return(cciptypes.CurseInfo{}, nil).Maybe()

	// Create a simplified plugin structure that will test the eligibility failure
	// This removes the dependency on actual cache implementations
	p := &Plugin{
		homeChain:       mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{}, // Empty map to trigger the eligibility check error
		lggr:            lggr,
		ocrTypeCodec:    ocrTypeCodec,
		ccipReader:      mockCCIPReader,
		commitRootsCache: cache.NewCommitRootsCache(
			lggr,
			8*time.Hour,
			5*time.Minute,
		),
		commitReportCache: &noopCommitReportCache{},
	}

	// Run the test - this should fail with the eligibility check error
	_, err := p.Observation(t.Context(), ocr3types.OutcomeContext{}, nil)
	require.Error(t, err)
	assert.Contains(t,
		err.Error(),
		"unable to determine if the destination chain is supported: "+
			"error getting supported chains: oracle ID 0 not found in oracleIDToP2pID")
}

// noopCommitReportCache is a simple implementation of CommitReportCache that does nothing
type noopCommitReportCache struct{}

func (c *noopCommitReportCache) RefreshCache(ctx context.Context) error {
	return nil
}

func (c *noopCommitReportCache) GetReportsToQueryFromTimestamp() time.Time {
	return time.Now().Add(-8 * time.Hour)
}

func (c *noopCommitReportCache) GetCachedReports(fromTimestamp time.Time) []cciptypes.CommitPluginReportWithMeta {
	return nil
}

func (c *noopCommitReportCache) DeduplicateReports(
	reports []cciptypes.CommitPluginReportWithMeta,
) []cciptypes.CommitPluginReportWithMeta {
	return reports // Simply return the reports without any deduplication
}

func TestPlugin_Outcome_DestFChainNotAvailable(t *testing.T) {
	ctx := t.Context()
	fChainMap := map[cciptypes.ChainSelector]int{
		1: 1,
		2: 2,
	}

	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}
	observation, err := ocrTypeCodec.EncodeObservation(exectypes.Observation{
		Contracts: dt.Observation{}, FChain: fChainMap,
	})
	require.NoError(t, err)
	_, err = p.Outcome(ctx, ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "destination chain 0 is not in FChain")
}

func TestPlugin_Outcome_BadObservationEncoding(t *testing.T) {
	ctx := t.Context()
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}
	_, err := p.Outcome(ctx, ocr3types.OutcomeContext{}, nil,
		[]types.AttributedObservation{
			{
				Observation: []byte("not a valid observation"),
				Observer:    0,
			},
		})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observations: proto unmarshal ExecObservation")
}

func TestPlugin_Outcome_BelowF(t *testing.T) {
	ctx := t.Context()
	fChainMap := map[cciptypes.ChainSelector]int{
		0: 1,
		2: 2,
	}
	p := &Plugin{
		reportingCfg: ocr3types.ReportingPluginConfig{
			F: 1,
		},
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}
	observation, err := ocrTypeCodec.EncodeObservation(exectypes.Observation{
		Contracts: dt.Observation{}, FChain: fChainMap,
	})
	require.NoError(t, err)
	_, err = p.Outcome(ctx, ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.Error(t, err)
	// because fChain observations doesn't reach consensus with the low number of observations
	assert.Contains(t, err.Error(), "destination chain 0 is not in FChain")
}

func TestPlugin_Outcome_CommitReportsMergeMissingValidator_Skips(t *testing.T) {
	ctx := t.Context()
	fChainMap := map[cciptypes.ChainSelector]int{
		10: 20,
		0:  3,
	}

	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}

	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {},
	}
	observation, err := ocrTypeCodec.EncodeObservation(exectypes.Observation{
		CommitReports: commitReports,
		Contracts:     dt.Observation{},
		FChain:        fChainMap,
	})
	require.NoError(t, err)
	outcome, err := p.Outcome(ctx, ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.NoError(t, err)
	require.Len(t, outcome, 0)
}

func TestPlugin_Outcome_MessagesMergeError(t *testing.T) {
	ctx := t.Context()
	fChainMap := map[cciptypes.ChainSelector]int{
		0:  3,
		10: 20,
	}

	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}

	// map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message
	messages := map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
		1: {
			1: {
				Header: cciptypes.RampMessageHeader{
					SourceChainSelector: 1,
				},
			},
		},
	}
	observation, err := ocrTypeCodec.EncodeObservation(exectypes.Observation{
		Messages: messages, Contracts: dt.Observation{}, FChain: fChainMap,
	})
	require.NoError(t, err)
	outcome, err := p.Outcome(ctx, ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.NoError(t, err)
	require.Len(t, outcome, 0)
}

func TestPlugin_Reports_MultipleReports(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	mockReportCodec := codec_mocks.NewMockExecutePluginCodec(t)
	mockReportCodec.On("Encode", mock.Anything, mock.MatchedBy(func(report cciptypes.ExecutePluginReport) bool {
		return len(report.ChainReports) > 0
	})).Return([]byte("encoded_report_1"), nil).Once()
	mockReportCodec.On("Encode", mock.Anything, mock.MatchedBy(func(report cciptypes.ExecutePluginReport) bool {
		return len(report.ChainReports) > 0
	})).Return([]byte("encoded_report_2"), nil).Once()

	// Mock chain support for transmission schedule
	mockChainSupport := plugincommon_mock.NewMockChainSupport(t)
	mockChainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil)

	// Create plugin instance with necessary mocks
	p := &Plugin{
		lggr:         lggr,
		reportCodec:  mockReportCodec,
		ocrTypeCodec: ocrTypeCodec,
		chainSupport: mockChainSupport,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			1: {},
		},
	}

	// Create test outcome with two reports
	commitReports := []exectypes.CommitData{
		{
			SourceChain: 100,
			MerkleRoot:  cciptypes.Bytes32{1, 2, 3},
		},
		{
			SourceChain: 200,
			MerkleRoot:  cciptypes.Bytes32{4, 5, 6},
		},
	}

	reports := []cciptypes.ExecutePluginReport{
		{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{
					SourceChainSelector: 100,
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{SequenceNumber: 1}},
					},
				},
			},
		},
		{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{
					SourceChainSelector: 200,
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
					},
				},
			},
		},
	}

	outcome := exectypes.NewOutcome(exectypes.Filter, commitReports, reports)
	encodedOutcome, err := ocrTypeCodec.EncodeOutcome(outcome)
	require.NoError(t, err)

	reportPlus, err := p.Reports(ctx, 1, encodedOutcome)

	require.NoError(t, err)
	require.Len(t, reportPlus, 2, "Should return 2 reports")

	// First report
	assert.Equal(t, types.Report("encoded_report_1"), reportPlus[0].ReportWithInfo.Report)
	assert.NotEmpty(t, reportPlus[0].ReportWithInfo.Info, "Report info should not be empty")
	assert.NotEmpty(t, reportPlus[0].TransmissionScheduleOverride, "Transmission schedule should not be empty")

	// Second report
	assert.Equal(t, types.Report("encoded_report_2"), reportPlus[1].ReportWithInfo.Report)
	assert.NotEmpty(t, reportPlus[1].ReportWithInfo.Info, "Report info should not be empty")
	assert.NotEmpty(t, reportPlus[1].TransmissionScheduleOverride, "Transmission schedule should not be empty")
}

func TestPlugin_Reports_UnableToParse(t *testing.T) {
	ctx := t.Context()
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
	}
	_, err := p.Reports(ctx, 0, ocr3types.Outcome("not a valid observation"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode outcome")
}

func TestPlugin_Reports_UnableToEncode(t *testing.T) {
	ctx := t.Context()
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Encode", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("test error"))
	// Mock chain support for transmission schedule
	chainSupport := plugincommon_mock.NewMockChainSupport(t)
	chainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil)
	p := &Plugin{
		reportCodec:  codec,
		lggr:         logger.Test(t),
		ocrTypeCodec: ocrTypeCodec,
		chainSupport: chainSupport,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
	}
	report, err := ocrTypeCodec.EncodeOutcome(exectypes.NewOutcome(
		exectypes.Unknown, []exectypes.CommitData{{}, {}},
		[]cciptypes.ExecutePluginReport{{ChainReports: []cciptypes.ExecutePluginReportSingleChain{
			{},
			{},
		}}}))
	require.NoError(t, err)

	_, err = p.Reports(ctx, 0, report)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to encode report")
}

func TestPlugin_ShouldAcceptAttestedReport_DoesNotDecode(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, fmt.Errorf("test error"))

	cs := plugincommon_mock.NewMockChainSupport(t)
	cs.EXPECT().SupportsDestChain(1).Return(true, nil).Maybe()

	homeChain := reader_mock.NewMockHomeChain(t)
	homeChain.EXPECT().GetSupportedChainsForPeer(libocrtypes.PeerID{1}).
		Return(mapset.NewSet[cciptypes.ChainSelector](0), nil).Maybe()

	p := &Plugin{
		reportCodec:     codec,
		lggr:            logger.Test(t),
		chainSupport:    cs,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{0: {1}},
		homeChain:       homeChain,
	}

	_, err := p.ShouldAcceptAttestedReport(t.Context(), 0, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("will not decode"), // faked out, see mock above
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "report validation error: decode exec plugin report: test error")
}

func TestPlugin_ShouldAcceptAttestedReport_NoReports(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, nil)

	cs := plugincommon_mock.NewMockChainSupport(t)
	cs.EXPECT().SupportsDestChain(1).Return(true, nil).Maybe()

	homeChain := reader_mock.NewMockHomeChain(t)
	homeChain.EXPECT().GetSupportedChainsForPeer(libocrtypes.PeerID{1}).
		Return(mapset.NewSet[cciptypes.ChainSelector](0), nil).Maybe()

	p := &Plugin{
		lggr:            logger.Test(t),
		reportCodec:     codec,
		chainSupport:    cs,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{0: {1}},
		homeChain:       homeChain,
	}
	result, err := p.ShouldAcceptAttestedReport(t.Context(), 0, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("empty report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.False(t, result)
}

func TestPlugin_ShouldAcceptAttestedReport_ShouldAccept(t *testing.T) {
	destChain := rand.RandomInt64()
	sourceChain := rand.RandomInt64()
	seqNum := cciptypes.SeqNum(1)
	configDigest := [32]byte{0xde, 0xad, 0xbe, 0xef}

	type depFunc func() (
		*codec_mocks.MockExecutePluginCodec,
		*reader_mock.MockHomeChain,
		*readerpkg_mock.MockCCIPReader)

	basicMockCodec := func() *codec_mocks.MockExecutePluginCodec {
		codec := codec_mocks.NewMockExecutePluginCodec(t)
		codec.EXPECT().Decode(mock.Anything, mock.Anything).
			Return(cciptypes.ExecutePluginReport{
				ChainReports: []cciptypes.ExecutePluginReportSingleChain{
					{
						SourceChainSelector: cciptypes.ChainSelector(sourceChain),
						Messages: []cciptypes.Message{
							{
								Header: cciptypes.RampMessageHeader{
									SequenceNumber:      seqNum,
									SourceChainSelector: cciptypes.ChainSelector(sourceChain),
								},
							},
						},
					},
				},
			}, nil)
		return codec
	}

	basicHomeChain := func() *reader_mock.MockHomeChain {
		homeChainMock := reader_mock.NewMockHomeChain(t)
		homeChainMock.EXPECT().GetChainConfig(cciptypes.ChainSelector(destChain)).Return(reader.ChainConfig{
			SupportedNodes: mapset.NewSet[libocrtypes.PeerID]([32]byte{1}),
		}, nil)
		return homeChainMock
	}

	basicCCIPReader := func() *readerpkg_mock.MockCCIPReader {
		ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
		ccipReaderMock.
			EXPECT().
			GetOffRampConfigDigest(
				mock.Anything,
				consts.PluginTypeExecute,
			).Return(
			configDigest,
			nil,
		)
		// by default, the message is not executed.
		ccipReaderMock.
			EXPECT().
			ExecutedMessages(
				mock.Anything,
				map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					cciptypes.ChainSelector(sourceChain): {cciptypes.NewSeqNumRange(seqNum, seqNum)}},
				primitives.Unconfirmed,
			).
			Return(nil, nil)
		return ccipReaderMock
	}

	testcases := []struct {
		name        string
		getDeps     depFunc
		assertions  func(*testing.T, bool, error)
		logsContain []string
	}{
		{
			name: "should accept",
			getDeps: func() (*codec_mocks.MockExecutePluginCodec,
				*reader_mock.MockHomeChain,
				*readerpkg_mock.MockCCIPReader,
			) {
				mockReader := basicCCIPReader()
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(cciptypes.CurseInfo{
						CursedSourceChains: map[cciptypes.ChainSelector]bool{
							cciptypes.ChainSelector(sourceChain): false,
						}}, nil)

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.NoError(t, err)
				require.True(t, b)
			},
		},
		{
			name: "rmn curse info error",
			getDeps: func() (*codec_mocks.MockExecutePluginCodec,
				*reader_mock.MockHomeChain,
				*readerpkg_mock.MockCCIPReader,
			) {
				mockReader := basicCCIPReader()
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(cciptypes.CurseInfo{}, fmt.Errorf("test error"))

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.ErrorContains(t, err, "error while fetching curse info")
				require.False(t, b)
			},
			logsContain: []string{"report not accepted due to curse checking error"},
		},
		{
			name: "rmn global curse error",
			getDeps: func() (*codec_mocks.MockExecutePluginCodec,
				*reader_mock.MockHomeChain,
				*readerpkg_mock.MockCCIPReader,
			) {
				mockReader := basicCCIPReader()
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(cciptypes.CurseInfo{GlobalCurse: true}, nil)

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.NoError(t, err) // error is logged, but not returned
				require.False(t, b)
			},
			logsContain: []string{"report not accepted due to RMN curse"},
		},
		{
			name: "rmn destination curse error",
			getDeps: func() (*codec_mocks.MockExecutePluginCodec,
				*reader_mock.MockHomeChain,
				*readerpkg_mock.MockCCIPReader,
			) {
				mockReader := basicCCIPReader()
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(cciptypes.CurseInfo{CursedDestination: true}, nil)

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.NoError(t, err) // error is logged, but not returned
				require.False(t, b)
			},
			logsContain: []string{"report not accepted due to RMN curse"},
		},
		{
			name: "rmn source curse error",
			getDeps: func() (*codec_mocks.MockExecutePluginCodec,
				*reader_mock.MockHomeChain,
				*readerpkg_mock.MockCCIPReader,
			) {
				mockReader := basicCCIPReader()
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything).
					Return(cciptypes.CurseInfo{CursedSourceChains: map[cciptypes.ChainSelector]bool{
						cciptypes.ChainSelector(sourceChain): true,
					},
					}, nil)

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.NoError(t, err) // error is logged, but not returned
				require.False(t, b)
			},
			logsContain: []string{"source chains were cursed during report generation"},
		},
		{
			name: "message already executed",
			getDeps: func() (*codec_mocks.MockExecutePluginCodec,
				*reader_mock.MockHomeChain,
				*readerpkg_mock.MockCCIPReader,
			) {
				mockReader := basicCCIPReader()

				// remove the call from index idx from the expected calls slice
				// so that we can override the expectation from the basicCCIPReader call.
				idx := slices.IndexFunc(mockReader.ExpectedCalls, func(e *mock.Call) bool {
					return e.Method == "ExecutedMessages"
				})
				require.GreaterOrEqual(t, idx, 0)
				mockReader.ExpectedCalls = append(mockReader.ExpectedCalls[:idx], mockReader.ExpectedCalls[idx+1:]...)

				mockReader.EXPECT().
					ExecutedMessages(
						mock.Anything,
						map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
							cciptypes.ChainSelector(sourceChain): {cciptypes.NewSeqNumRange(seqNum, seqNum)}},
						primitives.Unconfirmed,
					).Return(
					map[cciptypes.ChainSelector][]cciptypes.SeqNum{
						cciptypes.ChainSelector(sourceChain): {seqNum},
					},
					nil)

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.NoError(t, err) // error is logged, but not returned
				require.False(t, b)
			},
			logsContain: []string{"all messages from source already executed, checking next source"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			codec, homeChain, ccipReader := tc.getDeps()
			lggr, obs := logger.TestObserved(t, zapcore.DebugLevel)

			cs := plugincommon.NewChainSupport(
				logger.Test(t),
				homeChain,
				map[commontypes.OracleID]libocrtypes.PeerID{
					1: [32]byte{1},
				},
				1,
				cciptypes.ChainSelector(destChain),
			)
			homeChain.EXPECT().GetSupportedChainsForPeer(libocrtypes.PeerID{1}).
				Return(mapset.NewSet[cciptypes.ChainSelector](0), nil).Maybe()

			p := &Plugin{
				lggr:            lggr,
				reportCodec:     codec,
				homeChain:       homeChain,
				ccipReader:      ccipReader,
				chainSupport:    cs,
				oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{1: {1}},
				reportingCfg: ocr3types.ReportingPluginConfig{
					OracleID:     1,
					ConfigDigest: configDigest,
				},
			}

			result, err := p.ShouldAcceptAttestedReport(t.Context(), 0, ocr3types.ReportWithInfo[[]byte]{
				Report: []byte("report"), // faked out, see mock above
			})

			tc.assertions(t, result, err)

			for _, expLog := range tc.logsContain {
				found := false
				for _, log := range obs.All() {
					if strings.Contains(log.Message, expLog) {
						found = true
						break
					}
				}
				if !found {
					assert.Fail(t, "expected log not found", expLog)
				}
			}
		})
	}
}

func TestPlugin_ShouldTransmitAcceptReport_NilReport(t *testing.T) {
	lggr := logger.Test(t)

	p := &Plugin{
		lggr: lggr,
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.False(t, shouldTransmit)
}

func TestPlugin_ShouldTransmitAcceptedReport_DecodeFailure(t *testing.T) {
	lggr := logger.Test(t)

	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).Return(cciptypes.ExecutePluginReport{}, fmt.Errorf("test error"))

	p := &Plugin{
		lggr:        lggr,
		reportCodec: codec,
	}

	_, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("will not decode"), // faked out, see mock above
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "report validation error: decode exec plugin report: test error")
}

func TestPlugin_ShouldTransmitAcceptReport_SupportsDestChainCheckFails(t *testing.T) {
	lggr := logger.Test(t)
	oracleID := commontypes.OracleID(1)

	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).Return(cciptypes.ExecutePluginReport{
		ChainReports: []cciptypes.ExecutePluginReportSingleChain{
			{}, {},
		},
	}, nil)

	chainSupport := plugincommon_mock.NewMockChainSupport(t)
	chainSupport.EXPECT().SupportsDestChain(oracleID).Return(false, errors.New("test error"))

	p := &Plugin{
		lggr:         lggr,
		chainSupport: chainSupport,
		reportingCfg: ocr3types.ReportingPluginConfig{
			OracleID: oracleID,
		},
		reportCodec: codec,
	}

	_, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "test error")
}

func TestPlugin_ShouldTransmitAcceptReport_DontSupportDestChain(t *testing.T) {
	lggr := logger.Test(t)
	oracleID := commontypes.OracleID(1)

	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).Return(cciptypes.ExecutePluginReport{
		ChainReports: []cciptypes.ExecutePluginReportSingleChain{
			{}, {},
		},
	}, nil)

	chainSupport := plugincommon_mock.NewMockChainSupport(t)
	chainSupport.EXPECT().SupportsDestChain(oracleID).Return(false, nil)

	p := &Plugin{
		lggr:         lggr,
		chainSupport: chainSupport,
		reportingCfg: ocr3types.ReportingPluginConfig{
			OracleID: oracleID,
		},
		reportCodec: codec,
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.False(t, shouldTransmit)
}

func TestPlugin_ShouldTransmitAcceptedReport_MismatchingConfigDigests(t *testing.T) {
	lggr := logger.Test(t)
	destChain := rand.RandomInt64()
	configDigest := [32]byte{0xde, 0xad, 0xbe, 0xef}
	onchainConfigDigest := [32]byte{0xca, 0xfe, 0xba, 0xbe}
	oracleID := commontypes.OracleID(1)
	peerID := libocrtypes.PeerID{1}

	homeChainMock := reader_mock.NewMockHomeChain(t)
	homeChainMock.EXPECT().GetChainConfig(cciptypes.ChainSelector(destChain)).Return(reader.ChainConfig{
		SupportedNodes: mapset.NewSet(peerID),
	}, nil)

	ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
	ccipReaderMock.
		EXPECT().
		GetOffRampConfigDigest(mock.Anything, consts.PluginTypeExecute).
		Return(onchainConfigDigest, nil)

	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).Return(cciptypes.ExecutePluginReport{
		ChainReports: []cciptypes.ExecutePluginReportSingleChain{
			{}, {},
		},
	}, nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: homeChainMock,
		chainSupport: plugincommon.NewChainSupport(
			logger.Test(t),
			homeChainMock,
			map[commontypes.OracleID]libocrtypes.PeerID{
				oracleID: peerID,
			},
			oracleID,
			cciptypes.ChainSelector(destChain),
		),
		reportingCfg: ocr3types.ReportingPluginConfig{
			OracleID:     oracleID,
			ConfigDigest: configDigest,
		},
		reportCodec: codec,
		ccipReader:  ccipReaderMock,
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.False(t, shouldTransmit)
}

func TestPlugin_ShouldTransmitAcceptReport_Success(t *testing.T) {
	lggr := logger.Test(t)
	destChain := rand.RandomInt64()
	configDigest := [32]byte{0xde, 0xad, 0xbe, 0xef}
	oracleID := commontypes.OracleID(1)
	peerID := libocrtypes.PeerID{1}

	homeChainMock := reader_mock.NewMockHomeChain(t)
	homeChainMock.EXPECT().GetChainConfig(cciptypes.ChainSelector(destChain)).Return(reader.ChainConfig{
		SupportedNodes: mapset.NewSet(peerID),
	}, nil)

	codec := codec_mocks.NewMockExecutePluginCodec(t)
	reports := genRandomChainReports(1, 1)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).Return(cciptypes.ExecutePluginReport{
		ChainReports: reports,
	}, nil)

	ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
	ccipReaderMock.
		EXPECT().
		GetOffRampConfigDigest(
			mock.Anything,
			consts.PluginTypeExecute).
		Return(configDigest, nil)
	ccipReaderMock.
		EXPECT().
		ExecutedMessages(
			mock.Anything,
			map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				reports[0].SourceChainSelector: {cciptypes.NewSeqNumRange(
					reports[0].Messages[0].Header.SequenceNumber,
					reports[0].Messages[0].Header.SequenceNumber,
				)},
			},
			primitives.Unconfirmed,
		).Return(nil, nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: homeChainMock,
		chainSupport: plugincommon.NewChainSupport(
			logger.Test(t),
			homeChainMock,
			map[commontypes.OracleID]libocrtypes.PeerID{
				oracleID: peerID,
			},
			oracleID,
			cciptypes.ChainSelector(destChain),
		),
		reportingCfg: ocr3types.ReportingPluginConfig{
			OracleID:     oracleID,
			ConfigDigest: configDigest,
		},
		reportCodec: codec,
		ccipReader:  ccipReaderMock,
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.True(t, shouldTransmit)
}

func TestPlugin_ShouldTransmitAcceptReport_Failure_AlreadyExecuted(t *testing.T) {
	lggr := logger.Test(t)
	destChain := rand.RandomInt64()
	configDigest := [32]byte{0xde, 0xad, 0xbe, 0xef}
	oracleID := commontypes.OracleID(1)
	peerID := libocrtypes.PeerID{1}

	homeChainMock := reader_mock.NewMockHomeChain(t)
	homeChainMock.EXPECT().GetChainConfig(cciptypes.ChainSelector(destChain)).Return(reader.ChainConfig{
		SupportedNodes: mapset.NewSet(peerID),
	}, nil)

	codec := codec_mocks.NewMockExecutePluginCodec(t)
	reports := genRandomChainReports(1, 1)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).Return(cciptypes.ExecutePluginReport{
		ChainReports: reports,
	}, nil)

	ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
	ccipReaderMock.
		EXPECT().
		GetOffRampConfigDigest(
			mock.Anything,
			consts.PluginTypeExecute).
		Return(configDigest, nil)
	ccipReaderMock.
		EXPECT().
		ExecutedMessages(
			mock.Anything,
			map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				reports[0].SourceChainSelector: {
					cciptypes.NewSeqNumRange(
						reports[0].Messages[0].Header.SequenceNumber,
						reports[0].Messages[0].Header.SequenceNumber,
					)},
			},
			primitives.Unconfirmed, // Changed from Finalized to Unconfirmed
		).Return(
		map[cciptypes.ChainSelector][]cciptypes.SeqNum{
			reports[0].SourceChainSelector: {reports[0].Messages[0].Header.SequenceNumber},
		}, nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: homeChainMock,
		chainSupport: plugincommon.NewChainSupport(
			logger.Test(t),
			homeChainMock,
			map[commontypes.OracleID]libocrtypes.PeerID{
				oracleID: peerID,
			},
			oracleID,
			cciptypes.ChainSelector(destChain),
		),
		reportingCfg: ocr3types.ReportingPluginConfig{
			OracleID:     oracleID,
			ConfigDigest: configDigest,
		},
		reportCodec: codec,
		ccipReader:  ccipReaderMock,
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(t.Context(), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.False(t, shouldTransmit)
}

func mustDecodeBase64(t *testing.T, base64Str string) []byte {
	t.Helper()
	b, err := base64.StdEncoding.DecodeString(base64Str)
	require.NoError(t, err)
	return b
}

// This test is a real-world example of a report that was observed by 2 observers.
// {
// "chainSelector": 14767482510784807000,
// "OnRampAddress": "0x00000000000000000000000045004439953d8f2f52f7350b567ae644fbda469e",
// "timestamp": "2025-02-20T15:37:55Z",
// "blockNum": 125637894,
// "merkleRoot": "0xafbbb6789a9d008d0438210e0d9ab1cde75f811b09b8b6d658822c00cd18c9e6",
// "sequenceNumberRange": [
// 168,
// 179
// ],
// "executedMessages": [
// 168,
// 169,
// 170,
// 171,
// 172,
// 173,
// 174,
// 175
// ],
// "messages": null,
// "messageHashes": null,
// "costlyMessages": null,
// "messageTokenData": null
// }
// And 2 observers observe the root
// 0xafbbb6789a9d008d0438210e0d9ab1cde75f811b09b8b6d658822c00cd18c9e6 (same as above) with no messages executed:
//
// {
// "chainSelector": 14767482510784807000,
// "OnRampAddress": "0x00000000000000000000000045004439953d8f2f52f7350b567ae644fbda469e",
// "timestamp": "2025-02-20T15:37:55Z",
// "blockNum": 125637894,
// "merkleRoot": "0xafbbb6789a9d008d0438210e0d9ab1cde75f811b09b8b6d658822c00cd18c9e6",
// "sequenceNumberRange": [
// 168,
// 179
// ],
// "executedMessages": null,
// "messages": null,
// "messageHashes": null,
// "costlyMessages": null,
// "messageTokenData": null
// }
func TestPlugin_Outcome_RealworldObservation(t *testing.T) {
	var attObs = []types.AttributedObservation{
		{
			//nolint:lll
			Observation: mustDecodeBase64(t, "eyJjb21taXRSZXBvcnRzIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjpudWxsLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6W3siY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjM3OjU1WiIsImJsb2NrTnVtIjoxMjU2Mzc4OTQsIm1lcmtsZVJvb3QiOiIweGFmYmJiNjc4OWE5ZDAwOGQwNDM4MjEwZTBkOWFiMWNkZTc1ZjgxMWIwOWI4YjZkNjU4ODIyYzAwY2QxOGM5ZTYiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMTY4LDE3OV0sImV4ZWN1dGVkTWVzc2FnZXMiOlsxNjgsMTY5LDE3MCwxNzEsMTcyLDE3MywxNzQsMTc1XSwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTozOToxOVoiLCJibG9ja051bSI6MTI1NjM4MjA0LCJtZXJrbGVSb290IjoiMHhkMWZhZjgyYWM5ZGNhNzZjY2FiZmEzNDM2NzEzZWRlZGE2NzgzYmJlZDlmMWVmZDFjNjk0ZmEzYWFmZWYyMzRlIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzE4MCwxOTBdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQwOjQ1WiIsImJsb2NrTnVtIjoxMjU2Mzg1MzgsIm1lcmtsZVJvb3QiOiIweDkyNGFhNDA0ZWY1YWZlNzA0ZjNlYWIyMTAxOTkwMDU2ODQ5MDU0NDEyYzg3ZDdlZGI0YTFkZjkyYWJmZmQzY2YiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMTkxLDIwMV0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6NDI6MDRaIiwiYmxvY2tOdW0iOjEyNTYzODgzNywibWVya2xlUm9vdCI6IjB4MWE2OTJmZDAxYmRlODM2NzUwYmVlYjdlNGRmM2Q5NGE0ZjJhNDU1YmJkYjczNjZiNjNmNjJlYWYwNWE2NGEzYSIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsyMDIsMjE0XSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0MzoyOVoiLCJibG9ja051bSI6MTI1NjM5MTcyLCJtZXJrbGVSb290IjoiMHg5MmEyN2M4MjhjNTEyOTIwNDJhMjBjMjYwOWU2NTQxMzJjYjUxY2M5MWM5ZDg5YTY2ZmZlZTA0OGU0NmFhZGM0Iiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzIxNSwyMjddLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQ0OjUzWiIsImJsb2NrTnVtIjoxMjU2Mzk1MDIsIm1lcmtsZVJvb3QiOiIweDEzMmUzYjMyN2ExNzdkMGE4MDIzNTA0ZGIxZTk4YzAxYmM1MmRmNTYxYjU4NzNkOGViOGZjZTBlNjNkOWUyNzIiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjI4LDIzNl0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH1dLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6bnVsbCwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOm51bGwsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOm51bGx9LCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwidG9rZW5EYXRhT2JzZXJ2YXRpb25zIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibm9uY2VzIjpudWxsLCJjb250cmFjdHMiOnsiRkNoYWluIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoxLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6MSwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOjEsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoxLCIzNDc4NDg3MjM4NTI0NTEyMTA2IjoxLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoxfSwiQWRkcmVzc2VzIjp7IkZlZVF1b3RlciI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHhlYWI3NDI5N2U2YmIzMGEwNjNlYmE4ZmUxYWE0NTA1ODkyOWJmZjJjIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweGZlZTcxOWZmYWQwZGM2MTI0NmI4MmQ5YjgwZmEzNWMyZjhiZjkzODciLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweDMzNzEyMzg2MTUxZTU3YTAxZmQ3MWFiOTBlNTljMDAxMGQ4ZDEwZDAiLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoiMHg5NDVkOTg0NWJjMTRhNWU0ZmY2NDQ1N2JiNjc3MGFhYzAyOGNlZDM5In0sIk5vbmNlTWFuYWdlciI6eyIzNDc4NDg3MjM4NTI0NTEyMTA2IjoiMHhlNTMyZjZjYWJmODg5Mjk5YTI5NTRjMGFlODA2MzU4MmNkNjcyMDY1In0sIk9uUmFtcCI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAyNGE2ODA0YzBhZmI5NzE4OWFlZmZkZTUzNzFlNjgzNWM1N2M2ZDMiLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDIwYzg2NDlmZGU0OGZjNzk1MTU0ZTFmOWY1ZDlkZWVmMDcxOTY5M2MifSwiUk1OUmVtb3RlIjp7IjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweGFlNWFlYjJjMTE5MDUzNTgzOTRlYmY0NGEzMjU0MjYyN2Y0OTJhODMifSwiUm91dGVyIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMTQ3Njc0ODI1MTA3ODQ4MDYwNDMiOiIweDgwNmNjY2M1ZmQzZWRiOGNiMjRhNzRmZGE3ZGUyNGQ4NGNlMGQxZmIiLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6IjB4NDBjOWRmNmUyYmU3ZWQwNjY5NGUxMGQ5NDU1OTA0OWNjYzIzOGIxNCIsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMzQ3ODQ4NzIzODUyNDUxMjEwNiI6IjB4ZWNiMmQ0MDdjOTU1MWUyZTEwOGNjYWYwMjZlOTEzNGYyYjM0MTZkNyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweGQzZTE5MGYzODFmMDZkYzBkMjg5NTkwZmQ0NTJjNDJmYTJkYWM1ODYifX19LCJmQ2hhaW4iOnsiMTAzNDQ5NzEyMzU4NzQ0NjUwODAiOjEsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoxLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6MSwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOjEsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOjEsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOjF9fQ=="),
			Observer:    3,
		},
		{
			//nolint:lll
			Observation: mustDecodeBase64(t, "eyJjb21taXRSZXBvcnRzIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjpudWxsLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6W3siY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjM3OjU1WiIsImJsb2NrTnVtIjoxMjU2Mzc4OTQsIm1lcmtsZVJvb3QiOiIweGFmYmJiNjc4OWE5ZDAwOGQwNDM4MjEwZTBkOWFiMWNkZTc1ZjgxMWIwOWI4YjZkNjU4ODIyYzAwY2QxOGM5ZTYiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMTY4LDE3OV0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6Mzk6MTlaIiwiYmxvY2tOdW0iOjEyNTYzODIwNCwibWVya2xlUm9vdCI6IjB4ZDFmYWY4MmFjOWRjYTc2Y2NhYmZhMzQzNjcxM2VkZWRhNjc4M2JiZWQ5ZjFlZmQxYzY5NGZhM2FhZmVmMjM0ZSIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsxODAsMTkwXSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0MDo0NVoiLCJibG9ja051bSI6MTI1NjM4NTM4LCJtZXJrbGVSb290IjoiMHg5MjRhYTQwNGVmNWFmZTcwNGYzZWFiMjEwMTk5MDA1Njg0OTA1NDQxMmM4N2Q3ZWRiNGExZGY5MmFiZmZkM2NmIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzE5MSwyMDFdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQyOjA0WiIsImJsb2NrTnVtIjoxMjU2Mzg4MzcsIm1lcmtsZVJvb3QiOiIweDFhNjkyZmQwMWJkZTgzNjc1MGJlZWI3ZTRkZjNkOTRhNGYyYTQ1NWJiZGI3MzY2YjYzZjYyZWFmMDVhNjRhM2EiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjAyLDIxNF0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6NDM6MjlaIiwiYmxvY2tOdW0iOjEyNTYzOTE3MiwibWVya2xlUm9vdCI6IjB4OTJhMjdjODI4YzUxMjkyMDQyYTIwYzI2MDllNjU0MTMyY2I1MWNjOTFjOWQ4OWE2NmZmZWUwNDhlNDZhYWRjNCIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsyMTUsMjI3XSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0NDo1M1oiLCJibG9ja051bSI6MTI1NjM5NTAyLCJtZXJrbGVSb290IjoiMHgxMzJlM2IzMjdhMTc3ZDBhODAyMzUwNGRiMWU5OGMwMWJjNTJkZjU2MWI1ODczZDhlYjhmY2UwZTYzZDllMjcyIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzIyOCwyMzZdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQ0OjUzWiIsImJsb2NrTnVtIjoxMjU2Mzk1MDIsIm1lcmtsZVJvb3QiOiIweDEzMmUzYjMyN2ExNzdkMGE4MDIzNTA0ZGIxZTk4YzAxYmM1MmRmNTYxYjU4NzNkOGViOGZjZTBlNjNkOWUyNzIiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjI4LDIzNl0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH1dLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6bnVsbCwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOm51bGwsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOm51bGx9LCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwidG9rZW5EYXRhT2JzZXJ2YXRpb25zIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibm9uY2VzIjpudWxsLCJjb250cmFjdHMiOnsiRkNoYWluIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoxLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6MSwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOjEsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoxLCIzNDc4NDg3MjM4NTI0NTEyMTA2IjoxLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoxfSwiQWRkcmVzc2VzIjp7IkZlZVF1b3RlciI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHhlYWI3NDI5N2U2YmIzMGEwNjNlYmE4ZmUxYWE0NTA1ODkyOWJmZjJjIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweGZlZTcxOWZmYWQwZGM2MTI0NmI4MmQ5YjgwZmEzNWMyZjhiZjkzODciLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweDMzNzEyMzg2MTUxZTU3YTAxZmQ3MWFiOTBlNTljMDAxMGQ4ZDEwZDAiLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoiMHg5NDVkOTg0NWJjMTRhNWU0ZmY2NDQ1N2JiNjc3MGFhYzAyOGNlZDM5In0sIk5vbmNlTWFuYWdlciI6eyIzNDc4NDg3MjM4NTI0NTEyMTA2IjoiMHhlNTMyZjZjYWJmODg5Mjk5YTI5NTRjMGFlODA2MzU4MmNkNjcyMDY1In0sIk9uUmFtcCI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAyNGE2ODA0YzBhZmI5NzE4OWFlZmZkZTUzNzFlNjgzNWM1N2M2ZDMiLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDIwYzg2NDlmZGU0OGZjNzk1MTU0ZTFmOWY1ZDlkZWVmMDcxOTY5M2MifSwiUk1OUmVtb3RlIjp7IjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweGFlNWFlYjJjMTE5MDUzNTgzOTRlYmY0NGEzMjU0MjYyN2Y0OTJhODMifSwiUm91dGVyIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMTQ3Njc0ODI1MTA3ODQ4MDYwNDMiOiIweDgwNmNjY2M1ZmQzZWRiOGNiMjRhNzRmZGE3ZGUyNGQ4NGNlMGQxZmIiLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6IjB4NDBjOWRmNmUyYmU3ZWQwNjY5NGUxMGQ5NDU1OTA0OWNjYzIzOGIxNCIsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMzQ3ODQ4NzIzODUyNDUxMjEwNiI6IjB4ZWNiMmQ0MDdjOTU1MWUyZTEwOGNjYWYwMjZlOTEzNGYyYjM0MTZkNyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweGQzZTE5MGYzODFmMDZkYzBkMjg5NTkwZmQ0NTJjNDJmYTJkYWM1ODYifX19LCJmQ2hhaW4iOnsiMTAzNDQ5NzEyMzU4NzQ0NjUwODAiOjEsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoxLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6MSwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOjEsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOjEsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOjF9fQ=="),
			Observer:    1,
		},
		{
			//nolint:lll
			Observation: mustDecodeBase64(t, "eyJjb21taXRSZXBvcnRzIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjpudWxsLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6W3siY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjM3OjU1WiIsImJsb2NrTnVtIjoxMjU2Mzc4OTQsIm1lcmtsZVJvb3QiOiIweGFmYmJiNjc4OWE5ZDAwOGQwNDM4MjEwZTBkOWFiMWNkZTc1ZjgxMWIwOWI4YjZkNjU4ODIyYzAwY2QxOGM5ZTYiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMTY4LDE3OV0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6Mzk6MTlaIiwiYmxvY2tOdW0iOjEyNTYzODIwNCwibWVya2xlUm9vdCI6IjB4ZDFmYWY4MmFjOWRjYTc2Y2NhYmZhMzQzNjcxM2VkZWRhNjc4M2JiZWQ5ZjFlZmQxYzY5NGZhM2FhZmVmMjM0ZSIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsxODAsMTkwXSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0MDo0NVoiLCJibG9ja051bSI6MTI1NjM4NTM4LCJtZXJrbGVSb290IjoiMHg5MjRhYTQwNGVmNWFmZTcwNGYzZWFiMjEwMTk5MDA1Njg0OTA1NDQxMmM4N2Q3ZWRiNGExZGY5MmFiZmZkM2NmIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzE5MSwyMDFdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQyOjA0WiIsImJsb2NrTnVtIjoxMjU2Mzg4MzcsIm1lcmtsZVJvb3QiOiIweDFhNjkyZmQwMWJkZTgzNjc1MGJlZWI3ZTRkZjNkOTRhNGYyYTQ1NWJiZGI3MzY2YjYzZjYyZWFmMDVhNjRhM2EiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjAyLDIxNF0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6NDM6MjlaIiwiYmxvY2tOdW0iOjEyNTYzOTE3MiwibWVya2xlUm9vdCI6IjB4OTJhMjdjODI4YzUxMjkyMDQyYTIwYzI2MDllNjU0MTMyY2I1MWNjOTFjOWQ4OWE2NmZmZWUwNDhlNDZhYWRjNCIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsyMTUsMjI3XSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0NDo1M1oiLCJibG9ja051bSI6MTI1NjM5NTAyLCJtZXJrbGVSb290IjoiMHgxMzJlM2IzMjdhMTc3ZDBhODAyMzUwNGRiMWU5OGMwMWJjNTJkZjU2MWI1ODczZDhlYjhmY2UwZTYzZDllMjcyIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzIyOCwyMzZdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQ0OjUzWiIsImJsb2NrTnVtIjoxMjU2Mzk1MDIsIm1lcmtsZVJvb3QiOiIweDEzMmUzYjMyN2ExNzdkMGE4MDIzNTA0ZGIxZTk4YzAxYmM1MmRmNTYxYjU4NzNkOGViOGZjZTBlNjNkOWUyNzIiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjI4LDIzNl0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH1dLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6bnVsbCwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOm51bGwsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOm51bGx9LCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwidG9rZW5EYXRhT2JzZXJ2YXRpb25zIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibm9uY2VzIjpudWxsLCJjb250cmFjdHMiOnsiRkNoYWluIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoxLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6MSwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOjEsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoxLCIzNDc4NDg3MjM4NTI0NTEyMTA2IjoxLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoxfSwiQWRkcmVzc2VzIjp7IkZlZVF1b3RlciI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHhlYWI3NDI5N2U2YmIzMGEwNjNlYmE4ZmUxYWE0NTA1ODkyOWJmZjJjIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweGZlZTcxOWZmYWQwZGM2MTI0NmI4MmQ5YjgwZmEzNWMyZjhiZjkzODciLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweDMzNzEyMzg2MTUxZTU3YTAxZmQ3MWFiOTBlNTljMDAxMGQ4ZDEwZDAiLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoiMHg5NDVkOTg0NWJjMTRhNWU0ZmY2NDQ1N2JiNjc3MGFhYzAyOGNlZDM5In0sIk5vbmNlTWFuYWdlciI6eyIzNDc4NDg3MjM4NTI0NTEyMTA2IjoiMHhlNTMyZjZjYWJmODg5Mjk5YTI5NTRjMGFlODA2MzU4MmNkNjcyMDY1In0sIk9uUmFtcCI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAyNGE2ODA0YzBhZmI5NzE4OWFlZmZkZTUzNzFlNjgzNWM1N2M2ZDMiLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDIwYzg2NDlmZGU0OGZjNzk1MTU0ZTFmOWY1ZDlkZWVmMDcxOTY5M2MifSwiUk1OUmVtb3RlIjp7IjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweGFlNWFlYjJjMTE5MDUzNTgzOTRlYmY0NGEzMjU0MjYyN2Y0OTJhODMifSwiUm91dGVyIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMTQ3Njc0ODI1MTA3ODQ4MDYwNDMiOiIweDgwNmNjY2M1ZmQzZWRiOGNiMjRhNzRmZGE3ZGUyNGQ4NGNlMGQxZmIiLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6IjB4NDBjOWRmNmUyYmU3ZWQwNjY5NGUxMGQ5NDU1OTA0OWNjYzIzOGIxNCIsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMzQ3ODQ4NzIzODUyNDUxMjEwNiI6IjB4ZWNiMmQ0MDdjOTU1MWUyZTEwOGNjYWYwMjZlOTEzNGYyYjM0MTZkNyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweGQzZTE5MGYzODFmMDZkYzBkMjg5NTkwZmQ0NTJjNDJmYTJkYWM1ODYifX19LCJmQ2hhaW4iOnsiMTAzNDQ5NzEyMzU4NzQ0NjUwODAiOjEsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoxLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6MSwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOjEsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOjEsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOjF9fQ=="),
			Observer:    2,
		},
		{
			//nolint:lll
			Observation: mustDecodeBase64(t, "eyJjb21taXRSZXBvcnRzIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjpudWxsLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6W3siY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjM3OjU1WiIsImJsb2NrTnVtIjoxMjU2Mzc4OTQsIm1lcmtsZVJvb3QiOiIweGFmYmJiNjc4OWE5ZDAwOGQwNDM4MjEwZTBkOWFiMWNkZTc1ZjgxMWIwOWI4YjZkNjU4ODIyYzAwY2QxOGM5ZTYiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMTY4LDE3OV0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6Mzk6MTlaIiwiYmxvY2tOdW0iOjEyNTYzODIwNCwibWVya2xlUm9vdCI6IjB4ZDFmYWY4MmFjOWRjYTc2Y2NhYmZhMzQzNjcxM2VkZWRhNjc4M2JiZWQ5ZjFlZmQxYzY5NGZhM2FhZmVmMjM0ZSIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsxODAsMTkwXSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0MDo0NVoiLCJibG9ja051bSI6MTI1NjM4NTM4LCJtZXJrbGVSb290IjoiMHg5MjRhYTQwNGVmNWFmZTcwNGYzZWFiMjEwMTk5MDA1Njg0OTA1NDQxMmM4N2Q3ZWRiNGExZGY5MmFiZmZkM2NmIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzE5MSwyMDFdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQyOjA0WiIsImJsb2NrTnVtIjoxMjU2Mzg4MzcsIm1lcmtsZVJvb3QiOiIweDFhNjkyZmQwMWJkZTgzNjc1MGJlZWI3ZTRkZjNkOTRhNGYyYTQ1NWJiZGI3MzY2YjYzZjYyZWFmMDVhNjRhM2EiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjAyLDIxNF0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH0seyJjaGFpblNlbGVjdG9yIjoxNDc2NzQ4MjUxMDc4NDgwNjA0MywiT25SYW1wQWRkcmVzcyI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwNDUwMDQ0Mzk5NTNkOGYyZjUyZjczNTBiNTY3YWU2NDRmYmRhNDY5ZSIsInRpbWVzdGFtcCI6IjIwMjUtMDItMjBUMTU6NDM6MjlaIiwiYmxvY2tOdW0iOjEyNTYzOTE3MiwibWVya2xlUm9vdCI6IjB4OTJhMjdjODI4YzUxMjkyMDQyYTIwYzI2MDllNjU0MTMyY2I1MWNjOTFjOWQ4OWE2NmZmZWUwNDhlNDZhYWRjNCIsInNlcXVlbmNlTnVtYmVyUmFuZ2UiOlsyMTUsMjI3XSwiZXhlY3V0ZWRNZXNzYWdlcyI6bnVsbCwibWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VIYXNoZXMiOm51bGwsImNvc3RseU1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlVG9rZW5EYXRhIjpudWxsfSx7ImNoYWluU2VsZWN0b3IiOjE0NzY3NDgyNTEwNzg0ODA2MDQzLCJPblJhbXBBZGRyZXNzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwidGltZXN0YW1wIjoiMjAyNS0wMi0yMFQxNTo0NDo1M1oiLCJibG9ja051bSI6MTI1NjM5NTAyLCJtZXJrbGVSb290IjoiMHgxMzJlM2IzMjdhMTc3ZDBhODAyMzUwNGRiMWU5OGMwMWJjNTJkZjU2MWI1ODczZDhlYjhmY2UwZTYzZDllMjcyIiwic2VxdWVuY2VOdW1iZXJSYW5nZSI6WzIyOCwyMzZdLCJleGVjdXRlZE1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwiY29zdGx5TWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VUb2tlbkRhdGEiOm51bGx9LHsiY2hhaW5TZWxlY3RvciI6MTQ3Njc0ODI1MTA3ODQ4MDYwNDMsIk9uUmFtcEFkZHJlc3MiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDQ1MDA0NDM5OTUzZDhmMmY1MmY3MzUwYjU2N2FlNjQ0ZmJkYTQ2OWUiLCJ0aW1lc3RhbXAiOiIyMDI1LTAyLTIwVDE1OjQ0OjUzWiIsImJsb2NrTnVtIjoxMjU2Mzk1MDIsIm1lcmtsZVJvb3QiOiIweDEzMmUzYjMyN2ExNzdkMGE4MDIzNTA0ZGIxZTk4YzAxYmM1MmRmNTYxYjU4NzNkOGViOGZjZTBlNjNkOWUyNzIiLCJzZXF1ZW5jZU51bWJlclJhbmdlIjpbMjI4LDIzNl0sImV4ZWN1dGVkTWVzc2FnZXMiOm51bGwsIm1lc3NhZ2VzIjpudWxsLCJtZXNzYWdlSGFzaGVzIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibWVzc2FnZVRva2VuRGF0YSI6bnVsbH1dLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6bnVsbCwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOm51bGwsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOm51bGx9LCJtZXNzYWdlcyI6bnVsbCwibWVzc2FnZUhhc2hlcyI6bnVsbCwidG9rZW5EYXRhT2JzZXJ2YXRpb25zIjpudWxsLCJjb3N0bHlNZXNzYWdlcyI6bnVsbCwibm9uY2VzIjpudWxsLCJjb250cmFjdHMiOnsiRkNoYWluIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoxLCIxNDc2NzQ4MjUxMDc4NDgwNjA0MyI6MSwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOjEsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoxLCIzNDc4NDg3MjM4NTI0NTEyMTA2IjoxLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoxfSwiQWRkcmVzc2VzIjp7IkZlZVF1b3RlciI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHhlYWI3NDI5N2U2YmIzMGEwNjNlYmE4ZmUxYWE0NTA1ODkyOWJmZjJjIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweGZlZTcxOWZmYWQwZGM2MTI0NmI4MmQ5YjgwZmEzNWMyZjhiZjkzODciLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4OTQ1ZDk4NDViYzE0YTVlNGZmNjQ0NTdiYjY3NzBhYWMwMjhjZWQzOSIsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweDMzNzEyMzg2MTUxZTU3YTAxZmQ3MWFiOTBlNTljMDAxMGQ4ZDEwZDAiLCI1MjI0NDczMjc3MjM2MzMxMjk1IjoiMHg5NDVkOTg0NWJjMTRhNWU0ZmY2NDQ1N2JiNjc3MGFhYzAyOGNlZDM5In0sIk5vbmNlTWFuYWdlciI6eyIzNDc4NDg3MjM4NTI0NTEyMTA2IjoiMHhlNTMyZjZjYWJmODg5Mjk5YTI5NTRjMGFlODA2MzU4MmNkNjcyMDY1In0sIk9uUmFtcCI6eyIxMDM0NDk3MTIzNTg3NDQ2NTA4MCI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoiMHgwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDA0NTAwNDQzOTk1M2Q4ZjJmNTJmNzM1MGI1NjdhZTY0NGZiZGE0NjllIiwiMTYwMTUyODY2MDE3NTc4MjU3NTMiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAyNGE2ODA0YzBhZmI5NzE4OWFlZmZkZTUzNzFlNjgzNWM1N2M2ZDMiLCIxNjI4MTcxMTM5MTY3MDYzNDQ0NSI6IjB4MDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMjBjODY0OWZkZTQ4ZmM3OTUxNTRlMWY5ZjVkOWRlZWYwNzE5NjkzYyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDIwYzg2NDlmZGU0OGZjNzk1MTU0ZTFmOWY1ZDlkZWVmMDcxOTY5M2MifSwiUk1OUmVtb3RlIjp7IjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOiIweGFlNWFlYjJjMTE5MDUzNTgzOTRlYmY0NGEzMjU0MjYyN2Y0OTJhODMifSwiUm91dGVyIjp7IjEwMzQ0OTcxMjM1ODc0NDY1MDgwIjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMTQ3Njc0ODI1MTA3ODQ4MDYwNDMiOiIweDgwNmNjY2M1ZmQzZWRiOGNiMjRhNzRmZGE3ZGUyNGQ4NGNlMGQxZmIiLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6IjB4NDBjOWRmNmUyYmU3ZWQwNjY5NGUxMGQ5NDU1OTA0OWNjYzIzOGIxNCIsIjE2MjgxNzExMzkxNjcwNjM0NDQ1IjoiMHhkM2UxOTBmMzgxZjA2ZGMwZDI4OTU5MGZkNDUyYzQyZmEyZGFjNTg2IiwiMzQ3ODQ4NzIzODUyNDUxMjEwNiI6IjB4ZWNiMmQ0MDdjOTU1MWUyZTEwOGNjYWYwMjZlOTEzNGYyYjM0MTZkNyIsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOiIweGQzZTE5MGYzODFmMDZkYzBkMjg5NTkwZmQ0NTJjNDJmYTJkYWM1ODYifX19LCJmQ2hhaW4iOnsiMTAzNDQ5NzEyMzU4NzQ0NjUwODAiOjEsIjE0NzY3NDgyNTEwNzg0ODA2MDQzIjoxLCIxNjAxNTI4NjYwMTc1NzgyNTc1MyI6MSwiMTYyODE3MTEzOTE2NzA2MzQ0NDUiOjEsIjM0Nzg0ODcyMzg1MjQ1MTIxMDYiOjEsIjUyMjQ0NzMyNzcyMzYzMzEyOTUiOjF9fQ=="),
			Observer:    0,
		},
	}

	jsonCodec := ocrtypecodec.NewExecCodecJSON()

	ctx := t.Context()
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonCodec,
		destChain:    3478487238524512106,
		observer:     &metrics.Noop{},
	}

	prevOutcomeBytes := []byte(`{"State":"Initialized","commitReports":[],"report":{"chainReports":[]}}`)

	outcomeBytes, err := p.Outcome(ctx, ocr3types.OutcomeContext{
		SeqNr:           190942,
		PreviousOutcome: prevOutcomeBytes,
	}, nil, attObs)
	require.NoError(t, err)

	decodedOutcome, err := jsonCodec.DecodeOutcome(outcomeBytes)
	require.NoError(t, err)
	// assert uniqueness of merkle roots in commitreports
	merkleRoots := make(map[string]struct{})
	for _, commitReport := range decodedOutcome.CommitReports {
		merkleRoots[commitReport.MerkleRoot.String()] = struct{}{}
	}
	require.Equal(t, len(merkleRoots), len(decodedOutcome.CommitReports))
}

func TestCommitRootsCache_SkippedRootScenario(t *testing.T) {
	lggr := logger.Test(t)

	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	// Create chain selector and roots
	selector := cciptypes.ChainSelector(1)
	root1 := cciptypes.Bytes32{1}
	root2 := cciptypes.Bytes32{2}
	root3 := cciptypes.Bytes32{3}

	// Create the cache
	cache := cache.NewCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
	)

	t.Run("Root2 is not missed when Root1 and Root3 are executed", func(t *testing.T) {
		// Initially all roots should be executable
		assert.True(t, cache.CanExecute(selector, root1),
			"Root1 should be executable initially")
		assert.True(t, cache.CanExecute(selector, root2),
			"Root2 should be executable initially")
		assert.True(t, cache.CanExecute(selector, root3),
			"Root3 should be executable initially")

		// Execute Root1 and Root3, but not Root2
		cache.MarkAsExecuted(selector, root1)
		cache.MarkAsExecuted(selector, root3)

		// Verify Root2 is still marked as executable
		assert.True(t, cache.CanExecute(selector, root2),
			"Root2 should still be executable")
		assert.False(t, cache.CanExecute(selector, root1),
			"Root1 should not be executable")
		assert.False(t, cache.CanExecute(selector, root3),
			"Root3 should not be executable")

		// Now execute Root2
		cache.MarkAsExecuted(selector, root2)

		// Verify all roots are now non-executable
		assert.False(t, cache.CanExecute(selector, root1),
			"Root1 should not be executable")
		assert.False(t, cache.CanExecute(selector, root2),
			"Root2 should not be executable")
		assert.False(t, cache.CanExecute(selector, root3),
			"Root3 should not be executable")
	})
}

// configurableCommitReportCache is a mock for testing specific cache behaviors.
// Define this near Test_getPendingReportsForExecution or in a test helper section
type configurableCommitReportCache struct {
	reportsToQueryFrom      time.Time
	cachedReportsToReturn   []cciptypes.CommitPluginReportWithMeta
	lggr                    logger.Logger
	deduplicateShouldReturn []cciptypes.CommitPluginReportWithMeta // Control what DeduplicateReports returns
}

func (c *configurableCommitReportCache) RefreshCache(ctx context.Context) error { return nil }
func (c *configurableCommitReportCache) GetReportsToQueryFromTimestamp() time.Time {
	return c.reportsToQueryFrom
}
func (c *configurableCommitReportCache) GetCachedReports(
	fromTimestamp time.Time) []cciptypes.CommitPluginReportWithMeta {
	return c.cachedReportsToReturn
}
func (c *configurableCommitReportCache) DeduplicateReports(
	reports []cciptypes.CommitPluginReportWithMeta,
) []cciptypes.CommitPluginReportWithMeta {
	if c.deduplicateShouldReturn != nil {
		return c.deduplicateShouldReturn
	}
	// Fallback to real deduplication if not overridden, requires lggr to be set.
	if c.lggr != nil {
		return cache.DeduplicateReports(c.lggr, reports)
	}
	return reports // Or panic, or handle error, if lggr is essential but not provided
}
