package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func (a *SolanaAdapter) SetOCR3Config() *cldf_ops.Sequence[deployops.SetOCR3ConfigInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return SetOCR3Config
}

var SetOCR3Config = cldf_ops.NewSequence(
	"setocr3config",
	semver.MustParse("1.6.0"),
	"Set OCR3 Config on Solana chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input deployops.SetOCR3ConfigInput) (output sequences.OnChainOutput, err error) {
		var result sequences.OnChainOutput
		a := &SolanaAdapter{}
		offRampAddr, err := a.GetOffRampAddress(input.Datastore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		// Set OCR3 Config
		out, err := operations.ExecuteOperation(b, offrampops.SetOcr3, chains.SolanaChains()[input.ChainSelector], offrampops.SetOcr3Params{
			OffRamp:            solana.PublicKeyFromBytes(offRampAddr),
			SetOCR3ConfigInput: input,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize OffRamp: %w", err)
		}
		result.Addresses = append(result.Addresses, out.Output.Addresses...)
		result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)
		return result, nil
	},
)
