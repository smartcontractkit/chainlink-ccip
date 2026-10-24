package tokens

import (
	"bytes"
	"errors"
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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	lockbox_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	siloed_lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

// MigrateLockReleasePoolLiquiditySink is the SINK half of a lock-release liquidity migration. It funds
// a v2.0.0 lockbox-based pool's lockbox(es) from the timelock with the amounts a version-specific
// SOURCE already extracted (see tokens.TokenAdapter.MigrateLockReleasePoolLiquiditySequence). It never
// reads the old pool or computes basis points: Amount (non-siloed) or SiloExactAmounts/
// UnsiloedExactAmount (siloed) carry the resolved amounts.
var MigrateLockReleasePoolLiquiditySink = cldf_ops.NewSequence(
	"migrate-lock-release-pool-liquidity-sink",
	semver.MustParse("2.0.0"),
	"Funds a v2.0.0 lockbox-based pool's lockbox(es) from the timelock with liquidity extracted by a source",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
		evmChain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		if err := validateSinkInput(input); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid sink input: %w", err)
		}

		newPoolAddr := common.HexToAddress(input.NewPoolAddress)
		timelockAddr := common.HexToAddress(input.TimelockAddress)

		newTVReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       newPoolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get typeAndVersion from new pool %s: %w", newPoolAddr, err)
		}
		if !utils.IsLockReleasePoolType(string(newTVReport.Output.Type)) || !newTVReport.Output.Version.Equal(utils.Version_2_0_0) {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"new pool %s is a %s %s, not a %s lock-release pool; the migration destination must be a v2.0.0 lockbox-based pool",
				newPoolAddr, newTVReport.Output.Type, newTVReport.Output.Version, utils.Version_2_0_0)
		}

		tokenReport, err := cldf_ops.ExecuteOperation(b, token_pool_ops.GetToken, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       newPoolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from new pool %s: %w", newPoolAddr, err)
		}
		tokenAddr := tokenReport.Output

		// The setPool step uses the caller-supplied TokenRef, which is not otherwise tied to the
		// pools. TokenAdminRegistry.setPool would reject a pool that does not support the token
		// (InvalidTokenPoolToken), but the extraction, funding, and setPool writes are merged into
		// one batch, so that revert would take the whole batch down at timelock execution and leave
		// a dead proposal. Fail at plan time instead.
		if input.SetPoolConfig != nil {
			if setPoolToken := common.HexToAddress(input.SetPoolConfig.TokenAddress); setPoolToken != tokenAddr {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"SetPoolConfig.TokenAddress %s does not match the token managed by the pools %s; setPool would revert with InvalidTokenPoolToken and take the whole migration batch with it",
					setPoolToken, tokenAddr,
				)
			}
		}

		// The resolved amounts must match the new pool's shape; otherwise the extracted liquidity would
		// have no lockbox to land in.
		if string(newTVReport.Output.Type) == utils.SiloedLockReleaseTokenPool.String() {
			if input.Amount != nil {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"new pool %s is siloed but the resolved amounts are non-siloed; a siloed pool must be funded per silo", newPoolAddr)
			}
			return fundSiloedPool(b, evmChain, input, newPoolAddr, tokenAddr, timelockAddr)
		}
		if input.Amount == nil {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"new pool %s is not siloed but the resolved amounts are per silo; a non-siloed pool must be funded with a single amount", newPoolAddr)
		}
		return fundUnsiloedPool(b, evmChain, input, newPoolAddr, tokenAddr, timelockAddr)
	},
)

// validateSinkInput checks that the caller passed the addresses and the amounts a source resolved.
func validateSinkInput(input tokens.MigrateLockReleasePoolLiquidityInput) error {
	if input.NewPoolAddress == "" {
		return errors.New("NewPoolAddress must be provided")
	}
	if input.TimelockAddress == "" {
		return errors.New("TimelockAddress must be provided")
	}
	if input.Amount == nil && len(input.SiloExactAmounts) == 0 && input.UnsiloedExactAmount == nil {
		return errors.New("resolved amounts must be provided: Amount, or SiloExactAmounts/UnsiloedExactAmount")
	}
	if input.Amount != nil && (len(input.SiloExactAmounts) > 0 || input.UnsiloedExactAmount != nil) {
		return errors.New("Amount is mutually exclusive with SiloExactAmounts/UnsiloedExactAmount")
	}
	if input.BasisPoints != nil {
		return errors.New("BasisPoints must be resolved to an amount by the source")
	}
	return nil
}

// resolveUnsiloedLockBox validates the supplied destination in the input for the unsiloed (shared)
// balance against the lockboxes actually mapped on the new pool.
//
// An empty address is allowed here and reported later, but only if there turns out to be unsiloed
// liquidity to move - a siloed pool with no shared balance does not need a destination.
func resolveUnsiloedLockBox(
	address string,
	configuredLockBoxes map[common.Address]bool,
	newPoolAddr common.Address,
) (common.Address, error) {
	if address == "" {
		return common.Address{}, nil
	}

	lockBox := common.HexToAddress(address)
	if lockBox == (common.Address{}) {
		return common.Address{}, fmt.Errorf("UnsiloedLockBoxAddress %q is not a valid address", address)
	}

	if !configuredLockBoxes[lockBox] {
		return common.Address{}, fmt.Errorf(
			"UnsiloedLockBoxAddress %s is not one of the lockboxes configured on new pool %s; configured lockboxes are %v",
			lockBox, newPoolAddr, sortedAddresses(configuredLockBoxes),
		)
	}

	return lockBox, nil
}

// sortedAddresses returns the set in a stable order so error messages are reproducible.
func sortedAddresses(set map[common.Address]bool) []common.Address {
	out := make([]common.Address, 0, len(set))

	for addr := range set {
		out = append(out, addr)
	}
	slices.SortFunc(out, func(a, b common.Address) int { return bytes.Compare(a.Bytes(), b.Bytes()) })

	return out
}

func fundUnsiloedPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokens.MigrateLockReleasePoolLiquidityInput,
	newPoolAddr, tokenAddr, timelockAddr common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	var ops []evm_contract.WriteOutput

	if input.Amount.Sign() <= 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("resolved migration amount must be positive, got %s", input.Amount)
	}

	lockboxReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v170.GetLockBox, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       newPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get lockbox from new pool %s: %w", newPoolAddr, err)
	}

	ops, err = appendFundingOps(b, evmChain, chainSel, lockboxReport.Output, tokenAddr, timelockAddr, input.Amount, 0, input.UsePlainTransfer, ops)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	if input.SetPoolConfig != nil {
		ops, err = appendSetPool(b, evmChain, chainSel, input.SetPoolConfig, newPoolAddr, ops)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
	}

	batchOp, err := evm_contract.NewBatchOperationFromWrites(ops)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	return sequences.OnChainOutput{
		BatchOps: []mcms_types.BatchOperation{batchOp},
	}, nil
}

func fundSiloedPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokens.MigrateLockReleasePoolLiquidityInput,
	newPoolAddr, tokenAddr, timelockAddr common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	var ops []evm_contract.WriteOutput

	lockboxConfigsReport, err := cldf_ops.ExecuteOperation(b, siloed_lrtp_ops_v170.GetAllLockBoxConfigs, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       newPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get lockbox configs from new pool %s: %w", newPoolAddr, err)
	}

	lockboxByChain := make(map[uint64]common.Address)
	configuredLockBoxes := make(map[common.Address]bool)
	for _, config := range lockboxConfigsReport.Output {
		if config.LockBox == (common.Address{}) {
			continue
		}
		lockboxByChain[config.RemoteChainSelector] = config.LockBox
		configuredLockBoxes[config.LockBox] = true
	}

	unsiloedLockBox, err := resolveUnsiloedLockBox(input.UnsiloedLockBoxAddress, configuredLockBoxes, newPoolAddr)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	unsiloedAmount := input.UnsiloedExactAmount
	if unsiloedAmount == nil {
		unsiloedAmount = big.NewInt(0)
	}
	if unsiloedAmount.Sign() > 0 && unsiloedLockBox == (common.Address{}) {
		return sequences.OnChainOutput{}, fmt.Errorf(
			"%s unsiloed liquidity is being migrated but UnsiloedLockBoxAddress was not set; "+
				"the shared balance backs the pool's non-siloed chains and its destination cannot be inferred - "+
				"set it to the lockbox serving those chains on new pool %s",
			unsiloedAmount, newPoolAddr,
		)
	}

	// Check every silo has a lockbox before emitting any write.
	var chainsWithoutLockBox []uint64
	for _, sa := range input.SiloExactAmounts {
		if _, ok := lockboxByChain[sa.ChainSelector]; !ok {
			chainsWithoutLockBox = append(chainsWithoutLockBox, sa.ChainSelector)
		}
	}
	if len(chainsWithoutLockBox) > 0 {
		slices.Sort(chainsWithoutLockBox)
		return sequences.OnChainOutput{}, fmt.Errorf(
			"new siloed pool %s has no lockbox configured for siloed chains %v; the new pool's lockBoxGroups must cover every siloed chain being migrated",
			newPoolAddr, chainsWithoutLockBox,
		)
	}

	for _, sa := range input.SiloExactAmounts {
		if sa.Amount == nil || sa.Amount.Sign() == 0 {
			continue
		}
		ops, err = appendFundingOps(b, evmChain, chainSel, lockboxByChain[sa.ChainSelector], tokenAddr, timelockAddr, sa.Amount, sa.ChainSelector, input.UsePlainTransfer, ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build funding ops for siloed chain %d: %w", sa.ChainSelector, err)
		}
	}

	if unsiloedAmount.Sign() > 0 {
		ops, err = appendFundingOps(b, evmChain, chainSel, unsiloedLockBox, tokenAddr, timelockAddr, unsiloedAmount, 0, input.UsePlainTransfer, ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build funding ops for unsiloed liquidity: %w", err)
		}
	}

	if input.SetPoolConfig != nil {
		ops, err = appendSetPool(b, evmChain, chainSel, input.SetPoolConfig, newPoolAddr, ops)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
	}

	batchOp, err := evm_contract.NewBatchOperationFromWrites(ops)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	return sequences.OnChainOutput{
		BatchOps: []mcms_types.BatchOperation{batchOp},
	}, nil
}

// appendFundingOps appends the operations that fund the new pool's lockbox with migrated
// liquidity. By default it uses the lockbox's deposit() path (emitting the Deposit event). Use
// usePlainTransfer=true as a break-glass option to transfer the tokens directly, bypassing the
// Deposit event.
func appendFundingOps(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	lockboxAddr, tokenAddr, timelockAddr common.Address,
	amount *big.Int,
	remoteChainSelector uint64,
	usePlainTransfer bool,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	if usePlainTransfer {
		transferReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.TransferProposalOnly, evmChain, evm_contract.FunctionInput[erc20_ops.TransferArgs]{
			ChainSelector: chainSel,
			Address:       tokenAddr,
			Args: erc20_ops.TransferArgs{
				Receiver: lockboxAddr,
				Amount:   amount,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to transfer tokens to lockbox %s: %w", lockboxAddr, err)
		}
		return append(ops, transferReport.Output), nil
	}

	// Liquidity migration is done in tranches, so the authorize step must be idempotent: only append
	// the authorized-caller update if the timelock isn't already an authorized caller, avoiding a
	// redundant MCMS batch op and a spurious AuthorizedCallerAdded event on every subsequent tranche.
	authCallers, err := cldf_ops.ExecuteOperation(
		b,
		lockbox_ops.GetAllAuthorizedCallers,
		evmChain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       lockboxAddr,
		},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[struct{}], evm.Chain](),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get authorized callers on lockbox %s: %w", lockboxAddr, err)
	}
	if !slices.Contains(authCallers.Output, timelockAddr) {
		// Deposit path: make the timelock an authorized caller so it can deposit, approve the lockbox,
		// then deposit via the lockbox's deposit() function (emits the Deposit event). The framework
		// routes the authorize step automatically based on lockbox ownership: EOA when deployer-owned
		// or MCMS-batched when timelock-owned. The timelock remains an authorized caller afterward;
		// this is harmless since the timelock is governance and typically owns the lockbox.
		addAuthReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.ApplyAuthorizedCallerUpdates, evmChain, evm_contract.FunctionInput[lockbox_ops.AuthorizedCallerArgs]{
			ChainSelector: chainSel,
			Address:       lockboxAddr,
			Args: lockbox_ops.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{timelockAddr},
				RemovedCallers: []common.Address{},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to add timelock as authorized caller on lockbox %s: %w", lockboxAddr, err)
		}
		ops = append(ops, addAuthReport.Output)
	}

	approveReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.ApproveProposalOnly, evmChain, evm_contract.FunctionInput[erc20_ops.ApproveArgs]{
		ChainSelector: chainSel,
		Address:       tokenAddr,
		Args: erc20_ops.ApproveArgs{
			Spender: lockboxAddr,
			Value:   amount,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to approve lockbox %s to spend tokens: %w", lockboxAddr, err)
	}
	ops = append(ops, approveReport.Output)

	depositReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.DepositProposalOnly, evmChain, evm_contract.FunctionInput[lockbox_ops.DepositArgs]{
		ChainSelector: chainSel,
		Address:       lockboxAddr,
		Args: lockbox_ops.DepositArgs{
			Token:               tokenAddr,
			RemoteChainSelector: remoteChainSelector,
			Amount:              amount,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deposit into lockbox %s: %w", lockboxAddr, err)
	}
	ops = append(ops, depositReport.Output)

	return ops, nil
}

func appendSetPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	config *tokens.MigrationSetPoolConfig,
	newPoolAddr common.Address,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	setPoolReport, err := cldf_ops.ExecuteOperation(b, tar_ops.SetPool, evmChain, evm_contract.FunctionInput[tar_ops.SetPoolArgs]{
		ChainSelector: chainSel,
		Address:       common.HexToAddress(config.RegistryAddress),
		Args: tar_ops.SetPoolArgs{
			TokenAddress:     common.HexToAddress(config.TokenAddress),
			TokenPoolAddress: newPoolAddr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set pool on token admin registry: %w", err)
	}
	ops = append(ops, setPoolReport.Output)

	return ops, nil
}
