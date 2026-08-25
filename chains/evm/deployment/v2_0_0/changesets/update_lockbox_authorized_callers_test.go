package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	lbops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	cs_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	mcms_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

const lockboxTestChainSel = uint64(5009297550715157269)

// lockboxTestQualifier mirrors production, where every lockbox is deployed under a
// qualifier. The changeset no longer addresses by qualifier, but the fixture keeps one so
// the datastore refs look like the real thing.
const lockboxTestQualifier = "USDC"

// deployLockboxForTest deploys a lockbox, optionally under a qualifier. Pass nil for an
// unqualified ref.
func deployLockboxForTest(
	t *testing.T,
	e *cldf_deployment.Environment,
	chainSelector uint64,
	qualifier *string,
) (datastore.AddressRef, common.Address) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSelector]

	ref, err := contract.MaybeDeployContract(
		e.OperationsBundle, lbops.Deploy, chain,
		contract.DeployInput[lbops.ConstructorArgs]{
			TypeAndVersion: cldf_deployment.NewTypeAndVersion(lbops.ContractType, *lbops.Version),
			ChainSelector:  chainSelector,
			Args:           lbops.ConstructorArgs{Token: chain.DeployerKey.From},
			Qualifier:      qualifier,
		}, nil)
	require.NoError(t, err)

	return ref, common.HexToAddress(ref.Address)
}

func readOnChainAuthorizedCallers(
	t *testing.T,
	e *cldf_deployment.Environment,
	chainSelector uint64,
	lockboxAddr common.Address,
) []common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSelector]
	readBundle := testsetup.BundleWithFreshReporter(e.OperationsBundle)
	report, err := cldf_ops.ExecuteOperation(
		readBundle, lbops.GetAllAuthorizedCallers, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       lockboxAddr,
			Args:          struct{}{},
		})
	require.NoError(t, err)

	return report.Output
}

// setupLockboxWithTimelock deploys a lockbox, deploys the MCMS instance, seals both
// sets of refs into the DataStore, and transfers the lockbox to the timelock so that
// applyAuthorizedCallerUpdates writes land in MCMS proposals instead of executing directly.
func setupLockboxWithTimelock(
	t *testing.T,
	extraMCMSQualifiers ...string,
) (*cldf_deployment.Environment, common.Address, common.Address) {
	t.Helper()
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{lockboxTestChainSel}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[lockboxTestChainSel]
	b := e.OperationsBundle

	qualifier := lockboxTestQualifier
	ref, lockboxAddr := deployLockboxForTest(t, e, lockboxTestChainSel, &qualifier)

	timelockAddr, mcmsAddrs := deployMCMSInstanceForTest(
		t, b, chain, chain.DeployerKey.From, common_utils.CLLQualifier)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(ref))
	for _, addr := range mcmsAddrs {
		require.NoError(t, ds.Addresses().Add(addr))
	}
	// Additional MCMS instances make an unqualified timelock lookup ambiguous, which is
	// what production looks like - every chain carries CLLCCIP, RMNMCMS and UltraFastCurse.
	for _, q := range extraMCMSQualifiers {
		_, extraAddrs := deployMCMSInstanceForTest(t, b, chain, chain.DeployerKey.From, q)
		for _, addr := range extraAddrs {
			require.NoError(t, ds.Addresses().Add(addr))
		}
	}
	e.DataStore = ds.Seal()

	transferRampsToTimelockForTest(t, b, chain, e, []mcms_ops.OpTransferOwnershipInput{{
		ChainSelector:   lockboxTestChainSel,
		Address:         lockboxAddr,
		ProposedOwner:   timelockAddr,
		ContractType:    lbops.ContractType,
		TimelockAddress: timelockAddr,
	}})

	return e, timelockAddr, lockboxAddr
}

// mcmsConfig is validMCMSInput with the CLL qualifier, which is the qualifier
// setupLockboxWithTimelock deploys the MCMS instance under.
func mcmsConfig() mcms.Input {
	in := validMCMSInput()
	in.Qualifier = common_utils.CLLQualifier

	return in
}

// lockboxCfg builds a single-entry config at the lockbox's real version.
func lockboxCfg(
	lockboxAddr common.Address,
	appends, removes []common.Address,
) changesets.UpdateLockboxAuthorizedCallersCfg {
	return changesets.UpdateLockboxAuthorizedCallersCfg{
		Version: lbops.Version,
		Input: []changesets.LockboxCallerUpdate{{
			Selector: lockboxTestChainSel,
			Address:  lockboxAddr,
			Appends:  appends,
			Removes:  removes,
		}},
	}
}

func applyLockbox(
	t *testing.T,
	e *cldf_deployment.Environment,
	cfg changesets.UpdateLockboxAuthorizedCallersCfg,
) (cldf_deployment.ChangesetOutput, error) {
	t.Helper()

	return changesets.UpdateLockboxAuthorizedCallers(cs_changesets.GetRegistry()).Apply(*e,
		cs_changesets.WithMCMS[changesets.UpdateLockboxAuthorizedCallersCfg]{
			MCMS: mcmsConfig(),
			Cfg:  cfg,
		})
}

func verifyLockbox(
	t *testing.T,
	e *cldf_deployment.Environment,
	mcmsInput mcms.Input,
	cfg changesets.UpdateLockboxAuthorizedCallersCfg,
) error {
	t.Helper()

	return changesets.UpdateLockboxAuthorizedCallers(cs_changesets.GetRegistry()).VerifyPreconditions(*e,
		cs_changesets.WithMCMS[changesets.UpdateLockboxAuthorizedCallersCfg]{
			MCMS: mcmsInput,
			Cfg:  cfg,
		})
}

func TestUpdateLockboxAuthorizedCallers_AddAndRemove(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	chain := e.BlockChains.EVMChains()[lockboxTestChainSel]
	deployerKey := chain.DeployerKey.From
	otherAddr := common.HexToAddress("0x0000000000000000000000000000000000000001")

	out, err := applyLockbox(t, e, lockboxCfg(lockboxAddr,
		[]common.Address{deployerKey, otherAddr}, nil))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	callers := readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr)
	require.ElementsMatch(t, []common.Address{deployerKey, otherAddr}, callers)

	out, err = applyLockbox(t, e, lockboxCfg(lockboxAddr,
		nil, []common.Address{deployerKey, otherAddr}))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	require.Empty(t, readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))
}

// TestUpdateLockboxAuthorizedCallers_IdempotentAdd asserts the second apply of an
// already-applied add produces no proposal. The lockbox is timelock-owned, so dropping the
// filter would re-propose the write and fail the final assertion.
func TestUpdateLockboxAuthorizedCallers_IdempotentAdd(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	caller := common.HexToAddress("0x0000000000000000000000000000000000000001")

	out, err := applyLockbox(t, e, lockboxCfg(lockboxAddr, []common.Address{caller}, nil))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1, "adding an absent caller must produce a proposal")
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	require.Equal(t, []common.Address{caller},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))

	out, err = applyLockbox(t, e, lockboxCfg(lockboxAddr, []common.Address{caller}, nil))
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals, "no-op add must produce no proposal")

	require.Equal(t, []common.Address{caller},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))
}

// TestUpdateLockboxAuthorizedCallers_IdempotentRemove asserts removing an absent caller
// produces no proposal. On-chain the removal is a set no-op that does not revert, so only
// timelock ownership makes the absence of a proposal meaningful.
func TestUpdateLockboxAuthorizedCallers_IdempotentRemove(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	absentAddr := common.HexToAddress("0x0000000000000000000000000000000000000001")

	out, err := applyLockbox(t, e, lockboxCfg(lockboxAddr, nil, []common.Address{absentAddr}))
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals,
		"removing an absent caller must produce no proposal")

	require.Empty(t, readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))
}

// TestUpdateLockboxAuthorizedCallers_IdempotentMixed asserts a config whose adds are all
// present and whose removes are all absent produces no proposal.
func TestUpdateLockboxAuthorizedCallers_IdempotentMixed(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	present := common.HexToAddress("0x0000000000000000000000000000000000000001")
	absent := common.HexToAddress("0x0000000000000000000000000000000000000002")

	out, err := applyLockbox(t, e, lockboxCfg(lockboxAddr, []common.Address{present}, nil))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	out, err = applyLockbox(t, e, lockboxCfg(lockboxAddr,
		[]common.Address{present}, []common.Address{absent}))
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals, "fully filtered updates must produce no proposal")

	require.Equal(t, []common.Address{present},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))
}

// TestUpdateLockboxAuthorizedCallers_PartialFilter asserts that a config mixing an
// already-applied add with a new one proposes only the new one.
func TestUpdateLockboxAuthorizedCallers_PartialFilter(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	first := common.HexToAddress("0x0000000000000000000000000000000000000001")
	second := common.HexToAddress("0x0000000000000000000000000000000000000002")

	out, err := applyLockbox(t, e, lockboxCfg(lockboxAddr, []common.Address{first}, nil))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	out, err = applyLockbox(t, e, lockboxCfg(lockboxAddr,
		[]common.Address{first, second}, nil))
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1, "the new caller must still be proposed")
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	require.ElementsMatch(t, []common.Address{first, second},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))
}

// TestUpdateLockboxAuthorizedCallers_AddressNotInDataStore asserts an address absent from
// the datastore is rejected rather than written to blindly.
func TestUpdateLockboxAuthorizedCallers_AddressNotInDataStore(t *testing.T) {
	e, _, _ := setupLockboxWithTimelock(t)
	unknown := common.HexToAddress("0x00000000000000000000000000000000000000FF")

	cfg := lockboxCfg(unknown,
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil)

	err := verifyLockbox(t, e, mcmsConfig(), cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected exactly 1 datastore ref")

	_, err = applyLockbox(t, e, cfg)
	require.Error(t, err, "apply must reject it too, since it can run without verify")
}

// TestUpdateLockboxAuthorizedCallers_AddressIsNotALockbox asserts pointing at a real
// datastore ref of the wrong contract type is rejected. The timelock is a convenient
// non-lockbox ref that setupLockboxWithTimelock already seeds.
func TestUpdateLockboxAuthorizedCallers_AddressIsNotALockbox(t *testing.T) {
	e, timelockAddr, _ := setupLockboxWithTimelock(t)

	cfg := lockboxCfg(timelockAddr,
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil)

	err := verifyLockbox(t, e, mcmsConfig(), cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "not an ERC20LockBox")
}

// TestUpdateLockboxAuthorizedCallers_VersionMismatch asserts a config written for a
// different lockbox version is rejected, so a future 2.0.1 lockbox is not acted on by a
// config that meant 2.0.0.
func TestUpdateLockboxAuthorizedCallers_VersionMismatch(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	cfg := lockboxCfg(lockboxAddr,
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil)
	cfg.Version = semver.MustParse("2.0.1")

	err := verifyLockbox(t, e, mcmsConfig(), cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "is version 2.0.0, but config specifies 2.0.1")
}

func TestUpdateLockboxAuthorizedCallers_VerifyRequiresVersion(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	cfg := lockboxCfg(lockboxAddr,
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil)
	cfg.Version = nil

	err := verifyLockbox(t, e, mcmsConfig(), cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "version is required")
}

// TestUpdateLockboxAuthorizedCallers_VerifyRejectsDuplicateLockbox covers what the old
// map-keyed config made structurally impossible: the same lockbox listed twice.
func TestUpdateLockboxAuthorizedCallers_VerifyRejectsDuplicateLockbox(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	entry := changesets.LockboxCallerUpdate{
		Selector: lockboxTestChainSel,
		Address:  lockboxAddr,
		Appends:  []common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")},
	}

	err := verifyLockbox(t, e, mcmsConfig(), changesets.UpdateLockboxAuthorizedCallersCfg{
		Version: lbops.Version,
		Input:   []changesets.LockboxCallerUpdate{entry, entry},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "duplicate entry for lockbox")
}

// TestUpdateLockboxAuthorizedCallers_TwoLockboxesOneChain covers the layout that exists in
// CCV prod_testnet, where ethereum-sepolia holds two lockboxes. Two entries sharing a
// selector must be accepted, and their writes must be merged into a single batch operation
// so the chain's updates execute atomically.
func TestUpdateLockboxAuthorizedCallers_TwoLockboxesOneChain(t *testing.T) {
	e, timelockAddr, firstLockbox := setupLockboxWithTimelock(t)

	// A second lockbox on the same chain, under its own qualifier, added to the datastore
	// alongside the existing refs.
	secondQualifier := "LINK"
	secondRef, secondLockbox := deployLockboxForTest(t, e, lockboxTestChainSel, &secondQualifier)
	require.NotEqual(t, firstLockbox, secondLockbox)

	existing, err := e.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	ds := datastore.NewMemoryDataStore()
	for _, ref := range existing {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	require.NoError(t, ds.Addresses().Add(secondRef))
	e.DataStore = ds.Seal()

	// The second lockbox must also belong to the timelock, or its write executes inline
	// with the deployer key and never reaches the batch.
	transferRampsToTimelockForTest(t, e.OperationsBundle,
		e.BlockChains.EVMChains()[lockboxTestChainSel], e,
		[]mcms_ops.OpTransferOwnershipInput{{
			ChainSelector:   lockboxTestChainSel,
			Address:         secondLockbox,
			ProposedOwner:   timelockAddr,
			ContractType:    lbops.ContractType,
			TimelockAddress: timelockAddr,
		}})

	callerA := common.HexToAddress("0x0000000000000000000000000000000000000001")
	callerB := common.HexToAddress("0x0000000000000000000000000000000000000002")

	cfg := changesets.UpdateLockboxAuthorizedCallersCfg{
		Version: lbops.Version,
		Input: []changesets.LockboxCallerUpdate{
			{Selector: lockboxTestChainSel, Address: firstLockbox, Appends: []common.Address{callerA}},
			{Selector: lockboxTestChainSel, Address: secondLockbox, Appends: []common.Address{callerB}},
		},
	}

	require.NoError(t, verifyLockbox(t, e, mcmsConfig(), cfg), "two lockboxes on one chain must be allowed")

	out, err := applyLockbox(t, e, cfg)
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)

	// One batch operation for the chain, carrying both lockbox calls.
	ops := out.MCMSTimelockProposals[0].Operations
	require.Len(t, ops, 1, "both lockbox writes must merge into one batch op for the chain")
	require.Len(t, ops[0].Transactions, 2, "batch op must carry a transaction per lockbox")

	targets := []common.Address{
		common.HexToAddress(ops[0].Transactions[0].To),
		common.HexToAddress(ops[0].Transactions[1].To),
	}
	require.ElementsMatch(t, []common.Address{firstLockbox, secondLockbox}, targets)

	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
	require.Equal(t, []common.Address{callerA},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, firstLockbox))
	require.Equal(t, []common.Address{callerB},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, secondLockbox))
}

func TestUpdateLockboxAuthorizedCallers_VerifyRejectsAppendRemoveOverlap(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	addr := common.HexToAddress("0x0000000000000000000000000000000000000001")

	err := verifyLockbox(t, e, mcmsConfig(),
		lockboxCfg(lockboxAddr, []common.Address{addr}, []common.Address{addr}))
	require.Error(t, err)
	require.Contains(t, err.Error(), "is in both appends and removes")
}

func TestUpdateLockboxAuthorizedCallers_VerifyRejectsDuplicates(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)
	addr := common.HexToAddress("0x0000000000000000000000000000000000000001")

	err := verifyLockbox(t, e, mcmsConfig(),
		lockboxCfg(lockboxAddr, []common.Address{addr, addr}, nil))
	require.Error(t, err)
	require.Contains(t, err.Error(), "duplicate address")
}

func TestUpdateLockboxAuthorizedCallers_VerifyZeroLockboxAddress(t *testing.T) {
	e, _, _ := setupLockboxWithTimelock(t)

	err := verifyLockbox(t, e, mcmsConfig(), lockboxCfg(common.Address{},
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil))
	require.Error(t, err)
	require.Contains(t, err.Error(), "zero lockbox address")
}

func TestUpdateLockboxAuthorizedCallers_VerifyZeroCallerAddress(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	err := verifyLockbox(t, e, mcmsConfig(),
		lockboxCfg(lockboxAddr, []common.Address{{}}, nil))
	require.Error(t, err)
	require.Contains(t, err.Error(), "zero address in appends")
}

func TestUpdateLockboxAuthorizedCallers_VerifyNoInput(t *testing.T) {
	e, _, _ := setupLockboxWithTimelock(t)

	err := verifyLockbox(t, e, mcmsConfig(), changesets.UpdateLockboxAuthorizedCallersCfg{
		Version: lbops.Version,
		Input:   []changesets.LockboxCallerUpdate{},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "at least one entry is required")
}

func TestUpdateLockboxAuthorizedCallers_VerifyInvalidChainSelector(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	err := verifyLockbox(t, e, mcmsConfig(), changesets.UpdateLockboxAuthorizedCallersCfg{
		Version: lbops.Version,
		Input: []changesets.LockboxCallerUpdate{{
			Selector: 99999999,
			Address:  lockboxAddr,
			Appends:  []common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")},
		}},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid chain selector")
}

// TestUpdateLockboxAuthorizedCallers_VerifyMCMSDefaults asserts an empty MCMS block passes:
// the qualifier defaults to CLLCCIP and ValidUntil is populated before validation, so
// operators need not supply either.
func TestUpdateLockboxAuthorizedCallers_VerifyMCMSDefaults(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	err := verifyLockbox(t, e, mcms.Input{}, lockboxCfg(lockboxAddr,
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil))
	require.NoError(t, err)
}

// TestUpdateLockboxAuthorizedCallers_VerifyRejectsExpiredMCMS asserts a ValidUntil the
// operator did supply is still rejected when it is already expired.
func TestUpdateLockboxAuthorizedCallers_VerifyRejectsExpiredMCMS(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t)

	expired := mcmsConfig()
	expired.ValidUntil = 1

	err := verifyLockbox(t, e, expired, lockboxCfg(lockboxAddr,
		[]common.Address{common.HexToAddress("0x0000000000000000000000000000000000000001")}, nil))
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid MCMS input")
}

// TestUpdateLockboxAuthorizedCallers_OmittedQualifierResolvesCLLCCIP asserts an operator
// can leave the MCMS qualifier unset and still get a working proposal against the CLLCCIP
// timelock. The default comes from EVMMCMSReader.GetTimelockRef, not from this changeset;
// this pins the end-to-end behaviour operators rely on. A second MCMS instance is seeded
// under another qualifier so the datastore holds more than one timelock, as production does.
func TestUpdateLockboxAuthorizedCallers_OmittedQualifierResolvesCLLCCIP(t *testing.T) {
	e, _, lockboxAddr := setupLockboxWithTimelock(t, "RMNMCMS")
	caller := common.HexToAddress("0x0000000000000000000000000000000000000001")

	noQualifier := validMCMSInput()
	require.Empty(t, noQualifier.Qualifier, "fixture must leave the qualifier unset")

	out, err := changesets.UpdateLockboxAuthorizedCallers(cs_changesets.GetRegistry()).Apply(*e,
		cs_changesets.WithMCMS[changesets.UpdateLockboxAuthorizedCallersCfg]{
			MCMS: noQualifier,
			Cfg:  lockboxCfg(lockboxAddr, []common.Address{caller}, nil),
		})
	require.NoError(t, err)
	require.Len(t, out.MCMSTimelockProposals, 1)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	require.Equal(t, []common.Address{caller},
		readOnChainAuthorizedCallers(t, e, lockboxTestChainSel, lockboxAddr))
}
