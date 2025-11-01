package tokens

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// DeployTokenPoolInput is the input for the DeployTokenPool sequence.
type DeployTokenPoolInput struct {
	// ChainSel is the chain selector for the chain being configured.
	ChainSel uint64
	// TokenPoolType is the type of the token pool to deploy.
	TokenPoolType datastore.ContractType
	// TokenPoolVersion is the version of the token pool to deploy.
	TokenPoolVersion *semver.Version
	// TokenSymbol is the symbol of the token to be configured.
	// This symbol will be stored in the returned AddressRef.
	TokenSymbol string
	// RateLimitAdmin is an additional address allowed to set rate limiters.
	// If left empty, setRateLimitAdmin will not be attempted.
	RateLimitAdmin common.Address
	// ConstructorArgs are the constructor arguments for the token pool.
	ConstructorArgs token_pool.ConstructorArgs
}

func (c DeployTokenPoolInput) ChainSelector() uint64 {
	return c.ChainSel
}

func (c DeployTokenPoolInput) Validate(chain evm.Chain) error {
	if c.ChainSel != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSel, chain)
	}
	if c.TokenSymbol == "" {
		return errors.New("token symbol must be defined")
	}
	if c.TokenPoolType == "" {
		return errors.New("token pool type must be defined")
	}
	if c.TokenPoolVersion == nil {
		return errors.New("token pool version must be defined")
	}
	if c.ConstructorArgs.Token == (common.Address{}) {
		return errors.New("token address must be defined")
	}
	if c.ConstructorArgs.RMNProxy == (common.Address{}) {
		return errors.New("rmn proxy address must be defined")
	}
	if c.ConstructorArgs.Router == (common.Address{}) {
		return errors.New("router address must be defined")
	}

	return nil
}

var DeployTokenPool = cldf_ops.NewSequence(
	"deploy-token-pool",
	semver.MustParse("1.7.0"),
	"Deploys a token pool to an EVM chain",
	func(b operations.Bundle, chain evm.Chain, input DeployTokenPoolInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}
		typeAndVersion := deployment.NewTypeAndVersion(
			deployment.ContractType(input.TokenPoolType),
			*input.TokenPoolVersion,
		)
		deployReport, err := cldf_ops.ExecuteOperation(b, token_pool.Deploy, chain, evm_contract.DeployInput[token_pool.ConstructorArgs]{
			ChainSelector:  input.ChainSel,
			TypeAndVersion: typeAndVersion,
			Args:           input.ConstructorArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy %s to %s: %w", typeAndVersion, chain, err)
		}
		deployReport.Output.Qualifier = input.TokenSymbol // Use the token symbol as the qualifier.

		configureReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSel,
			TokenPoolAddress: common.HexToAddress(deployReport.Output.Address),
			RateLimitAdmin:   input.RateLimitAdmin,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s: %w", deployReport.Output.Address, chain, err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{deployReport.Output},
			BatchOps:  configureReport.Output.BatchOps,
		}, nil
	},
)
