package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type RouterUpdater struct{}

// UpdateRouter is a sequence that updates the Router to use the existing Ramps.
//
// It fetches the existing Router address and Ramp addresses from the provided ExistingAddresses, then calls the necessary functions to update the Router to use the existing Ramps.
// This sequence assumes that there is only one onRamp and offRamp per remote chain selector, and
// will not work with 1.5 system where there are multiple onRamps and offRamps per chain for each remote chain selector.
func (u *RouterUpdater) UpdateRouter() *cldf_ops.Sequence[deploy.RouterUpdaterConfig, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"router-updater:sequence-update-router-with-ramps",
		semver.MustParse("1.2.0"),
		"Updates the Router contract to use the existing Ramps",
		func(b cldf_ops.Bundle, chains chain.BlockChains, input deploy.RouterUpdaterConfig) (output sequences.OnChainOutput, err error) {
			c, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("EVM chain with selector %d not found in environment", input.ChainSelector)
			}
			var writes []contract.WriteOutput
			ds := datastore.NewMemoryDataStore()
			for _, addrRef := range input.ExistingAddresses {
				if err := ds.Addresses().Add(addrRef); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error adding address ref to temp datastore: %w", err)
				}
			}
			tempDS := ds.Seal()
			// get router from existing addresses
			routerAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(routerops.ContractType),
					Version:       routerops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding router address ref: %w", err)
			}
			onRampAddr, err := evm_datastore_utils.ToEVMAddress(input.OnRamp)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error formatting onRamp address: %w", err)
			}
			offRampAddr, err := evm_datastore_utils.ToEVMAddress(input.OffRamp)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error formatting offRamp address: %w", err)
			}
			onRampAdds := make([]routerops.OnRamp, 0)
			offRampAdds := make([]routerops.OffRamp, 0)
			// this assumes that there is only one onRamp and offRamp per c,
			// it will not work with 1.5 system where there can be multiple onRamps and offRamps per c
			// for each remote c selector
			for _, remoteChainSelector := range input.RemoteChainSelectors {
				onRampAdds = append(onRampAdds, routerops.OnRamp{
					DestChainSelector: remoteChainSelector,
					OnRamp:            onRampAddr,
				})
				offRampAdds = append(offRampAdds, routerops.OffRamp{
					SourceChainSelector: remoteChainSelector,
					OffRamp:             offRampAddr,
				})
			}
			out, err := cldf_ops.ExecuteOperation(
				b, routerops.ApplyRampUpdates, c, contract.FunctionInput[routerops.ApplyRampsUpdatesArgs]{
					ChainSelector: input.ChainSelector,
					Args: routerops.ApplyRampsUpdatesArgs{
						OnRampUpdates:  onRampAdds,
						OffRampRemoves: []routerops.OffRamp{},
						OffRampAdds:    offRampAdds,
					},
					Address: routerAddr,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error executing ApplyRampUpdates operation: %w", err)
			}
			writes = append(writes, out.Output)
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error creating batch operation: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})
}
