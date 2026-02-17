package sequences

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/registry_module_owner_custom"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/proxy"
	proxy_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
)

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
	USDPerLINK                     *big.Int
	USDPerWETH                     *big.Int
}

type ExecutorParams struct {
	Version       *semver.Version
	MaxCCVsPerMsg uint8
	DynamicConfig executor.SetDynamicConfigArgs
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
}

var DeployChainContracts = cldf_ops.NewSequence(
	"deploy-chain-contracts",
	semver.MustParse("1.7.0"),
	"Deploys all required contracts for CCIP 1.7.0 to an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input DeployChainContractsInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

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
		linkRef, err := contract_utils.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain, contract_utils.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link_token.ContractType, *link_token.Version),
			ChainSelector:  chain.Selector,
			Args: burn_mint_erc20_with_drip.ConstructorArgs{
				Name:   "LINK",
				Symbol: "LINK",
			},
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

		// Set the RMNRemote on the RMNProxy
		// Included in case the RMNRemote got deployed but the RMNProxy already existed.
		// In this case, we would not have set the RMNRemote in the constructor.
		// We would need to update the RMN on the existing RMNProxy.
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
		feeQuoterRef, err := contract_utils.MaybeDeployContract(b, fee_quoter.Deploy, chain, contract_utils.DeployInput[fee_quoter.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(fee_quoter.ContractType, *input.ContractParams.FeeQuoter.Version),
			ChainSelector:  chain.Selector,
			Args: fee_quoter.ConstructorArgs{
				StaticConfig: fee_quoter.StaticConfig{
					MaxFeeJuelsPerMsg: input.ContractParams.FeeQuoter.MaxFeeJuelsPerMsg,
					LinkToken:         common.HexToAddress(linkRef.Address),
				},
				PriceUpdaters: []common.Address{
					// Price updates via protocol are out of scope for initial launch.
					// TODO: Add Timelock here when MCMS support is needed.
					chain.DeployerKey.From,
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
		updatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract_utils.FunctionInput[fee_quoter.PriceUpdates]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(feeQuoterRef.Address),
			Args: fee_quoter.PriceUpdates{
				TokenPriceUpdates: []fee_quoter.TokenPriceUpdate{
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

			// Deploy ExecutorProxy via CREATE2
			var executorProxyRef *datastore.AddressRef
			for _, ref := range input.ExistingAddresses {
				if ref.Type == datastore.ContractType(executor.ProxyType) &&
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
					Type:           datastore.ContractType(executor.ProxyType),
					Version:        executor.Version,
					CREATE2Factory: input.CREATE2Factory,
					ABI:            proxy_latest.ProxyABI,
					BIN:            proxy_latest.ProxyBin,
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
				acceptOwnershipReport, err := cldf_ops.ExecuteOperation(b, proxy.AcceptOwnership, chain, contract_utils.FunctionInput[proxy.AcceptOwnershipArgs]{
					ChainSelector: chain.Selector,
					Address:       common.HexToAddress(executorProxyRef.Address),
					Args: proxy.AcceptOwnershipArgs{
						IsProposedOwner: true,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to accept ownership of ExecutorProxy: %w", err)
				}
				writes = append(writes, acceptOwnershipReport.Output)
			}

			// Set target on the ExecutorProxy
			setTargetReport, err := cldf_ops.ExecuteOperation(b, proxy.SetTarget, chain, contract_utils.FunctionInput[common.Address]{
				ChainSelector: chain.Selector,
				Address:       common.HexToAddress(executorProxyRef.Address),
				Args:          common.HexToAddress(executorRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set target on ExecutorProxy: %w", err)
			}
			writes = append(writes, setTargetReport.Output)

			// Set fee aggregator on the ExecutorProxy
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
					RequiredVerifiers: requiredVerifiers,
					OptionalVerifiers: optionalVerifiers,
					OptionalThreshold: mockReceiverParams.OptionalThreshold,
				},
				Qualifier: qualifierPtr,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy MockReceiver: %w", err)
			}
			addresses = append(addresses, deployReceiverReport.Output)
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
