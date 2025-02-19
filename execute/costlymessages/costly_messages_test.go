package costlymessages

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	gasmock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	b1 = mustMakeBytes("0x01")
	b2 = mustMakeBytes("0x02")
	b3 = mustMakeBytes("0x03")
)

func mustMakeBytes(s string) ccipocr3.Bytes32 {
	b, err := ccipocr3.NewBytes32FromString(s)
	if err != nil {
		panic(err)
	}
	return b
}

func TestCCIPCostlyMessageObserver_Observe(t *testing.T) {
	tests := []struct {
		name                               string
		notEnabled                         bool
		messageIDs                         []ccipocr3.Bytes32
		messageFees                        map[ccipocr3.Bytes32]plugintypes.USD18
		messageCosts                       map[ccipocr3.Bytes32]plugintypes.USD18
		disableAvailableFeeUsdCheckByChain map[ccipocr3.ChainSelector]bool
		sourceSelectorsByMessageID         map[ccipocr3.Bytes32]ccipocr3.ChainSelector
		want                               []ccipocr3.Bytes32
		wantErr                            assert.ErrorAssertionFunc
	}{
		{
			name:         "disabled observer returns empty slice",
			notEnabled:   true,
			messageIDs:   []ccipocr3.Bytes32{b1, b2, b3},
			messageFees:  map[ccipocr3.Bytes32]plugintypes.USD18{},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name:         "empty",
			messageIDs:   []ccipocr3.Bytes32{},
			messageFees:  map[ccipocr3.Bytes32]plugintypes.USD18{},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    []ccipocr3.Bytes32{},
			wantErr: assert.NoError,
		},
		{
			name:         "missing fees",
			messageIDs:   []ccipocr3.Bytes32{b1, b2, b2},
			messageFees:  map[ccipocr3.Bytes32]plugintypes.USD18{},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:       "missing costs",
			messageIDs: []ccipocr3.Bytes32{b1, b2, b2},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(10),
				b2: plugintypes.NewUSD18(20),
				b3: plugintypes.NewUSD18(30),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name:       "happy path",
			messageIDs: []ccipocr3.Bytes32{b1, b2, b3},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(10),
				b2: plugintypes.NewUSD18(20),
				b3: plugintypes.NewUSD18(30),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(5),
				b2: plugintypes.NewUSD18(25),
				b3: plugintypes.NewUSD18(15),
			},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    []ccipocr3.Bytes32{b2},
			wantErr: assert.NoError,
		},
		{
			name:       "happy path disable fee check explicitly set to false",
			messageIDs: []ccipocr3.Bytes32{b1, b2, b3},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(10),
				b2: plugintypes.NewUSD18(20),
				b3: plugintypes.NewUSD18(30),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(5),
				b2: plugintypes.NewUSD18(25),
				b3: plugintypes.NewUSD18(15),
			},
			disableAvailableFeeUsdCheckByChain: map[ccipocr3.ChainSelector]bool{
				1: false,
				2: false,
				3: false,
			},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    []ccipocr3.Bytes32{b2},
			wantErr: assert.NoError,
		},
		{
			name:       "disable available fee USD check for two source chains",
			messageIDs: []ccipocr3.Bytes32{b1, b2, b3},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(10),
				b2: plugintypes.NewUSD18(20),
				b3: plugintypes.NewUSD18(30),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{
				// All costs exceed messageFees
				b1: plugintypes.NewUSD18(35),
				b2: plugintypes.NewUSD18(35),
				b3: plugintypes.NewUSD18(35),
			},
			disableAvailableFeeUsdCheckByChain: map[ccipocr3.ChainSelector]bool{
				1: true,
				2: true,
				3: false,
			},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
				b2: 2,
				b3: 3,
			},
			want:    []ccipocr3.Bytes32{b3},
			wantErr: assert.NoError,
		},
		{
			name:       "disable available fee USD check - message fee is zero",
			messageIDs: []ccipocr3.Bytes32{b1},
			messageFees: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(0),
			},
			messageCosts: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(35),
			},
			disableAvailableFeeUsdCheckByChain: map[ccipocr3.ChainSelector]bool{
				1: true,
			},
			sourceSelectorsByMessageID: map[ccipocr3.Bytes32]ccipocr3.ChainSelector{
				b1: 1,
			},
			want:    []ccipocr3.Bytes32{},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			feeCalculator := NewStaticMessageFeeUSD18Calculator(tt.messageFees)
			execCostCalculator := NewStaticMessageExecCostUSD18Calculator(tt.messageCosts)
			observer := &observer{
				lggr:                               logger.Test(t),
				enabled:                            !tt.notEnabled,
				feeCalculator:                      feeCalculator,
				execCostCalculator:                 execCostCalculator,
				disableAvailableFeeUsdCheckByChain: tt.disableAvailableFeeUsdCheckByChain,
			}
			messages := make([]ccipocr3.Message, 0)
			for _, id := range tt.messageIDs {
				sourceSelector := tt.sourceSelectorsByMessageID[id]
				messages = append(messages, ccipocr3.Message{Header: ccipocr3.RampMessageHeader{
					MessageID:           id,
					SourceChainSelector: sourceSelector,
				}})
			}

			got, err := observer.Observe(ctx, messages, nil)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestWaitBoostedFee(t *testing.T) {
	lggr := logger.Test(t)

	tests := []struct {
		name                     string
		sendTimeDiff             time.Duration
		fee                      *big.Int
		diff                     *big.Int
		relativeBoostPerWaitHour float64
	}{
		{
			"wait 10s",
			time.Second * 10,
			big.NewInt(6e18), // Fee:   6    LINK

			big.NewInt(1166666666665984), // Boost: 0.01 LINK
			0.07,
		},
		{
			"wait 5m",
			time.Minute * 5,
			big.NewInt(6e18),  // Fee:   6    LINK
			big.NewInt(35e15), // Boost: 0.35 LINK
			0.07,
		},
		{
			"wait 7m",
			time.Minute * 7,
			big.NewInt(6e18),  // Fee:   6    LINK
			big.NewInt(49e15), // Boost: 0.49 LINK
			0.07,
		},
		{
			"wait 12m",
			time.Minute * 12,
			big.NewInt(6e18),  // Fee:   6    LINK
			big.NewInt(84e15), // Boost: 0.84 LINK
			0.07,
		},
		{
			"wait 25m",
			time.Minute * 25,
			big.NewInt(6e18),               // Fee:   6 LINK
			big.NewInt(174999999999998976), // Boost: 1.75 LINK
			0.07,
		},
		{
			"wait 1h",
			time.Hour * 1,
			big.NewInt(6e18),   // Fee:   6 LINK
			big.NewInt(420e15), // Boost: 4.2 LINK
			0.07,
		},
		{
			"wait 5h",
			time.Hour * 5,
			big.NewInt(6e18),                // Fee:   6 LINK
			big.NewInt(2100000000000001024), // Boost: 21LINK
			0.07,
		},
		{
			"wait 24h",
			time.Hour * 24,
			big.NewInt(6e18), // Fee:   6 LINK
			big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1008e15)), // Boost: 100LINK
			0.07,
		},
		{
			"high boost wait 10s",
			time.Second * 10,
			big.NewInt(5e18),
			big.NewInt(9722222222222336), // 1e16
			0.7,
		},
		{
			"high boost wait 5m",
			time.Minute * 5,
			big.NewInt(5e18),
			big.NewInt(291666666666667008), // 1e18
			0.7,
		},
		{
			"high boost wait 25m",
			time.Minute * 25,
			big.NewInt(5e18),
			big.NewInt(1458333333333334016), // 1e19
			0.7,
		},
		{
			"high boost wait 5h",
			time.Hour * 5,
			big.NewInt(5e18),
			big.NewInt(0).Mul(big.NewInt(10), big.NewInt(175e16)), // 1e20
			0.7,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			boosted := waitBoostedFee(lggr, tc.sendTimeDiff, tc.fee, tc.relativeBoostPerWaitHour)
			diff := big.NewInt(0).Sub(boosted, tc.fee)
			assert.Equal(t, diff, tc.diff)
		})
	}
}

func TestCCIPMessageFeeE18USDCalculator_MessageFeeE18USD(t *testing.T) {
	mockNow := func() time.Time {
		return time.Date(2023, time.January, 1, 14, 0, 0, 0, time.UTC)
	}
	t1 := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2023, time.January, 1, 13, 0, 0, 0, time.UTC)
	t3 := time.Date(2023, time.January, 1, 14, 0, 0, 0, time.UTC)

	tests := []struct {
		name                     string
		messages                 []ccipocr3.Message
		messageTimeStamps        map[ccipocr3.Bytes32]time.Time
		linkPrice                ccipocr3.BigInt
		relativeBoostPerWaitHour float64
		want                     map[ccipocr3.Bytes32]plugintypes.USD18
		wantErr                  assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			messages: []ccipocr3.Message{
				{
					Header:        ccipocr3.RampMessageHeader{MessageID: b1},
					FeeValueJuels: ccipocr3.NewBigIntFromInt64(14e14),
				},
				{
					Header:        ccipocr3.RampMessageHeader{MessageID: b2},
					FeeValueJuels: ccipocr3.NewBigIntFromInt64(25e14),
				},
				{
					Header:        ccipocr3.RampMessageHeader{MessageID: b3},
					FeeValueJuels: ccipocr3.NewBigIntFromInt64(36e14),
				},
			},
			messageTimeStamps: map[ccipocr3.Bytes32]time.Time{
				b1: t1,
				b2: t2,
				b3: t3,
			},
			linkPrice:                ccipocr3.NewBigIntFromInt64(8e18), // 8 USD
			relativeBoostPerWaitHour: 0.5,
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(2.24e16),
				b2: plugintypes.NewUSD18(3e16),
				b3: plugintypes.NewUSD18(2.88e16),
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			lggr := logger.Test(t)

			mockReader := readerpkg_mock.NewMockCCIPReader(t)
			mockReader.EXPECT().LinkPriceUSD(ctx).Return(tt.linkPrice, nil)

			calculator := &CCIPMessageFeeUSD18Calculator{
				lggr:                     lggr,
				ccipReader:               mockReader,
				relativeBoostPerWaitHour: tt.relativeBoostPerWaitHour,
				now:                      mockNow,
			}

			got, err := calculator.MessageFeeUSD18(ctx, tt.messages, tt.messageTimeStamps)
			tt.wantErr(t, err, "MessageFeeUSD18(...)")
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCCIPMessageExecCostUSD18Calculator_MessageExecCostUSD18(t *testing.T) {
	destChainSelector := ccipocr3.ChainSelector(1)
	nativeTokenPrice := ccipocr3.BigInt{
		Int: new(big.Int).Mul(big.NewInt(9), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))}

	tests := []struct {
		name                string
		messages            []ccipocr3.Message
		messageGases        []uint64
		executionFee        *big.Int
		dataAvailabilityFee *big.Int
		feeComponentsError  error
		daGasConfig         ccipocr3.DataAvailabilityGasConfig
		want                map[ccipocr3.Bytes32]plugintypes.USD18
		wantErr             bool
	}{
		{
			name:                "empty messages slice results in error",
			messages:            []ccipocr3.Message{},
			messageGases:        []uint64{},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: big.NewInt(0),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1,
				DestGasPerDataAvailabilityByte:    1,
				DestDataAvailabilityMultiplierBps: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "happy path, no DA cost",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1, DestChainSelector: destChainSelector},
				},
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b2, DestChainSelector: destChainSelector},
				},
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b3, DestChainSelector: destChainSelector},
				},
			},
			messageGases:        []uint64{100, 200, 300},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: big.NewInt(0),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1,
				DestGasPerDataAvailabilityByte:    1,
				DestDataAvailabilityMultiplierBps: 1,
			},
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(18000000000000),
				b2: plugintypes.NewUSD18(36000000000000),
				b3: plugintypes.NewUSD18(54000000000000),
			},
			wantErr: false,
		},
		{
			name: "happy path, with DA cost",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1, DestChainSelector: destChainSelector},
				},
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b2, DestChainSelector: destChainSelector},
				},
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b3, DestChainSelector: destChainSelector},
				},
			},
			messageGases:        []uint64{100, 200, 300},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1200,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 200,
			},
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(119700000000000),
				b2: plugintypes.NewUSD18(137700000000000),
				b3: plugintypes.NewUSD18(155700000000000),
			},
			wantErr: false,
		},
		{
			name: "message with token amounts affects DA gas calculation",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1, DestChainSelector: destChainSelector},
					TokenAmounts: []ccipocr3.RampTokenAmount{
						{
							SourcePoolAddress: []byte("source_pool"),
							DestTokenAddress:  []byte("dest_token"),
							ExtraData:         []byte("extra"),
							DestExecData:      []byte("exec_data"),
							Amount:            ccipocr3.NewBigInt(big.NewInt(1)),
						},
					},
					Data:      []byte("some_data"),
					Sender:    []byte("sender"),
					Receiver:  []byte("receiver"),
					ExtraArgs: []byte("extra_args"),
				},
			},
			messageGases:        []uint64{100},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 200,
			},
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(173700000000000),
			},
			wantErr: false,
		},
		{
			name: "zero DA multiplier results in only overhead gas",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1, DestChainSelector: destChainSelector},
					Data:   []byte("some_data"),
				},
			},
			messageGases:        []uint64{100},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 0, // Zero multiplier
			},
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(18000000000000), // Only exec cost, DA cost is 0
			},
			wantErr: false,
		},
		{
			name: "large message with multiple tokens",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1, DestChainSelector: destChainSelector},
					TokenAmounts: []ccipocr3.RampTokenAmount{
						{
							SourcePoolAddress: make([]byte, 100), // Large token data
							DestTokenAddress:  make([]byte, 100),
							ExtraData:         make([]byte, 100),
							DestExecData:      make([]byte, 100),
							Amount:            ccipocr3.NewBigInt(big.NewInt(1)),
						},
						{
							SourcePoolAddress: make([]byte, 100), // Second token
							DestTokenAddress:  make([]byte, 100),
							ExtraData:         make([]byte, 100),
							DestExecData:      make([]byte, 100),
							Amount:            ccipocr3.NewBigInt(big.NewInt(1)),
						},
					},
					Data:      make([]byte, 1000), // Large message data
					Sender:    make([]byte, 100),
					Receiver:  make([]byte, 100),
					ExtraArgs: make([]byte, 100),
				},
			},
			messageGases:        []uint64{100},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 200,
			},
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(489600000000000),
			},
			wantErr: false,
		},
		{
			name: "fee components error",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1},
				},
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b2},
				},
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b3},
				},
			},
			messageGases:        []uint64{100, 200, 300},
			executionFee:        big.NewInt(100),
			dataAvailabilityFee: big.NewInt(0),
			feeComponentsError:  fmt.Errorf("error"),
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1,
				DestGasPerDataAvailabilityByte:    1,
				DestDataAvailabilityMultiplierBps: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "minimal message - only constant parts",
			messages: []ccipocr3.Message{
				{
					Header: ccipocr3.RampMessageHeader{MessageID: b1, DestChainSelector: destChainSelector},
				},
			},
			messageGases:        []uint64{100},
			executionFee:        new(big.Int).Mul(big.NewInt(20), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			dataAvailabilityFee: new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(9), nil)),
			feeComponentsError:  nil,
			daGasConfig: ccipocr3.DataAvailabilityGasConfig{
				DestDataAvailabilityOverheadGas:   1000,
				DestGasPerDataAvailabilityByte:    10,
				DestDataAvailabilityMultiplierBps: 200,
			},
			want: map[ccipocr3.Bytes32]plugintypes.USD18{
				b1: plugintypes.NewUSD18(116100000000000),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			lggr := logger.Test(t)

			mockReader := readerpkg_mock.NewMockCCIPReader(t)
			feeComponents := types.ChainFeeComponents{
				ExecutionFee:        tt.executionFee,
				DataAvailabilityFee: tt.dataAvailabilityFee,
			}
			if len(tt.messages) > 0 {
				mockReader.EXPECT().GetDestChainFeeComponents(ctx).Return(feeComponents, tt.feeComponentsError)
				mockReader.EXPECT().GetWrappedNativeTokenPriceUSD(
					ctx,
					[]ccipocr3.ChainSelector{destChainSelector},
				).Return(
					map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						destChainSelector: nativeTokenPrice,
					},
				).Maybe()
				if !tt.wantErr {
					mockReader.EXPECT().GetMedianDataAvailabilityGasConfig(ctx).Return(tt.daGasConfig, nil)
				}
			}

			ep := gasmock.NewMockEstimateProvider(t)
			if !tt.wantErr {
				for _, messageGas := range tt.messageGases {
					ep.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(messageGas).Once()
				}
			}

			calculator := CCIPMessageExecCostUSD18Calculator{
				lggr:             lggr,
				ccipReader:       mockReader,
				estimateProvider: ep,
			}

			got, err := calculator.MessageExecCostUSD18(ctx, tt.messages)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestValidatePositive(t *testing.T) {
	tests := []struct {
		name    string
		fee     *big.Int
		wantErr bool
	}{
		{
			name:    "nil fee",
			fee:     nil,
			wantErr: true,
		},
		{
			name:    "negative fee",
			fee:     big.NewInt(-1),
			wantErr: true,
		},
		{
			name:    "zero fee",
			fee:     big.NewInt(0),
			wantErr: true,
		},
		{
			name:    "positive fee",
			fee:     big.NewInt(1),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePositive(tt.fee)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePositive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalculateETA(t *testing.T) {
	tests := []struct {
		name       string
		currentFee *big.Int
		targetFee  *big.Int
		boostRate  float64
		wantETA    *time.Duration
	}{
		{
			name:       "zero boost rate returns nil",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(200),
			boostRate:  0,
			wantETA:    nil,
		},
		{
			name:       "negative boost rate returns nil",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(200),
			boostRate:  -0.5,
			wantETA:    nil,
		},
		{
			name:       "target less than current returns nil",
			currentFee: big.NewInt(200),
			targetFee:  big.NewInt(100),
			boostRate:  0.5,
			wantETA:    nil,
		},
		{
			name:       "simple doubling with 0.5 boost rate",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(200),
			boostRate:  0.5,
			wantETA:    durationPtr(2 * time.Hour),
		},
		{
			name:       "50% increase with 0.25 boost rate",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(150),
			boostRate:  0.25,
			wantETA:    durationPtr(2 * time.Hour),
		},
		{
			name:       "large numbers - 2x with 0.5 boost",
			currentFee: new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil), // 10^30
			targetFee: new(big.Int).Mul(
				big.NewInt(2),
				new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil),
			), // 2 * 10^30
			boostRate: 0.5,
			wantETA:   durationPtr(2 * time.Hour),
		},
		{
			name:       "fractional increase - 10% with 0.1 boost",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(110),
			boostRate:  0.1,
			wantETA:    durationPtr(1 * time.Hour),
		},
		{
			name:       "many days ETA",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(1000),
			boostRate:  0.1,
			wantETA:    durationPtr(90 * time.Hour),
		},
		{
			name:       "very small boost rate",
			currentFee: big.NewInt(100),
			targetFee:  big.NewInt(150),
			boostRate:  0.01,
			wantETA:    durationPtr(50 * time.Hour),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotETA := calculateETA(tt.currentFee, tt.targetFee, tt.boostRate)

			if tt.wantETA == nil {
				assert.Nil(t, gotETA)
			} else {
				assert.NotNil(t, gotETA)
				// Allow for small floating point differences
				assert.InDelta(t, tt.wantETA.Hours(), gotETA.Hours(), 0.001)
			}
		})
	}
}

func TestFormatETA(t *testing.T) {
	tests := []struct {
		name     string
		duration *time.Duration
		want     string
	}{
		{
			name:     "nil duration",
			duration: nil,
			want:     "boost rate is zero or already at target",
		},
		{
			name:     "less than one hour",
			duration: durationPtr(30 * time.Minute),
			want:     "~30 minutes",
		},
		{
			name:     "exactly one hour",
			duration: durationPtr(1 * time.Hour),
			want:     "~1.0 hours",
		},
		{
			name:     "multiple hours",
			duration: durationPtr(5*time.Hour + 30*time.Minute),
			want:     "~5.5 hours",
		},
		{
			name:     "exactly one day",
			duration: durationPtr(24 * time.Hour),
			want:     "~1.0 days",
		},
		{
			name:     "multiple days",
			duration: durationPtr(72*time.Hour + 12*time.Hour),
			want:     "~3.5 days",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatETA(tt.duration)
			assert.Equal(t, tt.want, got)
		})
	}
}

func durationPtr(d time.Duration) *time.Duration {
	return &d
}
