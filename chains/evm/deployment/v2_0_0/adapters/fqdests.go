package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fqdests"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var _ fqdests.FQDestsAdapter = (*FQDestsV2Adapter)(nil)

type FQDestsV2Adapter struct {
	evm *evmseq.EVMAdapter
}

func NewFQDestsV2Adapter(evmAdapter *evmseq.EVMAdapter) *FQDestsV2Adapter {
	return &FQDestsV2Adapter{evm: evmAdapter}
}

func (a *FQDestsV2Adapter) GetFeeContractRef(e cldf.Environment, src, dst uint64) (datastore.AddressRef, error) {
	ds := e.DataStore
	fqAddr, err := a.evm.GetFQAddressDynamic(ds, src, e.BlockChains)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get FeeQuoter address for chain selector %d: %w", src, err)
	}

	filter := datastore.AddressRef{
		Type:          datastore.ContractType(fqops.ContractType),
		Address:       common.BytesToAddress(fqAddr).Hex(),
		ChainSelector: src,
	}
	ref, err := datastore_utils.FindAndFormatRef(ds, filter, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find FeeQuoter address ref for chain selector %d: %w", src, err)
	}
	return ref, nil
}

func (a *FQDestsV2Adapter) GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig {
	return a.evm.GetFeeQuoterDestChainConfig()
}

func (a *FQDestsV2Adapter) GetOnchainDestChainConfig(e cldf.Environment, src, dst uint64) (lanes.FeeQuoterDestChainConfig, error) {
	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	fqRef, err := a.GetFeeContractRef(e, src, dst)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, err
	}

	fqAddr := common.HexToAddress(fqRef.Address)
	fq, err := fqops.NewFeeQuoterContract(fqAddr, chain.Client)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to instantiate FeeQuoter at %s on chain %d: %w", fqAddr.Hex(), src, err)
	}

	cfg, err := fq.GetDestChainConfig(&bind.CallOpts{Context: e.GetContext()}, dst)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to get dest chain config from FeeQuoter at %s for src %d, dst %d: %w", fqAddr.Hex(), src, dst, err)
	}

	return lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   cfg.IsEnabled,
		MaxDataBytes:                cfg.MaxDataBytes,
		MaxPerMsgGasLimit:           cfg.MaxPerMsgGasLimit,
		DestGasOverhead:             cfg.DestGasOverhead,
		DestGasPerPayloadByteBase:   cfg.DestGasPerPayloadByteBase,
		ChainFamilySelector:         bytesToUint32(cfg.ChainFamilySelector),
		DefaultTokenFeeUSDCents:     cfg.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead: cfg.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:           cfg.DefaultTxGasLimit,
		NetworkFeeUSDCents:          cfg.NetworkFeeUSDCents,
		V2Params: &lanes.FeeQuoterV2Params{
			LinkFeeMultiplierPercent: cfg.LinkFeeMultiplierPercent,
		},
	}, nil
}

func (a *FQDestsV2Adapter) ApplyDestChainConfigUpdates(e cldf.Environment) *operations.Sequence[
	fqdests.ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains,
] {
	return operations.NewSequence(
		"FQDestsApplyDestChainConfigUpdates",
		semver.MustParse("2.0.0"),
		"Applies FeeQuoter 2.0 destination chain config updates",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input fqdests.ApplyDestChainConfigSequenceInput) (sequences.OnChainOutput, error) {
			var result sequences.OnChainOutput

			fqRef, err := a.GetFeeContractRef(e, input.Selector, 0)
			if err != nil {
				return result, fmt.Errorf("failed to get FeeQuoter address for chain %d: %w", input.Selector, err)
			}
			fqAddr := common.HexToAddress(fqRef.Address)

			chain, ok := chains.EVMChains()[input.Selector]
			if !ok {
				return result, fmt.Errorf("chain with selector %d not defined", input.Selector)
			}

			args := make([]fqops.DestChainConfigArgs, 0, len(input.Settings))
			for dst, cfg := range input.Settings {
				args = append(args, fqops.DestChainConfigArgs{
					DestChainSelector: dst,
					DestChainConfig:   evmseq.TranslateFQtoV2(cfg),
				})
			}

			report, err := operations.ExecuteOperation(
				b, fqops.ApplyDestChainConfigUpdates, chain,
				contract.FunctionInput[[]fqops.DestChainConfigArgs]{
					ChainSelector: chain.Selector,
					Address:       fqAddr,
					Args:          args,
				},
			)
			if err != nil {
				return result, fmt.Errorf("failed to apply dest chain config updates on FeeQuoter 2.0 for chain %d: %w", input.Selector, err)
			}

			batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return result, err
			}
			result.BatchOps = append(result.BatchOps, batch)

			return result, nil
		},
	)
}

func bytesToUint32(b [4]byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}
