package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

var ConfigureNonCanonicalUSDCForLanes = cldf_ops.NewSequence(
	"configure-non-canonical-usdc-for-lanes",
	semver.MustParse("1.6.1"),
	"Configures the non-canonical USDC contracts on a chain for multiple remote chains",
	func(b cldf_ops.Bundle, dep adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)

		// Resolve token admin registry address.
		tokenAdminRegistryRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(token_admin_registry.ContractType),
			Version: token_admin_registry.Version,
		}, input.ChainSelector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find token admin registry ref on chain %d: %w", input.ChainSelector, err)
		}

		// Find BurnMintWithLockReleaseFlagTokenPool.
		// Expect exactly one BurnMintWithLockReleaseFlagTokenPool deployed on any given non-canonical chain.
		existingPools := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByType(datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType)),
		)
		if len(existingPools) != 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("expected exactly one BurnMintWithLockReleaseFlagTokenPool to exist on chain %d, got %d", input.ChainSelector, len(existingPools))
		}
		burnMintWithLockReleaseFlagTokenPoolRef := existingPools[0]

		remoteChains := make(map[uint64]tokens.RemoteChainConfig[[]byte, string])
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			remoteToken, err := dep.RemoteChains[remoteChainSelector].TokenAddress(dep.DataStore, dep.BlockChains, remoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token address: %w", err)
			}
			remotePool, err := dep.RemoteChains[remoteChainSelector].PoolAddress(dep.DataStore, dep.BlockChains, remoteChainSelector, input.RemoteRegisteredPoolRefs[remoteChainSelector])
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pool address: %w", err)
			}
			remoteToken = common.LeftPadBytes(remoteToken, 32)
			remotePool = common.LeftPadBytes(remotePool, 32)

			remoteChains[remoteChainSelector] = tokens.RemoteChainConfig[[]byte, string]{
				TokenTransferFeeConfig:                   remoteChainConfig.TokenTransferFeeConfig,
				RemoteToken:                              remoteToken,
				RemotePool:                               remotePool,
				DefaultFinalityInboundRateLimiterConfig:  remoteChainConfig.DefaultFinalityInboundRateLimiterConfig,
				DefaultFinalityOutboundRateLimiterConfig: remoteChainConfig.DefaultFinalityOutboundRateLimiterConfig,
				CustomFinalityInboundRateLimiterConfig:   remoteChainConfig.CustomFinalityInboundRateLimiterConfig,
				CustomFinalityOutboundRateLimiterConfig:  remoteChainConfig.CustomFinalityOutboundRateLimiterConfig,
			}
		}

		// Configure the token pool for transfers.
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenForTransfers, dep.BlockChains, tokens.ConfigureTokenForTransfersInput{
			ChainSelector:     input.ChainSelector,
			TokenPoolAddress:  burnMintWithLockReleaseFlagTokenPoolRef.Address,
			RemoteChains:      remoteChains,
			RegistryAddress:   tokenAdminRegistryRef.Address,
			ExistingDataStore: dep.DataStore,
			PoolType:          string(burn_mint_with_lock_release_flag_token_pool.ContractType),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool: %w", err)
		}
		batchOps = append(batchOps, configureTokenPoolReport.Output.BatchOps...)

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
		}, nil
	},
)
