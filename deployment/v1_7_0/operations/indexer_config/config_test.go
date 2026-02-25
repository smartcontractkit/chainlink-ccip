package indexer_config_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/indexer_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	testChainSel       = uint64(16015286601757825753)
	committeeQualifier = "default"
	addrOther          = "0x1111111111111111111111111111111111111111"
	addr1_7_0          = "0x958F44bbA928E294D5199870e330c4f30E5E5Ed4"
)

var otherVersion = semver.MustParse("10.0.0")

func TestBuildConfig_ReturnsOnlyCommitteeVerifier1_7_0AddressesWhenDatastoreHasBothVersions(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSel

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Qualifier:     committeeQualifier,
		Address:       addrOther,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Qualifier:     committeeQualifier,
		Address:       addr1_7_0,
		Version:       committee_verifier.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(
		env.OperationsBundle,
		indexer_config.BuildConfig,
		indexer_config.BuildConfigDeps{Env: env},
		indexer_config.BuildConfigInput{
			ServiceIdentifier:                "test-indexer",
			CommitteeVerifierNameToQualifier: map[string]string{"committee": committeeQualifier},
			ChainSelectors:                   []uint64{chainSel},
		},
	)
	require.NoError(t, err)
	require.NotNil(t, report.Output.Verifiers)

	var committee *indexer_config.GeneratedVerifier
	for i := range report.Output.Verifiers {
		if report.Output.Verifiers[i].Name == "committee" {
			committee = &report.Output.Verifiers[i]
			break
		}
	}
	require.NotNil(t, committee, "expected verifier named committee")
	assert.Contains(t, committee.IssuerAddresses, addr1_7_0, "BuildConfig must resolve CommitteeVerifier 1.7.0 when both versions exist")
	assert.NotContains(t, committee.IssuerAddresses, addrOther, "BuildConfig must not include 1.6.0 address when filtering by version")
}
