package changesets_test

import (
	"testing"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_home"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	capabilities_registry "github.com/smartcontractkit/chainlink-evm/gethwrappers/keystone/generated/capabilities_registry_1_1_0"
	"github.com/stretchr/testify/require"
)

func TestDeployHomeChain_Apply(t *testing.T) {
	t.Parallel()
	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	out, err := changesets.DeployHomeChain.Apply(*e, sequences.DeployHomeChainConfig{
		HomeChainSel: chainSel,
		RMNStaticConfig: rmn_home.RMNHomeStaticConfig{
			Nodes:          []rmn_home.RMNHomeNode{},
			OffchainConfig: []byte("static config"),
		},
		RMNDynamicConfig: rmn_home.RMNHomeDynamicConfig{
			SourceChains:   []rmn_home.RMNHomeSourceChain{},
			OffchainConfig: []byte("dynamic config"),
		},
		NodeOperators: []capabilities_registry.CapabilitiesRegistryNodeOperator{
			{
				Admin: e.BlockChains.EVMChains()[chainSel].DeployerKey.From,
				Name:  "TestNodeOperator",
			},
		},
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")

	newAddrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err, "Failed to fetch addresses from datastore")
	require.Len(t, newAddrs, 3, "Expected 3 deployed contracts")
}
