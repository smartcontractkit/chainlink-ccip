package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
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
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{5009297550715157269, 4356164186791070119}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			mcmsRegistry := cs_core.NewMCMSReaderRegistry()

			// Deploy both chains
			runningDataStore := datastore.NewMemoryDataStore()
			for _, evmChain := range e.BlockChains.EVMChains() {
				out, err := changesets.DeployChainContracts(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
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

			_, err = changesets.ConfigureChainForLanes(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ConfigureChainForLanesCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.ConfigureChainForLanesCfg{
					ChainSel: 5009297550715157269,
					RemoteChains: map[uint64]changesets.RemoteChainConfig{
						4356164186791070119: {
							AllowTrafficFrom: true,
							CCIPMessageDest: datastore.AddressRef{
								Type:    datastore.ContractType(offramp.ContractType),
								Version: semver.MustParse("1.7.0"),
							},
							CCIPMessageSource: datastore.AddressRef{
								Type:    datastore.ContractType(onramp.ContractType),
								Version: semver.MustParse("1.7.0"),
							},
							DefaultInboundCCVs: []datastore.AddressRef{
								{Type: datastore.ContractType(committee_verifier.ContractType), Version: semver.MustParse("1.7.0")},
							},
							DefaultOutboundCCVs: []datastore.AddressRef{
								{Type: datastore.ContractType(committee_verifier.ContractType), Version: semver.MustParse("1.7.0")},
							},
							DefaultExecutor: datastore.AddressRef{
								Type:    datastore.ContractType(executor.ContractType),
								Version: semver.MustParse("1.7.0"),
							},
							CommitteeVerifierDestChainConfig: sequences.CommitteeVerifierDestChainConfig{
								AllowlistEnabled:          false,
								AddedAllowlistedSenders:   nil,
								RemovedAllowlistedSenders: nil,
								FeeUSDCents:               50,
								GasForVerification:        50_000,
								PayloadSizeBytes:          6*64 + 2*32,
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
