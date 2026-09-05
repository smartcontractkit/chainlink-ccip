package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

// deployMCMSInstanceForTest deploys a Proposer/Bypasser/Canceller MCMS, a Timelock, and a
// CallProxy under the given qualifier, and grants the CallProxy the executor role on the
// Timelock. This is what changesets.NewOutputBuilder's EVMMCMSReader needs on-chain in order to
// resolve chain metadata / timelock addresses for a real MCMS proposal.
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
			TimelockMinDelay: big.NewInt(1),
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
