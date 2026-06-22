package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func applyDeployMCMSAndAcceptOwnership(
	t *testing.T,
	env *cldf_deployment.Environment,
	cs cldf_deployment.ChangeSetV2[deployops.MCMSDeploymentConfig],
	cfg deployops.MCMSDeploymentConfig,
) cldf_deployment.ChangesetOutput {
	t.Helper()
	output, err := cs.Apply(*env, cfg)
	require.NoError(t, err)
	if len(output.MCMSTimelockProposals) > 0 {
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	}
	return output
}

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

	evmDeployer := &adapters.EVMDeployer{}
	dReg := deployops.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deployops.MCMSVersion, evmDeployer)
	cs := deployops.DeployMCMS(dReg, nil)

	output := applyDeployMCMSAndAcceptOwnership(t, env, cs, deployops.MCMSDeploymentConfig{
		AdapterVersion: deployops.MCMSVersion,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(utils.CLLQualifier),
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String(utils.CLLQualifier),
			},
		},
	})

	require.Greater(t, len(output.Reports), 0)
	env.DataStore = output.DataStore.Seal()
	// filter addresses for the test chain selector
	proposerRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.ProposerManyChainMultisig)),
		datastore.AddressRefByQualifier(utils.CLLQualifier),
	)
	require.Len(t, proposerRef, 1)
	require.NotEqual(t, common.Address{}, proposerRef[0].Address)

	bypasserRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.BypasserManyChainMultisig)),
		datastore.AddressRefByQualifier(utils.CLLQualifier),
	)
	require.Len(t, bypasserRef, 1)
	require.NotEqual(t, common.Address{}, bypasserRef[0].Address)

	cancellerRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.CancellerManyChainMultisig)),
		datastore.AddressRefByQualifier(utils.CLLQualifier),
	)
	require.Len(t, cancellerRef, 1)
	require.NotEqual(t, common.Address{}, cancellerRef[0].Address)

	timelockRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
		datastore.AddressRefByQualifier(utils.CLLQualifier),
	)
	require.Len(t, timelockRef, 1)
	require.NotEqual(t, common.Address{}, timelockRef[0].Address)

	callProxyRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector1),
		datastore.AddressRefByType(datastore.ContractType(utils.CallProxy)),
		datastore.AddressRefByQualifier(utils.CLLQualifier),
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

	// CLLCCIP timelock is self-governed.
	adminRoleAdmin, err := timelockC.GetRoleAdmin(&bind.CallOpts{Context: t.Context()}, ops.ADMIN_ROLE.ID)
	require.NoError(t, err)
	timelockAddr := common.HexToAddress(timelockRef[0].Address)
	selfAdmin, err := timelockC.HasRole(&bind.CallOpts{Context: t.Context()}, adminRoleAdmin, timelockAddr)
	require.NoError(t, err)
	require.True(t, selfAdmin, "CLLCCIP timelock should be self-governed")

	for _, ref := range [][]datastore.AddressRef{proposerRef, bypasserRef, cancellerRef} {
		owner, _, err := seq.LoadOwnableContract(common.HexToAddress(ref[0].Address), evmChain1.Client)
		require.NoError(t, err)
		require.Equal(t, timelockAddr, owner, "MCM %s should be owned by CLLCCIP timelock", ref[0].Type)
	}
}

// TestDeployMCMS_TimelockAdminRoleTransfer verifies the admin-role logic inside DeployMCMS:
//   - non-CLLCCIP qualifier with no CLLCCIP deployed → changeset fails with an error
//   - CLLCCIP qualifier (bootstrap) → timelock itself holds ADMIN_ROLE (self-sovereign)
func TestDeployMCMS_TimelockAdminRoleTransfer(t *testing.T) {
	t.Parallel()
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector}),
	)
	require.NoError(t, err)
	env.Logger = logger.Test(t)
	evmChain := env.BlockChains.EVMChains()[selector]

	evmDeployer := &adapters.EVMDeployer{}
	dReg := deployops.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deployops.MCMSVersion, evmDeployer)
	cs := deployops.DeployMCMS(dReg, nil)

	t.Run("non-CLLCCIP with no CLLCCIP deployed fails", func(t *testing.T) {
		_, err := cs.Apply(*env, deployops.MCMSDeploymentConfig{
			AdapterVersion: deployops.MCMSVersion,
			Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
				selector: {
					Canceller:        testhelpers.SingleGroupMCMS(),
					Bypasser:         testhelpers.SingleGroupMCMS(),
					Proposer:         testhelpers.SingleGroupMCMS(),
					TimelockMinDelay: big.NewInt(0),
					Qualifier:        ptr.String("no-cllccip"),
				},
			},
		})
		require.Error(t, err)
		require.ErrorContains(t, err, "CLLCCIP RBACTimelock must be deployed first")
	})

	t.Run("CLLCCIP bootstrap is self-governed", func(t *testing.T) {
		output := applyDeployMCMSAndAcceptOwnership(t, env, cs, deployops.MCMSDeploymentConfig{
			AdapterVersion: deployops.MCMSVersion,
			Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
				selector: {
					Canceller:        testhelpers.SingleGroupMCMS(),
					Bypasser:         testhelpers.SingleGroupMCMS(),
					Proposer:         testhelpers.SingleGroupMCMS(),
					TimelockMinDelay: big.NewInt(0),
					Qualifier:        ptr.String(utils.CLLQualifier),
				},
			},
		})
		env.DataStore = output.DataStore.Seal()

		timelockRef := env.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(selector),
			datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
			datastore.AddressRefByQualifier(utils.CLLQualifier),
		)
		require.Len(t, timelockRef, 1)
		timelockAddr := common.HexToAddress(timelockRef[0].Address)

		timelockC, err := bindings.NewRBACTimelock(timelockAddr, evmChain.Client)
		require.NoError(t, err)

		adminRoleAdmin, err := timelockC.GetRoleAdmin(&bind.CallOpts{Context: t.Context()}, ops.ADMIN_ROLE.ID)
		require.NoError(t, err)

		timelockIsSelfAdmin, err := timelockC.HasRole(&bind.CallOpts{Context: t.Context()}, adminRoleAdmin, timelockAddr)
		require.NoError(t, err)
		require.True(t, timelockIsSelfAdmin, "CLLCCIP timelock should be self-governed")

		deployerStillAdmin, err := timelockC.HasRole(&bind.CallOpts{Context: t.Context()}, adminRoleAdmin, evmChain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, deployerStillAdmin, "deployer should never hold ADMIN_ROLE after deployment")
	})
}

// TestDeployMCMS_DefaultsToExistingCLLCCIPTimelock verifies that when a new non-CLLCCIP
// MCMS is deployed, the deployer automatically uses the
// existing CLLCCIP RBACTimelock as the admin (so it is governed by CLL from the start).
func TestDeployMCMS_DefaultsToExistingCLLCCIPTimelock(t *testing.T) {
	t.Parallel()
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector}),
	)
	require.NoError(t, err)
	env.Logger = logger.Test(t)
	evmChain := env.BlockChains.EVMChains()[selector]

	evmDeployer := &adapters.EVMDeployer{}
	dReg := deployops.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deployops.MCMSVersion, evmDeployer)
	cs := deployops.DeployMCMS(dReg, nil)

	// Step 1: deploy the CLLCCIP MCMS instance
	cllOutput := applyDeployMCMSAndAcceptOwnership(t, env, cs, deployops.MCMSDeploymentConfig{
		AdapterVersion: deployops.MCMSVersion,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new(utils.CLLQualifier),
			},
		},
	})
	require.NoError(t, cllOutput.DataStore.Merge(env.DataStore))
	env.DataStore = cllOutput.DataStore.Seal()

	// Resolve the CLLCCIP timelock address for later assertion.
	cllTimelockRefs := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
		datastore.AddressRefByQualifier(utils.CLLQualifier),
	)
	require.Len(t, cllTimelockRefs, 1)
	cllTimelockAddr := common.HexToAddress(cllTimelockRefs[0].Address)

	// Step 2: deploy a second (non-CLLCCIP) MCMS instance.
	// The deployer should automatically pick the CLLCCIP timelock as the admin.
	rmnOutput := applyDeployMCMSAndAcceptOwnership(t, env, cs, deployops.MCMSDeploymentConfig{
		AdapterVersion: deployops.MCMSVersion,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new(utils.RMNTimelockQualifier),
			},
		},
	})
	require.NoError(t, rmnOutput.DataStore.Merge(env.DataStore))
	env.DataStore = rmnOutput.DataStore.Seal()

	rmnTimelockRefs := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
		datastore.AddressRefByQualifier(utils.RMNTimelockQualifier),
	)
	require.Len(t, rmnTimelockRefs, 1)
	rmnTimelockAddr := common.HexToAddress(rmnTimelockRefs[0].Address)

	rmnTimelock, err := bindings.NewRBACTimelock(rmnTimelockAddr, evmChain.Client)
	require.NoError(t, err)

	adminRoleAdmin, err := rmnTimelock.GetRoleAdmin(&bind.CallOpts{Context: t.Context()}, ops.ADMIN_ROLE.ID)
	require.NoError(t, err)

	// The CLLCCIP timelock should be the admin of the new RMNMCMS timelock.
	cllIsAdmin, err := rmnTimelock.HasRole(&bind.CallOpts{Context: t.Context()}, adminRoleAdmin, cllTimelockAddr)
	require.NoError(t, err)
	require.True(t, cllIsAdmin, "CLLCCIP timelock should be admin of the RMNMCMS timelock")

	// The deployer EOA must have renounced ADMIN_ROLE.
	deployerIsAdmin, err := rmnTimelock.HasRole(&bind.CallOpts{Context: t.Context()}, adminRoleAdmin, evmChain.DeployerKey.From)
	require.NoError(t, err)
	require.False(t, deployerIsAdmin, "deployer should have renounced ADMIN_ROLE after transfer to CLLCCIP timelock")
}

func TestUpdateMCMSConfig(t *testing.T) {
	t.Parallel()
	selector1 := chainsel.TEST_90000001.Selector
	selector2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector1, selector2}),
	)
	require.NoError(t, err)
	env.Logger = logger.Test(t)

	evmDeployer := &adapters.EVMDeployer{}
	dReg := deployops.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deployops.MCMSVersion, evmDeployer)

	// deploy one set of timelock and MCMS contracts on each chain
	deployMCMS := deployops.DeployMCMS(dReg, nil)
	output := applyDeployMCMSAndAcceptOwnership(t, env, deployMCMS, deployops.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new("CLLCCIP"),
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new("CLLCCIP"),
			},
		},
	})
	require.Greater(t, len(output.Reports), 0)
	env.DataStore = output.DataStore.Seal()

	// get recently deployed MCMS addresses
	mcmsRefs := make(map[uint64][]datastore.AddressRef)
	partialMCMSRefs := make(map[uint64][]datastore.AddressRef)
	for _, sel := range []uint64{selector1, selector2} {
		cancellerRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.CancellerManyChainMultisig),
			Qualifier:     "CLLCCIP",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		bypasserRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.BypasserManyChainMultisig),
			Qualifier:     "CLLCCIP",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		proposerRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.ProposerManyChainMultisig),
			Qualifier:     "CLLCCIP",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		mcmsRefs[sel] = append(mcmsRefs[sel], cancellerRef, bypasserRef, proposerRef)
		for _, ref := range mcmsRefs[sel] {
			partialMCMSRefs[sel] = append(partialMCMSRefs[sel], datastore.AddressRef{
				Type:      ref.Type,
				Qualifier: ref.Qualifier,
				Version:   ref.Version,
			})
		}
	}

	// check that deployed config is correct
	for _, sel := range []uint64{selector1, selector2} {
		for _, ref := range mcmsRefs[sel] {
			evmChain := env.BlockChains.EVMChains()[sel]
			mcmsContract, err := bindings.NewManyChainMultiSig(common.HexToAddress(ref.Address), evmChain.Client)
			require.NoError(t, err)

			// binding is done, now check config
			config, err := mcmsContract.GetConfig(&bind.CallOpts{
				Context: t.Context(),
			})
			require.NoError(t, err)

			numOfSigners := len(config.Signers)
			require.Equal(t, numOfSigners, len(testhelpers.SingleGroupMCMS().Signers)) // should be 1
		}
	}

	// update the config for each MCMS contract
	updateMcmsConfigMCMS := deployops.UpdateMCMSConfig(dReg, nil)
	output, err = updateMcmsConfigMCMS.Apply(*env, deployops.UpdateMCMSConfigInput{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deployops.UpdateMCMSConfigInputPerChain{
			selector1: {
				MCMConfig:    testhelpers.SingleGroupMCMSTwoSigners(),
				MCMContracts: partialMCMSRefs[selector1],
			},
			selector2: {
				MCMConfig:    testhelpers.SingleGroupMCMSTwoSigners(),
				MCMContracts: partialMCMSRefs[selector2],
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "CLLCCIP",
			Description:          "update mcms config test",
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.NotEmpty(t, output.MCMSTimelockProposals)
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// check that MCMS configs are updated correctly
	for _, sel := range []uint64{selector1, selector2} {
		for _, ref := range mcmsRefs[sel] {
			evmChain := env.BlockChains.EVMChains()[sel]
			mcmsContract, err := bindings.NewManyChainMultiSig(common.HexToAddress(ref.Address), evmChain.Client)
			require.NoError(t, err)
			config, err := mcmsContract.GetConfig(&bind.CallOpts{
				Context: t.Context(),
			})
			require.NoError(t, err)

			numOfSigners := len(config.Signers)
			require.Equal(t, numOfSigners, len(testhelpers.SingleGroupMCMSTwoSigners().Signers)) // should be 2
		}
	}
}

// TestDeployMCMS_CLLCCIPAutoAdminMultiChain verifies that on multiple chains, when the CLLCCIP
// instance is deployed first and a second MCMS set is deployed afterwards, DeployMCMS
// automatically sets the CLLCCIP timelock as the admin of the new timelock.
func TestDeployMCMS_CLLCCIPAutoAdminMultiChain(t *testing.T) {
	t.Parallel()
	selector1 := chainsel.TEST_90000001.Selector
	selector2 := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector1, selector2}),
	)
	require.NoError(t, err)
	env.Logger = logger.Test(t)

	evmDeployer := &adapters.EVMDeployer{}
	dReg := deployops.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deployops.MCMSVersion, evmDeployer)

	deployMCMS := deployops.DeployMCMS(dReg, nil)

	// Step 1: deploy CLLCCIP on both chains (bootstrap — each self-governed).
	cllOutput := applyDeployMCMSAndAcceptOwnership(t, env, deployMCMS, deployops.MCMSDeploymentConfig{
		AdapterVersion: deployops.MCMSVersion,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new(utils.CLLQualifier),
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new(utils.CLLQualifier),
			},
		},
	})
	require.Greater(t, len(cllOutput.Reports), 0)
	require.NoError(t, cllOutput.DataStore.Merge(env.DataStore))
	env.DataStore = cllOutput.DataStore.Seal()

	// Collect CLLCCIP timelock addresses for assertion.
	cllTimelockAddrs := make(map[uint64]common.Address)
	for _, sel := range []uint64{selector1, selector2} {
		refs := env.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(sel),
			datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
			datastore.AddressRefByQualifier(utils.CLLQualifier),
		)
		require.Len(t, refs, 1)
		cllTimelockAddrs[sel] = common.HexToAddress(refs[0].Address)
	}

	// Step 2: deploy a second MCMS set (testQual) with CLLCCIP already in datastore.
	// DeployMCMS should automatically use CLLCCIP timelock as admin.
	testOutput := applyDeployMCMSAndAcceptOwnership(t, env, deployMCMS, deployops.MCMSDeploymentConfig{
		AdapterVersion: deployops.MCMSVersion,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new("testQual"),
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        new("testQual"),
			},
		},
	})
	require.Greater(t, len(testOutput.Reports), 0)
	require.NoError(t, testOutput.DataStore.Merge(env.DataStore))
	env.DataStore = testOutput.DataStore.Seal()

	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]

		testQualRefs := env.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(sel),
			datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
			datastore.AddressRefByQualifier("testQual"),
		)
		require.Len(t, testQualRefs, 1)
		testQualTimelockAddr := common.HexToAddress(testQualRefs[0].Address)

		timelock, err := bindings.NewRBACTimelock(testQualTimelockAddr, evmChain.Client)
		require.NoError(t, err)
		roleAdmin, err := timelock.GetRoleAdmin(&bind.CallOpts{Context: t.Context()}, ops.ADMIN_ROLE.ID)
		require.NoError(t, err)

		// CLLCCIP timelock must be admin of testQual timelock.
		cllIsAdmin, err := timelock.HasRole(&bind.CallOpts{Context: t.Context()}, roleAdmin, cllTimelockAddrs[sel])
		require.NoError(t, err)
		require.True(t, cllIsAdmin, "CLLCCIP timelock should be admin of testQual timelock on chain %d", sel)
		t.Logf("CLLCCIP timelock %s is admin of testQual timelock %s on chain %d", cllTimelockAddrs[sel], testQualTimelockAddr, sel)

		// Deployer must NOT be admin.
		deployerIsAdmin, err := timelock.HasRole(&bind.CallOpts{Context: t.Context()}, roleAdmin, evmChain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, deployerIsAdmin, "deployer should not hold ADMIN_ROLE on chain %d", sel)
	}
}
