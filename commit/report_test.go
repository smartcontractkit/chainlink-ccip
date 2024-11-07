package commit

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestPluginReports(t *testing.T) {
	testCases := []struct {
		name          string
		outc          Outcome
		expErr        bool
		expReports    []ccipocr3.CommitPluginReport
		expReportInfo ReportInfo
	}{
		{
			name: "wrong outcome type gives an empty report but no error",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportIntervalsSelected,
				},
			},
			expErr: false,
		},
		{
			name: "correct outcome type but empty data",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportGenerated,
				},
			},
			expErr: false,
		},
		{
			name: "token prices reported but no merkle roots so report is empty",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportTransmissionFailed,
				},
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: []ccipocr3.TokenPrice{
						{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
					},
				},
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expErr: false,
		},
		{
			name: "token prices reported but no merkle roots so report is empty",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportGenerated,
					RootsToReport: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      3,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
						},
					},
					RMNRemoteCfg: rmntypes.RemoteConfig{F: 123},
				},
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: []ccipocr3.TokenPrice{
						{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
					},
				},
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expReports: []ccipocr3.CommitPluginReport{
				{
					MerkleRoots: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      3,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
						},
					},
					PriceUpdates: ccipocr3.PriceUpdates{
						TokenPriceUpdates: []ccipocr3.TokenPrice{
							{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
						},
						GasPriceUpdates: []ccipocr3.GasPriceChain{
							{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
						},
					},
					RMNSignatures: nil,
				},
			},
			expReportInfo: ReportInfo{RemoteF: 123},
		},
	}

	ctx := tests.Context(t)
	lggr := logger.Test(t)
	reportCodec := mocks.NewCommitPluginJSONReportCodec()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs := plugincommon.NewMockChainSupport(t)
			p := Plugin{
				lggr:            lggr,
				reportCodec:     reportCodec,
				oracleIDToP2PID: map[commontypes.OracleID]libocrtypes.PeerID{1: {1}},
				chainSupport:    cs,
			}
			cs.EXPECT().SupportsDestChain(commontypes.OracleID(1)).Return(true, nil).Maybe()

			outcomeBytes, err := tc.outc.Encode()
			require.NoError(t, err)

			reports, err := p.Reports(ctx, 0, outcomeBytes)
			if tc.expErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Len(t, reports, len(tc.expReports))
			for i, report := range reports {
				expEncodedReport, err := reportCodec.Encode(ctx, tc.expReports[i])
				require.NoError(t, err)
				require.Equal(t, expEncodedReport, []byte(report.ReportWithInfo.Report))

				expReportInfoBytes, err := tc.expReportInfo.Encode()
				require.NoError(t, err)
				require.Equal(t, expReportInfoBytes, report.ReportWithInfo.Info)
			}
		})
	}
}

func TestPluginReports_InvalidOutcome(t *testing.T) {
	lggr := logger.Test(t)
	p := Plugin{lggr: lggr}
	_, err := p.Reports(tests.Context(t), 0, []byte("invalid json"))
	require.Error(t, err)
}

func Test_IsCandidateCheck(t *testing.T) {
	rb := rand.RandomBytes32()
	digest := types.ConfigDigest(rb[:])
	donID := uint32(3)
	allTests := []struct {
		name          string
		makePlugin    func(t *testing.T, hc *reader_mock.MockHomeChain) *Plugin
		makeHomeChain func(t *testing.T) *reader_mock.MockHomeChain
		wantOutput    bool
		wantError     bool
	}{
		{
			name: "Should return true if digest matches",
			makePlugin: func(t *testing.T, hc *reader_mock.MockHomeChain) *Plugin {
				p := &Plugin{
					homeChain: hc,
					reportingCfg: ocr3types.ReportingPluginConfig{
						ConfigDigest: digest,
					},
				}
				return p
			},
			makeHomeChain: func(t *testing.T) *reader_mock.MockHomeChain {
				h := reader_mock.NewMockHomeChain(t)
				h.On("GetOCRConfigs", mock.Anything, mock.Anything, consts.PluginTypeCommit).
					Return(reader.ActiveAndCandidate{
						ActiveConfig: reader.OCR3ConfigWithMeta{},
						CandidateConfig: reader.OCR3ConfigWithMeta{
							ConfigDigest: digest,
						},
					}, nil)
				return h
			},
			wantOutput: true,
			wantError:  false,
		},
		{
			name: "Should return false if digest doesn't match",
			makePlugin: func(t *testing.T, hc *reader_mock.MockHomeChain) *Plugin {
				p := &Plugin{
					homeChain: hc,
					reportingCfg: ocr3types.ReportingPluginConfig{
						ConfigDigest: types.ConfigDigest(rand.RandomBytes32()),
					},
				}
				return p
			},
			makeHomeChain: func(t *testing.T) *reader_mock.MockHomeChain {
				h := reader_mock.NewMockHomeChain(t)
				h.On("GetOCRConfigs", mock.Anything, mock.Anything, consts.PluginTypeCommit).
					Return(reader.ActiveAndCandidate{
						ActiveConfig: reader.OCR3ConfigWithMeta{},
						CandidateConfig: reader.OCR3ConfigWithMeta{
							ConfigDigest: digest,
						},
					}, nil)
				return h
			},
			wantOutput: false,
			wantError:  false,
		},
		{
			name: "Should work as expected without candidate instance",
			makePlugin: func(t *testing.T, hc *reader_mock.MockHomeChain) *Plugin {
				p := &Plugin{
					homeChain: hc,
					reportingCfg: ocr3types.ReportingPluginConfig{
						ConfigDigest: types.ConfigDigest(rand.RandomBytes32()),
					},
				}
				return p
			},
			makeHomeChain: func(t *testing.T) *reader_mock.MockHomeChain {
				h := reader_mock.NewMockHomeChain(t)
				h.On("GetOCRConfigs", mock.Anything, mock.Anything, consts.PluginTypeCommit).
					Return(reader.ActiveAndCandidate{
						ActiveConfig:    reader.OCR3ConfigWithMeta{},
						CandidateConfig: reader.OCR3ConfigWithMeta{},
					}, nil)
				return h
			},
			wantOutput: false,
			wantError:  false,
		},
		{
			name: "Should throw error if donID doesn't exist",
			makePlugin: func(t *testing.T, hc *reader_mock.MockHomeChain) *Plugin {
				p := &Plugin{
					homeChain: hc,
					donID:     donID,
					reportingCfg: ocr3types.ReportingPluginConfig{
						ConfigDigest: types.ConfigDigest(rand.RandomBytes32()),
					},
				}
				return p
			},
			makeHomeChain: func(t *testing.T) *reader_mock.MockHomeChain {
				h := reader_mock.NewMockHomeChain(t)
				h.On("GetOCRConfigs", mock.Anything, donID, consts.PluginTypeCommit).
					Return(reader.ActiveAndCandidate{
						ActiveConfig:    reader.OCR3ConfigWithMeta{},
						CandidateConfig: reader.OCR3ConfigWithMeta{},
					}, fmt.Errorf("DonID does not exist"))
				return h
			},
			wantOutput: false,
			wantError:  true,
		},
	}
	for _, tt := range allTests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tests.Context(t)
			hc := tt.makeHomeChain(t)
			p := tt.makePlugin(t, hc)
			actualOutput, actualError := p.isCandidateInstance(ctx)
			assert.Equal(t, tt.wantOutput, actualOutput)
			if tt.wantError {
				require.Error(t, actualError)
			} else {
				require.NoError(t, actualError)
			}
		})
	}
}
