package tokenprice

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	commonmock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

var (
	oneBig                = cciptypes.NewBigInt(big.NewInt(1))
	twoBig                = cciptypes.NewBigInt(big.NewInt(2))
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
				"0x2": oneBig,
				"0x3": oneBig,
				"0xa": oneBig,
			},
			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				"0x1": {},
				"0x2": {},
				"0x3": {},
				"0xa": {},
			},
			expErr: false,
		},
		{
			name: "nil price",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x2": oneBig,
				"0x3": nilBig, // nil price
				"0xa": oneBig,
			},
			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				"0x1": {},
				"0x2": {},
				"0x3": {},
				"0xa": {},
			},
			expErr: true,
		},
		{
			name: "negative price",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x2": oneBig,
				"0x3": negativeOneBig, // negative price
				"0xa": oneBig,
			},
			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				"0x1": {},
				"0x2": {},
				"0x3": {},
				"0xa": {},
			},
			expErr: true,
		},
		{
			name: "zero price",
			tokenPrices: cciptypes.TokenPriceMap{
				"0x1": oneBig,
				"0x2": oneBig,
				"0x3": zeroBig, // zero price
				"0xa": oneBig,
			},
			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				"0x1": {},
				"0x2": {},
				"0x3": {},
				"0xa": {},
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedTokenPrices(tc.tokenPrices, tc.tokensToQuery)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
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
		FeeQuoterTokenUpdates: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
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
				obs.FeeQuoterTokenUpdates = map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

//
//func TestValidateObservedTokenPrices(t *testing.T) {
//	testCases := []struct {
//		name          string
//		tokenPrices   cciptypes.TokenPriceMap
//		tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
//		expErr        bool
//	}{
//		{
//			name:        "empty is valid",
//			tokenPrices: cciptypes.TokenPriceMap{},
//			expErr:      false,
//		},
//		{
//			name: "all valid",
//			tokenPrices: cciptypes.TokenPriceMap{
//				"0x1": oneBig,
//				"0x2": oneBig,
//				"0x3": oneBig,
//				"0xa": oneBig,
//			},
//			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
//				"0x1": {},
//				"0x2": {},
//				"0x3": {},
//				"0xa": {},
//			},
//			expErr: false,
//		},
//		{
//			name: "nil price",
//			tokenPrices: cciptypes.TokenPriceMap{
//				"0x1": oneBig,
//				"0x2": oneBig,
//				"0x3": nil, // nil price
//				"0xa": oneBig,
//			},
//			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
//				"0x1": {},
//				"0x2": {},
//				"0x3": {},
//				"0xa": {},
//			},
//			expErr: true,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			err := validateObservedTokenPrices(tc.tokenPrices, tc.tokensToQuery)
//			if tc.expErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//			}
//		})
//	}
//}
//
//func TestValidateObservedTokenUpdates(t *testing.T) {
//	testCases := []struct {
//		name          string
//		tokenUpdates  map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig
//		tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
//		expErr        bool
//	}{
//		{
//			name:         "empty is valid",
//			tokenUpdates: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{},
//			expErr:       false,
//		},
//		{
//			name: "all valid",
//			tokenUpdates: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
//				"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0x2": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0x3": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0xa": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//			},
//			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
//				"0x1": {},
//				"0x2": {},
//				"0x3": {},
//				"0xa": {},
//			},
//			expErr: false,
//		},
//		{
//			name: "nil value",
//			tokenUpdates: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
//				"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0x2": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0x3": {Value: nil, Timestamp: time.Now().Add(-time.Hour)}, // nil value
//				"0xa": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//			},
//			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
//				"0x1": {},
//				"0x2": {},
//				"0x3": {},
//				"0xa": {},
//			},
//			expErr: true,
//		},
//		{
//			name: "invalid timestamp",
//			tokenUpdates: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
//				"0x1": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0x2": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//				"0x3": {Value: oneBig, Timestamp: time.Now().Add(time.Hour)}, // future timestamp
//				"0xa": {Value: oneBig, Timestamp: time.Now().Add(-time.Hour)},
//			},
//			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
//				"0x1": {},
//				"0x2": {},
//				"0x3": {},
//				"0xa": {},
//			},
//			expErr: true,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			err := validateObservedTokenUpdates(tc.tokenUpdates, tc.tokensToQuery)
//			if tc.expErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//			}
//		})
//	}
//}
//
//// Mock implementations for testing
//type mockChainSupport struct{}
//
//func (m mockChainSupport) SupportedChains(oracleID string) ([]string, error) {
//	return []string{"validChain", "destChain"}, nil
//}
//
