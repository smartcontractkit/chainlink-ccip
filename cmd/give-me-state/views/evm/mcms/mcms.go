package mcms

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	gethCommon "github.com/ethereum/go-ethereum/common"
)

const mcmsABIJson = `[{"inputs":[],"name":"getConfig","outputs":[{"components":[{"components":[{"internalType":"address","name":"addr","type":"address"},{"internalType":"uint8","name":"index","type":"uint8"},{"internalType":"uint8","name":"group","type":"uint8"}],"internalType":"struct ManyChainMultiSig.Signer[]","name":"signers","type":"tuple[]"},{"internalType":"uint8[32]","name":"groupQuorums","type":"uint8[32]"},{"internalType":"uint8[32]","name":"groupParents","type":"uint8[32]"}],"internalType":"struct ManyChainMultiSig.Config","name":"","type":"tuple"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`

// Function selectors for MCMS contracts (ManyChainMultiSig)
var (
	mcmsABI           abi.ABI
	selectorOwner     []byte
	selectorGetConfig []byte
)

func init() {
	var err error
	mcmsABI, err = abi.JSON(strings.NewReader(mcmsABIJson))
	if err != nil {
		panic("Failed to parse mcmsABI: " + err.Error())
	}

	// Extract selectors dynamically
	selectorOwner = mcmsABI.Methods["owner"].ID
	selectorGetConfig = mcmsABI.Methods["getConfig"].ID
}

// Signer represents a signer in the MCMS config
type Signer struct {
	Address gethCommon.Address `json:"addr"`
	Index   uint8              `json:"index"`
	Group   uint8              `json:"group"`
}

// MCMSConfig represents the MCMS contract configuration
type MCMSConfig struct {
	Signers      []Signer  `json:"signers"`
	GroupQuorums [32]uint8 `json:"groupQuorums"`
	GroupParents [32]uint8 `json:"groupParents"`
}

// ViewMCMS generates a view of an MCMS contract (bypasser, canceller, proposer).
func ViewMCMS(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	// Get owner
	if owner, err := getMCMSOwner(ctx); err == nil {
		result["owner"] = owner
	} else {
		result["owner_error"] = err.Error()
	}

	// Get config
	if config, err := getMCMSConfig(ctx); err == nil {
		result["config"] = config
	} else {
		result["config_error"] = err.Error()
	}

	return result, nil
}

// getMCMSOwner fetches the owner address.
func getMCMSOwner(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorOwner)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getMCMSConfig fetches and decodes the MCMS configuration.
func getMCMSConfig(ctx *views.ViewContext) (*MCMSConfig, error) {
	data, err := common.ExecuteCall(ctx, selectorGetConfig)
	if err != nil {
		return nil, err
	}

	return decodeMCMSConfig(data)
}

// ViewCallProxy generates a view of a CallProxy contract.
// CallProxy has no public read functions, so we just return the address.
func ViewCallProxy(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	// CallProxy has no owner or other readable state
	result["owner"] = "0x0000000000000000000000000000000000000000"

	return result, nil
}

// decodeMCMSConfig decodes the getConfig() return data.
// Returns: (Signer[] signers, uint8[32] groupQuorums, uint8[32] groupParents)
// Signer struct: (address addr, uint8 index, uint8 group)
func decodeMCMSConfig(data []byte) (*MCMSConfig, error) {
	result := &MCMSConfig{
		Signers:      []Signer{},
		GroupQuorums: [32]uint8{},
		GroupParents: [32]uint8{},
	}

	// Unpack raw data into MCMSConfig
	err := mcmsABI.UnpackIntoInterface(&result, "getConfig", data)

	if err != nil {
		panic(err)
	}

	return result, nil
}
