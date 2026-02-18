package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	offrampops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
)

const (
	DefaultTxGasLimit        uint32 = 200_000
	DefaultMaxPerMsgGasLimit        = 8_000_000
	DefaultMaxDataBytes             = 32_000
)

type LaneMigrator struct{}

// UpdateVersionWithRouter is a sequence that updates Ramps to use the new Router and also updates the fee quoter dest chain config with default tx gas limit as 8M
//
// It fetches the existing onRamp and offRamp addresses from the provided ExistingAddresses, then calls the necessary functions to update the onRamp and offRamp to use the new Router.
//
// This sequence assumes that the destChainConfig on OnRamp and SourceChainConfig on OffRamp do not need to be updated, and only updates the Router address used by the Ramps.
// If you need to update the destChainConfig or sourceChainConfig, please use the ConfigureChainForLanes sequence instead.
func (r *LaneMigrator) UpdateVersionWithRouter() *cldf_ops.Sequence[deploy.RampUpdaterConfig, sequences.OnChainOutput, chain.BlockChains] {
	return cldf_ops.NewSequence(
		"ramp-updater:sequence-update-ramps-with-router",
		semver.MustParse("1.7.0"),
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
			feequoterAddr, err := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(fqops.ContractType),
					Version:       fqops.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error finding feequoter address ref: %w", err)
			}
			feeQuoterContract, err := fee_quoter.NewFeeQuoter(feequoterAddr, c.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error creating fee quoter contract instance: %w", err)
			}
			routerRef := input.RouterAddr
			routerAddr, err := evm_datastore_utils.ToEVMAddress(routerRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error formatting router address ref: %w", err)
			}
			var onRampArgs []onrampops.DestChainConfigArgs
			var offRampArgs []offrampops.SourceChainConfigArgs
			var fqArgs []fqops.DestChainConfigArgs
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
				existingDestChainCfg := existingDestChainCfgOut.Output
				if existingDestChainCfg.AddressBytesLength == 0 {
					return sequences.OnChainOutput{}, fmt.Errorf("no destchain config is set for remote chain %d on chain %d on onRamp."+
						" configure lanes with test router first before migrating", remoteChainSelector, input.ChainSelector)
				}
				// update router on onRamp for the remote c
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
				// update router on offRamp for the remote c
				existingSrcChainCfg.Router = routerAddr
				onRampArgs = append(onRampArgs, onrampops.DestChainConfigArgs{
					DestChainSelector:         remoteChainSelector,
					Router:                    routerAddr,
					AddressBytesLength:        existingDestChainCfg.AddressBytesLength,
					TokenReceiverAllowed:      existingDestChainCfg.TokenReceiverAllowed,
					MessageNetworkFeeUSDCents: existingDestChainCfg.MessageNetworkFeeUSDCents,
					TokenNetworkFeeUSDCents:   existingDestChainCfg.TokenNetworkFeeUSDCents,
					BaseExecutionGasCost:      existingDestChainCfg.BaseExecutionGasCost,
					DefaultCCVs:               existingDestChainCfg.DefaultCCVs,
					LaneMandatedCCVs:          existingDestChainCfg.LaneMandatedCCVs,
					DefaultExecutor:           existingDestChainCfg.DefaultExecutor,
					OffRamp:                   existingDestChainCfg.OffRamp,
				})

				offRampArgs = append(offRampArgs, offrampops.SourceChainConfigArgs{
					SourceChainSelector: remoteChainSelector,
					Router:              routerAddr,
					IsEnabled:           existingSrcChainCfg.IsEnabled,
					OnRamps:             existingSrcChainCfg.OnRamps,
					DefaultCCVs:         existingSrcChainCfg.DefaultCCVs,
					LaneMandatedCCVs:    existingSrcChainCfg.LaneMandatedCCVs,
				})
				// this also needs feequoter update for dest chain config,
				// fetch existing destChainConfig for feequoter
				dstChainCfg, err := feeQuoterContract.GetDestChainConfig(&bind.CallOpts{
					Context: b.GetContext(),
				}, remoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing destChainConfig for fee quoter: %w", err)
				}
				fqArgs = append(fqArgs, fqops.DestChainConfigArgs{
					DestChainSelector: remoteChainSelector,
					DestChainConfig: adapters.FeeQuoterDestChainConfig{
						IsEnabled:                   dstChainCfg.IsEnabled,
						MaxDataBytes:                DefaultMaxDataBytes,
						MaxPerMsgGasLimit:           DefaultMaxPerMsgGasLimit,
						DestGasOverhead:             dstChainCfg.DestGasOverhead,
						DestGasPerPayloadByteBase:   dstChainCfg.DestGasPerPayloadByteBase,
						ChainFamilySelector:         dstChainCfg.ChainFamilySelector,
						DefaultTokenFeeUSDCents:     dstChainCfg.DefaultTokenFeeUSDCents,
						DefaultTokenDestGasOverhead: dstChainCfg.DefaultTokenDestGasOverhead,
						DefaultTxGasLimit:           DefaultTxGasLimit,
						NetworkFeeUSDCents:          dstChainCfg.NetworkFeeUSDCents,
						LinkFeeMultiplierPercent:    dstChainCfg.LinkFeeMultiplierPercent,
					},
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
			// update fq 1.7 to have defaultTxLimit set to 8M
			fqDestChainUpdateRep, err := cldf_ops.ExecuteOperation(b, fqops.ApplyDestChainConfigUpdates, c, contract.FunctionInput[[]fqops.DestChainConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       feequoterAddr,
				Args:          fqArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying destChainConfig update to fee quoter: %w", err)
			}
			writes = append(writes, fqDestChainUpdateRep.Output)
			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batchOp},
			}, nil
		})
}
