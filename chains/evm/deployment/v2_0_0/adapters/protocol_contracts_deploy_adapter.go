package adapters

import (
	"fmt"

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

	if raw, ok := extras[ProtocolContractsExecutorMaxCCVsPerMsgExtra]; ok {
		n, ok := asInt64(raw)
		if !ok {
			return overrides, fmt.Errorf("FamilyExtras[%q] must be an integer, got %T", ProtocolContractsExecutorMaxCCVsPerMsgExtra, raw)
		}
		if n < 0 || n > 255 {
			return overrides, fmt.Errorf("FamilyExtras[%q] must be in [0, 255], got %d", ProtocolContractsExecutorMaxCCVsPerMsgExtra, n)
		}
		maxCCVs := uint8(n)
		overrides.maxCCVsPerMsg = &maxCCVs
	}

	if raw, ok := extras[ProtocolContractsExecutorCcvAllowlistEnabledExtra]; ok {
		enabled, ok := raw.(bool)
		if !ok {
			return overrides, fmt.Errorf("FamilyExtras[%q] must be a bool, got %T", ProtocolContractsExecutorCcvAllowlistEnabledExtra, raw)
		}
		overrides.ccvAllowlistEnabled = &enabled
	}

	return overrides, nil
}

// protocolContractsExecutorFinality reads the optional executor allowed finality
// config from FamilyExtras. It returns nil when neither key is set (deploy
// default applies). When either key is present it builds the finality.Config
// from ExecutorBlockDepth (integer, 0-65535) and ExecutorWaitForSafe (bool).
func protocolContractsExecutorFinality(extras map[string]any) (*finality.Config, error) {
	rawDepth, hasDepth := extras[ProtocolContractsExecutorBlockDepthExtra]
	rawSafe, hasSafe := extras[ProtocolContractsExecutorWaitForSafeExtra]
	if !hasDepth && !hasSafe {
		return nil, nil
	}
	cfg := finality.Config{}
	if hasDepth {
		depth, ok := asInt64(rawDepth)
		if !ok {
			return nil, fmt.Errorf("FamilyExtras[%q] must be an integer, got %T", ProtocolContractsExecutorBlockDepthExtra, rawDepth)
		}
		if depth < 0 || depth > 65535 {
			return nil, fmt.Errorf("FamilyExtras[%q] must be in [0, 65535], got %d", ProtocolContractsExecutorBlockDepthExtra, depth)
		}
		cfg.BlockDepth = uint16(depth)
	}
	if hasSafe {
		safe, ok := rawSafe.(bool)
		if !ok {
			return nil, fmt.Errorf("FamilyExtras[%q] must be a bool, got %T", ProtocolContractsExecutorWaitForSafeExtra, rawSafe)
		}
		cfg.WaitForSafe = safe
	}
	return &cfg, nil
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
