package shared

import (
	"encoding/hex"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mr-tron/base58"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
)

// GetAddressFromBytes renders a raw on-chain address in the textual form of the family that
// owns chainSelector.
//
// Reproduced from chainlink/deployment/ccip/view/shared/common.go. The rest of that package is
// not copied: it is built on the EVM RegistryModuleOwnerCustom and TokenAdminRegistry
// gethwrappers, which the Solana views never touch.
func GetAddressFromBytes(chainSelector uint64, address []byte) string {
	family, err := chain_selectors.GetSelectorFamily(chainSelector)
	if err != nil {
		return "invalid chain selector"
	}

	switch family {
	case chain_selectors.FamilyEVM:
		return strings.ToLower(common.BytesToAddress(address).Hex())
	case chain_selectors.FamilySolana:
		return base58.Encode(address)
	case chain_selectors.FamilyAptos:
		return "0x" + hex.EncodeToString(common.LeftPadBytes(address, 32))
	default:
		return "unsupported chain family"
	}
}
