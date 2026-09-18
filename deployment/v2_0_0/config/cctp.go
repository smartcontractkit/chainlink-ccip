package config

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"
)

// CCTPChainDefaults holds the Circle-defined CCTP contract addresses and the
// Circle domain identifier for a chain. These values are sourced from Circle's
// public documentation:
//
//   - https://developers.circle.com/cctp/references/contract-addresses
//   - https://developers.circle.com/cctp/v1/evm-smart-contracts
//   - https://developers.circle.com/cctp/v1/solana-programs
//   - https://developers.circle.com/stablecoins/usdc-contract-addresses
//
// They are treated as fallbacks only: any value explicitly provided in a
// changeset input takes precedence. This allows test environments (which use
// their own chain selectors and mock USDC/CCTP contracts) to override them, or
// to be unaffected because no defaults exist for their chain selectors.
type CCTPChainDefaults struct {
	// DomainIdentifier is the Circle CCTP domain identifier for the chain.
	DomainIdentifier uint32
	// TokenMessengerV1 is the CCTP V1 TokenMessenger contract/program address.
	TokenMessengerV1 string
	// TokenMessengerV2 is the CCTP V2 TokenMessenger contract address.
	TokenMessengerV2 string
	// USDCToken is the canonical USDC token contract/mint address.
	USDCToken string
}

// CCTPChainDefaultsBySelector maps chain selectors to their Circle-defined CCTP
// defaults. Chains that are documented by Circle but not yet present in the
// pinned chain-selectors version (for example HyperEVM and Injective) are
// intentionally omitted; their inputs must be provided explicitly.
var CCTPChainDefaultsBySelector = map[uint64]CCTPChainDefaults{
	// --- EVM mainnets ---
	chain_selectors.ETHEREUM_MAINNET.Selector: {
		DomainIdentifier: 0,
		TokenMessengerV1: "0xBd3fa81B58Ba92a82136038B25aDec7066af3155",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
	},
	chain_selectors.AVALANCHE_MAINNET.Selector: {
		DomainIdentifier: 1,
		TokenMessengerV1: "0x6B25532e1060CE10cc3B0A99e5683b91BFDe6982",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E",
	},
	chain_selectors.ETHEREUM_MAINNET_OPTIMISM_1.Selector: {
		DomainIdentifier: 2,
		TokenMessengerV1: "0x2B4069517957735bE00ceE0fadAE88a26365528f",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
	},
	chain_selectors.ETHEREUM_MAINNET_ARBITRUM_1.Selector: {
		DomainIdentifier: 3,
		TokenMessengerV1: "0x19330d10D9Cc8751218eaf51E8885D058642E08A",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xaf88d065e77c8cC2239327C5EDb3A432268e5831",
	},
	chain_selectors.ETHEREUM_MAINNET_BASE_1.Selector: {
		DomainIdentifier: 6,
		TokenMessengerV1: "0x1682Ae6375C4E4A97e4B583BC394c861A46D8962",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
	},
	chain_selectors.POLYGON_MAINNET.Selector: {
		DomainIdentifier: 7,
		TokenMessengerV1: "0x9daF8c91AEFAE50b9c0E69629D3F6Ca40cA3B3FE",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359",
	},
	chain_selectors.ETHEREUM_MAINNET_UNICHAIN_1.Selector: {
		DomainIdentifier: 10,
		TokenMessengerV1: "0x4e744b28E787c3aD0e810eD65A24461D4ac5a762",
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x078D782b760474a361dDA0AF3839290b0EF57AD6",
	},
	chain_selectors.ETHEREUM_MAINNET_LINEA_1.Selector: {
		DomainIdentifier: 11,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x176211869cA2b568f2A7D4EE941E073a821EE1ff",
	},
	chain_selectors.CODEX_MAINNET.Selector: {
		DomainIdentifier: 12,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xd996633a415985DBd7D6D12f4A4343E31f5037cf",
	},
	chain_selectors.SONIC_MAINNET.Selector: {
		DomainIdentifier: 13,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x29219dd400f2Bf60E5a23d13Be72B486D4038894",
	},
	chain_selectors.ETHEREUM_MAINNET_WORLDCHAIN_1.Selector: {
		DomainIdentifier: 14,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x79A02482A880bCe3F13E09da970dC34dB4cD24D1",
	},
	chain_selectors.MONAD_MAINNET.Selector: {
		DomainIdentifier: 15,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x754704Bc059F8C67012fEd69BC8A327a5aafb603",
	},
	chain_selectors.SEI_MAINNET.Selector: {
		DomainIdentifier: 16,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xe15fC38F6D8c56aF07bbCBe3BAf5708A2Bf42392",
	},
	chain_selectors.XDC_MAINNET.Selector: {
		DomainIdentifier: 18,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xfA2958CB79b0491CC627c1557F441eF849Ca8eb1",
	},
	chain_selectors.ETHEREUM_MAINNET_INK_1.Selector: {
		DomainIdentifier: 21,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x2D270e6886d130D724215A266106e6832161EAEd",
	},
	chain_selectors.PLUME_MAINNET.Selector: {
		DomainIdentifier: 22,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x222365EF19F7947e5484218551B56bb3965Aa7aF",
	},
	chain_selectors.EDGE_MAINNET.Selector: {
		DomainIdentifier: 28,
		TokenMessengerV2: "0x98706A006bc632Df31CAdFCBD43F38887ce2ca5c",
		USDCToken:        "0x98d2919b9A214E6Fa5384AC81E6864bA686Ad74c",
	},
	chain_selectors.MORPH_MAINNET.Selector: {
		DomainIdentifier: 30,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xCfb1186F4e93D60E60a8bDd997427D1F33bc372B",
	},
	chain_selectors.PHAROS_MAINNET.Selector: {
		DomainIdentifier: 31,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xC879C018dB60520F4355C26eD1a6D572cdAC1815",
	},
	chain_selectors.CRONOS_MAINNET.Selector: {
		DomainIdentifier: 32,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x3D7F2C478aAfdB65542BCB44bCeeC05849999d2D",
	},
	chain_selectors.PLASMA_MAINNET.Selector: {
		DomainIdentifier: 33,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0x2d661C89D812261039AF9764eceaAee884f5F67F",
	},
	chain_selectors.ETHEREUM_MAINNET_XLAYER_1.Selector: {
		DomainIdentifier: 37,
		TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
		USDCToken:        "0xB6CEceAB302E2E4948951eE7843FC24E92933061",
	},

	// --- Solana mainnet/devnet ---
	chain_selectors.SOLANA_MAINNET.Selector: {
		DomainIdentifier: 5,
		TokenMessengerV1: "CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3",
		USDCToken:        "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
	},
	chain_selectors.SOLANA_DEVNET.Selector: {
		DomainIdentifier: 5,
		TokenMessengerV1: "CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3",
		USDCToken:        "4zMMC9srt5Ri5X14GAgXhaHii3GnPAEERYPJgZJDncDU",
	},

	// --- EVM testnets ---
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector: {
		DomainIdentifier: 0,
		TokenMessengerV1: "0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
	},
	chain_selectors.AVALANCHE_TESTNET_FUJI.Selector: {
		DomainIdentifier: 1,
		TokenMessengerV1: "0xeb08f243E5d3FCFF26A9E38Ae5520A669f4019d0",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x5425890298aed601595a70AB815c96711a31Bc65",
	},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_OPTIMISM_1.Selector: {
		DomainIdentifier: 2,
		TokenMessengerV1: "0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x5fd84259d66Cd46123540766Be93DFE6D43130D7",
	},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector: {
		DomainIdentifier: 3,
		TokenMessengerV1: "0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d",
	},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_BASE_1.Selector: {
		DomainIdentifier: 6,
		TokenMessengerV1: "0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x036CbD53842c5426634e7929541eC2318f3dCF7e",
	},
	chain_selectors.POLYGON_TESTNET_AMOY.Selector: {
		DomainIdentifier: 7,
		TokenMessengerV1: "0x9f3B8679c73C2Fef8b59B4f3444d4e156fb70AA5",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x41E94Eb019C0762f9Bfcf9Fb1E58725BfB0e7582",
	},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_UNICHAIN_1.Selector: {
		DomainIdentifier: 10,
		TokenMessengerV1: "0x8ed94B8dAd2Dc5453862ea5e316A8e71AAed9782",
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x31d0220469e10c4E71834a79b1f276d740d3768F",
	},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_LINEA_1.Selector: {
		DomainIdentifier: 11,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xFEce4462D57bD51A6A552365A011b95f0E16d9B7",
	},
	chain_selectors.CODEX_TESTNET.Selector: {
		DomainIdentifier: 12,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x6d7f141b6819C2c9CC2f818e6ad549E7Ca090F8f",
	},
	chain_selectors.SONIC_TESTNET_BLAZE.Selector: {
		DomainIdentifier: 13,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xA4879Fed32Ecbef99399e5cbC247E533421C4eC6",
	},
	chain_selectors.ETHEREUM_TESTNET_SEPOLIA_WORLDCHAIN_1.Selector: {
		DomainIdentifier: 14,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x66145f38cBAC35Ca6F1Dfb4914dF98F1614aeA88",
	},
	chain_selectors.MONAD_TESTNET.Selector: {
		DomainIdentifier: 15,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x534b2f3A21130d7a60830c2Df862319e593943A3",
	},
	chain_selectors.SEI_TESTNET_ATLANTIC.Selector: {
		DomainIdentifier: 16,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x4fCF1784B31630811181f670Aea7A7bEF803eaED",
	},
	chain_selectors.XDC_TESTNET.Selector: {
		DomainIdentifier: 18,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xb5AB69F7bBada22B28e79C8FFAECe55eF1c771D4",
	},
	chain_selectors.INK_TESTNET_SEPOLIA.Selector: {
		DomainIdentifier: 21,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xFabab97dCE620294D2B0b0e46C68964e326300Ac",
	},
	chain_selectors.PLUME_TESTNET.Selector: {
		DomainIdentifier: 22,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xcB5f30e335672893c7eb944B374c196392C19D18",
	},
	chain_selectors.ARC_TESTNET.Selector: {
		DomainIdentifier: 26,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x3600000000000000000000000000000000000000",
	},
	chain_selectors.EDGE_TESTNET.Selector: {
		DomainIdentifier: 28,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x2d9F7CAD728051AA35Ecdc472a14cf8cDF5CFD6B",
	},
	chain_selectors.ETHEREUM_TESTNET_HOODI_MORPH.Selector: {
		DomainIdentifier: 30,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0x7433b41C6c5e1d58D4Da99483609520255ab661B",
	},
	chain_selectors.PHAROS_TESTNET.Selector: {
		DomainIdentifier: 31,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xcfC8330f4BCAB529c625D12781b1C19466A9Fc8B",
	},
	chain_selectors.CRONOS_TESTNET.Selector: {
		DomainIdentifier: 32,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xEb33dc5fac03833e132593659e1dE7256aB59794",
	},
	chain_selectors.PLASMA_TESTNET.Selector: {
		DomainIdentifier: 33,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xE67Fb267022cBA8064Dd388CC2FED724F3120D9D",
	},
	chain_selectors.XLAYER_TESTNET.Selector: {
		DomainIdentifier: 37,
		TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
		USDCToken:        "0xDec90b78111Ba2fc6FC6d84d8B9ec159A2d4b9B3",
	},
}

// GetCCTPChainDefaults returns the Circle-defined CCTP defaults for the given
// chain selector. The boolean return value indicates whether defaults exist.
func GetCCTPChainDefaults(chainSelector uint64) (CCTPChainDefaults, bool) {
	defaults, ok := CCTPChainDefaultsBySelector[chainSelector]
	return defaults, ok
}
