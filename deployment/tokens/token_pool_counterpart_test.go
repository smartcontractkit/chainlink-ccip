package tokens

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
)

// counterpartTestAdapter derives a per-token pool address, as Solana does with the pool config PDA.
type counterpartTestAdapter struct{ reverseTestAdapter }

func (a *counterpartTestAdapter) DeriveTokenPoolCounterpart(_ cldf.Environment, _ uint64, tokenPool []byte, token []byte) ([]byte, error) {
	return append(append(append([]byte{}, tokenPool...), ':'), token...), nil
}

func TestTokenPoolCounterpartAddress(t *testing.T) {
	family, err := chainsel.GetSelectorFamily(reverseTestRemoteSel)
	require.NoError(t, err)
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(family, reverseTestIdentityNormalizer{})

	identityVersion := semver.MustParse("9.9.1")
	derivedVersion := semver.MustParse("9.9.2")
	GetTokenAdapterRegistry().RegisterTokenAdapter(family, identityVersion, &reverseTestAdapter{})
	GetTokenAdapterRegistry().RegisterTokenAdapter(family, derivedVersion, &counterpartTestAdapter{})

	token := datastore.AddressRef{Address: "mint"}

	got, err := TokenPoolCounterpartAddress(cldf.Environment{}, reverseTestRemoteSel, datastore.AddressRef{Address: "pool", Version: identityVersion}, token)
	require.NoError(t, err)
	require.Equal(t, "pool", got, "a pool that is its own counterpart keeps its address")

	got, err = TokenPoolCounterpartAddress(cldf.Environment{}, reverseTestRemoteSel, datastore.AddressRef{Address: "pool", Version: derivedVersion}, token)
	require.NoError(t, err)
	require.Equal(t, "pool:mint", got, "a per-token counterpart is used when the family derives one")
}
