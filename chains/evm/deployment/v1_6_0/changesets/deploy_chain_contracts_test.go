package changesets_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts_VerifyPreconditions(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chain_selectors.ETHEREUM_MAINNET.Selector}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	tests := []struct {
		desc        string
		input       deployops.ContractDeploymentConfig
		expectedErr string
	}{
		{
			desc: "valid input",
			input: deployops.ContractDeploymentConfig{
				MCMS: mcms.Input{},
				Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
					chain_selectors.ETHEREUM_MAINNET.Selector: {
						Version: semver.MustParse("1.6.0"),
					},
				},
			},
		},
		{
			desc: "invalid version nil",
			input: deployops.ContractDeploymentConfig{
				MCMS: mcms.Input{},
				Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
					chain_selectors.ETHEREUM_MAINNET.Selector: {
						Version: nil,
					},
				},
			},
			expectedErr: "no version specified for chain with selector 5009297550715157269",
		},
		{
			desc: "invalid chain selector",
			input: deployops.ContractDeploymentConfig{
				MCMS: mcms.Input{},
				Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
					12345: {},
				},
			},
			expectedErr: "no selector 12345 found in environment: unknown chain selector 12345",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			dReg := deployops.GetRegistry()
			err := deployops.DeployContracts(dReg).VerifyPreconditions(*e, test.input)
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
					ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
					Type:          datastore.ContractType(link.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x01").Hex(),
				})
				_ = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
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
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chain_selectors.ETHEREUM_MAINNET.Selector}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			ds := test.makeDatastore()
			existingAddrs, err := ds.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")
			e.DataStore = ds.Seal() // Override datastore in environment to include existing addresses

			dReg := deployops.GetRegistry()
			version := semver.MustParse("1.6.0")
			out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
				MCMS: mcms.Input{},
				Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
					chain_selectors.ETHEREUM_MAINNET.Selector: {
						Version: version,
						// FEE QUOTER CONFIG
						MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
						TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
						LinkPremiumMultiplier:        9e17, // 0.9 ETH
						NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
						// OFFRAMP CONFIG
						PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
						GasForCallExactCheck:                    uint16(5000),
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
