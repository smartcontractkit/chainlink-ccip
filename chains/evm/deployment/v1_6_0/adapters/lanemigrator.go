package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"

	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	offbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	orbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type LaneMigrator struct{}

func (r *LaneMigrator) VerifyPreconditions(_ deployment.Environment, _ deploy.LaneMigratorConfig, _ changesets.MCMSReader) error {
	return nil
}

// UpdateVersionWithRouter is a sequence that updates Ramps to use the new Router
//
// It fetches the existing onRamp and offRamp addresses from the provided ExistingAddresses, then calls the necessary functions to update the onRamp and offRamp to use the new Router.
//
// This sequence assumes that the destChainConfig on OnRamp and SourceChainConfig on OffRamp do not need to be updated, and only updates the Router address used by the Ramps.
// This should not be used to set preliminary dest or source chain config on ramps
func (r *LaneMigrator) UpdateVersionWithRouter() *cldf_ops.Sequence[deploy.RampUpdaterConfig, sequences.OnChainOutput, chain.BlockChains] {
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

			onRamp, err := orbind.NewOnRamp(onRampAddr, c.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to bind OnRamp at %s: %w", onRampAddr.Hex(), err)
			}
			offRamp, err := offbind.NewOffRamp(offRampAddr, c.Client)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to bind OffRamp at %s: %w", offRampAddr.Hex(), err)
			}

			routerRef := input.RouterAddr
			routerAddr, err := evm_datastore_utils.ToEVMAddress(routerRef)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error formatting router address ref: %w", err)
			}
			var onRampArgs []orbind.OnRampDestChainConfigArgs
			var offRampArgs []offbind.OffRampSourceChainConfigArgs
			for _, remoteChainSelector := range input.RemoteChainSelectors {
				existingDestChainCfgOut, err := cldf_ops.ExecuteOperation(b, onrampops.NewReadGetDestChainConfig(onRamp), c, ops2contract.FunctionInput[uint64]{
					Args: remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing destChainConfig for onRamp: %w", err)
				}
				existingDestChainCfg := existingDestChainCfgOut.Output
				if existingDestChainCfg.Router == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("no destchain config is set for remote chain %d on chain %d on onRamp."+
						" configure lanes with test router first before migrating", remoteChainSelector, input.ChainSelector)
				}

				srcChainCfgOut, err := cldf_ops.ExecuteOperation(b, offrampops.NewReadGetSourceChainConfig(offRamp), c, ops2contract.FunctionInput[uint64]{
					Args: remoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error fetching existing sourceChainConfig for offRamp: %w", err)
				}
				existingSrcChainCfg := srcChainCfgOut.Output
				if existingSrcChainCfg.Router == (common.Address{}) {
					return sequences.OnChainOutput{}, fmt.Errorf("no source chain config is set for remote chain %d on chain %d on offRamp."+
						" configure lanes with test router first before migrating", remoteChainSelector, input.ChainSelector)
				}

				onRampArgs = append(onRampArgs, orbind.OnRampDestChainConfigArgs{
					DestChainSelector: remoteChainSelector,
					Router:            routerAddr,
					AllowlistEnabled:  existingDestChainCfg.AllowlistEnabled,
				})

				offRampArgs = append(offRampArgs, offbind.OffRampSourceChainConfigArgs{
					SourceChainSelector:       remoteChainSelector,
					Router:                    routerAddr,
					IsEnabled:                 existingSrcChainCfg.IsEnabled,
					IsRMNVerificationDisabled: true,
					OnRamp:                    existingSrcChainCfg.OnRamp,
				})
			}
			writeOutputOnRamp, err := cldf_ops.ExecuteOperation(b, onrampops.NewWriteApplyDestChainConfigUpdates(onRamp), c, ops2contract.FunctionInput[[]orbind.OnRampDestChainConfigArgs]{
				Args: onRampArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying destChainConfig update to onRamp: %w", err)
			}
			writes = append(writes, writeOutputOps2ToLegacy(writeOutputOnRamp.Output))

			writeOutputOffRamp, err := cldf_ops.ExecuteOperation(
				b, offrampops.NewWriteApplySourceChainConfigUpdates(offRamp), c,
				ops2contract.FunctionInput[[]offbind.OffRampSourceChainConfigArgs]{
					Args: offRampArgs,
				})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error applying sourceChainConfig update to offRamp: %w", err)
			}
			writes = append(writes, writeOutputOps2ToLegacy(writeOutputOffRamp.Output))
			batchOp, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batchOp},
			}, nil
		})
}
