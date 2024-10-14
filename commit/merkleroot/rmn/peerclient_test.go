package rmn

import (
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
)

func Test_writePrefix(t *testing.T) {
	prefix := 5130 // 0x140a -> {20, 10}
	h := cciptypes.Bytes32{1, 2, 3, 4, 5}
	exp := cciptypes.Bytes32{20, 10, 3, 4, 5}
	got := writePrefix(ocr2types.ConfigDigestPrefix(prefix), h)
	assert.Equal(t, exp, got)
}