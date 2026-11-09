package changesets_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/changesets"
	rmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/registry_module_external_access_control"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/sequences"
	tar_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const registryModuleTestChainSel = uint64(5009297550715157269)

// deployTokenAdminRegistry deploys a TokenAdminRegistry owned by the deployer key and returns its address ref.
func deployTokenAdminRegistry(t *testing.T, e *deployment.Environment) datastore.AddressRef {
	t.Helper()

	chain := e.BlockChains.EVMChains()[registryModuleTestChainSel]
	tarRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle,
		tarops.Deploy,
		chain,
		contract_utils.DeployInput[tarops.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(tarops.ContractType, *tarops.Version),
			ChainSelector:  registryModuleTestChainSel,
		},
		nil,
	)
	require.NoError(t, err, "Failed to deploy TokenAdminRegistry")

	return tarRef
}

func findModuleRef(t *testing.T, refs []datastore.AddressRef) datastore.AddressRef {
	t.Helper()

	for _, ref := range refs {
		if ref.Type == datastore.ContractType(rmops.ContractType) {
			return ref
		}
	}
	require.Fail(t, "RegistryModuleExternalAccessControl not found in datastore")

	return datastore.AddressRef{}
}

func TestDeployRegistryModuleExternalAccessControl_Apply(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{registryModuleTestChainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chain := e.BlockChains.EVMChains()[registryModuleTestChainSel]
	tarRef := deployTokenAdminRegistry(t, e)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(tarRef))
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.DeployRegistryModuleExternalAccessControl(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployRegistryModuleExternalAccessControlCfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.DeployRegistryModuleExternalAccessControlCfg{ChainSel: registryModuleTestChainSel},
	})
	require.NoError(t, err, "Failed to apply DeployRegistryModuleExternalAccessControl changeset")

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err, "Failed to fetch addresses from datastore")

	moduleRef := findModuleRef(t, addrs)
	require.Equal(t, registryModuleTestChainSel, moduleRef.ChainSelector, "Chain selector should match")
	require.Equal(t, rmops.Version.String(), moduleRef.Version.String(), "Version should match")
	moduleAddress := common.HexToAddress(moduleRef.Address)
	require.NotEqual(t, common.Address{}, moduleAddress, "Module address should not be empty")

	// The module should be registered on the TokenAdminRegistry.
	tar, err := tar_bindings.NewTokenAdminRegistry(common.HexToAddress(tarRef.Address), chain.Client)
	require.NoError(t, err, "Failed to bind TokenAdminRegistry")
	isModule, err := tar.IsRegistryModule(&bind.CallOpts{Context: t.Context()}, moduleAddress)
	require.NoError(t, err, "Failed to call IsRegistryModule")
	require.True(t, isModule, "Module should be registered on the TokenAdminRegistry")
}

func TestDeployRegistryModuleExternalAccessControl_ApplyIsIdempotent(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{registryModuleTestChainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chain := e.BlockChains.EVMChains()[registryModuleTestChainSel]
	tarRef := deployTokenAdminRegistry(t, e)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(tarRef))
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	cfg := cs_core.WithMCMS[changesets.DeployRegistryModuleExternalAccessControlCfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.DeployRegistryModuleExternalAccessControlCfg{ChainSel: registryModuleTestChainSel},
	}

	out1, err := changesets.DeployRegistryModuleExternalAccessControl(mcmsRegistry).Apply(*e, cfg)
	require.NoError(t, err, "First apply failed")

	addrs1, err := out1.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	firstModuleRef := findModuleRef(t, addrs1)

	// Seed the environment with the first run's output plus the TAR so the second run treats both as existing.
	ds2 := datastore.NewMemoryDataStore()
	require.NoError(t, ds2.Addresses().Add(tarRef))
	for _, ref := range addrs1 {
		require.NoError(t, ds2.Addresses().Add(ref))
	}
	e.DataStore = ds2.Seal()

	out2, err := changesets.DeployRegistryModuleExternalAccessControl(mcmsRegistry).Apply(*e, cfg)
	require.NoError(t, err, "Second apply failed")

	addrs2, err := out2.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	secondModuleRef := findModuleRef(t, addrs2)

	require.Equal(t, firstModuleRef.Address, secondModuleRef.Address, "Module address should not change on re-run")

	// Still registered exactly once on the TAR.
	tar, err := tar_bindings.NewTokenAdminRegistry(common.HexToAddress(tarRef.Address), chain.Client)
	require.NoError(t, err, "Failed to bind TokenAdminRegistry")
	isModule, err := tar.IsRegistryModule(&bind.CallOpts{Context: t.Context()}, common.HexToAddress(secondModuleRef.Address))
	require.NoError(t, err, "Failed to call IsRegistryModule")
	require.True(t, isModule, "Module should remain registered on the TokenAdminRegistry")
}

func TestDeployRegistryModuleExternalAccessControl_PreparesMCMSProposalWhenDeployerNotOwner(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t, []uint64{registryModuleTestChainSel}, onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chain := e.BlockChains.EVMChains()[registryModuleTestChainSel]
	require.NotEmpty(t, chain.Users, "Expected an additional simulated account")
	tarRef := deployTokenAdminRegistry(t, e)

	// Hand the TAR to a different account so the deployer key is no longer allowed to add registry modules. This mimics
	// a live chain where the TAR is owned by the timelock. TokenAdminRegistry uses two-step ownership, so the new owner
	// must accept.
	tar, err := tar_bindings.NewTokenAdminRegistry(common.HexToAddress(tarRef.Address), chain.Client)
	require.NoError(t, err, "Failed to bind TokenAdminRegistry")
	newOwner := chain.Users[0]
	transferTx, err := tar.TransferOwnership(chain.DeployerKey, newOwner.From)
	require.NoError(t, err, "Failed to transfer TokenAdminRegistry ownership")
	_, err = chain.Confirm(transferTx)
	require.NoError(t, err, "Failed to confirm ownership transfer")

	acceptTx, err := tar.AcceptOwnership(newOwner)
	require.NoError(t, err, "Failed to accept TokenAdminRegistry ownership")
	_, err = chain.Confirm(acceptTx)
	require.NoError(t, err, "Failed to confirm ownership acceptance")

	owner, err := tar.Owner(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to read TokenAdminRegistry owner")
	require.Equal(t, newOwner.From, owner, "TokenAdminRegistry should be owned by the new account")

	// Execute the sequence directly so the raw BatchOps can be inspected before the changeset wraps them in an MCMS
	// proposal.
	report, err := cldf_ops.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployRegistryModuleExternalAccessControl,
		chain,
		sequences.DeployRegistryModuleExternalAccessControlInput{
			ChainSelector:      registryModuleTestChainSel,
			TokenAdminRegistry: common.HexToAddress(tarRef.Address),
			ExistingAddresses:  []datastore.AddressRef{tarRef},
		},
	)
	require.NoError(t, err, "Failed to execute DeployRegistryModuleExternalAccessControl sequence")

	moduleRef := findModuleRef(t, report.Output.Addresses)
	moduleAddress := common.HexToAddress(moduleRef.Address)

	// The registration must be prepared as an MCMS transaction rather than executed.
	require.Len(t, report.Output.BatchOps, 1, "Expected a single batch operation with the registration write")
	require.Len(t, report.Output.BatchOps[0].Transactions, 1, "Expected a single prepared transaction")
	preparedTx := report.Output.BatchOps[0].Transactions[0]
	require.Equal(t, common.HexToAddress(tarRef.Address).Hex(), common.HexToAddress(preparedTx.To).Hex(), "Prepared transaction should target the TokenAdminRegistry")

	tarABI, err := tar_bindings.TokenAdminRegistryMetaData.GetAbi()
	require.NoError(t, err, "Failed to get TokenAdminRegistry ABI")
	expectedCalldata, err := tarABI.Pack("addRegistryModule", moduleAddress)
	require.NoError(t, err, "Failed to pack addRegistryModule calldata")
	require.Equal(t, expectedCalldata, []byte(preparedTx.Data), "Prepared transaction should call addRegistryModule with the module address")

	// The module is not registered on-chain yet because the write was only prepared.
	isModule, err := tar.IsRegistryModule(&bind.CallOpts{Context: t.Context()}, moduleAddress)
	require.NoError(t, err, "Failed to call IsRegistryModule")
	require.False(t, isModule, "Module should not be registered until the proposal executes")
}

func TestDeployRegistryModuleExternalAccessControl_VerifyPreconditions(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{registryModuleTestChainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	mcmsRegistry := cs_core.GetRegistry()
	changeset := changesets.DeployRegistryModuleExternalAccessControl(mcmsRegistry)
	cfg := cs_core.WithMCMS[changesets.DeployRegistryModuleExternalAccessControlCfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.DeployRegistryModuleExternalAccessControlCfg{ChainSel: registryModuleTestChainSel},
	}

	// Missing TokenAdminRegistry in the datastore should fail preconditions.
	err = changeset.VerifyPreconditions(*e, cfg)
	require.ErrorContains(t, err, "failed to find token admin registry address")

	// With the TokenAdminRegistry seeded, preconditions should pass.
	tarRef := deployTokenAdminRegistry(t, e)
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(tarRef))
	e.DataStore = ds.Seal()

	err = changeset.VerifyPreconditions(*e, cfg)
	require.NoError(t, err, "VerifyPreconditions should pass when the TokenAdminRegistry is present")
}
