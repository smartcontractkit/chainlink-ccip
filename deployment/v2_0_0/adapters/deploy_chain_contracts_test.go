package adapters

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

func TestMergeIfNotEmpty(t *testing.T) {
	v170 := semver.MustParse("2.0.0")

	t.Run("empty source returns base unchanged", func(t *testing.T) {
		base := DeployContractParams{
			RMNRemote: RMNRemoteDeployParams{Version: v170, LegacyRMN: "0xBase"},
			OnRamp:    OnRampDeployParams{Version: v170, FeeAggregator: "0xAgg"},
		}
		source := DeployContractParams{}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		assert.Equal(t, base, merged)
	})

	t.Run("source overwrites base for set struct fields", func(t *testing.T) {
		base := DeployContractParams{
			RMNRemote: RMNRemoteDeployParams{Version: v170, LegacyRMN: "0xBaseRMN"},
			OnRamp:    OnRampDeployParams{Version: v170, FeeAggregator: "0xBaseAgg"},
		}
		source := DeployContractParams{
			RMNRemote: RMNRemoteDeployParams{Version: v170, LegacyRMN: "0xSourceRMN"},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		assert.Equal(t, "0xSourceRMN", merged.RMNRemote.LegacyRMN, "RMNRemote should come from source")
		assert.Equal(t, "0xBaseAgg", merged.OnRamp.FeeAggregator, "OnRamp should be unchanged from base")
	})

	t.Run("empty source fields do not overwrite base", func(t *testing.T) {
		base := DeployContractParams{
			OnRamp: OnRampDeployParams{
				Version:               v170,
				FeeAggregator:         "0xBaseAgg",
				MaxUSDCentsPerMessage: 100,
			},
			OffRamp: OffRampDeployParams{
				Version:                   v170,
				GasForCallExactCheck:      2000,
				MaxGasBufferToUpdateState: 50,
			},
		}
		// Source has one non-empty field (FeeAggregator) and rest empty/zero
		source := DeployContractParams{
			OnRamp: OnRampDeployParams{
				FeeAggregator:         "0xSourceAgg",
				MaxUSDCentsPerMessage: 0, // zero = empty, should not overwrite base
			},
			OffRamp: OffRampDeployParams{
				GasForCallExactCheck:      0, // zero = empty, should not overwrite base
				MaxGasBufferToUpdateState: 0,
			},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		// Non-empty from source overwrites
		assert.Equal(t, "0xSourceAgg", merged.OnRamp.FeeAggregator)
		// Empty/zero from source does not overwrite base
		assert.Equal(t, uint32(100), merged.OnRamp.MaxUSDCentsPerMessage, "base value preserved when source is zero")
		assert.Equal(t, uint16(2000), merged.OffRamp.GasForCallExactCheck, "base value preserved when source is zero")
		assert.Equal(t, uint32(50), merged.OffRamp.MaxGasBufferToUpdateState, "base value preserved when source is zero")
	})

	t.Run("source overwrites base for OffRamp and FeeQuoter", func(t *testing.T) {
		base := DeployContractParams{
			OffRamp:   OffRampDeployParams{Version: v170, GasForCallExactCheck: 100},
			FeeQuoter: FeeQuoterDeployParams{Version: v170, MaxFeeJuelsPerMsg: big.NewInt(1e18)},
		}
		source := DeployContractParams{
			OffRamp:   OffRampDeployParams{Version: v170, GasForCallExactCheck: 200, MaxGasBufferToUpdateState: 50},
			FeeQuoter: FeeQuoterDeployParams{Version: v170, MaxFeeJuelsPerMsg: big.NewInt(2e18), USDPerLINK: big.NewInt(5e6)},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		assert.Equal(t, source.OffRamp.GasForCallExactCheck, merged.OffRamp.GasForCallExactCheck)
		assert.Equal(t, source.OffRamp.MaxGasBufferToUpdateState, merged.OffRamp.MaxGasBufferToUpdateState)
		require.NotNil(t, merged.FeeQuoter.MaxFeeJuelsPerMsg)
		assert.True(t, merged.FeeQuoter.MaxFeeJuelsPerMsg.Cmp(source.FeeQuoter.MaxFeeJuelsPerMsg) == 0)
		require.NotNil(t, merged.FeeQuoter.USDPerLINK)
		assert.True(t, merged.FeeQuoter.USDPerLINK.Cmp(source.FeeQuoter.USDPerLINK) == 0)
	})

	t.Run("source overwrites base for slice fields", func(t *testing.T) {
		base := DeployContractParams{
			CommitteeVerifiers: []CommitteeVerifierDeployParams{
				{Version: v170, Qualifier: "base"},
			},
			Executors: []ExecutorDeployParams{
				{Version: v170, Qualifier: "execBase"},
			},
		}
		source := DeployContractParams{
			CommitteeVerifiers: []CommitteeVerifierDeployParams{
				{Version: v170, Qualifier: "source1"},
				{Version: v170, Qualifier: "source2"},
			},
			Executors: []ExecutorDeployParams{
				{Version: v170, Qualifier: "execSource"},
			},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		require.Len(t, merged.CommitteeVerifiers, 2)
		assert.Equal(t, "source1", merged.CommitteeVerifiers[0].Qualifier)
		assert.Equal(t, "source2", merged.CommitteeVerifiers[1].Qualifier)
		require.Len(t, merged.Executors, 1)
		assert.Equal(t, "execSource", merged.Executors[0].Qualifier)
	})

	t.Run("source overwrites base for MockReceivers", func(t *testing.T) {
		base := DeployContractParams{
			MockReceivers: []MockReceiverDeployParams{
				{Version: v170, Qualifier: "mockBase"},
			},
		}
		source := DeployContractParams{
			MockReceivers: []MockReceiverDeployParams{
				{Version: v170, Qualifier: "mockSource", OptionalThreshold: 2},
			},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		require.Len(t, merged.MockReceivers, 1)
		assert.Equal(t, "mockSource", merged.MockReceivers[0].Qualifier)
		assert.Equal(t, uint8(2), merged.MockReceivers[0].OptionalThreshold)
	})

	t.Run("merge is idempotent when base and source are equal", func(t *testing.T) {
		params := DeployContractParams{
			RMNRemote: RMNRemoteDeployParams{Version: v170, LegacyRMN: "0xSame"},
			OnRamp:    OnRampDeployParams{Version: v170, MaxUSDCentsPerMessage: 100},
		}

		merged, err := params.MergeWithOverrideIfNotEmpty(params)
		require.NoError(t, err)
		assert.Equal(t, params, merged)
	})

	t.Run("MockReceivers with AddressRef slices merge correctly", func(t *testing.T) {
		ref := datastore.AddressRef{
			ChainSelector: 1,
			Type:          "Verifier",
			Version:       v170,
			Address:       "0xVerifier",
		}
		base := DeployContractParams{}
		source := DeployContractParams{
			MockReceivers: []MockReceiverDeployParams{
				{Version: v170, RequiredVerifiers: []datastore.AddressRef{ref}, Qualifier: "q"},
			},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)
		require.Len(t, merged.MockReceivers, 1)
		require.Len(t, merged.MockReceivers[0].RequiredVerifiers, 1)
		assert.Equal(t, uint64(1), merged.MockReceivers[0].RequiredVerifiers[0].ChainSelector)
		assert.Equal(t, "0xVerifier", merged.MockReceivers[0].RequiredVerifiers[0].Address)
	})

	// Test merge when source has the shape of output from importConfig (v1.5 RMN) + importConfigFromv1_6_0:
	// RMNRemote.LegacyRMN, OffRamp.GasForCallExactCheck, OnRamp.FeeAggregator, and optionally
	// CommitteeVerifiers/Executors with FeeAggregator set.
	t.Run("merge with importConfigFromv1_6_0-style source populates RMNRemote OffRamp CommitteeVerifiers OnRamp FeeQuoter Executors", func(t *testing.T) {
		// Base: full params as from topology/defaults
		base := DeployContractParams{
			RMNRemote: RMNRemoteDeployParams{
				Version:   v170,
				LegacyRMN: "0xBaseLegacyRMN",
			},
			OffRamp: OffRampDeployParams{
				Version:                   v170,
				GasForCallExactCheck:      1000,
				MaxGasBufferToUpdateState: 100,
			},
			CommitteeVerifiers: []CommitteeVerifierDeployParams{
				{Version: v170, Qualifier: "committee1", FeeAggregator: "0xBaseFeeAgg"},
				{Version: v170, Qualifier: "committee2", FeeAggregator: "0xBaseFeeAgg"},
			},
			OnRamp: OnRampDeployParams{
				Version:               v170,
				FeeAggregator:         "0xBaseOnRampFeeAgg",
				MaxUSDCentsPerMessage: 50,
			},
			FeeQuoter: FeeQuoterDeployParams{
				Version:                        v170,
				MaxFeeJuelsPerMsg:              big.NewInt(1e18),
				LINKPremiumMultiplierWeiPerEth: 1e18,
				WETHPremiumMultiplierWeiPerEth: 1e18,
			},
			Executors: []ExecutorDeployParams{
				{Version: v170, Qualifier: "exec1", DynamicConfig: ExecutorDynamicDeployConfig{FeeAggregator: "0xBaseExecFeeAgg"}},
				{Version: v170, Qualifier: "exec2", DynamicConfig: ExecutorDynamicDeployConfig{FeeAggregator: "0xBaseExecFeeAgg"}},
			},
		}

		// Source: values as populated by importConfig (RMNRemote.LegacyRMN from v1.5) and importConfigFromv1_6_0
		// (OnRamp.FeeAggregator, OffRamp.GasForCallExactCheck; and FeeAggregator on CommitteeVerifiers/Executors when those slices exist)
		importedLegacyRMN := "0xImportedLegacyRMN"
		importedGasForCallExactCheck := uint16(5000)
		importedFeeAggregator := "0xImportedFeeAggregator"
		source := DeployContractParams{
			RMNRemote: RMNRemoteDeployParams{
				LegacyRMN: importedLegacyRMN,
				// Version not set by import
			},
			OffRamp: OffRampDeployParams{
				GasForCallExactCheck: importedGasForCallExactCheck,
				// Version, MaxGasBufferToUpdateState not set by import
			},
			CommitteeVerifiers: []CommitteeVerifierDeployParams{
				{Version: v170, Qualifier: "importedCommittee", FeeAggregator: importedFeeAggregator},
			},
			OnRamp: OnRampDeployParams{
				FeeAggregator: importedFeeAggregator,
				// Version, MaxUSDCentsPerMessage not set by import
			},
			FeeQuoter: FeeQuoterDeployParams{
				Version:           v170,
				MaxFeeJuelsPerMsg: big.NewInt(2e18),
				USDPerLINK:        big.NewInt(5e6),
				USDPerWETH:        big.NewInt(6e6),
			},
			Executors: []ExecutorDeployParams{
				{Version: v170, Qualifier: "importedExec", DynamicConfig: ExecutorDynamicDeployConfig{FeeAggregator: importedFeeAggregator}},
			},
		}

		merged, err := base.MergeWithOverrideIfNotEmpty(source)
		require.NoError(t, err)

		// RMNRemote: merged should have source's LegacyRMN (as set by importConfig from v1.5)
		assert.Equal(t, importedLegacyRMN, merged.RMNRemote.LegacyRMN, "RMNRemote.LegacyRMN should come from import")

		// OffRamp: merged should have source's GasForCallExactCheck (as set by importConfigFromv1_6_0)
		assert.Equal(t, importedGasForCallExactCheck, merged.OffRamp.GasForCallExactCheck, "OffRamp.GasForCallExactCheck should come from import")

		// CommitteeVerifiers: merged should have source's slice (import-style: FeeAggregator set)
		require.Len(t, merged.CommitteeVerifiers, 1)
		assert.Equal(t, "importedCommittee", merged.CommitteeVerifiers[0].Qualifier)
		assert.Equal(t, importedFeeAggregator, merged.CommitteeVerifiers[0].FeeAggregator)

		// OnRamp: merged should have source's FeeAggregator (as set by importConfigFromv1_6_0)
		assert.Equal(t, importedFeeAggregator, merged.OnRamp.FeeAggregator, "OnRamp.FeeAggregator should come from import")

		// FeeQuoter: merged should have source's values
		require.NotNil(t, merged.FeeQuoter.MaxFeeJuelsPerMsg)
		assert.True(t, merged.FeeQuoter.MaxFeeJuelsPerMsg.Cmp(source.FeeQuoter.MaxFeeJuelsPerMsg) == 0)
		require.NotNil(t, merged.FeeQuoter.USDPerLINK)
		assert.True(t, merged.FeeQuoter.USDPerLINK.Cmp(source.FeeQuoter.USDPerLINK) == 0)
		require.NotNil(t, merged.FeeQuoter.USDPerWETH)
		assert.True(t, merged.FeeQuoter.USDPerWETH.Cmp(source.FeeQuoter.USDPerWETH) == 0)

		// Executors: merged should have source's slice (import-style: DynamicConfig.FeeAggregator set)
		require.Len(t, merged.Executors, 1)
		assert.Equal(t, "importedExec", merged.Executors[0].Qualifier)
		assert.Equal(t, importedFeeAggregator, merged.Executors[0].DynamicConfig.FeeAggregator)
	})
}

type transferOwnershipAwareAdapter struct{}

func (a *transferOwnershipAwareAdapter) GetDefaultDeployContractParams(_ uint64) DeployContractParams {
	return DeployContractParams{}
}

func (a *transferOwnershipAwareAdapter) ResolveDeployAddresses(_ cldf.Environment, _ uint64) (DeployChainResolvedAddresses, error) {
	return DeployChainResolvedAddresses{DeployerContract: "0x0000000000000000000000000000000000001234"}, nil
}

func (a *transferOwnershipAwareAdapter) BuildDeployContractParams(in BuildDeployContractParamsInput) (DeployContractParams, error) {
	return DeployContractParams{CommitteeVerifiers: in.CommitteeVerifiers}, nil
}

func (a *transferOwnershipAwareAdapter) DeployChainContracts() *cldf_ops.Sequence[DeployChainContractsInput, DeployChainContractsOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"transfer-ownership-aware-deploy",
		semver.MustParse("2.0.0"),
		"Returns ownership refs only when transfer is required",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, input DeployChainContractsInput) (DeployChainContractsOutput, error) {
			output := DeployChainContractsOutput{
				OnChainOutput: sequences.OnChainOutput{
					Addresses: []datastore.AddressRef{
						{
							ChainSelector: input.ChainSelector,
							Type:          "Router",
							Version:       semver.MustParse("2.0.0"),
							Address:       "0x00000000000000000000000000000000000000AA",
						},
					},
				},
			}
			if !input.DeployerKeyOwned {
				output.RefsToTransferOwnership = []datastore.AddressRef{
					{
						ChainSelector: input.ChainSelector,
						Type:          "OnRamp",
						Version:       semver.MustParse("2.0.0"),
						Address:       "0x00000000000000000000000000000000000000BB",
					},
					{
						ChainSelector: input.ChainSelector,
						Type:          "OffRamp",
						Version:       semver.MustParse("2.0.0"),
						Address:       "0x00000000000000000000000000000000000000CC",
					},
				}
			}
			return output, nil
		},
	)
}

func TestDeployChainContracts_TransferOwnershipRefsBehavior(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	adapter := &transferOwnershipAwareAdapter{}
	registry := NewDeployChainContractsRegistry()
	registry.Register(chainsel.FamilyEVM, adapter)

	resolved, err := registry.GetByChain(sel)
	require.NoError(t, err)

	deps := cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{
		sel: cldfevm.Chain{Selector: sel},
	})
	bundle := cldf_ops.NewBundle(
		func() context.Context { return context.Background() },
		logger.Test(t),
		cldf_ops.NewMemoryReporter(),
	)

	reportWithoutOwnership, err := cldf_ops.ExecuteSequence(
		bundle,
		resolved.DeployChainContracts(),
		deps,
		DeployChainContractsInput{
			ChainSelector:    sel,
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err)
	assert.Empty(t, reportWithoutOwnership.Output.RefsToTransferOwnership)

	reportWithOwnership, err := cldf_ops.ExecuteSequence(
		bundle,
		resolved.DeployChainContracts(),
		deps,
		DeployChainContractsInput{
			ChainSelector:    sel,
			DeployerKeyOwned: false,
		},
	)
	require.NoError(t, err)
	require.Len(t, reportWithOwnership.Output.RefsToTransferOwnership, 2)
	assert.Equal(t, "OnRamp", string(reportWithOwnership.Output.RefsToTransferOwnership[0].Type))
	assert.Equal(t, "OffRamp", string(reportWithOwnership.Output.RefsToTransferOwnership[1].Type))
}
