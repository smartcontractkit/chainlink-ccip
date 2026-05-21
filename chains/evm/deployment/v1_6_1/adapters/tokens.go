package adapters

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	seqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	tpOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	seqV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ tokensapi.TokenAdapter = &TokenAdapter{}

// TokenAdapter handles EVM token pools at version 1.6.1.
// It embeds EVMPoolAdapter for shared datastore/TAR/BnM logic and
// overrides only ConfigureTokenForTransfersSequence which delegates
// to the pre-built v1.6.1 sequence.
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a TokenAdapter with pre-wired PoolOps and
// the deploy-token-pool sequence.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops:                &poolOpsV161{},
			DeployTokenPoolSeq: seqV1_6_0.DeployTokenPool,
		},
	}
}

// ConfigureTokenForTransfersSequence wraps the v1.6.1 pre-built sequence,
// resolving the TAR address from the datastore when the caller leaves
// RegistryAddress empty (the top-level changeset relies on adapters for this).
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.6.1-adapter:configure-token-for-transfers",
		tpOps.Version,
		"Configure a v1.6.1 token pool for cross-chain transfers on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			if input.RegistryAddress == "" {
				tarAddr, err := t.EVMTokenBase.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve TAR address for chain %d: %w", input.ChainSelector, err)
				}
				input.RegistryAddress = tarAddr.Hex()
			}

			report, err := cldf_ops.ExecuteSequence(b, seqV1_6_1.ConfigureTokenForTransfers, chains, input)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			return report.Output, nil
		})
}

// poolOpsV161 implements PoolOps using v1.6.1 bindings.
type poolOpsV161 struct{}

func (p *poolOpsV161) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		tpOps.GetToken, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v1.6.1: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV161) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		tpOps.GetTokenDecimals, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.6.1: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV161) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := token_pool.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v1.6.1 contract: %w", err)
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

func (p *poolOpsV161) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error) {
	bucket, ok := input.GetBucketForFinality(false)
	if !ok {
		b.Logger.Warnf("skipping rate limiter config for token pool (%s) on chain %d since no default bucket was provided", datastore_utils.SprintRef(input.TokenPoolRef), input.ChainSelector)
		return nil, nil
	}

	report, err := cldf_ops.ExecuteOperation(b,
		tpOps.SetChainRateLimiterConfig, chain,
		evm_contract.FunctionInput[tpOps.SetChainRateLimiterConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: tpOps.SetChainRateLimiterConfigArgs{
				OutboundConfig: tpOps.Config{
					IsEnabled: bucket.OutboundRateLimiterConfig.IsEnabled,
					Capacity:  bucket.OutboundRateLimiterConfig.Capacity,
					Rate:      bucket.OutboundRateLimiterConfig.Rate,
				},
				InboundConfig: tpOps.Config{
					IsEnabled: bucket.InboundRateLimiterConfig.IsEnabled,
					Capacity:  bucket.InboundRateLimiterConfig.Capacity,
					Rate:      bucket.InboundRateLimiterConfig.Rate,
				},
				RemoteChainSelector: input.RemoteChainSelector,
			},
		})
	if err != nil {
		return nil, fmt.Errorf("SetChainRateLimiterConfig v1.6.1: %w", err)
	}
	return []evm_contract.WriteOutput{report.Output}, nil
}

func (p *poolOpsV161) SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) (evm_contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b,
		tpOps.SetRateLimitAdmin, chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args:          newAdmin,
		})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("SetRateLimitAdmin v1.6.1: %w", err)
	}
	return report.Output, nil
}

func (p *poolOpsV161) GetCurrentInboundRateLimit(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, _ bool) (tokensapi.RateLimiterConfig, error) {
	// Call the contract binding directly rather than cldf_ops Read: the framework caches read
	// reports by input hash, and earlier sequences in the same Apply run may have read this
	// same lane while it was still uninitialized — caching that stale result.
	tp, err := token_pool.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, fmt.Errorf("failed to instantiate v1.6.1 token pool contract at %s: %w", poolAddr.Hex(), err)
	}
	bucket, err := tp.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, remoteSelector)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, fmt.Errorf("failed to get inbound rate limiter state for remote chain %d: %w", remoteSelector, err)
	}
	return tokensapi.RateLimiterConfig{
		IsEnabled: bucket.IsEnabled,
		Capacity:  bucket.Capacity,
		Rate:      bucket.Rate,
	}, nil
}

func (p *poolOpsV161) Version() *semver.Version {
	return tpOps.Version
}
