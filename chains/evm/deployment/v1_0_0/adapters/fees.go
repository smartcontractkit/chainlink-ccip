package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

var _ fees.FeeResolver = (*EVMFeeResolver)(nil)

type EVMFeeResolver struct{}

func (a *EVMFeeResolver) GetOnRampRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
	srcChain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	// NOTE: for EVM, the router is rarely re-deployed and currently serves as the canonical source for the OnRamp
	// address. Therefore, it is not necessary for the user to know the OnRamp version when setting token transfer
	// fee configs between some source and destination chains - we can infer this purely from on-chain state.
	routerAddr, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:          datastore.ContractType(routerops.ContractType),
			Version:       routerops.Version,
			ChainSelector: src,
		},
		src,
		evm_datastore_utils.ToEVMAddress,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get Router address for chain selector %d: %w", src, err)
	}

	// NOTE: the returned OnRamp address could be either v1.5.0 or v1.6.0 - we make no assumptions about the
	// version being returned and do not import any version-specific code here to avoid any potential issues
	e.Logger.Infof("Found Router address %s for chain %s:", routerAddr.Hex(), srcChain.String())
	getOnRampReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		routerops.GetOnRamp,
		srcChain,
		contract.FunctionInput[uint64]{
			ChainSelector: src,
			Address:       routerAddr,
			Args:          dst,
		},
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to execute GetOnRamp operation for Router at %s on chain selector %d with dst %d: %w", routerAddr.Hex(), src, dst, err)
	}

	onRampAddr := getOnRampReport.Output
	if onRampAddr == (common.Address{}) {
		return datastore.AddressRef{}, fmt.Errorf("no OnRamp contract found for src %d and dst %d", src, dst)
	}

	// NOTE: the OnRamp ContractType varies between v1.5 and v1.6. For v1.5 it's `EVM2EVMOnRamp` and
	// for v1.6 it's `OnRamp`, so we should avoid filtering on the ContractType otherwise the search
	// may be too restrictive and this call will incorrectly fail
	e.Logger.Infof("Found OnRamp address %s for src %d and dst %d", onRampAddr.Hex(), src, dst)
	onRampRef, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{ChainSelector: src, Address: onRampAddr.Hex()},
		src,
		datastore_utils.FullRef,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find OnRamp address ref for address %s on chain selector %d: %w", getOnRampReport.Output.Hex(), src, err)
	}

	e.Logger.Infof(
		"OnRamp ref for src %d and dst %d: %s",
		src, dst, datastore_utils.SprintRef(onRampRef),
	)

	return onRampRef, nil
}
