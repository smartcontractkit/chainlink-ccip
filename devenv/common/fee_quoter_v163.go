package common

// This file overrides the v1.6.0 FeeQuoter deploy operation with v1.6.3
// bytecode. The v1.6.0 bytecode only validates EVM and SVM chain selectors,
// while v1.6.3 adds non-EVM family selector support.
//
// The ABI and constructor args are identical between versions — only the
// bytecode differs. This init()-based override means the standard
// DeployChainContracts sequence automatically uses v1.6.3 bytecode without
// any sequence duplication.

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
)

func init() {
	fqops.Deploy = contract.NewDeploy(contract.DeployParams[fqops.ConstructorArgs]{
		Name:        "fee-quoter:deploy",
		Version:     fqops.Version,
		Description: "Deploys FeeQuoter with v1.6.3 bytecode (non-EVM chain selector support)",
		ContractMetadata: &bind.MetaData{
			ABI: fqops.FeeQuoterABI,
			Bin: fq163ops.FeeQuoterBin,
		},
		BytecodeByTypeAndVersion: map[string]contract.Bytecode{
			deployment.NewTypeAndVersion(fqops.ContractType, *fqops.Version).String(): {
				EVM: common.FromHex(fq163ops.FeeQuoterBin),
			},
		},
		Validate: func(fqops.ConstructorArgs) error { return nil },
	})
}
