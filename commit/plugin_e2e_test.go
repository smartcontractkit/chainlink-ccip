package commit

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	helpers "github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	testCases := []struct {
		name                  string
		description           string
		nodes                 []nodeSetup
		expErr                func(*testing.T, error)
		expOutcome            plugintypes.CommitPluginOutcome
		expTransmittedReports []cciptypes.CommitPluginReport
		initialOutcome        plugintypes.CommitPluginOutcome
	}{
		{
			name:        "EmptyOutcome",
			description: "Empty observations are returned by all nodes which leads to an empty outcome.",
			nodes:       setupEmptyOutcome(ctx, t, lggr),
			expErr:      func(t *testing.T, err error) { assert.Equal(t, helpers.ErrEmptyOutcome, err) },
		},
		{
			name: "AllNodesReadAllChains",
			description: "Nodes observe the latest sequence numbers and new messages after those sequence numbers. " +
				"They also observe gas prices. In this setup all nodes can read all chains.",
			nodes: setupAllNodesReadAllChains(ctx, t, lggr),
			expOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: lastCommittedSeqNumA},
					{ChainSel: chainB, SeqNum: lastCommittedSeqNumB},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{
					{ChainSel: chainB, MerkleRoot: cciptypes.Bytes32{}, SeqNumsRange: cciptypes.NewSeqNumRange(21, 22)},
				},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices: []cciptypes.GasPriceChain{
					{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
					{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
				},
			},
			expTransmittedReports: []cciptypes.CommitPluginReport{
				{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{ChainSel: chainB, SeqNumsRange: cciptypes.NewSeqNumRange(21, 22)},
					},
					PriceUpdates: cciptypes.PriceUpdates{
						TokenPriceUpdates: []cciptypes.TokenPrice{},
						GasPriceUpdates: []cciptypes.GasPriceChain{
							{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
							{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
						},
					},
				},
			},
			initialOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: lastCommittedSeqNumA},
					{ChainSel: chainB, SeqNum: lastCommittedSeqNumB},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices:   []cciptypes.GasPriceChain{},
			},
		},
		{
			name:        "NodesDoNotAgreeOnMsgs",
			description: "Nodes do not agree on messages which leads to an outcome with empty merkle roots.",
			nodes:       setupNodesDoNotAgreeOnMsgs(ctx, t, lggr),
			expOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: lastCommittedSeqNumA},
					{ChainSel: chainB, SeqNum: lastCommittedSeqNumB},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices: []cciptypes.GasPriceChain{
					{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
					{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
				},
			},
			expTransmittedReports: []cciptypes.CommitPluginReport{
				{
					MerkleRoots: []cciptypes.MerkleRootChain{},
					PriceUpdates: cciptypes.PriceUpdates{
						TokenPriceUpdates: []cciptypes.TokenPrice{},
						GasPriceUpdates: []cciptypes.GasPriceChain{
							{ChainSel: chainA, GasPrice: cciptypes.NewBigIntFromInt64(1000)},
							{ChainSel: chainB, GasPrice: cciptypes.NewBigIntFromInt64(20_000)},
						},
					},
				},
			},
			initialOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: lastCommittedSeqNumA},
					{ChainSel: chainB, SeqNum: lastCommittedSeqNumB},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices:   []cciptypes.GasPriceChain{},
			},
		},
		{
			name:        "NodesDoNotReportGasPrices",
			description: "Nodes that don't have access to a contract writer do not submit gas price updates",
			nodes:       setupNodesDoNotReportGasPrices(ctx, t, lggr),
			expOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: lastCommittedSeqNumA},
					{ChainSel: chainB, SeqNum: lastCommittedSeqNumB},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{
					{ChainSel: chainB, MerkleRoot: cciptypes.Bytes32{}, SeqNumsRange: cciptypes.NewSeqNumRange(seqNumB, 22)},
				},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices:   []cciptypes.GasPriceChain{},
			},
			expTransmittedReports: []cciptypes.CommitPluginReport{
				{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{ChainSel: chainB, SeqNumsRange: cciptypes.NewSeqNumRange(seqNumB, 22)},
					},
					PriceUpdates: cciptypes.PriceUpdates{
						TokenPriceUpdates: []cciptypes.TokenPrice{},
						GasPriceUpdates:   []cciptypes.GasPriceChain{},
					},
				},
			},
			initialOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: chainA, SeqNum: lastCommittedSeqNumA},
					{ChainSel: chainB, SeqNum: lastCommittedSeqNumB},
				},
				MerkleRoots: []cciptypes.MerkleRootChain{},
				TokenPrices: []cciptypes.TokenPrice{},
				GasPrices:   []cciptypes.GasPriceChain{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
			t.Logf(">>> [%s]\n", tc.name)
			t.Logf(">>> %s\n", tc.description)
			defer t.Log("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")

			nodesSetup := tc.nodes
			nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
			for _, n := range nodesSetup {
				nodes = append(nodes, n.node)
			}

			nodeIDs := make([]commontypes.OracleID, 0, len(nodesSetup))
			for _, n := range nodesSetup {
				nodeIDs = append(nodeIDs, n.node.nodeID)
			}
			o, err := tc.initialOutcome.Encode()
			require.NoError(t, err)
			runner := helpers.NewOCR3Runner(nodes, nodeIDs, o)

			res, err := runner.RunRound(ctx)
			if tc.expErr != nil {
				tc.expErr(t, err)
			} else {
				require.NoErrorf(t, err, "failed to run round, testcase %s", tc.name)
			}

			if !reflect.DeepEqual(tc.expOutcome, plugintypes.CommitPluginOutcome{}) {
				outcome, err := plugintypes.DecodeCommitPluginOutcome(res.Outcome)
				require.NoError(t, err)
				assert.Equal(t, tc.expOutcome.TokenPrices, outcome.TokenPrices)
				assert.Equal(t, tc.expOutcome.MaxSeqNums, outcome.MaxSeqNums)
				assert.Equal(t, tc.expOutcome.GasPrices, outcome.GasPrices)

				assert.Equal(t, len(tc.expOutcome.MerkleRoots), len(outcome.MerkleRoots))
				for i, exp := range tc.expOutcome.MerkleRoots {
					assert.Equal(t, exp.ChainSel, outcome.MerkleRoots[i].ChainSel)
					assert.Equal(t, exp.SeqNumsRange, outcome.MerkleRoots[i].SeqNumsRange)
				}
			}

			assert.Equal(t, len(tc.expTransmittedReports), len(res.Transmitted))
			for i, exp := range tc.expTransmittedReports {
				actual, err := nodesSetup[0].reportCodec.Decode(ctx, res.Transmitted[i].Report)
				assert.NoError(t, err)
				assert.Equal(t, exp.PriceUpdates, actual.PriceUpdates)
				assert.Equal(t, len(exp.MerkleRoots), len(actual.MerkleRoots))
				for j, expRoot := range exp.MerkleRoots {
					assert.Equal(t, expRoot.ChainSel, actual.MerkleRoots[j].ChainSel)
					assert.Equal(t, expRoot.SeqNumsRange, actual.MerkleRoots[j].SeqNumsRange)
				}
			}
		})
	}
}

func setupEmptyOutcome(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	chainConfigInfos := []reader.ChainConfigInfo{
		helpers.SetupConfigInfo(destChain, pIDs1_2_3, fChainOne, cfgC),
	}

	homeChain := setupHomeChainPoller(lggr, chainConfigInfos)
	require.NoError(t, homeChain.Start(ctx))

	oracleIDToP2pID := helpers.CreateOracleIDToP2pID(1, 2, 3)
	nodes := []nodeSetup{
		newNode(ctx, t, lggr, 1, destCfg, homeChain, oracleIDToP2pID),
		newNode(ctx, t, lggr, 2, destCfg, homeChain, oracleIDToP2pID),
		newNode(ctx, t, lggr, 3, destCfg, homeChain, oracleIDToP2pID),
	}

	for _, n := range nodes {
		// All nodes have issue reading the latest sequence number, should lead to empty outcomes
		n.ccipReader.On(
			"NextSeqNum",
			ctx,
			mock.Anything,
		).Return([]cciptypes.SeqNum{}, nil)
	}

	require.NoError(t, homeChain.Close())
	return nodes
}

func setupAllNodesReadAllChains(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	chainConfigInfos := []reader.ChainConfigInfo{
		helpers.SetupConfigInfo(chainA, pIDs1_2_3, fChainOne, cfgA),
		helpers.SetupConfigInfo(chainB, pIDs1_2_3, fChainOne, cfgB),
		helpers.SetupConfigInfo(destChain, pIDs1_2_3, fChainOne, cfgC),
	}

	homeChain := setupHomeChainPoller(lggr, chainConfigInfos)
	require.NoError(t, homeChain.Start(ctx))

	oracleIDToP2pID := helpers.CreateOracleIDToP2pID(1, 2, 3)

	var nodes []nodeSetup
	for i := 1; i <= 3; i++ {
		n := newNode(ctx, t, lggr, i, destCfg, homeChain, oracleIDToP2pID)
		nodes = append(nodes, n)
		// then they fetch new msgs, there is nothing new on chainA
		mockMsgsBetweenSeqNums(ctx, n.ccipReader, chainA, seqNumA, emptyMsgs)
		// and there are two new message on chainB
		mockMsgsBetweenSeqNums(ctx, n.ccipReader, chainB, seqNumB, chainBDefaultMsgs)

		mockGasPrices(ctx, n.ccipReader, []cciptypes.ChainSelector{chainA, chainB}, []int64{1000, 20_000})

		// all nodes observe the same sequence numbers lastCommittedSeqNumA for chainA and lastCommittedSeqNumB for chainB
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.SeqNum{lastCommittedSeqNumA + 1, lastCommittedSeqNumB + 1}, nil)

		// transmission phase root staleness check passes
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainB}).
			Return([]cciptypes.SeqNum{lastCommittedSeqNumB + 1}, nil)
	}
	require.NoError(t, homeChain.Close())
	return nodes
}

func setupNodesDoNotAgreeOnMsgs(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	chainConfigInfos := []reader.ChainConfigInfo{
		helpers.SetupConfigInfo(chainA, pIDs1_2_3, fChainOne, cfgA),
		helpers.SetupConfigInfo(chainB, pIDs1_2_3, fChainOne, cfgB),
		helpers.SetupConfigInfo(destChain, pIDs1_2_3, fChainOne, cfgB),
	}

	homeChain := setupHomeChainPoller(lggr, chainConfigInfos)
	require.NoError(t, homeChain.Start(ctx))
	oracleIDToP2pID := helpers.CreateOracleIDToP2pID(1, 2, 3)

	var nodes []nodeSetup
	for i := 1; i <= 3; i++ {
		n := newNode(ctx, t, lggr, i, destCfg, homeChain, oracleIDToP2pID)
		nodes = append(nodes, n)
		// all nodes observe the same sequence numbers lastCommittedSeqNumA for chainA and lastCommittedSeqNumB for chainB
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.SeqNum{lastCommittedSeqNumA + 1, lastCommittedSeqNumB + 1}, nil)

		// then they fetch new msgs, there is nothing new on chainA
		mockMsgsBetweenSeqNums(ctx, n.ccipReader, chainA, seqNumA, emptyMsgs)

		var otherChainBMsgs = make([]cciptypes.Message, len(chainBDefaultMsgs))
		copy(otherChainBMsgs[:], chainBDefaultMsgs[:])
		otherChainBMsgs[0].Header.MessageID = messageIDFromInt(i + 1)
		otherChainBMsgs[0].Header.SequenceNumber = seqNumB + +cciptypes.SeqNum(i*int(lastCommittedSeqNumA))
		otherChainBMsgs[1].Header.MessageID = messageIDFromInt(i + 2)
		otherChainBMsgs[1].Header.SequenceNumber = 22 + +cciptypes.SeqNum(i*int(lastCommittedSeqNumA))
		mockMsgsBetweenSeqNums(ctx, n.ccipReader, chainB, seqNumB, otherChainBMsgs)

		mockGasPrices(ctx, n.ccipReader, []cciptypes.ChainSelector{chainA, chainB}, []int64{1000, 20_000})
	}

	require.NoError(t, homeChain.Close())
	return nodes
}

func setupNodesDoNotReportGasPrices(ctx context.Context, t *testing.T, lggr logger.Logger) []nodeSetup {
	chainConfigInfos := []reader.ChainConfigInfo{
		helpers.SetupConfigInfo(chainA, pIDs1_2_3, fChainOne, cfgA),
		helpers.SetupConfigInfo(chainB, pIDs1_2_3, fChainOne, cfgB),
		helpers.SetupConfigInfo(destChain, pIDs1_2_3, fChainOne, cfgB),
	}

	homeChain := setupHomeChainPoller(lggr, chainConfigInfos)
	require.NoError(t, homeChain.Start(ctx))

	oracleIDToP2pID := helpers.CreateOracleIDToP2pID(1, 2, 3)

	var nodes []nodeSetup
	for i := 1; i <= 3; i++ {
		n := newNode(ctx, t, lggr, i, destCfg, homeChain, oracleIDToP2pID)
		nodes = append(nodes, n)
		// then they fetch new msgs, there is nothing new on chainA
		mockMsgsBetweenSeqNums(ctx, n.ccipReader, chainA, seqNumA, emptyMsgs)
		// and there are two new message on chainB
		mockMsgsBetweenSeqNums(ctx, n.ccipReader, chainB, seqNumB, chainBDefaultMsgs)

		n.ccipReader.On("GasPrices", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.BigInt{}, fmt.Errorf("no gas prices available: %w", reader.ErrContractWriterNotFound))

		// all nodes observe the same sequence numbers lastCommittedSeqNumA for chainA and lastCommittedSeqNumB for chainB
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainA, chainB}).
			Return([]cciptypes.SeqNum{lastCommittedSeqNumA + 1, lastCommittedSeqNumB + 1}, nil)
		n.ccipReader.On("NextSeqNum", ctx, []cciptypes.ChainSelector{chainB}).
			Return([]cciptypes.SeqNum{lastCommittedSeqNumB + 1}, nil)
	}

	require.NoError(t, homeChain.Close())
	return nodes
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *mocks.CCIPReader
	priceReader *mocks.TokenPricesReader
	reportCodec *mocks.CommitPluginJSONReportCodec
	msgHasher   *mocks.MessageHasher
}

func newNode(
	_ context.Context,
	_ *testing.T,
	lggr logger.Logger,
	id int,
	cfg pluginconfig.CommitPluginConfig,
	homeChain reader.HomeChain,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
) nodeSetup {
	ccipReader := mocks.NewCCIPReader()
	priceReader := mocks.NewTokenPricesReader()
	reportCodec := mocks.NewCommitPluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()

	node1 := NewPlugin(
		context.Background(),
		commontypes.OracleID(id),
		oracleIDToP2pID,
		cfg,
		ccipReader,
		priceReader,
		reportCodec,
		msgHasher,
		lggr,
		homeChain,
	)

	return nodeSetup{
		node:        node1,
		ccipReader:  ccipReader,
		priceReader: priceReader,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
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
		10*time.Millisecond,
	)

	return homeChain
}

// mockGasPrices mocks the gas prices for the given chains
// the gas prices are returned in the same order as the chains
func mockGasPrices(
	ctx context.Context,
	ccipReader *mocks.CCIPReader,
	chains []cciptypes.ChainSelector,
	gasPrices []int64) {
	gasPricesBigInt := make([]cciptypes.BigInt, len(gasPrices))
	for i, gp := range gasPrices {
		gasPricesBigInt[i] = cciptypes.NewBigIntFromInt64(gp)
	}

	ccipReader.On("GasPrices", ctx, chains).
		Return(gasPricesBigInt, nil)
}

func mockMsgsBetweenSeqNums(
	ctx context.Context,
	ccipReader *mocks.CCIPReader,
	chain cciptypes.ChainSelector,
	seqNum cciptypes.SeqNum,
	msgs []cciptypes.Message) {
	ccipReader.On(
		"MsgsBetweenSeqNums",
		ctx,
		chain,
		cciptypes.NewSeqNumRange(seqNum, cciptypes.SeqNum(int(seqNum)+destCfg.NewMsgScanBatchSize)),
	).Return(msgs, nil)
}

var (
	chainA               = cciptypes.ChainSelector(1)
	cfgA                 = []byte("ChainA")
	lastCommittedSeqNumA = cciptypes.SeqNum(10)
	seqNumA              = cciptypes.SeqNum(11)
	chainB               = cciptypes.ChainSelector(2)
	cfgB                 = []byte("ChainB")
	lastCommittedSeqNumB = cciptypes.SeqNum(20)
	seqNumB              = cciptypes.SeqNum(21)

	emptyMsgs = []cciptypes.Message{}

	chainBDefaultMsgs = []cciptypes.Message{
		{
			Header: cciptypes.RampMessageHeader{
				MsgHash:             cciptypes.Bytes32{1},
				MessageID:           mustNewMessageID("0x01"),
				SourceChainSelector: chainB,
				SequenceNumber:      seqNumB,
			},
		},
		{
			Header: cciptypes.RampMessageHeader{
				MsgHash:             cciptypes.Bytes32{2},
				MessageID:           mustNewMessageID("0x02"),
				SourceChainSelector: chainB,
				SequenceNumber:      22,
			},
		},
	}

	destChain = cciptypes.ChainSelector(3)
	cfgC      = []byte("destChain")
	fChainOne = uint8(1)

	pIDs1_2_3 = []libocrtypes.PeerID{{1}, {2}, {3}}
	tokenX    = types.Account("tk_xxx")

	destCfg = pluginconfig.CommitPluginConfig{
		DestChain:           destChain,
		PricedTokens:        []types.Account{tokenX},
		TokenPricesObserver: false,
		NewMsgScanBatchSize: 256,
	}
)
