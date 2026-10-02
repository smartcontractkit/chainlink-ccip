package adapters

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	lrp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/operations/lock_release_token_pool"
	tp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/operations/token_pool"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/sequences/token_pool"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	tpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
	lrBindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/lock_release_token_pool"
	tpbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	_ tokensapi.TokenPoolMigrator             = &TokenAdapter{}
	_ tokensapi.TokenAdapter                  = &TokenAdapter{}
	_ tokensapi.TokenPoolDynamicConfigAdapter = &TokenAdapter{}
	_ tokensapi.RateLimitReaderAdapter        = &TokenAdapter{}
)

// TokenAdapter handles legacy EVM token pools at version 1.4.0.
//
// # Topology
//
// No v1.4.0 pool is reachable from CCIP directly any more. Every surviving
// v1.4.0 pool sits BEHIND a v1.5.0 *TokenPoolAndProxy, which is the contract
// registered in the TokenAdminRegistry and the one the ramps route through:
//
//	TokenAdminRegistry
//	  └─ BurnMintTokenPoolAndProxy 1.5.0   (registered; router = the CCIP Router)
//	       └─ getPreviousPool() -> BurnMintTokenPool 1.4.0   (router = the proxy)
//
// The legacy pool's router is repointed at the proxy, which is what lets the
// proxy satisfy the legacy pool's onlyRouter gate when it delegates. A proxy
// that has a lane configured on itself serves that lane directly and never
// delegates, so for configured lanes the legacy pool is dormant.
//
// # Scope
//
// This adapter is therefore deliberately narrow: it READS legacy pool state and
// supports draining/retiring it during migration. It does not deploy v1.4.0
// pools and does not re-point their router - both would only ever break a live
// proxy delegation.
//
// Writing rate limits to a legacy pool is supported, because the one legitimate
// reason to touch a v1.4.0 pool today is to replicate or tighten the limits
// already set on the proxy in front of it.
//
// Because the pool itself stores no remote token or remote pool, anything about
// the remote side of a lane is read from the fronting proxy (a v1.5.0 pool,
// which does store it).
//
// Pool types: BurnMintTokenPool and LockReleaseTokenPool. Both are driven
// through the shared TokenPool base ops, whose surface they share with
// byte-identical signatures, so nothing here branches on pool type.
//
// Lock-release liquidity is deliberately not handled here: the v2.0.0
// MigrateLockReleasePoolLiquidity sequence drives the old pool through the v1.6.1
// lock-release bindings, whose getRebalancer/setRebalancer/withdrawLiquidity
// signatures v1.4.0 shares, so a v1.4.0 lock-release pool migrates through that
// path unchanged.
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a TokenAdapter with pre-wired PoolOps.
//
// DeployTokenPoolSeq is intentionally left nil: deploying a new v1.4.0 pool is
// never correct today. EVMPoolAdapter.DeployTokenPoolForToken fails cleanly on a
// nil sequence.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops: &poolOpsV140{},
		},
	}
}

func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.4.0-adapter:configure-token-for-transfers",
		tp.Version,
		"Configure a v1.4.0 token pool for cross-chain transfers on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			if !common.IsHexAddress(input.TokenPoolAddress) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address %q is not a valid hex address", input.TokenPoolAddress)
			}

			tpAddr := common.HexToAddress(input.TokenPoolAddress)
			if tpAddr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("token pool address is zero address")
			}

			configureReport, err := cldf_ops.ExecuteSequence(
				b,
				tpSeq.ConfigureTokenPoolForRemoteChains, chain,
				tpSeq.ConfigureTokenPoolForRemoteChainsInput{
					ChainSelector:    input.ChainSelector,
					TokenPoolAddress: tpAddr,
					TokenPoolVersion: tp.Version,
					RemoteChains:     input.RemoteChains,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
			}
			result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)

			return result, nil
		},
	)
}

func (t *TokenAdapter) GetSupportedChains(e deployment.Environment, chainSelector uint64, poolAddr []byte) ([]uint64, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tp.GetSupportedChains, evmChain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr)},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get supported chains from pool %s on chain %d: %w", common.BytesToAddress(poolAddr).Hex(), chainSelector, err)
	}

	return report.Output, nil
}

// GetRemoteToken reads the remote token for a lane from the fronting proxy.
//
// A v1.4.0 pool stores no remote token. The proxy in front of it is a v1.5.0
// pool that does, and it is the contract actually serving the lane, so it is
// the authoritative source.
func (t *TokenAdapter) GetRemoteToken(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([]byte, error) {
	chain, proxy, err := t.frontingProxy(e, chainSelector, common.BytesToAddress(poolAddr))
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
// result is either empty or a single-element slice. Callers that rely on
// multiple remote pools for zero-downtime cutover (e.g.
// MigrationMetadata.LegacyRemotePools) therefore get a single entry here, and
// retargeting is a hard cutover.
func (t *TokenAdapter) GetRemotePools(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([][]byte, error) {
	chain, proxy, err := t.frontingProxy(e, chainSelector, common.BytesToAddress(poolAddr))
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

// frontingProxy resolves the v1.5.0 *TokenPoolAndProxy sitting in front of a
// legacy v1.4.0 pool, on the SAME chain.
//
// The link is only discoverable forwards (proxy -> getPreviousPool), so it is
// walked in that direction: the legacy pool's token is looked up in the
// TokenAdminRegistry to find the registered pool, and that pool's
// getPreviousPool is required to point back at the legacy pool. The round trip
// is what makes this safe. It rejects a registered pool that merely shares a
// token with the legacy pool but does not actually front it.
func (t *TokenAdapter) frontingProxy(e deployment.Environment, chainSelector uint64, poolAddr common.Address) (evm.Chain, common.Address, error) {
	chain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return evm.Chain{}, common.Address{}, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	pool, err := tpbindings.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return evm.Chain{}, common.Address{}, fmt.Errorf("failed to instantiate v1.4.0 token pool at %s: %w", poolAddr.Hex(), err)
	}
	token, err := pool.GetToken(&bind.CallOpts{Context: e.GetContext()})
	if err != nil {
		return evm.Chain{}, common.Address{}, fmt.Errorf("failed to get local token from pool %s on chain %d: %w", poolAddr.Hex(), chainSelector, err)
	}

	tarAddr, err := t.EVMTokenBase.GetTokenAdminRegistryAddress(e.DataStore, chainSelector)
	if err != nil {
		return evm.Chain{}, common.Address{}, fmt.Errorf("failed to resolve TokenAdminRegistry on chain %d: %w", chainSelector, err)
	}

	cfgReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
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
				"a v1.4.0 pool is expected to sit behind a v1.5.0 TokenPoolAndProxy", poolAddr.Hex(), chainSelector)
	}

	prevReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
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

// poolOpsV140 implements PoolOps against the shared v1.4.0 TokenPool base
// surface, so it serves both BurnMintTokenPool and LockReleaseTokenPool.
type poolOpsV140 struct{}

func (p *poolOpsV140) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(
		b,
		tp.GetToken, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v1.4.0: %w", err)
	}
	return res.Output, nil
}

// GetTokenDecimals reads the decimals from the token itself: v1.4.0 pools
// predate getTokenDecimals() (added in v1.5.1) and do not store the local
// token's decimals.
func (p *poolOpsV140) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	tokenAddr, err := p.GetToken(b, chain, poolAddr)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.4.0: %w", err)
	}
	res, err := cldf_ops.ExecuteOperation(
		b,
		erc20.GetDecimals, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       tokenAddr,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.4.0: failed to read decimals from token %s: %w", tokenAddr.Hex(), err)
	}
	return res.Output, nil
}

// GetPoolAdmins returns the pool owner and the rate limit admin.
//
// The rate limit admin concept only exists on LockReleaseTokenPool at v1.4.0:
// its setChainRateLimiterConfig is gated on `s_rateLimitAdmin || owner()`. The
// base TokenPool (and therefore BurnMintTokenPool) gates
// setChainRateLimiterConfig on onlyOwner, so the owner IS the only authorized
// caller and is returned as the rate limit admin. Returning the owner rather
// than the zero address keeps the SkipIfMissingPermissions guard in
// SetTokenPoolRateLimits correct: it treats a caller as authorized if it is
// either the rate limit admin or the pool owner.
func (p *poolOpsV140) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := tpbindings.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate v1.4.0 token pool contract at %s: %w", poolAddr.Hex(), err)
	}
	owner, err = pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}

	// Only LockReleaseTokenPool exposes getRateLimitAdmin. On BurnMintTokenPool
	// the selector is absent, so the call reverts and the owner is the sole
	// authorized caller. Only that case may fall back, a transport-level
	// failure must surface, or a flaky RPC would silently report the owner as
	// the rate limit admin and feed a wrong answer to permission checks.
	lrPool, err := lrBindings.NewLockReleaseTokenPool(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate v1.4.0 lock release token pool contract at %s: %w", poolAddr.Hex(), err)
	}
	rlAdmin, err = lrPool.GetRateLimitAdmin(&bind.CallOpts{Context: ctx})
	if err != nil {
		if isContractRevert(err) {
			return owner, owner, nil
		}
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get rate limit admin of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	return owner, rlAdmin, nil
}

// isContractRevert reports whether err is the contract rejecting the call
// (a revert, or no data returned because the selector does not exist) rather
// than a transport or RPC failure.
func isContractRevert(err error) bool {
	if _, ok := errors.AsType[rpc.DataError](err); ok {
		return true
	}
	return errors.Is(err, bind.ErrNoCode) || strings.Contains(err.Error(), "abi: attempting to unmarshall an empty string")
}

func (p *poolOpsV140) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error) {
	bucket, ok := input.GetBucketForFinality(false)
	if !ok {
		b.Logger.Warnf("skipping rate limiter config for token pool (%s) on chain %d since no default bucket was provided", datastore_utils.SprintRef(input.TokenPoolRef), input.ChainSelector)
		return nil, nil
	}

	// NOTE: v1.4.0 pools reject an enabled bucket whose rate and capacity are
	// both zero (InvalidRateLimitRate), same as v1.5.x.
	outbound, inbound := bucket.OutboundRateLimiterConfig, bucket.InboundRateLimiterConfig
	if outbound.IsEnabled && outbound.Capacity.Cmp(big.NewInt(0)) == 0 && outbound.Rate.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("outbound rate limiter config is enabled but rate and capacity are both zero")
	}
	if inbound.IsEnabled && inbound.Capacity.Cmp(big.NewInt(0)) == 0 && inbound.Rate.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("inbound rate limiter config is enabled but rate and capacity are both zero")
	}

	report, err := cldf_ops.ExecuteOperation(b,
		tp.SetChainRateLimiterConfig, chain,
		evm_contract.FunctionInput[tp.SetChainRateLimiterConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: tp.SetChainRateLimiterConfigArgs{
				OutboundConfig: tp.RateLimiterConfig{
					IsEnabled: outbound.IsEnabled,
					Capacity:  outbound.Capacity,
					Rate:      outbound.Rate,
				},
				InboundConfig: tp.RateLimiterConfig{
					IsEnabled: inbound.IsEnabled,
					Capacity:  inbound.Capacity,
					Rate:      inbound.Rate,
				},
				RemoteChainSelector: input.RemoteChainSelector,
			},
		})
	if err != nil {
		return nil, fmt.Errorf("SetChainRateLimiterConfig v1.4.0: %w", err)
	}
	return []evm_contract.WriteOutput{report.Output}, nil
}

// SetDynamicPoolConfigs updates the rate limit admin on a v1.4.0 pool.
//
// There is no fee admin concept before v2.0, so a non-nil feeAdmin is rejected
// rather than silently dropped. The rate limit admin is only settable on
// LockReleaseTokenPool. On BurnMintTokenPool the owner is the sole authorized
// caller, so a non-nil rlAdmin is rejected there.
//
// Router changes are refused outright. A surviving v1.4.0 pool has its router
// repointed at the v1.5.0 proxy in front of it, which is what lets the proxy
// pass the pool's onlyRouter gate when delegating. Calling setRouter here would
// sever that delegation and break the lane, so a non-nil router is an error
// rather than something to apply or quietly ignore.
func (p *poolOpsV140) SetDynamicPoolConfigs(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, router, rlAdmin, feeAdmin *common.Address) ([]evm_contract.WriteOutput, error) {
	if feeAdmin != nil {
		return nil, fmt.Errorf("fee admin is not supported on v1.4.0 token pools (pool %s on chain %d)", poolAddr.Hex(), chain.Selector)
	}
	if router != nil {
		return nil, fmt.Errorf(
			"refusing to set the router on v1.4.0 token pool %s on chain %d: its router is the v1.5.0 proxy that fronts it, "+
				"and changing it would break the proxy's delegation to this pool", poolAddr.Hex(), chain.Selector)
	}

	var writes []evm_contract.WriteOutput

	if rlAdmin != nil {
		lrPool, err := lrBindings.NewLockReleaseTokenPool(poolAddr, chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to instantiate v1.4.0 lock release token pool contract at %s: %w", poolAddr.Hex(), err)
		}
		current, err := lrPool.GetRateLimitAdmin(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return nil, fmt.Errorf("rate limit admin is not supported on this v1.4.0 token pool (pool %s on chain %d): %w", poolAddr.Hex(), chain.Selector, err)
		}
		if *rlAdmin == current {
			b.Logger.Infof("Rate limit admin already matches desired value for pool %s on chain %d; skipping", poolAddr.Hex(), chain.Selector)
		} else {
			report, err := cldf_ops.ExecuteOperation(b,
				lrp.SetRateLimitAdmin, chain,
				evm_contract.FunctionInput[lrp.SetRateLimitAdminArgs]{
					ChainSelector: chain.Selector,
					Address:       poolAddr,
					Args:          lrp.SetRateLimitAdminArgs{RateLimitAdmin: *rlAdmin},
				})
			if err != nil {
				return nil, fmt.Errorf("SetRateLimitAdmin v1.4.0: %w", err)
			}
			writes = append(writes, report.Output)
		}
	}

	return writes, nil
}

// RemoveRemotePools is not supported on v1.4.0 pools. The pool stores no remote
// pool at all. The pool for a lane is derived from the ramps, so there is
// never a second entry to remove. The only way to drop a lane's config is
// applyChainUpdates with allowed=false, which deletes the ENTIRE remote chain
// config (both rate limiters), not just a pool entry. Doing that silently under
// a "remove remote pools" pipeline would tear down more than the caller asked
// for, so this errors instead. Use ConfigureTokenForTransfers to retarget the
// lane, or drop the chain deliberately.
func (p *poolOpsV140) RemoveRemotePools(_ cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, _ []tokensapi.RemotePoolToRemove) ([]evm_contract.WriteOutput, error) {
	return nil, fmt.Errorf(
		"removing individual remote pools is not supported on v1.4.0 token pools (pool %s on chain %d): "+
			"the contract has no remote pool storage, and applyChainUpdates(allowed=false) would remove the whole remote chain config",
		poolAddr.Hex(), chain.Selector,
	)
}

func (p *poolOpsV140) GetCurrentRateLimits(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, ff bool) (tokensapi.OnchainRateLimits, error) {
	if ff {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("fast finality buckets are not supported on v1.4.0 token pools")
	}

	outboundReport, err := cldf_ops.ExecuteOperation(b,
		tp.GetCurrentOutboundRateLimiterState, chain,
		evm_contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          remoteSelector,
		},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to get outbound rate limiter state for remote chain %d: %w", remoteSelector, err)
	}
	inboundReport, err := cldf_ops.ExecuteOperation(b,
		tp.GetCurrentInboundRateLimiterState, chain,
		evm_contract.FunctionInput[uint64]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          remoteSelector,
		},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to get inbound rate limiter state for remote chain %d: %w", remoteSelector, err)
	}

	return tokensapi.OnchainRateLimits{
		Outbound: tokensapi.RateLimiterConfig{
			IsEnabled: outboundReport.Output.IsEnabled,
			Capacity:  outboundReport.Output.Capacity,
			Rate:      outboundReport.Output.Rate,
		},
		Inbound: tokensapi.RateLimiterConfig{
			IsEnabled: inboundReport.Output.IsEnabled,
			Capacity:  inboundReport.Output.Capacity,
			Rate:      inboundReport.Output.Rate,
		},
	}, nil
}

func (p *poolOpsV140) Version() *semver.Version {
	return tp.Version
}
