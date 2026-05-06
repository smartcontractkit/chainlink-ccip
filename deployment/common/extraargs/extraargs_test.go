package extraargs

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializeEVMExtraArgsV1(t *testing.T) {
	data := ClientEVMExtraArgsV1{GasLimit: big.NewInt(200_000)}
	result, err := SerializeEVMExtraArgsV1(data)
	require.NoError(t, err)

	assert.Equal(t, EVMExtraArgsV1Tag, result[:4], "tag mismatch")
	assert.Len(t, result, 4+32, "expected 4-byte tag + 32-byte uint256")
}

func TestSerializeClientGenericExtraArgsV2(t *testing.T) {
	data := ClientGenericExtraArgsV2{
		GasLimit:                 big.NewInt(300_000),
		AllowOutOfOrderExecution: true,
	}
	result, err := SerializeClientGenericExtraArgsV2(data)
	require.NoError(t, err)

	assert.Equal(t, GenericExtraArgsV2Tag, result[:4], "tag mismatch")
	// tuple(uint256, bool) = 32 + 32 = 64 bytes
	assert.Len(t, result, 4+64, "expected 4-byte tag + 64-byte encoded tuple")
}

func TestSerializeClientSVMExtraArgsV1(t *testing.T) {
	data := ClientSVMExtraArgsV1{
		ComputeUnits:             100_000,
		AccountIsWritableBitmap:  0xFF,
		AllowOutOfOrderExecution: false,
		TokenReceiver:            [32]byte{0x01},
		Accounts:                 [][32]byte{{0x02}, {0x03}},
	}
	result, err := SerializeClientSVMExtraArgsV1(data)
	require.NoError(t, err)

	assert.Equal(t, SVMExtraArgsV1Tag, result[:4], "tag mismatch")
	assert.Greater(t, len(result), 4, "result should have more than just the tag")
}

func TestSerializeClientSUIExtraArgsV1(t *testing.T) {
	data := ClientSuiExtraArgsV1{
		GasLimit:                 big.NewInt(500_000),
		AllowOutOfOrderExecution: true,
		TokenReceiver:            [32]byte{0xAA},
		ReceiverObjectIds:        [][32]byte{{0xBB}},
	}
	result, err := SerializeClientSUIExtraArgsV1(data)
	require.NoError(t, err)

	assert.Equal(t, SUIExtraArgsV1Tag, result[:4], "tag mismatch")
	assert.Greater(t, len(result), 4, "result should have more than just the tag")
}

func TestSerializeEVMExtraArgsV1_Deterministic(t *testing.T) {
	data := ClientEVMExtraArgsV1{GasLimit: big.NewInt(200_000)}
	r1, err := SerializeEVMExtraArgsV1(data)
	require.NoError(t, err)
	r2, err := SerializeEVMExtraArgsV1(data)
	require.NoError(t, err)
	assert.Equal(t, r1, r2, "serialization should be deterministic")
}

func TestSerializeClientGenericExtraArgsV2_OOOFalse(t *testing.T) {
	dataTrue := ClientGenericExtraArgsV2{
		GasLimit:                 big.NewInt(300_000),
		AllowOutOfOrderExecution: true,
	}
	dataFalse := ClientGenericExtraArgsV2{
		GasLimit:                 big.NewInt(300_000),
		AllowOutOfOrderExecution: false,
	}
	rTrue, err := SerializeClientGenericExtraArgsV2(dataTrue)
	require.NoError(t, err)
	rFalse, err := SerializeClientGenericExtraArgsV2(dataFalse)
	require.NoError(t, err)
	assert.NotEqual(t, rTrue, rFalse, "different bool values should produce different output")
}

func TestSerializeClientSVMExtraArgsV1_EmptyAccounts(t *testing.T) {
	data := ClientSVMExtraArgsV1{
		ComputeUnits:             50_000,
		AccountIsWritableBitmap:  0,
		AllowOutOfOrderExecution: true,
		TokenReceiver:            [32]byte{},
		Accounts:                 nil,
	}
	result, err := SerializeClientSVMExtraArgsV1(data)
	require.NoError(t, err)
	assert.Equal(t, SVMExtraArgsV1Tag, result[:4])
}
