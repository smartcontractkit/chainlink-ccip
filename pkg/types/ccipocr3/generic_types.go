package ccipocr3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

type TokenPrice struct {
	TokenID UnknownEncodedAddress `json:"tokenID"`
	Price   BigInt                `json:"price"`
}

type TokenPriceMap map[UnknownEncodedAddress]BigInt

func (t TokenPriceMap) ToSortedSlice() []TokenPrice {
	var res []TokenPrice
	for tokenID, price := range t {
		res = append(res, TokenPrice{tokenID, price})
	}

	// sort the token prices by tokenID
	sort.Slice(res, func(i, j int) bool {
		return res[i].TokenID < res[j].TokenID
	})

	return res
}

func NewTokenPrice(tokenID UnknownEncodedAddress, price *big.Int) TokenPrice {
	return TokenPrice{
		TokenID: tokenID,
		Price:   BigInt{price},
	}
}

type GasPriceChain struct {
	ChainSel ChainSelector `json:"chainSel"`
	GasPrice BigInt        `json:"gasPrice"`
}

func NewGasPriceChain(gasPrice *big.Int, chainSel ChainSelector) GasPriceChain {
	return GasPriceChain{
		ChainSel: chainSel,
		GasPrice: NewBigInt(gasPrice),
	}
}

type SeqNum uint64

func (s SeqNum) String() string {
	return strconv.FormatUint(uint64(s), 10)
}

func (s SeqNum) IsWithinRanges(ranges []SeqNumRange) bool {
	for _, r := range ranges {
		if r.Contains(s) {
			return true
		}
	}
	return false
}

func NewSeqNumRange(start, end SeqNum) SeqNumRange {
	return SeqNumRange{start, end}
}

// SeqNumRange defines an inclusive range of sequence numbers.
type SeqNumRange [2]SeqNum

func (s SeqNumRange) Start() SeqNum {
	return s[0]
}

func (s SeqNumRange) End() SeqNum {
	return s[1]
}

func (s *SeqNumRange) SetStart(v SeqNum) {
	s[0] = v
}

func (s *SeqNumRange) SetEnd(v SeqNum) {
	s[1] = v
}

// Limit returns a range limited up to n elements by truncating the end if necessary.
// Example: [1 -> 10].Limit(5) => [1 -> 5]
func (s *SeqNumRange) Limit(n uint64) SeqNumRange {
	limitedRange := NewSeqNumRange(s.Start(), s.End())

	numElems := s.End() - s.Start() + 1
	if numElems <= 0 {
		return limitedRange
	}

	if uint64(numElems) > n {
		newEnd := limitedRange.Start() + SeqNum(n) - 1
		if newEnd > limitedRange.End() { // overflow - do nothing
			return limitedRange
		}
		limitedRange.SetEnd(newEnd)
	}

	return limitedRange
}

// Overlaps returns true if the two ranges overlap.
func (s SeqNumRange) Overlaps(other SeqNumRange) bool {
	return s.Start() <= other.End() && other.Start() <= s.End()
}

// Contains returns true if the range contains the given sequence number.
func (s SeqNumRange) Contains(seq SeqNum) bool {
	return s.Start() <= seq && seq <= s.End()
}

// FilterSlice returns a slice of sequence numbers that are contained in the range.
func (s SeqNumRange) FilterSlice(seqNums []SeqNum) []SeqNum {
	var contained []SeqNum
	for _, seq := range seqNums {
		if s.Contains(seq) {
			contained = append(contained, seq)
		}
	}
	return contained
}

func (s SeqNumRange) String() string {
	return fmt.Sprintf("[%d -> %d]", s[0], s[1])
}

func (s SeqNumRange) Length() int {
	return int(s.End() - s.Start() + 1)
}

// ToSlice returns a slice of sequence numbers in the range.
func (s SeqNumRange) ToSlice() []SeqNum {
	var seqs []SeqNum
	for i := s.Start(); i <= s.End(); i++ {
		seqs = append(seqs, i)
	}
	return seqs
}

type ChainSelector uint64

func (c ChainSelector) String() string {
	return fmt.Sprintf("ChainSelector(%d)", c)
}

// Message is the generic Any2Any message type for CCIP messages.
// It represents, in particular, a message emitted by a CCIP onramp.
// The message is expected to be consumed by a CCIP offramp after
// translating it into the appropriate format for the destination chain.
type Message struct {
	// Header is the family-agnostic header for OnRamp and OffRamp messages.
	// This is always set on all CCIP messages.
	Header RampMessageHeader `json:"header"`
	// Sender address on the source chain.
	// i.e if the source chain is EVM, this is an abi-encoded EVM address.
	Sender UnknownAddress `json:"sender"`
	// Data is the arbitrary data payload supplied by the message sender.
	Data Bytes `json:"data"`
	// Receiver is the receiver address on the destination chain.
	// This is encoded in the destination chain family specific encoding.
	// i.e if the destination is EVM, this is abi.encode(receiver).
	Receiver UnknownAddress `json:"receiver"`
	// ExtraArgs is destination-chain specific extra args,
	// such as the gasLimit for EVM chains.
	// This field is encoded in the source chain encoding scheme.
	ExtraArgs Bytes `json:"extraArgs"`
	// FeeToken is the fee token address.
	// i.e if the source chain is EVM, len(FeeToken) == 20 (i.e, is not abi-encoded).
	FeeToken UnknownAddress `json:"feeToken"`
	// FeeTokenAmount is the amount of fee tokens paid.
	FeeTokenAmount BigInt `json:"feeTokenAmount"`
	// FeeValueJuels is the fee amount in Juels
	FeeValueJuels BigInt `json:"feeValueJuels"`
	// TokenAmounts is the array of tokens and amounts to transfer.
	TokenAmounts []RampTokenAmount `json:"tokenAmounts"`
}

func (m Message) CopyWithoutData() Message {
	return Message{
		Header:         m.Header,
		Sender:         m.Sender,
		Data:           []byte{},
		Receiver:       m.Receiver,
		ExtraArgs:      m.ExtraArgs,
		FeeToken:       m.FeeToken,
		FeeTokenAmount: m.FeeTokenAmount,
		FeeValueJuels:  m.FeeValueJuels,
		TokenAmounts:   m.TokenAmounts,
	}
}

func (m Message) String() string {
	js, _ := json.Marshal(m)
	return string(js)
}

// IsPseudoDeleted returns true when the message is stripped out of some fields that makes it usable. Message still
// contains some metaData like seqNumber and SourceChainSelector to be able to distinguish it from other messages while
// still in the pseudo deleted state.
func (m Message) IsPseudoDeleted() bool {
	return m.Header.DestChainSelector == 0 && m.Header.SourceChainSelector == 0 &&
		len(m.Header.OnRamp) == 0 && len(m.Receiver) == 0 && len(m.Sender) == 0
}

// RampMessageHeader is the family-agnostic header for OnRamp and OffRamp messages.
// The MessageID is not expected to match MsgHash, since it may originate from a different
// ramp family.
type RampMessageHeader struct {
	// MessageID is a unique identifier for the message, it should be unique across all chains.
	// It is generated on the chain that the CCIP send is requested (i.e. the source chain of a message).
	MessageID Bytes32 `json:"messageId"`
	// SourceChainSelector is the chain selector of the chain that the message originated from.
	SourceChainSelector ChainSelector `json:"sourceChainSelector,string"`
	// DestChainSelector is the chain selector of the chain that the message is destined for.
	DestChainSelector ChainSelector `json:"destChainSelector,string"`
	// SequenceNumber is an auto-incrementing sequence number for the message.
	// Not unique across lanes.
	SequenceNumber SeqNum `json:"seqNum,string"`
	// Nonce is the nonce for this lane for this sender, not unique across senders/lanes
	Nonce uint64 `json:"nonce"`

	// MsgHash is the hash of all the message fields.
	// NOTE: The field is expected to be empty, and will be populated by the plugin using the MsgHasher interface.
	MsgHash Bytes32 `json:"msgHash"` // populated

	// OnRamp is the address of the onramp that sent the message.
	// NOTE: This is populated by the ccip reader. Not emitted explicitly onchain.
	OnRamp UnknownAddress `json:"onRamp"`
}

// RampTokenAmount represents the family-agnostic token amounts used for both OnRamp & OffRamp messages.
type RampTokenAmount struct {
	// SourcePoolAddress is the source pool address, encoded according to source family native encoding scheme.
	// This value is trusted as it was obtained through the onRamp. It can be relied upon by the destination
	// pool to validate the source pool.
	SourcePoolAddress UnknownAddress `json:"sourcePoolAddress"`

	// DestTokenAddress is the address of the destination token, abi encoded in the case of EVM chains.
	// This value is UNTRUSTED as any pool owner can return whatever value they want.
	DestTokenAddress UnknownAddress `json:"destTokenAddress"`

	// ExtraData is optional pool data to be transferred to the destination chain. Be default this is capped at
	// CCIP_LOCK_OR_BURN_V1_RET_BYTES bytes. If more data is required, the TokenTransferFeeConfig.destBytesOverhead
	// has to be set for the specific token.
	ExtraData Bytes `json:"extraData"`

	// Amount is the amount of tokens to be transferred.
	Amount BigInt `json:"amount"`

	// DestExecData is destination chain specific execution data encoded in bytes.
	// For an EVM destination, it consists of the amount of gas available for the releaseOrMint
	// and transfer calls made by the offRamp.
	// NOTE: this must be decoded before providing it as an execution input to the destination chain
	// or hashing it. See Internal._hash(Any2EVMRampMessage) for more details as an example.
	DestExecData Bytes `json:"destExecData"`
}
