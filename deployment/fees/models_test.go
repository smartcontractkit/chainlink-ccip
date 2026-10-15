package fees

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
)

func TestUnresolvedFeeQuoterDestChainConfigResolve(t *testing.T) {
	var dst DestChainConfigForDst
	require.NoError(t, yaml.Unmarshal([]byte(`
selector: 16015286601757825753
config:
  destGasOverhead: 500000
  defaultTokenDestGasOverhead: 270000
`), &dst))

	onchain := lanes.FeeQuoterDestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                30_000,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		DestGasPerPayloadByteBase:   16,
		ChainFamilySelector:         0x2812d52c,
		DefaultTokenFeeUSDCents:     150,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		NetworkFeeUSDCents:          50,
		V1Params:                    &lanes.FeeQuoterV1Params{MaxNumberOfTokensPerMsg: 1, GasMultiplierWeiPerEth: 11e17},
		V2Params:                    &lanes.FeeQuoterV2Params{LinkFeeMultiplierPercent: 90, USDPerUnitGas: big.NewInt(1)},
	}
	want := onchain
	want.DestGasOverhead = 500_000
	want.DefaultTokenDestGasOverhead = 270_000

	require.False(t, dst.Config.IsEmpty())
	require.True(t, UnresolvedFeeQuoterDestChainConfig{}.IsEmpty())
	require.Equal(t, want, dst.Config.Resolve(onchain), "only the fields set in config change")
	require.Equal(t, onchain, UnresolvedFeeQuoterDestChainConfig{}.Resolve(onchain), "an empty config keeps the base config")
}
