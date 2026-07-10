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
	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	_ tokensapi.TokenPoolMigrator = &TokenAdapter{}
	_ tokensapi.TokenAdapter      = &TokenAdapter{}
)

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
		},
	)
}

func (t *TokenAdapter) GetSupportedChains(e deployment.Environment, chainSelector uint64, poolAddr []byte) ([]uint64, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := evmops.ExecuteRead(
		e.OperationsBundle,
		evmChain,
		common.BytesToAddress(poolAddr),
		seqV1_6_1.NewTokenPool,
		tpOps.NewReadGetSupportedChains,
		struct{}{},
		cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain](),
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

	report, err := evmops.ExecuteRead(
		e.OperationsBundle,
		evmChain,
		common.BytesToAddress(poolAddr),
		seqV1_6_1.NewTokenPool,
		tpOps.NewReadGetRemoteToken,
		remoteSelector,
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

func (t *TokenAdapter) GetRemotePools(e deployment.Environment, chainSelector uint64, poolAddr []byte, remoteSelector uint64) ([][]byte, error) {
	evmChain, ok := e.BlockChains.EVMChains()[chainSelector]
	if !ok {
		return nil, fmt.Errorf("chain with selector %d not found", chainSelector)
	}

	report, err := evmops.ExecuteRead(
		e.OperationsBundle,
		evmChain,
		common.BytesToAddress(poolAddr),
		seqV1_6_1.NewTokenPool,
		tpOps.NewReadGetRemotePools,
		remoteSelector,
		cldf_ops.WithForceExecute[contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get remote pools for chain %d from pool %s: %w", remoteSelector, common.BytesToAddress(poolAddr).Hex(), err)
	}

	return report.Output, nil
}

// poolOpsV161 implements PoolOps using v1.6.1 bindings.
type poolOpsV161 struct{}

func (p *poolOpsV161) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := evmops.ExecuteRead(b, chain, poolAddr, seqV1_6_1.NewTokenPool, tpOps.NewReadGetToken, struct{}{})
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v1.6.1: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV161) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	res, err := evmops.ExecuteRead(b, chain, poolAddr, seqV1_6_1.NewTokenPool, tpOps.NewReadGetTokenDecimals, struct{}{})
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.6.1: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV161) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := gobindings.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate v1.6.1 token pool contract at %s: %w", poolAddr.Hex(), err)
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

func (p *poolOpsV161) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]contract.WriteOutput, error) {
	bucket, ok := input.GetBucketForFinality(false)
	if !ok {
		b.Logger.Warnf("skipping rate limiter config for token pool (%s) on chain %d since no default bucket was provided", datastore_utils.SprintRef(input.TokenPoolRef), input.ChainSelector)
		return nil, nil
	}

	report, err := evmops.ExecuteWrite(b, chain, poolAddr, seqV1_6_1.NewTokenPool, tpOps.NewWriteSetChainRateLimiterConfig, tpOps.SetChainRateLimiterConfigArgs{
		OutboundConfig: gobindings.RateLimiterConfig{
			IsEnabled: bucket.OutboundRateLimiterConfig.IsEnabled,
			Capacity:  bucket.OutboundRateLimiterConfig.Capacity,
			Rate:      bucket.OutboundRateLimiterConfig.Rate,
		},
		InboundConfig: gobindings.RateLimiterConfig{
			IsEnabled: bucket.InboundRateLimiterConfig.IsEnabled,
			Capacity:  bucket.InboundRateLimiterConfig.Capacity,
			Rate:      bucket.InboundRateLimiterConfig.Rate,
		},
		RemoteChainSelector: input.RemoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("SetChainRateLimiterConfig v1.6.1: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (p *poolOpsV161) SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) ([]contract.WriteOutput, error) {
	report, err := evmops.ExecuteWrite(b, chain, poolAddr, seqV1_6_1.NewTokenPool, tpOps.NewWriteSetRateLimitAdmin, newAdmin)
	if err != nil {
		return nil, fmt.Errorf("SetRateLimitAdmin v1.6.1: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (p *poolOpsV161) GetCurrentRateLimits(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, ff bool) (tokensapi.OnchainRateLimits, error) {
	if ff {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("fast finality buckets are not supported on v1.6.x token pools")
	}

	outboundReport, err := evmops.ExecuteRead(
		b,
		chain,
		poolAddr,
		seqV1_6_1.NewTokenPool,
		tpOps.NewReadGetCurrentOutboundRateLimiterState,
		remoteSelector,
		cldf_ops.WithForceExecute[contract.FunctionInput[uint64], evm.Chain](),
	)
	if err != nil {
		return tokensapi.OnchainRateLimits{}, fmt.Errorf("failed to get outbound rate limiter state for remote chain %d: %w", remoteSelector, err)
	}
	inboundReport, err := evmops.ExecuteRead(
		b,
		chain,
		poolAddr,
		seqV1_6_1.NewTokenPool,
		tpOps.NewReadGetCurrentInboundRateLimiterState,
		remoteSelector,
		cldf_ops.WithForceExecute[contract.FunctionInput[uint64], evm.Chain](),
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

func (p *poolOpsV161) Version() *semver.Version {
	return tpOps.Version
}
