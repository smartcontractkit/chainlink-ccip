package usdc_token_pool_cctp_v2

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool_cctp_v2"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "USDCTokenPoolCCTPV2"
var Version *semver.Version = semver.MustParse("1.6.4")

type ConstructorArgs struct {
	TokenMessenger              common.Address
	CCTPMessageTransmitterProxy common.Address
	Token                       common.Address
	Allowlist                   []common.Address
	RMNProxy                    common.Address
	Router                      common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "usdc-token-pool-cctp-v2:deploy",
	Version:          Version,
	Description:      "Deploys the USDCTokenPoolCCTPV2 contract",
	ContractMetadata: usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(usdc_token_pool_cctp_v2.USDCTokenPoolCCTPV2Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
