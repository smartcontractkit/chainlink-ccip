package adapters_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	adapters1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	v1_7_0 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
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
			chainA := uint64(5009297550715157269)
			chainB := uint64(4949039107694359620)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainA, chainB}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			chainFamilyRegistry := lanes.GetLaneAdapterRegistry()
			mcmsRegistry := changesets.GetRegistry()

			// On each chain, deploy chain contracts
			ds := datastore.NewMemoryDataStore()
			for _, chainSel := range []uint64{chainA, chainB} {
				create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
					ChainSelector:  chainSel,
					Args: create2_factory.ConstructorArgs{
						AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
					},
				}, nil)
				require.NoError(t, err, "Failed to deploy CREATE2Factory")

				deployChainOut, err := v1_7_0.DeployChainContracts(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_7_0.DeployChainContractsCfg]{
					Cfg: v1_7_0.DeployChainContractsCfg{
						ChainSel:         chainSel,
						CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
						Params:           testsetup.CreateBasicContractParams(),
						DeployerKeyOwned: true,
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = ds.Merge(deployChainOut.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts changeset")
			}

			// Overwrite datastore in the environment
			e.DataStore = ds.Seal()

			// Configure chains for lanes
			e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
			_, err = lanes.ConnectChains(chainFamilyRegistry, mcmsRegistry).Apply(*e, lanes.ConnectChainsConfig{
				Lanes: []lanes.LaneConfig{
					{
						ChainA:  makeChainConfig(chainA, chainB),
						ChainB:  makeChainConfig(chainB, chainA),
						Version: semver.MustParse("2.0.0"),
					},
				},
			})
			require.NoError(t, err, "Failed to apply ConnectChains changeset")
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
			require.Equal(t, adapters1_7.DefaultMaxPerMsgGasLimit, destConfig.MaxPerMsgGasLimit)
			require.Equal(t, adapters1_7.DefaultMaxDataBytes, destConfig.MaxDataBytes)
		})
	}
}
