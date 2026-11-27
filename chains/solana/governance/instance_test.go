package governance

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

var (
	version160 = semver.MustParse("1.6.0")
	version200 = semver.MustParse("2.0.0")
)

const (
	testChainSel  = uint64(124615329519749607)
	testQualifier = "CLLCCIP"
	testSeedStr   = "BG7wilBWT4mc6p9yFnmfcu3yX7r9dazl"
)

func testRef(version *semver.Version, typ, address, qualifier string) cldf_datastore.AddressRef {
	return cldf_datastore.AddressRef{
		Address:       address,
		ChainSelector: testChainSel,
		Type:          cldf_datastore.ContractType(typ),
		Version:       version,
		Qualifier:     qualifier,
	}
}

func TestResolveMCMSInstance(t *testing.T) {
	t.Parallel()

	programID := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	var seed state.PDASeed
	copy(seed[:], testSeedStr)

	// the resolver must behave the same at every version: it only reads the version it is given
	for _, version := range []*semver.Version{version160, version200} {
		ref := func(typ, address, q string) cldf_datastore.AddressRef { return testRef(version, typ, address, q) }
		instanceRef := ref("RBACTimelock", programID.String()+"."+testSeedStr, testQualifier)
		seedRef := ref("RBACTimelockSeed", testSeedStr, testQualifier)
		programRef := ref("RBACTimelockProgram", programID.String(), "")

		tests := []struct {
			name        string
			refs        []cldf_datastore.AddressRef
			wantMissing []cldf_datastore.AddressRef
			wantErr     string
			notFound    bool
		}{
			{
				name: "both refs present",
				refs: []cldf_datastore.AddressRef{instanceRef, seedRef, programRef},
			},
			{
				name:        "seed ref missing is derived from the instance ref",
				refs:        []cldf_datastore.AddressRef{instanceRef, programRef},
				wantMissing: []cldf_datastore.AddressRef{seedRef},
			},
			{
				name:        "instance ref missing is rebuilt from program and seed refs",
				refs:        []cldf_datastore.AddressRef{seedRef, programRef},
				wantMissing: []cldf_datastore.AddressRef{instanceRef},
			},
			{
				name:     "no instance or seed ref",
				refs:     []cldf_datastore.AddressRef{programRef},
				wantErr:  "no RBACTimelock or RBACTimelockSeed ref",
				notFound: true,
			},
			{
				name:    "seed ref without program ref",
				refs:    []cldf_datastore.AddressRef{seedRef},
				wantErr: "no RBACTimelockProgram ref",
			},
			{
				name:    "instance and seed refs disagree",
				refs:    []cldf_datastore.AddressRef{instanceRef, ref("RBACTimelockSeed", "other", testQualifier)},
				wantErr: "disagree",
			},
			{
				name:    "malformed instance ref",
				refs:    []cldf_datastore.AddressRef{ref("RBACTimelock", "not-a-contract-address", testQualifier)},
				wantErr: "invalid RBACTimelock ref",
			},
			{
				name:     "other qualifier is not used",
				refs:     []cldf_datastore.AddressRef{ref("RBACTimelock", instanceRef.Address, "RMNMCMS"), programRef},
				wantErr:  "no RBACTimelock or RBACTimelockSeed ref",
				notFound: true,
			},
		}
		for _, tt := range tests {
			t.Run(version.String()+"/"+tt.name, func(t *testing.T) {
				t.Parallel()
				got, err := ResolveMCMSInstance(tt.refs, testChainSel, RBACTimelock, version, testQualifier)
				if tt.wantErr != "" {
					require.ErrorContains(t, err, tt.wantErr)
					require.Equal(t, tt.notFound, errors.Is(err, ErrMCMSInstanceNotFound))
					return
				}
				require.NoError(t, err)
				require.Equal(t, programID, got.ProgramID)
				require.Equal(t, seed, got.Seed)
				require.Equal(t, tt.wantMissing, got.MissingRefs)
			})
		}
	}

	t.Run("unsupported instance type", func(t *testing.T) {
		t.Parallel()
		_, err := ResolveMCMSInstance(nil, testChainSel, "Router", version160, testQualifier)
		require.ErrorContains(t, err, `unsupported solana mcms instance type "Router"`)
	})
}

// The resolver reads only the version it is given: a datastore holding the other version is empty to it.
func TestResolveMCMSInstance_VersionIsolation(t *testing.T) {
	t.Parallel()
	programID := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	addr := programID.String() + "." + testSeedStr

	only160 := []cldf_datastore.AddressRef{testRef(version160, "RBACTimelock", addr, testQualifier)}
	only200 := []cldf_datastore.AddressRef{testRef(version200, "RBACTimelock", addr, testQualifier)}

	_, err := ResolveMCMSInstance(only160, testChainSel, RBACTimelock, version160, testQualifier)
	require.NoError(t, err)
	_, err = ResolveMCMSInstance(only160, testChainSel, RBACTimelock, version200, testQualifier)
	require.ErrorIs(t, err, ErrMCMSInstanceNotFound)

	_, err = ResolveMCMSInstance(only200, testChainSel, RBACTimelock, version200, testQualifier)
	require.NoError(t, err)
	_, err = ResolveMCMSInstance(only200, testChainSel, RBACTimelock, version160, testQualifier)
	require.ErrorIs(t, err, ErrMCMSInstanceNotFound)
}

func TestGetTimelockSignerPDA(t *testing.T) {
	t.Parallel()

	programID := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	var seed state.PDASeed
	copy(seed[:], testSeedStr)
	for _, version := range []*semver.Version{version160, version200} {
		refs := []cldf_datastore.AddressRef{{
			Address:       programID.String() + "." + testSeedStr,
			ChainSelector: 1,
			Type:          RBACTimelock,
			Version:       version,
			Qualifier:     testQualifier,
		}}

		got, err := GetTimelockSignerPDA(refs, 1, version, testQualifier)
		require.NoError(t, err)
		require.Equal(t, state.GetTimelockSignerPDA(programID, seed), got)

		_, err = GetTimelockSignerPDA(refs, 2, version, testQualifier)
		require.ErrorIs(t, err, ErrMCMSInstanceNotFound)
	}
}

func TestGetMCMSignerPDA(t *testing.T) {
	t.Parallel()

	programID := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	var seed state.PDASeed
	copy(seed[:], testSeedStr)
	refs := []cldf_datastore.AddressRef{{
		Address:       programID.String() + "." + testSeedStr,
		ChainSelector: 1,
		Type:          ProposerManyChainMultisig,
		Version:       version160,
		Qualifier:     testQualifier,
	}}

	got, err := GetMCMSignerPDA(refs, 1, ProposerManyChainMultisig, version160, testQualifier)
	require.NoError(t, err)
	require.Equal(t, state.GetMCMSignerPDA(programID, seed), got)

	_, err = GetMCMSignerPDA(refs, 1, CancellerManyChainMultisig, version160, testQualifier)
	require.ErrorIs(t, err, ErrMCMSInstanceNotFound)
}
