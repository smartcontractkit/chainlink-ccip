// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package base_token_pool

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
	ErrInvalidInitPoolPermissions = &customErrorDef{
		code: 6000,
		msg:  "Pool authority does not match token mint owner",
		name: "InvalidInitPoolPermissions",
	}
	ErrUnauthorized = &customErrorDef{
		code: 6001,
		msg:  "Unauthorized",
		name: "Unauthorized",
	}
	ErrInvalidInputs = &customErrorDef{
		code: 6002,
		msg:  "Invalid inputs",
		name: "InvalidInputs",
	}
	ErrInvalidPoolCaller = &customErrorDef{
		code: 6003,
		msg:  "Caller is not ramp on router",
		name: "InvalidPoolCaller",
	}
	ErrInvalidSender = &customErrorDef{
		code: 6004,
		msg:  "Sender not allowed",
		name: "InvalidSender",
	}
	ErrInvalidSourcePoolAddress = &customErrorDef{
		code: 6005,
		msg:  "Invalid source pool address",
		name: "InvalidSourcePoolAddress",
	}
	ErrInvalidToken = &customErrorDef{
		code: 6006,
		msg:  "Invalid token",
		name: "InvalidToken",
	}
	ErrInvalidTokenAmountConversion = &customErrorDef{
		code: 6007,
		msg:  "Invalid token amount conversion",
		name: "InvalidTokenAmountConversion",
	}
	ErrRLBucketOverfilled = &customErrorDef{
		code: 6008,
		msg:  "RateLimit: bucket overfilled",
		name: "RLBucketOverfilled",
	}
	ErrRLMaxCapacityExceeded = &customErrorDef{
		code: 6009,
		msg:  "RateLimit: max capacity exceeded",
		name: "RLMaxCapacityExceeded",
	}
	ErrRLRateLimitReached = &customErrorDef{
		code: 6010,
		msg:  "RateLimit: rate limit reached",
		name: "RLRateLimitReached",
	}
	ErrRLInvalidRateLimitRate = &customErrorDef{
		code: 6011,
		msg:  "RateLimit: invalid rate limit rate",
		name: "RLInvalidRateLimitRate",
	}
	ErrRLDisabledNonZeroRateLimit = &customErrorDef{
		code: 6012,
		msg:  "RateLimit: disabled non-zero rate limit",
		name: "RLDisabledNonZeroRateLimit",
	}
	ErrLiquidityNotAccepted = &customErrorDef{
		code: 6013,
		msg:  "Liquidity not accepted",
		name: "LiquidityNotAccepted",
	}
	Errors = map[int]CustomError{
		6000: ErrInvalidInitPoolPermissions,
		6001: ErrUnauthorized,
		6002: ErrInvalidInputs,
		6003: ErrInvalidPoolCaller,
		6004: ErrInvalidSender,
		6005: ErrInvalidSourcePoolAddress,
		6006: ErrInvalidToken,
		6007: ErrInvalidTokenAmountConversion,
		6008: ErrRLBucketOverfilled,
		6009: ErrRLMaxCapacityExceeded,
		6010: ErrRLRateLimitReached,
		6011: ErrRLInvalidRateLimitRate,
		6012: ErrRLDisabledNonZeroRateLimit,
		6013: ErrLiquidityNotAccepted,
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
