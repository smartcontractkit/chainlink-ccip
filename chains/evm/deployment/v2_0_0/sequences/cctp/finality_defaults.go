package cctp

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
)

// fastTransferSourceSelectors is the set of EVM chain selectors that Circle lists as
// Fast Transfer source chains for CCTP V2. For these chains we use BlockDepth 1 on the
// CCTPVerifier and CCTP-through-CCV pool; everything else falls back to
// wait-for-finality.
//
// Source: https://developers.circle.com/cctp/concepts/supported-chains-and-domains
// Reviewed 2026-05-11. Keep in sync with the "Source (Fast transfer)" column;
// non-EVM (Solana, Starknet, Stellar) and BNB Smart Chain (USYC-only, not USDC) are
// intentionally excluded.
var fastTransferSourceSelectors = map[uint64]struct{}{
	// Ethereum
	chain_selectors.ETHEREUM_MAINNET.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector: {},
	// Arbitrum
	chain_selectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector: {},
	// Base
	chain_selectors.ETHEREUM_MAINNET_BASE_1.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_BASE_1.Selector: {},
	// OP Mainnet
	chain_selectors.ETHEREUM_MAINNET_OPTIMISM_1.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1.Selector: {},
	// Linea
	chain_selectors.ETHEREUM_MAINNET_LINEA_1.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_LINEA_1.Selector: {},
	// Unichain
	chain_selectors.ETHEREUM_MAINNET_UNICHAIN_1.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_UNICHAIN_1.Selector: {},
	// World Chain
	chain_selectors.ETHEREUM_MAINNET_WORLDCHAIN_1.Selector:         {},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_WORLDCHAIN_1.Selector: {},
	// Ink
	chain_selectors.ETHEREUM_MAINNET_INK_1.Selector: {},
	chain_selectors.INK_TESTNET_SEPOLIA.Selector:    {},
	// Codex
	chain_selectors.CODEX_MAINNET.Selector: {},
	chain_selectors.CODEX_TESTNET.Selector: {},
	// EDGE
	chain_selectors.EDGE_MAINNET.Selector: {},
	chain_selectors.EDGE_TESTNET.Selector: {},
	// Morph
	chain_selectors.MORPH_MAINNET.Selector: {},
	// Plume
	chain_selectors.PLUME_MAINNET.Selector:         {},
	chain_selectors.PLUME_TESTNET_SEPOLIA.Selector: {},

	// Test selectors - used in E2E tests
	chain_selectors.GETH_TESTNET.Selector:  {}, // chain ID 1337
	chain_selectors.GETH_DEVNET_2.Selector: {}, // chain ID 2337
	chain_selectors.GETH_DEVNET_3.Selector: {}, // chain ID 3337
}

// defaultAllowedFinalityForChain returns BlockDepth 1 when the source chain is in
// fastTransferSourceSelectors, otherwise wait-for-finality (on-chain 0x00).
func defaultAllowedFinalityForChain(chainSelector uint64) finality.Config {
	if _, ok := fastTransferSourceSelectors[chainSelector]; ok {
		return finality.Config{BlockDepth: 1}
	}
	return finality.Config{WaitForFinality: true}
}
