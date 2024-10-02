package execute

import (
	"context"
	"encoding/binary"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

const (
	randomEthAddress = "0x00000000000000000000000000001234"
)

type IntTest struct {
	t *testing.T

	donID uint32

	srcSelector cciptypes.ChainSelector
	dstSelector cciptypes.ChainSelector

	msgHasher  cciptypes.MessageHasher
	ccipReader *inmem.InMemoryCCIPReader
	server     *ConfigurableAttestationServer
}

func (it *IntTest) WithMessages(messages []inmem.MessagesWithMetadata, crBlockNumber uint64, crTimestamp time.Time) {
	mapped := slicelib.Map(messages, func(m inmem.MessagesWithMetadata) cciptypes.Message { return m.Message })
	reportData := exectypes.CommitData{
		SourceChain: it.srcSelector,
		SequenceNumberRange: cciptypes.NewSeqNumRange(
			messages[0].Header.SequenceNumber,
			messages[len(messages)-1].Header.SequenceNumber,
		),
		Messages: mapped,
	}

	tree, err := report.ConstructMerkleTree(tests.Context(it.t), it.msgHasher, reportData, logger.Test(it.t))
	require.NoError(it.t, err, "failed to construct merkle tree")

	it.ccipReader.Reports = append(it.ccipReader.Reports, plugintypes2.CommitPluginReportWithMeta{
		Report: cciptypes.CommitPluginReport{
			MerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     reportData.SourceChain,
					SeqNumsRange: reportData.SequenceNumberRange,
					MerkleRoot:   tree.Root(),
				},
			},
		},
		BlockNum:  crBlockNumber,
		Timestamp: crTimestamp,
	})

	it.ccipReader.Messages[it.srcSelector] = append(
		it.ccipReader.Messages[it.srcSelector],
		messages...,
	)
}

func SetupSimpleTest(ctx context.Context, t *testing.T, lggr logger.Logger, srcSelector, dstSelector cciptypes.ChainSelector) (*IntTest, *testhelpers.OCR3Runner[[]byte]) {
	donID := uint32(1)

	msgHasher := mocks.NewMessageHasher()
	ccipReader := inmem.InMemoryCCIPReader{
		Reports: []plugintypes2.CommitPluginReportWithMeta{},
		Messages: map[cciptypes.ChainSelector][]inmem.MessagesWithMetadata{
			srcSelector: {},
		},
		Dest: dstSelector,
	}

	server := newConfigurableAttestationServer(map[string]string{
		"0x0f43587da5355551d234a2ba24dde8edfe0e385346465d6d53653b6aa642992e": `{
			"status": "complete",
			"attestation": "0x720502893578a89a8a87982982ef781c18b193"
		}`,
	})

	cfg := pluginconfig.ExecutePluginConfig{
		OffchainConfig: pluginconfig.ExecuteOffchainConfig{
			MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
			BatchGasLimit:             100000000,
			TokenDataObservers: []pluginconfig.TokenDataObserverConfig{
				{
					Type:    "usdc-cctp",
					Version: "1",
					USDCCCTPObserverConfig: &pluginconfig.USDCCCTPObserverConfig{
						Tokens: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
							srcSelector: {
								SourcePoolAddress:            randomEthAddress,
								SourceMessageTransmitterAddr: randomEthAddress,
							},
						},
						AttestationAPI:         server.server.URL,
						AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Millisecond),
						AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Second),
					},
				},
			},
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
	err := homeChain.Start(ctx)
	require.NoError(t, err, "failed to start home chain poller")

	usdcEvents := []types.Sequence{
		{Data: newMessageSentEvent(0, 6, 1, []byte{1})},
		{Data: newMessageSentEvent(0, 6, 2, []byte{2})},
		{Data: newMessageSentEvent(0, 6, 3, []byte{3})},
	}

	r := readermock.NewMockContractReaderFacade(t)
	r.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil).Maybe()
	r.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(usdcEvents, nil).Maybe()

	tkObs, err := tokendata.NewConfigBasedCompositeObservers(
		lggr,
		cfg.DestChain,
		cfg.OffchainConfig.TokenDataObservers,
		testhelpers.TokenDataEncoderInstance,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			srcSelector: r,
			dstSelector: r,
		},
	)
	require.NoError(t, err)

	oracleIDToP2pID := GetP2pIDs(1, 2, 3)
	nodesSetup := []nodeSetup{
		newNode(donID, lggr, cfg, msgHasher, &ccipReader, homeChain, tkObs, oracleIDToP2pID, 1, 1),
		newNode(donID, lggr, cfg, msgHasher, &ccipReader, homeChain, tkObs, oracleIDToP2pID, 2, 1),
		newNode(donID, lggr, cfg, msgHasher, &ccipReader, homeChain, tkObs, oracleIDToP2pID, 3, 1),
	}

	err = homeChain.Close()
	if err != nil {
		return nil, nil
	}

	nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodes = append(nodes, n.node)
	}

	nodeIDs := make([]commontypes.OracleID, 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodeIDs = append(nodeIDs, n.node.ReportingCfg().OracleID)
	}

	it := &IntTest{
		t:           t,
		donID:       donID,
		msgHasher:   msgHasher,
		srcSelector: srcSelector,
		dstSelector: dstSelector,
		ccipReader:  &ccipReader,
		server:      server,
	}

	runner := testhelpers.NewOCR3Runner(nodes, nodeIDs, nil)

	return it, runner
}

func newNode(
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
		lggr,
	)

	return nodeSetup{
		node:        node1,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}

func makeMsgWithToken(
	seqNum cciptypes.SeqNum,
	src, dest cciptypes.ChainSelector,
	executed bool,
	tokens []cciptypes.RampTokenAmount,
) inmem.MessagesWithMetadata {
	msg := makeMsg(seqNum, src, dest, executed)
	msg.Message.TokenAmounts = tokens
	return msg
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

type ConfigurableAttestationServer struct {
	responses map[string]string
	server    *httptest.Server
}

func newConfigurableAttestationServer(responses map[string]string) *ConfigurableAttestationServer {
	c := &ConfigurableAttestationServer{
		responses: responses,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for url, response := range c.responses {
			if strings.Contains(r.RequestURI, url) {
				_, err := w.Write([]byte(response))
				if err != nil {
					panic(err)
				}
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	c.server = server
	return c
}

func (c *ConfigurableAttestationServer) AddResponse(url, response string) {
	c.responses[url] = response
}

func (c *ConfigurableAttestationServer) Close() {
	c.server.Close()
}

func newMessageSentEvent(
	sourceDomain uint32,
	destDomain uint32,
	nonce uint64,
	payload []byte,
) *readerpkg.MessageSentEvent {
	var buf []byte
	buf = binary.BigEndian.AppendUint32(buf, readerpkg.CCTPMessageVersion)
	buf = binary.BigEndian.AppendUint32(buf, sourceDomain)
	buf = binary.BigEndian.AppendUint32(buf, destDomain)
	buf = binary.BigEndian.AppendUint64(buf, nonce)

	senderBytes := [12]byte{}
	buf = append(buf, senderBytes[:]...)
	buf = append(buf, payload...)

	return &readerpkg.MessageSentEvent{Arg0: buf}
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

type nodeSetup struct {
	node        *Plugin
	reportCodec cciptypes.ExecutePluginCodec
	msgHasher   cciptypes.MessageHasher
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

func (it *IntTest) Close() {
	if it.server != nil {
		it.server.Close()
	}
}
