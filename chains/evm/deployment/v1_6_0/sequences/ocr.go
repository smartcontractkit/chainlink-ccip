package sequences

import (
	"github.com/Masterminds/semver/v3"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *EVMAdapter) SetOCR3Config() *cldf_ops.Sequence[deployops.SetOCR3ConfigInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return SetOCR3Config
}

var SetOCR3Config = cldf_ops.NewSequence(
	"set-ocr3-config",
	semver.MustParse("1.6.0"),
	"Sets the OCR3 configuration for CCIP 1.6.0 on an EVM chain",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input deployops.SetOCR3ConfigInput) (output sequences.OnChainOutput, err error) {
		return output, nil
	},
)
