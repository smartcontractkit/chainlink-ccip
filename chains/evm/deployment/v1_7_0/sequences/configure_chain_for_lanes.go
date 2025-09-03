package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_aggregator"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/commit_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter_v2"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type CommitOnRampDestChainConfig struct {
	// Whether or not to allow traffic TO the remote chain
	AllowlistEnabled bool
	// Addresses that are allowed to send messages TO the remote chain
	AddedAllowlistedSenders []common.Address
	// Addresses that are no longer allowed to send messages TO the remote chain
	RemovedAllowlistedSenders []common.Address
}

type RemoteChainConfig struct {
	// Whether or not to allow traffic FROM this remote chain
	AllowTrafficFrom bool
	// The address on the remote chain from which the message is emitted
	// For example, on EVM chains, this is the CCVProxy
	CCIPMessageSource []byte
	// The default CCVs that will be applied to messages FROM this remote chain if no receiver is specified
	DefaultCCVOffRamps []common.Address
	// Any CCVs that must always be used for messages FROM this remote chain
	LaneMandatedCCVOffRamps []common.Address
	// The CCVs that will be used for messages TO this remote chain if none are specified
	DefaultCCVOnRamps []common.Address
	// The CCVs that will always be applied to messages TO this remote chain
	LaneMandatedCCVOnRamps []common.Address
	// The executor that will be used for messages TO this remote chain if none is specified
	DefaultExecutor common.Address
	// CommitOnRampDestChainConfig configures the CommitOnRamp for this remote chain
	CommitOnRampDestChainConfig CommitOnRampDestChainConfig
	// FeeQuoterDestChainConfig configures the FeeQuoter for this remote chain
	FeeQuoterDestChainConfig fee_quoter_v2.DestChainConfig
}

type ConfigureChainForLanesInput struct {
	// The selector of the EVM chain being configured
	ChainSelector uint64
	// The router on the EVM chain being configured
	// We assume that all connections will use the same router, either test or production
	Router common.Address
	// The CCVProxy on the EVM chain being configured
	// Similarly, we assume that all connections will use the same CCVProxy
	CCVProxy common.Address
	// The CommitOnRamp on the EVM chain being configured
	CommitOnRamp common.Address
	// The FeeQuoter on the EVM chain being configured
	FeeQuoter common.Address
	// The CCVAggregator on the EVM chain being configured
	CCVAggregator common.Address
	// The configuration of each remote chain to configure
	RemoteChains map[uint64]RemoteChainConfig
}

type ConfigureChainForLanesOutput struct {
	Writes []call.WriteOutput
}

var ConfigureChainForLanes = cldf_ops.NewSequence(
	"configure-chain-for-lanes",
	semver.MustParse("1.7.0"),
	"Configures an EVM chain as a source & destination for multiple remote chains",
	func(b operations.Bundle, chain evm.Chain, input ConfigureChainForLanesInput) (output ConfigureChainForLanesOutput, err error) {
		writes := make([]call.WriteOutput, 0, 6) // 6 = number of ExecuteOperation calls

		// Create inputs for each operation
		ccvAggregatorArgs := make([]ccv_aggregator.SourceChainConfigArgs, 0, len(input.RemoteChains))
		ccvProxyArgs := make([]ccv_proxy.DestChainConfigArgs, 0, len(input.RemoteChains))
		commitOnRampDestConfigArgs := make([]commit_onramp.DestChainConfigArgs, 0, len(input.RemoteChains))
		commitOnRampAllowlistArgs := make([]commit_onramp.AllowlistConfigArgs, 0, len(input.RemoteChains))
		feeQuoterArgs := make([]fee_quoter_v2.DestChainConfigArgs, 0, len(input.RemoteChains))
		onRampAdds := make([]router.OnRamp, 0, len(input.RemoteChains))
		offRampAdds := make([]router.OffRamp, 0, len(input.RemoteChains))
		for remoteSelector, remoteConfig := range input.RemoteChains {
			ccvAggregatorArgs = append(ccvAggregatorArgs, ccv_aggregator.SourceChainConfigArgs{
				Router:              input.Router,
				SourceChainSelector: remoteSelector,
				IsEnabled:           remoteConfig.AllowTrafficFrom,
				OnRamp:              remoteConfig.CCIPMessageSource,
				DefaultCCV:          remoteConfig.DefaultCCVOffRamps,
				LaneMandatedCCVs:    remoteConfig.LaneMandatedCCVOffRamps,
			})
			ccvProxyArgs = append(ccvProxyArgs, ccv_proxy.DestChainConfigArgs{
				Router:            input.Router,
				DestChainSelector: remoteSelector,
				DefaultCCVs:       remoteConfig.DefaultCCVOnRamps,
				LaneMandatedCCVs:  remoteConfig.LaneMandatedCCVOnRamps,
				DefaultExecutor:   remoteConfig.DefaultExecutor,
			})
			commitOnRampDestConfigArgs = append(commitOnRampDestConfigArgs, commit_onramp.DestChainConfigArgs{
				CcvProxy:          input.CCVProxy,
				DestChainSelector: remoteSelector,
				AllowlistEnabled:  remoteConfig.CommitOnRampDestChainConfig.AllowlistEnabled,
			})
			commitOnRampAllowlistArgs = append(commitOnRampAllowlistArgs, commit_onramp.AllowlistConfigArgs{
				AllowlistEnabled:          remoteConfig.CommitOnRampDestChainConfig.AllowlistEnabled,
				AddedAllowlistedSenders:   remoteConfig.CommitOnRampDestChainConfig.AddedAllowlistedSenders,
				RemovedAllowlistedSenders: remoteConfig.CommitOnRampDestChainConfig.RemovedAllowlistedSenders,
			})
			feeQuoterArgs = append(feeQuoterArgs, fee_quoter_v2.DestChainConfigArgs{
				DestChainSelector: remoteSelector,
				DestChainConfig:   remoteConfig.FeeQuoterDestChainConfig,
			})
			onRampAdds = append(onRampAdds, router.OnRamp{
				DestChainSelector: remoteSelector,
				OnRamp:            input.CCVProxy,
			})
			offRampAdds = append(offRampAdds, router.OffRamp{
				SourceChainSelector: remoteSelector,
				OffRamp:             input.CCVAggregator,
			})
		}

		// ApplySourceChainConfigUpdates on CCVAggregator
		ccvAggregatorReport, err := cldf_ops.ExecuteOperation(b, ccv_aggregator.ApplySourceChainConfigUpdates, chain, call.Input[[]ccv_aggregator.SourceChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CCVAggregator,
			Args:          ccvAggregatorArgs,
		})
		if err != nil {
			return ConfigureChainForLanesOutput{}, fmt.Errorf("failed to apply source chain config updates to CCVAggregator(%s) on chain %s: %w", input.CCVAggregator, chain, err)
		}
		writes = append(writes, ccvAggregatorReport.Output)

		// ApplyDestChainConfigUpdates on CCVProxy
		ccvProxyReport, err := cldf_ops.ExecuteOperation(b, ccv_proxy.ApplyDestChainConfigUpdates, chain, call.Input[[]ccv_proxy.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CCVProxy,
			Args:          ccvProxyArgs,
		})
		if err != nil {
			return ConfigureChainForLanesOutput{}, fmt.Errorf("failed to apply dest chain config updates to CCVProxy(%s) on chain %s: %w", input.CCVProxy, chain, err)
		}
		writes = append(writes, ccvProxyReport.Output)

		// ApplyDestChainConfigUpdates on CommitOnRamp
		commitOnRampReport, err := cldf_ops.ExecuteOperation(b, commit_onramp.ApplyDestChainConfigUpdates, chain, call.Input[[]commit_onramp.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CommitOnRamp,
			Args:          commitOnRampDestConfigArgs,
		})
		if err != nil {
			return ConfigureChainForLanesOutput{}, fmt.Errorf("failed to apply dest chain config updates to CommitOnRamp(%s) on chain %s: %w", input.CommitOnRamp, chain, err)
		}
		writes = append(writes, commitOnRampReport.Output)

		// ApplyAllowlistUpdates on CommitOnRamp
		commitOnRampAllowlistReport, err := cldf_ops.ExecuteOperation(b, commit_onramp.ApplyAllowlistUpdates, chain, call.Input[[]commit_onramp.AllowlistConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.CommitOnRamp,
			Args:          commitOnRampAllowlistArgs,
		})
		if err != nil {
			return ConfigureChainForLanesOutput{}, fmt.Errorf("failed to apply allowlist updates to CommitOnRamp(%s) on chain %s: %w", input.CommitOnRamp, chain, err)
		}
		writes = append(writes, commitOnRampAllowlistReport.Output)

		// ApplyDestChainConfigUpdates on FeeQuoter
		feeQuoterReport, err := cldf_ops.ExecuteOperation(b, fee_quoter_v2.ApplyDestChainConfigUpdates, chain, call.Input[[]fee_quoter_v2.DestChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       input.FeeQuoter,
			Args:          feeQuoterArgs,
		})
		if err != nil {
			return ConfigureChainForLanesOutput{}, fmt.Errorf("failed to apply dest chain config updates to FeeQuoter(%s) on chain %s: %w", input.FeeQuoter, chain, err)
		}
		writes = append(writes, feeQuoterReport.Output)

		// ApplyRampUpdates on Router
		routerReport, err := cldf_ops.ExecuteOperation(b, router.ApplyRampUpdates, chain, call.Input[router.ApplyRampsUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.Router,
			Args: router.ApplyRampsUpdatesArgs{
				OnRampUpdates:  onRampAdds,
				OffRampRemoves: []router.OffRamp{}, // removals should be processed by a separate sequence responsible for disconnecting lanes
				OffRampAdds:    offRampAdds,
			},
		})
		if err != nil {
			return ConfigureChainForLanesOutput{}, fmt.Errorf("failed to apply ramp updates to Router(%s) on chain %s: %w", input.Router, chain, err)
		}
		writes = append(writes, routerReport.Output)

		return ConfigureChainForLanesOutput{
			Writes: writes,
		}, nil
	},
)
