package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

const localTokenDecimals = 6

var usdcQualifier = "USDC"

var DeployNonCanonicalCCTPChain = cldf_ops.NewSequence(
	"deploy-non-canonical-cctp-chain",
	semver.MustParse("1.6.1"),
	"Deploys & configures the non-canonical CCTP contracts on a chain",
	func(b cldf_ops.Bundle, dep adapters.DeployCCTPChainDeps, input adapters.DeployCCTPInput) (output sequences.OnChainOutput, err error) {
		addresses := make([]datastore.AddressRef, 0)
		batchOps := make([]mcms_types.BatchOperation, 0)
		usdcTokenAddress := common.HexToAddress(input.USDCToken)

		// Resolve chain
		chain, ok := dep.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		// Resolve RMN and router addresses
		rmnProxyRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(rmn_proxy.ContractType),
			Version: rmn_proxy.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find RMN proxy ref on chain %d: %w", chain.Selector, err)
		}
		routerRef, err := datastore_utils.FindAndFormatRef(dep.DataStore, datastore.AddressRef{
			Type:    datastore.ContractType(router.ContractType),
			Version: router.Version,
		}, chain.Selector, datastore_utils.FullRef)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to find router ref on chain %d: %w", chain.Selector, err)
		}

		// Search datastore for BurnMintWithLockReleaseFlagTokenPool
		// Expect 0 or 1 BurnMintWithLockReleaseFlagTokenPool deployed on any given non-canonical chain, regardless of version.
		existingPools := dep.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByType(datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType)),
		)
		if len(existingPools) > 1 {
			return sequences.OnChainOutput{}, fmt.Errorf("expected at most one BurnMintWithLockReleaseFlagTokenPool to exist on chain %d, got %d", input.ChainSelector, len(existingPools))
		}

		var burnMintWithLockReleaseFlagTokenPoolRef datastore.AddressRef
		if len(existingPools) == 1 {
			burnMintWithLockReleaseFlagTokenPoolRef = existingPools[0]
		} else {
			burnMintWithLockReleaseFlagTokenPoolReport, err := cldf_ops.ExecuteOperation(b, burn_mint_with_lock_release_flag_token_pool.Deploy, chain, contract_utils.DeployInput[burn_mint_with_lock_release_flag_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_with_lock_release_flag_token_pool.ContractType, *burn_mint_with_lock_release_flag_token_pool.Version),
				ChainSelector:  chain.Selector,
				Qualifier:      &usdcQualifier,
				Args: burn_mint_with_lock_release_flag_token_pool.ConstructorArgs{
					Token:              usdcTokenAddress,
					LocalTokenDecimals: localTokenDecimals,
					RmnProxy:           common.HexToAddress(rmnProxyRef.Address),
					Router:             common.HexToAddress(routerRef.Address),
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy BurnMintWithLockReleaseFlagTokenPool: %w", err)
			}
			burnMintWithLockReleaseFlagTokenPoolRef = burnMintWithLockReleaseFlagTokenPoolReport.Output
		}
		addresses = append(addresses, burnMintWithLockReleaseFlagTokenPoolRef)

		// Configure the token pool.
		configureTokenPoolReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPool, chain, ConfigureTokenPoolInput{
			ChainSelector:    input.ChainSelector,
			TokenPoolAddress: common.HexToAddress(burnMintWithLockReleaseFlagTokenPoolRef.Address),
			RouterAddress:    common.HexToAddress(routerRef.Address),
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
