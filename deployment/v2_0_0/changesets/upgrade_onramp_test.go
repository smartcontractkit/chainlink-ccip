package changesets_test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	changesetcore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	seq_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/mocks"
)

var (
	chainA = chainsel.TEST_90000001.Selector
	chainB = chainsel.TEST_90000002.Selector
)

const (
	oldOnRampAddr     = "0x1111111111111111111111111111111111111111"
	newOnRampAddr     = "0x2222222222222222222222222222222222222222"
	unknownOnRampAddr = "0x9999999999999999999999999999999999999999"
	timelockAddr      = "0x3333333333333333333333333333333333333333"
	mcmAddr           = "0x4444444444444444444444444444444444444444"
)

var (
	onRampType    = datastore.ContractType("OnRamp")
	onRampVersion = semver.MustParse("2.0.0")
)

// laneAdapter combines the ChainFamily, OffRampSourceOnRampSetter and
// OffRampSourceOnRampReader mocks so the composite satisfies the changesets' type
// assertions of a chain family adapter.
type laneAdapter struct {
	*mocks.MockChainFamily
	*mocks.MockOffRampSourceOnRampSetter
	*mocks.MockOffRampSourceOnRampReader
}

// laneMocks bundles the mocks behind the chain family adapter: the same adapter
// instance serves ChainFamily reads and OffRamp source onramp reads/writes.
type laneMocks struct {
	family *mocks.MockChainFamily
	setter *mocks.MockOffRampSourceOnRampSetter
	reader *mocks.MockOffRampSourceOnRampReader
}

func newLaneMocks(t *testing.T) (*adapters.ChainFamilyRegistry, laneMocks) {
	t.Helper()
	m := laneMocks{
		family: mocks.NewMockChainFamily(t),
		setter: mocks.NewMockOffRampSourceOnRampSetter(t),
		reader: mocks.NewMockOffRampSourceOnRampReader(t),
	}
	reg := adapters.NewChainFamilyRegistry()
	reg.RegisterChainFamily(chainsel.FamilyEVM, &laneAdapter{m.family, m.setter, m.reader})
	return reg, m
}

// noOpTransferSequence stands in for the on-chain ownership transfer sequence: it
// emits one canned batch op for the target chain without touching any chain.
var noOpTransferSequence = cldf_ops.NewSequence(
	"test-transfer-ownership",
	deploy.MCMSVersion,
	"emits a canned ownership-transfer batch op",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, in deploy.TransferOwnershipPerChainInput) (seq_core.OnChainOutput, error) {
		return seq_core.OnChainOutput{BatchOps: []mcms_types.BatchOperation{fakeBatchOp(in.ChainSelector, "0x0feee1")}}, nil
	},
)

func validMCMSInput() mcms.Input {
	return mcms.Input{
		TimelockAction: mcms_types.TimelockActionSchedule,
		ValidUntil:     uint32(time.Now().Add(48 * time.Hour).Unix()),
		Description:    "test proposal",
	}
}

func newTestEnv(t *testing.T) cldf.Environment {
	t.Helper()
	lggr := logger.Test(t)
	chains := map[uint64]cldf_chain.BlockChain{
		chainA: cldfevm.Chain{Selector: chainA},
		chainB: cldfevm.Chain{Selector: chainB},
	}
	return cldf.Environment{
		Name:        "test",
		Logger:      lggr,
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

// seedOnRamp adds an OnRamp ref to the environment datastore, simulating the
// canonical OnRamp as the changesets expect to find it.
func seedOnRamp(t *testing.T, e *cldf.Environment, addr string) {
	t.Helper()
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainA,
		Type:          onRampType,
		Version:       onRampVersion,
		Address:       addr,
	}))
	e.DataStore = ds.Seal()
}

func fakeTx(data string) mcms_types.Transaction {
	return mcms_types.Transaction{
		To:   newOnRampAddr,
		Data: []byte(data),
		// Valid EVM additional fields; empty json.RawMessage fails the
		// operations framework's serialization round-trip check.
		AdditionalFields: json.RawMessage(`{"value":0}`),
	}
}

func fakeBatchOp(chainSelector uint64, data string) mcms_types.BatchOperation {
	return mcms_types.BatchOperation{
		ChainSelector: mcms_types.ChainSelector(chainSelector),
		Transactions:  []mcms_types.Transaction{fakeTx(data)},
	}
}

func upgradeResult(destSelectors ...uint64) adapters.OnRampUpgradeResult {
	return adapters.OnRampUpgradeResult{
		NewOnRampRef: datastore.AddressRef{
			ChainSelector: chainA,
			Type:          onRampType,
			Version:       onRampVersion,
			Address:       newOnRampAddr,
		},
		LegacyOnRampRef: legacyOnRampRef(),
		// In the real EVM adapter, BatchOps only contain writes the deployer key
		// couldn't execute directly: the TestRouter wiring (applyRampUpdates) when
		// the TestRouter is timelock-owned. The dest chain config and dynamic
		// config writes on the fresh, deployer-owned OnRamp execute immediately.
		BatchOps: []mcms_types.BatchOperation{fakeBatchOp(chainA, "0xdeadbee1")},
	}
}

func legacyOnRampRef() datastore.AddressRef {
	return datastore.AddressRef{
		ChainSelector: chainA,
		Type:          onRampType,
		Version:       onRampVersion,
		Qualifier:     "legacy",
		Address:       oldOnRampAddr,
	}
}

// onRampAddressFromDS fakes adapters.ChainFamily.GetOnRampAddress: it reads the
// OnRamp ref from the given datastore and ABI-encodes it (32-byte left-padded).
func onRampAddressFromDS(ds datastore.DataStore, chainSelector uint64) ([]byte, error) {
	for _, ref := range ds.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector)) {
		if ref.Type == onRampType {
			return common.LeftPadBytes(common.HexToAddress(ref.Address).Bytes(), 32), nil
		}
	}
	return nil, fmt.Errorf("no OnRamp in datastore for chain %d", chainSelector)
}

func paddedAddrHex(addr string) string {
	return "0x" + hex.EncodeToString(common.LeftPadBytes(common.HexToAddress(addr).Bytes(), 32))
}

// mockMCMSReaderRegistry returns a fresh registry with a mock MCMS reader that
// serves canned timelock/MCM addresses for every chain.
func mockMCMSReaderRegistry(t *testing.T) *changesetcore.MCMSReaderRegistry {
	t.Helper()
	reader := mocks.NewMockMCMSReader(t)
	// Maybe: error paths and no-op runs never reach proposal building.
	reader.EXPECT().GetTimelockRef(mock.Anything, mock.Anything, mock.Anything).
		Return(datastore.AddressRef{Type: "Timelock", Version: semver.MustParse("1.0.0"), Address: timelockAddr}, nil).
		Maybe()
	reader.EXPECT().GetChainMetadata(mock.Anything, mock.Anything, mock.Anything).
		Return(mcms_types.ChainMetadata{StartingOpCount: 1, MCMAddress: mcmAddr}, nil).
		Maybe()
	reg := changesetcore.NewMCMSReaderRegistry()
	reg.RegisterMCMSReader(chainsel.FamilyEVM, reader)
	return reg
}

// mockTransferOwnershipRegistry returns a fresh registry with a mock ownership
// adapter registered for the EVM family.
func mockTransferOwnershipRegistry(t *testing.T) (*deploy.TransferOwnershipAdapterRegistry, *mocks.MockTransferOwnershipAdapter) {
	t.Helper()
	adapter := mocks.NewMockTransferOwnershipAdapter(t)
	reg := deploy.NewTransferOwnershipRegistry()
	reg.RegisterAdapter(chainsel.FamilyEVM, deploy.MCMSVersion, adapter)
	return reg, adapter
}

// mockUpgraderRegistry returns a registry with the given upgrader mock
// registered for the EVM family.
func mockUpgraderRegistry(upgrader *mocks.MockOnRampUpgrader) *adapters.OnRampUpgraderRegistry {
	reg := adapters.NewOnRampUpgraderRegistry()
	reg.Register(chainsel.FamilyEVM, upgrader)
	return reg
}

// expectNoUpgradeInProgress programs the pre-flight legacy-ref lookup: no legacy
// OnRamp exists yet.
func expectNoUpgradeInProgress(upgrader *mocks.MockOnRampUpgrader) {
	upgrader.EXPECT().LegacyOnRampRef(mock.Anything, chainA).
		Return(datastore.AddressRef{}, errors.New("not found")).Once()
}

// expectVerifyRequireUpgrade programs the pre-flight check that the canonical OnRamp is not yet upgraded to the new version.
func expectVerifyRequireUpgrade(upgrader *mocks.MockOnRampUpgrader, requireUpgrade bool) {
	var err error = nil
	if !requireUpgrade {
		err = fmt.Errorf("OnRamp is already upgraded: old=%s, new=%s", oldOnRampAddr, newOnRampAddr)
	}

	upgrader.EXPECT().VerifyOnrampRequireUpgrade(mock.Anything, chainA).
		Return(err).Once()
}

// expectPromotedToProdRouter programs the cleanup pre-flight promotion check.
func expectPromotedToProdRouter(upgrader *mocks.MockOnRampUpgrader) {
	upgrader.EXPECT().VerifyPromotedToProdRouter(mock.Anything, chainA, mock.Anything).
		Return(nil).Once()
}

// expectTimelockOwnership programs the Phase 3 pre-flight ownership check.
func expectTimelockOwnership(upgrader *mocks.MockOnRampUpgrader) {
	upgrader.EXPECT().VerifyNewOnRampOwner(mock.Anything, chainA, timelockAddr).
		Return(nil).Once()
}

// expectLegacyOnProdRouter programs the Phase 3 pre-flight prod Router check.
func expectLegacyOnProdRouter(upgrader *mocks.MockOnRampUpgrader) {
	upgrader.EXPECT().VerifyLegacyOnRampOnProdRouter(mock.Anything, chainA, mock.Anything).
		Return(nil).Once()
}

func TestUpgradeOnrampPhase1(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectVerifyRequireUpgrade(upgrader, true)
		expectNoUpgradeInProgress(upgrader)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		// chainB is TestRouter-class: Phase 1 must not deploy a test verifier for it.
		upgrader.EXPECT().DeployNewOnRamp(mock.Anything, chainA).
			Return(upgradeResult(chainB), nil).Once()

		familyReg, lane := newLaneMocks(t)
		// Called twice: old OnRamp from the env datastore, new OnRamp from the
		// temp datastore used to wire-encode the upgraded ref.
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS)
		// Pre-flight: the remote OffRamp whitelists exactly the old OnRamp.
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr)}, nil).Once()
		// The remote OffRamp must allow both the old and the new OnRamp.
		offRampUpdateOp := fakeBatchOp(chainB, "0x0ff9a")
		lane.setter.EXPECT().SetOffRampSourceOnRamps(mock.Anything, adapters.OffRampSetSourceOnRampsEntry{
			LocalChainSelector:  chainB,
			SourceChainSelector: chainA,
			OnRamps:             []string{paddedAddrHex(oldOnRampAddr), paddedAddrHex(newOnRampAddr)},
		}).Return(&offRampUpdateOp, false, nil).Once()

		ownershipReg, ownershipAdapter := mockTransferOwnershipRegistry(t)
		ownershipAdapter.EXPECT().InitializeTimelockAddress(mock.Anything, mock.Anything).Return(nil).Once()
		ownershipAdapter.EXPECT().SequenceTransferOwnershipViaMCMS().Return(noOpTransferSequence).Once()
		// The new OnRamp must be transferred to the timelock.
		ownershipAdapter.EXPECT().ShouldAcceptOwnershipWithTransferOwnership(mock.Anything, mock.MatchedBy(func(in deploy.TransferOwnershipPerChainInput) bool {
			return in.ChainSelector == chainA &&
				len(in.ContractRef) == 1 &&
				in.ContractRef[0].Address == newOnRampAddr &&
				in.ProposedOwner == timelockAddr
		})).Return(false, nil).Once()

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			ownershipReg,
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			MCMS:                 validMCMSInput(),
			DestSelectorsInScope: []uint64{chainB},
		})
		require.NoError(t, err)

		// Datastore output: canonical new OnRamp ref + legacy-qualified old OnRamp ref.
		refs := out.DataStore.Seal().Addresses().Filter(datastore.AddressRefByChainSelector(chainA))
		require.Len(t, refs, 2)
		byQualifier := map[string]datastore.AddressRef{}
		for _, ref := range refs {
			byQualifier[ref.Qualifier] = ref
		}
		assert.Equal(t, newOnRampAddr, byQualifier[""].Address, "new OnRamp should be canonical")
		assert.Equal(t, oldOnRampAddr, byQualifier["legacy"].Address, "old OnRamp should be legacy-qualified")

		// One proposal containing: the upgrader's batch op (chainA — the
		// TestRouter wiring tx in practice), the ownership transfer batch op
		// (chainA), and the OffRamp source update batch op (chainB).
		require.Len(t, out.MCMSTimelockProposals, 1)
		prop := out.MCMSTimelockProposals[0]
		require.Len(t, prop.Operations, 3)
		txCounts := map[mcms_types.ChainSelector]int{}
		for _, op := range prop.Operations {
			txCounts[op.ChainSelector] += len(op.Transactions)
		}
		assert.Equal(t, 2, txCounts[mcms_types.ChainSelector(chainA)], "chainA should have TestRouter wiring + ownership transfer txs")
		assert.Equal(t, 1, txCounts[mcms_types.ChainSelector(chainB)], "chainB should have the OffRamp update tx")
		assert.Equal(t, timelockAddr, prop.TimelockAddresses[mcms_types.ChainSelector(chainA)])
		assert.Equal(t, timelockAddr, prop.TimelockAddresses[mcms_types.ChainSelector(chainB)])
	})

	t.Run("SuccessWithProdRouterDests", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectVerifyRequireUpgrade(upgrader, true)
		expectNoUpgradeInProgress(upgrader)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		// chainB is ProdRouter-class: Phase 1 must validate test verifier infra.
		upgrader.EXPECT().DeployNewOnRamp(mock.Anything, chainA).
			Return(upgradeResult(chainB), nil).Once()

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS)
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr)}, nil).Once()
		offRampUpdateOp := fakeBatchOp(chainB, "0x0ff9a")
		lane.setter.EXPECT().SetOffRampSourceOnRamps(mock.Anything, adapters.OffRampSetSourceOnRampsEntry{
			LocalChainSelector:  chainB,
			SourceChainSelector: chainA,
			OnRamps:             []string{paddedAddrHex(oldOnRampAddr), paddedAddrHex(newOnRampAddr)},
		}).Return(&offRampUpdateOp, false, nil).Once()

		ownershipReg, ownershipAdapter := mockTransferOwnershipRegistry(t)
		ownershipAdapter.EXPECT().InitializeTimelockAddress(mock.Anything, mock.Anything).Return(nil).Once()
		ownershipAdapter.EXPECT().SequenceTransferOwnershipViaMCMS().Return(noOpTransferSequence).Once()
		ownershipAdapter.EXPECT().ShouldAcceptOwnershipWithTransferOwnership(mock.Anything, mock.MatchedBy(func(in deploy.TransferOwnershipPerChainInput) bool {
			return in.ChainSelector == chainA &&
				len(in.ContractRef) == 1 &&
				in.ContractRef[0].Address == newOnRampAddr &&
				in.ProposedOwner == timelockAddr
		})).Return(false, nil).Once()

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			ownershipReg,
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			MCMS:                 validMCMSInput(),
			DestSelectorsInScope: []uint64{chainB},
		})
		require.NoError(t, err)

		refs := out.DataStore.Seal().Addresses().Filter(datastore.AddressRefByChainSelector(chainA))
		require.Len(t, refs, 2)
		byQualifier := map[string]datastore.AddressRef{}
		for _, ref := range refs {
			byQualifier[ref.Qualifier] = ref
		}
		assert.Equal(t, newOnRampAddr, byQualifier[""].Address, "new OnRamp should be canonical")
		assert.Equal(t, oldOnRampAddr, byQualifier["legacy"].Address, "old OnRamp should be legacy-qualified")

		require.Len(t, out.MCMSTimelockProposals, 1)
		prop := out.MCMSTimelockProposals[0]
		require.Len(t, prop.Operations, 3)
	})

	t.Run("InvalidMCMSInput", func(t *testing.T) {
		e := newTestEnv(t)
		upgrader := mocks.NewMockOnRampUpgrader(t) // no calls expected
		familyReg, _ := newLaneMocks(t)            // no calls expected
		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector: chainA,
			// PopulateDefaults fills in a valid TimelockAction/ValidUntil for a zero
			// Input, so an explicitly invalid TimelockAction is needed to fail Validate.
			MCMS:                 mcms.Input{TimelockAction: "bogus"},
			DestSelectorsInScope: []uint64{chainB},
		})
		require.ErrorContains(t, err, "invalid timelock action")
	})

	t.Run("NoUpgraderRegistered", func(t *testing.T) {
		e := newTestEnv(t)
		familyReg, _ := newLaneMocks(t)
		cs := changesets.UpgradeOnrampPhase1(
			adapters.NewOnRampUpgraderRegistry(),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "no OnRampUpgrader registered")
	})

	t.Run("UpgradeError", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectNoUpgradeInProgress(upgrader)
		expectVerifyRequireUpgrade(upgrader, true)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		upgrader.EXPECT().DeployNewOnRamp(mock.Anything, chainA).
			Return(adapters.OnRampUpgradeResult{}, errors.New("boom")).Once()

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Once()
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr)}, nil).Once()

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "upgrade OnRamp")
		require.ErrorContains(t, err, "boom")
	})

	t.Run("TransferOwnershipError", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectNoUpgradeInProgress(upgrader)
		expectVerifyRequireUpgrade(upgrader, true)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		upgrader.EXPECT().DeployNewOnRamp(mock.Anything, chainA).
			Return(upgradeResult(chainB), nil).Once()

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS)
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr)}, nil).Once()

		ownershipReg, ownershipAdapter := mockTransferOwnershipRegistry(t)
		ownershipAdapter.EXPECT().InitializeTimelockAddress(mock.Anything, mock.Anything).
			Return(errors.New("boom")).Once()

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			ownershipReg,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "transfer ownership")
		require.ErrorContains(t, err, "boom")
	})

	t.Run("UnexpectedOnRampOnRemoteOffRamp", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectNoUpgradeInProgress(upgrader)
		expectVerifyRequireUpgrade(upgrader, true)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		// DeployNewOnRamp must NOT be called: the pre-flight check fails first.

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Once()
		// The remote OffRamp whitelists an onramp the changeset doesn't expect;
		// setting [old, new] would silently drop it.
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr), paddedAddrHex(unknownOnRampAddr)}, nil).Once()

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "pre-flight OffRamp whitelist check")
		require.ErrorContains(t, err, paddedAddrHex(unknownOnRampAddr))
	})

	t.Run("OldOnRampMissingOnRemoteOffRamp", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectNoUpgradeInProgress(upgrader)
		expectVerifyRequireUpgrade(upgrader, true)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Once()
		// The remote OffRamp was never configured for this source chain.
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{}, nil).Once()

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "does not whitelist expected onramp")
	})

	t.Run("UpgradeAlreadyInProgress", func(t *testing.T) {
		e := newTestEnv(t)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		// A legacy-qualified OnRamp exists: a previous upgrade is still in flight.
		upgrader.EXPECT().LegacyOnRampRef(mock.Anything, chainA).
			Return(legacyOnRampRef(), nil).Once()
		// Nothing else must be called: the guard fails before any read or mutation.

		familyReg, _ := newLaneMocks(t)
		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "upgrade is already in progress")
	})

	t.Run("NoUpgradeRequired", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, oldOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectNoUpgradeInProgress(upgrader)
		expectVerifyRequireUpgrade(upgrader, false)

		familyReg, _ := newLaneMocks(t)

		cs := changesets.UpgradeOnrampPhase1(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
			nil,
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "OnRamp is already upgraded")
	})
}

func TestUpgradeOnrampPhase2(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := newTestEnv(t)

		stageOp := mcms_types.BatchOperation{ // onramp:apply-dest-chain-config-updates (OnRamp -> TestRouter) + router:apply-ramp-updates (TestRouter -> OnRamp)
			ChainSelector: mcms_types.ChainSelector(chainA),
			Transactions: []mcms_types.Transaction{
				fakeTx("0xdeadbee3"),
				fakeTx("0xdeadbee4"),
			},
		}
		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		expectTimelockOwnership(upgrader)
		upgrader.EXPECT().PromoteOnrampToTestRouter(mock.Anything, chainA, []uint64{chainB}).
			Return([]mcms_types.BatchOperation{stageOp}, nil).Once()

		cs := changesets.UpgradeOnrampPhase2(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.NoError(t, err)

		require.Len(t, out.MCMSTimelockProposals, 1)
		prop := out.MCMSTimelockProposals[0]
		require.Len(t, prop.Operations, 1)
		assert.Equal(t, timelockAddr, prop.TimelockAddresses[mcms_types.ChainSelector(chainA)])
	})

	t.Run("NoOpWhenNoProdRouterDests", func(t *testing.T) {
		e := newTestEnv(t)

		// A chain whose lanes are all TestRouter-class has nothing to stage.
		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{TestRouterDests: []uint64{chainB}}, nil).Once()
		// SetDefaultCCVsOnNewOnRamp and PromoteOnrampToTestRouter must NOT be called.

		cs := changesets.UpgradeOnrampPhase2(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.NoError(t, err)
		assert.Empty(t, out.MCMSTimelockProposals)
	})

	t.Run("StageError", func(t *testing.T) {
		e := newTestEnv(t)
		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		expectTimelockOwnership(upgrader)
		upgrader.EXPECT().PromoteOnrampToTestRouter(mock.Anything, chainA, []uint64{chainB}).
			Return(nil, errors.New("boom")).Once()

		cs := changesets.UpgradeOnrampPhase2(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "stage on TestRouter")
		require.ErrorContains(t, err, "boom")
	})

	t.Run("OwnershipNotTransferred", func(t *testing.T) {
		e := newTestEnv(t)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		// The new OnRamp is still deployer-owned: Phase 1's ownership proposal
		// hasn't executed, so staging must not proceed.
		upgrader.EXPECT().VerifyNewOnRampOwner(mock.Anything, chainA, timelockAddr).
			Return(errors.New("boom")).Once()
		// SetDefaultCCVsOnNewOnRamp and PromoteOnrampToTestRouter must NOT be called.

		cs := changesets.UpgradeOnrampPhase2(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "pre-flight ownership check")
		require.ErrorContains(t, err, "boom")
	})
}

func TestUpgradeOnrampPhase3(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := newTestEnv(t)

		promoteOp := mcms_types.BatchOperation{
			ChainSelector: mcms_types.ChainSelector(chainA),
			Transactions: []mcms_types.Transaction{
				fakeTx("0xdeadbee6"), // onramp:apply-dest-chain-config-updates (OnRamp -> prod Router)
				fakeTx("0xdeadbee7"), // router:apply-ramp-updates (prod Router -> OnRamp)
			},
		}
		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectTimelockOwnership(upgrader)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		upgrader.EXPECT().VerifyLegacyOnRampOnProdRouter(mock.Anything, chainA, []uint64{chainB}).
			Return(nil).Once()
		upgrader.EXPECT().PromoteOnrampToProdRouter(mock.Anything, chainA, []uint64{chainB}).
			Return([]mcms_types.BatchOperation{promoteOp}, nil).Once()

		cs := changesets.UpgradeOnrampPhase3(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.NoError(t, err)

		require.Len(t, out.MCMSTimelockProposals, 1)
		prop := out.MCMSTimelockProposals[0]
		require.Len(t, prop.Operations, 1)
		assert.Equal(t, timelockAddr, prop.TimelockAddresses[mcms_types.ChainSelector(chainA)])
	})

	t.Run("TestRouterClassPromotesDirectly", func(t *testing.T) {
		e := newTestEnv(t)

		promoteOp := fakeBatchOp(chainA, "0xdeadbee8")
		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectTimelockOwnership(upgrader)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{TestRouterDests: []uint64{chainB}}, nil).Once()
		// prod-router-class checks/writes must NOT run.
		upgrader.EXPECT().PromoteOnrampToTestRouter(mock.Anything, chainA, []uint64{chainB}).
			Return([]mcms_types.BatchOperation{promoteOp}, nil).Once()

		cs := changesets.UpgradeOnrampPhase3(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.NoError(t, err)
		require.Len(t, out.MCMSTimelockProposals, 1)
	})

	t.Run("PromoteError", func(t *testing.T) {
		e := newTestEnv(t)
		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectTimelockOwnership(upgrader)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		expectLegacyOnProdRouter(upgrader)
		upgrader.EXPECT().PromoteOnrampToProdRouter(mock.Anything, chainA, []uint64{chainB}).
			Return(nil, errors.New("boom")).Once()

		cs := changesets.UpgradeOnrampPhase3(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "promote to prod Router")
		require.ErrorContains(t, err, "boom")
	})

	t.Run("OwnershipNotTransferred", func(t *testing.T) {
		e := newTestEnv(t)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		// The new OnRamp is still deployer-owned: Phase 1's ownership proposal
		// hasn't executed, so promotion must not proceed.
		upgrader.EXPECT().VerifyNewOnRampOwner(mock.Anything, chainA, timelockAddr).
			Return(errors.New("boom")).Once()
		// ClassifyDestChains and every promotion call must NOT be called.

		cs := changesets.UpgradeOnrampPhase3(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "pre-flight ownership check")
		require.ErrorContains(t, err, "boom")
	})

	t.Run("LegacyOnRampNotOnProdRouter", func(t *testing.T) {
		e := newTestEnv(t)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		expectTimelockOwnership(upgrader)
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		// The prod Router no longer routes through the legacy OnRamp: either Phase 3
		// already ran or an unexpected OnRamp is live.
		upgrader.EXPECT().VerifyLegacyOnRampOnProdRouter(mock.Anything, chainA, []uint64{chainB}).
			Return(errors.New("boom")).Once()
		// CopyDefaultCCVsFromLegacyOnRamp and PromoteOnrampToProdRouter must NOT be called.

		cs := changesets.UpgradeOnrampPhase3(
			mockUpgraderRegistry(upgrader),
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "pre-flight prod Router check")
		require.ErrorContains(t, err, "boom")
	})
}

func TestUpgradeOnrampCleanup(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		e := newTestEnv(t)
		// Post-merge state: the new OnRamp is the canonical one in the env datastore.
		seedOnRamp(t, &e, newOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		// chainB is prod-router class, so cleanup also un-wires the TestRouter for it.
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		upgrader.EXPECT().VerifyPromotedToProdRouter(mock.Anything, chainA, []uint64{chainB}).
			Return(nil).Once()
		upgrader.EXPECT().LegacyOnRampRef(mock.Anything, chainA).
			Return(legacyOnRampRef(), nil).Once()
		familyReg, lane := newLaneMocks(t)
		// Called twice: new OnRamp from the env datastore, legacy OnRamp from the
		// temp datastore used to wire-encode the legacy ref.
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Twice()
		// Pre-flight: the remote OffRamp whitelists exactly the old and new OnRamps.
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr), paddedAddrHex(newOnRampAddr)}, nil).Once()
		// Cleanup allows only the new OnRamp on the remote OffRamp, dropping the old one.
		cleanupOp := fakeBatchOp(chainB, "0x0ff9b")
		lane.setter.EXPECT().SetOffRampSourceOnRamps(mock.Anything, adapters.OffRampSetSourceOnRampsEntry{
			LocalChainSelector:  chainB,
			SourceChainSelector: chainA,
			OnRamps:             []string{paddedAddrHex(newOnRampAddr)},
		}).Return(&cleanupOp, false, nil).Once()

		cs := changesets.UpgradeOnrampCleanup(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			MCMS:                 validMCMSInput(),
			DestSelectorsInScope: []uint64{chainB},
		})
		require.NoError(t, err)

		require.Len(t, out.MCMSTimelockProposals, 1)
		prop := out.MCMSTimelockProposals[0]
		require.Len(t, prop.Operations, 1)
	})

	t.Run("DestChainSelectorsError", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, newOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return(nil, errors.New("boom")).Once()

		// The changeset resolves the OnRamp address before dest chain selectors.
		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Once()

		cs := changesets.UpgradeOnrampCleanup(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "resolve dest chain selectors")
		require.ErrorContains(t, err, "boom")
	})

	t.Run("UnexpectedOnRampOnRemoteOffRamp", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, newOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{TestRouterDests: []uint64{chainB}}, nil).Once()
		expectPromotedToProdRouter(upgrader)
		upgrader.EXPECT().LegacyOnRampRef(mock.Anything, chainA).
			Return(legacyOnRampRef(), nil).Once()

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Twice()
		// The remote OffRamp whitelists an unknown onramp alongside the new one;
		// setting [new] would silently drop it.
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(newOnRampAddr), paddedAddrHex(unknownOnRampAddr)}, nil).Once()
		// The setter must NOT be called: the pre-flight check fails first.

		cs := changesets.UpgradeOnrampCleanup(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "pre-flight OffRamp whitelist check")
		require.ErrorContains(t, err, paddedAddrHex(unknownOnRampAddr))
	})

	t.Run("NoBatchOps", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, newOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		// chainB is TestRouter-class here, so UnwireTestRouter must NOT be called —
		// otherwise this test's "no batch ops" assertion wouldn't hold.
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{TestRouterDests: []uint64{chainB}}, nil).Once()
		expectPromotedToProdRouter(upgrader)
		upgrader.EXPECT().LegacyOnRampRef(mock.Anything, chainA).
			Return(legacyOnRampRef(), nil).Once()

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Twice()
		lane.reader.EXPECT().GetOffRampSourceOnRamps(mock.Anything, chainB, chainA).
			Return([]string{paddedAddrHex(oldOnRampAddr), paddedAddrHex(newOnRampAddr)}, nil).Once()
		// The remote OffRamp already allows only the new OnRamp: update is skipped.
		lane.setter.EXPECT().SetOffRampSourceOnRamps(mock.Anything, mock.Anything).
			Return(nil, true, nil).Once()

		cs := changesets.UpgradeOnrampCleanup(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
		)
		out, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.NoError(t, err)
		assert.Empty(t, out.MCMSTimelockProposals, "no proposal expected when all updates are skipped")
	})

	t.Run("NotPromotedToProdRouter", func(t *testing.T) {
		e := newTestEnv(t)
		seedOnRamp(t, &e, newOnRampAddr)

		upgrader := mocks.NewMockOnRampUpgrader(t)
		upgrader.EXPECT().DestChainSelectors(mock.Anything, chainA).
			Return([]uint64{chainB}, nil).Once()
		upgrader.EXPECT().ClassifyDestChains(mock.Anything, chainA).
			Return(adapters.LaneClass{ProdRouterDests: []uint64{chainB}}, nil).Once()
		// The prod Router still routes through the old OnRamp: Phase 3 hasn't
		// executed, so removing the old OnRamp from whitelists would drop traffic.
		upgrader.EXPECT().VerifyPromotedToProdRouter(mock.Anything, chainA, []uint64{chainB}).
			Return(errors.New("boom")).Once()
		// LegacyOnRampRef and the OffRamp setter must NOT be called.

		familyReg, lane := newLaneMocks(t)
		lane.family.EXPECT().GetOnRampAddress(mock.Anything, chainA).
			RunAndReturn(onRampAddressFromDS).Once()

		cs := changesets.UpgradeOnrampCleanup(
			mockUpgraderRegistry(upgrader),
			familyReg,
			mockMCMSReaderRegistry(t),
		)
		_, err := cs.Apply(e, changesets.UpgradeOnrampConfig{
			ChainSelector:        chainA,
			DestSelectorsInScope: []uint64{chainB},
			MCMS:                 validMCMSInput(),
		})
		require.ErrorContains(t, err, "pre-flight promotion check")
		require.ErrorContains(t, err, "boom")
	})
}
