package v0_1_0

import (
	"encoding/json"
	"errors"
	"testing"

	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
	solanastateview "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared/stateview/solana"
)

func TestE2ETokenPoolRejectsEVMStep(t *testing.T) {
	t.Parallel()
	// existing pipeline inputs carry the field as null
	var cfg E2ETokenPoolConfig
	require.NoError(t, json.Unmarshal([]byte(`{"ConfigureTokenPoolContractsChangesets": null}`), &cfg))
	require.Empty(t, cfg.ConfigureTokenPoolContractsChangesets)

	require.NoError(t, json.Unmarshal([]byte(`{"ConfigureTokenPoolContractsChangesets": [{"MCMS": null}]}`), &cfg))
	_, err := E2ETokenPool(cldf.Environment{Logger: logger.Test(t)}, cfg)
	require.ErrorContains(t, err, "ConfigureTokenPoolContractsChangesets is no longer supported")
}

// fakeEVMRefResolver and fakeEVMTokenAdapter stand in for the EVM family's registered adapters; see
// the changesets package tests.
type fakeEVMRefResolver struct{}

func (fakeEVMRefResolver) ResolveTokenPoolRef(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, sel uint64, address string) (datastore.AddressRef, error) {
	return datastore.AddressRef{ChainSelector: sel, Address: address, Version: semver.MustParse("1.5.1"), Qualifier: testToken}, nil
}

func (fakeEVMRefResolver) ResolveTokenRef(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ datastore.DataStore, _ uint64, address string) (datastore.AddressRef, error) {
	return datastore.AddressRef{Address: address, Qualifier: "TEST"}, nil
}

type fakeEVMTokenAdapter struct {
	tokensapi.TokenAdapter
	poolHasDecimals bool
}

func (f *fakeEVMTokenAdapter) DeriveTokenDecimals(_ cldf.Environment, _ uint64, _ datastore.AddressRef, token []byte) (uint8, error) {
	if token != nil {
		return 6, nil
	}
	if !f.poolHasDecimals {
		return 0, errors.New("getTokenDecimals reverted")
	}
	return 18, nil
}

const (
	testPool  = "0x00000000000000000000000000000000000000a1"
	testToken = "0x00000000000000000000000000000000000000b1"
)

// The token adapter registry is global; this is the only test in the package that registers.
func TestEVMRemoteConfigUsesLegacyLookup(t *testing.T) {
	evmSel := chainsel.TEST_90000001.Selector
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{ChainSelector: evmSel, Type: datastore.ContractType(shared.BurnMintTokenPool), Address: testPool, Version: semver.MustParse("1.5.1"), Qualifier: testPool}))
	e := cldf.Environment{
		DataStore:   ds.Seal(),
		BlockChains: cldf_chain.NewBlockChainsFromSlice([]cldf_chain.BlockChain{cldf_evm.Chain{Selector: evmSel}}),
	}
	adapter := &fakeEVMTokenAdapter{poolHasDecimals: true}
	tokensapi.GetTokenAdapterRegistry().RegisterTokenRefResolver(chainsel.FamilyEVM, fakeEVMRefResolver{})
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, semver.MustParse("1.5.1"), adapter)
	cfg := EVMRemoteConfig{TokenSymbol: "TEST", PoolType: shared.BurnMintTokenPool, PoolVersion: *semver.MustParse("1.5.1")}
	state := solanastateview.CCIPOnChainState{}

	require.NoError(t, cfg.Validate(e, state, evmSel))
	remote, err := getOnChainEVMPoolConfig(e, state, evmSel, cfg)
	require.NoError(t, err)
	require.Equal(t, append(make([]byte, 31), 0xb1), remote.TokenAddress.Address)
	require.Equal(t, append(make([]byte, 19), 0xa1), remote.PoolAddresses[0].Address)
	require.Equal(t, uint8(18), remote.Decimals)

	// v0_1_0 read the decimals from the pool only: a pool that can't report them is an error,
	// never the token's decimals
	adapter.poolHasDecimals = false
	_, err = getOnChainEVMPoolConfig(e, state, evmSel, cfg)
	require.ErrorContains(t, err, "failed to get token decimals")

	missing := cfg
	missing.TokenSymbol = "MISSING"
	require.ErrorContains(t, missing.Validate(e, state, evmSel), "token pool does not exist")
}
