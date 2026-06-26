package sequences

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// =============================================================================
// Input Types
// =============================================================================

// ConfigureFeeSweepV150Input configures an OnRamp to send all LINK fees to a treasury address.
type ConfigureFeeSweepV150Input struct {
	ChainSelector     uint64
	OnRampAddress     common.Address
	Treasury          common.Address
	AllowedRecipients map[uint64]common.Address
}

func (c ConfigureFeeSweepV150Input) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	if c.OnRampAddress == (common.Address{}) {
		return fmt.Errorf("OnRampAddress cannot be zero address")
	}
	if err := validateTreasuryAddress(c.AllowedRecipients, c.ChainSelector, c.Treasury); err != nil {
		return fmt.Errorf("treasury validation failed: %w", err)
	}
	return nil
}

// SweepLinkFeesV150Input triggers LINK fee payout after verifying NOP configuration.
type SweepLinkFeesV150Input struct {
	ChainSelector     uint64
	OnRampAddress     common.Address
	ExpectedTreasury  common.Address
	AllowedRecipients map[uint64]common.Address
}

func (c SweepLinkFeesV150Input) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	if c.OnRampAddress == (common.Address{}) {
		return fmt.Errorf("OnRampAddress cannot be zero address")
	}
	if err := validateTreasuryAddress(c.AllowedRecipients, c.ChainSelector, c.ExpectedTreasury); err != nil {
		return fmt.Errorf("expected treasury validation failed: %w", err)
	}
	return nil
}

// SweepNonLinkFeesV150Input withdraws multiple non-LINK fee tokens to treasury.
// If WETH9Address is non-zero, tokens matching WETH9 trigger atomic 3-step unwrap
// (withdraw to MCMS → unwrap WETH → transfer native ETH to treasury).
type SweepNonLinkFeesV150Input struct {
	ChainSelector     uint64
	OnRampAddress     common.Address
	FeeTokens         []common.Address
	Treasury          common.Address
	WETH9Address      common.Address // Zero = no WETH auto-detection
	MCMSAddress       common.Address // RBACTimelock for WETH unwrap destination
	WETHBalance       *big.Int       // WETH balance of OnRamp (for unwrap amount)
	AllowedRecipients map[uint64]common.Address
}

func (c SweepNonLinkFeesV150Input) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	if c.OnRampAddress == (common.Address{}) {
		return fmt.Errorf("OnRampAddress cannot be zero address")
	}
	if err := validateTreasuryAddress(c.AllowedRecipients, c.ChainSelector, c.Treasury); err != nil {
		return fmt.Errorf("treasury validation failed: %w", err)
	}
	if c.WETH9Address != (common.Address{}) && c.MCMSAddress == (common.Address{}) {
		return fmt.Errorf("MCMSAddress must be set when WETH9Address is provided")
	}
	return nil
}

// SweepAllOnRampsV150Input is the mega flow input: configure NOPs, sweep LINK,
// sweep non-LINK (with WETH auto-unwrap) for ALL OnRamps on a chain.
type SweepAllOnRampsV150Input struct {
	ChainSelector       uint64
	Treasury            common.Address
	OnRamps             []common.Address
	NonLinkFeeTokens    []common.Address                               // Non-LINK tokens for this chain
	WETH9Address        common.Address                                 // Zero = no WETH unwrap
	MCMSAddress         common.Address                                 // RBACTimelock, zero if no WETH unwrap
	OnRampTokenBalances map[common.Address]map[common.Address]*big.Int // onRamp → token → balance
	OnRampLINKFees      map[common.Address]*big.Int                    // onRamp → LINK fee balance (juels)
	MinSweepAmount      *big.Int                                       // Skip tokens below this (wei)
	AllowedRecipients   map[uint64]common.Address
	// SkipNopsCheck skips the on-chain NOP pre-check when true.
	// Use for dry-run testing only. The SetNops tx is always included
	// in the batch regardless of this setting.
	SkipNopsCheck bool
}

func (c SweepAllOnRampsV150Input) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	if err := validateTreasuryAddress(c.AllowedRecipients, c.ChainSelector, c.Treasury); err != nil {
		return fmt.Errorf("treasury validation failed: %w", err)
	}
	if len(c.OnRamps) == 0 {
		return fmt.Errorf("no OnRamps provided")
	}
	if c.WETH9Address != (common.Address{}) && c.MCMSAddress == (common.Address{}) {
		return fmt.Errorf("MCMSAddress must be set when WETH9Address is provided")
	}
	if c.MinSweepAmount == nil {
		return fmt.Errorf("MinSweepAmount must not be nil (use 0 to sweep everything)")
	}
	if c.MinSweepAmount.Sign() < 0 {
		return fmt.Errorf("MinSweepAmount must be non-negative, got %s", c.MinSweepAmount.String())
	}
	for i, addr := range c.OnRamps {
		if addr == (common.Address{}) {
			return fmt.Errorf("OnRamps[%d] cannot be zero address", i)
		}
	}
	return nil
}

// =============================================================================
// Treasury Address Validation
// =============================================================================

func validateTreasuryAddress(allowedRecipients map[uint64]common.Address, chainSelector uint64, address common.Address) error {
	if address == (common.Address{}) {
		return fmt.Errorf("treasury address cannot be zero address")
	}
	if allowedRecipients == nil {
		return fmt.Errorf("allowedRecipients map is nil; must be provided from CLD inputs")
	}
	allowedAddr, ok := allowedRecipients[chainSelector]
	if !ok {
		return fmt.Errorf("chain selector %d is not in the allowed recipients list", chainSelector)
	}
	if allowedAddr != address {
		return fmt.Errorf("address %s is not the approved treasury for chain %d; expected %s",
			address.Hex(), chainSelector, allowedAddr.Hex())
	}
	return nil
}

// =============================================================================
// Shared Helper Functions
// =============================================================================

// checkNopConfig reads the current NOP config and returns true if it matches the
// expected treasury setup: exactly 1 NOP, address == treasury, weight == 65535,
// weightsTotal == 65535.
func checkNopConfig(b cldf_ops.Bundle, chain evm.Chain, onRampAddr, treasury common.Address) (bool, error) {
	nopsResult, err := cldf_ops.ExecuteOperation(b, onramp.OnRampGetNops, chain, evm_contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       onRampAddr,
	})
	if err != nil {
		return false, fmt.Errorf("failed to get NOP config for OnRamp %s: %w", onRampAddr.Hex(), err)
	}

	nops := nopsResult.Output.NopsAndWeights
	if len(nops) != 1 {
		return false, nil
	}
	if nops[0].Nop != treasury {
		return false, nil
	}
	if nops[0].Weight != 65535 {
		return false, nil
	}
	if nopsResult.Output.WeightsTotal == nil || nopsResult.Output.WeightsTotal.Cmp(big.NewInt(65535)) != 0 {
		return false, nil
	}
	return true, nil
}

// buildSetNopsTx builds a SetNops transaction to configure treasury as sole NOP with 100% weight.
func buildSetNopsTx(b cldf_ops.Bundle, chain evm.Chain, onRampAddr, treasury common.Address) (mcms_types.Transaction, error) {
	report, err := cldf_ops.ExecuteOperation(b, onramp.OnRampSetNops, chain, evm_contract.FunctionInput[onramp.SetNopsInput]{
		ChainSelector: chain.Selector,
		Address:       onRampAddr,
		Args: onramp.SetNopsInput{
			NopsAndWeights: []onramp.NopAndWeight{
				{Nop: treasury, Weight: 65535},
			},
		},
	})
	if err != nil {
		return mcms_types.Transaction{}, fmt.Errorf("failed to build SetNops tx for OnRamp %s: %w", onRampAddr.Hex(), err)
	}
	return report.Output.Tx, nil
}

// buildPayNopsTx builds a PayNops transaction. Caller must verify NOP config first.
func buildPayNopsTx(b cldf_ops.Bundle, chain evm.Chain, onRampAddr common.Address) (mcms_types.Transaction, error) {
	report, err := cldf_ops.ExecuteOperation(b, onramp.OnRampPayNops, chain, evm_contract.FunctionInput[any]{
		ChainSelector: chain.Selector,
		Address:       onRampAddr,
	})
	if err != nil {
		return mcms_types.Transaction{}, fmt.Errorf("failed to build PayNops tx for OnRamp %s: %w", onRampAddr.Hex(), err)
	}
	return report.Output.Tx, nil
}

// buildNonLinkSweepTxs builds transactions for sweeping non-LINK fee tokens.
// If a token matches weth9Addr (and weth9Addr is non-zero), it does atomic 3-step
// unwrap: withdraw WETH to MCMS → unwrap to native ETH → transfer to treasury.
// Otherwise does direct WithdrawNonLinkFees to treasury.
func buildNonLinkSweepTxs(b cldf_ops.Bundle, chain evm.Chain, onRampAddr common.Address,
	feeTokens []common.Address, treasury, weth9Addr, mcmsAddr common.Address,
	wethBalance *big.Int) ([]mcms_types.Transaction, error) {

	var txs []mcms_types.Transaction

	for _, token := range feeTokens {
		if weth9Addr != (common.Address{}) && token == weth9Addr {
			// Skip WETH if balance is nil/zero
			if wethBalance == nil || wethBalance.Sign() <= 0 {
				continue
			}

			// Atomic 3-step WETH unwrap

			// Step 1: Withdraw WETH from OnRamp to MCMS timelock
			withdrawReport, err := cldf_ops.ExecuteOperation(b, onramp.OnRampWithdrawNonLinkFees, chain,
				evm_contract.FunctionInput[onramp.WithdrawNonLinkFeesInput]{
					ChainSelector: chain.Selector,
					Address:       onRampAddr,
					Args: onramp.WithdrawNonLinkFeesInput{
						FeeToken: weth9Addr,
						To:       mcmsAddr,
					},
				})
			if err != nil {
				return nil, fmt.Errorf("failed to build WETH withdraw tx for OnRamp %s: %w", onRampAddr.Hex(), err)
			}
			if !withdrawReport.Output.Executed() {
				txs = append(txs, withdrawReport.Output.Tx)
			}

			// Step 2: MCMS calls WETH9.withdraw() to unwrap WETH to native ETH
			unwrapReport, err := cldf_ops.ExecuteOperation(b, weth.Withdraw, chain,
				evm_contract.FunctionInput[weth.WithdrawInput]{
					ChainSelector: chain.Selector,
					Address:       weth9Addr,
					Args:          weth.WithdrawInput{Amount: wethBalance},
				})
			if err != nil {
				return nil, fmt.Errorf("failed to build WETH unwrap tx for OnRamp %s: %w", onRampAddr.Hex(), err)
			}
			if !unwrapReport.Output.Executed() {
				txs = append(txs, unwrapReport.Output.Tx)
			}

			// Step 3: Transfer native ETH from MCMS to treasury
			ethTx, err := weth.CreateNativeETHTransferTx(treasury, wethBalance)
			if err != nil {
				return nil, fmt.Errorf("failed to build ETH transfer tx for OnRamp %s: %w", onRampAddr.Hex(), err)
			}
			txs = append(txs, ethTx)
		} else {
			// Direct withdraw to treasury
			report, err := cldf_ops.ExecuteOperation(b, onramp.OnRampWithdrawNonLinkFees, chain,
				evm_contract.FunctionInput[onramp.WithdrawNonLinkFeesInput]{
					ChainSelector: chain.Selector,
					Address:       onRampAddr,
					Args: onramp.WithdrawNonLinkFeesInput{
						FeeToken: token,
						To:       treasury,
					},
				})
			if err != nil {
				return nil, fmt.Errorf("failed to build WithdrawNonLinkFees tx for token %s on OnRamp %s: %w",
					token.Hex(), onRampAddr.Hex(), err)
			}
			if !report.Output.Executed() {
				txs = append(txs, report.Output.Tx)
			}
		}
	}

	return txs, nil
}

// getTokenBalance safely retrieves a token balance from a nested map.
func getTokenBalance(balances map[common.Address]map[common.Address]*big.Int, onRamp, token common.Address) *big.Int {
	if balances == nil {
		return nil
	}
	tokenBalances, ok := balances[onRamp]
	if !ok {
		return nil
	}
	return tokenBalances[token]
}

// =============================================================================
// Sequences
// =============================================================================

// ConfigureFeeSweepV150Sequence sets up an OnRamp to send all LINK fees to the treasury.
var ConfigureFeeSweepV150Sequence = cldf_ops.NewSequence(
	"onramp:configure-fee-sweep",
	semver.MustParse("1.5.0"),
	"Configures the OnRamp 1.5.0 to send all LINK fees to the treasury address",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureFeeSweepV150Input) (sequences.OnChainOutput, error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		tx, err := buildSetNopsTx(b, chain, input.OnRampAddress, input.Treasury)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		batch := mcms_types.BatchOperation{
			ChainSelector: mcms_types.ChainSelector(chain.Selector),
			Transactions:  []mcms_types.Transaction{tx},
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)

// SweepLinkFeesV150Sequence triggers LINK fee payout after verifying NOP configuration.
var SweepLinkFeesV150Sequence = cldf_ops.NewSequence(
	"onramp:sweep-link-fees",
	semver.MustParse("1.5.0"),
	"Sweeps accumulated LINK fees from the OnRamp 1.5.0 to configured NOPs (treasury)",
	func(b cldf_ops.Bundle, chain evm.Chain, input SweepLinkFeesV150Input) (sequences.OnChainOutput, error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		// Full NOP config validation before PayNops
		isCorrect, err := checkNopConfig(b, chain, input.OnRampAddress, input.ExpectedTreasury)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		if !isCorrect {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"NOP configuration is not correct for OnRamp %s; expected sole NOP %s with weight 65535; run configure_fee_sweep_v150_evm first",
				input.OnRampAddress.Hex(), input.ExpectedTreasury.Hex())
		}

		tx, err := buildPayNopsTx(b, chain, input.OnRampAddress)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		batch := mcms_types.BatchOperation{
			ChainSelector: mcms_types.ChainSelector(chain.Selector),
			Transactions:  []mcms_types.Transaction{tx},
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)

// SweepNonLinkFeesV150Sequence withdraws non-LINK fee tokens to treasury.
// Auto-detects WETH among fee tokens and does atomic unwrap if WETH9Address is set.
var SweepNonLinkFeesV150Sequence = cldf_ops.NewSequence(
	"onramp:sweep-non-link-fees",
	semver.MustParse("1.5.0"),
	"Sweeps accumulated non-LINK fee tokens from the OnRamp 1.5.0 to treasury",
	func(b cldf_ops.Bundle, chain evm.Chain, input SweepNonLinkFeesV150Input) (sequences.OnChainOutput, error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		txs, err := buildNonLinkSweepTxs(b, chain, input.OnRampAddress, input.FeeTokens,
			input.Treasury, input.WETH9Address, input.MCMSAddress, input.WETHBalance)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}

		if len(txs) == 0 {
			return sequences.OnChainOutput{}, nil
		}

		batch := mcms_types.BatchOperation{
			ChainSelector: mcms_types.ChainSelector(chain.Selector),
			Transactions:  txs,
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)

// SweepAllOnRampsV150Sequence is the mega flow: configure NOPs, sweep LINK,
// sweep non-LINK (with WETH auto-unwrap) for ALL OnRamps on a chain.
// All transactions are in a single atomic MCMS batch.
var SweepAllOnRampsV150Sequence = cldf_ops.NewSequence(
	"onramp:sweep-all-onramps",
	semver.MustParse("1.5.0"),
	"Configures NOPs and sweeps all fees (LINK + non-LINK + WETH unwrap) from ALL OnRamps v1.5.0",
	func(b cldf_ops.Bundle, chain evm.Chain, input SweepAllOnRampsV150Input) (sequences.OnChainOutput, error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}

		var allTxs []mcms_types.Transaction

		for _, onRampAddr := range input.OnRamps {
			// 1. Verify NOP config — fail if not configured (unless skipped for dry-run)
			if !input.SkipNopsCheck {
				isCorrect, err := checkNopConfig(b, chain, onRampAddr, input.Treasury)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				if !isCorrect {
					return sequences.OnChainOutput{}, fmt.Errorf(
						"NOP config for OnRamp %s is not set to treasury %s; "+
							"run configure_fee_sweep_v150_evm first",
						onRampAddr.Hex(), input.Treasury.Hex())
				}
			}
			// Defensive re-assertion: always include SetNops in the batch to guard
			// against concurrent proposals modifying NOP config between resolution
			// and execution.
			{
				tx, err := buildSetNopsTx(b, chain, onRampAddr, input.Treasury)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				allTxs = append(allTxs, tx)
			}

			// 2. Sweep LINK if above threshold
			linkFees := input.OnRampLINKFees[onRampAddr]
			if linkFees != nil && linkFees.Cmp(input.MinSweepAmount) >= 0 {
				tx, err := buildPayNopsTx(b, chain, onRampAddr)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				allTxs = append(allTxs, tx)
			}

			// 3. Sweep non-LINK tokens above threshold
			for _, token := range input.NonLinkFeeTokens {
				tokenBalance := getTokenBalance(input.OnRampTokenBalances, onRampAddr, token)
				if tokenBalance == nil || tokenBalance.Cmp(input.MinSweepAmount) < 0 {
					continue
				}

				var wethBal *big.Int
				if token == input.WETH9Address {
					wethBal = tokenBalance
				}

				tokenTxs, err := buildNonLinkSweepTxs(b, chain, onRampAddr,
					[]common.Address{token}, input.Treasury,
					input.WETH9Address, input.MCMSAddress, wethBal)
				if err != nil {
					return sequences.OnChainOutput{}, err
				}
				allTxs = append(allTxs, tokenTxs...)
			}
		}

		if len(allTxs) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("no transactions generated; all balances below minSweepAmount")
		}

		batch := mcms_types.BatchOperation{
			ChainSelector: mcms_types.ChainSelector(chain.Selector),
			Transactions:  allTxs,
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)
