package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	offbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

func (a *EVMAdapter) SetOCR3Config() *cldf_ops.Sequence[deployops.SetOCR3ConfigInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return SetOCR3Config
}

var SetOCR3Config = cldf_ops.NewSequence(
	"set-ocr3-config",
	semver.MustParse("1.6.0"),
	"Sets the OCR3 configuration for CCIP 1.6.0 on an EVM chain",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input deployops.SetOCR3ConfigInput) (output sequences.OnChainOutput, err error) {
		writes := make([]ops2contract.WriteOutput, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}
		e := &EVMAdapter{}
		offRampAddr, err := e.GetOffRampAddress(input.Datastore, input.ChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		offContract, err := offbind.NewOffRamp(common.BytesToAddress(offRampAddr), chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("bind offRamp: %w", err)
		}
		update := make([]offbind.MultiOCR3BaseOCRConfigArgs, 0)
		for _, cfg := range input.Configs {
			var signerAddresses []common.Address
			var transmitterAddresses []common.Address
			for _, s := range cfg.Signers {
				signerAddresses = append(signerAddresses, common.BytesToAddress(s))
			}
			for _, t := range cfg.Transmitters {
				transmitterAddresses = append(transmitterAddresses, common.BytesToAddress(t))
			}
			update = append(update, offbind.MultiOCR3BaseOCRConfigArgs{
				ConfigDigest:                   cfg.ConfigDigest,
				OcrPluginType:                  uint8(cfg.PluginType),
				F:                              cfg.F,
				IsSignatureVerificationEnabled: cfg.IsSignatureVerificationEnabled,
				Signers:                        signerAddresses,
				Transmitters:                   transmitterAddresses,
			})
		}
		report, err := operations.ExecuteOperation(b, offrampops.NewWriteSetOCR3Configs(offContract), chain, ops2contract.FunctionInput[[]offbind.MultiOCR3BaseOCRConfigArgs]{
			Args: update,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OffRampSetOcr3 on %s: %w", chain, err)
		}
		writes = append(writes, report.Output)
		batch, err := ops2contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)
