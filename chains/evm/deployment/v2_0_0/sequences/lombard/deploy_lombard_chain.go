package lombard

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	aph_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/advanced_pool_hooks"
	lombard_verifier_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lombard_verifier"
	vvr_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lombard_verifier"
	v2_0_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/verifier_tags"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
)

const (
	fallbackDecimals uint8 = 18
)

var (
	ContractQualifier = "Lombard"
)

var DeployLombardChain = cldf_ops.NewSequence(
	"deploy-lombard-chain",
	semver.MustParse("2.0.0"),
	"Deploys the Lombard chain with all required contracts and configurations",
	func(b cldf_ops.Bundle, dep adapters.DeployLombardChainDeps, input adapters.DeployLombardInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		existingAddressesOnChain := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
		)

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
		lombardVerifierRef, err := evmops.MaybeDeployContract(b, lombard_verifier.Deploy, chain, contract.DeployInput[lombard_verifier.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(lombard_verifier.ContractType, *lombard_verifier.Version),
			Qualifier:      &ContractQualifier,
			Args: lombard_verifier.ConstructorArgs{
				Bridge:          lombardBridgeAddress,
				StorageLocation: input.StorageLocations,
				DynamicConfig: lombard_verifier_bindings.LombardVerifierDynamicConfig{
					FeeAggregator: feeAggregatorAddress,
				},
				Rmn:        rmnAddress,
				VersionTag: verifier_tags.LombardVerifierV2(),
			},
		}, existingAddressesOnChain)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardVerifier: %w", err)
		}
		addresses = append(addresses, lombardVerifierRef)
		lombardVerifierAddress := common.HexToAddress(lombardVerifierRef.Address)

		_, err = evmops.ExecuteWrite(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewWriteUpdateSupportedTokens, lombard_verifier.UpdateSupportedTokensArgs{
			TokensToSet: []lombard_verifier_bindings.LombardVerifierSupportedTokenArgs{
				{
					LocalToken:   tokenAddress,
					LocalAdapter: localAdapterAddress,
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to update supported tokens on LombardVerifier: %w", err)
		}

		// Deploy LombardVerifierResolver if needed
		lombardVerifierResolverRefs := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chain.Selector),
			datastore.AddressRefByType(datastore.ContractType(versioned_verifier_resolver.LombardVerifierResolverType)),
			datastore.AddressRefByVersion(lombard_verifier.Version),
			datastore.AddressRefByQualifier(versioned_verifier_resolver.LombardVerifierResolverType.String()), // The qualifier for LombardVerifierResolver is the same as its type to avoid colliding with CommitteeVerifierResolver deployed with the "default" qualifier
		)
		var lombardVerifierResolverRef datastore.AddressRef
		if len(lombardVerifierResolverRefs) == 0 {
			if input.DeployerContract == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("deployer contract is required")
			}

			deployVerifierResolverViaCREATE2Report, err := cldf_ops.ExecuteSequence(b, v2_0_0_sequences.DeployVerifierResolverViaCREATE2, chain, v2_0_0_sequences.DeployVerifierResolverViaCREATE2Input{
				ChainSelector:  input.ChainSelector,
				Type:           datastore.ContractType(versioned_verifier_resolver.LombardVerifierResolverType),
				Version:        lombard_verifier.Version,
				CREATE2Factory: common.HexToAddress(input.DeployerContract),
				Qualifier:      versioned_verifier_resolver.LombardVerifierResolverType.String(), // The qualifier for LombardVerifierResolver is the same as its type to avoid colliding with CommitteeVerifierResolver deployed with the "default" qualifier
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

		versionTagReport, err := evmops.ExecuteRead(b, chain, lombardVerifierAddress, evmops.BindAs[lombard_verifier_bindings.LombardVerifierInterface](lombard_verifier_bindings.NewLombardVerifier), lombard_verifier.NewReadVersionTag, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get version tag from LombardVerifier: %w", err)
		}

		report, err := evmops.ExecuteWrite(b, chain, common.HexToAddress(lombardVerifierResolverRef.Address), vvr_bindings.NewVersionedVerifierResolver, versioned_verifier_resolver.NewWriteApplyInboundImplementationUpdates, []versioned_verifier_resolver.InboundImplementationArgs{
			{Version: versionTagReport.Output, Verifier: lombardVerifierAddress},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set inbound implementation on LombardVerifierResolver: %w", err)
		}
		writes = append(writes, report.Output)

		// There can be multiple pools / tokens and advancedPoolHooks for Lombard
		advancedPoolHooksRef, err := evmops.MaybeDeployContract(b, advanced_pool_hooks.Deploy, chain, contract.DeployInput[advanced_pool_hooks.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(advanced_pool_hooks.ContractType, *advanced_pool_hooks.Version),
			Qualifier:      tokenPoolQualifier(input.TokenQualifier),
			Args: advanced_pool_hooks.ConstructorArgs{
				Allowlist:                        []common.Address{}, // Empty allowlist
				ThresholdAmountForAdditionalCCVs: big.NewInt(0),
			},
		}, existingAddressesOnChain)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy AdvancedPoolHooks: %w", err)
		}
		addresses = append(addresses, advancedPoolHooksRef)
		advancedPoolHooksAddress := common.HexToAddress(advancedPoolHooksRef.Address)
		lombardVerifierResolverAddress := common.HexToAddress(lombardVerifierResolverRef.Address)

		lombardTokenPoolRef, err := evmops.MaybeDeployContract(b, lombard_token_pool.Deploy, chain, contract.DeployInput[lombard_token_pool.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(lombard_token_pool.ContractType, *lombard_token_pool.Version),
			Qualifier:      tokenPoolQualifier(input.TokenQualifier),
			Args: lombard_token_pool.ConstructorArgs{
				Token:             tokenAddress,
				Verifier:          lombardVerifierResolverAddress,
				Bridge:            lombardBridgeAddress,
				Adapter:           localAdapterAddress,
				AdvancedPoolHooks: advancedPoolHooksAddress,
				RmnProxy:          rmnAddress,
				Router:            routerAddress,
				FallbackDecimals:  fallbackDecimals,
			},
		}, existingAddressesOnChain)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy LombardTokenPool: %w", err)
		}
		addresses = append(addresses, lombardTokenPoolRef)
		lombardTokenPoolAddress := common.HexToAddress(lombardTokenPoolRef.Address)

		// Add the newly deployed token pool as an authorized caller on the hooks.
		{
			getAuthorizedCallersReport, err := evmops.ExecuteRead(b, chain, advancedPoolHooksAddress, evmops.BindAs[aph_bindings.AdvancedPoolHooksInterface](aph_bindings.NewAdvancedPoolHooks), advanced_pool_hooks.NewReadGetAllAuthorizedCallers, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers from advanced pool hooks %s on %d: %w", advancedPoolHooksAddress, input.ChainSelector, err)
			}

			if !slices.Contains(getAuthorizedCallersReport.Output, lombardTokenPoolAddress) {
				applyAuthorizedCallerUpdatesReport, err := evmops.ExecuteWrite(b, chain, advancedPoolHooksAddress, evmops.BindAs[aph_bindings.AdvancedPoolHooksInterface](aph_bindings.NewAdvancedPoolHooks), advanced_pool_hooks.NewWriteApplyAuthorizedCallerUpdates, aph_bindings.AuthorizedCallersAuthorizedCallerArgs{
					AddedCallers: []common.Address{lombardTokenPoolAddress},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize token pool %s on advanced pool hooks %s on %d: %w", lombardTokenPoolAddress, advancedPoolHooksAddress, input.ChainSelector, err)
				}

				batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{applyAuthorizedCallerUpdatesReport.Output})
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
