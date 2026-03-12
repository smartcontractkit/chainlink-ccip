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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/lombard_token_pool"
	usdc_token_pool_proxy "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/usdc_token_pool_proxy"
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
			// this assumes that there is only one onRamp and offRamp per chain,
			// it will not work with 1.5 system where there can be multiple onRamps and offRamps per chain
			// for each remote chain selector
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

			// If LombardTokenPool 2.0.0 is in existingAddresses, set it on the TokenAdminRegistry for LBTC (same batch). Token address is read from the pool via GetToken.
			lombardPoolAddr, lombardErr := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(lombard_token_pool.ContractType),
					Version:       lombard_token_pool.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if lombardErr == nil {
				tarAddr, tarErr := datastore_utils.FindAndFormatRef(
					tempDS,
					datastore.AddressRef{
						ChainSelector: input.ChainSelector,
						Type:          datastore.ContractType(token_admin_registry.ContractType),
						Version:       token_admin_registry.Version,
					},
					input.ChainSelector,
					evm_datastore_utils.ToEVMAddress,
				)
				if tarErr == nil {
					getTokenReport, getTokenErr := cldf_ops.ExecuteOperation(b, lombard_token_pool.GetToken, c, contract.FunctionInput[struct{}]{
						ChainSelector: input.ChainSelector,
						Address:       lombardPoolAddr,
					})
					if getTokenErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("error getting token from LombardTokenPool: %w", getTokenErr)
					}
					setPoolOut, setPoolErr := cldf_ops.ExecuteOperation(b, token_admin_registry.SetPool, c, contract.FunctionInput[token_admin_registry.SetPoolArgs]{
						ChainSelector: input.ChainSelector,
						Address:       tarAddr,
						Args: token_admin_registry.SetPoolArgs{
							TokenAddress:     getTokenReport.Output,
							TokenPoolAddress: lombardPoolAddr,
						},
					})
					if setPoolErr != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("error executing TokenAdminRegistry SetPool for LBTC: %w", setPoolErr)
					}
					writes = append(writes, setPoolOut.Output)
				}
			}

			// If USDCTokenPoolProxy 2.0.0 is in existingAddresses, update lockOrBurn mechanism to CCTP_V2_WITH_CCV for all remote chain selectors (same batch).
			usdcProxyAddr, usdcProxyErr := datastore_utils.FindAndFormatRef(
				tempDS,
				datastore.AddressRef{
					ChainSelector: input.ChainSelector,
					Type:          datastore.ContractType(usdc_token_pool_proxy.ContractType),
					Version:       usdc_token_pool_proxy.Version,
				},
				input.ChainSelector,
				evm_datastore_utils.ToEVMAddress,
			)
			if usdcProxyErr == nil && len(input.RemoteChainSelectors) > 0 {
				mechanisms := make([]uint8, len(input.RemoteChainSelectors))
				for i := range mechanisms {
					mechanisms[i] = 4 // CCTP_V2_WITH_CCV
				}
				updateMechOut, updateMechErr := cldf_ops.ExecuteOperation(b, usdc_token_pool_proxy.UpdateLockOrBurnMechanisms, c, contract.FunctionInput[usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs]{
					ChainSelector: input.ChainSelector,
					Address:       usdcProxyAddr,
					Args: usdc_token_pool_proxy.UpdateLockOrBurnMechanismsArgs{
						RemoteChainSelectors: input.RemoteChainSelectors,
						Mechanisms:           mechanisms,
					},
				})
				if updateMechErr != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("error executing USDCTokenPoolProxy UpdateLockOrBurnMechanisms: %w", updateMechErr)
				}
				writes = append(writes, updateMechOut.Output)
			}

			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("error creating batch operation: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})
}
