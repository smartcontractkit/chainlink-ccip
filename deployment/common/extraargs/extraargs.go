package extraargs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	EVMExtraArgsV1Tag     = []byte{0x97, 0xa6, 0x57, 0xc9}
	GenericExtraArgsV2Tag = []byte{0x18, 0x1d, 0xcf, 0x10}
	SVMExtraArgsV1Tag     = []byte{0x1f, 0x3b, 0x3a, 0xba}
	SUIExtraArgsV1Tag     = []byte{0x21, 0xea, 0x4c, 0xa9}
)

type ClientEVMExtraArgsV1 struct {
	GasLimit *big.Int
}

type ClientGenericExtraArgsV2 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
}

type ClientSVMExtraArgsV1 struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	Accounts                 [][32]byte
}

type ClientSuiExtraArgsV1 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	ReceiverObjectIds        [][32]byte
}

var (
	evmExtraArgsV1ABI = abi.Arguments{
		{Type: mustABITupleType([]abi.ArgumentMarshaling{
			{Name: "gasLimit", Type: "uint256"},
		})},
	}

	genericExtraArgsV2ABI = abi.Arguments{
		{Type: mustABITupleType([]abi.ArgumentMarshaling{
			{Name: "gasLimit", Type: "uint256"},
			{Name: "allowOutOfOrderExecution", Type: "bool"},
		})},
	}

	svmExtraArgsV1ABI = abi.Arguments{
		{Type: mustABITupleType([]abi.ArgumentMarshaling{
			{Name: "computeUnits", Type: "uint32"},
			{Name: "accountIsWritableBitmap", Type: "uint64"},
			{Name: "allowOutOfOrderExecution", Type: "bool"},
			{Name: "tokenReceiver", Type: "bytes32"},
			{Name: "accounts", Type: "bytes32[]"},
		})},
	}

	suiExtraArgsV1ABI = abi.Arguments{
		{Type: mustABITupleType([]abi.ArgumentMarshaling{
			{Name: "gasLimit", Type: "uint256"},
			{Name: "allowOutOfOrderExecution", Type: "bool"},
			{Name: "tokenReceiver", Type: "bytes32"},
			{Name: "receiverObjectIds", Type: "bytes32[]"},
		})},
	}
)

func mustABITupleType(fields []abi.ArgumentMarshaling) abi.Type {
	t, err := abi.NewType("tuple", "", fields)
	if err != nil {
		panic("extraargs: bad ABI type: " + err.Error())
	}
	return t
}

func serializeExtraArgs(tag []byte, args abi.Arguments, values ...any) ([]byte, error) {
	encoded, err := args.Pack(values...)
	if err != nil {
		return nil, err
	}
	return append(tag, encoded...), nil
}

func SerializeEVMExtraArgsV1(data ClientEVMExtraArgsV1) ([]byte, error) {
	return serializeExtraArgs(EVMExtraArgsV1Tag, evmExtraArgsV1ABI, struct {
		GasLimit *big.Int
	}{
		GasLimit: data.GasLimit,
	})
}

func SerializeClientGenericExtraArgsV2(data ClientGenericExtraArgsV2) ([]byte, error) {
	return serializeExtraArgs(GenericExtraArgsV2Tag, genericExtraArgsV2ABI, struct {
		GasLimit                 *big.Int
		AllowOutOfOrderExecution bool
	}{
		GasLimit:                 data.GasLimit,
		AllowOutOfOrderExecution: data.AllowOutOfOrderExecution,
	})
}

func SerializeClientSVMExtraArgsV1(data ClientSVMExtraArgsV1) ([]byte, error) {
	return serializeExtraArgs(SVMExtraArgsV1Tag, svmExtraArgsV1ABI, struct {
		ComputeUnits             uint32
		AccountIsWritableBitmap  uint64
		AllowOutOfOrderExecution bool
		TokenReceiver            [32]byte
		Accounts                 [][32]byte
	}{
		ComputeUnits:             data.ComputeUnits,
		AccountIsWritableBitmap:  data.AccountIsWritableBitmap,
		AllowOutOfOrderExecution: data.AllowOutOfOrderExecution,
		TokenReceiver:            data.TokenReceiver,
		Accounts:                 data.Accounts,
	})
}

func SerializeClientSUIExtraArgsV1(data ClientSuiExtraArgsV1) ([]byte, error) {
	return serializeExtraArgs(SUIExtraArgsV1Tag, suiExtraArgsV1ABI, struct {
		GasLimit                 *big.Int
		AllowOutOfOrderExecution bool
		TokenReceiver            [32]byte
		ReceiverObjectIds        [][32]byte
	}{
		GasLimit:                 data.GasLimit,
		AllowOutOfOrderExecution: data.AllowOutOfOrderExecution,
		TokenReceiver:            data.TokenReceiver,
		ReceiverObjectIds:        data.ReceiverObjectIds,
	})
}
