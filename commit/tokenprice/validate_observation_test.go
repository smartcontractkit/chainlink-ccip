package tokenprice

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	commonmock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var (
	oneBig                = cciptypes.NewBigInt(big.NewInt(1))
	negativeOneBig        = cciptypes.NewBigInt(big.NewInt(-1))
	zeroBig               = cciptypes.NewBigInt(big.NewInt(0))
	nilBig                = cciptypes.NewBigInt(nil)
	defaultOffChainConfig = pluginconfig.CommitOffchainConfig{
		PriceFeedChainSelector: feedChainSel,
		TokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
			"0x1": {},
			"0x2": {},
			"0x3": {},
			"0xa": {},
		},
	}
	defaultTokensToQuery = map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
		"0x1": {},
		"0x2": {},
		"0x3": {},
		"0xa": {},
	}
)

func Test_validateObservedTokenPrices(t *testing.T) {
	testCases := []struct {
		name          string
		tokenPrices   cciptypes.TokenPriceMap
		tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
		expErr        bool
	}{
		{
			name:        "empty is valid",
			tokenPrices: cciptypes.TokenPriceMap{},
			expErr:      false,
		},
		{
			name: "all valid",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x3": oneBig,
				"0xa": oneBig,
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        false,
		},
		{
			name: "nil price",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x2": oneBig,
				"0x3": nilBig, // nil price
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
		{
			name: "negative price",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x3": negativeOneBig, // negative price
				"0xa": oneBig,
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
		{
			name: "zero price",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x3": zeroBig, // zero price
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
		{
			name: "non queryable token",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x5": oneBig,
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedTokenPrices(tc.tokenPrices, tc.tokensToQuery)
			if tc.expErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestValidateObservation(t *testing.T) {
	prevOutcome := Outcome{}
	query := Query{}

	oracleID := commontypes.OracleID(1)
	chainSupport := commonmock.NewMockChainSupport(t)
	supportedChains := mapset.NewSet[cciptypes.ChainSelector]()
	supportedChains.Add(feedChainSel)
	supportedChains.Add(destChainSel)
	chainSupport.On("SupportedChains", oracleID).Return(supportedChains, nil)

	defaultObs := Observation{
		FChain: map[cciptypes.ChainSelector]int{
			feedChainSel: 1,
			destChainSel: 1,
		},
		FeedTokenPrices: cciptypes.TokenPriceMap{
			"0x1": oneBig,
		},
		FeeQuoterTokenUpdates: map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
			"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
		},
		Timestamp: time.Now().Add(-time.Hour),
	}

	testCases := []struct {
		name             string
		obs              func() Observation
		chainSupportMock func() *commonmock.MockChainSupport
		expErr           bool
	}{
		{
			name: "valid observation",
			obs: func() Observation {
				return defaultObs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				return chainSupport
			},
			expErr: false,
		},
		{
			name: "empty observation",
			obs: func() Observation {
				return Observation{}
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				return chainSupport
			},
			expErr: false,
		},
		{
			name: "invalid FChain",
			obs: func() Observation {
				obs := defaultObs
				obs.FChain = map[cciptypes.ChainSelector]int{
					destChainSel: -1,
				}
				return obs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				return chainSupport
			},
			expErr: true,
		},
		{
			name: "unsupported feed chain",
			obs: func() Observation {
				return defaultObs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				mock := commonmock.NewMockChainSupport(t)
				sc := mapset.NewSet[cciptypes.ChainSelector](destChainSel)
				mock.On("SupportedChains", oracleID).Return(sc, nil)
				return mock
			},
			expErr: true,
		},
		{
			name: "invalid token price",
			obs: func() Observation {
				obs := defaultObs
				obs.FeedTokenPrices = cciptypes.TokenPriceMap{
					"0x1": negativeOneBig,
				}
				return obs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				return chainSupport
			},
			expErr: true,
		},
		{
			name: "unsupported dest chain",
			obs: func() Observation {
				return defaultObs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				mock := commonmock.NewMockChainSupport(t)
				sc := mapset.NewSet[cciptypes.ChainSelector](feedChainSel)
				mock.On("SupportedChains", oracleID).Return(sc, nil).Maybe()
				return mock
			},
			expErr: true,
		},
		{
			name: "invalid token update",
			obs: func() Observation {
				obs := defaultObs
				obs.FeeQuoterTokenUpdates = map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
					"0x1": {Value: oneBig, Timestamp: time.Now().Add(time.Hour)}, // future timestamp
				}
				return obs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				return chainSupport
			},
			expErr: true,
		},
		{
			name: "invalid timestamp",
			obs: func() Observation {
				obs := defaultObs
				obs.Timestamp = time.Now().Add(time.Hour)
				return obs
			},
			chainSupportMock: func() *commonmock.MockChainSupport {
				return chainSupport
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			chainSupportMock := tc.chainSupportMock()
			p := &processor{
				chainSupport: chainSupportMock,
				offChainCfg:  defaultOffChainConfig,
				destChain:    destChainSel,
			}
			ao := plugincommon.AttributedObservation[Observation]{
				OracleID:    oracleID,
				Observation: tc.obs(),
			}
			err := p.ValidateObservation(prevOutcome, query, ao)
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateObservedTokenUpdates(t *testing.T) {
	testCases := []struct {
		name          string
		tokenUpdates  map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig
		tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
		expErr        bool
	}{
		{
			name:         "empty is valid",
			tokenUpdates: map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{},
			expErr:       false,
		},
		{
			name: "all valid",
			tokenUpdates: map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
				"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
				"0xa": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        false,
		},
		{
			name: "nil value",
			tokenUpdates: map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
				"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
				"0x3": {Value: nilBig, Timestamp: time.Now().Add(-time.Hour)}, // nil value
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
		{
			name: "invalid timestamp",
			tokenUpdates: map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
				"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
				"0x3": {Value: oneBig, Timestamp: time.Now().Add(time.Hour)}, // future timestamp
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
		{
			name: "non queryable token",
			tokenUpdates: map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
				"0x5": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
			},
			tokensToQuery: defaultTokensToQuery,
			expErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedTokenUpdates(tc.tokenUpdates, tc.tokensToQuery)
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
