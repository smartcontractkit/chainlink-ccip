package lombard

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/versioned_verifier_resolver"
	v1_7_0_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
)

const (
	fallbackDecimals uint8 = 18
)

var (
	ContractQualifier = "Lombard"
)

var DeployLombardChain = cldf_ops.NewSequence(
	"deploy-lombard-chain",
	semver.MustParse("1.7.0"),
	"Deploys the Lombard chain with all required contracts and configurations",
	func(b cldf_ops.Bundle, dep adapters.DeployLombardChainDeps, input adapters.DeployLombardInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		existingAddresses, err := dep.DataStore.Addresses().Fetch()
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to fetch all addresses: %w", err)
		}

		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		tokenAddress := common.HexToAddress(input.Token)
		localAdapterAddress := common.Address{}
		if input.LocalAdapter != "" {
			localAdapterAddress = common.HexToAddress(input.LocalAdapter)
		}
		lombardBridgeAddress := common.HexToAddress(input.Bridge)
		feeAggregatorAddress := common.HexToAddress(input.FeeAggregator)

		rmnRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: rmn_proxy.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find RMN proxy ref on chain %d: %w", chain.Selector, err)
		}
		rmnAddress := common.HexToAddress(rmnRef.Address)

		routerRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(router.ContractType),
			Version: router.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find router ref on chain %d: %w", chain.Selector, err)
		}
		routerAddress := common.HexToAddress(routerRef.Address)

		// Deploy LombardVerifier if needed
		lombardVerifierRef, err := contract_utils.MaybeDeployContract(b, lombard_verifier.Deploy, chain, contract_utils.DeployInput[lombard_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(lombard_verifier.ContractType, *lombard_verifier.Version),
			ChainSelector:  input.ChainSelector,
			Qualifier:      &ContractQualifier,
			Args: lombard_verifier.ConstructorArgs{
				Bridge:           lombardBridgeAddress,
				StorageLocations: input.StorageLocations,
				DynamicConfig: lombard_verifier.DynamicConfig{
					FeeAggregator: feeAggregatorAddress,
				},
				RMN: rmnAddress,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifier: %w", err)
		}
		addresses = append(addresses, lombardVerifierRef)
		lombardVerifierAddress := common.HexToAddress(lombardVerifierRef.Address)

		_, err = cldf_ops.ExecuteOperation(b, lombard_verifier.UpdateSupportedTokens, chain, contract_utils.FunctionInput[lombard_verifier.SupportedTokenArgs]{
			ChainSelector: input.ChainSelector,
			Address:       lombardVerifierAddress,
			Args: lombard_verifier.SupportedTokenArgs{
				TokensToSet: []lombard_verifier.SupportedTokensArgs{
					{
						LocalToken:   tokenAddress,
						LocalAdapter: localAdapterAddress,
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to update supported tokens on LombardVerifier: %w", err)
		}

		// Deploy LombardVerifierResolver if needed
		lombardVerifierResolverRefs := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByType(datastore.ContractType(lombard_verifier.ResolverType)),
			datastore.AddressRefByVersion(lombard_verifier.Version),
			datastore.AddressRefByQualifier(ContractQualifier),
		)
		var lombardVerifierResolverRef datastore.AddressRef
		if len(lombardVerifierResolverRefs) == 0 {
			if input.DeployerContract == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("deployer contract is required")
			}

			deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v1_7_0_sequences.DeployVerifierResolverViaCREATE2, chain, v1_7_0_sequences.DeployVerifierResolverViaCREATE2Input{
				ChainSelector:  input.ChainSelector,
				Qualifier:      ContractQualifier,
				Type:           datastore.ContractType(lombard_verifier.ResolverType),
				Version:        lombard_verifier.Version,
				CREATE2Factory: common.HexToAddress(input.DeployerContract),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifierResolver: %w", err)
			}
			if len(deployVerifierResolverViaCREATE2Report.Output.Addresses) != 1 {
				return sequences.OnChainOutput{}, fmt.Errorf("expected 1 LombardVerifierResolver address, got %d", len(deployVerifierResolverViaCREATE2Report.Output.Addresses))
			}
			lombardVerifierResolverRef = deployVerifierResolverViaCREATE2Report.Output.Addresses[0]
			addresses = append(addresses, deployVerifierResolverViaCREATE2Report.Output.Addresses...)
			writes = append(writes, deployVerifierResolverViaCREATE2Report.Output.Writes...)
		} else if len(lombardVerifierResolverRefs) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("multiple LombardVerifierResolver addresses found for chain %d", chain.Selector)
		} else {
			lombardVerifierResolverRef = lombardVerifierResolverRefs[0]
			addresses = append(addresses, lombardVerifierResolverRef)
		}

		versionTagReport, err := cldf_ops.ExecuteOperation(b, lombard_verifier.GetVersionTag, chain, contract_utils.FunctionInput[any]{
			ChainSelector: chain.Selector,
			Address:       lombardVerifierAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from LombardVerifier: %w", err)
		}

		report, err := cldf_ops.ExecuteOperation(b, versioned_verifier_resolver.ApplyInboundImplementationUpdates, chain, contract_utils.FunctionInput[[]versioned_verifier_resolver.InboundImplementationArgs]{
			ChainSelector: chain.Selector,
			Address:       common.HexToAddress(lombardVerifierResolverRef.Address),
			Args: []versioned_verifier_resolver.InboundImplementationArgs{
				{Version: versionTagReport.Output, Verifier: lombardVerifierAddress},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set inbound implementation on LombardVerifierResolver: %w", err)
		}
		writes = append(writes, report.Output)

		// There can be multiple pools / tokens and advancedPoolHooks for Lombard
		advancedPoolHooksRef, err := contract_utils.MaybeDeployContract(b, advanced_pool_hooks.Deploy, chain, contract_utils.DeployInput[advanced_pool_hooks.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *advanced_pool_hooks.Version),
			ChainSelector:  input.ChainSelector,
			Qualifier:      tokenPoolQualifier(input.TokenQualifier),
			Args: advanced_pool_hooks.ConstructorArgs{
				Allowlist:                        []common.Address{}, // Empty allowlist
				ThresholdAmountForAdditionalCCVs: big.NewInt(0),
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy AdvancedPoolHooks: %w", err)
		}
		addresses = append(addresses, advancedPoolHooksRef)
		advancedPoolHooksAddress := common.HexToAddress(advancedPoolHooksRef.Address)
		lombardVerifierResolverAddress := common.HexToAddress(lombardVerifierResolverRef.Address)

		lombardTokenPoolRef, err := contract_utils.MaybeDeployContract(b, lombard_token_pool.Deploy, chain, contract_utils.DeployInput[lombard_token_pool.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(lombard_token_pool.ContractType, *lombard_token_pool.Version),
			ChainSelector:  input.ChainSelector,
			Qualifier:      tokenPoolQualifier(input.TokenQualifier),
			Args: lombard_token_pool.ConstructorArgs{
				Token:             tokenAddress,
				Verifier:          lombardVerifierResolverAddress,
				Bridge:            lombardBridgeAddress,
				Adapter:           localAdapterAddress,
				AdvancedPoolHooks: advancedPoolHooksAddress,
				RMNProxy:          rmnAddress,
				Router:            routerAddress,
				FallbackDecimals:  fallbackDecimals,
			},
		}, existingAddresses)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardTokenPool: %w", err)
		}
		addresses = append(addresses, lombardTokenPoolRef)
		lombardTokenPoolAddress := common.HexToAddress(lombardTokenPoolRef.Address)

		// Add the newly deployed token pool as an authorized caller on the hooks.
		{
			getAuthorizedCallersReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetAllAuthorizedCallers, chain, contract_utils.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       advancedPoolHooksAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks %s on %d: %w", advancedPoolHooksAddress, input.ChainSelector, err)
			}

			if !slices.Contains(getAuthorizedCallersReport.Output, lombardTokenPoolAddress) {
				applyAuthorizedCallerUpdatesReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[advanced_pool_hooks.AuthorizedCallerArgs]{
					ChainSelector: input.ChainSelector,
					Address:       advancedPoolHooksAddress,
					Args: advanced_pool_hooks.AuthorizedCallerArgs{
						AddedCallers: []common.Address{lombardTokenPoolAddress},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize token pool %s on advanced pool hooks %s on %d: %w", lombardTokenPoolAddress, advancedPoolHooksAddress, input.ChainSelector, err)
				}

				batchOp, err := contract_utils.NewBatchOperationFromWrites([]contract_utils.WriteOutput{applyAuthorizedCallerUpdatesReport.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
				}
				batchOps = append(batchOps, batchOp)
			}
		}

		// Configure token pool
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPool, chain, tokens_sequences.ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: lombardTokenPoolAddress,
			RouterAddress:    routerAddress,
			RateLimitAdmin:   common.HexToAddress(input.RateLimitAdmin),
			FeeAggregator:    feeAggregatorAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool: %w", err)
		}
		batchOps = append(batchOps, configureTokenPoolReport.Output.BatchOps...)

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	})

func tokenPoolQualifier(tokenQualifier string) *string {
	qualifier := ContractQualifier + "_" + tokenQualifier
	return &qualifier
}
