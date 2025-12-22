package ccip_solana

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"

	ata "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/clclient"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/blockchain"
	"github.com/smartcontractkit/chainlink-testing-framework/framework/components/simple_node_set"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
	chainsel "github.com/smartcontractkit/chain-selectors"
	solconfig "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	solanaseqs "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	solccip "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	solCommonUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
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

type CCIP16Solana struct {
	e                      *deployment.Environment
	chainDetailsBySelector map[uint64]chainsel.ChainDetails
	ExpectedSeqNumRange map[SourceDestPair]ccipocr3common.SeqNumRange
	ExpectedSeqNumExec  map[SourceDestPair][]uint64
	MsgSentEvents       []*AnyMsgSentEvent
}

func NewEmptyCCIP16Solana() *CCIP16Solana {
	return &CCIP16Solana{
		chainDetailsBySelector: make(map[uint64]chainsel.ChainDetails),
		ExpectedSeqNumRange:    make(map[SourceDestPair]ccipocr3common.SeqNumRange),
		ExpectedSeqNumExec:     make(map[SourceDestPair][]uint64),
		MsgSentEvents:          make([]*AnyMsgSentEvent, 0),
	}
}

// NewCCIP16Solana creates new smart-contracts wrappers with utility functions for CCIP16Solana implementation.
func NewCCIP16Solana(ctx context.Context, e *deployment.Environment) (*CCIP16Solana, error) {
	_ = zerolog.Ctx(ctx)
	out := NewEmptyCCIP16Solana()
	out.e = e
	return out, nil
}

func (m *CCIP16Solana) SetCLDF(e *deployment.Environment) {
	m.e = e
}

func updatePrices(ds datastore.DataStore, src, dest uint64, srcChain cldf_solana.Chain) error {
	a := &solanaseqs.SolanaAdapter{}
	fqAddr, err := a.GetFQAddress(ds, src)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	linkAddr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: src,
		Type:          datastore.ContractType("LINK"),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to get LINK address: %w", err)
	}
	fqID := solana.PublicKeyFromBytes(fqAddr)
	fee_quoter.SetProgramID(solana.PublicKeyFromBytes(fqAddr))
	fqAllowedPriceUpdaterPDA, _, _ := state.FindFqAllowedPriceUpdaterPDA(srcChain.DeployerKey.PublicKey(), fqID)
	feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(fqID)
	ixn := fee_quoter.NewUpdatePricesInstruction(
		[]fee_quoter.TokenPriceUpdate{
			{
				SourceToken: solana.WrappedSol,
				UsdPerToken: [28]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 117, 223, 115, 252, 225, 110, 70, 208, 64, 0, 0},
			},
			{
				SourceToken: solana.MustPublicKeyFromBase58(linkAddr.Address),
				UsdPerToken: [28]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 117, 223, 115, 252, 225, 110, 70, 208, 64, 0, 0},
			},
		},
		[]fee_quoter.GasPriceUpdate{
			{
				DestChainSelector: dest,
				UsdPerUnitGas:     solCommonUtil.To28BytesBE(1),
			},
		},
		srcChain.DeployerKey.PublicKey(),
		fqAllowedPriceUpdaterPDA,
		feeQuoterConfigPDA,
	)
	billingTokenConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(solana.WrappedSol, fqID)
	ixn.Append(solana.Meta(billingTokenConfigPDA).WRITE())
	billingTokenConfigPDA, _, _ = state.FindFqBillingTokenConfigPDA(solana.MustPublicKeyFromBase58(linkAddr.Address), fqID)
	ixn.Append(solana.Meta(billingTokenConfigPDA).WRITE())
	fqDestPDA, _, _ := state.FindFqDestChainPDA(dest, fqID)
	ixn.Append(solana.Meta(fqDestPDA).WRITE())
	ix, err := ixn.ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to validate and build instruction: %w", err)
	}
	err = srcChain.Confirm([]solana.Instruction{ix})
	if err != nil {
		return fmt.Errorf("failed to confirm get fee tokens transaction: %w", err)
	}
	return nil
}

func (m *CCIP16Solana) SendMessage(ctx context.Context, src, dest uint64, fields any, opts any) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Sending CCIP message")
	a := &solanaseqs.SolanaAdapter{}
	receiver := common.LeftPadBytes(common.HexToAddress("0xdead").Bytes(), 32)
	msg := ccip_router.SVM2AnyMessage{
		Receiver:     receiver,
		Data:         []byte("hello eoa"),
		TokenAmounts: nil,
		FeeToken:     solana.PublicKey{},
		ExtraArgs:    nil,
	}
	rAddr, err := a.GetRouterAddress(m.e.DataStore, src)
	if err != nil {
		return fmt.Errorf("failed to get router address: %w", err)
	}
	fqAddr, err := a.GetFQAddress(m.e.DataStore, src)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	rmnAddr, err := datastore_utils.FindAndFormatRef(m.e.DataStore, datastore.AddressRef{
		ChainSelector: src,
		Type:          datastore.ContractType("RMNRemote"),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to get RMNRemote address: %w", err)
	}
	linkAddr, err := datastore_utils.FindAndFormatRef(m.e.DataStore, datastore.AddressRef{
		ChainSelector: src,
		Type:          datastore.ContractType("LINK"),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to get LINK address: %w", err)
	}
	fqID := solana.PublicKeyFromBytes(fqAddr)
	routerID := solana.PublicKeyFromBytes(rAddr)
	ccip_router.SetProgramID(routerID)
	l.Info().Msg("Got contract instances, preparing to send CCIP message")
	err = updatePrices(m.e.DataStore, src, dest, m.e.BlockChains.SolanaChains()[src])
	if err != nil {
		return fmt.Errorf("failed to update prices: %w", err)
	}
	feeToken := solana.SolMint
	sender := m.e.BlockChains.SolanaChains()[src].DeployerKey
	destinationChainStatePDA, _ := state.FindDestChainStatePDA(dest, routerID)
	noncePDA, _ := state.FindNoncePDA(dest, sender.PublicKey(), routerID)
	linkFqBillingConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(solana.MustPublicKeyFromBase58(linkAddr.Address), fqID)
	feeTokenFqBillingConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(feeToken, fqID)
	billingSignerPDA, _, _ := state.FindFeeBillingSignerPDA(routerID)
	feeTokenReceiverATA, _, _ := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, feeToken, billingSignerPDA)
	fqDestChainPDA, _, _ := state.FindFqDestChainPDA(dest, fqID)
	rmnRemoteCursesPDA, _, _ := state.FindRMNRemoteCursesPDA(solana.MustPublicKeyFromBase58(rmnAddr.Address))
	routerConfig, _, _ := state.FindConfigPDA(routerID)
	fqConfig, _, _ := state.FindConfigPDA(fqID)
	rmnConfig, _, _ := state.FindConfigPDA(solana.MustPublicKeyFromBase58(rmnAddr.Address))

	base := ccip_router.NewCcipSendInstruction(
		dest,
		msg,
		[]byte{}, // starting indices for accounts, calculated later
		routerConfig,
		destinationChainStatePDA,
		noncePDA,
		sender.PublicKey(),
		solana.SystemProgramID,
		solana.TokenProgramID,
		feeToken,
		solana.PublicKey{},
		feeTokenReceiverATA,
		billingSignerPDA,
		fqID,
		fqConfig,
		fqDestChainPDA,
		feeTokenFqBillingConfigPDA,
		linkFqBillingConfigPDA,
		solana.MustPublicKeyFromBase58(rmnAddr.Address),
		rmnRemoteCursesPDA,
		rmnConfig,
	)

	addressTables := map[solana.PublicKey]solana.PublicKeySlice{}

	tokenIndexes := []byte{}

	// set config.FeeQuoterProgram and CcipRouterProgram since they point to wrong addresses
	solconfig.FeeQuoterProgram = fqID
	solconfig.CcipRouterProgram = routerID

	base.SetTokenIndexes(tokenIndexes)

	tempIx, err := base.ValidateAndBuild()
	if err != nil {
		return err
	}
	ixData, err := tempIx.Data()
	if err != nil {
		return fmt.Errorf("failed to extract data payload from router ccip send instruction: %w", err)
	}
	ix := solana.NewInstruction(routerID, tempIx.Accounts(), ixData)

	// for some reason onchain doesn't see extraAccounts

	ixs := []solana.Instruction{ix}
	result, err := solCommonUtil.SendAndConfirmWithLookupTables(ctx, m.e.BlockChains.SolanaChains()[src].Client, ixs, *sender, solconfig.DefaultCommitment, addressTables, solCommonUtil.AddComputeUnitLimit(400_000))
	if err != nil {
		return fmt.Errorf("failed to send and confirm transaction: %w", err)
	}

	// check CCIP event
	ccipMessageSentEvent := solccip.EventCCIPMessageSent{}
	printEvents := true
	err = solCommonUtil.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, printEvents)
	if err != nil {
		return err
	}

	if len(msg.TokenAmounts) != len(ccipMessageSentEvent.Message.TokenAmounts) {
		return errors.New("token amounts mismatch")
	}

	// TODO: fee bumping?

	transactionID := "N/A"
	if tx, err := result.Transaction.GetTransaction(); err != nil {
		m.e.Logger.Warnf("could not obtain transaction details (err = %s)", err.Error())
	} else if len(tx.Signatures) == 0 {
		m.e.Logger.Warnf("transaction has no signatures: %v", tx)
	} else {
		transactionID = tx.Signatures[0].String()
	}

	m.e.Logger.Infof("CCIP message (id %s) sent from chain selector %d to chain selector %d tx %s seqNum %d nonce %d sender %s",
		common.Bytes2Hex(ccipMessageSentEvent.Message.Header.MessageId[:]),
		src,
		dest,
		transactionID,
		ccipMessageSentEvent.SequenceNumber,
		ccipMessageSentEvent.Message.Header.Nonce,
		ccipMessageSentEvent.Message.Sender.String(),
	)

	sourceDest := SourceDestPair{SourceChainSelector: src, DestChainSelector: dest}
	m.MsgSentEvents = append(m.MsgSentEvents, &AnyMsgSentEvent{
		SequenceNumber: ccipMessageSentEvent.SequenceNumber,
		RawEvent:       ccipMessageSentEvent,
	})
	m.ExpectedSeqNumRange[sourceDest] = ccipocr3common.SeqNumRange{
		ccipocr3common.SeqNum(m.MsgSentEvents[0].SequenceNumber),
		ccipocr3common.SeqNum(m.MsgSentEvents[0].SequenceNumber)}
	m.ExpectedSeqNumExec[sourceDest] = append(
		m.ExpectedSeqNumExec[sourceDest],
		ccipMessageSentEvent.SequenceNumber)

	return nil
}

func (m *CCIP16Solana) GetExpectedNextSequenceNumber(ctx context.Context, from, to uint64) (uint64, error) {
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
func (m *CCIP16Solana) WaitOneSentEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error) {
	l := zerolog.Ctx(ctx)
	a := &solanaseqs.SolanaAdapter{}
	offAddr, err := a.GetOffRampAddress(m.e.DataStore, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramp address: %w", err)
	}
	sourceDest := SourceDestPair{SourceChainSelector: from, DestChainSelector: to}
	seqRange, ok := m.ExpectedSeqNumRange[sourceDest]
	if !ok {
		return nil, fmt.Errorf("no expected sequence number range for source-dest pair %v", sourceDest)
	}
	seenMessages := NewCommitReportTracker(from, seqRange)
	done := make(chan any)
	defer close(done)
	sink, errCh := solEventEmitter[solCommonUtil.EventCommitReportAccepted](ctx, m.e.BlockChains.SolanaChains()[to].Client, solana.PublicKeyFromBytes(offAddr), consts.EventNameCommitReportAccepted, 0, done, time.NewTicker(2*time.Second))

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case eventWithTxn := <-sink:
			commitEvent := eventWithTxn.Event
			// if merkle root is zero, it only contains price updates
			if commitEvent.Report == nil {
				l.Info().Msg("Skipping commit report with nil report")
				continue
			}
			if commitEvent.Report.SourceChainSelector != from {
				l.Info().Uint64("SourceChainSelector", commitEvent.Report.SourceChainSelector).Msg("Skipping commit report from different source chain selector")
				return nil, fmt.Errorf("unexpected source chain selector in commit report: got %d, want %d", commitEvent.Report.SourceChainSelector, from)
			}

			// TODO: this logic is duplicated with verifyCommitReport, share
			mr := commitEvent.Report
			seenMessages.visitCommitReport(mr.SourceChainSelector, mr.MinSeqNr, mr.MaxSeqNr)
			if mr.SourceChainSelector == from &&
				uint64(seqRange.Start()) >= mr.MinSeqNr &&
				uint64(seqRange.End()) <= mr.MaxSeqNr {
				l.Info().Msgf("All sequence numbers committed in a single report [%d, %d]", seqRange.Start(), seqRange.End())
				return true, nil
			}

			if seenMessages.allCommited(from) {
				l.Info().Msgf("All sequence numbers already committed from range [%d, %d]", seqRange.Start(), seqRange.End())
				return true, nil
			}
		case err := <-errCh:
			if err != nil {
				return false, fmt.Errorf("error while fetching commit report events: %w", err)
			}
		case <-timer.C:
			return false, fmt.Errorf("timed out after waiting for commit report on chain selector %d from source selector %d expected seq nr range %s",
				to, from, seqRange.String())
		}
	}
}

type EventWithTxn[T any] struct {
	Event T
	Txn   *solRpc.GetTransactionResult
}

// Scan for events referencing address
func solEventEmitter[T any](ctx context.Context, client *solRpc.Client, address solana.PublicKey, eventType string, startSlot uint64, done chan any, ticker *time.Ticker) (<-chan EventWithTxn[T], <-chan error) {
	ch := make(chan EventWithTxn[T])
	errorCh := make(chan error)
	go func() {
		defer ticker.Stop()
		var until solana.Signature
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				// Scan for transactions referencing the address
				txSigs, err := client.GetSignaturesForAddressWithOpts(
					ctx,
					address,
					&solRpc.GetSignaturesForAddressOpts{
						Commitment: solRpc.CommitmentConfirmed,
						Until:      until,
					},
				)
				if err != nil {
					errorCh <- err
					return
				}

				if len(txSigs) == 0 {
					continue
				}

				// values are returned ordered newest to oldest, so we replay them backwards
				for _, txSig := range slices.Backward(txSigs) {
					if txSig.Err != nil {
						// We're not interested in failed transactions.
						continue
					}
					if txSig.Slot < startSlot {
						// Skip any signatures that are before the starting slot
						continue
					}
					v := uint64(0) // v0 = latest, supports address table lookups
					tx, err := client.GetTransaction(
						ctx,
						txSig.Signature,
						&solRpc.GetTransactionOpts{
							Commitment:                     solRpc.CommitmentConfirmed,
							Encoding:                       solana.EncodingBase64,
							MaxSupportedTransactionVersion: &v,
						},
					)
					if err != nil {
						errorCh <- err
						return
					}

					events, err := solCommonUtil.ParseMultipleEvents[T](tx.Meta.LogMessages, eventType, solconfig.PrintEvents)
					if err != nil && strings.Contains(err.Error(), "event not found") {
						continue
					}
					if err != nil {
						errorCh <- err
						return
					}

					for _, event := range events {
						select {
						case ch <- EventWithTxn[T]{
							Event: event,
							Txn:   tx,
						}:
						case <-done:
							return
						}
					}
				}
				// next scan should stop at the newest signature we've received
				until = txSigs[0].Signature
			}
		}
	}()

	return ch, errorCh
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
func (m *CCIP16Solana) WaitOneExecEventBySeqNo(ctx context.Context, from, to, seq uint64, timeout time.Duration) (any, error) {
	l := zerolog.Ctx(ctx)
	a := &solanaseqs.SolanaAdapter{}
	offAddr, err := a.GetOffRampAddress(m.e.DataStore, to)
	if err != nil {
		return nil, fmt.Errorf("failed to get off ramp address: %w", err)
	}
	offRampAddress := solana.PublicKeyFromBytes(offAddr)
	sourceDest := SourceDestPair{SourceChainSelector: from, DestChainSelector: to}
	seqRange, ok := m.ExpectedSeqNumRange[sourceDest]
	if !ok {
		return nil, fmt.Errorf("no expected sequence number range for source-dest pair %v", sourceDest)
	}

	executionStates := make(map[uint64]int)
	seqNrsToWatch := make(map[uint64]struct{})
	for seqNr := seqRange.Start(); seqNr <= seqRange.End(); seqNr++ {
		seqNrsToWatch[uint64(seqNr)] = struct{}{}
	}

	done := make(chan any)
	sink, errCh := solEventEmitter[solccip.EventExecutionStateChanged](ctx, m.e.BlockChains.SolanaChains()[to].Client, offRampAddress, consts.EventNameExecutionStateChanged, 0, done, time.NewTicker(2*time.Second))

	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case eventWithTxn := <-sink:
			execEvent := eventWithTxn.Event
			// TODO: share with EVM
			_, found := seqNrsToWatch[execEvent.SequenceNumber]
			if found && execEvent.SourceChainSelector == from {
				l.Log().Msgf("Received ExecutionStateChanged (state %s) on chain %d (offramp %s) from chain %d with expected sequence number %d",
					execEvent.State.String(), to, offRampAddress.String(), from, execEvent.SequenceNumber)
				if execEvent.State == ccip_offramp.InProgress_MessageExecutionState {
					// skip the in progress state, executed event should follow
					continue
				}
				executionStates[execEvent.SequenceNumber] = int(execEvent.State)
				delete(seqNrsToWatch, execEvent.SequenceNumber)
				if len(seqNrsToWatch) == 0 {
					return executionStates, nil
				}
			}
		case err := <-errCh:
			if err != nil {
				return nil, fmt.Errorf("error while fetching execution state changed events: %w", err)
			}
		case <-timer.C:
			return nil, fmt.Errorf("timed out waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence numbers %+v",
				to, offRampAddress.String(), from, seqRange)
		}
	}
}

func (m *CCIP16Solana) GetEOAReceiverAddress(ctx context.Context, chainSelector uint64) ([]byte, error) {
	_ = zerolog.Ctx(ctx)
	return nil, nil
}

func (m *CCIP16Solana) GetTokenBalance(ctx context.Context, chainSelector uint64, address, tokenAddress []byte) (*big.Int, error) {
	_ = zerolog.Ctx(ctx)
	return big.NewInt(0), nil
}

func (m *CCIP16Solana) ExposeMetrics(
	ctx context.Context,
	source, dest uint64,
	chainIDs []string,
	wsURLs []string,
) ([]string, *prometheus.Registry, error) {
	return nil, nil, nil
}

func (m *CCIP16Solana) DeployLocalNetwork(ctx context.Context, bc *blockchain.Input) (*blockchain.Output, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Deploying Solana networks")
	d, err := chainsel.GetChainDetailsByChainIDAndFamily(bc.ChainID, chainsel.FamilySolana)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain details: %w", err)
	}
	err = PreloadSolanaEnvironment(bc.ContractsDir, d.ChainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to preload Solana environment: %w", err)
	}
	bc.SolanaPrograms = solanaProgramIDs
	out, err := blockchain.NewBlockchainNetwork(bc)
	if err != nil {
		return nil, fmt.Errorf("failed to create blockchain network: %w", err)
	}
	return out, nil
}

var solanaProgramIDs = map[string]string{
	"ccip_router": "Ccip842gzYHhvdDkSyi2YVCoAWPbYJoApMFzSxQroE9C",
	// "test_token_pool":           "JuCcZ4smxAYv9QHJ36jshA7pA3FuQ3vQeWLUeAtZduJ",
	"burnmint_token_pool":    "41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB",
	"lockrelease_token_pool": "8eqh8wppT9c5rw4ERqNCffvU6cNFJWff9WmkcYtmGiqC",
	"fee_quoter":             "FeeQPGkKDeRV1MgoYfMH6L8o3KeuYjwUZrgn4LRKfjHi",
	"test_ccip_receiver":     "EvhgrPhTDt4LcSPS2kfJgH6T6XWZ6wT3X9ncDGLT1vui",
	"ccip_offramp":           "offqSMQWgQud6WJz694LRzkeN5kMYpCHTpXQr3Rkcjm",
	"mcm":                    "5vNJx78mz7KVMjhuipyr9jKBKcMrKYGdjGkgE4LUmjKk",
	"timelock":               "DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA",
	"access_controller":      "6KsN58MTnRQ8FfPaXHiFPPFGDRioikj9CdPvPxZJdCjb",
	// "external_program_cpi_stub": "2zZwzyptLqwFJFEFxjPvrdhiGpH9pJ3MfrrmZX6NTKxm",
	"rmn_remote": "RmnXLft1mSEwDgMKu2okYuHkiazxntFFcZFrrcXxYg7",
	// "cctp_token_pool":           "CCiTPESGEevd7TBU8EGBKrcxuRq7jx3YtW6tPidnscaZ",
}

var solanaContracts = map[string]datastore.ContractType{
	"ccip_router":        datastore.ContractType("Router"),
	"fee_quoter":         datastore.ContractType("FeeQuoter"),
	"ccip_offramp":       datastore.ContractType("OffRamp"),
	"rmn_remote":         datastore.ContractType("RMNRemote"),
	"mcm":                datastore.ContractType(utils.McmProgramType),
	"timelock":           datastore.ContractType(utils.TimelockProgramType),
	"access_controller":  datastore.ContractType(utils.AccessControllerProgramType),
	"test_ccip_receiver": datastore.ContractType("TestReceiver"),
}

func PreloadSolanaEnvironment(programsPath string, chainSelector uint64) error {
	err := utils.DownloadSolanaCCIPProgramArtifacts(context.Background(), programsPath, utils.VersionToShortCommitSHA[utils.VersionSolanaV0_1_1])
	if err != nil {
		return err
	}
	return nil
}

// Populates datastore with the predeployed program addresses
// pass map [programName]:ContractType of contracts to populate datastore with
func populateDatastore(ds *datastore.MemoryAddressRefStore, contracts map[string]datastore.ContractType, version *semver.Version, qualifier string, chainSel uint64) error {
	for programName, programID := range solanaProgramIDs {
		ct, ok := contracts[programName]
		if !ok {
			continue
		}

		err := ds.Add(datastore.AddressRef{
			Address:       programID,
			ChainSelector: chainSel,
			Qualifier:     qualifier,
			Type:          ct,
			Version:       version,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *CCIP16Solana) ConfigureNodes(ctx context.Context, bc *blockchain.Input) (string, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Configuring CL nodes for solana")
	name := fmt.Sprintf("node-solana-%s", uuid.New().String()[0:5])
	return fmt.Sprintf(`
	[[Solana]]
	ChainID = '%s'
	Enabled = true
	TxRetryTimeout = '90s'
	TxConfirmTimeout = '5m'
	TxExpirationRebroadcast = true
	TxRetentionTimeout = '8h'
	FeeBumpPeriod = '0s'
	FeeEstimatorMode = 'blockhistory'
	ComputeUnitPriceMin = 1
	ComputeUnitPriceMax = 500000
	BlockHistorySize = 150
	LogPollerStartingLookback = '24h'

	[Solana.MultiNode]
	VerifyChainID = false
	Enabled = true
	SyncThreshold = 170
	PollFailureThreshold = 5
	PollInterval = "15s"
	NewHeadsPollInterval = "5s"
	SelectionMode = "PriorityLevel"
	LeaseDuration = "1m"
	NodeIsSyncingEnabled = false
	FinalizedBlockPollInterval = "5s"
	EnforceRepeatableRead = true
	DeathDeclarationDelay = "20s"
	NodeNoNewHeadsThreshold = "20s"
	NoNewFinalizedHeadsThreshold = "20s"
	FinalityTagEnabled = true
	FinalityDepth = 0
	FinalizedBlockOffset = 50

	[[Solana.Nodes]]
	Name = '%s'
	URL = '%s'`,
		bc.ChainID,
		name,
		bc.Out.Nodes[0].InternalHTTPUrl,
	), nil
}

func (m *CCIP16Solana) PreDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	ds := datastore.NewMemoryDataStore()
	ds.Merge(env.DataStore)
	err := populateDatastore(ds.AddressRefStore, solanaContracts, semver.MustParse("1.6.0"), "", selector)
	if err != nil {
		return err
	}
	env.DataStore = ds.Seal()
	return nil
}

func (m *CCIP16Solana) PostDeployContractsForSelector(ctx context.Context, env *deployment.Environment, cls []*simple_node_set.Input, selector uint64, ccipHomeSelector uint64, crAddr string) error {
	// post contract setup for testing
	a := &solanaseqs.SolanaAdapter{}
	linkAddr, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType("LINK"),
	}, selector, datastore_utils.FullRef)
	if err != nil {
		return fmt.Errorf("failed to get LINK address: %w", err)
	}
	fqAddr, err := a.GetFQAddress(env.DataStore, selector)
	if err != nil {
		return fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	routerAddr, err := a.GetRouterAddress(env.DataStore, selector)
	if err != nil {
		return fmt.Errorf("failed to get router address: %w", err)
	}
	chain := env.BlockChains.SolanaChains()[selector]
	fqID := solana.PublicKeyFromBytes(fqAddr)
	routerID := solana.PublicKeyFromBytes(routerAddr)
	fee_quoter.SetProgramID(solana.PublicKeyFromBytes(fqAddr))
	fqAllowedPriceUpdaterPDA, _, _ := state.FindFqAllowedPriceUpdaterPDA(chain.DeployerKey.PublicKey(), fqID)
	feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(fqID)
	ixn, err := fee_quoter.NewAddPriceUpdaterInstruction(
		chain.DeployerKey.PublicKey(),
		fqAllowedPriceUpdaterPDA,
		feeQuoterConfigPDA,
		chain.DeployerKey.PublicKey(),
		solana.SystemProgramID,
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("failed to get fee tokens from fee quoter: %w", err)
	}
	err = chain.Confirm([]solana.Instruction{ixn})
	if err != nil {
		return fmt.Errorf("failed to add price updater: %w", err)
	}
	for _, tokenPubKey := range []solana.PublicKey{
		solana.WrappedSol,
		solana.MustPublicKeyFromBase58(linkAddr.Address),
	} {
		billingTokenConfig := fee_quoter.BillingTokenConfig{
			Enabled: true,
			Mint:    tokenPubKey,
			UsdPerToken: fee_quoter.TimestampedPackedU224{
				Timestamp: time.Now().Unix(),
				Value:     [28]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 46, 117, 223, 115, 252, 225, 110, 70, 208, 64, 0, 0},
			},
			PremiumMultiplierWeiPerEth: 1e18,
		}
		tokenBillingPDA, _, _ := state.FindFqBillingTokenConfigPDA(tokenPubKey, fqID)
		// we dont need to handle test router here because we explicitly create this and token Receiver for test router
		billingSignerPDA, _, _ := state.FindFeeBillingSignerPDA(routerID)
		tokenProgramID := solana.TokenProgramID
		tokenReceiver, _, _ := tokens.FindAssociatedTokenAddress(tokenProgramID, tokenPubKey, billingSignerPDA)
		feeQuoterConfigPDA, _, _ := state.FindFqConfigPDA(fqID)
		ixn, err := fee_quoter.NewAddBillingTokenConfigInstruction(
			billingTokenConfig,
			feeQuoterConfigPDA,
			tokenBillingPDA,
			tokenProgramID,
			tokenPubKey,
			tokenReceiver,
			chain.DeployerKey.PublicKey(),
			billingSignerPDA,
			ata.ProgramID,
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err != nil {
			return fmt.Errorf("failed to build add billing token config instruction: %w", err)
		}
		err = chain.Confirm([]solana.Instruction{ixn})
		if err != nil {
			return fmt.Errorf("failed to add price updater: %w", err)
		}
	}
	return nil
}

func (m *CCIP16Solana) FundNodes(ctx context.Context, ns []*simple_node_set.Input, nodeKeyBundles map[string]clclient.NodeKeysBundle, bc *blockchain.Input, linkAmount, nativeAmount *big.Int) error {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Funding CL nodes with ETH and LINK")
	var keys []solana.PublicKey
	for _, nk := range nodeKeyBundles {
		keys = append(keys, solana.MustPublicKeyFromBase58(nk.TXKey.Data.Attributes.PublicKey))
	}
	client := solRpc.New(bc.Out.Nodes[0].ExternalHTTPUrl)
	err := utils.FundSolanaAccounts(
		ctx,
		keys,
		10,
		client,
	)
	if err != nil {
		return fmt.Errorf("funding solana accounts: %w", err)
	}
	return nil
}
