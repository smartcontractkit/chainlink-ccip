package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	mock_receiver "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type RMNRemoteParams struct {
	Version   *semver.Version
	LegacyRMN common.Address
}

type CCVAggregatorParams struct {
	Version              *semver.Version
	GasForCallExactCheck uint16
}

type CommitteeVerifierParams struct {
	Version             *semver.Version
	AllowlistAdmin      common.Address
	FeeAggregator       common.Address
	SignatureConfigArgs committee_verifier.SetSignatureConfigArgs
	StorageLocation     string
}

type CCVProxyParams struct {
	Version       *semver.Version
	FeeAggregator common.Address
}

type FeeQuoterParams struct {
	Version                        *semver.Version
	MaxFeeJuelsPerMsg              *big.Int
	TokenPriceStalenessThreshold   uint32
	LINKPremiumMultiplierWeiPerEth uint64
	WETHPremiumMultiplierWeiPerEth uint64
	USDPerLINK                     *big.Int
	USDPerWETH                     *big.Int
}

type ExecutorOnRampParams struct {
	Version       *semver.Version
	MaxCCVsPerMsg uint8
}

type ContractParams struct {
	RMNRemote         RMNRemoteParams
	CCVAggregator     CCVAggregatorParams
	CommitteeVerifier CommitteeVerifierParams
	CCVProxy          CCVProxyParams
	FeeQuoter         FeeQuoterParams
	ExecutorOnRamp    ExecutorOnRampParams
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
		writes := make([]contract.WriteOutput, 0, 3)     // 3 calls

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

		// Deploy WETH
		wethRef, err := maybeDeployContract(b, weth.Deploy, weth.ContractType, chain, evm_contract.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(weth.ContractType), *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := maybeDeployContract(b, link.Deploy, link.ContractType, chain, evm_contract.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(link.ContractType), *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		if input.ContractParams.RMNRemote.Version == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("rmnRemote version must be specified")
		}
		rmnRemoteRef, err := maybeDeployContract(b, rmn_remote.Deploy, rmn_remote.ContractType, chain, evm_contract.DeployInput[rmn_remote.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(rmn_remote.ContractType), *input.ContractParams.RMNRemote.Version),
			ChainSelector:  chain.Selector,
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
		rmnProxyRef, err := maybeDeployContract(b, rmn_proxy.Deploy, rmn_proxy.ContractType, chain, evm_contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(rmn_proxy.ContractType), *semver.MustParse("1.0.0")),
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
		setRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.SetRMN, chain, evm_contract.FunctionInput[rmn_proxy.SetRMNArgs]{
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
		routerRef, err := maybeDeployContract(b, router.Deploy, router.ContractType, chain, evm_contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(router.ContractType), *semver.MustParse("1.2.0")),
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

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := maybeDeployContract(b, token_admin_registry.Deploy, token_admin_registry.ContractType, chain, evm_contract.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(token_admin_registry.ContractType), *semver.MustParse("1.5.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy FeeQuoter
		if input.ContractParams.FeeQuoter.Version == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("feeQuoter version must be specified")
		}
		feeQuoterRef, err := maybeDeployContract(b, fee_quoter_v2.Deploy, fee_quoter_v2.ContractType, chain, evm_contract.DeployInput[fee_quoter_v2.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(fee_quoter_v2.ContractType), *input.ContractParams.FeeQuoter.Version),
			ChainSelector:  chain.Selector,
			Args: fee_quoter_v2.ConstructorArgs{
				StaticConfig: fee_quoter_v2.StaticConfig{
					MaxFeeJuelsPerMsg:            input.ContractParams.FeeQuoter.MaxFeeJuelsPerMsg,
					TokenPriceStalenessThreshold: input.ContractParams.FeeQuoter.TokenPriceStalenessThreshold,
					LinkToken:                    common.HexToAddress(linkRef.Address),
				},
				PriceUpdaters: []common.Address{
					// Price updates via protocol are out of scope for initial launch.
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
		updatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter_v2.UpdatePrices, chain, evm_contract.FunctionInput[fee_quoter_v2.PriceUpdates]{
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
		if input.ContractParams.CCVAggregator.Version == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("ccvAggregator version must be specified")
		}
		ccvAggregatorRef, err := maybeDeployContract(b, ccv_aggregator.Deploy, ccv_aggregator.ContractType, chain, evm_contract.DeployInput[ccv_aggregator.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(ccv_aggregator.ContractType), *input.ContractParams.CCVAggregator.Version),
			ChainSelector:  chain.Selector,
			Args: ccv_aggregator.ConstructorArgs{
				StaticConfig: ccv_aggregator.StaticConfig{
					LocalChainSelector:   chain.Selector,
					RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
					GasForCallExactCheck: input.ContractParams.CCVAggregator.GasForCallExactCheck,
					TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CCVAggregator: %w", err)
		}
		addresses = append(addresses, ccvAggregatorRef)

		// Deploy CCVProxy
		if input.ContractParams.CCVProxy.Version == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("ccvProxy version must be specified")
		}
		ccvProxyRef, err := maybeDeployContract(b, ccv_proxy.Deploy, ccv_proxy.ContractType, chain, evm_contract.DeployInput[ccv_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(ccv_proxy.ContractType), *input.ContractParams.CCVProxy.Version),
			ChainSelector:  chain.Selector,
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

		// Deploy CommitteeVerifier
		if input.ContractParams.CommitteeVerifier.Version == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("committeeVerifier version must be specified")
		}
		committeeVerifierRef, err := maybeDeployContract(b, committee_verifier.Deploy, committee_verifier.ContractType, chain, evm_contract.DeployInput[committee_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(committee_verifier.ContractType), *input.ContractParams.CommitteeVerifier.Version),
			ChainSelector:  chain.Selector,
			Args: committee_verifier.ConstructorArgs{
				DynamicConfig: committee_verifier.DynamicConfig{
					FeeQuoter:      common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator:  input.ContractParams.CommitteeVerifier.FeeAggregator,
					AllowlistAdmin: input.ContractParams.CommitteeVerifier.AllowlistAdmin,
				},
				StorageLocation: input.ContractParams.CommitteeVerifier.StorageLocation,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
		}
		addresses = append(addresses, committeeVerifierRef)

		// Set signature config on the CommitteeVerifier
		setSignatureConfigReport, err := cldf_ops.ExecuteOperation(b, committee_verifier.SetSignatureConfigs, chain, evm_contract.FunctionInput[committee_verifier.SetSignatureConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(committeeVerifierRef.Address),
			Args:          input.ContractParams.CommitteeVerifier.SignatureConfigArgs,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set signature config on CommitteeVerifier: %w", err)
		}
		writes = append(writes, setSignatureConfigReport.Output)

		// Deploy CommitteeVerifierProxy
		committeeVerifierProxyRef, err := maybeDeployContract(b, committee_verifier.DeployProxy, committee_verifier.ProxyType, chain, evm_contract.DeployInput[committee_verifier.ProxyConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(committee_verifier.ProxyType), *semver.MustParse("1.7.0")),
			ChainSelector:  chain.Selector,
			Args: committee_verifier.ProxyConstructorArgs{
				RampAddress: common.HexToAddress(committeeVerifierRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifierProxy: %w", err)
		}
		addresses = append(addresses, committeeVerifierProxyRef)

		// Deploy ExecutorOnRamp
		if input.ContractParams.ExecutorOnRamp.Version == nil {
			return sequences.OnChainOutput{}, fmt.Errorf("executorOnRamp version must be specified")
		}
		executorOnRampRef, err := maybeDeployContract(b, executor_onramp.Deploy, executor_onramp.ContractType, chain, evm_contract.DeployInput[executor_onramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(executor_onramp.ContractType), *input.ContractParams.ExecutorOnRamp.Version),
			ChainSelector:  chain.Selector,
			Args: executor_onramp.ConstructorArgs{
				MaxCCVsPerMsg: input.ContractParams.ExecutorOnRamp.MaxCCVsPerMsg,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ExecutorOnRamp: %w", err)
		}
		addresses = append(addresses, executorOnRampRef)

		// Deploy MockReceiver (defines committee verifier as required)
		// TODO: Replace with a more configurable receiver contract.
		mockReceiver, err := maybeDeployContract(b, mock_receiver.Deploy, mock_receiver.ContractType, chain, evm_contract.DeployInput[mock_receiver.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(deployment.ContractType(mock_receiver.ContractType), *semver.MustParse("1.7.0")),
			ChainSelector:  chain.Selector,
			Args: mock_receiver.ConstructorArgs{
				RequiredVerifiers: []common.Address{common.HexToAddress(committeeVerifierRef.Address)},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy MockReceiver: %w", err)
		}
		addresses = append(addresses, mockReceiver)

		return sequences.OnChainOutput{
			Addresses: addresses,
			Writes:    writes,
		}, nil
	},
)

func maybeDeployContract[ARGS any](
	b operations.Bundle,
	op *operations.Operation[evm_contract.DeployInput[ARGS], datastore.AddressRef, evm.Chain],
	contractType cldf_deployment.ContractType,
	chain evm.Chain,
	input evm_contract.DeployInput[ARGS],
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
