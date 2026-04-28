package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	onrampOpsV15 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	onrampOpsV16 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	onrampOpsV20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	routerbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// EVMFeeContractResolver implements fees.FeeContractResolver for EVM chains by
// reading the live Router on the source chain to infer the on-ramp version.
type EVMFeeContractResolver struct{}

// ResolveFeeContractRef discovers the AddressRef of the contract that holds
// token-transfer fee config for a given (src, dst) lane, by:
//  1. Loading the v1.2.0 Router for src from the datastore.
//  2. Calling Router.getOnRamp(dst).
//  3. Reverse-looking-up the returned on-ramp address in the datastore for its
//     Type and Version.
//  4. For an EVM2EVMOnRamp (v1.5) the on-ramp itself is the fee contract.
//     For an OnRamp (v1.6+) the FeeQuoter — reachable via getDynamicConfig — is
//     the fee contract; the v1.6 and v2.0 OnRamp ABIs both expose a FeeQuoter
//     field on their dynamic config but the generated bindings differ, so we
//     dispatch on the on-ramp's major version.
//
// The Router (v1.2.0) is a stable hub that maps dst chain selectors to live
// on-ramp addresses, so the operator does not need to know which CCIP version
// each lane is on.
func (EVMFeeContractResolver) ResolveFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
	ds := e.DataStore

	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("EVM chain with selector %d not found", src)
	}

	routerRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(routerops.ContractType),
		Version: routerops.Version,
	}, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to find Router (v%s) for src %d: %w", routerops.Version.String(), src, err)
	}
	if !common.IsHexAddress(routerRef.Address) {
		return datastore.AddressRef{}, fmt.Errorf("invalid Router address %q for src %d", routerRef.Address, src)
	}

	routerContract, err := routerbind.NewRouter(common.HexToAddress(routerRef.Address), chain.Client)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to bind Router at %s on src %d: %w", routerRef.Address, src, err)
	}

	onRampAddr, err := routerContract.GetOnRamp(&bind.CallOpts{Context: e.GetContext()}, dst)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to call Router.getOnRamp(dst=%d) on src %d at %s: %w", dst, src, routerRef.Address, err)
	}
	if onRampAddr == (common.Address{}) {
		return datastore.AddressRef{}, fmt.Errorf("Router.getOnRamp(dst=%d) on src %d returned the zero address (no live lane)", dst, src)
	}

	onRampRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Address: onRampAddr.Hex(),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("on-ramp address %s returned by Router.getOnRamp(dst=%d) on src %d is not present in the datastore: %w", onRampAddr.Hex(), dst, src, err)
	}

	switch onRampRef.Type {
	case datastore.ContractType(onrampOpsV15.ContractType):
		return onRampRef, nil

	case datastore.ContractType(onrampOpsV16.ContractType):
		if onRampRef.Version == nil {
			return datastore.AddressRef{}, fmt.Errorf("on-ramp at %s on src %d has no Version metadata in datastore", onRampAddr.Hex(), src)
		}

		var fqAddr common.Address
		switch onRampRef.Version.Major() {
		case 1:
			c, err := onrampOpsV16.NewOnRampContract(onRampAddr, chain.Client)
			if err != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to bind v1.6 OnRamp at %s on src %d: %w", onRampAddr.Hex(), src, err)
			}
			cfg, err := c.GetDynamicConfig(&bind.CallOpts{Context: e.GetContext()})
			if err != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to call v1.6 OnRamp.getDynamicConfig at %s on src %d: %w", onRampAddr.Hex(), src, err)
			}
			fqAddr = cfg.FeeQuoter
		case 2:
			c, err := onrampOpsV20.NewOnRampContract(onRampAddr, chain.Client)
			if err != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to bind v2.0 OnRamp at %s on src %d: %w", onRampAddr.Hex(), src, err)
			}
			cfg, err := c.GetDynamicConfig(&bind.CallOpts{Context: e.GetContext()})
			if err != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to call v2.0 OnRamp.getDynamicConfig at %s on src %d: %w", onRampAddr.Hex(), src, err)
			}
			fqAddr = cfg.FeeQuoter
		default:
			return datastore.AddressRef{}, fmt.Errorf("unsupported OnRamp major version %d (%s) at %s on src %d", onRampRef.Version.Major(), onRampRef.Version.String(), onRampAddr.Hex(), src)
		}

		if fqAddr == (common.Address{}) {
			return datastore.AddressRef{}, fmt.Errorf("FeeQuoter address is zero in OnRamp.getDynamicConfig at %s on src %d", onRampAddr.Hex(), src)
		}

		fqRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
			Address: fqAddr.Hex(),
		}, src, datastore_utils.FullRef)
		if err != nil {
			return datastore.AddressRef{}, fmt.Errorf("FeeQuoter address %s reported by OnRamp at %s on src %d is not present in the datastore: %w", fqAddr.Hex(), onRampAddr.Hex(), src, err)
		}
		return fqRef, nil

	default:
		return datastore.AddressRef{}, fmt.Errorf("unsupported on-ramp type %q at address %s on src %d", onRampRef.Type, onRampAddr.Hex(), src)
	}
}
