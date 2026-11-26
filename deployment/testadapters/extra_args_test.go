package testadapters

import (
	"errors"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fakeDest struct {
	fallbackCalls int
	fallbackSrc   string
	fallbackOpts  []ExtraArgOpt
}

func (d *fakeDest) CCIPReceiver() []byte { return []byte("receiver") }

func (d *fakeDest) GetExtraArgs(_ []byte, sourceFamily string, opts ...ExtraArgOpt) ([]byte, error) {
	d.fallbackCalls++
	d.fallbackSrc = sourceFamily
	d.fallbackOpts = opts
	return []byte("fallback"), nil
}

type fakeDataDest struct {
	fakeDest
	data ExtraArgsDataProvider
	err  error
}

func (d *fakeDataDest) ExtraArgsData(_ []byte, opts ...ExtraArgOpt) (ExtraArgsDataProvider, error) {
	if d.err != nil {
		return nil, d.err
	}
	gasLimit, ooo, err := ParseGasLimitAndOOOExtraArgs(big.NewInt(1), true, opts...)
	if err != nil {
		return nil, err
	}
	if d.data != nil {
		return d.data, nil
	}
	return EVMExtraArgsV2Data{GasLimit: gasLimit, AllowOutOfOrderExecution: ooo}, nil
}

type fakeSrc struct{}

func (fakeSrc) Family() string { return "src-family" }

type fakeEVMV2Src struct {
	fakeSrc
	got ExtraArgsDataProvider
}

func (s *fakeEVMV2Src) BuildEVMExtraArgsV2(provider ExtraArgsDataProvider) ([]byte, error) {
	s.got = provider
	return []byte("evm-v2"), nil
}

type fakeSVMV1Src struct {
	fakeSrc
}

func (fakeSVMV1Src) BuildSVMExtraArgsV1(ExtraArgsDataProvider) ([]byte, error) {
	return []byte("svm-v1"), nil
}

func TestBuildExtraArgs(t *testing.T) {
	gasOpt := NewGasLimitExtraArg(big.NewInt(42))
	oooOpt := NewOutOfOrderExtraArg(false)

	t.Run("dest without data source falls back to GetExtraArgs", func(t *testing.T) {
		dest := &fakeDest{}
		got, err := BuildExtraArgs(&fakeEVMV2Src{}, dest, []byte("r"), gasOpt)
		require.NoError(t, err)
		assert.Equal(t, []byte("fallback"), got)
		assert.Equal(t, 1, dest.fallbackCalls)
		assert.Equal(t, "src-family", dest.fallbackSrc)
		assert.Equal(t, []ExtraArgOpt{gasOpt}, dest.fallbackOpts)
	})

	t.Run("source encodes destination data", func(t *testing.T) {
		dest := &fakeDataDest{}
		src := &fakeEVMV2Src{}
		got, err := BuildExtraArgs(src, dest, []byte("r"), gasOpt, oooOpt)
		require.NoError(t, err)
		assert.Equal(t, []byte("evm-v2"), got)
		assert.Equal(t, EVMExtraArgsV2Data{GasLimit: big.NewInt(42), AllowOutOfOrderExecution: false}, src.got)
		assert.Zero(t, dest.fallbackCalls)
	})

	t.Run("SVM data goes to the SVM encoder", func(t *testing.T) {
		dest := &fakeDataDest{data: SVMExtraArgsV1Data{ComputeUnits: 1}}
		got, err := BuildExtraArgs(fakeSVMV1Src{}, dest, []byte("r"))
		require.NoError(t, err)
		assert.Equal(t, []byte("svm-v1"), got)
	})

	t.Run("source without matching encoder falls back to GetExtraArgs", func(t *testing.T) {
		dest := &fakeDataDest{data: SVMExtraArgsV1Data{ComputeUnits: 1}}
		got, err := BuildExtraArgs(&fakeEVMV2Src{}, dest, []byte("r"))
		require.NoError(t, err)
		assert.Equal(t, []byte("fallback"), got)
		assert.Equal(t, 1, dest.fallbackCalls)
	})

	t.Run("source without any encoder falls back to GetExtraArgs", func(t *testing.T) {
		dest := &fakeDataDest{}
		got, err := BuildExtraArgs(fakeSrc{}, dest, []byte("r"))
		require.NoError(t, err)
		assert.Equal(t, []byte("fallback"), got)
	})

	t.Run("destination data error is returned without fallback", func(t *testing.T) {
		dest := &fakeDataDest{err: errors.New("boom")}
		_, err := BuildExtraArgs(&fakeEVMV2Src{}, dest, []byte("r"))
		require.ErrorContains(t, err, "boom")
		assert.Zero(t, dest.fallbackCalls)
	})

	t.Run("unknown destination data type is an error", func(t *testing.T) {
		dest := &fakeDataDest{data: unknownData{}}
		_, err := BuildExtraArgs(&fakeEVMV2Src{}, dest, []byte("r"))
		require.ErrorContains(t, err, "unsupported extra args data type")
	})
}

var _ ExtraArgsDataProvider = (*unknownData)(nil)

type unknownData struct{}

func (unknownData) IsExtraArgsDataProvider() {}

func TestParseGasLimitAndOOOExtraArgs(t *testing.T) {
	defaultGas := big.NewInt(100)

	t.Run("defaults", func(t *testing.T) {
		gas, ooo, err := ParseGasLimitAndOOOExtraArgs(defaultGas, true)
		require.NoError(t, err)
		assert.Equal(t, defaultGas, gas)
		assert.True(t, ooo)
	})

	t.Run("overrides, ignoring unsupported options", func(t *testing.T) {
		gas, ooo, err := ParseGasLimitAndOOOExtraArgs(defaultGas, true,
			NewGasLimitExtraArg(big.NewInt(7)),
			NewOutOfOrderExtraArg(false),
			NewFinalityExtraArg(1),
		)
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(7), gas)
		assert.False(t, ooo)
	})

	t.Run("wrong gas limit type", func(t *testing.T) {
		_, _, err := ParseGasLimitAndOOOExtraArgs(defaultGas, true, ExtraArgOpt{Name: ExtraArgGasLimit, Value: uint64(1)})
		require.ErrorContains(t, err, "expected *big.Int")
	})

	t.Run("wrong OOO type", func(t *testing.T) {
		_, _, err := ParseGasLimitAndOOOExtraArgs(defaultGas, true, ExtraArgOpt{Name: ExtraArgOOO, Value: "false"})
		require.ErrorContains(t, err, "expected bool")
	})
}
