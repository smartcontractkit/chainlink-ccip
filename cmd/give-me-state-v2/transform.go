package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// formatOutput transforms the current state.json structure to match goal_state.json.
// allChainSelectors is the set of chain selectors from the address refs (used for NOP filtering).
func formatOutput(stateJSON []byte, chainInfos map[uint64]*ChainInfo, allChainSelectors []uint64) ([]byte, error) {
	// Use json.Decoder with UseNumber() to preserve uint64 chain selector precision.
	// Default json.Unmarshal converts all numbers to float64, which truncates values > 2^53.
	var state map[string]any
	dec := json.NewDecoder(bytes.NewReader(stateJSON))
	dec.UseNumber()
	if err := dec.Decode(&state); err != nil {
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

	// Transform NOPs section: filter to relevant nodes, reshape to goal_state.json format.
	if nodeOps, ok := state["nodeOperators"].(map[string]any); ok {
		nops := transformNOPs(nodeOps, allChainSelectors, chainInfos)
		if nops != nil {
			legacy["nops"] = nops

			// Enrich committeeVerifier signer addresses in-place with NOP name + CSA key.
			signerToNOP := buildGlobalSignerToNOP(nops)
			enrichCommitteeVerifierSigners(legacy, signerToNOP)

			// Build per-network diff: JD vs on-chain committeeVerifier signers.
			nopsDiff := buildNopsDiff(legacy, nops)
			if nopsDiff != nil {
				legacy["nopsDiff"] = nopsDiff
			}
		}
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
		"capabilitiesRegistry": make(map[string]any),
		"priceRegistry":        make(map[string]any), "commitStore": make(map[string]any),
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
		"cciphome": "ccipHome", "rmnhome": "rmnHome", "capabilitiesregistry": "capabilitiesRegistry",
		"priceregistry": "priceRegistry", "commitstore": "commitStore", "rmn": "rmn",
		"committeeverifier": "committeeVerifier",
		"aptosccip":         "ccip", "aptosrouter": "router", "aptosonramp": "onRamp", "aptosofframp": "offRamp",
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
	// Don't fall back to info.Name -- it often holds the RPC provider name (e.g. "Alchemy")
	return fmt.Sprintf("chain-%d", selector)
}

func getChainID(selector uint64, info *ChainInfo) string {
	if id, ok := ChainSelectorToChainID[selector]; ok {
		return id
	}
	return ""
}

// ChainSelectorToName maps chain selectors to human-readable names (for live display and transform).
// Sourced from mainnet.yaml comments and well-known testnets.
var ChainSelectorToName = map[uint64]string{
	// --- Mainnets (from mainnet.yaml) ---
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
	2442541497099098535:  "mantle-mainnet",
	8805746078405598895:  "blast-mainnet",
	3577778157919314504:  "abstract-mainnet",
	9027416829622342829:  "sei-mainnet",
	4741433654826277614:  "aptos-mainnet",
	124615329519749607:   "solana-mainnet",
	3849287863852499584:  "bob-mainnet",
	4560701533377838164:  "botanix-mainnet",
	1456215246176062136:  "cronos-mainnet",
	8788096068760390840:  "cronos-zkevm-mainnet",
	13624601974233774587: "etherlink-mainnet",
	1804312132722180201:  "hemi-mainnet",
	15293031020466096408: "lisk-mainnet",
	13447077090413146373: "metal-mainnet",
	11690709103138290329: "mind-mainnet",
	17164792800244661392: "mint-mainnet",
	7222032299962346917:  "neox-mainnet",
	17912061998839310979: "plume-mainnet",
	2459028469735686113:  "katana-mainnet",
	11964252391146578476: "rootstock-mainnet",
	1673871237479749969:  "sonic-mainnet",
	470401360549526817:   "superseed-mainnet",
	16468599424800719238: "taiko-alethia-mainnet",
	3555797439612589184:  "zora-mainnet",
	14894068710063348487: "apechain-mainnet",
	6422105447186081193:  "astar-mainnet",
	1294465214383781161:  "berachain-mainnet",
	7937294810946806131:  "bitlayer-mainnet",
	5406759801798337480:  "bsquared-mainnet",
	1224752112135636129:  "core-mainnet",
	9043146809313071210:  "corn-mainnet",
	1462016016387883143:  "fraxtal-mainnet",
	7613811247471741961:  "hashkey-mainnet",
	3229138320728879060:  "hedera-mainnet",
	3461204551265785888:  "ink-mainnet",
	5608378062013572713:  "lens-mainnet",
	4627098889531055414:  "linea-mainnet",
	1556008542357238666:  "mantle-mainnet-v2",
	241851231317828981:   "merlin-mainnet",
	465944652040885897:   "opbnb-mainnet",
	4348158687435793198:  "polygon-zkevm-mainnet",
	6916147374840168594:  "ronin-mainnet",
	13204309965629103672: "scroll-mainnet-v2",
	3993510008929295315:  "shibarium-mainnet",
	12505351618335765396: "soneium-mainnet",
	1923510103922296319:  "unichain-mainnet",
	2049429975587534727:  "worldchain-mainnet",
	3016212468291539606:  "xlayer-mainnet",
	17198166215261833993: "zircuit-mainnet",
	18164309074156128038: "morph-mainnet",
	17529533435026248318: "sui-mainnet",
	5936861837188149645:  "tac-mainnet",
	17673274061779414707: "xdc-mainnet",
	9335212494177455608:  "plasma-mainnet",
	6473245816409426016:  "memento-mainnet",
	4426351306075016396:  "0g-mainnet",
	2135107236357186872:  "bittensor-mainnet",
	8481857512324358265:  "monad-mainnet",
	9723842205701363942:  "everclear-mainnet",
	4829375610284793157:  "ab-mainnet",
	12657445206920369324: "nexon-mainnet",
	1523760397290643893:  "jovay-mainnet",
	16978377838628290997: "stable-mainnet",
	6093540873831549674:  "megaeth-mainnet",
	7801139999541420232:  "pharos-mainnet",
	// --- Testnets ---
	16015286601757825753: "ethereum-sepolia",
	14767482510784806043: "avalanche-fuji",
	10344971235874465080: "base-sepolia",
	3478487238524512106:  "arbitrum-sepolia",
	16281711391670634445: "polygon-amoy",
	16423721717087811551: "solana-devnet",
	743186221051783445:   "aptos-testnet",
}

var ChainSelectorToChainID = map[uint64]string{
	// --- EVM Mainnets ---
	5009297550715157269:  "1",        // ethereum-mainnet
	4949039107694359620:  "42161",    // arbitrum-mainnet
	3734403246176062136:  "10",       // optimism-mainnet
	4051577828743386545:  "137",      // polygon-mainnet
	6433500567565415381:  "43114",    // avalanche-mainnet
	15971525489660198786: "8453",     // base-mainnet
	11344663589394136015: "56",       // bnb-mainnet
	465200170687744372:   "100",      // gnosis-mainnet
	1346049177634351622:  "1111",     // wemix-mainnet
	5142893604156789321:  "255",      // kroma-mainnet
	3719320017875267166:  "42220",    // celo-mainnet
	7264351850409363825:  "1088",     // metis-mainnet
	4411394078118774322:  "1284",     // moonbeam-mainnet
	3777822886988675105:  "324",      // zksync-mainnet
	1562403441176082196:  "534352",   // scroll-mainnet
	2442541497099098535:  "5000",     // mantle-mainnet
	8805746078405598895:  "81457",    // blast-mainnet
	3577778157919314504:  "2741",     // abstract-mainnet
	9027416829622342829:  "1329",     // sei-mainnet
	3849287863852499584:  "60808",    // bob-mainnet
	4560701533377838164:  "3637",     // botanix-mainnet
	1456215246176062136:  "25",       // cronos-mainnet
	8788096068760390840:  "388",      // cronos-zkevm-mainnet
	13624601974233774587: "42793",    // etherlink-mainnet
	1804312132722180201:  "43111",    // hemi-mainnet
	15293031020466096408: "1135",     // lisk-mainnet
	13447077090413146373: "1750",     // metal-mainnet
	11690709103138290329: "228",      // mind-mainnet
	17164792800244661392: "185",      // mint-mainnet
	7222032299962346917:  "47763",    // neox-mainnet
	17912061998839310979: "98866",    // plume-mainnet
	2459028469735686113:  "5765",     // katana-mainnet
	11964252391146578476: "30",       // rootstock-mainnet
	1673871237479749969:  "146",      // sonic-mainnet
	470401360549526817:   "5330",     // superseed-mainnet
	16468599424800719238: "167000",   // taiko-alethia-mainnet
	3555797439612589184:  "7777777",  // zora-mainnet
	14894068710063348487: "33139",    // apechain-mainnet
	6422105447186081193:  "592",      // astar-mainnet
	1294465214383781161:  "80094",    // berachain-mainnet
	7937294810946806131:  "200901",   // bitlayer-mainnet
	5406759801798337480:  "223",      // bsquared-mainnet
	1224752112135636129:  "1116",     // core-mainnet
	9043146809313071210:  "21000000", // corn-mainnet
	1462016016387883143:  "252",      // fraxtal-mainnet
	7613811247471741961:  "177",      // hashkey-mainnet
	3229138320728879060:  "295",      // hedera-mainnet
	3461204551265785888:  "57073",    // ink-mainnet
	5608378062013572713:  "232",      // lens-mainnet
	4627098889531055414:  "59144",    // linea-mainnet
	1556008542357238666:  "5000",     // mantle-mainnet-v2 (same chain ID as mantle)
	241851231317828981:   "4200",     // merlin-mainnet
	465944652040885897:   "204",      // opbnb-mainnet
	4348158687435793198:  "1101",     // polygon-zkevm-mainnet
	6916147374840168594:  "2020",     // ronin-mainnet
	13204309965629103672: "534352",   // scroll-mainnet-v2 (same chain ID as scroll)
	3993510008929295315:  "109",      // shibarium-mainnet
	12505351618335765396: "1868",     // soneium-mainnet
	1923510103922296319:  "130",      // unichain-mainnet
	2049429975587534727:  "480",      // worldchain-mainnet
	3016212468291539606:  "196",      // xlayer-mainnet
	17198166215261833993: "48900",    // zircuit-mainnet
	18164309074156128038: "2818",     // morph-mainnet
	5936861837188149645:  "2390",     // tac-mainnet
	17673274061779414707: "50",       // xdc-mainnet
	9335212494177455608:  "79",       // plasma-mainnet
	6473245816409426016:  "13371337", // memento-mainnet
	4426351306075016396:  "16600",    // 0g-mainnet
	2135107236357186872:  "964",      // bittensor-mainnet
	8481857512324358265:  "143",      // monad-mainnet
	9723842205701363942:  "25327",    // everclear-mainnet
	4829375610284793157:  "223718",   // ab-mainnet
	12657445206920369324: "7300",     // nexon-mainnet
	1523760397290643893:  "8998",     // jovay-mainnet
	16978377838628290997: "88999",    // stable-mainnet
	6093540873831549674:  "6342",     // megaeth-mainnet
	7801139999541420232:  "7001",     // pharos-mainnet
	// --- Non-EVM Mainnets ---
	124615329519749607:   "solana-mainnet", // solana-mainnet
	4741433654826277614:  "1",              // aptos-mainnet (Aptos chain ID)
	17529533435026248318: "sui-mainnet",    // sui-mainnet
	// --- EVM Testnets ---
	16015286601757825753: "11155111", // ethereum-sepolia
	14767482510784806043: "43113",    // avalanche-fuji
	10344971235874465080: "84532",    // base-sepolia
	3478487238524512106:  "421614",   // arbitrum-sepolia
	16281711391670634445: "80002",    // polygon-amoy
	// --- Non-EVM Testnets ---
	16423721717087811551: "solana-devnet", // solana-devnet
	743186221051783445:   "2",             // aptos-testnet
}

func SortedKeys(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// buildChainIDToSelector builds a reverse map from chain ID string to chain selector.
func buildChainIDToSelector() map[string]uint64 {
	m := make(map[string]uint64, len(ChainSelectorToChainID))
	for selector, chainID := range ChainSelectorToChainID {
		// If multiple selectors map to the same chain ID (e.g. mantle v1/v2),
		// keep the first one seen.
		if _, exists := m[chainID]; !exists {
			m[chainID] = selector
		}
	}
	return m
}

// transformNOPs filters and reshapes the raw JD node data into the goal_state.json nops format.
// Inclusion criteria (either satisfies inclusion):
//   - Node has label product=ccip
//   - Node has at least one chain config whose chainId maps to a selector in allChainSelectors
//
// ocrKeys are populated for ALL chain configs the node has (not just matching ones),
// so the user can see the full picture of how each NOP is configured.
func transformNOPs(nodeOps map[string]any, allChainSelectors []uint64, chainInfos map[uint64]*ChainInfo) map[string]any {
	nodesRaw, ok := nodeOps["nodes"]
	if !ok {
		return nil
	}
	nodes, ok := nodesRaw.([]any)
	if !ok {
		return nil
	}

	chainIDToSel := buildChainIDToSelector()

	// Build set of active chain selectors for quick lookup.
	activeSelectors := make(map[uint64]struct{}, len(allChainSelectors))
	for _, sel := range allChainSelectors {
		activeSelectors[sel] = struct{}{}
	}

	result := make(map[string]any)

	for _, nodeRaw := range nodes {
		node, ok := nodeRaw.(map[string]any)
		if !ok {
			continue
		}

		// Get node name (used as the NOP key).
		name, _ := node["name"].(string)
		if name == "" {
			// Try nopFriendlyName as fallback.
			name, _ = node["nopFriendlyName"].(string)
		}
		if name == "" {
			name, _ = node["id"].(string)
		}

		// Determine isBootstrap from labels or ocr2Config.
		isBootstrap := hasLabel(node, "type", "bootstrap")

		// Extract peerID from p2pKeyBundles.
		peerID := extractPeerID(node)

		// Build ocrKeys from chain configs -- only for non-bootstrap nodes.
		ocrKeys := make(map[string]any)

		if chainConfigs, ok := node["chainConfigs"].([]any); ok {
			for _, ccRaw := range chainConfigs {
				cc, ok := ccRaw.(map[string]any)
				if !ok {
					continue
				}

				chainID, _ := cc["chainId"].(string)
				if chainID == "" {
					continue
				}

				// Look up chain selector from chain ID.
				selector, found := chainIDToSel[chainID]

				// Check if this chain config indicates bootstrap.
				if ocr2, ok := cc["ocr2Config"].(map[string]any); ok {
					if isBoot, ok := ocr2["isBootstrap"].(bool); ok && isBoot {
						isBootstrap = true
					}
				}

				// Only include ocrKeys for chains that are in the active set.
				if !found {
					continue
				}
				if _, active := activeSelectors[selector]; !active {
					continue
				}

				// Determine network name: prefer canonical name, fall back to chainType-chainId.
				networkName := ChainSelectorToName[selector]
				if networkName == "" {
					chainType, _ := cc["chainType"].(string)
					if chainType == "" {
						chainType = "unknown"
					}
					networkName = fmt.Sprintf("%s-%s", strings.ToLower(chainType), chainID)
				}

				// For non-bootstrap nodes, build the ocrKey entry.
				if !isBootstrap {
					ocrEntry := buildOCRKeyEntry(cc, peerID)
					if ocrEntry != nil {
						ocrKeys[networkName] = ocrEntry
					}
				}
			}
		}

		// Inclusion filter: only include NOPs that have at least one chain config
		// whose chainId maps to a selector in allChainSelectors, or have label product=ccip.
		hasCCIPLabel := hasLabel(node, "product", "ccip") || hasLabel(node, "product", "ccv")
		hasMatchingChain := false
		if chainConfigs, ok := node["chainConfigs"].([]any); ok {
			for _, ccRaw := range chainConfigs {
				cc, ok := ccRaw.(map[string]any)
				if !ok {
					continue
				}
				chainID, _ := cc["chainId"].(string)
				if chainID == "" {
					continue
				}
				if sel, found := chainIDToSel[chainID]; found {
					if _, active := activeSelectors[sel]; active {
						hasMatchingChain = true
						break
					}
				}
			}
		}
		if !hasCCIPLabel && !hasMatchingChain {
			continue
		}

		// Build the NOP entry.
		nop := map[string]any{
			"nodeID":      node["id"],
			"peerID":      peerID,
			"isBootstrap": isBootstrap,
			"ocrKeys":     ocrKeys,
			"csaKey":      node["publicKey"],
			"isConnected": node["isConnected"],
			"isEnabled":   node["isEnabled"],
			"version":     node["version"],
			"labels":      node["labels"],
		}

		if wk, ok := node["workflowKey"].(string); ok && wk != "" {
			nop["workflowKey"] = wk
		}

		// Build approvedJobspecs from jobs data.
		approvedJobspecs := buildApprovedJobspecs(node)
		if len(approvedJobspecs) > 0 {
			nop["approvedJobspecs"] = approvedJobspecs
		}

		result[name] = nop
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

// hasLabel checks whether a node has a label with the given key and value.
func hasLabel(node map[string]any, key, value string) bool {
	labels, ok := node["labels"].([]any)
	if !ok {
		return false
	}
	for _, lRaw := range labels {
		l, ok := lRaw.(map[string]any)
		if !ok {
			continue
		}
		k, _ := l["key"].(string)
		v, _ := l["value"].(string)
		if k == key && v == value {
			return true
		}
	}
	return false
}

// extractPeerID extracts the first peerID from a node's p2pKeyBundles.
func extractPeerID(node map[string]any) string {
	bundles, ok := node["p2pKeyBundles"].([]any)
	if !ok || len(bundles) == 0 {
		return ""
	}
	first, ok := bundles[0].(map[string]any)
	if !ok {
		return ""
	}
	peerID, _ := first["peerId"].(string)
	return peerID
}

// buildOCRKeyEntry builds an ocrKeys entry from a chain config.
func buildOCRKeyEntry(chainConfig map[string]any, peerID string) map[string]any {
	ocr2ConfigRaw, ok := chainConfig["ocr2Config"]
	if !ok {
		return nil
	}
	ocr2Config, ok := ocr2ConfigRaw.(map[string]any)
	if !ok {
		return nil
	}

	ocrKeyBundleRaw, ok := ocr2Config["ocrKeyBundle"]
	if !ok {
		return nil
	}
	ocrKeyBundle, ok := ocrKeyBundleRaw.(map[string]any)
	if !ok {
		return nil
	}

	entry := map[string]any{
		"offchainPublicKey":         ocrKeyBundle["offchainPublicKey"],
		"onchainPublicKey":          ocrKeyBundle["onchainSigningAddress"],
		"peerID":                    peerID,
		"transmitAccount":           chainConfig["accountAddress"],
		"configEncryptionPublicKey": ocrKeyBundle["configPublicKey"],
		"keyBundleID":               ocrKeyBundle["bundleId"],
	}

	return entry
}

// buildApprovedJobspecs extracts approved job specs from the node's jobs data.
func buildApprovedJobspecs(node map[string]any) map[string]any {
	jobsRaw, ok := node["jobs"]
	if !ok {
		return nil
	}
	jobs, ok := jobsRaw.([]any)
	if !ok {
		return nil
	}

	result := make(map[string]any)
	for _, jobRaw := range jobs {
		job, ok := jobRaw.(map[string]any)
		if !ok {
			continue
		}

		jobID, _ := job["id"].(string)
		if jobID == "" {
			continue
		}

		proposalRaw, ok := job["approvedProposal"]
		if !ok {
			continue
		}
		proposal, ok := proposalRaw.(map[string]any)
		if !ok {
			continue
		}

		uuid, _ := job["uuid"].(string)

		entry := map[string]any{
			"proposal_id": proposal["id"],
			"uuid":        uuid,
			"spec":        proposal["spec"],
			"revision":    proposal["revision"],
		}

		result[jobID] = entry
	}

	if len(result) == 0 {
		return nil
	}
	return result
}

// nopInfo holds identifying information for a NOP, keyed by signer address.
type nopInfo struct {
	Name   string
	CSAKey string
}

// buildGlobalSignerToNOP builds a reverse map from lowercased signer address (no 0x) to NOP info.
func buildGlobalSignerToNOP(nops map[string]any) map[string]nopInfo {
	m := make(map[string]nopInfo)
	for nopName, nopRaw := range nops {
		nop, ok := nopRaw.(map[string]any)
		if !ok {
			continue
		}
		csaKey, _ := nop["csaKey"].(string)
		ocrKeysRaw, ok := nop["ocrKeys"]
		if !ok {
			continue
		}
		ocrKeys, ok := ocrKeysRaw.(map[string]any)
		if !ok {
			continue
		}
		for _, keyEntryRaw := range ocrKeys {
			keyEntry, ok := keyEntryRaw.(map[string]any)
			if !ok {
				continue
			}
			onchainPK, _ := keyEntry["onchainPublicKey"].(string)
			if onchainPK == "" {
				continue
			}
			onchainPK = strings.ToLower(strings.TrimPrefix(onchainPK, "0x"))
			m[onchainPK] = nopInfo{Name: nopName, CSAKey: csaKey}
		}
	}
	return m
}

// enrichCommitteeVerifierSigners walks all chains in the legacy output and replaces
// each committeeVerifier signatureConfigs "signers" array of plain address strings
// with an array of objects: { "address": ..., "nopName": ..., "csaKey": ... }.
// If a signer address is not found in the NOP map, nopName and csaKey are omitted.
func enrichCommitteeVerifierSigners(legacy map[string]any, signerToNOP map[string]nopInfo) {
	chainsRaw, ok := legacy["chains"]
	if !ok {
		return
	}
	chains, ok := chainsRaw.(map[string]any)
	if !ok {
		return
	}

	for _, chainDataRaw := range chains {
		chainData, ok := chainDataRaw.(map[string]any)
		if !ok {
			continue
		}

		cvRaw, ok := chainData["committeeVerifier"]
		if !ok {
			continue
		}
		cvContracts, ok := cvRaw.(map[string]any)
		if !ok {
			continue
		}

		for _, contractRaw := range cvContracts {
			contract, ok := contractRaw.(map[string]any)
			if !ok {
				continue
			}

			sigConfigsRaw, ok := contract["signatureConfigs"]
			if !ok {
				continue
			}
			sigConfigs, ok := sigConfigsRaw.([]any)
			if !ok {
				continue
			}

			for _, scRaw := range sigConfigs {
				sc, ok := scRaw.(map[string]any)
				if !ok {
					continue
				}

				signersRaw, ok := sc["signers"]
				if !ok {
					continue
				}
				signers, ok := signersRaw.([]any)
				if !ok {
					continue
				}

				enriched := make([]any, 0, len(signers))
				for _, signerRaw := range signers {
					addr, ok := signerRaw.(string)
					if !ok {
						enriched = append(enriched, signerRaw)
						continue
					}

					normalized := strings.ToLower(strings.TrimPrefix(addr, "0x"))
					entry := map[string]any{
						"address": addr,
					}
					if info, found := signerToNOP[normalized]; found {
						entry["nopName"] = info.Name
						entry["csaKey"] = info.CSAKey
					}
					enriched = append(enriched, entry)
				}

				sc["signers"] = enriched
			}
		}
	}
}

// buildNopsDiff compares JD NOP-to-network assignments against on-chain committeeVerifier
// signer data and returns a per-network diff.
//
// For each network in legacy["chains"]:
//   - jdNotOnChain: NOPs whose ocrKeys say they're on this network, but whose
//     onchainPublicKey is not a signer in any committeeVerifier on that chain.
//   - onChainNotInJD: Signer addresses in committeeVerifier that have no matching
//     NOP claiming to be on this network in JD.
//
// Networks with no diff (both lists empty) are omitted.
func buildNopsDiff(legacy map[string]any, nops map[string]any) map[string]any {
	chainsRaw, ok := legacy["chains"]
	if !ok {
		return nil
	}
	chains, ok := chainsRaw.(map[string]any)
	if !ok {
		return nil
	}

	result := make(map[string]any)

	for networkName, chainDataRaw := range chains {
		chainData, ok := chainDataRaw.(map[string]any)
		if !ok {
			continue
		}

		// Collect on-chain signer addresses from all committeeVerifier contracts on this chain.
		// Key: lowercased address without 0x prefix. Value: enriched signer entry.
		onChainSigners := make(map[string]map[string]any)
		if cvRaw, ok := chainData["committeeVerifier"]; ok {
			if cvContracts, ok := cvRaw.(map[string]any); ok {
				for _, contractRaw := range cvContracts {
					contract, ok := contractRaw.(map[string]any)
					if !ok {
						continue
					}
					sigConfigsRaw, ok := contract["signatureConfigs"]
					if !ok {
						continue
					}
					sigConfigs, ok := sigConfigsRaw.([]any)
					if !ok {
						continue
					}
					for _, scRaw := range sigConfigs {
						sc, ok := scRaw.(map[string]any)
						if !ok {
							continue
						}
						signersRaw, ok := sc["signers"]
						if !ok {
							continue
						}
						signers, ok := signersRaw.([]any)
						if !ok {
							continue
						}
						for _, signerRaw := range signers {
							signer, ok := signerRaw.(map[string]any)
							if !ok {
								continue
							}
							addr, _ := signer["address"].(string)
							if addr == "" {
								continue
							}
							normalized := strings.ToLower(strings.TrimPrefix(addr, "0x"))
							onChainSigners[normalized] = signer
						}
					}
				}
			}
		}

		// Collect JD-side NOPs that claim to be on this network.
		// Key: lowercased onchainPublicKey without 0x prefix. Value: nop name.
		jdNOPsOnNetwork := make(map[string]string)
		for nopName, nopRaw := range nops {
			nop, ok := nopRaw.(map[string]any)
			if !ok {
				continue
			}
			ocrKeysRaw, ok := nop["ocrKeys"]
			if !ok {
				continue
			}
			ocrKeys, ok := ocrKeysRaw.(map[string]any)
			if !ok {
				continue
			}
			keyEntry, ok := ocrKeys[networkName]
			if !ok {
				continue
			}
			keyMap, ok := keyEntry.(map[string]any)
			if !ok {
				continue
			}
			onchainPK, _ := keyMap["onchainPublicKey"].(string)
			if onchainPK == "" {
				continue
			}
			normalized := strings.ToLower(strings.TrimPrefix(onchainPK, "0x"))
			jdNOPsOnNetwork[normalized] = nopName
		}

		// Skip networks with no committeeVerifier data and no JD NOPs.
		if len(onChainSigners) == 0 && len(jdNOPsOnNetwork) == 0 {
			continue
		}

		// Diff: JD says NOP is on this network but not found on-chain.
		var jdNotOnChain []map[string]any
		for signerAddr, nopName := range jdNOPsOnNetwork {
			if _, onChain := onChainSigners[signerAddr]; !onChain {
				jdNotOnChain = append(jdNotOnChain, map[string]any{
					"nopName":       nopName,
					"signerAddress": "0x" + signerAddr,
				})
			}
		}

		// Diff: On-chain signer with no matching JD NOP on this network.
		var onChainNotInJD []map[string]any
		for signerAddr, signerEntry := range onChainSigners {
			if _, inJD := jdNOPsOnNetwork[signerAddr]; !inJD {
				entry := map[string]any{
					"address": signerEntry["address"],
				}
				if nopName, ok := signerEntry["nopName"].(string); ok && nopName != "" {
					entry["nopName"] = nopName
				}
				onChainNotInJD = append(onChainNotInJD, entry)
			}
		}

		// Only include networks with actual diffs.
		if len(jdNotOnChain) == 0 && len(onChainNotInJD) == 0 {
			continue
		}

		networkDiff := make(map[string]any)
		if len(jdNotOnChain) > 0 {
			networkDiff["jdNotOnChain"] = jdNotOnChain
		}
		if len(onChainNotInJD) > 0 {
			networkDiff["onChainNotInJD"] = onChainNotInJD
		}
		result[networkName] = networkDiff
	}

	if len(result) == 0 {
		return nil
	}
	return result
}
