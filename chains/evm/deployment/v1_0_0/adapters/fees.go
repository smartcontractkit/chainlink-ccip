package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

var _ fees.FeeResolver = (*EVMFeeResolver)(nil)

type EVMFeeResolver struct{}

func (a *EVMFeeResolver) GetOnRampRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, src uint64, dst uint64) (datastore.AddressRef, error) {
	srcChain, ok := chains.EVMChains()[src]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("chain with selector %d not defined", src)
	}

	routerAddr, err := datastore_utils.FindAndFormatRef(
		ds,
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

	b.Logger.Infof("Found Router address %s for chain %s:", routerAddr.Hex(), srcChain.String())
	getOnRampReport, err := evmops.ExecuteRead(
		b,
		srcChain,
		routerAddr,
		evmops.BindAs[router.RouterInterface](router.NewRouter),
		routerops.NewReadGetOnRamp,
		dst,
	)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to execute GetOnRamp operation for Router at %s on chain selector %d with dst %d: %w", routerAddr.Hex(), src, dst, err)
	}

	onRampAddr := getOnRampReport.Output
	if onRampAddr == (common.Address{}) {
		return datastore.AddressRef{}, fmt.Errorf("no OnRamp contract found for src %d and dst %d", src, dst)
	}

	// NOTE: the OnRamp ContractType varies between v1.5 and v1.6. For v1.5 it's `EVM2EVMOnRamp` and
	// for v1.6 it's `OnRamp`, so we use `typeAndVersion()` to resolve the type and version directly
	// from on-chain data rather than relying on the datastore.
	b.Logger.Infof("Found OnRamp address %s for src %d and dst %d", onRampAddr.Hex(), src, dst)
	tvReport, err := evmops.ExecuteRead(b, srcChain, onRampAddr, type_and_version.NewTypeAndVersionContract, type_and_version.NewReadGetTypeAndVersion, struct{}{})
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to get typeAndVersion for OnRamp at %s on chain selector %d: %w", onRampAddr.Hex(), src, err)
	}

	onRampRef := datastore.AddressRef{
		ChainSelector: src,
		Qualifier:     cciputils.DefaultQualifier(onRampAddr.Hex(), tvReport.Output.Type),
		Type:          datastore.ContractType(tvReport.Output.Type),
		Version:       tvReport.Output.Version,
		Address:       onRampAddr.Hex(),
	}
	b.Logger.Infof(
		"OnRamp ref for src %d and dst %d: %s",
		src, dst, datastore_utils.SprintRef(onRampRef),
	)

	return onRampRef, nil
}
