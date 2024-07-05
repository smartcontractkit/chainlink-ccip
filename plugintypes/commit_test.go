package plugintypes

import (
	"math"
	"math/big"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
)

func TestCommitPluginObservation_EncodeAndDecode(t *testing.T) {
	obs := NewCommitPluginObservation(
		[]cciptypes.CCIPMsgBaseDetails{
			{MsgHash: cciptypes.Bytes32{1}, ID: "1", SourceChain: math.MaxUint64, SeqNum: 123},
			{MsgHash: cciptypes.Bytes32{2}, ID: "2", SourceChain: 321, SeqNum: math.MaxUint64},
		},
		[]cciptypes.GasPriceChain{}, // todo: populate this
		[]cciptypes.TokenPrice{},
		[]SeqNumChain{},
		map[cciptypes.ChainSelector]int{},
	)

	b, err := obs.Encode()
	assert.NoError(t, err)
	// nolint:lll
	assert.Equal(t, `{"newMsgs":[{"id":"1","sourceChain":"18446744073709551615","seqNum":"123","msgHash":"0x0100000000000000000000000000000000000000000000000000000000000000"},{"id":"2","sourceChain":"321","seqNum":"18446744073709551615","msgHash":"0x0200000000000000000000000000000000000000000000000000000000000000"}],"gasPrices":[],"tokenPrices":[],"maxSeqNums":[],"fChain":{}}`, string(b))

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
	assert.Equal(t, `{"maxSeqNums":[{"chainSel":1,"seqNum":20},{"chainSel":2,"seqNum":25}],"merkleRoots":[{"chain":1,"seqNumsRange":[21,22],"merkleRoot":"0x0100000000000000000000000000000000000000000000000000000000000000"},{"chain":2,"seqNumsRange":[25,35],"merkleRoot":"0x0200000000000000000000000000000000000000000000000000000000000000"}],"tokenPrices":[{"tokenID":"0x123","price":"1234"},{"tokenID":"0x125","price":"999999999998000000000001"}],"gasPrices":[{"gasPrice":"1234","chainSel":1},{"gasPrice":"999999999998000000000001","chainSel":2}]}`, string(b))

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
