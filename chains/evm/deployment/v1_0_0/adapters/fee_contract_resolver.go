package adapters

import (
	"context"
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	routerbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// OnRampFeeContractOps abstracts the version-specific call that, given an
// on-ramp address, returns the address of the contract that holds
// token-transfer fee config.
//
// For v1.5 (EVM2EVMOnRamp) the on-ramp is itself the fee contract.
// For v1.6 and v2.0 (OnRamp) the FeeQuoter is the fee contract, reachable via
// OnRamp.getDynamicConfig(). The two OnRamp versions expose the same field
// name on different generated DynamicConfig structs, so each version supplies
// its own implementation that imports its own bindings.
type OnRampFeeContractOps interface {
	GetFeeContractAddress(ctx context.Context, chain evm.Chain, onRampAddr common.Address) (common.Address, error)
}

type onRampOpsKey struct {
	Type    datastore.ContractType
	Version string
}

func newOnRampOpsKey(t datastore.ContractType, v *semver.Version) onRampOpsKey {
	return onRampOpsKey{Type: t, Version: cciputils.StripPatchVersion(v).String()}
}

// EVMFeeContractResolver implements fees.FeeContractResolver for EVM lanes by
// reading the live Router on the source chain. Per-version on-ramp logic is
// supplied via RegisterOnRampOps so this file imports only the v1.2.0 Router
// and never the per-version on-ramp packages.
type EVMFeeContractResolver struct {
	mu      sync.RWMutex
	onRamps map[onRampOpsKey]OnRampFeeContractOps
}

func newEVMFeeContractResolver() *EVMFeeContractResolver {
	return &EVMFeeContractResolver{onRamps: make(map[onRampOpsKey]OnRampFeeContractOps)}
}

// RegisterOnRampOps wires a per-version OnRampFeeContractOps into the resolver.
// Patch components of the version are stripped before keying, so 1.6.x all map
// to the same Ops. First registration wins.
func (r *EVMFeeContractResolver) RegisterOnRampOps(t datastore.ContractType, v *semver.Version, ops OnRampFeeContractOps) {
	key := newOnRampOpsKey(t, v)
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.onRamps[key]; !exists {
		r.onRamps[key] = ops
	}
}

func (r *EVMFeeContractResolver) ResolveFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
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

	rc, err := routerbind.NewRouter(common.HexToAddress(routerRef.Address), chain.Client)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to bind Router at %s on src %d: %w", routerRef.Address, src, err)
	}

	onRampAddr, err := rc.GetOnRamp(&bind.CallOpts{Context: e.GetContext()}, dst)
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
	if onRampRef.Version == nil {
		return datastore.AddressRef{}, fmt.Errorf("on-ramp at %s on src %d has no Version metadata in datastore", onRampAddr.Hex(), src)
	}

	key := newOnRampOpsKey(onRampRef.Type, onRampRef.Version)
	r.mu.RLock()
	ops, ok := r.onRamps[key]
	r.mu.RUnlock()
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("no OnRampFeeContractOps registered for type=%q version=%s", onRampRef.Type, onRampRef.Version.String())
	}

	feeContractAddr, err := ops.GetFeeContractAddress(e.GetContext(), chain, onRampAddr)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to resolve fee-contract address for on-ramp %s on src %d: %w", onRampAddr.Hex(), src, err)
	}
	if feeContractAddr == (common.Address{}) {
		return datastore.AddressRef{}, fmt.Errorf("OnRampFeeContractOps returned the zero address for on-ramp %s on src %d", onRampAddr.Hex(), src)
	}

	if feeContractAddr == onRampAddr {
		return onRampRef, nil
	}

	feeRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Address: feeContractAddr.Hex(),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("fee-contract address %s reported by OnRamp at %s on src %d is not present in the datastore: %w", feeContractAddr.Hex(), onRampAddr.Hex(), src, err)
	}
	return feeRef, nil
}

var (
	evmFeeContractResolverOnce sync.Once
	evmFeeContractResolver     *EVMFeeContractResolver
)

// GetEVMFeeContractResolver returns the singleton EVMFeeContractResolver. The
// per-version /adapters packages register their OnRampFeeContractOps into this
// instance from their own init() functions.
func GetEVMFeeContractResolver() *EVMFeeContractResolver {
	evmFeeContractResolverOnce.Do(func() {
		evmFeeContractResolver = newEVMFeeContractResolver()
	})
	return evmFeeContractResolver
}
