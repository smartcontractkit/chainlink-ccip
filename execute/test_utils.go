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

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/gas/evm"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

type IntTest struct {
	t *testing.T

	donID uint32

	srcSelector cciptypes.ChainSelector
	dstSelector cciptypes.ChainSelector

	msgHasher           cciptypes.MessageHasher
	ccipReader          *inmem.InMemoryCCIPReader
	server              *ConfigurableAttestationServer
	tokenObserverConfig []pluginconfig.TokenDataObserverConfig
	tokenChainReader    map[cciptypes.ChainSelector]contractreader.ContractReaderFacade
	feeCalculator       *exectypes.CCIPMessageFeeUSD18Calculator
	execCostCalculator  *exectypes.StaticMessageExecCostUSD18Calculator
}

func SetupSimpleTest(t *testing.T, srcSelector, dstSelector cciptypes.ChainSelector) *IntTest {
	donID := uint32(1)

	msgHasher := mocks.NewMessageHasher()
	ccipReader := inmem.InMemoryCCIPReader{
		Reports: []plugintypes2.CommitPluginReportWithMeta{},
		Messages: map[cciptypes.ChainSelector][]inmem.MessagesWithMetadata{
			srcSelector: {},
		},
		Dest: dstSelector,
	}

	return &IntTest{
		t:                   t,
		donID:               donID,
		msgHasher:           msgHasher,
		srcSelector:         srcSelector,
		dstSelector:         dstSelector,
		ccipReader:          &ccipReader,
		tokenObserverConfig: []pluginconfig.TokenDataObserverConfig{},
		tokenChainReader:    map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{},
	}
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

func (it *IntTest) WithCustomFeeBoosting(
	relativeBoostPerWaitHour float64,
	now func() time.Time,
	messageCost map[cciptypes.Bytes32]plugintypes.USD18,
) {
	it.feeCalculator = exectypes.NewCCIPMessageFeeUSD18Calculator(
		logger.Test(it.t),
		it.ccipReader,
		relativeBoostPerWaitHour,
		now,
	)
	it.execCostCalculator = exectypes.NewStaticMessageExecCostUSD18Calculator(messageCost)
}

func (it *IntTest) WithUSDC(
	sourcePoolAddress string,
	attestations map[string]string,
	events []*readerpkg.MessageSentEvent,
) {
	it.server = newConfigurableAttestationServer(attestations)
	it.tokenObserverConfig = []pluginconfig.TokenDataObserverConfig{
		{
			Type:    "usdc-cctp",
			Version: "1",
			USDCCCTPObserverConfig: &pluginconfig.USDCCCTPObserverConfig{
				Tokens: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
					it.srcSelector: {
						SourcePoolAddress:            sourcePoolAddress,
						SourceMessageTransmitterAddr: sourcePoolAddress,
					},
				},
				AttestationAPI:         it.server.server.URL,
				AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Millisecond),
				AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Second),
			},
		},
	}

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
		it.srcSelector: r,
		it.dstSelector: r,
	}
}

//nolint:lll
func (it *IntTest) Start() *testhelpers.OCR3Runner[[]byte] {
	cfg := pluginconfig.ExecuteOffchainConfig{
		MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
		BatchGasLimit:             100000000,
	}
	chainConfigInfos := []reader.ChainConfigInfo{
		{
			ChainSelector: it.srcSelector,
			ChainConfig: reader.HomeChainConfigMapper{
				FChain: 1,
				Readers: []libocrtypes.PeerID{
					{1}, {2}, {3},
				},
				Config: mustEncodeChainConfig(chainconfig.ChainConfig{}),
			},
		}, {
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

	homeChain := setupHomeChainPoller(it.t, it.donID, logger.Test(it.t), chainConfigInfos)
	ctx := tests.Context(it.t)
	err := homeChain.Start(ctx)
	require.NoError(it.t, err, "failed to start home chain poller")

	tkObs, err := tokendata.NewConfigBasedCompositeObservers(
		ctx,
		logger.Test(it.t),
		it.dstSelector,
		it.tokenObserverConfig,
		testhelpers.TokenDataEncoderInstance,
		it.tokenChainReader,
	)
	require.NoError(it.t, err)

	var feeCalculator exectypes.MessageFeeE18USDCalculator
	if it.feeCalculator != nil {
		feeCalculator = it.feeCalculator
	} else {
		feeCalculator = &exectypes.ZeroMessageFeeUSD18Calculator{}
	}

	var execCostCalculator exectypes.MessageExecCostUSD18Calculator
	if it.execCostCalculator != nil {
		execCostCalculator = it.execCostCalculator
	} else {
		execCostCalculator = &exectypes.ZeroMessageExecCostUSD18Calculator{}
	}

	costlyMessageObserver := exectypes.NewCostlyMessageObserver(
		logger.Test(it.t),
		true,
		feeCalculator,
		execCostCalculator,
	)

	oracleIDToP2pID := testhelpers.CreateOracleIDToP2pID(1, 2, 3)
	nodesSetup := []nodeSetup{
		newNode(it.donID, logger.Test(it.t), cfg, it.dstSelector, it.msgHasher, it.ccipReader, homeChain, tkObs, costlyMessageObserver, oracleIDToP2pID, 1, 1),
		newNode(it.donID, logger.Test(it.t), cfg, it.dstSelector, it.msgHasher, it.ccipReader, homeChain, tkObs, costlyMessageObserver, oracleIDToP2pID, 2, 1),
		newNode(it.donID, logger.Test(it.t), cfg, it.dstSelector, it.msgHasher, it.ccipReader, homeChain, tkObs, costlyMessageObserver, oracleIDToP2pID, 3, 1),
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
	if it.server != nil {
		it.server.Close()
	}
}

func (it *IntTest) UpdateExecutionCost(id cciptypes.Bytes32, val int64) {
	it.execCostCalculator.UpdateCosts(id, plugintypes.NewUSD18(val))
}

func newNode(
	donID plugintypes.DonID,
	lggr logger.Logger,
	cfg pluginconfig.ExecuteOffchainConfig,
	destChain cciptypes.ChainSelector,
	msgHasher cciptypes.MessageHasher,
	ccipReader readerpkg.CCIPReader,
	homeChain reader.HomeChain,
	tokenDataObserver tokendata.TokenDataObserver,
	costlyMessageObserver exectypes.CostlyMessageObserver,
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
		destChain,
		oracleIDToP2pID,
		ccipReader,
		reportCodec,
		msgHasher,
		homeChain,
		tokenDataObserver,
		evm.EstimateProvider{},
		lggr,
		costlyMessageObserver,
	)

	return nodeSetup{
		node:        node1,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
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

type msgOption func(*cciptypes.Message)

func withFeeValueJuels(fee cciptypes.BigInt) msgOption {
	return func(m *cciptypes.Message) {
		m.FeeValueJuels = fee
	}
}

func withTokens(tokenAmounts []cciptypes.RampTokenAmount) msgOption {
	return func(m *cciptypes.Message) {
		m.TokenAmounts = tokenAmounts
	}
}

func makeMsg(seqNum cciptypes.SeqNum, src, dest cciptypes.ChainSelector, executed bool, opts ...msgOption) inmem.MessagesWithMetadata {
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
				*returnVal.(*reader.ActiveAndCandidate) = reader.ActiveAndCandidate{}
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

func extractSequenceNumbers(messages []cciptypes.Message) []cciptypes.SeqNum {
	sequenceNumbers := slicelib.Map(messages, func(m cciptypes.Message) cciptypes.SeqNum {
		return m.Header.SequenceNumber
	})
	return sequenceNumbers
}

type timeMachine struct {
	now time.Time
}

func (t *timeMachine) Now() time.Time {
	return t.now
}

func (t *timeMachine) SetNow(now time.Time) {
	t.now = now
}
