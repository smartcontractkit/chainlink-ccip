package sequences

import (
	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *EVMAdapter) InitializeTimelockAddress(e deployment.Environment, input mcms.Input) error {
	evmDeployer := &evm1_0_0.EVMTransferOwnershipAdapter{}
	return evmDeployer.InitializeTimelockAddress(e, input)
}

func (a *EVMAdapter) SequenceTransferOwnershipViaMCMS() *cldf_ops.Sequence[api.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMTransferOwnershipAdapter{}
	return evmDeployer.SequenceTransferOwnershipViaMCMS()
}

func (a *EVMAdapter) ShouldAcceptOwnershipWithTransferOwnership(e deployment.Environment, in api.TransferOwnershipPerChainInput) (bool, error) {
	evmDeployer := &evm1_0_0.EVMTransferOwnershipAdapter{}
	return evmDeployer.ShouldAcceptOwnershipWithTransferOwnership(e, in)
}

func (a *EVMAdapter) SequenceAcceptOwnership() *cldf_ops.Sequence[api.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	evmDeployer := &evm1_0_0.EVMTransferOwnershipAdapter{}
	return evmDeployer.SequenceAcceptOwnership()
}
