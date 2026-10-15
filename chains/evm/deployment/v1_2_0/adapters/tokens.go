package adapters

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	tp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/token_pool"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	_ tokensapi.TokenPoolMigrator             = &TokenAdapter{}
	_ tokensapi.TokenAdapter                  = &TokenAdapter{}
	_ tokensapi.TokenPoolDynamicConfigAdapter = &TokenAdapter{}
	_ tokensapi.RateLimitReaderAdapter        = &TokenAdapter{}
)

// TokenAdapter handles legacy EVM token pools at version 1.2.0.
//
// # Topology
//
// Same shape as v1.4.0: a surviving v1.2.0 pool is never registered in the
// TokenAdminRegistry - the registry did not exist until v1.5.0 - so it sits
// behind a v1.5.0 *TokenPoolAndProxy, which is the registered contract and the
// one the ramps route through:
//
//	TokenAdminRegistry
//	  └─ BurnMintTokenPoolAndProxy 1.5.0   (registered; router = the CCIP Router)
//	       └─ getPreviousPool() -> BurnMintTokenPool 1.2.0
//
// # Why this adapter reads almost nothing off the legacy pool
//
// A v1.2.0 pool is RAMP-keyed, not chain-keyed. It has no applyChainUpdates, no
// getSupportedChains, no isSupportedChain, no setChainRateLimiterConfig and no
// router at all. Lanes are expressed as onRamp/offRamp ADDRESSES
// (applyRampUpdates, getOnRamps, getOffRamps), and a remote chain selector
// appears nowhere on the contract.
//
// Rate limits are stored per ramp. Behind a proxy the proxy is the registered
// ramp, so there is exactly ONE outbound and ONE inbound limiter shared across
// every lane. There is no per-remote-chain granularity to read or write. (The
// v1.5.0 migration path already relies on this; see readV12Limits and the
// shared-limiter warning in v1_5_0/adapters/tokens.go.)
//
// So everything lane-shaped, supported chains, remote token, remote pool, rate
// limits is served by the FRONTING PROXY, which is both the authoritative
// lane contract and the one actually in the registry. The legacy pool is
// consulted only for what is genuinely its own: its token, its owner, and (for
// lock-release) its provided liquidity.
//
// # Scope
//
// Read-only. No v1.2.0 pool is deployed, re-pointed or reconfigured here:
//
//   - no deploy sequence (a new v1.2.0 pool is never correct)
//   - no rate limit writes. The per-ramp limiter is shared across every lane,
//     so a per-lane write would silently retune all of them. The generated
//     v1.2.0 operations package exposes no write operation at all.
//   - no router or admin writes. Neither concept exists on the contract.
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a TokenAdapter with pre-wired PoolOps.
//
// DeployTokenPoolSeq is intentionally nil. EVMPoolAdapter.DeployTokenPoolForToken
// fails cleanly on a nil sequence.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops: &poolOpsV120{},
		},
	}
}

// ConfigureTokenForTransfersSequence is unsupported: a v1.2.0 pool has no
// writable per-lane configuration. Lanes live on the fronting proxy, which the
// v1.5.0 adapter configures.
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.2.0-adapter:configure-token-for-transfers",
		tp.Version,
		"Unsupported: v1.2.0 token pools carry no per-lane configuration",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"configuring transfers is not supported on v1.2.0 token pools (pool %s on chain %d): "+
					"the pool is ramp-keyed and holds no per-chain lane config; configure the v1.5.0 proxy in front of it instead",
				input.TokenPoolAddress, input.ChainSelector)
		},
	)
}

// MigrateLockReleasePoolLiquiditySequence is unsupported on v1.2.0.
//
// Unlike v1.4.0, whose lock-release pool shares getRebalancer/setRebalancer/
// withdrawLiquidity with v1.6.1 and so migrates through the shared v2.0.0
// sequence unchanged, a v1.2.0 lock-release pool exposes only
// addLiquidity/removeLiquidity/getProvidedLiquidity and has no rebalancer at
// all. The shared sequence would revert on its first getRebalancer call, so
// this fails up front with an explanation instead.
//
// Use GetProvidedLiquidity to check whether a legacy pool still holds funds.
func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.2.0-adapter:migrate-lock-release-pool-liquidity",
		tp.Version,
		"Unsupported: v1.2.0 lock-release pools have no rebalancer",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input tokensapi.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"lock-release liquidity migration is not supported on v1.2.0 pools (pool %s on chain %d): "+
					"the contract has no rebalancer and exposes addLiquidity/removeLiquidity rather than "+
					"provideLiquidity/withdrawLiquidity, so the shared v2.0.0 migration sequence cannot drive it",
				input.OldPoolAddress, input.ChainSelector)
		},
	)
}

// GetSupportedChains delegates to the fronting proxy.
//
// A v1.2.0 pool stores ramp addresses, not chain selectors, and behind a proxy
// the only registered ramp is the proxy itself, so the pool's own view carries
// no lane information.
func (t *TokenAdapter) GetSupportedChains(e deployment.Environment, chainSelector uint64, poolAddr, tokenAddr []byte) ([]uint64, error) {
	chain, proxy, err := t.frontingProxy(e.OperationsBundle, e.BlockChains, e.DataStore, chainSelector, common.BytesToAddress(poolAddr), tokenAddr)
	if err != nil {
		return nil, err
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tpap.GetSupportedChains, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSelector, Address: proxy},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get supported chains from proxy %s on chain %d: %w", proxy.Hex(), chainSelector, err)
	}

	return report.Output, nil
}

// GetRemoteToken reads the remote token for a lane from the fronting proxy.
func (t *TokenAdapter) GetRemoteToken(e deployment.Environment, chainSelector uint64, poolAddr, tokenAddr []byte, remoteSelector uint64) ([]byte, error) {
	chain, proxy, err := t.frontingProxy(e.OperationsBundle, e.BlockChains, e.DataStore, chainSelector, common.BytesToAddress(poolAddr), tokenAddr)
	if err != nil {
		return nil, err
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tpap.GetRemoteToken, chain,
		evm_contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: proxy, Args: remoteSelector},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote token for chain %d from proxy %s: %w", remoteSelector, proxy.Hex(), err)
	}
	if len(report.Output) == 0 {
		return nil, fmt.Errorf("proxy %s on chain %d has no remote token registered for chain %d", proxy.Hex(), chainSelector, remoteSelector)
	}

	return report.Output, nil
}

// GetRemotePools reads the remote pool for a lane from the fronting proxy.
//
// The proxy is a v1.5.0 pool and holds at most ONE remote pool per lane, so the
// result is either empty or a single-element slice, and retargeting is a hard
// cutover.
func (t *TokenAdapter) GetRemotePools(e deployment.Environment, chainSelector uint64, poolAddr, tokenAddr []byte, remoteSelector uint64) ([][]byte, error) {
	chain, proxy, err := t.frontingProxy(e.OperationsBundle, e.BlockChains, e.DataStore, chainSelector, common.BytesToAddress(poolAddr), tokenAddr)
	if err != nil {
		return nil, err
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tpap.GetRemotePool, chain,
		evm_contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: proxy, Args: remoteSelector},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote pool for chain %d from proxy %s: %w", remoteSelector, proxy.Hex(), err)
	}
	if len(report.Output) == 0 {
		return [][]byte{}, nil
	}

	return [][]byte{report.Output}, nil
}

// GetOnchainRateLimits reads the lane's rate limits from the fronting proxy.
//
// This overrides EVMPoolAdapter's version, which would delegate to
// PoolOps.GetCurrentRateLimits  and that has no datastore to resolve the proxy
// with. The legacy pool's own limiter is per-ramp and therefore shared across
// every lane, so it cannot answer a per-lane question at all, the proxy can.
func (t *TokenAdapter) GetOnchainRateLimits(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	chainSelector uint64,
	poolRef datastore.AddressRef,
	tokenRef datastore.AddressRef,
	remoteSelector uint64,
	fastFinality bool,
) (tokensapi.OnchainRateLimits, error) {
	if fastFinality {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("fast finality buckets are not supported on v1.2.0 token pools")
	}

	poolAddr, err := t.EVMTokenBase.ParseNonZeroAddressRef(ds, poolRef, chainSelector)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to find token pool address for ref on chain %d: %w", chainSelector, err)
	}

	var tokenAddr []byte
	if addr, err := t.EVMTokenBase.ParseNonZeroAddressRef(ds, tokenRef, chainSelector); err == nil {
		tokenAddr = addr.Bytes()
	}

	chain, proxy, err := t.frontingProxy(b, chains, ds, chainSelector, poolAddr, tokenAddr)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, err
	}

	outbound, err := cldf_ops.ExecuteOperation(b,
		tpap.GetCurrentOutboundRateLimiterState, chain,
		evm_contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: proxy, Args: remoteSelector},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to get outbound rate limiter state for remote chain %d from proxy %s: %w", remoteSelector, proxy.Hex(), err)
	}
	inbound, err := cldf_ops.ExecuteOperation(b,
		tpap.GetCurrentInboundRateLimiterState, chain,
		evm_contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: proxy, Args: remoteSelector},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to get inbound rate limiter state for remote chain %d from proxy %s: %w", remoteSelector, proxy.Hex(), err)
	}

	return tokensapi.OnchainRateLimits{
		Outbound: tokensapi.RateLimiterConfig{IsEnabled: outbound.Output.IsEnabled, Capacity: outbound.Output.Capacity, Rate: outbound.Output.Rate},
		Inbound:  tokensapi.RateLimiterConfig{IsEnabled: inbound.Output.IsEnabled, Capacity: inbound.Output.Capacity, Rate: inbound.Output.Rate},
	}, nil
}

// frontingProxy resolves the v1.5.0 *TokenPoolAndProxy sitting in front of a
// legacy v1.2.0 pool, on the SAME chain.
//
// The link is only discoverable forwards (proxy -> getPreviousPool), so it is
// walked in that direction: the legacy pool's token is looked up in the
// TokenAdminRegistry to find the registered pool, and that pool's
// getPreviousPool is required to point back at the legacy pool. The round trip
// rejects a registered pool that merely shares a token with the legacy pool but
// does not actually front it.
func (t *TokenAdapter) frontingProxy(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	chainSelector uint64,
	poolAddr common.Address,
	tokenAddr []byte,
) (evm.Chain, common.Address, error) {
	chain, ok := chains.EVMChains()[chainSelector]
	if !ok {
		return evm.Chain{}, common.Address{}, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	token := common.BytesToAddress(tokenAddr)
	if token == (common.Address{}) {
		pool, err := tp.NewTokenPoolContract(poolAddr, chain.Client)
		if err != nil {
			return evm.Chain{}, common.Address{}, fmt.Errorf("failed to instantiate v1.2.0 token pool at %s: %w", poolAddr.Hex(), err)
		}
		token, err = pool.GetToken(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return evm.Chain{}, common.Address{}, fmt.Errorf("failed to get local token from pool %s on chain %d: %w", poolAddr.Hex(), chainSelector, err)
		}
	}

	tarAddr, err := t.EVMTokenBase.GetTokenAdminRegistryAddress(ds, chainSelector)
	if err != nil {
		return evm.Chain{}, common.Address{}, fmt.Errorf("failed to resolve TokenAdminRegistry on chain %d: %w", chainSelector, err)
	}

	cfgReport, err := cldf_ops.ExecuteOperation(b,
		tarops.GetTokenConfig, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSelector, Address: tarAddr, Args: token},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[common.Address], evm.Chain](),
	)
	if err != nil {
		return evm.Chain{}, common.Address{}, fmt.Errorf("failed to get token config for token %s from registry %s on chain %d: %w", token.Hex(), tarAddr.Hex(), chainSelector, err)
	}
	proxy := cfgReport.Output.TokenPool
	if proxy == (common.Address{}) {
		return evm.Chain{}, common.Address{}, fmt.Errorf("no pool registered in TokenAdminRegistry %s for token %s on chain %d", tarAddr.Hex(), token.Hex(), chainSelector)
	}
	if proxy == poolAddr {
		return evm.Chain{}, common.Address{}, fmt.Errorf(
			"pool %s on chain %d is registered in the TokenAdminRegistry directly and has no fronting proxy; "+
				"a v1.2.0 pool is expected to sit behind a v1.5.0 TokenPoolAndProxy", poolAddr.Hex(), chainSelector)
	}

	prevReport, err := cldf_ops.ExecuteOperation(b,
		tpap.GetPreviousPool, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSelector, Address: proxy},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return evm.Chain{}, common.Address{}, fmt.Errorf("failed to read getPreviousPool from registered pool %s on chain %d: %w", proxy.Hex(), chainSelector, err)
	}
	if prevReport.Output != poolAddr {
		return evm.Chain{}, common.Address{}, fmt.Errorf(
			"registered pool %s on chain %d does not front legacy pool %s (its previous pool is %s)",
			proxy.Hex(), chainSelector, poolAddr.Hex(), prevReport.Output.Hex())
	}

	return chain, proxy, nil
}

// poolOpsV120 implements PoolOps against the shared v1.2.0 TokenPool surface,
// which BurnMintTokenPool and LockReleaseTokenPool share with byte-identical
// signatures, so nothing here branches on pool type.
type poolOpsV120 struct{}

func (p *poolOpsV120) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		tp.GetToken, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: poolAddr},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v1.2.0: %w", err)
	}
	return res.Output, nil
}

// GetTokenDecimals reads decimals from the token itself: v1.2.0 pools predate
// getTokenDecimals() and store nothing about the local token beyond its address.
func (p *poolOpsV120) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	tokenAddr, err := p.GetToken(b, chain, poolAddr)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.2.0: %w", err)
	}
	res, err := cldf_ops.ExecuteOperation(b,
		erc20.GetDecimals, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chain.Selector, Address: tokenAddr},
	)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.2.0: failed to read decimals from token %s: %w", tokenAddr.Hex(), err)
	}
	return res.Output, nil
}

// GetPoolAdmins returns the owner for both values.
//
// v1.2.0 has no rate limit admin on either pool type - every mutating function
// is onlyOwner - so the owner IS the only authorized caller. Returning it rather
// than the zero address keeps the SkipIfMissingPermissions guard in
// SetTokenPoolRateLimits correct.
func (p *poolOpsV120) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := tp.NewTokenPoolContract(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate v1.2.0 token pool contract at %s: %w", poolAddr.Hex(), err)
	}
	owner, err = pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	return owner, owner, nil
}

// SetRateLimiterConfig is unsupported. A v1.2.0 pool stores one limiter per RAMP,
// and behind a proxy that is a single limiter shared by every lane, so writing a
// per-lane value would silently retune all of them. The limits that actually
// bind are the proxy's; set them there.
func (p *poolOpsV120) SetRateLimiterConfig(_ cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, _ tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error) {
	return nil, fmt.Errorf(
		"setting rate limits is not supported on v1.2.0 token pools (pool %s on chain %d): "+
			"limits are stored per ramp and are shared across every lane, so a per-lane write would affect all of them; "+
			"set the limits on the v1.5.0 proxy in front of this pool instead",
		poolAddr.Hex(), chain.Selector)
}

// SetDynamicPoolConfigs is unsupported: a v1.2.0 pool has no router, no rate
// limit admin and no fee admin. There is nothing dynamic to set.
func (p *poolOpsV120) SetDynamicPoolConfigs(_ cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, router, rlAdmin, feeAdmin *common.Address) ([]evm_contract.WriteOutput, error) {
	switch {
	case router != nil:
		return nil, fmt.Errorf("router is not settable on v1.2.0 token pools (pool %s on chain %d): the contract has no setRouter and reaches the ramps directly", poolAddr.Hex(), chain.Selector)
	case rlAdmin != nil:
		return nil, fmt.Errorf("rate limit admin is not supported on v1.2.0 token pools (pool %s on chain %d)", poolAddr.Hex(), chain.Selector)
	case feeAdmin != nil:
		return nil, fmt.Errorf("fee admin is not supported on v1.2.0 token pools (pool %s on chain %d)", poolAddr.Hex(), chain.Selector)
	}
	return nil, nil
}

// RemoveRemotePools is unsupported: a v1.2.0 pool has no remote pool storage at
// all. Lane membership is expressed as ramp addresses, and removing one via
// applyRampUpdates would tear down the lane rather than a pool entry.
func (p *poolOpsV120) RemoveRemotePools(_ cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, _ []tokensapi.RemotePoolToRemove) ([]evm_contract.WriteOutput, error) {
	return nil, fmt.Errorf(
		"removing individual remote pools is not supported on v1.2.0 token pools (pool %s on chain %d): "+
			"the contract has no remote pool storage and is keyed by ramp address",
		poolAddr.Hex(), chain.Selector)
}

// GetCurrentRateLimits is unreachable in practice: TokenAdapter overrides
// GetOnchainRateLimits so the read resolves the fronting proxy, which this
// signature has no datastore to do. It errors rather than returning the pool's
// own shared per-ramp limiter, which cannot answer a per-lane question.
func (p *poolOpsV120) GetCurrentRateLimits(_ cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, _ uint64, _ bool) (tokensapi.OnchainRateLimits, error) {
	return tokensapi.OnchainRateLimits{}, fmt.Errorf(
		"per-lane rate limits cannot be read from a v1.2.0 token pool (pool %s on chain %d): "+
			"its limiter is per ramp and shared across lanes; read them from the fronting proxy",
		poolAddr.Hex(), chain.Selector)
}

func (p *poolOpsV120) Version() *semver.Version {
	return tp.Version
}
