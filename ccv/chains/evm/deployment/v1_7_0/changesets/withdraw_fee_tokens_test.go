package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

const testChainSelector = 5009297550715157269

// TestWithdrawFeeTokens_VerifyPreconditions tests that the changeset rejects
// invalid configurations during precondition validation (before any on-chain work).
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
			expectedErr: "failed to resolve contract ref",
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

// TestWithdrawFeeTokens_Apply deploys a full set of chain contracts on a simulated
// chain, then verifies the changeset can successfully withdraw fee tokens from
// OnRamp, CommitteeVerifier, and multiple contracts at once.
func TestWithdrawFeeTokens_Apply(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)

	mcmsRegistry := cs_core.GetRegistry()

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy,
		e.BlockChains.EVMChains()[testChainSelector],
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
			ChainSelector:  testChainSelector,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{e.BlockChains.EVMChains()[testChainSelector].DeployerKey.From},
			},
		}, nil,
	)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")

	deployOut, err := changesets.DeployChainContracts(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsCfg{
			ChainSel:       testChainSelector,
			CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
			Params:         testsetup.CreateBasicContractParams(),
		},
	})
	require.NoError(t, err, "Failed to deploy chain contracts")

	deployedAddrs, err := deployOut.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	var wethAddr common.Address
	for _, ref := range deployedAddrs {
		if ref.Type == "WETH9" {
			wethAddr = common.HexToAddress(ref.Address)
			break
		}
	}
	require.NotEqual(t, common.Address{}, wethAddr, "WETH should be deployed")

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

// TestWithdrawFeeTokens_TokenPoolRequiresRecipient verifies that the sequence rejects
// TokenPool withdrawals when no recipient is specified. Covers both the generic
// TokenPool type and a concrete subtype (BurnMintTokenPool) to ensure all pool
// variants are routed through the same validation path.
func TestWithdrawFeeTokens_TokenPoolRequiresRecipient(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{testChainSelector}),
	)
	require.NoError(t, err)

	mcmsRegistry := cs_core.GetRegistry()

	tests := []struct {
		desc         string
		contractType datastore.ContractType
		version      *semver.Version
	}{
		{
			desc:         "generic TokenPool",
			contractType: datastore.ContractType(token_pool.ContractType),
			version:      token_pool.Version,
		},
		{
			desc:         "BurnMintTokenPool subtype",
			contractType: datastore.ContractType(burn_mint_token_pool.ContractType),
			version:      burn_mint_token_pool.Version,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ds := datastore.NewMemoryDataStore()
			err := ds.Addresses().Add(datastore.AddressRef{
				ChainSelector: testChainSelector,
				Type:          tt.contractType,
				Version:       tt.version,
				Address:       common.HexToAddress("0xDEAD").Hex(),
			})
			require.NoError(t, err)
			e.DataStore = ds.Seal()

			_, err = changesets.WithdrawFeeTokens(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.WithdrawFeeTokensCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.WithdrawFeeTokensCfg{
					ChainSel: testChainSelector,
					ContractRefs: []datastore.AddressRef{
						{
							Type:    tt.contractType,
							Version: tt.version,
						},
					},
					FeeTokens: []common.Address{common.HexToAddress("0x01")},
				},
			})
			require.ErrorContains(t, err, "recipient is required")
		})
	}
}
