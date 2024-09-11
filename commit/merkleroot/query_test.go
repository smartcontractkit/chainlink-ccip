package merkleroot

import (
	"fmt"
	"testing"
	"time"

	ct "github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
	rmnmocks "github.com/smartcontractkit/chainlink-ccip/mocks/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
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
		Signatures: []*rmnpb.EcdsaSignature{
			{R: []byte("r1"), S: []byte("s1")},
			{R: []byte("r2"), S: []byte("s2")},
		},
		LaneUpdates: []*rmnpb.FixedDestLaneUpdate{
			{
				LaneSource: &rmnpb.LaneSource{
					SourceChainSelector: uint64(srcChain1),
					OnrampAddress:       contractAddrs[srcChain1][consts.ContractNameOnRamp],
				},
			},
			{
				LaneSource: &rmnpb.LaneSource{
					SourceChainSelector: uint64(srcChain2),
					OnrampAddress:       contractAddrs[srcChain2][consts.ContractNameOnRamp],
				},
			},
		},
	}

	testCases := []struct {
		name              string
		prevOutcome       ct.MerkleRootOutcome
		contractAddresses map[ccipocr3.ChainSelector]map[string][]byte
		cfg               pluginconfig.CommitPluginConfig
		rmnClient         func(t *testing.T) *rmnmocks.MockClient
		expQuery          ct.MerkleRootQuery
		expErr            bool
	}{
		{
			name: "happy path",
			prevOutcome: ct.MerkleRootOutcome{
				OutcomeType: ct.ReportIntervalsSelected,
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
			rmnClient: func(t *testing.T) *rmnmocks.MockClient {
				cl := rmnmocks.NewMockClient(t)
				cl.EXPECT().
					ComputeReportSignatures(
						mock.Anything,
						&rmnpb.LaneDest{
							DestChainSelector: uint64(dstChain),
							OfframpAddress:    contractAddrs[dstChain][consts.ContractNameOffRamp],
						},
						[]*rmnpb.FixedDestLaneUpdateRequest{
							{
								LaneSource: &rmnpb.LaneSource{
									SourceChainSelector: uint64(srcChain1),
									OnrampAddress:       contractAddrs[srcChain1][consts.ContractNameOnRamp],
								},
								ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
							},
							{
								LaneSource: &rmnpb.LaneSource{
									SourceChainSelector: uint64(srcChain2),
									OnrampAddress:       contractAddrs[srcChain2][consts.ContractNameOnRamp],
								},
								ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 50, MaxMsgNr: 51},
							},
						},
					).
					Return(expSigs1, nil)
				return cl
			},
			expQuery: ct.MerkleRootQuery{
				RetryRMNSignatures: false,
				RMNSignatures:      expSigs1,
			},
			expErr: false,
		},
		{
			name: "rmn timeout",
			prevOutcome: ct.MerkleRootOutcome{
				OutcomeType: ct.ReportIntervalsSelected,
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
			rmnClient: func(t *testing.T) *rmnmocks.MockClient {
				cl := rmnmocks.NewMockClient(t)
				time.Sleep(time.Millisecond)
				cl.EXPECT().ComputeReportSignatures(mock.Anything, mock.Anything, mock.Anything).
					Return(expSigs1, rmn.ErrTimeout) // <------------------------------------ timeout error
				return cl
			},
			expQuery: ct.MerkleRootQuery{
				RetryRMNSignatures: true,
				RMNSignatures:      nil,
			},
			expErr: false,
		},
		{
			name: "rmn random error",
			prevOutcome: ct.MerkleRootOutcome{
				OutcomeType: ct.ReportIntervalsSelected,
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
			rmnClient: func(t *testing.T) *rmnmocks.MockClient {
				cl := rmnmocks.NewMockClient(t)
				time.Sleep(time.Millisecond)
				cl.EXPECT().ComputeReportSignatures(mock.Anything, mock.Anything, mock.Anything).
					Return(expSigs1, fmt.Errorf("some error")) // <------------------------- some random error
				return cl
			},
			expQuery: ct.MerkleRootQuery{},
			expErr:   true,
		},
		{
			name: "not in building reports state",
			prevOutcome: ct.MerkleRootOutcome{
				OutcomeType: ct.ReportInFlight,
			},
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           true,
				RMNSignaturesTimeout: 5 * time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmnmocks.MockClient { return rmnmocks.NewMockClient(t) },
			expQuery:  ct.MerkleRootQuery{},
			expErr:    false,
		},
		{
			name: "rmn sig checks disabled",
			prevOutcome: ct.MerkleRootOutcome{
				OutcomeType: ct.ReportIntervalsSelected,
			},
			cfg: pluginconfig.CommitPluginConfig{
				RMNEnabled:           false, // <-------------- disabled
				RMNSignaturesTimeout: 5 * time.Second,
				DestChain:            dstChain,
			},
			rmnClient: func(t *testing.T) *rmnmocks.MockClient { return rmnmocks.NewMockClient(t) },
			expQuery:  ct.MerkleRootQuery{},
			expErr:    false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ccipReader := reader.NewMockCCIPReader(t)
			for chainSel, contracts := range testCase.contractAddresses {
				for name, addr := range contracts {
					ccipReader.EXPECT().GetContractAddress(name, chainSel).Return(addr, nil)
				}
			}

			w := Processor{
				cfg:        testCase.cfg,
				ccipReader: ccipReader,
				rmnClient:  testCase.rmnClient(t),
				lggr:       logger.Test(t),
			}

			q, err := w.Query(ctx, ct.Outcome{MerkleRootOutcome: testCase.prevOutcome})
			if testCase.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, testCase.expQuery, q.MerkleRootQuery)
		})
	}
}
