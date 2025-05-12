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
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	mock_home_chain "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	mock_reader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
	expectedNonceManager := cciptypes.UnknownAddress("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		source: cciptypes.UnknownAddress("onRamp"),
	}
	expectedFeeQuoter := map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		source: cciptypes.UnknownAddress("from_source_onramp"),
		dest:   cciptypes.UnknownAddress("from_dest_offramp"),
	}
	expectedRMNRemote := cciptypes.UnknownAddress("rmnRemote")
	expectedRouter := cciptypes.UnknownAddress("router")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp:    expectedOnRamp,
		consts.ContractNameFeeQuoter: expectedFeeQuoter,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedRMNRemote,
		},
		consts.ContractNameRouter: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			source: expectedRouter,
		},
	}
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, mock.Anything).
		Return(expectedContracts, nil)

	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	mockHomeChain.EXPECT().GetAllChainConfigs().Return(nil, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test,
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

	cdp := internalNewContractDiscoveryProcessor(
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
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, mock.Anything).
		Return(nil, nil)

	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	mockHomeChain.EXPECT().GetAllChainConfigs().Return(nil, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := internalNewContractDiscoveryProcessor(
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
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, mock.Anything).
		Return(nil, discoveryErr)
	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	mockHomeChain.EXPECT().GetAllChainConfigs().Return(nil, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := internalNewContractDiscoveryProcessor(
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
	expectedNonceManager := cciptypes.UnknownAddress("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		dest: cciptypes.UnknownAddress("onRamp"),
	}
	expectedRMNRemote := cciptypes.UnknownAddress("rmnRemote")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedRMNRemote,
		},
		consts.ContractNameFeeQuoter: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			source1: cciptypes.UnknownAddress("feeQuoter1"),
			source2: cciptypes.UnknownAddress("feeQuoter2"),
		},
		consts.ContractNameRouter: {},
	}
	mockReader.
		EXPECT().
		Sync(mock.Anything, expectedContracts).
		Return(nil)
	defer mockReader.AssertExpectations(t)

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	obsSrc := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
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
	expectedNonceManager := cciptypes.UnknownAddress("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		dest: cciptypes.UnknownAddress("onRamp"),
	}
	expectedRMNRemote := cciptypes.UnknownAddress("rmnRemote")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedRMNRemote,
		},
		consts.ContractNameRouter: {
			dest: cciptypes.UnknownAddress("router"),
		},
		consts.ContractNameFeeQuoter: {}, // no consensus
	}
	mockReader.
		EXPECT().
		Sync(mock.Anything, expectedContracts).
		Return(nil)
	defer mockReader.AssertExpectations(t)

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	fChainObs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			consts.ContractNameFeeQuoter: {
				source1: cciptypes.UnknownAddress("fee_quoter_1"),
				source2: cciptypes.UnknownAddress("fee_quoter_2"),
			},
		},
	}
	destObs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			consts.ContractNameOnRamp: expectedOnRamp,
			consts.ContractNameNonceManager: {
				dest: expectedNonceManager,
			},
			consts.ContractNameRMNRemote: {
				dest: expectedRMNRemote,
			},
			consts.ContractNameRouter: {
				dest: cciptypes.UnknownAddress("router"),
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
	observedNonceManager := cciptypes.UnknownAddress("nonceManager")
	observedRMNRemote := cciptypes.UnknownAddress("rmnRemote")
	observedOnRamp := map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		source1: cciptypes.UnknownAddress("onRamp"),
		source2: cciptypes.UnknownAddress("onRamp"),
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

	cdp := internalNewContractDiscoveryProcessor(
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
		Addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
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
	expectedNonceManager := cciptypes.UnknownAddress("nonceManager")
	expectedOnRamp := map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		dest: cciptypes.UnknownAddress("onRamp"),
	}
	expectedRMNRemote := cciptypes.UnknownAddress("rmnRemote")
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
		consts.ContractNameRMNRemote: map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
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

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
		nil, // oracleIDToP2PID, not needed for this test
	)

	obs := discoverytypes.Observation{
		FChain: expectedFChain,
		Addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
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
	require.NoError(t, err)
	require.Equal(t, outcome, discoverytypes.Outcome{})
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

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID:    oracleID,
		Observation: dummyObservation,
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

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID:    oracleID,
		Observation: dummyObservation,
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

	cdp := internalNewContractDiscoveryProcessor(
		lggr,
		nil, // reader, not needed for this test
		mockHomeChain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
	)

	ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
		OracleID:    oracleID,
		Observation: dummyObservation,
	}

	err := cdp.ValidateObservation(discoverytypes.Outcome{}, discoverytypes.Query{}, ao)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), fmt.Sprintf("unable to get supported chains for Oracle %d", ao.OracleID))
}

func TestContractDiscoveryProcessor_ValidateObservation_OracleNotAllowedToObserve(t *testing.T) {
	dest := cciptypes.ChainSelector(1)

	cases := []struct {
		name            string
		supportedChains []cciptypes.ChainSelector
		addresses       reader.ContractAddresses
		fChain          map[cciptypes.ChainSelector]int
		errStr          string
	}{
		{
			name: "no observations no error",
		},
		{
			name: "onramps are only discovered on dest (error)",
			addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
				consts.ContractNameOnRamp: {
					dest + 1: cciptypes.UnknownAddress("1"),
					dest + 2: cciptypes.UnknownAddress("2"),
					dest + 3: cciptypes.UnknownAddress("3"),
				},
			},
			fChain: defaultFChain,
			errStr: "oracle 1 is not allowed to observe contract (OnRamp) on the destination chain ChainSelector(1)",
		},
		{
			name:            "onramps are only discovered on dest (pass)",
			supportedChains: []cciptypes.ChainSelector{dest},
			addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
				consts.ContractNameOnRamp: {
					dest + 1: cciptypes.UnknownAddress("1"),
					dest + 2: cciptypes.UnknownAddress("2"),
					dest + 3: cciptypes.UnknownAddress("3"),
				},
			},
			fChain: defaultFChain,
		},
		{
			name:            "FeeQuoter is discovered on the same chain (error)",
			supportedChains: []cciptypes.ChainSelector{dest},
			addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
				consts.ContractNameFeeQuoter: {
					dest + 1: cciptypes.UnknownAddress("1"),
					dest + 2: cciptypes.UnknownAddress("2"),
					dest + 3: cciptypes.UnknownAddress("3"),
				},
			},
			fChain: defaultFChain,
			errStr: "oracle 1 is not allowed to observe chain ChainSelector",
		},
		{
			name:            "FeeQuoter is discovered on the same chain (pass)",
			supportedChains: []cciptypes.ChainSelector{dest, dest + 1, dest + 2, dest + 3},
			addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
				consts.ContractNameFeeQuoter: {
					dest + 1: cciptypes.UnknownAddress("1"),
					dest + 2: cciptypes.UnknownAddress("2"),
					dest + 3: cciptypes.UnknownAddress("3"),
				},
				// Need to observe onramps to be able to observe feequoters
				consts.ContractNameOnRamp: {
					dest + 1: cciptypes.UnknownAddress("a"),
					dest + 2: cciptypes.UnknownAddress("b"),
					dest + 3: cciptypes.UnknownAddress("c"),
				},
			},
			fChain: defaultFChain,
		},
		{
			name:            "Invalid FChain (error)",
			supportedChains: []cciptypes.ChainSelector{dest, dest + 1},
			addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
				consts.ContractNameOnRamp: {
					dest + 1: cciptypes.UnknownAddress("1"),
				},
			},
			fChain: map[cciptypes.ChainSelector]int{
				dest: -2,
			},
			errStr: "fChain for chain 1 is not positive: -2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			fRoleDON := 1
			oracleID := commontypes.OracleID(1)
			peerID := ragep2ptypes.PeerID([32]byte{1, 2, 3})

			oracleIDToP2PID := map[commontypes.OracleID]ragep2ptypes.PeerID{
				oracleID: peerID,
			}

			mockHomeChain := mock_home_chain.NewMockHomeChain(t)
			mockHomeChain.EXPECT().GetSupportedChainsForPeer(peerID).Return(mapset.NewSet(tc.supportedChains...),
				nil).Maybe()
			defer mockHomeChain.AssertExpectations(t)

			cdp := internalNewContractDiscoveryProcessor(
				lggr,
				nil, // reader, not needed for this test
				mockHomeChain,
				dest,
				fRoleDON,
				oracleIDToP2PID,
			)

			ao := plugincommon.AttributedObservation[discoverytypes.Observation]{
				OracleID: oracleID,
				Observation: discoverytypes.Observation{
					Addresses: tc.addresses,
					FChain:    tc.fChain,
				},
			}

			err := cdp.ValidateObservation(discoverytypes.Outcome{}, discoverytypes.Query{}, ao)
			if tc.errStr == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tc.errStr)
			}
		})
	}
}

var defaultFChain = map[cciptypes.ChainSelector]int{
	1: 1,
}
var dummyObservation = discoverytypes.Observation{
	FChain: map[cciptypes.ChainSelector]int{
		1: 1,
	},
	Addresses: map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
		consts.ContractNameOnRamp: {
			1: cciptypes.UnknownAddress("someaddress"),
		},
	},
}

func internalNewContractDiscoveryProcessor(
	lggr logger.Logger,
	reader *reader.CCIPReader,
	homechain reader.HomeChain,
	dest cciptypes.ChainSelector,
	fRoleDON int,
	oracleIDToP2PID map[commontypes.OracleID]ragep2ptypes.PeerID,
) plugincommon.PluginProcessor[discoverytypes.Query, discoverytypes.Observation, discoverytypes.Outcome] {
	return NewContractDiscoveryProcessor(
		lggr,
		reader,
		homechain,
		dest,
		fRoleDON,
		oracleIDToP2PID,
		plugincommon.NoopReporter{},
	)
}
