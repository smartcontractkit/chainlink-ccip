package mock_receiver

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_receiver_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "MockReceiver"

type ConstructorArgs struct {
	RequiredVerifiers []common.Address
	OptionalVerifiers []common.Address
	OptionalThreshold uint8
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "mock-receiver-v2:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the MockReceiverV2 contract",
	ContractMetadata: mock_receiver_v2.MockReceiverV2MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(mock_receiver_v2.MockReceiverV2Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
