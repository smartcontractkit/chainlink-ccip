package commit

import (
	"context"
	"encoding/binary"
	"math/big"
	"slices"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func Test_observeMaxSeqNumsPerChain(t *testing.T) {
	testCases := []struct {
		name               string
		prevOutcome        plugintypes.CommitPluginOutcome
		onChainNextSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum
		readChains         []cciptypes.ChainSelector
		destChain          cciptypes.ChainSelector
		expErr             bool
		expSeqNumsInSync   bool
		expMaxSeqNums      []plugintypes.SeqNumChain
	}{
		{
			name: "report on chain seq num and can read dest",
			onChainNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
				1: 11,
				2: 21,
			},
			readChains: []cciptypes.ChainSelector{1, 2, 3},
			destChain:  3,
			expErr:     false,
			expMaxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
		},
		{
			name: "cannot read dest",
			prevOutcome: plugintypes.CommitPluginOutcome{
				MaxSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1, SeqNum: 11}, // for chain 1 previous outcome is higher than on-chain state
					{ChainSel: 2, SeqNum: 19}, // for chain 2 previous outcome is behind on-chain state
				},
			},
			onChainNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
				1: 11,
				2: 21,
			},
			readChains:    []cciptypes.ChainSelector{1, 2},
			destChain:     3,
			expErr:        false,
			expMaxSeqNums: []plugintypes.SeqNumChain{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockReader := mocks.NewCCIPReader()
			knownSourceChains := slicelib.Filter(
				tc.readChains,
				func(ch cciptypes.ChainSelector) bool { return ch != tc.destChain },
			)
			lggr := logger.Test(t)

			onChainSeqNums := make([]cciptypes.SeqNum, 0)
			for _, chain := range knownSourceChains {
				if v, ok := tc.onChainNextSeqNums[chain]; !ok {
					t.Fatalf("invalid test case missing on chain seq num expectation for %d", chain)
				} else {
					onChainSeqNums = append(onChainSeqNums, v)
				}
			}
			mockReader.On("NextSeqNum", ctx, knownSourceChains).Return(onChainSeqNums, nil)

			seqNums, err := observeLatestCommittedSeqNums(
				ctx,
				lggr,
				mockReader,
				mapset.NewSet(tc.readChains...),
				tc.destChain,
				knownSourceChains,
			)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expMaxSeqNums, seqNums)
		})
	}
}

func Test_observeNewMsgs(t *testing.T) {
	testCases := []struct {
		name               string
		maxSeqNumsPerChain []plugintypes.SeqNumChain
		readChains         []cciptypes.ChainSelector
		destChain          cciptypes.ChainSelector
		msgScanBatchSize   int
		newMsgs            map[cciptypes.ChainSelector][]cciptypes.Message
		expMsgs            []cciptypes.Message
		expErr             bool
	}{
		{
			name: "no new messages",
			maxSeqNumsPerChain: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			readChains:       []cciptypes.ChainSelector{1, 2},
			msgScanBatchSize: 256,
			newMsgs: map[cciptypes.ChainSelector][]cciptypes.Message{
				1: {},
				2: {},
			},
			expMsgs: []cciptypes.Message{},
			expErr:  false,
		},
		{
			name: "new messages",
			maxSeqNumsPerChain: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			readChains:       []cciptypes.ChainSelector{1, 2},
			msgScanBatchSize: 256,
			newMsgs: map[cciptypes.ChainSelector][]cciptypes.Message{
				1: {
					{Header: cciptypes.RampMessageHeader{
						MessageID:           mustNewMessageID("0x01"),
						SourceChainSelector: 1,
						SequenceNumber:      11}},
				},
				2: {
					{Header: cciptypes.RampMessageHeader{
						MessageID:           mustNewMessageID("0x02"),
						SourceChainSelector: 2,
						SequenceNumber:      21}},
					{Header: cciptypes.RampMessageHeader{
						MessageID:           mustNewMessageID("0x03"),
						SourceChainSelector: 2,
						SequenceNumber:      22}},
				},
			},
			expMsgs: []cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}},
				{Header: cciptypes.RampMessageHeader{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 2,
					SequenceNumber:      21}},
				{Header: cciptypes.RampMessageHeader{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 2,
					SequenceNumber:      22}},
			},
			expErr: false,
		},
		{
			name: "new messages but one chain is not readable",
			maxSeqNumsPerChain: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			readChains:       []cciptypes.ChainSelector{2},
			msgScanBatchSize: 256,
			newMsgs: map[cciptypes.ChainSelector][]cciptypes.Message{
				2: {
					{Header: cciptypes.RampMessageHeader{
						MessageID:           mustNewMessageID("0x02"),
						SourceChainSelector: 2,
						SequenceNumber:      21}},
					{Header: cciptypes.RampMessageHeader{
						MessageID:           mustNewMessageID("0x03"),
						SourceChainSelector: 2,
						SequenceNumber:      22}},
				},
			},
			expMsgs: []cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 2,
					SequenceNumber:      21}},
				{Header: cciptypes.RampMessageHeader{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 2,
					SequenceNumber:      22}},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockReader := mocks.NewCCIPReader()
			msgHasher := mocks.NewMessageHasher()
			for i := range tc.expMsgs { // make sure the hashes are populated
				h, err := msgHasher.Hash(ctx, tc.expMsgs[i])
				assert.NoError(t, err)
				tc.expMsgs[i].Header.MsgHash = h
			}

			lggr := logger.Test(t)

			for _, seqNumChain := range tc.maxSeqNumsPerChain {
				if slices.Contains(tc.readChains, seqNumChain.ChainSel) {
					mockReader.On(
						"MsgsBetweenSeqNums",
						ctx,
						seqNumChain.ChainSel,
						cciptypes.NewSeqNumRange(seqNumChain.SeqNum+1, seqNumChain.SeqNum+cciptypes.SeqNum(1+tc.msgScanBatchSize)),
					).Return(tc.newMsgs[seqNumChain.ChainSel], nil)
				}
			}

			msgs, err := observeNewMsgs(
				ctx,
				lggr,
				mockReader,
				msgHasher,
				mapset.NewSet(tc.readChains...),
				tc.maxSeqNumsPerChain,
				tc.msgScanBatchSize,
			)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expMsgs, msgs)
			mockReader.AssertExpectations(t)
		})
	}
}

func Benchmark_observeNewMsgs(b *testing.B) {
	const (
		numChains       = 5
		readerDelayMS   = 100
		newMsgsPerChain = 256
	)

	readChains := make([]cciptypes.ChainSelector, numChains)
	maxSeqNumsPerChain := make([]plugintypes.SeqNumChain, numChains)
	for i := 0; i < numChains; i++ {
		readChains[i] = cciptypes.ChainSelector(i + 1)
		maxSeqNumsPerChain[i] = plugintypes.SeqNumChain{ChainSel: cciptypes.ChainSelector(i + 1), SeqNum: cciptypes.SeqNum(1)}
	}

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		lggr, _ := logger.New()
		ccipReader := mocks.NewCCIPReader()
		msgHasher := mocks.NewMessageHasher()

		expNewMsgs := make([]cciptypes.Message, 0, newMsgsPerChain*numChains)
		for _, seqNumChain := range maxSeqNumsPerChain {
			newMsgs := make([]cciptypes.Message, 0, newMsgsPerChain)
			for msgSeqNum := 1; msgSeqNum <= newMsgsPerChain; msgSeqNum++ {
				newMsgs = append(newMsgs, cciptypes.Message{
					Header: cciptypes.RampMessageHeader{
						MessageID:           messageIDFromInt(msgSeqNum),
						SourceChainSelector: seqNumChain.ChainSel,
						SequenceNumber:      cciptypes.SeqNum(msgSeqNum),
					},
				})
			}

			ccipReader.On(
				"MsgsBetweenSeqNums",
				ctx,
				[]cciptypes.ChainSelector{seqNumChain.ChainSel},
				cciptypes.NewSeqNumRange(
					seqNumChain.SeqNum+1,
					seqNumChain.SeqNum+cciptypes.SeqNum(1+newMsgsPerChain),
				),
			).Run(func(args mock.Arguments) {
				time.Sleep(time.Duration(readerDelayMS) * time.Millisecond)
			}).Return(newMsgs, nil)
			expNewMsgs = append(expNewMsgs, newMsgs...)
		}

		msgs, err := observeNewMsgs(
			ctx,
			lggr,
			ccipReader,
			msgHasher,
			mapset.NewSet(readChains...),
			maxSeqNumsPerChain,
			newMsgsPerChain,
		)
		assert.NoError(b, err)
		assert.Equal(b, expNewMsgs, msgs)

		// (old)     sequential: 509.345 ms/op   (numChains * readerDelayMS)
		// (current) parallel:   102.543 ms/op     (readerDelayMS)
	}
}

func Test_observeTokenPrices(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		priceReader := mocks.NewTokenPricesReader()
		tokens := []types.Account{"0x1", "0x2", "0x3"}
		mockPrices := []*big.Int{big.NewInt(10), big.NewInt(20), big.NewInt(30)}
		priceReader.On("GetTokenPricesUSD", ctx, tokens).Return(mockPrices, nil)
		prices, err := observeTokenPrices(ctx, priceReader, tokens)
		assert.NoError(t, err)
		assert.Equal(t, []cciptypes.TokenPrice{
			cciptypes.NewTokenPrice("0x1", big.NewInt(10)),
			cciptypes.NewTokenPrice("0x2", big.NewInt(20)),
			cciptypes.NewTokenPrice("0x3", big.NewInt(30)),
		}, prices)
	})

	t.Run("price reader internal issue", func(t *testing.T) {
		priceReader := mocks.NewTokenPricesReader()
		tokens := []types.Account{"0x1", "0x2", "0x3"}
		mockPrices := []*big.Int{big.NewInt(10), big.NewInt(20)} // returned two prices for three tokens
		priceReader.On("GetTokenPricesUSD", ctx, tokens).Return(mockPrices, nil)
		_, err := observeTokenPrices(ctx, priceReader, tokens)
		assert.Error(t, err)
	})

}

func Test_observeGasPrices(t *testing.T) {
	ctx := context.Background()

	t.Run("happy path", func(t *testing.T) {
		mockReader := mocks.NewCCIPReader()
		chains := []cciptypes.ChainSelector{1, 2, 3}
		mockGasPrices := []cciptypes.BigInt{
			cciptypes.NewBigIntFromInt64(10),
			cciptypes.NewBigIntFromInt64(20),
			cciptypes.NewBigIntFromInt64(30),
		}
		mockReader.On("GasPrices", ctx, chains).Return(mockGasPrices, nil)
		gasPrices, err := observeGasPrices(ctx, mockReader, chains)
		assert.NoError(t, err)
		assert.Equal(t, []cciptypes.GasPriceChain{
			cciptypes.NewGasPriceChain(mockGasPrices[0].Int, chains[0]),
			cciptypes.NewGasPriceChain(mockGasPrices[1].Int, chains[1]),
			cciptypes.NewGasPriceChain(mockGasPrices[2].Int, chains[2]),
		}, gasPrices)
	})

	t.Run("gas reader internal issue", func(t *testing.T) {
		mockReader := mocks.NewCCIPReader()
		chains := []cciptypes.ChainSelector{1, 2, 3}
		mockGasPrices := []cciptypes.BigInt{
			cciptypes.NewBigIntFromInt64(10),
			cciptypes.NewBigIntFromInt64(20),
		} // return 2 prices for 3 chains
		mockReader.On("GasPrices", ctx, chains).Return(mockGasPrices, nil)
		_, err := observeGasPrices(ctx, mockReader, chains)
		assert.Error(t, err)
	})
}

func Test_validateObservedSequenceNumbers(t *testing.T) {
	testCases := []struct {
		name       string
		msgs       []cciptypes.RampMessageHeader
		maxSeqNums []plugintypes.SeqNumChain
		expErr     bool
	}{
		{
			name:       "empty",
			msgs:       nil,
			maxSeqNums: nil,
			expErr:     false,
		},
		{
			name: "dup seq num observation",
			msgs: nil,
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
				{ChainSel: 1, SeqNum: 10},
			},
			expErr: true,
		},
		{
			name: "seq nums ok",
			msgs: nil,
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			expErr: false,
		},
		{
			name: "dup msg seq num",
			msgs: []cciptypes.RampMessageHeader{
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 13},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 14},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 13}, // dup
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			expErr: true,
		},
		{
			name: "msg seq nums ok",
			msgs: []cciptypes.RampMessageHeader{
				{MsgHash: cciptypes.Bytes32{1}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MsgHash: cciptypes.Bytes32{2}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 13},
				{MsgHash: cciptypes.Bytes32{3}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 14},
				{MsgHash: cciptypes.Bytes32{4}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 2, SequenceNumber: 21},
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			expErr: false,
		},
		{
			name: "msg seq nums does not match observed max seq num",
			msgs: []cciptypes.RampMessageHeader{
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 13},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 10}, // max seq num is already 10
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 2, SequenceNumber: 21},
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			expErr: true,
		},
		{
			name: "max seq num not found",
			msgs: []cciptypes.RampMessageHeader{
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 13},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 14},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 2, SequenceNumber: 21}, // max seq num not reported
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
			},
			expErr: true,
		},
		{
			name: "msg hashes ok",
			msgs: []cciptypes.RampMessageHeader{
				{MsgHash: cciptypes.Bytes32{123}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MsgHash: cciptypes.Bytes32{99}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 13},
				{MsgHash: cciptypes.Bytes32{12}, MessageID: mustNewMessageID("0x01"), SourceChainSelector: 300, SequenceNumber: 23},
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
				{ChainSel: 300, SeqNum: 22},
			},
			expErr: false,
		},
		{
			name: "dup msg hashes",
			msgs: []cciptypes.RampMessageHeader{
				{
					MsgHash:             cciptypes.Bytes32{123},
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      12,
				},
				{
					MsgHash:             cciptypes.Bytes32{99},
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      13,
				},
				{
					MsgHash:             cciptypes.Bytes32{123},
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 300,
					SequenceNumber:      23,
				}, // dup hash
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
				{ChainSel: 300, SeqNum: 22},
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedSequenceNumbers(tc.msgs, tc.maxSeqNums)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateObserverReadingEligibility(t *testing.T) {
	testCases := []struct {
		name                string
		observer            libocrtypes.PeerID
		msgs                []cciptypes.RampMessageHeader
		seqNums             []plugintypes.SeqNumChain
		nodeSupportedChains mapset.Set[cciptypes.ChainSelector]
		destChain           cciptypes.ChainSelector
		expErr              bool
	}{
		{
			name:     "observer can read all chains",
			observer: libocrtypes.PeerID{10},
			msgs: []cciptypes.RampMessageHeader{
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x03"), SourceChainSelector: 2, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 3, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x02"), SourceChainSelector: 3, SequenceNumber: 12},
			},
			nodeSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 2, 3),
			destChain:           1,
			expErr:              false,
		},
		{
			name:     "observer is a writer so can observe seq nums",
			observer: libocrtypes.PeerID{10},
			msgs:     []cciptypes.RampMessageHeader{},
			seqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 12},
			},
			nodeSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 3),
			destChain:           1,
			expErr:              false,
		},
		{
			name:     "observer is not a writer so cannot observe seq nums",
			observer: libocrtypes.PeerID{10},
			msgs:     []cciptypes.RampMessageHeader{},
			seqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 12},
			},
			nodeSupportedChains: mapset.NewSet[cciptypes.ChainSelector](3),
			destChain:           1,
			expErr:              true,
		},
		{
			name:     "observer cfg not found",
			observer: libocrtypes.PeerID{10},
			msgs: []cciptypes.RampMessageHeader{
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 1, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x03"), SourceChainSelector: 2, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x01"), SourceChainSelector: 3, SequenceNumber: 12},
				{MessageID: mustNewMessageID("0x02"), SourceChainSelector: 3, SequenceNumber: 12},
			},
			nodeSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 3), // observer 10 not found
			destChain:           1,
			expErr:              true,
		},
		{
			name:                "no msgs",
			observer:            libocrtypes.PeerID{10},
			msgs:                []cciptypes.RampMessageHeader{},
			nodeSupportedChains: mapset.NewSet[cciptypes.ChainSelector](1, 3),
			expErr:              false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObserverReadingEligibility(tc.msgs, tc.seqNums, tc.nodeSupportedChains, tc.destChain)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateObservedTokenPrices(t *testing.T) {
	testCases := []struct {
		name        string
		tokenPrices []cciptypes.TokenPrice
		expErr      bool
	}{
		{
			name:        "empty is valid",
			tokenPrices: []cciptypes.TokenPrice{},
			expErr:      false,
		},
		{
			name: "all valid",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x3", big.NewInt(1)),
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			expErr: false,
		},
		{
			name: "dup price",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)), // dup
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			expErr: true,
		},
		{
			name: "nil price",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x3", nil), // nil price
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateObservedTokenPrices(tc.tokenPrices)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})

	}
}

func Test_validateObservedGasPrices(t *testing.T) {
	testCases := []struct {
		name      string
		gasPrices []cciptypes.GasPriceChain
		expErr    bool
	}{
		{
			name:      "empty is valid",
			gasPrices: []cciptypes.GasPriceChain{},
			expErr:    false,
		},
		{
			name: "all valid",
			gasPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(20), 2),
				cciptypes.NewGasPriceChain(big.NewInt(1312), 3),
			},
			expErr: false,
		},
		{
			name: "duplicate gas price",
			gasPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(20), 2),
				cciptypes.NewGasPriceChain(big.NewInt(1312), 1), // notice we already have a gas price for chain 1
			},
			expErr: true,
		},
		{
			name: "empty gas price",
			gasPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(20), 2),
				cciptypes.NewGasPriceChain(nil, 3), // nil
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateObservedGasPrices(tc.gasPrices)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_newMsgsConsensusForChain(t *testing.T) {
	testCases := []struct {
		name           string
		maxSeqNums     []plugintypes.SeqNumChain
		observations   []plugintypes.CommitPluginObservation
		expMerkleRoots []cciptypes.MerkleRootChain
		fChain         map[cciptypes.ChainSelector]int
		expErr         bool
	}{
		{
			name:           "empty",
			maxSeqNums:     []plugintypes.SeqNumChain{},
			observations:   nil,
			expMerkleRoots: []cciptypes.MerkleRootChain{},
			expErr:         false,
		},
		{
			name: "one message but not reaching 2fChain+1 observations",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
			},
			observations: []plugintypes.CommitPluginObservation{
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1, SequenceNumber: 11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1, SequenceNumber: 11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1, SequenceNumber: 11}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{},
			expErr:         false,
		},
		{
			name: "one message reaching 2fChain+1 observations",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
			maxSeqNums: []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 10}},
			observations: []plugintypes.CommitPluginObservation{
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(11, 11),
				},
			},
			expErr: false,
		},
		{
			name: "multiple messages all of them reaching 2fChain+1 observations",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
			},
			observations: []plugintypes.CommitPluginObservation{
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(11, 13),
				},
			},
			expErr: false,
		},
		{
			name: "one message sequence number is lower than consensus max seq num",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
			},
			observations: []plugintypes.CommitPluginObservation{
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      10}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      10}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      10}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      10}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      10}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(12, 13),
				},
			},
			expErr: false,
		},
		{
			name: "multiple messages some of them not reaching 2fChain+1 observations",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
			},
			observations: []plugintypes.CommitPluginObservation{
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x02"),
					SourceChainSelector: 1,
					SequenceNumber:      12}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 1,
					SequenceNumber:      13}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(11, 11), // we stop at 11 because there is a gap for going to 13
				},
			},
			expErr: false,
		},
		{
			name: "multiple messages on different chains",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
				2: 1,
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
				{ChainSel: 2, SeqNum: 20},
			},
			observations: []plugintypes.CommitPluginObservation{
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x01"),
					SourceChainSelector: 1,
					SequenceNumber:      11}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 2,
					SequenceNumber:      21}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 2,
					SequenceNumber:      21}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x03"),
					SourceChainSelector: 2,
					SequenceNumber:      21}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x04"),
					SourceChainSelector: 2,
					SequenceNumber:      22}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x04"),
					SourceChainSelector: 2,
					SequenceNumber:      22}}},
				{NewMsgs: []cciptypes.RampMessageHeader{{
					MessageID:           mustNewMessageID("0x04"),
					SourceChainSelector: 2,
					SequenceNumber:      22}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(11, 11), // we stop at 11 because there is a gap for going to 13
				},
				{
					ChainSel:     2,
					SeqNumsRange: cciptypes.NewSeqNumRange(21, 22), // we stop at 11 because there is a gap for going to 13
				},
			},
			expErr: false,
		},
		{
			name: "one message seq num with multiple reported ids",
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
			maxSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 1, SeqNum: 10},
			},
			observations: []plugintypes.CommitPluginObservation{
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x01"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x01"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x01"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x01"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x01"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x10"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x10"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x1101"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x1101"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x03"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x02"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x02"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x02"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x02"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
				{
					NewMsgs: []cciptypes.RampMessageHeader{
						{
							MessageID:           mustNewMessageID("0x02"),
							SourceChainSelector: 1, SequenceNumber: 11,
						}}},
			},
			expMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(11, 11),
				},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			merkleRoots, err := newMsgsConsensus(lggr, tc.maxSeqNums, tc.observations, tc.fChain)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, len(tc.expMerkleRoots), len(merkleRoots))
			for i, exp := range tc.expMerkleRoots {
				assert.Equal(t, exp.ChainSel, merkleRoots[i].ChainSel)
				assert.Equal(t, exp.SeqNumsRange, merkleRoots[i].SeqNumsRange)
			}
		})
	}
}

func Test_maxSeqNumsConsensus(t *testing.T) {
	testCases := []struct {
		name         string
		observations []plugintypes.CommitPluginObservation
		fChain       int
		expSeqNums   []plugintypes.SeqNumChain
	}{
		{
			name:         "empty observations",
			observations: []plugintypes.CommitPluginObservation{},
			fChain:       2,
			expSeqNums:   []plugintypes.SeqNumChain{},
		},
		{
			name: "one chain all followers agree",
			observations: []plugintypes.CommitPluginObservation{
				{
					MaxSeqNums: []plugintypes.SeqNumChain{
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
					},
				},
			},
			fChain: 2,
			expSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 2, SeqNum: 20},
			},
		},
		{
			name: "one chain all followers agree but not enough observations",
			observations: []plugintypes.CommitPluginObservation{
				{
					MaxSeqNums: []plugintypes.SeqNumChain{
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
					},
				},
			},
			fChain:     3,
			expSeqNums: []plugintypes.SeqNumChain{},
		},
		{
			name: "one chain 3 followers not in sync, 4 in sync",
			observations: []plugintypes.CommitPluginObservation{
				{
					MaxSeqNums: []plugintypes.SeqNumChain{
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 19},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 19},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 19},
						{ChainSel: 2, SeqNum: 20},
					},
				},
			},
			fChain: 3,
			expSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 2, SeqNum: 20},
			},
		},
		{
			name: "two chains",
			observations: []plugintypes.CommitPluginObservation{
				{
					MaxSeqNums: []plugintypes.SeqNumChain{
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},
						{ChainSel: 2, SeqNum: 20},

						{ChainSel: 3, SeqNum: 30},
						{ChainSel: 3, SeqNum: 30},
						{ChainSel: 3, SeqNum: 30},
						{ChainSel: 3, SeqNum: 30},
						{ChainSel: 3, SeqNum: 30},
					},
				},
			},
			fChain: 2,
			expSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: 2, SeqNum: 20},
				{ChainSel: 3, SeqNum: 30},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			seqNums := maxSeqNumsConsensus(lggr, tc.fChain, tc.observations)
			assert.Equal(t, tc.expSeqNums, seqNums)
		})
	}
}

func Test_tokenPricesConsensus(t *testing.T) {
	testCases := []struct {
		name         string
		observations []plugintypes.CommitPluginObservation
		fChain       int
		expPrices    []cciptypes.TokenPrice
	}{
		{
			name:         "empty",
			observations: make([]plugintypes.CommitPluginObservation, 0),
			fChain:       2,
			expPrices:    make([]cciptypes.TokenPrice, 0),
		},
		{
			name: "happy flow",
			observations: []plugintypes.CommitPluginObservation{
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(10)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(20)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(11)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(11)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(10)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(11)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(20)),
					},
				},
			},
			fChain: 2,
			expPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(11)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
			},
		},
		{
			name: "not enough observations for some token",
			observations: []plugintypes.CommitPluginObservation{
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x2", big.NewInt(20)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(11)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(11)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(10)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
					},
				},
				{
					TokenPrices: []cciptypes.TokenPrice{
						cciptypes.NewTokenPrice("0x1", big.NewInt(10)),
						cciptypes.NewTokenPrice("0x2", big.NewInt(20)),
					},
				},
			},
			fChain: 2,
			expPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x2", big.NewInt(21)),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prices := tokenPricesConsensus(tc.observations, tc.fChain)
			assert.Equal(t, tc.expPrices, prices)
		})
	}
}

func Test_gasPricesConsensus(t *testing.T) {
	testCases := []struct {
		name         string
		observations []plugintypes.CommitPluginObservation
		fChain       int
		expPrices    []cciptypes.GasPriceChain
	}{
		{
			name:         "empty",
			observations: make([]plugintypes.CommitPluginObservation, 0),
			fChain:       2,
			expPrices:    make([]cciptypes.GasPriceChain, 0),
		},
		{
			name: "one chain happy path",
			observations: []plugintypes.CommitPluginObservation{
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(20), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(11), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
			},
			fChain: 2,
			expPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
			},
		},
		{
			name: "one chain no consensus",
			observations: []plugintypes.CommitPluginObservation{
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(20), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(11), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
			},
			fChain:    3, // notice fChain is 3, means we need at least 2*3+1=7 observations
			expPrices: []cciptypes.GasPriceChain{},
		},
		{
			name: "two chains determinism check",
			observations: []plugintypes.CommitPluginObservation{
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(20), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(11), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(10), 1)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(200), 10)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(100), 10)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(100), 10)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(110), 10)}},
				{GasPrices: []cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(100), 10)}},
			},
			fChain: 2,
			expPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(100), 10),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			prices := gasPricesConsensus(lggr, tc.observations, tc.fChain)
			assert.Equal(t, tc.expPrices, prices)
		})
	}
}

func Test_validateMerkleRootsState(t *testing.T) {
	testCases := []struct {
		name               string
		reportSeqNums      []plugintypes.SeqNumChain
		onchainNextSeqNums []cciptypes.SeqNum
		expValid           bool
		expErr             bool
	}{
		{
			name: "happy path",
			reportSeqNums: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			onchainNextSeqNums: []cciptypes.SeqNum{100, 200},
			expValid:           true,
			expErr:             false,
		},
		{
			name: "one root is stale",
			reportSeqNums: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			onchainNextSeqNums: []cciptypes.SeqNum{100, 201}, // <- 200 is already on chain
			expValid:           false,
			expErr:             false,
		},
		{
			name: "one root has gap",
			reportSeqNums: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 101), // <- onchain 99 but we submit 101 instead of 100
				plugintypes.NewSeqNumChain(20, 200),
			},
			onchainNextSeqNums: []cciptypes.SeqNum{100, 200},
			expValid:           false,
			expErr:             false,
		},
		{
			name: "reader returned wrong number of seq nums",
			reportSeqNums: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			onchainNextSeqNums: []cciptypes.SeqNum{100, 200, 300},
			expValid:           false,
			expErr:             true,
		},
	}

	ctx := context.Background()
	lggr := logger.Test(t)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := mocks.NewCCIPReader()
			rep := cciptypes.CommitPluginReport{}
			chains := make([]cciptypes.ChainSelector, 0, len(tc.reportSeqNums))
			for _, snc := range tc.reportSeqNums {
				rep.MerkleRoots = append(rep.MerkleRoots, cciptypes.MerkleRootChain{
					ChainSel:     snc.ChainSel,
					SeqNumsRange: cciptypes.NewSeqNumRange(snc.SeqNum, snc.SeqNum+10),
				})
				chains = append(chains, snc.ChainSel)
			}
			reader.On("NextSeqNum", ctx, chains).Return(tc.onchainNextSeqNums, nil)
			valid, err := validateMerkleRootsState(ctx, lggr, rep, reader)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expValid, valid)
		})
	}
}

func mustNewMessageID(msgIDHex string) cciptypes.Bytes32 {
	msgID, err := cciptypes.NewBytes32FromString(msgIDHex)
	if err != nil {
		panic(err)
	}
	return msgID
}

func messageIDFromInt(i int) cciptypes.Bytes32 {
	var msgIDBytes cciptypes.Bytes32
	binary.BigEndian.PutUint64(msgIDBytes[:], uint64(i))
	return msgIDBytes
}
