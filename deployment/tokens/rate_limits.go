package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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

type RateLimitConfig struct {
	RateLimit    RateLimiterConfigFloatInput `yaml:"rateLimit" json:"rateLimit"`
	FastFinality bool                        `yaml:"fastFinality" json:"fastFinality"`
}

// RemoteOutbounds holds the outbound rate limit configuration for a given remote chain. RateLimit is a
// backwards-compatible alias for the default bucket (FastFinality = false); Outbounds carries explicit
// per-bucket rows.
//
// Fast-finality buckets (Outbounds entries with FastFinality = true) target CCIP v2; pre-v2 adapters
// ignore those rows for legacy TPRL writes (they only touch default-lane scalars).
type RemoteOutbounds struct {
	// RateLimit is the backwards compatible alias for the default rate limit bucket. This field has
	// lower precedence than Outbounds: when resolving the default bucket, Outbounds will be checked
	// first for any entry with FastFinality = false. If such an entry exists then it'll be used for
	// the default bucket, otherwise this field will be used when non-nil.
	RateLimit *RateLimiterConfigFloatInput `yaml:"rateLimit" json:"rateLimit"`

	// Outbounds is the primary source of truth for outbound rate limit configuration and can be used
	// for all token pool versions. The slice should only contain up to two entries (one per finality
	// flag). Pre-v2 adapters ignore fast-finality rows and only apply default buckets if they exist.
	Outbounds []RateLimitConfig `yaml:"outbounds" json:"outbounds"`

	// OutboundOnly, when true, configures only the outbound side of this lane on the local chain.
	// The counterpart chain's TPRLConfig need not provide a matching RemoteOutbounds entry; the
	// changeset will read the counterpart's on-chain inbound for the same lane and reject the
	// update if it is not at least 110% of the new outbound (capacity AND rate). The counterpart
	// chain's TPRLConfig must still exist with TokenPoolRef + TokenRef so the changeset can
	// resolve the counterpart pool and its decimals for the validation read.
	OutboundOnly bool `yaml:"outboundOnly,omitempty" json:"outboundOnly,omitempty"`
}

// DefaultBucket gets the default lane (FastFinality = false) RateLimitConfig from Outbounds
// falling back to the legacy RateLimit alias if no such Outbounds entry exists. The boolean
// return indicates whether a bucket was found.
func (ro RemoteOutbounds) DefaultBucket() (RateLimitConfig, bool) {
	return ro.BucketForFinality(false)
}

// FastFinalityBucket gets the fast-finality lane (FastFinality = true) RateLimitConfig from
// Outbounds. The boolean return indicates whether a bucket was found. This is only relevant
// for CCIP v2 adapters.
func (ro RemoteOutbounds) FastFinalityBucket() (RateLimitConfig, bool) {
	return ro.BucketForFinality(true)
}

// BucketForFinality fetches the outbound rate limit bucket for the given FastFinality setting.
// Callers should prefer [RemoteOutbounds.DefaultBucket] / [RemoteOutbounds.FastFinalityBucket]
// at fixed call sites; this method is useful when the lane is parameterized (e.g. in tests).
func (ro RemoteOutbounds) BucketForFinality(fastFinality bool) (RateLimitConfig, bool) {
	for _, ob := range ro.Outbounds {
		if ob.FastFinality == fastFinality {
			return ob, true
		}
	}
	if !fastFinality && ro.RateLimit != nil {
		return RateLimitConfig{RateLimit: *ro.RateLimit, FastFinality: false}, true
	}
	return RateLimitConfig{}, false
}

// Validate checks structural rules on operator input: validates the legacy RateLimit alias when set, at most
// two Outbounds buckets, and at most one per FastFinality value—aligned with TPRL verify preconditions.
func (ro RemoteOutbounds) Validate() error {
	if ro.RateLimit != nil {
		if err := ro.RateLimit.Validate(); err != nil {
			return fmt.Errorf("rate limit alias: %w", err)
		}
	}
	if len(ro.Outbounds) > 2 {
		return fmt.Errorf("at most two rate limit buckets allowed")
	}

	defaultCount, fastFinCount := 0, 0
	for _, rl := range ro.Outbounds {
		if err := rl.RateLimit.Validate(); err != nil {
			return fmt.Errorf("rate limit bucket: %w", err)
		}
		if rl.FastFinality {
			fastFinCount++
		} else {
			defaultCount++
		}
	}
	if defaultCount > 1 {
		return fmt.Errorf("multiple rate limit buckets with fastFinality=false")
	}
	if fastFinCount > 1 {
		return fmt.Errorf("multiple rate limit buckets with fastFinality=true")
	}
	return nil
}

type TPRLConfig struct {
	ChainAdapterVersion      *semver.Version            `yaml:"chainAdapterVersion" json:"chainAdapterVersion"`
	TokenRef                 datastore.AddressRef       `yaml:"tokenRef" json:"tokenRef"`
	TokenPoolRef             datastore.AddressRef       `yaml:"tokenPoolRef" json:"tokenPoolRef"`
	AllowedFinalityConfig    finality.Config            `yaml:"allowedFinalityConfig" json:"allowedFinalityConfig"`
	RemoteOutbounds          map[uint64]RemoteOutbounds `yaml:"remoteOutbounds" json:"remoteOutbounds"`
	SkipIfMissingPermissions bool                       `yaml:"skipIfMissingPermissions" json:"skipIfMissingPermissions"`
}

// TPRLRateLimitBucket is one outbound/inbound pair after scaling for a given TokenPool fastFinality bucket.
type TPRLRateLimitBucket struct {
	OutboundRateLimiterConfig RateLimiterConfig
	InboundRateLimiterConfig  RateLimiterConfig
	FastFinality              bool
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

	// RateLimitBuckets carries built TPRL buckets (default and optional fast-finality). CCIP v2 adapters apply
	// all entries. Pre-v2 EVM pool adapters apply default-lane outbound/inbound scalars only when a default RL
	// bucket exists otherwise a warning is emitted (fast-finality-only inputs are ignored).
	RateLimitBuckets []TPRLRateLimitBucket

	// If true, the changeset will check if timelock or the deployer key has sufficient permissions to set rate limits
	// on the token pool. If both accounts are missing permissions (i.e. not the pool owner or rate limit admin), then
	// a warning will be logged and the changeset will NOT perform the rate limit update since it has a high chance of
	// failure. This flag is disabled by default so that it still allows flexibility for callers to schedule both rate
	// limit permission updates AND token pool rate limit updates in parallel / in the same batch. At the time of this
	// writing, this flag is only applicable for EVM, but can be extended to other chains in the future if needed.
	SkipIfMissingPermissions bool
}

// GetBucketForFinality gets the TPRLRateLimitBucket for the given finality flag.
// Returns a boolean indicating whether a bucket was found for that finality flag.
func (r TPRLRemotes) GetBucketForFinality(fastFinality bool) (TPRLRateLimitBucket, bool) {
	for _, b := range r.RateLimitBuckets {
		if b.FastFinality == fastFinality {
			return b, true
		}
	}
	return TPRLRateLimitBucket{}, false
}

// SetTokenPoolRateLimits returns a changeset that sets rate limits for token pools on multiple chains.
func SetTokenPoolRateLimits() cldf.ChangeSetV2[TPRLInput] {
	return cldf.CreateChangeSet(setTokenPoolRateLimitsApply(), setTokenPoolRateLimitsVerify())
}

func setTokenPoolRateLimitsVerify() func(cldf.Environment, TPRLInput) error {
	return func(e cldf.Environment, cfg TPRLInput) error {
		for localSelector, config := range cfg.Configs {
			for remoteSelector, localOutbound := range config.RemoteOutbounds {
				if err := localOutbound.Validate(); err != nil {
					return fmt.Errorf("outbound rate limiter config for remote chain %d: %w", remoteSelector, err)
				}

				// Counterpart config must always exist: when OutboundOnly is set we still need the
				// counterpart's TokenPoolRef/TokenRef to resolve its pool address and decimals for
				// the on-chain inbound validation.
				remote, ok := cfg.Configs[remoteSelector]
				if !ok {
					return fmt.Errorf("no config provided for remote chain with selector %d", remoteSelector)
				}

				if localOutbound.OutboundOnly {
					// In OutboundOnly mode the counterpart's RemoteOutbounds[localSelector] is not
					// required and symmetry checks do not apply: chain B is read-only for the
					// changeset and its rate limit will be validated against on-chain state at
					// apply time, not user input.
					if _, ok := localOutbound.DefaultBucket(); !ok {
						if _, ffOK := localOutbound.FastFinalityBucket(); !ffOK {
							return fmt.Errorf("outbound-only lane from chain %d to %d has no outbound buckets", localSelector, remoteSelector)
						}
					}
					continue
				}

				remoteOutbound, ok := remote.RemoteOutbounds[localSelector]
				if !ok {
					return fmt.Errorf("no inputs provided for remote chain with selector %d to chain with selector %d", remoteSelector, localSelector)
				}

				// Rate limit must be valid on both sides
				if err := remoteOutbound.Validate(); err != nil {
					return fmt.Errorf("outbound rate limiter config from chain %d toward %d: %w", remoteSelector, localSelector, err)
				}

				// Fast-finality rate limit must either be absent on both sides or present on both sides; it cannot be asymmetric
				_, remoteFastFinalityRateLimitExists := remoteOutbound.FastFinalityBucket()
				_, localFastFinalityRateLimitExists := localOutbound.FastFinalityBucket()
				if localFastFinalityRateLimitExists != remoteFastFinalityRateLimitExists {
					return fmt.Errorf(
						"both local and remote buckets must be provided for fastFinality=true or neither can be provided for chain selector %d and remote selector %d",
						localSelector, remoteSelector,
					)
				}

				// Default rate limit must either be absent on both sides or present on both sides; it cannot be asymmetric
				_, remoteDefaultRateLimitExists := remoteOutbound.DefaultBucket()
				_, localDefaultRateLimitExists := localOutbound.DefaultBucket()
				if localDefaultRateLimitExists != remoteDefaultRateLimitExists {
					return fmt.Errorf(
						"both local and remote buckets must be provided for fastFinality=false or neither can be provided for chain selector %d and remote selector %d",
						localSelector, remoteSelector,
					)
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
			tokenPoolAdapter, family, tokenPool, tokenFull, err := ResolveAdapterAndRefs(e, tokenPoolRegistry, selector, config.TokenPoolRef, config.TokenRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool and token refs on chain with selector %d: %w", selector, err)
			}
			tokenBytes, err := tokenPoolAdapter.AddressRefToBytes(tokenFull)
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

				counterpart.TokenPoolRef, err = TryNormalizeAddressRef(remoteSelector, counterpart.TokenPoolRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to normalize token pool ref on chain with selector %d: %w", remoteSelector, err)
				}

				counterPartAdapter, _, remoteTokenPool, remoteToken, err := ResolveAdapterAndRefs(e, tokenPoolRegistry, remoteSelector, counterpart.TokenPoolRef, counterpart.TokenRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token pool and token refs on chain with selector %d: %w", remoteSelector, err)
				}
				remoteTokenBytes, err := counterPartAdapter.AddressRefToBytes(remoteToken)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve token ref on chain with selector %d: %w", remoteSelector, err)
				}
				remoteDecimals, err := counterPartAdapter.DeriveTokenDecimals(e, remoteSelector, remoteTokenPool, remoteTokenBytes)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get token decimals for token on chain with selector %d: %w", remoteSelector, err)
				}

				// In OutboundOnly mode the counterpart's RemoteOutbounds[selector] is not required; we
				// pass an empty RemoteOutbounds into the build so no inbound config is generated and
				// then overwrite each bucket's InboundRateLimiterConfig with the local pool's current
				// on-chain inbound below.
				var remoteInputs RemoteOutbounds
				if !inputs.OutboundOnly {
					remoteInputs, ok = counterpart.RemoteOutbounds[selector]
					if !ok {
						return cldf.ChangesetOutput{}, fmt.Errorf("no inputs provided for remote chain with selector %d to chain with selector %d", selector, remoteSelector)
					}
				}

				tprlRemote.OutboundRateLimiterConfig, tprlRemote.InboundRateLimiterConfig, tprlRemote.RateLimitBuckets, err = buildTPRLRemotesForSetRateLimitsLane(
					family,
					tokenPool,
					selector,
					decimals,
					inputs,
					remoteSelector,
					remoteDecimals,
					remoteInputs,
				)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to generate TPRL configs for chain selector %d and remote selector %d: %w", selector, remoteSelector, err)
				}

				if inputs.OutboundOnly {
					// The local pool setter is atomic over outbound+inbound, so we read the current
					// on-chain inbound for each finality bucket and pass it through unchanged.
					localReader, ok := tokenPoolAdapter.(RateLimitReaderAdapter)
					if !ok {
						return cldf.ChangesetOutput{}, fmt.Errorf(
							"adapter for local chain selector %d (family %s, version %v) does not implement RateLimitReaderAdapter; outbound-only mode unsupported for this lane",
							selector, family, tokenPool.Version,
						)
					}
					for i, bucket := range tprlRemote.RateLimitBuckets {
						current, err := localReader.GetOnchainInboundRateLimit(
							e, selector, tokenPool, tokenFull, remoteSelector, bucket.FastFinality,
						)
						if err != nil {
							return cldf.ChangesetOutput{}, fmt.Errorf("failed to read current inbound rate limit for pass-through on outbound-only update (chain %d, remote %d, fastFinality=%v): %w", selector, remoteSelector, bucket.FastFinality, err)
						}
						tprlRemote.RateLimitBuckets[i].InboundRateLimiterConfig = current
					}

					if err := validateOutboundOnlyAgainstCounterpartInbound(
						e,
						counterPartAdapter,
						counterpartFamily,
						remoteTokenPool,
						counterpart.TokenRef,
						remoteSelector,
						remoteDecimals,
						selector,
						decimals,
						inputs,
					); err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("outbound-only validation failed for chain selector %d and remote selector %d: %w", selector, remoteSelector, err)
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

// buildTPRLRemotesForSetRateLimitsLane constructs the TPRL rate limiter configs for both outbound and
// inbound directions for a given local-remote chain pair based on user input. It validates that both
// sides of the lane are configured correctly (i.e. both must specify a bucket for a given finality flag,
// or neither can specify a bucket for that finality flag) and returns an error if not. It also returns
// a slice of TPRLRateLimitBucket which includes one entry per configured lane (default and optional
// fast-finality) with the outbound/inbound configs already scaled and converted to big.Int.
func buildTPRLRemotesForSetRateLimitsLane(
	chainFamily string,
	tokenPoolRef datastore.AddressRef,
	localSelector uint64,
	localDecimals uint8,
	localOutbounds RemoteOutbounds,
	remoteSelector uint64,
	remoteDecimals uint8,
	remoteOutbounds RemoteOutbounds,
) (RateLimiterConfig, RateLimiterConfig, []TPRLRateLimitBucket, error) {
	buckets := []TPRLRateLimitBucket{}
	outboundOnly := localOutbounds.OutboundOnly

	localFastFinalityBucket, localFastFinalityExists := localOutbounds.FastFinalityBucket()
	remoteFastFinalityBucket, remoteFastFinalityExists := remoteOutbounds.FastFinalityBucket()
	if !outboundOnly && localFastFinalityExists != remoteFastFinalityExists {
		return RateLimiterConfig{}, RateLimiterConfig{}, nil, fmt.Errorf(
			"both local and remote buckets must be provided for fastFinality=true or neither can be provided for chain selector %d and remote selector %d",
			localSelector, remoteSelector,
		)
	}

	localDefaultBucket, localDefaultExists := localOutbounds.DefaultBucket()
	remoteDefaultBucket, remoteDefaultExists := remoteOutbounds.DefaultBucket()
	if !outboundOnly && localDefaultExists != remoteDefaultExists {
		return RateLimiterConfig{}, RateLimiterConfig{}, nil, fmt.Errorf(
			"both local and remote buckets must be provided for fastFinality=false or neither can be provided for chain selector %d and remote selector %d",
			localSelector, remoteSelector,
		)
	}

	// When outboundOnly, inbound is left zero here; the changeset overwrites each bucket's
	// InboundRateLimiterConfig with the local pool's current on-chain inbound before dispatch.
	var fastFinalityOutboundRL, fastFinalityInboundRL RateLimiterConfig
	if localFastFinalityExists {
		if outboundOnly {
			fastFinalityOutboundRL, _ = GenerateTPRLConfigs(
				localFastFinalityBucket.RateLimit, RateLimiterConfigFloatInput{}, localDecimals, remoteDecimals,
				chainFamily, tokenPoolRef.Version, tokenPoolRef.Type.String(),
			)
		} else {
			fastFinalityOutboundRL, fastFinalityInboundRL = GenerateTPRLConfigs(
				localFastFinalityBucket.RateLimit, remoteFastFinalityBucket.RateLimit, localDecimals, remoteDecimals,
				chainFamily, tokenPoolRef.Version, tokenPoolRef.Type.String(),
			)
		}
		buckets = append(buckets, TPRLRateLimitBucket{
			FastFinality:              true,
			OutboundRateLimiterConfig: fastFinalityOutboundRL,
			InboundRateLimiterConfig:  fastFinalityInboundRL,
		})
	}

	var defaultOutboundRL, defaultInboundRL RateLimiterConfig
	if localDefaultExists {
		if outboundOnly {
			defaultOutboundRL, _ = GenerateTPRLConfigs(
				localDefaultBucket.RateLimit, RateLimiterConfigFloatInput{}, localDecimals, remoteDecimals,
				chainFamily, tokenPoolRef.Version, tokenPoolRef.Type.String(),
			)
		} else {
			defaultOutboundRL, defaultInboundRL = GenerateTPRLConfigs(
				localDefaultBucket.RateLimit, remoteDefaultBucket.RateLimit, localDecimals, remoteDecimals,
				chainFamily, tokenPoolRef.Version, tokenPoolRef.Type.String(),
			)
		}
		buckets = append(buckets, TPRLRateLimitBucket{
			FastFinality:              false,
			OutboundRateLimiterConfig: defaultOutboundRL,
			InboundRateLimiterConfig:  defaultInboundRL,
		})
	}

	return defaultOutboundRL, defaultInboundRL, buckets, nil
}

// validateOutboundOnlyAgainstCounterpartInbound reads the counterpart chain's on-chain
// inbound rate limit (one per FastFinality bucket present in localOutbounds) and checks
// that each is at least 110% of the new outbound being set on the local chain (both
// Capacity and Rate). Per the design, IsEnabled is ignored when checking — adapters are
// expected to return a zero RateLimiterConfig for unconfigured lanes, which will cause
// the check to reject any positive outbound.
func validateOutboundOnlyAgainstCounterpartInbound(
	e cldf.Environment,
	counterpartAdapter TokenAdapter,
	counterpartFamily string,
	counterpartPoolRef datastore.AddressRef,
	counterpartTokenRef datastore.AddressRef,
	counterpartSelector uint64,
	counterpartDecimals uint8,
	localSelector uint64,
	localDecimals uint8,
	localOutbounds RemoteOutbounds,
) error {
	reader, ok := counterpartAdapter.(RateLimitReaderAdapter)
	if !ok {
		return fmt.Errorf(
			"adapter for counterpart chain selector %d (family %s, version %v) does not implement RateLimitReaderAdapter; outbound-only mode unsupported for this lane",
			counterpartSelector, counterpartFamily, counterpartPoolRef.Version,
		)
	}

	check := func(bucket RateLimitConfig) error {
		// Compute what the counterpart's inbound would be if it were derived from the
		// new outbound the normal way: GenerateTPRLConfigs scales inboundInput by the
		// counterpart pool's decimals/scaling rules with +10% applied. We pass the
		// counterpart's family / pool version / pool type because the inbound is being
		// stored on the counterpart pool, and the legacy EVM pools (<1.6.1) scale by
		// the source (here: local) decimals while newer pools scale by their own.
		_, requiredInbound := GenerateTPRLConfigs(
			RateLimiterConfigFloatInput{}, bucket.RateLimit,
			counterpartDecimals, localDecimals,
			counterpartFamily, counterpartPoolRef.Version, counterpartPoolRef.Type.String(),
		)
		// If the new outbound is disabled there is nothing to validate against.
		if !bucket.RateLimit.IsEnabled {
			return nil
		}
		onchainInbound, err := reader.GetOnchainInboundRateLimit(
			e, counterpartSelector, counterpartPoolRef, counterpartTokenRef, localSelector, bucket.FastFinality,
		)
		if err != nil {
			return fmt.Errorf("failed to read on-chain inbound rate limit on counterpart chain selector %d (fastFinality=%v): %w", counterpartSelector, bucket.FastFinality, err)
		}
		if !onchainInbound.IsEnabled {
			return nil // If the counterpart's on-chain inbound is disabled, we allow any outbound to be set since it won't be enforced
		}
		if onchainInbound.Capacity == nil || onchainInbound.Capacity.Cmp(requiredInbound.Capacity) < 0 {
			return fmt.Errorf(
				"on-chain inbound capacity (%v) on counterpart chain selector %d for lane from %d is below required 110%% of new outbound capacity (%v) for fastFinality=%v",
				onchainInbound.Capacity, counterpartSelector, localSelector, requiredInbound.Capacity, bucket.FastFinality,
			)
		}
		if onchainInbound.Rate == nil || onchainInbound.Rate.Cmp(requiredInbound.Rate) < 0 {
			return fmt.Errorf(
				"on-chain inbound rate (%v) on counterpart chain selector %d for lane from %d is below required 110%% of new outbound rate (%v) for fastFinality=%v",
				onchainInbound.Rate, counterpartSelector, localSelector, requiredInbound.Rate, bucket.FastFinality,
			)
		}
		return nil
	}

	if defaultBucket, ok := localOutbounds.DefaultBucket(); ok {
		if err := check(defaultBucket); err != nil {
			return err
		}
	}
	if fastFinalityBucket, ok := localOutbounds.FastFinalityBucket(); ok {
		if err := check(fastFinalityBucket); err != nil {
			return err
		}
	}
	return nil
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
	localDecimals uint8,
	remoteDecimals uint8,
	chainFamily string,
	poolVersion *semver.Version,
	poolType string,
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
		outboundConfig.Capacity = ScaleFloatToBigInt(outboundInput.Capacity, int(localDecimals), 0)
		outboundConfig.Rate = ScaleFloatToBigInt(outboundInput.Rate, int(localDecimals), 0)
	}

	if !inboundInput.IsEnabled {
		inboundConfig.IsEnabled = false
		inboundConfig.Capacity = big.NewInt(0)
		inboundConfig.Rate = big.NewInt(0)
	} else {
		// We set the inbound capacity to be 1.1x the outbound capacity of the counterpart to avoid accidentally hitting the rate limit due to minor timing differences in refilling
		scaleByDecimals := localDecimals

		// https://github.com/smartcontractkit/chainlink-deployments/blob/cce886554ca0587492955784381321ce817fb6bb/domains/ccip/shared/tokendefaults.go#L1904
		// Only old EVM pools need to scale by remote decimals on inbound. Newer pools and non-EVM pools handle all conversions in local decimals.
		// This is a hack. Avoiding it would require refactoring the token pool adapters to handle rate limit configs in a more structured way instead of
		// just passing them as bytes through the registry, so for now we can live with this special case for old EVM pools since we're moving towards newer versions and non-EVM chains where this isn't an issue.
		if chainFamily == chain_selectors.FamilyEVM && poolVersion.LessThan(utils.Version_1_6_1) {
			// These custom contracts actually scale by local decimals:
			//   BurnMintWithExternalMinterTokenPool: https://explorer.plume.org/address/0x770318D51052871DeF5Eb5c452F4fd28B7960C4e?tab=contract
			//   HybridWithExternalMinterTokenPool: https://etherscan.io/address/0x36a72eD0096B414521C45E3ddC9ed657d1D9c141#code
			isBurnMintWithExternalMinterTokenPool := poolType == utils.BurnMintWithExternalMinterTokenPool.String()
			isHybridWithExternalMinterTokenPool := poolType == utils.HybridWithExternalMinterTokenPool.String()
			if poolVersion.Equal(utils.Version_1_6_0) && (isBurnMintWithExternalMinterTokenPool || isHybridWithExternalMinterTokenPool) {
				scaleByDecimals = localDecimals
			} else {
				scaleByDecimals = remoteDecimals
			}
		}
		inboundConfig.IsEnabled = true
		inboundConfig.Capacity = ScaleFloatToBigInt(inboundInput.Capacity, int(scaleByDecimals), .10)
		inboundConfig.Rate = ScaleFloatToBigInt(inboundInput.Rate, int(scaleByDecimals), .10)
	}
	return outboundConfig, inboundConfig
}
