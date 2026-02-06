package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// formatOutput transforms the current state.json structure to match goal_state.json
// This provides backward compatibility for downstream consumers
func formatOutput(stateJSON []byte, chainInfos map[uint64]*ChainInfo) ([]byte, error) {
	// Parse the current state
	var state map[string]any
	if err := json.Unmarshal(stateJSON, &state); err != nil {
		return nil, fmt.Errorf("failed to parse state: %w", err)
	}

	// Build the legacy structure
	legacy := make(map[string]any)

	// Process chains grouped by family
	evmChains := make(map[string]any)
	solChains := make(map[string]any)
	aptosChains := make(map[string]any)

	if chains, ok := state["chains"].(map[string]any); ok {
		for selectorStr, contracts := range chains {
			contractList, ok := contracts.([]any)
			if !ok {
				continue
			}

			// Get chain info to determine family and name
			var selector uint64
			fmt.Sscanf(selectorStr, "%d", &selector)

			chainInfo := chainInfos[selector]
			chainName := getChainName(selector, chainInfo)
			chainID := getChainID(selector, chainInfo)
			family := "evm"
			if chainInfo != nil {
				family = string(chainInfo.Family)
			}

			// Build the chain object
			chainObj := buildChainObject(selector, chainID, contractList)

			// Add to appropriate family
			switch family {
			case "svm":
				solChains[chainName] = chainObj
			case "aptos":
				aptosChains[chainName] = chainObj
			default:
				evmChains[chainName] = chainObj
			}
		}
	}

	// Add chains to legacy structure (sorted for consistent output)
	if len(evmChains) > 0 {
		legacy["chains"] = evmChains
	}
	if len(solChains) > 0 {
		legacy["solChains"] = solChains
	}
	if len(aptosChains) > 0 {
		legacy["aptosChains"] = aptosChains
	}

	// NOPs placeholder - we don't have this data from on-chain sources
	// The CapabilitiesRegistry provides partial data but not the full NOP structure
	// legacy["nops"] = make(map[string]any)

	// Marshal with sorted keys for deterministic output
	return json.MarshalIndent(legacy, "", "    ")
}

// buildChainObject creates the legacy chain structure from a list of contracts
func buildChainObject(chainSelector uint64, chainID string, contracts []any) map[string]any {
	chainObj := map[string]any{
		"chainSelector": chainSelector,
		"chainID":       chainID,
	}

	// Group contracts by category
	categories := map[string]map[string]any{
		"rmnProxy":           make(map[string]any),
		"router":             make(map[string]any),
		"tokenAdminRegistry": make(map[string]any),
		"tokenPoolFactory":   make(map[string]any),
		"registryModules":    make(map[string]any),
		"feeQuoter":          make(map[string]any),
		"nonceManager":       make(map[string]any),
		"rmnRemote":          make(map[string]any),
		"onRamp":             make(map[string]any),
		"offRamp":            make(map[string]any),
		"linkToken":          make(map[string]any),
		"staticLinkToken":    make(map[string]any),
		"ccipHome":           make(map[string]any),
		"rmnHome":            make(map[string]any),
		"priceRegistry":      make(map[string]any),
		"commitStore":        make(map[string]any),
		"rmn":                make(map[string]any),
		"committeeVerifier":  make(map[string]any),
	}

	// MCMS contracts need special handling
	mcmsContracts := make(map[string]any)

	// Token pools grouped by token
	poolsByToken := make(map[string]map[string]any)

	// Aptos-specific
	ccipModule := make(map[string]any)

	for _, c := range contracts {
		contract, ok := c.(map[string]any)
		if !ok {
			continue
		}

		contractType, _ := contract["_type"].(string)
		address, _ := contract["address"].(string)
		if address == "" {
			// Try to get from top level for some contract types
			if addr, ok := contract["address"].(string); ok {
				address = addr
			}
		}

		// Remove internal metadata fields from the output
		cleanContract := cleanContractData(contract)

		// Categorize the contract
		category := categorizeContract(contractType)

		switch category {
		case "pool":
			// Group pools by token
			tokenName := getTokenNameFromContract(contract)
			if poolsByToken[tokenName] == nil {
				poolsByToken[tokenName] = make(map[string]any)
			}
			poolsByToken[tokenName][address] = cleanContract

		case "mcms":
			// MCMS contracts go into mcmsWithTimelock
			mcmsKey := getMCMSKey(contractType)
			if mcmsKey != "" {
				mcmsContracts[mcmsKey] = cleanContract
			}

		case "ccip":
			// Aptos CCIP module
			ccipModule[address] = cleanContract

		default:
			// Regular categorized contract
			if cat, ok := categories[category]; ok {
				cat[address] = cleanContract
			}
		}
	}

	// Add non-empty categories to chain object
	for category, contracts := range categories {
		if len(contracts) > 0 {
			chainObj[category] = contracts
		}
	}

	// Add pools by token if any
	if len(poolsByToken) > 0 {
		chainObj["poolByTokens"] = poolsByToken
	}

	// Add MCMS with timelock if any
	if len(mcmsContracts) > 0 {
		chainObj["mcmsWithTimelock"] = mcmsContracts
	}

	// Add CCIP module for Aptos
	if len(ccipModule) > 0 {
		chainObj["ccip"] = ccipModule
	}

	return chainObj
}

// categorizeContract determines which category a contract belongs to
func categorizeContract(contractType string) string {
	// Normalize the type
	t := strings.ToLower(contractType)

	// Pool types
	if strings.Contains(t, "tokenpool") || strings.Contains(t, "pool") && !strings.Contains(t, "factory") {
		return "pool"
	}

	// MCMS types
	if strings.Contains(t, "mcms") || strings.Contains(t, "multisig") ||
		strings.Contains(t, "timelock") || strings.Contains(t, "callproxy") {
		return "mcms"
	}

	// Direct mappings
	typeMap := map[string]string{
		"armproxy":                  "rmnProxy",
		"rmnproxy":                  "rmnProxy",
		"router":                    "router",
		"tokenadminregistry":        "tokenAdminRegistry",
		"tokenpoolfactory":          "tokenPoolFactory",
		"registrymoduleownercustom": "registryModules",
		"feequoter":                 "feeQuoter",
		"noncemanager":              "nonceManager",
		"rmnremote":                 "rmnRemote",
		"onramp":                    "onRamp",
		"offramp":                   "offRamp",
		"linktoken":                 "linkToken",
		"staticlinktoken":           "staticLinkToken",
		"cciphome":                  "ccipHome",
		"rmnhome":                   "rmnHome",
		"priceregistry":             "priceRegistry",
		"commitstore":               "commitStore",
		"rmn":                       "rmn",
		"committeeverifier":         "committeeVerifier",
		// Aptos types
		"aptosccip":    "ccip",
		"aptosrouter":  "router",
		"aptosonramp":  "onRamp",
		"aptosofframp": "offRamp",
	}

	for pattern, category := range typeMap {
		if strings.Contains(t, pattern) {
			return category
		}
	}

	return ""
}

// getMCMSKey returns the key for MCMS contracts
func getMCMSKey(contractType string) string {
	t := strings.ToLower(contractType)

	if strings.Contains(t, "bypasser") {
		if strings.Contains(t, "multisig") {
			return "bypasser"
		}
		return "bypasserAccessController"
	}
	if strings.Contains(t, "canceller") {
		if strings.Contains(t, "multisig") {
			return "canceller"
		}
		return "cancellerAccessController"
	}
	if strings.Contains(t, "proposer") {
		if strings.Contains(t, "multisig") {
			return "proposer"
		}
		return "proposerAccessController"
	}
	if strings.Contains(t, "timelock") {
		return "timelock"
	}
	if strings.Contains(t, "callproxy") {
		return "callProxy"
	}

	return ""
}

// getTokenNameFromContract attempts to extract a token name for pool grouping
func getTokenNameFromContract(contract map[string]any) string {
	// First, check if we have the symbol directly from on-chain call
	if symbol, ok := contract["symbol"].(string); ok && symbol != "" {
		return symbol
	}

	// Try to get from qualifier which often contains token info
	if qualifier, ok := contract["_qualifier"].(string); ok {
		// Format: "address-Type" or might contain token name
		parts := strings.Split(qualifier, "-")
		if len(parts) > 1 {
			// Check if any part looks like a token name
			for _, part := range parts[1:] {
				if isLikelyTokenName(part) {
					return part
				}
			}
		}
	}

	// Try to get token address and use it as identifier
	if token, ok := contract["token"].(string); ok && token != "" {
		// Use abbreviated token address as fallback name
		if len(token) > 10 {
			return token[:10] + "..."
		}
		return token
	}

	return "unknown"
}

// isLikelyTokenName checks if a string looks like a token name
func isLikelyTokenName(s string) bool {
	// Token names are usually short uppercase strings
	if len(s) < 2 || len(s) > 20 {
		return false
	}

	// Common token name patterns
	knownTokens := []string{
		"LINK", "WETH", "USDC", "USDT", "DAI", "WBTC", "ETH",
		"BnM", "LnM", "CCIP-BnM", "CCIP-LnM",
	}

	for _, t := range knownTokens {
		if strings.EqualFold(s, t) {
			return true
		}
	}

	// Check if it's a typical token pool type that shouldn't be used as name
	poolTypes := []string{
		"BurnMintTokenPool", "LockReleaseTokenPool", "TokenPool",
		"Pool", "Factory", "Registry",
	}

	for _, pt := range poolTypes {
		if strings.Contains(s, pt) {
			return false
		}
	}

	return false
}

// cleanContractData removes internal metadata fields and normalizes the contract data
func cleanContractData(contract map[string]any) map[string]any {
	clean := make(map[string]any)

	// Fields to exclude from output
	excludeFields := map[string]bool{
		"_qualifier":        true,
		"_requestedVersion": true,
		"_type":             true,
	}

	// Copy all non-excluded fields
	for k, v := range contract {
		if !excludeFields[k] {
			clean[k] = v
		}
	}

	// Add typeAndVersion if we have the data
	if contractType, ok := contract["_type"].(string); ok {
		if version, ok := contract["_requestedVersion"].(string); ok {
			clean["typeAndVersion"] = fmt.Sprintf("%s %s", contractType, version)
		}
	}

	return clean
}

// getChainName returns a human-readable chain name
func getChainName(selector uint64, info *ChainInfo) string {
	// Check well-known chains first
	if name, ok := ChainSelectorToName[selector]; ok {
		return name
	}

	// Try to use the chain info name
	if info != nil && info.Name != "" {
		return info.Name
	}

	// Fall back to selector string
	return fmt.Sprintf("chain-%d", selector)
}

// getChainID returns the chain ID as a string
func getChainID(selector uint64, info *ChainInfo) string {
	if id, ok := ChainSelectorToChainID[selector]; ok {
		return id
	}
	return ""
}

// ChainSelectorToName maps chain selectors to human-readable names
// This is a comprehensive list based on known CCIP deployments
var ChainSelectorToName = map[uint64]string{
	// Mainnets
	5009297550715157269:  "ethereum-mainnet",
	4949039107694359620:  "arbitrum-mainnet",
	3734403246176062136:  "optimism-mainnet",
	4051577828743386545:  "polygon-mainnet",
	6433500567565415381:  "avalanche-mainnet",
	15971525489660198786: "base-mainnet",
	11344663589394136015: "bnb-mainnet",
	465200170687744372:   "gnosis-mainnet",
	1346049177634351622:  "wemix-mainnet",
	5142893604156789321:  "kroma-mainnet",
	3719320017875267166:  "celo-mainnet",
	7264351850409363825:  "metis-mainnet",
	4411394078118774322:  "moonbeam-mainnet",
	3777822886988675105:  "zksync-mainnet",
	1562403441176082196:  "scroll-mainnet",
	6916147374840168594:  "linea-mainnet",
	2442541497099098535:  "mantle-mainnet",
	8805746078405598895:  "blast-mainnet",
	11155420:             "mode-mainnet",
	4426351306075016396:  "0g-mainnet",
	7801139999541420232:  "xlayer-mainnet",
	1994725729866924022:  "boba-mainnet",
	7518481528524156446:  "ronin-mainnet",
	3577778157919314504:  "abstract-mainnet",
	9027416829622342829:  "sei-mainnet",
	4893902685675875560:  "apechain-mainnet",
	5765866055508860288:  "ink-mainnet",
	8481857512324358265:  "sonic-mainnet",
	11690709103138290329: "iotex-mainnet",
	12657445206920369324: "nexon-mainnet",
	1456215246176062136:  "zircuit-mainnet",
	6051068554852905209:  "soneium-mainnet",
	1479926114862168794:  "worldchain-mainnet",
	9284632837123596123:  "corn-mainnet",
	12433265701047361059: "lisk-mainnet",
	7717148896336251131:  "fraxtal-mainnet",
	5990477251245693094:  "unichain-mainnet",
	4741433654826277614:  "taiko-mainnet",
	1923510103922296319:  "immutablezkevm-mainnet",
	3016212468291539606:  "morph-mainnet",
	6180750485974931498:  "treasure-mainnet",
	1224752112135636129:  "berachain-mainnet",
	9435876195998811891:  "hyperliquid-mainnet",

	// Solana
	124615329519749607: "solana-mainnet",

	// Aptos
	4741433654826277615: "aptos-mainnet",

	// Testnets
	16015286601757825753: "ethereum-sepolia",
	14767482510784806043: "avalanche-fuji",
	10344971235874465080: "base-sepolia",
	3478487238524512106:  "arbitrum-sepolia",
	16281711391670634445: "polygon-amoy",
	16423721717087811551: "solana-devnet",
	743186221051783445:   "aptos-testnet",
}

// ChainSelectorToChainID maps chain selectors to chain IDs
var ChainSelectorToChainID = map[uint64]string{
	5009297550715157269:  "1",
	4949039107694359620:  "42161",
	3734403246176062136:  "10",
	4051577828743386545:  "137",
	6433500567565415381:  "43114",
	15971525489660198786: "8453",
	11344663589394136015: "56",
	465200170687744372:   "100",
	4426351306075016396:  "16661",
	3577778157919314504:  "2741",
	9027416829622342829:  "1329",
	124615329519749607:   "solana-mainnet",
	4741433654826277615:  "aptos-mainnet",
}

// SortedKeys returns sorted keys of a map for deterministic output
func SortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
