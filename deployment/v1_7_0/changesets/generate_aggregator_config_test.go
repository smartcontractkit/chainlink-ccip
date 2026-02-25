package changesets_test

import (
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

func TestGenerateAggregatorConfig_ValidatesServiceIdentifier(t *testing.T) {
	changeset := changesets.GenerateAggregatorConfig()

	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateAggregatorConfigCfg{
		ServiceIdentifier:  "",
		CommitteeQualifier: "default",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "service identifier is required")
}

func TestGenerateAggregatorConfig_ValidatesCommitteeQualifier(t *testing.T) {
	changeset := changesets.GenerateAggregatorConfig()

	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateAggregatorConfigCfg{
		ServiceIdentifier:  "default-aggregator",
		CommitteeQualifier: "",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "committee qualifier is required")
}

func TestGenerateAggregatorConfig_ValidatesChainSelectors(t *testing.T) {
	changeset := changesets.GenerateAggregatorConfig()

	env := createTestEnvironmentForValidation(t)

	err := changeset.VerifyPreconditions(env, changesets.GenerateAggregatorConfigCfg{
		ServiceIdentifier:  "default-aggregator",
		CommitteeQualifier: "default",
		ChainSelectors:     []uint64{1234},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "selector 1234 is not available in environment")
}

func createTestEnvironmentForValidation(t *testing.T) deployment.Environment {
	t.Helper()
	bundle := testutils.NewTestBundle()

	return deployment.Environment{
		OperationsBundle: bundle,
	}
}

func TestGenerateAggregatorConfig_GeneratesCorrectConfigFromOnChainState(t *testing.T) {
	selectors := []uint64{
		chainsel.TEST_90000001.Selector,
		chainsel.TEST_90000002.Selector,
		chainsel.TEST_90000003.Selector,
	}
	targetCommittee := testCommittee
	otherCommittee := "other-committee"

	env, evmChains := testutils.NewSimulatedEVMEnvironment(t, selectors)
	require.Len(t, evmChains, 3)

	chain1, chain2, chain3 := evmChains[0], evmChains[1], evmChains[2]
	sel1, sel2, sel3 := selectors[0], selectors[1], selectors[2]

	// Use consistent signers for each source chain across all verifiers
	// This ensures deterministic expected config
	signersForSel1 := generateTestSigners(t, 3)
	signersForSel2 := generateTestSigners(t, 2)
	signersForSel3 := generateTestSigners(t, 2)
	otherSigners := generateTestSigners(t, 2)
	threshold := uint8(2)

	ds := datastore.NewMemoryDataStore()

	// Deploy target committee verifiers - all use same signers for same source chain
	// chain1 - accepts messages from chain2 and chain3
	addr1 := deployAndConfigureVerifier(t, env.OperationsBundle, chain1, targetCommittee, []sourceChainConfig{
		{selector: sel2, signers: signersForSel2, threshold: threshold},
		{selector: sel3, signers: signersForSel3, threshold: threshold},
	})
	addVerifierToDatastore(t, ds, sel1, targetCommittee, addr1)

	// chain2 - accepts messages from chain1 and chain3
	addr2 := deployAndConfigureVerifier(t, env.OperationsBundle, chain2, targetCommittee, []sourceChainConfig{
		{selector: sel1, signers: signersForSel1, threshold: threshold},
		{selector: sel3, signers: signersForSel3, threshold: threshold},
	})
	addVerifierToDatastore(t, ds, sel2, targetCommittee, addr2)

	// chain3 - accepts messages from chain1 and chain2
	addr3 := deployAndConfigureVerifier(t, env.OperationsBundle, chain3, targetCommittee, []sourceChainConfig{
		{selector: sel1, signers: signersForSel1, threshold: threshold},
		{selector: sel2, signers: signersForSel2, threshold: threshold},
	})
	addVerifierToDatastore(t, ds, sel3, targetCommittee, addr3)

	// Deploy other committee verifiers on chain1 and chain2 (not chain3)
	// These should NOT appear in the generated config for targetCommittee
	otherAddr1 := deployAndConfigureVerifier(t, env.OperationsBundle, chain1, otherCommittee, []sourceChainConfig{
		{selector: sel2, signers: otherSigners, threshold: 1},
	})
	addVerifierToDatastore(t, ds, sel1, otherCommittee, otherAddr1)

	otherAddr2 := deployAndConfigureVerifier(t, env.OperationsBundle, chain2, otherCommittee, []sourceChainConfig{
		{selector: sel1, signers: otherSigners, threshold: 1},
	})
	addVerifierToDatastore(t, ds, sel2, otherCommittee, otherAddr2)

	env.DataStore = ds.Seal()

	cs := changesets.GenerateAggregatorConfig()
	output, err := cs.Apply(env, changesets.GenerateAggregatorConfigCfg{
		ServiceIdentifier:  "test-aggregator",
		CommitteeQualifier: targetCommittee,
		ChainSelectors:     selectors,
	})
	require.NoError(t, err)

	require.NotNil(t, output.DataStore)
	cfg, err := ccv.GetAggregatorConfig(output.DataStore.Seal(), "test-aggregator")
	require.NoError(t, err)

	sel1Str := strconv.FormatUint(sel1, 10)
	sel2Str := strconv.FormatUint(sel2, 10)
	sel3Str := strconv.FormatUint(sel3, 10)

	// Build expected config using model.Committee
	// - DestinationVerifiers: verifier address on each destination chain
	// - QuorumConfigs: for each source chain, SourceVerifierAddress is verifier on that source chain
	expectedConfig := &offchain.Committee{
		DestinationVerifiers: map[offchain.DestinationSelector]string{
			sel1Str: addr1.Hex(),
			sel2Str: addr2.Hex(),
			sel3Str: addr3.Hex(),
		},
		QuorumConfigs: map[offchain.SourceSelector]*offchain.QuorumConfig{
			sel1Str: {
				SourceVerifierAddress: addr1.Hex(),
				Signers:               toModelSigners(signersForSel1),
				Threshold:             threshold,
			},
			sel2Str: {
				SourceVerifierAddress: addr2.Hex(),
				Signers:               toModelSigners(signersForSel2),
				Threshold:             threshold,
			},
			sel3Str: {
				SourceVerifierAddress: addr3.Hex(),
				Signers:               toModelSigners(signersForSel3),
				Threshold:             threshold,
			},
		},
	}

	// Compare full config - verifies targetCommittee is used, not otherCommittee
	assert.Equal(t, expectedConfig.DestinationVerifiers, cfg.DestinationVerifiers)
	assert.Equal(t, expectedConfig.QuorumConfigs, cfg.QuorumConfigs)
}

func toModelSigners(addrs []common.Address) []offchain.Signer {
	signers := make([]offchain.Signer, 0, len(addrs))
	for _, addr := range addrs {
		signers = append(signers, offchain.Signer{Address: addr.Hex()})
	}
	return signers
}

type sourceChainConfig struct {
	selector  uint64
	signers   []common.Address
	threshold uint8
}

func deployAndConfigureVerifier(
	t *testing.T,
	bundle operations.Bundle,
	chain cldfevm.Chain,
	qualifier string,
	sourceConfigs []sourceChainConfig,
) common.Address {
	t.Helper()

	rmnAddress := common.HexToAddress("0x0000000000000000000000000000000000000001")

	deployResult, err := operations.ExecuteOperation(
		bundle,
		committee_verifier.Deploy,
		chain,
		contract.DeployInput[committee_verifier.ConstructorArgs]{
			ChainSelector:  chain.Selector,
			TypeAndVersion: deployment.NewTypeAndVersion(committee_verifier.ContractType, *committee_verifier.Version),
			Args: committee_verifier.ConstructorArgs{
				DynamicConfig: committee_verifier.DynamicConfig{
					FeeAggregator:  chain.DeployerKey.From,
					AllowlistAdmin: chain.DeployerKey.From,
				},
				RMN: rmnAddress,
			},
			Qualifier: &qualifier,
		},
	)
	require.NoError(t, err)

	contractAddr := common.HexToAddress(deployResult.Output.Address)

	signatureUpdates := make([]committee_verifier.SignatureConfig, 0, len(sourceConfigs))
	for _, sc := range sourceConfigs {
		signatureUpdates = append(signatureUpdates, committee_verifier.SignatureConfig{
			SourceChainSelector: sc.selector,
			Threshold:           sc.threshold,
			Signers:             sc.signers,
		})
	}

	_, err = operations.ExecuteOperation(
		bundle,
		committee_verifier.ApplySignatureConfigs,
		chain,
		contract.FunctionInput[committee_verifier.SignatureConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       contractAddr,
			Args: committee_verifier.SignatureConfigArgs{
				SignatureConfigUpdates: signatureUpdates,
			},
		},
	)
	require.NoError(t, err)

	return contractAddr
}

func addVerifierToDatastore(t *testing.T, ds *datastore.MemoryDataStore, selector uint64, qualifier string, addr common.Address) {
	t.Helper()

	err := ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Qualifier:     qualifier,
		Type:          datastore.ContractType(committee_verifier.ContractType),
		Address:       addr.Hex(),
		Version:       committee_verifier.Version,
	})
	require.NoError(t, err)

	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: selector,
		Qualifier:     qualifier,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Address:       addr.Hex(),
		Version:       committee_verifier.Version,
	})
	require.NoError(t, err)
}

func generateTestSigners(t *testing.T, count int) []common.Address {
	t.Helper()
	signers := make([]common.Address, 0, count)
	for range count {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		signers = append(signers, crypto.PubkeyToAddress(key.PublicKey))
	}
	return signers
}

func TestGenerateAggregatorConfig_PreservesExistingIndexerConfig(t *testing.T) {
	selectors := []uint64{
		chainsel.TEST_90000001.Selector,
		chainsel.TEST_90000002.Selector,
	}
	committee := testCommittee

	env, evmChains := testutils.NewSimulatedEVMEnvironment(t, selectors)
	require.Len(t, evmChains, 2)

	chain1, chain2 := evmChains[0], evmChains[1]
	sel1, sel2 := selectors[0], selectors[1]

	signers := generateTestSigners(t, 2)
	threshold := uint8(1)

	ds := datastore.NewMemoryDataStore()

	addr1 := deployAndConfigureVerifier(t, env.OperationsBundle, chain1, committee, []sourceChainConfig{
		{selector: sel2, signers: signers, threshold: threshold},
	})
	addVerifierToDatastore(t, ds, sel1, committee, addr1)

	addr2 := deployAndConfigureVerifier(t, env.OperationsBundle, chain2, committee, []sourceChainConfig{
		{selector: sel1, signers: signers, threshold: threshold},
	})
	addVerifierToDatastore(t, ds, sel2, committee, addr2)

	existingIndexerConfig := &offchain.IndexerGeneratedConfig{
		Verifier: []offchain.IndexerGeneratedVerifierConfig{
			{
				Name:            "Existing Verifier",
				IssuerAddresses: []string{"0xissuer_1001", "0xissuer_1002"},
			},
		},
	}
	err := ccv.SaveIndexerConfig(ds, "existing-indexer", existingIndexerConfig)
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.GenerateAggregatorConfig()
	output, err := cs.Apply(env, changesets.GenerateAggregatorConfigCfg{
		ServiceIdentifier:  "new-aggregator",
		CommitteeQualifier: committee,
		ChainSelectors:     selectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	_, err = ccv.GetAggregatorConfig(outputSealed, "new-aggregator")
	require.NoError(t, err, "new aggregator config should be present")

	retrievedIndexerConfig, err := ccv.GetIndexerConfig(outputSealed, "existing-indexer")
	require.NoError(t, err, "existing indexer config should be preserved")
	assert.Equal(t, existingIndexerConfig, retrievedIndexerConfig, "indexer config should be unchanged")
}
