package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// =====================================================
// Address References (from address_refs.json)
// =====================================================

// AddressRef represents a contract deployment from address_refs.json
type AddressRef struct {
	Address       string   `json:"address"`
	ChainSelector uint64   `json:"chainSelector"`
	Labels        []string `json:"labels"`
	Qualifier     string   `json:"qualifier"`
	Type          string   `json:"type"`
	Version       string   `json:"version"`
}

// LoadAddressRefs loads contract addresses from a JSON file
func LoadAddressRefs(path string) ([]AddressRef, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read address refs file: %w", err)
	}

	var refs []AddressRef
	if err := json.Unmarshal(data, &refs); err != nil {
		return nil, fmt.Errorf("failed to parse address refs: %w", err)
	}

	return refs, nil
}

// =====================================================
// Network Configuration (from testnet.yaml)
// =====================================================

// NetworkConfig represents the root of the YAML configuration
type NetworkConfig struct {
	PreferredURLScheme string    `yaml:"preferred_url_scheme"`
	Networks           []Network `yaml:"networks"`
}

// Network represents a single network/chain configuration
type Network struct {
	Type           string         `yaml:"type"`
	ChainSelector  uint64         `yaml:"chain_selector"`
	RPCs           []RPCEndpoint  `yaml:"rpcs"`
	BlockExplorer  *BlockExplorer `yaml:"block_explorer,omitempty"`
	Metadata       *Metadata      `yaml:"metadata,omitempty"`
}

// RPCEndpoint represents an RPC endpoint configuration
type RPCEndpoint struct {
	RPCName            string `yaml:"rpc_name"`
	PreferredURLScheme string `yaml:"preferred_url_scheme"`
	HTTPURL            string `yaml:"http_url"`
	WSURL              string `yaml:"ws_url,omitempty"`
}

// BlockExplorer represents block explorer configuration
type BlockExplorer struct {
	Type   string `yaml:"type"`
	APIKey string `yaml:"api_key"`
	URL    string `yaml:"url"`
}

// Metadata contains optional chain metadata
type Metadata struct {
	AnvilConfig *AnvilConfig `yaml:"anvil_config,omitempty"`
}

// AnvilConfig represents Anvil fork configuration
type AnvilConfig struct {
	Image          string `yaml:"image"`
	Port           int    `yaml:"port"`
	ArchiveHTTPURL string `yaml:"archive_http_url"`
}

// LoadNetworkConfig loads network configuration from a YAML file
func LoadNetworkConfig(path string) (*NetworkConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read network config file: %w", err)
	}

	var config NetworkConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse network config: %w", err)
	}

	return &config, nil
}

// =====================================================
// Chain Family Detection
// =====================================================

// ChainFamily represents the VM/chain type
type ChainFamily string

const (
	ChainFamilyEVM    ChainFamily = "evm"
	ChainFamilySVM    ChainFamily = "svm"    // Solana
	ChainFamilyAptos  ChainFamily = "aptos"
	ChainFamilyTON    ChainFamily = "ton"
	ChainFamilySui    ChainFamily = "sui"
	ChainFamilyUnknown ChainFamily = "unknown"
)

// ChainInfo holds information about a chain
type ChainInfo struct {
	Selector uint64
	Family   ChainFamily
	Name     string
	RPCs     []RPCEndpoint
}

// ChainRegistry maps chain selectors to their info
type ChainRegistry struct {
	chains map[uint64]*ChainInfo
}

// NewChainRegistry creates a new chain registry from network config
func NewChainRegistry(config *NetworkConfig) *ChainRegistry {
	registry := &ChainRegistry{
		chains: make(map[uint64]*ChainInfo),
	}

	for _, network := range config.Networks {
		family := detectChainFamily(network)
		name := extractChainName(network)

		registry.chains[network.ChainSelector] = &ChainInfo{
			Selector: network.ChainSelector,
			Family:   family,
			Name:     name,
			RPCs:     network.RPCs,
		}
	}

	return registry
}

// GetChain returns chain info for a selector, or nil if not found
func (r *ChainRegistry) GetChain(selector uint64) *ChainInfo {
	return r.chains[selector]
}

// GetRPCURLs returns HTTP RPC URLs for a chain
func (r *ChainRegistry) GetRPCURLs(selector uint64) []string {
	chain := r.chains[selector]
	if chain == nil {
		return nil
	}

	urls := make([]string, 0, len(chain.RPCs))
	for _, rpc := range chain.RPCs {
		if rpc.HTTPURL != "" {
			urls = append(urls, rpc.HTTPURL)
		}
	}
	return urls
}

// detectChainFamily determines the chain family based on RPC URLs and patterns
// The YAML file uses comments like "# ethereum-testnet-sepolia (evm:11155111)"
// but we need to detect from available data
func detectChainFamily(network Network) ChainFamily {
	// Check if any RPC URL contains chain-specific patterns
	for _, rpc := range network.RPCs {
		url := strings.ToLower(rpc.HTTPURL)

		// Solana detection
		if strings.Contains(url, "solana") || strings.Contains(url, "sol-") {
			return ChainFamilySVM
		}

		// Aptos detection
		if strings.Contains(url, "aptos") {
			return ChainFamilyAptos
		}

		// TON detection
		if strings.Contains(url, "liteserver://") {
			return ChainFamilyTON
		}

		// Sui detection
		if strings.Contains(url, "sui") {
			return ChainFamilySui
		}
	}

	// Check block explorer type
	if network.BlockExplorer != nil {
		switch strings.ToLower(network.BlockExplorer.Type) {
		case "solana-explorer":
			return ChainFamilySVM
		case "aptos-explorer":
			return ChainFamilyAptos
		}
	}

	// Default to EVM (most chains in the config are EVM)
	return ChainFamilyEVM
}

// extractChainName attempts to extract a readable name for the chain
func extractChainName(network Network) string {
	// Use the first RPC name as a hint, or generate from selector
	if len(network.RPCs) > 0 && network.RPCs[0].RPCName != "" {
		return network.RPCs[0].RPCName
	}
	return fmt.Sprintf("chain-%d", network.ChainSelector)
}

// =====================================================
// Known Chain Selectors (for reference)
// =====================================================

// WellKnownChains maps chain selectors to human-readable names
var WellKnownChains = map[uint64]string{
	16015286601757825753: "ethereum-sepolia",
	14767482510784806043: "avalanche-fuji",
	10344971235874465080: "base-sepolia",
	3478487238524512106:  "arbitrum-sepolia",
	16281711391670634445: "polygon-amoy",
	16423721717087811551: "solana-devnet",
	743186221051783445:   "aptos-testnet",
	1399300952838017768:  "ton-testnet",
	9762610643973837292:  "sui-testnet",
}
