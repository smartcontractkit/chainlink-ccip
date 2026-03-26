package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	versioned_verifier_resolver_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
)

// TestDeployVerifierResolverViaCREATE2_DeploysAndAcceptsOwnership verifies that
// the sequence completes the full three-step flow (compute address → deploy via
// CREATE2 → accept ownership) without error and returns the expected address ref.
func TestDeployVerifierResolverViaCREATE2_DeploysAndAcceptsOwnership(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{create2TestChainSel}),
	)
	require.NoError(t, err)

	create2Addr := deployTestCreate2Factory(t, e)
	chain := e.BlockChains.EVMChains()[create2TestChainSel]

	input := sequences.DeployVerifierResolverViaCREATE2Input{
		CREATE2Factory: common.HexToAddress(create2Addr),
		ChainSelector:  create2TestChainSel,
		Qualifier:      "test-vvr",
		Type:           datastore.ContractType("VersionedVerifierResolver"),
		Version:        semver.MustParse("2.0.0"),
	}

	report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployVerifierResolverViaCREATE2, chain, input)
	require.NoError(t, err)

	require.Len(t, report.Output.Addresses, 1)
	addr := report.Output.Addresses[0]
	assert.Equal(t, create2TestChainSel, addr.ChainSelector)
	assert.Equal(t, "test-vvr", addr.Qualifier)
	assert.NotEmpty(t, addr.Address, "deployed address should be non-empty")

	// Two writes: CreateAndTransferOwnership + AcceptOwnership
	assert.Len(t, report.Output.Writes, 2)

	// Verify that AcceptOwnership completed: the deployer key must be the owner.
	deployedAddr := common.HexToAddress(addr.Address)
	resolver, err := versioned_verifier_resolver_bindings.NewVersionedVerifierResolver(deployedAddr, chain.Client)
	require.NoError(t, err)
	owner, err := resolver.Owner(nil)
	require.NoError(t, err)
	assert.Equal(t, chain.DeployerKey.From, owner, "owner should be the deployer key after AcceptOwnership")
}

// TestDeployVerifierResolverViaCREATE2_DifferentQualifiersProduceDifferentAddresses
// confirms that the CREATE2 salt (qualifier) drives address uniqueness, consistent
// with the lower-level DeployContractViaCREATE2 behaviour.
func TestDeployVerifierResolverViaCREATE2_DifferentQualifiersProduceDifferentAddresses(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{create2TestChainSel}),
	)
	require.NoError(t, err)

	create2Addr := deployTestCreate2Factory(t, e)
	chain := e.BlockChains.EVMChains()[create2TestChainSel]

	deployWith := func(qualifier string) string {
		input := sequences.DeployVerifierResolverViaCREATE2Input{
			CREATE2Factory: common.HexToAddress(create2Addr),
			ChainSelector:  create2TestChainSel,
			Qualifier:      qualifier,
			Type:           datastore.ContractType("VersionedVerifierResolver"),
			Version:        semver.MustParse("2.0.0"),
		}
		report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployVerifierResolverViaCREATE2, chain, input)
		require.NoError(t, err)
		require.Len(t, report.Output.Addresses, 1)
		return report.Output.Addresses[0].Address
	}

	addr1 := deployWith("resolver-alpha")
	addr2 := deployWith("resolver-beta")
	assert.NotEqual(t, addr1, addr2, "different qualifiers must produce different CREATE2 addresses")
}
