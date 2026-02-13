package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type TPRLInput struct {
	Configs map[uint64]TPRLConfig `yaml:"configs" json:"configs"`
	MCMS    mcms.Input            `yaml:"mcms,omitempty" json:"mcms"`
}

type TPRLConfig struct {
	ChainSelector       uint64                 `yaml:"chain-selector" json:"chainSelector"`
	ChainAdapterVersion *semver.Version        `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
	TokenRef            datastore.AddressRef   `yaml:"token-ref" json:"tokenRef"`
	TokenPoolQualifier  string                 `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType            string                 `yaml:"pool-type" json:"poolType"`
	Inputs              map[uint64]TPRLRemotes `yaml:"inputs" json:"inputs"`
}

type TPRLRemotes struct {
	OutboundRateLimiterConfig RateLimiterConfig `yaml:"outbound-rate-limiter-config" json:"outboundRateLimiterConfig"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	InboundRateLimiterConfig RateLimiterConfig
	ChainSelector            uint64
	RemoteChainSelector      uint64
	TokenRef                 datastore.AddressRef
	TokenPoolQualifier       string
	PoolType                 string
	ExistingDataStore        datastore.DataStore
}

// SetTokenPoolRateLimits returns a changeset that sets rate limits for token pools on multiple chains.
func SetTokenPoolRateLimits() cldf.ChangeSetV2[TPRLInput] {
	return cldf.CreateChangeSet(setTokenPoolRateLimitsApply(), setTokenPoolRateLimitsVerify())
}

func setTokenPoolRateLimitsVerify() func(cldf.Environment, TPRLInput) error {
	return func(e cldf.Environment, cfg TPRLInput) error {
		for _, config := range cfg.Configs {
			for remoteSelector, input := range config.Inputs {
				if input.OutboundRateLimiterConfig.IsEnabled {
					if input.OutboundRateLimiterConfig.Capacity == nil || input.OutboundRateLimiterConfig.Rate == nil ||
						input.OutboundRateLimiterConfig.Capacity.Sign() <= 0 || input.OutboundRateLimiterConfig.Rate.Sign() <= 0 {
						return fmt.Errorf("outbound rate limiter config for remote chain %d is enabled but capacity or rate is nil", remoteSelector)
					}
				}
			}
		}
		return nil
	}
}

func setTokenPoolRateLimitsApply() func(cldf.Environment, TPRLInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg TPRLInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		tokenPoolRegistry := GetTokenAdapterRegistry()
		mcmsRegistry := changesets.GetRegistry()

		family, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, cfg.ChainAdapterVersion)
		if !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
		}
		for remoteSelector, inputs := range cfg.Inputs {
			inputs.ChainSelector = cfg.ChainSelector
			inputs.TokenRef = cfg.TokenRef
			inputs.TokenPoolQualifier = cfg.TokenPoolQualifier
			inputs.PoolType = cfg.PoolType
			inputs.RemoteChainSelector = remoteSelector
			inputs.ExistingDataStore = e.DataStore
			if !inputs.InboundRateLimiterConfig.IsEnabled {
				// If the rate limiter is not enabled, set capacity and rate to 0 to disable it on chain.
				inputs.InboundRateLimiterConfig.Capacity = big.NewInt(0)
				inputs.InboundRateLimiterConfig.Rate = big.NewInt(0)
			}
			if !inputs.OutboundRateLimiterConfig.IsEnabled {
				// If the rate limiter is not enabled, set capacity and rate to 0 to disable it on chain.
				inputs.OutboundRateLimiterConfig.Capacity = big.NewInt(0)
				inputs.OutboundRateLimiterConfig.Rate = big.NewInt(0)
			}
			rateLimitReport, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle, tokenPoolAdapter.SetTokenPoolRateLimits(), e.BlockChains, inputs)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to set rate limits for token pool %d on remote chain %d: %w", cfg.ChainSelector, remoteSelector, err)
			}
			batchOps = append(batchOps, rateLimitReport.Output.BatchOps...)
			reports = append(reports, rateLimitReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
