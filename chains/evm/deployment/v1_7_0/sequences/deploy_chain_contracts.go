package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/nonce_manager"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	mock_receiver "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type RMNRemoteParams struct {
	LegacyRMN common.Address
}

type CCVAggregatorParams struct {
	GasForCallExactCheck uint16
}

type CommitOnRampParams struct {
	AllowlistAdmin common.Address
	FeeAggregator  common.Address
}

type CommitOffRampParams struct {
	SignatureConfigArgs commit_offramp.SetSignatureConfigArgs
}

type CCVProxyParams struct {
	FeeAggregator common.Address
}

type FeeQuoterParams struct {
	MaxFeeJuelsPerMsg              *big.Int
	TokenPriceStalenessThreshold   uint32
	LINKPremiumMultiplierWeiPerEth uint64
	WETHPremiumMultiplierWeiPerEth uint64
	USDPerLINK                     *big.Int
	USDPerWETH                     *big.Int
}

type ExecutorOnRampParams struct {
	MaxCCVsPerMsg uint8
}

type ContractParams struct {
	RMNRemote      RMNRemoteParams
	CCVAggregator  CCVAggregatorParams
	CommitOnRamp   CommitOnRampParams
	CommitOffRamp  CommitOffRampParams
	CCVProxy       CCVProxyParams
	FeeQuoter      FeeQuoterParams
	ExecutorOnRamp ExecutorOnRampParams
}

type DeployChainContractsInput struct {
	ChainSelector     uint64 // Only exists to differentiate sequence runs on different chains
	ExistingAddresses []datastore.AddressRef
	ContractParams    ContractParams
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("1.7.0"),
	"Deploys all required contracts for CCIP 1.7.0 to an EVM chain",
	func(b operations.Bundle, chain evm.Chain, input DeployChainContractsInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0, 13) // 13 = number of maybeDeployContract calls
		writes := make([]contract.WriteOutput, 0, 4)     // 4 = number of ExecuteOperation calls

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

		// Deploy WETH
		wethRef, err := maybeDeployContract(b, weth.Deploy, weth.ContractType, chain, contract.DeployInput[weth.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := maybeDeployContract(b, link.Deploy, link.ContractType, chain, contract.DeployInput[link.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := maybeDeployContract(b, rmn_remote.Deploy, rmn_remote.ContractType, chain, contract.DeployInput[rmn_remote.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: rmn_remote.ConstructorArgs{
				LocalChainSelector: chain.Selector,
				LegacyRMN:          input.ContractParams.RMNRemote.LegacyRMN,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnRemoteRef)

		// Deploy RMNProxy
		rmnProxyRef, err := maybeDeployContract(b, rmn_proxy.Deploy, rmn_proxy.ContractType, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: rmn_proxy.ConstructorArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnProxyRef)

		// Set the RMNRemote on the RMNProxy
		// Included in case the RMNRemote got deployed but the RMNProxy already existed.
		// In this case, we would not have set the RMNRemote in the constructor.
		// We would need to update the RMN on the existing RMNProxy.
		setRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.SetRMN, chain, contract.FunctionInput[rmn_proxy.SetRMNArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(rmnProxyRef.Address),
			Args: rmn_proxy.SetRMNArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		writes = append(writes, setRMNReport.Output)

		// Deploy Router
		routerRef, err := maybeDeployContract(b, router.Deploy, router.ContractType, chain, contract.DeployInput[router.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, routerRef)

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := maybeDeployContract(b, token_admin_registry.Deploy, token_admin_registry.ContractType, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy FeeQuoter
		feeQuoterRef, err := maybeDeployContract(b, fee_quoter_v2.Deploy, fee_quoter_v2.ContractType, chain, contract.DeployInput[fee_quoter_v2.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: fee_quoter_v2.ConstructorArgs{
				StaticConfig: fee_quoter_v2.StaticConfig{
					MaxFeeJuelsPerMsg:            input.ContractParams.FeeQuoter.MaxFeeJuelsPerMsg,
					TokenPriceStalenessThreshold: input.ContractParams.FeeQuoter.TokenPriceStalenessThreshold,
					LinkToken:                    common.HexToAddress(linkRef.Address),
				},
				PriceUpdaters: []common.Address{
					// NOTE: CommitOffRamp not set here in 1.7.0, as price updates are out of scope for launch.
					// TODO: Add Timelock here when MCMS support is needed.
					chain.DeployerKey.From,
				},
				PremiumMultiplierWeiPerEthArgs: []fee_quoter_v2.PremiumMultiplierWeiPerEthArgs{
					{
						Token:                      common.HexToAddress(linkRef.Address),
						PremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoter.LINKPremiumMultiplierWeiPerEth,
					},
					{
						Token:                      common.HexToAddress(wethRef.Address),
						PremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoter.WETHPremiumMultiplierWeiPerEth,
					},
				},
				FeeTokens: []common.Address{
					common.HexToAddress(linkRef.Address),
					common.HexToAddress(wethRef.Address),
				},
				// Skipped fields:
				// - TokenPriceFeeds (will not be used in 1.7.0)
				// - TokenTransferFeeConfigArgs (token+lane-specific config, set elsewhere)
				// - DestChainConfigArgs (lane-specific config, set elsewhere)
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, feeQuoterRef)

		// Set initial prices on FeeQuoter
		updatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter_v2.UpdatePrices, chain, contract.FunctionInput[fee_quoter_v2.PriceUpdates]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(feeQuoterRef.Address),
			Args: fee_quoter_v2.PriceUpdates{
				TokenPriceUpdates: []fee_quoter_v2.TokenPriceUpdate{
					{
						SourceToken: common.HexToAddress(linkRef.Address),
						UsdPerToken: input.ContractParams.FeeQuoter.USDPerLINK,
					},
					{
						SourceToken: common.HexToAddress(wethRef.Address),
						UsdPerToken: input.ContractParams.FeeQuoter.USDPerWETH,
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set initial prices on FeeQuoter: %w", err)
		}
		writes = append(writes, updatePricesReport.Output)

		// Deploy CCVAggregator
		ccvAggregatorRef, err := maybeDeployContract(b, ccv_aggregator.Deploy, ccv_aggregator.ContractType, chain, contract.DeployInput[ccv_aggregator.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: ccv_aggregator.ConstructorArgs{
				LocalChainSelector:   chain.Selector,
				RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
				GasForCallExactCheck: input.ContractParams.CCVAggregator.GasForCallExactCheck,
				TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCVAggregator: %w", err)
		}
		addresses = append(addresses, ccvAggregatorRef)

		// Deploy CCVProxy
		ccvProxyRef, err := maybeDeployContract(b, ccv_proxy.Deploy, ccv_proxy.ContractType, chain, contract.DeployInput[ccv_proxy.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: ccv_proxy.ConstructorArgs{
				StaticConfig: ccv_proxy.StaticConfig{
					ChainSelector:      chain.Selector,
					RmnRemote:          common.HexToAddress(rmnRemoteRef.Address),
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
				},
				DynamicConfig: ccv_proxy.DynamicConfig{
					FeeQuoter:     common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator: input.ContractParams.CCVProxy.FeeAggregator,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCVProxy: %w", err)
		}
		addresses = append(addresses, ccvProxyRef)

		// Deploy NonceManager
		nonceManagerRef, err := maybeDeployContract(b, nonce_manager.Deploy, nonce_manager.ContractType, chain, contract.DeployInput[nonce_manager.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy NonceManager: %w", err)
		}
		addresses = append(addresses, nonceManagerRef)

		// Deploy CommitOnRamp
		commitOnRampRef, err := maybeDeployContract(b, commit_onramp.Deploy, commit_onramp.ContractType, chain, contract.DeployInput[commit_onramp.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: commit_onramp.ConstructorArgs{
				DynamicConfig: commit_onramp.DynamicConfig{
					FeeQuoter:      common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator:  input.ContractParams.CommitOnRamp.FeeAggregator,
					AllowlistAdmin: input.ContractParams.CommitOnRamp.AllowlistAdmin,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitOnRamp: %w", err)
		}
		addresses = append(addresses, commitOnRampRef)

		// Deploy ExecutorOnRamp
		executorOnRampRef, err := maybeDeployContract(b, executor_onramp.Deploy, executor_onramp.ContractType, chain, contract.DeployInput[executor_onramp.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: executor_onramp.ConstructorArgs{
				MaxCCVsPerMsg: input.ContractParams.ExecutorOnRamp.MaxCCVsPerMsg,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ExecutorOnRamp: %w", err)
		}
		addresses = append(addresses, executorOnRampRef)

		// Deploy CommitOffRamp
		commitOffRampRef, err := maybeDeployContract(b, commit_offramp.Deploy, commit_offramp.ContractType, chain, contract.DeployInput[commit_offramp.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args:          commit_offramp.ConstructorArgs{},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitOffRamp: %w", err)
		}
		addresses = append(addresses, commitOffRampRef)

		mockReceiver, err := maybeDeployContract(b, mock_receiver.Deploy, mock_receiver.ContractType, chain, contract.DeployInput[mock_receiver.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: mock_receiver.ConstructorArgs{
				RequiredVerifiers: []common.Address{common.HexToAddress(commitOffRampRef.Address)},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy MockReceiver: %w", err)
		}
		addresses = append(addresses, mockReceiver)

		// Set signature config on the CommitOffRamp
		setSignatureConfigReport, err := cldf_ops.ExecuteOperation(b, commit_offramp.SetSignatureConfigs, chain, contract.FunctionInput[commit_offramp.SetSignatureConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(commitOffRampRef.Address),
			Args:          input.ContractParams.CommitOffRamp.SignatureConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set signature config on CommitOffRamp: %w", err)
		}
		writes = append(writes, setSignatureConfigReport.Output)

		// Add CommitOnRamp and CommitOffRamp as AuthorizedCallers on NonceManager
		applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, nonce_manager.ApplyAuthorizedCallerUpdates, chain, contract.FunctionInput[nonce_manager.AuthorizedCallerArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(nonceManagerRef.Address),
			Args: nonce_manager.AuthorizedCallerArgs{
				AddedCallers: []common.Address{
					common.HexToAddress(commitOnRampRef.Address),
					common.HexToAddress(commitOffRampRef.Address),
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to add CommitOnRamp and CommitOffRamp as AuthorizedCallers to NonceManager: %w", err)
		}
		writes = append(writes, applyAuthorizedCallerUpdatesReport.Output)

		return sequences.OnChainOutput{
			Addresses: addresses,
			Writes:    writes,
		}, nil
	},
)

func maybeDeployContract[ARGS any](
	b operations.Bundle,
	op *operations.Operation[contract.DeployInput[ARGS], datastore.AddressRef, evm.Chain],
	contractType cldf_deployment.ContractType,
	chain evm.Chain,
	input contract.DeployInput[ARGS],
	existingAddresses []datastore.AddressRef,
) (datastore.AddressRef, error) {
	for _, ref := range existingAddresses {
		if ref.Type == datastore.ContractType(contractType) {
			return ref, nil
		}
	}
	report, err := cldf_ops.ExecuteOperation(b, op, chain, input)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s %s: %w", contractType, op.Def().Version, err)
	}
	return report.Output, nil
}
