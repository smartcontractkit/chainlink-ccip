package testadapter

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	ton_onramp "github.com/smartcontractkit/chainlink-ton/pkg/ccip/bindings/onramp"
	"github.com/xssnick/tonutils-go/tlb"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/onramp"

	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/common/extraargs"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

func init() {
	testadapters.GetTestAdapterRegistry().RegisterForkCCIPSendTestAdapter(chain_selectors.FamilyEVM, semver.MustParse("2.0.0"), NewEVMForkCCIPSendTestAdapter)
	testadapters.GetTestAdapterRegistry().RegisterTestAdapterForFamily(chain_selectors.FamilyEVM, semver.MustParse("2.0.0"), NewEVMTestAdapterForFamily)
}

type EVMAdapter struct {
	state testadapters.StateProvider
	cldf_evm.Chain
}

func NewEVMForkCCIPSendTestAdapter(env *deployment.Environment, selector uint64) testadapters.ForkCCIPSendTestAdapter {
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

func NewEVMTestAdapterForFamily(ds datastore.DataStore, selector uint64) testadapters.TestAdapterForFamily {
	return &EVMAdapter{
		state: &testadapters.DataStoreStateProvider{Selector: selector, DS: ds},
	}
}

var ErrNoAddressFound = errors.New("no address found")

func (a *EVMAdapter) getAddress(ty datastore.ContractType) (common.Address, error) {
	addr, err := a.state.GetAddress(ty)
	if err != nil {
		if strings.HasPrefix(err.Error(), "expected to find exactly 1 ref with criteria") &&
			strings.HasSuffix(err.Error(), ", found 0") {
			return common.Address{}, ErrNoAddressFound
		}
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

func (a *EVMAdapter) getRouter() (*router.Router, error) {
	// if TestRouter env var is set use Test Router instead
	// We are leveraging env var here to avoid making change to the SendMessage method signature
	// TestRouter is only valid for evm
	contractType := datastore.ContractType(routerops.ContractType)
	isTest, err := strconv.ParseBool(strings.TrimSpace(os.Getenv("TestRouter")))
	if err == nil && isTest {
		fmt.Println("Using Test Router for sending message")
		contractType = datastore.ContractType(routerops.TestRouterContractType)
	}
	rAddr, err := a.getAddress(contractType)
	if err != nil {
		return nil, fmt.Errorf("failed to get router address: %w", err)
	}
	return router.NewRouter(
		rAddr,
		a.Client)
}

func (a *EVMAdapter) SendMessage(ctx context.Context, destChainSelector uint64, m any) (uint64, string, error) {
	l := zerolog.Ctx(ctx)
	l.Info().Msg("Sending CCIP message")

	messageID := ""

	msg, ok := m.(router.ClientEVM2AnyMessage)
	if !ok {
		return 0, messageID, errors.New("expected router.ClientEVM2AnyMessage")
	}
	r, err := a.getRouter()
	if err != nil {
		return 0, messageID, fmt.Errorf("failed to create router instance: %w", err)
	}
	const errCodeInsufficientFee = "0x07da6ee6"
	const cannotDecodeErrorReason = "could not decode error reason"
	const errMsgMissingTrieNode = "missing trie node"
	sender := a.DeployerKey
	defer func() { sender.Value = nil }()

	onRampAddr, err := r.GetOnRamp(nil, destChainSelector)
	if err != nil {
		return 0, messageID, fmt.Errorf("failed to get onramp address: %w", err)
	}
	onRamp, err := onramp.NewOnRamp(
		onRampAddr,
		a.Client)
	if err != nil {
		return 0, messageID, fmt.Errorf("failed to create onramp instance: %w", err)
	}
	l.Info().Msg("Got contract instances, preparing to send CCIP message")

	var retryCount int
	for {
		fee, err := r.GetFee(&bind.CallOpts{Context: ctx}, destChainSelector, msg)
		if err != nil {
			return 0, messageID, fmt.Errorf("failed to get EVM fee: %w", deployment.MaybeDataErr(err))
		}
		if msg.FeeToken == (common.Address{}) || msg.FeeToken == common.HexToAddress(a.NativeFeeToken()) {
			sender.Value = fee
		} else {
			err := a.AllowRouterToWithdrawTokens(ctx, msg.FeeToken.Hex(), new(big.Int).Add(fee, fee)) // approve 2x the fee to be safe
			if err != nil {
				return 0, messageID, fmt.Errorf("failed to approve tokens for fee: %w", err)
			}
		}

		tx, err := r.CcipSend(sender, destChainSelector, msg)
		if err != nil {
			return 0, messageID, fmt.Errorf("failed to send CCIP message: %w", err)
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
					return 0, messageID, fmt.Errorf("failed to confirm CCIP message after %d retries: %w", retryCount, deployment.MaybeDataErr(err))
				}
				retryCount++
				continue
			}

			return 0, messageID, fmt.Errorf("failed to confirm CCIP message: %w", deployment.MaybeDataErr(err))
		}
		fmt.Printf("CCIP message sent in block %d with tx %s", blockNum, tx.Hash().Hex())
		it, err := onRamp.FilterCCIPMessageSent(&bind.FilterOpts{
			Start:   blockNum,
			End:     &blockNum,
			Context: ctx,
		}, []uint64{destChainSelector}, []common.Address{sender.From}, nil)
		if err != nil {
			return 0, messageID, fmt.Errorf("failed to filter CCIPMessageSent events: %w", err)
		}

		if !it.Next() {
			return 0, messageID, fmt.Errorf("no CCIP message sent event found")
		}
		messageID = hex.EncodeToString(it.Event.MessageId[:])

		fmt.Printf("Sent CCIP message %+v id %s from chain %d to chain %d\n", msg, messageID, a.Selector, destChainSelector)
		return 0, messageID, nil
	}
}

func (a *EVMAdapter) receiverAddr() []byte {
	for _, typeStr := range []datastore.ContractType{"CCIPReceiver", "TestReceiver"} {
		receiverAddr, err := a.getAddress(typeStr)
		if err == nil {
			return common.LeftPadBytes(receiverAddr.Bytes(), 32)
		}
		if !errors.Is(err, ErrNoAddressFound) {
			panic(err)
		}
	}
	return nil
}

func (a *EVMAdapter) CCIPReceiver() []byte {
	result := a.receiverAddr()
	if result != nil {
		return result
	}
	// Fallback because receiver is not found in a.state for devenv. TODO investigate why and remove fallback
	return common.LeftPadBytes(common.HexToAddress("0xdead").Bytes(), 32)
}

func (a *EVMAdapter) EOAReceiver(t *testing.T) []byte {
	// Return the deployer's wallet address as the EOA receiver for testing purposes.
	return common.LeftPadBytes(a.DeployerKey.From.Bytes(), 32)
}

func (a *EVMAdapter) InvalidAddresses() [][]byte {
	return [][]byte{
		[]byte{99}, // invalid address
		common.LeftPadBytes(common.Address{}.Bytes(), 32), // evm zero address
	}
}

func (a *EVMAdapter) NativeFeeToken() string {
	return "0x0"
}

func (a *EVMAdapter) GetExtraArgs(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error) {
	switch sourceFamily {
	case chain_selectors.FamilyEVM:
		extraArgs := extraargs.ClientGenericExtraArgsV2{
			GasLimit:                 new(big.Int).SetUint64(100_000),
			AllowOutOfOrderExecution: true,
		}
		for _, opt := range opts {
			switch opt.Name {
			case testadapters.ExtraArgGasLimit:
				extraArgs.GasLimit = opt.Value.(*big.Int)
			case testadapters.ExtraArgOOO:
				extraArgs.AllowOutOfOrderExecution = opt.Value.(bool)
			default:
				// unsupported arg
			}
		}
		return extraargs.SerializeClientGenericExtraArgsV2(extraArgs)
	case chain_selectors.FamilySolana:
		// EVM allows empty extraArgs
		return nil, nil
	case chain_selectors.FamilyTon:
		// TODO: maybe for 1.6 we should look up the source adapter and use a 1.6 method to encode? would be good to avoid other chain SDKs
		extraArgs := ton_onramp.GenericExtraArgsV2{
			GasLimit:                 big.NewInt(1000000),
			AllowOutOfOrderExecution: true,
		}
		for _, opt := range opts {
			switch opt.Name {
			case testadapters.ExtraArgGasLimit:
				extraArgs.GasLimit = opt.Value.(*big.Int)
			case testadapters.ExtraArgOOO:
				extraArgs.AllowOutOfOrderExecution = opt.Value.(bool)
			default:
				// unsupported arg
			}
		}
		extraArgsCell, err := tlb.ToCell(extraArgs)
		if err != nil {
			return nil, err
		}
		return extraArgsCell.ToBOC(), nil
	default:
		// TODO: add support for other families
		return nil, fmt.Errorf("unsupported source family: %s", sourceFamily)
	}
}

func (a *EVMAdapter) AllowRouterToWithdrawTokens(ctx context.Context, tokenAddress string, amount *big.Int) error {
	if !common.IsHexAddress(tokenAddress) {
		return fmt.Errorf("invalid token address: %s", tokenAddress)
	}
	if amount.Cmp(big.NewInt(0)) <= 0 {
		return fmt.Errorf("amount must be greater than zero: %s", amount.String())
	}
	r, err := a.getRouter()
	if err != nil {
		return fmt.Errorf("failed to get router address: %w", err)
	}

	routerAddr := r.Address()
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
