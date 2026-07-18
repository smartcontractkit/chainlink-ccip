package tokens

import (
	"fmt"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ConfigureTokenForTransfers = cldf_ops.NewSequence(
	"configure-token-for-transfers",
	semver.MustParse("2.0.0"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (output sequences.OnChainOutput, err error) {
		ops := make([]mcms_types.BatchOperation, 0)
		evmChain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		registryAddress := common.Address{}
		if input.RegistryAddress != "" {
			registryAddress = common.HexToAddress(input.RegistryAddress)
		} else {
			filter := datastore.AddressRef{
				Type:          datastore.ContractType(tar_ops.ContractType),
				ChainSelector: evmChain.Selector,
				Version:       tar_ops.Version,
			}

			addr, err := datastore_utils.FindAndFormatRef(input.ExistingDataStore, filter, evmChain.Selector, datastore_utils_evm.ToEVMAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to find registry address in datastore: %w", err)
			}

			registryAddress = addr
		}

		var tokenAddress common.Address
		if input.TokenAddress != "" {
			tokenAddress = common.HexToAddress(input.TokenAddress)
		} else {
			tokenAddrReport, err := evmops.ExecuteRead(b, evmChain, common.HexToAddress(input.TokenPoolAddress), evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadGetToken, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
			}
			tokenAddress = tokenAddrReport.Output
		}
		tokenPoolAddress := common.HexToAddress(input.TokenPoolAddress)
		registryTokenPoolAddress := tokenPoolAddress
		if input.RegistryTokenPoolAddress != "" {
			registryTokenPoolAddress = common.HexToAddress(input.RegistryTokenPoolAddress)
		}

		// If a liquidity migration is requested, derive the old pool from the TAR and run the migration sequence.
		if input.LiquidityMigrationAmount != nil || input.LiquidityMigrationBasisPoints != nil {
			if input.TimelockAddress == "" {
				return sequences.OnChainOutput{}, fmt.Errorf("TimelockAddress is required when liquidity migration is requested")
			}

			tokenConfigReport, err := evmops.ExecuteRead(b, evmChain, registryAddress, tarbindings.NewTokenAdminRegistry, tar_ops.NewReadGetTokenConfig, tokenAddress)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config from registry for token %s: %w", tokenAddress, err)
			}
			oldPoolAddress := tokenConfigReport.Output.TokenPool
			if oldPoolAddress == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("no pool currently registered for token %s on chain %d", tokenAddress, input.ChainSelector)
			}

			migrationReport, err := cldf_ops.ExecuteSequence(b, MigrateLockReleasePoolLiquidity, chains, tokens.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   input.ChainSelector,
				OldPoolAddress:  oldPoolAddress.Hex(),
				NewPoolAddress:  input.TokenPoolAddress,
				TimelockAddress: input.TimelockAddress,
				Amount:          input.LiquidityMigrationAmount,
				BasisPoints:     input.LiquidityMigrationBasisPoints,
				SetPoolConfig:   nil,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute liquidity migration on chain %d: %w", input.ChainSelector, err)
			}
			ops = append(ops, migrationReport.Output.BatchOps...)
		}

		// Validate the pool supports the token
		isSupported, err := evmops.ExecuteRead(b, evmChain, tokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadIsSupportedToken, tokenAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if token %s is supported by token pool %s on %s: %w", tokenAddress, tokenPoolAddress, evmChain, err)
		}
		if !isSupported.Output {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s is not supported by token pool %s", tokenAddress, tokenPoolAddress)
		}

		if !input.AllowedFinalityConfig.IsZero() {
			desiredFinalityConfig := input.AllowedFinalityConfig.Raw()
			currentAllowedFinalityReport, err := evmops.ExecuteRead(b, evmChain, tokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadGetAllowedFinalityConfig, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allowed finality config from token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
			}
			if currentAllowedFinalityReport.Output != desiredFinalityConfig {
				configureMinBlockConfirmationReport, err := evmops.ExecuteWrite(b, evmChain, tokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewWriteSetAllowedFinalityConfig, desiredFinalityConfig)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to configure minimum block confirmation for token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
				}
				configureMinBlockConfirmationOps, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{configureMinBlockConfirmationReport.Output})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from write outputs: %w", err)
				}
				ops = append(ops, configureMinBlockConfirmationOps)
			}

			// Create a fresh bundle to avoid stale reads of allowedFinalityConfig in subsequent operations.
			// See: deployment/docs/style-guide.md#avoid-stale-reads-from-cached-operations
			b = cldf_ops.NewBundle(b.GetContext, b.Logger, cldf_ops.NewMemoryReporter())
		}

		// Get the advanced pool hooks address
		advancedPoolHooksAddress, err := evmops.ExecuteRead(b, evmChain, tokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadGetAdvancedPoolHooks, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get advanced pool hooks address from token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
		}

		// Configure token pool for all remote chains (validates active pool supported chains once when upgrading).
		configureTokenPoolForRemoteChainsReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChains, evmChain, ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:                      input.ChainSelector,
			TokenPoolAddress:                   tokenPoolAddress,
			AdvancedPoolHooks:                  advancedPoolHooksAddress.Output,
			RemoteChains:                       input.RemoteChains,
			RegistryAddress:                    registryAddress,
			TokenAddress:                       tokenAddress,
			SkipActivePoolSupportedChainsCheck: input.SkipActivePoolSupportedChainsCheck,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chains: %w", err)
		}
		ops = append(ops, configureTokenPoolForRemoteChainsReport.Output.BatchOps...)

		// Register the token with the token admin registry
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, v1_5_0.RegisterToken, evmChain, v1_5_0.RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenAddress:              tokenAddress,
			TokenPoolAddress:          registryTokenPoolAddress,
			ExternalAdmin:             common.HexToAddress(input.ExternalAdmin),
			TokenAdminRegistryAddress: registryAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to register token pool with address %s on %s: %w", input.TokenPoolAddress, evmChain, err)
		}
		ops = append(ops, registerTokenReport.Output.BatchOps...)

		return sequences.OnChainOutput{
			BatchOps: ops,
		}, nil
	},
)
