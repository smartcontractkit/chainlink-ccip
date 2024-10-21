// Package gen wraps an external type to generate a mock object.
package gen

import cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

// TODO: get rid of this.
// ExecutePluginCodec is defined in chainlink-common.
type ExecutePluginCodec interface {
	cciptypes.ExecutePluginCodec
}
