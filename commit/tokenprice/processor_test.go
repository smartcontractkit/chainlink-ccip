package tokenprice

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/mocks/shared"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	bi100              = big.NewInt(100)
	bi200              = big.NewInt(200)
	tokenA             = types.Account("0xAAAAAAAAAAAAAAAa75C1216873Ec4F88C11E57E3")
	tokenB             = types.Account("0xBBBBBBBBBBBBBBBb75C1216873Ec4F88C11E57E3")
	tokenPriceChainSel = cciptypes.ChainSelector(1)
	destChainSel       = cciptypes.ChainSelector(2)
)

func Test_Observation(t *testing.T) {
	fDestChain := 3
	timestamp := time.Now().UTC()
	feedTokenPrices := []cciptypes.TokenPrice{
		cciptypes.NewTokenPrice(tokenA, bi100),
		cciptypes.NewTokenPrice(tokenB, bi200),
	}
	feeQuoterTokenUpdates := map[types.Account]NumericalUpdate{
		tokenA: NewNumericalUpdate(bi100.Int64(), timestamp),
		tokenB: NewNumericalUpdate(bi200.Int64(), timestamp),
	}

	testCases := []struct {
		name         string
		getProcessor func(t *testing.T) *Processor
		expObs       Observation
		expErr       error
	}{
		{
			name: "Successful observation",
			getProcessor: func(t *testing.T) *Processor {
				chainSupport := shared.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportedChains(mock.Anything).Return(
					mapset.NewSet[cciptypes.ChainSelector](tokenPriceChainSel, destChainSel), nil,
				)
				chainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil)

				tokenPriceReader := readermock.NewMockPriceReader(t)
				tokenPriceReader.EXPECT().GetTokenFeedPricesUSD(mock.Anything, mock.Anything).Return([]*big.Int{bi100, bi200}, nil)
				tokenPriceReader.EXPECT().GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything).Return(
					map[types.Account]reader.NumericalUpdate{
						tokenA: reader.NewNumericalUpdate(bi100.Int64(), timestamp),
						tokenB: reader.NewNumericalUpdate(bi200.Int64(), timestamp),
					},
					nil,
				)

				homeChain := readermock.NewMockHomeChain(t)
				homeChain.EXPECT().GetFChain().Return(
					map[cciptypes.ChainSelector]int{destChainSel: fDestChain},
					nil,
				)

				return &Processor{
					oracleID:         1,
					lggr:             logger.Test(t),
					chainSupport:     chainSupport,
					tokenPriceReader: tokenPriceReader,
					homeChain:        homeChain,
					cfg:              defaultCfg,
				}
			},
			expObs: Observation{
				FeedTokenPrices:       feedTokenPrices,
				FeeQuoterTokenUpdates: feeQuoterTokenUpdates,
				FDestChain:            fDestChain,
				Timestamp:             time.Now().UTC(),
			},
			expErr: nil,
		},
		{
			name: "Failed to get FDestChain",
			getProcessor: func(t *testing.T) *Processor {
				homeChain := readermock.NewMockHomeChain(t)
				homeChain.EXPECT().GetFChain().Return(nil, errors.New("failed to get FChain"))

				chainSupport := shared.NewMockChainSupport(t)
				tokenPriceReader := readermock.NewMockPriceReader(t)

				return &Processor{
					oracleID:         1,
					lggr:             logger.Test(t),
					chainSupport:     chainSupport,
					tokenPriceReader: tokenPriceReader,
					homeChain:        homeChain,
					cfg:              defaultCfg,
				}
			},
			expObs: Observation{},
			expErr: errors.New("failed to get FChain"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			p := tc.getProcessor(t)

			actualObs, err := p.Observation(ctx, Outcome{}, Query{})
			if tc.expErr != nil {
				require.Error(t, err)
				assert.Equal(t, tc.expErr.Error(), err.Error())
				assert.Equal(t, Observation{}, actualObs)
			} else {
				require.NoError(t, err)
				// No need to check timestamp
				actualObs.Timestamp = tc.expObs.Timestamp
				assert.Equal(t, tc.expObs, actualObs)
			}
		})
	}
}

var defaultCfg = pluginconfig.CommitPluginConfig{
	DestChain: destChainSel,
	OffchainConfig: pluginconfig.CommitOffchainConfig{
		TokenInfo: map[types.Account]pluginconfig.TokenInfo{
			tokenA: {
				Decimals:          18,
				AggregatorAddress: "0x1111111111111111111111Ff18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
			},
			tokenB: {
				Decimals:          18,
				AggregatorAddress: "0x2222222222222222222222Ff18C45Df59775Fbb2",
				DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)}},
		},
		TokenPriceChainSelector: uint64(tokenPriceChainSel),
	},
}
