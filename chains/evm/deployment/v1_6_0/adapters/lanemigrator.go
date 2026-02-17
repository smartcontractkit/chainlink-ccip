package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"

	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type LaneMigrater struct{}

// UpdateVersionWithRouter is a sequence that updates Ramps to use the new Router
//
// It fetches the existing onRamp and offRamp addresses from the provided ExistingAddresses, then calls the necessary functions to update the onRamp and offRamp to use the new Router.
//
// This sequence assumes that the destChainConfig on OnRamp and SourceChainConfig on OffRamp do not need to be updated, and only updates the Router address used by the Ramps.
// This should not be used to set preliminary dest or source chain config on ramps
func (r *LaneMigrater) UpdateVersionWithRouter() *cldf_ops.Sequence[deploy.RampUpdaterConfig, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-updater:sequence-update-ramps-with-router",
		semver.MustParse("1.6.0"),
		"Updates Ramps contracts to use the updated Router contract",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input deploy.RampUpdaterConfig) (output sequences.OnChainOutput, err error) {
			var writes []contract.WriteOutput
			c, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("c with selector %d not found in environment", input.ChainSelector)
			}
			ds := datastore.NewMemoryDataStore()
			for _, addrRef := range input.ExistingAddresses {
				if err := ds.Addresses().Add(addrRef); err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error adding address ref to temp datastore: %w", err)
				}
			}
			tempDS := ds.Seal()
			// fetch onRamp and offRamp from the existing addresses
			onRampAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(onrampops.ContractType),
					Version:       onrampops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding onRamp address ref: %w", err)
			}
			offRampAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(offrampops.ContractType),
					Version:       offrampops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding offRamp address ref: %w", err)
			}

			routerRef := input.RouterAddr
			routerAddr, err := evm_datastore_utils.ToEVMAddress(routerRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error formatting router address ref: %w", err)
			}
			var onRampArgs []onrampops.DestChainConfigArgs
			var offRampArgs []offrampops.SourceChainConfigArgs
			for _, remoteChainSelector := range input.RemoteChainSelectors {
				// get existing destChainConfig for the onRamp
				existingDestChainCfgOut, err := cldf_ops.ExecuteOperation(b, onrampops.GetDestChainConfig, c, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       onRampAddr,
					Args:          remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing destChainConfig for onRamp: %w", err)
				}
				existingDestChainCfg := existingDestChainCfgOut.Output.(onramp.GetDestChainConfig)
				if existingDestChainCfg.Router == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("no destchain config is set for remote chain %d on chain %d on onRamp."+
						" configure lanes with test router first before migrating", remoteChainSelector, input.ChainSelector)
				}
				// update router on onRamp for the remote chain
				existingDestChainCfg.Router = routerAddr

				// get the sourceChainConfig for the offRamp
				srcChainCfgOut, err := cldf_ops.ExecuteOperation(b, offrampops.GetSourceChainConfig, c, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       offRampAddr,
					Args:          remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing sourceChainConfig for offRamp: %w", err)
				}
				existingSrcChainCfg := srcChainCfgOut.Output
				if existingSrcChainCfg.Router == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("no source chain config is set for remote chain %d on chain %d on offRamp."+
						" configure lanes with test router first before migrating", remoteChainSelector, input.ChainSelector)
				}
				// update router on offRamp for the remote chain
				existingSrcChainCfg.Router = routerAddr
				onRampArgs = append(onRampArgs, onrampops.DestChainConfigArgs{
					DestChainSelector: remoteChainSelector,
					Router:            routerAddr,
					AllowlistEnabled:  existingDestChainCfg.AllowlistEnabled,
				})

				offRampArgs = append(offRampArgs, offrampops.SourceChainConfigArgs{
					SourceChainSelector:       remoteChainSelector,
					Router:                    routerAddr,
					IsEnabled:                 existingSrcChainCfg.IsEnabled,
					IsRMNVerificationDisabled: true,
					OnRamp:                    existingSrcChainCfg.OnRamp,
				})
			}
			//  set the destChainConfig with the updated router
			writeOutputOnRamp, err := cldf_ops.ExecuteOperation(b, onrampops.ApplyDestChainConfigUpdates, c, contract.FunctionInput[[]onrampops.DestChainConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       onRampAddr,
				Args:          onRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying destChainConfig update to onRamp: %w", err)
			}
			writes = append(writes, writeOutputOnRamp.Output)
			// now set the sourceChainConfig with the updated router
			writeOutputOffRamp, err := cldf_ops.ExecuteOperation(
				b, offrampops.ApplySourceChainConfigUpdates, c,
				contract.FunctionInput[[]offrampops.SourceChainConfigArgs]{
					ChainSelector: input.ChainSelector,
					Address:       offRampAddr,
					Args:          offRampArgs,
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying sourceChainConfig update to offRamp: %w", err)
			}
			writes = append(writes, writeOutputOffRamp.Output)
			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batchOp},
			}, nil
		})
}
