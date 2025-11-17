package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type FeeQuoterSetTokenTransferFeeConfigSequenceInput struct {
	RemoteChainConfigs map[uint64]map[solana.PublicKey]fee_quoter.TokenTransferFeeConfig
	DataStore          datastore.DataStore
	FeeQuoter          solana.PublicKey
	Selector           uint64
}

func (a *SolanaAdapter) SetTokenTransferFee() *cldf_ops.Sequence[FeeQuoterSetTokenTransferFeeConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return SetTokenTransferFeeConfig
}

var SetTokenTransferFeeConfig = cldf_ops.NewSequence(
	"fee-quoter:set-token-transfer-fee-config",
	semver.MustParse("1.6.0"),
	"Sets token transfer fee configs on the FeeQuoter 1.6.0 contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input FeeQuoterSetTokenTransferFeeConfigSequenceInput) (sequences.OnChainOutput, error) {
		var result sequences.OnChainOutput
		chain, ok := chains.SolanaChains()[input.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("solana chain with selector %d not found", input.Selector)
		}
		fqAddrBytes, err := (&SolanaAdapter{}).GetFQAddress(input.DataStore, input.Selector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get FeeQuoter address: %w", err)
		}
		fqPubkey := solana.PublicKeyFromBytes(fqAddrBytes)
		for dst, dstCfg := range input.RemoteChainConfigs {
			for tok, feeCfg := range dstCfg {
				out, err := operations.ExecuteOperation(b,
					fqops.SetTokenTransferFeeConfig,
					chain,
					fqops.SetTokenTransferFeeConfigInput{
						SrcSelector: input.Selector,
						DstSelector: dst,
						FeeQuoter:   fqPubkey,
						Config:      feeCfg,
						Token:       tok,
					},
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set token transfer fee config for token %s and destination %d: %w", tok, dst, err)
				}
				result.Addresses = append(result.Addresses, out.Output.Addresses...)
				result.BatchOps = append(result.BatchOps, out.Output.BatchOps...)
			}
		}
		return result, nil
	},
)
