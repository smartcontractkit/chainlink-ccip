package execute

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	nodeSetups := setupSimpleTest(ctx, t, lggr, 1)
	//runner := testhelpers.NewOCR3Runner(nodes, nodeIDs, o)

	nodesSetup := nodeSetups
	nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodes = append(nodes, n.node)
	}

	nodeIDs := make([]commontypes.OracleID, 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodeIDs = append(nodeIDs, n.node.reportingCfg.OracleID)
	}

	//o, err := tc.initialOutcome.Encode()
	//require.NoError(t, err)
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
		10*time.Millisecond,
	)

	return homeChain
}

func setupSimpleTest(ctx context.Context, t *testing.T, lggr logger.Logger, selector cciptypes.ChainSelector) []nodeSetup {
	cfg := pluginconfig.ExecutePluginConfig{}
	chainConfigInfos := []reader.ChainConfigInfo{
		{
			ChainSelector: selector,
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

	ccipReader := mocks.InMemoryCCIPReader{
		Reports:  nil,
		Messages: nil,
		Dest:     0,
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

var ()
