package changesets

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

// LegacyEVMTokenPool is an EVM token pool found by FindLegacyEVMTokenPool.
type LegacyEVMTokenPool struct {
	PoolAddress  []byte
	TokenAddress []byte
	// PoolRef is the pool's datastore ref, for the decimals read.
	PoolRef datastore.AddressRef
}

// decimalsAdapterVersion is the EVM token adapter used to read pool decimals. The legacy code read
// every pool version through one token pool binding; getTokenDecimals has the same ABI in all of
// them, and the 1.5.1 adapter reads it through the 1.5.1 binding.
var decimalsAdapterVersion = semver.MustParse("1.5.1")

// legacyEVMPoolTiers returns the groups of datastore ref versions that chainlink's legacy EVM
// state loaded into the map searched for a pool type and requested version, in the order
// v1_5_1.GetTokenPoolAddressFromSymbolTypeAndVersion searched them. bySymbol is false for the USDC
// pools and proxies, which the legacy state keyed by version only.
func legacyEVMPoolTiers(poolType cldf.ContractType, version semver.Version) (tiers [][]string, bySymbol bool) {
	v161 := version.Equal(semver.MustParse("1.6.1"))
	switch poolType {
	case shared.BurnMintTokenPool:
		if v161 {
			return [][]string{{"1.6.1"}, {"1.5.1"}, {"1.5.0"}}, true
		}
		return [][]string{{"1.5.1"}, {"1.5.0"}}, true
	case shared.LockReleaseTokenPool:
		if v161 {
			return [][]string{{"1.6.1"}, {"1.5.1"}}, true
		}
		return [][]string{{"1.5.1"}}, true
	case "BurnMintFastTransferTokenPool":
		return [][]string{{"1.6.1", "1.6.3-dev"}}, true
	case "BurnMintWithExternalMinterFastTransferTokenPool", "HybridWithExternalMinterFastTransferTokenPool",
		"BurnMintWithExternalMinterTokenPool", "HybridWithExternalMinterTokenPool":
		return [][]string{{"1.6.0"}}, true
	case "BurnFromMintTokenPool", "BurnWithFromMintTokenPool":
		return [][]string{{"1.5.1"}}, true
	case shared.USDCTokenPool, "HybridLockReleaseUSDCTokenPool":
		return [][]string{{version.String()}}, false
	case "USDCTokenPoolProxy":
		return [][]string{{version.String()}}, false
	}
	return nil, false
}

// legacyEVMPoolRefTypes returns the datastore types feeding the legacy map for a pool type: the
// USDC and hybrid lock-release USDC pools shared one map.
func legacyEVMPoolRefTypes(poolType cldf.ContractType) []datastore.ContractType {
	switch poolType {
	case shared.USDCTokenPool, "HybridLockReleaseUSDCTokenPool":
		return []datastore.ContractType{datastore.ContractType(shared.USDCTokenPool), "HybridLockReleaseUSDCTokenPool"}
	}
	return []datastore.ContractType{datastore.ContractType(poolType)}
}

// FindLegacyEVMTokenPool finds an EVM token pool the way the legacy Solana changesets did through
// chainlink's EVM state and v1_5_1.GetTokenPoolAddressFromSymbolTypeAndVersion: among the
// datastore refs of the pool's type, the one whose token has the given on-chain ERC20 symbol and
// whose on-chain typeAndVersion has the given version. found is false, with a nil error, when no
// pool matches.
//
// The Solana module can't import EVM code, so the on-chain reads go through the EVM family's
// registered tokens.TokenRefResolver: ResolveTokenPoolRef reads the pool's typeAndVersion and token,
// and ResolveTokenRef reads the token's symbol. Importing the chainlink-ccip EVM adapters registers
// it.
//
// Differences from the legacy lookup, none of which affects the current datastores:
//   - Two live pools matching the same symbol and version are an error. The legacy maps kept
//     whichever was loaded last.
//   - A pool whose typeAndVersion isn't valid semver (such as 1.6.x-dev) is an error. The legacy
//     loader rewrote x to 3. No such pool is in the CLD datastores.
func FindLegacyEVMTokenPool(
	e cldf.Environment,
	chainSelector uint64,
	symbol shared.TokenSymbol,
	poolType cldf.ContractType,
	poolVersion semver.Version,
) (LegacyEVMTokenPool, bool, error) {
	if e.DataStore == nil {
		return LegacyEVMTokenPool{}, false, fmt.Errorf("datastore not available for chain %d", chainSelector)
	}
	resolver, ok := tokensapi.GetTokenAdapterRegistry().GetTokenRefResolver(chain_selectors.FamilyEVM)
	if !ok {
		return LegacyEVMTokenPool{}, false, errors.New("no token ref resolver registered for the EVM family; import the chainlink-ccip EVM adapters")
	}
	tiers, bySymbol := legacyEVMPoolTiers(poolType, poolVersion)
	refTypes := legacyEVMPoolRefTypes(poolType)
	for _, tier := range tiers {
		var matches []LegacyEVMTokenPool
		for _, ref := range e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector)) {
			if !isLegacyEVMPoolRef(ref, refTypes, tier) {
				continue
			}
			onchain, err := resolver.ResolveTokenPoolRef(e.OperationsBundle, e.BlockChains, e.DataStore, chainSelector, ref.Address)
			if err != nil {
				return LegacyEVMTokenPool{}, false, err
			}
			// ResolveTokenPoolRef puts the pool's token address in the qualifier, or a placeholder
			// when the pool's getToken call fails. The legacy loader failed in that case.
			tokenAddress, err := evmAddressBytes(onchain.Qualifier)
			if err != nil {
				return LegacyEVMTokenPool{}, false, fmt.Errorf("failed to get token address from pool with address %s: %w", ref.Address, err)
			}
			if bySymbol {
				if onchain.Version == nil || !onchain.Version.Equal(&poolVersion) {
					continue
				}
				tokenRef, symbolErr := resolver.ResolveTokenRef(e.OperationsBundle, e.BlockChains, e.DataStore, chainSelector, onchain.Qualifier)
				if symbolErr != nil {
					return LegacyEVMTokenPool{}, false, symbolErr
				}
				if tokenRef.Qualifier != string(symbol) {
					continue
				}
			}
			poolAddress, err := evmAddressBytes(ref.Address)
			if err != nil {
				return LegacyEVMTokenPool{}, false, err
			}
			matches = append(matches, LegacyEVMTokenPool{PoolAddress: poolAddress, TokenAddress: tokenAddress, PoolRef: ref})
		}
		switch len(matches) {
		case 0:
			continue
		case 1:
			return matches[0], true, nil
		default:
			return LegacyEVMTokenPool{}, false, fmt.Errorf("%d %s pools with symbol %q and version %s on chain %d; the pool is ambiguous", len(matches), poolType, symbol, poolVersion.String(), chainSelector)
		}
	}
	return LegacyEVMTokenPool{}, false, nil
}

func isLegacyEVMPoolRef(ref datastore.AddressRef, types []datastore.ContractType, versions []string) bool {
	if ref.Version == nil || ref.Labels.Contains(shared.SupersededLabel) {
		return false
	}
	typeOK := false
	for _, t := range types {
		typeOK = typeOK || ref.Type == t
	}
	if !typeOK {
		return false
	}
	for _, v := range versions {
		if ref.Version.Equal(semver.MustParse(v)) {
			return true
		}
	}
	return false
}

// GetLegacyEVMTokenPoolDecimals reads the decimals of the pool's token from the pool, as the legacy
// Solana changesets did. When the pool can't report them, it fails unless fallbackToToken is set,
// in which case it reads them from the token. Callers pass what their legacy changeset did: the
// v0_1_1 changesets fell back, the v0_1_0 ones failed.
//
// The read goes through the EVM family's registered 1.5.1 token adapter. Its DeriveTokenDecimals
// tries the token first and then the pool, so passing no token makes it read the pool only.
func GetLegacyEVMTokenPoolDecimals(e cldf.Environment, chainSelector uint64, pool LegacyEVMTokenPool, fallbackToToken bool) (uint8, error) {
	adapter, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdapter(chain_selectors.FamilyEVM, decimalsAdapterVersion)
	if !ok {
		return 0, errors.New("no 1.5.1 token adapter registered for the EVM family; import the chainlink-ccip EVM adapters")
	}
	decimals, err := adapter.DeriveTokenDecimals(e, chainSelector, pool.PoolRef, nil)
	if err == nil {
		return decimals, nil
	}
	if !fallbackToToken {
		return 0, fmt.Errorf("failed to get token decimals: %w", err)
	}
	decimals, err = adapter.DeriveTokenDecimals(e, chainSelector, pool.PoolRef, pool.TokenAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to get token decimals from token contract: %w", err)
	}
	return decimals, nil
}

// evmAddressBytes parses a 0x-prefixed 20-byte hex address. The Solana module can't import
// go-ethereum.
func evmAddressBytes(address string) ([]byte, error) {
	b, err := hex.DecodeString(strings.TrimPrefix(strings.TrimPrefix(address, "0x"), "0X"))
	if err != nil || len(b) != 20 {
		return nil, fmt.Errorf("%q is not an EVM address", address)
	}
	return b, nil
}

// leftPad32 reproduces go-ethereum's common.LeftPadBytes(b, 32).
func leftPad32(b []byte) []byte {
	if len(b) >= 32 {
		return b
	}
	padded := make([]byte, 32)
	copy(padded[32-len(b):], b)
	return padded
}
