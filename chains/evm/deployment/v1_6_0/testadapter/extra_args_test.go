package testadapter

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/common/extraargs"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

// GenericExtraArgsV2 tag || abi.encode(uint256 100_000, bool true).
const defaultEVMToEVMExtraArgsHex = "181dcf10" +
	"00000000000000000000000000000000000000000000000000000000000186a0" +
	"0000000000000000000000000000000000000000000000000000000000000001"

func TestEVMAdapterExtraArgsData(t *testing.T) {
	a := &EVMAdapter{}

	t.Run("defaults", func(t *testing.T) {
		data, err := a.ExtraArgsData(nil)
		require.NoError(t, err)
		assert.Equal(t, testadapters.EVMExtraArgsV2Data{GasLimit: big.NewInt(100_000), AllowOutOfOrderExecution: true}, data)
	})

	t.Run("overrides", func(t *testing.T) {
		data, err := a.ExtraArgsData(nil,
			testadapters.NewGasLimitExtraArg(big.NewInt(1)),
			testadapters.NewOutOfOrderExtraArg(false),
			testadapters.NewFinalityExtraArg(1), // not used by EVM destinations
		)
		require.NoError(t, err)
		assert.Equal(t, testadapters.EVMExtraArgsV2Data{GasLimit: big.NewInt(1), AllowOutOfOrderExecution: false}, data)
	})
}

func TestEVMAdapterBuildEVMExtraArgsV2(t *testing.T) {
	a := &EVMAdapter{}

	got, err := a.BuildEVMExtraArgsV2(testadapters.EVMExtraArgsV2Data{GasLimit: big.NewInt(100_000), AllowOutOfOrderExecution: true})
	require.NoError(t, err)
	assert.Equal(t, defaultEVMToEVMExtraArgsHex, hex.EncodeToString(got))

	_, err = a.BuildEVMExtraArgsV2(testadapters.SVMExtraArgsV1Data{})
	require.ErrorContains(t, err, "provider is not a EVMExtraArgsV2Data")
}

func TestEVMAdapterBuildSVMExtraArgsV1(t *testing.T) {
	a := &EVMAdapter{}
	data := testadapters.SVMExtraArgsV1Data{
		ComputeUnits:             100_000,
		AccountIsWritableBitmap:  3,
		AllowOutOfOrderExecution: true,
		TokenReceiver:            [32]byte{1},
		Accounts:                 [][32]byte{{2}, {3}},
	}

	got, err := a.BuildSVMExtraArgsV1(data)
	require.NoError(t, err)
	want, err := extraargs.SerializeClientSVMExtraArgsV1(extraargs.ClientSVMExtraArgsV1{
		ComputeUnits:             data.ComputeUnits,
		AccountIsWritableBitmap:  data.AccountIsWritableBitmap,
		AllowOutOfOrderExecution: data.AllowOutOfOrderExecution,
		TokenReceiver:            data.TokenReceiver,
		Accounts:                 data.Accounts,
	})
	require.NoError(t, err)
	assert.Equal(t, want, got)

	_, err = a.BuildSVMExtraArgsV1(testadapters.EVMExtraArgsV2Data{})
	require.ErrorContains(t, err, "provider is not a SVMExtraArgsV1Data")
}

func TestEVMToEVMExtraArgs(t *testing.T) {
	src, dest := &EVMAdapter{}, &EVMAdapter{}

	viaBuild, err := testadapters.BuildExtraArgs(src, dest, nil)
	require.NoError(t, err)
	viaLegacy, err := dest.GetExtraArgs(nil, chain_selectors.FamilyEVM)
	require.NoError(t, err)

	assert.Equal(t, defaultEVMToEVMExtraArgsHex, hex.EncodeToString(viaBuild))
	assert.Equal(t, viaBuild, viaLegacy)
}

// Stands in for the Solana adapter, which lives in another module.
type fakeSolanaSrc struct{}

func (fakeSolanaSrc) Family() string { return chain_selectors.FamilySolana }

// Solana->EVM intentionally keeps sending empty extra args until the Solana adapter implements the source side.
func TestSolanaToEVMExtraArgsFallback(t *testing.T) {
	got, err := testadapters.BuildExtraArgs(fakeSolanaSrc{}, &EVMAdapter{}, nil)
	require.NoError(t, err)
	assert.Empty(t, got)
}
