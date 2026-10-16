package adapters

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/lombard_verifier"
	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

const glamsterdamTokensTestChainSel = uint64(111)

var (
	lombardPoolAddr = common.HexToAddress("0x0000000000000000000000000000000000000A")
	usdcPoolAddr    = common.HexToAddress("0x0000000000000000000000000000000000000B")
	cctpPoolAddr    = common.HexToAddress("0x0000000000000000000000000000000000000C")
	unrelatedAddr   = common.HexToAddress("0x0000000000000000000000000000000000000D")
)

// glamsterdamTokensTestDataStore seeds a datastore with one of each Lombard/USDC-family token pool
// (each on its own chain-tagged ref), plus one unrelated contract type that must never be picked
// up as a candidate token. These methods only ever read the datastore — no real chain calls are
// involved — so they can be tested directly against the real EVM adapter without deploying any
// contracts.
func glamsterdamTokensTestDataStore(t *testing.T) datastore.DataStore {
	t.Helper()
	ds := datastore.NewMemoryDataStore()

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: glamsterdamTokensTestChainSel,
		Address:       lombardPoolAddr.Hex(),
		Type:          datastore.ContractType(lombard_token_pool.ContractType),
		Version:       lombard_token_pool.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: glamsterdamTokensTestChainSel,
		Address:       usdcPoolAddr.Hex(),
		Type:          datastore.ContractType(siloed_usdc_token_pool.ContractType),
		Version:       siloed_usdc_token_pool.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: glamsterdamTokensTestChainSel,
		Address:       cctpPoolAddr.Hex(),
		Type:          datastore.ContractType(cctp_through_ccv_token_pool.ContractType),
		Version:       cctp_through_ccv_token_pool.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: glamsterdamTokensTestChainSel,
		Address:       unrelatedAddr.Hex(),
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       fee_quoter.Version,
	}))

	return ds.Seal()
}

// TestGlamsterdamGasAdapter_DiscoverCandidateTokens is a regression test for the bug where
// DiscoverCandidateTokens filtered on a generic "TokenPool" ContractType that nothing in the
// datastore is ever tagged with, so it silently found zero pools. It must return exactly the
// Lombard/USDC-family pool addresses (as pool addresses, not token addresses) and ignore anything
// else in the datastore.
func TestGlamsterdamGasAdapter_DiscoverCandidateTokens(t *testing.T) {
	adapter := &GlamsterdamGasAdapter{}
	ds := glamsterdamTokensTestDataStore(t)

	tokens, err := adapter.DiscoverCandidateTokens(operations.Bundle{}, chain.BlockChains{}, ds, glamsterdamTokensTestChainSel)
	require.NoError(t, err)

	var got []common.Address
	for _, tok := range tokens {
		got = append(got, common.BytesToAddress(tok))
	}
	require.ElementsMatch(t, []common.Address{lombardPoolAddr, usdcPoolAddr, cctpPoolAddr}, got)
}

// TestGlamsterdamGasAdapter_DiscoverCandidateTokens_DedupesAcrossQualifiers verifies that the same
// pool address tracked under two different datastore qualifiers (a real scenario for Lombard,
// which supports multiple qualifier-scoped pools per chain) is only returned once.
func TestGlamsterdamGasAdapter_DiscoverCandidateTokens_DedupesAcrossQualifiers(t *testing.T) {
	adapter := &GlamsterdamGasAdapter{}
	ds := datastore.NewMemoryDataStore()

	for _, qualifier := range []string{"lane-a", "lane-b"} {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: glamsterdamTokensTestChainSel,
			Address:       lombardPoolAddr.Hex(),
			Type:          datastore.ContractType(lombard_token_pool.ContractType),
			Version:       lombard_token_pool.Version,
			Qualifier:     qualifier,
		}))
	}

	tokens, err := adapter.DiscoverCandidateTokens(operations.Bundle{}, chain.BlockChains{}, ds.Seal(), glamsterdamTokensTestChainSel)
	require.NoError(t, err)
	require.Len(t, tokens, 1)
	require.Equal(t, lombardPoolAddr, common.BytesToAddress(tokens[0]))
}

// TestGlamsterdamGasAdapter_TokenFieldSpec is the key regression test for the "always defaults to
// USDC" bug: TokenFieldSpec must distinguish a Lombard pool from either USDC pool implementation
// by its actual datastore ContractType, not guess.
func TestGlamsterdamGasAdapter_TokenFieldSpec(t *testing.T) {
	adapter := &GlamsterdamGasAdapter{}
	ds := glamsterdamTokensTestDataStore(t)

	lombardSpec, err := adapter.TokenFieldSpec(operations.Bundle{}, chain.BlockChains{}, ds, glamsterdamTokensTestChainSel, lombardPoolAddr.Bytes())
	require.NoError(t, err)
	require.Equal(t, v2_0_0_adapters.LombardTokenPoolDestGasOverhead.Name, lombardSpec.Name)

	usdcSpec, err := adapter.TokenFieldSpec(operations.Bundle{}, chain.BlockChains{}, ds, glamsterdamTokensTestChainSel, usdcPoolAddr.Bytes())
	require.NoError(t, err)
	require.Equal(t, v2_0_0_adapters.USDCTokenPoolDestGasOverhead.Name, usdcSpec.Name)

	cctpSpec, err := adapter.TokenFieldSpec(operations.Bundle{}, chain.BlockChains{}, ds, glamsterdamTokensTestChainSel, cctpPoolAddr.Bytes())
	require.NoError(t, err)
	require.Equal(t, v2_0_0_adapters.USDCTokenPoolDestGasOverhead.Name, cctpSpec.Name, "CCTPThroughCCVTokenPool is a USDC-family pool")

	_, err = adapter.TokenFieldSpec(operations.Bundle{}, chain.BlockChains{}, ds, glamsterdamTokensTestChainSel, unrelatedAddr.Bytes())
	require.Error(t, err)
}

// TestResolveAddressRefsAllVersions is a regression test for the LombardVerifier/CCTPVerifier
// single-address bug: a chain can have more than one verifier version tracked simultaneously
// during a migration window (e.g. v2.0.0 and v2.1.0 live at once), and every one of them must be
// discovered, not just whichever version the lookup happens to check first.
func TestResolveAddressRefsAllVersions(t *testing.T) {
	v210 := semver.MustParse("2.1.0")
	v200 := semver.MustParse("2.0.0")
	versions := []*semver.Version{v210, v200}

	addrA := common.HexToAddress("0x00000000000000000000000000000000000A01")
	addrB := common.HexToAddress("0x00000000000000000000000000000000000A02")

	t.Run("returns every version's address when both are tracked", func(t *testing.T) {
		addrs := []datastore.AddressRef{
			{ChainSelector: glamsterdamTokensTestChainSel, Address: addrA.Hex(), Type: "LombardVerifier", Version: v210},
			{ChainSelector: glamsterdamTokensTestChainSel, Address: addrB.Hex(), Type: "LombardVerifier", Version: v200},
		}
		got := resolveAddressRefsAllVersions(addrs, glamsterdamTokensTestChainSel, "LombardVerifier", versions)
		require.ElementsMatch(t, []common.Address{addrA, addrB}, got)
	})

	t.Run("dedupes when the same address is tagged under two versions", func(t *testing.T) {
		addrs := []datastore.AddressRef{
			{ChainSelector: glamsterdamTokensTestChainSel, Address: addrA.Hex(), Type: "LombardVerifier", Version: v210},
			{ChainSelector: glamsterdamTokensTestChainSel, Address: addrA.Hex(), Type: "LombardVerifier", Version: v200},
		}
		got := resolveAddressRefsAllVersions(addrs, glamsterdamTokensTestChainSel, "LombardVerifier", versions)
		require.Equal(t, []common.Address{addrA}, got)
	})

	t.Run("returns nothing when no version is tracked", func(t *testing.T) {
		got := resolveAddressRefsAllVersions(nil, glamsterdamTokensTestChainSel, "LombardVerifier", versions)
		require.Empty(t, got)
	})

	t.Run("real Lombard/CCTP version lists resolve their own contract types", func(t *testing.T) {
		addrs := []datastore.AddressRef{
			{ChainSelector: glamsterdamTokensTestChainSel, Address: addrA.Hex(), Type: datastore.ContractType(lombard_verifier.ContractType), Version: lombard_verifier.Version},
			{ChainSelector: glamsterdamTokensTestChainSel, Address: addrB.Hex(), Type: datastore.ContractType(lombard_verifier.ContractType), Version: v200},
		}
		got := resolveAddressRefsAllVersions(addrs, glamsterdamTokensTestChainSel, lombard_verifier.ContractType, supportedLombardVerifierVersions)
		require.ElementsMatch(t, []common.Address{addrA, addrB}, got)
	})
}
