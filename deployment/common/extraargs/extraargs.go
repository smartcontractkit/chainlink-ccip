package extraargs

import "math/big"

// Tag constants matching Solidity Client.sol selectors.
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

// TODO: implement without EVM contract/gobinding dependency using
// go-ethereum/accounts/abi with hardcoded ABI strings. The encoding
// is standard Solidity ABI: tag (4 bytes) + abi.encode(struct).

func SerializeEVMExtraArgsV1(data ClientEVMExtraArgsV1) ([]byte, error) {
	panic("extraargs: SerializeEVMExtraArgsV1 not yet implemented; use chains/evm/deployment/common.SerializeEVMExtraArgsV1")
}

func SerializeClientGenericExtraArgsV2(data ClientGenericExtraArgsV2) ([]byte, error) {
	panic("extraargs: SerializeClientGenericExtraArgsV2 not yet implemented; use chains/evm/deployment/common.SerializeClientGenericExtraArgsV2")
}

func SerializeClientSVMExtraArgsV1(data ClientSVMExtraArgsV1) ([]byte, error) {
	panic("extraargs: SerializeClientSVMExtraArgsV1 not yet implemented; use chains/evm/deployment/common.SerializeClientSVMExtraArgsV1")
}

func SerializeClientSUIExtraArgsV1(data ClientSuiExtraArgsV1) ([]byte, error) {
	panic("extraargs: SerializeClientSUIExtraArgsV1 not yet implemented; use chains/evm/deployment/common.SerializeClientSUIExtraArgsV1")
}
