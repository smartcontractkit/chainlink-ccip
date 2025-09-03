package common

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// ExtraArgs version tags
	EVMExtraArgsV1Tag    = crypto.Keccak256([]byte("CCIP EVMExtraArgsV1"))[:4]
	EVMExtraArgsV2Tag    = crypto.Keccak256([]byte("CCIP EVMExtraArgsV2"))[:4]
	EVMExtraArgsV3Tag, _ = hex.DecodeString("302326cb") // TODO: determine hash preimage
)

// EVMExtraArgsV1 represents the basic extra args format
type EVMExtraArgsV1 struct {
	GasLimit *big.Int
}

// ToBytes encodes EVMExtraArgsV1 to bytes
func (e *EVMExtraArgsV1) ToBytes() []byte {
	if e == nil {
		return nil
	}
	data := make([]byte, 0, len(EVMExtraArgsV1Tag)+32)
	data = append(data, EVMExtraArgsV1Tag...)
	if e.GasLimit != nil {
		gasLimitBytes := make([]byte, 32)
		e.GasLimit.FillBytes(gasLimitBytes)
		data = append(data, gasLimitBytes...)
	}
	return data
}

// FromBytes decodes EVMExtraArgsV1 from bytes
func (e *EVMExtraArgsV1) FromBytes(data []byte) error {
	if len(data) == 0 {
		e.GasLimit = big.NewInt(200_000)
		return nil
	}
	if !bytes.HasPrefix(data, EVMExtraArgsV1Tag) {
		return fmt.Errorf("invalid EVMExtraArgsV1 tag")
	}
	data = data[len(EVMExtraArgsV1Tag):]
	if len(data) < 32 {
		return fmt.Errorf("data too short")
	}
	e.GasLimit = new(big.Int).SetBytes(data[:32])
	return nil
}

// GenericExtraArgsV2 represents the v2 extra args format with out-of-order execution
type GenericExtraArgsV2 struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
}

// ToBytes encodes GenericExtraArgsV2 to bytes
func (e *GenericExtraArgsV2) ToBytes() []byte {
	if e == nil {
		return nil
	}
	data := make([]byte, 0, len(EVMExtraArgsV2Tag)+32+1) // tag + gasLimit + bool
	data = append(data, EVMExtraArgsV2Tag...)
	if e.GasLimit != nil {
		gasLimitBytes := make([]byte, 32)
		e.GasLimit.FillBytes(gasLimitBytes)
		data = append(data, gasLimitBytes...)
	}
	if e.AllowOutOfOrderExecution {
		data = append(data, 1)
	} else {
		data = append(data, 0)
	}
	return data
}

// FromBytes decodes GenericExtraArgsV2 from bytes
func (e *GenericExtraArgsV2) FromBytes(data []byte) error {
	if len(data) == 0 {
		e.GasLimit = big.NewInt(200_000)
		e.AllowOutOfOrderExecution = false
		return nil
	}
	if !bytes.HasPrefix(data, EVMExtraArgsV2Tag) {
		return fmt.Errorf("invalid GenericExtraArgsV2 tag")
	}
	data = data[len(EVMExtraArgsV2Tag):]
	if len(data) < 33 { // 32 bytes for gas limit + 1 byte for bool
		return fmt.Errorf("data too short")
	}
	e.GasLimit = new(big.Int).SetBytes(data[:32])
	e.AllowOutOfOrderExecution = data[32] == 1
	return nil
}

// CCV represents a Cross-Chain Verifier configuration
type CCV struct {
	CCVAddress UnknownAddress
	Args       []byte
}

// EVMExtraArgsV3 represents the v3 extra args format with modular security
type EVMExtraArgsV3 struct {
	RequiredCCV       []CCV
	OptionalCCV       []CCV
	OptionalThreshold uint8
	FinalityConfig    uint32
	Executor          UnknownAddress
	ExecutorArgs      []byte
	TokenArgs         []byte
}

// ToBytes encodes EVMExtraArgsV3 to bytes
func (e *EVMExtraArgsV3) ToBytes() []byte {
	if e == nil {
		return nil
	}
	data := make([]byte, 0, 1024) // Reasonable initial capacity
	data = append(data, EVMExtraArgsV3Tag...)

	// Encode required CCVs
	data = binary.BigEndian.AppendUint16(data, uint16(len(e.RequiredCCV)))
	for _, ccv := range e.RequiredCCV {
		data = append(data, ccv.CCVAddress...)
		data = binary.BigEndian.AppendUint16(data, uint16(len(ccv.Args)))
		data = append(data, ccv.Args...)
	}

	// Encode optional CCVs
	data = binary.BigEndian.AppendUint16(data, uint16(len(e.OptionalCCV)))
	for _, ccv := range e.OptionalCCV {
		data = append(data, ccv.CCVAddress...)
		data = binary.BigEndian.AppendUint16(data, uint16(len(ccv.Args)))
		data = append(data, ccv.Args...)
	}

	// Encode remaining fields
	data = append(data, e.OptionalThreshold)
	data = binary.BigEndian.AppendUint32(data, e.FinalityConfig)
	data = append(data, e.Executor...)
	data = binary.BigEndian.AppendUint16(data, uint16(len(e.ExecutorArgs)))
	data = append(data, e.ExecutorArgs...)
	data = binary.BigEndian.AppendUint16(data, uint16(len(e.TokenArgs)))
	data = append(data, e.TokenArgs...)

	return data
}

// FromBytes decodes EVMExtraArgsV3 from bytes
func (e *EVMExtraArgsV3) FromBytes(data []byte) error {
	if len(data) == 0 {
		// Set default values
		e.RequiredCCV = nil
		e.OptionalCCV = nil
		e.OptionalThreshold = 0
		e.FinalityConfig = 0
		e.Executor = make(UnknownAddress, 20)
		args := &EVMExtraArgsV1{GasLimit: big.NewInt(200_000)}
		e.ExecutorArgs = args.ToBytes()
		e.TokenArgs = nil
		return nil
	}

	if !bytes.HasPrefix(data, EVMExtraArgsV3Tag) {
		return fmt.Errorf("invalid EVMExtraArgsV3 tag")
	}
	data = data[len(EVMExtraArgsV3Tag):]

	// Helper function to read next bytes
	readNextBytes := func(n int) ([]byte, error) {
		if len(data) < n {
			return nil, fmt.Errorf("data too short")
		}
		result := data[:n]
		data = data[n:]
		return result, nil
	}

	// Read required CCVs
	requiredCCVCount, err := readNextBytes(2)
	if err != nil {
		return fmt.Errorf("failed to read required CCV count: %w", err)
	}
	count := binary.BigEndian.Uint16(requiredCCVCount)
	e.RequiredCCV = make([]CCV, count)
	for i := uint16(0); i < count; i++ {
		addr, err := readNextBytes(20)
		if err != nil {
			return fmt.Errorf("failed to read required CCV address: %w", err)
		}
		argsLen, err := readNextBytes(2)
		if err != nil {
			return fmt.Errorf("failed to read required CCV args length: %w", err)
		}
		args, err := readNextBytes(int(binary.BigEndian.Uint16(argsLen)))
		if err != nil {
			return fmt.Errorf("failed to read required CCV args: %w", err)
		}
		e.RequiredCCV[i] = CCV{
			CCVAddress: UnknownAddress(addr),
			Args:       args,
		}
	}

	// Read optional CCVs
	optionalCCVCount, err := readNextBytes(2)
	if err != nil {
		return fmt.Errorf("failed to read optional CCV count: %w", err)
	}
	count = binary.BigEndian.Uint16(optionalCCVCount)
	e.OptionalCCV = make([]CCV, count)
	for i := uint16(0); i < count; i++ {
		addr, err := readNextBytes(20)
		if err != nil {
			return fmt.Errorf("failed to read optional CCV address: %w", err)
		}
		argsLen, err := readNextBytes(2)
		if err != nil {
			return fmt.Errorf("failed to read optional CCV args length: %w", err)
		}
		args, err := readNextBytes(int(binary.BigEndian.Uint16(argsLen)))
		if err != nil {
			return fmt.Errorf("failed to read optional CCV args: %w", err)
		}
		e.OptionalCCV[i] = CCV{
			CCVAddress: UnknownAddress(addr),
			Args:       args,
		}
	}

	// Read remaining fields
	threshold, err := readNextBytes(1)
	if err != nil {
		return fmt.Errorf("failed to read optional threshold: %w", err)
	}
	e.OptionalThreshold = threshold[0]

	finality, err := readNextBytes(4)
	if err != nil {
		return fmt.Errorf("failed to read finality config: %w", err)
	}
	e.FinalityConfig = binary.BigEndian.Uint32(finality)

	executor, err := readNextBytes(20)
	if err != nil {
		return fmt.Errorf("failed to read executor address: %w", err)
	}
	e.Executor = UnknownAddress(executor)

	execArgsLen, err := readNextBytes(2)
	if err != nil {
		return fmt.Errorf("failed to read executor args length: %w", err)
	}
	e.ExecutorArgs, err = readNextBytes(int(binary.BigEndian.Uint16(execArgsLen)))
	if err != nil {
		return fmt.Errorf("failed to read executor args: %w", err)
	}

	tokenArgsLen, err := readNextBytes(2)
	if err != nil {
		return fmt.Errorf("failed to read token args length: %w", err)
	}
	e.TokenArgs, err = readNextBytes(int(binary.BigEndian.Uint16(tokenArgsLen)))
	if err != nil {
		return fmt.Errorf("failed to read token args: %w", err)
	}

	return nil
}
