package adapters_test

import (
	"math/big"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	v1_0 "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

func TestDeployMCMS(t *testing.T) {
	t.Parallel()
	selector1 := chainsel.TEST_90000001.Selector
	selector2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector1, selector2}),
	)
	require.NoError(t, err)
	env.Logger = logger.Test(t)
	evmChain1 := env.BlockChains.EVMChains()[selector1]
	evmChain2 := env.BlockChains.EVMChains()[selector2]

	evmDeployer := &adapters.EVMDeployer{}
	dReg := v1_0.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, v1_0.MCMSVersion, evmDeployer)
	cs := v1_0.DeployMCMS(dReg)
	output, err := cs.Apply(*env, v1_0.MCMSDeploymentConfig{
		AdapterVersion: v1_0.MCMSVersion,
		Chains: map[uint64]v1_0.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain2.DeployerKey.From,
			},
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	env.DataStore = output.DataStore.Seal()
	// filter addresses for the test chain selector
	proposerRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.ProposerManyChainMultisig)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, proposerRef, 1)
	require.NotEqual(t, common.Address{}, proposerRef[0].Address)

	bypasserRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.BypasserManyChainMultisig)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, bypasserRef, 1)
	require.NotEqual(t, common.Address{}, bypasserRef[0].Address)

	cancellerRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.CancellerManyChainMultisig)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, cancellerRef, 1)
	require.NotEqual(t, common.Address{}, cancellerRef[0].Address)

	timelockRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, timelockRef, 1)
	require.NotEqual(t, common.Address{}, timelockRef[0].Address)

	callProxyRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.CallProxy)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, callProxyRef, 1)
	require.NotEqual(t, common.Address{}, callProxyRef[0].Address)

	// query timelock and check the role assignments
	timelockC, err := bindings.NewRBACTimelock(
		common.HexToAddress(timelockRef[0].Address),
		evmChain1.Client)
	require.NoError(t, err)

	pRole, err := timelockC.PROPOSERROLE(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	hasRole, err := timelockC.HasRole(&bind.CallOpts{Context: t.Context()}, pRole, common.HexToAddress(proposerRef[0].Address))
	require.NoError(t, err)
	require.True(t, hasRole, "proposer MCM should have admin role for PROPOSER_ROLE")

	eRole, err := timelockC.EXECUTORROLE(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	hasRole, err = timelockC.HasRole(&bind.CallOpts{Context: t.Context()}, eRole, common.HexToAddress(callProxyRef[0].Address))
	require.NoError(t, err)
	require.True(t, hasRole, "Call Proxy should have admin role for EXECUTOR_ROLE")
}

func TestGrantAdminRoleToTimelock(t *testing.T) {
	t.Parallel()
	selector1 := chainsel.TEST_90000001.Selector
	selector2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector1, selector2}),
	)
	require.NoError(t, err)
	env.Logger = logger.Test(t)
	evmChain1 := env.BlockChains.EVMChains()[selector1]
	evmChain2 := env.BlockChains.EVMChains()[selector2]

	evmDeployer := &adapters.EVMDeployer{}
	dReg := deployops.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deployops.MCMSVersion, evmDeployer)

	// deploy two timelocks on each chain so we can set one as the admin of the other
	deployMCMS := deployops.DeployMCMS(dReg, nil)
	output, err := deployMCMS.Apply(*env, deployops.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("testQual"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("testQual"),
				TimelockAdmin:    evmChain2.DeployerKey.From,
			},
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	ds := output.DataStore

	output, err = deployMCMS.Apply(*env, deployops.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("CLLCCIP"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("CLLCCIP"),
				TimelockAdmin:    evmChain2.DeployerKey.From,
			},
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	addresses, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	for _, addr := range addresses {
		t.Logf("Adding address %s of type %s on chain %d to datastore", addr.Address, addr.Type, addr.ChainSelector)
		require.NoError(t, ds.Addresses().Add(addr))
	}
	env.DataStore = ds.Seal()

	// get recently deployed timelock addresses
	timelockAddrs := make(map[uint64]string)
	newAdminTimelockAddrs := make(map[uint64]string)
	for _, sel := range []uint64{selector1, selector2} {
		timelockRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.RBACTimelock),
			Qualifier:     "testQual",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		timelockAddrs[sel] = timelockRef.Address
		newAdminTimelockRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.RBACTimelock),
			Qualifier:     "CLLCCIP",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		newAdminTimelockAddrs[sel] = newAdminTimelockRef.Address
	}

	// check that admin of timelock is deployer eoa
	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		timelock, err := bindings.NewRBACTimelock(common.HexToAddress(timelockAddrs[sel]), evmChain.Client)
		require.NoError(t, err)
		roleAdmin, err := timelock.GetRoleAdmin(&bind.CallOpts{
			Context: t.Context(),
		}, ops.ADMIN_ROLE.ID)
		require.NoError(t, err)

		callerIsAdmin, err := timelock.HasRole(&bind.CallOpts{
			Context: t.Context(),
		}, roleAdmin, evmChain.DeployerKey.From)
		require.NoError(t, err)
		require.True(t, callerIsAdmin)
	}

	// grant admin role to timelock
	grantAdminRoleMCMS := deployops.GrantAdminRoleToTimelock(dReg, nil)
	output, err = grantAdminRoleMCMS.Apply(*env, deployops.GrantAdminRoleToTimelockConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deployops.GrantAdminRoleToTimelockConfigPerChain{
			selector1: {
				TimelockToTransferRef: datastore.AddressRef{
					Type:      datastore.ContractType(deploymentutils.RBACTimelock),
					Version:   semver.MustParse("1.0.0"),
					Qualifier: "testQual",
				},
				NewAdminTimelockRef: datastore.AddressRef{
					Type:      datastore.ContractType(deploymentutils.RBACTimelock),
					Version:   semver.MustParse("1.0.0"),
					Qualifier: "CLLCCIP",
				},
			},
			selector2: {
				TimelockToTransferRef: datastore.AddressRef{
					Type:      datastore.ContractType(deploymentutils.RBACTimelock),
					Version:   semver.MustParse("1.0.0"),
					Qualifier: "testQual",
				},
				NewAdminTimelockRef: datastore.AddressRef{
					Type:      datastore.ContractType(deploymentutils.RBACTimelock),
					Version:   semver.MustParse("1.0.0"),
					Qualifier: "CLLCCIP",
				},
			},
		},
	})

	// check that "CLLCCIP" timelocks have admin role
	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		timelock, err := bindings.NewRBACTimelock(common.HexToAddress(timelockAddrs[sel]), evmChain.Client)
		require.NoError(t, err)
		roleAdmin, err := timelock.GetRoleAdmin(&bind.CallOpts{
			Context: t.Context(),
		}, ops.ADMIN_ROLE.ID)
		require.NoError(t, err)

		timelockIsAdmin, err := timelock.HasRole(&bind.CallOpts{
			Context: t.Context(),
		}, roleAdmin, common.HexToAddress(newAdminTimelockAddrs[sel]))
		require.NoError(t, err)
		require.True(t, timelockIsAdmin, "timelock does not have admin role on chain %d", sel)
		t.Logf("Timelock with address %s on chain %d was successfully granted admin role of timelock with address %s", newAdminTimelockAddrs[sel], sel, timelockAddrs[sel])
	}

	// check that deployer eoa no longer has admin role
	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		timelock, err := bindings.NewRBACTimelock(common.HexToAddress(timelockAddrs[sel]), evmChain.Client)
		require.NoError(t, err)
		roleAdmin, err := timelock.GetRoleAdmin(&bind.CallOpts{
			Context: t.Context(),
		}, ops.ADMIN_ROLE.ID)
		require.NoError(t, err)

		timelockIsAdmin, err := timelock.HasRole(&bind.CallOpts{
			Context: t.Context(),
		}, roleAdmin, evmChain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, timelockIsAdmin, "deployer eoa still has admin role on chain %d", sel)
		t.Logf("The deployer EOA on chain %d has successfully renounced admin role of timelock with address %s", sel, timelockAddrs[sel])
	}
}
