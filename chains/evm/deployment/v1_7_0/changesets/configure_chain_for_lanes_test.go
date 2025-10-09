package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/ccv_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/executor_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/off_ramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func TestConfigureChainForLanes_Apply(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "valid input",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
				5009297550715157269: {NumAdditionalAccounts: 1},
				4356164186791070119: {NumAdditionalAccounts: 1},
			})
			require.NoError(t, err, "Failed to create test environment")
			evmChains := e.BlockChains.EVMChains()

			mcmsRegistry := cs_core.NewMCMSReaderRegistry()

			// Deploy both chains
			runningDataStore := datastore.NewMemoryDataStore()
			for _, evmChain := range evmChains {
				out, err := changesets.DeployChainContracts(mcmsRegistry).Apply(e, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
					MCMS: mcms.Input{},
					Cfg: changesets.DeployChainContractsCfg{
						ChainSel: evmChain.Selector,
						Params:   testsetup.CreateBasicContractParams(),
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = runningDataStore.Merge(out.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts")
			}
			e.DataStore = runningDataStore.Seal() // Override datastore in environment to include deployed contracts

			_, err = changesets.ConfigureChainForLanes(mcmsRegistry).Apply(e, cs_core.WithMCMS[changesets.ConfigureChainForLanesCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.ConfigureChainForLanesCfg{
					ChainSel: 5009297550715157269,
					RemoteChains: map[uint64]changesets.RemoteChainConfig{
						4356164186791070119: {
							AllowTrafficFrom: true,
							CCIPMessageDest: datastore.AddressRef{
								Type:    datastore.ContractType(off_ramp.ContractType),
								Version: semver.MustParse("1.7.0"),
							},
							CCIPMessageSource: datastore.AddressRef{
								Type:    datastore.ContractType(ccv_proxy.ContractType),
								Version: semver.MustParse("1.7.0"),
							},
							DefaultCCVOffRamps: []datastore.AddressRef{
								{Type: datastore.ContractType(committee_verifier.ContractType), Version: semver.MustParse("1.7.0")},
							},
							DefaultCCVOnRamps: []datastore.AddressRef{
								{Type: datastore.ContractType(committee_verifier.ContractType), Version: semver.MustParse("1.7.0")},
							},
							DefaultExecutor: datastore.AddressRef{
								Type:    datastore.ContractType(executor_onramp.ContractType),
								Version: semver.MustParse("1.7.0"),
							},
							CommitteeVerifierDestChainConfig: sequences.CommitteeVerifierDestChainConfig{
								AllowlistEnabled: false,
							},
							FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
						},
					},
				},
			})
			require.NoError(t, err, "Failed to apply ConfigureChainForLanes changeset")
		})
	}
}
