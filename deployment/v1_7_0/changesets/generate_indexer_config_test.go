package changesets_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

func TestGenerateIndexerConfig_ValidatesServiceIdentifier(t *testing.T) {
	changeset := changesets.GenerateIndexerConfig()

	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateIndexerConfigCfg{
		ServiceIdentifier:                "",
		CommitteeVerifierNameToQualifier: map[string]string{"verifier": "default"},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "service identifier is required")
}

func TestGenerateIndexerConfig_ValidatesVerifierNameToQualifier(t *testing.T) {
	changeset := changesets.GenerateIndexerConfig()

	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateIndexerConfigCfg{
		ServiceIdentifier:                "default-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one verifier name to qualifier mapping is required")
}

func TestGenerateIndexerConfig_ValidatesSourceChainSelectors(t *testing.T) {
	changeset := changesets.GenerateIndexerConfig()

	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateIndexerConfigCfg{
		ServiceIdentifier:                "default-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"verifier": "default"},
		ChainSelectors:                   []uint64{1234},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "selector 1234 is not available in environment")
}

func TestGenerateIndexerConfig_GeneratesCorrectConfigWithMultipleVerifiers(t *testing.T) {
	chainSelectors := []uint64{1001, 1002}
	verifierNameToQualifier := map[string]string{
		"Verifier A": "qualifier-a",
		"Verifier B": "qualifier-b",
	}

	ds := datastore.NewMemoryDataStore()
	for _, qualifier := range verifierNameToQualifier {
		for _, sel := range chainSelectors {
			err := ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: sel,
				Qualifier:     qualifier,
				Type:          datastore.ContractType(committee_verifier.ResolverType),
				Address:       fmt.Sprintf("0x%s_%d", qualifier, sel),
				Version:       committee_verifier.Version,
			})
			require.NoError(t, err)
		}
	}

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateIndexerConfig()
	output, err := cs.Apply(env, changesets.GenerateIndexerConfigCfg{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: verifierNameToQualifier,
		ChainSelectors:                   chainSelectors,
	})
	require.NoError(t, err)

	require.NotNil(t, output.DataStore)
	cfg, err := ccv.GetIndexerConfig(output.DataStore.Seal(), "test-indexer")
	require.NoError(t, err)

	assert.Len(t, cfg.Verifier, len(verifierNameToQualifier))

	verifierByName := make(map[string]bool)
	for _, v := range cfg.Verifier {
		verifierByName[v.Name] = true
		assert.Len(t, v.IssuerAddresses, len(chainSelectors),
			"verifier %s should have %d issuer addresses", v.Name, len(chainSelectors))
	}

	for name := range verifierNameToQualifier {
		require.True(t, verifierByName[name], "expected verifier config for name %s", name)
	}
}

func TestGenerateIndexerConfig_PreservesExistingAggregatorConfig(t *testing.T) {
	chainSelectors := []uint64{1001, 1002}
	committee := testCommittee

	ds := datastore.NewMemoryDataStore()
	for _, sel := range chainSelectors {
		err := ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Qualifier:     committee,
			Type:          datastore.ContractType(committee_verifier.ResolverType),
			Address:       fmt.Sprintf("0xverifier_%d", sel),
			Version:       committee_verifier.Version,
		})
		require.NoError(t, err)
	}

	existingAggregatorConfig := &offchain.Committee{
		DestinationVerifiers: map[offchain.DestinationSelector]string{
			"1001": "0xdest_1001",
			"1002": "0xdest_1002",
		},
		QuorumConfigs: map[offchain.SourceSelector]*offchain.QuorumConfig{
			"1001": {
				SourceVerifierAddress: "0xsource_1001",
				Signers:               []offchain.Signer{{Address: "0xsigner1"}},
				Threshold:             1,
			},
		},
	}
	err := ccv.SaveAggregatorConfig(ds, "existing-aggregator", existingAggregatorConfig)
	require.NoError(t, err)

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateIndexerConfig()
	output, err := cs.Apply(env, changesets.GenerateIndexerConfigCfg{
		ServiceIdentifier:                "new-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"Test Verifier": committee},
		ChainSelectors:                   chainSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	_, err = ccv.GetIndexerConfig(outputSealed, "new-indexer")
	require.NoError(t, err, "new indexer config should be present")

	retrievedAggregatorConfig, err := ccv.GetAggregatorConfig(outputSealed, "existing-aggregator")
	require.NoError(t, err, "existing aggregator config should be preserved")
	assert.Equal(t, existingAggregatorConfig, retrievedAggregatorConfig, "aggregator config should be unchanged")
}
