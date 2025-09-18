package mock_receiver

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_receiver_v2"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "MockReceiver"

type ConstructorArgs struct {
	RequiredVerifiers []common.Address
	OptionalVerifiers []common.Address
	OptionalThreshold uint8
}

var Deploy = contract.NewDeploy(
	"mock-receiver-v2:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the MockReceiverV2 contract",
	ContractType,
	mock_receiver_v2.MockReceiverV2ABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := mock_receiver_v2.DeployMockReceiverV2(opts, backend, args.RequiredVerifiers, args.OptionalVerifiers, args.OptionalThreshold)
			return address, tx, err
		},
	},
)
