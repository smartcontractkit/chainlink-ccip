package adapters_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_adapters "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	v1_7_0 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	v1_7_0_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func makeChainConfig(chainSelector uint64, remoteChainSelector uint64) v1_7_0_changesets.ChainConfig {
	return v1_7_0_changesets.ChainConfig{
		ChainSelector: chainSelector,
		Router: datastore.AddressRef{
			Type:    datastore.ContractType(router.ContractType),
			Version: semver.MustParse("1.2.0"),
		},
		OnRamp: datastore.AddressRef{
			Type:    datastore.ContractType(onramp.ContractType),
			Version: semver.MustParse("1.7.0"),
		},
		CommitteeVerifiers: []adapters.CommitteeVerifierConfig[datastore.AddressRef]{
			{
				CommitteeVerifier: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(committee_verifier.ContractType),
						Version: semver.MustParse("1.7.0"),
					},
					{
						Type:    datastore.ContractType(committee_verifier.ResolverType),
						Version: semver.MustParse("1.7.0"),
					},
				},
				RemoteChains: map[uint64]adapters.CommitteeVerifierRemoteChainConfig{
					remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
		FeeQuoter: datastore.AddressRef{
			Type:    datastore.ContractType(fee_quoter.ContractType),
			Version: semver.MustParse("1.7.0"),
		},
		OffRamp: datastore.AddressRef{
			Type:    datastore.ContractType(offramp.ContractType),
			Version: semver.MustParse("1.7.0"),
		},
		RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
			remoteChainSelector: {
				AllowTrafficFrom: true,
				OnRamps: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(onramp.ContractType),
						Version: semver.MustParse("1.7.0"),
					},
				},
				OffRamp: datastore.AddressRef{
					Type:    datastore.ContractType(offramp.ContractType),
					Version: semver.MustParse("1.7.0"),
				},
				DefaultInboundCCVs: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(committee_verifier.ContractType),
						Version: semver.MustParse("1.7.0"),
					},
				},
				DefaultOutboundCCVs: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(committee_verifier.ContractType),
						Version: semver.MustParse("1.7.0"),
					},
				},
				DefaultExecutor: datastore.AddressRef{
					Type:      datastore.ContractType(executor.ContractType),
					Version:   semver.MustParse("1.7.0"),
					Qualifier: "default",
				},
				FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
				ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
				AddressBytesLength:       20,
				BaseExecutionGasCost:     80_000,
			},
		},
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

			chainFamilyRegistry := adapters.NewChainFamilyRegistry()
			chainFamilyRegistry.RegisterChainFamily("evm", &evm_adapters.ChainFamilyAdapter{})
			mcmsRegistry := changesets.NewMCMSReaderRegistry()

			// On each chain, deploy chain contracts
			ds := datastore.NewMemoryDataStore()
			for _, chainSel := range []uint64{chainA, chainB} {
				create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
					ChainSelector:  chainSel,
					Args: create2_factory.ConstructorArgs{
						AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
					},
				}, nil)
				require.NoError(t, err, "Failed to deploy CREATE2Factory")

				deployChainOut, err := v1_7_0.DeployChainContracts(mcmsRegistry).Apply(*e, changesets.WithMCMS[v1_7_0.DeployChainContractsCfg]{
					Cfg: v1_7_0.DeployChainContractsCfg{
						ChainSel:       chainSel,
						CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
						Params:         testsetup.CreateBasicContractParams(),
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = ds.Merge(deployChainOut.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts changeset")
			}

			// Overwrite datastore in the environment
			e.DataStore = ds.Seal()

			// Configure chains for lanes
			_, err = v1_7_0_changesets.ConfigureChainsForLanes(chainFamilyRegistry, mcmsRegistry).Apply(*e, v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					makeChainConfig(chainA, chainB),
					makeChainConfig(chainB, chainA),
				},
			})
			require.NoError(t, err, "Failed to apply ConfigureChainsForLanes changeset")
		})
	}
}
