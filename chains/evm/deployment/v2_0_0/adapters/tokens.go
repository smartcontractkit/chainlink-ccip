package adapters

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	evm_tokens "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	_ tokens.TokenPoolMigrator = &TokenAdapter{}
	_ tokens.TokenFeeAdapter   = &TokenAdapter{}
	_ tokens.TokenAdapter      = &TokenAdapter{}
)

// TokenAdapter handles EVM token pools at version 2.0.0.
// It embeds EVMPoolAdapter for shared methods (DeriveTokenAddress,
// ManualRegistration) and overrides the methods that have genuinely
// different v2.0.0 logic (SetTokenPoolRateLimits with batched default
// + fast-finality TPRL buckets, ConfigureTokenForTransfersSequence with
// its own sequences).
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a v2.0.0 TokenAdapter with pre-wired PoolOps and
// the deploy-token-pool sequence.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops:                &poolOpsV200{},
			DeployTokenPoolSeq: evm_tokens.DeployTokenPool,
		},
	}
}

func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.ConfigureTokenForTransfers
}

func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokens.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.MigrateLockReleasePoolLiquidity
}

func (t *TokenAdapter) SetAllowedFinalityConfig(e *deployment.Environment) *cldf_ops.Sequence[tokens.SetAllowedFinalityConfigSequenceInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.SetAllowedFinalityConfigForTokenPools
}

func (t *TokenAdapter) SetTokenTransferFee(e *deployment.Environment) *cldf_ops.Sequence[tokens.SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, chain.BlockChains] {
	return evm_tokens.SetTokenTransferFeeConfigForTokenPools
}

func (t *TokenAdapter) GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) tokens.TokenTransferFeeConfig {
	return tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(src, dst)
}

func (t *TokenAdapter) GetOnchainTokenTransferFeeConfig(e deployment.Environment, poolAddress string, src uint64, dst uint64) (tokens.TokenTransferFeeConfig, error) {
	evmChain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return tokens.TokenTransferFeeConfig{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	args := token_pool.GetTokenTransferFeeConfigArgs{DestChainSelector: dst}
	if !common.IsHexAddress(poolAddress) {
		return tokens.TokenTransferFeeConfig{}, fmt.Errorf("invalid pool address: %s", poolAddress)
	}

	addr := common.HexToAddress(poolAddress)
	if addr == (common.Address{}) {
		return tokens.TokenTransferFeeConfig{}, errors.New("pool address cannot be the zero address")
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		token_pool.GetTokenTransferFeeConfig, evmChain,
		contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{ChainSelector: src, Address: addr, Args: args},
		cldf_ops.WithForceExecute[contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs], evm.Chain](),
	)
	if err != nil {
		return tokens.TokenTransferFeeConfig{}, fmt.Errorf("failed to get on-chain token transfer fee config for pool %s on chain selector %d for dest chain selector %d: %w", poolAddress, src, dst, err)
	}

	return tokens.TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: report.Output.FinalityTransferFeeBps,
		CustomFinalityTransferFeeBps:  report.Output.FastFinalityTransferFeeBps,
		DefaultFinalityFeeUSDCents:    report.Output.FinalityFeeUSDCents,
		CustomFinalityFeeUSDCents:     report.Output.FastFinalityFeeUSDCents,
		DestBytesOverhead:             report.Output.DestBytesOverhead,
		DestGasOverhead:               report.Output.DestGasOverhead,
		IsEnabled:                     report.Output.IsEnabled,
	}, nil
}

// GetSupportedChains returns the remote chain selectors the pool at poolAddr is configured for.
func (t *TokenAdapter) GetSupportedChains(e deployment.Environment, chainSelector uint64, poolAddr []byte) ([]uint64, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		token_pool.GetSupportedChains, evmChain,
		contract.FunctionInput[struct{}]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr)},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get supported chains from pool %s on chain %d: %w", common.BytesToAddress(poolAddr).Hex(), chainSelector, err)
	}

	return report.Output, nil
}

// GetRemoteToken returns the remote token (raw bytes) the pool at poolAddr uses for remoteSelector.
func (t *TokenAdapter) GetRemoteToken(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		token_pool.GetRemoteToken, evmChain,
		contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr), Args: remoteSelector},
		cldf_ops.WithForceExecute[contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote token for chain %d from pool %s: %w", remoteSelector, common.BytesToAddress(poolAddr).Hex(), err)
	}

	if len(report.Output) == 0 {
		return nil, fmt.Errorf("pool %s has no remote token registered for chain %d", common.BytesToAddress(poolAddr).Hex(), remoteSelector)
	}

	return report.Output, nil
}

// GetRemotePools returns the remote pools (raw bytes) the pool at poolAddr is linked to for remoteSelector.
func (t *TokenAdapter) GetRemotePools(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([][]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		token_pool.GetRemotePools, evmChain,
		contract.FunctionInput[uint64]{ChainSelector: chainSelector, Address: common.BytesToAddress(poolAddr), Args: remoteSelector},
		cldf_ops.WithForceExecute[contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote pools for chain %d from pool %s: %w", remoteSelector, common.BytesToAddress(poolAddr).Hex(), err)
	}

	return report.Output, nil
}

// poolOpsV200 implements PoolOps using v2.0.0 bindings.
type poolOpsV200 struct{}

func (p *poolOpsV200) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(
		b,
		token_pool.GetToken, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v2.0.0: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV200) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	res, err := cldf_ops.ExecuteOperation(
		b,
		token_pool.GetTokenDecimals, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v2.0.0: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV200) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (common.Address, common.Address, error) {
	pool, err := token_pool.NewTokenPoolContract(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v2.0.0 contract at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	owner, err := pool.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get owner of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	cfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to get dynamic config of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	return owner, cfg.RateLimitAdmin, nil
}

func (p *poolOpsV200) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokens.TPRLRemotes) ([]contract.WriteOutput, error) {
	var writes []contract.WriteOutput
	if !input.AllowedFinalityConfig.IsZero() {
		currentFinalityConfig, err := cldf_ops.ExecuteOperation(b, token_pool.GetAllowedFinalityConfig, chain, contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          struct{}{},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get allowed finality config for token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
		}
		if input.AllowedFinalityConfig.Raw() != currentFinalityConfig.Output {
			setFinalityReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetAllowedFinalityConfig, chain, contract.FunctionInput[[4]byte]{
				ChainSelector: chain.Selector,
				Address:       poolAddr,
				Args:          input.AllowedFinalityConfig.Raw(),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to set allowed finality config on token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
			}
			writes = append(writes, setFinalityReport.Output)
		}
	}

	args := make([]token_pool.RateLimitConfigArgs, 0, len(input.RateLimitBuckets))
	for _, bucket := range input.RateLimitBuckets {
		args = append(args, token_pool.RateLimitConfigArgs{
			RemoteChainSelector: input.RemoteChainSelector,
			FastFinality:        bucket.FastFinality,
			OutboundRateLimiterConfig: token_pool.Config{
				IsEnabled: bucket.OutboundRateLimiterConfig.IsEnabled,
				Capacity:  bucket.OutboundRateLimiterConfig.Capacity,
				Rate:      bucket.OutboundRateLimiterConfig.Rate,
			},
			InboundRateLimiterConfig: token_pool.Config{
				IsEnabled: bucket.InboundRateLimiterConfig.IsEnabled,
				Capacity:  bucket.InboundRateLimiterConfig.Capacity,
				Rate:      bucket.InboundRateLimiterConfig.Rate,
			},
		})
	}

	if len(args) > 0 {
		rateLimitsReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, chain, contract.FunctionInput[[]token_pool.RateLimitConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          args,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to set rate limit config on pool %s: %w", poolAddr, err)
		}
		writes = append(writes, rateLimitsReport.Output)
	}

	if len(writes) == 0 {
		return nil, nil
	}

	return writes, nil
}

func (p *poolOpsV200) SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) ([]contract.WriteOutput, error) {
	pool, err := token_pool.NewTokenPoolContract(poolAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate token pool v2.0.0 contract at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	cfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: b.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("failed to get dynamic config of token pool at %s on chain %d: %w", poolAddr.Hex(), chain.Selector, err)
	}
	if newAdmin == cfg.RateLimitAdmin {
		b.Logger.Info("Rate limit admin is already set to the desired address; no update needed")
		return nil, nil
	}

	report, err := cldf_ops.ExecuteOperation(
		b,
		token_pool.SetDynamicConfig, chain,
		contract.FunctionInput[token_pool.SetDynamicConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: token_pool.SetDynamicConfigArgs{
				RateLimitAdmin: newAdmin,
				FeeAdmin:       cfg.FeeAdmin,
				Router:         cfg.Router,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("SetDynamicConfig v2.0.0: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func (p *poolOpsV200) GetCurrentRateLimits(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, ff bool) (tokens.OnchainRateLimits, error) {
	report, err := cldf_ops.ExecuteOperation(
		b,
		token_pool.GetCurrentRateLimiterState, chain,
		contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: token_pool.GetCurrentRateLimiterStateArgs{
				RemoteChainSelector: remoteSelector,
				FastFinality:        ff,
			},
		},
		cldf_ops.WithForceExecute[contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs], evm.Chain](),
	)
	if err != nil {
		return tokens.OnchainRateLimits{}, fmt.Errorf("failed to get rate limiter state for remote chain %d (fastFinality=%t): %w", remoteSelector, ff, err)
	}
	return tokens.OnchainRateLimits{
		Outbound: tokens.RateLimiterConfig{
			IsEnabled: report.Output.OutboundRateLimiterState.IsEnabled,
			Capacity:  report.Output.OutboundRateLimiterState.Capacity,
			Rate:      report.Output.OutboundRateLimiterState.Rate,
		},
		Inbound: tokens.RateLimiterConfig{
			IsEnabled: report.Output.InboundRateLimiterState.IsEnabled,
			Capacity:  report.Output.InboundRateLimiterState.Capacity,
			Rate:      report.Output.InboundRateLimiterState.Rate,
		},
	}, nil
}

func (p *poolOpsV200) Version() *semver.Version {
	return cciputils.Version_2_0_0
}
