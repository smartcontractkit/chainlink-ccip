package testhelpers

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	TokenDataEncoderInstance cciptypes.TokenDataEncoder = TokenDataEncoder{}
	USDCEncoder                                         = TokenDataEncoderInstance.EncodeUSDC
)

// TokenDataEncoder is a fake encoder that just returns the attestation as is, it's used only for tests to verify
// if proper attestation is returned. Production implementation is going to be chain-specific
type TokenDataEncoder struct{}

func (t TokenDataEncoder) EncodeUSDC(
	_ context.Context,
	_ cciptypes.Bytes,
	attestation cciptypes.Bytes,
) (cciptypes.Bytes, error) {
	return attestation, nil
}
