package discovery

import (
	"errors"
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
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
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
	}
	mockReader.
		EXPECT().
		DiscoverContracts(mock.Anything, dest).
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
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.NoError(t, err)
	assert.Equal(t, expectedFChain, observation.FChain)
	assert.Equal(t, expectedOnRamp, observation.OnRamp)
	assert.Equal(t, expectedNonceManager, observation.DestNonceManager)
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
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.Error(t, err)
	assert.Empty(t, observation.DestNonceManager)
	assert.Empty(t, observation.OnRamp)
	assert.Empty(t, observation.FChain)
}

func TestContractDiscoveryProcessor_Observation_DontSupportDest_StillObserveFChain(t *testing.T) {
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
		DiscoverContracts(mock.Anything, dest).
		Return(nil, reader.ErrContractReaderNotFound)
	mockHomeChain.EXPECT().GetFChain().Return(expectedFChain, nil)
	defer mockReader.AssertExpectations(t)
	defer mockHomeChain.AssertExpectations(t)

	cdp := NewContractDiscoveryProcessor(
		lggr,
		&mockReaderIface,
		mockHomeChain,
		dest,
		fRoleDON,
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.NoError(t, err)
	assert.Equal(t, expectedFChain, observation.FChain)
	assert.Empty(t, observation.DestNonceManager)
	assert.Empty(t, observation.OnRamp)
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
		DiscoverContracts(mock.Anything, dest).
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
	)

	observation, err := cdp.Observation(ctx, discoverytypes.Outcome{}, discoverytypes.Query{})
	assert.Error(t, err)
	assert.Empty(t, observation.FChain)
	assert.Empty(t, observation.DestNonceManager)
	assert.Empty(t, observation.OnRamp)
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
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
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
	)

	obs := discoverytypes.Observation{
		FChain:           expectedFChain,
		OnRamp:           expectedOnRamp,
		DestNonceManager: expectedNonceManager,
	}
	// here we have:
	// 2*fRoleDON + 1 observations of fChain
	// 2*fDest + 1 observations of the onramps and the dest nonce manager
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: obs},
		{Observation: obs},
		{Observation: obs},
	}

	outcome, err := cdp.Outcome(discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
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
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
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
	)

	fChainObs := discoverytypes.Observation{
		FChain: expectedFChain,
	}
	destObs := discoverytypes.Observation{
		FChain:           expectedFChain,
		OnRamp:           expectedOnRamp,
		DestNonceManager: expectedNonceManager,
	}
	// here we have:
	// 2*fRoleDON + 1 == 7 observations of fChain
	// 2*fDest + 1 == 5 observations of the onramps and the dest nonce manager
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: fChainObs},
		{Observation: fChainObs},
	}

	outcome, err := cdp.Outcome(discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
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
	observedOnRamp := map[cciptypes.ChainSelector][]byte{
		source1: []byte("onRamp"),
		source2: []byte("onRamp"),
	}
	// we expect no contracts here due to not enough observations to come to consensus.
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{},
		consts.ContractNameOnRamp:       map[cciptypes.ChainSelector][]byte{},
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
	)

	fChainObs := discoverytypes.Observation{
		FChain: expectedFChain,
	}
	destObs := discoverytypes.Observation{
		FChain:           expectedFChain,
		OnRamp:           observedOnRamp,
		DestNonceManager: observedNonceManager,
	}
	// here we have:
	// 2*fRoleDON + 1 == 7 observations of fChain
	// 2*fDest == 4 observations of the onramps and the dest nonce manager (not enough).
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: destObs},
		{Observation: fChainObs},
		{Observation: fChainObs},
		{Observation: fChainObs},
	}

	outcome, err := cdp.Outcome(discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
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
	expectedContracts := reader.ContractAddresses{
		consts.ContractNameNonceManager: map[cciptypes.ChainSelector][]byte{
			dest: expectedNonceManager,
		},
		consts.ContractNameOnRamp: expectedOnRamp,
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
	)

	obs := discoverytypes.Observation{
		FChain:           expectedFChain,
		OnRamp:           expectedOnRamp,
		DestNonceManager: expectedNonceManager,
	}
	// here we have:
	// 2*fRoleDON + 1 observations of fChain
	// 2*fDest + 1 observations of the onramps and the dest nonce manager
	aos := []plugincommon.AttributedObservation[discoverytypes.Observation]{
		{Observation: obs},
		{Observation: obs},
		{Observation: obs},
	}

	outcome, err := cdp.Outcome(discoverytypes.Outcome{}, discoverytypes.Query{}, aos)
	assert.Error(t, err)
	assert.ErrorIs(t, err, syncErr)
	assert.Empty(t, outcome)
}
