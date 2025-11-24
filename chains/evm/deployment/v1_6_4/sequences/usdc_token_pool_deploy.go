package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type USDCTokenPoolDeploySequenceInput struct {
	ChainSelector               uint64
	TokenMessenger              common.Address
	CCTPMessageTransmitterProxy common.Address
	Token                       common.Address
	Allowlist                   []common.Address
	RMNProxy                    common.Address
	Router                      common.Address
	SupportedUSDCVersion        uint32
	MCMSAddress                 common.Address
}

var USDCTokenPoolDeploySequence = operations.NewSequence(
	"USDCTokenPoolDeploySequence",
	semver.MustParse("1.6.4"),
	"Deploys the USDCTokenPool contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input USDCTokenPoolDeploySequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}

		var addresses []datastore.AddressRef

		// Deploy USDCTokenPool using the deployed CCTPMessageTransmitterProxy address
		report, err := operations.ExecuteOperation(b, usdc_token_pool.Deploy, chain, contract.DeployInput[usdc_token_pool.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(
				usdc_token_pool.ContractType,
				*usdc_token_pool.Version,
			),
			Args: usdc_token_pool.ConstructorArgs{
				TokenMessenger:              input.TokenMessenger,
				CCTPMessageTransmitterProxy: input.CCTPMessageTransmitterProxy,
				Token:                       input.Token,
				Allowlist:                   input.Allowlist,
				RMNProxy:                    input.RMNProxy,
				Router:                      input.Router,
				SupportedUSDCVersion:        input.SupportedUSDCVersion,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPool on %s: %w", chain, err)
		}

		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure allowed callers for the CCTPMessageTransmitterProxy on %s: %w", chain, err)
		}

		// Begin transferring ownership of the token pool to MCMS
		_, err = operations.ExecuteOperation(b, usdc_token_pool_ops.USDCTokenPoolTransferOwnership, chain, contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(report.Output.Address),
			Args:          input.MCMSAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership of the token pool to MCMS on %s: %w", chain, err)
		}

		addresses = append(addresses, report.Output)
		return sequences.OnChainOutput{
			Addresses: addresses,
		}, nil
	})
