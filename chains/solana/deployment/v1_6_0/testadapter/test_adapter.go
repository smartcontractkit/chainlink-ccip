package testadapter

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	solrpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	solconfig "github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_0/ccip_router"
	solccip "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	solcommon "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	msg_hasher163 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/message_hasher"
	ccipcommon "github.com/smartcontractkit/chainlink-ccip/deployment/common"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// TODO: deduplicate
type CommitReportTracker struct {
	seenMessages map[uint64]map[uint64]bool
}

func NewCommitReportTracker(sourceChainSelector uint64, seqNrs ccipocr3.SeqNumRange) CommitReportTracker {
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

func init() {
	testadapters.GetTestAdapterRegistry().RegisterTestAdapter(chain_selectors.FamilySolana, semver.MustParse("1.6.0"), NewSVMAdapter)
}

type SVMAdapter[S testadapters.StateProvider] struct {
	state S
	cldf_solana.Chain
}

// state svm_stateview.CCIPChainState
// state, err := stateview.LoadOnchainStateSolana(env)
// if err != nil {
// 	panic(fmt.Sprintf("failed to load onchain state: %T", err))
// }
// NOTE: since this returns a copy, adapters shouldn't be constructed until everything is deployed
// s := state.SolChains[c.ChainSelector()]

func NewSVMAdapter(env *deployment.Environment, selector uint64) testadapters.TestAdapter {
	c, ok := env.BlockChains.SolanaChains()[selector]
	if !ok {
		panic(fmt.Sprintf("chain not found: %d", selector))
	}
	s := &testadapters.DataStoreStateProvider{Selector: selector, DS: env.DataStore}
	return &SVMAdapter[*testadapters.DataStoreStateProvider]{
		state: s,
		Chain: c,
	}
}

func (a *SVMAdapter[S]) BuildMessage(components testadapters.MessageComponents) (any, error) {
	feeToken := solana.PublicKey{}
	if len(components.FeeToken) > 0 {
		var err error
		feeToken, err = solana.PublicKeyFromBase58(components.FeeToken)
		if err != nil {
			return nil, err
		}
	}

	return ccip_router.SVM2AnyMessage{
		Receiver:     common.LeftPadBytes(components.Receiver, 32),
		TokenAmounts: nil,
		Data:         components.Data,
		FeeToken:     feeToken,
		ExtraArgs:    components.ExtraArgs,
	}, nil
}

func (a *SVMAdapter[S]) getAddress(ty datastore.ContractType) (solana.PublicKey, error) {
	addr, err := a.state.GetAddress(ty)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to get %v address: %w", ty, err)
	}
	return solana.PublicKeyFromBase58(addr)
}

// TODO: contractType constants should be extracted from core

func (a *SVMAdapter[S]) SendMessage(ctx context.Context, destChainSelector uint64, m any) (uint64, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Sending CCIP message")

	msg, ok := m.(ccip_router.SVM2AnyMessage)
	if !ok {
		return 0, errors.New("expected ccip_router.SVM2AnyMessage")
	}

	routerID, err := a.getAddress(datastore.ContractType("Router"))
	if err != nil {
		return 0, fmt.Errorf("failed to get router address: %w", err)
	}
	fqID, err := a.getAddress(datastore.ContractType("FeeQuoter"))
	if err != nil {
		return 0, fmt.Errorf("failed to get fee quoter address: %w", err)
	}
	rmnAddr, err := a.getAddress(datastore.ContractType("RMNRemote"))
	if err != nil {
		return 0, fmt.Errorf("failed to get RMNRemote address: %w", err)
	}
	linkAddr, err := a.getAddress(datastore.ContractType("LINK"))
	if err != nil {
		return 0, fmt.Errorf("failed to get LINK address: %w", err)
	}
	ccip_router.SetProgramID(routerID)
	l.Info().Msg("Got contract instances, preparing to send CCIP message")
	// err = updatePrices(m.e.DataStore, src, dest, m.e.BlockChains.SolanaChains()[src])
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to update prices: %w", err)
	// }

	feeToken := solana.SolMint
	sender := a.DeployerKey
	destinationChainStatePDA, _ := state.FindDestChainStatePDA(destChainSelector, routerID)
	noncePDA, _ := state.FindNoncePDA(destChainSelector, sender.PublicKey(), routerID)
	linkFqBillingConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(linkAddr, fqID)
	feeTokenFqBillingConfigPDA, _, _ := state.FindFqBillingTokenConfigPDA(feeToken, fqID)
	billingSignerPDA, _, _ := state.FindFeeBillingSignerPDA(routerID)
	feeTokenReceiverATA, _, _ := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, feeToken, billingSignerPDA)
	fqDestChainPDA, _, _ := state.FindFqDestChainPDA(destChainSelector, fqID)
	rmnRemoteCursesPDA, _, _ := state.FindRMNRemoteCursesPDA(rmnAddr)
	routerConfig, _, _ := state.FindConfigPDA(routerID)
	fqConfig, _, _ := state.FindConfigPDA(fqID)
	rmnConfig, _, _ := state.FindConfigPDA(rmnAddr)

	base := ccip_router.NewCcipSendInstruction(
		destChainSelector,
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
		rmnAddr,
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
		return 0, err
	}
	ixData, err := tempIx.Data()
	if err != nil {
		return 0, fmt.Errorf("failed to extract data payload from router ccip send instruction: %w", err)
	}
	ix := solana.NewInstruction(routerID, tempIx.Accounts(), ixData)

	// for some reason onchain doesn't see extraAccounts

	ixs := []solana.Instruction{ix}
	result, err := solcommon.SendAndConfirmWithLookupTables(ctx, a.Client, ixs, *sender, solconfig.DefaultCommitment, addressTables, solcommon.AddComputeUnitLimit(400_000))
	if err != nil {
		return 0, fmt.Errorf("failed to send and confirm transaction: %w", err)
	}

	// check CCIP event
	ccipMessageSentEvent := solccip.EventCCIPMessageSent{}
	printEvents := true
	err = solcommon.ParseEvent(result.Meta.LogMessages, "CCIPMessageSent", &ccipMessageSentEvent, printEvents)
	if err != nil {
		return 0, err
	}

	if len(msg.TokenAmounts) != len(ccipMessageSentEvent.Message.TokenAmounts) {
		return 0, errors.New("token amounts mismatch")
	}

	// TODO: fee bumping?

	transactionID := "N/A"
	if tx, err := result.Transaction.GetTransaction(); err != nil {
		l.Warn().Msgf("could not obtain transaction details (err = %s)", err.Error())
	} else if len(tx.Signatures) == 0 {
		l.Warn().Msgf("transaction has no signatures: %v", tx)
	} else {
		transactionID = tx.Signatures[0].String()
	}

	l.Info().Msgf("CCIP message (id %s) sent from chain selector %d to chain selector %d tx %s seqNum %d nonce %d sender %s",
		common.Bytes2Hex(ccipMessageSentEvent.Message.Header.MessageId[:]),
		a.Selector,
		destChainSelector,
		transactionID,
		ccipMessageSentEvent.SequenceNumber,
		ccipMessageSentEvent.Message.Header.Nonce,
		ccipMessageSentEvent.Message.Sender.String(),
	)
	return ccipMessageSentEvent.SequenceNumber, nil
}

func (a *SVMAdapter[S]) CCIPReceiver() []byte {
	receiver, err := a.getAddress("TestReceiver")
	if err != nil {
		panic(fmt.Sprintf("failed to get TestReceiver address: %v", err))
	}
	return receiver.Bytes()

}

func (a *SVMAdapter[S]) NativeFeeToken() string {
	return solana.SolMint.String()
}

func (a *SVMAdapter[S]) GetExtraArgs(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error) {
	receiverProgram := solana.PublicKeyFromBytes(receiver)
	receiverTargetAccountPDA, _, _ := solana.FindProgramAddress([][]byte{[]byte("counter")}, receiverProgram)
	receiverExternalExecutionConfigPDA, _, _ := solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, receiverProgram)
	accounts := [][32]byte{
		receiverExternalExecutionConfigPDA,
		receiverTargetAccountPDA,
		solana.SystemProgramID,
	}

	switch sourceFamily {
	case chain_selectors.FamilyEVM:
		return ccipcommon.SerializeClientSVMExtraArgsV1(msg_hasher163.ClientSVMExtraArgsV1{
			AccountIsWritableBitmap:  solccip.GenerateBitMapForIndexes([]int{0, 1}),
			Accounts:                 accounts,
			ComputeUnits:             80_000,
			AllowOutOfOrderExecution: true,
		})
	case chain_selectors.FamilySolana:
		panic("unimplemented GetExtraArgs(solana->solana)")
	default:
		// TODO: add support for other families
		return nil, fmt.Errorf("unsupported source family: %s", sourceFamily)
	}
}

func (a *SVMAdapter[S]) GetInboundNonce(ctx context.Context, sender []byte, srcSel uint64) (uint64, error) {
	chainSelectorLE := solcommon.Uint64ToLE(a.Selector)
	routerAddress, err := a.getAddress(datastore.ContractType("Router"))
	if err != nil {
		return 0, err
	}
	noncePDA, _, err := solana.FindProgramAddress([][]byte{[]byte("nonce"), chainSelectorLE, sender}, routerAddress)
	if err != nil {
		return 0, err
	}
	var nonceCounterAccount ccip_router.Nonce
	// we ignore the error because the account might not exist yet
	_ = solcommon.GetAccountDataBorshInto(ctx, a.Client, noncePDA, solconfig.DefaultCommitment, &nonceCounterAccount)
	return nonceCounterAccount.Counter, nil
}

func (a *SVMAdapter[S]) ValidateCommit(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNumRange ccipocr3.SeqNumRange) {
	var startSlot uint64
	if startBlock != nil {
		startSlot = *startBlock
	}
	offRampAddress, err := a.getAddress(datastore.ContractType("OffRamp"))
	_, err = confirmCommitWithExpectedSeqNumRangeSol(
		t,
		sourceSelector,
		a.Chain,
		offRampAddress,
		startSlot,
		seqNumRange,
		true,
	)
	require.NoError(t, err)
}

func (a *SVMAdapter[S]) ValidateExec(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNrs []uint64) (executionStates map[uint64]int) {
	var startSlot uint64
	if startBlock != nil {
		startSlot = *startBlock
	}
	offRampAddress, err := a.getAddress("OffRamp")
	require.NoError(t, err)
	executionStates, err = confirmExecWithSeqNrsSol(
		t,
		sourceSelector,
		a.Chain,
		offRampAddress,
		startSlot,
		seqNrs,
	)
	require.NoError(t, err)
	return executionStates
}

type EventWithTxn[T any] struct {
	Event T
	Txn   *solrpc.GetTransactionResult
}

// SolEventEmitter listens for events of type T emitted by the Solana program at the given address. Failed transactions
// can be included by setting the includeFailed flag to true.
func SolEventEmitter[T any](ctx context.Context, client *solrpc.Client, address solana.PublicKey, eventType string, startSlot uint64, done chan any, ticker *time.Ticker, includeFailed bool) (<-chan EventWithTxn[T], <-chan error) {
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
					&solrpc.GetSignaturesForAddressOpts{
						Commitment: solrpc.CommitmentConfirmed,
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
					if txSig.Err != nil && !includeFailed {
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
						&solrpc.GetTransactionOpts{
							Commitment:                     solrpc.CommitmentConfirmed,
							Encoding:                       solana.EncodingBase64,
							MaxSupportedTransactionVersion: &v,
						},
					)
					if err != nil {
						errorCh <- err
						return
					}

					events, err := solcommon.ParseMultipleEvents[T](tx.Meta.LogMessages, eventType, solconfig.PrintEvents)
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

func confirmCommitWithExpectedSeqNumRangeSol(
	t *testing.T,
	srcSelector uint64,
	dest cldf_solana.Chain,
	offrampAddress solana.PublicKey,
	startSlot uint64,
	expectedSeqNumRange ccipocr3.SeqNumRange,
	enforceSingleCommit bool,
) (bool, error) {
	seenMessages := NewCommitReportTracker(srcSelector, expectedSeqNumRange)

	done := make(chan any)
	defer close(done)
	sink, errCh := SolEventEmitter[solcommon.EventCommitReportAccepted](t.Context(), dest.Client, offrampAddress, consts.EventNameCommitReportAccepted, startSlot, done, time.NewTicker(2*time.Second), false)

	timeout := time.NewTimer(tests.WaitTimeout(t))
	defer timeout.Stop()

	for {
		select {
		case eventWithTxn := <-sink:
			commitEvent := eventWithTxn.Event
			// if merkle root is zero, it only contains price updates
			if commitEvent.Report == nil {
				t.Logf("Skipping CommitReportAccepted with only price updates")
				continue
			}
			// Check source chain selector match. We might see roots from other chains on repeated test runs
			if commitEvent.Report.SourceChainSelector != srcSelector {
				t.Logf("[Root] Source chain mismatch: got %d, expected %d",
					commitEvent.Report.SourceChainSelector, srcSelector)
				continue
			}
			require.Equal(t, srcSelector, commitEvent.Report.SourceChainSelector)

			// TODO: this logic is duplicated with verifyCommitReport, share
			mr := commitEvent.Report
			seenMessages.visitCommitReport(mr.SourceChainSelector, mr.MinSeqNr, mr.MaxSeqNr)
			if mr.SourceChainSelector == srcSelector &&
				uint64(expectedSeqNumRange.Start()) >= mr.MinSeqNr &&
				uint64(expectedSeqNumRange.End()) <= mr.MaxSeqNr {
				t.Logf("All sequence numbers committed in a single report [%d, %d]", expectedSeqNumRange.Start(), expectedSeqNumRange.End())
				return true, nil
			}

			if !enforceSingleCommit && seenMessages.allCommited(srcSelector) {
				t.Logf("All sequence numbers already committed from range [%d, %d]", expectedSeqNumRange.Start(), expectedSeqNumRange.End())
				return true, nil
			}
		case err := <-errCh:
			require.NoError(t, err)
		case <-timeout.C:
			return false, fmt.Errorf("timed out after waiting for commit report on chain selector %d from source selector %d expected seq nr range %s",
				dest.Selector, srcSelector, expectedSeqNumRange.String())
		}
	}
}

type MessageStateEvent struct {
	SequenceNumber uint64
	Block          uint64
	State          ccip_offramp.MessageExecutionState
}

// GetMessageStatesWithSeqNrsSol waits for execution state changes on the destination Solana chain with the expected
// sequence numbers. The timeout is configurable.
// If "inProgress" is true, and there are unexecutable messages, it will continue to watch for additional "InProgress"
// states for the entire timeoutDuration.
func GetMessageStatesWithSeqNrsSol(
	t *testing.T,
	timeoutDuration time.Duration,
	srcSelector uint64,
	dest cldf_solana.Chain,
	offrampAddress solana.PublicKey,
	startSlot uint64,
	expectedSeqNrs []uint64,
	inProgress bool,
) (executionStates map[uint64][]MessageStateEvent, err error) {
	// TODO: share with EVM
	// some state to efficiently track the execution states
	// of all the expected sequence numbers.
	executionStates = make(map[uint64][]MessageStateEvent)
	seqNrsInProgress := make(map[uint64]struct{})
	seqNrsToWatch := make(map[uint64]struct{})
	for _, seqNr := range expectedSeqNrs {
		seqNrsToWatch[seqNr] = struct{}{}
		seqNrsInProgress[seqNr] = struct{}{}
	}

	done := make(chan any)
	defer close(done)
	sink, errCh := SolEventEmitter[solccip.EventExecutionStateChanged](t.Context(), dest.Client, offrampAddress, consts.EventNameExecutionStateChanged, startSlot, done, time.NewTicker(2*time.Second), inProgress)

	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()

	for {
		select {
		case eventWithTxn := <-sink:
			execEvent := eventWithTxn.Event
			// TODO: share with EVM
			_, found := seqNrsToWatch[execEvent.SequenceNumber]
			if found && execEvent.SourceChainSelector == srcSelector {
				t.Logf("Received ExecutionStateChanged (state %s) on chain %d (offramp %s) from chain %d with expected sequence number %d",
					execEvent.State.String(), dest.Selector, offrampAddress.String(), srcSelector, execEvent.SequenceNumber)
				executionStates[execEvent.SequenceNumber] = append(executionStates[execEvent.SequenceNumber],
					MessageStateEvent{
						SequenceNumber: execEvent.SequenceNumber,
						Block:          eventWithTxn.Txn.Slot,
						State:          execEvent.State,
					})
				if execEvent.State == ccip_offramp.InProgress_MessageExecutionState {
					delete(seqNrsInProgress, execEvent.SequenceNumber)
					// continue watching for final state or timeout
					continue
				}
				delete(seqNrsToWatch, execEvent.SequenceNumber)
				delete(seqNrsInProgress, execEvent.SequenceNumber)
				if len(seqNrsToWatch) == 0 {
					return executionStates, nil
				}
			}
		case err := <-errCh:
			require.NoError(t, err)
		case <-timeout.C:
			// If we have some status for every seqNr, return what we have instead of an error.
			if len(seqNrsInProgress) == 0 {
				return executionStates, nil
			}
			// Otherwise, return a timeout error.
			return nil, fmt.Errorf("timed out waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence numbers %+v",
				dest.Selector, offrampAddress.String(), srcSelector, expectedSeqNrs)
		}
	}
}

// ConfirmExecWithSeqNrsSol waits for an execution state change on the destination Solana chain with the expected
// sequence numbers. The timeout is automatically set to 30 seconds or 90% of the test timeout (if there is one).
// This is a wrapper around the more general GetMessageStatesWithSeqNrsSol.
func confirmExecWithSeqNrsSol(
	t *testing.T,
	srcSelector uint64,
	dest cldf_solana.Chain,
	offrampAddress solana.PublicKey,
	startSlot uint64,
	expectedSeqNrs []uint64,
) (executionStates map[uint64]int, err error) {
	timeout := tests.WaitTimeout(t)
	states, err := GetMessageStatesWithSeqNrsSol(t, timeout, srcSelector, dest, offrampAddress, startSlot, expectedSeqNrs, false)
	if err != nil {
		return nil, err
	}

	executionStates = make(map[uint64]int)

	for seqNr, stateList := range states {
		if len(stateList) == 0 {
			return nil, fmt.Errorf("no execution states found for seqNr %d", seqNr)
		}
		// check that the final state is either success or failure
		state := stateList[len(stateList)-1].State
		if state != ccip_offramp.Success_MessageExecutionState && state != ccip_offramp.Failure_MessageExecutionState {
			return nil, fmt.Errorf("expected final execution state for seqNr %d, got %s", seqNr, state.String())
		}

		executionStates[seqNr] = int(state)
	}

	return executionStates, nil
}
