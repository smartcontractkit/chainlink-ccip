package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"

	ccvdeploymentadapters "github.com/smartcontractkit/chainlink-ccv/deployment/adapters"
)

// ProtocolContractsFeeAggregatorExtra is the FamilyExtras key (string hex
// address) used to set the fee aggregator on the OnRamp and on deployed
// executor proxies. It is optional: when absent the fee aggregator defaults to
// the zero address and is wired up in a later configuration step.
const ProtocolContractsFeeAggregatorExtra = "feeAggregator"

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

	params := defaultDeployContractParams()
	// Protocol-only deploy: committee verifiers and the mock receivers (which
	// reference the committee verifier resolver) are deployed separately.
	params.CommitteeVerifiers = nil
	params.MockReceivers = nil
	params.OnRamp.FeeAggregator = feeAggregator
	params.Executors = protocolExecutorParams(input.Executors, feeAggregator)

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
// EVM executor params, applying EVM defaults (dynamic config, MaxCCVsPerMsg) and
// the supplied fee aggregator. Version and Qualifier fall back to the defaults
// when the caller leaves them unset.
func protocolExecutorParams(executors []ccvdeploymentadapters.ExecutorDeployParams, feeAggregator string) []ccvadapters.ExecutorDeployParams {
	if len(executors) == 0 {
		return nil
	}
	defaults := defaultExecutorParams(feeAggregator)[0]
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
