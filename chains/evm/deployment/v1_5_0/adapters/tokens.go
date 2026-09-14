package adapters

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences/token_pool"
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
	// RateLimitReaderAdapter is load-bearing, not incidental: ConfigureTokensForTransfers hard
	// errors during auto-migrate discovery if an adapter implements TokenPoolMigrator without it.
	_ tokensapi.RateLimitReaderAdapter = &TokenAdapter{}
)

// TokenAdapter handles EVM token pools at version 1.5.0.
// It embeds EVMPoolAdapter for shared datastore/TAR/BnM logic and
// overrides only ConfigureTokenForTransfersSequence which inlines
// the v1.5.0-specific configure + register flow.
//
// Scope: BurnMintTokenPoolAndProxy only. The other v1.5.0 pool contracts
// (lock-release and rebasing *AndProxy variants) have generated bindings but no
// deployed footprint, and the deploy sequence rejects them.
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
		bmtpap.Version,
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
					TokenPoolAddress: tpAddr,
					TokenPoolVersion: bmtpap.Version,
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

func (t *TokenAdapter) GetSupportedChains(e deployment.Environment, chainSelector uint64, poolAddr []byte) ([]uint64, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		bmtpap.GetSupportedChains, evmChain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr)},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get supported chains from pool %s on chain %d: %w", common.BytesToAddress(poolAddr).Hex(), chainSelector, err)
	}

	return report.Output, nil
}

func (t *TokenAdapter) GetRemoteToken(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		bmtpap.GetRemoteToken, evmChain,
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
func (t *TokenAdapter) GetRemotePools(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([][]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		bmtpap.GetRemotePool, evmChain,
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

// poolOpsV150 implements PoolOps using v1.5.0 BurnMintTokenPoolAndProxy bindings.
type poolOpsV150 struct{}

func (p *poolOpsV150) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(
		b,
		bmtpap.GetToken, chain,
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
	pool, err := bmtpap.NewBurnMintTokenPoolAndProxyContract(poolAddr, chain.Client)
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
		bmtpap.SetChainRateLimiterConfig, chain,
		evm_contract.FunctionInput[bmtpap.SetChainRateLimiterConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: bmtpap.SetChainRateLimiterConfigArgs{
				OutboundConfig: bmtpap.Config{
					IsEnabled: outbound.IsEnabled,
					Capacity:  outbound.Capacity,
					Rate:      outbound.Rate,
				},
				InboundConfig: bmtpap.Config{
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

	pool, err := bmtpap.NewBurnMintTokenPoolAndProxyContract(poolAddr, chain.Client)
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
				bmtpap.SetRouter, chain,
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
				bmtpap.SetRateLimitAdmin, chain,
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

// RemoveRemotePools is not supported on v1.5.0 pools. The contract has no removeRemotePool: the
// only way to drop a remote pool entry is applyChainUpdates with allowed=false, which deletes the
// ENTIRE remote chain config (remote token and both rate limiters), not just the pool entry.
// Doing that silently under a "remove remote pools" pipeline would tear down more than the caller
// asked for, so this errors instead. Use ConfigureTokenForTransfers to retarget the lane, or drop
// the chain deliberately.
func (p *poolOpsV150) RemoveRemotePools(_ cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, _ []tokensapi.RemotePoolToRemove) ([]evm_contract.WriteOutput, error) {
	return nil, fmt.Errorf(
		"removing individual remote pools is not supported on v1.5.0 token pools (pool %s on chain %d): "+
			"the contract has no removeRemotePool, and applyChainUpdates(allowed=false) would remove the whole remote chain config",
		poolAddr.Hex(), chain.Selector,
	)
}

func (p *poolOpsV150) GetCurrentRateLimits(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, ff bool) (tokensapi.OnchainRateLimits, error) {
	if ff {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("fast finality buckets are not supported on v1.5.x token pools")
	}

	outboundReport, err := cldf_ops.ExecuteOperation(b,
		bmtpap.GetCurrentOutboundRateLimiterState, chain,
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
		bmtpap.GetCurrentInboundRateLimiterState, chain,
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
	return bmtpap.Version
}
