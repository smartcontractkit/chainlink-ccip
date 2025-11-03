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
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// TestTransferOwnership tests transferring ownership of deployed contracts via MCMS timelocks.
// It deploys MCMS contracts on two EVM chains, deploys routers, transfers ownership of routers from deployer key to the timelock,
// then transfers ownership from the first timelock to a second timelock.
func TestTransferOwnership(t *testing.T) {
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
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	deployMCMS := deploy.DeployMCMS(dReg)
	output, err := deployMCMS.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
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
	ds := output.DataStore

	// deploy another timelock so that later we can transfer ownership to it from first timelock
	output, err = deployMCMS.Apply(*env, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test1"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			selector2: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test1"),
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

	// deploy router on both chains, then transfer ownership to the timelock
	routerAddrs := make(map[uint64]common.Address)
	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		// mock wrapped native and rmnproxy address
		wNative := utils.RandomAddress()
		rmnProxy := utils.RandomAddress()
		deployRouterOp, err := cldf_ops.ExecuteOperation(env.OperationsBundle, routerops1_2.Deploy, evmChain, contract.DeployInput[routerops1_2.ConstructorArgs]{
			ChainSelector:  evmChain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(routerops1_2.ContractType, *semver.MustParse("1.2.0")),
			Args: routerops1_2.ConstructorArgs{
				WrappedNative: wNative,
				RMNProxy:      rmnProxy,
			},
		})
		require.NoError(t, err)
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			Type:          datastore.ContractType(routerops1_2.ContractType),
			Version:       semver.MustParse("1.2.0"),
			ChainSelector: sel,
			Address:       deployRouterOp.Output.Address,
		}))
		routerAddrs[sel] = common.HexToAddress(deployRouterOp.Output.Address)
	}
	env.DataStore = ds.Seal()
	timelockAddrs := make(map[uint64]string)
	newTimelockAddrs := make(map[uint64]string)
	for _, sel := range []uint64{selector1, selector2} {
		timelockRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.RBACTimelock),
			Qualifier:     "test",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		timelockAddrs[sel] = timelockRef.Address
		newTimelockRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(deploymentutils.RBACTimelock),
			Qualifier:     "test1",
			Version:       semver.MustParse("1.0.0"),
		}, sel, datastore_utils.FullRef)
		require.NoError(t, err)
		newTimelockAddrs[sel] = newTimelockRef.Address
	}
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs: []deploy.TransferOwnershipPerChainInput{
			{
				ChainSelector: selector1,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(routerops1_2.ContractType),
						Version: semver.MustParse("1.2.0"),
					},
				},
				ProposedOwner: timelockAddrs[selector1],
			},
			{
				ChainSelector: selector2,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(routerops1_2.ContractType),
						Version: semver.MustParse("1.2.0"),
					},
				},
				ProposedOwner: timelockAddrs[selector2],
			},
		},
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Transfer ownership test",
		},
	}
	// register chain adapter
	cr := deploy.GetTransferOwnershipRegistry()
	evmAdapter := &adapters.EVMTransferOwnershipAdapter{}
	cr.RegisterAdapter(chainsel.FamilyEVM, transferOwnershipInput.AdapterVersion, evmAdapter)
	mcmsRegistry := changesets.GetRegistry()
	evmMCMSReader := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, evmMCMSReader)
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// now check the owner of the routers is the timelock
	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		r, err := router.NewRouter(routerAddrs[sel], evmChain.Client)
		require.NoError(t, err)
		owner, err := r.Owner(&bind.CallOpts{
			Context: t.Context(),
		})
		require.NoError(t, err)
		require.Equal(t, common.HexToAddress(timelockAddrs[sel]), owner, "owner mismatch on chain %d", sel)
		t.Logf("Ownership of router on chain %d successfully transferred to timelock %s", sel, timelockAddrs[sel])
	}

	// now transfer ownership from first timelock to second timelock
	// the mcms input should denote the first timelock address
	transferOwnershipInput = deploy.TransferOwnershipInput{
		ChainInputs: []deploy.TransferOwnershipPerChainInput{
			{
				ChainSelector: selector1,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(routerops1_2.ContractType),
						Version: semver.MustParse("1.2.0"),
					},
				},
				ProposedOwner: newTimelockAddrs[selector1],
			},
			{
				ChainSelector: selector2,
				ContractRef: []datastore.AddressRef{
					{
						Type:    datastore.ContractType(routerops1_2.ContractType),
						Version: semver.MustParse("1.2.0"),
					},
				},
				ProposedOwner: newTimelockAddrs[selector2],
			},
		},
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Transfer ownership test",
		},
	}
	transferOwnershipChangeset = deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// now accept ownership from the new timelock, the mcms input should denote new timelock address
	transferOwnershipInput.MCMS = mcms.Input{
		OverridePreviousRoot: false,
		ValidUntil:           3759765795,
		TimelockDelay:        mcms_types.MustParseDuration("0s"),
		TimelockAction:       mcms_types.TimelockActionSchedule,
		Qualifier:            "test1",
		Description:          "Transfer ownership test",
	}
	acceptOwnershipChangeset := deploy.AcceptOwnershipChangeset(cr, mcmsRegistry)
	output, err = acceptOwnershipChangeset.Apply(*env, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// now check the owner of the routers is the new timelock
	for _, sel := range []uint64{selector1, selector2} {
		evmChain := env.BlockChains.EVMChains()[sel]
		r, err := router.NewRouter(routerAddrs[sel], evmChain.Client)
		require.NoError(t, err)
		owner, err := r.Owner(&bind.CallOpts{
			Context: t.Context(),
		})
		require.NoError(t, err)
		require.Equal(t, common.HexToAddress(newTimelockAddrs[sel]), owner, "owner mismatch on chain %d", sel)
		t.Logf("Ownership of router on chain %d successfully transferred to new timelock %s", sel, newTimelockAddrs[sel])
	}
}
