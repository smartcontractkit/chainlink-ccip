package tokens

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
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
}

// ConfigureTokenPoolForRemoteChains runs the supported-chains check once (when registry/token set) then calls
// ConfigureTokenPoolForRemoteChain for each remote chain.
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"configure-token-pool-for-remote-chains",
	semver.MustParse("1.7.0"),
	"Configures a token pool for all configured remote chains; validates active pool supported chains when upgrading.",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainsInput) (output sequences.OnChainOutput, err error) {
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
				supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, evm_contract.FunctionInput[any]{
					ChainSelector: input.ChainSelector,
					Address:       activePool,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains from active pool %s: %w", activePool, err)
				}
				supportedChains := supportedChainsReport.Output
				for _, sel := range supportedChains {
					if _, ok := input.RemoteChains[sel]; !ok {
						sort.Slice(supportedChains, func(i, j int) bool { return supportedChains[i] < supportedChains[j] })
						configuredSelectors := make([]uint64, 0, len(input.RemoteChains))
						for s := range input.RemoteChains {
							configuredSelectors = append(configuredSelectors, s)
						}
						sort.Slice(configuredSelectors, func(i, j int) bool { return configuredSelectors[i] < configuredSelectors[j] })
						return sequences.OnChainOutput{}, fmt.Errorf("remoteChains must include all active pool supported chains: pool has %v, remoteChains has %v (missing chain selector %d)",
							supportedChains, configuredSelectors, sel)
					}
				}
			}
		}
		ops := make([]mcms_types.BatchOperation, 0)
		supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, evm_contract.FunctionInput[any]{
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
				ChainSelector:                input.ChainSelector,
				TokenPoolAddress:             input.TokenPoolAddress,
				AdvancedPoolHooks:            input.AdvancedPoolHooks,
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
