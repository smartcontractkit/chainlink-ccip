package adapters_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"

	adapters1_7 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
)

func TestLaneMigrator(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "happy path",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainA := chainsel.ETHEREUM_MAINNET.Selector
			chainB := chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainA, chainB}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			mcmsRegistry := changesets.GetRegistry()

			// On each chain, deploy chain contracts
			ds := datastore.NewMemoryDataStore()
			e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
			deployLaneContractsToDatastore(t, e, chainA, ds)
			e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
			deployLaneContractsToDatastore(t, e, chainB, ds)

			// Overwrite datastore in the environment
			e.DataStore = ds.Seal()

			// Configure chains for lanes. The FeeQuoter dest chain values below are set
			// away from the migrator's defaults so the assertions further down observe
			// the migrator's write rather than the value the lane was built with.
			deployer := e.BlockChains.EVMChains()[chainA].DeployerKey.From.Hex()
			e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
			_, err = v2changesets.ConfigureChainsForLanesFromTopology(
				ccvadapters.GetCommitteeVerifierContractRegistry(),
				ccvadapters.GetChainFamilyRegistry(),
				mcmsRegistry,
			).Apply(*e, v2changesets.ConfigureChainsForLanesFromTopologyConfig{
				Topology: bidirectionalLaneTopology(deployer, chainA, chainB),
				BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
					Lanes: []v2changesets.CrossFamilyLanePair{
						{
							ChainA: chainA,
							ChainB: chainB,
							ChainAOverrides: &v2changesets.ChainOverrides{
								RemoteChainCfg: v2changesets.PartialRemoteChainConfig{
									FeeQuoterDestChainConfig: ccvadapters.FeeQuoterDestChainConfigOverrides{
										MaxDataBytes:              new(uint32(20_000)),
										MaxPerMsgGasLimit:         new(uint32(3_000_000)),
										DestGasPerPayloadByteBase: new(uint8(16)),
									},
								},
							},
						},
					},
					MCMS: mcms.Input{},
				},
			})
			require.NoError(t, err, "Failed to apply ConfigureChainsForLanesFromTopology changeset")
			// now apply the lane migrator
			mReg := deploy.GetLaneMigratorRegistry()
			e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
			cs := deploy.LaneMigrateToNewVersionChangeset(mReg, mcmsRegistry)
			_, err = cs.Apply(*e, deploy.LaneMigratorConfig{
				Input: map[uint64]deploy.LaneMigratorConfigPerChain{
					chainA: {
						RemoteChains:  []uint64{chainB},
						RouterVersion: semver.MustParse("1.2.0"),
						RampVersion:   semver.MustParse("2.0.0"),
					},
				},
			})
			require.NoError(t, err)
			evmChain1 := e.BlockChains.EVMChains()[chainA]
			routerAddr, err := datastore_utils.FindAndFormatRef(
				e.DataStore,
				datastore.AddressRef{
					Type:    "Router",
					Version: semver.MustParse("1.2.0"),
				}, chainA, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err)
			onRampAddr, err := datastore_utils.FindAndFormatRef(
				e.DataStore,
				datastore.AddressRef{
					Type:    "OnRamp",
					Version: semver.MustParse("2.0.0"),
				}, chainA, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err)
			offRampAddr, err := datastore_utils.FindAndFormatRef(
				e.DataStore,
				datastore.AddressRef{
					Type:    "OffRamp",
					Version: semver.MustParse("2.0.0"),
				}, chainA, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err)
			// query router
			routerC, err := router.NewRouter(routerAddr, evmChain1.Client)
			require.NoError(t, err)
			onRamp, err := routerC.GetOnRamp(nil, chainB)
			require.NoError(t, err)
			require.Equal(t, onRampAddr, onRamp)
			offRamps, err := routerC.GetOffRamps(nil)
			require.NoError(t, err)
			require.Contains(t, offRamps, router.RouterOffRamp{
				SourceChainSelector: chainB,
				OffRamp:             offRampAddr,
			})
			feeQuoterAddr, err := datastore_utils.FindAndFormatRef(
				e.DataStore,
				datastore.AddressRef{
					Type:    datastore.ContractType(fee_quoter.ContractType),
					Version: fee_quoter.Version,
				}, chainA, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err)
			opOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fee_quoter.GetDestChainConfig, evmChain1, contract_utils.FunctionInput[uint64]{
				ChainSelector: chainA,
				Address:       feeQuoterAddr,
				Args:          chainB,
			})
			require.NoError(t, err)
			destConfig := opOut.Output
			require.Equal(t, uint32(15_000_000), destConfig.MaxPerMsgGasLimit) // MaxPerMsgGasLimit should be 15 million for Arbitrum
			require.Equal(t, adapters1_7.DefaultMaxDataBytes, destConfig.MaxDataBytes)
			require.Equal(t, uint8(100), destConfig.DestGasPerPayloadByteBase)
		})
	}
}
