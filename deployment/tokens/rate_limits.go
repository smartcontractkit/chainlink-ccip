package tokens

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
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
	ChainAdapterVersion *semver.Version                        `yaml:"chain-adapter-version" json:"chainAdapterVersion"`
	TokenRef            datastore.AddressRef                   `yaml:"token-ref" json:"tokenRef"`
	TokenPoolRef        datastore.AddressRef                   `yaml:"token-pool-ref" json:"tokenPoolRef"`
	RemoteOutbounds     map[uint64]RateLimiterConfigFloatInput `yaml:"remote-outbounds" json:"remoteOutbounds"`
}

type TPRLRemotes struct {
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	ChainSelector             uint64
	RemoteChainSelector       uint64
	TokenRef                  datastore.AddressRef
	TokenPoolRef              datastore.AddressRef
	ExistingDataStore         datastore.DataStore
}

// SetTokenPoolRateLimits returns a changeset that sets rate limits for token pools on multiple chains.
func SetTokenPoolRateLimits() cldf.ChangeSetV2[TPRLInput] {
	return cldf.CreateChangeSet(setTokenPoolRateLimitsApply(), setTokenPoolRateLimitsVerify())
}

func setTokenPoolRateLimitsVerify() func(cldf.Environment, TPRLInput) error {
	return func(e cldf.Environment, cfg TPRLInput) error {
		for _, config := range cfg.Configs {
			for remoteSelector, input := range config.RemoteOutbounds {
				if input.IsEnabled {
					if input.Capacity <= 0 || input.Rate <= 0 {
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

		for selector, config := range cfg.Configs {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			tokenPoolAdapter, exists := tokenPoolRegistry.GetTokenAdapter(family, config.ChainAdapterVersion)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", family)
			}
			tokenPool, err := datastore_utils.FindAndFormatRef(e.DataStore, config.TokenPoolRef, selector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool ref on chain with selector %d: %w", selector, err)
			}
			tokenFull, err := datastore_utils.FindAndFormatRef(e.DataStore, config.TokenRef, selector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token ref on chain with selector %d: %w", selector, err)
			}
			tokenBytes, err := datastore_utils.FindAndFormatRef(e.DataStore, config.TokenRef, selector, tokenPoolAdapter.AddressRefToBytes)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token ref on chain with selector %d: %w", selector, err)
			}
			decimals, err := tokenPoolAdapter.DeriveTokenDecimals(e, selector, tokenPool, tokenBytes)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get token decimals for token on chain with selector %d: %w", selector, err)
			}
			for remoteSelector, inputs := range config.RemoteOutbounds {
				tprlRemote := TPRLRemotes{
					ChainSelector:       selector,
					RemoteChainSelector: remoteSelector,
					TokenRef:            tokenFull,
					TokenPoolRef:        tokenPool,
					ExistingDataStore:   e.DataStore,
				}

				// We derive the inbound rate limiter config from counterpart's outbound config for simplicity
				// My inbound rate limiter config should be the same as my counterpart's outbound config to avoid confusion,
				// and I can derive it programmatically so it's less error-prone for users than requiring them to specify both
				counterpart, ok := cfg.Configs[remoteSelector]
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no config provided for remote chain with selector %d", remoteSelector)
				}
				remoteInputs, ok := counterpart.RemoteOutbounds[selector]
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no inputs provided for remote chain with selector %d to chain with selector %d", selector, remoteSelector)
				}

				// remoteTokenPool, err := datastore_utils.FindAndFormatRef(e.DataStore, counterpart.TokenPoolRef, remoteSelector, datastore_utils.FullRef)
				// if err != nil {
				// 	return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool ref on chain with selector %d: %w", remoteSelector, err)
				// }
				// remoteToken, err := datastore_utils.FindAndFormatRef(e.DataStore, counterpart.TokenRef, remoteSelector, datastore_utils.FullRef)
				// if err != nil {
				// 	return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token ref on chain with selector %d: %w", remoteSelector, err)
				// }
				// remoteDecimals, err := tokenPoolAdapter.DeriveTokenDecimals(e, remoteSelector, remoteTokenPool, []byte(remoteToken.Address))
				// if err != nil {
				// 	return cldf.ChangesetOutput{}, fmt.Errorf("failed to get token decimals for token on chain with selector %d: %w", selector, err)
				// }
				if !inputs.IsEnabled {
					tprlRemote.OutboundRateLimiterConfig.IsEnabled = false
					tprlRemote.OutboundRateLimiterConfig.Capacity = big.NewInt(0)
					tprlRemote.OutboundRateLimiterConfig.Rate = big.NewInt(0)
				} else {
					// We scale the rate limiter configs by the token decimals to convert from
					// human-readable token amounts to the on-chain representation
					tprlRemote.OutboundRateLimiterConfig.IsEnabled = true
					tprlRemote.OutboundRateLimiterConfig.Capacity = big.NewInt(int64(inputs.Capacity * math.Pow10(int(decimals))))
					tprlRemote.OutboundRateLimiterConfig.Rate = big.NewInt(int64(inputs.Rate * math.Pow10(int(decimals))))
					if tprlRemote.OutboundRateLimiterConfig.Capacity.Cmp(big.NewInt(0)) <= 0 || tprlRemote.OutboundRateLimiterConfig.Rate.Cmp(big.NewInt(0)) <= 0 {
						return cldf.ChangesetOutput{}, fmt.Errorf("outbound rate limiter config for remote chain %d is enabled but capacity or rate is zero or negative after scaling by token decimals", remoteSelector)
					}
				}

				if !remoteInputs.IsEnabled {
					tprlRemote.InboundRateLimiterConfig.IsEnabled = false
					tprlRemote.InboundRateLimiterConfig.Capacity = big.NewInt(0)
					tprlRemote.InboundRateLimiterConfig.Rate = big.NewInt(0)
				} else {
					tprlRemote.InboundRateLimiterConfig.IsEnabled = true
					// We set the inbound capacity to be 1.1x the outbound capacity of the counterpart to avoid accidentally hitting the rate limit due to minor timing differences in refilling
					tprlRemote.InboundRateLimiterConfig.Capacity = big.NewInt(int64(remoteInputs.Capacity * 1.1 * math.Pow10(int(decimals))))
					tprlRemote.InboundRateLimiterConfig.Rate = big.NewInt(int64(remoteInputs.Rate * 1.1 * math.Pow10(int(decimals))))
					if tprlRemote.InboundRateLimiterConfig.Capacity.Cmp(big.NewInt(0)) <= 0 || tprlRemote.InboundRateLimiterConfig.Rate.Cmp(big.NewInt(0)) <= 0 {
						return cldf.ChangesetOutput{}, fmt.Errorf("inbound rate limiter config for remote chain %d is enabled but capacity or rate is zero or negative after scaling by token decimals", remoteSelector)
					}
				}
				rateLimitReport, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle, tokenPoolAdapter.SetTokenPoolRateLimits(), e.BlockChains, tprlRemote)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to set rate limits for token pool %d on remote chain %d: %w", selector, remoteSelector, err)
				}
				batchOps = append(batchOps, rateLimitReport.Output.BatchOps...)
				reports = append(reports, rateLimitReport.ExecutionReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
