package tokens

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type fakeTokenTransferFeeReader struct {
	onchain  TokenTransferFeeConfig
	defaults TokenTransferFeeConfig
	err      error
}

func (f fakeTokenTransferFeeReader) GetOnchainTokenTransferFeeConfig(deployment.Environment, string, uint64, uint64) (TokenTransferFeeConfig, error) {
	return f.onchain, f.err
}

func (f fakeTokenTransferFeeReader) GetDefaultTokenTransferFeeConfig(uint64, uint64) TokenTransferFeeConfig {
	return f.defaults
}

func TestInferTokenTransferFeeArgsSkipsMatchingOnchainConfig(t *testing.T) {
	onchain := TokenTransferFeeConfig{
		DestGasOverhead:               90_000,
		DestBytesOverhead:             32,
		DefaultFinalityFeeUSDCents:    25,
		CustomFinalityFeeUSDCents:     100,
		DefaultFinalityTransferFeeBps: 5,
		CustomFinalityTransferFeeBps:  6,
		IsEnabled:                     true,
	}

	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: onchain},
		deployment.Environment{Logger: logger.Test(t)},
		"pool",
		1,
		2,
		TokenTransferFeeForDst{},
	)

	require.NoError(t, err)
	require.False(t, shouldApply)
	require.Nil(t, got)
}

func TestInferTokenTransferFeeArgsReturnsChangedConfig(t *testing.T) {
	onchain := TokenTransferFeeConfig{
		DestGasOverhead:               90_000,
		DestBytesOverhead:             32,
		DefaultFinalityFeeUSDCents:    25,
		CustomFinalityFeeUSDCents:     100,
		DefaultFinalityTransferFeeBps: 5,
		CustomFinalityTransferFeeBps:  6,
		IsEnabled:                     true,
	}

	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: onchain},
		deployment.Environment{Logger: logger.Test(t)},
		"pool",
		1,
		2,
		TokenTransferFeeForDst{
			Settings: UnresolvedTokenTransferFeeArgs{
				DefaultFinalityFeeUSDCents: utils.NewOptional(uint32(26)),
			},
		},
	)

	require.NoError(t, err)
	require.True(t, shouldApply)
	require.NotNil(t, got)
	require.Equal(t, uint32(26), got.DefaultFinalityFeeUSDCents)
	require.Equal(t, onchain.CustomFinalityFeeUSDCents, got.CustomFinalityFeeUSDCents)
}

func TestInferTokenTransferFeeArgsSkipsResetWhenAlreadyDisabled(t *testing.T) {
	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: TokenTransferFeeConfig{IsEnabled: false}},
		deployment.Environment{Logger: logger.Test(t)},
		"pool",
		1,
		2,
		TokenTransferFeeForDst{IsReset: true},
	)

	require.NoError(t, err)
	require.False(t, shouldApply)
	require.Nil(t, got)
}

func TestInferTokenTransferFeeArgsReturnsResetWhenEnabled(t *testing.T) {
	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: TokenTransferFeeConfig{IsEnabled: true}},
		deployment.Environment{Logger: logger.Test(t)},
		"pool",
		1,
		2,
		TokenTransferFeeForDst{IsReset: true},
	)

	require.NoError(t, err)
	require.True(t, shouldApply)
	require.Nil(t, got)
}

func TestInferTokenTransferFeeArgsReturnsOnchainReadError(t *testing.T) {
	readErr := errors.New("read failed")

	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{err: readErr},
		deployment.Environment{Logger: logger.Test(t)},
		"pool",
		1,
		2,
		TokenTransferFeeForDst{},
	)

	require.ErrorIs(t, err, readErr)
	require.False(t, shouldApply)
	require.Nil(t, got)
}
