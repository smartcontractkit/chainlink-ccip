package execute

import (
	"context"
	"encoding/binary"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/observer"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	cciptypesmocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type IntTest struct {
	t *testing.T

	lggr  logger.Logger
	donID uint32

	srcSelectors []cciptypes.ChainSelector
	dstSelector  cciptypes.ChainSelector

	msgHasher           cciptypes.MessageHasher
	ccipReader          *inmem.InMemoryCCIPReader
	usdcServer          *ConfigurableAttestationServer
	tokenObserverConfig []pluginconfig.TokenDataObserverConfig
	tokenChainReader    map[cciptypes.ChainSelector]contractreader.ContractReaderFacade
}

func SetupSimpleTest(t *testing.T,
	lggr logger.Logger,
	srcSelectors []cciptypes.ChainSelector,
	dstSelector cciptypes.ChainSelector,
) *IntTest {
	donID := uint32(1)

	messagesMap := make(map[cciptypes.ChainSelector][]inmem.MessagesWithMetadata)

	for _, src := range srcSelectors {
		messagesMap[src] = []inmem.MessagesWithMetadata{}

	}
	msgHasher := mocks.NewMessageHasher()
	ccipReader := inmem.InMemoryCCIPReader{
		UnfinalizedReports: []cciptypes.CommitPluginReportWithMeta{},
		Messages:           messagesMap,
		Dest:               dstSelector,
	}

	return &IntTest{
		t:                   t,
		lggr:                lggr,
		donID:               donID,
		msgHasher:           msgHasher,
		srcSelectors:        srcSelectors,
		dstSelector:         dstSelector,
		ccipReader:          &ccipReader,
		tokenObserverConfig: []pluginconfig.TokenDataObserverConfig{},
		tokenChainReader:    map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{},
	}
}

func (it *IntTest) WithMessages(
	messages []inmem.MessagesWithMetadata,
	crBlockNumber uint64,
	crTimestamp time.Time,
	numReports int,
	srcSelector cciptypes.ChainSelector) {
	mapped := slicelib.Map(messages,
		func(m inmem.MessagesWithMetadata) cciptypes.Message {
			return m.Message
		},
	)
	totalMessages := len(mapped)
	messagesPerReport := totalMessages / numReports

	for i := 0; i < numReports; i++ {
		startIndex := i * messagesPerReport
		endIndex := startIndex + messagesPerReport
		if i == numReports-1 {
			endIndex = totalMessages // Ensure the last report includes any remaining messages
		}

		msgs := mapped[startIndex:endIndex]
		hashes := make([]cciptypes.Bytes32, len(msgs))
		for i, m := range msgs {
			hash, err := it.msgHasher.Hash(context.Background(), m)
			require.NoError(it.t, err, "failed to hash message")
			hashes[i] = hash
		}
		reportData := exectypes.CommitData{
			SourceChain: srcSelector,
			SequenceNumberRange: cciptypes.NewSeqNumRange(
				mapped[startIndex].Header.SequenceNumber,
				mapped[endIndex-1].Header.SequenceNumber,
			),
			Messages:         msgs,
			Hashes:           hashes,
			MessageTokenData: make([]exectypes.MessageTokenData, len(msgs)),
			Timestamp:        crTimestamp,
		}

		tree, err := report.ConstructMerkleTree(reportData, logger.Test(it.t))
		require.NoError(it.t, err, "failed to construct merkle tree")

		it.ccipReader.UnfinalizedReports = append(it.ccipReader.UnfinalizedReports, cciptypes.CommitPluginReportWithMeta{
			Report: cciptypes.CommitPluginReport{
				BlessedMerkleRoots: []cciptypes.MerkleRootChain{
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
	}

	it.ccipReader.Messages[srcSelector] = append(
		it.ccipReader.Messages[srcSelector],
		messages...,
	)
}

func (it *IntTest) WithUSDC(
	sourcePoolAddress string,
	attestations map[string]string,
	events []*readerpkg.MessageSentEvent,
	srcSelector cciptypes.ChainSelector,
) {
	it.usdcServer = newConfigurableAttestationServer(attestations)
	it.tokenObserverConfig = append(it.tokenObserverConfig, pluginconfig.TokenDataObserverConfig{
		Type:    "usdc-cctp",
		Version: "1",
		USDCCCTPObserverConfig: &pluginconfig.USDCCCTPObserverConfig{
			AttestationConfig: pluginconfig.AttestationConfig{
				AttestationAPI:         it.usdcServer.server.URL,
				AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Millisecond),
				AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Second),
			},
			AttestationAPICooldown: commonconfig.MustNewDuration(5 * time.Minute),
			Tokens: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				srcSelector: {
					SourcePoolAddress:            sourcePoolAddress,
					SourceMessageTransmitterAddr: sourcePoolAddress,
				},
			},
		},
	})

	usdcEvents := make([]types.Sequence, len(events))
	for i, e := range events {
		usdcEvents[i] = types.Sequence{Data: e}
	}

	r := readermock.NewMockContractReaderFacade(it.t)
	r.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil).Maybe()
	r.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(usdcEvents, nil).Maybe()

	it.tokenChainReader = map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
		srcSelector:    r,
		it.dstSelector: r,
	}
}

func (it *IntTest) Start() *testhelpers.OCR3Runner[[]byte] {
	cfg := pluginconfig.ExecuteOffchainConfig{
		MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
		BatchGasLimit:             100000000,
	}
	chainConfigInfos := []reader.ChainConfigInfo{
		{
			ChainSelector: it.dstSelector,
			ChainConfig: reader.HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					{1}, {2}, {3},
				},
				Config: mustEncodeChainConfig(chainconfig.ChainConfig{}),
			},
		},
	}
	// Add config for all srcSelectors
	for _, src := range it.srcSelectors {
		chainConfigInfos = append(chainConfigInfos, reader.ChainConfigInfo{
			ChainSelector: src,
			ChainConfig: reader.HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					{1}, {2}, {3},
				},
				Config: mustEncodeChainConfig(chainconfig.ChainConfig{}),
			},
		})
	}

	homeChain := setupHomeChainPoller(it.t, it.lggr, chainConfigInfos)
	ctx := tests.Context(it.t)
	err := homeChain.Start(ctx)
	require.NoError(it.t, err, "failed to start home chain poller")
	mockAddrCodec := internal.NewMockAddressCodecHex(it.t)
	tkObs, err := observer.NewConfigBasedCompositeObservers(
		ctx,
		it.lggr,
		it.dstSelector,
		it.tokenObserverConfig,
		testhelpers.TokenDataEncoderInstance,
		it.tokenChainReader,
		mockAddrCodec,
	)
	require.NoError(it.t, err)

	ep := cciptypesmocks.NewMockEstimateProvider(it.t)
	ep.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(uint64(0)).Maybe()
	ep.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(uint64(0)).Maybe()

	oracleIDToP2pID := testhelpers.CreateOracleIDToP2pID(1, 2, 3)
	nodesSetup := []nodeSetup{
		it.newNode(cfg, homeChain, ep, tkObs, oracleIDToP2pID, 1, 1, [32]byte{0xde, 0xad}, mockAddrCodec),
		it.newNode(cfg, homeChain, ep, tkObs, oracleIDToP2pID, 2, 1, [32]byte{0xde, 0xad}, mockAddrCodec),
		it.newNode(cfg, homeChain, ep, tkObs, oracleIDToP2pID, 3, 1, [32]byte{0xde, 0xad}, mockAddrCodec),
	}

	require.NoError(it.t, homeChain.Close())

	nodes := make([]ocr3types.ReportingPlugin[[]byte], 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodes = append(nodes, n.node)
	}

	nodeIDs := make([]commontypes.OracleID, 0, len(nodesSetup))
	for _, n := range nodesSetup {
		nodeIDs = append(nodeIDs, n.node.reportingCfg.OracleID)
	}

	return testhelpers.NewOCR3Runner(nodes, nodeIDs, nil)
}

func (it *IntTest) Close() {
	if it.usdcServer != nil {
		it.usdcServer.Close()
	}
}

func (it *IntTest) newNode(
	cfg pluginconfig.ExecuteOffchainConfig,
	homeChain reader.HomeChain,
	ep cciptypes.EstimateProvider,
	tokenDataObserver observer.TokenDataObserver,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	id int,
	N int,
	configDigest [32]byte,
	mockCodec *cciptypesmocks.MockAddressCodec,
) nodeSetup {
	reportCodec := mocks.NewExecutePluginJSONReportCodec()
	rCfg := ocr3types.ReportingPluginConfig{
		N:            N,
		OracleID:     commontypes.OracleID(id),
		ConfigDigest: configDigest,
	}

	it.ccipReader.ConfigDigest = configDigest
	node1 := NewPlugin(
		it.donID,
		rCfg,
		cfg,
		it.dstSelector,
		oracleIDToP2pID,
		it.ccipReader,
		reportCodec,
		it.msgHasher,
		homeChain,
		tokenDataObserver,
		ep,
		it.lggr,
		&metrics.Noop{},
		mockCodec,
	)

	// FIXME: Test should not rely on the specific type of the plugin but rather than that on
	// the interface type
	p := node1.(*TrackedPlugin)
	pp := p.ReportingPlugin.(*Plugin)

	return nodeSetup{
		node:        pp,
		reportCodec: reportCodec,
		msgHasher:   it.msgHasher,
	}
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
		if r.Method == http.MethodGet {
			for url, response := range c.responses {
				if strings.Contains(r.RequestURI, url) {
					_, err := w.Write([]byte(response))
					if err != nil {
						panic(err)
					}
					return
				}
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	c.server = server
	return c
}

func (c *ConfigurableAttestationServer) AddResponse(key, response string) {
	c.responses[key] = response
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

type msgOption func(*cciptypes.Message)

func withTokens(tokenAmounts ...cciptypes.RampTokenAmount) msgOption {
	return func(m *cciptypes.Message) {
		m.TokenAmounts = tokenAmounts
	}
}

func makeMsgWithMetadata(
	seqNum cciptypes.SeqNum,
	src, dest cciptypes.ChainSelector,
	executed bool,
	opts ...msgOption,
) inmem.MessagesWithMetadata {
	msg := cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			SourceChainSelector: src,
			SequenceNumber:      seqNum,
			MessageID:           rand.RandomBytes32(),
		},
		FeeValueJuels: cciptypes.NewBigIntFromInt64(100),
	}

	for _, opt := range opts {
		opt(&msg)
	}

	return inmem.MessagesWithMetadata{
		Message:     msg,
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

func extractSequenceNumbers(messages []cciptypes.Message) []cciptypes.SeqNum {
	sequenceNumbers := slicelib.Map(messages, func(m cciptypes.Message) cciptypes.SeqNum {
		return m.Header.SequenceNumber
	})
	return sequenceNumbers
}

func emptyMessagesForRange(start, end uint64) []cciptypes.Message {
	messages := make([]cciptypes.Message, end-start+1)
	for i := start; i <= end; i++ {
		messages[i-start] = cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				MessageID:      cciptypes.Bytes32{byte(i)},
				SequenceNumber: cciptypes.SeqNum(i),
			},
		}
	}
	return messages
}

func emptyMessagesMapForRange(start, end uint64) map[cciptypes.SeqNum]cciptypes.Message {
	messages := make(map[cciptypes.SeqNum]cciptypes.Message)
	for i := start; i <= end; i++ {
		messages[cciptypes.SeqNum(i)] = cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				MessageID:      cciptypes.Bytes32{byte(i)},
				SequenceNumber: cciptypes.SeqNum(i),
			},
		}
	}
	return messages
}

func emptyMessagesMapForRanges(ranges []cciptypes.SeqNumRange) map[cciptypes.SeqNum]cciptypes.Message {
	messages := make(map[cciptypes.SeqNum]cciptypes.Message)
	for _, r := range ranges {
		for i := r.Start(); i <= r.End(); i++ {
			messages[i] = cciptypes.Message{
				Header: cciptypes.RampMessageHeader{
					MessageID:      cciptypes.Bytes32{byte(i)},
					SequenceNumber: i,
				},
			}
		}
	}
	return messages
}

func NewMessage(
	msgID int,
	seqNum int,
	sourceChainSelector int,
	destChainSelector int) cciptypes.Message {
	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:           cciptypes.Bytes32{byte(msgID)},
			SourceChainSelector: cciptypes.ChainSelector(sourceChainSelector),
			DestChainSelector:   cciptypes.ChainSelector(destChainSelector),
			SequenceNumber:      cciptypes.SeqNum(seqNum),
		},
	}
}
