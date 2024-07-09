package execute

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
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

	res, err := runner.RunRound(ctx)
	fmt.Println(res, err)
}

type nodeSetup struct {
	node        *Plugin
	priceReader *mocks.TokenPricesReader
	reportCodec *mocks.ExecutePluginJSONReportCodec
	msgHasher   *mocks.MessageHasher
}

func setupHomeChainPoller(lggr logger.Logger, chainConfigInfos []reader.ChainConfigInfo) reader.HomeChain {
	homeChainReader := mocks.NewContractReaderMock()
	homeChainReader.On(
		"GetLatestValue", mock.Anything, "CCIPConfig", "getAllChainConfigs", mock.Anything, mock.Anything,
	).Run(
		func(args mock.Arguments) {
			arg := args.Get(4).(*[]reader.ChainConfigInfo)
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

func makeMsg(seqNum cciptypes.SeqNum, dest cciptypes.ChainSelector, executed bool) inmem.MessagesWithMetadata {
	return inmem.MessagesWithMetadata{
		Message: cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				SequenceNumber: seqNum,
			},
		},
		Destination: dest,
		Executed:    executed,
	}
}

func setupSimpleTest(
	ctx context.Context, t *testing.T, lggr logger.Logger, srcSelector, dstSelector cciptypes.ChainSelector,
) []nodeSetup {
	// Initialize reader with some data
	ccipReader := inmem.InMemoryCCIPReader{
		Dest: dstSelector,
		Reports: []plugintypes.CommitPluginReportWithMeta{
			{
				Report: cciptypes.CommitPluginReport{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{
							ChainSel:     srcSelector,
							SeqNumsRange: cciptypes.NewSeqNumRange(100, 105),
							//MerkleRoot:   [],
						},
					},
				},
				BlockNum:  1000,
				Timestamp: time.Now().Add(-4 * time.Hour),
			},
		},
		Messages: map[cciptypes.ChainSelector][]inmem.MessagesWithMetadata{
			srcSelector: {
				makeMsg(100, dstSelector, true),
				makeMsg(101, dstSelector, true),
				makeMsg(102, dstSelector, false),
				makeMsg(103, dstSelector, false),
				makeMsg(104, dstSelector, false),
				makeMsg(105, dstSelector, false),
			},
		},
	}

	cfg := pluginconfig.ExecutePluginConfig{
		MessageVisibilityInterval: 8 * time.Hour,
		DestChain:                 dstSelector,
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
	err := homeChain.Start(ctx)
	if err != nil {
		return nil
	}

	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	nodes := []nodeSetup{
		newNode(ctx, t, lggr, cfg, ccipReader, homeChain, oracleIDToP2pID, 1, 1),
		newNode(ctx, t, lggr, cfg, ccipReader, homeChain, oracleIDToP2pID, 2, 1),
		newNode(ctx, t, lggr, cfg, ccipReader, homeChain, oracleIDToP2pID, 3, 1),
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
	ccipReader reader.CCIP,
	homeChain reader.HomeChain,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	id int,
	N int,
) nodeSetup {
	priceReader := mocks.NewTokenPricesReader()
	reportCodec := mocks.NewExecutePluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()

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
		lggr)

	return nodeSetup{
		node:        node1,
		priceReader: priceReader,
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
