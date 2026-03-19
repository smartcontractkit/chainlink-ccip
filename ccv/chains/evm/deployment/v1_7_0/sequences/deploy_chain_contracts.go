package sequences

import (
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

	proxy_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/mock_receiver_v2"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/proxy"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/registry_module_owner_custom"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type proxyAcceptOwnershipArgs struct {
	IsProposedOwner bool
}

var proxyAcceptOwnership = contract_utils.NewWrite(contract_utils.WriteParams[proxyAcceptOwnershipArgs, *proxy_bindings.Proxy]{
	Name:         "proxy:accept-ownership",
	Version:      proxy.Version,
	Description:  "Accept ownership of the proxy",
	ContractType: proxy.ContractType,
	ContractABI:  proxy.ProxyABI,
	NewContract:  proxy_bindings.NewProxy,
	IsAllowedCaller: func(_ *proxy_bindings.Proxy, _ *bind.CallOpts, _ common.Address, args proxyAcceptOwnershipArgs) (bool, error) {
		return args.IsProposedOwner, nil
	},
	Validate: func(proxyAcceptOwnershipArgs) error { return nil },
	CallContract: func(p *proxy_bindings.Proxy, opts *bind.TransactOpts, _ proxyAcceptOwnershipArgs) (*types.Transaction, error) {
		return p.AcceptOwnership(opts)
	},
})

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
	// MinimumBlockConfirmations is the minimum block depth that the mock receiver will accept.
	MinimumBlockConfirmations uint16
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
	DynamicConfig executor.DynamicConfig
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
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployChainContractsInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		ownableContracts := make([]ownableContract, 0)

		var cllccipTimelockAddr, rmnTimelockAddr common.Address
		if !input.DeployerKeyOwned {
			var mcmContracts []ownableContract
			cllccipTimelockAddr, rmnTimelockAddr, mcmContracts, err = ResolveOwnershipDeps(
				input.ExistingAddresses, chain.Selector,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to resolve ownership dependencies: %w", err)
			}
			ownableContracts = append(ownableContracts, mcmContracts...)
		}

		// Deploy WETH
		wethRef, err := contract_utils.MaybeDeployContract(b, weth.Deploy, chain, contract_utils.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(weth.ContractType, *weth.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := contract_utils.MaybeDeployContract(b, link.Deploy, chain, contract_utils.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link.ContractType, *link.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := contract_utils.MaybeDeployContract(b, rmn_remote.Deploy, chain, contract_utils.DeployInput[rmn_remote.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_remote.ContractType, *input.ContractParams.RMNRemote.Version),
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
		rmnProxyRef, err := contract_utils.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract_utils.DeployInput[rmn_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
			ChainSelector:  chain.Selector,
			Args: rmn_proxy.ConstructorArgs{
				RMN: common.HexToAddress(rmnRemoteRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, rmnProxyRef)

		// Fetch the RMN contract address set on the RMNProxy
		rmnAddressReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.GetRMN, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(rmnProxyRef.Address),
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Set the RMNRemote on the RMNProxy if diff exists
		if rmnAddressReport.Output != common.HexToAddress(rmnRemoteRef.Address) {
			setRMNReport, err := cldf_ops.ExecuteOperation(b, rmn_proxy.SetRMN, chain, contract_utils.FunctionInput[rmn_proxy.SetRMNArgs]{
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
		}

		// Deploy Router
		routerRef, err := contract_utils.MaybeDeployContract(b, router.Deploy, chain, contract_utils.DeployInput[router.ConstructorArgs]{
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

		// Fetch the wrapped native address set on the Router
		wrappedNativeAddressReport, err := cldf_ops.ExecuteOperation(b, router.GetWrappedNative, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(routerRef.Address),
		})
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Set wrapped native on the Router if diff exists
		if wrappedNativeAddressReport.Output != common.HexToAddress(wethRef.Address) {
			setWrappedNativeReport, err := cldf_ops.ExecuteOperation(b, router.SetWrappedNative, chain, contract_utils.FunctionInput[common.Address]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(routerRef.Address),
				Args:          common.HexToAddress(wethRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			writes = append(writes, setWrappedNativeReport.Output)
		}

		// Deploy Test Router
		if input.DeployTestRouter {
			testRouterRef, err := contract_utils.MaybeDeployContract(b, router.DeployTestRouter, chain, contract_utils.DeployInput[router.ConstructorArgs]{
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

			// Fetch the wrapped native address set on the Test Router
			wrappedNativeAddressReport, err = cldf_ops.ExecuteOperation(b, router.GetWrappedNative, chain, contract_utils.FunctionInput[struct{}]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(testRouterRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			// Set wrapped native on the Test Router if diff exists
			if wrappedNativeAddressReport.Output != common.HexToAddress(wethRef.Address) {
				setWrappedNativeReport, err := cldf_ops.ExecuteOperation(b, router.SetWrappedNative, chain, contract_utils.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(testRouterRef.Address),
					Args:          common.HexToAddress(wethRef.Address),
				})
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				writes = append(writes, setWrappedNativeReport.Output)
			}
		}

		// Deploy TokenAdminRegistry
		tokenAdminRegistryRef, err := contract_utils.MaybeDeployContract(b, token_admin_registry.Deploy, chain, contract_utils.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy RegistryModuleOwnerCustom
		registryModuleOwnerCustomRef, err := contract_utils.MaybeDeployContract(b, registry_module_owner_custom.Deploy, chain, contract_utils.DeployInput[registry_module_owner_custom.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(registry_module_owner_custom.ContractType, *registry_module_owner_custom.Version),
			ChainSelector:  chain.Selector,
			Args: registry_module_owner_custom.ConstructorArgs{
				TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
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
			return sequences.OnChainOutput{}, err
		}
		// Only append to writes if a transaction was actually created (i.e., module wasn't already registered).
		if hasOnchainDiff {
			writes = append(writes, addRegistryModuleReport)
		}

		// Deploy FeeQuoter
		priceUpdaters := []common.Address{chain.DeployerKey.From}
		if cllccipTimelockAddr != (common.Address{}) {
			priceUpdaters = append(priceUpdaters, cllccipTimelockAddr)
		}
		feeQuoterRef, err := contract_utils.MaybeDeployContract(b, fee_quoter.Deploy, chain, contract_utils.DeployInput[fee_quoter.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(fee_quoter.ContractType, *input.ContractParams.FeeQuoter.Version),
			ChainSelector:  chain.Selector,
			Args: fee_quoter.ConstructorArgs{
				StaticConfig: fee_quoter.StaticConfig{
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
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, feeQuoterRef)

		var tokenPriceUpdates []fee_quoter.TokenPriceUpdate
		if input.ContractParams.FeeQuoter.USDPerLINK != nil {
			tokenPriceUpdates = append(tokenPriceUpdates, fee_quoter.TokenPriceUpdate{
				SourceToken: common.HexToAddress(linkRef.Address),
				UsdPerToken: input.ContractParams.FeeQuoter.USDPerLINK,
			})
		}
		if input.ContractParams.FeeQuoter.USDPerWETH != nil {
			tokenPriceUpdates = append(tokenPriceUpdates, fee_quoter.TokenPriceUpdate{
				SourceToken: common.HexToAddress(wethRef.Address),
				UsdPerToken: input.ContractParams.FeeQuoter.USDPerWETH,
			})
		}
		if len(tokenPriceUpdates) > 0 {
			updatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract_utils.FunctionInput[fee_quoter.PriceUpdates]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(feeQuoterRef.Address),
				Args: fee_quoter.PriceUpdates{
					TokenPriceUpdates: tokenPriceUpdates,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to update token prices on FeeQuoter: %w", err)
			}
			writes = append(writes, updatePricesReport.Output)
		}

		// Deploy OffRamp
		offRampRef, err := contract_utils.MaybeDeployContract(b, offramp.Deploy, chain, contract_utils.DeployInput[offramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(offramp.ContractType, *input.ContractParams.OffRamp.Version),
			ChainSelector:  chain.Selector,
			Args: offramp.ConstructorArgs{
				StaticConfig: offramp.StaticConfig{
					LocalChainSelector:        chain.Selector,
					RmnRemote:                 common.HexToAddress(rmnProxyRef.Address),
					GasForCallExactCheck:      input.ContractParams.OffRamp.GasForCallExactCheck,
					TokenAdminRegistry:        common.HexToAddress(tokenAdminRegistryRef.Address),
					MaxGasBufferToUpdateState: input.ContractParams.OffRamp.MaxGasBufferToUpdateState,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OffRamp: %w", err)
		}
		addresses = append(addresses, offRampRef)

		// Deploy OnRamp
		onRampRef, err := contract_utils.MaybeDeployContract(b, onramp.Deploy, chain, contract_utils.DeployInput[onramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(onramp.ContractType, *input.ContractParams.OnRamp.Version),
			ChainSelector:  chain.Selector,
			Args: onramp.ConstructorArgs{
				StaticConfig: onramp.StaticConfig{
					ChainSelector:         chain.Selector,
					RmnRemote:             common.HexToAddress(rmnRemoteRef.Address),
					TokenAdminRegistry:    common.HexToAddress(tokenAdminRegistryRef.Address),
					MaxUSDCentsPerMessage: input.ContractParams.OnRamp.MaxUSDCentsPerMessage,
				},
				DynamicConfig: onramp.DynamicConfig{
					FeeQuoter:     common.HexToAddress(feeQuoterRef.Address),
					FeeAggregator: input.ContractParams.OnRamp.FeeAggregator,
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OnRamp: %w", err)
		}
		addresses = append(addresses, onRampRef)

		// Fetch the dynamic config on the OnRamp
		dynamicConfigReport, err := cldf_ops.ExecuteOperation(b, onramp.GetDynamicConfig, chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(onRampRef.Address),
			Args:          struct{}{},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config on OnRamp: %w", err)
		}

		// Set dynamic config on the OnRamp if there is a diff
		desiredFeeAggregator := dynamicConfigReport.Output.FeeAggregator
		if input.ContractParams.OnRamp.FeeAggregator != (common.Address{}) {
			desiredFeeAggregator = input.ContractParams.OnRamp.FeeAggregator
		}
		if dynamicConfigReport.Output.FeeQuoter != common.HexToAddress(feeQuoterRef.Address) || desiredFeeAggregator != dynamicConfigReport.Output.FeeAggregator {
			desiredDynamicConfig := onramp.DynamicConfig{
				FeeQuoter:              common.HexToAddress(feeQuoterRef.Address),
				ReentrancyGuardEntered: false, // This should never be true.
				FeeAggregator:          input.ContractParams.OnRamp.FeeAggregator,
			}
			setDynamicConfigReport, err := cldf_ops.ExecuteOperation(b, onramp.SetDynamicConfig, chain, contract_utils.FunctionInput[onramp.DynamicConfig]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(onRampRef.Address),
				Args:          desiredDynamicConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set dynamic config on OnRamp: %w", err)
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
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
			}
			addresses = append(addresses, report.Output.Addresses...)
			committeeVerifierBatchOps = append(committeeVerifierBatchOps, report.Output.BatchOps...)
		}

		// Deploy Executors
		for _, executorParam := range input.ContractParams.Executors {
			var qualifierPtr *string
			if executorParam.Qualifier != "" {
				qualifierPtr = &executorParam.Qualifier
			}
			executorRef, err := contract_utils.MaybeDeployContract(b, executor.Deploy, chain, contract_utils.DeployInput[executor.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(executor.ContractType, *executorParam.Version),
				ChainSelector:  chain.Selector,
				Args: executor.ConstructorArgs{
					MaxCCVsPerMsg: executorParam.MaxCCVsPerMsg,
					DynamicConfig: executorParam.DynamicConfig,
				},
				Qualifier: qualifierPtr,
			}, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy Executor: %w, params: %+v", err, executorParam)
			}
			addresses = append(addresses, executorRef)
			ownableContracts = append(ownableContracts, ownableContract{common.HexToAddress(executorRef.Address), executor.ContractType, []common.Address{cllccipTimelockAddr}})

			// Fetch the dynamic config on the Executor
			dynamicConfigReport, err := cldf_ops.ExecuteOperation(b, executor.GetDynamicConfig, chain, contract_utils.FunctionInput[struct{}]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(executorRef.Address),
				Args:          struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config on Executor: %w", err)
			}

			// Set dynamic config on the Executor if diff exists
			desiredFeeAggregator := dynamicConfigReport.Output.FeeAggregator
			if executorParam.DynamicConfig.FeeAggregator != (common.Address{}) {
				desiredFeeAggregator = executorParam.DynamicConfig.FeeAggregator
			}
			if desiredFeeAggregator != dynamicConfigReport.Output.FeeAggregator ||
				dynamicConfigReport.Output.MinBlockConfirmations != executorParam.DynamicConfig.MinBlockConfirmations ||
				dynamicConfigReport.Output.CcvAllowlistEnabled != executorParam.DynamicConfig.CcvAllowlistEnabled {
				setDynamicConfigReport, err := cldf_ops.ExecuteOperation(b, executor.SetDynamicConfig, chain, contract_utils.FunctionInput[executor.DynamicConfig]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(executorRef.Address),
					Args: executor.DynamicConfig{
						FeeAggregator:         executorParam.DynamicConfig.FeeAggregator,
						MinBlockConfirmations: executorParam.DynamicConfig.MinBlockConfirmations,
						CcvAllowlistEnabled:   executorParam.DynamicConfig.CcvAllowlistEnabled,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set dynamic config on Executor: %w", err)
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
					return sequences.OnChainOutput{}, fmt.Errorf("CREATE2Factory is required to deploy ExecutorProxy")
				}
				deployExecutorProxyViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, DeployContractViaCREATE2, chain, DeployContractViaCREATE2Input{
					ChainSelector:  chain.Selector,
					Qualifier:      *qualifierPtr,
					Type:           datastore.ContractType(ExecutorProxyType),
					Version:        executor.Version,
					CREATE2Factory: input.CREATE2Factory,
					ABI:            proxy.ProxyABI,
					BIN:            proxy.ProxyBin,
					ConstructorArgs: []any{
						// To ensure consistent addresses, we have to deploy with the same constructor args on every chain.
						// Instead of setting in the constructor, we set the target and fee aggregator after deployment.
						common.HexToAddress("0x01"), // Target (will revert if target is 0, so we use a dummy address)
						common.Address{},            // Fee Aggregator
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ExecutorProxy: %w", err)
				}
				addresses = append(addresses, deployExecutorProxyViaCREATE2Report.Output.Addresses...)
				writes = append(writes, deployExecutorProxyViaCREATE2Report.Output.Writes...)

				if len(deployExecutorProxyViaCREATE2Report.Output.Addresses) != 1 {
					return sequences.OnChainOutput{}, fmt.Errorf("expected 1 ExecutorProxy address, got %d", len(deployExecutorProxyViaCREATE2Report.Output.Addresses))
				}
				executorProxyRef = &deployExecutorProxyViaCREATE2Report.Output.Addresses[0]

				// Accept ownership of the ExecutorProxy
				acceptOwnershipReport, err := cldf_ops.ExecuteOperation(b, proxyAcceptOwnership, chain, contract_utils.FunctionInput[proxyAcceptOwnershipArgs]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(executorProxyRef.Address),
					Args: proxyAcceptOwnershipArgs{
						IsProposedOwner: true,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership of ExecutorProxy: %w", err)
				}
				writes = append(writes, acceptOwnershipReport.Output)
			}
			ownableContracts = append(ownableContracts, ownableContract{common.HexToAddress(executorProxyRef.Address), ExecutorProxyType, []common.Address{cllccipTimelockAddr}})

			// Fetch the target on the ExecutorProxy
			targetReport, err := cldf_ops.ExecuteOperation(b, proxy.GetTarget, chain, contract_utils.FunctionInput[struct{}]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(executorProxyRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get target on ExecutorProxy: %w", err)
			}

			// Set target on the ExecutorProxy if diff exists
			if targetReport.Output != common.HexToAddress(executorRef.Address) {
				setTargetReport, err := cldf_ops.ExecuteOperation(b, proxy.SetTarget, chain, contract_utils.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(executorProxyRef.Address),
					Args:          common.HexToAddress(executorRef.Address),
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set target on ExecutorProxy: %w", err)
				}
				writes = append(writes, setTargetReport.Output)
			}

			// Fetch the fee aggregator on the ExecutorProxy
			feeAggregatorReport, err := cldf_ops.ExecuteOperation(b, proxy.GetFeeAggregator, chain, contract_utils.FunctionInput[struct{}]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(executorProxyRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get fee aggregator on ExecutorProxy: %w", err)
			}

			// Set fee aggregator on the ExecutorProxy if diff exists
			if feeAggregatorReport.Output != executorParam.DynamicConfig.FeeAggregator {
				setFeeAggregatorReport, err := cldf_ops.ExecuteOperation(b, proxy.SetFeeAggregator, chain, contract_utils.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(executorProxyRef.Address),
					Args:          executorParam.DynamicConfig.FeeAggregator,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set fee aggregator on ExecutorProxy: %w", err)
				}
				writes = append(writes, setFeeAggregatorReport.Output)
			}
		}

		for _, mockReceiverParams := range input.ContractParams.MockReceivers {
			requiredVerifiers, optionalVerifiers, err := getMockReceiverVerifiers(mockReceiverParams, addresses, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get mock receiver verifiers: %w", err)
			}
			var qualifierPtr *string
			if mockReceiverParams.Qualifier != "" {
				qualifierPtr = &mockReceiverParams.Qualifier
			}
			deployReceiverReport, err := cldf_ops.ExecuteOperation(b, mock_receiver.Deploy, chain, contract_utils.DeployInput[mock_receiver.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(mock_receiver.ContractType, *mockReceiverParams.Version),
				ChainSelector:  chain.Selector,
				Args: mock_receiver.ConstructorArgs{
					Required:  requiredVerifiers,
					Optional:  optionalVerifiers,
					Threshold: mockReceiverParams.OptionalThreshold,
				},
				Qualifier: qualifierPtr,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy MockReceiver: %w", err)
			}
			addresses = append(addresses, deployReceiverReport.Output)

			// Set minimum block depth on the MockReceiver if diff exists
			if mockReceiverParams.MinimumBlockConfirmations != 0 {
				// Get the minimum block depth on the MockReceiver
				minimumBlockConfirmations, err := cldf_ops.ExecuteOperation(b, mock_receiver_v2.GetCCVsAndMinBlockConfirmations, chain, contract_utils.FunctionInput[mock_receiver_v2.GetCCVsAndMinBlockConfirmationsArgs]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(deployReceiverReport.Output.Address),
					Args: mock_receiver_v2.GetCCVsAndMinBlockConfirmationsArgs{
						Arg0: chain.Selector,
						Arg1: []byte{},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get minimum block depth on MockReceiver: %w", err)
				}
				if minimumBlockConfirmations.Output.MinBlockDepth != mockReceiverParams.MinimumBlockConfirmations {
					// Set the minimum block depth on the MockReceiver
					setMinimumBlockConfirmationsReport, err := cldf_ops.ExecuteOperation(b, mock_receiver_v2.SetMinBlockConfirmations, chain, contract_utils.FunctionInput[uint16]{
						ChainSelector: chain.Selector,
						Address:       common.HexToAddress(deployReceiverReport.Output.Address),
						Args:          mockReceiverParams.MinimumBlockConfirmations,
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set minimum block depth on MockReceiver: %w", err)
					}
					writes = append(writes, setMinimumBlockConfirmationsReport.Output)
				}
			}
		}

		// Transfer ownership of MCM and product contracts to the CLLCCIP timelock.
		// MCM contracts (Proposer, Bypasser, Canceller) were added to ownableContracts
		// during the MCMS validation above; product contracts are added here.
		if cllccipTimelockAddr != (common.Address{}) {
			ownableContracts = append(ownableContracts,
				ownableContract{common.HexToAddress(rmnRemoteRef.Address), rmn_remote.ContractType, []common.Address{cllccipTimelockAddr, rmnTimelockAddr}},
				ownableContract{common.HexToAddress(routerRef.Address), router.ContractType, []common.Address{cllccipTimelockAddr}},
				ownableContract{common.HexToAddress(tokenAdminRegistryRef.Address), token_admin_registry.ContractType, []common.Address{cllccipTimelockAddr}},
				// ownableContract{common.HexToAddress(registryModuleOwnerCustomRef.Address), registry_module_owner_custom.ContractType},
				ownableContract{common.HexToAddress(feeQuoterRef.Address), fee_quoter.ContractType, []common.Address{cllccipTimelockAddr}},
				ownableContract{common.HexToAddress(offRampRef.Address), offramp.ContractType, []common.Address{cllccipTimelockAddr}},
				ownableContract{common.HexToAddress(onRampRef.Address), onramp.ContractType, []common.Address{cllccipTimelockAddr}},
			)
			if err := transferContractsOwnership(b, chain, ownableContracts, cllccipTimelockAddr); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership to CLLCCIP timelock: %w", err)
			}

			// Ensure both timelocks are self-governed: the CLLCCIP timelock should be
			// admin of both itself and the RMNMCMS timelock. If the deployer still holds
			// the admin role on either, fix it now so incomplete earlier setups are healed.
			if err := ensureTimelockSelfGoverned(b, chain, cllccipTimelockAddr, cllccipTimelockAddr, []common.Address{cllccipTimelockAddr}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to ensure CLLCCIP timelock is self-governed: %w", err)
			}
			if err := ensureTimelockSelfGoverned(b, chain, rmnTimelockAddr, cllccipTimelockAddr, []common.Address{cllccipTimelockAddr, rmnTimelockAddr}); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to ensure RMNMCMS timelock is governed by CLLCCIP timelock: %w", err)
			}
		}

		var batchOps []mcms_types.BatchOperation
		batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		batchOps = append(batchOps, batchOp)
		batchOps = append(batchOps, committeeVerifierBatchOps...)

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	},
)

type ownableContract struct {
	Address      common.Address
	ContractType deployment.ContractType
	// AcceptableOwners lists addresses that are considered valid owners besides
	// the target newOwner. If the contract is already owned by any of these, no
	// transfer is performed. For example, RMN MCM contracts may be acceptably
	// owned by either the CLL timelock or the RMN timelock.
	AcceptableOwners []common.Address
}

// ResolveOwnershipDeps looks up the MCMS contracts required for ownership
// transfer from existingAddresses. It returns the CLL and RMN timelock
// addresses together with the MCM contracts (Proposer, Bypasser, Canceller)
// wrapped as ownableContracts so the caller can include them in the
// transfer-ownership pass.
func ResolveOwnershipDeps(
	existingAddresses []datastore.AddressRef,
	chainSelector uint64,
) (cllccipTimelockAddr, rmnTimelockAddr common.Address, mcmContracts []ownableContract, err error) {
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
	for _, qualifier := range []string{common_utils.CLLQualifier, common_utils.RMNTimelockQualifier} {
		acceptableOwners := []common.Address{cllccipTimelockAddr}
		if qualifier == common_utils.RMNTimelockQualifier {
			acceptableOwners = append(acceptableOwners, rmnTimelockAddr)
		}
		for _, ct := range mcmTypes {
			addr, err := datastore_utils.FindAndFormatRef(mcmsDS, datastore.AddressRef{
				Type:      datastore.ContractType(ct),
				Qualifier: qualifier,
			}, chainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return common.Address{}, common.Address{}, nil,
					fmt.Errorf("ownership transfer requires MCM contract (type=%s, qualifier=%s) in ExistingAddresses: %w",
						ct, qualifier, err)
			}
			mcmContracts = append(mcmContracts, ownableContract{
				Address:          addr,
				ContractType:     ct,
				AcceptableOwners: acceptableOwners,
			})
		}
	}

	return cllccipTimelockAddr, rmnTimelockAddr, mcmContracts, nil
}

// transferContractsOwnership transfers ownership of the given contracts to newOwner.
// For each contract:
//   - If already owned by newOwner or any of the contract's AcceptableOwners, skip.
//   - If owned by the deployer, transfer to newOwner.
//   - Otherwise, error (unexpected owner).
func transferContractsOwnership(
	b cldf_ops.Bundle,
	chain evm.Chain,
	contracts []ownableContract,
	newOwner common.Address,
) error {
	for _, c := range contracts {
		currentOwner, ownable, err := mcms_seq.LoadOwnableContract(c.Address, chain.Client)
		if err != nil {
			return fmt.Errorf("failed to load ownable contract %s (%s): %w", c.Address, c.ContractType, err)
		}
		if currentOwner == newOwner {
			b.Logger.Infof("Contract %s (%s) already owned by %s, skipping transfer", c.Address, c.ContractType, newOwner)
			continue
		}
		acceptable := false
		for _, ao := range c.AcceptableOwners {
			if currentOwner == ao {
				acceptable = true
				break
			}
		}
		if acceptable {
			b.Logger.Infof("Contract %s (%s) owned by acceptable owner %s, skipping transfer", c.Address, c.ContractType, currentOwner)
			continue
		}
		if currentOwner != chain.DeployerKey.From {
			return fmt.Errorf(
				"contract %s (%s) is owned by %s, which is neither the deployer %s nor the target owner %s; cannot transfer",
				c.Address, c.ContractType, currentOwner, chain.DeployerKey.From, newOwner,
			)
		}
		deps := mcms_ops.OpEVMOwnershipDeps{
			Chain:    chain,
			OwnableC: ownable,
		}
		_, err = cldf_ops.ExecuteOperation(b, mcms_ops.OpTransferOwnership, deps, mcms_ops.OpTransferOwnershipInput{
			ChainSelector:   chain.Selector,
			Address:         c.Address,
			ProposedOwner:   newOwner,
			ContractType:    c.ContractType,
			TimelockAddress: newOwner,
		})
		if err != nil {
			return fmt.Errorf("failed to transfer ownership of %s (%s) to %s: %w", c.Address, c.ContractType, newOwner, err)
		}
	}
	return nil
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
		_, err = cldf_ops.ExecuteOperation(b, mcms_ops.OpGrantRoleTimelock, chain, contract_utils.FunctionInput[mcms_ops.OpGrantRoleTimelockInput]{
			ChainSelector: chain.Selector,
			Address:       timelockAddr,
			Args: mcms_ops.OpGrantRoleTimelockInput{
				RoleID:  mcms_ops.ADMIN_ROLE.ID,
				Account: newAdmin,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to grant ADMIN_ROLE to %s on timelock %s: %w", newAdmin, timelockAddr, err)
		}
		b.Logger.Infof("Granted ADMIN_ROLE on timelock %s to %s", timelockAddr, newAdmin)
	}

	_, err = cldf_ops.ExecuteOperation(b, mcms_ops.OpRenounceRoleTimelock, chain, contract_utils.FunctionInput[mcms_ops.OpRenounceRoleTimelockInput]{
		ChainSelector: chain.Selector,
		Address:       timelockAddr,
		Args: mcms_ops.OpRenounceRoleTimelockInput{
			RoleID: mcms_ops.ADMIN_ROLE.ID,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to renounce deployer ADMIN_ROLE on timelock %s: %w", timelockAddr, err)
	}
	b.Logger.Infof("Renounced deployer ADMIN_ROLE on timelock %s", timelockAddr)

	return nil
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
) (contract_utils.WriteOutput, bool, error) {
	// Check if the module is already registered.
	isRegisteredReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.IsRegistryModule, chain, contract_utils.FunctionInput[common.Address]{
		ChainSelector: chain.Selector,
		Address:       tokenAdminRegistryAddress,
		Args:          moduleAddress,
	})
	if err != nil {
		return contract_utils.WriteOutput{}, false, fmt.Errorf("failed to check if module is registered: %w", err)
	}

	// If already registered, return without performing a write.
	if isRegisteredReport.Output {
		return contract_utils.WriteOutput{}, false, nil
	}

	// Add the module to the registry.
	addRegistryModuleReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.AddRegistryModule, chain, contract_utils.FunctionInput[common.Address]{
		ChainSelector: chain.Selector,
		Address:       tokenAdminRegistryAddress,
		Args:          moduleAddress,
	})
	if err != nil {
		return contract_utils.WriteOutput{}, false, fmt.Errorf("failed to add registry module: %w", err)
	}

	return addRegistryModuleReport.Output, true, nil
}
