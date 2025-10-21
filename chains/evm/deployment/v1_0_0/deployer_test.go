package v1_0_0

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	mcmstypes "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_0"
)

var (
	// TestXXXMCMSSigner is a throwaway private key used for signing MCMS proposals.
	// in tests.
	TestXXXMCMSSigner *ecdsa.PrivateKey
)

func init() {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	TestXXXMCMSSigner = key
}

func SingleGroupTimelockConfigV2(t *testing.T) v1_0.MCMSDeploymentConfigPerChain {
	return v1_0.MCMSDeploymentConfigPerChain{
		Canceller:        SingleGroupMCMSV2(t),
		Bypasser:         SingleGroupMCMSV2(t),
		Proposer:         SingleGroupMCMSV2(t),
		TimelockMinDelay: big.NewInt(0),
	}
}

func SingleGroupMCMSV2(t *testing.T) mcmstypes.Config {
	publicKey := TestXXXMCMSSigner.Public().(*ecdsa.PublicKey)
	// Convert the public key to an Ethereum address
	address := crypto.PubkeyToAddress(*publicKey)
	c, err := mcmstypes.NewConfig(1, []common.Address{address}, []mcmstypes.Config{})
	require.NoError(t, err)
	return c
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
	evmChain2 := env.BlockChains.EVMChains()[selector2]

	evmDeployer := &EVMDeployer{}
	dReg := v1_0.NewDeployerRegistry()
	dReg.RegisterDeployer(chainsel.FamilyEVM, v1_0.MCMSVersion, evmDeployer)
	cs := v1_0.DeployMCMS(dReg)
	output, err := cs.Apply(*env, v1_0.MCMSDeploymentConfig{
		Chains: map[uint64]v1_0.MCMSDeploymentConfigPerChain{
			selector1: {
				Canceller:        SingleGroupMCMSV2(t),
				Bypasser:         SingleGroupMCMSV2(t),
				Proposer:         SingleGroupMCMSV2(t),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain1.DeployerKey.From,
			},
			selector2: {
				Canceller:        SingleGroupMCMSV2(t),
				Bypasser:         SingleGroupMCMSV2(t),
				Proposer:         SingleGroupMCMSV2(t),
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
