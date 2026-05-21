package adapters

import (
	"context"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/token_pool"
	seqV1_5_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var _ tokensapi.TokenAdapter = &TokenAdapter{}

// TokenAdapter handles EVM token pools at version 1.5.1.
// It embeds EVMPoolAdapter for shared datastore/TAR/BnM logic and
// overrides only ConfigureTokenForTransfersSequence which inlines
// the v1.5.1-specific configure + register flow.
type TokenAdapter struct {
	evm1_0_0.EVMPoolAdapter
}

// NewTokenAdapter constructs a TokenAdapter with pre-wired PoolOps and
// the deploy-token-pool sequence.
func NewTokenAdapter() *TokenAdapter {
	return &TokenAdapter{
		EVMPoolAdapter: evm1_0_0.EVMPoolAdapter{
			Ops:                &poolOpsV151{},
			DeployTokenPoolSeq: seqV1_5_1.DeployTokenPool,
		},
	}
}

func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.1-adapter:configure-token-for-transfers",
		tpOps.Version,
		"Configure a v1.5.1 token pool for cross-chain transfers on an EVM chain",
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

			configureReport, err := cldf_ops.ExecuteSequence(b,
				tpSeq.ConfigureTokenPoolForRemoteChains, chain,
				tpSeq.ConfigureTokenPoolForRemoteChainsInput{
					TokenPoolAddress: tpAddr,
					TokenPoolVersion: tpOps.Version,
					RemoteChains:     input.RemoteChains,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for transfers on chain %d: %w", input.ChainSelector, err)
			}
			result.Addresses = append(result.Addresses, configureReport.Output.Addresses...)
			result.BatchOps = append(result.BatchOps, configureReport.Output.BatchOps...)

			registerReport, err := cldf_ops.ExecuteSequence(b,
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
		})
}

// poolOpsV151 implements PoolOps using v1.5.1 bindings.
type poolOpsV151 struct{}

func (p *poolOpsV151) GetToken(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (common.Address, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		tpOps.GetToken, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return common.Address{}, fmt.Errorf("GetToken v1.5.1: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV151) GetTokenDecimals(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	res, err := cldf_ops.ExecuteOperation(b,
		tpOps.GetTokenDecimals, chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("GetTokenDecimals v1.5.1: %w", err)
	}
	return res.Output, nil
}

func (p *poolOpsV151) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := token_pool.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
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

func (p *poolOpsV151) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, input tokensapi.TPRLRemotes) ([]evm_contract.WriteOutput, error) {
	bucket, ok := input.GetBucketForFinality(false)
	if !ok {
		b.Logger.Warnf("skipping rate limiter config for token pool (%s) on chain %d since no default bucket was provided", datastore_utils.SprintRef(input.TokenPoolRef), input.ChainSelector)
		return nil, nil
	}

	// NOTE: EVM v1.5.1 pools have slightly different rate limit validation rules than v1.6.1+ pools.
	// See: https://basescan.org/address/0x5192Bd10f28A0206211CcBB66671118f85c2E539#code#F12#L119
	outbound, inbound := bucket.OutboundRateLimiterConfig, bucket.InboundRateLimiterConfig
	if outbound.IsEnabled && outbound.Capacity.Cmp(big.NewInt(0)) == 0 && outbound.Rate.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("outbound rate limiter config is enabled but rate and capacity are both zero")
	}
	if inbound.IsEnabled && inbound.Capacity.Cmp(big.NewInt(0)) == 0 && inbound.Rate.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("inbound rate limiter config is enabled but rate and capacity are both zero")
	}

	report, err := cldf_ops.ExecuteOperation(b,
		tpOps.SetChainRateLimiterConfig, chain,
		evm_contract.FunctionInput[tpOps.SetChainRateLimiterConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: tpOps.SetChainRateLimiterConfigArgs{
				OutboundRateLimitConfig: token_pool.RateLimiterConfig{
					IsEnabled: outbound.IsEnabled,
					Capacity:  outbound.Capacity,
					Rate:      outbound.Rate,
				},
				InboundRateLimitConfig: token_pool.RateLimiterConfig{
					IsEnabled: inbound.IsEnabled,
					Capacity:  inbound.Capacity,
					Rate:      inbound.Rate,
				},
				RemoteChainSelector: input.RemoteChainSelector,
			},
		})
	if err != nil {
		return nil, fmt.Errorf("SetChainRateLimiterConfig v1.5.1: %w", err)
	}
	return []evm_contract.WriteOutput{report.Output}, nil
}

func (p *poolOpsV151) SetRateLimitAdmin(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, newAdmin common.Address) (evm_contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b,
		tpOps.SetRateLimitAdmin, chain,
		evm_contract.FunctionInput[tpOps.SetRateLimitAdminArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: tpOps.SetRateLimitAdminArgs{
				NewAdmin: newAdmin,
			},
		})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("SetRateLimitAdmin v1.5.1: %w", err)
	}
	return report.Output, nil
}

func (p *poolOpsV151) GetCurrentInboundRateLimit(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteSelector uint64, _ bool) (tokensapi.RateLimiterConfig, error) {
	tp, err := token_pool.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return tokensapi.RateLimiterConfig{}, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
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

func (p *poolOpsV151) Version() *semver.Version {
	return tpOps.Version
}
