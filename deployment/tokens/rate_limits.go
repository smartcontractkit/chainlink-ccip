package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type RateLimiterConfigInput struct {
	ChainSelector       uint64                             `yaml:"chain-selector" json:"chainSelector"`
	ChainAdapterVersion *semver.Version                    `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
	TokenSymbol         string                             `yaml:"token-symbol" json:"tokenSymbol"`
	TokenPoolQualifier  string                             `yaml:"token-pool-qualifier" json:"tokenPoolQualifier"`
	PoolType            string                             `yaml:"pool-type" json:"poolType"`
	Inputs              map[uint64]RateLimiterConfigInputs `yaml:"inputs" json:"inputs"`
	MCMS                mcms.Input                         `yaml:"mcms,omitempty" json:"mcms"`
}

type RateLimiterConfigInputs struct {
	InboundRateLimiterConfig  RateLimiterConfig `yaml:"inbound-rate-limiter-config" json:"inboundRateLimiterConfig"`
	OutboundRateLimiterConfig RateLimiterConfig `yaml:"outbound-rate-limiter-config" json:"outboundRateLimiterConfig"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector       uint64
	RemoteChainSelector uint64
	TokenSymbol         string
	TokenPoolQualifier  string
	PoolType            string
	ExistingDataStore   datastore.DataStore
}

// SetTokenPoolRateLimits returns a changeset that sets rate limits for token pools on multiple chains.
func SetTokenPoolRateLimits() cldf.ChangeSetV2[RateLimiterConfigInput] {
	return cldf.CreateChangeSet(setTokenPoolRateLimitsApply(), setTokenPoolRateLimitsVerify())
}

func setTokenPoolRateLimitsVerify() func(cldf.Environment, RateLimiterConfigInput) error {
	return func(e cldf.Environment, cfg RateLimiterConfigInput) error {
		for _, input := range cfg.Inputs {
			if input.InboundRateLimiterConfig.IsEnabled {
				if input.InboundRateLimiterConfig.Capacity == nil || input.InboundRateLimiterConfig.Rate == nil ||
					input.InboundRateLimiterConfig.Capacity.Sign() <= 0 || input.InboundRateLimiterConfig.Rate.Sign() <= 0 {
					return fmt.Errorf("inbound rate limiter config for remote chain %d is enabled but capacity or rate is nil", input.RemoteChainSelector)
				}
			} else {
				// if the rate limiter is not enabled, the capacity and rate should be nil or zero
				if (input.InboundRateLimiterConfig.Capacity != nil && input.InboundRateLimiterConfig.Capacity.Sign() > 0) ||
					(input.InboundRateLimiterConfig.Rate != nil && input.InboundRateLimiterConfig.Rate.Sign() > 0) {
					return fmt.Errorf("inbound rate limiter config for remote chain %d is not enabled but capacity or rate is set", input.RemoteChainSelector)
				}
			}
			if input.OutboundRateLimiterConfig.IsEnabled {
				if input.OutboundRateLimiterConfig.Capacity == nil || input.OutboundRateLimiterConfig.Rate == nil ||
					input.OutboundRateLimiterConfig.Capacity.Sign() <= 0 || input.OutboundRateLimiterConfig.Rate.Sign() <= 0 {
					return fmt.Errorf("outbound rate limiter config for remote chain %d is enabled but capacity or rate is nil", input.RemoteChainSelector)
				}
			} else {
				// if the rate limiter is not enabled, the capacity and rate should be nil or zero
				if (input.OutboundRateLimiterConfig.Capacity != nil && input.OutboundRateLimiterConfig.Capacity.Sign() > 0) ||
					(input.OutboundRateLimiterConfig.Rate != nil && input.OutboundRateLimiterConfig.Rate.Sign() > 0) {
					return fmt.Errorf("outbound rate limiter config for remote chain %d is not enabled but capacity or rate is set", input.RemoteChainSelector)
				}
			}
		}
		return nil
	}
}

func setTokenPoolRateLimitsApply() func(cldf.Environment, RateLimiterConfigInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg RateLimiterConfigInput) (cldf.ChangesetOutput, error) {
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
			inputs.TokenSymbol = cfg.TokenSymbol
			inputs.TokenPoolQualifier = cfg.TokenPoolQualifier
			inputs.PoolType = cfg.PoolType
			inputs.RemoteChainSelector = remoteSelector
			inputs.ExistingDataStore = e.DataStore
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
