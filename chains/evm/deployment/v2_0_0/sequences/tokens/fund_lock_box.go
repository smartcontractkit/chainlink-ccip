package tokens

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	lockbox_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	siloed_lrtp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

// LockBoxKind declares whether a lockbox is a silo (dedicated to one remote chain) or the shared
// (unsiloed) bucket. It is declared explicitly by the operator and verified on-chain against the
// pool's lockbox mapping before any write is emitted.
type LockBoxKind string

const (
	// LockBoxKindSiloed is a lockbox mapped to exactly one remote chain selector.
	LockBoxKindSiloed LockBoxKind = "siloed"
	// LockBoxKindUnsiloed is a lockbox shared by more than one remote chain selector.
	LockBoxKindUnsiloed LockBoxKind = "unsiloed"
)

// IsValid reports whether k is a recognised lockbox kind.
func (k LockBoxKind) IsValid() bool {
	return k == LockBoxKindSiloed || k == LockBoxKindUnsiloed
}

// LockBoxDeposit is a single deposit into a lockbox bucket.
type LockBoxDeposit struct {
	// RemoteChainSelector identifies the bucket to fund. Zero designates the unsiloed (shared)
	// bucket; any other value designates the silo for that remote chain.
	RemoteChainSelector uint64
	// Amount is the amount to deposit, in raw base units.
	Amount *big.Int
}

// FundLockBoxInput is the input for the lockbox funding sequence.
type FundLockBoxInput struct {
	ChainSelector uint64
	// PoolAddress is the v2.0 SiloedLockReleaseTokenPool that maps LockBoxAddress. Required: the
	// siloed/unsiloed distinction is defined by the pool's lockbox mapping, so the pool is read to
	// verify Kind on-chain.
	PoolAddress string
	// LockBoxAddress is the ERC20LockBox to fund.
	LockBoxAddress string
	// Kind declares whether LockBoxAddress is a silo or the shared (unsiloed) bucket. Verified
	// on-chain against the pool's lockbox mapping before any write is emitted.
	Kind LockBoxKind
	// TokenAddress is the token to deposit into the lockbox.
	TokenAddress string
	// TimelockAddress is the MCMS timelock address that will execute the funding operations.
	// Required because the timelock must be an authorized caller on the lockbox to deposit.
	TimelockAddress string
	// Deposits are the individual deposits to make into the lockbox. Each entry targets one bucket:
	// a siloed bucket (RemoteChainSelector set to the remote chain) or the unsiloed shared bucket
	// (RemoteChainSelector zero).
	Deposits []LockBoxDeposit
	// UsePlainTransfer, when true, transfers tokens directly to the lockbox via ERC20.transfer
	// instead of using the lockbox's deposit() function. This bypasses the Deposit event emission
	// and the per-bucket accounting. Use only as a break-glass option.
	UsePlainTransfer bool
}

// FundLockBox deposits tokens into a v2.0 ERC20LockBox, funding either a siloed bucket (keyed by
// remote chain selector) or the unsiloed (shared) bucket. It is the standalone counterpart to the
// funding step inside MigrateLockReleasePoolLiquidity: it requires no legacy pool to migrate from,
// so an operator can top up a lockbox directly.
//
// All writes are proposal-only: the timelock must be an authorized caller on the lockbox to
// deposit, so the sequence authorizes it (idempotently) and then approves + deposits. The tokens
// are expected to be held by the timelock, since the deposit pulls from the caller's balance.
var FundLockBox = cldf_ops.NewSequence(
	"fund-lock-box",
	semver.MustParse("2.0.0"),
	"Deposits tokens into a v2.0 ERC20LockBox (siloed and/or unsiloed buckets)",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input FundLockBoxInput) (sequences.OnChainOutput, error) {
		evmChain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		if err := validateFundLockBoxInput(input); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid fund lockbox input: %w", err)
		}

		poolAddr := common.HexToAddress(input.PoolAddress)
		lockBoxAddr := common.HexToAddress(input.LockBoxAddress)
		tokenAddr := common.HexToAddress(input.TokenAddress)
		timelockAddr := common.HexToAddress(input.TimelockAddress)

		mappedSelectors, err := verifyLockBoxKind(b, evmChain, input.ChainSelector, poolAddr, lockBoxAddr, input.Kind)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		if input.Kind == LockBoxKindSiloed {
			for _, deposit := range input.Deposits {
				if deposit.RemoteChainSelector != mappedSelectors[0] {
					return sequences.OnChainOutput{}, fmt.Errorf(
						"deposit targets remote chain %d but silo lockbox %s on pool %s is mapped to chain %d",
						deposit.RemoteChainSelector, lockBoxAddr, poolAddr, mappedSelectors[0])
				}
			}
		}

		var ops []evm_contract.WriteOutput

		// The deposit path pulls tokens from the caller, so the timelock must be an authorized
		// caller on the lockbox. The authorize step is idempotent: only append it when the timelock
		// isn't already authorized, avoiding a redundant MCMS batch op and a spurious
		// AuthorizedCallerAdded event on every subsequent top-up.
		if !input.UsePlainTransfer {
			authCallers, err := cldf_ops.ExecuteOperation(
				b,
				lockbox_ops.GetAllAuthorizedCallers,
				evmChain,
				evm_contract.FunctionInput[struct{}]{
					ChainSelector: input.ChainSelector,
					Address:       lockBoxAddr,
				},
				cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers on lockbox %s: %w", lockBoxAddr, err)
			}
			if !slices.Contains(authCallers.Output, timelockAddr) {
				addAuthReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.ApplyAuthorizedCallerUpdates, evmChain, evm_contract.FunctionInput[lockbox_ops.AuthorizedCallerArgs]{
					ChainSelector: input.ChainSelector,
					Address:       lockBoxAddr,
					Args: lockbox_ops.AuthorizedCallerArgs{
						AddedCallers:   []common.Address{timelockAddr},
						RemovedCallers: []common.Address{},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to add timelock as authorized caller on lockbox %s: %w", lockBoxAddr, err)
				}
				ops = append(ops, addAuthReport.Output)
			}
		}

		for _, deposit := range input.Deposits {
			if input.UsePlainTransfer {
				transferReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.TransferProposalOnly, evmChain, evm_contract.FunctionInput[erc20_ops.TransferArgs]{
					ChainSelector: input.ChainSelector,
					Address:       tokenAddr,
					Args: erc20_ops.TransferArgs{
						Receiver: lockBoxAddr,
						Amount:   deposit.Amount,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer tokens to lockbox %s: %w", lockBoxAddr, err)
				}
				ops = append(ops, transferReport.Output)
				continue
			}

			approveReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.ApproveProposalOnly, evmChain, evm_contract.FunctionInput[erc20_ops.ApproveArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenAddr,
				Args: erc20_ops.ApproveArgs{
					Spender: lockBoxAddr,
					Value:   deposit.Amount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to approve lockbox %s to spend tokens: %w", lockBoxAddr, err)
			}
			ops = append(ops, approveReport.Output)

			depositReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.DepositProposalOnly, evmChain, evm_contract.FunctionInput[lockbox_ops.DepositArgs]{
				ChainSelector: input.ChainSelector,
				Address:       lockBoxAddr,
				Args: lockbox_ops.DepositArgs{
					Token:               tokenAddr,
					RemoteChainSelector: deposit.RemoteChainSelector,
					Amount:              deposit.Amount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deposit into lockbox %s: %w", lockBoxAddr, err)
			}
			ops = append(ops, depositReport.Output)
		}

		batchOp, err := evm_contract.NewBatchOperationFromWrites(ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

// verifyLockBoxKind reads the pool's lockbox mapping and verifies that lockBoxAddr matches the
// declared kind, returning the remote chain selectors the lockbox is mapped to (sorted).
//
// This is the only on-chain definition of the siloed/unsiloed distinction: the ERC20LockBox
// contract has no silo concept and ignores the remoteChainSelector passed to deposit(). The
// distinction lives entirely in the pool's s_lockBoxes mapping (configureLockBoxes), where a silo
// is mapped to exactly one remote chain selector and the shared (unsiloed) bucket is mapped to
// more than one.
func verifyLockBoxKind(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSelector uint64,
	poolAddr, lockBoxAddr common.Address,
	kind LockBoxKind,
) ([]uint64, error) {
	configsReport, err := cldf_ops.ExecuteOperation(
		b,
		siloed_lrtp_ops.GetAllLockBoxConfigs,
		evmChain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       poolAddr,
		},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to read lockbox configs from pool %s: %w", poolAddr, err)
	}

	var mappedSelectors []uint64
	for _, config := range configsReport.Output {
		if config.LockBox == lockBoxAddr {
			mappedSelectors = append(mappedSelectors, config.RemoteChainSelector)
		}
	}
	if len(mappedSelectors) == 0 {
		return nil, fmt.Errorf(
			"lockbox %s is not configured on pool %s; configure it with configureLockBoxes before funding",
			lockBoxAddr, poolAddr)
	}
	slices.Sort(mappedSelectors)

	switch kind {
	case LockBoxKindSiloed:
		if len(mappedSelectors) != 1 {
			return nil, fmt.Errorf(
				"lockbox %s is declared siloed but pool %s maps it to %d chain selectors %v; a silo is mapped to exactly one",
				lockBoxAddr, poolAddr, len(mappedSelectors), mappedSelectors)
		}
	case LockBoxKindUnsiloed:
		if len(mappedSelectors) < 2 {
			return nil, fmt.Errorf(
				"lockbox %s is declared unsiloed but pool %s maps it to only %d chain selector(s) %v; the shared bucket is mapped to more than one",
				lockBoxAddr, poolAddr, len(mappedSelectors), mappedSelectors)
		}
	default:
		return nil, fmt.Errorf("unknown lockbox kind %q", kind)
	}

	return mappedSelectors, nil
}

func validateFundLockBoxInput(input FundLockBoxInput) error {
	if input.PoolAddress == "" {
		return fmt.Errorf("PoolAddress must be provided")
	}
	if input.LockBoxAddress == "" {
		return fmt.Errorf("LockBoxAddress must be provided")
	}
	if !input.Kind.IsValid() {
		return fmt.Errorf("Kind must be %q or %q, got %q", LockBoxKindSiloed, LockBoxKindUnsiloed, input.Kind)
	}
	if input.TokenAddress == "" {
		return fmt.Errorf("TokenAddress must be provided")
	}
	if input.TimelockAddress == "" {
		return fmt.Errorf("TimelockAddress must be provided")
	}
	if len(input.Deposits) == 0 {
		return fmt.Errorf("at least one deposit is required")
	}

	seen := make(map[uint64]bool, len(input.Deposits))
	for i, deposit := range input.Deposits {
		if deposit.Amount == nil || deposit.Amount.Sign() <= 0 {
			return fmt.Errorf("deposits[%d]: Amount must be positive", i)
		}
		if seen[deposit.RemoteChainSelector] {
			return fmt.Errorf("deposits[%d]: duplicate RemoteChainSelector %d", i, deposit.RemoteChainSelector)
		}
		seen[deposit.RemoteChainSelector] = true
	}

	return nil
}
