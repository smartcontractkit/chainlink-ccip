package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

const testChainSelector = 5009297550715157269

func TestWithdrawFeeTokens_VerifyPreconditions(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)

	tests := []struct {
		desc        string
		input       cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]
		expectedErr string
	}{
		{
			desc: "valid input with OnRamp ref",
			input: cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: testChainSelector,
					ContractRefs: []datastore.AddressRef{
						{
							Type:    datastore.ContractType(onramp.ContractType),
							Version: onramp.Version,
						},
					},
					FeeTokens: []common.Address{common.HexToAddress("0x01")},
				},
			},
			// Will fail because the datastore is empty, not because of validation
			expectedErr: "expected to find exactly 1 ref",
		},
		{
			desc: "invalid chain selector",
			input: cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: 12345,
					ContractRefs: []datastore.AddressRef{
						{
							Type:    datastore.ContractType(onramp.ContractType),
							Version: onramp.Version,
						},
					},
					FeeTokens: []common.Address{common.HexToAddress("0x01")},
				},
			},
			expectedErr: "no EVM chain with selector 12345 found in environment",
		},
		{
			desc: "unsupported contract type",
			input: cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: testChainSelector,
					ContractRefs: []datastore.AddressRef{
						{
							Type:    "UnsupportedContract",
							Version: semver.MustParse("1.0.0"),
						},
					},
					FeeTokens: []common.Address{common.HexToAddress("0x01")},
				},
			},
			expectedErr: "not a supported FeeTokenHandler",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mcmsRegistry := cs_core.GetRegistry()
			err := changesets.WithdrawFeeTokens(mcmsRegistry).VerifyPreconditions(*e, test.input)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWithdrawFeeTokens_Apply(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)

	// Deploy chain contracts first to get OnRamp, CommitteeVerifier, etc.
	mcmsRegistry := cs_core.GetRegistry()
	deployOut, err := changesets.DeployChainContracts(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			ChainSel: testChainSelector,
			Params:   testsetup.CreateBasicContractParams(),
		},
	})
	require.NoError(t, err, "Failed to deploy chain contracts")

	deployedAddrs, err := deployOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	// Find the WETH address to use as fee token
	var wethAddr common.Address
	for _, ref := range deployedAddrs {
		if ref.Type == "WETH9" {
			wethAddr = common.HexToAddress(ref.Address)
			break
		}
	}
	require.NotEqual(t, common.Address{}, wethAddr, "WETH should be deployed")

	// Copy deployed addresses into a new MemoryDataStore and seal it for the environment.
	ds := datastore.NewMemoryDataStore()
	for _, ref := range deployedAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	tests := []struct {
		desc  string
		input cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]
	}{
		{
			desc: "withdraw from OnRamp",
			input: cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: testChainSelector,
					ContractRefs: []datastore.AddressRef{
						{
							Type:    datastore.ContractType(onramp.ContractType),
							Version: onramp.Version,
						},
					},
					FeeTokens: []common.Address{wethAddr},
				},
			},
		},
		{
			desc: "withdraw from CommitteeVerifier",
			input: cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: testChainSelector,
					ContractRefs: []datastore.AddressRef{
						{
							Type:      datastore.ContractType(committee_verifier.ContractType),
							Version:   committee_verifier.Version,
							Qualifier: "alpha",
						},
					},
					FeeTokens: []common.Address{wethAddr},
				},
			},
		},
		{
			desc: "withdraw from multiple contracts",
			input: cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: testChainSelector,
					ContractRefs: []datastore.AddressRef{
						{
							Type:    datastore.ContractType(onramp.ContractType),
							Version: onramp.Version,
						},
						{
							Type:      datastore.ContractType(committee_verifier.ContractType),
							Version:   committee_verifier.Version,
							Qualifier: "alpha",
						},
					},
					FeeTokens: []common.Address{wethAddr},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			_, err := changesets.WithdrawFeeTokens(mcmsRegistry).Apply(*e, test.input)
			require.NoError(t, err)
		})
	}
}

func TestWithdrawFeeTokens_TokenPoolRequiresRecipient(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)

	// Manually add a TokenPool ref to the datastore
	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: testChainSelector,
		Type:          datastore.ContractType(token_pool.ContractType),
		Version:       token_pool.Version,
		Address:       common.HexToAddress("0xDEAD").Hex(),
	})
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()

	// Without recipient should fail
	_, err = changesets.WithdrawFeeTokens(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.WithdrawFeeTokensCfg{
			ChainSel: testChainSelector,
			ContractRefs: []datastore.AddressRef{
				{
					Type:    datastore.ContractType(token_pool.ContractType),
					Version: token_pool.Version,
				},
			},
			FeeTokens: []common.Address{common.HexToAddress("0x01")},
		},
	})
	require.ErrorContains(t, err, "recipient is required")
}
