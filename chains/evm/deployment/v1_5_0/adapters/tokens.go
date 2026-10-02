package adapters

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	v1_4_0_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"
	v1_2_0_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/burn_mint_token_pool"
	v1_0_0_lock_release_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/lock_release_token_pool"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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
	// RateLimitReaderAdapter is load-bearing, not incidental: ConfigureTokensForTransfers hard
	// errors during auto-migrate discovery if an adapter implements TokenPoolMigrator without it.
	_ tokensapi.RateLimitReaderAdapter = &TokenAdapter{}
)

// TokenAdapter handles EVM token pools at version 1.5.0.
// It embeds EVMPoolAdapter for shared datastore/TAR/BnM logic and
// overrides only ConfigureTokenForTransfersSequence which inlines
// the v1.5.0-specific configure + register flow.
//
// Scope: BurnMintTokenPoolAndProxy and LockReleaseTokenPoolAndProxy. Both are driven through
// the shared TokenPoolAndProxy base ops, whose surface they share with byte-identical
// signatures, so nothing here branches on pool type. The remaining v1.5.0 pool contracts
// (BurnWithFromMintTokenPoolAndProxy and BurnWithFromMintRebasingTokenPool) have generated
// bindings but no deployed footprint, and the deploy sequence rejects them.
//
// Lock-release liquidity is deliberately not handled here: the v2.0.0
// MigrateLockReleasePoolLiquidity sequence drives the old pool through the v1.6.1 lock-release
// bindings, whose getRebalancer/setRebalancer/withdrawLiquidity signatures v1.5.0 shares, so a
// v1.5.0 lock-release pool migrates through that path unchanged.
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a TokenAdapter with pre-wired PoolOps and
// the deploy-token-pool sequence.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops:                &poolOpsV150{},
			DeployTokenPoolSeq: tpSeq.DeployTokenPool,
		},
	}
}

func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.0-adapter:configure-token-for-transfers",
		tpap.Version,
		"Configure a v1.5.0 token pool for cross-chain transfers on an EVM chain",
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

			externalAdmin := common.Address{}
			if input.ExternalAdmin != "" {
				if !common.IsHexAddress(input.ExternalAdmin) {
					return sequences.OnChainOutput{}, fmt.Errorf("external admin address %q is not a valid hex address", input.ExternalAdmin)
				}
				externalAdmin = common.HexToAddress(input.ExternalAdmin)
			}

			tarAddress, err := t.EVMTokenBase.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token admin registry address for chain %d: %w", input.ChainSelector, err)
			}

			tokenAddress, err := t.Ops.GetToken(b, chain, tpAddr)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from pool at %s: %w", tpAddr, err)
			}

			configureReport, err := cldf_ops.ExecuteSequence(
				b,
				tpSeq.ConfigureTokenPoolForRemoteChains, chain,
				tpSeq.ConfigureTokenPoolForRemoteChainsInput{
					ChainSelector:    input.ChainSelector,
					TokenPoolAddress: tpAddr,
					TokenPoolVersion: tpap.Version,
					RemoteChains:     input.RemoteChains,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
			}
			result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)

			registerReport, err := cldf_ops.ExecuteSequence(
				b,
				tarseq.RegisterToken, chain,
				tarseq.RegisterTokenInput{
					ChainSelector:             input.ChainSelector,
					TokenAdminRegistryAddress: tarAddress,
					TokenPoolAddress:          tpAddr,
					ExternalAdmin:             externalAdmin,
					TokenAddress:              tokenAddress,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to register token on chain %d: %w", input.ChainSelector, err)
			}
			result.Addresses = append(result.Addresses, registerReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, registerReport.Output.BatchOps...)

			return result, nil
		},
	)
}

func (t *TokenAdapter) GetSupportedChains(e deployment.Environment, chainSelector uint64, poolAddr, _ []byte) ([]uint64, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tpap.GetSupportedChains, evmChain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr)},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get supported chains from pool %s on chain %d: %w", common.BytesToAddress(poolAddr).Hex(), chainSelector, err)
	}

	return report.Output, nil
}

func (t *TokenAdapter) GetRemoteToken(e deployment.Environment, chainSelector uint64, poolAddr, _ []byte, remoteSelector uint64) ([]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tpap.GetRemoteToken, evmChain,
		evm_contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr), Args: remoteSelector},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote token for chain %d from pool %s: %w", remoteSelector, common.BytesToAddress(poolAddr).Hex(), err)
	}

	if len(report.Output) == 0 {
		return nil, fmt.Errorf("pool %s has no remote token registered for chain %d", common.BytesToAddress(poolAddr).Hex(), remoteSelector)
	}

	return report.Output, nil
}

// GetRemotePools adapts v1.5.0's singular getRemotePool to the plural TokenPoolMigrator contract.
// A v1.5.0 pool holds at most ONE remote pool per lane, so the result is either empty or a
// single-element slice. Callers that rely on multiple remote pools for zero-downtime cutover
// (e.g. MigrationMetadata.LegacyRemotePools) therefore get a single entry here, and retargeting a
// v1.5.0 pool is a hard cutover - see the note on ConfigureTokenPoolForRemoteChain.
func (t *TokenAdapter) GetRemotePools(e deployment.Environment, chainSelector uint64, poolAddr, _ []byte, remoteSelector uint64) ([][]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		tpap.GetRemotePool, evmChain,
		evm_contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr), Args: remoteSelector},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote pool for chain %d from pool %s: %w", remoteSelector, common.BytesToAddress(poolAddr).Hex(), err)
	}

	if len(report.Output) == 0 {
		return [][]byte{}, nil
	}

	return [][]byte{report.Output}, nil
}

// poolOpsV150 implements PoolOps against the shared v1.5.0 TokenPoolAndProxy base surface,
// so it serves both BurnMintTokenPoolAndProxy and LockReleaseTokenPoolAndProxy.
type poolOpsV150 struct{}

func (p *poolOpsV150) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(
		b,
		tpap.GetToken, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v1.5.0: %w", err)
	}
	return res.Output, nil
}

// GetTokenDecimals reads the decimals from the token itself: v1.5.0 pools predate
// getTokenDecimals() (added in v1.5.1) and do not store the local token's decimals.
func (p *poolOpsV150) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	tokenAddr, err := p.GetToken(b, chain, poolAddr)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.5.0: %w", err)
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
		return 0, fmt.Errorf("GetTokenDecimals v1.5.0: failed to read decimals from token %s: %w", tokenAddr.Hex(), err)
	}
	return res.Output, nil
}

func (p *poolOpsV150) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := tpap.NewTokenPoolAndProxyContract(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate v1.5.0 token pool contract at %s: %w", poolAddr.Hex(), err)
	}
	owner, err = pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	rlAdmin, err = pool.GetRateLimitAdmin(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get rate limit admin of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	return owner, rlAdmin, nil
}

func (p *poolOpsV150) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error) {
	bucket, ok := input.GetBucketForFinality(false)
	if !ok {
		b.Logger.Warnf("skipping rate limiter config for token pool (%s) on chain %d since no default bucket was provided", datastore_utils.SprintRef(input.TokenPoolRef), input.ChainSelector)
		return nil, nil
	}

	// NOTE: v1.5.0 pools reject an enabled bucket whose rate and capacity are both zero
	// (InvalidRateLimitRate), same as v1.5.1.
	outbound, inbound := bucket.OutboundRateLimiterConfig, bucket.InboundRateLimiterConfig
	if outbound.IsEnabled && outbound.Capacity.Cmp(big.NewInt(0)) == 0 && outbound.Rate.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("outbound rate limiter config is enabled but rate and capacity are both zero")
	}
	if inbound.IsEnabled && inbound.Capacity.Cmp(big.NewInt(0)) == 0 && inbound.Rate.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("inbound rate limiter config is enabled but rate and capacity are both zero")
	}

	report, err := cldf_ops.ExecuteOperation(b,
		tpap.SetChainRateLimiterConfig, chain,
		evm_contract.FunctionInput[tpap.SetChainRateLimiterConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: tpap.SetChainRateLimiterConfigArgs{
				OutboundConfig: tpap.Config{
					IsEnabled: outbound.IsEnabled,
					Capacity:  outbound.Capacity,
					Rate:      outbound.Rate,
				},
				InboundConfig: tpap.Config{
					IsEnabled: inbound.IsEnabled,
					Capacity:  inbound.Capacity,
					Rate:      inbound.Rate,
				},
				RemoteChainSelector: input.RemoteChainSelector,
			},
		})
	if err != nil {
		return nil, fmt.Errorf("SetChainRateLimiterConfig v1.5.0: %w", err)
	}
	return []evm_contract.WriteOutput{report.Output}, nil
}

// SetDynamicPoolConfigs updates the router and rate limit admin on a v1.5.0 pool. There is no
// fee admin concept before v2.0, so a non-nil feeAdmin is rejected rather than silently dropped.
func (p *poolOpsV150) SetDynamicPoolConfigs(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, router, rlAdmin, feeAdmin *common.Address) ([]evm_contract.WriteOutput, error) {
	if feeAdmin != nil {
		return nil, fmt.Errorf("fee admin is not supported on v1.5.0 token pools (pool %s on chain %d)", poolAddr.Hex(), chain.Selector)
	}

	pool, err := tpap.NewTokenPoolAndProxyContract(poolAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate v1.5.0 token pool contract at %s: %w", poolAddr.Hex(), err)
	}

	var writes []evm_contract.WriteOutput

	if router != nil {
		current, err := pool.GetRouter(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return nil, fmt.Errorf("failed to get router of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
		}
		if *router == current {
			b.Logger.Infof("Router already matches desired value for pool %s on chain %d; skipping", poolAddr.Hex(), chain.Selector)
		} else {
			report, err := cldf_ops.ExecuteOperation(b,
				tpap.SetRouter, chain,
				evm_contract.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       poolAddr,
					Args:          *router,
				})
			if err != nil {
				return nil, fmt.Errorf("SetRouter v1.5.0: %w", err)
			}
			writes = append(writes, report.Output)
		}
	}

	if rlAdmin != nil {
		current, err := pool.GetRateLimitAdmin(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return nil, fmt.Errorf("failed to get rate limit admin of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
		}
		if *rlAdmin == current {
			b.Logger.Infof("Rate limit admin already matches desired value for pool %s on chain %d; skipping", poolAddr.Hex(), chain.Selector)
		} else {
			report, err := cldf_ops.ExecuteOperation(b,
				tpap.SetRateLimitAdmin, chain,
				evm_contract.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       poolAddr,
					Args:          *rlAdmin,
				})
			if err != nil {
				return nil, fmt.Errorf("SetRateLimitAdmin v1.5.0: %w", err)
			}
			writes = append(writes, report.Output)
		}
	}

	return writes, nil
}

// RemoveRemotePools removes remote pool entries from a v1.5.0 pool. v1.5.0 stores a single remote
// pool per remote chain and has no removeRemotePool, so a removal clears that slot with
// setRemotePool(remoteChainSelector, <empty bytes>) when it holds the requested pool. An empty slot
// makes releaseOrMint reject every source pool from that chain (its configured-pool length check),
// the same effect as removeRemotePool on later versions, while the remote chain config (remote
// token, rate limits) is kept. Empty bytes is used rather than an encoded address(0) so the slot
// reads back as "no remote pool" instead of a zero-address pool. A slot that holds a different
// pool, or is already empty, is skipped with a warning so re-runs are idempotent.
func (p *poolOpsV150) RemoveRemotePools(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remotes []tokensapi.RemotePoolToRemove) ([]evm_contract.WriteOutput, error) {
	var writes []evm_contract.WriteOutput
	for _, remote := range remotes {
		poolReport, err := cldf_ops.ExecuteOperation(
			b, tpap.GetRemotePool, chain,
			evm_contract.FunctionInput[uint64]{ChainSelector: chain.Selector, Address: poolAddr, Args: remote.Selector},
			cldf_ops.WithForceExecute[evm_contract.FunctionInput[uint64], evm.Chain](),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to get remote pool for remote chain %d from pool %s on chain %d: %w", remote.Selector, poolAddr.Hex(), chain.Selector, err)
		}

		// The single slot is cleared when it holds the remote pool in any of its encodings.
		var matches [][]byte
		if len(poolReport.Output) > 0 {
			matches, err = evm1_0_0.MatchingRemotePools([][]byte{poolReport.Output}, remote.Selector, remote.Remote.Address)
			if err != nil {
				return nil, err
			}
		}
		if len(matches) == 0 {
			b.Logger.Warnf("skipping removal of remote pool %s for remote chain %d from pool %s on chain %d: pairing already absent", remote.Remote.Address, remote.Selector, poolAddr.Hex(), chain.Selector)
			continue
		}

		clearReport, err := cldf_ops.ExecuteOperation(
			b, tpap.SetRemotePool, chain,
			evm_contract.FunctionInput[tpap.SetRemotePoolArgs]{
				ChainSelector: chain.Selector,
				Address:       poolAddr,
				Args: tpap.SetRemotePoolArgs{
					RemoteChainSelector: remote.Selector,
					RemotePoolAddress:   []byte{},
				},
			},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to clear remote pool %s for remote chain %d on pool %s on chain %d: %w", remote.Remote.Address, remote.Selector, poolAddr.Hex(), chain.Selector, err)
		}

		writes = append(writes, clearReport.Output)
	}

	return writes, nil
}

func (p *poolOpsV150) GetCurrentRateLimits(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, ff bool) (tokensapi.OnchainRateLimits, error) {
	if ff {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("fast finality buckets are not supported on v1.5.x token pools")
	}

	outboundReport, err := cldf_ops.ExecuteOperation(b,
		tpap.GetCurrentOutboundRateLimiterState, chain,
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
		tpap.GetCurrentInboundRateLimiterState, chain,
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

func (p *poolOpsV150) Version() *semver.Version {
	return tpap.Version
}

// EffectiveMigrationRateLimits resolves the rate limits that were actually in
// force for a lane. A v1.5.0 *AndProxy pool ("proxy") forwards every transfer to
// its getPreviousPool() ("previous"), and BOTH apply their own limiter, so the
// effective limit per direction is the tighter of the two.
//
// Exists only for the v1.5.0 *AndProxy migration path; do not copy this pattern elsewhere.
func EffectiveMigrationRateLimits(
	b cldf_ops.Bundle,
	chain evm.Chain,
	proxy common.Address,
	remoteSelector uint64,
	proxyOut, proxyIn tokensapi.RateLimiterConfig,
	remoteDecimals, localDecimals uint8,
) (tokensapi.OnchainRateLimits, error) {
	unchanged := tokensapi.OnchainRateLimits{Outbound: proxyOut, Inbound: proxyIn}
	if proxy == (common.Address{}) {
		return unchanged, nil
	}

	opts := &bind.CallOpts{Context: b.GetContext()}

	previous, err := previousPool(chain, proxy, opts)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to get previous pool for proxy pool %s: %w", proxy.Hex(), err)
	}
	if previous == (common.Address{}) {
		return unchanged, nil
	}

	prevVersion, prevType, err := previousTypeAndVersion(b, chain, previous)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, err
	}

	var prevOut, prevIn tokensapi.RateLimiterConfig
	switch {
	case prevVersion.GreaterThanEqual(utils.Version_1_4_0) && prevVersion.LessThan(utils.Version_1_5_0):
		prevOut, prevIn, err = readV14Limits(chain, previous, remoteSelector, opts)
	case prevVersion.GreaterThanEqual(utils.Version_1_2_0) && prevVersion.LessThan(utils.Version_1_4_0):
		prevOut, prevIn, err = readV12Limits(chain, previous, proxy, prevType, opts)
		if err == nil && (prevOut.IsEnabled || prevIn.IsEnabled) {
			b.Logger.Warnf(
				"v1.5.0 *AndProxy pool %s lane %d: previous v1.2 pool %s applies an enabled, "+
					"per-proxy (shared) rate limit; the aggregate limit across lanes cannot be "+
					"preserved and may increase after migration",
				proxy.Hex(), remoteSelector, previous.Hex(),
			)
		}
	default:
		// NOTE: a v1.5.0 *AndProxy pool can also point at another v1.5.0 pool (e.g. a
		// v1.5.0 BnM pool). That is out of scope for now; extend here if it comes up.
		return tokensapi.OnchainRateLimits{}, fmt.Errorf(
			"unsupported previous pool %s version %s for v1.5.0 *AndProxy pool %s",
			prevType, prevVersion.String(), proxy.Hex(),
		)
	}
	if err != nil {
		return tokensapi.OnchainRateLimits{}, err
	}

	prevIn = rebasePreviousInbound(prevIn, remoteDecimals, localDecimals)

	return tokensapi.OnchainRateLimits{
		Outbound: effectiveRateLimiter(proxyOut, prevOut),
		Inbound:  effectiveRateLimiter(proxyIn, prevIn),
	}, nil
}

// rebasePreviousInbound rebases a previous pool's inbound bucket to local decimals, mirroring the
// exact guard LegacyRateLimitsForAutoMigrate already applied to proxyIn upstream (only when
// remoteDecimals != 0) so the two are compared like for like. v1.2/v1.4 EVM previous pools always
// use remote decimals, so the DoesPoolUseLocalDecimals check reduces to the remoteDecimals
// sentinel.
func rebasePreviousInbound(prevIn tokensapi.RateLimiterConfig, remoteDecimals, localDecimals uint8) tokensapi.RateLimiterConfig {
	if remoteDecimals == 0 {
		return prevIn
	}
	return tokensapi.RebaseRateLimiterConfig(prevIn, remoteDecimals, localDecimals)
}

// effectiveRateLimiter returns the constraint that actually binds: the enabled
// side wins; when both are enabled the smaller capacity/rate wins. A disabled
// limiter imposes no constraint (it is NOT equivalent to numeric zero).
func effectiveRateLimiter(proxy, previous tokensapi.RateLimiterConfig) tokensapi.RateLimiterConfig {
	switch {
	case !proxy.IsEnabled && !previous.IsEnabled:
		return tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
	case !previous.IsEnabled:
		return proxy
	case !proxy.IsEnabled:
		return previous
	default:
		return tokensapi.RateLimiterConfig{
			IsEnabled: true,
			Capacity:  minBigInt(proxy.Capacity, previous.Capacity),
			Rate:      minBigInt(proxy.Rate, previous.Rate),
		}
	}
}

// previousPool reads getPreviousPool() from the v1.5.0 *AndProxy pool at proxy. The value is
// technically mutable, so this is a raw call rather than a cached operation.
func previousPool(chain evm.Chain, proxy common.Address, opts *bind.CallOpts) (common.Address, error) {
	contract, err := tpap.NewTokenPoolAndProxyContract(proxy, chain.Client)
	if err != nil {
		return common.Address{}, err
	}
	return contract.GetPreviousPool(opts)
}

// previousTypeAndVersion resolves the type and version of the previous pool via the shared
// typeAndVersion operation. The result is immutable per pool address, so it is safe to cache.
func previousTypeAndVersion(b cldf_ops.Bundle, chain evm.Chain, previous common.Address) (*semver.Version, string, error) {
	report, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       previous,
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to get type and version of previous pool %s: %w", previous.Hex(), err)
	}
	return report.Output.Version, report.Output.Type.String(), nil
}

// readV14Limits reads the per-remote-chain outbound/inbound buckets from a v1.4 previous pool.
// Bucket state is mutable, so this is a raw call rather than a cached operation.
func readV14Limits(
	chain evm.Chain, previous common.Address, remoteSelector uint64, opts *bind.CallOpts,
) (outbound, inbound tokensapi.RateLimiterConfig, err error) {
	caller, err := v1_4_0_token_pool.NewTokenPoolCaller(previous, chain.Client)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, err
	}
	out, err := caller.GetCurrentOutboundRateLimiterState(opts, remoteSelector)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
			"failed to get outbound rate limiter state for v1.4 previous pool %s: %w", previous.Hex(), err,
		)
	}
	in, err := caller.GetCurrentInboundRateLimiterState(opts, remoteSelector)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
			"failed to get inbound rate limiter state for v1.4 previous pool %s: %w", previous.Hex(), err,
		)
	}
	return tokensapi.RateLimiterConfig{IsEnabled: out.IsEnabled, Capacity: out.Capacity, Rate: out.Rate},
		tokensapi.RateLimiterConfig{IsEnabled: in.IsEnabled, Capacity: in.Capacity, Rate: in.Rate},
		nil
}

// TODO: if v1.2 has a shared `TokenPool` base contract that BnM and LnR pools inherit
// from, then this function should be refactored such that it re-uses the shared bindings
// for both pool types similar to `readV14Limits`.
//
// readV12Limits reads the per-proxy-address outbound/inbound buckets from a v1.2 previous pool.
// v1.2 pools key rate limiter state by onRamp (outbound)/offRamp (inbound) address, shared across
// every lane the proxy pool serves (see the aggregate-bucket caveat logged by the caller). Bucket
// state is mutable, so this is a raw call rather than a cached operation.
func readV12Limits(
	chain evm.Chain, previous, proxy common.Address, prevType string, opts *bind.CallOpts,
) (outbound, inbound tokensapi.RateLimiterConfig, err error) {
	switch deployment.ContractType(prevType) {
	case utils.BurnMintTokenPool:
		caller, err := v1_2_0_burn_mint_token_pool.NewBurnMintTokenPoolCaller(previous, chain.Client)
		if err != nil {
			return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, err
		}
		out, err := caller.CurrentOnRampRateLimiterState(opts, proxy)
		if err != nil {
			return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
				"failed to get onRamp rate limiter state for v1.2 BurnMintTokenPool %s: %w", previous.Hex(), err,
			)
		}
		in, err := caller.CurrentOffRampRateLimiterState(opts, proxy)
		if err != nil {
			return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
				"failed to get offRamp rate limiter state for v1.2 BurnMintTokenPool %s: %w", previous.Hex(), err,
			)
		}
		return tokensapi.RateLimiterConfig{IsEnabled: out.IsEnabled, Capacity: out.Capacity, Rate: out.Rate},
			tokensapi.RateLimiterConfig{IsEnabled: in.IsEnabled, Capacity: in.Capacity, Rate: in.Rate},
			nil

	case utils.LockReleaseTokenPool:
		// No official v1.2.0 LockRelease gobindings exist; the v1.0.0 binding's read selectors
		// are identical, so it is reused here. Revisit if official v1.2.0 bindings become available.
		caller, err := v1_0_0_lock_release_token_pool.NewLockReleaseTokenPoolCaller(previous, chain.Client)
		if err != nil {
			return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, err
		}
		out, err := caller.CurrentOnRampRateLimiterState(opts, proxy)
		if err != nil {
			return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
				"failed to get onRamp rate limiter state for v1.2 LockReleaseTokenPool %s: %w", previous.Hex(), err,
			)
		}
		in, err := caller.CurrentOffRampRateLimiterState(opts, proxy)
		if err != nil {
			return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
				"failed to get offRamp rate limiter state for v1.2 LockReleaseTokenPool %s: %w", previous.Hex(), err,
			)
		}
		return tokensapi.RateLimiterConfig{IsEnabled: out.IsEnabled, Capacity: out.Capacity, Rate: out.Rate},
			tokensapi.RateLimiterConfig{IsEnabled: in.IsEnabled, Capacity: in.Capacity, Rate: in.Rate},
			nil

	default:
		return tokensapi.RateLimiterConfig{}, tokensapi.RateLimiterConfig{}, fmt.Errorf(
			"unsupported v1.2 previous pool type %q for pool %s", prevType, previous.Hex(),
		)
	}
}

// minBigInt returns the smaller of a and b, treating nil as zero.
func minBigInt(a, b *big.Int) *big.Int {
	if a == nil {
		a = big.NewInt(0)
	}
	if b == nil {
		b = big.NewInt(0)
	}
	if a.Cmp(b) <= 0 {
		return a
	}
	return b
}
