package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

var _ types.RemoteCodec = &ExampleStructJSONCodec{}

type ExampleStructJSONCodec struct{}

type OnChainStruct struct {
	Aa int64
	Bb string
	Cc bool
	Dd string
	Ee int64
	Ff string
}

func (ExampleStructJSONCodec) Encode(_ context.Context, item any, _ string) ([]byte, error) {
	return json.Marshal(item)
}

func (ExampleStructJSONCodec) GetMaxEncodingSize(_ context.Context, n int, _ string) (int, error) {
	// not used in the example, and not really valid for json.
	return math.MaxInt32, nil
}

func (ExampleStructJSONCodec) Decode(_ context.Context, raw []byte, into any, _ string) error {
	err := json.Unmarshal(raw, into)
	if err != nil {
		return fmt.Errorf("%w: %w", types.ErrInvalidType, err)
	}
	return nil
}

func (ExampleStructJSONCodec) GetMaxDecodingSize(ctx context.Context, n int, _ string) (int, error) {
	// not used in the example, and not really valid for json.
	return math.MaxInt32, nil
}

func (ExampleStructJSONCodec) CreateType(_ string, _ bool) (any, error) {
	// parameters here are unused in the example, but can be used to determine what type to expect.
	// this allows remote execution to know how to decode the incoming message
	// and for [codec.NewModifierCodec] to know what type to expect for intermediate phases.
	return &OnChainStruct{}, nil
}
