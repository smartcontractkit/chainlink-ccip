package cctp

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
)

type DeploySiloedUSDCLockReleaseInput struct {
	ChainSelector uint64
	USDCToken     string
	Router        string
	RMN           string
	TokenDecimals uint8
	// Existing siloed pool address; optional.
	SiloedUSDCTokenPool       string
	LockReleaseChainSelectors []uint64
	// Remote chain configs for lock-release chains; used to run ConfigureTokenPoolForRemoteChain (1.7.0) on the siloed pool.
	RemoteChainConfigs map[uint64]tokens.RemoteChainConfig[[]byte, string]
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
	"Deploys SiloedUSDCTokenPool and per-chain ERC20LockBox contracts",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input DeploySiloedUSDCLockReleaseInput) (output DeploySiloedUSDCLockReleaseOutput, err error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		siloedPoolAddr := input.SiloedUSDCTokenPool
		// Deploy siloed USDC token pool if not provided
		if siloedPoolAddr == "" {
			poolReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.Deploy, chain, contract_utils.DeployInput[siloed_usdc_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(siloed_usdc_token_pool.ContractType, *siloed_usdc_token_pool.Version),
				ChainSelector:  chain.Selector,
				Args: siloed_usdc_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(input.USDCToken),
					LocalTokenDecimals: input.TokenDecimals,
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
		siloedPoolAddress := common.HexToAddress(siloedPoolAddr)

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
				Address:       siloedPoolAddress,
				Args:          configs,
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to configure lockboxes on pool: %w", err)
			}
			writes = append(writes, cfgReport.Output)
		}

		// Authorize siloed pool on each lockbox
		for sel := range lockBoxes {
			lbAddr := lockBoxes[sel]
			lockBoxAddress := common.HexToAddress(lbAddr)
			// Check if already authorized
			callersReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.GetAllAuthorizedCallers, chain, contract_utils.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       lockBoxAddress,
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to get authorized callers on lockbox %s (chain %d): %w", lbAddr, sel, err)
			}
			// If not authorized, authorize it
			if !slices.Contains(callersReport.Output, siloedPoolAddress) {
				authReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
					ChainSelector: input.ChainSelector,
					Address:       lockBoxAddress,
					Args: erc20_lock_box.AuthorizedCallerArgs{
						AddedCallers: []common.Address{siloedPoolAddress},
					},
				})
				if err != nil {
					return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to authorize siloed pool on lockbox %s (chain %d): %w", lbAddr, sel, err)
				}
				writes = append(writes, authReport.Output)
			}
		}

		batchOps := make([]mcms_types.BatchOperation, 0)
		if len(writes) > 0 {
			batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			batchOps = append(batchOps, batchOp)
		}

		// Configure remote chains on the siloed pool (1.7.0 sequence)
		for remoteChainSelector, remoteChainConfig := range input.RemoteChainConfigs {
			report, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPoolForRemoteChain, chain, tokens_sequences.ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    siloedPoolAddress,
				AdvancedPoolHooks:   common.Address{},
				RemoteChainSelector: remoteChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to configure siloed pool for remote chain %d: %w", remoteChainSelector, err)
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
		}

		return DeploySiloedUSDCLockReleaseOutput{
			SiloedPoolAddress: siloedPoolAddr,
			LockBoxes:         lockBoxes,
			BatchOps:          batchOps,
			Addresses:         addresses,
		}, nil
	},
)
