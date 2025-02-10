// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package fee_quoter

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
	ErrInvalidSequenceInterval = &customErrorDef{
		code: 6000,
		msg:  "The given sequence interval is invalid",
		name: "InvalidSequenceInterval",
	}
	ErrRootNotCommitted = &customErrorDef{
		code: 6001,
		msg:  "The given Merkle Root is missing",
		name: "RootNotCommitted",
	}
	ErrExistingMerkleRoot = &customErrorDef{
		code: 6002,
		msg:  "The given Merkle Root is already committed",
		name: "ExistingMerkleRoot",
	}
	ErrUnauthorized = &customErrorDef{
		code: 6003,
		msg:  "The signer is unauthorized",
		name: "Unauthorized",
	}
	ErrInvalidInputs = &customErrorDef{
		code: 6004,
		msg:  "Invalid inputs",
		name: "InvalidInputs",
	}
	ErrUnsupportedSourceChainSelector = &customErrorDef{
		code: 6005,
		msg:  "Source chain selector not supported",
		name: "UnsupportedSourceChainSelector",
	}
	ErrUnsupportedDestinationChainSelector = &customErrorDef{
		code: 6006,
		msg:  "Destination chain selector not supported",
		name: "UnsupportedDestinationChainSelector",
	}
	ErrInvalidProof = &customErrorDef{
		code: 6007,
		msg:  "Invalid Proof for Merkle Root",
		name: "InvalidProof",
	}
	ErrInvalidMessage = &customErrorDef{
		code: 6008,
		msg:  "Invalid message format",
		name: "InvalidMessage",
	}
	ErrReachedMaxSequenceNumber = &customErrorDef{
		code: 6009,
		msg:  "Reached max sequence number",
		name: "ReachedMaxSequenceNumber",
	}
	ErrManualExecutionNotAllowed = &customErrorDef{
		code: 6010,
		msg:  "Manual execution not allowed",
		name: "ManualExecutionNotAllowed",
	}
	ErrInvalidInputsTokenIndices = &customErrorDef{
		code: 6011,
		msg:  "Invalid pool account account indices",
		name: "InvalidInputsTokenIndices",
	}
	ErrInvalidInputsPoolAccounts = &customErrorDef{
		code: 6012,
		msg:  "Invalid pool accounts",
		name: "InvalidInputsPoolAccounts",
	}
	ErrInvalidInputsTokenAccounts = &customErrorDef{
		code: 6013,
		msg:  "Invalid token accounts",
		name: "InvalidInputsTokenAccounts",
	}
	ErrInvalidInputsConfigAccounts = &customErrorDef{
		code: 6014,
		msg:  "Invalid config account",
		name: "InvalidInputsConfigAccounts",
	}
	ErrInvalidInputsTokenAdminRegistryAccounts = &customErrorDef{
		code: 6015,
		msg:  "Invalid Token Admin Registry account",
		name: "InvalidInputsTokenAdminRegistryAccounts",
	}
	ErrInvalidInputsLookupTableAccounts = &customErrorDef{
		code: 6016,
		msg:  "Invalid LookupTable account",
		name: "InvalidInputsLookupTableAccounts",
	}
	ErrInvalidInputsLookupTableAccountWritable = &customErrorDef{
		code: 6017,
		msg:  "Invalid LookupTable account writable access",
		name: "InvalidInputsLookupTableAccountWritable",
	}
	ErrInvalidInputsTokenAmount = &customErrorDef{
		code: 6018,
		msg:  "Cannot send zero tokens",
		name: "InvalidInputsTokenAmount",
	}
	ErrOfframpReleaseMintBalanceMismatch = &customErrorDef{
		code: 6019,
		msg:  "Release or mint balance mismatch",
		name: "OfframpReleaseMintBalanceMismatch",
	}
	ErrOfframpInvalidDataLength = &customErrorDef{
		code: 6020,
		msg:  "Invalid data length",
		name: "OfframpInvalidDataLength",
	}
	ErrStaleCommitReport = &customErrorDef{
		code: 6021,
		msg:  "Stale commit report",
		name: "StaleCommitReport",
	}
	ErrDestinationChainDisabled = &customErrorDef{
		code: 6022,
		msg:  "Destination chain disabled",
		name: "DestinationChainDisabled",
	}
	ErrFeeTokenDisabled = &customErrorDef{
		code: 6023,
		msg:  "Fee token disabled",
		name: "FeeTokenDisabled",
	}
	ErrMessageTooLarge = &customErrorDef{
		code: 6024,
		msg:  "Message exceeds maximum data size",
		name: "MessageTooLarge",
	}
	ErrUnsupportedNumberOfTokens = &customErrorDef{
		code: 6025,
		msg:  "Message contains an unsupported number of tokens",
		name: "UnsupportedNumberOfTokens",
	}
	ErrUnsupportedChainFamilySelector = &customErrorDef{
		code: 6026,
		msg:  "Chain family selector not supported",
		name: "UnsupportedChainFamilySelector",
	}
	ErrInvalidEVMAddress = &customErrorDef{
		code: 6027,
		msg:  "Invalid EVM address",
		name: "InvalidEVMAddress",
	}
	ErrInvalidEncoding = &customErrorDef{
		code: 6028,
		msg:  "Invalid encoding",
		name: "InvalidEncoding",
	}
	ErrInvalidInputsAtaAddress = &customErrorDef{
		code: 6029,
		msg:  "Invalid Associated Token Account address",
		name: "InvalidInputsAtaAddress",
	}
	ErrInvalidInputsAtaWritable = &customErrorDef{
		code: 6030,
		msg:  "Invalid Associated Token Account writable flag",
		name: "InvalidInputsAtaWritable",
	}
	ErrInvalidTokenPrice = &customErrorDef{
		code: 6031,
		msg:  "Invalid token price",
		name: "InvalidTokenPrice",
	}
	ErrStaleGasPrice = &customErrorDef{
		code: 6032,
		msg:  "Stale gas price",
		name: "StaleGasPrice",
	}
	ErrInsufficientLamports = &customErrorDef{
		code: 6033,
		msg:  "Insufficient lamports",
		name: "InsufficientLamports",
	}
	ErrInsufficientFunds = &customErrorDef{
		code: 6034,
		msg:  "Insufficient funds",
		name: "InsufficientFunds",
	}
	ErrUnsupportedToken = &customErrorDef{
		code: 6035,
		msg:  "Unsupported token",
		name: "UnsupportedToken",
	}
	ErrInvalidInputsMissingTokenConfig = &customErrorDef{
		code: 6036,
		msg:  "Inputs are missing token configuration",
		name: "InvalidInputsMissingTokenConfig",
	}
	ErrMessageFeeTooHigh = &customErrorDef{
		code: 6037,
		msg:  "Message fee is too high",
		name: "MessageFeeTooHigh",
	}
	ErrSourceTokenDataTooLarge = &customErrorDef{
		code: 6038,
		msg:  "Source token data is too large",
		name: "SourceTokenDataTooLarge",
	}
	ErrMessageGasLimitTooHigh = &customErrorDef{
		code: 6039,
		msg:  "Message gas limit too high",
		name: "MessageGasLimitTooHigh",
	}
	ErrExtraArgOutOfOrderExecutionMustBeTrue = &customErrorDef{
		code: 6040,
		msg:  "Extra arg out of order execution must be true",
		name: "ExtraArgOutOfOrderExecutionMustBeTrue",
	}
	ErrInvalidExtraArgsTag = &customErrorDef{
		code: 6041,
		msg:  "Invalid extra args tag",
		name: "InvalidExtraArgsTag",
	}
	ErrInvalidChainFamilySelector = &customErrorDef{
		code: 6042,
		msg:  "Invalid chain family selector",
		name: "InvalidChainFamilySelector",
	}
	ErrInvalidTokenReceiver = &customErrorDef{
		code: 6043,
		msg:  "Invalid token receiver",
		name: "InvalidTokenReceiver",
	}
	ErrInvalidSVMAddress = &customErrorDef{
		code: 6044,
		msg:  "Invalid SVM address",
		name: "InvalidSVMAddress",
	}
	ErrUnauthorizedPriceUpdater = &customErrorDef{
		code: 6045,
		msg:  "The caller is not an authorized price updater",
		name: "UnauthorizedPriceUpdater",
	}
	Errors = map[int]CustomError{
		6000: ErrInvalidSequenceInterval,
		6001: ErrRootNotCommitted,
		6002: ErrExistingMerkleRoot,
		6003: ErrUnauthorized,
		6004: ErrInvalidInputs,
		6005: ErrUnsupportedSourceChainSelector,
		6006: ErrUnsupportedDestinationChainSelector,
		6007: ErrInvalidProof,
		6008: ErrInvalidMessage,
		6009: ErrReachedMaxSequenceNumber,
		6010: ErrManualExecutionNotAllowed,
		6011: ErrInvalidInputsTokenIndices,
		6012: ErrInvalidInputsPoolAccounts,
		6013: ErrInvalidInputsTokenAccounts,
		6014: ErrInvalidInputsConfigAccounts,
		6015: ErrInvalidInputsTokenAdminRegistryAccounts,
		6016: ErrInvalidInputsLookupTableAccounts,
		6017: ErrInvalidInputsLookupTableAccountWritable,
		6018: ErrInvalidInputsTokenAmount,
		6019: ErrOfframpReleaseMintBalanceMismatch,
		6020: ErrOfframpInvalidDataLength,
		6021: ErrStaleCommitReport,
		6022: ErrDestinationChainDisabled,
		6023: ErrFeeTokenDisabled,
		6024: ErrMessageTooLarge,
		6025: ErrUnsupportedNumberOfTokens,
		6026: ErrUnsupportedChainFamilySelector,
		6027: ErrInvalidEVMAddress,
		6028: ErrInvalidEncoding,
		6029: ErrInvalidInputsAtaAddress,
		6030: ErrInvalidInputsAtaWritable,
		6031: ErrInvalidTokenPrice,
		6032: ErrStaleGasPrice,
		6033: ErrInsufficientLamports,
		6034: ErrInsufficientFunds,
		6035: ErrUnsupportedToken,
		6036: ErrInvalidInputsMissingTokenConfig,
		6037: ErrMessageFeeTooHigh,
		6038: ErrSourceTokenDataTooLarge,
		6039: ErrMessageGasLimitTooHigh,
		6040: ErrExtraArgOutOfOrderExecutionMustBeTrue,
		6041: ErrInvalidExtraArgsTag,
		6042: ErrInvalidChainFamilySelector,
		6043: ErrInvalidTokenReceiver,
		6044: ErrInvalidSVMAddress,
		6045: ErrUnauthorizedPriceUpdater,
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
