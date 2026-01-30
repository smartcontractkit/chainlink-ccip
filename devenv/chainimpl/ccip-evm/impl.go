package ccip_evm

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagliardetto/solana-go"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/xssnick/tonutils-go/address"

	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/devenv/blockchainutils"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-ton/pkg/ccip/codec"

	evmseqs "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	msg_hasher163 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/message_hasher"
	solccip "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	devenvcommon "github.com/smartcontractkit/chainlink-ccip/devenv/common"
)

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
	ExpectedSeqNumRange    map[SourceDestPair]ccipocr3common.SeqNumRange
	ExpectedSeqNumExec     map[SourceDestPair][]uint64
	MsgSentEvents          []*AnyMsgSentEvent
}

func NewEmptyCCIP16EVM() *CCIP16EVM {
	return &CCIP16EVM{
		chainDetailsBySelector: make(map[uint64]chainsel.ChainDetails),
		ethClients:             make(map[uint64]*ethclient.Client),
		ExpectedSeqNumRange:    make(map[SourceDestPair]ccipocr3common.SeqNumRange),
		ExpectedSeqNumExec:     make(map[SourceDestPair][]uint64),
		MsgSentEvents:          make([]*AnyMsgSentEvent, 0),
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
	var receiver []byte
	family, err := chainsel.GetSelectorFamily(dest)
	if err != nil {
		return fmt.Errorf("failed to get chain family: %w", err)
	}
	var extraArgs []byte
	switch family {
	case chainsel.FamilyEVM:
		receiver = common.LeftPadBytes(common.HexToAddress("0xdead").Bytes(), 32)
	case chainsel.FamilySolana:
		receiverAddr, err := datastore_utils.FindAndFormatRef(m.e.DataStore, datastore.AddressRef{
			ChainSelector: dest,
			Type:          datastore.ContractType("TestReceiver"),
		}, dest, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("failed to get TestReceiver address: %w", err)
		}
		receiverProgram := solana.MustPublicKeyFromBase58(receiverAddr.Address)
		receiver = receiverProgram.Bytes()
		receiverTargetAccountPDA, _, _ := solana.FindProgramAddress([][]byte{[]byte("counter")}, receiverProgram)
		receiverExternalExecutionConfigPDA, _, _ := solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, receiverProgram)
		accounts := [][32]byte{
			receiverExternalExecutionConfigPDA,
			receiverTargetAccountPDA,
			solana.SystemProgramID,
		}

		extraArgs, err = devenvcommon.SerializeClientSVMExtraArgsV1(msg_hasher163.ClientSVMExtraArgsV1{
			AccountIsWritableBitmap:  solccip.GenerateBitMapForIndexes([]int{0, 1}),
			Accounts:                 accounts,
			ComputeUnits:             80_000,
			AllowOutOfOrderExecution: true,
		})
		if err != nil {
			return fmt.Errorf("failed to serialize SVM extra args: %w", err)
		}
	case chainsel.FamilyTon:
		receiverAddr, err := datastore_utils.FindAndFormatRef(m.e.DataStore, datastore.AddressRef{
			ChainSelector: dest,
			Type:          datastore.ContractType("Receiver"),
		}, dest, datastore_utils.FullRef)
		if err != nil {
			return fmt.Errorf("failed to get TonReceiver address: %w", err)
		}
		tonreceiver, err := address.ParseAddr(receiverAddr.Address)
		if err != nil {
			return fmt.Errorf("failed to parse TON receiver address: %w", err)
		}
		ac := codec.NewAddressCodec()
		receiver, err = ac.AddressStringToBytes(tonreceiver.String())
		if err != nil {
			return fmt.Errorf("failed to convert TON address to bytes: %w", err)
		}
		extraArgs, err = devenvcommon.SerializeClientGenericExtraArgsV2(msg_hasher163.ClientGenericExtraArgsV2{
			GasLimit:                 new(big.Int).SetUint64(100_000_000),
			AllowOutOfOrderExecution: true,
		})
		if err != nil {
			return fmt.Errorf("failed to serialize TON extra args: %w", err)
		}
	default:
		return fmt.Errorf("unsupported chain family: %s", family)
	}
	msg := router.ClientEVM2AnyMessage{
		Receiver:     receiver,
		Data:         []byte("hello eoa"),
		TokenAmounts: nil,
		FeeToken:     common.HexToAddress("0x0"),
		ExtraArgs:    extraArgs,
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
		m.MsgSentEvents = append(m.MsgSentEvents, &AnyMsgSentEvent{
			SequenceNumber: it.Event.SequenceNumber,
			RawEvent:       it.Event,
		})
		m.ExpectedSeqNumRange[sourceDest] = ccipocr3common.SeqNumRange{
			ccipocr3common.SeqNum(m.MsgSentEvents[0].SequenceNumber),
			ccipocr3common.SeqNum(m.MsgSentEvents[len(m.MsgSentEvents)-1].SequenceNumber)}
		m.ExpectedSeqNumExec[sourceDest] = append(
			m.ExpectedSeqNumExec[sourceDest],
			it.Event.SequenceNumber)
		messageID := hex.EncodeToString(it.Event.Message.Header.MessageId[:])
		fmt.Printf("Sent CCIP message id %s seq %d from chain %d to chain %d\n", messageID, it.Event.SequenceNumber, src, dest)

		return nil
	}
}

func (m *CCIP16EVM) GetExpectedNextSequenceNumber(ctx context.Context, from, to uint64) (uint64, error) {
	_ = zerolog.Ctx(ctx)
	sourceDest := SourceDestPair{SourceChainSelector: from, DestChainSelector: to}
	seqRange, ok := m.ExpectedSeqNumRange[sourceDest]
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
	seqRange := ccipocr3common.SeqNumRange{ccipocr3common.SeqNum(seq), ccipocr3common.SeqNum(seq)}

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
	seqRange := ccipocr3common.SeqNumRange{ccipocr3common.SeqNum(seq), ccipocr3common.SeqNum(seq)}

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
	l.Info().Msg("Configuring CL nodes for evm")
	name := fmt.Sprintf("node-evm-%s", uuid.New().String()[0:5])

	// Check if this is an external chain (user pre-configured the Out section in TOML)
	// External chains are detected by checking if the URLs are external (not localhost/docker)
	isExternalChain := bc.Out != nil && len(bc.Out.Nodes) > 0 &&
		!strings.Contains(bc.Out.Nodes[0].ExternalHTTPUrl, "host.docker.internal") &&
		!strings.Contains(bc.Out.Nodes[0].ExternalHTTPUrl, "localhost") &&
		!strings.Contains(bc.Out.Nodes[0].ExternalHTTPUrl, "127.0.0.1")

	if isExternalChain {
		// For external chains (testnets/mainnets), don't generate any EVM config.
		// The user must provide the full [[EVM]] config including [[EVM.Nodes]] via node_config_overrides.
		// This avoids duplicate ChainID errors when both auto-generated and user configs exist.
		l.Info().Str("ChainID", bc.ChainID).Msg("External chain detected - skipping auto-generated EVM config (user provides via node_config_overrides)")
		return "", nil
	}

	// For local chains, generate full EVM config
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
WSURL = '%s'
HTTPURL = '%s'`,
		bc.ChainID,
		finality,
		name,
		bc.Out.Nodes[0].InternalWSUrl,
		bc.Out.Nodes[0].InternalHTTPUrl,
	), nil
}

func (m *CCIP16EVM) PreDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	return nil
}

func (m *CCIP16EVM) PostDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	return nil
}

func (m *CCIP16EVM) FundNodes(ctx context.Context, ns []*simple_node_set.Input, nodeKeyBundles map[string]clclient.NodeKeysBundle, bc *blockchain.Input, linkAmount, nativeAmount *big.Int) error {
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
	// Use WS URL if available, otherwise fallback to HTTP URL for HTTP-only mode
	rpcURL := bc.Out.Nodes[0].ExternalWSUrl
	if rpcURL == "" {
		rpcURL = bc.Out.Nodes[0].ExternalHTTPUrl
		l.Info().Str("URL", rpcURL).Msg("Using HTTP URL for ETH client (HTTP-only mode)")
	}
	clientSrc, _, _, err := blockchainutils.ETHClient(ctx, rpcURL, &blockchainutils.GasSettings{
		FeeCapMultiplier: 2,
		TipCapMultiplier: 2,
	})
	if err != nil {
		return fmt.Errorf("could not create basic eth client: %w", err)
	}
	// Use default Anvil key for local chain 1337, otherwise use PRIVATE_KEY env var
	privateKey := blockchainutils.GetNetworkPrivateKey()
	if bc.ChainID == "1337" {
		privateKey = blockchainutils.DefaultAnvilKey
	}

	// nativeAmount is in ETH units (integer) - use directly for FundNodeEIP1559
	// EVM-specific conversion: FundNodeEIP1559 expects ETH as float64
	nativeAmountETH := float64(nativeAmount.Int64())

	for _, addr := range ethKeyAddressesSrc {
		if err := blockchainutils.FundNodeEIP1559(ctx, clientSrc, privateKey, addr, nativeAmountETH); err != nil {
			return fmt.Errorf("failed to fund CL nodes on src chain: %w", err)
		}
	}
	return nil
}
