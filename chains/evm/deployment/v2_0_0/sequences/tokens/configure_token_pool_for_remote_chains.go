package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
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
}

// ConfigureTokenPoolForRemoteChains calls ConfigureTokenPoolForRemoteChain for each remote chain.
//
// TODO: re-enable the active-pool upgrade-safety validation that previously lived here. It compared
// input.RemoteChains against getSupportedChains() of the currently-registered pool on the
// TokenAdminRegistry. That check fired incorrectly when configuring an auxiliary pool (e.g. the
// CCTP-through-CCV pool) whose chain set is intentionally a strict subset of the registered pool's
// (non-EVM remotes like Solana are excluded from the CCV pool). It needs to be reintroduced in a
// way that only runs when the pool being configured is actually being installed as the new active
// pool — see configure_cctp_chain_for_lanes.go where input.TokenPoolAddress (CCV pool) intentionally
// differs from the RegistryTokenPoolAddress passed to ConfigureTokenForTransfers.
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"configure-token-pool-for-remote-chains",
	semver.MustParse("2.0.0"),
	"Configures a token pool for all configured remote chains.",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainsInput) (output sequences.OnChainOutput, err error) {
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
