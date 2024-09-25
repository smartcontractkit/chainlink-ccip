package discovery

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	mock_home_chain "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	mock_reader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func TestContractDiscoveryProcessor_Observation(t *testing.T) {
	t.Run("support the dest, happy path", func(t *testing.T) {
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
	})

	t.Run("error getting fchain", func(t *testing.T) {
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
	})

	t.Run("don't support the dest, still observe fChain", func(t *testing.T) {
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
	})

	t.Run("error discovering contracts", func(t *testing.T) {
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
	})

}
