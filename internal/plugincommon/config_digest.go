package plugincommon

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"

	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// ConfigDigestsMatch compares the expected config digest (from the home chain / OCR config) against the
// offramp's on-chain config digest for the given plugin type.
// Returns true when the digests match, false when they differ.
func ConfigDigestsMatch(
	ctx context.Context,
	ccipReader readerpkg.CCIPReader,
	pluginType uint8,
	expectedDigest [32]byte,
) (match bool, offRampDigest [32]byte, err error) {
	offRampDigest, err = ccipReader.GetOffRampConfigDigest(ctx, pluginType)
	if err != nil {
		return false, offRampDigest, fmt.Errorf("get offramp config digest: %w", err)
	}

	return bytes.Equal(offRampDigest[:], expectedDigest[:]), offRampDigest, nil
}

// FormatConfigDigest returns a hex-encoded string representation of a config digest.
func FormatConfigDigest(digest [32]byte) string {
	return hex.EncodeToString(digest[:])
}
