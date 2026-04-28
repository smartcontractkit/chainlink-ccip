package fees

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type fakeTokenTransferFeeReader struct {
	onchain  TokenTransferFeeArgs
	defaults TokenTransferFeeArgs
	err      error
}

func (f fakeTokenTransferFeeReader) GetOnchainTokenTransferFeeConfig(cldf.Environment, uint64, uint64, string) (TokenTransferFeeArgs, error) {
	return f.onchain, f.err
}

func (f fakeTokenTransferFeeReader) GetDefaultTokenTransferFeeConfig(uint64, uint64) TokenTransferFeeArgs {
	return f.defaults
}

func TestInferTokenTransferFeeArgsSkipsMatchingOnchainConfig(t *testing.T) {
	onchain := TokenTransferFeeArgs{
		DestBytesOverhead: 32,
		DestGasOverhead:   90_000,
		MinFeeUSDCents:    25,
		MaxFeeUSDCents:    100,
		DeciBps:           5,
		IsEnabled:         true,
	}

	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: onchain},
		cldf.Environment{Logger: logger.Test(t)},
		1,
		2,
		TokenTransferFee{Address: "token"},
	)

	require.NoError(t, err)
	require.False(t, shouldApply)
	require.Nil(t, got)
}

func TestInferTokenTransferFeeArgsReturnsChangedConfig(t *testing.T) {
	onchain := TokenTransferFeeArgs{
		DestBytesOverhead: 32,
		DestGasOverhead:   90_000,
		MinFeeUSDCents:    25,
		MaxFeeUSDCents:    100,
		DeciBps:           5,
		IsEnabled:         true,
	}

	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: onchain},
		cldf.Environment{Logger: logger.Test(t)},
		1,
		2,
		TokenTransferFee{
			Address: "token",
			FeeArgs: UnresolvedTokenTransferFeeArgs{
				MinFeeUSDCents: utils.NewOptional(uint32(26)),
			},
		},
	)

	require.NoError(t, err)
	require.True(t, shouldApply)
	require.NotNil(t, got)
	require.Equal(t, uint32(26), got.MinFeeUSDCents)
	require.Equal(t, onchain.MaxFeeUSDCents, got.MaxFeeUSDCents)
}

func TestInferTokenTransferFeeArgsSkipsResetWhenAlreadyDisabled(t *testing.T) {
	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: TokenTransferFeeArgs{IsEnabled: false}},
		cldf.Environment{Logger: logger.Test(t)},
		1,
		2,
		TokenTransferFee{Address: "token", IsReset: true},
	)

	require.NoError(t, err)
	require.False(t, shouldApply)
	require.Nil(t, got)
}

func TestInferTokenTransferFeeArgsReturnsResetWhenEnabled(t *testing.T) {
	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{onchain: TokenTransferFeeArgs{IsEnabled: true}},
		cldf.Environment{Logger: logger.Test(t)},
		1,
		2,
		TokenTransferFee{Address: "token", IsReset: true},
	)

	require.NoError(t, err)
	require.True(t, shouldApply)
	require.Nil(t, got)
}

func TestInferTokenTransferFeeArgsReturnsOnchainReadError(t *testing.T) {
	readErr := errors.New("read failed")

	got, shouldApply, err := inferTokenTransferFeeArgs(
		fakeTokenTransferFeeReader{err: readErr},
		cldf.Environment{Logger: logger.Test(t)},
		1,
		2,
		TokenTransferFee{Address: "token"},
	)

	require.ErrorIs(t, err, readErr)
	require.False(t, shouldApply)
	require.Nil(t, got)
}
