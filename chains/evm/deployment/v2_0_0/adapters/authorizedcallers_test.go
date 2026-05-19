package adapters_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/rmn"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/authorizedcallers"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// TestAuthorizedCallersAdapter_OperatorFlow verifies the end-to-end operator workflow:
//  1. Deploy RMN with no initial curse admins.
//  2. Initialize the adapter.
//  3. GetAllAuthorizedCallers → empty.
//  4. Apply an add via ConfigureAuthorizedCallersChangeset.
//  5. GetAllAuthorizedCallers → deployer present (reads directly from chain, always fresh).
//  6. Apply a remove → GetAllAuthorizedCallers → empty.
func TestAuthorizedCallersAdapter_OperatorFlow(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	e.DataStore = datastore.NewMemoryDataStore().Seal()
	mcmsRegistry := cs_core.GetRegistry()

	// Deploy RMN with no initial curse admins.
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

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnAddrs[0]))
	e.DataStore = ds.Seal()

	// Initialize the adapter.
	adapter := adapters.NewRMNAuthorizedCallersAdapter()
	applyIn := api.ApplyInput{
		ChainSelector: chainSelector,
		ContractType:  rmnops.ContractType,
		Version:       rmnops.Version,
	}
	require.NoError(t, adapter.Initialize(*e, applyIn))

	// Step 3 — empty.
	initial, err := adapter.GetAllAuthorizedCallers(*e, chainSelector, rmnops.ContractType, rmnops.Version)
	require.NoError(t, err)
	require.Empty(t, initial)

	// Step 4 — add deployer via the chain-agnostic changeset (operator-facing path).
	// The changeset blank-imports this adapters package via init(), so the registry
	// already has the (EVM, RMN, 2.0.0) adapter registered.
	reg := api.GetAuthorizedCallersRegistry()
	_, err = api.ConfigureAuthorizedCallersChangeset(reg, mcmsRegistry).Apply(*e, api.Config{
		Force: true,
		Updates: []api.ApplyInput{
			{
				ChainSelector: chainSelector,
				ContractType:  rmnops.ContractType,
				Version:       rmnops.Version,
				Update:        api.CallerUpdate{AddedCallers: []api.Caller{deployer.Bytes()}},
			},
		},
	})
	require.NoError(t, err)

	// Step 5 — GetAllAuthorizedCallers bypasses OperationsBundle; always fresh.
	afterAdd, err := adapter.GetAllAuthorizedCallers(*e, chainSelector, rmnops.ContractType, rmnops.Version)
	require.NoError(t, err)
	require.Len(t, afterAdd, 1)
	require.Equal(t, deployer.Bytes(), []byte(afterAdd[0]))

	// Step 6 — remove.
	_, err = api.ConfigureAuthorizedCallersChangeset(reg, mcmsRegistry).Apply(*e, api.Config{
		Force: true,
		Updates: []api.ApplyInput{
			{
				ChainSelector: chainSelector,
				ContractType:  rmnops.ContractType,
				Version:       rmnops.Version,
				Update:        api.CallerUpdate{RemovedCallers: []api.Caller{deployer.Bytes()}},
			},
		},
	})
	require.NoError(t, err)

	afterRemove, err := adapter.GetAllAuthorizedCallers(*e, chainSelector, rmnops.ContractType, rmnops.Version)
	require.NoError(t, err)
	require.Empty(t, afterRemove)
}

// TestConfigureAuthorizedCallersChangeset_Force confirms that Force=false pre-filters
// already-present adds and already-absent removes, producing no batch ops (noop).
func TestConfigureAuthorizedCallersChangeset_Force(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	e.DataStore = datastore.NewMemoryDataStore().Seal()
	mcmsRegistry := cs_core.GetRegistry()

	// Deploy RMN with deployer pre-added as curse admin.
	deployOut, err := changesets.DeployRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployRMNCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployRMNCfg{
			ChainSel:    chainSelector,
			CurseAdmins: []common.Address{deployer},
		},
	})
	require.NoError(t, err)

	rmnAddrs, err := deployOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnAddrs[0]))
	e.DataStore = ds.Seal()

	reg := api.GetAuthorizedCallersRegistry()
	// Force=false: add that is already present and remove of non-present address.
	// Changeset should produce no MCMS proposal (all ops filtered out).
	csOut, err := api.ConfigureAuthorizedCallersChangeset(reg, mcmsRegistry).Apply(*e, api.Config{
		Force: false,
		Updates: []api.ApplyInput{
			{
				ChainSelector: chainSelector,
				ContractType:  rmnops.ContractType,
				Version:       rmnops.Version,
				Update: api.CallerUpdate{
					AddedCallers:   []api.Caller{deployer.Bytes()},
					RemovedCallers: []api.Caller{common.HexToAddress("0x1234").Bytes()},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Nil(t, csOut.MCMSTimelockProposals, "no-op updates should produce no MCMS proposal")
}

// TestConfigureAuthorizedCallersChangeset_MultiTarget verifies that two ApplyInput entries
// on the same chain with different contract types are validated independently by the
// changeset (VerifyPreconditions), proving the grouping key includes ContractType.
func TestConfigureAuthorizedCallersChangeset_MultiTarget(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From

	e.DataStore = datastore.NewMemoryDataStore().Seal()
	mcmsRegistry := cs_core.GetRegistry()

	// Deploy a real RMN for the primary type.
	deployOut, err := changesets.DeployRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployRMNCfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.DeployRMNCfg{ChainSel: chainSelector, CurseAdmins: []common.Address{}},
	})
	require.NoError(t, err)
	rmnAddrs, err := deployOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)

	// Register a second adapter for a stub type so validation resolves both entries.
	const secondType = "AuthorizedCallersV2"
	reg := api.GetAuthorizedCallersRegistry()
	reg.RegisterAdapter("evm", secondType, rmnops.Version, adapters.NewRMNAuthorizedCallersAdapter())

	// Insert a stub datastore ref for the second contract type.
	secondRef := datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          secondType,
		Version:       rmnops.Version,
		Address:       common.HexToAddress("0x0000000000000000000000000000000000000001").Hex(),
	}
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(rmnAddrs[0]))
	require.NoError(t, ds.Addresses().Add(secondRef))
	e.DataStore = ds.Seal()

	cfg := api.Config{
		Force: true,
		Updates: []api.ApplyInput{
			{
				ChainSelector: chainSelector,
				ContractType:  rmnops.ContractType,
				Version:       rmnops.Version,
				Update:        api.CallerUpdate{AddedCallers: []api.Caller{deployer.Bytes()}},
			},
			{
				ChainSelector: chainSelector,
				ContractType:  secondType,
				Version:       rmnops.Version,
				Update:        api.CallerUpdate{AddedCallers: []api.Caller{deployer.Bytes()}},
			},
		},
	}

	// VerifyPreconditions must pass: both contract types have registered adapters.
	err = api.ConfigureAuthorizedCallersChangeset(reg, mcmsRegistry).VerifyPreconditions(*e, cfg)
	require.NoError(t, err, "both contract types should resolve to registered adapters")
}
