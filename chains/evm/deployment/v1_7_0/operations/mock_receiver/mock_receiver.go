package mock_receiver

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
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

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "mock-receiver-v2:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the MockReceiverV2 contract",
	ContractType:     ContractType,
	ContractMetadata: mock_receiver_v2.MockReceiverV2MetaData,
	BytecodeByVersion: map[string]contract.Bytecode{
		semver.MustParse("1.7.0").String(): {EVM: common.FromHex(mock_receiver_v2.MockReceiverV2Bin)},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
