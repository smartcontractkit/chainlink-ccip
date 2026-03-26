package sequences_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	versioned_verifier_resolver_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

const create2TestChainSel = uint64(5009297550715157269)

// deployTestCreate2Factory deploys a CREATE2Factory and returns its address.
func deployTestCreate2Factory(t *testing.T, e *deployment.Environment) string {
	t.Helper()
	chain := e.BlockChains.EVMChains()[create2TestChainSel]
	ref, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle,
		create2_factory.Deploy,
		chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  create2TestChainSel,
			// Deployer must be in the allow list to call CreateAndTransferOwnership.
			Args: create2_factory.ConstructorArgs{AllowList: []common.Address{chain.DeployerKey.From}},
		},
		nil,
	)
	require.NoError(t, err, "failed to deploy CREATE2Factory")
	return ref.Address
}

// TestDeployContractViaCREATE2_DeploysAndReturnsAddress verifies that the sequence
// deploys the contract to a deterministic address and returns it in the output.
func TestDeployContractViaCREATE2_DeploysAndReturnsAddress(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{create2TestChainSel}),
	)
	require.NoError(t, err)

	create2Addr := deployTestCreate2Factory(t, e)
	chain := e.BlockChains.EVMChains()[create2TestChainSel]

	input := sequences.DeployContractViaCREATE2Input{
		CREATE2Factory:  common.HexToAddress(create2Addr),
		ChainSelector:   create2TestChainSel,
		Qualifier:       "test-qualifier",
		Type:            datastore.ContractType("VersionedVerifierResolver"),
		Version:         semver.MustParse("2.0.0"),
		ABI:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.ABI,
		BIN:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.Bin,
		ConstructorArgs: []any{},
	}

	report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployContractViaCREATE2, chain, input)
	require.NoError(t, err)

	require.Len(t, report.Output.Addresses, 1)
	addr := report.Output.Addresses[0]
	assert.Equal(t, create2TestChainSel, addr.ChainSelector)
	assert.Equal(t, "test-qualifier", addr.Qualifier)
	assert.NotEmpty(t, addr.Address, "deployed address should be non-empty")
}

// TestDeployContractViaCREATE2_DifferentQualifiersProduceDifferentAddresses verifies
// that the CREATE2 salt (qualifier) drives address uniqueness.
func TestDeployContractViaCREATE2_DifferentQualifiersProduceDifferentAddresses(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{create2TestChainSel}),
	)
	require.NoError(t, err)

	create2Addr := deployTestCreate2Factory(t, e)
	chain := e.BlockChains.EVMChains()[create2TestChainSel]

	deployWith := func(qualifier string) string {
		input := sequences.DeployContractViaCREATE2Input{
			CREATE2Factory:  common.HexToAddress(create2Addr),
			ChainSelector:   create2TestChainSel,
			Qualifier:       qualifier,
			Type:            datastore.ContractType("VersionedVerifierResolver"),
			Version:         semver.MustParse("2.0.0"),
			ABI:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.ABI,
			BIN:             versioned_verifier_resolver_latest.VersionedVerifierResolverMetaData.Bin,
			ConstructorArgs: []any{},
		}
		report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployContractViaCREATE2, chain, input)
		require.NoError(t, err)
		require.Len(t, report.Output.Addresses, 1)
		return report.Output.Addresses[0].Address
	}

	addr1 := deployWith("qualifier-alpha")
	addr2 := deployWith("qualifier-beta")
	assert.NotEqual(t, addr1, addr2, "different qualifiers must produce different CREATE2 addresses")
}
