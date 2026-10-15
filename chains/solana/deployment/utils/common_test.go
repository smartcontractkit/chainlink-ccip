package utils

import (
	"errors"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestResolveMCMSInstance(t *testing.T) {
	t.Parallel()

	const (
		chainSel  = uint64(124615329519749607)
		qualifier = "CLLCCIP"
		seedStr   = "BG7wilBWT4mc6p9yFnmfcu3yX7r9dazl"
	)
	programID := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	var seed state.PDASeed
	copy(seed[:], seedStr)

	ref := func(typ, address, q string) cldf_datastore.AddressRef {
		return cldf_datastore.AddressRef{
			Address:       address,
			ChainSelector: chainSel,
			Type:          cldf_datastore.ContractType(typ),
			Version:       common_utils.Version_1_6_0,
			Qualifier:     q,
		}
	}
	instanceRef := ref("RBACTimelock", programID.String()+"."+seedStr, qualifier)
	seedRef := ref("RBACTimelockSeed", seedStr, qualifier)
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
			refs:    []cldf_datastore.AddressRef{instanceRef, ref("RBACTimelockSeed", "other", qualifier)},
			wantErr: "disagree",
		},
		{
			name:    "malformed instance ref",
			refs:    []cldf_datastore.AddressRef{ref("RBACTimelock", "not-a-contract-address", qualifier)},
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ResolveMCMSInstance(tt.refs, chainSel, common_utils.RBACTimelock, qualifier)
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

func TestGetTimelockSignerPDA(t *testing.T) {
	t.Parallel()

	programID := solana.MustPublicKeyFromBase58("DoajfR5tK24xVw51fWcawUZWhAXD8yrBJVacc13neVQA")
	const seedStr = "BG7wilBWT4mc6p9yFnmfcu3yX7r9dazl"
	var seed state.PDASeed
	copy(seed[:], seedStr)
	refs := []cldf_datastore.AddressRef{{
		Address:       programID.String() + "." + seedStr,
		ChainSelector: 1,
		Type:          cldf_datastore.ContractType(common_utils.RBACTimelock),
		Version:       common_utils.Version_1_6_0,
		Qualifier:     common_utils.CLLQualifier,
	}}

	got, err := GetTimelockSignerPDA(refs, 1, common_utils.CLLQualifier)
	require.NoError(t, err)
	require.Equal(t, state.GetTimelockSignerPDA(programID, seed), got)

	_, err = GetTimelockSignerPDA(refs, 2, common_utils.CLLQualifier)
	require.ErrorIs(t, err, ErrMCMSInstanceNotFound)
}
