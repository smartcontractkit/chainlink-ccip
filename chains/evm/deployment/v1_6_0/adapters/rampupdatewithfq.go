package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type RampUpdateWithFQ struct{}

func (ru RampUpdateWithFQ) ResolveRampsInput(e cldf.Environment, input deploy.UpdateRampsInput) (deploy.UpdateRampsInput, error) {
	// fetch address of Ramps
	onRampAddr := datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByType(datastore.ContractType(onrampops.ContractType)),
			datastore.AddressRefByVersion(onrampops.Version),
		),
		input.ChainSelector,
		onrampops.ContractType,
		onrampops.Version,
		"",
	)
	if datastore_utils.IsAddressRefEmpty(onRampAddr) {
		return input, fmt.Errorf("onramp address not found for chain selector %d", input.ChainSelector)
	}
	input.OnRampAddressRef = onRampAddr

	offRampAddr := datastore_utils.GetAddressRef(
		e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(input.ChainSelector),
			datastore.AddressRefByType(datastore.ContractType(offrampops.ContractType)),
			datastore.AddressRefByVersion(offrampops.Version),
		),
		input.ChainSelector,
		offrampops.ContractType,
		offrampops.Version,
		"",
	)
	if datastore_utils.IsAddressRefEmpty(offRampAddr) {
		return input, fmt.Errorf("offramp address not found for chain selector %d", input.ChainSelector)
	}
	input.OffRampAddressRef = offRampAddr
	return input, nil
}

// SequenceUpdateRampsWithFeeQuoter updates OnRamp and OffRamp contracts to use the new FeeQuoter contract
func (ru RampUpdateWithFQ) SequenceUpdateRampsWithFeeQuoter() *cldf_ops.Sequence[deploy.UpdateRampsInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-update-with-fq:sequence-update-ramps-with-fee-quoter",
		semver.MustParse("1.6.0"),
		"Updates Ramps contracts to use the new FeeQuoter contract",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.UpdateRampsInput) (output sequences.OnChainOutput, err error) {
			var writes []contract.WriteOutput
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found in environment", input.ChainSelector)
			}
			onDCfgReport, err := cldf_ops.ExecuteOperation(b, onrampops.GetDynamicConfig, chain, contract.FunctionInput[struct{}]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.OnRampAddressRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			existingDynamicConfig := onDCfgReport.Output
			if existingDynamicConfig.FeeQuoter != common.HexToAddress(input.FeeQuoterAddress.Address) {
				// Update OnRamp's FeeQuoter address
				existingDynamicConfig.FeeQuoter = common.HexToAddress(input.FeeQuoterAddress.Address)
				onRampReport, err := cldf_ops.ExecuteOperation(b, onrampops.SetDynamicConfig, chain, contract.FunctionInput[onrampops.DynamicConfig]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(input.OnRampAddressRef.Address),
					Args:          existingDynamicConfig,
				})
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				writes = append(writes, onRampReport.Output)
			}
			// Similarly, update OffRamp's FeeQuoter address
			offDCfgReport, err := cldf_ops.ExecuteOperation(b, offrampops.GetDynamicConfig, chain, contract.FunctionInput[struct{}]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.OffRampAddressRef.Address),
			})
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			existingOffDynamicConfig := offDCfgReport.Output
			if existingOffDynamicConfig.FeeQuoter != common.HexToAddress(input.FeeQuoterAddress.Address) {
				existingOffDynamicConfig.FeeQuoter = common.HexToAddress(input.FeeQuoterAddress.Address)
				offRampReport, err := cldf_ops.ExecuteOperation(b, offrampops.SetDynamicConfig, chain, contract.FunctionInput[offrampops.DynamicConfig]{
					ChainSelector: input.ChainSelector,
					Address:       common.HexToAddress(input.OffRampAddressRef.Address),
					Args:          existingOffDynamicConfig,
				})
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				writes = append(writes, offRampReport.Output)
			}
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		},
	)
}
