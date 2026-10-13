package solana

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/shared"
)

func TestLoadOnchainStateSolana(t *testing.T) {
	t.Parallel()
	const solSel, otherSol = uint64(124615329519749607), uint64(16423721717087811551)
	link := solana.NewWallet().PublicKey()
	otherLink := solana.NewWallet().PublicKey()

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: solSel, Address: link.String(), Type: datastore.ContractType(shared.LinkToken), Version: semver.MustParse("1.6.0"),
	}))
	// a chain that isn't in the environment is not loaded
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: otherSol, Address: otherLink.String(), Type: datastore.ContractType(shared.LinkToken), Version: semver.MustParse("1.6.0"),
	}))
	e := cldf.Environment{
		DataStore:   ds.Seal(),
		BlockChains: cldf_chain.NewBlockChainsFromSlice([]cldf_chain.BlockChain{cldf_solana.Chain{Selector: solSel}}),
	}

	state, err := LoadOnchainStateSolana(e)
	require.NoError(t, err)
	require.Len(t, state.SolChains, 1)
	require.Equal(t, link, state.SolChains[solSel].LinkToken)
	require.Equal(t, map[uint64]struct{}{solSel: {}}, state.SupportedChains())

	_, err = LoadOnchainStateSolana(cldf.Environment{})
	require.Error(t, err)
}
