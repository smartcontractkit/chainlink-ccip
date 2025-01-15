package mocks

import (
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type MockExtraDataCodec struct{}

func NewMockExtraDataCodec() *MockExtraDataCodec {
	return &MockExtraDataCodec{}
}

func (m *MockExtraDataCodec) DecodeExtraData(ExtraArgs cciptypes.Bytes, sourceChainSelector cciptypes.ChainSelector) (map[string]any, error) {
	// simply return an empty map
	return map[string]any{}, nil
}
