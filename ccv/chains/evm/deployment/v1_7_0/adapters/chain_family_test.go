package adapters_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	v1_7_0 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
)

func makeChainConfig(chainSelector uint64, remoteChainSelector uint64) lanes.ChainDefinition {
	return lanes.ChainDefinition{
		Selector: chainSelector,
		CommitteeVerifiers: []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
			{
				CommitteeVerifier: []datastore.AddressRef{
					{
						ChainSelector: chainSelector,
						Type:          datastore.ContractType(committee_verifier.ContractType),
						Version:       committee_verifier.Version,
					},
					{
						ChainSelector: chainSelector,
						Type:          datastore.ContractType(sequences.CommitteeVerifierResolverType),
						Version:       semver.MustParse("2.0.0"),
					},
				},
				RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
					remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
		DefaultInboundCCVs: []datastore.AddressRef{
			{
				ChainSelector: chainSelector,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       committee_verifier.Version,
			},
		},
		DefaultOutboundCCVs: []datastore.AddressRef{
			{
				ChainSelector: chainSelector,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       committee_verifier.Version,
			},
		},
		DefaultExecutor: datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(sequences.ExecutorProxyType),
			Version:       executor.Version,
			Qualifier:     "default",
		},
		FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
		ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
		AddressBytesLength:       20,
		BaseExecutionGasCost:     80_000,
	}
}

func TestChainFamilyAdapter(t *testing.T) {
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
		})
	}
}
