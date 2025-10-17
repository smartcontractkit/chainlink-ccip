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
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/onramp"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
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
	Version              *semver.Version
	GasForCallExactCheck uint16
}

type OnRampParams struct {
	Version       *semver.Version
	FeeAggregator common.Address
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
}

type ContractParams struct {
	RMNRemote         RMNRemoteParams
	OffRamp           OffRampParams
	CommitteeVerifier []CommitteeVerifierParams
	OnRamp            OnRampParams
	FeeQuoter         FeeQuoterParams
	Executor          ExecutorParams
	MockReceivers     []MockReceiverParams
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
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)

		// TODO: Deploy MCMS (Timelock, MCM contracts) when MCMS support is needed.

		// Deploy WETH
		wethRef, err := contract_utils.MaybeDeployContract(b, weth.Deploy, chain, contract.DeployInput[weth.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(weth.ContractType, *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, wethRef)

		// Deploy LINK
		linkRef, err := contract_utils.MaybeDeployContract(b, link.Deploy, chain, contract.DeployInput[link.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(link.ContractType, *semver.MustParse("1.0.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, linkRef)

		// Deploy RMNRemote
		rmnRemoteRef, err := contract_utils.MaybeDeployContract(b, rmn_remote.Deploy, chain, contract.DeployInput[rmn_remote.ConstructorArgs]{
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
		rmnProxyRef, err := contract_utils.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
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
		routerRef, err := contract_utils.MaybeDeployContract(b, router.Deploy, chain, contract.DeployInput[router.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(router.ContractType, *semver.MustParse("1.2.0")),
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
		tokenAdminRegistryRef, err := contract_utils.MaybeDeployContract(b, token_admin_registry.Deploy, chain, contract.DeployInput[token_admin_registry.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *semver.MustParse("1.5.0")),
			ChainSelector:  chain.Selector,
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		addresses = append(addresses, tokenAdminRegistryRef)

		// Deploy FeeQuoter
		feeQuoterRef, err := contract_utils.MaybeDeployContract(b, fee_quoter.Deploy, chain, contract.DeployInput[fee_quoter.ConstructorArgs]{
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
				FeeTokens: []fee_quoter.FeeTokenArgs{
					{
						Token:                      common.HexToAddress(linkRef.Address),
						PremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoter.LINKPremiumMultiplierWeiPerEth,
					},
					{
						Token:                      common.HexToAddress(wethRef.Address),
						PremiumMultiplierWeiPerEth: input.ContractParams.FeeQuoter.WETHPremiumMultiplierWeiPerEth,
					},
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
		updatePricesReport, err := cldf_ops.ExecuteOperation(b, fee_quoter.UpdatePrices, chain, contract.FunctionInput[fee_quoter.PriceUpdates]{
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
		offRampRef, err := contract_utils.MaybeDeployContract(b, offramp.Deploy, chain, contract.DeployInput[offramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(offramp.ContractType, *input.ContractParams.OffRamp.Version),
			ChainSelector:  chain.Selector,
			Args: offramp.ConstructorArgs{
				StaticConfig: offramp.StaticConfig{
					LocalChainSelector:   chain.Selector,
					RmnRemote:            common.HexToAddress(rmnProxyRef.Address),
					GasForCallExactCheck: input.ContractParams.OffRamp.GasForCallExactCheck,
					TokenAdminRegistry:   common.HexToAddress(tokenAdminRegistryRef.Address),
				},
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy OffRamp: %w", err)
		}
		addresses = append(addresses, offRampRef)

		// Deploy OnRamp
		onRampRef, err := contract_utils.MaybeDeployContract(b, onramp.Deploy, chain, contract.DeployInput[onramp.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(onramp.ContractType, *input.ContractParams.OnRamp.Version),
			ChainSelector:  chain.Selector,
			Args: onramp.ConstructorArgs{
				StaticConfig: onramp.StaticConfig{
					ChainSelector:      chain.Selector,
					RmnRemote:          common.HexToAddress(rmnRemoteRef.Address),
					TokenAdminRegistry: common.HexToAddress(tokenAdminRegistryRef.Address),
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
		var committeeVerifierRefs []datastore.AddressRef
		var committeeVerifierBatchOps []mcms_types.BatchOperation
		for _, committeeVerifierParams := range input.ContractParams.CommitteeVerifier {
			report, err := operations.ExecuteSequence(b, DeployCommitteeVerifier, chain, DeployCommitteeVerifierInput{
				ChainSelector:     chain.Selector,
				ExistingAddresses: input.ExistingAddresses,
				Params:            committeeVerifierParams,
				FeeQuoter:         common.HexToAddress(feeQuoterRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy CommitteeVerifier: %w", err)
			}
			addresses = append(addresses, report.Output.Addresses...)
			for _, addr := range report.Output.Addresses {
				if addr.Type == datastore.ContractType(committee_verifier.ContractType) {
					committeeVerifierRefs = append(committeeVerifierRefs, addr)
				}
			}
			committeeVerifierBatchOps = append(committeeVerifierBatchOps, report.Output.BatchOps...)
		}

		// Deploy Executor
		ExecutorRef, err := contract_utils.MaybeDeployContract(b, executor.Deploy, chain, contract.DeployInput[executor.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(executor.ContractType, *input.ContractParams.Executor.Version),
			ChainSelector:  chain.Selector,
			Args: executor.ConstructorArgs{
				MaxCCVsPerMsg: input.ContractParams.Executor.MaxCCVsPerMsg,
			},
		}, input.ExistingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy Executor: %w", err)
		}
		addresses = append(addresses, ExecutorRef)

		for _, mockReceiverParams := range input.ContractParams.MockReceivers {
			requiredVerifiers, optionalVerifiers, err := getMockReceiverVerifiers(mockReceiverParams, addresses, input.ExistingAddresses)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get mock receiver verifiers: %w", err)
			}
			var qualifierPtr *string
			if mockReceiverParams.Qualifier != "" {
				qualifierPtr = &mockReceiverParams.Qualifier
			}
			deployReceiverReport, err := operations.ExecuteOperation(b, mock_receiver.Deploy, chain, contract.DeployInput[mock_receiver.ConstructorArgs]{
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
		batchOp, err := contract.NewBatchOperationFromWrites(writes)
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
