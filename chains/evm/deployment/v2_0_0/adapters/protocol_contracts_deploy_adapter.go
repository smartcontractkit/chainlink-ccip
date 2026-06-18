package adapters

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Masterminds/semver/v3"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

const (
	// ProtocolContractsFeeAggregatorExtra is the FamilyExtras key (string hex
	// address) used to set the fee aggregator on the OnRamp and on deployed
	// executor proxies. It is optional: when absent the fee aggregator defaults
	// to the zero address and is wired up in a later configuration step.
	ProtocolContractsFeeAggregatorExtra = "feeAggregator"

	// ProtocolContractsExecutorBlockDepthExtra and
	// ProtocolContractsExecutorWaitForSafeExtra are the optional FamilyExtras
	// keys used to set the executors' allowed finality config at deploy time.
	// Executors validate the requested finality in getFee, and (unlike committee
	// verifiers and token pools) their allowed finality is never reconfigured by
	// the lane-configuration changesets — it can only be set on the deployed
	// DynamicConfig. When both keys are absent the deploy default applies; when
	// either is present the executors are deployed with the resulting
	// finality.Config. ExecutorBlockDepth is an integer; ExecutorWaitForSafe is
	// a bool.
	ProtocolContractsExecutorBlockDepthExtra  = "executorBlockDepth"
	ProtocolContractsExecutorWaitForSafeExtra = "executorWaitForSafe"

	// ProtocolContractsExecutorMaxCCVsPerMsgExtra (integer, 0-255) and
	// ProtocolContractsExecutorCcvAllowlistEnabledExtra (bool) are the optional
	// FamilyExtras keys used to override the executors' MaxCCVsPerMsg and
	// CcvAllowlistEnabled at deploy time. Each falls back to the deploy default
	// (10 and false respectively) when its key is absent.
	ProtocolContractsExecutorMaxCCVsPerMsgExtra       = "executorMaxCCVsPerMsg"
	ProtocolContractsExecutorCcvAllowlistEnabledExtra = "executorCcvAllowlistEnabled"

	// The remaining keys override individual contract deploy params. All are
	// optional; an absent key leaves the corresponding deploy default in place.
	// The FeeQuoter price fields are big.Int and must be passed as base-10
	// strings because TOML integers are limited to int64.
	ProtocolContractsRMNRemoteLegacyRMNExtra = "rmnRemoteLegacyRmn"

	ProtocolContractsOffRampGasForCallExactCheckExtra      = "offRampGasForCallExactCheck"
	ProtocolContractsOffRampMaxGasBufferToUpdateStateExtra = "offRampMaxGasBufferToUpdateState"

	ProtocolContractsOnRampMaxUSDCentsPerMessageExtra = "onRampMaxUsdCentsPerMessage"

	ProtocolContractsFeeQuoterMaxFeeJuelsPerMsgExtra              = "feeQuoterMaxFeeJuelsPerMsg"
	ProtocolContractsFeeQuoterLINKPremiumMultiplierWeiPerEthExtra = "feeQuoterLinkPremiumMultiplierWeiPerEth"
	ProtocolContractsFeeQuoterWETHPremiumMultiplierWeiPerEthExtra = "feeQuoterWethPremiumMultiplierWeiPerEth"
	ProtocolContractsFeeQuoterUSDPerLINKExtra                     = "feeQuoterUsdPerLink"
	ProtocolContractsFeeQuoterUSDPerWETHExtra                     = "feeQuoterUsdPerWeth"
)

// EVMProtocolContractsDeployAdapter implements
// ccvdeploymentadapters.ProtocolContractsDeployAdapter for EVM chains.
//
// It deploys the core CCIP protocol contracts (RMNRemote, RMNProxy, Router,
// FeeQuoter, OnRamp, OffRamp, TokenAdminRegistry, Executors, …) by reusing the
// existing EVM DeployChainContracts sequence with committee verifiers and mock
// receivers omitted. Those are deployed separately — committee verifiers via
// EVMCommitteeVerifierDeployAdapter, which the topology-free OnboardChain
// changeset runs as a second phase against the protocol addresses deployed
// here. The underlying sequence is idempotent, so re-running reconciles drift
// rather than redeploying.
type EVMProtocolContractsDeployAdapter struct{}

var _ ccvdeploymentadapters.ProtocolContractsDeployAdapter = (*EVMProtocolContractsDeployAdapter)(nil)

var evmDeployProtocolContracts = cldf_ops.NewSequence(
	"evm-deploy-protocol-contracts",
	semver.MustParse("2.0.0"),
	"Chain-agnostic wrapper around the EVM DeployChainContracts sequence that deploys protocol contracts only (no committee verifiers)",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input ccvdeploymentadapters.ProtocolContractsDeployInput) (ccvdeploymentadapters.ProtocolContractsDeployOutput, error) {
		evmChains := chains.EVMChains()
		evmChain, ok := evmChains[input.ChainSelector]
		if !ok {
			return ccvdeploymentadapters.ProtocolContractsDeployOutput{},
				fmt.Errorf("EVM chain not found for selector %d", input.ChainSelector)
		}

		evmInput, err := toEVMProtocolDeployInput(input)
		if err != nil {
			return ccvdeploymentadapters.ProtocolContractsDeployOutput{},
				fmt.Errorf("failed to convert protocol contracts deploy input to EVM types: %w", err)
		}

		report, err := cldf_ops.ExecuteSequence(b, sequences.DeployChainContracts, evmChain, evmInput)
		if err != nil {
			return ccvdeploymentadapters.ProtocolContractsDeployOutput{},
				fmt.Errorf("EVM DeployChainContracts (protocol-only) failed: %w", err)
		}

		return ccvdeploymentadapters.ProtocolContractsDeployOutput{
			Addresses:               report.Output.Addresses,
			BatchOps:                report.Output.BatchOps,
			RefsToTransferOwnership: report.Output.RefsToTransferOwnership,
		}, nil
	},
)

func (a *EVMProtocolContractsDeployAdapter) DeployProtocolContracts() *cldf_ops.Sequence[ccvdeploymentadapters.ProtocolContractsDeployInput, ccvdeploymentadapters.ProtocolContractsDeployOutput, cldf_chain.BlockChains] {
	return evmDeployProtocolContracts
}

// toEVMProtocolDeployInput builds the EVM DeployChainContracts input for a
// protocol-only deploy. It starts from the EVM defaults, clears the committee
// verifiers and mock receivers (deployed separately), wires the executors from
// the caller, and threads the optional fee aggregator from FamilyExtras through
// to the OnRamp and executors. It reuses toEVMDeployInput so the address
// parsing and EVM conversion logic stays identical to the bundled deploy path.
func toEVMProtocolDeployInput(input ccvdeploymentadapters.ProtocolContractsDeployInput) (sequences.DeployChainContractsInput, error) {
	feeAggregator, err := protocolContractsFeeAggregator(input.FamilyExtras)
	if err != nil {
		return sequences.DeployChainContractsInput{}, err
	}
	executorOverrides, err := protocolContractsExecutorOverrides(input.FamilyExtras)
	if err != nil {
		return sequences.DeployChainContractsInput{}, err
	}

	params := defaultDeployContractParams()
	// Protocol-only deploy: committee verifiers and the mock receivers (which
	// reference the committee verifier resolver) are deployed separately.
	params.CommitteeVerifiers = nil
	params.MockReceivers = nil
	params.OnRamp.FeeAggregator = feeAggregator
	params.Executors = protocolExecutorParams(input.Executors, feeAggregator, executorOverrides)
	if err := applyProtocolContractParamOverrides(&params, input.FamilyExtras); err != nil {
		return sequences.DeployChainContractsInput{}, err
	}

	return toEVMDeployInput(ccvadapters.DeployChainContractsInput{
		ChainSelector:     input.ChainSelector,
		DeployerContract:  input.DeployerContract,
		DeployTestRouter:  input.DeployTestRouter,
		ExistingAddresses: input.ExistingAddresses,
		ContractParams:    params,
		DeployerKeyOwned:  input.DeployerKeyOwned,
	})
}

// protocolExecutorParams maps the chain-agnostic executor deploy params onto the
// EVM executor params, applying EVM defaults (dynamic config, MaxCCVsPerMsg), the
// supplied fee aggregator, and any executor overrides from FamilyExtras. Version
// and Qualifier fall back to the defaults when the caller leaves them unset; each
// override field, when set, replaces the corresponding default on every deployed
// executor.
func protocolExecutorParams(executors []ccvdeploymentadapters.ExecutorDeployParams, feeAggregator string, overrides executorDeployOverrides) []ccvadapters.ExecutorDeployParams {
	if len(executors) == 0 {
		return nil
	}
	defaults := defaultExecutorParams(feeAggregator)[0]
	if overrides.allowedFinality != nil {
		defaults.DynamicConfig.AllowedFinalityConfig = *overrides.allowedFinality
	}
	if overrides.maxCCVsPerMsg != nil {
		defaults.MaxCCVsPerMsg = *overrides.maxCCVsPerMsg
	}
	if overrides.ccvAllowlistEnabled != nil {
		defaults.DynamicConfig.CcvAllowlistEnabled = *overrides.ccvAllowlistEnabled
	}
	result := make([]ccvadapters.ExecutorDeployParams, 0, len(executors))
	for _, e := range executors {
		params := defaults
		if e.Version != nil {
			params.Version = e.Version
		}
		if e.Qualifier != "" {
			params.Qualifier = e.Qualifier
		}
		result = append(result, params)
	}
	return result
}

// protocolContractsFeeAggregator extracts the optional fee aggregator address
// from FamilyExtras, returning the empty string when it is not set.
func protocolContractsFeeAggregator(extras map[string]any) (string, error) {
	raw, ok := extras[ProtocolContractsFeeAggregatorExtra]
	if !ok {
		return "", nil
	}
	feeAgg, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("FamilyExtras[%q] must be a string hex address, got %T", ProtocolContractsFeeAggregatorExtra, raw)
	}
	return feeAgg, nil
}

// executorDeployOverrides carries the optional executor deploy-config overrides
// parsed from FamilyExtras. A nil field means "use the deploy default".
type executorDeployOverrides struct {
	allowedFinality     *finality.Config
	maxCCVsPerMsg       *uint8
	ccvAllowlistEnabled *bool
}

// protocolContractsExecutorOverrides reads the optional executor deploy-config
// overrides from FamilyExtras. Absent keys leave the corresponding field nil so
// the deploy default applies.
func protocolContractsExecutorOverrides(extras map[string]any) (executorDeployOverrides, error) {
	var overrides executorDeployOverrides

	allowedFinality, err := protocolContractsExecutorFinality(extras)
	if err != nil {
		return overrides, err
	}
	overrides.allowedFinality = allowedFinality

	if v, ok, err := extraBoundedInt(extras, ProtocolContractsExecutorMaxCCVsPerMsgExtra, math.MaxUint8); err != nil {
		return overrides, err
	} else if ok {
		maxCCVs := uint8(v)
		overrides.maxCCVsPerMsg = &maxCCVs
	}

	if enabled, ok, err := extraBool(extras, ProtocolContractsExecutorCcvAllowlistEnabledExtra); err != nil {
		return overrides, err
	} else if ok {
		overrides.ccvAllowlistEnabled = &enabled
	}

	return overrides, nil
}

// protocolContractsExecutorFinality reads the optional executor allowed finality
// config from FamilyExtras. It returns nil when neither key is set (deploy
// default applies). When either key is present it builds the finality.Config
// from ExecutorBlockDepth (integer, 0-65535) and ExecutorWaitForSafe (bool).
func protocolContractsExecutorFinality(extras map[string]any) (*finality.Config, error) {
	depth, hasDepth, err := extraBoundedInt(extras, ProtocolContractsExecutorBlockDepthExtra, math.MaxUint16)
	if err != nil {
		return nil, err
	}
	safe, hasSafe, err := extraBool(extras, ProtocolContractsExecutorWaitForSafeExtra)
	if err != nil {
		return nil, err
	}
	if !hasDepth && !hasSafe {
		return nil, nil
	}
	return &finality.Config{
		BlockDepth:  uint16(depth),
		WaitForSafe: safe,
	}, nil
}

// applyProtocolContractParamOverrides applies the optional contract-parameter
// FamilyExtras overrides onto the default DeployContractParams. Each override is
// applied only when its key is present; absent keys leave the deploy default.
func applyProtocolContractParamOverrides(params *ccvadapters.DeployContractParams, extras map[string]any) error {
	if v, ok, err := extraString(extras, ProtocolContractsRMNRemoteLegacyRMNExtra); err != nil {
		return err
	} else if ok {
		params.RMNRemote.LegacyRMN = v
	}

	if v, ok, err := extraBoundedInt(extras, ProtocolContractsOffRampGasForCallExactCheckExtra, math.MaxUint16); err != nil {
		return err
	} else if ok {
		params.OffRamp.GasForCallExactCheck = uint16(v)
	}

	if v, ok, err := extraBoundedInt(extras, ProtocolContractsOffRampMaxGasBufferToUpdateStateExtra, math.MaxUint32); err != nil {
		return err
	} else if ok {
		params.OffRamp.MaxGasBufferToUpdateState = uint32(v)
	}

	if v, ok, err := extraBoundedInt(extras, ProtocolContractsOnRampMaxUSDCentsPerMessageExtra, math.MaxUint32); err != nil {
		return err
	} else if ok {
		params.OnRamp.MaxUSDCentsPerMessage = uint32(v)
	}

	if v, ok, err := extraBoundedInt(extras, ProtocolContractsFeeQuoterLINKPremiumMultiplierWeiPerEthExtra, math.MaxInt64); err != nil {
		return err
	} else if ok {
		params.FeeQuoter.LINKPremiumMultiplierWeiPerEth = uint64(v)
	}

	if v, ok, err := extraBoundedInt(extras, ProtocolContractsFeeQuoterWETHPremiumMultiplierWeiPerEthExtra, math.MaxInt64); err != nil {
		return err
	} else if ok {
		params.FeeQuoter.WETHPremiumMultiplierWeiPerEth = uint64(v)
	}

	if v, ok, err := extraBigInt(extras, ProtocolContractsFeeQuoterMaxFeeJuelsPerMsgExtra); err != nil {
		return err
	} else if ok {
		params.FeeQuoter.MaxFeeJuelsPerMsg = v
	}

	if v, ok, err := extraBigInt(extras, ProtocolContractsFeeQuoterUSDPerLINKExtra); err != nil {
		return err
	} else if ok {
		params.FeeQuoter.USDPerLINK = v
	}

	if v, ok, err := extraBigInt(extras, ProtocolContractsFeeQuoterUSDPerWETHExtra); err != nil {
		return err
	} else if ok {
		params.FeeQuoter.USDPerWETH = v
	}

	return nil
}

// extraBoundedInt reads an optional integer FamilyExtras value, validating it is
// within [0, max]. The bool return is false when the key is absent.
func extraBoundedInt(extras map[string]any, key string, max int64) (int64, bool, error) {
	raw, present := extras[key]
	if !present {
		return 0, false, nil
	}
	n, ok := asInt64(raw)
	if !ok {
		return 0, false, fmt.Errorf("FamilyExtras[%q] must be an integer, got %T", key, raw)
	}
	if n < 0 || n > max {
		return 0, false, fmt.Errorf("FamilyExtras[%q] must be in [0, %d], got %d", key, max, n)
	}
	return n, true, nil
}

// extraBigInt reads an optional base-10 integer string FamilyExtras value into a
// big.Int. Wei-denominated values that exceed int64 must be passed as strings
// because TOML integers are limited to int64. The bool return is false when the
// key is absent.
func extraBigInt(extras map[string]any, key string) (*big.Int, bool, error) {
	raw, present := extras[key]
	if !present {
		return nil, false, nil
	}
	s, ok := raw.(string)
	if !ok {
		return nil, false, fmt.Errorf("FamilyExtras[%q] must be a base-10 integer string, got %T", key, raw)
	}
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil, false, fmt.Errorf("FamilyExtras[%q] must be a base-10 integer string, got %q", key, s)
	}
	return v, true, nil
}

// extraString reads an optional string FamilyExtras value. The bool return is
// false when the key is absent.
func extraString(extras map[string]any, key string) (string, bool, error) {
	raw, present := extras[key]
	if !present {
		return "", false, nil
	}
	s, ok := raw.(string)
	if !ok {
		return "", false, fmt.Errorf("FamilyExtras[%q] must be a string, got %T", key, raw)
	}
	return s, true, nil
}

// extraBool reads an optional bool FamilyExtras value. The bool ok return is
// false when the key is absent.
func extraBool(extras map[string]any, key string) (value bool, ok bool, err error) {
	raw, present := extras[key]
	if !present {
		return false, false, nil
	}
	b, isBool := raw.(bool)
	if !isBool {
		return false, false, fmt.Errorf("FamilyExtras[%q] must be a bool, got %T", key, raw)
	}
	return b, true, nil
}

// asInt64 coerces the common numeric types produced by TOML/JSON decoders.
func asInt64(v any) (int64, bool) {
	switch n := v.(type) {
	case int64:
		return n, true
	case int:
		return int64(n), true
	case float64:
		return int64(n), true
	default:
		return 0, false
	}
}
