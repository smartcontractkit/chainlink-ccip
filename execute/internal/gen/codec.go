// Package gen wraps an external type to generate a mock object.
package gen

import cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

// ExecutePluginCodec is defined in chainlink-common.
type ExecutePluginCodec interface {
	cciptypes.ExecutePluginCodec
}
