package merkleroot

import (
	"fmt"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmn2 "github.com/smartcontractkit/chainlink-ccip/mocks/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func TestProcessor_Query(t *testing.T) {
	srcChain1 := ccipocr3.ChainSelector(1)
	srcChain2 := ccipocr3.ChainSelector(2)
	dstChain := ccipocr3.ChainSelector(3)
	ctx := tests.Context(t)

	contractAddrs := map[ccipocr3.ChainSelector]map[string][]byte{
		srcChain1: {consts.ContractNameOnRamp: []byte("0x1234567890123456789012345678901234567890")},
		srcChain2: {consts.ContractNameOnRamp: []byte("0x1234567890123456789012345678901234567891")},
		dstChain:  {consts.ContractNameOffRamp: []byte("0x1234567890123456789012345678901234567892")},
	}

	expSigs1 := &rmn.ReportSignatures{
<<<<<<< Updated upstream
		Signatures: []rmn.ECDSASignature{
			{R: []byte("r1"), S: []byte("s1")},
			{R: []byte("r2"), S: []byte("s2")},
		},
		LaneUpdates: []rmn.FixedDestLaneUpdate{
			{SourceChain: rmn.SourceChainInfo{Chain: srcChain1}},
			{SourceChain: rmn.SourceChainInfo{Chain: srcChain2}},
		},
=======
		Signatures: []*rmnpb.EcdsaSignature{
			{R: []byte("r1"), S: []byte("s1")},
			{R: []byte("r2"), S: []byte("s2")},
		}
>>>>>>> Stashed changes
	}

	testCases := []struct {
		name              string
		prevOutcome       Outcome
		contractAddresses map[ccipocr3.ChainSelector]map[string][]byte
		cfg               pluginconfig.CommitPluginConfig
		rmnClient         func(t *testing.T) *rmn2.MockClient
		expQuery          Query
		expErr            bool
	}{
		{
			name: "happy path",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: srcChain1, SeqNumRange: ccipocr3.NewSeqNumRange(10, 20)},
					{ChainSel: srcChain2, SeqNumRange: ccipocr3.NewSeqNumRange(50, 51)},
				},
			},
			contractAddresses: contractAddrs,
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           true,
				RMNSignaturesTimeout: 5 * time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmn2.MockClient {
				cl := rmn2.NewMockClient(t)
				cl.EXPECT().
					ComputeReportSignatures(
						mock.Anything,
						rmn.DestChainInfo{
							Chain:          dstChain,
							OffRampAddress: contractAddrs[dstChain][consts.ContractNameOffRamp],
						},
						[]rmn.FixedDestLaneUpdateRequest{
							{
								SourceChainInfo: rmn.SourceChainInfo{
									Chain:         srcChain1,
									OnRampAddress: contractAddrs[srcChain1][consts.ContractNameOnRamp],
								},
								SeqNumRange: ccipocr3.NewSeqNumRange(10, 20),
							},
							{
								SourceChainInfo: rmn.SourceChainInfo{
									Chain:         srcChain2,
									OnRampAddress: contractAddrs[srcChain2][consts.ContractNameOnRamp],
								},
								SeqNumRange: ccipocr3.NewSeqNumRange(50, 51),
							},
						},
					).
					Return(expSigs1, nil)
				return cl
			},
			expQuery: Query{
				RetryRMNSignatures: false,
				RMNSignatures:      expSigs1,
			},
			expErr: false,
		},
		{
			name: "rmn timeout",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: srcChain1, SeqNumRange: ccipocr3.NewSeqNumRange(10, 20)},
					{ChainSel: srcChain2, SeqNumRange: ccipocr3.NewSeqNumRange(50, 51)},
				},
			},
			contractAddresses: contractAddrs,
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           true,
				RMNSignaturesTimeout: time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmn2.MockClient {
				cl := rmn2.NewMockClient(t)
				time.Sleep(time.Millisecond)
				cl.EXPECT().ComputeReportSignatures(mock.Anything, mock.Anything, mock.Anything).
					Return(expSigs1, rmn.ErrTimeout) // <------------------------------------ timeout error
				return cl
			},
			expQuery: Query{
				RetryRMNSignatures: true,
				RMNSignatures:      nil,
			},
			expErr: false,
		},
		{
			name: "rmn random error",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: srcChain1, SeqNumRange: ccipocr3.NewSeqNumRange(10, 20)},
					{ChainSel: srcChain2, SeqNumRange: ccipocr3.NewSeqNumRange(50, 51)},
				},
			},
			contractAddresses: contractAddrs,
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           true,
				RMNSignaturesTimeout: time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmn2.MockClient {
				cl := rmn2.NewMockClient(t)
				time.Sleep(time.Millisecond)
				cl.EXPECT().ComputeReportSignatures(mock.Anything, mock.Anything, mock.Anything).
					Return(expSigs1, fmt.Errorf("some error")) // <------------------------- some random error
				return cl
			},
			expQuery: Query{},
			expErr:   true,
		},
		{
			name: "not in building reports state",
			prevOutcome: Outcome{
				OutcomeType: ReportInFlight,
			},
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           true,
				RMNSignaturesTimeout: 5 * time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmn2.MockClient { return rmn2.NewMockClient(t) },
			expQuery:  Query{},
			expErr:    false,
		},
		{
			name: "rmn sig checks disabled",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
			},
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           false, // <-------------- disabled
				RMNSignaturesTimeout: 5 * time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmn2.MockClient { return rmn2.NewMockClient(t) },
			expQuery:  Query{},
			expErr:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ccipReader := reader.NewMockCCIP(t)
			for chainSel, contracts := range tc.contractAddresses {
				for name, addr := range contracts {
					ccipReader.EXPECT().GetContractAddress(name, chainSel).Return(addr, nil)
				}
			}

			w := Processor{
				cfg:        tc.cfg,
				ccipReader: ccipReader,
				rmnClient:  tc.rmnClient(t),
				lggr:       logger.Test(t),
			}

			q, err := w.Query(ctx, tc.prevOutcome)
			if tc.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expQuery, q)
		})
	}
}
