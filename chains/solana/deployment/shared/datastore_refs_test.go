package shared

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcmscontracts "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/contracts/mcms"
)

func testRef(chainSel uint64, addr string, t cldf.ContractType, ver *semver.Version, qualifier string, labels ...string) datastore.AddressRef {
	ref := datastore.AddressRef{
		ChainSelector: chainSel,
		Address:       addr,
		Type:          datastore.ContractType(t),
		Qualifier:     qualifier,
	}
	if ver != nil {
		v := *ver
		ref.Version = &v
	}
	if len(labels) > 0 {
		ref.Labels = datastore.NewLabelSet(labels...)
	}
	return ref
}

func TestMCMSBundleRefs_Isolation(t *testing.T) {
	v10 := *semver.MustParse("1.0.0")
	refs := []datastore.AddressRef{
		testRef(1, "0xQualifiedTimelock", mcmscontracts.RBACTimelock, &v10, DefaultMCMSQualifier),
		testRef(1, "0xQualifiedProposer", mcmscontracts.ProposerManyChainMultisig, &v10, DefaultMCMSQualifier),
		testRef(1, "0xEmptyTimelock", mcmscontracts.RBACTimelock, &v10, ""),
		testRef(1, "0xEmptyCallProxy", mcmscontracts.CallProxy, &v10, ""),
		testRef(1, "0xVersionless", mcmscontracts.RBACTimelock, nil, DefaultMCMSQualifier),
		testRef(1, "0xSuperseded", mcmscontracts.CallProxy, &v10, DefaultMCMSQualifier, SupersededLabel),
	}

	bundle, err := MCMSBundleRefs(refs, 1, DefaultMCMSQualifier)
	require.NoError(t, err)
	require.Len(t, bundle, 2) // qualified rows only; versionless and superseded dropped
	for _, ref := range bundle {
		require.Equal(t, DefaultMCMSQualifier, ref.Qualifier)
	}

	// custom qualifier with no rows fails closed — no empty-qualifier fallback
	_, err = MCMSBundleRefs(refs, 1, "RMNMCMS")
	require.ErrorContains(t, err, "no mcms refs for chain 1")

	// default qualifier with no rows yields an empty bundle (MaybeLoad semantics)
	bundle, err = MCMSBundleRefs(refs, 7, DefaultMCMSQualifier)
	require.NoError(t, err)
	require.Empty(t, bundle)

	// An unqualified bundle IS the legacy fallback for the default qualifier
	// (singletons and DeployMCMSWithTimelockV2 test deployments are empty-qualified);
	// dedicated bundles (RMNMCMS etc.) never fall back.
	emptyOnly := []datastore.AddressRef{testRef(1, "0xEmptyOnly", mcmscontracts.RBACTimelock, &v10, "")}
	bundle, err = MCMSBundleRefs(emptyOnly, 1, DefaultMCMSQualifier)
	require.NoError(t, err)
	require.Len(t, bundle, 1)
	require.Empty(t, bundle[0].Qualifier)

	// the fallback never fires for a custom qualifier
	_, err = MCMSBundleRefs(emptyOnly, 1, "RMNMCMS")
	require.ErrorContains(t, err, "no mcms refs for chain 1")

	// two active refs, one identity
	dup := append([]datastore.AddressRef(nil), refs...)
	dup = append(dup, testRef(1, "0xQualifiedTimelock2", mcmscontracts.RBACTimelock, &v10, DefaultMCMSQualifier))
	_, err = MCMSBundleRefs(dup, 1, DefaultMCMSQualifier)
	require.ErrorContains(t, err, "both")
}
