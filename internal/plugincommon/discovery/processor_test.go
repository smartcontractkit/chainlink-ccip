package discovery

import (
	"errors"
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	mock_home_chain "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	mock_reader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func TestContractDiscoveryProcessor_Observation_SupportsDest_HappyPath(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source := cciptypes.ChainSelector(2)
	fRoleDON := 1

	ctx := tests.Context(t)
	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:   1,
		source: 2,
	}
	expectedNonceManager := []byte("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector][]byte{
		source: []byte("onRamp"),
	}
	expectedFeeQuoter := map[cciptypes.ChainSelector][]byte{
		source: []byte("from_source_onramp"),
		dest:   []byte("from_dest_offramp"),
	}
	expectedRMNRemote := []byte("rmnRemote")
	expectedRouter := []byte("router")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp:    expectedOnRamp,
		consts.ContractNameFeeQuoter: expectedFeeQuoter,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector][]byte{
			dest: expectedRMNRemote,
		},
		consts.ContractNameRouter: map[cciptypes.ChainSelector][]byte{
			source: expectedRouter,
		},
	}
	var emptySelectors []cciptypes.ChainSelector
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, emptySelectors).
		Return(expectedContracts, nil)

	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.NoError(t, err)
	assert.Equal(t, expectedFChain, observation.FChain)
	assert.Equal(t, expectedOnRamp, observation.Addresses[consts.ContractNameOnRamp])
	assert.Equal(t, expectedFeeQuoter, observation.Addresses[consts.ContractNameFeeQuoter])

	require.Len(t, observation.Addresses[consts.ContractNameNonceManager], 1)
	assert.Equal(t, expectedNonceManager, observation.Addresses[consts.ContractNameNonceManager][dest])
	require.Len(t, observation.Addresses[consts.ContractNameRMNRemote], 1)
	assert.Equal(t, expectedRMNRemote, observation.Addresses[consts.ContractNameRMNRemote][dest])
	require.Len(t, observation.Addresses[consts.ContractNameRouter], 1)
	assert.Equal(t, expectedRouter, observation.Addresses[consts.ContractNameRouter][source])
}

func TestContractDiscoveryProcessor_Observation_ErrorGettingFChain(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	fRoleDON := 1

	ctx := tests.Context(t)
	expectedErr := fmt.Errorf("error getting fchain")
	mockHomeChain.EXPECT().GetFChain().Return(nil, expectedErr)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.Error(t, err)
	assert.Empty(t, observation.Addresses)
	assert.Empty(t, observation.FChain)
}

// No dest reader and source readers not ready, still observes fChain from home chain.
func TestContractDiscoveryProcessor_Observation_SourceReadersNotReady(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source := cciptypes.ChainSelector(2)
	fRoleDON := 1

	ctx := tests.Context(t)
	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:   1,
		source: 2,
	}
	var emptySelectors []cciptypes.ChainSelector
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, emptySelectors).
		Return(nil, nil)

	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.NoError(t, err)
	assert.Equal(t, expectedFChain, observation.FChain)
	assert.Empty(t, observation.Addresses)
}

func TestContractDiscoveryProcessor_Observation_ErrorDiscoveringContracts(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source := cciptypes.ChainSelector(2)
	fRoleDON := 1

	ctx := tests.Context(t)
	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:   1,
		source: 2,
	}
	discoveryErr := fmt.Errorf("discovery error")
	var emptySelectors []cciptypes.ChainSelector
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, emptySelectors).
		Return(nil, discoveryErr)
	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.Error(t, err)
	assert.Empty(t, observation.FChain)
	assert.Empty(t, observation.Addresses)
}

func TestContractDiscoveryProcessor_Outcome_HappyPath(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source1 := cciptypes.ChainSelector(2)
	source2 := cciptypes.ChainSelector(3)
	fRoleDON := 1
	fDest := 1
	fSource1 := 1
	fSource2 := 1

	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:    fDest,
		source1: fSource1,
		source2: fSource2,
	}
	expectedNonceManager := []byte("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector][]byte{
		dest: []byte("onRamp"),
	}
	expectedRMNRemote := []byte("rmnRemote")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector][]byte{
			dest: expectedRMNRemote,
		},
		consts.ContractNameFeeQuoter: map[cciptypes.ChainSelector][]byte{
			source1: []byte("feeQuoter1"),
			source2: []byte("feeQuoter2"),
		},
		consts.ContractNameRouter: map[cciptypes.ChainSelector][]byte{
			dest: []byte("dest_router"),
		},
	}
	mockReader.
		EXPECT().
		Sync(mock.Anything, expectedContracts).
		Return(nil)
	defer mockReader.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	obsSrc := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector][]byte{
			consts.ContractNameOnRamp: expectedOnRamp,
			consts.ContractNameNonceManager: {
				dest: expectedNonceManager,
			},
			consts.ContractNameRMNRemote: {
				dest: expectedRMNRemote,
			},
			consts.ContractNameFeeQuoter: {
				source1: expectedContracts[consts.ContractNameFeeQuoter][source1],
				source2: expectedContracts[consts.ContractNameFeeQuoter][source2],
			},
			consts.ContractNameRouter: {
				source1: expectedContracts[consts.ContractNameRouter][dest],
				source2: expectedContracts[consts.ContractNameRouter][dest],
			},
		},
	}

	// here we have:
	// 2*fRoleDON + 1 observations of fChain
	// 2*fDest + 1 observations of the onramps and the dest nonce manager
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: obsSrc},
		{Observation: obsSrc},
		{Observation: obsSrc},
	}

	ctx := tests.Context(t)
	outcome, err := cdp.Outcome(ctx, discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
	assert.NoError(t, err)
	assert.Empty(t, outcome)
}

func TestContractDiscovery_Outcome_HappyPath_FRoleDONAndFDestChainAreDifferent(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source1 := cciptypes.ChainSelector(2)
	source2 := cciptypes.ChainSelector(3)
	fRoleDON := 3
	fDest := 2
	fSource1 := 1
	fSource2 := 1

	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:    fDest,
		source1: fSource1,
		source2: fSource2,
	}
	expectedNonceManager := []byte("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector][]byte{
		dest: []byte("onRamp"),
	}
	expectedRMNRemote := []byte("rmnRemote")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector][]byte{
			dest: expectedRMNRemote,
		},
		consts.ContractNameFeeQuoter: {}, // no consensus
		consts.ContractNameRouter:    {}, // no consensus
	}
	mockReader.
		EXPECT().
		Sync(mock.Anything, expectedContracts).
		Return(nil)
	defer mockReader.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	fChainObs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector][]byte{
			consts.ContractNameRouter: {
				source1: []byte("router"),
				source2: []byte("router"),
			},
			consts.ContractNameFeeQuoter: {
				source1: []byte("fee_quoter_1"),
				source2: []byte("fee_quoter_2"),
			},
		},
	}
	destObs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector][]byte{
			consts.ContractNameOnRamp: expectedOnRamp,
			consts.ContractNameNonceManager: {
				dest: expectedNonceManager,
			},
			consts.ContractNameRMNRemote: {
				dest: expectedRMNRemote,
			},
		},
	}
	// here we have:
	// 2*fRoleDON + 1 == 7 observations of fChain
	// 2*fDest + 1 == 5 observations of the onramps, the dest nonce manager, and the RMNRemote
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: fChainObs},
		{Observation: fChainObs}, // no consensus on fChainObs
	}

	ctx := tests.Context(t)
	outcome, err := cdp.Outcome(ctx, discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
	assert.NoError(t, err)
	assert.Empty(t, outcome)
}

func TestContractDiscoveryProcessor_Outcome_NotEnoughObservations(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source1 := cciptypes.ChainSelector(2)
	source2 := cciptypes.ChainSelector(3)
	fRoleDON := 3
	fDest := 2
	fSource1 := 1
	fSource2 := 1

	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:    fDest,
		source1: fSource1,
		source2: fSource2,
	}
	observedNonceManager := []byte("nonceManager")
	observedRMNRemote := []byte("rmnRemote")
	observedOnRamp := map[cciptypes.ChainSelector][]byte{
		source1: []byte("onRamp"),
		source2: []byte("onRamp"),
	}
	// we expect no contracts here due to not enough observations to come to consensus.
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: {},
		consts.ContractNameOnRamp:       {},
		consts.ContractNameRMNRemote:    {},
		consts.ContractNameFeeQuoter:    {},
		consts.ContractNameRouter:       {},
	}
	mockReader.
		EXPECT().
		Sync(mock.Anything, expectedContracts).
		Return(nil)
	defer mockReader.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	fChainObs := discoverytypes.Observation{
		FChain: expectedFChain,
	}
	destObs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector][]byte{
			consts.ContractNameOnRamp: observedOnRamp,
			consts.ContractNameNonceManager: {
				dest: observedNonceManager,
			},
			consts.ContractNameRMNRemote: {
				dest: observedRMNRemote,
			},
			consts.ContractNameFeeQuoter: {},
			consts.ContractNameRouter:    {},
		},
	}
	// here we have:
	// 2*fRoleDON + 1 == 7 observations of fChain
	// 2*fDest == 4 observations of the onramps, the dest nonce manager and the RMNRemote (not enough)
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs}, // dest requires 2*f+1, for f=2 we need 5 observations
		{Observation: fChainObs},
		{Observation: fChainObs},
		{Observation: fChainObs},
	}

	ctx := tests.Context(t)
	outcome, err := cdp.Outcome(ctx, discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
	assert.NoError(t, err)
	assert.Empty(t, outcome)
}

func TestContractDiscoveryProcessor_Outcome_ErrorSyncingContracts(t *testing.T) {
	mockReader := mock_reader.NewMockCCIPReader(t)
	mockReaderIface := reader.CCIPReader(mockReader)
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	source1 := cciptypes.ChainSelector(2)
	source2 := cciptypes.ChainSelector(3)
	fRoleDON := 1
	fDest := 1
	fSource1 := 1
	fSource2 := 1

	expectedFChain := map[cciptypes.ChainSelector]int{
		dest:    fDest,
		source1: fSource1,
		source2: fSource2,
	}
	expectedNonceManager := []byte("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector][]byte{
		dest: []byte("onRamp"),
	}
	expectedRMNRemote := []byte("rmnRemote")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector][]byte{
			dest: expectedRMNRemote,
		},
		consts.ContractNameFeeQuoter: {},
		consts.ContractNameRouter:    {},
	}
	syncErr := errors.New("sync error")
	mockReader.
		EXPECT().
		Sync(mock.Anything, expectedContracts).
		Return(syncErr)
	defer mockReader.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	obs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector][]byte{
			consts.ContractNameOnRamp: expectedOnRamp,
			consts.ContractNameNonceManager: {
				dest: expectedNonceManager,
			},
			consts.ContractNameRMNRemote: {
				dest: expectedRMNRemote,
			},
		},
	}
	// here we have:
	// 2*fRoleDON + 1 observations of fChain
	// 2*fDest + 1 observations of the onramps, the dest nonce manager and the RMNRemote
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: obs},
		{Observation: obs},
		{Observation: obs},
	}

	ctx := tests.Context(t)
	outcome, err := cdp.Outcome(ctx, discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
	assert.Error(t, err)
	assert.ErrorIs(t, err, syncErr)
	assert.Empty(t, outcome)
}

func TestContractDiscoveryProcessor_ValidateObservation_HappyPath(t *testing.T) {
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	fRoleDON := 1
	oracleID := commontypes.OracleID(1)
	peerID := ragep2ptypes.PeerID([32]byte{1, 2, 3})
	supportedChains := mapset.NewSet(dest)

	oracleIDToP2PID := map[commontypes.OracleID]ragep2ptypes.PeerID{
		oracleID: peerID,
	}

	mockHomeChain.EXPECT().GetSupportedChainsForPeer(peerID).Return(supportedChains, nil)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID: oracleID,
	}

	err := cdp.ValidateObservation(discoverytypes.Outcome{}, discoverytypes.Query{}, ao)
	assert.NoError(t, err)
}

func TestContractDiscoveryProcessor_ValidateObservation_NoPeerID(t *testing.T) {
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	fRoleDON := 1
	oracleID := commontypes.OracleID(1)

	oracleIDToP2PID := map[commontypes.OracleID]ragep2ptypes.PeerID{}

	cdp := NewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID: oracleID,
	}

	err := cdp.ValidateObservation(discoverytypes.Outcome{}, discoverytypes.Query{}, ao)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("no peer ID found for Oracle %d", ao.OracleID))
}

func TestContractDiscoveryProcessor_ValidateObservation_ErrorGettingSupportedChains(t *testing.T) {
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	fRoleDON := 1
	oracleID := commontypes.OracleID(1)
	peerID := ragep2ptypes.PeerID([32]byte{1, 2, 3})
	expectedErr := fmt.Errorf("error getting supported chains")

	oracleIDToP2PID := map[commontypes.OracleID]ragep2ptypes.PeerID{
		oracleID: peerID,
	}

	mockHomeChain.EXPECT().GetSupportedChainsForPeer(peerID).Return(nil, expectedErr)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID: oracleID,
	}

	err := cdp.ValidateObservation(discoverytypes.Outcome{}, discoverytypes.Query{}, ao)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("unable to get supported chains for Oracle %d", ao.OracleID))
}

func TestContractDiscoveryProcessor_ValidateObservation_OracleNotAllowedToObserve(t *testing.T) {
	mockHomeChain := mock_home_chain.NewMockHomeChain(t)
	lggr := logger.Test(t)
	dest := cciptypes.ChainSelector(1)
	fRoleDON := 1
	oracleID := commontypes.OracleID(1)
	peerID := ragep2ptypes.PeerID([32]byte{1, 2, 3})
	supportedChains := mapset.NewSet(cciptypes.ChainSelector(2)) // Different chain

	oracleIDToP2PID := map[commontypes.OracleID]ragep2ptypes.PeerID{
		oracleID: peerID,
	}

	mockHomeChain.EXPECT().GetSupportedChainsForPeer(peerID).Return(supportedChains, nil)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID: oracleID,
	}

	err := cdp.ValidateObservation(discoverytypes.Outcome{}, discoverytypes.Query{}, ao)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("oracle %d is not allowed to observe chain %s", ao.OracleID, cdp.dest))
}

func Test_getRouterConsensus(t *testing.T) {
	tests := []struct {
		name        string
		routerAddrs map[cciptypes.ChainSelector][][]byte
		want        []byte
		err         string
	}{
		{
			name: "no inputs",
			err:  "no consensus on router",
		},
		{
			name: "total agreement",
			routerAddrs: map[cciptypes.ChainSelector][][]byte{
				1: {{1, 2, 3}, {1, 2, 3}},
				2: {{1, 2, 3}, {1, 2, 3}},
				3: {{1, 2, 3}, {1, 2, 3}},
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "threshold 2 not reached",
			routerAddrs: map[cciptypes.ChainSelector][][]byte{
				1: {{1, 2, 3}},
				2: {{1, 2, 3}},
				3: {{1, 2, 3}},
			},
			err: "no consensus on router, routerAddrs",
		},
		{
			name: "3 way tie",
			routerAddrs: map[cciptypes.ChainSelector][][]byte{
				1: {{1, 1, 1}, {1, 1, 1}},
				2: {{2, 2, 2}, {2, 2, 2}},
				3: {{3, 3, 3}, {3, 3, 3}},
			},
			err: "no consensus on router, there is a tie multiple routers were seen 1 times",
		},
		{
			name: "majority",
			routerAddrs: map[cciptypes.ChainSelector][][]byte{
				1: {{1, 1, 1}, {1, 1, 1}},
				2: {{2, 2, 2}, {2, 2, 2}},
				3: {{1, 1, 1}, {1, 1, 1}},
			},
			want: []byte{1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			lggr := logger.Test(t)

			thresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](2)
			got, err := getOnRampDestRouterConsensus(lggr, tt.routerAddrs, thresh)
			if tt.err != "" {
				assert.ErrorContains(t, err, tt.err)
				return
			}

			require.NoError(t, err)
			assert.Equalf(t, tt.want, got, "getOnRampDestRouterConsensus(...)")
		})
	}
}
