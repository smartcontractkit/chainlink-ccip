package cctp

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	lockboxbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	sutpbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	evm_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	tokens_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
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
	// Existing address refs from the datastore; used to skip re-deploying already-deployed lockboxes.
	ExistingAddresses []datastore.AddressRef
	// Remote chain configs for lock-release chains; used to run ConfigureTokenPoolForRemoteChain (2.0.0) on the siloed pool.
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
	semver.MustParse("2.0.0"),
	"Deploys SiloedUSDCTokenPool and per-chain ERC20LockBox contracts",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input DeploySiloedUSDCLockReleaseInput) (output DeploySiloedUSDCLockReleaseOutput, err error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)

		siloedPoolAddr := input.SiloedUSDCTokenPool
		if siloedPoolAddr == "" {
			poolReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.Deploy, chain, ops2contract.DeployInput[siloed_usdc_token_pool.ConstructorArgs]{
				TypeAndVersion: deployment.NewTypeAndVersion(siloed_usdc_token_pool.ContractType, *siloed_usdc_token_pool.Version),
				Args: siloed_usdc_token_pool.ConstructorArgs{
					Token:              common.HexToAddress(input.USDCToken),
					LocalTokenDecimals: input.TokenDecimals,
					AdvancedPoolHooks:  common.Address{},
					RmnProxy:           common.HexToAddress(input.RMN),
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
		if len(input.LockReleaseChainSelectors) > 0 {
			configs := make([]sutpbind.SiloedLockReleaseTokenPoolLockBoxConfig, 0, len(input.LockReleaseChainSelectors))
			for _, sel := range input.LockReleaseChainSelectors {
				qualifier := fmt.Sprintf("remoteChainSelector(%d)", sel)

				existingRef := datastore_utils.GetAddressRef(
					input.ExistingAddresses,
					input.ChainSelector,
					erc20_lock_box.ContractType,
					erc20_lock_box.Version,
					qualifier,
				)

				var lbAddr string
				if !datastore_utils.IsAddressRefEmpty(existingRef) {
					lbAddr = existingRef.Address
				} else {
					lbReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.Deploy, chain, ops2contract.DeployInput[erc20_lock_box.ConstructorArgs]{
						TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
						Qualifier:      &qualifier,
						Args: erc20_lock_box.ConstructorArgs{
							Token: common.HexToAddress(input.USDCToken),
						},
					})
					if err != nil {
						return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to deploy ERC20LockBox for chain %d: %w", sel, err)
					}
					addresses = append(addresses, lbReport.Output)
					lbAddr = lbReport.Output.Address
				}

				lockBoxes[sel] = lbAddr
				configs = append(configs, sutpbind.SiloedLockReleaseTokenPoolLockBoxConfig{
					RemoteChainSelector: sel,
					LockBox:             common.HexToAddress(lbAddr),
				})
			}

			siloedPool, err := sutpbind.NewSiloedUSDCTokenPool(siloedPoolAddress, chain.Client)
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to bind SiloedUSDCTokenPool at %s: %w", siloedPoolAddress.Hex(), err)
			}
			cfgReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.NewWriteConfigureLockBoxes(siloedPool), chain, ops2contract.FunctionInput[[]sutpbind.SiloedLockReleaseTokenPoolLockBoxConfig]{
				Args: configs,
			})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to configure lockboxes on pool: %w", err)
			}
			writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(cfgReport.Output))
		}

		for sel := range lockBoxes {
			lbAddr := lockBoxes[sel]
			lockBoxAddress := common.HexToAddress(lbAddr)
			lockBox, err := lockboxbind.NewERC20LockBox(lockBoxAddress, chain.Client)
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to bind ERC20LockBox at %s: %w", lockBoxAddress.Hex(), err)
			}
			callersReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.NewReadGetAllAuthorizedCallers(lockBox), chain, ops2contract.FunctionInput[struct{}]{})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to get authorized callers on lockbox %s (chain %d): %w", lbAddr, sel, err)
			}
			if !slices.Contains(callersReport.Output, siloedPoolAddress) {
				authReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.NewWriteApplyAuthorizedCallerUpdates(lockBox), chain, ops2contract.FunctionInput[lockboxbind.AuthorizedCallersAuthorizedCallerArgs]{
					Args: lockboxbind.AuthorizedCallersAuthorizedCallerArgs{
						AddedCallers: []common.Address{siloedPoolAddress},
					},
				})
				if err != nil {
					return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to authorize siloed pool on lockbox %s (chain %d): %w", lbAddr, sel, err)
				}
				writes = append(writes, tokens_sequences.WriteOutputOps2ToLegacy(authReport.Output))
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

		if len(input.RemoteChainConfigs) > 0 {
			siloedTP, err := tokens_sequences.BindTokenPool(siloedPoolAddress, chain)
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, err
			}
			supportedChainsReport, err := cldf_ops.ExecuteOperation(b, evm_token_pool.NewReadGetSupportedChains(siloedTP), chain, ops2contract.FunctionInput[struct{}]{})
			if err != nil {
				return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to get supported chains on siloed pool: %w", err)
			}
			supportedChains := supportedChainsReport.Output

			for remoteChainSelector, remoteChainConfig := range input.RemoteChainConfigs {
				report, err := cldf_ops.ExecuteSequence(b, tokens_sequences.ConfigureTokenPoolForRemoteChain, chain, tokens_sequences.ConfigureTokenPoolForRemoteChainInput{
					ChainSelector:               input.ChainSelector,
					TokenPoolAddress:            siloedPoolAddress,
					AdvancedPoolHooks:           common.Address{},
					RemoteChainSelector:         remoteChainSelector,
					RemoteChainConfig:           remoteChainConfig,
					RemoteChainAlreadySupported: slices.Contains(supportedChains, remoteChainSelector),
				})
				if err != nil {
					return DeploySiloedUSDCLockReleaseOutput{}, fmt.Errorf("failed to configure siloed pool for remote chain %d: %w", remoteChainSelector, err)
				}
				batchOps = append(batchOps, report.Output.BatchOps...)
			}
		}

		return DeploySiloedUSDCLockReleaseOutput{
			SiloedPoolAddress: siloedPoolAddr,
			LockBoxes:         lockBoxes,
			BatchOps:          batchOps,
			Addresses:         addresses,
		}, nil
	},
)
