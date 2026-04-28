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
	feesapi "github.com/smartcontractkit/chainlink-ccip/deployment/fees"
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
	if t == "" {
		panic("RegisterOnRampOps: empty ContractType")
	}
	if v == nil {
		panic(fmt.Sprintf("RegisterOnRampOps: nil version for type=%q", t))
	}
	if ops == nil {
		panic(fmt.Sprintf("RegisterOnRampOps: nil ops for type=%q version=%s", t, v.String()))
	}
	key := newOnRampOpsKey(t, v)
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.onRamps[key]; !exists {
		r.onRamps[key] = ops
	}
}

// ResolveFeeContractRef returns the AddressRef of the contract that holds
// token-transfer fee config for the (src, dst) lane. Callers select the
// per-version FeeAdapter from the returned AddressRef.Version after
// StripPatchVersion; the returned Address is informational because each
// adapter re-derives its own write target.
func (r *EVMFeeContractResolver) ResolveFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error) {
	ds := e.DataStore

	chain, ok := e.BlockChains.EVMChains()[src]
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf("EVM chain with selector %d not found", src)
	}
	if chain.Client == nil {
		return datastore.AddressRef{}, fmt.Errorf("EVM chain %d has nil Client; cannot read live Router state", src)
	}

	routerRef, err := findRouterRef(ds, src)
	if err != nil {
		return datastore.AddressRef{}, err
	}
	if !common.IsHexAddress(routerRef.Address) {
		return datastore.AddressRef{}, fmt.Errorf("invalid Router address %q for src %d", routerRef.Address, src)
	}
	e.Logger.Infof("EVMFeeContractResolver: src=%d using %s at %s (v%s) to resolve OnRamp for dst=%d", src, routerRef.Type, routerRef.Address, routerRef.Version.String(), dst)

	rc, err := routerbind.NewRouter(common.HexToAddress(routerRef.Address), chain.Client)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to bind Router at %s on src %d: %w", routerRef.Address, src, err)
	}

	onRampAddr, err := rc.GetOnRamp(&bind.CallOpts{Context: e.GetContext()}, dst)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to call Router.GetOnRamp(dst=%d) on src %d at %s: %w", dst, src, routerRef.Address, err)
	}
	if onRampAddr == (common.Address{}) {
		return datastore.AddressRef{}, fmt.Errorf("Router.GetOnRamp(dst=%d) on src %d at %s returned the zero address: %w", dst, src, routerRef.Address, feesapi.ErrNoLiveLane)
	}

	onRampRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Address: onRampAddr.Hex(),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("OnRamp address %s returned by Router.GetOnRamp(dst=%d) on src %d is not present in the datastore: %w", onRampAddr.Hex(), dst, src, err)
	}
	if onRampRef.Version == nil {
		return datastore.AddressRef{}, fmt.Errorf("OnRamp at %s on src %d has no Version metadata in datastore", onRampAddr.Hex(), src)
	}

	key := newOnRampOpsKey(onRampRef.Type, onRampRef.Version)
	r.mu.RLock()
	ops, ok := r.onRamps[key]
	r.mu.RUnlock()
	if !ok {
		return datastore.AddressRef{}, fmt.Errorf(
			"no OnRampFeeContractOps registered for type=%q version=%s (lookup key %s) at OnRamp %s on src %d, dst %d",
			onRampRef.Type,
			onRampRef.Version.String(),
			cciputils.StripPatchVersion(onRampRef.Version).String(),
			onRampAddr.Hex(),
			src,
			dst,
		)
	}

	feeContractAddr, err := ops.GetFeeContractAddress(e.GetContext(), chain, onRampAddr)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to resolve fee-contract address for OnRamp %s on src %d: %w", onRampAddr.Hex(), src, err)
	}
	if feeContractAddr == (common.Address{}) {
		return datastore.AddressRef{}, fmt.Errorf("OnRampFeeContractOps returned the zero address for OnRamp %s on src %d", onRampAddr.Hex(), src)
	}

	// v1.5 short-circuit: the EVM2EVMOnRamp itself holds fee config, so the
	// onramp ref already points at the fee contract. Skip the second datastore
	// lookup; the returned ref's Type will be the OnRamp's type
	// (e.g. EVM2EVMOnRamp), which is intentional for v1.5.
	if feeContractAddr == onRampAddr {
		return onRampRef, nil
	}

	feeRef, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(cciputils.FeeQuoter),
		Address: feeContractAddr.Hex(),
	}, src, datastore_utils.FullRef)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("FeeQuoter address %s reported by OnRamp at %s on src %d is not present in the datastore (filtered by Type=FeeQuoter): %w", feeContractAddr.Hex(), onRampAddr.Hex(), src, err)
	}
	if feeRef.Version == nil {
		return datastore.AddressRef{}, fmt.Errorf("FeeQuoter at %s on src %d (reported by OnRamp %s) has no Version metadata in datastore", feeContractAddr.Hex(), src, onRampAddr.Hex())
	}
	return feeRef, nil
}

// findRouterRef looks up the active Router for src in the datastore, falling
// back to TestRouter when the production Router is not registered. Production
// is preferred when both exist; this matches the broader codebase pattern of
// keeping Router and TestRouter as separate ContractTypes that callers select
// between explicitly (see v1_6_0/sequences/adapter.go GetRouter / GetTestRouter).
func findRouterRef(ds datastore.DataStore, src uint64) (datastore.AddressRef, error) {
	ref, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(routerops.ContractType),
		Version: routerops.Version,
	}, src, datastore_utils.FullRef)
	if err == nil {
		return ref, nil
	}
	testRef, testErr := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		Type:    datastore.ContractType(routerops.TestRouterContractType),
		Version: routerops.Version,
	}, src, datastore_utils.FullRef)
	if testErr == nil {
		return testRef, nil
	}
	return datastore.AddressRef{}, fmt.Errorf("no Router or TestRouter (v%s) for src %d: router lookup error: %w; testRouter lookup error: %v", routerops.Version.String(), src, err, testErr)
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
