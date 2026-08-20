package changesets_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	mcms_seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	offramp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	onramp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	offramp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/offramp"
	onramp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/onramp"
	resolver_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

const redeployResolverValidUntil = uint32(4128029039)

// transferRampsToTimelockForTest gives the ramps to the timelock, so their writes
// land in an MCMS proposal instead of executing with the deployer key.
func transferRampsToTimelockForTest(
	t *testing.T,
	b operations.Bundle,
	chain evm.Chain,
	e *deployment.Environment,
	targets []mcms_ops.OpTransferOwnershipInput,
) {
	t.Helper()
	batchOps, err := mcms_seq.TransferAndAcceptOwnership(b, chain, targets)
	require.NoError(t, err)
	if len(batchOps) == 0 {
		return
	}
	out, err := cs_changesets.NewOutputBuilder(*e, cs_changesets.GetRegistry()).
		WithBatchOps(batchOps).
		Build(mcms.Input{
			Qualifier:      common_utils.CLLQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     redeployResolverValidUntil,
		})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
}

// TestRedeployCommitteeVerifierResolver_Apply_WithMCMS exercises the production
// path: the ramps belong to the timelock, so their writes go into the proposal,
// and the new resolver ends up owned by the timelock.
func TestRedeployCommitteeVerifierResolver_Apply_WithMCMS(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{redeployResolverChainSel}))
	require.NoError(t, err)
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]
	b := e.OperationsBundle

	f := deployRedeployResolverFixture(t, e)
	timelockAddr, mcmsAddrs := deployMCMSInstanceForTest(t, b, chain, chain.DeployerKey.From, common_utils.CLLQualifier)

	ds := datastore.NewMemoryDataStore()
	existing, err := redeployResolverDataStore(t, f, true).Addresses().Fetch()
	require.NoError(t, err)
	for _, ref := range existing {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	for _, ref := range mcmsAddrs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	e.DataStore = ds.Seal()

	transferRampsToTimelockForTest(t, b, chain, e, []mcms_ops.OpTransferOwnershipInput{
		{ChainSelector: chain.Selector, Address: f.canonicalOnRamp, ProposedOwner: timelockAddr, ContractType: onramp_ops.ContractType, TimelockAddress: timelockAddr},
		{ChainSelector: chain.Selector, Address: f.legacyOnRamp, ProposedOwner: timelockAddr, ContractType: onramp_ops.ContractType, TimelockAddress: timelockAddr},
		{ChainSelector: chain.Selector, Address: f.offRamp, ProposedOwner: timelockAddr, ContractType: offramp_ops.ContractType, TimelockAddress: timelockAddr},
	})

	input := redeployResolverInput(f)
	input.Cfg.DisableTransferOwnership = false
	input.MCMS = mcms.Input{
		Qualifier:      common_utils.CLLQualifier,
		TimelockAction: mcms_types.TimelockActionSchedule,
		ValidUntil:     redeployResolverValidUntil,
	}

	out, err := changesets.RedeployCommitteeVerifierResolver(
		cs_changesets.GetRegistry(), deploy.GetTransferOwnershipRegistry(),
	).Apply(*e, input)
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1, "ramp writes and the ownership transfer share one proposal")

	refs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	var newResolver common.Address
	for _, ref := range refs {
		if ref.Type == datastore.ContractType(versioned_verifier_resolver.CommitteeVerifierResolverType) {
			newResolver = common.HexToAddress(ref.Address)
		}
	}
	require.NotEqual(t, common.Address{}, newResolver)

	// The proposal must carry the ramp writes and the resolver acceptOwnership.
	targets := map[common.Address]int{}
	for _, op := range out.MCMSTimelockProposals[0].Operations {
		for _, tx := range op.Transactions {
			targets[common.HexToAddress(tx.To)]++
		}
	}
	assert.Positive(t, targets[f.canonicalOnRamp], "canonical OnRamp write must be in the proposal")
	assert.Positive(t, targets[f.legacyOnRamp], "legacy OnRamp write must be in the proposal")
	assert.Positive(t, targets[f.offRamp], "OffRamp write must be in the proposal")
	assert.Positive(t, targets[newResolver], "resolver acceptOwnership must be in the proposal")

	// Ownership is a two-step transfer, so the deployer still owns it until the
	// timelock accepts.
	resolver, err := resolver_bindings.NewVersionedVerifierResolver(newResolver, chain.Client)
	require.NoError(t, err)
	owner, err := resolver.Owner(nil)
	require.NoError(t, err)
	assert.Equal(t, chain.DeployerKey.From, owner, "ownership only moves once the proposal executes")

	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	owner, err = resolver.Owner(nil)
	require.NoError(t, err)
	assert.Equal(t, timelockAddr, owner, "the timelock owns the new resolver")

	// The ramp writes took effect through the proposal.
	onRamp, err := onramp_bindings.NewOnRamp(f.canonicalOnRamp, chain.Client)
	require.NoError(t, err)
	destCfg, err := onRamp.GetDestChainConfig(nil, redeployResolverDestSel)
	require.NoError(t, err)
	assert.Equal(t, []common.Address{newResolver}, destCfg.DefaultCCVs)

	offRamp, err := offramp_bindings.NewOffRamp(f.offRamp, chain.Client)
	require.NoError(t, err)
	sourceCfg, err := offRamp.GetSourceChainConfig(nil, redeployResolverDestSel)
	require.NoError(t, err)
	assert.Equal(t, []common.Address{newResolver}, sourceCfg.DefaultCCVs)
}
