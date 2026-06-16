package tokens

import (
	"fmt"
	"maps"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
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
	// AutoMigrate, when true, carries forward remote chains supported by the active pool but not
	// explicitly listed in RemoteChains. Explicitly listed chains take precedence.
	AutoMigrate bool
}

// ConfigureTokenPoolForRemoteChains runs the supported-chains check once (when registry/token set) then calls
// ConfigureTokenPoolForRemoteChain for each remote chain.
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"configure-token-pool-for-remote-chains",
	semver.MustParse("2.0.0"),
	"Configures a token pool for all configured remote chains; validates active pool supported chains when upgrading.",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input ConfigureTokenPoolForRemoteChainsInput) (output sequences.OnChainOutput, err error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		// Active-pool validation: when RegistryAddress and TokenAddress are set, the registry's
		// "active" pool for that token is queried. For USDC/CCTP the registered pool is the
		// USDCTokenPoolProxy (so the router uses the proxy). The proxy does not implement
		// getSupportedChains like a 2.0.0 TokenPool and may revert—so we treat this check as
		// best-effort and skip validation when the call fails.
		if input.RegistryAddress != (common.Address{}) && input.TokenAddress != (common.Address{}) {
			tokenConfigReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetTokenConfig, chain, evm_contract.FunctionInput[common.Address]{
				ChainSelector: input.ChainSelector,
				Address:       input.RegistryAddress,
				Args:          input.TokenAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config from registry for supported-chains check: %w", err)
			}
			activePool := tokenConfigReport.Output.TokenPool
			if activePool != (common.Address{}) {
				supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, evm_contract.FunctionInput[struct{}]{
					ChainSelector: input.ChainSelector,
					Address:       activePool,
				})
				if err == nil {
					// For each chain the active pool supports but the operator did not list:
					//   - AutoMigrate=false (default): error — the operator must list every active-pool chain (upgrade safety).
					//   - AutoMigrate=true: discover the chain's remote token, remote pool, and remote decimals from the
					//     active pool and carry it forward. Rate limits and remote pools are imported per-chain downstream
					//     by ConfigureTokenPoolForRemoteChain (importConfigFromActivePool), so we only supply the fields
					//     that import does not provide.
					supportedChains := supportedChainsReport.Output
					for _, sel := range supportedChains {
						if _, ok := input.RemoteChains[sel]; ok {
							continue
						}
						if !input.AutoMigrate {
							slices.Sort(supportedChains)
							return sequences.OnChainOutput{}, fmt.Errorf(
								"remoteChains must include all active pool supported chains: pool has %v, remoteChains has %v",
								supportedChains, slices.Sorted(maps.Keys(input.RemoteChains)),
							)
						}
						discovered, err := discoverRemoteChainFromActivePool(b, chains, chain, input.ChainSelector, activePool, sel)
						if err != nil {
							return sequences.OnChainOutput{}, fmt.Errorf("auto-migrate: failed to discover config for active pool chain %d: %w", sel, err)
						}
						if input.RemoteChains == nil {
							input.RemoteChains = make(map[uint64]tokens.RemoteChainConfig[[]byte, string])
						}
						input.RemoteChains[sel] = discovered
					}
				}
				// If GetSupportedChains failed (e.g. active pool is USDCTokenPoolProxy and reverts), skip validation.
			}
		}
		ops := make([]mcms_types.BatchOperation, 0)
		supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
		})
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

// discoverRemoteChainFromActivePool builds a RemoteChainConfig for a chain that the active pool supports but
// the operator did not explicitly list (AutoMigrate). It reads the remote token and remote pool from the active
// pool (on the local chain) and the remote token's decimals from the remote chain. Rate limits are intentionally
// omitted: ConfigureTokenPoolForRemoteChain imports them (and the active pool's remote pools) per-chain via
// importConfigFromActivePool, which also rebases pre-1.6.1 inbound limits using RemoteDecimals.
func discoverRemoteChainFromActivePool(
	b cldf_ops.Bundle,
	chains chain.BlockChains,
	localChain evm.Chain,
	localChainSelector uint64,
	activePool common.Address,
	remoteChainSelector uint64,
) (tokens.RemoteChainConfig[[]byte, string], error) {
	var zero tokens.RemoteChainConfig[[]byte, string]
	remoteChain, ok := chains.EVMChains()[remoteChainSelector]
	if !ok {
		return zero, fmt.Errorf("remote chain with selector %d not found in environment", remoteChainSelector)
	}

	remoteTokenReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemoteToken, localChain, evm_contract.FunctionInput[uint64]{
		ChainSelector: localChainSelector,
		Address:       activePool,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return zero, fmt.Errorf("failed to get remote token from active pool: %w", err)
	}

	remoteToken := remoteTokenReport.Output
	if len(remoteToken) == 0 {
		return zero, fmt.Errorf("active pool has no remote token registered for chain %d", remoteChainSelector)
	}

	remotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemotePools, localChain, evm_contract.FunctionInput[uint64]{
		ChainSelector: localChainSelector,
		Address:       activePool,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return zero, fmt.Errorf("failed to get remote pools from active pool: %w", err)
	}

	remotePools := remotePoolsReport.Output
	if len(remotePools) == 0 {
		return zero, fmt.Errorf("active pool has no remote pool registered for chain %d", remoteChainSelector)
	}

	remotePool := remotePools[0]
	if len(remotePool) == 0 {
		return zero, fmt.Errorf("active pool's remote pool for chain %d is empty", remoteChainSelector)
	}

	decimalsReport, err := cldf_ops.ExecuteOperation(b, erc20.GetDecimals, remoteChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: remoteChainSelector,
		Address:       common.BytesToAddress(remoteToken),
	})
	if err != nil {
		return zero, fmt.Errorf(
			"failed to get decimals for remote token %s on chain %d (only ERC20-compatible tokens are supported for auto-migrate): %w",
			common.BytesToAddress(remoteToken).Hex(), remoteChainSelector, err,
		)
	}

	// NOTE: getRemotePools can return >1 address during a remote-side upgrade. `RemotePool` only needs
	// one valid (already-registered) address — ConfigureTokenPoolForRemoteChain re-reads and registers
	// the full list (deduping this one), so picking [0] drops nothing.
	return tokens.RemoteChainConfig[[]byte, string]{
		RemoteDecimals: decimalsReport.Output,
		RemoteToken:    remoteToken,
		RemotePool:     remotePool,
	}, nil
}
