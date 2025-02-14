package execute

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	plugincommon_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	codec_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec"
	reader2 "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

var jsonOcrTypeCodec = ocrtypecodec.NewExecCodecJSON()

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

func Test_getMinMaxSeqNrRangesBySource(t *testing.T) {
	chainReports := genRandomChainReports(10, 15)
	minMaxSeqNrRanges := getSeqNrRangeBySource(chainReports)
	require.Len(t, minMaxSeqNrRanges, len(chainReports))
	for _, chainReport := range chainReports {
		seqNrRange, ok := minMaxSeqNrRanges[chainReport.SourceChainSelector]
		require.True(t, ok)
		require.NotNil(t, seqNrRange)
		expectedMin := chainReport.Messages[0].Header.SequenceNumber
		expectedMax := chainReport.Messages[len(chainReport.Messages)-1].Header.SequenceNumber
		require.Equal(t, expectedMin, seqNrRange[0])
		require.Equal(t, expectedMax, seqNrRange[1])
	}
}

func Test_checkAlreadyExecuted(t *testing.T) {
	chainReports := genRandomChainReports(10, 15)
	minMaxSeqNrRanges := getSeqNrRangeBySource(chainReports)

	t.Run("some executed", func(t *testing.T) {
		ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
		// need to setup assertions like this because map iteration
		// order is undefined.
		// Basically there is one chain report that is executed and the
		// rest are not.
		midPoint := len(minMaxSeqNrRanges) / 2
		i := 0
		for sourceSel, seqNrRange := range minMaxSeqNrRanges {
			if i == midPoint {
				ccipReaderMock.EXPECT().ExecutedMessages(
					mock.Anything,
					sourceSel,
					seqNrRange,
				).Return(seqNrRange.ToSlice(), nil)
			} else {
				ccipReaderMock.EXPECT().ExecutedMessages(
					mock.Anything,
					sourceSel,
					seqNrRange,
				).Return(nil, nil).Maybe() // not executed
			}
			i++
		}

		p := &Plugin{
			ccipReader: ccipReaderMock,
			lggr:       logger.Test(t),
		}

		err := p.checkAlreadyExecuted(tests.Context(t), p.lggr, minMaxSeqNrRanges)
		require.Error(t, err)
		require.ErrorContains(t, err, "already executed range")
	})

	t.Run("all unexecuted", func(t *testing.T) {
		ccipReaderMock := readerpkg_mock.NewMockCCIPReader(t)
		for sourceSel, seqNrRange := range minMaxSeqNrRanges {
			ccipReaderMock.EXPECT().ExecutedMessages(
				mock.Anything,
				sourceSel,
				seqNrRange,
			).Return(nil, nil)
		}

		p := &Plugin{
			ccipReader: ccipReaderMock,
			lggr:       logger.Test(t),
		}

		err := p.checkAlreadyExecuted(tests.Context(t), p.lggr, minMaxSeqNrRanges)
		require.NoError(t, err)
	})
}

func Test_getPendingExecutedReports(t *testing.T) {
	canExecute := func(ret bool) CanExecuteHandle {
		return func(cciptypes.ChainSelector, cciptypes.Bytes32) bool { return ret }
	}

	tcs := []struct {
		name         string
		reports      []plugintypes2.CommitPluginReportWithMeta
		ranges       map[cciptypes.ChainSelector][]cciptypes.SeqNum
		canExec      CanExecuteHandle
		wantObs      exectypes.CommitObservations
		wantExecuted []exectypes.CommitData
		wantErr      assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			reports: nil,
			ranges:  nil,
			wantObs: exectypes.CommitObservations{},
			wantErr: assert.NoError,
		},
		{
			name: "single non-executed report",
			reports: []plugintypes2.CommitPluginReportWithMeta{
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
			wantErr: assert.NoError,
		},
		{
			name: "single half-executed report",
			reports: []plugintypes2.CommitPluginReportWithMeta{
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
				1: {1, 2, 3, 7, 8},
			},
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
			wantErr: assert.NoError,
		},
		{
			name: "last timestamp",
			reports: []plugintypes2.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report:    cciptypes.CommitPluginReport{},
				},
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(9999999999999999),
					Report:    cciptypes.CommitPluginReport{},
				},
			},
			ranges:  map[cciptypes.ChainSelector][]cciptypes.SeqNum{},
			wantObs: exectypes.CommitObservations{},
			wantErr: assert.NoError,
		},
		{
			name: "single fully-executed report",
			reports: []plugintypes2.CommitPluginReportWithMeta{
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
			wantObs: exectypes.CommitObservations{
				1: nil,
			},
			wantExecuted: []exectypes.CommitData{
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
			name: "single known-executed report",
			reports: []plugintypes2.CommitPluginReportWithMeta{
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
			canExec: canExecute(false),
			ranges:  nil,
			wantObs: exectypes.CommitObservations{
				1: nil,
			},
			wantExecuted: nil, // the executed message is not returned.
			wantErr:      assert.NoError,
		},
	}
	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.canExec == nil {
				tt.canExec = canExecute(true)
			}
			mockReader := readerpkg_mock.NewMockCCIPReader(t)
			mockReader.On(
				"CommitReportsGTETimestamp", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			).Return(tt.reports, nil)
			for k, v := range tt.ranges {
				mockReader.On("ExecutedMessages", mock.Anything, k, mock.Anything, mock.Anything).Return(v, nil)
			}

			// CCIP Reader mocks:
			// once:
			//      CommitReportsGTETimestamp(ctx, dest, ts, 1000) -> ([]cciptypes.CommitPluginReportWithMeta, error)
			// for each chain selector:
			//      ExecutedMessages(ctx, selector, dest, seqRange) -> ([]cciptypes.SeqNum, error)
			got, got2, err := getPendingExecutedReports(
				tests.Context(t),
				mockReader,
				tt.canExec,
				time.Now(),
				logger.Test(t),
			)
			if !tt.wantErr(t, err, "getPendingExecutedReports(...)") {
				return
			}
			assert.Equalf(t, tt.wantObs, got, "getPendingExecutedReports(...)")
			assert.Equalf(t, tt.wantExecuted, got2, "getPendingExecutedReports(...)")
		})
	}
}

func TestPlugin_Close(t *testing.T) {
	p := &Plugin{tokenDataObserver: &tokendata.NoopTokenDataObserver{}}
	require.NoError(t, p.Close())
}

func TestPlugin_Query(t *testing.T) {
	p := &Plugin{}
	q, err := p.Query(tests.Context(t), ocr3types.OutcomeContext{})
	require.NoError(t, err)
	require.Equal(t, types.Query{}, q)
}

func TestPlugin_ObservationQuorum(t *testing.T) {
	ctx := tests.Context(t)
	p := &Plugin{
		reportingCfg: ocr3types.ReportingPluginConfig{F: 1},
	}
	got, err := p.ObservationQuorum(ctx, ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{Observation: []byte{}},
		{Observation: []byte{}},
	})
	require.NoError(t, err)
	assert.Equal(t, true, got)
}

func TestPlugin_ValidateObservation_NonDecodable(t *testing.T) {
	ctx := tests.Context(t)
	p := &Plugin{ocrTypeCodec: jsonOcrTypeCodec}
	err := p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte("not a valid observation"),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observation")
}

func TestPlugin_ValidateObservation_SupportedChainsError(t *testing.T) {
	ctx := tests.Context(t)
	p := &Plugin{
		ocrTypeCodec: jsonOcrTypeCodec,
	}
	err := p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte(`{"oracleID": "0xdeadbeef"}`),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error finding supported chains by node: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_ValidateObservation_IneligibleMessageObserver(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet[cciptypes.ChainSelector](), nil)
	defer mockHomeChain.AssertExpectations(t)

	p := &Plugin{
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		ocrTypeCodec: jsonOcrTypeCodec,
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
	}, nil, nil, nil, dt.Observation{}, nil)
	encoded, err := jsonOcrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)

	prevOutcome := exectypes.Outcome{
		State: exectypes.GetCommitReports,
	}
	encodedPrevOutcome, err := jsonOcrTypeCodec.EncodeOutcome(prevOutcome)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{PreviousOutcome: encodedPrevOutcome}, types.Query{},
		types.AttributedObservation{
			Observation: encoded,
		})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observer reading eligibility: observer not allowed to read from chain 0")
}

func TestPlugin_ValidateObservation_IneligibleCommitReportsObserver(t *testing.T) {
	ctx := tests.Context(t)
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
		ocrTypeCodec: jsonOcrTypeCodec,
	}

	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{
				MerkleRoot:          cciptypes.Bytes32{},
				SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2),
				SourceChain:         1,
			},
		},
	}
	observation := exectypes.NewObservation(
		commitReports, nil, nil, nil, nil, dt.Observation{}, nil,
	)
	encoded, err := jsonOcrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t,
		err.Error(),
		"validate commit reports reading eligibility: observer not allowed to read from chain 1")
}

func TestPlugin_ValidateObservation_ValidateObservedSeqNum_Error(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		ocrTypeCodec: jsonOcrTypeCodec,
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
		commitReports, nil, nil, nil, nil, dt.Observation{}, nil,
	)
	encoded, err := jsonOcrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observed sequence numbers: duplicate merkle root")
}

func TestPlugin_ValidateObservation_CallsDiscoveryValidateObservation(t *testing.T) {
	ctx := tests.Context(t)
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
		ocrTypeCodec: jsonOcrTypeCodec,
	}

	// Reports with duplicate roots.
	root := cciptypes.Bytes32{}
	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{MerkleRoot: root, SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2), SourceChain: 1},
		},
	}
	observation := exectypes.NewObservation(
		commitReports, nil, nil, nil, nil, dt.Observation{}, nil,
	)
	encoded, err := jsonOcrTypeCodec.EncodeObservation(observation)
	require.NoError(t, err)
	err = p.ValidateObservation(ctx, ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.NoError(t, err)
}

func TestPlugin_Observation_BadPreviousOutcome(t *testing.T) {
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
	}
	_, err := p.Observation(tests.Context(t), ocr3types.OutcomeContext{
		PreviousOutcome: []byte("not a valid observation"),
	}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode previous outcome: invalid character")
}

func TestPlugin_Observation_EligibilityCheckFailure(t *testing.T) {
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetFChain().Return(map[cciptypes.ChainSelector]int{}, nil)

	p := &Plugin{
		homeChain:       mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{},
		lggr:            lggr,
		ocrTypeCodec:    jsonOcrTypeCodec,
	}

	_, err := p.Observation(tests.Context(t), ocr3types.OutcomeContext{}, nil)
	require.Error(t, err)
	//nolint:lll // error message
	assert.Contains(t, err.Error(), "unable to determine if the destination chain is supported: error getting supported chains: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_Outcome_DestFChainNotAvailable(t *testing.T) {
	ctx := tests.Context(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		1: 1,
		2: 2,
	}

	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
	}
	observation, err := jsonOcrTypeCodec.EncodeObservation(exectypes.Observation{
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
	ctx := tests.Context(t)
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
	}
	_, err := p.Outcome(ctx, ocr3types.OutcomeContext{}, nil,
		[]types.AttributedObservation{
			{
				Observation: []byte("not a valid observation"),
				Observer:    0,
			},
		})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observations: invalid character")
}

func TestPlugin_Outcome_BelowF(t *testing.T) {
	ctx := tests.Context(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		0: 1,
		2: 2,
	}
	p := &Plugin{
		reportingCfg: ocr3types.ReportingPluginConfig{
			F: 1,
		},
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
	}
	observation, err := jsonOcrTypeCodec.EncodeObservation(exectypes.Observation{
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
	ctx := tests.Context(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		10: 20,
		0:  3,
	}

	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
	}

	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {},
	}
	observation, err := jsonOcrTypeCodec.EncodeObservation(exectypes.Observation{
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
	ctx := tests.Context(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		0:  3,
		10: 20,
	}

	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
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
	observation, err := jsonOcrTypeCodec.EncodeObservation(exectypes.Observation{
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

func TestPlugin_Reports_UnableToParse(t *testing.T) {
	ctx := tests.Context(t)
	p := &Plugin{
		lggr:         logger.Test(t),
		ocrTypeCodec: jsonOcrTypeCodec,
	}
	_, err := p.Reports(ctx, 0, ocr3types.Outcome("not a valid observation"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode outcome")
}

func TestPlugin_Reports_UnableToEncode(t *testing.T) {
	ctx := tests.Context(t)
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Encode", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("test error"))
	p := &Plugin{reportCodec: codec, lggr: logger.Test(t), ocrTypeCodec: jsonOcrTypeCodec}
	report, err := jsonOcrTypeCodec.EncodeOutcome(exectypes.NewOutcome(
		exectypes.Unknown, nil, cciptypes.ExecutePluginReport{}))
	require.NoError(t, err)

	_, err = p.Reports(ctx, 0, report)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to encode report: test error")
}

func TestPlugin_ShouldAcceptAttestedReport_DoesNotDecode(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, fmt.Errorf("test error"))

	p := &Plugin{
		reportCodec: codec,
		lggr:        logger.Test(t),
	}

	_, err := p.ShouldAcceptAttestedReport(tests.Context(t), 0, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("will not decode"), // faked out, see mock above
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate exec report: decode exec plugin report: test error")
}

func TestPlugin_ShouldAcceptAttestedReport_NoReports(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.EXPECT().Decode(mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, nil)

	p := &Plugin{
		lggr:        logger.Test(t),
		reportCodec: codec,
	}
	result, err := p.ShouldAcceptAttestedReport(tests.Context(t), 0, ocr3types.ReportWithInfo[[]byte]{
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
		ccipReaderMock.
			EXPECT().
			ExecutedMessages(
				mock.Anything,
				cciptypes.ChainSelector(sourceChain),
				cciptypes.NewSeqNumRange(seqNum, seqNum),
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
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything, mock.Anything).
					Return(&reader2.CurseInfo{
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
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything, mock.Anything).
					Return(&reader2.CurseInfo{}, fmt.Errorf("test error"))

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
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything, mock.Anything).
					Return(&reader2.CurseInfo{GlobalCurse: true}, nil)

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
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything, mock.Anything).
					Return(&reader2.CurseInfo{CursedDestination: true}, nil)

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
				mockReader.EXPECT().GetRmnCurseInfo(mock.Anything, mock.Anything).
					Return(&reader2.CurseInfo{CursedSourceChains: map[cciptypes.ChainSelector]bool{
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
						cciptypes.ChainSelector(sourceChain),
						cciptypes.NewSeqNumRange(seqNum, seqNum),
					).Return(
					[]cciptypes.SeqNum{seqNum},
					nil)

				homeChain := basicHomeChain()
				codec := basicMockCodec()
				return codec, homeChain, mockReader
			},
			assertions: func(t *testing.T, b bool, err error) {
				require.NoError(t, err) // error is logged, but not returned
				require.False(t, b)
			},
			logsContain: []string{"some messages in report already executed"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			codec, homeChain, ccipReader := tc.getDeps()
			lggr, obs := logger.TestObserved(t, zapcore.DebugLevel)
			p := &Plugin{
				lggr:        lggr,
				reportCodec: codec,
				homeChain:   homeChain,
				ccipReader:  ccipReader,
				chainSupport: plugincommon.NewChainSupport(
					logger.Test(t),
					homeChain,
					map[commontypes.OracleID]libocrtypes.PeerID{
						1: [32]byte{1},
					},
					1,
					cciptypes.ChainSelector(destChain),
				),
				reportingCfg: ocr3types.ReportingPluginConfig{
					OracleID:     1,
					ConfigDigest: configDigest,
				},
			}

			result, err := p.ShouldAcceptAttestedReport(tests.Context(t), 0, ocr3types.ReportWithInfo[[]byte]{
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

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{})
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

	_, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("will not decode"), // faked out, see mock above
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "validate exec report: decode exec plugin report: test error")
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

	_, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{
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

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{
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

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{
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
			reports[0].SourceChainSelector,
			cciptypes.NewSeqNumRange(
				reports[0].Messages[0].Header.SequenceNumber,
				reports[0].Messages[0].Header.SequenceNumber,
			),
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

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{
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
			reports[0].SourceChainSelector,
			cciptypes.NewSeqNumRange(
				reports[0].Messages[0].Header.SequenceNumber,
				reports[0].Messages[0].Header.SequenceNumber,
			),
		).Return([]cciptypes.SeqNum{
		reports[0].Messages[0].Header.SequenceNumber,
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

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(tests.Context(t), 1, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.False(t, shouldTransmit)
}
