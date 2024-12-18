package contracts

import (
	ag_binary "github.com/gagliardetto/binary"
)

type Ocr3Error ag_binary.BorshEnum

// This Errors should be automatically generated by Anchor-Go but they only support one error per program
const (
	Ocr3ErrorInvalidConfigFMustBePositive Ocr3Error = iota
	Ocr3ErrorInvalidConfigTooManyTransmitters
	Ocr3ErrorInvalidConfigTooManySigners
	Ocr3ErrorInvalidConfigFIsTooHigh
	Ocr3ErrorInvalidConfigRepeatedOracle
	Ocr3ErrorWrongMessageLength
	Ocr3ErrorConfigDigestMismatch
	Ocr3ErrorWrongNumberOfSignatures
	Ocr3ErrorUnauthorizedTransmitter
	Ocr3ErrorUnauthorizedSigner
	Ocr3ErrorNonUniqueSignatures
	Ocr3ErrorOracleCannotBeZeroAddress
	Ocr3ErrorStaticConfigCannotBeChanged
	Ocr3ErrorInvalidPluginType
	Ocr3ErrorInvalidSignature
)

func (value Ocr3Error) String() string {
	switch value {
	case Ocr3ErrorInvalidConfigFMustBePositive:
		return "InvalidConfigFMustBePositive"
	case Ocr3ErrorInvalidConfigTooManyTransmitters:
		return "InvalidConfigTooManyTransmitters"
	case Ocr3ErrorInvalidConfigTooManySigners:
		return "InvalidConfigTooManySigners"
	case Ocr3ErrorInvalidConfigFIsTooHigh:
		return "InvalidConfigFIsTooHigh"
	case Ocr3ErrorInvalidConfigRepeatedOracle:
		return "InvalidConfigRepeatedOracle"
	case Ocr3ErrorWrongMessageLength:
		return "WrongMessageLength"
	case Ocr3ErrorConfigDigestMismatch:
		return "ConfigDigestMismatch"
	case Ocr3ErrorWrongNumberOfSignatures:
		return "WrongNumberOfSignatures"
	case Ocr3ErrorUnauthorizedTransmitter:
		return "UnauthorizedTransmitter"
	case Ocr3ErrorUnauthorizedSigner:
		return "UnauthorizedSigner"
	case Ocr3ErrorNonUniqueSignatures:
		return "NonUniqueSignatures"
	case Ocr3ErrorOracleCannotBeZeroAddress:
		return "OracleCannotBeZeroAddress"
	case Ocr3ErrorStaticConfigCannotBeChanged:
		return "StaticConfigCannotBeChanged"
	case Ocr3ErrorInvalidPluginType:
		return "InvalidPluginType"
	case Ocr3ErrorInvalidSignature:
		return "InvalidSignature"
	default:
		return ""
	}
}
