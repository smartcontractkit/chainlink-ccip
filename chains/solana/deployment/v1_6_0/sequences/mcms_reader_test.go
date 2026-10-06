package sequences

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

const testSelector uint64 = 124615329519749607

// newMultiMCMSEnv mirrors Solana mainnet: several MCMS instances, distinguished by qualifier, sharing
// unqualified program refs. Addresses are "<type>-<qualifier>" so the test can tell them apart.
func newMultiMCMSEnv(t *testing.T) deployment.Environment {
	ds := cldf_datastore.NewMemoryDataStore()
	add := func(contractType deployment.ContractType, qualifier string) {
		require.NoError(t, ds.Addresses().Add(cldf_datastore.AddressRef{
			Address:       string(contractType) + "-" + qualifier,
			ChainSelector: testSelector,
			Type:          cldf_datastore.ContractType(contractType),
			Version:       common_utils.Version_1_6_0,
			Qualifier:     qualifier,
		}))
	}
	add(utils.McmProgramType, "")
	// Non-default qualifiers are added first so a qualifier-agnostic lookup would pick them.
	for _, q := range []string{"UltraFastCurse", "RMNMCMS", common_utils.CLLQualifier} {
		add(common_utils.RBACTimelock, q)
		add(utils.ProposerAccessControllerAccount, q)
		add(utils.CancellerAccessControllerAccount, q)
		add(utils.BypasserAccessControllerAccount, q)
	}
	return deployment.Environment{DataStore: ds.Seal()}
}

func TestMCMSQualifierResolution(t *testing.T) {
	e := newMultiMCMSEnv(t)
	a := &SolanaAdapter{}

	tests := []struct {
		name          string
		qualifier     string
		wantQualifier string
	}{
		{name: "empty defaults to CLLCCIP", qualifier: "", wantQualifier: common_utils.CLLQualifier},
		{name: "CLLCCIP", qualifier: common_utils.CLLQualifier, wantQualifier: common_utils.CLLQualifier},
		{name: "UltraFastCurse", qualifier: "UltraFastCurse", wantQualifier: "UltraFastCurse"},
		{name: "RMNMCMS", qualifier: "RMNMCMS", wantQualifier: "RMNMCMS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := mcms_utils.Input{Qualifier: tt.qualifier}
			q := mcmsQualifier(input)
			require.Equal(t, tt.wantQualifier, q)

			timelock, err := a.GetTimelockRef(e, testSelector, input)
			require.NoError(t, err)
			require.Equal(t, tt.wantQualifier, timelock.Qualifier)

			for _, ct := range []deployment.ContractType{
				utils.ProposerAccessControllerAccount,
				utils.CancellerAccessControllerAccount,
				utils.BypasserAccessControllerAccount,
			} {
				ref, err := getMCMSAccountRef(e, testSelector, ct, q)
				require.NoError(t, err)
				require.Equal(t, tt.wantQualifier, ref.Qualifier, "%s resolved from the wrong MCMS instance", ct)
			}

			// The MCM program is shared and unqualified; it must resolve regardless of the input qualifier.
			mcm, err := a.GetMCMSRef(e, testSelector, input)
			require.NoError(t, err)
			require.Equal(t, string(utils.McmProgramType)+"-", mcm.Address)
		})
	}
}

func TestMCMSQualifierResolution_UnknownQualifierFails(t *testing.T) {
	e := newMultiMCMSEnv(t)

	_, err := getMCMSAccountRef(e, testSelector, utils.ProposerAccessControllerAccount, "DoesNotExist")
	require.ErrorContains(t, err, `qualifier "DoesNotExist" not found`)
}

// Callers (e.g. token expansion) rely on GetTimelockRef/GetMCMSRef returning an empty ref, not an error,
// on chains where MCMS is not deployed.
func TestMCMSRefs_NoMCMSReturnsEmptyRef(t *testing.T) {
	e := deployment.Environment{DataStore: cldf_datastore.NewMemoryDataStore().Seal()}
	a := &SolanaAdapter{}

	timelock, err := a.GetTimelockRef(e, testSelector, mcms_utils.Input{})
	require.NoError(t, err)
	require.Empty(t, timelock.Address)

	mcm, err := a.GetMCMSRef(e, testSelector, mcms_utils.Input{})
	require.NoError(t, err)
	require.Empty(t, mcm.Address)
}
