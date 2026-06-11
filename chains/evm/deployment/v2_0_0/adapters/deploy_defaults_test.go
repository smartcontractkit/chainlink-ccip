package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

func TestEVMDeployChainContractsAdapter_BuildDeployContractParams(t *testing.T) {
	adapter := &adapters.EVMDeployChainContractsAdapter{}
	defaults := adapter.GetDefaultDeployContractParams(1)
	require.NotNil(t, defaults.FeeQuoter.MaxFeeJuelsPerMsg)
	assert.Equal(t, offramp.Version, defaults.OffRamp.Version)

	built, err := adapter.BuildDeployContractParams(ccvadapters.BuildDeployContractParamsInput{
		ChainSelector: 1,
		CommitteeVerifiers: []ccvadapters.CommitteeVerifierDeployParams{
			{
				Version:       committee_verifier.Version,
				FeeAggregator: "0x0000000000000000000000000000000000000001",
				Qualifier:     "default",
			},
		},
		Defaults: defaults,
	})
	require.NoError(t, err)
	assert.Equal(t, "0x0000000000000000000000000000000000000001", built.OnRamp.FeeAggregator)
	require.Len(t, built.Executors, 1)
	assert.Equal(t, "default", built.Executors[0].Qualifier)
	assert.Equal(t, executor.Version, built.Executors[0].Version)
	require.Len(t, built.MockReceivers, 1)
}

func TestEVMDeployChainContractsAdapter_BuildDeployContractParams_AppliesPointerOverrides(t *testing.T) {
	adapter := &adapters.EVMDeployChainContractsAdapter{}
	maxUSDCents := uint32(50_00)

	built, err := adapter.BuildDeployContractParams(ccvadapters.BuildDeployContractParamsInput{
		ChainSelector: 1,
		CommitteeVerifiers: []ccvadapters.CommitteeVerifierDeployParams{
			{
				Version:       committee_verifier.Version,
				FeeAggregator: "0x0000000000000000000000000000000000000001",
				Qualifier:     "default",
			},
		},
		Overrides: &ccvadapters.DeployContractParamsOverrides{
			OnRamp: &ccvadapters.OnRampDeployParamsOverrides{
				MaxUSDCentsPerMessage: &maxUSDCents,
			},
		},
	})
	require.NoError(t, err)
	assert.Equal(t, uint32(50_00), built.OnRamp.MaxUSDCentsPerMessage)
	assert.Equal(t, uint32(100_00), adapter.GetDefaultDeployContractParams(1).OnRamp.MaxUSDCentsPerMessage)
}

func TestApplyDeployContractParamsOverrides_UsesDefaultsWhenOverrideNil(t *testing.T) {
	defaults := ccvadapters.DeployContractParams{
		OffRamp: ccvadapters.OffRampDeployParams{GasForCallExactCheck: 5_000},
	}
	merged := ccvadapters.ApplyDeployContractParamsOverrides(defaults, nil)
	assert.Equal(t, uint16(5_000), merged.OffRamp.GasForCallExactCheck)
}

func TestApplyDeployContractParamsOverrides_OverridesOnlySetFields(t *testing.T) {
	defaults := ccvadapters.DeployContractParams{
		OffRamp: ccvadapters.OffRampDeployParams{
			GasForCallExactCheck:      5_000,
			MaxGasBufferToUpdateState: 12_000,
		},
	}
	gas := uint16(9_000)
	merged := ccvadapters.ApplyDeployContractParamsOverrides(defaults, &ccvadapters.DeployContractParamsOverrides{
		OffRamp: &ccvadapters.OffRampDeployParamsOverrides{
			GasForCallExactCheck: &gas,
		},
	})
	assert.Equal(t, uint16(9_000), merged.OffRamp.GasForCallExactCheck)
	assert.Equal(t, uint32(12_000), merged.OffRamp.MaxGasBufferToUpdateState)
}
