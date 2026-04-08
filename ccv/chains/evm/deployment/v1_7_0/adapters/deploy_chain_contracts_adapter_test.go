package adapters_test

import (
	"sort"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	rmnops1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	offrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	seq1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
)

// importConfigAddressRefs are address refs for SetContractParamsFromImportedConfig.
// Chain 5009297550715157269 has OnRamp 1.6.0 + OffRamp 1.6.0 (and optionally RMN 1.5.0).
// Chain 4949039107694359620 has only RMN 1.5.0 (v1.6 path returns empty; v1.5 path fills LegacyRMN and GasForCallExactCheck).
var (
	legacyRMNAddress1       = common.HexToAddress("0x8888888888888888888888888888888888888888")
	legacyRMNAddress2       = common.HexToAddress("0x9999999999999999999999999999999999999999")
	feeAggregatorAddress    = common.HexToAddress("0xFEE0000000000000000000000000000000000001")
	importConfigAddressRefs = []datastore.AddressRef{
		{Address: "0x6666666666666666666666666666666666666666", ChainSelector: 5009297550715157269, Type: datastore.ContractType(onrampops_v160.ContractType), Version: onrampops_v160.Version},
		{Address: "0x7777777777777777777777777777777777777777", ChainSelector: 5009297550715157269, Type: datastore.ContractType(offrampops_v160.ContractType), Version: offrampops_v160.Version},
		{Address: legacyRMNAddress1.String(), ChainSelector: 5009297550715157269, Type: datastore.ContractType(rmnops1_5.ContractType), Version: semver.MustParse("1.5.0")},
		{Address: legacyRMNAddress2.String(), ChainSelector: 4949039107694359620, Type: datastore.ContractType(rmnops1_5.ContractType), Version: semver.MustParse("1.5.0")},
	}

	// importConfigContractMetadata holds OnRamp 1.6.0 and OffRamp 1.6.0 metadata for the chain that has both.
	importConfigContractMetadata = []datastore.ContractMetadata{
		{
			Address:       "0x6666666666666666666666666666666666666666",
			ChainSelector: 5009297550715157269,
			Metadata: seq1_6.OnRampImportConfigSequenceOutput{
				DestChainCfgs: map[uint64]onrampops_v160.GetDestChainConfigResult{},
				StaticConfig: onrampops_v160.StaticConfig{
					ChainSelector:      5009297550715157269,
					RmnRemote:          common.HexToAddress("0x8888888888888888888888888888888888888888"),
					NonceManager:       common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
					TokenAdminRegistry: common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
				},
				DynamicConfig: onrampops_v160.DynamicConfig{
					FeeQuoter:          common.HexToAddress("0x1111111111111111111111111111111111111111"),
					FeeAggregator:      feeAggregatorAddress,
					AllowlistAdmin:     common.Address{},
					MessageInterceptor: common.Address{},
				},
			},
		},
		{
			Address:       "0x7777777777777777777777777777777777777777",
			ChainSelector: 5009297550715157269,
			Metadata: seq1_6.OffRampImportConfigSequenceOutput{
				SourceChainCfgs: map[uint64]offrampops_v160.SourceChainConfig{},
				StaticConfig: offrampops_v160.StaticConfig{
					ChainSelector:        5009297550715157269,
					GasForCallExactCheck: 6000,
					RmnRemote:            common.HexToAddress("0x8888888888888888888888888888888888888888"),
					TokenAdminRegistry:   common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
					NonceManager:         common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
				},
				DynamicConfig: offrampops_v160.DynamicConfig{},
			},
		},
	}
)

// dummyDeployContractParams returns params with Executors, CommitteeVerifiers, OnRamp, and OffRamp
// populated with dummy values. OnRamp/OffRamp use zero/empty so that when merged as source into
// generated config they do not overwrite the imported OnRamp/OffRamp/RMN (per MergeWithOverrideIfNotEmpty behavior).
func dummyDeployContractParams() ccvadapters.DeployContractParams {
	v170 := semver.MustParse("2.0.0")
	return ccvadapters.DeployContractParams{
		OnRamp: ccvadapters.OnRampDeployParams{
			Version:       v170,
			FeeAggregator: "0xDummyOnRampFeeAgg",
		},
		OffRamp: ccvadapters.OffRampDeployParams{
			Version: v170,
			// GasForCallExactCheck etc. left zero so merge keeps generated (imported) values.
		},
		Executors: []ccvadapters.ExecutorDeployParams{
			{
				Version:       v170,
				MaxCCVsPerMsg: 3,
				DynamicConfig: ccvadapters.ExecutorDynamicDeployConfig{
					FeeAggregator:         "0xDummyExecutorFeeAgg",
					MinBlockConfirmations: 5,
					CcvAllowlistEnabled:   true,
				},
				Qualifier: "dummy-exec",
			},
		},
		CommitteeVerifiers: []ccvadapters.CommitteeVerifierDeployParams{
			{
				Version:          v170,
				FeeAggregator:    "0xDummyCVFeeAgg",
				StorageLocations: []string{"https://dummy.store"},
				Qualifier:        "dummy-cv",
			},
		},
	}
}

func TestSetContractParamsFromImportedConfig(t *testing.T) {
	chainSelectors := make(map[uint64]bool)
	ds := datastore.NewMemoryDataStore()
	for _, ref := range importConfigAddressRefs {
		chainSelectors[ref.ChainSelector] = true
		err := ds.Addresses().Add(ref)
		require.NoError(t, err, "Failed to add address ref %+v to datastore", ref)
	}
	err := sequences.WriteMetadataToDatastore(ds, sequences.Metadata{Contracts: importConfigContractMetadata})
	require.NoError(t, err, "Failed to write contract metadata to datastore")

	chainSelectorList := make([]uint64, 0, len(chainSelectors))
	for selector := range chainSelectors {
		chainSelectorList = append(chainSelectorList, selector)
	}
	sort.Slice(chainSelectorList, func(i, j int) bool { return chainSelectorList[i] < chainSelectorList[j] })

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chainSelectorList),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = ds.Seal()

	adapter := &adapters.EVMDeployChainContractsAdapter{}

	for _, chainSelector := range chainSelectorList {
		_, ok := e.BlockChains.EVMChains()[chainSelector]
		require.True(t, ok, "Chain with selector %d should exist", chainSelector)

		existingAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector))
		contractMeta := e.DataStore.ContractMetadata().Filter(datastore.ContractMetadataByChainSelector(chainSelector))
		dummyBase := dummyDeployContractParams()
		input := ccvadapters.DeployChainConfigCreatorInput{
			ChainSelector:      chainSelector,
			ExistingAddresses:  existingAddresses,
			ContractMeta:       contractMeta,
			UserProvidedConfig: dummyBase,
		}

		report, err := cldf_ops.ExecuteSequence(
			e.OperationsBundle,
			adapter.SetContractParamsFromImportedConfig(),
			e.BlockChains,
			input,
		)

		require.NoError(t, err, "SetContractParamsFromImportedConfig should not error for chain %d", chainSelector)
		require.NotNil(t, report, "Report should not be nil for chain %d", chainSelector)

		generated := report.Output
		// Merge generated (imported) config with dummy base: generated.MergeWithOverrideIfNotEmpty(dummy).
		// the non-empty fields in generated (imported config) should overwrite dummy, while empty fields in generated should keep dummy's values after merge.
		merged, err := dummyBase.MergeWithOverrideIfNotEmpty(generated)
		require.NoError(t, err, "MergeWithOverrideIfNotEmpty should not error")
		switch chainSelector {
		case 5009297550715157269:
			// This chain has OnRamp 1.6.0 + OffRamp 1.6.0 + RMN 1.5.0: v1.6 and v1.5 merge.
			// v1.5's GasForCallExactCheck (5000) is overwritten by v1.6's value.
			require.Equal(t, feeAggregatorAddress, common.HexToAddress(generated.OnRamp.FeeAggregator),
				"OnRamp.FeeAggregator should come from v1.6 OnRamp DynamicConfig")
			require.Equal(t, uint16(6000), generated.OffRamp.GasForCallExactCheck,
				"OffRamp.GasForCallExactCheck comes from v1.6 path after merge (v1.6 overwrites v1.5)")
			require.Equal(t, legacyRMNAddress1, common.HexToAddress(generated.RMNRemote.LegacyRMN),
				"RMNRemote.LegacyRMN should come from v1.5 address ref")
			require.Equal(t, uint16(6000), merged.OffRamp.GasForCallExactCheck,
				"merged OffRamp.GasForCallExactCheck should come from generated")
			require.Equal(t, legacyRMNAddress1, common.HexToAddress(merged.RMNRemote.LegacyRMN),
				"merged RMNRemote.LegacyRMN should come from generated")
			// OnRamp, OffRamp, RMNRemote unchanged from generated (imported config).
			require.Equal(t, feeAggregatorAddress, common.HexToAddress(merged.OnRamp.FeeAggregator),
				"merged OnRamp.FeeAggregator should come from generated (imported config)")

			for i, exec := range dummyBase.Executors {
				exec.DynamicConfig.FeeAggregator = feeAggregatorAddress.String()
				require.Equal(t, exec, merged.Executors[i], "merged Executors should keep dummy values for empty fields in generated (imported) config")
			}
			for i, cv := range dummyBase.CommitteeVerifiers {
				cv.FeeAggregator = feeAggregatorAddress.String()
				require.Equal(t, cv, merged.CommitteeVerifiers[i], "merged CommitteeVerifiers should keep dummy values for empty fields in generated (imported) config")
			}

		case 4949039107694359620:
			// This chain has only RMN 1.5.0: v1.6 path returns empty (no OnRamp 1.6.0), v1.5 fills LegacyRMN and default GasForCallExactCheck.
			require.Equal(t, legacyRMNAddress2, common.HexToAddress(generated.RMNRemote.LegacyRMN),
				"RMNRemote.LegacyRMN should come from v1.5 address ref")
			require.Equal(t, uint16(5000), generated.OffRamp.GasForCallExactCheck,
				"OffRamp.GasForCallExactCheck should be default from v1.5 path when no v1.6 OffRamp")
			require.Equal(t, uint16(5000), merged.OffRamp.GasForCallExactCheck,
				"merged OffRamp.GasForCallExactCheck should come from generated")
			require.Equal(t, legacyRMNAddress2, common.HexToAddress(merged.RMNRemote.LegacyRMN),
				"merged RMNRemote.LegacyRMN should come from generated")
			require.ElementsMatch(t, dummyBase.Executors, merged.Executors,
				"merged Executors should be the same as dummy since generated (imported config) is empty for this chain")
			require.ElementsMatch(t, dummyBase.CommitteeVerifiers, merged.CommitteeVerifiers,
				"merged CommitteeVerifiers should be the same as dummy since generated (imported config) is empty for this chain")
		}

		t.Logf("Successfully executed SetContractParamsFromImportedConfig for chain %d", chainSelector)
	}
}
