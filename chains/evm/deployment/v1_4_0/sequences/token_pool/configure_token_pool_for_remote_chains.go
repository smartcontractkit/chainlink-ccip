package token_pool

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evmutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/operations/token_pool"
	tpbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureTokenPoolForRemoteChainsInput struct {
	// ChainSelector identifies the chain the pool lives on. It must be carried in
	// the input (not only the executor dependency) so that two invocations with
	// otherwise identical inputs on different chains do not collide in the
	// operations report cache, which keys on the sequence input alone. Without
	// it, configuring two chains whose pool addresses are identical reuses the
	// first chain's cached report and emits ops against the wrong chain.
	ChainSelector    uint64
	TokenPoolAddress common.Address
	TokenPoolVersion *semver.Version
	RemoteChains     map[uint64]tokensapi.RemoteChainConfig[[]byte, string]
}

type ConfigureTokenPoolForRemoteChainInput struct {
	ChainSelector       uint64
	TokenPoolAddress    common.Address
	RemoteChainSelector uint64
	RemoteChainConfig   tokensapi.RemoteChainConfig[[]byte, string]
}

func (c ConfigureTokenPoolForRemoteChainInput) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	return evmutils.Wrap(
		c.RemoteChainConfig.Validate(),
		fmt.Sprintf("invalid remote chain config for remote chain selector %d", c.RemoteChainSelector),
	)
}

// ConfigureTokenPoolForRemoteChains configures a legacy v1.4.0 token pool on an
// EVM chain for cross-chain token transfers with other remote chains. It's
// capable of configuring multiple remote chains with a single invocation.
//
// Scope note: a surviving v1.4.0 pool sits behind a v1.5.0 TokenPoolAndProxy
// that serves the lane. The only reason to configure the legacy pool today is to
// replicate or tighten the rate limits already set on that proxy, which is what
// this sequence does - it allows a remote chain and sets its rate limits, and
// nothing else.
//
// This is the v1.4.0 analogue of the v1.5.0 sequence of the same name. It cannot
// share that implementation because the v1.4.0 pool ABI differs in three
// load-bearing ways:
//
//   - there is no getTokenDecimals(); local decimals are read from the token's
//     ERC20 decimals()
//   - the pool stores NO remote token and NO remote pool - the fronting proxy
//     holds those - so applyChainUpdates carries only the selector and rate
//     limits, and there is no setRemotePool step
//   - applyChainUpdates takes one ChainUpdate[] and expresses removals as
//     allowed=false entries, rather than v1.5.1's (selectorsToRemove, chainsToAdd)
//
// Lane removal is deliberately not implemented: dropping a chain here means
// applyChainUpdates(allowed=false), which deletes the whole remote chain config,
// and retiring a legacy pool is a migration concern rather than a configuration
// one.
//
// Pool-type agnostic: every call below is on the shared TokenPool base surface,
// which BurnMintTokenPool and LockReleaseTokenPool share with byte-identical
// signatures. Type-specific state (lock-release liquidity and rebalancer) is not
// touched here.
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"token-pool:configure-token-pool-for-remote-chains",
	tp.Version,
	"Configure a v1.4.0 token pool on an EVM chain for cross-chain transfers",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainsInput) (sequences.OnChainOutput, error) {
		tokenPool, err := tpbindings.NewTokenPool(input.TokenPoolAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token pool contract: %w", err)
		}

		batchOps := make([]mcms_types.BatchOperation, 0)
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			report, err := cldf_ops.ExecuteSequence(b,
				ConfigureTokenPoolForRemoteChain,
				chain,
				ConfigureTokenPoolForRemoteChainInput{
					ChainSelector:       input.ChainSelector,
					TokenPoolAddress:    tokenPool.Address(),
					RemoteChainSelector: remoteChainSelector,
					RemoteChainConfig:   remoteChainConfig,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chain %d: %w", remoteChainSelector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
		}

		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	})

// ConfigureTokenPoolForRemoteChain is a helper sequence that performs the logic
// for configuring a token pool for a SINGLE remote chain. The sequence allows
// the upper level ConfigureTokenPoolForRemoteChains sequence to handle multiple
// remote chains.
var ConfigureTokenPoolForRemoteChain = cldf_ops.NewSequence(
	"token-pool:configure-token-pool-for-remote-chain",
	tp.Version,
	"Configures a v1.4.0 token pool on an EVM chain for transfers with other chains",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (sequences.OnChainOutput, error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, err
		}

		tokenPool, err := tpbindings.NewTokenPool(input.TokenPoolAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token pool contract: %w", err)
		}
		supportedChains, err := tokenPool.GetSupportedChains(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains: %w", err)
		}
		chainSupported := slices.Contains(supportedChains, input.RemoteChainSelector)

		localTokenAddr, err := tokenPool.GetToken(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token from token pool: %w", err)
		}
		decimalsReport, err := cldf_ops.ExecuteOperation(b, erc20.GetDecimals, chain, contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       localTokenAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get decimals for token %s: %w", localTokenAddr.Hex(), err)
		}
		localDecimals := decimalsReport.Output

		tvReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args:          struct{}{},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get type and version of token pool: %w", err)
		}

		outboundRL, outboundOk := input.RemoteChainConfig.GetOutboundRateLimitBuckets().DefaultBucket()
		inboundRL, inboundOk := input.RemoteChainConfig.GetInboundRateLimitBuckets().DefaultBucket()

		var inputORL, inputIRL tokensapi.RateLimiterConfig
		switch {
		case outboundOk && inboundOk:
			inputORL, inputIRL = tokensapi.GenerateTPRLConfigs(
				outboundRL.RateLimit,
				inboundRL.RateLimit,
				localDecimals,
				input.RemoteChainConfig.RemoteDecimals,
				chain.Family(),
				tvReport.Output.Version,
				tvReport.Output.Type.String(),
			)

		case !outboundOk && !inboundOk:
			if chainSupported {
				onchainOutboundBucket, err := tokenPool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound rate limiter state for remote chain %d: %w", input.RemoteChainSelector, err)
				}
				onchainInboundBucket, err := tokenPool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound rate limiter state for remote chain %d: %w", input.RemoteChainSelector, err)
				}

				inputORL = tokensapi.RateLimiterConfig{
					IsEnabled: onchainOutboundBucket.IsEnabled,
					Capacity:  onchainOutboundBucket.Capacity,
					Rate:      onchainOutboundBucket.Rate,
				}
				inputIRL = tokensapi.RateLimiterConfig{
					IsEnabled: onchainInboundBucket.IsEnabled,
					Capacity:  onchainInboundBucket.Capacity,
					Rate:      onchainInboundBucket.Rate,
				}
			} else {
				inputORL = tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
				inputIRL = tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf(
				"default outbound and inbound rate limits must both be specified together or both omitted for remote chain %d",
				input.RemoteChainSelector,
			)
		}

		// NOTE: v1.4.0 pools share v1.5.x's rate limiter validation: an enabled
		// bucket with both rate and capacity at zero is rejected
		// (InvalidRateLimitRate), and a disabled bucket with a non-zero rate or
		// capacity is rejected (DisabledNonZeroRateLimit).
		if inputORL.IsEnabled && inputORL.Capacity.Cmp(big.NewInt(0)) == 0 && inputORL.Rate.Cmp(big.NewInt(0)) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("outbound rate limiter config is enabled but rate and capacity are both zero")
		}
		if inputIRL.IsEnabled && inputIRL.Capacity.Cmp(big.NewInt(0)) == 0 && inputIRL.Rate.Cmp(big.NewInt(0)) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("inbound rate limiter config is enabled but rate and capacity are both zero")
		}

		reportWrites := []contract.WriteOutput{}
		if chainSupported {
			onchainORL, err := tokenPool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound rate limiter state: %w", err)
			}
			onchainIRL, err := tokenPool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound rate limiter state: %w", err)
			}

			isOutboundEqual := inputORL.IsEnabled == onchainORL.IsEnabled &&
				inputORL.Capacity.Cmp(onchainORL.Capacity) == 0 &&
				inputORL.Rate.Cmp(onchainORL.Rate) == 0
			isInboundEqual := inputIRL.IsEnabled == onchainIRL.IsEnabled &&
				inputIRL.Capacity.Cmp(onchainIRL.Capacity) == 0 &&
				inputIRL.Rate.Cmp(onchainIRL.Rate) == 0

			if isOutboundEqual && isInboundEqual {
				return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{}}, nil
			}

			report, err := cldf_ops.ExecuteOperation(b, tp.SetChainRateLimiterConfig, chain, contract.FunctionInput[tp.SetChainRateLimiterConfigArgs]{
				ChainSelector: chain.Selector,
				Address:       input.TokenPoolAddress,
				Args: tp.SetChainRateLimiterConfigArgs{
					RemoteChainSelector: input.RemoteChainSelector,
					OutboundConfig:      tp.Config{IsEnabled: inputORL.IsEnabled, Capacity: inputORL.Capacity, Rate: inputORL.Rate},
					InboundConfig:       tp.Config{IsEnabled: inputIRL.IsEnabled, Capacity: inputIRL.Capacity, Rate: inputIRL.Rate},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
			}
			reportWrites = append(reportWrites, report.Output)
		} else {
			report, err := cldf_ops.ExecuteOperation(b, tp.ApplyChainUpdates, chain, contract.FunctionInput[[]tp.ChainUpdate]{
				ChainSelector: chain.Selector,
				Address:       input.TokenPoolAddress,
				Args: []tp.ChainUpdate{{
					RemoteChainSelector:       input.RemoteChainSelector,
					Allowed:                   true,
					OutboundRateLimiterConfig: tp.Config{IsEnabled: inputORL.IsEnabled, Capacity: inputORL.Capacity, Rate: inputORL.Rate},
					InboundRateLimiterConfig:  tp.Config{IsEnabled: inputIRL.IsEnabled, Capacity: inputIRL.Capacity, Rate: inputIRL.Rate},
				}},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
			}
			reportWrites = append(reportWrites, report.Output)
		}

		if len(reportWrites) == 0 {
			return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{}}, nil
		}

		batchOp, err := contract.NewBatchOperationFromWrites(reportWrites)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})
