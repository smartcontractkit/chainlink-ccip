package cctp

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeploySiloedUSDCLockReleaseInput struct {
	ChainSelector uint64
	USDCToken     string
	Router        string
	RMN           string
	// Existing proxy and siloed pool addresses; proxy is required, pool optional.
	USDCTokenPoolProxy        string
	SiloedUSDCTokenPool       string
	LockReleaseChainSelectors []uint64
}

type DeploySiloedUSDCLockReleaseOutput struct {
	SiloedPoolAddress string
	LockBoxes         map[uint64]string
	BatchOps          []mcms_types.BatchOperation
	Addresses         []datastore.AddressRef
}

var DeploySiloedUSDCLockRelease = cldf_ops.NewSequence(
	"deploy-siloed-usdc-lock-release",
	semver.MustParse("1.7.0"),
	"Deploys SiloedUSDCTokenPool and per-chain ERC20LockBox contracts, wiring them to the USDCTokenPoolProxy for lock-release lanes",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input DeploySiloedUSDCLockReleaseInput) (output DeploySiloedUSDCLockReleaseOutput, err error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		if input.USDCTokenPoolProxy == "" {
			return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("usdctokenpoolproxy address is required")
		}

		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		siloedPoolAddr := input.SiloedUSDCTokenPool
		// Deploy siloed USDC token pool if not provided
		if siloedPoolAddr == "" {
			poolReport, err := cldf_ops.ExecuteOperation(b, lock_release_token_pool.Deploy, chain, contract_utils.DeployInput[lock_release_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(lock_release_token_pool.SiloedUSDCTokenPoolContractType, *lock_release_token_pool.Version),
				ChainSelector:  chain.Selector,
				Args: lock_release_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(input.USDCToken),
					LocalTokenDecimals: localTokenDecimals,
					AdvancedPoolHooks:  common.Address{},
					RMNProxy:           common.HexToAddress(input.RMN),
					Router:             common.HexToAddress(input.Router),
				},
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to deploy SiloedUSDCTokenPool: %w", err)
			}
			addresses = append(addresses, poolReport.Output)
			siloedPoolAddr = poolReport.Output.Address
		}

		lockBoxes := make(map[uint64]string, len(input.LockReleaseChainSelectors))
		// Deploy lockboxes and configure them on the pool
		if len(input.LockReleaseChainSelectors) > 0 {
			configs := make([]siloed_usdc_token_pool.LockBoxConfig, 0, len(input.LockReleaseChainSelectors))
			for _, sel := range input.LockReleaseChainSelectors {
				lbReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.Deploy, chain, contract_utils.DeployInput[erc20_lock_box.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
					ChainSelector:  chain.Selector,
					Args: erc20_lock_box.ConstructorArgs{
						Token: common.HexToAddress(input.USDCToken),
					},
				})
				if err != nil {
					return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to deploy ERC20LockBox for chain %d: %w", sel, err)
				}
				addresses = append(addresses, lbReport.Output)
				lockBoxes[sel] = lbReport.Output.Address
				configs = append(configs, siloed_usdc_token_pool.LockBoxConfig{
					RemoteChainSelector: sel,
					LockBox:             common.HexToAddress(lbReport.Output.Address),
				})
			}

			cfgReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.ConfigureLockBoxes, chain, contract_utils.FunctionInput[[]siloed_usdc_token_pool.LockBoxConfig]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(siloedPoolAddr),
				Args:          configs,
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to configure lockboxes on pool: %w", err)
			}
			writes = append(writes, cfgReport.Output)
		}

		// Authorize proxy on pool
		poolAuthReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[siloed_usdc_token_pool.AuthorizedCallerArgs]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(siloedPoolAddr),
			Args: siloed_usdc_token_pool.AuthorizedCallerArgs{
				AddedCallers: []common.Address{common.HexToAddress(input.USDCTokenPoolProxy)},
			},
		})
		if err != nil {
			return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to authorize proxy on siloed pool: %w", err)
		}
		writes = append(writes, poolAuthReport.Output)

		// Authorize siloed pool on each lockbox
		for sel := range lockBoxes {
			lbAddr := lockBoxes[sel]
			authReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(lbAddr),
				Args: erc20_lock_box.AuthorizedCallerArgs{
					AddedCallers: []common.Address{common.HexToAddress(siloedPoolAddr)},
				},
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to authorize siloed pool on lockbox %s (chain %d): %w", lbAddr, sel, err)
			}
			writes = append(writes, authReport.Output)
		}

		// Update proxy pool addresses and lock-release mechanisms
		if len(input.LockReleaseChainSelectors) > 0 {
			poolsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.GetPools, chain, contract_utils.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.USDCTokenPoolProxy),
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to get existing proxy pools: %w", err)
			}
			currentPools := poolsReport.Output
			if currentPools.SiloedLockReleasePool.Hex() != common.HexToAddress(siloedPoolAddr).Hex() {
				updatePoolsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdatePoolAddresses, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.PoolAddresses]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(input.USDCTokenPoolProxy),
					Args: usdc_token_pool_proxy.PoolAddresses{
						CctpV1Pool:            currentPools.CctpV1Pool,
						CctpV2Pool:            currentPools.CctpV2Pool,
						CctpV2PoolWithCCV:     currentPools.CctpV2PoolWithCCV,
						SiloedLockReleasePool: common.HexToAddress(siloedPoolAddr),
					},
				})
				if err != nil {
					return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to update proxy pool addresses: %w", err)
				}
				writes = append(writes, updatePoolsReport.Output)
			}

			selectors := make([]uint64, 0, len(input.LockReleaseChainSelectors))
			mechanisms := make([]uint8, 0, len(input.LockReleaseChainSelectors))
			lockReleaseMechanism, err := convertMechanismToUint8("LOCK_RELEASE")
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to resolve LOCK_RELEASE mechanism: %w", err)
			}
			for _, sel := range input.LockReleaseChainSelectors {
				selectors = append(selectors, sel)
				mechanisms = append(mechanisms, lockReleaseMechanism)
			}
			updateMechanismsReport, err := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockOrBurnMechanisms, chain, contract_utils.FunctionInput[usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.USDCTokenPoolProxy),
				Args: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
					RemoteChainSelectors: selectors,
					Mechanisms:           mechanisms,
				},
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to set lock-release mechanisms on proxy: %w", err)
			}
			writes = append(writes, updateMechanismsReport.Output)
		}

		batchOps := make([]mcms_types.BatchOperation, 0)
		if len(writes) > 0 {
			batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			batchOps = append(batchOps, batchOp)
		}

		return DeploySiloedUSDCLockReleaseOutput{
			SiloedPoolAddress: siloedPoolAddr,
			LockBoxes:         lockBoxes,
			BatchOps:          batchOps,
			Addresses:         addresses,
		}, nil
	},
)
