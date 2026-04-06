package adapters

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tarseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/token_pool"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences/token_pool"
	v1_5_1_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
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
			DeployTokenPoolSeq: v1_5_1_seq.DeployTokenPool,
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

			tarAddress, err := evm1_0_0.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector, &t.EVMTokenBase)
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

func (p *poolOpsV151) GetTokenDecimals(ctx context.Context, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	pool, err := token_pool.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to instantiate token pool v1.5.1 contract: %w", err)
	}
	return pool.GetTokenDecimals(&bind.CallOpts{Context: ctx})
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

func (p *poolOpsV151) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteChainSelector uint64, outbound, inbound tokensapi.RateLimiterConfig) (evm_contract.WriteOutput, error) {
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
				RemoteChainSelector: remoteChainSelector,
			},
		})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("SetChainRateLimiterConfig v1.5.1: %w", err)
	}
	return report.Output, nil
}

func (p *poolOpsV151) Version() *semver.Version {
	return tpOps.Version
}
