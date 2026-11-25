package testadapter

import (
	"math"
	"math/big"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/common/extraargs"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

func TestSVMAdapterExtraArgsData(t *testing.T) {
	a := &SVMAdapter{}
	receiver := solana.NewWallet().PublicKey()
	counterPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("counter")}, receiver)
	require.NoError(t, err)
	execConfigPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, receiver)
	require.NoError(t, err)

	t.Run("defaults", func(t *testing.T) {
		data, err := a.ExtraArgsData(receiver.Bytes())
		require.NoError(t, err)
		assert.Equal(t, testadapters.SVMExtraArgsV1Data{
			ComputeUnits:             100_000,
			AccountIsWritableBitmap:  0b11, // accounts 0 and 1 are writable
			AllowOutOfOrderExecution: true,
			TokenReceiver:            receiver,
			Accounts:                 [][32]byte{execConfigPDA, counterPDA, solana.SystemProgramID},
		}, data)
	})

	t.Run("overrides", func(t *testing.T) {
		data, err := a.ExtraArgsData(receiver.Bytes(),
			testadapters.NewGasLimitExtraArg(big.NewInt(5)),
			testadapters.NewOutOfOrderExtraArg(false),
		)
		require.NoError(t, err)
		svmData := data.(testadapters.SVMExtraArgsV1Data)
		assert.Equal(t, uint32(5), svmData.ComputeUnits)
		assert.False(t, svmData.AllowOutOfOrderExecution)
	})

	t.Run("compute units must fit in uint32", func(t *testing.T) {
		_, err := a.ExtraArgsData(receiver.Bytes(), testadapters.NewGasLimitExtraArg(big.NewInt(math.MaxUint32+1)))
		require.ErrorContains(t, err, "ComputeUnits is larger than uint32")
		_, err = a.ExtraArgsData(receiver.Bytes(), testadapters.NewGasLimitExtraArg(big.NewInt(math.MaxInt64)))
		require.ErrorContains(t, err, "ComputeUnits is larger than uint32")
	})
}

func TestSVMAdapterGetExtraArgsFromEVM(t *testing.T) {
	a := &SVMAdapter{}
	receiver := solana.NewWallet().PublicKey()

	provider, err := a.ExtraArgsData(receiver.Bytes())
	require.NoError(t, err)
	data := provider.(testadapters.SVMExtraArgsV1Data)
	want, err := extraargs.SerializeClientSVMExtraArgsV1(extraargs.ClientSVMExtraArgsV1{
		ComputeUnits:             data.ComputeUnits,
		AccountIsWritableBitmap:  data.AccountIsWritableBitmap,
		AllowOutOfOrderExecution: data.AllowOutOfOrderExecution,
		TokenReceiver:            data.TokenReceiver,
		Accounts:                 data.Accounts,
	})
	require.NoError(t, err)

	got, err := a.GetExtraArgs(receiver.Bytes(), chain_selectors.FamilyEVM)
	require.NoError(t, err)
	assert.Equal(t, want, got)
}
