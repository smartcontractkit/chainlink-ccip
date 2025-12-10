package changesets

import (
	"fmt"
	"time"

	"dario.cat/mergo"
	"github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type CCIPOCRParams struct {
	// OCRParameters contains the parameters for the OCR plugin.
	OCRParameters OCRParameters `json:"ocrParameters"`
	// CommitOffChainConfig contains pointers to Arb feeds for prices.
	CommitOffChainConfig *pluginconfig.CommitOffchainConfig `json:"commitOffChainConfig,omitempty"`
	// ExecuteOffChainConfig contains USDC config.
	ExecuteOffChainConfig *pluginconfig.ExecuteOffchainConfig `json:"executeOffChainConfig,omitempty"`
}

func (c CCIPOCRParams) Copy() CCIPOCRParams {
	newC := CCIPOCRParams{
		OCRParameters: c.OCRParameters,
	}
	if c.CommitOffChainConfig != nil {
		commit := *c.CommitOffChainConfig
		newC.CommitOffChainConfig = &commit
	}
	if c.ExecuteOffChainConfig != nil {
		exec := *c.ExecuteOffChainConfig
		newC.ExecuteOffChainConfig = &exec
	}
	return newC
}

var (
	DefaultOCRParamsForCommitForNonETH = CCIPOCRParams{
		OCRParameters:        CommitOCRParams,
		CommitOffChainConfig: &DefaultCommitOffChainCfg,
	}

	DefaultOCRParamsForCommitForETH = CCIPOCRParams{
		OCRParameters:        CommitOCRParamsForEthereum,
		CommitOffChainConfig: &DefaultCommitOffChainCfgForEth,
	}

	DefaultOCRParamsForExecForNonETH = CCIPOCRParams{
		OCRParameters:         ExecOCRParams,
		ExecuteOffChainConfig: &DefaultExecuteOffChainCfg,
	}

	DefaultOCRParamsForExecForETH = CCIPOCRParams{
		OCRParameters:         ExecOCRParamsForEthereum,
		ExecuteOffChainConfig: &DefaultExecuteOffChainCfg,
	}

	// Used for only testing with simulated chains
	OcrParamsForTest = CCIPOCRParams{
		OCRParameters: OCRParameters{
			DeltaProgress:                           30 * time.Second, // Lower DeltaProgress can lead to timeouts when running tests locally
			DeltaResend:                             10 * time.Second,
			DeltaInitial:                            20 * time.Second,
			DeltaRound:                              2 * time.Second,
			DeltaGrace:                              2 * time.Second,
			DeltaCertifiedCommitRequest:             10 * time.Second,
			DeltaStage:                              10 * time.Second,
			Rmax:                                    50,
			MaxDurationQuery:                        10 * time.Second,
			MaxDurationObservation:                  10 * time.Second,
			MaxDurationShouldAcceptAttestedReport:   10 * time.Second,
			MaxDurationShouldTransmitAcceptedReport: 10 * time.Second,
		},
		CommitOffChainConfig: &pluginconfig.CommitOffchainConfig{
			RemoteGasPriceBatchWriteFrequency:  *config.MustNewDuration(RemoteGasPriceBatchWriteFrequency),
			TokenPriceBatchWriteFrequency:      *config.MustNewDuration(TokenPriceBatchWriteFrequency),
			NewMsgScanBatchSize:                merklemulti.MaxNumberTreeLeaves,
			MaxReportTransmissionCheckAttempts: 5,
			RMNEnabled:                         false,
			RMNSignaturesTimeout:               30 * time.Minute,
			MaxMerkleTreeSize:                  merklemulti.MaxNumberTreeLeaves,
			SignObservationPrefix:              "chainlink ccip 1.6 rmn observation",
			MerkleRootAsyncObserverDisabled:    false,
			MerkleRootAsyncObserverSyncFreq:    4 * time.Second,
			MerkleRootAsyncObserverSyncTimeout: 12 * time.Second,
			ChainFeeAsyncObserverSyncFreq:      10 * time.Second,
			ChainFeeAsyncObserverSyncTimeout:   12 * time.Second,
			DonBreakingChangesVersion:          pluginconfig.DonBreakingChangesVersion1RoleDonSupport,
		},
		ExecuteOffChainConfig: &pluginconfig.ExecuteOffchainConfig{
			BatchGasLimit:             BatchGasLimit,
			InflightCacheExpiry:       *config.MustNewDuration(InflightCacheExpiry),
			RootSnoozeTime:            *config.MustNewDuration(RootSnoozeTime),
			MessageVisibilityInterval: *config.MustNewDuration(PermissionLessExecutionThreshold),
			BatchingStrategyID:        BatchingStrategyID,
			MaxCommitReportsToFetch:   MaxCommitReportsToFetch,
		},
	}
)

type OCRConfigChainType int

const (
	Default OCRConfigChainType = iota + 1
	Ethereum
	// SimulationTest is kept only for backward compatibility. Tests probably should
	// migrate to using Default or Ethereum
	SimulationTest
)

func (c OCRConfigChainType) CommitOCRParams() CCIPOCRParams {
	switch c {
	case Ethereum:
		return DefaultOCRParamsForCommitForETH.Copy()
	case Default:
		return DefaultOCRParamsForCommitForNonETH.Copy()
	case SimulationTest:
		return OcrParamsForTest.Copy()
	default:
		panic("unknown OCRConfigChainType")
	}
}

func (c OCRConfigChainType) ExecuteOCRParams() CCIPOCRParams {
	switch c {
	case Ethereum:
		return DefaultOCRParamsForExecForETH.Copy()
	case Default:
		return DefaultOCRParamsForExecForNonETH.Copy()
	case SimulationTest:
		return OcrParamsForTest.Copy()
	default:
		panic("unknown OCRConfigChainType")
	}
}

func DeriveOCRParamsForCommit(
	ocrChainType OCRConfigChainType,
	feedChain uint64,
	feeTokenInfo map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo,
	override func(params CCIPOCRParams) CCIPOCRParams,
) CCIPOCRParams {
	params := ocrChainType.CommitOCRParams()
	params.CommitOffChainConfig.TokenInfo = feeTokenInfo
	params.CommitOffChainConfig.PriceFeedChainSelector = ccipocr3.ChainSelector(feedChain)
	if override == nil {
		return params
	}
	return override(params)
}

func DeriveOCRParamsForExec(
	ocrChainType OCRConfigChainType,
	observerConfig []pluginconfig.TokenDataObserverConfig,
	override func(params CCIPOCRParams) CCIPOCRParams,
) CCIPOCRParams {
	params := ocrChainType.ExecuteOCRParams()
	params.ExecuteOffChainConfig.TokenDataObservers = observerConfig
	if override == nil {
		return params
	}
	return override(params)
}

type ConfigType string

const (
	ConfigTypeActive    ConfigType = "active"
	ConfigTypeCandidate ConfigType = "candidate"
	// ========= Changeset Defaults =========
	PermissionLessExecutionThreshold  = 1 * time.Hour
	RemoteGasPriceBatchWriteFrequency = 20 * time.Minute
	TokenPriceBatchWriteFrequency     = 2 * time.Hour
	// Building batches with 6.5m and transmit with 8m to account for overhead.
	BatchGasLimit               = 6_500_000
	InflightCacheExpiry         = 1 * time.Minute
	RootSnoozeTime              = 5 * time.Minute
	BatchingStrategyID          = 0
	OptimisticConfirmations     = 1
	TransmissionDelayMultiplier = 15 * time.Second
	MaxCommitReportsToFetch     = 250
	// ======================================

	// ========= Onchain consts =========
	// CCIPLockOrBurnV1RetBytes Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES
	// Reference: https://github.com/smartcontractkit/chainlink/blob/develop/contracts/src/v0.8/ccip/libraries/Pool.sol#L17
	CCIPLockOrBurnV1RetBytes = 32
	// ======================================
)

var (
	// DefaultCommitOffChainCfg represents the default offchain configuration for the Commit plugin
	// on _most_ chains. This should be used as a base for all chains, with overrides only where necessary.
	// Notable overrides are for Ethereum, which has a slower block time.
	DefaultCommitOffChainCfg = pluginconfig.CommitOffchainConfig{
		RemoteGasPriceBatchWriteFrequency:  *config.MustNewDuration(RemoteGasPriceBatchWriteFrequency),
		TokenPriceBatchWriteFrequency:      *config.MustNewDuration(TokenPriceBatchWriteFrequency),
		NewMsgScanBatchSize:                merklemulti.MaxNumberTreeLeaves,
		MaxReportTransmissionCheckAttempts: 10,
		RMNSignaturesTimeout:               6900 * time.Millisecond,
		RMNEnabled:                         true,
		MaxMerkleTreeSize:                  merklemulti.MaxNumberTreeLeaves,
		SignObservationPrefix:              "chainlink ccip 1.6 rmn observation",
		// TransmissionDelayMultiplier for non-ETH (i.e, typically fast) chains should be pretty aggressive.
		// e.g assuming a 2s blocktime, 15 seconds is ~8 blocks.
		TransmissionDelayMultiplier:        TransmissionDelayMultiplier,
		InflightPriceCheckRetries:          10,
		MerkleRootAsyncObserverDisabled:    false,
		MerkleRootAsyncObserverSyncFreq:    4 * time.Second,
		MerkleRootAsyncObserverSyncTimeout: 12 * time.Second,

		// Disabling the chainfee + tokenprice async observers because the low cache TTL + low timeout
		// is currently not a viable combo.
		// Super aggressive frequency and timeout causes rpc timeouts more frequently.
		ChainFeeAsyncObserverDisabled: true,
		// TODO: revisit
		// ChainFeeAsyncObserverSyncFreq:      1*time.Second + 500*time.Millisecond,
		// ChainFeeAsyncObserverSyncTimeout:   1 * time.Second,
		TokenPriceAsyncObserverDisabled: true,
		// TODO: revisit
		// TokenPriceAsyncObserverSyncFreq:    *config.MustNewDuration(1*time.Second + 500*time.Millisecond),
		// TokenPriceAsyncObserverSyncTimeout: *config.MustNewDuration(1 * time.Second),

		// Remaining fields cannot be statically set:
		// PriceFeedChainSelector: , // Must be configured in CLD
		// TokenInfo: , // Must be configured in CLD

		DonBreakingChangesVersion: pluginconfig.DonBreakingChangesVersion1RoleDonSupport,
	}

	// DefaultExecuteOffChainCfg represents the default offchain configuration for the Execute plugin
	// on _most_ chains. This should be used as a base for all chains, with overrides only where necessary.
	// Notable overrides are for Ethereum, which has a slower block time.
	DefaultExecuteOffChainCfg = pluginconfig.ExecuteOffchainConfig{
		BatchGasLimit:               BatchGasLimit,
		InflightCacheExpiry:         *config.MustNewDuration(InflightCacheExpiry),
		RootSnoozeTime:              *config.MustNewDuration(RootSnoozeTime),
		MessageVisibilityInterval:   *config.MustNewDuration(8 * time.Hour),
		BatchingStrategyID:          BatchingStrategyID,
		TransmissionDelayMultiplier: TransmissionDelayMultiplier,
		MaxReportMessages:           0,
		MaxSingleChainReports:       0,
		MaxCommitReportsToFetch:     MaxCommitReportsToFetch,
		// Remaining fields cannot be statically set:
		// TokenDataObservers: , // Must be configured in CLD
	}

	DefaultCommitOffChainCfgForEth = withCommitOffchainConfigOverrides(
		DefaultCommitOffChainCfg,
		pluginconfig.CommitOffchainConfig{
			RemoteGasPriceBatchWriteFrequency: *config.MustNewDuration(2 * time.Hour),
			TokenPriceBatchWriteFrequency:     *config.MustNewDuration(12 * time.Hour),
		},
	)
)

func withCommitOffchainConfigOverrides(base pluginconfig.CommitOffchainConfig, overrides pluginconfig.CommitOffchainConfig) pluginconfig.CommitOffchainConfig {
	outcome := base
	if err := mergo.Merge(&outcome, overrides, mergo.WithOverride); err != nil {
		panic(fmt.Sprintf("error while building an OCR config %v", err))
	}
	return outcome
}
