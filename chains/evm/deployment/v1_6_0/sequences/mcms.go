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

func (a *EVMAdapter) getTransferOwnershipAdapter() *evm1_0_0.EVMTransferOwnershipAdapter {
	if a.transferOwnershipAdapter != nil {
		return a.transferOwnershipAdapter
	}
	return &evm1_0_0.EVMTransferOwnershipAdapter{}
}

func (a *EVMAdapter) InitializeTimelockAddress(e deployment.Environment, input mcms.Input) error {
	return a.getTransferOwnershipAdapter().InitializeTimelockAddress(e, input)
}

func (a *EVMAdapter) SequenceTransferOwnershipViaMCMS() *cldf_ops.Sequence[api.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return a.getTransferOwnershipAdapter().SequenceTransferOwnershipViaMCMS()
}

func (a *EVMAdapter) ShouldAcceptOwnershipWithTransferOwnership(e deployment.Environment, in api.TransferOwnershipPerChainInput) (bool, error) {
	return a.getTransferOwnershipAdapter().ShouldAcceptOwnershipWithTransferOwnership(e, in)
}

func (a *EVMAdapter) SequenceAcceptOwnership() *cldf_ops.Sequence[api.TransferOwnershipPerChainInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return a.getTransferOwnershipAdapter().SequenceAcceptOwnership()
}
