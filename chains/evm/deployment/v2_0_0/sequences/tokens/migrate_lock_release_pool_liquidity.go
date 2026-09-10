package tokens

import (
	"bytes"
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
	lrtp_ops_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	siloed_ops_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"
	lockbox_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	siloed_lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

var MigrateLockReleasePoolLiquidity = cldf_ops.NewSequence(
	"migrate-lock-release-pool-liquidity",
	semver.MustParse("2.0.0"),
	"Migrates liquidity from a legacy LockReleaseTokenPool (v1.5.1/v1.6.1) to a v2.0 lockbox-based pool",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
		evmChain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		if err := validateMigrationInput(input); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid migration input: %w", err)
		}

		oldPoolAddr := common.HexToAddress(input.OldPoolAddress)
		newPoolAddr := common.HexToAddress(input.NewPoolAddress)
		timelockAddr := common.HexToAddress(input.TimelockAddress)

		tvReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       oldPoolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get typeAndVersion from old pool %s: %w", oldPoolAddr, err)
		}
		oldPoolType := string(tvReport.Output.Type)

		tokenReport, err := cldf_ops.ExecuteOperation(b, token_pool_ops.GetToken, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       newPoolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from new pool %s: %w", newPoolAddr, err)
		}
		tokenAddr := tokenReport.Output

		// Only the generic siloed lock-release pool is handled here. A substring match on "Siloed"
		// would also catch SiloedUSDCTokenPool, which migrates through the CCTP hybrid path instead.
		isSiloed := oldPoolType == utils.SiloedLockReleaseTokenPool.String()

		if isSiloed {
			return migrateSiloedPool(b, evmChain, input, oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr)
		}
		if isExactSiloMode(input) {
			return sequences.OnChainOutput{}, fmt.Errorf("SiloExactAmounts/UnsiloedExactAmount are only supported for siloed pools")
		}
		return migrateUnsiloedPool(b, evmChain, input, oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr)
	},
)

// isExactSiloMode reports whether the input requests exact-amount migration (per-silo and/or
// unsiloed shared bucket) rather than the legacy Amount/BasisPoints proportional mode.
func isExactSiloMode(input tokens.MigrateLockReleasePoolLiquidityInput) bool {
	return len(input.SiloExactAmounts) > 0 || input.UnsiloedExactAmount != nil
}

// siloExactAmount returns the exact amount configured for the given remote chain's silo, if any.
func siloExactAmount(input tokens.MigrateLockReleasePoolLiquidityInput, chainSel uint64) (*big.Int, bool) {
	for _, sa := range input.SiloExactAmounts {
		if sa.ChainSelector == chainSel {
			return sa.Amount, true
		}
	}
	return nil, false
}

func validateMigrationInput(input tokens.MigrateLockReleasePoolLiquidityInput) error {
	exactMode := isExactSiloMode(input)
	legacyMode := input.Amount != nil || input.BasisPoints != nil

	if exactMode && legacyMode {
		return fmt.Errorf("SiloExactAmounts/UnsiloedExactAmount are mutually exclusive with Amount/BasisPoints")
	}
	if !exactMode && !legacyMode {
		return fmt.Errorf("one of Amount, BasisPoints, or SiloExactAmounts/UnsiloedExactAmount must be provided")
	}
	if input.Amount != nil && input.BasisPoints != nil {
		return fmt.Errorf("Amount and BasisPoints are mutually exclusive")
	}
	if input.BasisPoints != nil {
		bp := *input.BasisPoints
		if bp == 0 || bp > 10000 {
			return fmt.Errorf("BasisPoints must be between 1 and 10000, got %d", bp)
		}
	}
	if input.Amount != nil && input.Amount.Sign() <= 0 {
		return fmt.Errorf("Amount must be positive")
	}
	if exactMode {
		// UnsiloedExactAmount is a companion to SiloExactAmounts, not a standalone mode: migrating
		// only the shared bucket while leaving every silo untouched isn't supported by this loop
		// (it would migrate all siloed chains too), so reject it outright rather than silently
		// draining silos the operator didn't intend to touch.
		if input.UnsiloedExactAmount != nil && len(input.SiloExactAmounts) == 0 {
			return fmt.Errorf("UnsiloedExactAmount requires SiloExactAmounts to also be set; exact mode cannot migrate the unsiloed bucket alone")
		}
		seen := make(map[uint64]bool, len(input.SiloExactAmounts))
		for i, sa := range input.SiloExactAmounts {
			if seen[sa.ChainSelector] {
				return fmt.Errorf("duplicate ChainSelector %d in SiloExactAmounts", sa.ChainSelector)
			}
			seen[sa.ChainSelector] = true
			if sa.Amount == nil || sa.Amount.Sign() < 0 {
				return fmt.Errorf("SiloExactAmounts[%d].Amount must be positive", i)
			}
		}
		if input.UnsiloedExactAmount != nil && input.UnsiloedExactAmount.Sign() <= 0 {
			return fmt.Errorf("UnsiloedExactAmount must be positive")
		}
	}
	if input.OldPoolAddress == "" || input.NewPoolAddress == "" {
		return fmt.Errorf("OldPoolAddress and NewPoolAddress must be provided")
	}
	if input.TimelockAddress == "" {
		return fmt.Errorf("TimelockAddress must be provided")
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

func computeAmount(balance *big.Int, input tokens.MigrateLockReleasePoolLiquidityInput) *big.Int {
	if input.Amount != nil {
		return new(big.Int).Set(input.Amount)
	}
	bp := *input.BasisPoints
	if bp == 10000 {
		return new(big.Int).Set(balance)
	}
	amount := new(big.Int).Mul(balance, big.NewInt(int64(bp)))
	return amount.Div(amount, big.NewInt(10000))
}

func migrateUnsiloedPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokens.MigrateLockReleasePoolLiquidityInput,
	oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	var ops []evm_contract.WriteOutput

	lockboxReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v170.GetLockBox, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       newPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get lockbox from new pool %s: %w", newPoolAddr, err)
	}
	lockboxAddr := lockboxReport.Output

	balanceReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.BalanceOf, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddr,
		Args:          oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get balance of old pool %s: %w", oldPoolAddr, err)
	}
	balance := balanceReport.Output

	amount := computeAmount(balance, input)
	if amount.Sign() == 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("computed migration amount is zero")
	}
	if amount.Cmp(balance) > 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("migration amount %s exceeds old pool balance %s", amount, balance)
	}

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get rebalancer from old pool %s: %w", oldPoolAddr, err)
	}
	originalRebalancer := rebalancerReport.Output

	ops, err = appendSetRebalancerAndWithdraw(b, evmChain, chainSel, oldPoolAddr, timelockAddr, amount, ops)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	ops, err = appendFundingOps(b, evmChain, chainSel, lockboxAddr, tokenAddr, timelockAddr, amount, 0, input.UsePlainTransfer, ops)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	restoreRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          originalRebalancer,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to restore rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, restoreRebalancerReport.Output)

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

func migrateSiloedPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokens.MigrateLockReleasePoolLiquidityInput,
	oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	var ops []evm_contract.WriteOutput

	chainsReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetSupportedChains, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains from old siloed pool %s: %w", oldPoolAddr, err)
	}
	supportedChains := chainsReport.Output

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

	// Resolve the destination for the unsiloed (shared) balance up front, so a missing or wrong
	// address fails before any write is emitted rather than partway through the batch.
	unsiloedLockBox, err := resolveUnsiloedLockBox(input.UnsiloedLockBoxAddress, configuredLockBoxes, newPoolAddr)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	// Resolve which of the old pool's chains are siloed once; used both for the coverage check below
	// and for the rebalancer handover further down.
	isSiloedByChain := make(map[uint64]bool, len(supportedChains))
	for _, remoteChain := range supportedChains {
		isSiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.IsSiloed, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if chain %d is siloed on old pool %s: %w", remoteChain, oldPoolAddr, err)
		}
		isSiloedByChain[remoteChain] = isSiloedReport.Output
	}

	// Check lockbox coverage before emitting any writes. Without this the batch is built chain by
	// chain and a gap surfaces partway through - after the rebalancer has already been repointed at
	// the timelock - leaving the engineer to work out which silo was missing.
	var chainsWithoutLockBox []uint64
	for _, remoteChain := range supportedChains {
		if !isSiloedByChain[remoteChain] {
			continue
		}
		if _, ok := lockboxByChain[remoteChain]; !ok {
			chainsWithoutLockBox = append(chainsWithoutLockBox, remoteChain)
		}
	}
	if len(chainsWithoutLockBox) > 0 {
		slices.Sort(chainsWithoutLockBox)
		return sequences.OnChainOutput{}, fmt.Errorf(
			"new siloed pool %s has no lockbox configured for siloed chains %v of old pool %s; the new pool's lockBoxGroups must cover every siloed chain being migrated",
			newPoolAddr, chainsWithoutLockBox, oldPoolAddr,
		)
	}

	exactMode := isExactSiloMode(input)

	// In exact mode, every siloed chain must be explicit - a chain silently falling back to zero
	// migration is more likely to be an oversight than intent, so surface it before any write is
	// emitted rather than migrating a partial set of silos.
	if exactMode {
		validSiloedChains := make(map[uint64]bool, len(supportedChains))
		for _, remoteChain := range supportedChains {
			if isSiloedByChain[remoteChain] {
				validSiloedChains[remoteChain] = true
			}
		}

		// A ChainSelector that isn't actually a siloed chain of the old pool would otherwise be
		// silently ignored by the withdraw loop below, leaving no trace in the execution logs that
		// the entry was never applied - fail loudly instead so a typo'd or stale selector surfaces
		// immediately.
		var unknownSilos []uint64
		for _, sa := range input.SiloExactAmounts {
			if !validSiloedChains[sa.ChainSelector] {
				unknownSilos = append(unknownSilos, sa.ChainSelector)
			}
		}
		if len(unknownSilos) > 0 {
			slices.Sort(unknownSilos)
			return sequences.OnChainOutput{}, fmt.Errorf(
				"SiloExactAmounts references chain selectors %v that are not siloed chains on old pool %s",
				unknownSilos, oldPoolAddr,
			)
		}

		var missingSilos []uint64
		for remoteChain := range validSiloedChains {
			if _, ok := siloExactAmount(input, remoteChain); !ok {
				missingSilos = append(missingSilos, remoteChain)
			}
		}
		if len(missingSilos) > 0 {
			slices.Sort(missingSilos)
			return sequences.OnChainOutput{}, fmt.Errorf(
				"SiloExactAmounts is missing entries for siloed chains %v of old pool %s; exact mode requires every siloed chain to be explicit",
				missingSilos, oldPoolAddr,
			)
		}
	}

	// Read the shared balance here, for the same reason as the coverage check above: a missing
	// destination must surface before the rebalancer is repointed at the timelock, not once the
	// silos have already been drained.
	unsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetUnsiloedLiquidity, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get unsiloed liquidity from old pool %s: %w", oldPoolAddr, err)
	}

	var unsiloedAmount *big.Int
	if exactMode {
		if unsiloedReport.Output.Sign() > 0 && input.UnsiloedExactAmount == nil {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"old pool %s holds %s unsiloed liquidity to migrate but UnsiloedExactAmount was not set; "+
					"exact mode requires an explicit amount for the shared bucket",
				oldPoolAddr, unsiloedReport.Output,
			)
		}
		if input.UnsiloedExactAmount != nil {
			unsiloedAmount = new(big.Int).Set(input.UnsiloedExactAmount)
		} else {
			unsiloedAmount = big.NewInt(0)
		}
		if unsiloedAmount.Cmp(unsiloedReport.Output) > 0 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"UnsiloedExactAmount %s exceeds old pool %s unsiloed balance %s",
				unsiloedAmount, oldPoolAddr, unsiloedReport.Output,
			)
		}
	} else {
		unsiloedAmount = computeAmount(unsiloedReport.Output, input)
	}
	if unsiloedAmount.Sign() > 0 && unsiloedLockBox == (common.Address{}) {
		return sequences.OnChainOutput{}, fmt.Errorf(
			"old pool %s holds %s unsiloed liquidity to migrate but UnsiloedLockBoxAddress was not set; "+
				"the shared balance backs the pool's non-siloed chains and its destination cannot be inferred - "+
				"set it to the lockbox serving those chains on new pool %s",
			oldPoolAddr, unsiloedAmount, newPoolAddr,
		)
	}

	// Resolve every siloed chain's migration amount and validate it against the on-chain balance
	// up front, before any write is emitted (the rebalancer handover below is the first write).
	siloAmounts := make(map[uint64]*big.Int, len(supportedChains))
	for _, remoteChain := range supportedChains {
		if !isSiloedByChain[remoteChain] {
			continue
		}

		availableReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetAvailableTokens, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get available tokens for chain %d: %w", remoteChain, err)
		}
		siloBalance := availableReport.Output

		var siloAmount *big.Int
		if exactMode {
			amt, _ := siloExactAmount(input, remoteChain) // presence guaranteed by the missing-silo check above
			siloAmount = amt
		} else {
			siloAmount = computeAmount(siloBalance, input)
		}
		if siloAmount.Cmp(siloBalance) > 0 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"migration amount %s for chain %d exceeds silo balance %s", siloAmount, remoteChain, siloBalance,
			)
		}
		siloAmounts[remoteChain] = siloAmount
	}

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get unsiloed rebalancer from old pool %s: %w", oldPoolAddr, err)
	}
	originalUnsiloedRebalancer := rebalancerReport.Output

	setRebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          timelockAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to set unsiloed rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, setRebalancerReport.Output)

	type chainRebalancerInfo struct {
		chainSelector      uint64
		originalRebalancer common.Address
		isSiloed           bool
	}
	var siloInfos []chainRebalancerInfo

	for _, remoteChain := range supportedChains {
		if isSiloedByChain[remoteChain] {
			chainRebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetChainRebalancer, evmChain, evm_contract.FunctionInput[uint64]{
				ChainSelector: chainSel,
				Address:       oldPoolAddr,
				Args:          remoteChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get chain rebalancer for chain %d: %w", remoteChain, err)
			}

			setSiloReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetSiloRebalancer, evmChain, evm_contract.FunctionInput[siloed_ops_v161.SetSiloRebalancerArgs]{
				ChainSelector: chainSel,
				Address:       oldPoolAddr,
				Args: siloed_ops_v161.SetSiloRebalancerArgs{
					RemoteChainSelector: remoteChain,
					NewRebalancer:       timelockAddr,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set silo rebalancer for chain %d: %w", remoteChain, err)
			}
			ops = append(ops, setSiloReport.Output)

			siloInfos = append(siloInfos, chainRebalancerInfo{
				chainSelector:      remoteChain,
				originalRebalancer: chainRebalancerReport.Output,
				isSiloed:           true,
			})
		} else {
			siloInfos = append(siloInfos, chainRebalancerInfo{
				chainSelector: remoteChain,
				isSiloed:      false,
			})
		}
	}

	for _, info := range siloInfos {
		if !info.isSiloed {
			continue
		}

		lockbox, ok := lockboxByChain[info.chainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("no lockbox configured for chain %d on new siloed pool", info.chainSelector)
		}

		siloAmount := siloAmounts[info.chainSelector]
		if siloAmount.Sign() == 0 {
			continue
		}

		withdrawReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.WithdrawSiloedLiquidity, evmChain, evm_contract.FunctionInput[siloed_ops_v161.WithdrawSiloedLiquidityArgs]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args: siloed_ops_v161.WithdrawSiloedLiquidityArgs{
				RemoteChainSelector: info.chainSelector,
				Amount:              siloAmount,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw siloed liquidity for chain %d: %w", info.chainSelector, err)
		}
		ops = append(ops, withdrawReport.Output)

		ops, err = appendFundingOps(b, evmChain, chainSel, lockbox, tokenAddr, timelockAddr, siloAmount, info.chainSelector, input.UsePlainTransfer, ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build funding ops for siloed chain %d: %w", info.chainSelector, err)
		}
	}

	if unsiloedAmount.Sign() > 0 {
		withdrawUnsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.WithdrawLiquidity, evmChain, evm_contract.FunctionInput[*big.Int]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          unsiloedAmount,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw unsiloed liquidity: %w", err)
		}
		ops = append(ops, withdrawUnsiloedReport.Output)

		ops, err = appendFundingOps(b, evmChain, chainSel, unsiloedLockBox, tokenAddr, timelockAddr, unsiloedAmount, 0, input.UsePlainTransfer, ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build funding ops for unsiloed liquidity: %w", err)
		}
	}

	for _, info := range siloInfos {
		if info.isSiloed {
			restoreReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetSiloRebalancer, evmChain, evm_contract.FunctionInput[siloed_ops_v161.SetSiloRebalancerArgs]{
				ChainSelector: chainSel,
				Address:       oldPoolAddr,
				Args: siloed_ops_v161.SetSiloRebalancerArgs{
					RemoteChainSelector: info.chainSelector,
					NewRebalancer:       info.originalRebalancer,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to restore silo rebalancer for chain %d: %w", info.chainSelector, err)
			}
			ops = append(ops, restoreReport.Output)
		}
	}

	restoreUnsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          originalUnsiloedRebalancer,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to restore unsiloed rebalancer: %w", err)
	}
	ops = append(ops, restoreUnsiloedReport.Output)

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

func appendSetRebalancerAndWithdraw(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	oldPoolAddr, timelockAddr common.Address,
	amount *big.Int,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	setRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          timelockAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, setRebalancerReport.Output)

	withdrawReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.WithdrawLiquidity, evmChain, evm_contract.FunctionInput[*big.Int]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          amount,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw liquidity from old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, withdrawReport.Output)

	return ops, nil
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
