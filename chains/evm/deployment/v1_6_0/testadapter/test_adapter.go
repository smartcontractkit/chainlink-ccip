package testadapter

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/nonce_manager"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	msg_hasher163 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/message_hasher"
	ccipcommon "github.com/smartcontractkit/chainlink-ccip/deployment/common"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	commonutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

func getExecutionState(t *testing.T, sourceSelector uint64, offRamp offramp.OffRampInterface, expectedSeqNr uint64) (offramp.OffRampSourceChainConfig, uint8) {
	scc, err := offRamp.GetSourceChainConfig(nil, sourceSelector)
	require.NoError(t, err)
	executionState, err := offRamp.GetExecutionState(nil, sourceSelector, expectedSeqNr)
	require.NoError(t, err)
	return scc, executionState
}

func init() {
	testadapters.GetTestAdapterRegistry().RegisterTestAdapter(chain_selectors.FamilyEVM, semver.MustParse("1.6.0"), NewEVMAdapter)
}

type EVMAdapter struct {
	state testadapters.StateProvider
	cldf_evm.Chain
}

func NewEVMAdapter(env *deployment.Environment, selector uint64) testadapters.TestAdapter {
	// TODO: tron needs to use TronChains
	c, ok := env.BlockChains.EVMChains()[selector]
	if !ok {
		panic(fmt.Sprintf("chain not found: %d", selector))
	}

	s := &testadapters.DataStoreStateProvider{Selector: selector, DS: env.DataStore}
	return &EVMAdapter{
		state: s,
		Chain: c,
	}
}

func (a *EVMAdapter) getAddress(ty datastore.ContractType) (common.Address, error) {
	addr, err := a.state.GetAddress(ty)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get %v address: %w", ty, err)
	}
	return common.HexToAddress(addr), nil
}

func (a *EVMAdapter) BuildMessage(components testadapters.MessageComponents) (any, error) {
	receiver := common.LeftPadBytes(components.Receiver, 32)
	feeToken := common.HexToAddress(a.NativeFeeToken())
	if len(components.FeeToken) > 0 {
		feeToken = common.HexToAddress(components.FeeToken)
	}

	tokenAmounts := []router.ClientEVMTokenAmount{}
	for i, ta := range components.TokenAmounts {
		if !common.IsHexAddress(ta.Token) {
			return nil, fmt.Errorf("invalid token address at index %d: %s", i, ta.Token)
		}

		tokenAmounts = append(tokenAmounts,
			router.ClientEVMTokenAmount{
				Token:  common.HexToAddress(ta.Token),
				Amount: ta.Amount,
			},
		)
	}

	return router.ClientEVM2AnyMessage{
		Receiver:     receiver,
		Data:         components.Data,
		TokenAmounts: tokenAmounts,
		FeeToken:     feeToken,
		ExtraArgs:    components.ExtraArgs,
	}, nil
}

func (a *EVMAdapter) SendMessage(ctx context.Context, destChainSelector uint64, m any) (uint64, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Sending CCIP message")

	msg, ok := m.(router.ClientEVM2AnyMessage)
	if !ok {
		return 0, errors.New("expected router.ClientEVM2AnyMessage")
	}
	// case chainsel.FamilyTon:
	// 	receiverAddr, err := datastore_utils.FindAndFormatRef(m.e.DataStore, datastore.AddressRef{
	// 		ChainSelector: dest,
	// 		Type:          datastore.ContractType("Receiver"),
	// 	}, dest, datastore_utils.FullRef)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to get TonReceiver address: %w", err)
	// 	}
	// 	tonreceiver, err := address.ParseAddr(receiverAddr.Address)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to parse TON receiver address: %w", err)
	// 	}
	// 	ac := codec.NewAddressCodec()
	// 	receiver, err = ac.AddressStringToBytes(tonreceiver.String())
	// 	if err != nil {
	// 		return fmt.Errorf("failed to convert TON address to bytes: %w", err)
	// 	}

	const errCodeInsufficientFee = "0x07da6ee6"
	const cannotDecodeErrorReason = "could not decode error reason"
	const errMsgMissingTrieNode = "missing trie node"
	sender := a.DeployerKey
	defer func() { sender.Value = nil }()
	rAddr, err := a.getAddress(datastore.ContractType("Router"))
	if err != nil {
		return 0, fmt.Errorf("failed to get router address: %w", err)
	}
	r, err := router.NewRouter(
		rAddr,
		a.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to create router instance: %w", err)
	}
	onRampAddr, err := r.GetOnRamp(nil, destChainSelector)
	if err != nil {
		return 0, fmt.Errorf("failed to get onramp address: %w", err)
	}
	onRamp, err := onramp.NewOnRamp(
		onRampAddr,
		a.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to create onramp instance: %w", err)
	}
	l.Info().Msg("Got contract instances, preparing to send CCIP message")
	// TODO: why?
	// err = updatePrices(m.e.DataStore, src, dest, m.e.BlockChains.EVMChains()[src])
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to update prices: %w", err)
	// }

	var retryCount int
	for {
		fee, err := r.GetFee(&bind.CallOpts{Context: ctx}, destChainSelector, msg)
		if err != nil {
			return 0, fmt.Errorf("failed to get EVM fee: %w", deployment.MaybeDataErr(err))
		}

		sender.Value = fee

		tx, err := r.CcipSend(sender, destChainSelector, msg)
		if err != nil {
			return 0, fmt.Errorf("failed to send CCIP message: %w", err)
		}

		blockNum, err := a.Confirm(tx)
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
					return 0, fmt.Errorf("failed to confirm CCIP message after %d retries: %w", retryCount, deployment.MaybeDataErr(err))
				}
				retryCount++
				continue
			}

			return 0, fmt.Errorf("failed to confirm CCIP message: %w", deployment.MaybeDataErr(err))
		}
		it, err := onRamp.FilterCCIPMessageSent(&bind.FilterOpts{
			Start:   blockNum,
			End:     &blockNum,
			Context: context.Background(),
		}, []uint64{destChainSelector}, []uint64{})
		if err != nil {
			return 0, fmt.Errorf("failed to filter CCIPMessageSent events: %w", err)
		}

		if !it.Next() {
			return 0, fmt.Errorf("no CCIP message sent event found")
		}

		messageID := hex.EncodeToString(it.Event.Message.Header.MessageId[:])
		fmt.Printf("Sent CCIP message id %s seq %d from chain %d to chain %d\n", messageID, it.Event.SequenceNumber, a.Selector, destChainSelector)
		return it.Event.SequenceNumber, nil
	}
}

func (a *EVMAdapter) CCIPReceiver() []byte {
	return common.LeftPadBytes(common.HexToAddress("0xdead").Bytes(), 32)
}

func (a *EVMAdapter) NativeFeeToken() string {
	return "0x0"
}

func (a *EVMAdapter) GetExtraArgs(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error) {
	switch sourceFamily {
	case chain_selectors.FamilyEVM:
		return ccipcommon.SerializeClientGenericExtraArgsV2(msg_hasher163.ClientGenericExtraArgsV2{
			GasLimit:                 new(big.Int).SetUint64(100_000),
			AllowOutOfOrderExecution: true,
		})
	case chain_selectors.FamilySolana:
		// EVM allows empty extraArgs
		return nil, nil
	default:
		// TODO: add support for other families
		return nil, fmt.Errorf("unsupported source family: %s", sourceFamily)
	}
}

func (a *EVMAdapter) GetInboundNonce(ctx context.Context, sender []byte, srcSel uint64) (uint64, error) {
	nonceManagerAddress, err := a.getAddress("NonceManager")
	if err != nil {
		return 0, err
	}
	nonceManager, err := nonce_manager.NewNonceManager(nonceManagerAddress, a.Client)
	if err != nil {
		return 0, err
	}
	return nonceManager.GetInboundNonce(&bind.CallOpts{Context: ctx}, srcSel, sender)
}

func (a *EVMAdapter) ValidateCommit(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNumRange ccipocr3.SeqNumRange) {
	offRampAddress, err := a.getAddress(datastore.ContractType("OffRamp"))
	require.NoError(t, err)
	offRamp, err := offramp.NewOffRamp(
		offRampAddress,
		a.Client)
	require.NoError(t, err)
	_, err = ConfirmCommitWithExpectedSeqNumRange(
		t,
		sourceSelector,
		a.Chain,
		offRamp,
		startBlock,
		seqNumRange,
		true,
	)
	require.NoError(t, err)
}

func (a *EVMAdapter) ValidateExec(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNrs []uint64) (executionStates map[uint64]int) {
	offRampAddress, err := a.getAddress("OffRamp")
	require.NoError(t, err)
	offRamp, err := offramp.NewOffRamp(
		offRampAddress,
		a.Client)
	require.NoError(t, err)
	executionStates, err = ConfirmExecWithSeqNrs(
		t,
		sourceSelector,
		a.Chain,
		offRamp,
		startBlock,
		seqNrs,
	)
	require.NoError(t, err)
	return executionStates
}

func (a *EVMAdapter) AllowRouterToWithdrawTokens(ctx context.Context, tokenAddress string, amount *big.Int) error {
	if !common.IsHexAddress(tokenAddress) {
		return fmt.Errorf("invalid token address: %s", tokenAddress)
	}
	if amount.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("amount must be greater than zero: %s", amount.String())
	}

	routerAddr, err := a.getAddress(datastore.ContractType("Router"))
	if err != nil {
		return fmt.Errorf("failed to get router address: %w", err)
	}

	tokenAddr := common.HexToAddress(tokenAddress)
	if tokenAddr == (common.Address{}) {
		return errors.New("cannot approve zero address token")
	}

	// NOTE: GetTokenExpansionConfig uses BurnMintERC20 as the token
	// type so we need to be consistent about using it here as well.
	token, err := burn_mint_erc20.NewBurnMintERC20(tokenAddr, a.Chain.Client)
	if err != nil {
		return fmt.Errorf("failed to create burn mint erc20 instance: %w", err)
	}

	tx, err := token.Approve(a.DeployerKey, routerAddr, amount)
	if err != nil {
		return fmt.Errorf(
			"failed to send approve tokens tx (token = %q, deployer = %q, router = %q): %w",
			tokenAddr.Hex(), a.DeployerKey.From.Hex(), routerAddr.Hex(), err,
		)
	}

	_, err = a.Chain.Confirm(tx)
	if err != nil {
		return fmt.Errorf(
			"failed to confirm approve tokens tx (token = %q, deployer = %q, router = %q): %w",
			tokenAddr.Hex(), a.DeployerKey.From.Hex(), routerAddr.Hex(), err,
		)
	}

	return nil
}

func (a *EVMAdapter) GetTokenBalance(ctx context.Context, tokenAddress string, ownerAddress []byte) (*big.Int, error) {
	if !common.IsHexAddress(tokenAddress) {
		return nil, fmt.Errorf("invalid token address: %s", tokenAddress)
	}

	ownerAddr := common.BytesToAddress(ownerAddress)
	if ownerAddr == (common.Address{}) {
		return nil, errors.New("cannot get balance of zero address owner")
	}

	tokenAddr := common.HexToAddress(tokenAddress)
	if tokenAddr == (common.Address{}) {
		return nil, errors.New("cannot get balance of zero address token")
	}

	// NOTE: GetTokenExpansionConfig uses BurnMintERC20 as the token
	// type so we need to be consistent about using it here as well.
	token, err := burn_mint_erc20.NewBurnMintERC20(tokenAddr, a.Chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create burn mint erc20 instance for address %q: %w", tokenAddr.Hex(), err)
	}

	balance, err := token.BalanceOf(&bind.CallOpts{Context: ctx}, ownerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance of token %q for owner %q: %w", tokenAddr.Hex(), ownerAddr.Hex(), err)
	}

	return balance, nil
}

func (a *EVMAdapter) GetTokenExpansionConfig() tokensapi.TokenExpansionInputPerChain {
	suffix := strconv.FormatUint(a.Selector, 10) + "-" + a.Family()
	admin := a.Chain.DeployerKey.From.Hex()
	deci := uint8(18)

	oneToken := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(deci)), nil)
	mintAmnt := new(big.Int).Mul(oneToken, big.NewInt(1_000_000)) // pre-mint 1 million tokens

	return tokensapi.TokenExpansionInputPerChain{
		TokenPoolVersion:        cciputils.Version_1_5_1,
		TokenPoolRateLimitAdmin: admin,
		TokenPoolAdmin:          admin,
		TARAdmin:                admin,
		DeployTokenInput: &tokensapi.DeployTokenInput{
			Decimals:               deci,
			Symbol:                 "TEST_TOKEN_" + suffix,
			Name:                   "TEST TOKEN " + suffix,
			Type:                   bnmERC20ops.ContractType, // BnM ERC20 is the most common
			Supply:                 big.NewInt(0),            // unlimited supply
			PreMint:                mintAmnt,                 // pre-mint some tokens for transfers
			Senders:                []string{admin},          // use deployer as sender
			ExternalAdmin:          "",                       // not needed for tests
			DisableFreezeAuthority: false,                    // not applicable for EVM
			TokenPrivKey:           "",                       // not applicable for EVM
			CCIPAdmin:              admin,                    // deployer is the admin (if empty defaults to timelock)
		},
		DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
			PoolType:           cciputils.BurnMintTokenPool.String(),
			TokenPoolQualifier: "TEST TOKEN POOL " + suffix,
		},
		// optional fields left empty, but included here for completeness
		RemoteCounterpartUpdates: map[uint64]tokensapi.RateLimiterConfig{},
		RemoteCounterpartDeletes: []uint64{},
	}
}

func (a *EVMAdapter) GetRegistryAddress() (string, error) {
	addr, err := a.getAddress("TokenAdminRegistry")
	if err != nil {
		return "", fmt.Errorf("failed to get TokenAdminRegistry address: %w", err)
	}

	return addr.Hex(), nil
}

// Co// ConfirmCommitWithExpectedSeqNumRange waits for a commit report on the destination chain with the expected sequence number range.
// startBlock is the block number to start watching from.
// If startBlock is nil, it will start watching from the latest block.
func ConfirmCommitWithExpectedSeqNumRange(
	t *testing.T,
	srcSelector uint64,
	dest cldf_evm.Chain,
	offRamp offramp.OffRampInterface,
	startBlock *uint64,
	expectedSeqNumRange ccipocr3.SeqNumRange,
	enforceSingleCommit bool,
) (*offramp.OffRampCommitReportAccepted, error) {
	sink := make(chan *offramp.OffRampCommitReportAccepted)
	subscription, err := offRamp.WatchCommitReportAccepted(&bind.WatchOpts{
		Context: context.Background(),
		Start:   startBlock,
	}, sink)
	if err != nil {
		return nil, fmt.Errorf("error to subscribe CommitReportAccepted : %w", err)
	}

	seenMessages := testadapters.NewCommitReportTracker(srcSelector, expectedSeqNumRange)

	verifyCommitReport := func(report *offramp.OffRampCommitReportAccepted) bool {
		t.Logf("Verifying commit report: blessed roots=%d, unblessed roots=%d",
			len(report.BlessedMerkleRoots), len(report.UnblessedMerkleRoots))

		processRoots := func(roots []offramp.InternalMerkleRoot, rootType string) bool {
			t.Logf("Processing %d %s merkle roots", len(roots), rootType)
			for i, mr := range roots {
				t.Logf(
					"[%s Root #%d] Received commit report for [%d, %d] on selector %d from source selector %d expected seq nr range %s, token prices: %v",
					rootType, i+1, mr.MinSeqNr, mr.MaxSeqNr, dest.Selector, srcSelector, expectedSeqNumRange.String(), report.PriceUpdates.TokenPriceUpdates,
				)
				seenMessages.VisitCommitReport(srcSelector, mr.MinSeqNr, mr.MaxSeqNr)

				// Check source chain selector match
				if mr.SourceChainSelector != srcSelector {
					t.Logf("[%s Root #%d] Source chain mismatch: got %d, expected %d",
						rootType, i+1, mr.SourceChainSelector, srcSelector)
					continue
				}

				// Check sequence number range
				expectedStart := uint64(expectedSeqNumRange.Start())
				expectedEnd := uint64(expectedSeqNumRange.End())
				if expectedStart < mr.MinSeqNr || expectedEnd > mr.MaxSeqNr {
					t.Logf("[%s Root #%d] Sequence range mismatch: expected [%d, %d], got [%d, %d]",
						rootType, i+1, expectedStart, expectedEnd, mr.MinSeqNr, mr.MaxSeqNr)
					continue
				}

				t.Logf(
					"[%s Root #%d] ✅ All sequence numbers committed in a single report [%d, %d]",
					rootType, i+1, expectedSeqNumRange.Start(), expectedSeqNumRange.End(),
				)
				return true
			}

			// Check if all messages committed across multiple reports (if enforceSingleCommit is false)
			if !enforceSingleCommit && seenMessages.AllCommitted(srcSelector) {
				t.Logf(
					"✅ All sequence numbers already committed from range [%d, %d] across multiple reports",
					expectedSeqNumRange.Start(), expectedSeqNumRange.End(),
				)
				return true
			}

			t.Logf("No matching %s roots found for expected criteria", rootType)
			return false
		}

		blessedResult := processRoots(report.BlessedMerkleRoots, "Blessed")
		if blessedResult {
			return true
		}

		unblessedResult := processRoots(report.UnblessedMerkleRoots, "Unblessed")
		return unblessedResult
	}

	defer subscription.Unsubscribe()
	timeoutDuration := tests.WaitTimeout(t)
	startTime := time.Now()
	t.Logf("Starting commit report wait with timeout: %s", timeoutDuration)
	timeout := time.NewTimer(timeoutDuration)
	defer timeout.Stop()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-t.Context().Done():
			return nil, nil
		case <-ticker.C:
			elapsed := time.Since(startTime)
			remaining := timeoutDuration - elapsed
			t.Logf("Waiting for commit report on chain selector %d from source selector %d expected seq nr range %s (elapsed: %s, remaining: %s)",
				dest.Selector, srcSelector, expectedSeqNumRange.String(), elapsed.Round(time.Second), remaining.Round(time.Second))

			// Need to do this because the subscription sometimes fails to get the event.
			t.Logf("Creating FilterCommitReportAccepted iterator for offramp %s", offRamp.Address().String())
			iter, err := offRamp.FilterCommitReportAccepted(&bind.FilterOpts{
				Context: t.Context(),
			})
			// In some test case the test ends while the filter is still running resulting in a context.Canceled error.
			if err != nil {
				if errors.Is(err, context.Canceled) {
					t.Logf("FilterCommitReportAccepted context was canceled, continuing...")
				} else {
					t.Logf("FilterCommitReportAccepted failed with error: %v", err)
					require.NoError(t, err)
				}
				continue // Skip to next ticker iteration if filter creation failed
			}

			eventCount := 0
			t.Logf("Starting to iterate through FilterCommitReportAccepted events...")
			for iter.Next() {
				eventCount++
				event := iter.Event
				t.Logf("Processing commit report event #%d: blessed roots=%d, unblessed roots=%d",
					eventCount, len(event.BlessedMerkleRoots), len(event.UnblessedMerkleRoots))

				verified := verifyCommitReport(event)
				if verified {
					t.Logf("Commit report verified successfully after processing %d events", eventCount)
					return event, nil
				}
				t.Logf("Commit report event #%d did not match expected criteria", eventCount)
			}

			// Check for iteration errors
			if err := iter.Error(); err != nil {
				t.Logf("Iterator error after processing %d events: %v", eventCount, err)
			} else if eventCount == 0 {
				t.Logf("No commit report events found in this iteration")
			} else {
				t.Logf("Processed %d commit report events, none matched expected criteria", eventCount)
			}
		case subErr := <-subscription.Err():
			return nil, fmt.Errorf("subscription error: %w", subErr)
		case <-timeout.C:
			return nil, fmt.Errorf("timed out after waiting for commit report on chain selector %d from source selector %d expected seq nr range %s",
				dest.Selector, srcSelector, expectedSeqNumRange.String())
		case report := <-sink:
			t.Logf("Received commit report via subscription: blessed roots=%d, unblessed roots=%d",
				len(report.BlessedMerkleRoots), len(report.UnblessedMerkleRoots))
			verified := verifyCommitReport(report)
			if verified {
				t.Logf("Subscription commit report verified successfully")
				return report, nil
			} else {
				t.Logf("Subscription commit report did not match expected criteria")
			}
		}
	}
}

// ConfirmExecWithSeqNrs waits for an execution state change on the destination chain with the expected sequence number.
// startBlock is the block number to start watching from.
// If startBlock is nil, it will start watching from the latest block.
// Returns a map that maps the expected sequence number to its execution state.
func ConfirmExecWithSeqNrs(
	t *testing.T,
	sourceSelector uint64,
	dest cldf_evm.Chain,
	offRamp offramp.OffRampInterface,
	startBlock *uint64,
	expectedSeqNrs []uint64,
) (executionStates map[uint64]int, err error) {
	if len(expectedSeqNrs) == 0 {
		return nil, errors.New("no expected sequence numbers provided")
	}

	timeout := time.NewTimer(tests.WaitTimeout(t))
	defer timeout.Stop()
	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()
	sink := make(chan *offramp.OffRampExecutionStateChanged)
	subscription, err := offRamp.WatchExecutionStateChanged(&bind.WatchOpts{
		Context: context.Background(),
		Start:   startBlock,
	}, sink, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error to subscribe ExecutionStateChanged : %w", err)
	}
	defer subscription.Unsubscribe()

	// some state to efficiently track the execution states
	// of all the expected sequence numbers.
	executionStates = make(map[uint64]int)
	seqNrsToWatch := make(map[uint64]struct{})
	for _, seqNr := range expectedSeqNrs {
		seqNrsToWatch[seqNr] = struct{}{}
	}
	for {
		select {
		case <-tick.C:
			for expectedSeqNr := range seqNrsToWatch {
				scc, executionState := getExecutionState(t, sourceSelector, offRamp, expectedSeqNr)
				t.Logf("Waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence number %d, current onchain minSeqNr: %d, execution state: %s",
					dest.Selector, offRamp.Address().String(), sourceSelector, expectedSeqNr, scc.MinSeqNr, commonutils.ExecutionStateToString(executionState))
				if executionState == commonutils.EXECUTION_STATE_SUCCESS || executionState == commonutils.EXECUTION_STATE_FAILURE {
					t.Logf("Observed %s execution state on chain %d (offramp %s) from chain %d with expected sequence number %d",
						commonutils.ExecutionStateToString(executionState), dest.Selector, offRamp.Address().String(), sourceSelector, expectedSeqNr)
					executionStates[expectedSeqNr] = int(executionState)
					delete(seqNrsToWatch, expectedSeqNr)
					if len(seqNrsToWatch) == 0 {
						return executionStates, nil
					}
				}
			}
		case execEvent := <-sink:
			t.Logf("Received ExecutionStateChanged (state %s) for seqNum %d on chain %d (offramp %s) from chain %d",
				commonutils.ExecutionStateToString(execEvent.State), execEvent.SequenceNumber, dest.Selector, offRamp.Address().String(),
				sourceSelector,
			)

			_, found := seqNrsToWatch[execEvent.SequenceNumber]
			if found && execEvent.SourceChainSelector == sourceSelector {
				t.Logf("Received ExecutionStateChanged (state %s) on chain %d (offramp %s) from chain %d with expected sequence number %d",
					commonutils.ExecutionStateToString(execEvent.State), dest.Selector, offRamp.Address().String(), sourceSelector, execEvent.SequenceNumber)
				executionStates[execEvent.SequenceNumber] = int(execEvent.State)
				delete(seqNrsToWatch, execEvent.SequenceNumber)
				if len(seqNrsToWatch) == 0 {
					return executionStates, nil
				}
			}
		case <-timeout.C:
			return nil, fmt.Errorf("timed out waiting for ExecutionStateChanged on chain %d (offramp %s) from chain %d with expected sequence numbers %+v",
				dest.Selector, offRamp.Address().String(), sourceSelector, expectedSeqNrs)
		case subErr := <-subscription.Err():
			return nil, fmt.Errorf("subscription error: %w", subErr)
		}
	}
}
