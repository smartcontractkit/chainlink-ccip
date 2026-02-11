package advanced_pool_hooks_extractor

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/advanced_pool_hooks_extractor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "AdvancedPoolHooksExtractor"

var Version = semver.MustParse("2.0.0")

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "advanced-pool-hooks-extractor:deploy",
	Version:          Version,
	Description:      "Deploys the AdvancedPoolHooksExtractor contract",
	ContractMetadata: advanced_pool_hooks_extractor.AdvancedPoolHooksExtractorMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(advanced_pool_hooks_extractor.AdvancedPoolHooksExtractorBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
