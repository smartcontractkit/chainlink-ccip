package adapters

import (
	"context"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	v1_6_0_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	evm_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences"
	tpOpsV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	tpV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
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
			DeployTokenPoolSeq: v1_6_0_seq.DeployTokenPool,
		},
	}
}

// ConfigureTokenForTransfersSequence wraps the v1.6.1 pre-built sequence,
// resolving the TAR address from the datastore when the caller leaves
// RegistryAddress empty (the top-level changeset relies on adapters for this).
func (t *TokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokensapi.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.6.1-adapter:configure-token-for-transfers",
		tpOpsV1_6_1.Version,
		"Configure a v1.6.1 token pool for cross-chain transfers on an EVM chain",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			if input.RegistryAddress == "" {
				tarAddr, err := evm1_0_0.GetTokenAdminRegistryAddress(input.ExistingDataStore, input.ChainSelector, &t.EVMTokenBase)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve TAR address for chain %d: %w", input.ChainSelector, err)
				}
				input.RegistryAddress = tarAddr.Hex()
			}

			report, err := cldf_ops.ExecuteSequence(b, evm_seq.ConfigureTokenForTransfers, chains, input)
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
		tpOpsV1_6_1.GetToken, chain,
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

func (p *poolOpsV161) GetTokenDecimals(ctx context.Context, chain evm.Chain, poolAddr common.Address) (uint8, error) {
	pool, err := tpV1_6_1.NewTokenPool(poolAddr, chain.Client)
	if err != nil {
		return 0, fmt.Errorf("failed to instantiate token pool v1.6.1 contract: %w", err)
	}
	return pool.GetTokenDecimals(&bind.CallOpts{Context: ctx})
}

func (p *poolOpsV161) GetPoolAdmins(ctx context.Context, chain *evm.Chain, poolAddr common.Address) (owner, rlAdmin common.Address, err error) {
	pool, err := tpV1_6_1.NewTokenPool(poolAddr, chain.Client)
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

func (p *poolOpsV161) SetRateLimiterConfig(b cldf_ops.Bundle, chain evm.Chain, poolAddr common.Address, remoteChainSelector uint64, outbound, inbound tokensapi.RateLimiterConfig) (evm_contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b,
		tpOpsV1_6_1.SetChainRateLimiterConfig, chain,
		evm_contract.FunctionInput[tpOpsV1_6_1.SetChainRateLimiterConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       poolAddr,
			Args: tpOpsV1_6_1.SetChainRateLimiterConfigArgs{
				OutboundConfig: tpOpsV1_6_1.Config{
					IsEnabled: outbound.IsEnabled,
					Capacity:  outbound.Capacity,
					Rate:      outbound.Rate,
				},
				InboundConfig: tpOpsV1_6_1.Config{
					IsEnabled: inbound.IsEnabled,
					Capacity:  inbound.Capacity,
					Rate:      inbound.Rate,
				},
				RemoteChainSelector: remoteChainSelector,
			},
		})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("SetChainRateLimiterConfig v1.6.1: %w", err)
	}
	return report.Output, nil
}

func (p *poolOpsV161) Version() *semver.Version {
	return tpOpsV1_6_1.Version
}
