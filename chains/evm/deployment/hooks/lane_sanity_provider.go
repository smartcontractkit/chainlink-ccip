package hooks

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/latest/burn_mint_erc20"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	fq20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	adapters1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	cciphooks "github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

func init() {
	cciphooks.GetPostProposalLaneSanityRegistry().Register(chain_selectors.FamilyEVM, &EVMLaneSanityProvider{})
}

var _ cciphooks.PostProposalLaneSanity = (*EVMLaneSanityProvider)(nil)

// EVMLaneSanityProvider implements PostProposalLaneSanity for EVM chains running
// FeeQuoter v2.0. It embeds EVMPostProposalCCIPSend and overrides the methods
// that need v2.0-specific behaviour or that must filter to v2.0 lanes only.
type EVMLaneSanityProvider struct {
	EVMPostProposalCCIPSend
}

// ApplySenderPrivateKey configures every EVM chain in env to send from senderKey.
func (e *EVMLaneSanityProvider) ApplySenderPrivateKey(
	ctx context.Context,
	lggr logger.Logger,
	env *cldf.Environment,
	senderKey string,
) error {
	senderKey = strings.TrimSpace(senderKey)
	if senderKey == "" {
		return nil
	}

	privateKey, err := parseSenderPrivateKey(senderKey)
	if err != nil {
		return err
	}

	senderAddr := crypto.PubkeyToAddress(privateKey.PublicKey)
	lggr.Infof("lane-sanity-check: sending from %s", senderAddr.Hex())
	if len(env.BlockChains.ListChainSelectors(cldf_chain.WithFamily(chain_selectors.FamilyEVM))) == 0 {
		return fmt.Errorf("no EVM chains in environment")
	}
	updatedChains := make(map[uint64]cldf_chain.BlockChain)
	for sel, chain := range env.BlockChains.All() {
		evmChain, ok := chain.(cldf_evm.Chain)
		if !ok {
			updatedChains[sel] = chain
			continue
		}

		chainMeta, ok := chain_selectors.ChainBySelector(sel)
		if !ok {
			return fmt.Errorf("unknown chain selector %d", sel)
		}
		chainID := big.NewInt(int64(chainMeta.EvmChainID))

		transactor, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			return fmt.Errorf("transactor selector=%d: %w", sel, err)
		}
		if evmChain.DeployerKey != nil {
			transactor.GasLimit = evmChain.DeployerKey.GasLimit
			transactor.GasFeeCap = evmChain.DeployerKey.GasFeeCap
			transactor.GasTipCap = evmChain.DeployerKey.GasTipCap
			transactor.Context = evmChain.DeployerKey.Context
		}
		if transactor.Context == nil {
			transactor.Context = ctx
		}

		evmChain.DeployerKey = transactor
		updatedChains[sel] = evmChain
	}

	env.BlockChains = cldf_chain.NewBlockChains(updatedChains)
	return nil
}

func parseSenderPrivateKey(senderKey string) (*ecdsa.PrivateKey, error) {
	keyHex := strings.TrimPrefix(strings.TrimSpace(senderKey), "0x")
	if keyHex == "" {
		return nil, fmt.Errorf("sender private key is empty")
	}
	privateKey, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		return nil, fmt.Errorf("parse sender private key: %w", err)
	}
	return privateKey, nil
}

// AvailableTransferTokens returns transfer tokens with an enabled token-transfer
// fee config on source→dest. Keys are ERC20 symbols; values are source-chain
// token addresses.
func (e *EVMLaneSanityProvider) AvailableTransferTokens(
	env cldf.Environment,
	source, dest uint64,
) (map[string]string, error) {
	chain, ok := env.BlockChains.EVMChains()[source]
	if !ok {
		return nil, fmt.Errorf("source chain %d not found in environment EVM chains", source)
	}
	tarAddr, err := datastore_utils.FindAndFormatRef(
		env.DataStore,
		datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}, source,
		evm_datastore_utils.ToEVMAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("finding token_admin_registry contract address on source %d: %w", source, err)
	}
	tokensPerChain, err := adapters1_5.GetSupportedTokensPerRemoteChain(env.GetContext(), env.Logger, tarAddr, chain, []uint64{dest})
	if err != nil {
		return nil, fmt.Errorf("getting supported tokens per remote chain from token admin registry on source %d to %d: %w", source, dest, err)
	}
	TransferTokensForDest, ok := tokensPerChain[dest]
	if !ok {
		return nil, fmt.Errorf("no supported tokens found for destination %d in token admin registry on source %d", dest, source)
	}

	result := make(map[string]string)
	for _, token := range TransferTokensForDest {
		tokenC, err := erc20.NewERC20(token, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to create ERC20 client for token %s on chain %d: %w", token.Hex(), source, err)
		}
		name, err := tokenC.Symbol(&bind.CallOpts{Context: env.GetContext()})
		if err != nil {
			return nil, err
		}
		result[name] = token.Hex()
	}

	return result, nil
}

// EncodeReceiverAddress ABI-encodes an EVM address as a 32-byte CCIP receiver.
func (e *EVMLaneSanityProvider) EncodeReceiverAddress(
	env cldf.Environment,
	destSel uint64,
	receiverAddress string,
) ([]byte, error) {
	if !common.IsHexAddress(receiverAddress) {
		return nil, fmt.Errorf("invalid EVM receiver address: %s", receiverAddress)
	}
	return common.LeftPadBytes(common.HexToAddress(receiverAddress).Bytes(), 32), nil
}

// SupportedDestinations overrides the embedded implementation to return only
// destination selectors whose lane version is v2.0, since all methods in this
// provider require FeeQuoter v2.0.
func (e *EVMLaneSanityProvider) SupportedDestinations(env cldf.Environment, srcSel uint64) ([]uint64, error) {
	allDests, err := e.EVMPostProposalCCIPSend.SupportedDestinations(env, srcSel)
	if err != nil {
		return nil, err
	}
	var v2Dests []uint64
	for _, destSel := range allDests {
		ver, err := e.EVMPostProposalCCIPSend.AdapterVersionForLane(env, srcSel, destSel)
		if err != nil {
			env.Logger.Debugf("lane-sanity: AdapterVersionForLane src=%d dest=%d: %v (skipping)", srcSel, destSel, err)
			continue
		}
		if ver.Major() == 2 {
			v2Dests = append(v2Dests, destSel)
		}
	}
	return v2Dests, nil
}

// MockReceiverAddress looks up the MockReceiverV2 contract for chainSel in the
// datastore and returns its address left-padded to 32 bytes (EVM ABI encoding).
// Returns nil, nil when no MockReceiverV2 is deployed.
func (e *EVMLaneSanityProvider) MockReceiverAddress(
	env cldf.Environment,
	chainSel uint64,
) ([]byte, error) {
	refs := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSel),
		datastore.AddressRefByType(datastore.ContractType(mock_receiver.ContractType)),
	)
	if len(refs) == 0 {
		return nil, nil //nolint:nilnil
	}
	if len(refs) > 1 {
		env.Logger.Warnf("lane-sanity: multiple MockReceiverV2 on chain %d, using first", chainSel)
	}
	addr := common.HexToAddress(refs[0].Address)
	return common.LeftPadBytes(addr.Bytes(), 32), nil
}

// GetMessageFee calls the source chain Router's GetFee for the given message.
// The msg parameter must be a router.ClientEVM2AnyMessage as returned by the
// EVM adapter's BuildMessage. Returns an empty string without error when the
// cast fails (e.g. cross-family lane) so the caller omits the fee line.
func (e *EVMLaneSanityProvider) GetMessageFee(
	ctx context.Context,
	env cldf.Environment,
	srcSel, destSel uint64,
	msg any,
) (string, error) {
	evm2anyMsg, ok := msg.(router.ClientEVM2AnyMessage)
	if !ok {
		return "", nil //nolint:nilnil
	}

	chain, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return "", fmt.Errorf("chain %d not in environment", srcSel)
	}

	routerAddr, err := resolveEVMRouterAddress(env.DataStore, srcSel)
	if err != nil {
		return "", fmt.Errorf("router address chain=%d: %w", srcSel, err)
	}

	r, err := router.NewRouter(routerAddr, chain.Client)
	if err != nil {
		return "", fmt.Errorf("router binding chain=%d: %w", srcSel, err)
	}

	fee, err := r.GetFee(&bind.CallOpts{Context: ctx}, destSel, evm2anyMsg)
	if err != nil {
		return "", fmt.Errorf("GetFee src=%d dest=%d: %w", srcSel, destSel, err)
	}

	return fee.String(), nil
}

// SupportedFeeTokens overrides the embedded EVMPostProposalCCIPSend method so
// that a nil forkContext (real environment / lane-sanity path) is handled
// without impersonation: the v2.0 FeeQuoter is queried directly and all
// configured fee tokens are returned, trusting the deployer is already funded.
// A non-nil forkContext delegates to the embedded implementation which funds
// the deployer via Anvil impersonation before returning the token list.
func (e *EVMLaneSanityProvider) SupportedFeeTokens(
	env cldf.Environment,
	srcSel uint64,
	forkContext cldf_changeset.ForkContext,
) ([]string, error) {
	if forkContext != nil {
		return e.EVMPostProposalCCIPSend.SupportedFeeTokens(env, srcSel, forkContext)
	}
	return e.queryFeeTokensFromChain(env, srcSel)
}

// queryFeeTokensFromChain queries the v2.0 FeeQuoter on srcSel and returns
// configured ERC20 fee tokens with native ("") appended last.
func (e *EVMLaneSanityProvider) queryFeeTokensFromChain(
	env cldf.Environment,
	srcSel uint64,
) ([]string, error) {
	fq, err := e.feeQuoterV2(env, srcSel)
	if err != nil {
		return nil, err
	}

	allFeeTokens, err := fq.GetFeeTokens(&bind.CallOpts{Context: env.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("GetFeeTokens src=%d: %w", srcSel, err)
	}
	out := make([]string, 0, len(allFeeTokens)+1)
	for _, token := range allFeeTokens {
		out = append(out, token.String())
	}
	out = append(out, "") // native last
	return out, nil
}

// FundAndApproveTransferToken verifies that the sender on srcSel holds at
// least one smallest token unit (1/10^decimals of a whole token) of tokenAddress,
// then approves the Router to spend that amount.
//
// Lane sanity checks always run against a real environment where the sender is
// pre-funded. If the balance is insufficient an error is returned immediately so
// the caller fails fast rather than silently skipping the transfer.
func (e *EVMLaneSanityProvider) FundAndApproveTransferToken(
	ctx context.Context,
	env cldf.Environment,
	srcSel uint64,
	tokenAddress string,
) (*big.Int, error) {
	if !common.IsHexAddress(tokenAddress) {
		return nil, fmt.Errorf("invalid EVM token address: %s", tokenAddress)
	}

	chain, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return nil, fmt.Errorf("chain %d not in environment", srcSel)
	}

	tokenAddr := common.HexToAddress(tokenAddress)
	token, err := burn_mint_erc20.NewBurnMintERC20(tokenAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("BurnMintERC20 binding %s: %w", tokenAddress, err)
	}

	// 1 smallest unit = 1/10^decimals of a whole token (e.g. 1e-18 for 18 decimals).
	minTransferAmount := big.NewInt(1)

	senderAddr := chain.DeployerKey.From
	balance, err := token.BalanceOf(&bind.CallOpts{Context: env.GetContext()}, senderAddr)
	if err != nil {
		return nil, fmt.Errorf("balanceOf sender token=%s chain=%d: %w", tokenAddress, srcSel, err)
	}
	if balance.Cmp(minTransferAmount) < 0 {
		return nil, fmt.Errorf(
			"sender %s has insufficient balance for token %s on chain %d: have %s, need %s (1 smallest unit); ensure the sender is funded before running lane sanity checks",
			senderAddr.Hex(), tokenAddress, srcSel, balance.String(), minTransferAmount.String(),
		)
	}

	routerAddr, err := resolveEVMRouterAddress(env.DataStore, srcSel)
	if err != nil {
		return nil, fmt.Errorf("router address chain=%d: %w", srcSel, err)
	}

	approveTx, err := token.Approve(chain.DeployerKey, routerAddr, minTransferAmount)
	if err != nil {
		return nil, fmt.Errorf("approve router=%s token=%s chain=%d: %w",
			routerAddr.Hex(), tokenAddress, srcSel, err)
	}
	if _, err := chain.Confirm(approveTx); err != nil {
		return nil, fmt.Errorf("confirm approve token=%s chain=%d: %w", tokenAddress, srcSel, err)
	}

	return minTransferAmount, nil
}

// feeQuoterV2 resolves the v2.0 FeeQuoter binding for srcSel. Returns an error
// when the chain's FeeQuoter is not v2.0 so callers can skip non-v2.0 chains.
func (e *EVMLaneSanityProvider) feeQuoterV2(env cldf.Environment, srcSel uint64) (*fq20.FeeQuoter, error) {
	chain, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return nil, fmt.Errorf("chain %d not in environment", srcSel)
	}
	fqAddr, fqVer, err := sequences.GetFeeQuoterAddressAndVersionFromOnRamp(env.DataStore, srcSel, env.BlockChains)
	if err != nil {
		return nil, fmt.Errorf("fee quoter on chain %d: %w", srcSel, err)
	}
	if fqVer.Major() != 2 {
		return nil, fmt.Errorf("lane sanity checks require FeeQuoter v2.0, got v%d on chain %d", fqVer.Major(), srcSel)
	}
	fq, err := fq20.NewFeeQuoter(fqAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("fee quoter v2.0 binding src=%d: %w", srcSel, err)
	}
	return fq, nil
}

// resolveEVMRouterAddress returns the production Router address for chainSel.
func resolveEVMRouterAddress(ds datastore.DataStore, chainSel uint64) (common.Address, error) {
	refs := ds.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSel),
		datastore.AddressRefByType(datastore.ContractType("Router")),
	)
	if len(refs) == 0 {
		return common.Address{}, fmt.Errorf("no Router found for chain %d", chainSel)
	}
	if !common.IsHexAddress(refs[0].Address) {
		return common.Address{}, fmt.Errorf("invalid Router address %q for chain %d", refs[0].Address, chainSel)
	}
	return common.HexToAddress(refs[0].Address), nil
}
