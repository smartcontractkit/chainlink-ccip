package execute

import (
	"context"
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	codec_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/execute/internal_/gen"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func Test_getPendingExecutedReports(t *testing.T) {
	tests := []struct {
		name    string
		reports []plugintypes2.CommitPluginReportWithMeta
		ranges  map[cciptypes.ChainSelector][]cciptypes.SeqNumRange
		want    exectypes.CommitObservations
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			reports: nil,
			ranges:  nil,
			want:    exectypes.CommitObservations{},
			wantErr: assert.NoError,
		},
		{
			name: "single non-executed report",
			reports: []plugintypes2.CommitPluginReportWithMeta{
				{
					BlockNum:  999,
					Timestamp: time.UnixMilli(10101010101),
					Report: cciptypes.CommitPluginReport{
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				1: nil,
			},
			want: exectypes.CommitObservations{
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
						MerkleRoots: []cciptypes.MerkleRootChain{
							{
								ChainSel:     1,
								SeqNumsRange: cciptypes.NewSeqNumRange(1, 10),
							},
						},
					},
				},
			},
			ranges: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
				1: {
					cciptypes.NewSeqNumRange(1, 3),
					cciptypes.NewSeqNumRange(7, 8),
				},
			},
			want: exectypes.CommitObservations{
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
			ranges:  map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{},
			want:    exectypes.CommitObservations{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReader := readerpkg_mock.NewMockCCIPReader(t)
			mockReader.On(
				"CommitReportsGTETimestamp", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			).Return(tt.reports, nil)
			for k, v := range tt.ranges {
				mockReader.On("ExecutedMessageRanges", mock.Anything, k, mock.Anything, mock.Anything).Return(v, nil)
			}

			// CCIP Reader mocks:
			// once:
			//      CommitReportsGTETimestamp(ctx, dest, ts, 1000) -> ([]cciptypes.CommitPluginReportWithMeta, error)
			// for each chain selector:
			//      ExecutedMessageRanges(ctx, selector, dest, seqRange) -> ([]cciptypes.SeqNumRange, error)
			got, err := getPendingExecutedReports(
				context.Background(),
				mockReader,
				123,
				time.Now(),
				logger.Test(t),
			)
			if !tt.wantErr(t, err, "getPendingExecutedReports(...)") {
				return
			}
			assert.Equalf(t, tt.want, got, "getPendingExecutedReports(...)")
		})
	}
}

func TestPlugin_Close(t *testing.T) {
	mockReader := readerpkg_mock.NewMockCCIPReader(t)
	mockReader.On("Close", mock.Anything).Return(nil)

	lggr := logger.Test(t)
	p := &Plugin{lggr: lggr, ccipReader: mockReader}
	require.NoError(t, p.Close())
}

func TestPlugin_Query(t *testing.T) {
	p := &Plugin{}
	q, err := p.Query(context.Background(), ocr3types.OutcomeContext{})
	require.NoError(t, err)
	require.Equal(t, types.Query{}, q)
}

func TestPlugin_ObservationQuorum(t *testing.T) {
	p := &Plugin{}
	got, err := p.ObservationQuorum(ocr3types.OutcomeContext{}, nil)
	require.NoError(t, err)
	assert.Equal(t, ocr3types.QuorumFPlusOne, got)
}

func TestPlugin_ValidateObservation_NonDecodable(t *testing.T) {
	p := &Plugin{}
	err := p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte("not a valid observation"),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode observation")
}

func TestPlugin_ValidateObservation_SupportedChainsError(t *testing.T) {
	p := &Plugin{}
	err := p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: []byte(`{"oracleID": "0xdeadbeef"}`),
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error finding supported chains by node: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_ValidateObservation_IneligibleObserver(t *testing.T) {
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet[cciptypes.ChainSelector](), nil)
	defer mockHomeChain.AssertExpectations(t)

	p := &Plugin{
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
		lggr: lggr,
	}

	observation := exectypes.NewObservation(nil, exectypes.MessageObservations{
		0: map[cciptypes.SeqNum]cciptypes.Message{
			1: {
				Header: cciptypes.RampMessageHeader{
					SourceChainSelector: 1,
				},
			},
		},
	}, nil, nil, dt.Observation{})
	encoded, err := observation.Encode()
	require.NoError(t, err)
	err = p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observer reading eligibility: observer not allowed to read from chain 0")
}

func TestPlugin_ValidateObservation_ValidateObservedSeqNum_Error(t *testing.T) {
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(0)), nil)

	p := &Plugin{
		lggr:      lggr,
		homeChain: mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			0: {},
		},
	}

	// Reports with duplicate roots.
	root := cciptypes.Bytes32{}
	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{MerkleRoot: root},
			{MerkleRoot: root},
		},
	}
	observation := exectypes.NewObservation(commitReports, nil, nil, nil, dt.Observation{})
	encoded, err := observation.Encode()
	require.NoError(t, err)
	err = p.ValidateObservation(ocr3types.OutcomeContext{}, types.Query{}, types.AttributedObservation{
		Observation: encoded,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "validate observed sequence numbers: duplicate merkle root")
}

func TestPlugin_Observation_BadPreviousOutcome(t *testing.T) {
	p := &Plugin{}
	_, err := p.Observation(context.Background(), ocr3types.OutcomeContext{
		PreviousOutcome: []byte("not a valid observation"),
	}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode previous outcome: invalid character")
}

func TestPlugin_Observation_EligibilityCheckFailure(t *testing.T) {
	lggr := logger.Test(t)

	mockHomeChain := reader_mock.NewMockHomeChain(t)

	p := &Plugin{
		homeChain:       mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{},
		lggr:            lggr,
	}

	_, err := p.Observation(context.Background(), ocr3types.OutcomeContext{}, nil)
	require.Error(t, err)
	// nolint:lll // error message
	assert.Contains(t, err.Error(), "unable to determine if the destination chain is supported: error getting supported chains: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_Outcome_BadObservationEncoding(t *testing.T) {
	p := &Plugin{lggr: logger.Test(t)}
	_, err := p.Outcome(ocr3types.OutcomeContext{}, nil,
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
	homeChain := reader_mock.NewMockHomeChain(t)
	homeChain.EXPECT().GetFChain().Return(nil, nil)
	p := &Plugin{
		homeChain: homeChain,
		reportingCfg: ocr3types.ReportingPluginConfig{
			F: 1,
		},
		lggr: logger.Test(t),
	}
	_, err := p.Outcome(ocr3types.OutcomeContext{}, nil,
		[]types.AttributedObservation{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "below F threshold")
}

func TestPlugin_Outcome_HomeChainError(t *testing.T) {
	homeChain := reader_mock.NewMockHomeChain(t)
	homeChain.On("GetFChain", mock.Anything).Return(nil, fmt.Errorf("test error"))

	p := &Plugin{
		homeChain: homeChain,
		lggr:      logger.Test(t),
	}
	_, err := p.Outcome(ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to get FChain: test error")
}

func TestPlugin_Outcome_CommitReportsMergeError(t *testing.T) {
	homeChain := reader_mock.NewMockHomeChain(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		10: 20,
	}
	homeChain.On("GetFChain", mock.Anything).Return(fChainMap, nil)

	p := &Plugin{
		homeChain: homeChain,
		lggr:      logger.Test(t),
	}

	commitReports := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {},
	}
	observation, err := exectypes.NewObservation(commitReports, nil, nil, nil, dt.Observation{}).Encode()
	require.NoError(t, err)
	_, err = p.Outcome(ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to merge commit report observations: no validator")
}

func TestPlugin_Outcome_MessagesMergeError(t *testing.T) {
	homeChain := reader_mock.NewMockHomeChain(t)
	fChainMap := map[cciptypes.ChainSelector]int{
		10: 20,
	}
	homeChain.On("GetFChain", mock.Anything).Return(fChainMap, nil)

	p := &Plugin{
		homeChain: homeChain,
		lggr:      logger.Test(t),
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
	observation, err := exectypes.NewObservation(nil, messages, nil, nil, dt.Observation{}).Encode()
	require.NoError(t, err)
	_, err = p.Outcome(ocr3types.OutcomeContext{}, nil, []types.AttributedObservation{
		{
			Observation: observation,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to merge message observations: no validator")
}

func TestPlugin_Reports_UnableToParse(t *testing.T) {
	p := &Plugin{}
	_, err := p.Reports(0, ocr3types.Outcome("not a valid observation"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to decode outcome")
}

func TestPlugin_Reports_UnableToEncode(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Encode", mock.Anything, mock.Anything).
		Return(nil, fmt.Errorf("test error"))
	p := &Plugin{reportCodec: codec}
	report, err := exectypes.NewOutcome(exectypes.Unknown, nil, cciptypes.ExecutePluginReport{}).Encode()
	require.NoError(t, err)

	_, err = p.Reports(0, report)
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
	_, err := p.ShouldAcceptAttestedReport(context.Background(), 0, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("will not decode"), // faked out, see mock above
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "decode commit plugin report: test error")
}

func TestPlugin_ShouldAcceptAttestedReport_NoReports(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, nil)
	p := &Plugin{
		lggr:        logger.Test(t),
		reportCodec: codec,
	}
	result, err := p.ShouldAcceptAttestedReport(context.Background(), 0, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("empty report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.False(t, result)
}

func TestPlugin_ShouldAcceptAttestedReport_ShouldAccept(t *testing.T) {
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{},
			},
		}, nil)
	p := &Plugin{
		lggr:        logger.Test(t),
		reportCodec: codec,
	}
	result, err := p.ShouldAcceptAttestedReport(context.Background(), 0, ocr3types.ReportWithInfo[[]byte]{
		Report: []byte("report"), // faked out, see mock above
	})
	require.NoError(t, err)
	require.True(t, result)
}

func TestPlugin_ShouldTransmitAcceptReport_ElegibilityCheckFailure(t *testing.T) {
	lggr := logger.Test(t)

	p := &Plugin{
		lggr:            lggr,
		homeChain:       reader_mock.NewMockHomeChain(t),
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{},
	}

	_, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.Error(t, err)
	// nolint:lll // error message
	assert.Contains(t, err.Error(), "unable to determine if the destination chain is supported: error getting supported chains: oracle ID 0 not found in oracleIDToP2pID")
}

func TestPlugin_ShouldTransmitAcceptReport_Ineligible(t *testing.T) {
	lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)

	mockHomeChain := reader_mock.NewMockHomeChain(t)
	mockHomeChain.EXPECT().GetSupportedChainsForPeer(mock.Anything).Return(mapset.NewSet[cciptypes.ChainSelector](), nil)
	defer mockHomeChain.AssertExpectations(t)

	p := &Plugin{
		lggr:         lggr,
		cfg:          pluginconfig.ExecutePluginConfig{DestChain: 1},
		reportingCfg: ocr3types.ReportingPluginConfig{OracleID: 2},
		homeChain:    mockHomeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			2: {},
		},
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.False(t, shouldTransmit)

	messages := slicelib.Map(logs.All(), func(e observer.LoggedEntry) string {
		return e.Message
	})
	require.ElementsMatch(t, messages, []string{"not a destination writer, skipping report transmission"})
}

func TestPlugin_ShouldTransmitAcceptReport_DecodeFailure(t *testing.T) {
	const donID = uint32(1)
	homeChain := reader_mock.NewMockHomeChain(t)
	homeChain.On("GetSupportedChainsForPeer", mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)
	homeChain.
		EXPECT().
		GetOCRConfigs(mock.Anything, donID, consts.PluginTypeExecute).
		Return([]reader.OCR3ConfigWithMeta{{}}, nil)
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, fmt.Errorf("test error"))

	p := &Plugin{
		donID:        donID,
		lggr:         logger.Test(t),
		cfg:          pluginconfig.ExecutePluginConfig{DestChain: 1},
		reportingCfg: ocr3types.ReportingPluginConfig{OracleID: 2},
		reportCodec:  codec,
		homeChain:    homeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			2: {1},
		},
	}

	_, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.Error(t, err)
	require.ErrorContains(t, err, "decode commit plugin report: test error")
}

func TestPlugin_ShouldTransmitAcceptReport_Success(t *testing.T) {
	const donID = uint32(1)
	lggr, logs := logger.TestObserved(t, zapcore.DebugLevel)
	homeChain := reader_mock.NewMockHomeChain(t)
	homeChain.On("GetSupportedChainsForPeer", mock.Anything).Return(mapset.NewSet(cciptypes.ChainSelector(1)), nil)
	homeChain.
		EXPECT().
		GetOCRConfigs(mock.Anything, donID, consts.PluginTypeExecute).
		Return([]reader.OCR3ConfigWithMeta{{}}, nil)
	codec := codec_mocks.NewMockExecutePluginCodec(t)
	codec.On("Decode", mock.Anything, mock.Anything).
		Return(cciptypes.ExecutePluginReport{}, nil)

	p := &Plugin{
		donID:        donID,
		lggr:         lggr,
		cfg:          pluginconfig.ExecutePluginConfig{DestChain: 1},
		reportingCfg: ocr3types.ReportingPluginConfig{OracleID: 2},
		reportCodec:  codec,
		homeChain:    homeChain,
		oracleIDToP2pID: map[commontypes.OracleID]libocrtypes.PeerID{
			2: {1},
		},
	}

	shouldTransmit, err := p.ShouldTransmitAcceptedReport(context.Background(), 1, ocr3types.ReportWithInfo[[]byte]{})
	require.NoError(t, err)
	require.True(t, shouldTransmit)

	messages := slicelib.Map(logs.All(), func(e observer.LoggedEntry) string {
		return e.Message
	})
	require.ElementsMatch(t, messages, []string{"transmitting report"})
}
