package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	rmn_proxy_bind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestActivateRMN_Apply(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle

	legacyARM := deployer
	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: legacyARM},
	}, nil)
	require.NoError(t, err)

	ultraFastTimelockAddr, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)
	rmnTimelockAddr, rmnMCMSAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.RMNTimelockQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range rmnMCMSAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{
			Qualifier:      common_utils.RMNTimelockQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			TimelockDelay:  mcms_types.MustParseDuration("0s"),
			ValidUntil:     3759765795,
		},
		Cfg: changesets.ActivateRMNCfg{
			ChainSels: []uint64{chainSelector},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, out.MCMSTimelockProposals, "accept ownership on RMNMCMS timelock should be proposed via MCMS")
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	rmnAddrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, rmnAddrs, 1)
	require.Equal(t, datastore.ContractType(rmnops.ContractType), rmnAddrs[0].Type)

	newRMN := common.HexToAddress(rmnAddrs[0].Address)

	rmnC, err := rmnops.NewRMNContract(newRMN, chain.Client)
	require.NoError(t, err)
	callers, err := rmnC.GetAllAuthorizedCallers(nil)
	require.NoError(t, err)
	require.Contains(t, callers, ultraFastTimelockAddr, "Ultra Fast Curse timelock must be a curse admin")

	owner, err := rmnC.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, rmnTimelockAddr, owner, "RMN should be owned by RMNMCMS timelock after proposal execution")

	proxyAddr := common.HexToAddress(proxyRef.Address)
	proxyC, err := rmn_proxy_bind.NewRMNProxy(proxyAddr, chain.Client)
	require.NoError(t, err)
	proxyOwner, err := proxyC.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, deployer, proxyOwner, "ARMProxy should be deployer-owned in this test")

	armReport, err := operations.ExecuteOperation(b, rmn_proxy.GetRMN, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       proxyAddr,
	})
	require.NoError(t, err)
	require.Equal(t, newRMN, armReport.Output, "ARMProxy must point at the new RMN implementation")
	require.NotEqual(t, legacyARM, armReport.Output)
}

func TestActivateRMN_MissingRMNMCMS(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSelector]
	deployer := chain.DeployerKey.From
	b := e.OperationsBundle

	proxyRef, err := contract.MaybeDeployContract(b, rmn_proxy.Deploy, chain, contract.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: deployer},
	}, nil)
	require.NoError(t, err)

	_, ultraFastAddrs := deployMCMSInstanceForTest(t, b, chain, deployer, common_utils.UltraFastCurseMCMSQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(proxyRef))
	for _, ref := range ultraFastAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	_, err = changesets.ActivateRMN(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.ActivateRMNCfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.ActivateRMNCfg{ChainSels: []uint64{chainSelector}},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, common_utils.RMNTimelockQualifier)
}

func deployMCMSInstanceForTest(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	deployer common.Address,
	qualifier string,
) (timelockAddr common.Address, addresses []datastore.AddressRef) {
	t.Helper()
	mcmsCfg := testhelpers.SingleGroupMCMS()
	qualifierPtr := &qualifier

	proposerReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.ProposerManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	require.NotEmpty(t, proposerReport.Output.Addresses)
	addresses = append(addresses, proposerReport.Output.Addresses...)

	bypasserReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.BypasserManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	addresses = append(addresses, bypasserReport.Output.Addresses...)

	cancellerReport, err := operations.ExecuteSequence(b, mcms_seq.SeqDeployMCMWithConfig, chain, mcms_seq.SeqMCMSDeploymentCfg{
		ChainSelector: chain.Selector,
		ContractType:  common_utils.CancellerManyChainMultisig,
		MCMConfig:     &mcmsCfg,
		Qualifier:     qualifierPtr,
	})
	require.NoError(t, err)
	addresses = append(addresses, cancellerReport.Output.Addresses...)

	timelockRef, err := contract.MaybeDeployContract(b, mcms_ops.OpDeployTimelock, chain, contract.DeployInput[mcms_ops.OpDeployTimelockInput]{
		ChainSelector:  chain.Selector,
		Qualifier:      qualifierPtr,
		TypeAndVersion: deployment.NewTypeAndVersion(common_utils.RBACTimelock, *mcms_ops.MCMSVersion),
		Args: mcms_ops.OpDeployTimelockInput{
			TimelockMinDelay: big.NewInt(0),
			Proposers:        []common.Address{common.HexToAddress(proposerReport.Output.Addresses[0].Address)},
			Bypassers:        []common.Address{common.HexToAddress(bypasserReport.Output.Addresses[0].Address)},
			Cancellers:       []common.Address{common.HexToAddress(cancellerReport.Output.Addresses[0].Address)},
			Admin:            deployer,
			Executors:        []common.Address{},
		},
	}, nil)
	require.NoError(t, err)
	timelockAddr = common.HexToAddress(timelockRef.Address)
	addresses = append(addresses, timelockRef)

	callProxyRef, err := contract.MaybeDeployContract(b, mcms_ops.OpDeployCallProxy, chain, contract.DeployInput[mcms_ops.OpDeployCallProxyInput]{
		ChainSelector:  chain.Selector,
		Qualifier:      qualifierPtr,
		TypeAndVersion: deployment.NewTypeAndVersion(common_utils.CallProxy, *mcms_ops.MCMSVersion),
		Args: mcms_ops.OpDeployCallProxyInput{
			TimelockAddress: timelockAddr,
		},
	}, nil)
	require.NoError(t, err)
	addresses = append(addresses, callProxyRef)
	callProxyAddr := common.HexToAddress(callProxyRef.Address)

	_, err = operations.ExecuteOperation(b, mcms_ops.OpGrantRoleTimelock, chain, contract.FunctionInput[mcms_ops.OpGrantRoleTimelockInput]{
		ChainSelector: chain.Selector,
		Address:       timelockAddr,
		Args: mcms_ops.OpGrantRoleTimelockInput{
			RoleID:  mcms_ops.EXECUTOR_ROLE.ID,
			Account: callProxyAddr,
		},
	})
	require.NoError(t, err)

	return timelockAddr, addresses
}

