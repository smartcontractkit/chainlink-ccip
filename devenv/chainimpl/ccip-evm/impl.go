package ccip_evm

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	pingpongdapp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/ping_pong_dapp"
	evmseqs "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/devenv/changesets"
	"github.com/smartcontractkit/chainlink-ccip/devenv/sequences"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
)

var ccipMessageSentTopic = onramp.OnRampCCIPMessageSent{}.Topic()

type SourceDestPair struct {
	SourceChainSelector uint64
	DestChainSelector   uint64
}

type AnyMsgSentEvent struct {
	SequenceNumber uint64
	// RawEvent contains the raw event depending on the chain:
	//  EVM:   *onramp.OnRampCCIPMessageSent
	//  Aptos: module_onramp.CCIPMessageSent
	RawEvent any
}

type CCIP16EVM struct {
	e                      *deployment.Environment
	chainDetailsBySelector map[uint64]chainsel.ChainDetails
	ethClients             map[uint64]*ethclient.Client
	expectedSeqNumRange    map[SourceDestPair]ccipocr3common.SeqNumRange
	expectedSeqNumExec     map[SourceDestPair][]uint64
	msgSentEvents          []*AnyMsgSentEvent
}

func NewEmptyCCIP16EVM() *CCIP16EVM {
	return &CCIP16EVM{
		chainDetailsBySelector: make(map[uint64]chainsel.ChainDetails),
		ethClients:             make(map[uint64]*ethclient.Client),
		expectedSeqNumRange:    make(map[SourceDestPair]ccipocr3common.SeqNumRange),
		expectedSeqNumExec:     make(map[SourceDestPair][]uint64),
		msgSentEvents:          make([]*AnyMsgSentEvent, 0),
	}
}

// NewCCIP16EVM creates new smart-contracts wrappers with utility functions for CCIP16EVM implementation.
func NewCCIP16EVM(ctx context.Context, e *deployment.Environment) (*CCIP16EVM, error) {
	_ = zerolog.Ctx(ctx)
	out := NewEmptyCCIP16EVM()
	out.e = e
	return out, nil
}

func (m *CCIP16EVM) SetCLDF(e *deployment.Environment) {
	m.e = e
}

func updatePrices(datastore datastore.DataStore, src, dest uint64, srcChain cldf_evm.Chain) error {
	a := &evmseqs.EVMAdapter{}
	fqAddr, err := a.GetFQAddress(datastore, src)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	fq, err := fee_quoter.NewFeeQuoter(
		common.BytesToAddress(fqAddr),
		srcChain.Client)
	if err != nil {
		return fmt.Errorf("failed to create fee quoter instance: %w", err)
	}
	feeTokens, err := fq.GetFeeTokens(nil)
	if err != nil {
		return fmt.Errorf("failed to get fee tokens from fee quoter: %w", err)
	}
	sender := srcChain.DeployerKey
	tx, err := fq.UpdatePrices(sender, fee_quoter.InternalPriceUpdates{
		TokenPriceUpdates: []fee_quoter.InternalTokenPriceUpdate{
			{
				SourceToken: feeTokens[0],
				UsdPerToken: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2000)),
			},
			{
				SourceToken: feeTokens[1],
				UsdPerToken: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(2000)),
			},
		},
		GasPriceUpdates: []fee_quoter.InternalGasPriceUpdate{
			{
				DestChainSelector: dest,
				UsdPerUnitGas:     big.NewInt(20000e9),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to update prices: %w", err)
	}
	_, err = srcChain.Confirm(tx)
	if err != nil {
		return fmt.Errorf("failed to confirm update prices transaction: %w", err)
	}
	return nil
}

func (m *CCIP16EVM) SendMessage(ctx context.Context, src, dest uint64, fields any, opts any) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Sending CCIP message")
	a := &evmseqs.EVMAdapter{}
	receiver := common.LeftPadBytes(common.HexToAddress("0xdead").Bytes(), 32)
	msg := router.ClientEVM2AnyMessage{
		Receiver:     receiver,
		Data:         []byte("hello eoa"),
		TokenAmounts: nil,
		FeeToken:     common.HexToAddress("0x0"),
		ExtraArgs:    nil,
	}
	const errCodeInsufficientFee = "0x07da6ee6"
	const cannotDecodeErrorReason = "could not decode error reason"
	const errMsgMissingTrieNode = "missing trie node"
	sender := m.e.BlockChains.EVMChains()[src].DeployerKey
	defer func() { sender.Value = nil }()
	rAddr, err := a.GetRouterAddress(m.e.DataStore, src)
	if err != nil {
		return fmt.Errorf("failed to get router address: %w", err)
	}
	r, err := router.NewRouter(
		common.BytesToAddress(rAddr),
		m.e.BlockChains.EVMChains()[src].Client)
	if err != nil {
		return fmt.Errorf("failed to create router instance: %w", err)
	}
	onRampAddr, err := r.GetOnRamp(nil, dest)
	if err != nil {
		return fmt.Errorf("failed to get onramp address: %w", err)
	}
	onRamp, err := onramp.NewOnRamp(
		onRampAddr,
		m.e.BlockChains.EVMChains()[src].Client)
	if err != nil {
		return fmt.Errorf("failed to create onramp instance: %w", err)
	}
	l.Info().Msg("Got contract instances, preparing to send CCIP message")
	err = updatePrices(m.e.DataStore, src, dest, m.e.BlockChains.EVMChains()[src])
	if err != nil {
		return fmt.Errorf("failed to update prices: %w", err)
	}

	var retryCount int
	for {
		fee, err := r.GetFee(&bind.CallOpts{Context: ctx}, dest, msg)
		if err != nil {
			return fmt.Errorf("failed to get EVM fee: %w", deployment.MaybeDataErr(err))
		}

		sender.Value = fee

		tx, err := r.CcipSend(sender, dest, msg)
		if err != nil {
			return fmt.Errorf("failed to send CCIP message: %w", err)
		}

		blockNum, err := m.e.BlockChains.EVMChains()[src].Confirm(tx)
		if err != nil {
			if strings.Contains(err.Error(), errCodeInsufficientFee) {
				// Don't count insufficient fee as part of the retry count
				// because this is expected and we need to adjust the fee
				continue
			} else if strings.Contains(err.Error(), cannotDecodeErrorReason) ||
				strings.Contains(err.Error(), errMsgMissingTrieNode) {
				// If the error reason cannot be decoded, we retry to avoid transient issues. The retry behavior is disabled by default
				// It is configured in the CCIPSendReqConfig.
				// This retry was originally added to solve transient failure in end to end tests
				if retryCount >= 5 {
					return fmt.Errorf("failed to confirm CCIP message after %d retries: %w", retryCount, deployment.MaybeDataErr(err))
				}
				retryCount++
				continue
			}

			return fmt.Errorf("failed to confirm CCIP message: %w", deployment.MaybeDataErr(err))
		}
		it, err := onRamp.FilterCCIPMessageSent(&bind.FilterOpts{
			Start:   blockNum,
			End:     &blockNum,
			Context: context.Background(),
		}, []uint64{dest}, []uint64{})
		if err != nil {
			return fmt.Errorf("failed to filter CCIPMessageSent events: %w", err)
		}

		if !it.Next() {
			return fmt.Errorf("no CCIP message sent event found")
		}

		sourceDest := SourceDestPair{SourceChainSelector: src, DestChainSelector: dest}
		m.msgSentEvents = append(m.msgSentEvents, &AnyMsgSentEvent{
			SequenceNumber: it.Event.SequenceNumber,
			RawEvent:       it.Event,
		})
		m.expectedSeqNumRange[sourceDest] = ccipocr3common.SeqNumRange{
			ccipocr3common.SeqNum(m.msgSentEvents[0].SequenceNumber),
			ccipocr3common.SeqNum(m.msgSentEvents[len(m.msgSentEvents)-1].SequenceNumber)}
		m.expectedSeqNumExec[sourceDest] = append(
			m.expectedSeqNumExec[sourceDest],
			it.Event.SequenceNumber)

		return nil
	}
}

func (m *CCIP16EVM) GetExpectedNextSequenceNumber(ctx context.Context, from, to uint64) (uint64, error) {
	_ = zerolog.Ctx(ctx)
	sourceDest := SourceDestPair{SourceChainSelector: from, DestChainSelector: to}
	seqRange, ok := m.expectedSeqNumRange[sourceDest]
	if !ok {
		return 0, fmt.Errorf("no expected sequence number range for source-dest pair %v", sourceDest)
	}
	return uint64(seqRange.End()), nil
}

type CommitReportTracker struct {
	seenMessages map[uint64]map[uint64]bool
}

func NewCommitReportTracker(sourceChainSelector uint64, seqNrs ccipocr3common.SeqNumRange) CommitReportTracker {
	seenMessages := make(map[uint64]map[uint64]bool)
	seenMessages[sourceChainSelector] = make(map[uint64]bool)

	for i := seqNrs.Start(); i <= seqNrs.End(); i++ {
		seenMessages[sourceChainSelector][uint64(i)] = false
	}
	return CommitReportTracker{seenMessages: seenMessages}
}

func (c *CommitReportTracker) visitCommitReport(sourceChainSelector uint64, minSeqNr uint64, maxSeqNr uint64) {
	if _, ok := c.seenMessages[sourceChainSelector]; !ok {
		return
	}

	for i := minSeqNr; i <= maxSeqNr; i++ {
		c.seenMessages[sourceChainSelector][i] = true
	}
}

func (c *CommitReportTracker) allCommited(sourceChainSelector uint64) bool {
	for _, v := range c.seenMessages[sourceChainSelector] {
		if !v {
			return false
		}
	}
	return true
}

// WaitOneSentEventBySeqNo wait and fetch strictly one CCIPMessageSent event by selector and sequence number and selector.
func (m *CCIP16EVM) WaitOneSentEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error) {
	l := zerolog.Ctx(ctx)
	a := &evmseqs.EVMAdapter{}
	offAddr, err := a.GetOffRampAddress(m.e.DataStore, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramp address: %w", err)
	}
	offRamp, err := offramp.NewOffRamp(
		common.BytesToAddress(offAddr),
		m.e.BlockChains.EVMChains()[to].Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create off ramp instance: %w", err)
	}
	sourceDest := SourceDestPair{SourceChainSelector: from, DestChainSelector: to}
	seqRange, ok := m.expectedSeqNumRange[sourceDest]
	if !ok {
		return nil, fmt.Errorf("no expected sequence number range for source-dest pair %v", sourceDest)
	}

	seenMessages := NewCommitReportTracker(from, seqRange)

	verifyCommitReport := func(report *offramp.OffRampCommitReportAccepted) bool {
		processRoots := func(roots []offramp.InternalMerkleRoot) bool {
			for _, mr := range roots {
				l.Info().Msgf(
					"Received commit report for [%d, %d] on selector %d from source selector %d expected seq nr range %s, token prices: %v",
					mr.MinSeqNr, mr.MaxSeqNr, to, from, seqRange.String(), report.PriceUpdates.TokenPriceUpdates,
				)
				fmt.Printf(
					"Received commit report for [%d, %d] on selector %d from source selector %d expected seq nr range %s, token prices: %v\n",
					mr.MinSeqNr, mr.MaxSeqNr, to, from, seqRange.String(), report.PriceUpdates.TokenPriceUpdates,
				)
				seenMessages.visitCommitReport(from, mr.MinSeqNr, mr.MaxSeqNr)

				if mr.SourceChainSelector == from &&
					uint64(seqRange.Start()) >= mr.MinSeqNr &&
					uint64(seqRange.End()) <= mr.MaxSeqNr {
					l.Info().Msgf(
						"All sequence numbers committed in a single report [%d, %d]",
						seqRange.Start(), seqRange.End(),
					)
					return true
				}

				if seenMessages.allCommited(from) {
					l.Info().Msgf(
						"All sequence numbers already committed from range [%d, %d]",
						seqRange.Start(), seqRange.End(),
					)
					return true
				}
			}
			return false
		}

		return processRoots(report.BlessedMerkleRoots) || processRoots(report.UnblessedMerkleRoots)
	}

	// defer subscription.Unsubscribe()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, nil
		case <-ticker.C:
			l.Info().Msgf("Waiting for commit report on chain selector %d from source selector %d expected seq nr range %s",
				to, from, seqRange.String())

			// Need to do this because the subscription sometimes fails to get the event.
			iter, err := offRamp.FilterCommitReportAccepted(&bind.FilterOpts{
				Context: ctx,
			})

			// In some test case the test ends while the filter is still running resulting in a context.Canceled error.
			if err != nil && !errors.Is(err, context.Canceled) {
				return nil, fmt.Errorf("error filtering CommitReportAccepted: %w", err)
			}
			for iter.Next() {
				event := iter.Event
				verified := verifyCommitReport(event)
				if verified {
					return event, nil
				}
			}
		case <-timer.C:
			return nil, fmt.Errorf("timed out after waiting for commit report on chain selector %d from source selector %d expected seq nr range %s",
				to, from, seqRange.String())
		}
	}
}

const (
	EXECUTION_STATE_UNTOUCHED  = 0
	EXECUTION_STATE_INPROGRESS = 1
	EXECUTION_STATE_SUCCESS    = 2
	EXECUTION_STATE_FAILURE    = 3
)

func executionStateToString(state uint8) string {
	switch state {
	case EXECUTION_STATE_UNTOUCHED:
		return "UNTOUCHED"
	case EXECUTION_STATE_INPROGRESS:
		return "IN_PROGRESS"
	case EXECUTION_STATE_SUCCESS:
		return "SUCCESS"
	case EXECUTION_STATE_FAILURE:
		return "FAILURE"
	default:
		return "UNKNOWN"
	}
}

// WaitOneExecEventBySeqNo wait and fetch strictly one ExecutionStateChanged event by sequence number and selector.
func (m *CCIP16EVM) WaitOneExecEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error) {
	l := zerolog.Ctx(ctx)
	a := &evmseqs.EVMAdapter{}
	offAddr, err := a.GetOffRampAddress(m.e.DataStore, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramp address: %w", err)
	}
	offRamp, err := offramp.NewOffRamp(
		common.BytesToAddress(offAddr),
		m.e.BlockChains.EVMChains()[to].Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create off ramp instance: %w", err)
	}
	sourceDest := SourceDestPair{SourceChainSelector: from, DestChainSelector: to}
	seqRange, ok := m.expectedSeqNumRange[sourceDest]
	if !ok {
		return nil, fmt.Errorf("no expected sequence number range for source-dest pair %v", sourceDest)
	}

	executionStates := make(map[uint64]int)
	seqNrsToWatch := make(map[uint64]struct{})
	for seqNr := seqRange.Start(); seqNr <= seqRange.End(); seqNr++ {
		seqNrsToWatch[uint64(seqNr)] = struct{}{}
	}

	// defer subscription.Unsubscribe()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, nil
		case <-ticker.C:
			for expectedSeqNr := range seqNrsToWatch {
				scc, executionState := getExecutionState(from, offRamp, expectedSeqNr)
				l.Info().Msgf("Waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence number %d, current onchain minSeqNr: %d, execution state: %s",
					to, offRamp.Address().String(), from, expectedSeqNr, scc.MinSeqNr, executionStateToString(executionState))
				fmt.Printf("Waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence number %d, current onchain minSeqNr: %d, execution state: %s\n",
					to, offRamp.Address().String(), from, expectedSeqNr, scc.MinSeqNr, executionStateToString(executionState))
				if executionState == EXECUTION_STATE_SUCCESS || executionState == EXECUTION_STATE_FAILURE {
					l.Info().Msgf("Observed %s execution state on chain %d (offramp %s) from chain %d with expected sequence number %d",
						executionStateToString(executionState), to, offRamp.Address().String(), from, expectedSeqNr)
					fmt.Printf("Observed %s execution state on chain %d (offramp %s) from chain %d with expected sequence number %d\n",
						executionStateToString(executionState), to, offRamp.Address().String(), from, expectedSeqNr)
					executionStates[expectedSeqNr] = int(executionState)
					delete(seqNrsToWatch, expectedSeqNr)
					if len(seqNrsToWatch) == 0 {
						return executionStates, nil
					}
				}
			}
		case <-timer.C:
			return nil, fmt.Errorf("timed out waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence numbers %+v",
				to, offRamp.Address().String(), from, seqNrsToWatch)
		}
	}
}

func getExecutionState(sourceSelector uint64, offRamp offramp.OffRampInterface, expectedSeqNr uint64) (offramp.OffRampSourceChainConfig, uint8) {
	scc, err := offRamp.GetSourceChainConfig(nil, sourceSelector)
	if err != nil {
		panic(fmt.Errorf("failed to get source chain config: %w", err))
	}
	executionState, err := offRamp.GetExecutionState(nil, sourceSelector, expectedSeqNr)
	if err != nil {
		panic(fmt.Errorf("failed to get execution state: %w", err))
	}
	return scc, executionState
}

func (m *CCIP16EVM) GetEOAReceiverAddress(ctx context.Context, chainSelector uint64) ([]byte, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

func (m *CCIP16EVM) GetTokenBalance(ctx context.Context, chainSelector uint64, address, tokenAddress []byte) (*big.Int, error) {
	_ = zerolog.Ctx(ctx)
	return big.NewInt(0), nil
}

func (m *CCIP16EVM) ExposeMetrics(
	ctx context.Context,
	source, dest uint64,
	chainIDs []string,
	wsURLs []string,
) ([]string, *prometheus.Registry, error) {
	msgSentTotal.Reset()
	msgExecTotal.Reset()
	srcDstLatency.Reset()

	reg := prometheus.NewRegistry()
	reg.MustRegister(msgSentTotal, msgExecTotal, srcDstLatency)

	lp := NewLokiPusher()
	err := ProcessLaneEvents(ctx, m, lp, &LaneStreamConfig{
		FromSelector:      source,
		ToSelector:        dest,
		AggregatorAddress: "localhost:50051",
		AggregatorSince:   0,
	})
	if err != nil {
		return nil, nil, err
	}
	err = ProcessLaneEvents(ctx, m, lp, &LaneStreamConfig{
		FromSelector:      dest,
		ToSelector:        source,
		AggregatorAddress: "localhost:50051",
		AggregatorSince:   0,
	})
	if err != nil {
		return nil, nil, err
	}
	return []string{}, reg, nil
}

func (m *CCIP16EVM) DeployLocalNetwork(ctx context.Context, bc *blockchain.Input) (*blockchain.Output, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Deploying EVM networks")
	out, err := blockchain.NewBlockchainNetwork(bc)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain network: %w", err)
	}
	return out, nil
}

func (m *CCIP16EVM) ConfigureNodes(ctx context.Context, bc *blockchain.Input) (string, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring CL nodes")
	name := fmt.Sprintf("node-evm-%s", uuid.New().String()[0:5])
	finality := 1
	return fmt.Sprintf(`
       [[EVM]]
       LogPollInterval = '1s'
       BlockBackfillDepth = 100
       ChainID = '%s'
       MinIncomingConfirmations = 1
       MinContractPayment = '0.0000001 link'
       FinalityDepth = %d

       [[EVM.Nodes]]
       Name = '%s'
       WsUrl = '%s'
       HttpUrl = '%s'`,
		bc.ChainID,
		finality,
		name,
		bc.Out.Nodes[0].InternalWSUrl,
		bc.Out.Nodes[0].InternalHTTPUrl,
	), nil
}

func (m *CCIP16EVM) DeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) (datastore.DataStore, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring contracts for selector")
	l.Info().Any("Selector", selector).Msg("Deploying for chain selectors")
	runningDS := datastore.NewMemoryDataStore()

	l.Info().Uint64("Selector", selector).Msg("Configuring per-chain contracts bundle")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		env.Logger,
		operations.NewMemoryReporter(),
	)
	env.OperationsBundle = bundle

	chain, ok := env.BlockChains.EVMChains()[selector]
	if !ok {
		return nil, fmt.Errorf("evm chain not found for selector %d", selector)
	}
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	out, err := deployops.DeployContracts(dReg).Apply(*env, deployops.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
			chain.Selector: {
				Version: version,
				// FEE QUOTER CONFIG
				MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
				TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
				LinkPremiumMultiplier:        9e17, // 0.9 ETH
				NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
				// OFFRAMP CONFIG
				PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
				GasForCallExactCheck:                    uint16(5000),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deploy contracts: %w", err)
	}
	nodeClients, err := clclient.New(cls[0].Out.CLNodes)
	if err != nil {
		return nil, fmt.Errorf("connecting to CL nodes: %w", err)
	}
	// bootstrap is 0
	workerNodes := nodeClients[1:]
	if selector == ccipHomeSelector {
		var readers [][32]byte
		for _, node := range workerNodes {
			nodeP2PIds, err := node.MustReadP2PKeys()
			if err != nil {
				return nil, fmt.Errorf("reading worker node P2P keys: %w", err)
			}
			l.Info().Str("Node", node.Config.URL).Str("PeerID", nodeP2PIds.Data[0].Attributes.PeerID).Msg("Adding reader peer ID")
			id := changesets.MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
			readers = append(readers, id)
			l.Info().Msgf("peerID: %+v", id)
			l.Info().Msgf("peer ID from bytes: %s", id.Raw())
		}
		ccipHomeOut, err := changesets.DeployHomeChain.Apply(*env, sequences.DeployHomeChainConfig{
			HomeChainSel: selector,
			CapReg:       common.HexToAddress(crAddr),
			RMNStaticConfig: rmn_home.RMNHomeStaticConfig{
				Nodes:          []rmn_home.RMNHomeNode{},
				OffchainConfig: []byte("static config"),
			},
			RMNDynamicConfig: rmn_home.RMNHomeDynamicConfig{
				SourceChains:   []rmn_home.RMNHomeSourceChain{},
				OffchainConfig: []byte("dynamic config"),
			},
			NodeOperators: []capabilities_registry.CapabilitiesRegistryNodeOperator{
				{
					Admin: chain.DeployerKey.From,
					Name:  "NodeOperator",
				},
			},
			NodeP2PIDsPerNodeOpAdmin: map[string][][32]byte{"NodeOperator": readers},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to deploy home chain contracts: %w", err)
		}
		out.DataStore.Merge(ccipHomeOut.DataStore.Seal())
		out.DataStore.Addresses().Add(
			datastore.AddressRef{
				ChainSelector: selector,
				Type:          datastore.ContractType(utils.CapabilitiesRegistry),
				Version:       semver.MustParse("1.6.0"),
				Address:       crAddr,
			},
		)
	}

	env.DataStore = out.DataStore.Seal()
	runningDS.Merge(env.DataStore)

	return runningDS.Seal(), nil
}

func (m *CCIP16EVM) ConnectContractsWithSelectors(ctx context.Context, e *deployment.Environment, selector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("FromSelector", selector).Any("ToSelectors", remoteSelectors).Msg("Connecting contracts with selectors")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	// we're assuming all dest chains are EVM for this implementation
	evmEncoded, err := hex.DecodeString(cciputils.EVMFamilySelector)
	if err != nil {
		return fmt.Errorf("encoding EVM family selector: %w", err)
	}
	mcmsRegistry := changesetscore.GetRegistry()
	version := semver.MustParse("1.6.0")
	chainA := lanesapi.ChainDefinition{
		Selector:                 selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, evmEncoded),
	}
	for _, destSelector := range remoteSelectors {
		chainB := lanesapi.ChainDefinition{
			Selector:                 destSelector,
			GasPrice:                 big.NewInt(1e9),
			FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, evmEncoded),
		}
		l.Info().Uint64("ChainASelector", chainA.Selector).Uint64("ChainBSelector", chainB.Selector).Msg("Connecting chain pairs")
		_, err = lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
			Lanes: []lanesapi.LaneConfig{
				{
					Version: version,
					ChainA:  chainA,
					ChainB:  chainB,
				},
			},
		})
		if err != nil {
			return fmt.Errorf("connecting chains %d and %d: %w", chainA.Selector, chainB.Selector, err)
		}
		err = updatePrices(e.DataStore, selector, destSelector, e.BlockChains.EVMChains()[selector])
		if err != nil {
			return fmt.Errorf("failed to update prices: %w", err)
		}
		err = updatePrices(e.DataStore, destSelector, selector, e.BlockChains.EVMChains()[destSelector])
		if err != nil {
			return fmt.Errorf("failed to update prices: %w", err)
		}
	}

	return nil
}

func (m *CCIP16EVM) LinkPingPongContracts(ctx context.Context, e *deployment.Environment, selector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("FromSelector", selector).Any("ToSelectors", remoteSelectors).Msg("Linking PingPongDemo contracts")

	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	a := &evmseqs.EVMAdapter{}

	// Get the PingPongDemo address on this chain
	localPingPongAddr, err := a.GetPingPongDemoAddress(e.DataStore, selector)
	if err != nil {
		return fmt.Errorf("failed to get PingPongDemo address for selector %d: %w", selector, err)
	}

	chain := e.BlockChains.EVMChains()[selector]

	for _, remoteSelector := range remoteSelectors {
		// Get the PingPongDemo address on the remote chain
		remotePingPongAddr, err := a.GetPingPongDemoAddress(e.DataStore, remoteSelector)
		if err != nil {
			return fmt.Errorf("failed to get PingPongDemo address for remote selector %d: %w", remoteSelector, err)
		}

		// CCIP requires addresses to be 32-byte left-padded for cross-chain messaging
		paddedRemoteAddr := common.LeftPadBytes(remotePingPongAddr, 32)

		l.Info().
			Uint64("LocalSelector", selector).
			Uint64("RemoteSelector", remoteSelector).
			Str("LocalPingPong", common.BytesToAddress(localPingPongAddr).Hex()).
			Str("RemotePingPong", common.BytesToAddress(remotePingPongAddr).Hex()).
			Msg("Setting counterpart for PingPongDemo")

		// Set the counterpart on the local PingPongDemo contract
		_, err = operations.ExecuteOperation(bundle, pingpongdapp.SetCounterpart, chain, contract.FunctionInput[pingpongdapp.SetCounterpartArgs]{
			ChainSelector: selector,
			Address:       common.BytesToAddress(localPingPongAddr),
			Args: pingpongdapp.SetCounterpartArgs{
				CounterpartChainSelector: remoteSelector,
				CounterpartAddress:       paddedRemoteAddr,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to set counterpart on PingPongDemo for selector %d -> %d: %w", selector, remoteSelector, err)
		}
	}

	return nil
}

func (m *CCIP16EVM) ConfigureContractsForSelectors(ctx context.Context, e *deployment.Environment, cls []*simple_node_set.Input, ccipHomeSelector uint64, remoteSelectors []uint64) error {
	l := zerolog.Ctx(ctx)
	l.Info().Uint64("HomeChainSelector", ccipHomeSelector).Msg("Configuring contracts for home chain selector")
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	// Build the CCIPHome chain configs.
	chainConfigs := make(map[uint64]changesets.ChainConfig)
	commitOCRConfigs := make(map[uint64]changesets.CCIPOCRParams)
	execOCRConfigs := make(map[uint64]changesets.CCIPOCRParams)
	nodeClients, err := clclient.New(cls[0].Out.CLNodes)
	if err != nil {
		return fmt.Errorf("connecting to CL nodes: %w", err)
	}
	// bootstrap is 0
	workerNodes := nodeClients[1:]
	var readers [][32]byte
	for _, node := range workerNodes {
		nodeP2PIds, err := node.MustReadP2PKeys()
		if err != nil {
			return fmt.Errorf("reading worker node P2P keys: %w", err)
		}
		l.Info().Str("Node", node.Config.URL).Str("PeerID", nodeP2PIds.Data[0].Attributes.PeerID).Msg("Adding reader peer ID")
		id := changesets.MustPeerIDFromString(nodeP2PIds.Data[0].Attributes.PeerID)
		readers = append(readers, id)
	}
	for _, chain := range remoteSelectors {
		ocrOverride := func(ocrParams changesets.CCIPOCRParams) changesets.CCIPOCRParams {
			if ocrParams.CommitOffChainConfig != nil {
				ocrParams.CommitOffChainConfig.RMNEnabled = false
			}
			return ocrParams
		}
		commitOCRConfigs[chain] = changesets.DeriveOCRParamsForCommit(changesets.SimulationTest, ccipHomeSelector, nil, ocrOverride)
		execOCRConfigs[chain] = changesets.DeriveOCRParamsForExec(changesets.SimulationTest, nil, ocrOverride)

		l.Info().Msgf("setting readers for chain %d to %v due to no topology", chain, len(readers))
		chainConfigs[chain] = changesets.ChainConfig{
			Readers: readers,
			FChain:  uint8(len(readers) / 3),
			EncodableChainConfig: chainconfig.ChainConfig{
				GasPriceDeviationPPB:      ccipocr3common.BigInt{Int: big.NewInt(1000)},
				DAGasPriceDeviationPPB:    ccipocr3common.BigInt{Int: big.NewInt(1000)},
				OptimisticConfirmations:   changesets.OptimisticConfirmations,
				ChainFeeDeviationDisabled: false,
			},
		}
	}

	_, err = changesets.UpdateChainConfig.Apply(*e, changesets.UpdateChainConfigConfig{
		HomeChainSelector: ccipHomeSelector,
		RemoteChainAdds:   chainConfigs,
	})
	if err != nil {
		return fmt.Errorf("updating chain config for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = changesets.AddDONAndSetCandidate.Apply(*e, changesets.AddDonAndSetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: changesets.SetCandidatePluginInfo{
			OCRConfigPerRemoteChainSelector: commitOCRConfigs,
			PluginType:                      ccipocr3common.PluginTypeCCIPCommit,
		},
		NonBootstraps: workerNodes,
	})
	if err != nil {
		return fmt.Errorf("adding DON and setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = changesets.SetCandidate.Apply(*e, changesets.SetCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		FeedChainSelector: ccipHomeSelector,
		PluginInfo: []changesets.SetCandidatePluginInfo{
			{
				OCRConfigPerRemoteChainSelector: execOCRConfigs,
				PluginType:                      ccipocr3common.PluginTypeCCIPExec,
			},
		},
		NonBootstraps: workerNodes,
	})
	if err != nil {
		return fmt.Errorf("setting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	_, err = changesets.PromoteCandidate.Apply(*e, changesets.PromoteCandidateChangesetConfig{
		HomeChainSelector: ccipHomeSelector,
		PluginInfo: []changesets.PromoteCandidatePluginInfo{
			{
				PluginType:           ccipocr3common.PluginTypeCCIPCommit,
				RemoteChainSelectors: remoteSelectors,
			},
			{
				PluginType:           ccipocr3common.PluginTypeCCIPExec,
				RemoteChainSelectors: remoteSelectors,
			},
		},
		NonBootstraps: workerNodes,
	})
	if err != nil {
		return fmt.Errorf("promoting candidate for selector %d: %w", ccipHomeSelector, err)
	}
	dReg := deployops.GetRegistry()
	mcmsRegistry := changesetscore.GetRegistry()
	_, err = deployops.SetOCR3Config(dReg, mcmsRegistry).Apply(*e, deployops.SetOCR3ConfigArgs{
		HomeChainSel:    ccipHomeSelector,
		RemoteChainSels: remoteSelectors,
		ConfigType:      cciputils.ConfigTypeActive,
	})
	if err != nil {
		return fmt.Errorf("setting OCR3 config for selector %d: %w", ccipHomeSelector, err)
	}
	return nil
}

func (m *CCIP16EVM) FundNodes(ctx context.Context, ns []*simple_node_set.Input, bc *blockchain.Input, linkAmount, nativeAmount *big.Float) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Funding CL nodes with ETH and LINK")
	nodeClients, err := clclient.New(ns[0].Out.CLNodes)
	if err != nil {
		return fmt.Errorf("connecting to CL nodes: %w", err)
	}
	ethKeyAddressesSrc := make([]string, 0)
	for i, nc := range nodeClients {
		addrSrc, err := nc.ReadPrimaryETHKey(bc.ChainID)
		if err != nil {
			return fmt.Errorf("getting primary ETH key from OCR node %d (src chain): %w", i, err)
		}
		ethKeyAddressesSrc = append(ethKeyAddressesSrc, addrSrc.Attributes.Address)
		l.Info().
			Int("Idx", i).
			Str("ETHKeySrc", addrSrc.Attributes.Address).
			Msg("Node info")
	}
	clientSrc, _, _, err := ETHClient(ctx, bc.Out.Nodes[0].ExternalWSUrl, &GasSettings{
		FeeCapMultiplier: 2,
		TipCapMultiplier: 2,
	})
	if err != nil {
		return fmt.Errorf("could not create basic eth client: %w", err)
	}
	// Use default Anvil key for local chain 1337, otherwise use PRIVATE_KEY env var
	privateKey := getNetworkPrivateKey()
	if bc.ChainID == "1337" {
		privateKey = DefaultAnvilKey
	}

	// nativeAmount is in ETH units - FundNodeEIP1559 converts to wei internally
	nativeAmountFloat, _ := nativeAmount.Float64()
	for _, addr := range ethKeyAddressesSrc {
		if err := FundNodeEIP1559(ctx, clientSrc, privateKey, addr, nativeAmountFloat); err != nil {
			return fmt.Errorf("failed to fund CL nodes on src chain: %w", err)
		}
	}
	return nil
}
