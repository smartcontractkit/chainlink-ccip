package ccipocr3

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSeqNumRange(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		rng := NewSeqNumRange(1, 2)
		assert.Equal(t, SeqNum(1), rng.Start())
		assert.Equal(t, SeqNum(2), rng.End())
	})

	t.Run("empty", func(t *testing.T) {
		rng := SeqNumRange{}
		assert.Equal(t, SeqNum(0), rng.Start())
		assert.Equal(t, SeqNum(0), rng.End())
	})

	t.Run("override start and end", func(t *testing.T) {
		rng := NewSeqNumRange(1, 2)
		rng.SetStart(10)
		rng.SetEnd(20)
		assert.Equal(t, SeqNum(10), rng.Start())
		assert.Equal(t, SeqNum(20), rng.End())
	})

	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "[1 -> 2]", NewSeqNumRange(1, 2).String())
		assert.Equal(t, "[0 -> 0]", SeqNumRange{}.String())
	})
}

func TestSeqNumRange_Overlap(t *testing.T) {
	testCases := []struct {
		name string
		r1   SeqNumRange
		r2   SeqNumRange
		exp  bool
	}{
		{"OverlapMiddle", SeqNumRange{5, 10}, SeqNumRange{8, 12}, true},
		{"OverlapStart", SeqNumRange{5, 10}, SeqNumRange{10, 15}, true},
		{"OverlapEnd", SeqNumRange{5, 10}, SeqNumRange{0, 5}, true},
		{"NoOverlapBefore", SeqNumRange{5, 10}, SeqNumRange{0, 4}, false},
		{"NoOverlapAfter", SeqNumRange{5, 10}, SeqNumRange{11, 15}, false},
		{"SameRange", SeqNumRange{5, 10}, SeqNumRange{5, 10}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, tc.r1.Overlaps(tc.r2))
		})
	}
}

func TestSeqNumRange_Contains(t *testing.T) {
	tests := []struct {
		name     string
		r        SeqNumRange
		seq      SeqNum
		expected bool
	}{
		{"ContainsMiddle", SeqNumRange{5, 10}, SeqNum(7), true},
		{"ContainsStart", SeqNumRange{5, 10}, SeqNum(5), true},
		{"ContainsEnd", SeqNumRange{5, 10}, SeqNum(10), true},
		{"BeforeRange", SeqNumRange{5, 10}, SeqNum(4), false},
		{"AfterRange", SeqNumRange{5, 10}, SeqNum(11), false},
		{"EmptyRange", SeqNumRange{5, 5}, SeqNum(5), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.r.Contains(tt.seq))
		})
	}
}

func TestSeqNumRangeLimit(t *testing.T) {
	testCases := []struct {
		name string
		rng  SeqNumRange
		n    uint64
		want SeqNumRange
	}{
		{
			name: "no truncation",
			rng:  NewSeqNumRange(0, 10),
			n:    11,
			want: NewSeqNumRange(0, 10),
		},
		{
			name: "no truncation 2",
			rng:  NewSeqNumRange(100, 110),
			n:    11,
			want: NewSeqNumRange(100, 110),
		},
		{
			name: "truncation",
			rng:  NewSeqNumRange(0, 10),
			n:    10,
			want: NewSeqNumRange(0, 9),
		},
		{
			name: "truncation 2",
			rng:  NewSeqNumRange(100, 110),
			n:    10,
			want: NewSeqNumRange(100, 109),
		},
		{
			name: "empty",
			rng:  NewSeqNumRange(0, 0),
			n:    0,
			want: NewSeqNumRange(0, 0),
		},
		{
			name: "wrong range",
			rng:  NewSeqNumRange(20, 15),
			n:    3,
			want: NewSeqNumRange(20, 15),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rng.Limit(tc.n)
			if got != tc.want {
				t.Errorf("SeqNumRangeLimit(%v, %v) = %v; want %v", tc.rng, tc.n, got, tc.want)
			}
		})
	}
}

func TestSeqNumFilterSlice(t *testing.T) {
	testCases := []struct {
		name     string
		r        SeqNumRange
		seqs     []SeqNum
		expected []SeqNum
	}{
		{
			"none",
			SeqNumRange{0, 0},
			[]SeqNum{1, 2, 3},
			nil,
		},
		{
			"zero in range",
			SeqNumRange{0, 0},
			[]SeqNum{1, 2, 0, 3},
			[]SeqNum{0},
		},
		{
			"inclusive",
			SeqNumRange{10, 20},
			[]SeqNum{1, 10, 2, 0, 4, 13, 3, 20},
			[]SeqNum{10, 13, 20},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.r.FilterSlice(tc.seqs))
		})
	}
}

func TestCCIPMsg_String(t *testing.T) {
	tests := []struct {
		name     string
		c        Message
		expected string
	}{
		{
			"base",
			Message{
				Header: RampMessageHeader{
					MessageID:           mustNewBytes32(t, "0x01"),
					SourceChainSelector: ChainSelector(1),
					DestChainSelector:   ChainSelector(2),
					SequenceNumber:      2,
					Nonce:               1,

					MsgHash: mustNewBytes32(t, "0x23"),
					OnRamp:  mustNewUnknownAddress(t, "0x04D4cC5972ad487F71b85654d48b27D32b13a22F"),
				},
			},
			//nolint:lll // test input
			`{"header":{"messageId":"0x0100000000000000000000000000000000000000000000000000000000000000","sourceChainSelector":"1","destChainSelector":"2","seqNum":"2","nonce":1,"msgHash":"0x2300000000000000000000000000000000000000000000000000000000000000","onRamp":"0x04d4cc5972ad487f71b85654d48b27d32b13a22f"},"sender":"0x","data":"0x","receiver":"0x","extraArgs":"0x","feeToken":"0x","feeTokenAmount":null,"feeValueJuels":null,"tokenAmounts":null}`,
		},
		{
			"with evm ramp message",
			Message{
				Header: RampMessageHeader{
					MessageID:           mustNewBytes32(t, "0x01"),
					SourceChainSelector: ChainSelector(1),
					DestChainSelector:   ChainSelector(2),
					SequenceNumber:      2,
					Nonce:               1,

					MsgHash: mustNewBytes32(t, "0x23"),
					OnRamp:  mustNewUnknownAddress(t, "0x04D4cC5972ad487F71b85654d48b27D32b13a22F"),
				},
				Sender:         mustNewUnknownAddress(t, "0x04D4cC5972ad487F71b85654d48b27D32b13a22F"),
				Receiver:       mustNewUnknownAddress(t, "0x101112131415"), // simulate a non-evm receiver
				Data:           []byte("some data"),
				ExtraArgs:      []byte("extra args"),
				FeeToken:       mustNewUnknownAddress(t, "0xB5fCC870d2aC8745054b4ba99B1f176B93382162"),
				FeeTokenAmount: BigInt{Int: big.NewInt(1000)},
				FeeValueJuels:  BigInt{Int: big.NewInt(287)},
				TokenAmounts: []RampTokenAmount{
					{
						SourcePoolAddress: mustNewUnknownAddress(t, "0x3E8456720B88A1DAdce8E2808C9Bf73dfFFd807c"),
						DestTokenAddress:  []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, // simulate a non-evm token
						ExtraData:         []byte("extra token data"),
						Amount:            BigInt{Int: big.NewInt(2000)},
						DestExecData:      []byte("extra token data"),
					},
				},
			},
			//nolint:lll // test input
			`{"header":{"messageId":"0x0100000000000000000000000000000000000000000000000000000000000000","sourceChainSelector":"1","destChainSelector":"2","seqNum":"2","nonce":1,"msgHash":"0x2300000000000000000000000000000000000000000000000000000000000000","onRamp":"0x04d4cc5972ad487f71b85654d48b27d32b13a22f"},"sender":"0x04d4cc5972ad487f71b85654d48b27d32b13a22f","data":"0x736f6d652064617461","receiver":"0x101112131415","extraArgs":"0x65787472612061726773","feeToken":"0xb5fcc870d2ac8745054b4ba99b1f176b93382162","feeTokenAmount":"1000","feeValueJuels":"287","tokenAmounts":[{"sourcePoolAddress":"0x3e8456720b88a1dadce8e2808c9bf73dfffd807c","destTokenAddress":"0x0102030405060708090a","extraData":"0x657874726120746f6b656e2064617461","amount":"2000","destExecData":"0x657874726120746f6b656e2064617461"}]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.c.String())
		})
	}
}

func TestNewTokenPrice(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		tp := NewTokenPrice("link", big.NewInt(1000))
		assert.Equal(t, "link", string(tp.TokenID))
		assert.Equal(t, uint64(1000), tp.Price.Int.Uint64())
	})
}

func TestNewGasPriceChain(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		gpc := NewGasPriceChain(big.NewInt(1000), ChainSelector(1))
		assert.Equal(t, uint64(1000), (gpc.GasPrice).Uint64())
		assert.Equal(t, ChainSelector(1), gpc.ChainSel)
	})
}

func TestMerkleRoot(t *testing.T) {
	t.Run("str", func(t *testing.T) {
		mr := Bytes32([32]byte{1})
		assert.Equal(t, "0x0100000000000000000000000000000000000000000000000000000000000000", mr.String())
	})

	t.Run("json", func(t *testing.T) {
		mr := Bytes32([32]byte{1})
		b, err := json.Marshal(mr)
		assert.NoError(t, err)
		assert.Equal(t, `"0x0100000000000000000000000000000000000000000000000000000000000000"`, string(b))

		mr2 := Bytes32{}
		err = json.Unmarshal(b, &mr2)
		assert.NoError(t, err)
		assert.Equal(t, mr, mr2)

		mr3 := Bytes32{}
		err = json.Unmarshal([]byte(`"123"`), &mr3)
		assert.Error(t, err)

		err = json.Unmarshal([]byte(`""`), &mr3)
		assert.Error(t, err)
	})
}

func mustNewBytes32(t *testing.T, s string) Bytes32 {
	b32, err := NewBytes32FromString(s)
	require.NoError(t, err)
	return b32
}

func mustNewUnknownAddress(t *testing.T, s string) UnknownAddress {
	b, err := NewUnknownAddressFromHex(s)
	require.NoError(t, err)
	return b
}
