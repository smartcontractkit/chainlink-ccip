package plugintypes

import (
	"math"
	"math/big"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCommitPluginObservation_EncodeAndDecode(t *testing.T) {
	obs := NewCommitPluginObservation(
		[]cciptypes.RampMessageHeader{
			{
				MessageID:           mustNewBytes32(t, "0x01"),
				SourceChainSelector: math.MaxUint64,
				DestChainSelector:   cciptypes.ChainSelector(123),
				SequenceNumber:      123,
				Nonce:               1,

				MsgHash: cciptypes.Bytes32{1},
				OnRamp:  mustNewBytes(t, "0x010203"),
			},
			{
				MessageID:           mustNewBytes32(t, "0x02"),
				SourceChainSelector: 321,
				DestChainSelector:   cciptypes.ChainSelector(456),
				SequenceNumber:      math.MaxUint64,
				Nonce:               0,

				MsgHash: cciptypes.Bytes32{2},
				OnRamp:  mustNewBytes(t, "0x040506"),
			},
		},
		[]cciptypes.GasPriceChain{
			cciptypes.NewGasPriceChain(big.NewInt(1234), cciptypes.ChainSelector(math.MaxUint64)),
		},
		[]cciptypes.TokenPrice{},
		[]SeqNumChain{},
		map[cciptypes.ChainSelector]int{},
	)

	b, err := obs.Encode()
	assert.NoError(t, err)
	// nolint:lll
	assert.Equal(t, `{"newMsgs":[{"messageId":"0x0100000000000000000000000000000000000000000000000000000000000000","sourceChainSelector":"18446744073709551615","destChainSelector":"123","seqNum":"123","nonce":1,"msgHash":"0x0100000000000000000000000000000000000000000000000000000000000000","onRamp":"0x010203"},{"messageId":"0x0200000000000000000000000000000000000000000000000000000000000000","sourceChainSelector":"321","destChainSelector":"456","seqNum":"18446744073709551615","nonce":0,"msgHash":"0x0200000000000000000000000000000000000000000000000000000000000000","onRamp":"0x040506"}],"gasPrices":[{"UsdPerUnitGas":"1234","DestChainSelector":18446744073709551615}],"tokenPrices":[],"maxSeqNums":[],"fChain":{}}`, string(b))

	obs2, err := DecodeCommitPluginObservation(b)
	assert.NoError(t, err)
	assert.Equal(t, obs, obs2)
}

func TestCommitPluginOutcome_EncodeAndDecode(t *testing.T) {
	o := NewCommitPluginOutcome(
		[]SeqNumChain{
			NewSeqNumChain(cciptypes.ChainSelector(1), cciptypes.SeqNum(20)),
			NewSeqNumChain(cciptypes.ChainSelector(2), cciptypes.SeqNum(25)),
		},
		[]cciptypes.MerkleRootChain{
			cciptypes.NewMerkleRootChain(cciptypes.ChainSelector(1), cciptypes.NewSeqNumRange(21, 22), [32]byte{1}),
			cciptypes.NewMerkleRootChain(cciptypes.ChainSelector(2), cciptypes.NewSeqNumRange(25, 35), [32]byte{2}),
		},
		[]cciptypes.TokenPrice{
			cciptypes.NewTokenPrice("0x123", big.NewInt(1234)),
			cciptypes.NewTokenPrice("0x125", big.NewInt(0).Mul(big.NewInt(999999999999), big.NewInt(999999999999))),
		},
		[]cciptypes.GasPriceChain{
			cciptypes.NewGasPriceChain(big.NewInt(1234), cciptypes.ChainSelector(1)),
			cciptypes.NewGasPriceChain(big.NewInt(0).Mul(big.NewInt(999999999999), big.NewInt(999999999999)),
				cciptypes.ChainSelector(2)),
		},
	)

	b, err := o.Encode()
	assert.NoError(t, err)
	// nolint:lll
	assert.Equal(t, `{"maxSeqNums":[{"chainSel":1,"seqNum":20},{"chainSel":2,"seqNum":25}],"merkleRoots":[{"SourceChainSelector":1,"Interval":{"Min":21,"Max":22},"MerkleRoot":"0x0100000000000000000000000000000000000000000000000000000000000000"},{"SourceChainSelector":2,"Interval":{"Min":25,"Max":35},"MerkleRoot":"0x0200000000000000000000000000000000000000000000000000000000000000"}],"tokenPrices":[{"SourceToken":"0x123","UsdPerToken":"1234"},{"SourceToken":"0x125","UsdPerToken":"999999999998000000000001"}],"gasPrices":[{"UsdPerUnitGas":"1234","DestChainSelector":1},{"UsdPerUnitGas":"999999999998000000000001","DestChainSelector":2}]}`, string(b))

	o2, err := DecodeCommitPluginOutcome(b)
	assert.NoError(t, err)
	assert.Equal(t, o, o2)
	// nolint:lll
	assert.Equal(t, `{MaxSeqNums: [{ChainSelector(1) 20} {ChainSelector(2) 25}], MerkleRoots: [{ChainSelector(1) [21 -> 22] 0x0100000000000000000000000000000000000000000000000000000000000000} {ChainSelector(2) [25 -> 35] 0x0200000000000000000000000000000000000000000000000000000000000000}]}`, o.String())
}

func TestCommitPluginOutcome_IsEmpty(t *testing.T) {
	o := NewCommitPluginOutcome(nil, nil, nil, nil)
	assert.True(t, o.IsEmpty())

	o = NewCommitPluginOutcome(nil, nil, nil,
		[]cciptypes.GasPriceChain{cciptypes.NewGasPriceChain(big.NewInt(1), cciptypes.ChainSelector(1))})
	assert.False(t, o.IsEmpty())

	o = NewCommitPluginOutcome(nil, nil,
		[]cciptypes.TokenPrice{cciptypes.NewTokenPrice("0x123", big.NewInt(123))}, nil)
	assert.False(t, o.IsEmpty())

	o = NewCommitPluginOutcome(nil, []cciptypes.MerkleRootChain{
		cciptypes.NewMerkleRootChain(
			cciptypes.ChainSelector(1), cciptypes.NewSeqNumRange(1, 2), [32]byte{1})}, nil, nil)
	assert.False(t, o.IsEmpty())

	o = NewCommitPluginOutcome([]SeqNumChain{
		NewSeqNumChain(cciptypes.ChainSelector(1), cciptypes.SeqNum(1))}, nil, nil, nil)
	assert.False(t, o.IsEmpty())
}

func mustNewBytes32(t *testing.T, s string) cciptypes.Bytes32 {
	b, err := cciptypes.NewBytes32FromString(s)
	require.NoError(t, err)
	return b
}

func mustNewBytes(t *testing.T, s string) cciptypes.Bytes {
	b, err := cciptypes.NewBytesFromString(s)
	require.NoError(t, err)
	return b
}
