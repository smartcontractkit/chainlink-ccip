// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package ccip_offramp

import (
	"encoding/json"
	"errors"
	"fmt"
	ag_jsonrpc "github.com/gagliardetto/solana-go/rpc/jsonrpc"
)

var (
	_ *json.Encoder        = nil
	_ *ag_jsonrpc.RPCError = nil
	_ fmt.Formatter        = nil
	_                      = errors.ErrUnsupported
)
var (
	ErrInvalidConfigFMustBePositive = &customErrorDef{
		code: 6000,
		msg:  "Invalid config: F must be positive",
		name: "InvalidConfigFMustBePositive",
	}
	ErrInvalidConfigTooManyTransmitters = &customErrorDef{
		code: 6001,
		msg:  "Invalid config: Too many transmitters",
		name: "InvalidConfigTooManyTransmitters",
	}
	ErrInvalidConfigTooManySigners = &customErrorDef{
		code: 6002,
		msg:  "Invalid config: Too many signers",
		name: "InvalidConfigTooManySigners",
	}
	ErrInvalidConfigFIsTooHigh = &customErrorDef{
		code: 6003,
		msg:  "Invalid config: F is too high",
		name: "InvalidConfigFIsTooHigh",
	}
	ErrInvalidConfigRepeatedOracle = &customErrorDef{
		code: 6004,
		msg:  "Invalid config: Repeated oracle address",
		name: "InvalidConfigRepeatedOracle",
	}
	ErrWrongMessageLength = &customErrorDef{
		code: 6005,
		msg:  "Wrong message length",
		name: "WrongMessageLength",
	}
	ErrConfigDigestMismatch = &customErrorDef{
		code: 6006,
		msg:  "Config digest mismatch",
		name: "ConfigDigestMismatch",
	}
	ErrWrongNumberOfSignatures = &customErrorDef{
		code: 6007,
		msg:  "Wrong number signatures",
		name: "WrongNumberOfSignatures",
	}
	ErrUnauthorizedTransmitter = &customErrorDef{
		code: 6008,
		msg:  "Unauthorized transmitter",
		name: "UnauthorizedTransmitter",
	}
	ErrUnauthorizedSigner = &customErrorDef{
		code: 6009,
		msg:  "Unauthorized signer",
		name: "UnauthorizedSigner",
	}
	ErrNonUniqueSignatures = &customErrorDef{
		code: 6010,
		msg:  "Non unique signatures",
		name: "NonUniqueSignatures",
	}
	ErrOracleCannotBeZeroAddress = &customErrorDef{
		code: 6011,
		msg:  "Oracle cannot be zero address",
		name: "OracleCannotBeZeroAddress",
	}
	ErrStaticConfigCannotBeChanged = &customErrorDef{
		code: 6012,
		msg:  "Static config cannot be changed",
		name: "StaticConfigCannotBeChanged",
	}
	ErrInvalidPluginType = &customErrorDef{
		code: 6013,
		msg:  "Incorrect plugin type",
		name: "InvalidPluginType",
	}
	ErrInvalidSignature = &customErrorDef{
		code: 6014,
		msg:  "Invalid signature",
		name: "InvalidSignature",
	}
	ErrSignaturesOutOfRegistration = &customErrorDef{
		code: 6015,
		msg:  "Signatures out of registration",
		name: "SignaturesOutOfRegistration",
	}
	Errors = map[int]CustomError{
		6000: ErrInvalidConfigFMustBePositive,
		6001: ErrInvalidConfigTooManyTransmitters,
		6002: ErrInvalidConfigTooManySigners,
		6003: ErrInvalidConfigFIsTooHigh,
		6004: ErrInvalidConfigRepeatedOracle,
		6005: ErrWrongMessageLength,
		6006: ErrConfigDigestMismatch,
		6007: ErrWrongNumberOfSignatures,
		6008: ErrUnauthorizedTransmitter,
		6009: ErrUnauthorizedSigner,
		6010: ErrNonUniqueSignatures,
		6011: ErrOracleCannotBeZeroAddress,
		6012: ErrStaticConfigCannotBeChanged,
		6013: ErrInvalidPluginType,
		6014: ErrInvalidSignature,
		6015: ErrSignaturesOutOfRegistration,
	}
)

type CustomError interface {
	Code() int
	Name() string
	Error() string
}

type customErrorDef struct {
	code int
	name string
	msg  string
}

func (e *customErrorDef) Code() int {
	return e.code
}

func (e *customErrorDef) Name() string {
	return e.name
}

func (e *customErrorDef) Error() string {
	return fmt.Sprintf("%s(%d): %s", e.name, e.code, e.msg)
}

func DecodeCustomError(rpcErr error) (err error, ok bool) {
	if errCode, o := decodeErrorCode(rpcErr); o {
		if customErr, o := Errors[errCode]; o {
			err = customErr
			ok = true
			return
		}
	}
	return
}

func decodeErrorCode(rpcErr error) (errorCode int, ok bool) {
	var jErr *ag_jsonrpc.RPCError
	if errors.As(rpcErr, &jErr) && jErr.Data != nil {
		if root, o := jErr.Data.(map[string]interface{}); o {
			if rootErr, o := root["err"].(map[string]interface{}); o {
				if rootErrInstructionError, o := rootErr["InstructionError"]; o {
					if rootErrInstructionErrorItems, o := rootErrInstructionError.([]interface{}); o {
						if len(rootErrInstructionErrorItems) == 2 {
							if v, o := rootErrInstructionErrorItems[1].(map[string]interface{}); o {
								if v2, o := v["Custom"].(json.Number); o {
									if code, err := v2.Int64(); err == nil {
										ok = true
										errorCode = int(code)
									}
								} else if v2, o := v["Custom"].(float64); o {
									ok = true
									errorCode = int(v2)
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
