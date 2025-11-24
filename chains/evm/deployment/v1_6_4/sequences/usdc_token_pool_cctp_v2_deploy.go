package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/cctp_message_transmitter_proxy"
	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_cctp_v2"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type USDCTokenPoolCCTPV2DeploySequenceInput struct {
	ChainSelector  uint64
	TokenMessenger common.Address
	Token          common.Address
	Allowlist      []common.Address
	RMNProxy       common.Address
	Router         common.Address
	MCMSAddress    common.Address
}

var USDCTokenPoolCCTPV2DeploySequence = operations.NewSequence(
	"USDCTokenPoolCCTPV2DeploySequence",
	semver.MustParse("1.6.4"),
	"Deploys the CCTP V2 pool on a USDCTokenPool contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input USDCTokenPoolCCTPV2DeploySequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}

		// Deploy CCTPMessageTransmitterProxy first so that it can be used by the USDCTokenPoolCCTPV2 contract
		cctpProxyReport, err := operations.ExecuteOperation(b, cctp_message_transmitter_proxy.Deploy, chain, contract.DeployInput[cctp_message_transmitter_proxy.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(
				cctp_message_transmitter_proxy.ContractType,
				*cctp_message_transmitter_proxy.Version,
			),
			Args: cctp_message_transmitter_proxy.ConstructorArgs{
				TokenMessenger: input.TokenMessenger,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCTPMessageTransmitterProxy on %s: %w", chain, err)
		}

		// Get the deployed address from the report
		cctpProxyAddress := common.HexToAddress(cctpProxyReport.Output.Address)

		// Deploy USDCTokenPoolCCTPV2 using the deployed CCTPMessageTransmitterProxy address
		report, err := operations.ExecuteOperation(b, usdc_token_pool_cctp_v2.Deploy, chain, contract.DeployInput[usdc_token_pool_cctp_v2.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(
				usdc_token_pool_cctp_v2.ContractType,
				*usdc_token_pool_cctp_v2.Version,
			),
			Args: usdc_token_pool_cctp_v2.ConstructorArgs{
				TokenMessenger:              input.TokenMessenger,
				CCTPMessageTransmitterProxy: cctpProxyAddress,
				Token:                       input.Token,
				Allowlist:                   input.Allowlist,
				RMNProxy:                    input.RMNProxy,
				Router:                      input.Router,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolCCTPV2 on %s: %w", chain, err)
		}

		// Configure the allowed callers for the CCTPMessageTransmitterProxy
		_, err = operations.ExecuteOperation(b, cctp_message_transmitter_proxy.CCTPMessageTransmitterProxyConfigureAllowedCallers, chain, contract.FunctionInput[[]cctp_message_transmitter_proxy.AllowedCallerConfigArgs]{
			ChainSelector: input.ChainSelector,
			Address:       cctpProxyAddress,
			Args: []cctp_message_transmitter_proxy.AllowedCallerConfigArgs{
				{
					Caller:  common.HexToAddress(report.Output.Address),
					Allowed: true,
				},
			},
		})

		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure allowed callers for the CCTPMessageTransmitterProxy on %s: %w", chain, err)
		}

		// Begin transferring ownership to MCMS. A separate changeset will be used to accept ownership.
		_, err = operations.ExecuteOperation(b, usdc_token_pool_ops.USDCTokenPoolTransferOwnership, chain, contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       cctpProxyAddress,
			Args:          input.MCMSAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership of the CCTPMessageTransmitterProxy to MCMS on %s: %w", chain, err)
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

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{report.Output, cctpProxyReport.Output},
		}, nil
	})
