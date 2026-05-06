package changesets_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/changesets"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestConfigureRMNCurseAdmins_Apply(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	chain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy RMN with no initial curse admins.
	e.DataStore = datastore.NewMemoryDataStore().Seal()
	mcmsRegistry := cs_core.GetRegistry()
	deployOut, err := changesets.DeployRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployRMNCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployRMNCfg{
			ChainSel:    chainSelector,
			CurseAdmins: []common.Address{},
		},
	})
	require.NoError(t, err)

	rmnAddrs, err := deployOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)
	rmnRef := rmnAddrs[0]

	// Seed the environment datastore with the deployed RMN address.
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnRef))
	e.DataStore = ds.Seal()

	newCurseAdmin := chain.DeployerKey.From

	// Apply the ConfigureRMNCurseAdmins changeset to add a curse admin.
	_, err = changesets.ConfigureRMNCurseAdmins(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ConfigureRMNCurseAdminsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.ConfigureRMNCurseAdminsCfg{
			ChainSel: chainSelector,
			RMN:      datastore.AddressRef{Type: rmnRef.Type, Version: rmnRef.Version},
			Args: rmnops.AuthorizedCallerArgs{
				AddedCallers: []common.Address{newCurseAdmin},
			},
		},
	})
	require.NoError(t, err)
}

func TestConfigureRMNCurseAdmins_VerifyPreconditions(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	rmnRef := datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(rmnops.ContractType),
		Version:       rmnops.Version,
		Address:       common.HexToAddress("0x01").Hex(),
	}
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnRef))
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()

	tests := []struct {
		desc        string
		cfg         changesets.ConfigureRMNCurseAdminsCfg
		expectedErr string
	}{
		{
			desc: "valid input",
			cfg: changesets.ConfigureRMNCurseAdminsCfg{
				ChainSel: chainSelector,
				RMN:      datastore.AddressRef{Type: rmnRef.Type, Version: rmnRef.Version},
				Args: rmnops.AuthorizedCallerArgs{
					AddedCallers: []common.Address{common.HexToAddress("0x02")},
				},
			},
		},
		{
			desc: "RMN ref not in datastore",
			cfg: changesets.ConfigureRMNCurseAdminsCfg{
				ChainSel: chainSelector,
				RMN:      datastore.AddressRef{Type: "UnknownContract", Version: rmnops.Version},
				Args: rmnops.AuthorizedCallerArgs{
					AddedCallers: []common.Address{common.HexToAddress("0x02")},
				},
			},
			expectedErr: "failed to resolve RMN ref",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := changesets.ConfigureRMNCurseAdmins(mcmsRegistry).VerifyPreconditions(*e, cs_core.WithMCMS[changesets.ConfigureRMNCurseAdminsCfg]{
				MCMS: mcms.Input{},
				Cfg:  test.cfg,
			})
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
