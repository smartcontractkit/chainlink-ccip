package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/lombard_verifier"
	onrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

const (
	testServiceIdentifier = "test-token-verifier"
	testCCTPQualifier     = "cctp-v1"
	testLombardQualifier  = "lombard-v1"
	testCCTPVerifierID    = "cctp-custom"
	testLombardVerifierID = "lombard-custom"
	testPyroscopeURL      = "http://pyroscope:4040"
	verifierTypeCCTP      = "cctp"
	verifierTypeLombard   = "lombard"
)

func TestGenerateTokenVerifierConfig_ValidatesServiceIdentifier(t *testing.T) {
	changeset := changesets.GenerateTokenVerifierConfig()
	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: "",
		CCTP: sequences.CCTPConfigInput{
			Qualifier: testCCTPQualifier,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "service identifier is required")
}

func TestGenerateTokenVerifierConfig_ValidatesChainSelectors(t *testing.T) {
	changeset := changesets.GenerateTokenVerifierConfig()
	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    []uint64{999999},
		CCTP: sequences.CCTPConfigInput{
			Qualifier: testCCTPQualifier,
		},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "selector 999999 is not available in environment")
}

func TestGenerateTokenVerifierConfig_AllowsEmptyQualifiers(t *testing.T) {
	changeset := changesets.GenerateTokenVerifierConfig()
	env := createTestEnvironmentForValidation(t)

	// Empty qualifiers should not cause validation error
	err := changeset.VerifyPreconditions(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		CCTP:              sequences.CCTPConfigInput{},
		Lombard:           sequences.LombardConfigInput{},
	})
	require.NoError(t, err)
}

func TestGenerateTokenVerifierConfig_GeneratesConfigWithBothVerifiers(t *testing.T) {
	chainSelectors := []uint64{1001, 1002}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, true)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		PyroscopeURL:      testPyroscopeURL,
		Monitoring:        createTestMonitoringInput(),
		CCTP: sequences.CCTPConfigInput{
			Qualifier:  testCCTPQualifier,
			VerifierID: testCCTPVerifierID,
		},
		Lombard: sequences.LombardConfigInput{
			Qualifier:  testLombardQualifier,
			VerifierID: testLombardVerifierID,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Verify basic config
	assert.Equal(t, testPyroscopeURL, cfg.PyroscopeURL)
	assert.Len(t, cfg.OnRampAddresses, len(chainSelectors))
	assert.Len(t, cfg.RMNRemoteAddresses, len(chainSelectors))

	// Verify monitoring config
	assert.True(t, cfg.Monitoring.Enabled)
	assert.Equal(t, "beholder", cfg.Monitoring.Type)

	// Verify token verifiers
	assert.Len(t, cfg.TokenVerifiers, 2)

	// Find CCTP and Lombard verifiers
	var cctpVerifier, lombardVerifier *offchain.TokenVerifierPluginConfig
	for i := range cfg.TokenVerifiers {
		v := &cfg.TokenVerifiers[i]
		switch v.Type {
		case verifierTypeCCTP:
			cctpVerifier = v
		case verifierTypeLombard:
			lombardVerifier = v
		}
	}

	require.NotNil(t, cctpVerifier, "CCTP verifier should be present")
	require.NotNil(t, lombardVerifier, "Lombard verifier should be present")
}

func TestGenerateTokenVerifierConfig_GeneratesConfigWithOnlyCCTP(t *testing.T) {
	chainSelectors := []uint64{1001, 1002}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, false)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		CCTP: sequences.CCTPConfigInput{
			Qualifier: testCCTPQualifier,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Should have only CCTP verifier
	assert.Len(t, cfg.TokenVerifiers, 1)
	assert.Equal(t, verifierTypeCCTP, cfg.TokenVerifiers[0].Type)
}

func TestGenerateTokenVerifierConfig_GeneratesConfigWithOnlyLombard(t *testing.T) {
	chainSelectors := []uint64{1001, 1002}
	ds := setupTokenVerifierDatastore(t, chainSelectors, false, true)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		Lombard: sequences.LombardConfigInput{
			Qualifier: testLombardQualifier,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Should have only Lombard verifier
	assert.Len(t, cfg.TokenVerifiers, 1)
	assert.Equal(t, verifierTypeLombard, cfg.TokenVerifiers[0].Type)
}

func TestGenerateTokenVerifierConfig_UsesDefaultVerifierIDWhenNotProvided(t *testing.T) {
	chainSelectors := []uint64{1001}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, true)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		CCTP: sequences.CCTPConfigInput{
			Qualifier: testCCTPQualifier,
			// Not specifying VerifierID
		},
		Lombard: sequences.LombardConfigInput{
			Qualifier: testLombardQualifier,
			// Not specifying VerifierID
		},
	})
	require.NoError(t, err)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Verify default verifier IDs are generated from qualifiers
	for _, v := range cfg.TokenVerifiers {
		switch v.Type {
		case verifierTypeCCTP:
			assert.Equal(t, "cctp-"+testCCTPQualifier, v.VerifierID)
		case verifierTypeLombard:
			assert.Equal(t, "lombard-"+testLombardQualifier, v.VerifierID)
		}
	}
}

func TestGenerateTokenVerifierConfig_UsesCustomVerifierIDWhenProvided(t *testing.T) {
	chainSelectors := []uint64{1001}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, true)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		CCTP: sequences.CCTPConfigInput{
			Qualifier:  testCCTPQualifier,
			VerifierID: testCCTPVerifierID,
		},
		Lombard: sequences.LombardConfigInput{
			Qualifier:  testLombardQualifier,
			VerifierID: testLombardVerifierID,
		},
	})
	require.NoError(t, err)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Verify custom verifier IDs are used
	for _, v := range cfg.TokenVerifiers {
		switch v.Type {
		case verifierTypeCCTP:
			assert.Equal(t, testCCTPVerifierID, v.VerifierID)
		case verifierTypeLombard:
			assert.Equal(t, testLombardVerifierID, v.VerifierID)
		}
	}
}

func TestGenerateTokenVerifierConfig_PreservesExistingConfigs(t *testing.T) {
	chainSelectors := []uint64{1001}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, false)

	// Add existing aggregator config
	existingAggregatorConfig := &offchain.Committee{
		DestinationVerifiers: map[offchain.DestinationSelector]string{
			"1001": "0xdest1001",
		},
		QuorumConfigs: map[offchain.SourceSelector]*offchain.QuorumConfig{
			"1001": {
				SourceVerifierAddress: "0xsource1001",
				Signers:               []offchain.Signer{{Address: "0xsigner1"}},
				Threshold:             1,
			},
		},
	}
	err := ccv.SaveAggregatorConfig(ds, "existing-aggregator", existingAggregatorConfig)
	require.NoError(t, err)

	// Add existing indexer config
	existingIndexerConfig := &offchain.IndexerGeneratedConfig{
		Verifier: []offchain.IndexerGeneratedVerifierConfig{
			{
				Name:            "test-verifier",
				IssuerAddresses: []string{"0xissuer1"},
			},
		},
	}
	err = ccv.SaveIndexerConfig(ds, "existing-indexer", existingIndexerConfig)
	require.NoError(t, err)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		CCTP: sequences.CCTPConfigInput{
			Qualifier: testCCTPQualifier,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	// Verify token verifier config was saved
	_, err = ccv.GetTokenVerifierConfig(outputSealed, testServiceIdentifier)
	require.NoError(t, err)

	// Verify existing configs are preserved
	retrievedAggregatorConfig, err := ccv.GetAggregatorConfig(outputSealed, "existing-aggregator")
	require.NoError(t, err)
	assert.Equal(t, existingAggregatorConfig, retrievedAggregatorConfig)

	retrievedIndexerConfig, err := ccv.GetIndexerConfig(outputSealed, "existing-indexer")
	require.NoError(t, err)
	assert.Equal(t, existingIndexerConfig, retrievedIndexerConfig)
}

func TestGenerateTokenVerifierConfig_UsesDefaultVerifierVersion(t *testing.T) {
	chainSelectors := []uint64{1001}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, true)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		CCTP: sequences.CCTPConfigInput{
			Qualifier: testCCTPQualifier,
			// Not specifying VerifierVersion
		},
		Lombard: sequences.LombardConfigInput{
			Qualifier: testLombardQualifier,
			// Not specifying VerifierVersion
		},
	})
	require.NoError(t, err)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Verify default verifier versions are used
	for _, v := range cfg.TokenVerifiers {
		switch v.Type {
		case verifierTypeCCTP:
			require.NotNil(t, v.CCTPConfig, "CCTP config should not be nil")
			assert.NotEmpty(t, v.CCTPConfig.VerifierVersion, "CCTP verifier version should not be empty")
			// Should use the default CCTP version (0x8e1d1a9d)
			assert.Equal(t, []byte{0x8e, 0x1d, 0x1a, 0x9d}, []byte(v.CCTPConfig.VerifierVersion))
		case verifierTypeLombard:
			require.NotNil(t, v.LombardConfig, "Lombard config should not be nil")
			assert.NotEmpty(t, v.LombardConfig.VerifierVersion, "Lombard verifier version should not be empty")
			// Should use the default Lombard version (0xf0f3a135)
			assert.Equal(t, []byte{0xf0, 0xf3, 0xa1, 0x35}, []byte(v.LombardConfig.VerifierVersion))
		}
	}
}

func TestGenerateTokenVerifierConfig_UsesCustomVerifierVersion(t *testing.T) {
	chainSelectors := []uint64{1001}
	ds := setupTokenVerifierDatastore(t, chainSelectors, true, true)

	env := deployment.Environment{
		Name:             "testnet",
		OperationsBundle: testutils.NewTestBundle(),
		BlockChains:      testutils.NewStubBlockChains(chainSelectors),
		DataStore:        ds.Seal(),
	}

	customCCTPVersion := []byte{0x11, 0x22, 0x33, 0x44}
	customLombardVersion := []byte{0xaa, 0xbb, 0xcc, 0xdd}

	cs := changesets.GenerateTokenVerifierConfig()
	output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigCfg{
		ServiceIdentifier: testServiceIdentifier,
		ChainSelectors:    chainSelectors,
		CCTP: sequences.CCTPConfigInput{
			Qualifier:       testCCTPQualifier,
			VerifierVersion: customCCTPVersion,
		},
		Lombard: sequences.LombardConfigInput{
			Qualifier:       testLombardQualifier,
			VerifierVersion: customLombardVersion,
		},
	})
	require.NoError(t, err)

	cfg, err := ccv.GetTokenVerifierConfig(output.DataStore.Seal(), testServiceIdentifier)
	require.NoError(t, err)

	// Verify custom verifier versions are used
	for _, v := range cfg.TokenVerifiers {
		switch v.Type {
		case verifierTypeCCTP:
			require.NotNil(t, v.CCTPConfig, "CCTP config should not be nil")
			assert.Equal(t, customCCTPVersion, []byte(v.CCTPConfig.VerifierVersion))
		case verifierTypeLombard:
			require.NotNil(t, v.LombardConfig, "Lombard config should not be nil")
			assert.Equal(t, customLombardVersion, []byte(v.LombardConfig.VerifierVersion))
		}
	}
}

// Helper functions

func setupTokenVerifierDatastore(t *testing.T, chainSelectors []uint64, includeCCTP, includeLombard bool) *datastore.MemoryDataStore {
	t.Helper()
	ds := datastore.NewMemoryDataStore()

	for _, sel := range chainSelectors {
		// Add required contracts (OnRamp and RMN Remote)
		err := ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(onrampoperations.ContractType),
			Address:       createTestAddress("onramp", sel),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: sel,
			Type:          datastore.ContractType(rmn_remote.ContractType),
			Address:       createTestAddress("rmn", sel),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// Add CCTP contracts if requested
		if includeCCTP {
			err = ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: sel,
				Qualifier:     testCCTPQualifier,
				Type:          datastore.ContractType(cctp_verifier.ContractType),
				Address:       createTestAddress("cctp-verifier", sel),
				Version:       semver.MustParse("1.0.0"),
			})
			require.NoError(t, err)

			err = ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: sel,
				Qualifier:     testCCTPQualifier,
				Type:          datastore.ContractType(cctp_verifier.ResolverType),
				Address:       createTestAddress("cctp-resolver", sel),
				Version:       semver.MustParse("1.0.0"),
			})
			require.NoError(t, err)
		}

		// Add Lombard contracts if requested
		if includeLombard {
			err = ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: sel,
				Qualifier:     testLombardQualifier,
				Type:          datastore.ContractType(lombard_verifier.ResolverType),
				Address:       createTestAddress("lombard-resolver", sel),
				Version:       semver.MustParse("1.0.0"),
			})
			require.NoError(t, err)
		}
	}

	return ds
}

func createTestAddress(prefix string, chainSelector uint64) string {
	return "0x" + prefix + "_" + string(rune(chainSelector))
}

func createTestMonitoringInput() shared.MonitoringInput {
	return shared.MonitoringInput{
		Enabled: true,
		Type:    "beholder",
		Beholder: shared.BeholderInput{
			InsecureConnection:       false,
			CACertFile:               "/path/to/ca.crt",
			OtelExporterGRPCEndpoint: "beholder:4317",
			OtelExporterHTTPEndpoint: "beholder:4318",
			LogStreamingEnabled:      true,
			MetricReaderInterval:     60,
			TraceSampleRatio:         0.1,
			TraceBatchTimeout:        5,
		},
	}
}
