package testadapters

import (
	"fmt"
	"math/big"
)

// Extra args follow the same split as chainlink-ccv (build/devenv/cciptestinterfaces/extra_args.go): the
// destination adapter decides what to send (ExtraArgsDataSource) and the source adapter decides how to encode
// it (EVMExtraArgsV2, SVMExtraArgsV1). Each family's encoding stays in that family's adapter, so no adapter
// needs another family's SDK or bindings to build extra args.

type ExtraArgsDataProvider interface {
	IsExtraArgsDataProvider()
}

var (
	_ ExtraArgsDataProvider = (*EVMExtraArgsV2Data)(nil)
	_ ExtraArgsDataProvider = (*SVMExtraArgsV1Data)(nil)
)

// EVMExtraArgsV2Data is the data for GenericExtraArgsV2 (tag 0x181dcf10 = bytes4(keccak256("CCIP EVMExtraArgsV2"))).
// Despite the name it is also what TON destinations expect.
type EVMExtraArgsV2Data struct {
	GasLimit                 *big.Int
	AllowOutOfOrderExecution bool
}

func (EVMExtraArgsV2Data) IsExtraArgsDataProvider() {}

type SVMExtraArgsV1Data struct {
	ComputeUnits             uint32
	AccountIsWritableBitmap  uint64
	AllowOutOfOrderExecution bool
	TokenReceiver            [32]byte
	Accounts                 [][32]byte
}

func (SVMExtraArgsV1Data) IsExtraArgsDataProvider() {}

// ExtraArgsDataSource is implemented by destination adapters. Defaults must depend only on the destination,
// never on the source family, so every source sends the same values.
type ExtraArgsDataSource interface {
	ExtraArgsData(receiver []byte, opts ...ExtraArgOpt) (ExtraArgsDataProvider, error)
}

// EVMExtraArgsV2 is implemented by source adapters that can send to EVM and TON chains.
type EVMExtraArgsV2 interface {
	BuildEVMExtraArgsV2(provider ExtraArgsDataProvider) ([]byte, error)
}

type SVMExtraArgsV1 interface {
	BuildSVMExtraArgsV1(provider ExtraArgsDataProvider) ([]byte, error)
}

// ExtraArgsSource is all BuildExtraArgs needs from the source adapter; how it encodes is discovered through
// the optional EVMExtraArgsV2 / SVMExtraArgsV1 interfaces.
type ExtraArgsSource interface {
	Family() string
}

// BuildExtraArgs falls back to dest.GetExtraArgs when either side does not implement its part, so families
// can migrate one at a time.
// TODO: remove the fallback (and GetExtraArgs) once Solana implements the source side. Until then
// Solana->EVM intentionally keeps sending empty extra args through it.
func BuildExtraArgs(src ExtraArgsSource, dest TestAdapterForFamily, receiver []byte, opts ...ExtraArgOpt) ([]byte, error) {
	dataSource, ok := dest.(ExtraArgsDataSource)
	if !ok {
		return dest.GetExtraArgs(receiver, src.Family(), opts...)
	}
	data, err := dataSource.ExtraArgsData(receiver, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to get extra args data from destination: %w", err)
	}

	switch data.(type) {
	case EVMExtraArgsV2Data:
		if enc, ok := src.(EVMExtraArgsV2); ok {
			return enc.BuildEVMExtraArgsV2(data)
		}
	case SVMExtraArgsV1Data:
		if enc, ok := src.(SVMExtraArgsV1); ok {
			return enc.BuildSVMExtraArgsV1(data)
		}
	default:
		return nil, fmt.Errorf("unsupported extra args data type %T", data)
	}
	return dest.GetExtraArgs(receiver, src.Family(), opts...)
}

// ParseGasLimitAndOOOExtraArgs ignores other options rather than failing, because callers pass the same
// options to every destination (e.g. the lane sanity hook always sets finality, which EVM does not use).
func ParseGasLimitAndOOOExtraArgs(gasLimit *big.Int, allowOOO bool, opts ...ExtraArgOpt) (*big.Int, bool, error) {
	for _, opt := range opts {
		switch opt.Name {
		case ExtraArgGasLimit:
			v, ok := opt.Value.(*big.Int)
			if !ok {
				return nil, false, fmt.Errorf("extra arg %q: expected *big.Int, got %T", opt.Name, opt.Value)
			}
			gasLimit = v
		case ExtraArgOOO:
			v, ok := opt.Value.(bool)
			if !ok {
				return nil, false, fmt.Errorf("extra arg %q: expected bool, got %T", opt.Name, opt.Value)
			}
			allowOOO = v
		}
	}
	return gasLimit, allowOOO, nil
}
