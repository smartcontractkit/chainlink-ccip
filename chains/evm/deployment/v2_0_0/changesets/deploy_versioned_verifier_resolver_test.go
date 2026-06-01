package changesets_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	versioned_verifier_resolver_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

const deployVVRTestChainSel = uint64(5009297550715157269)

func deployTestCREATE2Factory(t *testing.T, e *deployment.Environment) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[deployVVRTestChainSel]
	ref, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, chain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *create2_factory.Version),
		ChainSelector:  deployVVRTestChainSel,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{chain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err)
	return common.HexToAddress(ref.Address)
}

func TestDeployVersionedVerifierResolver_VerifyPreconditions(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{deployVVRTestChainSel}))
	require.NoError(t, err)

	create2Addr := deployTestCREATE2Factory(t, e)
	mcmsRegistry := cs_changesets.GetRegistry()

	tests := []struct {
		desc        string
		input       cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]
		expectedErr string
	}{
		{
			desc: "valid committee resolver config",
			input: cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployVersionedVerifierResolverCfg{
					Chains: map[uint64]changesets.DeployVersionedVerifierResolverChainCfg{
						deployVVRTestChainSel: {
							ResolverType:   changesets.CommitteeVerifierResolver,
							CREATE2Factory: create2Addr,
						},
					},
				},
			},
		},
		{
			desc: "unsupported resolver type",
			input: cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployVersionedVerifierResolverCfg{
					Chains: map[uint64]changesets.DeployVersionedVerifierResolverChainCfg{
						deployVVRTestChainSel: {
							ResolverType:   changesets.VerifierResolverType("CCTPVerifierResolver"),
							CREATE2Factory: create2Addr,
						},
					},
				},
			},
			expectedErr: "unsupported verifier resolver type",
		},
		{
			desc: "missing CREATE2 factory",
			input: cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployVersionedVerifierResolverCfg{
					Chains: map[uint64]changesets.DeployVersionedVerifierResolverChainCfg{
						deployVVRTestChainSel: {
							ResolverType: changesets.LombardVerifierResolver,
						},
					},
				},
			},
			expectedErr: "CREATE2Factory is required",
		},
		{
			desc: "empty chains map",
			input: cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]{
				MCMS: mcms.Input{},
				Cfg:  changesets.DeployVersionedVerifierResolverCfg{Chains: map[uint64]changesets.DeployVersionedVerifierResolverChainCfg{}},
			},
			expectedErr: "at least one chain must be configured",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := changesets.DeployVersionedVerifierResolver(mcmsRegistry).VerifyPreconditions(*e, test.input)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeployVersionedVerifierResolver_Apply_DeploysAndIsIdempotent(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{deployVVRTestChainSel}))
	require.NoError(t, err)

	create2Addr := deployTestCREATE2Factory(t, e)
	mcmsRegistry := cs_changesets.GetRegistry()
	input := cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployVersionedVerifierResolverCfg{
			Chains: map[uint64]changesets.DeployVersionedVerifierResolverChainCfg{
				deployVVRTestChainSel: {
					ResolverType:   changesets.CommitteeVerifierResolver,
					CREATE2Factory: create2Addr,
				},
			},
		},
	}

	out1, err := changesets.DeployVersionedVerifierResolver(mcmsRegistry).Apply(*e, input)
	require.NoError(t, err)
	addrs1, err := out1.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, addrs1, 1)
	assert.Equal(t, string(changesets.CommitteeVerifierResolver), addrs1[0].Qualifier)
	assert.Equal(t, datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType), addrs1[0].Type)

	chain := e.BlockChains.EVMChains()[deployVVRTestChainSel]
	resolver, err := versioned_verifier_resolver_bindings.NewVersionedVerifierResolver(common.HexToAddress(addrs1[0].Address), chain.Client)
	require.NoError(t, err)
	owner, err := resolver.Owner(nil)
	require.NoError(t, err)
	assert.Equal(t, chain.DeployerKey.From, owner)

	dsSeed := datastore.NewMemoryDataStore()
	require.NoError(t, dsSeed.Addresses().Add(addrs1[0]))
	e.DataStore = dsSeed.Seal()

	out2, err := changesets.DeployVersionedVerifierResolver(mcmsRegistry).Apply(*e, input)
	require.NoError(t, err)
	addrs2, err := out2.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	require.Len(t, addrs2, 1)
	assert.Equal(t, addrs1[0].Address, addrs2[0].Address)
	assert.Empty(t, out2.Reports, "second apply should skip deployment and produce no new reports")
}

func TestDeployVersionedVerifierResolver_Apply_DifferentTypesProduceDifferentAddresses(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{deployVVRTestChainSel}))
	require.NoError(t, err)

	create2Addr := deployTestCREATE2Factory(t, e)
	mcmsRegistry := cs_changesets.GetRegistry()

	deployType := func(resolverType changesets.VerifierResolverType) string {
		out, err := changesets.DeployVersionedVerifierResolver(mcmsRegistry).Apply(*e, cs_changesets.WithMCMS[changesets.DeployVersionedVerifierResolverCfg]{
			MCMS: mcms.Input{},
			Cfg: changesets.DeployVersionedVerifierResolverCfg{
				Chains: map[uint64]changesets.DeployVersionedVerifierResolverChainCfg{
					deployVVRTestChainSel: {
						ResolverType:   resolverType,
						CREATE2Factory: create2Addr,
					},
				},
			},
		})
		require.NoError(t, err)
		addrs, err := out.DataStore.Addresses().Fetch()
		require.NoError(t, err)
		require.Len(t, addrs, 1)
		dsSeed := datastore.NewMemoryDataStore()
		require.NoError(t, dsSeed.Addresses().Add(addrs[0]))
		e.DataStore = dsSeed.Seal()
		return addrs[0].Address
	}

	committeeAddr := deployType(changesets.CommitteeVerifierResolver)
	lombardAddr := deployType(changesets.LombardVerifierResolver)
	assert.NotEqual(t, committeeAddr, lombardAddr)
}
