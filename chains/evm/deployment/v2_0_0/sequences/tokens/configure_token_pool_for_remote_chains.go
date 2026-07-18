package tokens

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"fmt"
	"maps"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// ConfigureTokenPoolForRemoteChainsInput is the input for the ConfigureTokenPoolForRemoteChains sequence.
type ConfigureTokenPoolForRemoteChainsInput struct {
	ChainSelector     uint64
	TokenPoolAddress  common.Address
	AdvancedPoolHooks common.Address
	RemoteChains      map[uint64]tokens.RemoteChainConfig[[]byte, string]
	RegistryAddress   common.Address
	TokenAddress      common.Address
	// SkipActivePoolSupportedChainsCheck disables the upgrade-safety check that requires remoteChains
	// to cover all chains the currently-registered pool supports. Set this when the pool being configured
	// is not a direct replacement for the registered pool (e.g. CCTP-through-CCV alongside USDCTokenPoolProxy).
	SkipActivePoolSupportedChainsCheck bool
}

// ConfigureTokenPoolForRemoteChains runs the supported-chains check once (when registry/token set) then calls
// ConfigureTokenPoolForRemoteChain for each remote chain.
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"configure-token-pool-for-remote-chains",
	semver.MustParse("2.0.0"),
	"Configures a token pool for all configured remote chains; validates active pool supported chains when upgrading.",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainsInput) (output sequences.OnChainOutput, err error) {
		// Active-pool validation: when RegistryAddress and TokenAddress are set, the registry's
		// "active" pool for that token is queried. For USDC/CCTP the registered pool is the
		// USDCTokenPoolProxy (so the router uses the proxy). The proxy does not implement
		// getSupportedChains like a 2.0.0 TokenPool and may revert—so we treat this check as
		// best-effort and skip validation when the call fails.
		//
		// This best-effort check should only fire in the case of pool upgrades NOT pool extensions.
		// If the pool is already migrated and we only want to connect new pools to it then the code
		// should NOT require the operator to list every single remote chain. Instead we should only
		// need to specify the pools we want to add. With that said this best-effort check will only
		// fire if #1 the active pool is non-zero, #2 the active pool is different from the provided
		// pool, and #3 the active pool supports `getSupportedChains`. The operator can override the
		// supported chains check entirely by setting SkipActivePoolSupportedChainsCheck to true.
		if input.RegistryAddress != (common.Address{}) && input.TokenAddress != (common.Address{}) {
			tokenConfigReport, err := evmops.ExecuteRead(b, chain, input.RegistryAddress, tarbindings.NewTokenAdminRegistry, token_admin_registry.NewReadGetTokenConfig, input.TokenAddress, cldf_ops.WithForceExecute[contract.FunctionInput[common.Address], evm.Chain]())
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config from registry for supported-chains check: %w", err)
			}
			activePool := tokenConfigReport.Output.TokenPool
			if !input.SkipActivePoolSupportedChainsCheck && activePool != (common.Address{}) && activePool != input.TokenPoolAddress {
				supportedChainsReport, err := evmops.ExecuteRead(b, chain, activePool, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadGetSupportedChains, struct{}{}, cldf_ops.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
				if err == nil {
					supportedChains := supportedChainsReport.Output
					for _, sel := range supportedChains {
						if _, ok := input.RemoteChains[sel]; !ok {
							slices.Sort(supportedChains)
							return sequences.OnChainOutput{}, fmt.Errorf("remoteChains must include all active pool supported chains: pool has %v, remoteChains has %v",
								supportedChains, slices.Sorted(maps.Keys(input.RemoteChains)))
						}
					}
				}
				// If GetSupportedChains failed (e.g. active pool is USDCTokenPoolProxy and reverts), skip validation.
			}
		}
		ops := make([]mcms_types.BatchOperation, 0)
		supportedChainsReport, err := evmops.ExecuteRead(b, chain, input.TokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadGetSupportedChains, struct{}{})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains from pool: %w", err)
		}
		supportedSet := make(map[uint64]struct{}, len(supportedChainsReport.Output))
		for _, s := range supportedChainsReport.Output {
			supportedSet[s] = struct{}{}
		}
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			_, alreadySupported := supportedSet[remoteChainSelector]
			report, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:               input.ChainSelector,
				TokenPoolAddress:            input.TokenPoolAddress,
				AdvancedPoolHooks:           input.AdvancedPoolHooks,
				RemoteChainSelector:         remoteChainSelector,
				RemoteChainConfig:           remoteChainConfig,
				RegistryAddress:             input.RegistryAddress,
				TokenAddress:                input.TokenAddress,
				RemoteChainAlreadySupported: alreadySupported,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chain %d: %w", remoteChainSelector, err)
			}
			ops = append(ops, report.Output.BatchOps...)
		}
		return sequences.OnChainOutput{BatchOps: ops}, nil
	},
)
