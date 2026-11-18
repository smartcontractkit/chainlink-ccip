package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// A Solana pool program is one address registered under several qualifiers: the program's own
// qualifier (e.g. CLLCCIP) and the generic TokenPoolProgram alias, both at the program's version.
// ResolveTokenPoolRef must pick the latest version instead of failing on the multi-hit lookup.
func TestResolveTokenPoolRef_SolanaSharedProgramRefs(t *testing.T) {
	t.Parallel()

	const (
		solanaMainnet = uint64(124615329519749607)
		programID     = "41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB"
	)

	ref := func(version, qualifier string) datastore.AddressRef {
		return datastore.AddressRef{
			ChainSelector: solanaMainnet,
			Address:       programID,
			Type:          datastore.ContractType("BurnMintTokenPool"),
			Version:       semver.MustParse(version),
			Qualifier:     qualifier,
		}
	}
	newEnv := func(t *testing.T, refs ...datastore.AddressRef) deployment.Environment {
		t.Helper()

		ds := datastore.NewMemoryDataStore()
		for _, r := range refs {
			require.NoError(t, ds.Addresses().Add(r))
		}

		return deployment.Environment{DataStore: ds.Seal()}
	}

	tests := []struct {
		name         string
		refs         []datastore.AddressRef
		expectedErr  string
		expectedVer  string
		expectedQual string
	}{
		{
			name:         "single hit returns the ref",
			refs:         []datastore.AddressRef{ref("2.1.0", "TokenPoolProgram")},
			expectedVer:  "2.1.0",
			expectedQual: "TokenPoolProgram",
		},
		{
			name: "multi hit agreeing on type and address resolves to the latest version",
			refs: []datastore.AddressRef{
				ref("2.0.0", "TESTTR"),
				ref("2.1.0", "CLLCCIP"),
				ref("2.1.0", "TokenPoolProgram"),
			},
			expectedVer: "2.1.0",
		},
		{
			name: "multi hit disagreeing on type errors",
			refs: []datastore.AddressRef{
				ref("2.1.0", "CLLCCIP"),
				{
					ChainSelector: solanaMainnet,
					Address:       programID,
					Type:          datastore.ContractType("USDCTokenPool"),
					Version:       semver.MustParse("2.1.0"),
					Qualifier:     "USDC",
				},
			},
			expectedErr: "multiple refs found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tokens.ResolveTokenPoolRef(
				newEnv(t, tt.refs...),
				tokens.GetTokenAdapterRegistry(),
				solanaMainnet,
				datastore.AddressRef{ChainSelector: solanaMainnet, Address: programID},
			)
			if tt.expectedErr != "" {
				require.ErrorContains(t, err, tt.expectedErr)

				return
			}
			require.NoError(t, err)
			require.Equal(t, programID, got.Address)
			require.Equal(t, tt.expectedVer, got.Version.String())
			if tt.expectedQual != "" {
				require.Equal(t, tt.expectedQual, got.Qualifier)
			}
		})
	}
}
