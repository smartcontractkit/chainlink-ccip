package execute

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas/evm"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	mock_types "github.com/smartcontractkit/chainlink-ccip/mocks/execute/exectypes"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	nodesSetup := setupSimpleTest(ctx, t, lggr, 1, 2)

	nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodes = append(nodes, n.node)
	}

	nodeIDs := make([]commontypes.OracleID, 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodeIDs = append(nodeIDs, n.node.reportingCfg.OracleID)
	}

	runner := testhelpers.NewOCR3Runner(nodes, nodeIDs, nil)

	// Contract Discovery round.
	res, err := runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err := exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	// One pending commit report only.
	// Two of the messages are executed which should be indicated in the Outcome.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)
	require.ElementsMatch(t, outcome.PendingCommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// Round 2 - Get Messages
	// Messages now attached to the pending commit.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 3 - Filter
	// An execute report with the following messages executed: 102, 103, 104, 105.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	sequenceNumbers := slicelib.Map(outcome.Report.ChainReports[0].Messages, func(m cciptypes.Message) cciptypes.SeqNum {
		return m.Header.SequenceNumber
	})
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})

}

type nodeSetup struct {
	node            *Plugin
	reportCodec     cciptypes.ExecutePluginCodec
	msgHasher       cciptypes.MessageHasher
	TokenDataReader *mock_types.MockTokenDataReader
}

func setupHomeChainPoller(
	t *testing.T,
	donID plugintypes.DonID,
	lggr logger.Logger,
	chainConfigInfos []reader.ChainConfigInfo,
) reader.HomeChain {
	const ccipConfigAddress = "0xCCIPConfigFakeAddress"

	homeChainReader := readermock.NewMockContractReaderFacade(t)
	var firstCall = true
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.MatchedBy(func(input map[string]interface{}) bool {
			_, pageIndexExists := input["pageIndex"]
			_, pageSizeExists := input["pageSize"]
			return pageIndexExists && pageSizeExists
		}),
		mock.Anything,
	).Run(
		func(args mock.Arguments) {
			arg := args.Get(4).(*[]reader.ChainConfigInfo)
			if firstCall {
				*arg = chainConfigInfos
				firstCall = false
			} else {
				*arg = []reader.ChainConfigInfo{} // return empty for other pages
			}
		}).Return(nil)

	homeChainReader.EXPECT().
		GetLatestValue(mock.Anything, types.BoundContract{
			Address: ccipConfigAddress,
			Name:    consts.ContractNameCCIPConfig,
		}.ReadIdentifier(consts.MethodNameGetOCRConfig), primitives.Unconfirmed, map[string]any{
			"donId":      donID,
			"pluginType": consts.PluginTypeExecute,
		}, mock.Anything).
		Run(
			func(
				ctx context.Context,
				readIdentifier string,
				confidenceLevel primitives.ConfidenceLevel,
				params,
				returnVal interface{},
			) {
				*returnVal.(*[]reader.OCR3ConfigWithMeta) = []reader.OCR3ConfigWithMeta{{}}
			}).
		Return(nil)

	homeChain := reader.NewHomeChainConfigPoller(
		homeChainReader,
		lggr,
		// to prevent linting error because of logging after finishing tests, we close the poller after each test, having
		// lower polling interval make it catch up faster
		time.Minute,
		types.BoundContract{
			Address: ccipConfigAddress,
			Name:    consts.ContractNameCCIPConfig,
		},
	)

	return homeChain
}

func makeMsg(seqNum cciptypes.SeqNum, src, dest cciptypes.ChainSelector, executed bool) inmem.MessagesWithMetadata {
	return inmem.MessagesWithMetadata{
		Message: cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				SourceChainSelector: src,
				SequenceNumber:      seqNum,
			},
		},
		Destination: dest,
		Executed:    executed,
	}
}

func setupSimpleTest(
	ctx context.Context, t *testing.T, lggr logger.Logger, srcSelector, dstSelector cciptypes.ChainSelector,
) []nodeSetup {
	donID := uint32(1)

	msgHasher := mocks.NewMessageHasher()

	messages := []inmem.MessagesWithMetadata{
		makeMsg(100, srcSelector, dstSelector, true),
		makeMsg(101, srcSelector, dstSelector, true),
		makeMsg(102, srcSelector, dstSelector, false),
		makeMsg(103, srcSelector, dstSelector, false),
		makeMsg(104, srcSelector, dstSelector, false),
		makeMsg(105, srcSelector, dstSelector, false),
	}

	mapped := slicelib.Map(messages, func(m inmem.MessagesWithMetadata) cciptypes.Message { return m.Message })
	reportData := exectypes.CommitData{
		SourceChain:         srcSelector,
		SequenceNumberRange: cciptypes.NewSeqNumRange(100, 105),
		Messages:            mapped,
	}

	tree, err := report.ConstructMerkleTree(context.Background(), msgHasher, reportData, logger.Test(t))
	require.NoError(t, err, "failed to construct merkle tree")

	// Initialize reader with some data
	ccipReader := inmem.InMemoryCCIPReader{
		Dest: dstSelector,
		Reports: []plugintypes2.CommitPluginReportWithMeta{
			{
				Report: cciptypes.CommitPluginReport{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{
							ChainSel:     reportData.SourceChain,
							SeqNumsRange: reportData.SequenceNumberRange,
							MerkleRoot:   tree.Root(),
						},
					},
				},
				BlockNum:  1000,
				Timestamp: time.Now().Add(-4 * time.Hour),
			},
		},
		Messages: map[cciptypes.ChainSelector][]inmem.MessagesWithMetadata{
			srcSelector: {
				makeMsg(100, srcSelector, dstSelector, true),
				makeMsg(101, srcSelector, dstSelector, true),
				makeMsg(102, srcSelector, dstSelector, false),
				makeMsg(103, srcSelector, dstSelector, false),
				makeMsg(104, srcSelector, dstSelector, false),
				makeMsg(105, srcSelector, dstSelector, false),
			},
		},
	}

	cfg := pluginconfig.ExecutePluginConfig{
		OffchainConfig: pluginconfig.ExecuteOffchainConfig{
			MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
			BatchGasLimit:             100000000,
		},
		DestChain: dstSelector,
	}
	chainConfigInfos := []reader.ChainConfigInfo{
		{
			ChainSelector: srcSelector,
			ChainConfig: reader.HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					{1}, {2}, {3},
				},
				Config: mustEncodeChainConfig(chainconfig.ChainConfig{}),
			},
		}, {
			ChainSelector: dstSelector,
			ChainConfig: reader.HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					{1}, {2}, {3},
				},
				Config: mustEncodeChainConfig(chainconfig.ChainConfig{}),
			},
		},
	}

	homeChain := setupHomeChainPoller(t, donID, lggr, chainConfigInfos)
	err = homeChain.Start(ctx)
	require.NoError(t, err, "failed to start home chain poller")

	tokenDataObserver := &tokendata.NoopTokenDataObserver{}

	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	nodes := []nodeSetup{
		newNode(ctx, t, donID, lggr, cfg, msgHasher, ccipReader, homeChain, tokenDataObserver, oracleIDToP2pID, 1, 1),
		newNode(ctx, t, donID, lggr, cfg, msgHasher, ccipReader, homeChain, tokenDataObserver, oracleIDToP2pID, 2, 1),
		newNode(ctx, t, donID, lggr, cfg, msgHasher, ccipReader, homeChain, tokenDataObserver, oracleIDToP2pID, 3, 1),
	}

	err = homeChain.Close()
	if err != nil {
		return nil
	}
	return nodes
}

func newNode(
	_ context.Context,
	_ *testing.T,
	donID plugintypes.DonID,
	lggr logger.Logger,
	cfg pluginconfig.ExecutePluginConfig,
	msgHasher cciptypes.MessageHasher,
	ccipReader readerpkg.CCIPReader,
	homeChain reader.HomeChain,
	tokenDataObserver tokendata.TokenDataObserver,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	id int,
	N int,
) nodeSetup {
	reportCodec := mocks.NewExecutePluginJSONReportCodec()

	rCfg := ocr3types.ReportingPluginConfig{
		N:        N,
		OracleID: commontypes.OracleID(id),
	}

	node1 := NewPlugin(
		donID,
		rCfg,
		cfg,
		oracleIDToP2pID,
		ccipReader,
		reportCodec,
		msgHasher,
		homeChain,
		tokenDataObserver,
		evm.EstimateProvider{},
		lggr)

	return nodeSetup{
		node:        node1,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}

func GetP2pIDs(ids ...int) map[commontypes.OracleID]libocrtypes.PeerID {
	res := make(map[commontypes.OracleID]libocrtypes.PeerID)
	for _, id := range ids {
		res[commontypes.OracleID(id)] = libocrtypes.PeerID{byte(id)}
	}
	return res
}

func mustEncodeChainConfig(cc chainconfig.ChainConfig) []byte {
	encoded, err := chainconfig.EncodeChainConfig(cc)
	if err != nil {
		panic(err)
	}
	return encoded
}
