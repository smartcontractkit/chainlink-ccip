package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
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

type RemoteOutbounds struct {
	// RateLimit the rate limiter config. For CCIP v2.0, there are two types of rate limits: custom and default.
	// Originally, this struct had two separate fields (one for each rate limit type) which isn't ideal because:
	// ---
	//   1. Only one of these rate limits is ever active at a time depending on the value of `allowedFinality`
	//   2. It forces the user to configure more than what they actually need onchain leading to higher risk of misconfiguration
	//   3. The user if forced to provide both rate limits up front - failure to provide one of them will lead to the unspecified one being reset by mistake
	// ---
	// For v2.0.0 it's possible to avoid these hazards entirely by collapsing both fields into one. With one
	// rate limit config, we simply need to read the allowedFinalityConfig from the chain and if it is zero,
	// then we update the default rate limit, otherwise we update the custom rate limit. This not only makes
	// things this safer and easier to run but allows us to maintain full backwards compatibility with 1.6.0
	// tooling as well.
	RateLimit RateLimiterConfigFloatInput `yaml:"rateLimit" json:"rateLimit"`
}

type TPRLConfig struct {
	ChainAdapterVersion      *semver.Version            `yaml:"chainAdapterVersion" json:"chainAdapterVersion"`
	TokenRef                 datastore.AddressRef       `yaml:"tokenRef" json:"tokenRef"`
	TokenPoolRef             datastore.AddressRef       `yaml:"tokenPoolRef" json:"tokenPoolRef"`
	AllowedFinalityConfig    finality.Config            `yaml:"allowedFinalityConfig" json:"allowedFinalityConfig"`
	RemoteOutbounds          map[uint64]RemoteOutbounds `yaml:"remoteOutbounds" json:"remoteOutbounds"`
	SkipIfMissingPermissions bool                       `yaml:"skipIfMissingPermissions" json:"skipIfMissingPermissions"`
}

type TPRLRemotes struct {
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	AllowedFinalityConfig     finality.Config
	ChainSelector             uint64
	RemoteChainSelector       uint64
	TokenRef                  datastore.AddressRef
	TokenPoolRef              datastore.AddressRef
	ExistingDataStore         datastore.DataStore

	// If true, the changeset will check if timelock or the deployer key has sufficient permissions to set rate limits
	// on the token pool. If both accounts are missing permissions (i.e. not the pool owner or rate limit admin), then
	// a warning will be logged and the changeset will NOT perform the rate limit update since it has a high chance of
	// failure. This flag is disabled by default so that it still allows flexibility for callers to schedule both rate
	// limit permission updates AND token pool rate limit updates in parallel / in the same batch. At the time of this
	// writing, this flag is only applicable for EVM, but can be extended to other chains in the future if needed.
	SkipIfMissingPermissions bool
}

// SetTokenPoolRateLimits returns a changeset that sets rate limits for token pools on multiple chains.
func SetTokenPoolRateLimits() cldf.ChangeSetV2[TPRLInput] {
	return cldf.CreateChangeSet(setTokenPoolRateLimitsApply(), setTokenPoolRateLimitsVerify())
}

func setTokenPoolRateLimitsVerify() func(cldf.Environment, TPRLInput) error {
	return func(e cldf.Environment, cfg TPRLInput) error {
		for _, config := range cfg.Configs {
			for remoteSelector, input := range config.RemoteOutbounds {
				if input.RateLimit.IsEnabled {
					if input.RateLimit.Capacity <= 0 || input.RateLimit.Rate <= 0 {
						return fmt.Errorf("outbound rate limiter config for remote chain %d is enabled but capacity or rate is invalid", remoteSelector)
					}
					if input.RateLimit.Rate > input.RateLimit.Capacity {
						return fmt.Errorf("outbound rate limiter config for remote chain %d has rate greater than capacity", remoteSelector)
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
					SkipIfMissingPermissions: config.SkipIfMissingPermissions,
					AllowedFinalityConfig:    config.AllowedFinalityConfig,
					ChainSelector:            selector,
					RemoteChainSelector:      remoteSelector,
					TokenRef:                 tokenFull,
					TokenPoolRef:             tokenPool,
					ExistingDataStore:        e.DataStore,
				}

				// We derive the inbound rate limiter config from counterpart's outbound config for simplicity
				// My inbound rate limiter config should be the same as my counterpart's outbound config to avoid confusion,
				// and I can derive it programmatically so it's less error-prone for users than requiring them to specify both
				counterpart, ok := cfg.Configs[remoteSelector]
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no config provided for remote chain with selector %d", remoteSelector)
				}
				counterpartFamily, err := chain_selectors.GetSelectorFamily(remoteSelector)
				if err != nil {
					return cldf.ChangesetOutput{}, err
				}
				counterPartAdapter, exists := tokenPoolRegistry.GetTokenAdapter(counterpartFamily, counterpart.ChainAdapterVersion)
				if !exists {
					return cldf.ChangesetOutput{}, fmt.Errorf("no TokenPoolAdapter registered for chain family '%s'", counterpartFamily)
				}
				remoteInputs, ok := counterpart.RemoteOutbounds[selector]
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no inputs provided for remote chain with selector %d to chain with selector %d", selector, remoteSelector)
				}

				remoteTokenPool, err := datastore_utils.FindAndFormatRef(e.DataStore, counterpart.TokenPoolRef, remoteSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool ref on chain with selector %d: %w", remoteSelector, err)
				}
				remoteTokenBytes, err := datastore_utils.FindAndFormatRef(e.DataStore, counterpart.TokenRef, remoteSelector, counterPartAdapter.AddressRefToBytes)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token ref on chain with selector %d: %w", selector, err)
				}
				remoteDecimals, err := counterPartAdapter.DeriveTokenDecimals(e, remoteSelector, remoteTokenPool, remoteTokenBytes)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get token decimals for token on chain with selector %d: %w", selector, err)
				}
				tprlRemote.OutboundRateLimiterConfig, tprlRemote.InboundRateLimiterConfig = GenerateTPRLConfigs(inputs.RateLimit, remoteInputs.RateLimit, decimals, remoteDecimals, family, tokenPool.Version)
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

// AI generated code below
// scaleFloatToBigInt converts a floating‑point value (capacity or rate)
// to a *big.Int* after applying two scalings:
//
//  1. decimalFactor = 10^decimals          (e.g. decimals = 6 → 1_000_000)
//  2. staticFactor  = 1 + extraPercent (e.g. extraPercent = 0.10 → 1.10)
//
// The function never overflows because all arithmetic is done with
// arbitrary‑precision types (big.Rat → big.Int).
func ScaleFloatToBigInt(value float64, decimals int, extraPercent float64) *big.Int {
	// -------------------------------------------------------------
	// Turn the float into an *exact* rational.
	// -------------------------------------------------------------
	//
	// big.NewFloat(value) creates a *big.Float* that holds the exact binary
	// representation of the float64.
	//
	// This path never fails for a finite float64, so we don’t need the
	// ok‑check that SetString requires.
	floatValue := new(big.Float).SetFloat64(value)

	// -------------------------------------------------------------
	// Multiply by 10^decimals (the “decimal” factor).
	// -------------------------------------------------------------
	//
	// Use big.Int.Exp so the power can be arbitrarily large.
	tenPow := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil) // 10^decimals
	floatValue.Mul(floatValue, new(big.Float).SetInt(tenPow))

	// -------------------------------------------------------------
	// Apply the optional static factor (e.g. +10 % → ×1.10).
	// -------------------------------------------------------------
	if extraPercent != 0 {
		num := big.NewFloat(1 + extraPercent)
		floatValue.Mul(floatValue, num)
	}

	// -------------------------------------------------------------
	// Return the *big.Int*.
	// -------------------------------------------------------------
	tmp := new(big.Int)
	floatValue.SetMode(big.AwayFromZero) // Round half up to avoid underestimating the rate limit
	out, _ := floatValue.Int(tmp)        // big.Float.Int sets tmp to the integer part of floatValue and returns tmp
	return out
}

func GenerateTPRLConfigs(
	outboundInput RateLimiterConfigFloatInput,
	inboundInput RateLimiterConfigFloatInput,
	decimals uint8,
	remoteDecimals uint8,
	chainFamily string,
	tokenPoolVersion *semver.Version,
) (RateLimiterConfig, RateLimiterConfig) {
	outboundConfig := RateLimiterConfig{}
	inboundConfig := RateLimiterConfig{}
	if !outboundInput.IsEnabled {
		outboundConfig.IsEnabled = false
		outboundConfig.Capacity = big.NewInt(0)
		outboundConfig.Rate = big.NewInt(0)
	} else {
		// We scale the rate limiter configs by the token decimals to convert from
		// human-readable token amounts to the on-chain representation
		outboundConfig.IsEnabled = true
		outboundConfig.Capacity = ScaleFloatToBigInt(outboundInput.Capacity, int(decimals), 0)
		outboundConfig.Rate = ScaleFloatToBigInt(outboundInput.Rate, int(decimals), 0)
	}

	if !inboundInput.IsEnabled {
		inboundConfig.IsEnabled = false
		inboundConfig.Capacity = big.NewInt(0)
		inboundConfig.Rate = big.NewInt(0)
	} else {
		// We set the inbound capacity to be 1.1x the outbound capacity of the counterpart to avoid accidentally hitting the rate limit due to minor timing differences in refilling
		scaleByDecimals := decimals
		// https://github.com/smartcontractkit/chainlink-deployments/blob/cce886554ca0587492955784381321ce817fb6bb/domains/ccip/shared/tokendefaults.go#L1904
		// Only old EVM pools need to scale by remote deciamls on inbound. Newer pools and non-EVM pools handle all conversions in local decimals.
		// This is a hack. Avoiding it would require refactoring the token pool adapters to handle rate limit configs in a more structured way instead of
		// just passing them as bytes through the registry, so for now we can live with this special case for old EVM pools since we're moving towards newer versions and non-EVM chains where this isn't an issue.
		if chainFamily == chain_selectors.FamilyEVM && tokenPoolVersion.LessThan(semver.MustParse("1.6.1")) {
			scaleByDecimals = remoteDecimals
		}
		inboundConfig.IsEnabled = true
		inboundConfig.Capacity = ScaleFloatToBigInt(inboundInput.Capacity, int(scaleByDecimals), .10)
		inboundConfig.Rate = ScaleFloatToBigInt(inboundInput.Rate, int(scaleByDecimals), .10)
	}
	return outboundConfig, inboundConfig
}
