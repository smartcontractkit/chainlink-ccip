package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// formatOutput transforms the current state.json structure to match goal_state.json
func formatOutput(stateJSON []byte, chainInfos map[uint64]*ChainInfo) ([]byte, error) {
	var state map[string]any
	if err := json.Unmarshal(stateJSON, &state); err != nil {
		return nil, fmt.Errorf("failed to parse state: %w", err)
	}
	legacy := make(map[string]any)
	evmChains := make(map[string]any)
	solChains := make(map[string]any)
	aptosChains := make(map[string]any)

	if chains, ok := state["chains"].(map[string]any); ok {
		for selectorStr, contracts := range chains {
			contractList, ok := contracts.([]any)
			if !ok {
				continue
			}
			var selector uint64
			fmt.Sscanf(selectorStr, "%d", &selector)
			chainInfo := chainInfos[selector]
			chainName := getChainName(selector, chainInfo)
			chainID := getChainID(selector, chainInfo)
			family := "evm"
			if chainInfo != nil {
				family = string(chainInfo.Family)
			}
			chainObj := buildChainObject(selector, chainID, contractList)
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
	if len(evmChains) > 0 {
		legacy["chains"] = evmChains
	}
	if len(solChains) > 0 {
		legacy["solChains"] = solChains
	}
	if len(aptosChains) > 0 {
		legacy["aptosChains"] = aptosChains
	}
	return json.MarshalIndent(legacy, "", "    ")
}

func buildChainObject(chainSelector uint64, chainID string, contracts []any) map[string]any {
	chainObj := map[string]any{
		"chainSelector": chainSelector,
		"chainID":       chainID,
	}
	categories := map[string]map[string]any{
		"rmnProxy": make(map[string]any), "router": make(map[string]any),
		"tokenAdminRegistry": make(map[string]any), "tokenPoolFactory": make(map[string]any),
		"registryModules": make(map[string]any), "feeQuoter": make(map[string]any),
		"nonceManager": make(map[string]any), "rmnRemote": make(map[string]any),
		"onRamp": make(map[string]any), "offRamp": make(map[string]any),
		"linkToken": make(map[string]any), "staticLinkToken": make(map[string]any),
		"ccipHome": make(map[string]any), "rmnHome": make(map[string]any),
		"priceRegistry": make(map[string]any), "commitStore": make(map[string]any),
		"rmn": make(map[string]any), "committeeVerifier": make(map[string]any),
	}
	mcmsContracts := make(map[string]any)
	poolsByToken := make(map[string]map[string]any)
	ccipModule := make(map[string]any)

	for _, c := range contracts {
		contract, ok := c.(map[string]any)
		if !ok {
			continue
		}
		contractType, _ := contract["_type"].(string)
		address, _ := contract["address"].(string)
		cleanContract := cleanContractData(contract)
		category := categorizeContract(contractType)
		switch category {
		case "pool":
			tokenName := getTokenNameFromContract(contract)
			if poolsByToken[tokenName] == nil {
				poolsByToken[tokenName] = make(map[string]any)
			}
			poolsByToken[tokenName][address] = cleanContract
		case "mcms":
			mcmsKey := getMCMSKey(contractType)
			if mcmsKey != "" {
				mcmsContracts[mcmsKey] = cleanContract
			}
		case "ccip":
			ccipModule[address] = cleanContract
		default:
			if cat, ok := categories[category]; ok {
				cat[address] = cleanContract
			}
		}
	}
	for category, contracts := range categories {
		if len(contracts) > 0 {
			chainObj[category] = contracts
		}
	}
	if len(poolsByToken) > 0 {
		chainObj["poolByTokens"] = poolsByToken
	}
	if len(mcmsContracts) > 0 {
		chainObj["mcmsWithTimelock"] = mcmsContracts
	}
	if len(ccipModule) > 0 {
		chainObj["ccip"] = ccipModule
	}
	return chainObj
}

func categorizeContract(contractType string) string {
	t := strings.ToLower(contractType)
	if strings.Contains(t, "tokenpool") || (strings.Contains(t, "pool") && !strings.Contains(t, "factory")) {
		return "pool"
	}
	if strings.Contains(t, "mcms") || strings.Contains(t, "multisig") ||
		strings.Contains(t, "timelock") || strings.Contains(t, "callproxy") {
		return "mcms"
	}
	typeMap := map[string]string{
		"armproxy": "rmnProxy", "rmnproxy": "rmnProxy", "router": "router",
		"tokenadminregistry": "tokenAdminRegistry", "tokenpoolfactory": "tokenPoolFactory",
		"registrymoduleownercustom": "registryModules", "feequoter": "feeQuoter",
		"noncemanager": "nonceManager", "rmnremote": "rmnRemote", "onramp": "onRamp",
		"offramp": "offRamp", "linktoken": "linkToken", "staticlinktoken": "staticLinkToken",
		"cciphome": "ccipHome", "rmnhome": "rmnHome", "priceregistry": "priceRegistry",
		"commitstore": "commitStore", "rmn": "rmn", "committeeverifier": "committeeVerifier",
		"aptosccip": "ccip", "aptosrouter": "router", "aptosonramp": "onRamp", "aptosofframp": "offRamp",
	}
	for pattern, category := range typeMap {
		if strings.Contains(t, pattern) {
			return category
		}
	}
	return ""
}

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

func getTokenNameFromContract(contract map[string]any) string {
	if symbol, ok := contract["symbol"].(string); ok && symbol != "" {
		return symbol
	}
	if qualifier, ok := contract["_qualifier"].(string); ok {
		parts := strings.Split(qualifier, "-")
		for _, part := range parts[1:] {
			if isLikelyTokenName(part) {
				return part
			}
		}
	}
	if token, ok := contract["token"].(string); ok && token != "" {
		if len(token) > 10 {
			return token[:10] + "..."
		}
		return token
	}
	return "unknown"
}

func isLikelyTokenName(s string) bool {
	if len(s) < 2 || len(s) > 20 {
		return false
	}
	knownTokens := []string{"LINK", "WETH", "USDC", "USDT", "DAI", "WBTC", "ETH", "BnM", "LnM", "CCIP-BnM", "CCIP-LnM"}
	for _, t := range knownTokens {
		if strings.EqualFold(s, t) {
			return true
		}
	}
	poolTypes := []string{"BurnMintTokenPool", "LockReleaseTokenPool", "TokenPool", "Pool", "Factory", "Registry"}
	for _, pt := range poolTypes {
		if strings.Contains(s, pt) {
			return false
		}
	}
	return false
}

func cleanContractData(contract map[string]any) map[string]any {
	excludeFields := map[string]bool{"_qualifier": true, "_requestedVersion": true, "_type": true}
	clean := make(map[string]any)
	for k, v := range contract {
		if !excludeFields[k] {
			clean[k] = v
		}
	}
	if contractType, ok := contract["_type"].(string); ok {
		if version, ok := contract["_requestedVersion"].(string); ok {
			clean["typeAndVersion"] = fmt.Sprintf("%s %s", contractType, version)
		}
	}
	return clean
}

func getChainName(selector uint64, info *ChainInfo) string {
	if name, ok := ChainSelectorToName[selector]; ok {
		return name
	}
	if info != nil && info.Name != "" {
		return info.Name
	}
	return fmt.Sprintf("chain-%d", selector)
}

func getChainID(selector uint64, info *ChainInfo) string {
	if id, ok := ChainSelectorToChainID[selector]; ok {
		return id
	}
	return ""
}

// ChainSelectorToName maps chain selectors to human-readable names (for live display and transform).
var ChainSelectorToName = map[uint64]string{
	5009297550715157269: "ethereum-mainnet", 4949039107694359620: "arbitrum-mainnet",
	3734403246176062136: "optimism-mainnet", 4051577828743386545: "polygon-mainnet",
	6433500567565415381: "avalanche-mainnet", 15971525489660198786: "base-mainnet",
	11344663589394136015: "bnb-mainnet", 465200170687744372: "gnosis-mainnet",
	1346049177634351622: "wemix-mainnet", 5142893604156789321: "kroma-mainnet",
	3719320017875267166: "celo-mainnet", 7264351850409363825: "metis-mainnet",
	4411394078118774322: "moonbeam-mainnet", 3777822886988675105: "zksync-mainnet",
	1562403441176082196: "scroll-mainnet", 6916147374840168594: "linea-mainnet",
	2442541497099098535: "mantle-mainnet", 8805746078405598895: "blast-mainnet",
	3577778157919314504: "abstract-mainnet", 9027416829622342829: "sei-mainnet",
	1479926114862168794: "worldchain-mainnet", 4741433654826277614: "taiko-mainnet",
	124615329519749607: "solana-mainnet", 4741433654826277615: "aptos-mainnet",
	16015286601757825753: "ethereum-sepolia", 14767482510784806043: "avalanche-fuji",
	10344971235874465080: "base-sepolia", 3478487238524512106: "arbitrum-sepolia",
	16281711391670634445: "polygon-amoy", 16423721717087811551: "solana-devnet",
	743186221051783445: "aptos-testnet",
}

var ChainSelectorToChainID = map[uint64]string{
	5009297550715157269: "1", 4949039107694359620: "42161", 3734403246176062136: "10",
	4051577828743386545: "137", 6433500567565415381: "43114", 15971525489660198786: "8453",
	11344663589394136015: "56", 465200170687744372: "100", 124615329519749607: "solana-mainnet",
	4741433654826277615: "aptos-mainnet",
}

func SortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
