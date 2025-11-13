package create2_factory_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets/create2_factory"
	create2_factory_ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestDeployCREATE2Factory_Apply(t *testing.T) {
	tests := []struct {
		desc      string
		chainSel  uint64
		allowList []common.Address
	}{
		{
			desc:     "deploy with empty allow list",
			chainSel: 5009297550715157269,
		},
		{
			desc:     "deploy with single address in allow list",
			chainSel: 5009297550715157269,
			allowList: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
		},
		{
			desc:     "deploy with multiple addresses in allow list",
			chainSel: 5009297550715157269,
			allowList: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
				common.HexToAddress("0x0987654321098765432109876543210987654321"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{test.chainSel}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			// Apply the changeset
			out, err := create2_factory.DeployCREATE2Factory.Apply(*e, create2_factory.DeployCREATE2FactoryCfg{
				ChainSel:  test.chainSel,
				AllowList: test.allowList,
			})
			require.NoError(t, err, "Failed to apply DeployCREATE2Factory changeset")
			require.NotNil(t, out.DataStore, "DataStore should be returned")
			require.Len(t, out.Reports, 1, "Should have one report")

			// Verify the contract was deployed
			addrs, err := out.DataStore.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")
			require.Len(t, addrs, 1, "Should have deployed one contract")
			require.Equal(t, datastore.ContractType(create2_factory_ops.ContractType), addrs[0].Type, "Contract type should match")
			require.Equal(t, test.chainSel, addrs[0].ChainSelector, "Chain selector should match")
			require.NotEmpty(t, addrs[0].Address, "Address should not be empty")
		})
	}
}

func TestDeployCREATE2Factory_Idempotency(t *testing.T) {
	chainSel := uint64(5009297550715157269)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	allowList := []common.Address{
		e.BlockChains.EVMChains()[chainSel].DeployerKey.From,
	}

	// First deployment
	out1, err := create2_factory.DeployCREATE2Factory.Apply(*e, create2_factory.DeployCREATE2FactoryCfg{
		ChainSel:  chainSel,
		AllowList: allowList,
	})
	require.NoError(t, err, "Failed to apply DeployCREATE2Factory changeset (first time)")

	addrs1, err := out1.DataStore.Addresses().Fetch()
	require.NoError(t, err, "Failed to fetch addresses from datastore")
	require.Len(t, addrs1, 1, "Should have deployed one contract")

	// Update environment with the deployed contract
	e.DataStore = out1.DataStore.Seal()

	// Second deployment (should use existing)
	out2, err := create2_factory.DeployCREATE2Factory.Apply(*e, create2_factory.DeployCREATE2FactoryCfg{
		ChainSel:  chainSel,
		AllowList: allowList,
	})
	require.NoError(t, err, "Failed to apply DeployCREATE2Factory changeset (second time)")

	addrs2, err := out2.DataStore.Addresses().Fetch()
	require.NoError(t, err, "Failed to fetch addresses from datastore")
	require.Len(t, addrs2, 1, "Should still have one contract")

	// Verify addresses are the same (idempotency)
	require.Equal(t, addrs1[0].Address, addrs2[0].Address, "Contract address should be the same on repeated deployment")
}

func TestDeployCREATE2Factory_InvalidChain(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{5009297550715157269}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	// Try to deploy on a chain that doesn't exist
	_, err = create2_factory.DeployCREATE2Factory.Apply(*e, create2_factory.DeployCREATE2FactoryCfg{
		ChainSel:  99999999,
		AllowList: []common.Address{},
	})
	require.Error(t, err, "Should fail when deploying to non-existent chain")
	require.ErrorContains(t, err, "chain with selector 99999999 not found", "Error should mention chain not found")
}
