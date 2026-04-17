package hooks

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/latest/burn_mint_erc20"
	"golang.org/x/exp/maps"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	fq16 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	fq20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/testhelpers"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	cciphooks "github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

var _ cciphooks.PostProposalCCIPSend = (*EVMPostProposalCCIPSend)(nil)

// fund deployer with at least one token unit so forked sends can pay fees.
var feeTokenFundingAmount = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(20)) // 10 tokens with 18 decimals, adjust as needed for tokens with different decimals

func init() {
	cciphooks.GetPostProposalCCIPSendRegistry().Register(chain_selectors.FamilyEVM, &EVMPostProposalCCIPSend{})
}

// EVMPostProposalCCIPSend provides EVM-specific discovery of CCIP lanes and fee tokens for post-proposal verify.
type EVMPostProposalCCIPSend struct{}

// SkipSend reports whether post-proposal CCIP send verification should be skipped.
// It returns true when the hook environment is not running on an EVM fork.
func (e *EVMPostProposalCCIPSend) SkipSend(env cldf_changeset.ProposalHookEnv) bool {
	envContext := env.ForkContext
	if envContext == nil {
		return true
	}
	evmForkContext, ok := envContext.(*cldf_changeset.EVMForkContext)
	return !ok || evmForkContext == nil
}

// PreSendValidation verifies that srcSel is a valid selector and that the
// source chain exists in the EVM chains configured on env.
func (e *EVMPostProposalCCIPSend) PreSendValidation(env cldf.Environment, srcSel uint64) error {
	_, err := chain_selectors.GetSelectorFamily(srcSel)
	if err != nil {
		return fmt.Errorf("get selector family: %w", err)
	}
	_, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return fmt.Errorf("chain '%d' not found in chain selectors", srcSel)
	}
	return nil
}

// SupportedDestinations returns all destination selectors that have a
// resolvable lane version from the source selector.
func (e *EVMPostProposalCCIPSend) SupportedDestinations(env cldf.Environment, srcSel uint64) ([]uint64, error) {
	allDests, err := e.supportedRemoteChainsWithVersions(env, srcSel)
	if err != nil {
		return nil, err
	}
	return maps.Keys(allDests), nil
}

// AdapterVersionForLane returns the adapter version for the srcSel -> destSel lane.
func (e *EVMPostProposalCCIPSend) AdapterVersionForLane(env cldf.Environment, srcSel, destSel uint64) (*semver.Version, error) {
	allDests, err := e.supportedRemoteChainsWithVersions(env, srcSel)
	if err != nil {
		return nil, err
	}
	for sel, version := range allDests {
		if sel == destSel {
			return version, nil
		}
	}
	return nil, fmt.Errorf("no adapter version found for src %d -> dest %d", srcSel, destSel)
}

// supportedRemoteChainsWithVersions resolves destination selectors and
// their lane versions for a given source selector.
func (e *EVMPostProposalCCIPSend) supportedRemoteChainsWithVersions(env cldf.Environment, srcSel uint64) (map[uint64]*semver.Version, error) {
	resolver := &adapters1_2.LaneVersionResolver{}
	alldests, _, err := resolver.DeriveLaneVersionsForChain(env, srcSel)
	if err != nil {
		return nil, fmt.Errorf("derive all lane versions from source %d: %w", srcSel, err)
	}
	return alldests, nil
}

// SupportedFeeTokens discovers fee tokens configured on the source chain's
// FeeQuoter and ensures the deployer is funded for each token on forked chains.
// The first returned token is always the empty string, representing native fee payment.
func (e *EVMPostProposalCCIPSend) SupportedFeeTokens(env cldf.Environment, srcSel uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
	evmForkContext, ok := forkContext.(*cldf_changeset.EVMForkContext)
	if !ok {
		return nil, fmt.Errorf("invalid fork context type: expected *EVMForkContext, got %T", forkContext)
	}
	if evmForkContext == nil {
		return nil, errors.New("invalid fork context: no fork context found")
	}
	if evmForkContext.ChainConfig.HTTPRPCs == nil || len(evmForkContext.ChainConfig.HTTPRPCs) == 0 {
		return nil, errors.New("invalid fork context: no http rpcs found")
	}
	rpcUrl := evmForkContext.ChainConfig.HTTPRPCs[0].External
	ec, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to eth client for chain %d at rpc %s: %w", srcSel, evmForkContext.ChainConfig.HTTPRPCs[0].External, err)
	}
	defer ec.Close()
	fqAddr, fqVer, err := sequences.GetFeeQuoterAddressAndVersionFromOnRamp(env.DataStore, srcSel, env.BlockChains)
	if err != nil {
		return nil, err
	}
	// we are not using chain.client here to avoid using multi client
	// multi client attempts a lot of retries on failed transactions which causes significant delay in this loop
	// when the node is not fully synced or has issues processing the event filters with anvil
	chain, ok := env.BlockChains.EVMChains()[srcSel]
	if !ok {
		return nil, fmt.Errorf("chain %d not in environment EVM chains", srcSel)
	}
	chain.Client = ec

	// get supported chains
	chains, err := e.SupportedDestinations(env, srcSel)
	if err != nil {
		return nil, err
	}
	feeTokenFundingAmount = new(big.Int).Mul(big.NewInt(int64(len(chains))), feeTokenFundingAmount)
	var addrs []common.Address
	// FeeQuoter bindings differ by major version; select the matching wrapper at runtime.
	switch fqVer.Major() {
	case 1:
		fq, err := fq16.NewFeeQuoter(fqAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("fee quoter 1.x binding: %w", err)
		}
		addrs, err = fq.GetFeeTokens(nil)
		if err != nil {
			return nil, fmt.Errorf("getFeeTokens: %w", err)
		}
	case 2:
		fq, err := fq20.NewFeeQuoter(fqAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("fee quoter 2.x binding: %w", err)
		}
		addrs, err = fq.GetFeeTokens(nil)
		if err != nil {
			return nil, fmt.Errorf("getFeeTokens: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported fee quoter major version %d for chain %d", fqVer.Major(), srcSel)
	}

	var filteredFeeTokens []common.Address
	// Best-effort funding: try to give the deployer fee token balances by impersonating each token owner
	// on forked chains. Tokens that fail discovery, transfer construction, or impersonated send are skipped
	// and excluded from filteredFeeTokens.
	for _, addr := range addrs {
		ctx, cancel := context.WithTimeout(env.GetContext(), 1*time.Minute)
		defer cancel()
		env.Logger.Infof("Processing fee token %s on chain %d", addr.Hex(), srcSel)
		token, err := burn_mint_erc20.NewBurnMintERC20(addr, ec)
		if err != nil {
			return nil, fmt.Errorf("failed to create burn mint erc20 instance: %w", err)
		}
		deployerBal, err := token.BalanceOf(nil, chain.DeployerKey.From)
		if err == nil && deployerBal.Cmp(feeTokenFundingAmount) >= 0 {
			filteredFeeTokens = append(filteredFeeTokens, addr)
			continue
		}
		env.Logger.Infof("Deployer balance for token %s on chain %d is %s, needs funding", addr.Hex(), srcSel, deployerBal.String())
		// Prefer owner() when available; otherwise infer a likely funded account from token events.
		tokenOwner, err := discoverFeeTokenFundingAccount(ctx, ec, token, addr, feeTokenFundingAmount)
		if err != nil {
			// in case of error continue
			env.Logger.Warnf("Failed to discover fee token funding account for token %s on chain %d, continuing without it: %v", addr.Hex(), srcSel, err)
			continue
		}
		tx, err := token.Transfer(cldf.SimTransactOpts(), chain.DeployerKey.From, feeTokenFundingAmount)
		if err != nil {
			// in case of error continue
			env.Logger.Warnf("Failed to create transfer transaction for fee token %s from token owner %s to deployer %s on chain %d, "+
				"continuing without it: %v", addr.Hex(), tokenOwner.Hex(), chain.DeployerKey.From.Hex(), srcSel, err)
			continue
		}
		if err := testhelpers.SendImpersonatedTx(env.GetContext(), ec, rpcUrl, tokenOwner.Hex(), addr.Hex(), tx.Data()); err != nil {
			// in case of error continue
			env.Logger.Warnf("Failed to send impersonated transfer transaction for fee token %s from token owner %s to deployer %s on chain %d, "+
				"continuing without it: %v", addr.Hex(), tokenOwner.Hex(), chain.DeployerKey.From.Hex(), srcSel, err)
			continue
		}
		filteredFeeTokens = append(filteredFeeTokens, addr)
	}
	out := make([]string, 0, len(filteredFeeTokens)+1)
	// Keep native token first (empty string) to mirror adapter expectations.
	out = append(out, "") // native (empty encodes to wrapped native in adapter)
	for _, a := range filteredFeeTokens {
		out = append(out, a.Hex())
	}
	return out, nil
}

// discoverFeeTokenFundingAccount returns an account that can fund fundingAmount
// of tokenAddr, preferring owner() and falling back to event-derived candidates.
func discoverFeeTokenFundingAccount(
	ctx context.Context,
	backend bind.ContractBackend,
	token *burn_mint_erc20.BurnMintERC20,
	tokenAddr common.Address,
	fundingAmount *big.Int,
) (common.Address, error) {
	owner, err := optionalAddressGetter(ctx, backend, tokenAddr, "owner")
	if err != nil {
		return findFundingSenderFromTokenEvents(ctx, backend, token, tokenAddr, fundingAmount)
	}
	ownerBal, err := token.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, owner)
	if err != nil {
		return common.Address{}, fmt.Errorf("fetch owner balance for token %s: %w", tokenAddr.Hex(), err)
	}
	if ownerBal.Cmp(fundingAmount) >= 0 {
		return owner, nil
	}
	// if owner does not have sufficient balance, fall back to finding sender from token events
	return findFundingSenderFromTokenEvents(ctx, backend, token, tokenAddr, fundingAmount)
}

// optionalAddressGetter calls a no-arg address getter on contractAddr and
// returns its result.
func optionalAddressGetter(
	ctx context.Context,
	backend bind.ContractBackend,
	contractAddr common.Address,
	getter string,
) (common.Address, error) {
	// Build a minimal ABI dynamically so this works with tokens that may or may not expose the getter.
	parsed, err := abi.JSON(strings.NewReader(fmt.Sprintf(
		`[{"inputs":[],"name":"%s","outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view","type":"function"}]`,
		getter,
	)))
	if err != nil {
		return common.Address{}, err
	}
	contract := bind.NewBoundContract(contractAddr, parsed, backend, backend, backend)
	var out []any
	if err := contract.Call(&bind.CallOpts{
		Context: ctx,
	}, &out, getter); err != nil {
		return common.Address{}, err
	}
	if len(out) == 0 {
		return common.Address{}, errors.New("empty call result")
	}
	return *abi.ConvertType(out[0], new(common.Address)).(*common.Address), nil
}

// findFundingSenderFromTokenEvents searches recent Transfer/Approval logs and
// returns a sender with at least fundingAmount balance.
func findFundingSenderFromTokenEvents(
	ctx context.Context,
	backend bind.ContractBackend,
	token *burn_mint_erc20.BurnMintERC20,
	tokenAddr common.Address,
	fundingAmount *big.Int,
) (common.Address, error) {
	const (
		chunkSize   = uint64(20_000)
		maxLookback = uint64(500_000)
	)
	header, err := backend.HeaderByNumber(ctx, nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("get latest block header: %w", err)
	}
	if header == nil || header.Number == nil {
		return common.Address{}, errors.New("latest block header is nil")
	}
	latestBlock := header.Number.Uint64()
	startBlock := uint64(0)
	if latestBlock > maxLookback {
		startBlock = latestBlock - maxLookback
	}

	transferSig := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	approvalSig := crypto.Keccak256Hash([]byte("Approval(address,address,uint256)"))

	// Walk newest-to-oldest logs in bounded chunks to quickly find an address that can fund the deployer.
	for toBlock := latestBlock; ; {
		fromBlock := uint64(0)
		if toBlock > chunkSize {
			fromBlock = toBlock - chunkSize + 1
		}
		if fromBlock < startBlock {
			fromBlock = startBlock
		}
		logs, err := backend.FilterLogs(ctx, ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(fromBlock),
			ToBlock:   new(big.Int).SetUint64(toBlock),
			Addresses: []common.Address{tokenAddr},
			Topics:    [][]common.Hash{{transferSig, approvalSig}},
		})
		if err != nil {
			return common.Address{}, fmt.Errorf("filter token logs from %d to %d: %w", fromBlock, toBlock, err)
		}
		// Iterate logs backwards so recent token movement is checked first.
		for i := len(logs) - 1; i >= 0; i-- {
			lg := logs[i]
			if len(lg.Topics) < 2 {
				continue
			}
			sender := common.BytesToAddress(lg.Topics[1].Bytes())
			if sender == (common.Address{}) {
				continue
			}
			bal, err := token.BalanceOf(nil, sender)
			if err != nil {
				continue
			}
			if bal.Cmp(fundingAmount) >= 0 {
				return sender, nil
			}
		}

		if fromBlock == startBlock || fromBlock == 0 {
			break
		}
		toBlock = fromBlock - 1
	}
	return common.Address{}, fmt.Errorf("no sender in recent approval/transfer logs has required balance %s", fundingAmount.String())
}
