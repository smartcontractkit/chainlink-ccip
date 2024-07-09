package pluginconfig

import (
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
)

func TestCommitPluginConfigValidate(t *testing.T) {
	testCases := []struct {
		name   string
		input  CommitPluginConfig
		expErr bool
	}{
		{
			name: "valid cfg",
			input: CommitPluginConfig{
				DestChain: cciptypes.ChainSelector(1),
				PricedTokens: []types.Account{
					types.Account("0x123"),
					types.Account("0x124"),
				},
				NewMsgScanBatchSize: 256,
				TokenPricesObserver: true,
			},
			expErr: false,
		},
		{
			name: "dest chain is empty",
			input: CommitPluginConfig{
				PricedTokens: []types.Account{
					types.Account("0x123"),
					types.Account("0x124"),
				},
				NewMsgScanBatchSize: 256,
				TokenPricesObserver: true,
			},
			expErr: true,
		},
		{
			name: "zero priced tokens",
			input: CommitPluginConfig{
				DestChain:           cciptypes.ChainSelector(1),
				NewMsgScanBatchSize: 256,
				TokenPricesObserver: true,
			},
			expErr: true,
		},
		{
			name: "empty batch scan size",
			input: CommitPluginConfig{
				DestChain: cciptypes.ChainSelector(1),
				PricedTokens: []types.Account{
					types.Account("0x123"),
					types.Account("0x124"),
				},
				TokenPricesObserver: true,
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.input.Validate()
			if tc.expErr {
				assert.Error(t, actual)
				return
			}
			assert.NoError(t, actual)
		})
	}
}
