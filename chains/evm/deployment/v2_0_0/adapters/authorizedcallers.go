package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	api "github.com/smartcontractkit/chainlink-ccip/deployment/authorizedcallers"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	sequtil "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

const evmCallerLen = 20

// EVMAuthorizedCallersAdapter implements api.AuthorizedCallersAdapter for EVM chains.
// The adapter is not tied to any specific contract: it is parameterized at construction
// time with per-contract generated operations (applyOp, getAllOp) and a buildArgs
// helper so that MCMS batch metadata (ContractType, ABI label, method selector) comes
// from the real contract binding, not a generic shim.
//
// Registered adapters may be singletons (see init.go). Initialize must refresh the cached
// address from the provided Environment.DataStore every time — not skip when the
// (selector, type, version) key was seen before — so a new simulated chain / new deployment
// in another test does not reuse a stale RMN address.
//
// addrCache memoizes the resolved address within a single deployment flow after Initialize.
//
// GetAllAuthorizedCallers uses getAllOp with operations.NewBundle(..., NewMemoryReporter())
// on each call: deployment.Environment.OperationsBundle deduplicates ExecuteOperation reads
// (same op + same input hash returns the first successful cached result). For
// getAllAuthorizedCallers the input is always the same contract address on a chain, so
// without a fresh reporter the read would stay stuck at the first observed state.
type EVMAuthorizedCallersAdapter struct {
	addrCache map[string]common.Address
	getAllOp  *cldf_ops.Operation[contract.FunctionInput[struct{}], []common.Address, cldf_evm.Chain]
	// execApply executes the contract-specific applyAuthorizedCallerUpdates operation
	// through the ops bundle so MCMS metadata is accurate.
	execApply func(b cldf_ops.Bundle, chain cldf_evm.Chain, addr common.Address, added, removed []common.Address) (sequtil.OnChainOutput, error)
}

// NewEVMAuthorizedCallersAdapter constructs an EVMAuthorizedCallersAdapter backed by
// per-contract generated operations. applyOp must be the generated
// applyAuthorizedCallerUpdates write operation; getAllOp the generated
// getAllAuthorizedCallers read operation; buildArgs converts chain-agnostic address
// slices into the contract-specific ARGS struct.
//
// Example (adapters/init.go):
//
//	NewEVMAuthorizedCallersAdapter(
//	    rmnops.ApplyAuthorizedCallerUpdates,
//	    rmnops.GetAllAuthorizedCallers,
//	    func(added, removed []common.Address) rmnops.AuthorizedCallerArgs {
//	        return rmnops.AuthorizedCallerArgs{AddedCallers: added, RemovedCallers: removed}
//	    },
//	)
func NewEVMAuthorizedCallersAdapter[ARGS any](
	applyOp *cldf_ops.Operation[contract.FunctionInput[ARGS], contract.WriteOutput, cldf_evm.Chain],
	getAllOp *cldf_ops.Operation[contract.FunctionInput[struct{}], []common.Address, cldf_evm.Chain],
	buildArgs func(added, removed []common.Address) ARGS,
) *EVMAuthorizedCallersAdapter {
	return &EVMAuthorizedCallersAdapter{
		addrCache: make(map[string]common.Address),
		getAllOp:  getAllOp,
		execApply: func(b cldf_ops.Bundle, chain cldf_evm.Chain, addr common.Address, added, removed []common.Address) (sequtil.OnChainOutput, error) {
			report, err := cldf_ops.ExecuteOperation(b, applyOp, chain, contract.FunctionInput[ARGS]{
				ChainSelector: chain.Selector,
				Address:       addr,
				Args:          buildArgs(added, removed),
			})
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("applyAuthorizedCallerUpdates on %s: %w", addr.Hex(), err)
			}
			batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{report.Output})
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("failed to create batch from writes: %w", err)
			}
			return sequtil.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batch}}, nil
		},
	}
}

// Initialize resolves and caches the target contract address from e.DataStore for the given ApplyInput.
// Called for every changeset invocation; always re-reads the datastore so singleton adapters
// do not retain an RMN address from a prior Environment (e.g. another test's simulated chain).
// Must still be invoked before GetAllAuthorizedCallers or ApplyAuthorizedCallerUpdates for a given triple.
func (a *EVMAuthorizedCallersAdapter) Initialize(e cldf.Environment, in api.ApplyInput) error {
	if err := api.ValidateApplyInput(in); err != nil {
		return err
	}
	key := addrCacheKey(in.ChainSelector, in.ContractType, in.Version)
	ref := datastore.AddressRef{
		Type:    datastore.ContractType(in.ContractType),
		Version: in.Version,
	}
	addr, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, in.ChainSelector, evmds.ToEVMAddress)
	if err != nil {
		return fmt.Errorf(
			"failed to resolve %q v%s address on chain %d: %w",
			in.ContractType, in.Version.String(), in.ChainSelector, err)
	}
	a.addrCache[key] = addr
	return nil
}

// GetAllAuthorizedCallers reads the current set of authorized callers from the chain
// via the injected getAllOp. Returns each caller as a 20-byte EVM address slice.
func (a *EVMAuthorizedCallersAdapter) GetAllAuthorizedCallers(
	e cldf.Environment,
	selector uint64,
	contractType cldf.ContractType,
	version *semver.Version,
) ([]api.Caller, error) {
	addr, err := a.cachedAddr(selector, contractType, version)
	if err != nil {
		return nil, err
	}
	chain, ok := e.BlockChains.EVMChains()[selector]
	if !ok {
		return nil, fmt.Errorf("no EVM chain found for selector %d", selector)
	}
	readBundle := cldf_ops.NewBundle(e.GetContext, e.Logger, cldf_ops.NewMemoryReporter())
	report, err := cldf_ops.ExecuteOperation(readBundle, a.getAllOp, chain, contract.FunctionInput[struct{}]{
		ChainSelector: chain.Selector,
		Address:       addr,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("getAllAuthorizedCallers at %s on chain %d: %w", addr.Hex(), selector, err)
	}
	callers := make([]api.Caller, len(report.Output))
	for i, c := range report.Output {
		callers[i] = c.Bytes()
	}
	return callers, nil
}

// ApplyAuthorizedCallerUpdates returns the sequence that calls applyAuthorizedCallerUpdates
// on the target contract via the injected applyOp. Each api.Caller is decoded to a
// common.Address (exactly 20 bytes, non-zero).
func (a *EVMAuthorizedCallersAdapter) ApplyAuthorizedCallerUpdates() *cldf_ops.Sequence[api.ApplyInput, sequtil.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm:authorized-callers:apply-updates",
		semver.MustParse("1.0.0"),
		"Applies authorized caller updates on an AuthorizedCallers-inheriting EVM contract",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in api.ApplyInput) (sequtil.OnChainOutput, error) {
			if err := api.ValidateApplyInput(in); err != nil {
				return sequtil.OnChainOutput{}, err
			}
			chain, ok := chains.EVMChains()[in.ChainSelector]
			if !ok {
				return sequtil.OnChainOutput{}, fmt.Errorf("chain %d not found", in.ChainSelector)
			}
			addr, err := a.cachedAddr(in.ChainSelector, in.ContractType, in.Version)
			if err != nil {
				return sequtil.OnChainOutput{}, err
			}
			added, err := toEVMAddresses(in.Update.AddedCallers)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("invalid addedCallers: %w", err)
			}
			removed, err := toEVMAddresses(in.Update.RemovedCallers)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf("invalid removedCallers: %w", err)
			}
			out, err := a.execApply(b, chain, addr, added, removed)
			if err != nil {
				return sequtil.OnChainOutput{}, fmt.Errorf(
					"apply on %s (contract %q v%s) chain %d: %w",
					addr.Hex(), in.ContractType, in.Version.String(), in.ChainSelector, err)
			}
			return out, nil
		},
	)
}

func (a *EVMAuthorizedCallersAdapter) cachedAddr(
	selector uint64,
	contractType cldf.ContractType,
	version *semver.Version,
) (common.Address, error) {
	if version == nil {
		return common.Address{}, fmt.Errorf("no cached address lookup: version is nil for %q on chain %d", contractType, selector)
	}
	if contractType == "" {
		return common.Address{}, fmt.Errorf("no cached address lookup: contractType is empty on chain %d", selector)
	}
	key := addrCacheKey(selector, contractType, version)
	addr, ok := a.addrCache[key]
	if !ok {
		return common.Address{}, fmt.Errorf(
			"no cached address for %q v%s on chain %d; call Initialize first",
			contractType, version.String(), selector)
	}
	return addr, nil
}

func addrCacheKey(selector uint64, contractType cldf.ContractType, version *semver.Version) string {
	ver := ""
	if version != nil {
		ver = version.String()
	}
	return fmt.Sprintf("%d|%s|%s", selector, contractType, ver)
}

func toEVMAddresses(callers []api.Caller) ([]common.Address, error) {
	out := make([]common.Address, len(callers))
	for i, c := range callers {
		if len(c) != evmCallerLen {
			return nil, fmt.Errorf("caller at index %d has length %d, expected %d (EVM address)", i, len(c), evmCallerLen)
		}
		addr := common.BytesToAddress(c)
		if addr == (common.Address{}) {
			return nil, fmt.Errorf("caller at index %d is the zero address", i)
		}
		out[i] = addr
	}
	return out, nil
}
