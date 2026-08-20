package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	committee_verifier_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	mock_receiver_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/mock_receiver"
	offramp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	onramp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	committee_verifier_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/committee_verifier"
	mock_receiver_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/mock_receiver_v2"
	offramp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/offramp"
	onramp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/onramp"
	resolver_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	contract_utils "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

const redeployResolverQualifier = "default"

var (
	redeployResolverChainSel = chainsel.TEST_90000001.Selector
	redeployResolverDestSel  = chainsel.TEST_90000002.Selector
	// redeployResolverInboundOnlySel is a source of an inbound lane that is not a
	// dest of any outbound lane. It proves that the OffRamp update does not depend
	// on the OnRamp dest set.
	redeployResolverInboundOnlySel = chainsel.TEST_90000003.Selector

	redeployResolverVersionTag = [4]byte{0xe9, 0xa0, 0x5a, 0x20}

	// redeployResolverExtraVerifier is a second required verifier of a
	// MockReceiver. The changeset must keep it.
	redeployResolverExtraVerifier = common.HexToAddress("0x0A")

	// redeployResolverFeeAggregator is the fee aggregator of the old resolver.
	// The replacement must keep it.
	redeployResolverFeeAggregator = common.HexToAddress("0x0B")
)

// redeployResolverFixture holds the contracts of the simulated chain.
type redeployResolverFixture struct {
	verifier          common.Address
	oldResolver       common.Address
	canonicalOnRamp   common.Address
	legacyOnRamp      common.Address
	offRamp           common.Address
	mockReceiver      common.Address
	unrelatedReceiver common.Address
	create2Factory    common.Address
}

// deployRedeployResolverFixture builds a chain in the broken state: a resolver
// that was deployed with plain CREATE, and ramps and a MockReceiver that point at
// it.
func deployRedeployResolverFixture(t *testing.T, e *deployment.Environment) redeployResolverFixture {
	t.Helper()
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]

	verifier := deployTestCommitteeVerifier(t, chain)

	// Plain CREATE, not the CREATE2 factory. This is the defect that the
	// changeset repairs.
	oldResolver, tx, resolver, err := resolver_bindings.DeployVersionedVerifierResolver(chain.DeployerKey, chain.Client)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// The old resolver is configured, so the lane works before the change.
	tx, err = resolver.ApplyOutboundImplementationUpdates(chain.DeployerKey, []resolver_bindings.VersionedVerifierResolverOutboundImplementationArgs{
		{DestChainSelector: redeployResolverDestSel, Verifier: verifier},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	tx, err = resolver.ApplyInboundImplementationUpdates(chain.DeployerKey, []resolver_bindings.VersionedVerifierResolverInboundImplementationArgs{
		{Version: redeployResolverVersionTag, Verifier: verifier},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// The resolver collects fee tokens as an OnRamp receipt issuer, so the
	// replacement must keep the fee aggregator. It has no constructor, so a new
	// resolver starts with a zero aggregator and withdrawFeeTokens reverts.
	tx, err = resolver.SetFeeAggregator(chain.DeployerKey, redeployResolverFeeAggregator)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// This receiver requires two verifiers. Only the resolver may change; the
	// second entry must survive the redeployment.
	mockReceiver, tx, _, err := mock_receiver_bindings.DeployMockReceiverV2(
		chain.DeployerKey, chain.Client,
		[]common.Address{oldResolver, redeployResolverExtraVerifier}, []common.Address{}, 0,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// This receiver never names the resolver, so it must be left alone.
	unrelatedReceiver, tx, _, err := mock_receiver_bindings.DeployMockReceiverV2(
		chain.DeployerKey, chain.Client,
		[]common.Address{redeployResolverExtraVerifier}, []common.Address{}, 0,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	return redeployResolverFixture{
		verifier:          verifier,
		oldResolver:       oldResolver,
		canonicalOnRamp:   deployTestOnRamp(t, chain, oldResolver),
		legacyOnRamp:      deployTestOnRamp(t, chain, oldResolver),
		offRamp:           deployTestOffRamp(t, chain, oldResolver),
		mockReceiver:      mockReceiver,
		unrelatedReceiver: unrelatedReceiver,
		create2Factory:    deployRedeployResolverCREATE2Factory(t, e),
	}
}

func deployRedeployResolverCREATE2Factory(t *testing.T, e *deployment.Environment) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]
	ref, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, chain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *create2_factory.Version),
		ChainSelector:  redeployResolverChainSel,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{chain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err)
	return common.HexToAddress(ref.Address)
}

func deployTestCommitteeVerifier(t *testing.T, chain evm.Chain) common.Address {
	t.Helper()
	addr, tx, _, err := committee_verifier_bindings.DeployCommitteeVerifier(
		chain.DeployerKey, chain.Client,
		committee_verifier_bindings.CommitteeVerifierDynamicConfig{
			FeeAggregator:  chain.DeployerKey.From,
			AllowlistAdmin: chain.DeployerKey.From,
		},
		[]string{"https://example.invalid"},
		common.HexToAddress("0x01"),
		redeployResolverVersionTag,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)
	return addr
}

func deployTestOnRamp(t *testing.T, chain evm.Chain, resolver common.Address) common.Address {
	t.Helper()
	addr, tx, contract, err := onramp_bindings.DeployOnRamp(
		chain.DeployerKey, chain.Client,
		onramp_bindings.OnRampStaticConfig{
			ChainSelector:         chain.Selector,
			RmnRemote:             common.HexToAddress("0x01"),
			MaxUSDCentsPerMessage: 1_000_000,
			TokenAdminRegistry:    common.HexToAddress("0x02"),
		},
		onramp_bindings.OnRampDynamicConfig{
			FeeQuoter:     common.HexToAddress("0x03"),
			FeeAggregator: chain.DeployerKey.From,
		},
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	tx, err = contract.ApplyDestChainConfigUpdates(chain.DeployerKey, []onramp_bindings.OnRampDestChainConfigArgs{{
		DestChainSelector:    redeployResolverDestSel,
		Router:               common.HexToAddress("0x04"),
		AddressBytesLength:   20,
		TokenReceiverAllowed: true,
		BaseExecutionGasCost: 100_000,
		DefaultCCVs:          []common.Address{resolver},
		LaneMandatedCCVs:     []common.Address{},
		DefaultExecutor:      common.HexToAddress("0x05"),
		OffRamp:              common.HexToAddress("0x06").Bytes(),
	}})
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	return addr
}

func deployTestOffRamp(t *testing.T, chain evm.Chain, resolver common.Address) common.Address {
	t.Helper()
	addr, tx, contract, err := offramp_bindings.DeployOffRamp(
		chain.DeployerKey, chain.Client,
		offramp_bindings.OffRampStaticConfig{
			LocalChainSelector:        chain.Selector,
			GasForCallExactCheck:      5_000,
			RmnRemote:                 common.HexToAddress("0x01"),
			TokenAdminRegistry:        common.HexToAddress("0x02"),
			MaxGasBufferToUpdateState: 100_000,
		},
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	sourceArgs := make([]offramp_bindings.OffRampSourceChainConfigArgs, 0, 2)
	for _, sourceSel := range []uint64{redeployResolverDestSel, redeployResolverInboundOnlySel} {
		sourceArgs = append(sourceArgs, offramp_bindings.OffRampSourceChainConfigArgs{
			Router:              common.HexToAddress("0x04"),
			SourceChainSelector: sourceSel,
			IsEnabled:           true,
			OnRamps:             [][]byte{common.HexToAddress("0x07").Bytes()},
			DefaultCCVs:         []common.Address{resolver},
			LaneMandatedCCVs:    []common.Address{},
		})
	}
	tx, err = contract.ApplySourceChainConfigUpdates(chain.DeployerKey, sourceArgs)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	return addr
}

// redeployResolverDataStore seeds the datastore with the refs from before the change.
func redeployResolverDataStore(t *testing.T, f redeployResolverFixture, withLegacyOnRamp bool) datastore.DataStore {
	t.Helper()
	ds := datastore.NewMemoryDataStore()
	add := func(contractType datastore.ContractType, version *semver.Version, qualifier string, address common.Address) {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: redeployResolverChainSel,
			Type:          contractType,
			Version:       version,
			Qualifier:     qualifier,
			Address:       address.Hex(),
		}))
	}
	add(datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType),
		versioned_verifier_resolver.Version, redeployResolverQualifier, f.oldResolver)
	add(datastore.ContractType(committee_verifier_ops.ContractType),
		committee_verifier_ops.Version, redeployResolverQualifier, f.verifier)
	add(datastore.ContractType(onramp_ops.ContractType), onramp_ops.Version, "", f.canonicalOnRamp)
	add(datastore.ContractType(offramp_ops.ContractType), offramp_ops.Version, "", f.offRamp)
	add(datastore.ContractType(mock_receiver_ops.ContractType), mock_receiver_ops.Version,
		redeployResolverQualifier, f.mockReceiver)
	add(datastore.ContractType(mock_receiver_ops.ContractType), mock_receiver_ops.Version,
		"unrelated", f.unrelatedReceiver)
	add(datastore.ContractType(create2_factory.ContractType), create2_factory.Version, "", f.create2Factory)
	if withLegacyOnRamp {
		add(datastore.ContractType(onramp_ops.ContractType), onramp_ops.Version, "legacy", f.legacyOnRamp)
	}
	return ds.Seal()
}

func redeployResolverInput(f redeployResolverFixture) cs_changesets.WithMCMS[changesets.RedeployCommitteeVerifierResolverCfg] {
	return cs_changesets.WithMCMS[changesets.RedeployCommitteeVerifierResolverCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.RedeployCommitteeVerifierResolverCfg{
			ChainSelectors:           []uint64{redeployResolverChainSel},
			CanonicalCREATE2Factory:  f.create2Factory,
			DisableTransferOwnership: true,
		},
	}
}

func newRedeployResolverEnv(t *testing.T, withLegacyOnRamp bool) (*deployment.Environment, redeployResolverFixture) {
	t.Helper()
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{redeployResolverChainSel}))
	require.NoError(t, err)
	f := deployRedeployResolverFixture(t, e)
	e.DataStore = redeployResolverDataStore(t, f, withLegacyOnRamp)
	return e, f
}

func TestRedeployCommitteeVerifierResolver_VerifyPreconditions(t *testing.T) {
	e, f := newRedeployResolverEnv(t, true)
	cs := changesets.RedeployCommitteeVerifierResolver(cs_changesets.GetRegistry(), nil)

	tests := []struct {
		desc        string
		mutate      func(cfg *changesets.RedeployCommitteeVerifierResolverCfg)
		expectedErr string
	}{
		{
			desc:   "valid config",
			mutate: func(*changesets.RedeployCommitteeVerifierResolverCfg) {},
		},
		{
			desc: "no chain selectors",
			mutate: func(cfg *changesets.RedeployCommitteeVerifierResolverCfg) {
				cfg.ChainSelectors = nil
			},
			expectedErr: "at least one chain must be configured",
		},
		{
			desc: "duplicate chain selector",
			mutate: func(cfg *changesets.RedeployCommitteeVerifierResolverCfg) {
				cfg.ChainSelectors = []uint64{redeployResolverChainSel, redeployResolverChainSel}
			},
			expectedErr: "appears more than once",
		},
		{
			desc: "canonical factory not set",
			mutate: func(cfg *changesets.RedeployCommitteeVerifierResolverCfg) {
				cfg.CanonicalCREATE2Factory = common.Address{}
			},
			expectedErr: "CanonicalCREATE2Factory is required",
		},
		{
			// This is the Polygon and Base case. The chain has a CREATE2 factory at
			// a non-canonical address, so every CREATE2 address on it differs.
			desc: "factory is not at the canonical address",
			mutate: func(cfg *changesets.RedeployCommitteeVerifierResolverCfg) {
				cfg.CanonicalCREATE2Factory = common.HexToAddress("0xdead")
			},
			expectedErr: "but the canonical factory is at",
		},
		{
			desc: "chain not in environment",
			mutate: func(cfg *changesets.RedeployCommitteeVerifierResolverCfg) {
				cfg.ChainSelectors = []uint64{redeployResolverInboundOnlySel}
			},
			expectedErr: "not found in environment",
		},
		{
			desc: "unknown committee qualifier",
			mutate: func(cfg *changesets.RedeployCommitteeVerifierResolverCfg) {
				cfg.CommitteeQualifier = "secondary"
			},
			expectedErr: `no CommitteeVerifierResolver with qualifier "secondary" found`,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			input := redeployResolverInput(f)
			test.mutate(&input.Cfg)
			err := cs.VerifyPreconditions(*e, input)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestRedeployCommitteeVerifierResolver_Apply(t *testing.T) {
	e, f := newRedeployResolverEnv(t, true)
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]

	out, err := changesets.RedeployCommitteeVerifierResolver(cs_changesets.GetRegistry(), nil).
		Apply(*e, redeployResolverInput(f))
	require.NoError(t, err)

	refs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	var resolverRefs, mockRefs []datastore.AddressRef
	for _, ref := range refs {
		switch ref.Type {
		case datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType):
			resolverRefs = append(resolverRefs, ref)
		case datastore.ContractType(mock_receiver_ops.ContractType):
			mockRefs = append(mockRefs, ref)
		}
	}

	// The resolver ref keeps its key and carries the new address. No legacy ref
	// exists, because a second resolver ref breaks the lane config helpers.
	require.Len(t, resolverRefs, 1)
	assert.Equal(t, redeployResolverQualifier, resolverRefs[0].Qualifier)
	newResolver := common.HexToAddress(resolverRefs[0].Address)
	assert.NotEqual(t, f.oldResolver, newResolver, "the resolver must move to a new address")

	// Only the receiver that names the old resolver is redeployed. Its ref keeps
	// the qualifier and carries a new address.
	require.Len(t, mockRefs, 1, "the unrelated MockReceiver must not be redeployed")
	assert.Equal(t, redeployResolverQualifier, mockRefs[0].Qualifier)
	newMockReceiver := common.HexToAddress(mockRefs[0].Address)
	assert.NotEqual(t, f.mockReceiver, newMockReceiver)

	// The new resolver resolves both directions.
	resolver, err := resolver_bindings.NewVersionedVerifierResolver(newResolver, chain.Client)
	require.NoError(t, err)
	outbound, err := resolver.GetAllOutboundImplementations(nil)
	require.NoError(t, err)
	require.Len(t, outbound, 1)
	assert.Equal(t, redeployResolverDestSel, outbound[0].DestChainSelector)
	assert.Equal(t, f.verifier, outbound[0].Verifier)

	inbound, err := resolver.GetAllInboundImplementations(nil)
	require.NoError(t, err)
	require.Len(t, inbound, 1)
	assert.Equal(t, redeployResolverVersionTag, inbound[0].Version)
	assert.Equal(t, f.verifier, inbound[0].Verifier)

	// The fee aggregator carries over, so withdrawFeeTokens keeps working.
	feeAggregator, err := resolver.GetFeeAggregator(nil)
	require.NoError(t, err)
	assert.Equal(t, redeployResolverFeeAggregator, feeAggregator)

	// The MockReceiver swaps the resolver and keeps its other required verifier.
	receiver, err := mock_receiver_bindings.NewMockReceiverV2(newMockReceiver, chain.Client)
	require.NoError(t, err)
	ccvs, err := receiver.GetCCVsAndFinalityConfig(nil, redeployResolverChainSel, []byte{})
	require.NoError(t, err)
	assert.Equal(t, []common.Address{newResolver, redeployResolverExtraVerifier}, ccvs.RequiredVerifier)

	// The unrelated receiver is untouched on chain.
	unrelated, err := mock_receiver_bindings.NewMockReceiverV2(f.unrelatedReceiver, chain.Client)
	require.NoError(t, err)
	unrelatedCCVs, err := unrelated.GetCCVsAndFinalityConfig(nil, redeployResolverChainSel, []byte{})
	require.NoError(t, err)
	assert.Equal(t, []common.Address{redeployResolverExtraVerifier}, unrelatedCCVs.RequiredVerifier)

	// Both OnRamps and the OffRamp use the new resolver. The deployer owns the
	// ramps in this test, so the writes ran directly instead of going into a
	// proposal.
	assertOnRampDefaultCCV(t, chain, f.canonicalOnRamp, newResolver)
	assertOnRampDefaultCCV(t, chain, f.legacyOnRamp, newResolver)

	// Every source config is updated, including the inbound-only lane that is not
	// a dest of any OnRamp lane.
	offRamp, err := offramp_bindings.NewOffRamp(f.offRamp, chain.Client)
	require.NoError(t, err)
	for _, sourceSel := range []uint64{redeployResolverDestSel, redeployResolverInboundOnlySel} {
		sourceCfg, err := offRamp.GetSourceChainConfig(nil, sourceSel)
		require.NoError(t, err)
		assert.Equal(t, []common.Address{newResolver}, sourceCfg.DefaultCCVs,
			"source config %d must use the new resolver", sourceSel)
	}
}

func assertOnRampDefaultCCV(t *testing.T, chain evm.Chain, onRampAddr, want common.Address) {
	t.Helper()
	onRamp, err := onramp_bindings.NewOnRamp(onRampAddr, chain.Client)
	require.NoError(t, err)
	cfg, err := onRamp.GetDestChainConfig(nil, redeployResolverDestSel)
	require.NoError(t, err)
	assert.Equal(t, []common.Address{want}, cfg.DefaultCCVs)
}

func TestRedeployCommitteeVerifierResolver_Apply_WithoutLegacyOnRamp(t *testing.T) {
	e, f := newRedeployResolverEnv(t, false)
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]

	_, err := changesets.RedeployCommitteeVerifierResolver(cs_changesets.GetRegistry(), nil).
		Apply(*e, redeployResolverInput(f))
	require.NoError(t, err)

	// The legacy OnRamp keeps the old resolver, because the datastore has no
	// legacy ref for it.
	assertOnRampDefaultCCV(t, chain, f.legacyOnRamp, f.oldResolver)
}

func TestRedeployCommitteeVerifierResolver_Apply_FailsWhenAlreadyCanonical(t *testing.T) {
	e, f := newRedeployResolverEnv(t, true)

	out, err := changesets.RedeployCommitteeVerifierResolver(cs_changesets.GetRegistry(), nil).
		Apply(*e, redeployResolverInput(f))
	require.NoError(t, err)

	// Merge the result back. The resolver now sits at its canonical address, so a
	// second run has nothing to redeploy.
	merged := datastore.NewMemoryDataStore()
	existing, err := e.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	for _, ref := range existing {
		require.NoError(t, merged.Addresses().Add(ref))
	}
	newRefs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	for _, ref := range newRefs {
		require.NoError(t, merged.Addresses().Upsert(ref))
	}
	e.DataStore = merged.Seal()

	_, err = changesets.RedeployCommitteeVerifierResolver(cs_changesets.GetRegistry(), nil).
		Apply(*e, redeployResolverInput(f))
	require.ErrorContains(t, err, "already at the canonical CREATE2 address")
}

func TestRedeployCommitteeVerifierResolver_Apply_FailsWithoutResolverRef(t *testing.T) {
	e, f := newRedeployResolverEnv(t, true)
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	_, err := changesets.RedeployCommitteeVerifierResolver(cs_changesets.GetRegistry(), nil).
		Apply(*e, redeployResolverInput(f))
	require.ErrorContains(t, err, `no CommitteeVerifierResolver with qualifier "default" found`)
}
