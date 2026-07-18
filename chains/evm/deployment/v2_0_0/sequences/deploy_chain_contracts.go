package sequences

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	fq_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/fee_quoter"
	mock_receiver_v2_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/mock_receiver_v2"
	onramp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/onramp"
	offramp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/offramp"
	executor_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/executor"
	rmn_proxy_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	proxy_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/proxy"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	router_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/registry_module_owner_custom"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool_factory"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

type proxyAcceptOwnershipArgs struct {
	IsProposedOwner bool
}

func newWriteProxyAcceptOwnership(c proxy_bindings.ProxyInterface) *cldf_ops.Operation[contract.FunctionInput[proxyAcceptOwnershipArgs], contract.WriteOutput, evm.Chain] {
	return contract.NewWrite(contract.WriteParams[proxyAcceptOwnershipArgs, proxy_bindings.ProxyInterface]{
		Name:         "proxy:accept-ownership",
		Version:      proxy.Version,
		Description:  "Accept ownership of the proxy",
		ContractType: proxy.ContractType,
		ContractABI:  proxy_bindings.ProxyABI,
		Contract:     c,
		IsAllowedCaller: func(_ proxy_bindings.ProxyInterface, _ *bind.CallOpts, _ common.Address, args proxyAcceptOwnershipArgs) (bool, error) {
			return args.IsProposedOwner, nil
		},
		Validate: func(proxyAcceptOwnershipArgs) error { return nil },
		CallContract: func(p proxy_bindings.ProxyInterface, opts *bind.TransactOpts, _ proxyAcceptOwnershipArgs) (*types.Transaction, error) {
			return p.AcceptOwnership(opts)
		},
	})
}

type MockReceiverParams struct {
	Version *semver.Version
	// RequiredVerifiers are references to verifier contracts that are required.
	// Note that these could be references to contracts that are either already deployed
	// or will be deployed by the DeployChainContracts sequence.
	RequiredVerifiers []datastore.AddressRef
	// OptionalVerifiers are references to verifier contracts that are optional.
	// Note that these could be references to contracts that are either already deployed
	// or will be deployed by the DeployChainContracts sequence.
	OptionalVerifiers []datastore.AddressRef
	OptionalThreshold uint8
	// AllowedFinalityConfig is the finality config that the mock receiver will accept.
	AllowedFinalityConfig [4]byte
	// Qualifier distinguishes between multiple deployments of the mock receiver on the same chain.
	// If only one mock receiver is deployed this can be left blank.
	Qualifier string
}

type RMNRemoteParams struct {
	Version   *semver.Version
	LegacyRMN common.Address
}

type OffRampParams struct {
	Version                   *semver.Version
	GasForCallExactCheck      uint16
	MaxGasBufferToUpdateState uint32
}

type OnRampParams struct {
	Version               *semver.Version
	FeeAggregator         common.Address
	MaxUSDCentsPerMessage uint32
}

type FeeQuoterParams struct {
	Version                        *semver.Version
	MaxFeeJuelsPerMsg              *big.Int
	LINKPremiumMultiplierWeiPerEth uint64
	WETHPremiumMultiplierWeiPerEth uint64
	// USDPerLINK can be nil if setting a LINK price is not desired.
	USDPerLINK *big.Int
	// USDPerWETH can be nil if setting a WETH price is not desired.
	USDPerWETH *big.Int
}

type ExecutorParams struct {
	Version       *semver.Version
	MaxCCVsPerMsg uint8
	DynamicConfig executor_bindings.ExecutorDynamicConfig
	Qualifier     string
}

type ContractParams struct {
	RMNRemote          RMNRemoteParams
	OffRamp            OffRampParams
	CommitteeVerifiers []CommitteeVerifierParams
	OnRamp             OnRampParams
	FeeQuoter          FeeQuoterParams
	Executors          []ExecutorParams
	MockReceivers      []MockReceiverParams
}

type DeployChainContractsInput struct {
	ChainSelector     uint64 // Only exists to differentiate sequence runs on different chains
	CREATE2Factory    common.Address
	ExistingAddresses []datastore.AddressRef
	ContractParams    ContractParams
	DeployTestRouter  bool
	// DeployerKeyOwned, when true, skips the transfer-ownership step so that
	// contracts remain owned by the deployer key. By default (false) the
	// sequence looks up the existing CLLCCIP RBACTimelock in ExistingAddresses
	// and transfers ownership of product contracts to it, failing fast if the
	// required MCMS instances are not found.
	DeployerKeyOwned bool
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("2.0.0"),
	"Deploys all required contracts for CCIP 2.0.0 to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployChainContractsInput) (output ccvadapters.DeployChainContractsOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		ownableContracts := make([]datastore.AddressRef, 0)

		var cllccipTimelockAddr, rmnTimelockAddr common.Address
		if !input.DeployerKeyOwned {
			var mcmContracts []datastore.AddressRef
			cllccipTimelockAddr, rmnTimelockAddr, mcmContracts, err = ResolveOwnershipDeps(
				input.ExistingAddresses, input.ChainSelector,
			)
			if err != nil {
				return output, fmt.Errorf("failed to resolve ownership dependencies: %w", err)
			}
			ownableContracts = append(ownableContracts, mcmContracts...)
		}

		// Deploy WETH
		wethRef, err := evmops.MaybeDeployContract(b, weth.Deploy, chain, contract.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(weth.ContractType, *weth.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := evmops.MaybeDeployContract(b, link.Deploy, chain, contract.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link.ContractType, *link.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := evmops.MaybeDeployContract(b, rmn_remote.Deploy, chain, contract.DeployInput[rmn_remote.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_remote.ContractType, *input.ContractParams.RMNRemote.Version),
			Args: rmn_remote.ConstructorArgs{
				LocalChainSelector: chain.Selector,
				LegacyRMN:          input.ContractParams.RMNRemote.LegacyRMN,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, rmnRemoteRef)
		ownableContracts = append(ownableContracts, rmnRemoteRef)

		// Deploy RMNProxy
		rmnProxyRef, err := evmops.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
			Args: rmn_proxy.ConstructorArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, rmnProxyRef)
		ownableContracts = append(ownableContracts, rmnProxyRef)

		// Fetch the RMN contract address set on the RMNProxy
		rmnAddressReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(rmnProxyRef.Address), rmn_proxy_contract.NewRMNProxy, rmn_proxy.NewReadGetRMN, struct{}{})
		if err != nil {
			return output, err
		}

		// Set the RMNRemote on the RMNProxy if diff exists
		if rmnAddressReport.Output != common.HexToAddress(rmnRemoteRef.Address) {
			setRMNReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(rmnProxyRef.Address), rmn_proxy_contract.NewRMNProxy, rmn_proxy.NewWriteSetRMN, rmn_proxy.SetRMNArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			})
			if err != nil {
				return output, err
			}
			writes = append(writes, setRMNReport.Output)
		}

		// Deploy Router
		routerRef, err := evmops.MaybeDeployContract(b, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.ContractType, *router.Version),
			Args: router.ConstructorArgs{
				WrappedNative: common.HexToAddress(wethRef.Address),
				RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, routerRef)
		ownableContracts = append(ownableContracts, routerRef)

		// Fetch the wrapped native address set on the Router
		wrappedNativeAddressReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(routerRef.Address), router_bindings.NewRouter, router.NewReadGetWrappedNative, struct{}{})
		if err != nil {
			return output, err
		}

		// Set wrapped native on the Router if diff exists
		if wrappedNativeAddressReport.Output != common.HexToAddress(wethRef.Address) {
			setWrappedNativeReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(routerRef.Address), router_bindings.NewRouter, router.NewWriteSetWrappedNative, common.HexToAddress(wethRef.Address))
			if err != nil {
				return output, err
			}
			writes = append(writes, setWrappedNativeReport.Output)
		}

		// Deploy Test Router
		if input.DeployTestRouter {
			testRouterRef, err := evmops.MaybeDeployContract(b, router.DeployTestRouter, chain, contract.DeployInput[router.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(router.TestRouterContractType, *router.Version),
				Args: router.ConstructorArgs{
					WrappedNative: common.HexToAddress(wethRef.Address),
					RMNProxy:      common.HexToAddress(rmnProxyRef.Address),
				},
			}, input.ExistingAddresses)
			if err != nil {
				return output, err
			}
			addresses = append(addresses, testRouterRef)

			// Fetch the wrapped native address set on the Test Router
			wrappedNativeAddressReport, err = evmops.ExecuteRead(b, chain, common.HexToAddress(testRouterRef.Address), router_bindings.NewRouter, router.NewReadGetWrappedNative, struct{}{})
			if err != nil {
				return output, err
			}

			// Set wrapped native on the Test Router if diff exists
			if wrappedNativeAddressReport.Output != common.HexToAddress(wethRef.Address) {
				setWrappedNativeReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(testRouterRef.Address), router_bindings.NewRouter, router.NewWriteSetWrappedNative, common.HexToAddress(wethRef.Address))
				if err != nil {
					return output, err
				}
				writes = append(writes, setWrappedNativeReport.Output)
			}
		}

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := evmops.MaybeDeployContract(b, token_admin_registry.Deploy, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)
		ownableContracts = append(ownableContracts, tokenAdminRegistryRef)

		// Deploy RegistryModuleOwnerCustom
		registryModuleOwnerCustomRef, err := evmops.MaybeDeployContract(b, registry_module_owner_custom.Deploy, chain, contract.DeployInput[registry_module_owner_custom.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(registry_module_owner_custom.ContractType, *registry_module_owner_custom.Version),
			Args: registry_module_owner_custom.ConstructorArgs{
				TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, registryModuleOwnerCustomRef)

		// Add RegistryModuleOwnerCustom to TokenAdminRegistry
		addRegistryModuleReport, hasOnchainDiff, err := MaybeRegisterModuleOnTokenAdminRegistry(
			b,
			chain,
			common.HexToAddress(tokenAdminRegistryRef.Address),
			common.HexToAddress(registryModuleOwnerCustomRef.Address),
		)
		if err != nil {
			return output, err
		}
		// Only append to writes if a transaction was actually created (i.e., module wasn't already registered).
		if hasOnchainDiff {
			writes = append(writes, addRegistryModuleReport)
		}

		// deploy TokenPoolFactory
		tokenPoolFactoryRef, err := evmops.MaybeDeployContract(
			b, token_pool_factory.Deploy, chain,
			contract.DeployInput[token_pool_factory.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(token_pool_factory.ContractType, *token_pool_factory.Version),
				Args: token_pool_factory.ConstructorArgs{
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
					TokenAdminModule:   common.HexToAddress(registryModuleOwnerCustomRef.Address),
					RmnProxy:           common.HexToAddress(rmnProxyRef.Address),
					CcipRouter:         common.HexToAddress(routerRef.Address),
				},
			}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}

		addresses = append(addresses, tokenPoolFactoryRef)

		// Deploy FeeQuoter
		priceUpdaters := []common.Address{chain.DeployerKey.From}
		if cllccipTimelockAddr != (common.Address{}) {
			priceUpdaters = append(priceUpdaters, cllccipTimelockAddr)
		}
		feeQuoterRef, err := evmops.MaybeDeployContract(b, fee_quoter.Deploy, chain, contract.DeployInput[fee_quoter.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(fee_quoter.ContractType, *input.ContractParams.FeeQuoter.Version),
			Args: fee_quoter.ConstructorArgs{
				StaticConfig: fq_bindings.FeeQuoterStaticConfig{
					MaxFeeJuelsPerMsg: input.ContractParams.FeeQuoter.MaxFeeJuelsPerMsg,
					LinkToken:         common.HexToAddress(linkRef.Address),
				},
				PriceUpdaters: priceUpdaters,
				// Skipped fields:
				// - TokenPriceFeeds (will not be used in 2.0.0)
				// - TokenTransferFeeConfigArgs (token+lane-specific config, set elsewhere)
				// - DestChainConfigArgs (lane-specific config, set elsewhere)
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, err
		}
		addresses = append(addresses, feeQuoterRef)
		ownableContracts = append(ownableContracts, feeQuoterRef)

		var tokenPriceUpdates []fq_bindings.InternalTokenPriceUpdate
		if input.ContractParams.FeeQuoter.USDPerLINK != nil {
			tokenPriceUpdates = append(tokenPriceUpdates, fq_bindings.InternalTokenPriceUpdate{
				SourceToken: common.HexToAddress(linkRef.Address),
				UsdPerToken: input.ContractParams.FeeQuoter.USDPerLINK,
			})
		}
		if input.ContractParams.FeeQuoter.USDPerWETH != nil {
			tokenPriceUpdates = append(tokenPriceUpdates, fq_bindings.InternalTokenPriceUpdate{
				SourceToken: common.HexToAddress(wethRef.Address),
				UsdPerToken: input.ContractParams.FeeQuoter.USDPerWETH,
			})
		}
		if len(tokenPriceUpdates) > 0 {
			updatePricesReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(feeQuoterRef.Address), evmops.BindAs[fq_bindings.FeeQuoterInterface](fq_bindings.NewFeeQuoter), fee_quoter.NewWriteUpdatePrices, fq_bindings.InternalPriceUpdates{
				TokenPriceUpdates: tokenPriceUpdates,
			})
			if err != nil {
				return output, fmt.Errorf("failed to update token prices on FeeQuoter: %w", err)
			}
			writes = append(writes, updatePricesReport.Output)
		}

		// Deploy OffRamp
		offRampRef, err := evmops.MaybeDeployContract(b, offramp.Deploy, chain, contract.DeployInput[offramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(offramp.ContractType, *input.ContractParams.OffRamp.Version),
			Args: offramp.ConstructorArgs{
				StaticConfig: offramp_bindings.OffRampStaticConfig{
					LocalChainSelector:        chain.Selector,
					RmnRemote:                 common.HexToAddress(rmnProxyRef.Address),
					GasForCallExactCheck:      input.ContractParams.OffRamp.GasForCallExactCheck,
					TokenAdminRegistry:        common.HexToAddress(tokenAdminRegistryRef.Address),
					MaxGasBufferToUpdateState: input.ContractParams.OffRamp.MaxGasBufferToUpdateState,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, fmt.Errorf("failed to deploy OffRamp: %w", err)
		}
		addresses = append(addresses, offRampRef)
		ownableContracts = append(ownableContracts, offRampRef)

		// Deploy OnRamp
		onRampRef, err := evmops.MaybeDeployContract(b, onramp.Deploy, chain, contract.DeployInput[onramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(onramp.ContractType, *input.ContractParams.OnRamp.Version),
			Args: onramp.ConstructorArgs{
				StaticConfig: onramp_bindings.OnRampStaticConfig{
					ChainSelector:         chain.Selector,
					RmnRemote:             common.HexToAddress(rmnRemoteRef.Address),
					TokenAdminRegistry:    common.HexToAddress(tokenAdminRegistryRef.Address),
					MaxUSDCentsPerMessage: input.ContractParams.OnRamp.MaxUSDCentsPerMessage,
				},
				DynamicConfig: onramp_bindings.OnRampDynamicConfig{
					FeeQuoter:     common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator: input.ContractParams.OnRamp.FeeAggregator,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return output, fmt.Errorf("failed to deploy OnRamp: %w", err)
		}
		addresses = append(addresses, onRampRef)
		ownableContracts = append(ownableContracts, onRampRef)

		// Fetch the dynamic config on the OnRamp
		dynamicConfigReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(onRampRef.Address), evmops.BindAs[onramp_bindings.OnRampInterface](onramp_bindings.NewOnRamp), onramp.NewReadGetDynamicConfig, struct{}{})
		if err != nil {
			return output, fmt.Errorf("failed to get dynamic config on OnRamp: %w", err)
		}

		// Set dynamic config on the OnRamp if there is a diff
		desiredFeeAggregator := dynamicConfigReport.Output.FeeAggregator
		if input.ContractParams.OnRamp.FeeAggregator != (common.Address{}) {
			desiredFeeAggregator = input.ContractParams.OnRamp.FeeAggregator
		}
		if dynamicConfigReport.Output.FeeQuoter != common.HexToAddress(feeQuoterRef.Address) || desiredFeeAggregator != dynamicConfigReport.Output.FeeAggregator {
			desiredDynamicConfig := onramp_bindings.OnRampDynamicConfig{
				FeeQuoter:              common.HexToAddress(feeQuoterRef.Address),
				ReentrancyGuardEntered: false, // This should never be true.
				FeeAggregator:          input.ContractParams.OnRamp.FeeAggregator,
			}
			setDynamicConfigReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(onRampRef.Address), evmops.BindAs[onramp_bindings.OnRampInterface](onramp_bindings.NewOnRamp), onramp.NewWriteSetDynamicConfig, desiredDynamicConfig)
			if err != nil {
				return output, fmt.Errorf("failed to set dynamic config on OnRamp: %w", err)
			}
			writes = append(writes, setDynamicConfigReport.Output)
		}

		// TODO: validate prior to deploying that qualifiers are unique?
		var committeeVerifierBatchOps []mcms_types.BatchOperation
		for _, committeeVerifierParams := range input.ContractParams.CommitteeVerifiers {
			report, err := cldf_ops.ExecuteSequence(b, DeployCommitteeVerifier, chain, DeployCommitteeVerifierInput{
				ChainSelector:     chain.Selector,
				CREATE2Factory:    input.CREATE2Factory,
				ExistingAddresses: input.ExistingAddresses,
				Params:            committeeVerifierParams,
				RMN:               common.HexToAddress(rmnProxyRef.Address),
			})
			if err != nil {
				return output, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
			}
			addresses = append(addresses, report.Output.Addresses...)
			ownableContracts = append(ownableContracts, report.Output.Addresses...)
			committeeVerifierBatchOps = append(committeeVerifierBatchOps, report.Output.BatchOps...)
		}

		// Deploy Executors
		for _, executorParam := range input.ContractParams.Executors {
			var qualifierPtr *string
			if executorParam.Qualifier != "" {
				qualifierPtr = &executorParam.Qualifier
			}
			executorRef, err := evmops.MaybeDeployContract(b, executor.Deploy, chain, contract.DeployInput[executor.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(executor.ContractType, *executorParam.Version),
				Args: executor.ConstructorArgs{
					MaxCCVsPerMsg: executorParam.MaxCCVsPerMsg,
					DynamicConfig: executorParam.DynamicConfig,
				},
				Qualifier: qualifierPtr,
			}, input.ExistingAddresses)
			if err != nil {
				return output, fmt.Errorf("failed to deploy Executor: %w, params: %+v", err, executorParam)
			}
			addresses = append(addresses, executorRef)
			ownableContracts = append(ownableContracts, executorRef)

			// Fetch the dynamic config on the Executor
			dynamicConfigReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(executorRef.Address), evmops.BindAs[executor_bindings.ExecutorInterface](executor_bindings.NewExecutor), executor.NewReadGetDynamicConfig, struct{}{})
			if err != nil {
				return output, fmt.Errorf("failed to get dynamic config on Executor: %w", err)
			}

			// Set dynamic config on the Executor if diff exists
			desiredFeeAggregator := dynamicConfigReport.Output.FeeAggregator
			if executorParam.DynamicConfig.FeeAggregator != (common.Address{}) {
				desiredFeeAggregator = executorParam.DynamicConfig.FeeAggregator
			}
			if desiredFeeAggregator != dynamicConfigReport.Output.FeeAggregator ||
				dynamicConfigReport.Output.AllowedFinalityConfig != executorParam.DynamicConfig.AllowedFinalityConfig ||
				dynamicConfigReport.Output.CcvAllowlistEnabled != executorParam.DynamicConfig.CcvAllowlistEnabled {
				setDynamicConfigReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(executorRef.Address), evmops.BindAs[executor_bindings.ExecutorInterface](executor_bindings.NewExecutor), executor.NewWriteSetDynamicConfig, executor_bindings.ExecutorDynamicConfig{
					FeeAggregator:         executorParam.DynamicConfig.FeeAggregator,
						AllowedFinalityConfig: executorParam.DynamicConfig.AllowedFinalityConfig,
						CcvAllowlistEnabled:   executorParam.DynamicConfig.CcvAllowlistEnabled,
				})
				if err != nil {
					return output, fmt.Errorf("failed to set dynamic config on Executor: %w", err)
				}
				writes = append(writes, setDynamicConfigReport.Output)
			}

			// Deploy ExecutorProxy via CREATE2
			var executorProxyRef *datastore.AddressRef
			for _, ref := range input.ExistingAddresses {
				if ref.Type == datastore.ContractType(ExecutorProxyType) &&
					ref.Version.String() == executor.Version.String() &&
					(qualifierPtr == nil || ref.Qualifier == *qualifierPtr) {
					executorProxyRef = &ref
				}
			}
			if executorProxyRef != nil {
				addresses = append(addresses, *executorProxyRef)
			} else {
				if input.CREATE2Factory == (common.Address{}) {
					return output, fmt.Errorf("CREATE2Factory is required to deploy ExecutorProxy")
				}
				deployExecutorProxyViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, DeployContractViaCREATE2, chain, DeployContractViaCREATE2Input{
					Qualifier:      *qualifierPtr,
					Type:           datastore.ContractType(ExecutorProxyType),
					Version:        executor.Version,
					CREATE2Factory: input.CREATE2Factory,
					ABI:            proxy_bindings.ProxyMetaData.ABI,
					BIN:            proxy_bindings.ProxyMetaData.Bin,
					ConstructorArgs: []any{
						// To ensure consistent addresses, we have to deploy with the same constructor args on every chain.
						// Instead of setting in the constructor, we set the target and fee aggregator after deployment.
						common.HexToAddress("0x01"), // Target (will revert if target is 0, so we use a dummy address)
						common.Address{},            // Fee Aggregator
					},
				})
				if err != nil {
					return output, fmt.Errorf("failed to deploy ExecutorProxy: %w", err)
				}
				addresses = append(addresses, deployExecutorProxyViaCREATE2Report.Output.Addresses...)
				writes = append(writes, deployExecutorProxyViaCREATE2Report.Output.Writes...)

				if len(deployExecutorProxyViaCREATE2Report.Output.Addresses) != 1 {
					return output, fmt.Errorf("expected 1 ExecutorProxy address, got %d", len(deployExecutorProxyViaCREATE2Report.Output.Addresses))
				}
				executorProxyRef = &deployExecutorProxyViaCREATE2Report.Output.Addresses[0]

				// Accept ownership of the ExecutorProxy
				acceptOwnershipReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(executorProxyRef.Address), evmops.BindAs[proxy_bindings.ProxyInterface](proxy_bindings.NewProxy), newWriteProxyAcceptOwnership, proxyAcceptOwnershipArgs{
					IsProposedOwner: true,
				})
				if err != nil {
					return output, fmt.Errorf("failed to accept ownership of ExecutorProxy: %w", err)
				}
				writes = append(writes, acceptOwnershipReport.Output)
			}
			ownableContracts = append(ownableContracts, *executorProxyRef)

			// Fetch the target on the ExecutorProxy
			targetReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(executorProxyRef.Address), evmops.BindAs[proxy_bindings.ProxyInterface](proxy_bindings.NewProxy), proxy.NewReadGetTarget, struct{}{})
			if err != nil {
				return output, fmt.Errorf("failed to get target on ExecutorProxy: %w", err)
			}

			// Set target on the ExecutorProxy if diff exists
			if targetReport.Output != common.HexToAddress(executorRef.Address) {
				setTargetReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(executorProxyRef.Address), evmops.BindAs[proxy_bindings.ProxyInterface](proxy_bindings.NewProxy), proxy.NewWriteSetTarget, common.HexToAddress(executorRef.Address))
				if err != nil {
					return output, fmt.Errorf("failed to set target on ExecutorProxy: %w", err)
				}
				writes = append(writes, setTargetReport.Output)
			}

			// Fetch the fee aggregator on the ExecutorProxy
			feeAggregatorReport, err := evmops.ExecuteRead(b, chain, common.HexToAddress(executorProxyRef.Address), evmops.BindAs[proxy_bindings.ProxyInterface](proxy_bindings.NewProxy), proxy.NewReadGetFeeAggregator, struct{}{})
			if err != nil {
				return output, fmt.Errorf("failed to get fee aggregator on ExecutorProxy: %w", err)
			}

			// Set fee aggregator on the ExecutorProxy if diff exists
			if feeAggregatorReport.Output != executorParam.DynamicConfig.FeeAggregator {
				setFeeAggregatorReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(executorProxyRef.Address), evmops.BindAs[proxy_bindings.ProxyInterface](proxy_bindings.NewProxy), proxy.NewWriteSetFeeAggregator, executorParam.DynamicConfig.FeeAggregator)
				if err != nil {
					return output, fmt.Errorf("failed to set fee aggregator on ExecutorProxy: %w", err)
				}
				writes = append(writes, setFeeAggregatorReport.Output)
			}
		}

		for _, mockReceiverParams := range input.ContractParams.MockReceivers {
			requiredVerifiers, optionalVerifiers, err := getMockReceiverVerifiers(mockReceiverParams, addresses, input.ExistingAddresses)
			if err != nil {
				return output, fmt.Errorf("failed to get mock receiver verifiers: %w", err)
			}
			var qualifierPtr *string
			if mockReceiverParams.Qualifier != "" {
				qualifierPtr = &mockReceiverParams.Qualifier
			}
			deployReceiverReport, err := evmops.ExecuteDeploy(b, mock_receiver.Deploy, chain, contract.DeployInput[mock_receiver.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(mock_receiver.ContractType, *mockReceiverParams.Version),
				Args: mock_receiver.ConstructorArgs{
					Required:  requiredVerifiers,
					Optional:  optionalVerifiers,
					Threshold: mockReceiverParams.OptionalThreshold,
				},
				Qualifier: qualifierPtr,
			})
			if err != nil {
				return output, fmt.Errorf("failed to deploy MockReceiver: %w", err)
			}
			addresses = append(addresses, deployReceiverReport.Output)

			// Set finality config on the MockReceiver if diff exists
			if mockReceiverParams.AllowedFinalityConfig != finality.RawWaitForFinality {
				// Get the current finality config on the MockReceiver
				finalityConfigResult, err := evmops.ExecuteRead(b, chain, common.HexToAddress(deployReceiverReport.Output.Address), evmops.BindAs[mock_receiver_v2_bindings.MockReceiverV2Interface](mock_receiver_v2_bindings.NewMockReceiverV2), mock_receiver_v2.NewReadGetCCVsAndFinalityConfig, mock_receiver_v2.GetCCVsAndFinalityConfigArgs{
					Arg0: chain.Selector,
					Arg1: []byte{},
				})
				if err != nil {
					return output, fmt.Errorf("failed to get finality config on MockReceiver: %w", err)
				}
				if finalityConfigResult.Output.AllowedFinalityConfig != mockReceiverParams.AllowedFinalityConfig {
					// Set the finality config on the MockReceiver
					setFinalityConfigReport, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(deployReceiverReport.Output.Address), evmops.BindAs[mock_receiver_v2_bindings.MockReceiverV2Interface](mock_receiver_v2_bindings.NewMockReceiverV2), mock_receiver_v2.NewWriteSetAllowedFinalityConfig, mockReceiverParams.AllowedFinalityConfig)
					if err != nil {
						return output, fmt.Errorf("failed to set finality config on MockReceiver: %w", err)
					}
					writes = append(writes, setFinalityConfigReport.Output)
				}
			}
		}

		// Transfer ownership of MCM and product contracts to the CLLCCIP timelock.
		// MCM contracts (Proposer, Bypasser, Canceller) were added to ownableContracts
		// during the MCMS validation above; product contracts are added here.
		if cllccipTimelockAddr != (common.Address{}) {
			// Ensure both timelocks are self-governed: the CLLCCIP timelock should be
			// admin of both itself and the RMNMCMS timelock. If the deployer still holds
			// the admin role on either, fix it now so incomplete earlier setups are healed.
			if err := ensureTimelockSelfGoverned(b, chain, cllccipTimelockAddr, cllccipTimelockAddr, []common.Address{cllccipTimelockAddr}); err != nil {
				return output, fmt.Errorf("failed to ensure CLLCCIP timelock is self-governed: %w", err)
			}
			if err := ensureTimelockSelfGoverned(b, chain, rmnTimelockAddr, cllccipTimelockAddr, []common.Address{cllccipTimelockAddr, rmnTimelockAddr}); err != nil {
				return output, fmt.Errorf("failed to ensure RMNMCMS timelock is governed by CLLCCIP timelock: %w", err)
			}
		}

		var batchOps []mcms_types.BatchOperation
		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return output, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		batchOps = append(batchOps, batchOp)
		batchOps = append(batchOps, committeeVerifierBatchOps...)
		output.OnChainOutput = sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}
		if !input.DeployerKeyOwned {
			output.RefsToTransferOwnership, err = filterContractsNeedingOwnershipTransfer(chain, ownableContracts, cllccipTimelockAddr, rmnTimelockAddr)
			if err != nil {
				return output, fmt.Errorf("failed to filter ownable contracts: %w", err)
			}
		}
		return output, nil
	},
)

// ResolveOwnershipDeps looks up the MCMS contracts required for ownership
// transfer from existingAddresses. It returns the CLL and RMN timelock
// addresses together with the MCM contracts (Proposer, Bypasser, Canceller)
// wrapped as ownableContracts so the caller can include them in the
// transfer-ownership pass.
func ResolveOwnershipDeps(
	existingAddresses []datastore.AddressRef,
	chainSelector uint64,
) (cllccipTimelockAddr, rmnTimelockAddr common.Address, mcmContracts []datastore.AddressRef, err error) {
	existingDS := datastore.NewMemoryDataStore()
	for _, ref := range existingAddresses {
		if addErr := existingDS.Addresses().Add(ref); addErr != nil && !errors.Is(addErr, datastore.ErrAddressRefExists) {
			return common.Address{}, common.Address{}, nil,
				fmt.Errorf("failed to add existing address to datastore (%+v): %w", ref, addErr)
		}
	}
	mcmsDS := existingDS.Seal()

	cllccipTimelockAddr, err = datastore_utils.FindAndFormatRef(mcmsDS, datastore.AddressRef{
		Type:      datastore.ContractType(common_utils.RBACTimelock),
		Qualifier: common_utils.CLLQualifier,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return common.Address{}, common.Address{}, nil,
			fmt.Errorf("ownership transfer requires CLLCCIP RBACTimelock in ExistingAddresses: %w", err)
	}

	rmnTimelockAddr, err = datastore_utils.FindAndFormatRef(mcmsDS, datastore.AddressRef{
		Type:      datastore.ContractType(common_utils.RBACTimelock),
		Qualifier: common_utils.RMNTimelockQualifier,
	}, chainSelector, evm_datastore_utils.ToEVMAddress)
	if err != nil {
		return common.Address{}, common.Address{}, nil,
			fmt.Errorf("ownership transfer requires RMNMCMS RBACTimelock in ExistingAddresses: %w", err)
	}

	// Look up MCM contracts (Proposer, Bypasser, Canceller) for both qualifiers
	// so we can transfer their ownership to the CLL timelock.
	// RMN MCM contracts are also acceptably owned by the RMN timelock (since
	// proposals target a single MCMS instance); CLL MCMs must be owned by the
	// CLL timelock exclusively.
	mcmTypes := []deployment.ContractType{
		common_utils.ProposerManyChainMultisig,
		common_utils.BypasserManyChainMultisig,
		common_utils.CancellerManyChainMultisig,
	}
	requiredQualifiers := []string{
		common_utils.CLLQualifier,
		common_utils.RMNTimelockQualifier,
	}
	for _, ct := range mcmTypes {
		for _, qualifier := range requiredQualifiers {
			addresses := mcmsDS.Addresses().Filter(
				datastore.AddressRefByType(datastore.ContractType(ct)),
				datastore.AddressRefByQualifier(qualifier),
				datastore.AddressRefByChainSelector(chainSelector),
			)
			if len(addresses) == 0 {
				return common.Address{}, common.Address{}, nil,
					fmt.Errorf("ownership transfer requires MCM contract of type %s with qualifier %s in ExistingAddresses", ct, qualifier)
			}
			mcmContracts = append(mcmContracts, addresses...)
		}
	}

	return cllccipTimelockAddr, rmnTimelockAddr, mcmContracts, nil
}

// filterContractsNeedingOwnershipTransfer filters the provided datastore refs based on current ownership status
// it returns only the refs which would require ownership transfer to timelock and skips the one which are already owned by timelock
func filterContractsNeedingOwnershipTransfer(
	chain evm.Chain,
	refs []datastore.AddressRef,
	cllccipTimelockAddr, rmnTimelockAddr common.Address,
) ([]datastore.AddressRef, error) {
	var filtered []datastore.AddressRef
	for _, ref := range refs {
		currentOwner, _, err := mcms_seq.LoadOwnableContract(common.HexToAddress(ref.Address), chain.Client)
		if err != nil {
			return nil, fmt.Errorf("failed to load ownable contract %s (%s): %w", ref.Address, ref.Type, err)
		}
		switch ref.Type {
		case datastore.ContractType(rmn_remote.ContractType):
			if currentOwner != rmnTimelockAddr {
				filtered = append(filtered, ref)
			}
		case datastore.ContractType(common_utils.ProposerManyChainMultisig), datastore.ContractType(common_utils.BypasserManyChainMultisig), datastore.ContractType(common_utils.CancellerManyChainMultisig):
			if ref.Qualifier == common_utils.CLLQualifier {
				if currentOwner != cllccipTimelockAddr {
					filtered = append(filtered, ref)
				}
			} else if ref.Qualifier == common_utils.RMNTimelockQualifier {
				if currentOwner != rmnTimelockAddr {
					filtered = append(filtered, ref)
				}
			} else {
				if currentOwner != cllccipTimelockAddr && currentOwner != rmnTimelockAddr {
					filtered = append(filtered, ref)
				}
			}
		default:
			if currentOwner != cllccipTimelockAddr {
				filtered = append(filtered, ref)
			}
		}
	}
	return filtered, nil
}

// ensureTimelockSelfGoverned checks that newAdmin holds ADMIN_ROLE on the given
// timelock. If the deployer still has admin, the function grants the role to newAdmin
// and renounces the deployer's admin. This is idempotent: if newAdmin is already admin
// and the deployer is not, the function is a no-op.
func ensureTimelockSelfGoverned(
	b cldf_ops.Bundle,
	chain evm.Chain,
	timelockAddr common.Address,
	newAdmin common.Address,
	acceptableAdmins []common.Address,
) error {
	timelock, err := mcms_seq.LoadTimelockContract(timelockAddr, chain.Client)
	if err != nil {
		return fmt.Errorf("failed to load timelock contract %s: %w", timelockAddr, err)
	}

	var adminHasRole bool
	for _, acceptableAdmin := range acceptableAdmins {
		adminHasRole, err = timelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, acceptableAdmin)
		if err != nil {
			return fmt.Errorf("failed to check admin role for acceptable admin %s on timelock %s: %w", acceptableAdmin, timelockAddr, err)
		}
		if adminHasRole {
			b.Logger.Infof("Timelock %s is already governed by acceptable admin %s", timelockAddr, acceptableAdmin)
			break
		}
	}

	deployerHasRole, err := timelock.HasRole(nil, mcms_ops.ADMIN_ROLE.ID, chain.DeployerKey.From)
	if err != nil {
		return fmt.Errorf("failed to check admin role for deployer %s on timelock %s: %w", chain.DeployerKey.From, timelockAddr, err)
	}

	if adminHasRole && !deployerHasRole {
		b.Logger.Infof("Timelock %s is already governed by %s, deployer is not admin — no action needed", timelockAddr, newAdmin)
		return nil
	}

	if !deployerHasRole {
		return fmt.Errorf(
			"timelock %s: new admin %s does not have ADMIN_ROLE and deployer %s does not have ADMIN_ROLE; cannot fix governance",
			timelockAddr, newAdmin, chain.DeployerKey.From,
		)
	}

	if !adminHasRole {
		if err := mcms_seq.GrantTimelockAdminRole(b, chain, chain.Selector, timelockAddr, newAdmin); err != nil {
			return err
		}
	}

	return mcms_seq.RenounceDeployerTimelockAdmin(b, chain, chain.Selector, timelockAddr)
}

// getMockReceiverVerifiers finds the required and optional verifier addresses given the mock receiver
// params, the addresses of the newly deployed contracts, and the addresses of the existing contracts.
func getMockReceiverVerifiers(
	mockReceiverParams MockReceiverParams,
	addresses []datastore.AddressRef,
	existingAddresses []datastore.AddressRef,
) (requiredVerifiers []common.Address, optionalVerifiers []common.Address, err error) {
	// create a datastore in order to use the API.
	ds := datastore.NewMemoryDataStore()
	for _, ref := range append(addresses, existingAddresses...) {
		// We ignore ErrAddressRefExists since a partial deployment could result in an address being in both
		// addresses and existingAddresses due to the idempotent nature of the sequence.
		if err := ds.Addresses().Add(ref); err != nil && !errors.Is(err, datastore.ErrAddressRefExists) {
			return nil, nil, fmt.Errorf("failed to add address to datastore (%+v): %w", ref, err)
		}
	}
	sealed := ds.Seal()
	for _, ref := range mockReceiverParams.RequiredVerifiers {
		address, err := datastore_utils.FindAndFormatRef(sealed, ref, ref.ChainSelector, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to find required verifier (%+v) in datastore containing existing and newly deployed addresses: %w",
				ref,
				err)
		}
		requiredVerifiers = append(requiredVerifiers, address)
	}
	for _, ref := range mockReceiverParams.OptionalVerifiers {
		address, err := datastore_utils.FindAndFormatRef(sealed, ref, ref.ChainSelector, evm_datastore_utils.ToEVMAddress)
		if err != nil {
			return nil, nil, fmt.Errorf(
				"failed to find optional verifier (%+v) in datastore containing existing and newly deployed addresses: %w",
				ref,
				err)
		}
		optionalVerifiers = append(optionalVerifiers, address)
	}
	if len(requiredVerifiers) != len(mockReceiverParams.RequiredVerifiers) {
		return nil, nil, fmt.Errorf("not all required verifiers found, got %d (%+v), expected %d (%+v)",
			len(requiredVerifiers),
			requiredVerifiers,
			len(mockReceiverParams.RequiredVerifiers),
			mockReceiverParams.RequiredVerifiers)
	}
	if len(optionalVerifiers) != len(mockReceiverParams.OptionalVerifiers) {
		return nil, nil, fmt.Errorf("not all optional verifiers found, got %d (%+v), expected %d (%+v)",
			len(optionalVerifiers),
			optionalVerifiers,
			len(mockReceiverParams.OptionalVerifiers),
			mockReceiverParams.OptionalVerifiers)
	}
	return requiredVerifiers, optionalVerifiers, nil
}

// MaybeRegisterModuleOnTokenAdminRegistry checks if a module is already registered on the TokenAdminRegistry,
// and if not, adds it as a registry module.
// Returns the write output and a boolean indicating whether a write operation was performed.
func MaybeRegisterModuleOnTokenAdminRegistry(
	b cldf_ops.Bundle,
	chain evm.Chain,
	tokenAdminRegistryAddress common.Address,
	moduleAddress common.Address,
) (contract.WriteOutput, bool, error) {
	// Check if the module is already registered.
	isRegisteredReport, err := evmops.ExecuteRead(b, chain, tokenAdminRegistryAddress, tarbindings.NewTokenAdminRegistry, token_admin_registry.NewReadIsRegistryModule, moduleAddress)
	if err != nil {
		return contract.WriteOutput{}, false, fmt.Errorf("failed to check if module is registered: %w", err)
	}

	// If already registered, return without performing a write.
	if isRegisteredReport.Output {
		return contract.WriteOutput{}, false, nil
	}

	// Add the module to the registry.
	addRegistryModuleReport, err := evmops.ExecuteWrite(b, chain, tokenAdminRegistryAddress, tarbindings.NewTokenAdminRegistry, token_admin_registry.NewWriteAddRegistryModule, moduleAddress)
	if err != nil {
		return contract.WriteOutput{}, false, fmt.Errorf("failed to add registry module: %w", err)
	}

	return addRegistryModuleReport.Output, true, nil
}
