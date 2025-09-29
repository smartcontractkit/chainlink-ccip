package changesets_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts_VerifyPreconditions(t *testing.T) {
	lggr, err := logger.New()
	require.NoError(t, err, "Failed to create logger")

	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)

	chain, err := cldf_evm_provider.NewSimChainProvider(t, 5009297550715157269,
		cldf_evm_provider.SimChainProviderConfig{
			NumAdditionalAccounts: 1,
		},
	).Initialize(t.Context())
	require.NoError(t, err, "Failed to create SimChainProvider")

	chains := cldf_chain.NewBlockChainsFromSlice(
		[]cldf_chain.BlockChain{chain},
	)

	e := deployment.Environment{
		GetContext:       func() context.Context { return context.Background() },
		Logger:           lggr,
		OperationsBundle: bundle,
		BlockChains:      chains,
		DataStore:        datastore.NewMemoryDataStore().Seal(),
	}

	tests := []struct {
		desc        string
		input       changesets.DeployChainContractsCfg
		expectedErr string
	}{
		{
			desc: "valid input",
			input: changesets.DeployChainContractsCfg{
				ChainSel: 5009297550715157269,
				Params:   sequences.ContractParams{},
			},
		},
		{
			desc: "invalid chain selector",
			input: changesets.DeployChainContractsCfg{
				ChainSel: 12345,
				Params:   sequences.ContractParams{},
			},
			expectedErr: "no EVM chain with selector 12345 found in environment",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := changesets.DeployChainContracts.VerifyPreconditions(e, test.input)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr, "Expected error containing %q but got none", test.expectedErr)
			} else {
				require.NoError(t, err, "Did not expect error but got: %v", err)
			}
		})
	}
}

func TestDeployChainContracts_Apply(t *testing.T) {
	tests := []struct {
		desc          string
		makeDatastore func() *datastore.MemoryDataStore
	}{
		{
			desc: "empty datastore",
			makeDatastore: func() *datastore.MemoryDataStore {
				return datastore.NewMemoryDataStore()
			},
		},
		{
			desc: "non-empty datastore",
			makeDatastore: func() *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				_ = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(link.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x01").Hex(),
				})
				_ = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(weth.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x02").Hex(),
				})
				return ds
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")

			bundle := operations.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				operations.NewMemoryReporter(),
			)

			chain, err := cldf_evm_provider.NewSimChainProvider(t, 5009297550715157269,
				cldf_evm_provider.SimChainProviderConfig{
					NumAdditionalAccounts: 1,
				},
			).Initialize(t.Context())
			require.NoError(t, err, "Failed to create SimChainProvider")

			chains := cldf_chain.NewBlockChainsFromSlice(
				[]cldf_chain.BlockChain{chain},
			)

			ds := test.makeDatastore()
			existingAddrs, err := ds.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")

			e := deployment.Environment{
				GetContext:       func() context.Context { return context.Background() },
				Logger:           lggr,
				OperationsBundle: bundle,
				BlockChains:      chains,
				DataStore:        ds.Seal(),
			}

			usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10) // $15
			require.True(t, ok, "Failed to parse USDPerLINK")
			usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10) // $2000
			require.True(t, ok, "Failed to parse USDPerWETH")

			out, err := changesets.DeployChainContracts.Apply(e, changesets.DeployChainContractsCfg{
				ChainSel: 5009297550715157269,
				Params: sequences.ContractParams{
					RMNRemote:     sequences.RMNRemoteParams{},
					CCVAggregator: sequences.CCVAggregatorParams{},
					ExecutorOnRamp: sequences.ExecutorOnRampParams{
						MaxCCVsPerMsg: 10,
					},
					CommitteeVerifier: sequences.CommitteeVerifierParams{
						FeeAggregator: common.HexToAddress("0x01"),
						SignatureConfigArgs: committee_verifier.SetSignatureConfigArgs{
							Threshold: 1,
							Signers: []common.Address{
								common.HexToAddress("0x02"),
								common.HexToAddress("0x03"),
								common.HexToAddress("0x04"),
								common.HexToAddress("0x05"),
							},
						},
						StorageLocation: "https://test.chain.link.fake",
					},
					CCVProxy: sequences.CCVProxyParams{
						FeeAggregator: common.HexToAddress("0x01"),
					},
					FeeQuoter: sequences.FeeQuoterParams{
						MaxFeeJuelsPerMsg:              big.NewInt(0).Mul(big.NewInt(2e2), big.NewInt(1e18)),
						TokenPriceStalenessThreshold:   uint32(24 * 60 * 60),
						LINKPremiumMultiplierWeiPerEth: 9e17, // 0.9 ETH
						WETHPremiumMultiplierWeiPerEth: 1e18, // 1.0 ETH
						USDPerLINK:                     usdPerLink,
						USDPerWETH:                     usdPerWeth,
					},
				},
			})
			require.NoError(t, err, "Failed to apply DeployChainContracts changeset")

			newAddrs, err := out.DataStore.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")

			for _, addr := range existingAddrs {
				for _, newAddr := range newAddrs {
					if addr.Type == newAddr.Type {
						require.Equal(t, addr.Address, newAddr.Address, "Expected existing address for type %s to remain unchanged", addr.Type)
					}
				}
			}
		})
	}
}
