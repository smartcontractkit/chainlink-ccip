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
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	mock_types "github.com/smartcontractkit/chainlink-ccip/mocks/execute/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
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

	// In the first round there is a pending commit report only.
	// Two of the messages are executed which should be indicated in the Outcome.
	res, err := runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err := plugintypes.DecodeExecutePluginOutcome(res.Outcome)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)
	require.ElementsMatch(t, outcome.PendingCommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// In the second round there is an exec report and the pending commit report is removed.
	// The exec report should indicate the following messages are executed: 102, 103, 104, 105.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = plugintypes.DecodeExecutePluginOutcome(res.Outcome)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 1)
	require.Len(t, outcome.PendingCommitReports, 0)
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

func setupHomeChainPoller(lggr logger.Logger, chainConfigInfos []reader.ChainConfigInfo) reader.HomeChain {
	homeChainReader := mocks.NewContractReaderMock()
	homeChainReader.On(
		"GetLatestValue",
		mock.Anything,
		consts.ContractNameCCIPConfig,
		consts.MethodNameGetAllChainConfigs,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Run(
		func(args mock.Arguments) {
			arg := args.Get(5).(*[]reader.ChainConfigInfo)
			*arg = chainConfigInfos
		}).Return(nil)

	homeChain := reader.NewHomeChainConfigPoller(
		homeChainReader,
		lggr,
		// to prevent linting error because of logging after finishing tests, we close the poller after each test, having
		// lower polling interval make it catch up faster
		time.Minute,
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
	reportData := plugintypes.ExecutePluginCommitData{
		SourceChain:         srcSelector,
		SequenceNumberRange: cciptypes.NewSeqNumRange(100, 105),
		Messages:            mapped,
	}

	tree, err := report.ConstructMerkleTree(context.Background(), msgHasher, reportData, logger.Test(t))
	require.NoError(t, err, "failed to construct merkle tree")

	// Initialize reader with some data
	ccipReader := inmem.InMemoryCCIPReader{
		Dest: dstSelector,
		Reports: []plugintypes.CommitPluginReportWithMeta{
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
				Config: []byte{0},
			},
		}, {
			ChainSelector: dstSelector,
			ChainConfig: reader.HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					{1}, {2}, {3},
				},
				Config: []byte{0},
			},
		},
	}

	homeChain := setupHomeChainPoller(lggr, chainConfigInfos)
	err = homeChain.Start(ctx)
	require.NoError(t, err, "failed to start home chain poller")

	tokenDataReader := mock_types.NewMockTokenDataReader(t)
	tokenDataReader.On("ReadTokenData", mock.Anything, mock.Anything, mock.Anything).Return([][]byte{}, nil)

	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	nodes := []nodeSetup{
		newNode(ctx, t, lggr, cfg, msgHasher, ccipReader, homeChain, tokenDataReader, oracleIDToP2pID, 1, 1),
		newNode(ctx, t, lggr, cfg, msgHasher, ccipReader, homeChain, tokenDataReader, oracleIDToP2pID, 2, 1),
		newNode(ctx, t, lggr, cfg, msgHasher, ccipReader, homeChain, tokenDataReader, oracleIDToP2pID, 3, 1),
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
	lggr logger.Logger,
	cfg pluginconfig.ExecutePluginConfig,
	msgHasher cciptypes.MessageHasher,
	ccipReader reader.CCIP,
	homeChain reader.HomeChain,
	tokenDataReader types.TokenDataReader,
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
		rCfg,
		cfg,
		oracleIDToP2pID,
		ccipReader,
		reportCodec,
		msgHasher,
		homeChain,
		tokenDataReader,
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
