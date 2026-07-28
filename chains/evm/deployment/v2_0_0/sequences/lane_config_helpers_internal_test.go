package sequences

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
)

const testChainSelector = uint64(5009297550715157269)

func verifierRef(address string) datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: testChainSelector,
		Type:          datastore.ContractType(committee_verifier.ContractType),
		Address:       address,
	}
}

func resolverRef(address string) datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: testChainSelector,
		Type:          datastore.ContractType(CommitteeVerifierResolverType),
		Address:       address,
	}
}

func TestExtractCommitteeVerifierAddresses(t *testing.T) {
	const (
		verifierAddr = "0x0000000000000000000000000000000000000001"
		resolverAddr = "0x0000000000000000000000000000000000000002"
		otherAddr    = "0x0000000000000000000000000000000000000003"
	)

	t.Run("verifier and resolver", func(t *testing.T) {
		verifier, resolver, err := extractCommitteeVerifierAddresses(
			[]datastore.AddressRef{verifierRef(verifierAddr), resolverRef(resolverAddr)},
			testChainSelector,
		)
		require.NoError(t, err)
		require.Equal(t, verifierAddr, verifier)
		require.Equal(t, resolverAddr, resolver)
	})

	t.Run("unrelated refs are ignored", func(t *testing.T) {
		unrelated := datastore.AddressRef{
			ChainSelector: testChainSelector,
			Type:          datastore.ContractType("Router"),
			Address:       otherAddr,
		}
		verifier, resolver, err := extractCommitteeVerifierAddresses(
			[]datastore.AddressRef{unrelated, verifierRef(verifierAddr), resolverRef(resolverAddr)},
			testChainSelector,
		)
		require.NoError(t, err)
		require.Equal(t, verifierAddr, verifier)
		require.Equal(t, resolverAddr, resolver)
	})

	t.Run("missing verifier", func(t *testing.T) {
		_, _, err := extractCommitteeVerifierAddresses(
			[]datastore.AddressRef{resolverRef(resolverAddr)},
			testChainSelector,
		)
		require.ErrorContains(t, err, "committee verifier contract not found")
	})

	t.Run("missing resolver", func(t *testing.T) {
		_, _, err := extractCommitteeVerifierAddresses(
			[]datastore.AddressRef{verifierRef(verifierAddr)},
			testChainSelector,
		)
		require.ErrorContains(t, err, "committee verifier resolver contract not found")
	})

	t.Run("duplicate verifier", func(t *testing.T) {
		_, _, err := extractCommitteeVerifierAddresses(
			[]datastore.AddressRef{verifierRef(verifierAddr), verifierRef(otherAddr), resolverRef(resolverAddr)},
			testChainSelector,
		)
		require.ErrorContains(t, err, "duplicate committee verifier contract")
	})

	t.Run("duplicate resolver", func(t *testing.T) {
		_, _, err := extractCommitteeVerifierAddresses(
			[]datastore.AddressRef{verifierRef(verifierAddr), resolverRef(resolverAddr), resolverRef(otherAddr)},
			testChainSelector,
		)
		require.ErrorContains(t, err, "duplicate committee verifier resolver contract")
	})

	t.Run("no refs", func(t *testing.T) {
		_, _, err := extractCommitteeVerifierAddresses(nil, testChainSelector)
		require.ErrorContains(t, err, "committee verifier contract not found")
	})
}
