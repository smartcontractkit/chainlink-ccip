package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func basicDeployCommitteeVerifierParams() sequences.CommitteeVerifierParams {
	return sequences.CommitteeVerifierParams{
		Version:         semver.MustParse("1.7.0"),
		FeeAggregator:   common.HexToAddress("0x02"),
		AllowlistAdmin:  common.HexToAddress("0x03"),
		StorageLocation: "https://test.chain.link.fake",
		SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
			Threshold: 1,
			Signers: []common.Address{
				common.HexToAddress("0x02"),
				common.HexToAddress("0x03"),
				common.HexToAddress("0x04"),
				common.HexToAddress("0x05"),
			},
		},
	}
}

func TestDeployCommitteeVerifier_VerifyPreconditions(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{5009297550715157269}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	tests := []struct {
		desc        string
		input       cs_core.WithMCMS[changesets.DeployCommitteeVerifierCfg]
		expectedErr string
	}{
		{
			desc: "valid input",
			input: cs_core.WithMCMS[changesets.DeployCommitteeVerifierCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployCommitteeVerifierCfg{
					ChainSel: 5009297550715157269,
					Params:   basicDeployCommitteeVerifierParams(),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mcmsRegistry := cs_core.NewMCMSReaderRegistry()
			err := changesets.DeployCommitteeVerifier(mcmsRegistry).VerifyPreconditions(*e, test.input)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr, "Expected error containing %q but got none", test.expectedErr)
			} else {
				require.NoError(t, err, "Did not expect error but got: %v", err)
			}
		})
	}
}

func TestDeployCommitteeVerifier_Apply_MultipleQualifiersOnSameChain(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{5009297550715157269}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	fqAddressRef := datastore.AddressRef{
		ChainSelector: 5009297550715157269,
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       semver.MustParse("1.7.0"),
		Address:       common.HexToAddress("0x01").Hex(),
	}

	// Ensure environment has an initial (empty) datastore
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(fqAddressRef))

	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.NewMCMSReaderRegistry()

	// 1) First run with qualifier "alpha"
	paramsAlpha := basicDeployCommitteeVerifierParams()
	paramsAlpha.Qualifier = "alpha"
	out1, err := changesets.DeployCommitteeVerifier(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployCommitteeVerifierCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployCommitteeVerifierCfg{
			ChainSel: 5009297550715157269,
			Params:   paramsAlpha,
		},
	})
	require.NoError(t, err, "First apply failed")
	addrs1, err := out1.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	// Helper to find address by type and qualifier in a slice
	find := func(refs []datastore.AddressRef, ct datastore.ContractType, q string) (datastore.AddressRef, bool) {
		for _, r := range refs {
			if r.Type == ct && r.Qualifier == q {
				return r, true
			}
		}
		return datastore.AddressRef{}, false
	}

	alphaCV, ok := find(addrs1, datastore.ContractType(committee_verifier.ContractType), "alpha")
	require.True(t, ok, "committee verifier (alpha) not found in first run output")
	alphaProxy, ok := find(addrs1, datastore.ContractType(committee_verifier.ProxyType), "alpha")
	require.True(t, ok, "committee verifier proxy (alpha) not found in first run output")

	// 2) Second run with qualifier "beta"; seed env with first run addresses so they are considered existing
	dsSeed := datastore.NewMemoryDataStore()
	for _, r := range addrs1 {
		require.NoError(t, dsSeed.Addresses().Add(r))
	}
	require.NoError(t, dsSeed.Addresses().Add(fqAddressRef))
	e.DataStore = dsSeed.Seal()

	paramsBeta := basicDeployCommitteeVerifierParams()
	paramsBeta.Qualifier = "beta"
	out2, err := changesets.DeployCommitteeVerifier(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployCommitteeVerifierCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployCommitteeVerifierCfg{
			ChainSel: 5009297550715157269,
			Params:   paramsBeta,
		},
	})
	require.NoError(t, err, "Second apply failed")
	addrs2, err := out2.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	betaCV, ok := find(addrs2, datastore.ContractType(committee_verifier.ContractType), "beta")
	require.True(t, ok, "committee verifier (beta) not found in second run output")
	betaProxy, ok := find(addrs2, datastore.ContractType(committee_verifier.ProxyType), "beta")
	require.True(t, ok, "committee verifier proxy (beta) not found in second run output")

	// Ensure addresses differ across qualifiers
	require.NotEqual(t, alphaCV.Address, betaCV.Address, "expected different CommitteeVerifier addresses for different qualifiers")
	require.NotEqual(t, alphaProxy.Address, betaProxy.Address, "expected different CommitteeVerifierProxy addresses for different qualifiers")

	// 3) Third run re-using qualifier "alpha" should be idempotent (returns existing alpha addresses)
	dsUnion := datastore.NewMemoryDataStore()
	for _, r := range addrs1 {
		require.NoError(t, dsUnion.Addresses().Add(r))
	}
	for _, r := range addrs2 {
		require.NoError(t, dsUnion.Addresses().Add(r))
	}
	require.NoError(t, dsUnion.Addresses().Add(fqAddressRef))
	e.DataStore = dsUnion.Seal()

	out3, err := changesets.DeployCommitteeVerifier(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployCommitteeVerifierCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployCommitteeVerifierCfg{
			ChainSel: 5009297550715157269,
			Params:   paramsAlpha, // same qualifier as first run
		},
	})
	require.NoError(t, err, "Third apply (repeat qualifier) failed")
	addrs3, err := out3.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	reAlphaCV, ok := find(addrs3, datastore.ContractType(committee_verifier.ContractType), "alpha")
	require.True(t, ok)
	reAlphaProxy, ok := find(addrs3, datastore.ContractType(committee_verifier.ProxyType), "alpha")
	require.True(t, ok)

	// Should return the same addresses as first run for the same qualifier
	require.Equal(t, alphaCV.Address, reAlphaCV.Address, "expected same CommitteeVerifier address when reusing qualifier")
	require.Equal(t, alphaProxy.Address, reAlphaProxy.Address, "expected same CommitteeVerifierProxy address when reusing qualifier")
}
