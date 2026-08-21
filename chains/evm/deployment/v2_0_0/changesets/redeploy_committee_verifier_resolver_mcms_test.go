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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	contract_utils "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_lib "github.com/smartcontractkit/mcms"
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

// newRedeployResolverMCMSEnv builds the production layout: the ramps belong to
// the timelock, so their writes go into a proposal. It returns the timelock
// address.
func newRedeployResolverMCMSEnv(t *testing.T) (*deployment.Environment, redeployResolverFixture, common.Address) {
	t.Helper()
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

	return e, f, timelockAddr
}

// redeployResolverMCMSInput returns the config that builds a scheduled proposal.
func redeployResolverMCMSInput(f redeployResolverFixture) cs_changesets.WithMCMS[changesets.RedeployCommitteeVerifierResolverCfg] {
	input := redeployResolverInput(f)
	input.Cfg.DisableTransferOwnership = false
	input.MCMS = mcms.Input{
		Qualifier:      common_utils.CLLQualifier,
		TimelockAction: mcms_types.TimelockActionSchedule,
		ValidUntil:     redeployResolverValidUntil,
	}
	return input
}

// proposalTargets counts the transactions of a proposal per target address.
func proposalTargets(proposal mcms_lib.TimelockProposal) map[common.Address]int {
	targets := map[common.Address]int{}
	for _, op := range proposal.Operations {
		for _, tx := range op.Transactions {
			targets[common.HexToAddress(tx.To)]++
		}
	}
	return targets
}

// TestRedeployCommitteeVerifierResolver_Apply_WithMCMS exercises the production
// path: the ramps belong to the timelock, so their writes go into the proposal,
// and the new resolver ends up owned by the timelock.
func TestRedeployCommitteeVerifierResolver_Apply_WithMCMS(t *testing.T) {
	e, f, timelockAddr := newRedeployResolverMCMSEnv(t)
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]
	input := redeployResolverMCMSInput(f)

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
	targets := proposalTargets(out.MCMSTimelockProposals[0])
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

// addLateOffRampSourceForTest adds a source config to the OffRamp through the
// timelock. The config names the given resolver.
//
// The lane arrives after the first run of the changeset, so it makes the config
// of the first proposal stale.
func addLateOffRampSourceForTest(
	t *testing.T,
	e *deployment.Environment,
	chain evm.Chain,
	offRampAddr common.Address,
	resolver common.Address,
) {
	t.Helper()
	report, err := operations.ExecuteOperation(e.OperationsBundle, offramp_ops.ApplySourceChainConfigUpdates, chain,
		contract_utils.FunctionInput[[]offramp_ops.SourceChainConfigArgs]{
			ChainSelector: chain.Selector,
			Address:       offRampAddr,
			Args: []offramp_ops.SourceChainConfigArgs{{
				Router:              common.HexToAddress("0x04"),
				SourceChainSelector: redeployResolverLateSourceSel,
				IsEnabled:           true,
				OnRamps:             [][]byte{common.HexToAddress("0x07").Bytes()},
				DefaultCCVs:         []common.Address{resolver},
				LaneMandatedCCVs:    []common.Address{},
			}},
		})
	require.NoError(t, err)

	batch, err := contract_utils.NewBatchOperationFromWrites([]contract_utils.WriteOutput{report.Output})
	require.NoError(t, err)
	out, err := cs_changesets.NewOutputBuilder(*e, cs_changesets.GetRegistry()).
		WithBatchOps([]mcms_types.BatchOperation{batch}).
		Build(mcms.Input{
			Qualifier:      common_utils.CLLQualifier,
			TimelockAction: mcms_types.TimelockActionSchedule,
			ValidUntil:     redeployResolverValidUntil,
		})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
}

// acceptOwnershipSelector returns the 4-byte selector of acceptOwnership().
func acceptOwnershipSelector(t *testing.T) []byte {
	t.Helper()
	parsed, err := resolver_bindings.VersionedVerifierResolverMetaData.GetAbi()
	require.NoError(t, err)
	method, ok := parsed.Methods["acceptOwnership"]
	require.True(t, ok)
	return method.ID
}

// TestRedeployCommitteeVerifierResolver_Apply_WithMCMS_SkipRedeploy reproduces a
// dropped proposal.
//
// The first run deployed the resolver and built a proposal that held the ramp
// writes and the acceptOwnership of the resolver. The proposal was dropped, so
// the timelock never accepted the resolver and the ramps still name the old
// resolver. A lane then arrived, which made the config of the first proposal
// stale.
//
// The rerun must build a new proposal that holds acceptOwnership again, plus the
// ramp writes for the current lane set.
func TestRedeployCommitteeVerifierResolver_Apply_WithMCMS_SkipRedeploy(t *testing.T) {
	e, f, timelockAddr := newRedeployResolverMCMSEnv(t)
	chain := e.BlockChains.EVMChains()[redeployResolverChainSel]
	cs := changesets.RedeployCommitteeVerifierResolver(
		cs_changesets.GetRegistry(), deploy.GetTransferOwnershipRegistry(),
	)

	// The first run. Its proposal is dropped, so it is never processed.
	out, err := cs.Apply(*e, redeployResolverMCMSInput(f))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)

	// The pipeline keeps the address refs of the first run, because the contracts
	// exist on the chain.
	mergeRedeployResolverDataStore(t, e, out)
	newResolver := resolverRefAddress(t, e)
	require.NotEqual(t, f.oldResolver, newResolver)

	resolver, err := resolver_bindings.NewVersionedVerifierResolver(newResolver, chain.Client)
	require.NoError(t, err)
	owner, err := resolver.Owner(nil)
	require.NoError(t, err)
	require.Equal(t, chain.DeployerKey.From, owner, "the dropped proposal left the resolver with the deployer")

	// A new pipeline run has a new reporter.
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)

	// The lane that makes the first proposal stale.
	addLateOffRampSourceForTest(t, e, chain, f.offRamp, f.oldResolver)

	input := redeployResolverMCMSInput(f)
	input.Cfg.SkipRedeploy = true
	input.Cfg.PreviousResolvers = map[uint64]common.Address{redeployResolverChainSel: f.oldResolver}

	require.NoError(t, cs.VerifyPreconditions(*e, input))
	rerun, err := cs.Apply(*e, input)
	require.NoError(t, err)
	require.Len(t, rerun.MCMSTimelockProposals, 1)

	// The rerun deploys nothing.
	rerunRefs, err := rerun.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	assert.Empty(t, rerunRefs, "a rerun must not produce new address refs")

	// The new proposal holds acceptOwnership of the resolver and the ramp writes.
	targets := proposalTargets(rerun.MCMSTimelockProposals[0])
	assert.Positive(t, targets[newResolver], "resolver acceptOwnership must be in the new proposal")
	assert.Positive(t, targets[f.offRamp], "OffRamp write must be in the new proposal")
	assert.Positive(t, targets[f.canonicalOnRamp], "canonical OnRamp write must be in the new proposal")
	assert.Positive(t, targets[f.legacyOnRamp], "legacy OnRamp write must be in the new proposal")

	// Confirm the resolver transaction is acceptOwnership, and not some other call.
	var acceptCalls int
	for _, op := range rerun.MCMSTimelockProposals[0].Operations {
		for _, tx := range op.Transactions {
			if common.HexToAddress(tx.To) != newResolver {
				continue
			}
			if len(tx.Data) >= 4 && string(tx.Data[:4]) == string(acceptOwnershipSelector(t)) {
				acceptCalls++
			}
		}
	}
	assert.Equal(t, 1, acceptCalls, "the proposal must hold exactly one acceptOwnership of the resolver")

	testhelpers.ProcessTimelockProposals(t, *e, rerun.MCMSTimelockProposals, false)

	// The timelock now owns the resolver.
	owner, err = resolver.Owner(nil)
	require.NoError(t, err)
	assert.Equal(t, timelockAddr, owner, "the rerun completes the ownership transfer")

	// Every lane names the new resolver, the late lane included.
	offRamp, err := offramp_bindings.NewOffRamp(f.offRamp, chain.Client)
	require.NoError(t, err)
	for _, sourceSel := range []uint64{redeployResolverDestSel, redeployResolverInboundOnlySel, redeployResolverLateSourceSel} {
		sourceCfg, err := offRamp.GetSourceChainConfig(nil, sourceSel)
		require.NoError(t, err)
		assert.Equal(t, []common.Address{newResolver}, sourceCfg.DefaultCCVs,
			"source config %d must use the new resolver", sourceSel)
	}

	onRamp, err := onramp_bindings.NewOnRamp(f.canonicalOnRamp, chain.Client)
	require.NoError(t, err)
	destCfg, err := onRamp.GetDestChainConfig(nil, redeployResolverDestSel)
	require.NoError(t, err)
	assert.Equal(t, []common.Address{newResolver}, destCfg.DefaultCCVs)
}
