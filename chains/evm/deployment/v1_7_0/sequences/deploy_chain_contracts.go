package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/optypes/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/optypes/deployment"
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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	NUM_CONTRACTS = 12
	NUM_TXS       = 2
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

type CCVProxyParams struct {
	FeeAggregator common.Address
}

type FeeQuoterV2Params struct {
	MaxFeeJuelsPerMsg              *big.Int
	TokenPriceStalenessThreshold   uint32
	LINKPremiumMultiplierWeiPerEth uint64
	WETHPremiumMultiplierWeiPerEth uint64
}

type ContractParams struct {
	RMNRemote     RMNRemoteParams
	CCVAggregator CCVAggregatorParams
	CommitOnRamp  CommitOnRampParams
	CCVProxy      CCVProxyParams
	FeeQuoterV2   FeeQuoterV2Params
}

type DeployChainContractsInput struct {
	ExistingAddresses []datastore.AddressRef
	ContractParams    ContractParams
}

type DeployChainContractsOutput struct {
	Addresses []datastore.AddressRef
	Writes    []call.WriteOutput
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("1.7.0"),
	"Deploys all required contracts for CCIP 1.7.0 to a chain",
	func(b operations.Bundle, chain evm.Chain, input DeployChainContractsInput) (output DeployChainContractsOutput, err error) {
		addresses := make([]datastore.AddressRef, 0, NUM_CONTRACTS)
		writes := make([]call.WriteOutput, 0, NUM_TXS)

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

		// Deploy WETH
		wethRef, err := maybeDeployContract(b, weth.Deploy, weth.ContractType, chain, deployment.Input[weth.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := maybeDeployContract(b, link.Deploy, link.ContractType, chain, deployment.Input[link.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNProxy
		rmnProxyRef, err := maybeDeployContract(b, rmn_proxy.Deploy, rmn_proxy.ContractType, chain, deployment.Input[rmn_proxy.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args:          rmn_proxy.ConstructorArgs{},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, rmnProxyRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := maybeDeployContract(b, rmn_remote.Deploy, rmn_remote.ContractType, chain, deployment.Input[rmn_remote.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: rmn_remote.ConstructorArgs{
				LocalChainSelector: chain.Selector,
				LegacyRMN:          input.ContractParams.RMNRemote.LegacyRMN,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, rmnRemoteRef)

		// Set the RMNRemote on the RMNProxy
		setRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.SetRMN, chain, call.Input[rmn_proxy.SetRMNArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(rmnProxyRef.Address),
			Args: rmn_proxy.SetRMNArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			},
		})
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		writes = append(writes, setRMNReport.Output)

		// Deploy Router
		routerRef, err := maybeDeployContract(b, router.Deploy, router.ContractType, chain, deployment.Input[router.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, routerRef)

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := maybeDeployContract(b, token_admin_registry.Deploy, token_admin_registry.ContractType, chain, deployment.Input[token_admin_registry.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy FeeQuoterV2
		feeQuoterV2Ref, err := maybeDeployContract(b, fee_quoter_v2.Deploy, fee_quoter_v2.ContractType, chain, deployment.Input[fee_quoter_v2.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: fee_quoter_v2.ConstructorArgs{
				StaticConfig: fee_quoter_v2.StaticConfig{
					MaxFeeJuelsPerMsg:            input.ContractParams.FeeQuoterV2.MaxFeeJuelsPerMsg,
					TokenPriceStalenessThreshold: input.ContractParams.FeeQuoterV2.TokenPriceStalenessThreshold,
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
						PremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoterV2.LINKPremiumMultiplierWeiPerEth,
					},
					{
						Token:                      common.HexToAddress(wethRef.Address),
						PremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoterV2.WETHPremiumMultiplierWeiPerEth,
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
			return DeployChainContractsOutput{}, err
		}
		addresses = append(addresses, feeQuoterV2Ref)

		// Deploy CCVAggregator
		ccvAggregatorRef, err := maybeDeployContract(b, ccv_aggregator.Deploy, ccv_aggregator.ContractType, chain, deployment.Input[ccv_aggregator.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: ccv_aggregator.ConstructorArgs{
				LocalChainSelector:   chain.Selector,
				RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
				GasForCallExactCheck: input.ContractParams.CCVAggregator.GasForCallExactCheck,
				TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, fmt.Errorf("failed to deploy CCVAggregator: %w", err)
		}
		addresses = append(addresses, ccvAggregatorRef)

		// Deploy CCVProxy
		ccvProxyRef, err := maybeDeployContract(b, ccv_proxy.Deploy, ccv_proxy.ContractType, chain, deployment.Input[ccv_proxy.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: ccv_proxy.ConstructorArgs{
				StaticConfig: ccv_proxy.StaticConfig{
					ChainSelector:      chain.Selector,
					RmnRemote:          common.HexToAddress(rmnRemoteRef.Address),
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
				},
				DynamicConfig: ccv_proxy.DynamicConfig{
					FeeQuoter:     common.HexToAddress(feeQuoterV2Ref.Address),
					FeeAggregator: input.ContractParams.CCVProxy.FeeAggregator,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, fmt.Errorf("failed to deploy CCVProxy: %w", err)
		}
		addresses = append(addresses, ccvProxyRef)

		// Deploy NonceManager
		nonceManagerRef, err := maybeDeployContract(b, nonce_manager.Deploy, nonce_manager.ContractType, chain, deployment.Input[nonce_manager.ConstructorArgs]{
			ChainSelector: chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, fmt.Errorf("failed to deploy NonceManager: %w", err)
		}
		addresses = append(addresses, nonceManagerRef)

		// Deploy CommitOnRamp
		commitOnRampRef, err := maybeDeployContract(b, commit_onramp.Deploy, commit_onramp.ContractType, chain, deployment.Input[commit_onramp.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: commit_onramp.ConstructorArgs{
				RMNRemote:    common.HexToAddress(rmnRemoteRef.Address),
				NonceManager: common.HexToAddress(nonceManagerRef.Address),
				DynamicConfig: commit_onramp.DynamicConfig{
					FeeQuoter:      common.HexToAddress(feeQuoterV2Ref.Address),
					FeeAggregator:  input.ContractParams.CommitOnRamp.FeeAggregator,
					AllowlistAdmin: input.ContractParams.CommitOnRamp.AllowlistAdmin,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, fmt.Errorf("failed to deploy CommitOnRamp: %w", err)
		}
		addresses = append(addresses, commitOnRampRef)

		// Deploy CommitOffRamp
		commitOffRampRef, err := maybeDeployContract(b, commit_offramp.Deploy, commit_offramp.ContractType, chain, deployment.Input[commit_offramp.ConstructorArgs]{
			ChainSelector: chain.Selector,
			Args: commit_offramp.ConstructorArgs{
				NonceManager: common.HexToAddress(nonceManagerRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return DeployChainContractsOutput{}, fmt.Errorf("failed to deploy CommitOffRamp: %w", err)
		}
		addresses = append(addresses, commitOffRampRef)

		// Add CommitOnRamp and CommitOffRamp as AuthorizedCallers on NonceManager
		applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, nonce_manager.ApplyAuthorizedCallerUpdates, chain, call.Input[nonce_manager.AuthorizedCallerArgs]{
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
			return DeployChainContractsOutput{}, fmt.Errorf("failed to add CommitOnRamp and CommitOffRamp as AuthorizedCallers to NonceManager: %w", err)
		}
		writes = append(writes, applyAuthorizedCallerUpdatesReport.Output)

		return DeployChainContractsOutput{
			Addresses: addresses,
			Writes:    writes,
		}, nil
	},
)

func maybeDeployContract[ARGS any](
	b operations.Bundle,
	op *operations.Operation[deployment.Input[ARGS], datastore.AddressRef, evm.Chain],
	contractType cldf_deployment.ContractType,
	chain evm.Chain,
	input deployment.Input[ARGS],
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
