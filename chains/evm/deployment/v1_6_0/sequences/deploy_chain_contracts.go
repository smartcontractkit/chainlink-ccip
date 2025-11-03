package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/nonce_manager"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

func (a *EVMAdapter) DeployChainContracts() *cldf_ops.Sequence[deployops.ContractDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return DeployChainContracts
}

// just a wrapper around the v1.0.0 deployer for now
func (a *EVMAdapter) DeployMCMS() *cldf_ops.Sequence[deployops.MCMSDeploymentConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMDeployer{}
	return evmDeployer.DeployMCMS()
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("1.6.0"),
	"Deploys all required contracts for CCIP 1.6.0 to an EVM chain",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input deployops.ContractDeploymentConfigPerChainWithAddress) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		chain := chains.EVMChains()[input.ChainSelector]

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

		// Deploy WETH
		wethRef, err := contract.MaybeDeployContract(b, weth.Deploy, chain, contract.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(weth.ContractType, *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := contract.MaybeDeployContract(b, link.Deploy, chain, contract.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link.ContractType, *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := contract.MaybeDeployContract(b, rmn_remote.Deploy, chain, contract.DeployInput[rmn_remote.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_remote.ContractType, *rmn_remote.Version),
			ChainSelector:  chain.Selector,
			Args: rmn_remote.ConstructorArgs{
				LocalChainSelector: chain.Selector,
				LegacyRMN:          common.HexToAddress(input.LegacyRMN),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnRemoteRef)

		// Deploy RMNProxy
		rmnProxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
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
		routerRef, err := contract.MaybeDeployContract(b, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.ContractType, *router.Version),
			ChainSelector:  chain.Selector,
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, routerRef)

		// Deploy Test Router
		testRouterRef, err := contract.MaybeDeployContract(b, router.DeployTestRouter, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.TestRouterContractType, *router.Version),
			ChainSelector:  chain.Selector,
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, testRouterRef)

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := contract.MaybeDeployContract(b, token_admin_registry.Deploy, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy NonceManager
		nonceManagerRef, err := contract.MaybeDeployContract(b, nonce_manager.Deploy, chain, contract.DeployInput[nonce_manager.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(nonce_manager.ContractType, *nonce_manager.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, nonceManagerRef)

		// Deploy FeeQuoter
		feeQuoterRef, err := contract.MaybeDeployContract(b, fqops.Deploy, chain, contract.DeployInput[fqops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(fqops.ContractType, *fqops.Version),
			ChainSelector:  chain.Selector,
			Args: fqops.ConstructorArgs{
				StaticConfig: fee_quoter.FeeQuoterStaticConfig{
					MaxFeeJuelsPerMsg:            input.MaxFeeJuelsPerMsg,
					LinkToken:                    common.HexToAddress(linkRef.Address),
					TokenPriceStalenessThreshold: input.TokenPriceStalenessThreshold,
				},
				PriceUpdaters: []common.Address{
					// TODO: Add Timelock here when MCMS support is needed.
					chain.DeployerKey.From,
				},
				FeeTokens: []common.Address{
					common.HexToAddress(linkRef.Address),
					common.HexToAddress(wethRef.Address),
				},
				TokenPriceFeedUpdates:      []fee_quoter.FeeQuoterTokenPriceFeedUpdate{},
				TokenTransferFeeConfigArgs: []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs{},
				MorePremiumMultiplierWeiPerEth: []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs{
					{
						PremiumMultiplierWeiPerEth: input.LinkPremiumMultiplier,
						Token:                      common.HexToAddress(linkRef.Address),
					},
					{
						PremiumMultiplierWeiPerEth: input.NativeTokenPremiumMultiplier,
						Token:                      common.HexToAddress(wethRef.Address),
					},
				},
				DestChainConfigArgs: []fee_quoter.FeeQuoterDestChainConfigArgs{},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, feeQuoterRef)

		// Deploy OffRamp
		offRampRef, err := contract.MaybeDeployContract(b, offrampops.Deploy, chain, contract.DeployInput[offrampops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(offrampops.ContractType, *offrampops.Version),
			ChainSelector:  chain.Selector,
			Args: offrampops.ConstructorArgs{
				StaticConfig: offrampops.StaticConfig{
					ChainSelector:        chain.Selector,
					GasForCallExactCheck: input.GasForCallExactCheck,
					RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
					NonceManager:         common.HexToAddress(nonceManagerRef.Address),
					TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
				},
				DynamicConfig: offrampops.DynamicConfig{
					FeeQuoter:                               common.HexToAddress(feeQuoterRef.Address),
					PermissionLessExecutionThresholdSeconds: input.PermissionLessExecutionThresholdSeconds,
					MessageInterceptor:                      common.HexToAddress(input.MessageInterceptor),
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OffRamp: %w", err)
		}
		addresses = append(addresses, offRampRef)

		// Deploy OnRamp
		onRampRef, err := contract.MaybeDeployContract(b, onrampops.Deploy, chain, contract.DeployInput[onrampops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(onrampops.ContractType, *onrampops.Version),
			ChainSelector:  chain.Selector,
			Args: onrampops.ConstructorArgs{
				StaticConfig: onrampops.StaticConfig{
					ChainSelector:      chain.Selector,
					RmnRemote:          common.HexToAddress(rmnProxyRef.Address),
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
					NonceManager:       common.HexToAddress(nonceManagerRef.Address),
				},
				DynamicConfig: onrampops.DynamicConfig{
					FeeQuoter:     common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator: chain.DeployerKey.From,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OnRamp: %w", err)
		}
		addresses = append(addresses, onRampRef)

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)
