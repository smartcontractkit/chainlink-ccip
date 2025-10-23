package v1_0_0

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	"github.com/stretchr/testify/require"

	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func TestDeployMCMS(t *testing.T) {
	t.Parallel()
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selector}),
	)
	require.NoError(t, err)
	evmChain := env.BlockChains.EVMChains()[selector]

	evmDeployer := &EVMDeployer{}
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.0.0")
	dReg.RegisterDeployer(chainsel.FamilyEVM, version, evmDeployer)
	cs := deployops.DeployMCMS(dReg)
	output, err := cs.Apply(*env, deployops.MCMSDeploymentConfig{
		AdapterVersion: version,
		Chains: map[uint64]deployops.MCMSDeploymentConfigPerChain{
			selector: {
				Canceller:        deployops.SingleGroupMCMSV2(),
				Bypasser:         deployops.SingleGroupMCMSV2(),
				Proposer:         deployops.SingleGroupMCMSV2(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain.DeployerKey.From,
			},
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	env.DataStore = output.DataStore.Seal()
	// filter addresses for the test chain selector
	proposerRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.ProposerManyChainMultisig)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, proposerRef, 1)
	require.NotEqual(t, common.Address{}, proposerRef[0].Address)

	bypasserRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.BypasserManyChainMultisig)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, bypasserRef, 1)
	require.NotEqual(t, common.Address{}, bypasserRef[0].Address)

	cancellerRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.CancellerManyChainMultisig)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, cancellerRef, 1)
	require.NotEqual(t, common.Address{}, cancellerRef[0].Address)

	timelockRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.RBACTimelock)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, timelockRef, 1)
	require.NotEqual(t, common.Address{}, timelockRef[0].Address)

	callProxyRef := env.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selector),
		datastore.AddressRefByType(datastore.ContractType(utils.CallProxy)),
		datastore.AddressRefByQualifier("test"),
	)
	require.Len(t, callProxyRef, 1)
	require.NotEqual(t, common.Address{}, callProxyRef[0].Address)

	// query timelock and check the role assignments
	timelockC, err := bindings.NewRBACTimelock(
		common.HexToAddress(timelockRef[0].Address),
		evmChain.Client)
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
