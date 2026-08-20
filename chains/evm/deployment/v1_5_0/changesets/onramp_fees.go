package changesets

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	fee_quoter_binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evm_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	fee_quoter_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	v1_6_0_sequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/weth9"
)

// =============================================================================
// Configuration Changeset
// =============================================================================

// ConfigureFeeSweepV150Cfg configures an OnRamp to send all LINK fees to treasury
type ConfigureFeeSweepV150Cfg struct {
	ChainSel      uint64         `json:"chainSelector" yaml:"chainSelector"`
	OnRampAddress common.Address `json:"onRampAddress" yaml:"onRampAddress"`
	Treasury      common.Address `json:"treasury" yaml:"treasury"`
}

func (c ConfigureFeeSweepV150Cfg) ChainSelector() uint64 {
	return c.ChainSel
}

// ConfigureFeeSweepV150 returns a changeset that configures the OnRamp to send all LINK fees to treasury.
func ConfigureFeeSweepV150(allowedRecipients map[uint64]common.Address) func(*changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[ConfigureFeeSweepV150Cfg]] {
	return changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
		sequences.ConfigureFeeSweepV150Input,
		evm.Chain,
		ConfigureFeeSweepV150Cfg,
	]{
		Sequence: sequences.ConfigureFeeSweepV150Sequence,
		ResolveInput: func(e cldf_deployment.Environment, cfg ConfigureFeeSweepV150Cfg) (sequences.ConfigureFeeSweepV150Input, error) {
			return sequences.ConfigureFeeSweepV150Input{
				ChainSelector:     cfg.ChainSel,
				OnRampAddress:     cfg.OnRampAddress,
				Treasury:          cfg.Treasury,
				AllowedRecipients: allowedRecipients,
			}, nil
		},
		ResolveDep: evm_sequences.ResolveEVMChainDep[ConfigureFeeSweepV150Cfg],
	})
}

// =============================================================================
// LINK Sweep Changeset
// =============================================================================

// SweepLinkFeesV150Cfg triggers LINK fee payout with safety validation
type SweepLinkFeesV150Cfg struct {
	ChainSel         uint64         `json:"chainSelector" yaml:"chainSelector"`
	OnRampAddress    common.Address `json:"onRampAddress" yaml:"onRampAddress"`
	ExpectedTreasury common.Address `json:"expectedTreasury" yaml:"expectedTreasury"`
}

func (c SweepLinkFeesV150Cfg) ChainSelector() uint64 {
	return c.ChainSel
}

// SweepLinkFeesV150 returns a changeset that sweeps accumulated LINK fees to the configured NOPs.
func SweepLinkFeesV150(allowedRecipients map[uint64]common.Address) func(*changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[SweepLinkFeesV150Cfg]] {
	return changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
		sequences.SweepLinkFeesV150Input,
		evm.Chain,
		SweepLinkFeesV150Cfg,
	]{
		Sequence: sequences.SweepLinkFeesV150Sequence,
		ResolveInput: func(e cldf_deployment.Environment, cfg SweepLinkFeesV150Cfg) (sequences.SweepLinkFeesV150Input, error) {
			return sequences.SweepLinkFeesV150Input{
				ChainSelector:     cfg.ChainSel,
				OnRampAddress:     cfg.OnRampAddress,
				ExpectedTreasury:  cfg.ExpectedTreasury,
				AllowedRecipients: allowedRecipients,
			}, nil
		},
		ResolveDep: evm_sequences.ResolveEVMChainDep[SweepLinkFeesV150Cfg],
	})
}

// =============================================================================
// Shared helpers
// =============================================================================

// resolveWETH9AndTimelock looks up WETH9 and RBACTimelock from the datastore.
// Returns (zero, zero, nil) if WETH9 is not registered. Returns an error only
// if WETH9 is found but the RBACTimelock cannot be resolved.
func resolveWETH9AndTimelock(
	ds datastore.DataStore,
	chainSel uint64,
	timelockQualifier string,
) (weth9Addr, timelockAddr common.Address, err error) {
	resolved, err := datastore_utils.FindAndFormatRef(
		ds,
		datastore.AddressRef{Type: datastore.ContractType("WETH9")},
		chainSel,
		evm_datastore_utils.ToEVMAddress,
	)
	if err != nil {
		// WETH9 not registered — caller should treat all tokens as plain ERC20
		return common.Address{}, common.Address{}, nil
	}
	weth9Addr = resolved

	timelockAddr, err = datastore_utils.FindAndFormatRef(
		ds,
		datastore.AddressRef{
			Type:      datastore.ContractType("RBACTimelock"),
			Qualifier: timelockQualifier,
		},
		chainSel,
		evm_datastore_utils.ToEVMAddress,
	)
	if err != nil {
		return common.Address{}, common.Address{}, fmt.Errorf(
			"RBACTimelock not found for chain %d (qualifier: %q); needed for WETH unwrap: %w",
			chainSel, timelockQualifier, err)
	}
	return weth9Addr, timelockAddr, nil
}

// resolveFeeQuoterNonLinkTokens resolves non-LINK fee token addresses from the
// on-chain FeeQuoter contract. Returns an error if no FeeQuoter is found in the
// datastore. Returns an empty slice if all fee tokens are LINK.
func resolveFeeQuoterNonLinkTokens(
	ds datastore.DataStore,
	chainSel uint64,
	chain evm.Chain,
	callOpts *bind.CallOpts,
) ([]common.Address, error) {
	// Find FeeQuoter in datastore (picks latest version < 1.7.0)
	refs := ds.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(fee_quoter_ops.ContractType)),
		datastore.AddressRefByChainSelector(chainSel),
	)
	ref, err := v1_6_0_sequences.GetFeeQuoterAddress(refs, chainSel)
	if err != nil {
		return nil, fmt.Errorf("no FeeQuoter found for chain %d; ensure chain has v1.6+ contracts: %w", chainSel, err)
	}
	fqAddr := common.HexToAddress(ref.Address)

	// Create FeeQuoter binding and query on-chain state
	fq, err := fee_quoter_binding.NewFeeQuoter(fqAddr, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to create FeeQuoter binding for %s: %w", fqAddr.Hex(), err)
	}

	feeTokens, err := fq.GetFeeTokens(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get fee tokens from FeeQuoter %s: %w", fqAddr.Hex(), err)
	}

	staticCfg, err := fq.GetStaticConfig(callOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to get static config from FeeQuoter %s: %w", fqAddr.Hex(), err)
	}

	// Filter out LINK, return remaining non-LINK tokens.
	// Fail on unexpected data — don't silently fix bad on-chain state.
	seen := make(map[common.Address]bool)
	var nonLink []common.Address
	for _, token := range feeTokens {
		if token == (common.Address{}) {
			return nil, fmt.Errorf("FeeQuoter %s returned zero address in fee tokens list", fqAddr.Hex())
		}
		if seen[token] {
			return nil, fmt.Errorf("FeeQuoter %s returned duplicate fee token %s", fqAddr.Hex(), token.Hex())
		}
		seen[token] = true
		if token != staticCfg.LinkToken {
			nonLink = append(nonLink, token)
		}
	}
	return nonLink, nil
}

// =============================================================================
// Non-LINK Sweep Changeset (with WETH auto-detection)
// =============================================================================

// SweepNonLinkFeesV150Cfg withdraws non-LINK fee tokens to treasury.
// If any fee token is WETH9 (detected from datastore), automatically does atomic unwrap.
type SweepNonLinkFeesV150Cfg struct {
	ChainSel          uint64           `json:"chainSelector" yaml:"chainSelector"`
	OnRampAddress     common.Address   `json:"onRampAddress" yaml:"onRampAddress"`
	FeeTokens         []common.Address `json:"feeTokens" yaml:"feeTokens"`
	Treasury          common.Address   `json:"treasury" yaml:"treasury"`
	TimelockQualifier string           `json:"timelockQualifier" yaml:"timelockQualifier"` // Required for WETH unwrap
}

func (c SweepNonLinkFeesV150Cfg) ChainSelector() uint64 {
	return c.ChainSel
}

// SweepNonLinkFeesV150 returns a changeset that withdraws non-LINK fee tokens to treasury.
// Automatically detects WETH9 from datastore and does atomic unwrap if a fee token matches.
func SweepNonLinkFeesV150(allowedRecipients map[uint64]common.Address) func(*changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[SweepNonLinkFeesV150Cfg]] {
	return changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
		sequences.SweepNonLinkFeesV150Input,
		evm.Chain,
		SweepNonLinkFeesV150Cfg,
	]{
		Sequence: sequences.SweepNonLinkFeesV150Sequence,
		ResolveInput: func(e cldf_deployment.Environment, cfg SweepNonLinkFeesV150Cfg) (sequences.SweepNonLinkFeesV150Input, error) {
			result := sequences.SweepNonLinkFeesV150Input{
				ChainSelector:     cfg.ChainSel,
				OnRampAddress:     cfg.OnRampAddress,
				FeeTokens:         cfg.FeeTokens,
				Treasury:          cfg.Treasury,
				AllowedRecipients: allowedRecipients,
			}

			// Resolve WETH9 + RBACTimelock (graceful: zero addresses if WETH9 not registered)
			weth9Addr, timelockAddr, err := resolveWETH9AndTimelock(e.DataStore, cfg.ChainSel, cfg.TimelockQualifier)
			if err != nil {
				return sequences.SweepNonLinkFeesV150Input{}, err
			}
			result.WETH9Address = weth9Addr
			result.MCMSAddress = timelockAddr

			// If WETH9 is not among fee tokens, no unwrap needed — skip balance query
			hasWETH := false
			for _, token := range cfg.FeeTokens {
				if token == weth9Addr {
					hasWETH = true
					break
				}
			}
			if !hasWETH || weth9Addr == (common.Address{}) {
				return result, nil
			}

			// Query WETH balance of OnRamp
			chain, ok := e.BlockChains.EVMChains()[cfg.ChainSel]
			if !ok {
				return sequences.SweepNonLinkFeesV150Input{}, fmt.Errorf(
					"chain %d not found in environment", cfg.ChainSel)
			}
			wethContract, err := weth9.NewWETH9(weth9Addr, chain.Client)
			if err != nil {
				return sequences.SweepNonLinkFeesV150Input{}, fmt.Errorf(
					"failed to create WETH9 binding: %w", err)
			}
			// NOTE: WETH balance is queried at resolution time. If the balance changes
			// between resolution and MCMS proposal execution (e.g., more fees accumulate
			// or a concurrent proposal withdraws), the unwrap step may fail because it
			// uses this exact amount. The batch reverts atomically — no fund loss.
			// To fix: re-resolve and re-submit the proposal.
			balance, err := wethContract.BalanceOf(&bind.CallOpts{Context: e.GetContext()}, cfg.OnRampAddress)
			if err != nil {
				return sequences.SweepNonLinkFeesV150Input{}, fmt.Errorf(
					"failed to query WETH balance for OnRamp %s: %w", cfg.OnRampAddress.Hex(), err)
			}
			result.WETHBalance = balance

			return result, nil
		},
		ResolveDep: evm_sequences.ResolveEVMChainDep[SweepNonLinkFeesV150Cfg],
	})
}

// =============================================================================
// Chain-Wide Mega Flow Changeset
// =============================================================================

// SweepAllOnRampsV150Cfg is the user config for the mega flow: configure NOPs,
// sweep LINK, sweep non-LINK (with WETH auto-unwrap) for ALL OnRamps on a chain.
type SweepAllOnRampsV150Cfg struct {
	ChainSel          uint64         `json:"chainSelector" yaml:"chainSelector"`
	Treasury          common.Address `json:"treasury" yaml:"treasury"`
	TimelockQualifier string         `json:"timelockQualifier" yaml:"timelockQualifier"`
	MinSweepAmount    string         `json:"minSweepAmount" yaml:"minSweepAmount"` // Wei as string, "0" = sweep all
	// SkipNopsCheck skips the on-chain NOP configuration pre-check.
	// Default false. Set true for dry-run/testing when NOPs haven't been
	// configured on-chain yet. The SetNops tx is still included in the
	// batch regardless — this only skips the fail-fast verification.
	SkipNopsCheck bool `json:"skipNopsCheck,omitempty" yaml:"skipNopsCheck,omitempty"`
}

func (c SweepAllOnRampsV150Cfg) ChainSelector() uint64 {
	return c.ChainSel
}

// SweepAllOnRampsV150 creates the mega flow changeset.
// Non-LINK fee tokens are resolved dynamically from the on-chain FeeQuoter contract.
// WETH9 is resolved from datastore.
func SweepAllOnRampsV150(
	allowedRecipients map[uint64]common.Address,
) func(*changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[changesets.WithMCMS[SweepAllOnRampsV150Cfg]] {
	return changesets.NewFromOnChainSequence(changesets.NewFromOnChainSequenceParams[
		sequences.SweepAllOnRampsV150Input,
		evm.Chain,
		SweepAllOnRampsV150Cfg,
	]{
		Sequence: sequences.SweepAllOnRampsV150Sequence,
		ResolveInput: func(e cldf_deployment.Environment, cfg SweepAllOnRampsV150Cfg) (sequences.SweepAllOnRampsV150Input, error) {
			// Parse MinSweepAmount
			minSweepAmount := big.NewInt(0)
			if cfg.MinSweepAmount != "" {
				var ok bool
				minSweepAmount, ok = new(big.Int).SetString(cfg.MinSweepAmount, 10)
				if !ok {
					return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf("invalid minSweepAmount: %s", cfg.MinSweepAmount)
				}
			}

			// Get EVM chain for RPC calls
			chain, ok := e.BlockChains.EVMChains()[cfg.ChainSel]
			if !ok {
				return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
					"chain %d not found in environment", cfg.ChainSel)
			}

			// Resolve non-LINK fee tokens from on-chain FeeQuoter
			tokens, err := resolveFeeQuoterNonLinkTokens(
				e.DataStore, cfg.ChainSel, chain,
				&bind.CallOpts{Context: e.GetContext()},
			)
			if err != nil {
				return sequences.SweepAllOnRampsV150Input{}, err
			}

			// Resolve WETH9 + RBACTimelock (graceful: zero addresses if WETH9 not registered)
			weth9Addr, timelockAddr, err := resolveWETH9AndTimelock(e.DataStore, cfg.ChainSel, cfg.TimelockQualifier)
			if err != nil {
				return sequences.SweepAllOnRampsV150Input{}, err
			}

			// Discover OnRamps from datastore
			onRampRefs := e.DataStore.Addresses().Filter(
				datastore.AddressRefByType(datastore.ContractType(onramp.ContractType)),
				datastore.AddressRefByVersion(onramp.Version),
				datastore.AddressRefByChainSelector(cfg.ChainSel),
			)
			if len(onRampRefs) == 0 {
				return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
					"no v1.5.0 OnRamps found for chain %d", cfg.ChainSel)
			}

			// Convert refs to addresses
			onRamps := make([]common.Address, 0, len(onRampRefs))
			for _, ref := range onRampRefs {
				if !common.IsHexAddress(ref.Address) {
					return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
						"invalid OnRamp address in datastore: %s", ref.Address)
				}
				onRamps = append(onRamps, common.HexToAddress(ref.Address))
			}

			// Query balances for all OnRamps.
			// NOTE: All balances (LINK fees and token balances including WETH) are
			// queried at resolution time. If balances change between resolution and
			// MCMS proposal execution, the WETH unwrap step may fail because it uses
			// exact amounts. The batch reverts atomically — no fund loss.
			// To fix: re-resolve and re-submit the proposal.
			onRampTokenBalances := make(map[common.Address]map[common.Address]*big.Int)
			onRampLINKFees := make(map[common.Address]*big.Int)

			for _, onRampAddr := range onRamps {
				// Query LINK fees
				onRampContract, err := evm_2_evm_onramp.NewEVM2EVMOnRamp(onRampAddr, chain.Client)
				if err != nil {
					return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
						"failed to create OnRamp binding for %s: %w", onRampAddr.Hex(), err)
				}
				linkFees, err := onRampContract.GetNopFeesJuels(&bind.CallOpts{Context: e.GetContext()})
				if err != nil {
					return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
						"failed to query LINK fees for OnRamp %s: %w", onRampAddr.Hex(), err)
				}
				onRampLINKFees[onRampAddr] = linkFees

				// Query non-LINK token balances. We reuse the WETH9 binding as a generic
				// ERC20 wrapper since it has balanceOf — any ERC20-compatible binding works.
				tokenBalances := make(map[common.Address]*big.Int)
				for _, token := range tokens {
					tokenContract, err := weth9.NewWETH9(token, chain.Client)
					if err != nil {
						return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
							"failed to create ERC20 binding for token %s: %w", token.Hex(), err)
					}
					balance, err := tokenContract.BalanceOf(&bind.CallOpts{Context: e.GetContext()}, onRampAddr)
					if err != nil {
						return sequences.SweepAllOnRampsV150Input{}, fmt.Errorf(
							"failed to query balance of token %s for OnRamp %s: %w",
							token.Hex(), onRampAddr.Hex(), err)
					}
					tokenBalances[token] = balance
				}
				onRampTokenBalances[onRampAddr] = tokenBalances
			}

			return sequences.SweepAllOnRampsV150Input{
				ChainSelector:       cfg.ChainSel,
				Treasury:            cfg.Treasury,
				OnRamps:             onRamps,
				NonLinkFeeTokens:    tokens,
				WETH9Address:        weth9Addr,
				MCMSAddress:         timelockAddr,
				OnRampTokenBalances: onRampTokenBalances,
				OnRampLINKFees:      onRampLINKFees,
				MinSweepAmount:      minSweepAmount,
				AllowedRecipients:   allowedRecipients,
				SkipNopsCheck:       cfg.SkipNopsCheck,
			}, nil
		},
		ResolveDep: evm_sequences.ResolveEVMChainDep[SweepAllOnRampsV150Cfg],
	})
}
